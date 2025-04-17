package config

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

type NetworkConfig struct {
	RpcEndpoints    []string         `yaml:"rpc-endpoints"`
	Gateway         common.Address   `yaml:"gateway"`
	StartBlock      int              `yaml:"start-block"`
	EndBlock        int              `yaml:"end-block"`
	QueryBatchSize  int              `yaml:"query-batch-size"`
	Tokens          []common.Address `yaml:"tokens"`
	ExcludeTxHashes []common.Hash    `yaml:"exclude-tx-hashes"`
}

type Config struct {
	NetworkConfig map[string]NetworkConfig `yaml:"network-config"`
	Ronin         string                   `yaml:"ronin"`
	Mainchain     string                   `yaml:"mainchain"`
	PollInterval  int                      `yaml:"poll-interval"`
	LogLevel      string                   `yaml:"log-level"`
}

type AppConfig struct {
	Config
	Ctx          context.Context
	Cancel       context.CancelFunc
	PollInterval time.Duration
	LogLevel     slog.Level
	Ronin        NetworkConfig
	Mainchain    NetworkConfig
	Clients      map[string][]*ethclient.Client
}

func (a *AppConfig) String() string {
	return fmt.Sprintf("LogLevel=%s, PollInterval=%dms, RoninRpcs=%v, MainchainRpcs=%v", a.LogLevel, a.PollInterval, len(a.Config.NetworkConfig[a.Config.Ronin].RpcEndpoints), len(a.Config.NetworkConfig[a.Config.Mainchain].RpcEndpoints))
}

func NewAppConfig(filePath string) *AppConfig {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Warn("Failed to load .env file", "err", err)
	}

	raw, err := os.ReadFile(filePath)
	if err != nil {
		log.Crit("Read file Failed", err)
	}

	// Replace environment variables in raw config
	raw = []byte(os.ExpandEnv(string(raw)))

	var cfg Config
	err = yaml.Unmarshal(raw, &cfg)
	if err != nil {
		log.Crit("Unmarshal Failed", "Err", err)
	}

	var appCfg AppConfig
	appCfg.Config = cfg
	appCfg.Ctx, appCfg.Cancel = context.WithCancel(context.Background())

	if cfg.LogLevel == "debug" {
		appCfg.LogLevel = slog.LevelDebug
	} else if cfg.LogLevel == "info" {
		appCfg.LogLevel = slog.LevelInfo
	} else if cfg.LogLevel == "warn" {
		appCfg.LogLevel = slog.LevelWarn
	} else if cfg.LogLevel == "error" {
		appCfg.LogLevel = slog.LevelError
	} else {
		log.Crit("Invalid log level", "logLevel", cfg.LogLevel)
	}

	appCfg.Ronin = cfg.NetworkConfig[cfg.Ronin]
	appCfg.Mainchain = cfg.NetworkConfig[cfg.Mainchain]

	if appCfg.Ronin.RpcEndpoints == nil || len(appCfg.Ronin.RpcEndpoints) == 0 {
		log.Crit("Ronin RPC endpoints are not set or empty")
	}
	if appCfg.Mainchain.RpcEndpoints == nil || len(appCfg.Mainchain.RpcEndpoints) == 0 {
		log.Crit("Mainchain RPC endpoints are not set or empty")
	}

	appCfg.Clients = make(map[string][]*ethclient.Client)
	for _, rpc := range appCfg.Ronin.RpcEndpoints {
		client, err := ethclient.Dial(rpc)
		if err != nil {
			log.Crit("Failed to create Ronin client", "rpc", rpc, "err", err)
		}
		appCfg.Clients[cfg.Ronin] = append(appCfg.Clients[cfg.Ronin], client)
	}
	for _, rpc := range appCfg.Mainchain.RpcEndpoints {
		client, err := ethclient.Dial(rpc)
		if err != nil {
			log.Crit("Failed to create Mainchain client", "rpc", rpc, "err", err)
		}
		appCfg.Clients[cfg.Mainchain] = append(appCfg.Clients[cfg.Mainchain], client)
	}

	appCfg.PollInterval = time.Duration(cfg.PollInterval) * time.Millisecond
	if appCfg.PollInterval <= 0 {
		log.Crit("Poll interval must be greater than 0", "pollInterval", cfg.PollInterval)
	}

	return &appCfg
}
