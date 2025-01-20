package main

import (
	"context"
	"os"
	"ronin-bridge-snapshot/cmd/mainchain"
	//"ronin-bridge-snapshot/cmd/ronin"
	"ronin-bridge-snapshot/internal/config"
	"ronin-bridge-snapshot/internal/util"
	"sync"
	// "time"

	// "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

func main() {
	ctx := context.Background()

	// Logging setup
	logger := log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stdout, log.LevelTrace, true))
	log.SetDefault(logger)

	// Load configuration
	cfg := config.LoadConfigFromFile("chain-config.yml")
	mainchainConfig := cfg.NetworkConfig[cfg.Mainchain]

	clients := util.NewClients(mainchainConfig.RpcEndpoints)

	// Create trackers
	trackers := util.NewTrackers(mainchainConfig.Tokens)

	var wg sync.WaitGroup
	done := make(chan struct{})

	// Run record
	mainchain.RecordReleasedOnMainchain(ctx, &wg, done, trackers, clients, mainchainConfig)
}

// func main() {
// 	ctx := context.Background()

// 	// Logging setup
// 	logger := log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stdout, log.LevelDebug, true))
// 	log.SetDefault(logger)

// 	// Load configuration
// 	cfg := config.LoadConfigFromFile("chain-config.yml")
// 	roninConfig := cfg.NetworkConfig[cfg.Ronin]

// 	clients := util.NewClients(roninConfig.RpcEndpoints)

// 	// Create trackers
// 	trackers := util.NewTrackers(roninConfig.Tokens)

// 	var wg sync.WaitGroup
// 	done := make(chan struct{})

// 	// Run record
// 	// ronin.RecordWithdrawalsOnRonin(ctx, &wg, done, trackers, clients[0], roninConfig)
// 	ronin.RecordRequestWithdrawalsOnRonin(ctx, &wg, done, trackers, clients, roninConfig)

// }
