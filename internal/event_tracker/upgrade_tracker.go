package event_tracker

import (
	"context"
	"ronin-bridge-snapshot/generated/contract/transparent_proxy_v2"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

type UpgradeInfo struct {
	impl common.Address
}

type UpgradeTracker struct {
	*Tracker
	txHashes2Info map[common.Hash]*UpgradeInfo
}

func NewUpgradeTracker(wg *sync.WaitGroup, ctx context.Context, nWorker int, in <-chan *types.Log) *UpgradeTracker {
	u := &UpgradeTracker{
		Tracker:       NewTracker(wg, ctx, nWorker, in, nil),
		txHashes2Info: make(map[common.Hash]*UpgradeInfo),
	}

	u.Tracker.iface, _ = transparent_proxy_v2.TransparentProxyV2MetaData.GetAbi()
	u.Tracker.callback = u.Record

	return u
}

func (u *UpgradeTracker) Summarize() {
	log.Info("\nSummarizing Upgrade events")

	count := u.Total()
	if count > 0 {
		log.Info("Total Upgrade events recorded", "count", count)
	} else {
		log.Info("No Upgrade events recorded")
	}

	for _, re := range u.Tracker.rawEvts {
		if info, ok := u.txHashes2Info[re.TxHash]; ok {
			log.Info("Upgrade event", "txHash", re.TxHash.Hex(), "block", re.BlockNumber, "impl", info.impl.Hex())
		} else {
			log.Warn("Missing UpgradeInfo for txHash", "txHash", re.TxHash.Hex())
		}
	}
}

func (u *UpgradeTracker) Record(e *types.Log) error {
	log.Trace("Processing Upgrade event", "txHash", e.TxHash.Hex())
	event := new(transparent_proxy_v2.TransparentProxyV2Upgraded)
	parseEvent(u.iface, event, "Upgraded", e)
	event.Raw = *e

	if err := u.Tracker.Record(e); err != nil {
		return err
	}

	u.mu.Lock()
	defer u.mu.Unlock()
	u.txHashes2Info[e.TxHash] = &UpgradeInfo{
		impl: event.Implementation,
	}

	return nil
}
