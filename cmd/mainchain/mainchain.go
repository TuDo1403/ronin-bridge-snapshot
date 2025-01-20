package mainchain

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"ronin-bridge-snapshot/internal/abi/mainchain_gateway"
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
func RecordReleasedOnMainchain(
	ctx context.Context,
	wg *sync.WaitGroup,
	done chan struct{},
	trackers map[common.Address]*util.Tracker,
	clients []*ethclient.Client,
	nwCfg config.NetworkConfig,
) {
	// Calculate how many batches we'll do
	numBatch := (nwCfg.EndBlock-nwCfg.StartBlock)/uint64(nwCfg.QueryBatchSize) + 1
	var processedBatches int64
	log.Info("Config", "BatchCount", numBatch)

	// Concurrency limit for block-range queries
	concurrencyLimit := 2
	sem := make(chan struct{}, concurrencyLimit)

	rawLogChan := make(chan []types.Log, 10000)
	sanitizedLogChan := make(chan map[common.Address][]*mainchain_gateway.MainchainGatewayWithdrew, 10000)
	sanitizedTxHashesChan := make(chan map[common.Address][]common.Hash, 10000)
	receiptHashesChan := make(chan map[common.Address][]common.Hash, 10000)
	receiptIdsChan := make(chan map[common.Address][]*big.Int, 10000)
	quantitiesChan := make(chan map[common.Address][]*big.Int, 10000)

	// 1) Goroutine: Sanitize logs (consumes rawLogChan, produces sanitized logs/txhashes)
	wg.Add(1)
	go func() {
		defer wg.Done()
		util.SanitizeWithdrewLogs(
			ctx,
			done,
			rawLogChan,
			sanitizedLogChan,
			sanitizedTxHashesChan,
			receiptHashesChan,
			receiptIdsChan,
			quantitiesChan,
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

			case withdrawalEvents := <-sanitizedLogChan:
				// Update trackers for each token
				for token, withdrawalEventBatch := range withdrawalEvents {
					mu.Lock()
					util.RecordWithdrawals(trackers[token], withdrawalEventBatch)
					mu.Unlock()
				}

				atomic.AddInt64(&processedBatches, 1)
				log.Trace("Processed sanitized logs", "Count", len(withdrawalEvents), "BatchCount", atomic.LoadInt64(&processedBatches))

			case sanitizedTxHashes := <-sanitizedTxHashesChan:
				for token, sanitizedTxHashesBatch := range sanitizedTxHashes {
					mu.Lock()
					util.RecordSanitizedTxHashes(trackers[token], sanitizedTxHashesBatch)
					mu.Unlock()
				}

			case receiptHashes := <-receiptHashesChan:
				for token, receiptHashesBatch := range receiptHashes {
					mu.Lock()
					util.RecordReceiptHashes(trackers[token], receiptHashesBatch)
					mu.Unlock()
				}

			case receiptIds := <-receiptIdsChan:
				for token, receiptIdsBatch := range receiptIds {
					mu.Lock()
					util.RecordReceiptIds(trackers[token], receiptIdsBatch)
					mu.Unlock()
				}

			case quantities := <-quantitiesChan:
				for token, quantitiesBatch := range quantities {
					mu.Lock()
					util.RecordQuantities(trackers[token], quantitiesBatch)
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

			util.FilterWithdrew(
				ctx,
				done,
				clients,
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
	close(sanitizedTxHashesChan)

	names := make(map[common.Address]string, len(nwCfg.Tokens))
	symbols := make(map[common.Address]string, len(nwCfg.Tokens))
	decimals := make(map[common.Address]int, len(nwCfg.Tokens))
	symbol2Token := make(map[string]common.Address, len(nwCfg.Tokens))

	// Sort token symbols
	sort.Slice(nwCfg.Tokens, func(i, j int) bool {
		return symbols[nwCfg.Tokens[i]] < symbols[nwCfg.Tokens[j]]
	})

	// Fetch token metadata
	for _, token := range nwCfg.Tokens {
		name, symbol, dec := util.FetchERC20Metadata(ctx, clients[len(clients)-1], token)
		names[token] = name
		symbols[token] = symbol
		decimals[token] = dec
		symbol2Token[symbol] = token
	}

	// Log final state
	for token, tracker := range trackers {
		log.Info("Snapshot", "LastBlock", tracker.LastRecordedBlock, "Addr", token, "Symbol", symbols[token], "TxCount", tracker.TxCount, "TotalWithdraw", util.BigIntToFloat(tracker.AccAmount, decimals[token]))
	}

	// Write TxHashes to file
	f, err := os.Create("mainchain-weth-withdrawals.txt")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	wethTxHashes := trackers[symbol2Token["WETH"]].SanitizedTxHashes
	wethReceiptHashes := trackers[symbol2Token["WETH"]].ReceiptHashes
	wethReceiptIds := trackers[symbol2Token["WETH"]].ReceiptIds
	wethQuantities := trackers[symbol2Token["WETH"]].Quantities

	for i := range wethTxHashes {
		fmt.Fprintln(f, wethTxHashes[i], wethReceiptHashes[i], wethReceiptIds[i], wethQuantities[i])
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
}
