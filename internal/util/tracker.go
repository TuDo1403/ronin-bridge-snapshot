package util

import (
	"math/big"
	"ronin-bridge-snapshot/internal/erc20_transfer"
	"ronin-bridge-snapshot/internal/withdrawal"

	"github.com/ethereum/go-ethereum/common"
)

type Tracker struct {
	TxCount           int
	AccAmount         *big.Int
	LastRecordedBlock uint64
	MintTxHashes      []common.Hash
	SanitizedTxHashes []common.Hash
	ReceiptHashes     []common.Hash
	ReceiptIds        []int64
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

func (t *Tracker) RecordWithdrawalReceipts(withdrawalReceipts []*withdrawal.WithdrawalReceipt) {
	t.TxCount += len(withdrawalReceipts)

	for _, receipt := range withdrawalReceipts {
		t.AccAmount.Add(t.AccAmount, receipt.Quantity)

		if receipt.BlockNumber > t.LastRecordedBlock {
			t.LastRecordedBlock = receipt.BlockNumber
		}

		t.Quantities = append(t.Quantities, receipt.Quantity)
		t.ReceiptIds = append(t.ReceiptIds, receipt.ReceiptId)
		t.ReceiptHashes = append(t.ReceiptHashes, receipt.ReceiptHash)
		t.SanitizedTxHashes = append(t.SanitizedTxHashes, receipt.TxHash)
	}
}

func (t *Tracker) RecordTransfers(transfers []*erc20_transfer.Transfer) {
	t.TxCount += len(transfers)

	for _, transfer := range transfers {
		if transfer.BlockNumber > t.LastRecordedBlock {
			t.LastRecordedBlock = transfer.BlockNumber
		}

		if transfer.From == (common.Address{}) {
			t.MintTxHashes = append(t.MintTxHashes, transfer.TxHash)
		} else {
			t.AccAmount.Add(t.AccAmount, transfer.Value)
			t.Quantities = append(t.Quantities, transfer.Value)
			t.SanitizedTxHashes = append(t.SanitizedTxHashes, transfer.TxHash)
		}
	}
}
