package event_tracker

import (
	"context"
	"ronin-bridge-snapshot/generated/contract/transparent_proxy_v2"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

type ChangeAdminInfo struct {
	previousAdmin common.Address
	newAdmin      common.Address
	rawEvent      *types.Log
}

type ChangeAdminTracker struct {
	*Tracker
	txHashes2Info map[common.Hash]*ChangeAdminInfo
}

func NewChangeAdminTracker(ctx context.Context, nWorker int, in <-chan *types.Log) *ChangeAdminTracker {
	c := &ChangeAdminTracker{
		Tracker:       NewTracker(ctx, "ChangeAdmin", nWorker, in, nil),
		txHashes2Info: make(map[common.Hash]*ChangeAdminInfo),
	}

	c.Tracker.iface, _ = transparent_proxy_v2.TransparentProxyV2MetaData.GetAbi()
	c.Tracker.callback = c.Record

	return c
}

func (c *ChangeAdminTracker) Summarize() {
	log.Info("#### Summarizing ChangeAdmin events ####")

	count := c.Total()
	log.Info("Total ChangeAdmin events recorded", "count", count)

	if count == 0 {
		return
	}

	changeAdminInfos := make([]*ChangeAdminInfo, 0, len(c.txHashes2Info))
	for _, info := range c.txHashes2Info {
		changeAdminInfos = append(changeAdminInfos, info)
	}
	// sort by block number
	sort.Slice(changeAdminInfos, func(i, j int) bool {
		return cmp(changeAdminInfos[i].rawEvent, changeAdminInfos[j].rawEvent)
	})

	for _, info := range changeAdminInfos {
		log.Info("ChangeAdmin event", "previousAdmin", info.previousAdmin.Hex(), "newAdmin", info.newAdmin.Hex(), "txHash", info.rawEvent.TxHash.Hex(), "blockNumber", info.rawEvent.BlockNumber, "txIndex", info.rawEvent.TxIndex)
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
		rawEvent:      e,
	}

	return nil
}
