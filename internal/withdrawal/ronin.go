package withdrawal

import (
	"context"
	"fmt"
	"math/big"
	"ronin-bridge-snapshot/internal/abi/ronin_gateway"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Adapter for RoninGatewayWithdrawalRequested event
type RoninGatewayWithdrawalRequestedAdapter struct {
	Event *ronin_gateway.RoninGatewayWithdrawalRequested
}

func (r *RoninGatewayWithdrawalRequestedAdapter) GetTxBlockNumber() uint64 {
	return r.Event.Raw.BlockNumber
}
func (r *RoninGatewayWithdrawalRequestedAdapter) GetTxHash() common.Hash {
	return r.Event.Raw.TxHash
}
func (r *RoninGatewayWithdrawalRequestedAdapter) GetReceiptHash() common.Hash {
	return r.Event.ReceiptHash
}
func (r *RoninGatewayWithdrawalRequestedAdapter) GetReceiptId() int64 {
	return r.Event.Arg1.Id.Int64()
}
func (r *RoninGatewayWithdrawalRequestedAdapter) GetQuantity() *big.Int {
	return r.Event.Arg1.Info.Quantity
}
func (r *RoninGatewayWithdrawalRequestedAdapter) GetTokenAddr() common.Address {
	return r.Event.Arg1.Ronin.TokenAddr
}
func (r *RoninGatewayWithdrawalRequestedAdapter) IsRemoved() bool {
	return r.Event.Raw.Removed
}
func (r *RoninGatewayWithdrawalRequestedAdapter) GetKind() uint8 {
	return r.Event.Arg1.Kind
}
func (r *RoninGatewayWithdrawalRequestedAdapter) IsValidKind() bool {
	return r.Event.Arg1.Kind == 1
}
func (r *RoninGatewayWithdrawalRequestedAdapter) IsValidERC20() bool {
	return r.Event.Arg1.Info.Erc == 0
}

func (r *RoninGatewayWithdrawalRequestedAdapter) ValidateEvent(excludeTxHashes map[common.Hash]bool) (valid bool, reason string, err error) {
	if r.GetKind() != 1 {
		return false, "Invalid Withdrawal", fmt.Errorf("invalid Kind %d", r.GetKind())
	}
	if !r.IsValidERC20() {
		return false, "Non-ERC20 Withdrawal", nil
	}
	if r.GetQuantity().Int64() == 0 {
		return false, "Null Value Withdrawal", nil
	}
	if r.IsRemoved() {
		return false, "Log Removed", nil
	}
	if excludeTxHashes[r.GetTxHash()] {
		return false, "Excluded Tx Detected", nil
	}

	return true, "", nil
}

func FilterRequestWithdraw(
	ctx context.Context,
	clients []*ethclient.Client,
	gwAddr common.Address,
	startBlock,
	endBlock uint64,
	eventChan chan []*ronin_gateway.RoninGatewayWithdrawalRequested,
) {
	FilterEvents(
		ctx, clients, gwAddr, startBlock, endBlock,
		func(opts *bind.FilterOpts, client *ethclient.Client) (*ronin_gateway.RoninGatewayWithdrawalRequestedIterator, error) {
			filterer, err := ronin_gateway.NewRoninGatewayFilterer(gwAddr, client)
			if err != nil {
				return nil, err
			}
			return filterer.FilterWithdrawalRequested(opts)
		},
		func(iterator *ronin_gateway.RoninGatewayWithdrawalRequestedIterator) []*ronin_gateway.RoninGatewayWithdrawalRequested {
			var logs []*ronin_gateway.RoninGatewayWithdrawalRequested
			for iterator.Next() {
				logs = append(logs, iterator.Event)
			}
			return logs
		},
		eventChan,
	)
}

func SanitizeRequestWithdrawalEvents(
	events []*ronin_gateway.RoninGatewayWithdrawalRequested,
	excludeTxHashes map[common.Hash]bool,
	withdrawalReceiptChan chan map[common.Address][]*WithdrawalReceipt,
) {
	var adaptedEvents []EventProcessor
	for _, event := range events {
		adaptedEvents = append(adaptedEvents, &RoninGatewayWithdrawalRequestedAdapter{Event: event})
	}

	SanitizeEvents(
		adaptedEvents,
		excludeTxHashes,
		func(event EventProcessor) (bool, string, error) {
			return event.(*RoninGatewayWithdrawalRequestedAdapter).ValidateEvent(excludeTxHashes)
		},
		withdrawalReceiptChan,
	)
}
