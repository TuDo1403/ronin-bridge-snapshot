package event_filterer_test

import (
	"context"
	"os"
	"ronin-bridge-snapshot/generated/contract/erc20"
	"ronin-bridge-snapshot/generated/contract/ronin_gateway_v3"
	"ronin-bridge-snapshot/generated/contract/transparent_proxy_v2"
	"ronin-bridge-snapshot/internal/event_filterer"
	"ronin-bridge-snapshot/internal/event_handler"
	"ronin-bridge-snapshot/internal/event_tracker"
	"ronin-bridge-snapshot/internal/util"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

// func BenchmarkLegacyFilter(b *testing.B) {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	logger := log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stdout, log.LevelInfo, true))
// 	log.SetDefault(logger)

// 	client, err := ethclient.Dial("http://190.102.110.75:8545")
// 	if err != nil {
// 		b.Fatalf("dial RPC: %v", err)
// 	}

// 	numBatch := (41758294-14765762)/uint64(5000) + 1

// 	// Create a progress bar with the total set to the number of queries.
// 	bar := progressbar.NewOptions(int(numBatch),
// 		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
// 		progressbar.OptionSetDescription("[ronin] fetching events legacy"),
// 		progressbar.OptionSetWidth(15),
// 		progressbar.OptionEnableColorCodes(true),
// 		progressbar.OptionShowCount(),
// 		progressbar.OptionClearOnFinish(),
// 		progressbar.OptionSetTheme(progressbar.Theme{
// 			Saucer:        "[green]=[reset]",
// 			SaucerHead:    "[green]>[reset]",
// 			SaucerPadding: " ",
// 			BarStart:      "[",
// 			BarEnd:        "]",
// 		}),
// 	)

// 	// erc20ABI, _ := erc20.Erc20MetaData.GetAbi()
// 	targets := []common.Address{
// 		common.HexToAddress("0x0b7007c13325c48911f73a2dad5fa5dcbf808adc"),
// 		common.HexToAddress("0x18d2bdef572c67127e218c425f546fe64430a92c"),
// 		common.HexToAddress("0x1c306872bc82525d72bf3562e8f0aa3f8f26e857"),
// 		common.HexToAddress("0x294311a8c37f0744f99eb152c419d4d3d6fec1c7"),
// 		common.HexToAddress("0x7894b3088d069e70895effa4e8f7d2c243fd04c1"),
// 		common.HexToAddress("0x7e73630f81647bcfd7b1f2c04c1c662d17d4577e"),
// 		common.HexToAddress("0x7eae20d11ef8c779433eb24503def900b9d28ad7"),
// 		common.HexToAddress("0x97a9107c1793bc407d6f527b77e7fff4d812bece"),
// 		common.HexToAddress("0xa8754b9fa15fc18bb59458815510e40a12cd2014"),
// 		common.HexToAddress("0xc99a6a985ed2cac1ef41640596c5a5f9f4e19ef5"),
// 		common.HexToAddress("0xd61bbbb8369c46c15868ad9263a2710aced156c4"),
// 		common.HexToAddress("0xf80132fc0a86add011bffce3aedd60a86e3d704d"),
// 		common.HexToAddress("0x1a89ecd466a23e98f07111b0510a2d6c1cd5e400"),
// 	}
// 	// topics := [][]common.Hash{{erc20ABI.Events["Transfer"].ID}, {}, {common.BytesToHash(common.HexToAddress("0x0CF8fF40a508bdBc39fBe1Bb679dCBa64E65C7Df").Bytes())}}

// 	outCh := make(chan []types.Log, 10_000)
// 	wg := &sync.WaitGroup{}

// 	// consumer
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				return
// 			case _, ok := <-outCh:
// 				if !ok {
// 					b.Log("outCh closed")
// 					return
// 				}
// 			}
// 		}
// 	}()

// 	// Concurrency limit for block-range queries
// 	concurrencyLimit := 20
// 	sem := make(chan struct{}, concurrencyLimit)

// 	for startBlock := 14765762; startBlock < 41758294; startBlock += 5000 {
// 		endBlock := min(startBlock+5000-1, 41758294)

// 		sb, eb := startBlock, endBlock

// 		// Acquire concurrency slot
// 		sem <- struct{}{}

// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			defer func() { <-sem }()

// 			erc20_transfer.FilterERC20TransferForTokens(
// 				ctx,
// 				client,
// 				targets,
// 				common.HexToAddress("0x0CF8fF40a508bdBc39fBe1Bb679dCBa64E65C7Df"),
// 				uint64(sb),
// 				uint64(eb),
// 				outCh,
// 			)

// 			// Update progress bar
// 			bar.Add(1)
// 		}()
// 	}

// 	// Wait for all batch goroutines to finish
// 	wg.Wait()

// 	// Now we know no more logs will be produced, so we can close the channels.
// 	close(outCh)
// }

func TestFilterERC20Transfer(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stdout, log.LevelInfo, true))
	log.SetDefault(logger)

	client, err := ethclient.Dial("http://190.102.110.75:8545")
	if err != nil {
		t.Fatalf("dial RPC: %v", err)
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
	wg := &sync.WaitGroup{}

	filterer := event_filterer.NewEventFilterer(
		ctx,
		100*time.Millisecond,
		"[ronin] fetching events",
		[]*ethclient.Client{client},
		targets, topics,
		10000, // batch size
		100,   // workers
		14_765_762,
		// 41758294,
		14815762,
	)

	inCh := filterer.ReceiveOnlyCh()

	// consumer
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				t.Log("context done")
				return
			case _, ok := <-inCh:
				if !ok {
					t.Log("inCh closed")
					return
				}
			}
		}
	}()

	filterer.Start()
	defer filterer.Stop()

	wg.Wait()
}

