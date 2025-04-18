package event_tracker

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

var (
	errNoEventSignature       = errors.New("no event signature")
	errEventSignatureMismatch = errors.New("event signature mismatch")
)

type EventTrackerInterface interface {
	Record(e *types.Log) error
}

type Tracker struct {
	iface      *abi.ABI
	inCh       <-chan *types.Log
	rawEvts    []*types.Log
	mu         sync.Mutex
	txHash2Log map[common.Hash]*types.Log

	ctx      context.Context
	wg       *sync.WaitGroup
	nWorker  int
	callback func(e *types.Log) error
	desc     string
}

func NewTracker(ctx context.Context, desc string, nWorker int, in <-chan *types.Log, cb func(e *types.Log) error) *Tracker {
	return &Tracker{
		inCh:       in,
		rawEvts:    make([]*types.Log, 0),
		txHash2Log: make(map[common.Hash]*types.Log),
		ctx:        ctx,
		wg:         &sync.WaitGroup{},
		nWorker:    nWorker,
		callback:   cb,
		desc:       desc,
	}
}

func (tr *Tracker) GetAbi() *abi.ABI {
	return tr.iface
}

func (tr *Tracker) SetInCh(in <-chan *types.Log) {
	tr.inCh = in
}

func (tr *Tracker) Start() {
	tr.wg.Add(tr.nWorker)

	for i := range tr.nWorker {
		go tr.worker(i)
	}
}

func (tr *Tracker) Stop() {
	tr.wg.Wait()
	log.Info("tracker", fmt.Sprintf("Tracker stopped: %s Stopped!", tr.desc))
}

func (tr *Tracker) worker(i int) {
	defer tr.wg.Done()

	for {
		select {
		case <-tr.ctx.Done():
			log.Info("Worker stopped", "worker", i)
			return
		case e, ok := <-tr.inCh:
			if !ok {
				return
			}

			err := tr.callback(e)
			if err != nil {
				log.Crit("Failed to record event", "error", err, "event", e)
			}
		}
	}
}

func (tr *Tracker) Summarize() {
	// Log total number of events
	log.Info("Total events recorded", "count", tr.Total())
	// Log the oldest and latest transaction hashes
	oldest, latest := tr.OldestAndLatestTxHash()
	log.Info("Oldest and latest transaction hashes", "oldest", oldest.Hex(), "latest", latest.Hex())
}

func (tr *Tracker) Total() int {
	return len(tr.rawEvts)
}

func (tr *Tracker) OldestAndLatestTxHash() (oldest, latest common.Hash) {
	tr.mu.Lock()
	defer tr.mu.Unlock()

	if len(tr.rawEvts) == 0 {
		return common.Hash{}, common.Hash{}
	}

	// Create a copy of rawEvts to avoid modifying the original slice
	sortedLogs := make([]*types.Log, len(tr.rawEvts))
	copy(sortedLogs, tr.rawEvts)

	// Sort the logs using Cmp
	sort.Slice(sortedLogs, func(i, j int) bool {
		return cmp(sortedLogs[i], sortedLogs[j])
	})

	oldestLog := sortedLogs[0]
	latestLog := sortedLogs[len(sortedLogs)-1]

	return oldestLog.TxHash, latestLog.TxHash
}

func (tr *Tracker) Record(e *types.Log) error {
	tr.mu.Lock()
	tr.rawEvts = append(tr.rawEvts, e)
	tr.txHash2Log[e.TxHash] = e
	tr.mu.Unlock()

	return nil
}

// cmp compares two logs based on their block number and transaction index
func cmp(a, b *types.Log) bool {
	if a.BlockNumber == b.BlockNumber {
		return a.TxIndex < b.TxIndex
	}

	return a.BlockNumber < b.BlockNumber
}

func parseEvent(a *abi.ABI, out any, event string, rawEvent *types.Log) {
	err := unpackLog(a, out, event, *rawEvent)
	if err != nil {
		log.Error("Failed to unpack log", "error", err, "log", rawEvent)
	}
}

// UnpackLog unpacks a retrieved log into the provided output structure.
func unpackLog(a *abi.ABI, out any, event string, e types.Log) error {
	// Anonymous events are not supported.
	if len(e.Topics) == 0 {
		return errNoEventSignature
	}
	if e.Topics[0] != a.Events[event].ID {
		return errEventSignatureMismatch
	}
	if len(e.Data) > 0 {
		if err := a.UnpackIntoInterface(out, event, e.Data); err != nil {
			return err
		}
	}

	var indexed abi.Arguments
	for _, arg := range a.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}

	return abi.ParseTopics(out, indexed, e.Topics[1:])
}
