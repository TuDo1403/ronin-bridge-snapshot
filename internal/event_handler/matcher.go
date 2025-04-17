package event_handler

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

/* ───────────────────────────  Matcher  ──────────────────────────── */

type Matcher struct {
	froms     map[common.Address]struct{} // O(1) address check
	topic0    common.Hash
	topic1    *common.Hash
	topic2    *common.Hash
	topic3    *common.Hash
	outCh     chan *types.Log
	closeOnce sync.Once
}

func NewMatcher(addrs []common.Address, topic0 common.Hash, t1, t2, t3 *common.Hash) *Matcher {

	set := make(map[common.Address]struct{}, len(addrs))
	for _, a := range addrs {
		set[a] = struct{}{}
	}
	m := &Matcher{
		froms:  set,
		topic0: topic0,
		topic1: t1,
		topic2: t2,
		topic3: t3,
		outCh:  make(chan *types.Log, 10_000), // buffered channel
	}

	log.Info("Matcher created", "topic0", topic0.Hex(), "topic1", t1, "topic2", t2, "topic3", t3)

	return m
}

func (m *Matcher) Topic0() common.Hash {
	return m.topic0
}

func (m *Matcher) Topic1() *common.Hash {
	return m.topic1
}

func (m *Matcher) Topic2() *common.Hash {
	return m.topic2
}

func (m *Matcher) Topic3() *common.Hash {
	return m.topic3
}

func (m *Matcher) Froms() []common.Address {
	addrs := make([]common.Address, 0, len(m.froms))

	for addr := range m.froms {
		addrs = append(addrs, addr)
	}

	return addrs
}

func (m *Matcher) ReceiveOnlyCh() <-chan *types.Log {
	return m.outCh
}

func (m *Matcher) Close() {
	m.closeOnce.Do(func() {
		close(m.outCh)
	})
}
