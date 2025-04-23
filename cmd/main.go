package main

import (
	"fmt"
	"math/big"
	"os"
	"ronin-bridge-snapshot/generated/contract/mainchain_gateway_v3"
	"ronin-bridge-snapshot/generated/contract/ronin_gateway_v3"
	"ronin-bridge-snapshot/internal/config"
	"ronin-bridge-snapshot/internal/event_filterer"
	"ronin-bridge-snapshot/internal/event_handler"
	"ronin-bridge-snapshot/internal/event_tracker"
	"ronin-bridge-snapshot/internal/util"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

func main() {
	// Logging setup
	logger := log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stdout, log.LevelTrace, true))
	log.SetDefault(logger)

	cfg := config.NewAppConfig("chain-config.yml")
	log.Info("Config loaded", "Config", cfg.String())

	// Logging setup
	logger = log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stdout, cfg.LogLevel, true))
	log.SetDefault(logger)

	requestWithdrawalTracker := event_tracker.NewRequestWithdrawalTracker(cfg.Ctx, 10, cfg.Ronin.ExcludeTxHashes, nil)
	withdrawalRequested := event_handler.NewMatcher(
		util.ToSingletonArray(cfg.Ronin.Gateway),
		requestWithdrawalTracker.GetAbi().Events["WithdrawalRequested"].ID,
		nil,
		nil, nil,
	)
	requestWithdrawalTracker.SetInCh(withdrawalRequested.ReceiveOnlyCh())

	withdrawalTracker := event_tracker.NewWithdrewTracker(cfg.Ctx, 10, cfg.Mainchain.ExcludeTxHashes, nil)
	withdrewMatcher := event_handler.NewMatcher(
		util.ToSingletonArray(cfg.Mainchain.Gateway),
		withdrawalTracker.GetAbi().Events["Withdrew"].ID,
		nil,
		nil, nil,
	)
	withdrawalTracker.SetInCh(withdrewMatcher.ReceiveOnlyCh())

	roninFilterer := event_filterer.NewEventFilterer(
		cfg.Ctx,
		cfg.PollInterval,
		"[ronin] fetching events",
		cfg.Clients[cfg.Config.Ronin],
		util.ToSingletonArray(cfg.Ronin.Gateway),
		util.AggregateTopics([]common.Hash{
			withdrawalRequested.Topic0()}, nil, nil, nil),
		cfg.Ronin.QueryBatchSize,
		20,
		cfg.Ronin.StartBlock,
		cfg.Ronin.EndBlock,
	)
	mainchainFilterer := event_filterer.NewEventFilterer(
		cfg.Ctx,
		cfg.PollInterval,
		"[mainchain] fetching events",
		cfg.Clients[cfg.Config.Mainchain],
		util.ToSingletonArray(cfg.Mainchain.Gateway),
		util.AggregateTopics([]common.Hash{
			withdrewMatcher.Topic0()}, nil, nil, nil),
		cfg.Mainchain.QueryBatchSize,
		20,
		cfg.Mainchain.StartBlock,
		cfg.Mainchain.EndBlock,
	)

	eventHandler := event_handler.NewEventHandler(cfg.Ctx, nil, 50)
	eventHandler.AddMatcher(withdrawalRequested)
	eventHandler.AddMatcher(withdrewMatcher)
	eventHandler.AddInCh(roninFilterer.ReceiveOnlyCh())
	eventHandler.AddInCh(mainchainFilterer.ReceiveOnlyCh())

	roninDoneCh := roninFilterer.DoneCh()
	mainchainDoneCh := mainchainFilterer.DoneCh()

	// Start all components
	requestWithdrawalTracker.Start()
	withdrawalTracker.Start()
	eventHandler.Start()
	roninFilterer.Start()
	mainchainFilterer.Start()

	<-roninDoneCh
	<-mainchainDoneCh

	log.Info("Ronin/Mainchain filterer done")

	// Wait for a signal to stop
	mainchainFilterer.Stop()
	roninFilterer.Stop()
	eventHandler.Stop()
	withdrawalTracker.Stop()
	requestWithdrawalTracker.Stop()

	// Summarize the trackers
	mainchainERC20s, _, _, _ := withdrawalTracker.Summarize()
	roninERC20s, _, _, roninNullERC20s, _ := requestWithdrawalTracker.Summarize()

	generateReport(cfg, roninERC20s, roninNullERC20s, mainchainERC20s)
}

