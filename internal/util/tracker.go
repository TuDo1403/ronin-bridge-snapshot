package util

import (
	"math/big"
	"ronin-bridge-snapshot/internal/abi/erc20"
	"ronin-bridge-snapshot/internal/abi/mainchain_gateway"
	"ronin-bridge-snapshot/internal/abi/ronin_gateway"

	"github.com/ethereum/go-ethereum/common"
)

type Tracker struct {
	TxCount           int
	AccAmount         *big.Int
	LastRecordedBlock int
	MintTxHashes      []common.Hash
	SanitizedTxHashes []common.Hash
	ReceiptHashes     []common.Hash
	ReceiptIds        []*big.Int
	Quantities        []*big.Int
}

func NewTracker() *Tracker {
	return &Tracker{
		TxCount:           0,
		AccAmount:         big.NewInt(0),
		LastRecordedBlock: 0,
		MintTxHashes:      make([]common.Hash, 0),
		SanitizedTxHashes: make([]common.Hash, 0),
	}
}

func NewTrackers(tokens []common.Address) map[common.Address]*Tracker {
	trackers := make(map[common.Address]*Tracker)
	for _, token := range tokens {
		trackers[token] = NewTracker()
	}
	return trackers
}

func RecordTransfer(tracker *Tracker, logs []*erc20.Erc20Transfer) {
	tracker.TxCount += len(logs)
	for _, log := range logs {
		tracker.AccAmount.Add(tracker.AccAmount, log.Value)
		if int(log.Raw.BlockNumber) > tracker.LastRecordedBlock {
			tracker.LastRecordedBlock = int(log.Raw.BlockNumber)
		}
	}
}

func RecordRequestWithdrawals(tracker *Tracker, logs []*ronin_gateway.RoninGatewayWithdrawalRequested) {
	tracker.TxCount += len(logs)
	for _, log := range logs {
		tracker.AccAmount.Add(tracker.AccAmount, log.Arg1.Info.Quantity)
		if int(log.Raw.BlockNumber) > tracker.LastRecordedBlock {
			tracker.LastRecordedBlock = int(log.Raw.BlockNumber)
		}
	}
}

func RecordQuantities(tracker *Tracker, quantities []*big.Int) {
	tracker.Quantities = append(tracker.Quantities, quantities...)
}

func RecordWithdrawals(tracker *Tracker, logs []*mainchain_gateway.MainchainGatewayWithdrew) {
	tracker.TxCount += len(logs)
	for _, log := range logs {
		tracker.AccAmount.Add(tracker.AccAmount, log.Receipt.Info.Quantity)
		if int(log.Raw.BlockNumber) > tracker.LastRecordedBlock {
			tracker.LastRecordedBlock = int(log.Raw.BlockNumber)
		}
	}
}

func RecordReceiptHashes(tracker *Tracker, receiptHashes []common.Hash) {
	tracker.ReceiptHashes = append(tracker.ReceiptHashes, receiptHashes...)
}

func RecordReceiptIds(tracker *Tracker, receiptIds []*big.Int) {
	tracker.ReceiptIds = append(tracker.ReceiptIds, receiptIds...)
}

func RecordMintTxHashes(tracker *Tracker, mintTxHashes []common.Hash) {
	tracker.MintTxHashes = append(tracker.MintTxHashes, mintTxHashes...)
}

func RecordSanitizedTxHashes(tracker *Tracker, sanitizedTxHashes []common.Hash) {
	tracker.SanitizedTxHashes = append(tracker.SanitizedTxHashes, sanitizedTxHashes...)
}
