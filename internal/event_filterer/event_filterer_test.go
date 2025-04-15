package event_filterer_test

import (
	"context"
	"os"
	"ronin-bridge-snapshot/internal/abi/erc20"
	"ronin-bridge-snapshot/internal/event_filterer"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

func TestFilter(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stdout, log.LevelTrace, true))
	log.SetDefault(logger)

	client, err := ethclient.Dial("http://190.102.110.75:8545")
	if err != nil {
		t.Fatalf("dial RPC: %v", err)
	}

	erc20ABI, _ := erc20.Erc20MetaData.GetAbi()
	targets := []common.Address{common.HexToAddress("0xe514d9deb7966c8be0ca922de8a064264ea6bcd4")}
	topics := [][]common.Hash{{erc20ABI.Events["Transfer"].ID}}

	outCh := make(chan []*types.Log, 10_000)
	wg := &sync.WaitGroup{}

	// consumer
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case evts, ok := <-outCh:
				if !ok {
					t.Log("outCh closed")
					return
				}
				t.Logf("got %d logs", len(evts))
			}
		}
	}()

	filterer := event_filterer.NewEventFilterer(
		ctx, wg,
		[]*ethclient.Client{client},
		targets, topics,
		500, // batch size
		10,  // workers
		14_765_762, 14_775_762,
		outCh,
	)
	filterer.Start()

	wg.Wait() // <- crucial!
}