func TestFilterWithdrawalRequestedAndAdminChanged(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stdout, log.LevelInfo, true))
	log.SetDefault(logger)

	client, err := ethclient.Dial("http://190.102.110.75:8545")
	if err != nil {
		t.Fatalf("dial RPC: %v", err)

	}

	roninGatewayABI, _ := ronin_gateway_v3.RoninGatewayV3MetaData.GetAbi()
	tpABI, _ := transparent_proxy_v2.TransparentProxyV2MetaData.GetAbi()

	withdrawalRequestedMatcher := event_handler.NewMatcher(
		[]common.Address{
			common.HexToAddress("0x0CF8fF40a508bdBc39fBe1Bb679dCBa64E65C7Df"),
		},
		roninGatewayABI.Events["WithdrawalRequested"].ID,
		nil,
		nil,
		nil,
	)
	withdrawalRequestTracker := event_tracker.NewRequestWithdrawalTracker(
		ctx,
		50,
		[]common.Hash{
			common.HexToHash("0xa53c7c4402b40a63e485a1a8798d7b7ac6a151dce19c3486d9928575476ac3ce"),
			common.HexToHash("0x6af887fda63d7ec5538c43fdd3b1fa87ce5e2aaf824c5332f3bd8330d45f4c0b"),
			common.HexToHash("0x2392f70de576a6e90d95eb4686561ab6232fda7d972ca56fa94df20840f289c9"),
			common.HexToHash("0xc27bbc806abf466645a24f2be8bb7f082cba007f52767739eee0815153ae6fea"),
			common.HexToHash("0x78c3e08640228ff072749632f407617695a731f2dd377fa3b1a87f7697282897"),
			common.HexToHash("0x542c6ddaf6d3752d488153b212346b94bbe30d8b25eee94c80051baa4a8714b7"),
			common.HexToHash("0x2c4db38cafde1d399e07a754b6605bc5dd242119929dfd8382dfb9925d322379"),
			common.HexToHash("0x534c69d4c943b20f180774b942147e446e39e12920057e9795927defa547e2ab"),
			common.HexToHash("0xced4cf252ff100eb7cc96859fc0e85fcad97d97d26f96ed4ae78665301e09094"),
			common.HexToHash("0x28211787395779c94fca595000f3d079e86ebe12b92e4f005fad4dae5c44d2e9"),
			common.HexToHash("0x4887f384ba119800b85cb65e146bb559972cc15fab65e5d8ccafd32999b419c9"),
			common.HexToHash("0x1357e4fa4b955249724d437f634105195f1337edc2db8905a72e69b35ccbd870"),
		},
		withdrawalRequestedMatcher.ReceiveOnlyCh())

	upgradedMatcher := event_handler.NewMatcher(
		[]common.Address{
			common.HexToAddress("0x0CF8fF40a508bdBc39fBe1Bb679dCBa64E65C7Df"),
		},
		tpABI.Events["Upgraded"].ID,
		nil,
		nil,
		nil,
	)
	upgradedTracker := event_tracker.NewUpgradeTracker(
		ctx,
		50, upgradedMatcher.ReceiveOnlyCh())

	adminChangedMatcher := event_handler.NewMatcher(
		[]common.Address{
			common.HexToAddress("0x0CF8fF40a508bdBc39fBe1Bb679dCBa64E65C7Df"),
		},
		tpABI.Events["AdminChanged"].ID,
		nil,
		nil,
		nil,
	)
	adminChangedTracker := event_tracker.NewChangeAdminTracker(
		ctx,
		50, adminChangedMatcher.ReceiveOnlyCh())

	targets := util.AggregateAddresses(
		append(append(withdrawalRequestedMatcher.Froms(), upgradedMatcher.Froms()...), adminChangedMatcher.Froms()...)...,
	)
	topics := util.AggregateTopics(
		[]common.Hash{
			withdrawalRequestedMatcher.Topic0(),
			upgradedMatcher.Topic0(),
			adminChangedMatcher.Topic0(),
		},
		[]common.Hash{},
		[]common.Hash{},
		[]common.Hash{},
	)

	filterer := event_filterer.NewEventFilterer(
		ctx,
		300*time.Millisecond,
		"[ronin] fetching events",
		[]*ethclient.Client{client},
		targets, topics,
		5000,       // batch size
		20,         // workers
		14_765_762, // 14765762
		41758294,
		// 15_765_762,
	)

	handler := event_handler.NewEventHandler(
		ctx,
		[]<-chan []*types.Log{filterer.ReceiveOnlyCh()},
		10, // workers
	)

	handler.AddMatcher(withdrawalRequestedMatcher)
	handler.AddMatcher(upgradedMatcher)
	handler.AddMatcher(adminChangedMatcher)

	withdrawalRequestTracker.Start()
	adminChangedTracker.Start()
	upgradedTracker.Start()
	handler.Start()
	filterer.Start()

	filterer.Stop()
	handler.Stop()
	upgradedTracker.Stop()
	adminChangedTracker.Stop()
	withdrawalRequestTracker.Stop()

	log.Info("All workers finished")

	// Summarize the trackers
	withdrawalRequestTracker.Summarize()
	upgradedTracker.Summarize()
	adminChangedTracker.Summarize()
}
