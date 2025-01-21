package withdrawal

import (
	"context"
	"fmt"
	"math/big"
	"ronin-bridge-snapshot/internal/abi/mainchain_gateway"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Adapter for MainchainGatewayWithdrew event
type MainchainGatewayWithdrewAdapter struct {
	Event *mainchain_gateway.MainchainGatewayWithdrew
}

func (m *MainchainGatewayWithdrewAdapter) GetTxBlockNumber() uint64 {
	return m.Event.Raw.BlockNumber
}
func (m *MainchainGatewayWithdrewAdapter) GetTxHash() common.Hash {
	return m.Event.Raw.TxHash
}
func (m *MainchainGatewayWithdrewAdapter) GetReceiptHash() common.Hash {
	return m.Event.ReceiptHash
}
func (m *MainchainGatewayWithdrewAdapter) GetReceiptId() int64 {
	return m.Event.Receipt.Id.Int64()
}
func (m *MainchainGatewayWithdrewAdapter) GetQuantity() *big.Int {
	return m.Event.Receipt.Info.Quantity
}
func (m *MainchainGatewayWithdrewAdapter) GetTokenAddr() common.Address {
	return m.Event.Receipt.Mainchain.TokenAddr
}
func (m *MainchainGatewayWithdrewAdapter) IsRemoved() bool {
	return m.Event.Raw.Removed
}
func (m *MainchainGatewayWithdrewAdapter) IsValidKind() bool {
	return m.Event.Receipt.Kind == 1
}
func (m *MainchainGatewayWithdrewAdapter) IsValidERC20() bool {
	return m.Event.Receipt.Info.Erc == 0
}
func (m *MainchainGatewayWithdrewAdapter) GetKind() uint8 {
	return m.Event.Receipt.Kind
}
func (m *MainchainGatewayWithdrewAdapter) ValidateEvent(excludeTxHashes map[common.Hash]bool) (valid bool, reason string, err error) {

	if m.GetKind() != 1 {
		return false, "Invalid Withdrawal", fmt.Errorf("invalid Kind %d", m.GetKind())
	}
	if !m.IsValidERC20() {
		return false, "Non-ERC20 Withdrawal", nil
	}
	if m.GetQuantity().Int64() == 0 {
		return false, "Null Quantity Withdrawal", nil
	}
	if m.IsRemoved() {
		return false, "Log Removed", nil
	}

	if excludeTxHashes[m.GetTxHash()] {
		return false, "Excluded Tx Detected", nil
	}

	return true, "", nil
}

func FilterWithdrew(
	ctx context.Context,
	clients []*ethclient.Client,
	gwAddr common.Address,
	startBlock, endBlock uint64,
	eventChan chan []*mainchain_gateway.MainchainGatewayWithdrew,
) {
	FilterEvents(
		ctx, clients, gwAddr, startBlock, endBlock,
		func(opts *bind.FilterOpts, client *ethclient.Client) (*mainchain_gateway.MainchainGatewayWithdrewIterator, error) {
			filterer, err := mainchain_gateway.NewMainchainGatewayFilterer(gwAddr, client)
			if err != nil {
				return nil, err
			}
			return filterer.FilterWithdrew(opts)
		},
		func(iterator *mainchain_gateway.MainchainGatewayWithdrewIterator) []*mainchain_gateway.MainchainGatewayWithdrew {
			var logs []*mainchain_gateway.MainchainGatewayWithdrew
			for iterator.Next() {
				logs = append(logs, iterator.Event)
			}
			return logs
		},
		eventChan,
	)
}

func SanitizeWithdrewEvents(
	events []*mainchain_gateway.MainchainGatewayWithdrew,
	excludeTxHashes map[common.Hash]bool,
	withdrawalReceiptChan chan map[common.Address][]*WithdrawalReceipt,
) {
	var adaptedEvents []EventProcessor
	for _, event := range events {
		adaptedEvents = append(adaptedEvents, &MainchainGatewayWithdrewAdapter{Event: event})
	}

	SanitizeEvents(
		adaptedEvents,
		excludeTxHashes,
		func(event EventProcessor) (bool, string, error) {
			return event.(*MainchainGatewayWithdrewAdapter).ValidateEvent(excludeTxHashes)
		},
		withdrawalReceiptChan,
	)
}
