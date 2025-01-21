package util

import (
	"context"
	"math/big"
	"ronin-bridge-snapshot/internal/abi/erc20"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

func FilterERC20TransferForTokens(ctx context.Context, done <-chan struct{}, client *ethclient.Client, tokenAddrs []common.Address, toAddr common.Address, startBlock, endBlock uint64, rawLogChan chan []types.Log) {
	if len(tokenAddrs) == 0 {
		log.Error("no token addresses provided")
		return
	}

	erc20ABI, err := erc20.Erc20MetaData.GetAbi() // or however you load your ABI
	if err != nil {
		log.Error("failed to parse ABI", err)
		return
	}

	startTime := time.Now()

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(startBlock)),
		ToBlock:   big.NewInt(int64(endBlock)),
		Addresses: tokenAddrs,
		Topics: [][]common.Hash{
			{erc20ABI.Events["Transfer"].ID},
			{},
			{common.BytesToHash(toAddr.Bytes())},
		},
	}

	rawLogs, err := client.FilterLogs(ctx, query)
	if err != nil {
		log.Error("failed to filter logs for query", query, err)
	}

	responseTime := time.Since(startTime)
	if responseTime > 10*time.Second {
		log.Info("Low Response Time", "From", startBlock, "To", endBlock, "Time", responseTime)
		time.Sleep(60 * time.Second)
	}

	log.Debug("Raw Filter", "Count", len(rawLogs), "From", startBlock, "To", endBlock, "Size", endBlock-startBlock, "Time", responseTime)
	rawLogChan <- rawLogs
}

func SanitizeTransferLogs(ctx context.Context, done <-chan struct{}, toAddr common.Address, rawLogChan chan []types.Log, sanitizedLogChan chan map[common.Address][]*erc20.Erc20Transfer, mintTxHashesChan chan map[common.Address][]common.Hash, sanitizedTxHashesChan chan map[common.Address][]common.Hash) {
	// Load ABI once to avoid unnecessary recomputation
	erc20ABI, err := erc20.Erc20MetaData.GetAbi()
	if err != nil {
		log.Error("failed to parse ABI: %v", err)
		return
	}

	for {
		select {
		case <-ctx.Done():
			log.Trace("Sanitizer routine stopped.")
			return

		case <-done:
			log.Trace("Sanitizer routine received done signal. Exiting.")
			return

		case rawLogs := <-rawLogChan:
			sanitizedLogs := make(map[common.Address][]*erc20.Erc20Transfer)
			sanitizedTxHashes := make(map[common.Address][]common.Hash)
			mintTxHashes := make(map[common.Address][]common.Hash)

			for _, vLog := range rawLogs {
				tokenAddr := vLog.Address

				transferEvent := new(erc20.Erc20Transfer)
				err := erc20ABI.UnpackIntoInterface(transferEvent, "Transfer", vLog.Data)
				if err != nil {
					log.Error("failed to unpack transfer log", err)
					continue
				}

				transferEvent.From = common.BytesToAddress(vLog.Topics[1].Bytes())
				transferEvent.To = common.BytesToAddress(vLog.Topics[2].Bytes())

				// If "From" address is zero, it's a mint transaction
				if transferEvent.From == (common.Address{}) {
					log.Trace("Mint Transaction", "To", transferEvent.To, "TxHash", vLog.TxHash)
					mintTxHashes[tokenAddr] = append(mintTxHashes[tokenAddr], vLog.TxHash)
					continue
				}

				// If it's not sent to `toAddr`, ignore it
				if transferEvent.To != toAddr {
					log.Trace("Not Sent To Gateway", "To", transferEvent.To, "Gateway", toAddr, "TxHash", vLog.TxHash)
					continue
				}

				// Skip zero-value transfers
				if transferEvent.Value.Cmp(big.NewInt(0)) == 0 {
					log.Trace("Null Value Transfer", "From", transferEvent.From, "To", transferEvent.To, "TxHash", vLog.TxHash)
					continue
				}

				if vLog.Removed {
					log.Warn("Log Removed", "TxHash", vLog.TxHash)
					continue
				}

				if vLog.TxHash == common.HexToHash("0xa53c7c4402b40a63e485a1a8798d7b7ac6a151dce19c3486d9928575476ac3ce") ||
					vLog.TxHash == common.HexToHash("0x6af887fda63d7ec5538c43fdd3b1fa87ce5e2aaf824c5332f3bd8330d45f4c0b") ||
					vLog.TxHash == common.HexToHash("0x2392f70de576a6e90d95eb4686561ab6232fda7d972ca56fa94df20840f289c9") ||
					vLog.TxHash == common.HexToHash("0xc27bbc806abf466645a24f2be8bb7f082cba007f52767739eee0815153ae6fea") ||
					vLog.TxHash == common.HexToHash("0x78c3e08640228ff072749632f407617695a731f2dd377fa3b1a87f7697282897") ||
					vLog.TxHash == common.HexToHash("0x542c6ddaf6d3752d488153b212346b94bbe30d8b25eee94c80051baa4a8714b7") ||
					vLog.TxHash == common.HexToHash("0x2c4db38cafde1d399e07a754b6605bc5dd242119929dfd8382dfb9925d322379") ||
					vLog.TxHash == common.HexToHash("0x534c69d4c943b20f180774b942147e446e39e12920057e9795927defa547e2ab") ||
					vLog.TxHash == common.HexToHash("0xced4cf252ff100eb7cc96859fc0e85fcad97d97d26f96ed4ae78665301e09094") ||
					vLog.TxHash == common.HexToHash("0x28211787395779c94fca595000f3d079e86ebe12b92e4f005fad4dae5c44d2e9") ||
					vLog.TxHash == common.HexToHash("0x4887f384ba119800b85cb65e146bb559972cc15fab65e5d8ccafd32999b419c9") ||
					vLog.TxHash == common.HexToHash("0x1357e4fa4b955249724d437f634105195f1337edc2db8905a72e69b35ccbd870") {

					log.Warn("Hack Txs", "TxHash", vLog.TxHash.String(), "From", vLog.Address, "Token", tokenAddr, "Amount", transferEvent.Value)
					continue
				}

				log.Trace("Transfer Event", "From", transferEvent.From, "To", transferEvent.To, "Value", transferEvent.Value, "TxHash", vLog.TxHash)
				// Attach original log data
				transferEvent.Raw = vLog

				// Push to sanitized logs
				sanitizedLogs[tokenAddr] = append(sanitizedLogs[tokenAddr], transferEvent)
				sanitizedTxHashes[tokenAddr] = append(sanitizedTxHashes[tokenAddr], vLog.TxHash)
			}

			// Send sanitized data to respective channels
			sanitizedLogChan <- sanitizedLogs
			mintTxHashesChan <- mintTxHashes
			sanitizedTxHashesChan <- sanitizedTxHashes

			log.Trace("Sanitized Logs", "Count", len(sanitizedLogs))
		}
	}
}
