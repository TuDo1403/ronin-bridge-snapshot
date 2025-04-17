package event_handler

import (
	"context"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

/* ─────────────────────────  EventHandler  ───────────────────────── */

type EventHandler struct {
	inCh    <-chan []*types.Log // fan‑in from filterers
	ctx     context.Context
	wg      *sync.WaitGroup
	workers int

	// index: topic0  → []*Matcher
	byTopic map[common.Hash][]*Matcher
}

/* ---------- constructor & lifecycle ----------------------------------- */

func NewEventHandler(ctx context.Context, in <-chan []*types.Log, nWorker int) *EventHandler {
	eh := &EventHandler{
		inCh:    in,
		workers: nWorker,
		ctx:     ctx,
		wg:      &sync.WaitGroup{},
		byTopic: make(map[common.Hash][]*Matcher),
	}

	return eh
}

// AddMatcher may be called before Start().
// If you need hot‑plugging after Start, protect by a RW‑mutex.
func (eh *EventHandler) AddMatcher(m *Matcher) {
	eh.byTopic[m.topic0] = append(eh.byTopic[m.topic0], m)
}

func (eh *EventHandler) GetMatchers() []*Matcher {
	matchers := make([]*Matcher, 0)
	for _, m := range eh.byTopic {
		matchers = append(matchers, m...)
	}
	return matchers
}

func (eh *EventHandler) Stop() {
	eh.wg.Wait()
	log.Info("[EventHandler] All workers stopped")

	matchers := eh.GetMatchers()
	for _, m := range matchers {
		m.Close()
	}
	log.Info("[EventHandler] All matchers closed")
}

func (eh *EventHandler) Start() {
	eh.wg.Add(eh.workers)

	for i := range eh.workers {
		go eh.worker(i)
	}
}

/* ------------------------------ worker -------------------------------- */

func (eh *EventHandler) worker(id int) {
	defer eh.wg.Done()

	for {
		select {
		case <-eh.ctx.Done():
			return

		case batch, ok := <-eh.inCh:
			if !ok {
				return
			}

			if len(batch) == 0 {
				continue
			}

			for _, lg := range batch {
				eh.dispatch(lg)
			}
		}
	}
}

/* ----------------------------- dispatch ------------------------------- */

func (eh *EventHandler) dispatch(lg *types.Log) {
	if lg.Removed {
		log.Warn("event‑handler", "msg", "log removed (reorg)", "hash", lg.TxHash)
		return
	}
	if lg.TxHash == (common.Hash{}) {
		log.Warn("event‑handler", "msg", "log without txHash", "block", lg.BlockNumber, "logIndex", lg.Index)
		return
	}
	if lg.Address == (common.Address{}) {
		log.Warn("event‑handler", "msg", "log without address", "block", lg.BlockNumber, "logIndex", lg.Index)
		return
	}

	candidates := eh.byTopic[lg.Topics[0]]
	if len(candidates) == 0 {
		return // nothing registered for this topic0
	}

	for _, m := range candidates {
		// quick address vector check
		if _, ok := m.froms[lg.Address]; !ok {
			continue
		}

		if !matchTopic(lg, 1, m.topic1) || !matchTopic(lg, 2, m.topic2) || !matchTopic(lg, 3, m.topic3) {
			continue
		}

		m.outCh <- lg // send to the matcher
	}
}

/* --------------------------- helpers ---------------------------------- */

func matchTopic(lg *types.Log, idx int, want *common.Hash) bool {
	if want == nil {
		return true
	}
	// guard against short Topics slices
	if len(lg.Topics) <= idx {
		return false
	}
	return lg.Topics[idx] == *want
}

/* ──────────────────────────  usage sketch  ─────────────────────────────

tracker := NewRequestWithdrawalTracker(wg, ctx, 4)        // your previous code
handler := NewEventHandler(ctx, eventFilterer.Out(), 4)   // Out() => <-chan []*types.Log
handler.AddMatcher(
    NewMatcher(
        []common.Address{contractAddr},
        ronin_gateway.RoninGatewayABI.Events["WithdrawalRequested"].ID,
        nil, nil, nil,
        tracker.filteredCh,           // tracker will consume matched logs
))
handler.Start()
defer handler.Stop()

*/
