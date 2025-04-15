package withdrawal

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

type WithdrawalReceipt struct {
	TxHash      common.Hash
	ReceiptHash common.Hash
	ReceiptId   int64
	Quantity    *big.Int
	BlockNumber uint64
}

// EventProcessor interface (generic for both event types)
type EventProcessor interface {
	GetTxHash() common.Hash
	GetTxBlockNumber() uint64
	GetReceiptHash() common.Hash
	GetReceiptId() int64
	GetQuantity() *big.Int
	GetTokenAddr() common.Address
	GetKind() uint8
	IsRemoved() bool
	IsValidKind() bool
	IsValidERC20() bool
}

// Generic function to sanitize withdrawal events
func SanitizeEvents[T EventProcessor](
	events []T,
	excludeTxHashes map[common.Hash]bool,
	validateEvent func(EventProcessor) (bool, string, error),
	withdrawalReceipt chan map[common.Address][]*WithdrawalReceipt,
) {
	sanitizedReceiptMap := make(map[common.Address][]*WithdrawalReceipt)

	for _, event := range events {
		txHash := event.GetTxHash()
		tokenAddr := event.GetTokenAddr()

		valid, reason, err := validateEvent(event)
		if err != nil {
			log.Crit("Failed to validate event", err, "TxHash", txHash.String())
		}
		if !valid {
			log.Warn("Invalid Withdrawal", "TxHash", txHash.String(), "Reason", reason)
			continue
		}

		log.Trace("Processed Withdrawal Event:", "Token", tokenAddr.String(), "Amount", event.GetQuantity(), "TxHash", txHash.String())

		sanitizedReceiptMap[tokenAddr] = append(sanitizedReceiptMap[tokenAddr], &WithdrawalReceipt{
			TxHash:      txHash,
			ReceiptHash: event.GetReceiptHash(),
			ReceiptId:   event.GetReceiptId(),
			Quantity:    event.GetQuantity(),
			BlockNumber: event.GetTxBlockNumber(),
		})
	}

	withdrawalReceipt <- sanitizedReceiptMap
}

// Generic filtering function
func FilterEvents[T any, I any](
	ctx context.Context,
	clients []*ethclient.Client,
	gatewayAddr common.Address,
	startBlock, endBlock uint64,
	filterFunc func(*bind.FilterOpts, *ethclient.Client) (I, error),
	extractEvents func(I) []*T,
	eventChan chan []*T,
) {
	for i, client := range clients {
		iterator, err := filterFunc(&bind.FilterOpts{
			Start:   startBlock,
			End:     &endBlock,
			Context: ctx,
		}, client)
		if err != nil {
			if i == len(clients)-1 {
				log.Crit("Failed to filter logs", err)
			}
			log.Debug("Retrying", "From", startBlock, "To", endBlock, "Using Idx", i+1)
			continue
		}

		// Extract events from the iterator
		events := extractEvents(iterator)

		log.Debug("Filtered Events", "Count", len(events), "From", startBlock, "To", endBlock)
		eventChan <- events

		break
	}
}
