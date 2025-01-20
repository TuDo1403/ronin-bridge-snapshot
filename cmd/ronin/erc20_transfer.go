package ronin

import (
	"context"
	"fmt"
	"os"
	"ronin-bridge-snapshot/internal/abi/erc20"
	"ronin-bridge-snapshot/internal/config"
	"ronin-bridge-snapshot/internal/util"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

// record fetches and records transfer events for multiple tokens over a batch of blocks,
func RecordWithdrawalsOnRonin(
	ctx context.Context,
	wg *sync.WaitGroup,
	done chan struct{},
	trackers map[common.Address]*util.Tracker,
	client *ethclient.Client,
	nwCfg config.NetworkConfig,
) {
	// Calculate how many batches we'll do
	numBatch := (nwCfg.EndBlock-nwCfg.StartBlock)/uint64(nwCfg.QueryBatchSize) + 1
	var processedBatches int64
	log.Info("Config", "BatchCount", numBatch)

	// Concurrency limit for block-range queries
	concurrencyLimit := 10
	sem := make(chan struct{}, concurrencyLimit)

	// Channels for producer->consumer flow
	rawLogChan := make(chan []types.Log, 10000)
	sanitizedLogChan := make(chan map[common.Address][]*erc20.Erc20Transfer, 10000)
	mintTxHashesChan := make(chan map[common.Address][]common.Hash, 10000)
	sanitizedTxHashesChan := make(chan map[common.Address][]common.Hash, 10000)

	// 1) Goroutine: Sanitize logs (consumes rawLogChan, produces sanitized logs/txhashes)
	wg.Add(1)
	go func() {
		defer wg.Done()
		util.SanitizeTransferLogs(
			ctx,
			done,
			nwCfg.Gateway,
			rawLogChan,
			sanitizedLogChan,
			mintTxHashesChan,
			sanitizedTxHashesChan,
		)
	}()

	// 2) Goroutine: Process sanitized logs/txhashes
	wg.Add(1)
	go func() {
		defer wg.Done()

		var mu sync.Mutex // protect shared trackers

		for {
			select {
			case <-ctx.Done():
				log.Trace("Processor routine stopped (ctx canceled).")
				return

			case <-done:
				log.Trace("Processor routine received done signal. Exiting.")
				return

			case transferEvents := <-sanitizedLogChan:

				// Update trackers for each token
				for token, transferEventBatch := range transferEvents {
					mu.Lock()
					util.RecordTransfer(trackers[token], transferEventBatch)
					mu.Unlock()
				}

				atomic.AddInt64(&processedBatches, 1)
				log.Trace("Processed sanitized logs", "Count", len(transferEvents), "BatchCount", atomic.LoadInt64(&processedBatches))

			case mintTxHashes := <-mintTxHashesChan:
				// Update trackers for each token
				for token, mintTxHashBatch := range mintTxHashes {
					mu.Lock()
					util.RecordMintTxHashes(trackers[token], mintTxHashBatch)
					mu.Unlock()
				}
			case sanitizedTxHashes := <-sanitizedTxHashesChan:
				// Update trackers for each token
				for token, sanitizedTxHashBatch := range sanitizedTxHashes {
					mu.Lock()
					util.RecordSanitizedTxHashes(trackers[token], sanitizedTxHashBatch)
					mu.Unlock()
				}
			}
		}
	}()

	// 3) Main loop: spawn goroutines for each batch
	var batchIndex int
	var endBlock uint64

	wg.Add(1)
	// End condition
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(20)

		for {
			select {
			case <-ctx.Done():
				log.Trace("Processor routine stopped (ctx canceled).")
				return
			case <-done:
				log.Trace("Processor routine received done signal. Exiting.")
				return
			case <-ticker.C:
				if atomic.LoadInt64(&processedBatches) == int64(numBatch) {
					log.Info("All batch jobs finished. Closing 'done' channel.")
					close(done)
					return
				}
			}
		}
	}()

	for startBlock := nwCfg.StartBlock; startBlock < nwCfg.EndBlock; startBlock += uint64(nwCfg.QueryBatchSize) {
		endBlock = startBlock + uint64(nwCfg.QueryBatchSize) - 1
		if endBlock > nwCfg.EndBlock {
			endBlock = nwCfg.EndBlock
		}

		batchIndex++
		if batchIndex%100 == 0 {
			progress := float64(batchIndex) / float64(numBatch) * 100
			log.Info("Progress", "%", progress)
		}

		sb, eb := startBlock, endBlock

		// Acquire concurrency slot
		sem <- struct{}{}

		wg.Add(1)
		go func() {
			defer wg.Done()
			defer func() { <-sem }()

			util.FilterERC20TransferForTokens(
				ctx,
				done,
				client,
				nwCfg.Tokens,
				nwCfg.Gateway,
				sb,
				eb,
				rawLogChan,
			)
		}()
	}

	// Wait for all batch goroutines to finish
	wg.Wait()

	// Now we know no more logs will be produced, so we can close the channels.
	close(rawLogChan)
	close(sanitizedLogChan)
	close(mintTxHashesChan)
	close(sanitizedTxHashesChan)

	names := make(map[common.Address]string, len(nwCfg.Tokens))
	symbols := make(map[common.Address]string, len(nwCfg.Tokens))
	decimals := make(map[common.Address]int, len(nwCfg.Tokens))

	// Sort tokens
	sort.Slice(nwCfg.Tokens, func(i, j int) bool {
		return nwCfg.Tokens[i].Hex() < nwCfg.Tokens[j].Hex()
	})

	// Fetch token metadata
	for _, token := range nwCfg.Tokens {
		name, symbol, dec := util.FetchERC20Metadata(ctx, client, token)
		names[token] = name
		symbols[token] = symbol
		decimals[token] = dec
	}

	// Log final state
	for token, tracker := range trackers {
		log.Info("Snapshot", "LastBlock", tracker.LastRecordedBlock, "Addr", token, "Symbol", symbols[token], "TxCount", tracker.TxCount, "TotalWithdraw", util.BigIntToFloat(tracker.AccAmount, decimals[token]))
	}

	// Write TxHashes to file
	f, err := os.Create("ronin-weth-tx-hashes.txt")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	wethTxHashes := trackers[common.HexToAddress("0xc99a6A985eD2Cac1ef41640596C5A5f9F4E19Ef5")].SanitizedTxHashes

	for _, v := range wethTxHashes {
		fmt.Fprintln(f, v)
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
}
