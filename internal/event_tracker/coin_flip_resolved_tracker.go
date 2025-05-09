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

type CoinFlipResolvedInfo struct {
	Player    common.Address
	ConfigId  *big.Int
	ReqHash   common.Hash
	TxHash    common.Hash
	PlayerWin bool
	RawEvent  *types.Log
}

type CoinFlipResolvedTracker struct {
	*Tracker
	reqHashes2Info map[common.Hash]*CoinFlipResolvedInfo
}

func NewCoinFlipResolvedTracker(ctx context.Context, nWorker int, in <-chan *types.Log) *CoinFlipResolvedTracker {
	c := &CoinFlipResolvedTracker{
		Tracker:        NewTracker(ctx, "CoinFlipResolved", nWorker, in, nil),
		reqHashes2Info: make(map[common.Hash]*CoinFlipResolvedInfo),
	}

	c.Tracker.iface, _ = coin_flipper.CoinFlipperMetaData.GetAbi()
	c.Tracker.callback = c.Record

	return c
}

func (c *CoinFlipResolvedTracker) Summarize() []*CoinFlipResolvedInfo {
	count := c.Total()

	if count == 0 {
		return nil
	}

	CoinFlipResolvedInfos := make([]*CoinFlipResolvedInfo, 0, len(c.reqHashes2Info))
	for _, info := range c.reqHashes2Info {
		CoinFlipResolvedInfos = append(CoinFlipResolvedInfos, info)
	}
	// sort by block number
	sort.Slice(CoinFlipResolvedInfos, func(i, j int) bool {
		return cmp(CoinFlipResolvedInfos[i].RawEvent, CoinFlipResolvedInfos[j].RawEvent)
	})

	// for _, info := range CoinFlipResolvedInfos {
	// 	log.Info("CoinFlipResolved event", "Player", info.Player.Hex(), "ConfigId", info.ConfigId, "ReqHash", info.ReqHash, "PlayerWin", info.PlayerWin, "TxHash", info.RawEvent.TxHash.Hex(), "blockNumber", info.RawEvent.BlockNumber, "txIndex", info.RawEvent.TxIndex)
	// }

	type PlayerStat struct {
		Player common.Address
		won    int
		count  int
	}
	playerStats := make(map[common.Address]*PlayerStat)
	winCount, loseCount := 0, 0

	for _, info := range CoinFlipResolvedInfos {
		stat, exists := playerStats[info.Player]
		if !exists {
			stat = &PlayerStat{Player: info.Player}
			playerStats[info.Player] = stat
		}
		stat.count++
		if info.PlayerWin {
			winCount++
			log.Info("Win Request Hash", "ReqHash", info.ReqHash.Hex(), "Player", info.Player.Hex())
			stat.won++
		} else {
			loseCount++
		}
	}

	playerWinCounts := make([]PlayerStat, 0, len(playerStats))
	for _, stat := range playerStats {
		playerWinCounts = append(playerWinCounts, *stat)
	}

	sort.Slice(playerWinCounts, func(i, j int) bool {
		return playerWinCounts[i].won > playerWinCounts[j].won
	})

	log.Info("#### Summarizing CoinFlipResolved events ####")
	log.Info("Total CoinFlipResolved events recorded", "count", count)

	// Print top 50 players
	for i := 0; i < len(playerWinCounts) && i < 50; i++ {
		log.Info("Top Player", "Player", playerWinCounts[i].Player.Hex(), "winCount", playerWinCounts[i].won, "totalCount", playerWinCounts[i].count, "winRate", float64(playerWinCounts[i].won)/float64(playerWinCounts[i].count)*100)
	}

	log.Info("Summary", "totalWins", winCount, "totalLosses", loseCount, "totalPlayers", len(playerStats), "winRate", float64(winCount)/float64(count)*100)

	// Sort by total played
	sort.Slice(playerWinCounts, func(i, j int) bool {
		return playerWinCounts[i].count > playerWinCounts[j].count
	})
	log.Info("Top 50 Players by Total Played")
	for i := 0; i < len(playerWinCounts) && i < 50; i++ {
		log.Info("Top Player", "Player", playerWinCounts[i].Player.Hex(), "totalCount", playerWinCounts[i].count, "winRate", float64(playerWinCounts[i].won)/float64(playerWinCounts[i].count)*100)
	}

	return CoinFlipResolvedInfos
}

func (c *CoinFlipResolvedTracker) Record(e *types.Log) error {
	log.Trace("Processing CoinFlipResolved event", "TxHash", e.TxHash.Hex())

	event := new(coin_flipper.CoinFlipperCoinFlipResolved)
	parseEvent(c.iface, event, "CoinFlipResolved", e)
	event.Raw = *e

	if err := c.Tracker.Record(e); err != nil {
		return err
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.reqHashes2Info[event.ReqHash] = &CoinFlipResolvedInfo{
		Player:    event.Player,
		ConfigId:  event.ConfigId,
		ReqHash:   event.ReqHash,
		TxHash:    e.TxHash,
		PlayerWin: event.PlayerWin,
		RawEvent:  e,
	}

	return nil
}
