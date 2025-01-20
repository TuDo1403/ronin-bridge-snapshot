package util

import (
	"context"
	"log"
	"math"
	"math/big"
	"ronin-bridge-snapshot/internal/abi/erc20"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func NewClients(rpcURLs []string) []*ethclient.Client {
	clients := make([]*ethclient.Client, 0)
	for _, url := range rpcURLs {
		client, err := ethclient.Dial(url)
		if err != nil {
			log.Fatal("failed to create client", "url", url, "error", err)
		}
		clients = append(clients, client)
	}
	return clients
}

func BigIntToFloat(amount *big.Int, decimals int) float64 {
	// Convert the big.Int to a big.Float
	fAmount := new(big.Float).SetInt(amount)

	// Create a big.Float = 10^decimals
	divisor := new(big.Float).SetFloat64(math.Pow10(decimals))

	// Divide the amount by 10^decimals
	result := new(big.Float).Quo(fAmount, divisor)

	// Finally, convert big.Float -> float64 (possible precision loss!)
	f, _ := result.Float64()
	return f
}

func FetchERC20Metadata(ctx context.Context, client *ethclient.Client, tokenAddr common.Address) (name string, symbol string, decimals int) {
	if (tokenAddr == common.Address{}) {
		return "ETH", "ETH", 18
	}

	token, err := erc20.NewErc20Caller(tokenAddr, client)
	if err != nil {
		log.Fatal("failed to create token contract", "error", err)
	}

	name, err = token.Name(nil)
	if err != nil {
		log.Fatal("failed to fetch token name", "error", err)
	}

	symbol, err = token.Symbol(nil)
	if err != nil {
		log.Fatal("failed to fetch token symbol", "error", err)
	}

	raw, err := token.Decimals(nil)
	if err != nil {
		log.Fatal("failed to fetch token decimals", "error", err)
	}
	decimals = int(raw)

	return name, symbol, decimals
}
