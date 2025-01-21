package erc20_transfer

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

func FilterERC20TransferForTokens(ctx context.Context, client *ethclient.Client, tokenAddrs []common.Address, toAddr common.Address, startBlock, endBlock uint64, rawLogChan chan []types.Log) {
	if len(tokenAddrs) == 0 {
		log.Crit("No token addresses provided")
	}

	erc20ABI, err := erc20.Erc20MetaData.GetAbi() // or however you load your ABI
	if err != nil {
		log.Crit("Failed to parse ABI", err)
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

type Transfer struct {
	TxHash      common.Hash
	From        common.Address
	To          common.Address
	Value       *big.Int
	BlockNumber uint64
}

func SanitizeTransferLogs(
	rawLogs []types.Log,
	excludeTxHashes map[common.Hash]bool,
	toAddr common.Address,
	erc20TransferChan chan map[common.Address][]*Transfer,
) {
	// Load ABI once to avoid unnecessary recomputation
	erc20ABI, err := erc20.Erc20MetaData.GetAbi()
	if err != nil {
		log.Error("failed to parse ABI: %v", err)
		return
	}

	sanitizedTransferMap := make(map[common.Address][]*Transfer)

	for _, vLog := range rawLogs {
		tokenAddr := vLog.Address

		transferEvent := new(erc20.Erc20Transfer)
		err := erc20ABI.UnpackIntoInterface(transferEvent, "Transfer", vLog.Data)

		transferEvent.From = common.BytesToAddress(vLog.Topics[1].Bytes())
		transferEvent.To = common.BytesToAddress(vLog.Topics[2].Bytes())

		if err != nil {
			log.Crit("failed to unpack transfer log", err)
		}
		if vLog.Removed {
			log.Error("Log Removed", "TxHash", vLog.TxHash)
			continue
		}
		if transferEvent.From == (common.Address{}) {
			log.Warn("Mint Transaction", "Token", tokenAddr, "To", transferEvent.To, "TxHash", vLog.TxHash.String())
		}
		if transferEvent.Value.Int64() == 0 {
			log.Trace("Null Value Transfer", "Token", tokenAddr, "From", transferEvent.From, "To", transferEvent.To, "TxHash", vLog.TxHash.String())
			continue
		}
		if excludeTxHashes[vLog.TxHash] {
			log.Trace("Excluded TxHash", "TxHash", vLog.TxHash)
			continue
		}

		log.Trace("Transfer Event", "From", transferEvent.From, "To", transferEvent.To, "Value", transferEvent.Value, "TxHash", vLog.TxHash.String())
		// Attach original log data
		transferEvent.Raw = vLog

		sanitizedTransferMap[tokenAddr] = append(sanitizedTransferMap[tokenAddr], &Transfer{
			TxHash:      vLog.TxHash,
			From:        transferEvent.From,
			To:          transferEvent.To,
			Value:       transferEvent.Value,
			BlockNumber: vLog.BlockNumber,
		})
	}

	erc20TransferChan <- sanitizedTransferMap
}
