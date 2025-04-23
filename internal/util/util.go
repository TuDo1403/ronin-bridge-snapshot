package util

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"ronin-bridge-snapshot/generated/contract/erc20"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func BatchTokenAddr2Name(
	ctx context.Context,
	client *ethclient.Client,
	tokenAddrs []common.Address,
) map[common.Address]map[string]string {
	names := make(map[common.Address]map[string]string)
	for _, tokenAddr := range tokenAddrs {
		name, symbol, decimal := FetchERC20Metadata(ctx, client, tokenAddr)
		names[tokenAddr] = map[string]string{
			"name":     name,
			"symbol":   symbol,
			"decimals": fmt.Sprintf("%d", decimal),
		}
	}
	return names
}

func ToSingletonArray(addr common.Address) []common.Address {
	return []common.Address{addr}
}

func AggregateAddresses(addrs ...common.Address) []common.Address {
	// Remove duplicates
	addrSet := make(map[common.Address]struct{})
	for _, addr := range addrs {
		addrSet[addr] = struct{}{}
	}

	// Convert the set back to a slice
	result := make([]common.Address, 0, len(addrSet))
	for addr := range addrSet {
		result = append(result, addr)
	}
	return result
}

// AggregateTopics builds the [][]common.Hash “topics” matrix used in an
// ethereum.FilterQuery.  Each input slice represents the **OR‑set** for one
// topic position (topic0 … topic3).  A nil or empty slice means “match
// anything” in that position.
//
// The result:
//   - contains each position in order until the last **non‑nil** slot
//   - trailing nils are trimmed, so the matrix is a valid prefix per geth
//     rules (e.g. {{A},{}} becomes {{A}} because topic1 “anything” can be
//     omitted entirely).
//
// Special case: if every position is nil/empty the function returns nil, which
// tells the JSON‑RPC to match any topic list.
func AggregateTopics(
	topic0s, topic1s, topic2s, topic3s []common.Hash,
) [][]common.Hash {

	// Build a fixed‑length slice first.
	raw := [][]common.Hash{
		topic0s,
		topic1s,
		topic2s,
		topic3s,
	}

	topics := make([][]common.Hash, 4)
	for i, set := range raw {
		if len(set) > 0 {
			// Copy to avoid aliasing accidental mutations by caller.
			cp := make([]common.Hash, len(set))
			copy(cp, set)
			topics[i] = cp
		} else {
			// nil slot → “match any” in this position
			topics[i] = nil
		}
	}

	// Trim trailing nils so the slice is only as long as the last constraint.
	for i := len(topics) - 1; i >= 0; i-- {
		if topics[i] != nil {
			topics = topics[:i+1]
			break
		}
		if i == 0 { // all nil
			return nil
		}
	}

	return topics
}

func ToMap(txHashes []common.Hash) map[common.Hash]bool {
	m := make(map[common.Hash]bool)
	for _, txHash := range txHashes {
		m[txHash] = true
	}
	return m
}

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

func FetchTokenBalance(
	ctx context.Context,
	client *ethclient.Client,
	tokenAddr common.Address,
	accountAddr common.Address,
) *big.Int {
	if (tokenAddr == common.Address{} || IsWETH(tokenAddr)) {
		bal, err := client.BalanceAt(ctx, accountAddr, nil)
		if err != nil {
			log.Fatal("failed to fetch ETH balance", "error", err)
		}
		return bal
	}

	token, err := erc20.NewErc20Caller(tokenAddr, client)
	if err != nil {
		log.Fatal("failed to create token contract", "error", err)
	}

	balance, err := token.BalanceOf(nil, accountAddr)
	if err != nil {
		log.Fatal("failed to fetch token balance", "error", err)
	}

	return balance
}

func IsWETH(tokenAddr common.Address) bool {
	return tokenAddr == common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
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
