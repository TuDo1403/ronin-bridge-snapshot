package event_tracker

import (
	"context"
	"math/big"
	"ronin-bridge-snapshot/internal/abi/mainchain_gateway"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

type WithdrewTracker struct {
	*Tracker
	receiptHashes2Info map[common.Hash]*mainchain_gateway.TransferReceipt
	excludeTxHashes    map[common.Hash]bool
}

func NewWithdrewTracker(ctx context.Context, nWorker int, excludeTxHashes []common.Hash, in <-chan *types.Log) *WithdrewTracker {
	w := &WithdrewTracker{
		Tracker:            NewTracker(ctx, "Withdrew", nWorker, in, nil),
		receiptHashes2Info: make(map[common.Hash]*mainchain_gateway.TransferReceipt),
		excludeTxHashes:    make(map[common.Hash]bool),
	}

	w.Tracker.callback = w.Record
	return w
}

func (w *WithdrewTracker) GetReceiptHashes() []common.Hash {
	w.mu.Lock()
	defer w.mu.Unlock()

	receiptHashes := make([]common.Hash, 0, len(w.receiptHashes2Info))
	for receiptHash := range w.receiptHashes2Info {
		receiptHashes = append(receiptHashes, receiptHash)
	}
	return receiptHashes
}

func (w *WithdrewTracker) GetReceiptHashInfo(receiptHash common.Hash) *mainchain_gateway.TransferReceipt {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.receiptHashes2Info[receiptHash]
}

func (w *WithdrewTracker) Summarize() {
	log.Info("#### Summarizing Withdrew events ####")

	w.Tracker.mu.Lock()
	defer w.Tracker.mu.Unlock()

	count := w.Total()

	tokenAmounts := make(map[common.Address]*big.Int)
	txCount := make(map[common.Address]int)
	erc20Count := 0
	erc721Count := 0
	erc1155Count := 0

	for _, receipt := range w.receiptHashes2Info {
		if receipt.Info.Quantity == big.NewInt(0) || receipt.Info.Erc != 0 {
			if receipt.Info.Erc == 1 {
				erc721Count++
			}
			if receipt.Info.Erc == 2 {
				erc1155Count++
			}

			continue
		}

		tokenAddr := receipt.Ronin.TokenAddr
		if _, ok := tokenAmounts[tokenAddr]; !ok {
			tokenAmounts[tokenAddr] = new(big.Int)
		}

		erc20Count++
		txCount[tokenAddr]++
		tokenAmounts[tokenAddr].Add(tokenAmounts[tokenAddr], receipt.Info.Quantity)
	}

	log.Info("Total Withdrew events", "total", count, "erc20", erc20Count, "erc721", erc721Count, "erc1155", erc1155Count)

	for token, amount := range tokenAmounts {
		log.Info("Token", "address", token.Hex(), "txCount", txCount[token], "total", amount.String())
	}
}

func (w *WithdrewTracker) Record(e *types.Log) error {
	event := new(mainchain_gateway.MainchainGatewayWithdrew)
	parseEvent(w.iface, event, "Withdrew", e)
	event.Raw = *e

	if !w.validate(event) {
		log.Trace("Invalid event", "event", event)
		return nil
	}

	if err := w.Tracker.Record(e); err != nil {
		return err
	}

	// Assuming log is already filtered and valid for this tracker
	w.mu.Lock()
	defer w.mu.Unlock()
	w.receiptHashes2Info[event.ReceiptHash] = &event.Receipt

	return nil
}

func (w *WithdrewTracker) validate(e *mainchain_gateway.MainchainGatewayWithdrew) bool {
	if w.excludeTxHashes[e.Raw.TxHash] {
		log.Warn("Transaction hash is excluded", "TxHash", e.Raw.TxHash.String())
		return false
	}
	if e.Receipt.Kind != 1 {
		return false
	}

	// if e.Receipt.Info.Erc != 0 {
	// 	return false
	// }
	// if e.Receipt.Info.Quantity.Int64() == 0 {
	// 	return false
	// }

	return true
}
