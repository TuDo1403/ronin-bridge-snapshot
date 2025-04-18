package event_tracker

import (
	"context"
	"math/big"
	"ronin-bridge-snapshot/generated/contract/mainchain_gateway_v3"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

type WithdrewTracker struct {
	*Tracker
	receiptHashes2Info map[common.Hash]*mainchain_gateway_v3.TransferReceipt
	receiptHash2TxHash map[common.Hash]common.Hash
	excludeTxHashes    map[common.Hash]struct{}
}

func NewWithdrewTracker(ctx context.Context, nWorker int, excludeTxHashes []common.Hash, in <-chan *types.Log) *WithdrewTracker {
	excludeTxHashesMap := make(map[common.Hash]struct{})
	for _, txHash := range excludeTxHashes {
		excludeTxHashesMap[txHash] = struct{}{}
	}

	w := &WithdrewTracker{
		Tracker:            NewTracker(ctx, "Withdrew", nWorker, in, nil),
		receiptHashes2Info: make(map[common.Hash]*mainchain_gateway_v3.TransferReceipt),
		receiptHash2TxHash: make(map[common.Hash]common.Hash),
		excludeTxHashes:    excludeTxHashesMap,
	}

	w.Tracker.callback = w.Record
	w.Tracker.iface, _ = mainchain_gateway_v3.MainchainGatewayV3MetaData.GetAbi()
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

func (w *WithdrewTracker) GetReceiptHashInfo(receiptHash common.Hash) *mainchain_gateway_v3.TransferReceipt {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.receiptHashes2Info[receiptHash]
}

func (w *WithdrewTracker) Summarize() (erc20ReceiptHashes, erc721ReceiptHashes, erc1155ReceiptHashes map[common.Hash]*mainchain_gateway_v3.TransferReceipt, tokenAmounts map[common.Address]*big.Int) {
	log.Info("#### Summarizing Withdrew events ####")

	w.Tracker.mu.Lock()
	defer w.Tracker.mu.Unlock()

	count := w.Total()

	txCount := make(map[common.Address]int)

	tokenAmounts = make(map[common.Address]*big.Int)
	erc721ReceiptHashes = make(map[common.Hash]*mainchain_gateway_v3.TransferReceipt)
	erc1155ReceiptHashes = make(map[common.Hash]*mainchain_gateway_v3.TransferReceipt)
	erc20ReceiptHashes = make(map[common.Hash]*mainchain_gateway_v3.TransferReceipt)

	for receiptHash, receipt := range w.receiptHashes2Info {
		if receipt.Info.Quantity == big.NewInt(0) || receipt.Info.Erc != 0 {
			if receipt.Info.Erc == 1 {
				erc721ReceiptHashes[receiptHash] = receipt
			}
			if receipt.Info.Erc == 2 {
				erc1155ReceiptHashes[receiptHash] = receipt
			}

			continue
		}

		if receipt.Info.Quantity.Cmp(big.NewInt(0)) == 0 {
			log.Debug("Quantity is zero", "receiptHash", receiptHash.Hex(), "txHash", w.receiptHash2TxHash[receiptHash].Hex())
			continue
		}
		if receipt.Info.Erc != 0 {
			log.Debug("Not ERC20", "receiptHash", receiptHash.Hex(), "txHash", w.receiptHash2TxHash[receiptHash].Hex())
			continue
		}

		if receipt.Info.Quantity.Cmp(big.NewInt(0)) > 0 && receipt.Info.Erc == 0 {
			erc20ReceiptHashes[receiptHash] = receipt

			tokenAddr := receipt.Ronin.TokenAddr
			if _, ok := tokenAmounts[tokenAddr]; !ok {
				tokenAmounts[tokenAddr] = new(big.Int)
			}

			txCount[tokenAddr]++
			tokenAmounts[tokenAddr].Add(tokenAmounts[tokenAddr], receipt.Info.Quantity)
		}
	}

	log.Info("Total Withdrew events", "total", count, "erc20", len(erc20ReceiptHashes), "erc721", len(erc721ReceiptHashes), "erc1155", len(erc1155ReceiptHashes))

	for token, amount := range tokenAmounts {
		log.Info("Token", "address", token.Hex(), "txCount", txCount[token], "total", amount.String())
	}

	return erc20ReceiptHashes, erc721ReceiptHashes, erc1155ReceiptHashes, tokenAmounts
}

func (w *WithdrewTracker) Record(e *types.Log) error {
	event := new(mainchain_gateway_v3.MainchainGatewayV3Withdrew)
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
	w.receiptHash2TxHash[event.ReceiptHash] = e.TxHash

	return nil
}

func (w *WithdrewTracker) validate(e *mainchain_gateway_v3.MainchainGatewayV3Withdrew) bool {
	if _, ok := w.excludeTxHashes[e.Raw.TxHash]; ok {
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
