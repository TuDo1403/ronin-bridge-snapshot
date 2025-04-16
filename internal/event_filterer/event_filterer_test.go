package event_filterer_test

import (
	"context"
	"os"
	"ronin-bridge-snapshot/internal/abi/erc20"
	"ronin-bridge-snapshot/internal/erc20_transfer"
	"ronin-bridge-snapshot/internal/event_filterer"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
)

func BenchmarkLegacyFilter(b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stdout, log.LevelInfo, true))
	log.SetDefault(logger)

	client, err := ethclient.Dial("http://190.102.110.75:8545")
	if err != nil {
		b.Fatalf("dial RPC: %v", err)
	}

	numBatch := (41758294-14765762)/uint64(5000) + 1

	// Create a progress bar with the total set to the number of queries.
	bar := progressbar.NewOptions(int(numBatch),
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionSetDescription("[ronin] fetching events legacy"),
		progressbar.OptionSetWidth(15),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowCount(),
		progressbar.OptionClearOnFinish(),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
	)

	// erc20ABI, _ := erc20.Erc20MetaData.GetAbi()
	targets := []common.Address{
		common.HexToAddress("0x0b7007c13325c48911f73a2dad5fa5dcbf808adc"),
		common.HexToAddress("0x18d2bdef572c67127e218c425f546fe64430a92c"),
		common.HexToAddress("0x1c306872bc82525d72bf3562e8f0aa3f8f26e857"),
		common.HexToAddress("0x294311a8c37f0744f99eb152c419d4d3d6fec1c7"),
		common.HexToAddress("0x7894b3088d069e70895effa4e8f7d2c243fd04c1"),
		common.HexToAddress("0x7e73630f81647bcfd7b1f2c04c1c662d17d4577e"),
		common.HexToAddress("0x7eae20d11ef8c779433eb24503def900b9d28ad7"),
		common.HexToAddress("0x97a9107c1793bc407d6f527b77e7fff4d812bece"),
		common.HexToAddress("0xa8754b9fa15fc18bb59458815510e40a12cd2014"),
		common.HexToAddress("0xc99a6a985ed2cac1ef41640596c5a5f9f4e19ef5"),
		common.HexToAddress("0xd61bbbb8369c46c15868ad9263a2710aced156c4"),
		common.HexToAddress("0xf80132fc0a86add011bffce3aedd60a86e3d704d"),
		common.HexToAddress("0x1a89ecd466a23e98f07111b0510a2d6c1cd5e400"),
	}
	// topics := [][]common.Hash{{erc20ABI.Events["Transfer"].ID}, {}, {common.BytesToHash(common.HexToAddress("0x0CF8fF40a508bdBc39fBe1Bb679dCBa64E65C7Df").Bytes())}}

	outCh := make(chan []types.Log, 10_000)
	wg := &sync.WaitGroup{}

	// consumer
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case _, ok := <-outCh:
				if !ok {
					b.Log("outCh closed")
					return
				}
			}
		}
	}()

	// Concurrency limit for block-range queries
	concurrencyLimit := 20
	sem := make(chan struct{}, concurrencyLimit)

	for startBlock := 14765762; startBlock < 41758294; startBlock += 5000 {
		endBlock := min(startBlock+5000-1, 41758294)

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
				targets,
				common.HexToAddress("0x0CF8fF40a508bdBc39fBe1Bb679dCBa64E65C7Df"),
				uint64(sb),
				uint64(eb),
				outCh,
			)

			// Update progress bar
			bar.Add(1)
		}()
	}

	// Wait for all batch goroutines to finish
	wg.Wait()

	// Now we know no more logs will be produced, so we can close the channels.
	close(outCh)
}

func BenchmarkFilter(b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stdout, log.LevelInfo, true))
	log.SetDefault(logger)

	client, err := ethclient.Dial("http://190.102.110.75:8545")
	if err != nil {
		b.Fatalf("dial RPC: %v", err)
	}

	erc20ABI, _ := erc20.Erc20MetaData.GetAbi()
	targets := []common.Address{
		common.HexToAddress("0x0b7007c13325c48911f73a2dad5fa5dcbf808adc"),
		common.HexToAddress("0x18d2bdef572c67127e218c425f546fe64430a92c"),
		common.HexToAddress("0x1c306872bc82525d72bf3562e8f0aa3f8f26e857"),
		common.HexToAddress("0x294311a8c37f0744f99eb152c419d4d3d6fec1c7"),
		common.HexToAddress("0x7894b3088d069e70895effa4e8f7d2c243fd04c1"),
		common.HexToAddress("0x7e73630f81647bcfd7b1f2c04c1c662d17d4577e"),
		common.HexToAddress("0x7eae20d11ef8c779433eb24503def900b9d28ad7"),
		common.HexToAddress("0x97a9107c1793bc407d6f527b77e7fff4d812bece"),
		common.HexToAddress("0xa8754b9fa15fc18bb59458815510e40a12cd2014"),
		common.HexToAddress("0xc99a6a985ed2cac1ef41640596c5a5f9f4e19ef5"),
		common.HexToAddress("0xd61bbbb8369c46c15868ad9263a2710aced156c4"),
		common.HexToAddress("0xf80132fc0a86add011bffce3aedd60a86e3d704d"),
		common.HexToAddress("0x1a89ecd466a23e98f07111b0510a2d6c1cd5e400"),
	}
	topics := [][]common.Hash{{erc20ABI.Events["Transfer"].ID}, {}, {common.BytesToHash(common.HexToAddress("0x0CF8fF40a508bdBc39fBe1Bb679dCBa64E65C7Df").Bytes())}}

	outCh := make(chan []*types.Log, 10_000)
	wg := &sync.WaitGroup{}

	// consumer
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				b.Log("context done")
				return
			case _, ok := <-outCh:
				if !ok {
					b.Log("outCh closed")
					return
				}
			}
		}
	}()

	filterer := event_filterer.NewEventFilterer(
		ctx, wg,
		100*time.Millisecond,
		"[ronin] fetching events",
		[]*ethclient.Client{client},
		targets, topics,
		10000, // batch size
		100,   // workers
		14_765_762,
		41758294,
		// 14815762,
		outCh,
	)

	filterer.Start()

	wg.Wait()
}
