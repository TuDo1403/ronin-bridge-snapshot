package main

import (
	"os"
	"ronin-bridge-snapshot/internal/config"
	"ronin-bridge-snapshot/internal/event_filterer"
	"ronin-bridge-snapshot/internal/event_handler"
	"ronin-bridge-snapshot/internal/event_tracker"
	"ronin-bridge-snapshot/internal/util"

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
	withdrawalTracker.Summarize()
	requestWithdrawalTracker.Summarize()
}
