package event_filterer

import (
	"context"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
)

// EventFilterer is NOT safe for reuse – create one per batch run.
type EventFilterer struct {
	queries []*ethereum.FilterQuery
	clients []*ethclient.Client
	outCh   chan []*types.Log

	nWorker      int
	batchSize    int
	pollInterval time.Duration
	desc         string

	ctx context.Context
	wg  *sync.WaitGroup

	// fulfilled counts how many queries have been processed
	fulfilled atomic.Int32
	// nextQueryIndex atomically assigns the next query index to a worker.
	nextQueryIndex atomic.Int32
	// closeOnce ensures outCh is closed exactly one time.
	closeOnce sync.Once

	progressBar *progressbar.ProgressBar
}

// NewEventFilterer builds the list of filter queries once and spins workers later.
func NewEventFilterer(
	ctx context.Context,
	wg *sync.WaitGroup,
	pollInterval time.Duration,
	desc string,
	clients []*ethclient.Client,
	targets []common.Address,
	topics [][]common.Hash,
	batchSize, nWorker, start, end int,
	outCh chan []*types.Log,
) (ef *EventFilterer) {
	if batchSize <= 0 || start <= 0 || end < start {
		log.Crit("invalid block range", "start", start, "end", end, "batchSize", batchSize)
	}
	if len(clients) == 0 {
		log.Crit("no RPC clients provided")
	}
	if nWorker < 1 {
		log.Crit("workers must be >= 1")
	}

	nBatch := (end-start)/batchSize + 1
	queries := make([]*ethereum.FilterQuery, nBatch)
	for i := range nBatch {
		bs := start + i*batchSize
		be := min(bs+batchSize-1, end)
		queries[i] = buildQuery(targets, bs, be, topics)
	}

	// Ensure the number of workers does not exceed the number of batches.
	nWorker = min(nWorker, nBatch)

	// Create a progress bar with the total set to the number of queries.
	bar := progressbar.NewOptions(nBatch,
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionSetDescription(desc),
		progressbar.OptionSetWidth(15),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowCount(),
		progressbar.OptionShowIts(),
		progressbar.OptionShowElapsedTimeOnFinish(),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
	)

	ef = &EventFilterer{
		queries:      queries,
		clients:      clients,
		outCh:        outCh,
		nWorker:      nWorker,
		batchSize:    batchSize,
		pollInterval: pollInterval,
		desc:         desc,
		ctx:          ctx,
		wg:           wg,
		progressBar:  bar,
	}

	log.Info("EventFilterer", "batchSize", batchSize, "nBatch", nBatch, "nWorker", nWorker, "pollInterval", pollInterval, "from", start, "to", end)

	return ef
}

// Start launches worker goroutines. The caller must call wg.Wait() later.
func (ef *EventFilterer) Start() {
	ef.wg.Add(ef.nWorker)
	for i := range ef.nWorker {
		go ef.worker(i)
	}
}

// worker continually processes queries until none remain.
func (ef *EventFilterer) worker(i int) {
	defer ef.wg.Done()

	for {
		// Exit if context is canceled.
		select {
		case <-ef.ctx.Done():
			log.Debug("context canceled", "worker id", i)
			return
		default:
		}

		// Atomically fetch the next query index.
		idx := int(ef.nextQueryIndex.Add(1)) - 1
		if idx >= len(ef.queries) {
			log.Debug("no more queries to process", "worker", i)
			return
		}

		// Attempt to fetch logs for the current query.
		_ = ef.fetchLogs(idx, i)

		// Update the progress bar and count the processed query.
		ef.progressBar.Add(1)
		if ef.fulfilled.Add(1) == int32(len(ef.queries)) {
			ef.closeOnce.Do(func() {
				close(ef.outCh)
			})

			ef.progressBar.Finish()
			return
		}

		// Optionally throttle requests to avoid tight looping.
		time.Sleep(ef.pollInterval)
	}
}

// fetchLogs attempts to fetch logs for the query at index i using each client until one succeeds.
func (ef *EventFilterer) fetchLogs(i, wid int) bool {
	q := ef.queries[i]
	log.Trace("fetchLogs", "worker id", wid, "idx", i, "from", q.FromBlock, "to", q.ToBlock)

	var success bool
	// Try each client in order.
	for cIdx, cli := range ef.clients {
		res, err := cli.FilterLogs(ef.ctx, *q)
		if err != nil {
			log.Warn("filterLogs failed", "worker id", wid, "client", cIdx, "err", err)
			continue
		}

		// Send logs to the output channel.
		if len(res) != 0 {
			log.Trace("filterLogs succeeded", "worker id", wid, "client", cIdx, "count", len(res), "from", q.FromBlock, "to", q.ToBlock)
			out := make([]*types.Log, len(res))
			for j := range res {
				out[j] = &res[j]
			}

			ef.outCh <- out
		}

		success = true

		break
	}

	if !success {
		log.Crit("all clients failed to fetch logs", "worker id", wid, "idx", i, "from", q.FromBlock, "to", q.ToBlock)
	}

	return success
}

func buildQuery(addrs []common.Address, from, to int, topics [][]common.Hash) *ethereum.FilterQuery {
	if len(addrs) == 0 {
		log.Crit("no addresses provided")
	}
	if len(topics) == 0 {
		log.Crit("no topics provided")
	}
	if from <= 0 || to < from {
		log.Crit("invalid block range", "from", from, "to", to)
	}

	return &ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(from)),
		ToBlock:   big.NewInt(int64(to)),
		Addresses: addrs,
		Topics:    topics,
	}
}
