package event_tracker

import (
	"context"
	"ronin-bridge-snapshot/generated/contract/transparent_proxy_v2"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

type ChangeAdminInfo struct {
	previousAdmin common.Address
	newAdmin      common.Address
}

type ChangeAdminTracker struct {
	*Tracker
	txHashes2Info map[common.Hash]*ChangeAdminInfo
}

func NewChangeAdminTracker(wg *sync.WaitGroup, ctx context.Context, nWorker int, in <-chan *types.Log) *ChangeAdminTracker {
	c := &ChangeAdminTracker{
		Tracker:       NewTracker(wg, ctx, nWorker, in, nil),
		txHashes2Info: make(map[common.Hash]*ChangeAdminInfo),
	}

	c.Tracker.iface, _ = transparent_proxy_v2.TransparentProxyV2MetaData.GetAbi()
	c.Tracker.callback = c.Record

	return c
}

func (c *ChangeAdminTracker) Summarize() {
	log.Info("\nSummarizing ChangeAdmin events")

	count := c.Total()
	if count > 0 {
		log.Info("Total ChangeAdmin events recorded", "count", count)
	} else {
		log.Info("No ChangeAdmin events recorded")
	}

	for _, re := range c.Tracker.rawEvts {
		if info, ok := c.txHashes2Info[re.TxHash]; ok {
			log.Info("ChangeAdmin event", "txHash", re.TxHash.Hex(), "block", re.BlockNumber, "previousAdmin", info.previousAdmin.Hex(), "newAdmin", info.newAdmin.Hex())
		} else {
			log.Warn("Missing ChangeAdminInfo for txHash", "txHash", re.TxHash.Hex())
		}
	}

}

func (c *ChangeAdminTracker) Record(e *types.Log) error {
	log.Trace("Processing ChangeAdmin event", "txHash", e.TxHash.Hex())

	event := new(transparent_proxy_v2.TransparentProxyV2AdminChanged)
	parseEvent(c.iface, event, "AdminChanged", e)
	event.Raw = *e

	if err := c.Tracker.Record(e); err != nil {
		return err
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.txHashes2Info[e.TxHash] = &ChangeAdminInfo{
		previousAdmin: event.PreviousAdmin,
		newAdmin:      event.NewAdmin,
	}

	return nil
}
