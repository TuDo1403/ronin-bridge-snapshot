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
)

type Status uint8

const (
	None    Status = iota // query not yet picked up
	Started               // query is being processed
	Done                  // query completed
)

func (s Status) String() string {
	switch s {
	case None:
		return "None"
	case Started:
		return "Started"
	case Done:
		return "Done"
	default:
		return "Unknown"
	}
}

// EventFilterer is **NOT** safe for reuse – create one per batch run.
type EventFilterer struct {
	queries   []*ethereum.FilterQuery
	clients   []*ethclient.Client
	status    []Status
	outCh     chan []*types.Log
	nWorker   int
	batchSize int

	ctx context.Context
	wg  *sync.WaitGroup

	fulfilled atomic.Int32
	mu        sync.Mutex
}

// NewEventFilterer builds the list of filter queries once and spins workers later.
func NewEventFilterer(
	ctx context.Context,
	wg *sync.WaitGroup,
	clients []*ethclient.Client,
	targets []common.Address,
	topics [][]common.Hash,
	batchSize, nWorker, start, end int,
	outCh chan []*types.Log,
) *EventFilterer {
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

	ef := &EventFilterer{
		queries:   queries,
		clients:   clients,
		status:    make([]Status, nBatch),
		outCh:     outCh,
		nWorker:   nWorker,
		batchSize: batchSize,
		ctx:       ctx,
		wg:        wg,
	}
	log.Info("EventFilterer", "nBatch", nBatch, "batchSize", batchSize, "start", start, "end", end, "nWorker", nWorker)
	return ef
}

// Start launches the pool.  Caller must **wg.Wait()** later.
func (ef *EventFilterer) Start() {
	ef.wg.Add(ef.nWorker)
	for range ef.nWorker {
		go ef.worker()
	}
}

// -----------------------------------------------------------------------------
// internal helpers
// -----------------------------------------------------------------------------

func (ef *EventFilterer) worker() {
	defer ef.wg.Done()

	for {
		select {
		case <-ef.ctx.Done():
			return
		default:
		}

		if ef.processNextQuery() {
			log.Trace("worker finished all queries")
			// all batches are finished, exit this worker
			return
		}
		// small sleep prevents a tight spin once work is done
		time.Sleep(150 * time.Millisecond)
	}
}

// processNextQuery finds the next *unstarted* query and runs it.
// It returns true when **all** queries are Done (so worker can exit).
func (ef *EventFilterer) processNextQuery() (allDone bool) {
	var idx int = -1

	ef.mu.Lock()
	for i, st := range ef.status {
		if st == None {
			idx = i
			ef.status[i] = Started
			break
		}
	}
	allDone = ef.fulfilled.Load() == int32(len(ef.queries))
	ef.mu.Unlock()

	if allDone || idx == -1 {
		return allDone
	}

	ok := ef.fetchLogs(idx)
	if ok {
		if ef.fulfilled.Add(1) == int32(len(ef.queries)) {
			log.Trace("all queries fulfilled")
			// last query – close channel so downstream consumers know we're done
			close(ef.outCh)
		}
	}
	return ef.fulfilled.Load() == int32(len(ef.queries))
}

func (ef *EventFilterer) fetchLogs(i int) bool {
	log.Trace("fetchLogs", "from", ef.queries[i].FromBlock, "to", ef.queries[i].ToBlock, "idx", i)
	q := ef.queries[i]

	for cIdx, cli := range ef.clients {
		res, err := cli.FilterLogs(ef.ctx, *q)
		if err != nil {
			log.Warn("filterLogs failed", "client", cIdx, "err", err)
			continue
		}
		if len(res) == 0 {
			continue
		}

		// copy to avoid keeping the backing array of res alive
		out := make([]*types.Log, len(res))
		for j := range res {
			out[j] = &res[j]
		}
		ef.outCh <- out

		ef.mu.Lock()
		ef.status[i] = Done
		ef.mu.Unlock()

		return true
	}

	log.Warn("all clients failed to fetch logs", "idx", i)

	// no client returned anything – mark as Done so we don’t loop forever
	ef.mu.Lock()
	ef.status[i] = Done
	ef.mu.Unlock()
	return false
}

func buildQuery(addrs []common.Address, from, to int, topics [][]common.Hash) *ethereum.FilterQuery {
	if len(addrs) == 0 {
		log.Crit("no addresses provided")
	}
	if len(topics) == 0 {
		log.Crit("no topics provided")
	}
	return &ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(from)),
		ToBlock:   big.NewInt(int64(to)),
		Addresses: addrs,
		Topics:    topics,
	}
}
