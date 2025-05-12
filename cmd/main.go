package main

import (
	"encoding/csv"
	"fmt"
	"math/big"
	"os"
	"ronin-bridge-snapshot/generated/contract/coin_flipper"
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

	coinFlipResolvedTracker := event_tracker.NewCoinFlipResolvedTracker(cfg.Ctx, 10, nil)
	coinFlipResolved := event_handler.NewMatcher(
		util.ToSingletonArray(cfg.Ronin.Gateway),
		coinFlipResolvedTracker.GetAbi().Events["CoinFlipResolved"].ID,
		nil,
		nil, nil,
	)
	coinFlipResolvedTracker.SetInCh(coinFlipResolved.ReceiveOnlyCh())

	coinFlipInitiatedTracker := event_tracker.NewCoinFlipInitiatedTracker(cfg.Ctx, 10, nil)
	coinFlipInitiated := event_handler.NewMatcher(
		util.ToSingletonArray(cfg.Ronin.Gateway),
		coinFlipInitiatedTracker.GetAbi().Events["CoinFlipInitiated"].ID,
		nil,
		nil, nil,
	)
	coinFlipInitiatedTracker.SetInCh(coinFlipInitiated.ReceiveOnlyCh())

	// withdrawalTracker := event_tracker.NewWithdrewTracker(cfg.Ctx, 10, cfg.Mainchain.ExcludeTxHashes, nil)
	// withdrewMatcher := event_handler.NewMatcher(
	// 	util.ToSingletonArray(cfg.Mainchain.Gateway),
	// 	withdrawalTracker.GetAbi().Events["Withdrew"].ID,
	// 	nil,
	// 	nil, nil,
	// )
	// withdrawalTracker.SetInCh(withdrewMatcher.ReceiveOnlyCh())

	roninFilterer := event_filterer.NewEventFilterer(
		cfg.Ctx,
		cfg.PollInterval,
		"[ronin] fetching events",
		cfg.Clients[cfg.Config.Ronin],
		util.ToSingletonArray(cfg.Ronin.Gateway),
		util.AggregateTopics([]common.Hash{coinFlipResolved.Topic0(), coinFlipInitiated.Topic0()}, nil, nil, nil),
		cfg.Ronin.QueryBatchSize,
		20,
		cfg.Ronin.StartBlock,
		cfg.Ronin.EndBlock,
	)
	// mainchainFilterer := event_filterer.NewEventFilterer(
	// 	cfg.Ctx,
	// 	cfg.PollInterval,
	// 	"[mainchain] fetching events",
	// 	cfg.Clients[cfg.Config.Mainchain],
	// 	util.ToSingletonArray(cfg.Mainchain.Gateway),
	// 	util.AggregateTopics([]common.Hash{
	// 		withdrewMatcher.Topic0()}, nil, nil, nil),
	// 	cfg.Mainchain.QueryBatchSize,
	// 	20,
	// 	cfg.Mainchain.StartBlock,
	// 	cfg.Mainchain.EndBlock,
	// )

	eventHandler := event_handler.NewEventHandler(cfg.Ctx, nil, 50)
	eventHandler.AddMatcher(coinFlipResolved)
	eventHandler.AddMatcher(coinFlipInitiated)
	// eventHandler.AddMatcher(withdrewMatcher)
	eventHandler.AddInCh(roninFilterer.ReceiveOnlyCh())
	// eventHandler.AddInCh(mainchainFilterer.ReceiveOnlyCh())

	roninDoneCh := roninFilterer.DoneCh()
	// mainchainDoneCh := mainchainFilterer.DoneCh()

	// Start all components
	coinFlipResolvedTracker.Start()
	coinFlipInitiatedTracker.Start()
	// withdrawalTracker.Start()
	eventHandler.Start()
	roninFilterer.Start()
	// mainchainFilterer.Start()

	<-roninDoneCh
	// <-mainchainDoneCh

	log.Info("Ronin/Mainchain filterer done")

	// Wait for a signal to stop
	// mainchainFilterer.Stop()
	roninFilterer.Stop()
	eventHandler.Stop()
	// withdrawalTracker.Stop()
	coinFlipResolvedTracker.Stop()

	// Summarize the trackers
	resolveInfo := coinFlipResolvedTracker.Summarize()
	initiateInfo := coinFlipInitiatedTracker.Summarize()

	headWinCount, tailWinCount := 0, 0
	headOutcomeCount, tailOutcomeCount := 0, 0

	for _, info := range resolveInfo {
		choice := initiateInfo[info.ReqHash].Choice // true = head, false = tail
		outcome := info.PlayerWin
		if choice && outcome {
			headWinCount++
		}
		if !choice && !outcome {
			tailWinCount++
		}

		if choice {
			if outcome {
				headOutcomeCount++
			} else {
				tailOutcomeCount++
			}
		} else {
			if outcome {
				tailOutcomeCount++
			} else {
				headOutcomeCount++
			}
		}

	}
	log.Info("Logging win results")
	for _, info := range resolveInfo {
		choice := initiateInfo[info.ReqHash].Choice
		outcome := info.PlayerWin

		if choice {
			if outcome {
				log.Info("Outcome", "ReqHash", info.ReqHash.Hex(), "outcome", true)
			} else {
				log.Info("Outcome", "ReqHash", info.ReqHash.Hex(), "outcome", false)
			}
		} else {
			if outcome {
				log.Info("Outcome", "ReqHash", info.ReqHash.Hex(), "outcome", false)
			} else {
				log.Info("Outcome", "ReqHash", info.ReqHash.Hex(), "outcome", true)
			}
		}
	}

	log.Info("#### Summarizing CoinFlip events ####")
	log.Info("Total CoinFlipResolved events recorded", "count", len(resolveInfo))
	log.Info("Total CoinFlipInitiated events recorded", "count", len(initiateInfo))

	log.Info("CoinFlipResolved event summary", "headWinCount", headWinCount, "tailWinCount", tailWinCount)
	log.Info("CoinFlipResolved event summary", "headOutcomeCount", headOutcomeCount, "tailOutcomeCount", tailOutcomeCount)

	// mainchainERC20s, _, _, _ := withdrawalTracker.Summarize()
	// roninERC20s, _, _, roninNullERC20s, _ := coinFlipResolvedTracker.Summarize()

	// generateReport(cfg, roninERC20s, roninNullERC20s, mainchainERC20s)

	coinFlipContract, err := coin_flipper.NewCoinFlipperCaller(cfg.Ronin.Gateway, cfg.Clients[cfg.Config.Ronin][0])
	if err != nil {
		log.Crit("Failed to create CoinFlipper contract", "error", err)
	}

	configCount := 20
	configId2Amount := make(map[uint16]*big.Int)
	configId2Token := make(map[uint16]common.Address)
	for i := range configCount {
		res, err := coinFlipContract.CoinFlipConfigs(nil, uint16(i))
		if err != nil {
			log.Crit("Failed to get CoinFlipConfig", "error", err)
		}
		configId2Amount[uint16(i)] = res.TokenAmount
		configId2Token[uint16(i)] = res.TokenAddress
	}

	// targetAddress := common.HexToAddress("0x0E89Ee674a4929aDEBda0e54905Af9ca7827f13C")
	// targetBetSum := new(big.Int)
	// targetWinSum := new(big.Int)

	// for _, info := range resolveInfo {
	// 	if info.Player == targetAddress {

	// 		configId, ok := new(big.Int).SetString(info.ConfigId.String(), 10)
	// 		if !ok {
	// 			log.Crit("Failed to parse ConfigId", "ConfigId", info.ConfigId.String())
	// 		}
	// 		betAmount := configId2Amount[uint16(configId.Uint64())]
	// 		if betAmount == nil {
	// 			log.Crit("Failed to get bet amount", "ConfigId", configId.String())
	// 		}
	// 		if configId2Token[uint16(configId.Uint64())] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
	// 			continue
	// 		}

	// 		targetBetSum = new(big.Int).Add(targetBetSum, betAmount)
	// 		bankRollBlock := info.RawEvent.BlockNumber
	// 		balanceAtBlock, err := cfg.Clients[cfg.Config.Ronin][1].BalanceAt(cfg.Ctx, targetAddress, big.NewInt(int64(bankRollBlock)))
	// 		balanceAtBlock = new(big.Int).Div(balanceAtBlock, big.NewInt(1000000000000000000))
	// 		if err != nil {
	// 			log.Crit("Failed to get balance at block", "error", err)
	// 		}
	// 		if info.PlayerWin {
	// 			targetWinSum = new(big.Int).Add(targetWinSum, new(big.Int).Mul(betAmount, big.NewInt(2)))
	// 		}

	// 		log.Info("Player action", "Choice", initiateInfo[info.ReqHash].Choice, "BetAmount", new(big.Int).Div(betAmount, big.NewInt(1000000000000000000)).String(), "Won", info.PlayerWin, "TokenAddress", configId2Token[uint16(configId.Uint64())].Hex(), "TxHash", info.TxHash.Hex(), "ReqHash", info.ReqHash.Hex(), "BlockNumber", info.RawEvent.BlockNumber, "BankRoll", balanceAtBlock)
	// 	}
	// }
	// log.Info("Target address", "BetSum", new(big.Int).Div(targetBetSum, big.NewInt(1000000000000000000)).String(), "WinSum", new(big.Int).Div(targetWinSum, big.NewInt(1000000000000000000)).String())

	file, err := os.Create("coin_flip_results.csv")
	if err != nil {
		log.Crit("Failed to create CSV file", "error", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	if err := writer.Write([]string{"choice", "bet_amount", "won", "vrfChoice"}); err != nil {
		log.Crit("Failed to write CSV header", "error", err)
	}

	for _, info := range resolveInfo {
		configId, ok := new(big.Int).SetString(info.ConfigId.String(), 10)
		if !ok {
			log.Crit("Failed to parse ConfigId", "ConfigId", info.ConfigId.String())
		}
		// if configId2Token[uint16(configId.Uint64())] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
		// 	continue
		// }

		vrfChoice := false // true = head, false = tail

		if initiateInfo[info.ReqHash].Choice { // true = head, false = tail
			if info.PlayerWin {
				vrfChoice = true
			}
		} else { // false = tail
			if !info.PlayerWin {
				vrfChoice = true
			}
		}

		record := []string{
			fmt.Sprintf("%t", initiateInfo[info.ReqHash].Choice),
			new(big.Int).Div(configId2Amount[uint16(configId.Uint64())], big.NewInt(1000000000000000000)).String(),
			fmt.Sprintf("%t", info.PlayerWin),
			fmt.Sprintf("%t", vrfChoice),
		}

		if err := writer.Write(record); err != nil {
			log.Crit("Failed to write CSV record", "error", err)
		}
	}
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

		vDecimals, _ := strconv.Atoi(decimals)

		log.Info("Pending withdrawal", "addr", tokenAddr.Hex(), "metadata", fmt.Sprintf("%s (%s) - %v", name, symbol, decimals), "wei", amount.String(), "eth", util.BigIntToFloat(amount, vDecimals))
	}
}
