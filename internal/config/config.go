package config

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"gopkg.in/yaml.v2"
)

type NetworkConfig struct {
	ChainId          uint64           `yaml:"chain-id"`
	RpcEndpoints     []string         `yaml:"rpc-endpoints"`
	CompanionNetwork string           `yaml:"companion-network"`
	Gateway          common.Address   `yaml:"gateway"`
	StartBlock       uint64           `yaml:"start-block"`
	EndBlock         uint64           `yaml:"end-block"`
	QueryBatchSize   int              `yaml:"query-batch-size"`
	Tokens           []common.Address `yaml:"tokens"`
	ExcludeTxHashes  []common.Hash    `yaml:"exclude-tx-hashes"`
}

type Config struct {
	NetworkConfig map[string]NetworkConfig `yaml:"network-config"`
	Ronin         string                   `yaml:"ronin"`
	Mainchain     string                   `yaml:"mainchain"`
}

func LoadConfigFromFile(filePath string) *Config {
	raw, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Read file Failed", err)
	}

	var cfg Config
	err = yaml.Unmarshal(raw, &cfg)
	if err != nil {
		log.Fatal("Unmarshal Failed", "Err", err)
	}

	return &cfg
}
