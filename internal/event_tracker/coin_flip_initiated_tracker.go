package event_tracker

import (
	"context"
	"math/big"
	"ronin-bridge-snapshot/generated/contract/coin_flipper"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

type CoinFlipInitiatedInfo struct {
	Player   common.Address
	ConfigId *big.Int
	TxHash   common.Hash
	ReqHash  common.Hash
	Choice   bool
	NftId    *big.Int
	RawEvent *types.Log
}

type CoinFlipInitiatedTracker struct {
	*Tracker
	reqHashes2Info map[common.Hash]*CoinFlipInitiatedInfo
}

func NewCoinFlipInitiatedTracker(ctx context.Context, nWorker int, in <-chan *types.Log) *CoinFlipInitiatedTracker {
	c := &CoinFlipInitiatedTracker{
		Tracker:        NewTracker(ctx, "CoinFlipInitiated", nWorker, in, nil),
		reqHashes2Info: make(map[common.Hash]*CoinFlipInitiatedInfo),
	}

	c.Tracker.iface, _ = coin_flipper.CoinFlipperMetaData.GetAbi()
	c.Tracker.callback = c.Record

	return c
}

func (c *CoinFlipInitiatedTracker) Summarize() map[common.Hash]*CoinFlipInitiatedInfo {
	count := c.Total()

	if count == 0 {
		return nil
	}

	CoinFlipInitiatedInfos := make([]*CoinFlipInitiatedInfo, 0, len(c.reqHashes2Info))
	for _, info := range c.reqHashes2Info {
		CoinFlipInitiatedInfos = append(CoinFlipInitiatedInfos, info)
	}
	// sort by block number
	sort.Slice(CoinFlipInitiatedInfos, func(i, j int) bool {
		return cmp(CoinFlipInitiatedInfos[i].RawEvent, CoinFlipInitiatedInfos[j].RawEvent)
	})

	// for _, info := range CoinFlipInitiatedInfos {
	// 	log.Info("CoinFlipInitiated event", "Player", info.Player.Hex(), "ConfigId", info.ConfigId, "reqHash", info.reqHash, "choice", info.choice, "blockNumber", info.RawEvent.BlockNumber, "txIndex", info.RawEvent.TxIndex)
	// }

	// Log
	headCount, tailCount := 0, 0
	for _, info := range CoinFlipInitiatedInfos {
		if info.Choice {
			headCount++
		} else {
			tailCount++
		}
	}

	log.Info("#### Summarizing CoinFlipInitiated events ####")
	log.Info("Total CoinFlipInitiated events recorded", "count", count)
	log.Info("CoinFlipInitiated event summary", "headCount", headCount, "tailCount", tailCount)

	return c.reqHashes2Info
}

func (c *CoinFlipInitiatedTracker) Record(e *types.Log) error {
	log.Trace("Processing CoinFlipInitiated event", "TxHash", e.TxHash.Hex())

	event := new(coin_flipper.CoinFlipperCoinFlipInitiated)
	parseEvent(c.iface, event, "CoinFlipInitiated", e)
	event.Raw = *e

	if err := c.Tracker.Record(e); err != nil {
		return err
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.reqHashes2Info[event.ReqHash] = &CoinFlipInitiatedInfo{
		Player:   event.Player,
		ConfigId: event.ConfigId,
		TxHash:   e.TxHash,
		ReqHash:  event.ReqHash,
		Choice:   event.Choice,
		NftId:    event.NftId,
		RawEvent: e,
	}

	return nil
}