func generateReport(appCfg *config.AppConfig, roninERC20Receipts, roninNullERC20s map[common.Hash]*ronin_gateway_v3.TransferReceipt, mainchainERC20Receipts map[common.Hash]*mainchain_gateway_v3.TransferReceipt) {
	pendingWithdrawals := make(map[common.Hash]*ronin_gateway_v3.TransferReceipt, 0)
	pendingAmounts := make(map[common.Address]*big.Int)

	for receiptHash, receipt := range roninERC20Receipts {
		if _, ok := mainchainERC20Receipts[receiptHash]; !ok {
			pendingWithdrawals[receiptHash] = receipt
			if _, ok := pendingAmounts[receipt.Mainchain.TokenAddr]; !ok {
				pendingAmounts[receipt.Mainchain.TokenAddr] = new(big.Int)
			}

			pendingAmounts[receipt.Mainchain.TokenAddr].Add(pendingAmounts[receipt.Mainchain.TokenAddr], receipt.Info.Quantity)

			continue
		}
		if roninERC20Receipts[receiptHash].Mainchain.TokenAddr != mainchainERC20Receipts[receiptHash].Mainchain.TokenAddr {
			log.Crit("Token address mismatch", "roninReceiptHash", receiptHash.Hex(), "roninTxHash", receiptHash.Hex(), "mainchainTxHash", receiptHash.Hex())
		}
		if roninERC20Receipts[receiptHash].Ronin.TokenAddr != mainchainERC20Receipts[receiptHash].Ronin.TokenAddr {
			log.Crit("Token address mismatch", "roninReceiptHash", receiptHash.Hex(), "roninTxHash", receiptHash.Hex(), "mainchainTxHash", receiptHash.Hex())
		}
	}

	pendingWithdrawalsNull := make(map[common.Hash]*ronin_gateway_v3.TransferReceipt, 0)
	for receiptHash, receipt := range roninNullERC20s {
		if _, ok := mainchainERC20Receipts[receiptHash]; !ok {
			pendingWithdrawalsNull[receiptHash] = receipt
		}
	}

	tokenAddrMap := make([]common.Address, 0)
	for tokenAddr := range pendingAmounts {
		tokenAddrMap = append(tokenAddrMap, tokenAddr)
	}

	tokenInfoMap := util.BatchTokenAddr2Name(appCfg.Ctx, appCfg.Clients[appCfg.Config.Mainchain][0], tokenAddrMap)

	log.Info("Pending withdrawals", "count", len(pendingWithdrawals))
	log.Info("Pending withdrawals null", "count", len(pendingWithdrawalsNull))
	log.Info("Null ERC20s", "count", len(roninNullERC20s))

	for tokenAddr, amount := range pendingAmounts {
		info := tokenInfoMap[tokenAddr]
		name, nameOk := info["name"]
		symbol, symbolOk := info["symbol"]
		decimals, decimalsOk := info["decimals"]

		if !nameOk || !symbolOk || !decimalsOk {
			log.Warn("Missing token metadata", "addr", tokenAddr.Hex(), "metadata", info)
			continue
		}

		currBalance := util.FetchTokenBalance(appCfg.Ctx, appCfg.Clients[appCfg.Config.Mainchain][0], tokenAddr, appCfg.Mainchain.Gateway)
		migratableLiquidity := new(big.Int).Sub(currBalance, amount)

		vDecimals, _ := strconv.Atoi(decimals)

		log.Info("Pending withdrawal", "addr", tokenAddr.Hex(), "metadata", fmt.Sprintf("%s (%s) - %v", name, symbol, decimals), "wei", amount.String(), "eth", util.BigIntToFloat(amount, vDecimals))
		fmt.Println()
		log.Info("Current balance", "addr", tokenAddr.Hex(), "metadata", fmt.Sprintf("%s (%s) - %v", name, symbol, decimals), "wei", currBalance.String(), "eth", util.BigIntToFloat(currBalance, vDecimals))
		log.Info("Migratable liquidity", "addr", tokenAddr.Hex(), "metadata", fmt.Sprintf("%s (%s) - %v", name, symbol, decimals), "wei", migratableLiquidity, "eth", util.BigIntToFloat(migratableLiquidity, vDecimals))
		fmt.Println()
	}
}
