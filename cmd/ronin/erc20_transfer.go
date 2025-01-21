package ronin

import (
	"context"
	"fmt"
	"os"

	// "fmt"
	// "os"
	"ronin-bridge-snapshot/internal/config"
	"ronin-bridge-snapshot/internal/erc20_transfer"
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
	sanitizedTransferChan := make(chan map[common.Address][]*erc20_transfer.Transfer, 10000)

	// 1) Goroutine: Sanitize logs (consumes rawLogChan, produces sanitized logs/txHashes)
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				log.Trace("Sanitizer routine stopped (ctx canceled).")
				return
			case <-done:
				log.Trace("Sanitizer routine received done signal. Exiting.")
				return
			case rawLogs := <-rawLogChan:
				erc20_transfer.SanitizeTransferLogs(rawLogs, util.ToMap(nwCfg.ExcludeTxHashes), nwCfg.Gateway, sanitizedTransferChan)
			}
		}
	}()

	// 2) Goroutine: Process sanitized logs/txHashes
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

			case transferEvents := <-sanitizedTransferChan:
				for token, transfers := range transferEvents {
					mu.Lock()
					trackers[token].RecordTransfers(transfers)
					mu.Unlock()
				}

				atomic.AddInt64(&processedBatches, 1)
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

			erc20_transfer.FilterERC20TransferForTokens(
				ctx,
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
	close(sanitizedTransferChan)

	names := make(map[common.Address]string, len(nwCfg.Tokens))
	symbols := make(map[common.Address]string, len(nwCfg.Tokens))
	decimals := make(map[common.Address]int, len(nwCfg.Tokens))
	symbol2token := make(map[string]common.Address, len(nwCfg.Tokens))

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
		symbol2token[symbol] = token
	}

	// Log final state
	for token, tracker := range trackers {
		log.Info("Snapshot", "LastBlock", tracker.LastRecordedBlock, "Addr", token, "Symbol", symbols[token], "TxCount", tracker.TxCount, "TotalWithdraw", util.BigIntToFloat(tracker.AccAmount, decimals[token]))
	}

	// Write TxHashes to file
	f, err := os.Create("ronin-weth-erc20-tx-hashes.txt")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	wethTxHashes := trackers[symbol2token["WETH"]].SanitizedTxHashes
	wethQuantities := trackers[symbol2token["WETH"]].Quantities

	fmt.Fprintf(f, "tx_hash,quantity\n")

	for i := range wethTxHashes {
		fmt.Fprintf(f, "%s,%s\n", wethTxHashes[i], wethQuantities[i])
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
}
