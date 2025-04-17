package event_tracker

import (
	"context"
	"math/big"
	"ronin-bridge-snapshot/internal/abi/ronin_gateway"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

type RequestWithdrawalTracker struct {
	*Tracker
	receiptHash2Info   map[common.Hash]*ronin_gateway.TransferReceipt
	receiptHash2TxHash map[common.Hash]common.Hash
	excludeTxHashes    map[common.Hash]struct{}
}

func NewRequestWithdrawalTracker(ctx context.Context, nWorker int, excludeTxHashes []common.Hash, in <-chan *types.Log) *RequestWithdrawalTracker {
	excludeTxHashesMap := make(map[common.Hash]struct{})
	for _, txHash := range excludeTxHashes {
		excludeTxHashesMap[txHash] = struct{}{}
	}
	r := &RequestWithdrawalTracker{
		Tracker:            NewTracker(ctx, "WithdrawalRequested", nWorker, in, nil),
		receiptHash2Info:   make(map[common.Hash]*ronin_gateway.TransferReceipt),
		receiptHash2TxHash: make(map[common.Hash]common.Hash),
		excludeTxHashes:    excludeTxHashesMap,
	}

	r.Tracker.iface, _ = ronin_gateway.RoninGatewayMetaData.GetAbi()
	r.Tracker.callback = r.Record
	return r
}

func (r *RequestWithdrawalTracker) GetReceiptHashes() []common.Hash {
	r.mu.Lock()
	defer r.mu.Unlock()

	receiptHashes := make([]common.Hash, 0, len(r.receiptHash2Info))
	for receiptHash := range r.receiptHash2Info {
		receiptHashes = append(receiptHashes, receiptHash)
	}
	return receiptHashes
}

func (r *RequestWithdrawalTracker) GetReceiptHashInfo(receiptHash common.Hash) *ronin_gateway.TransferReceipt {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.receiptHash2Info[receiptHash]
}

func (r *RequestWithdrawalTracker) Summarize() {
	log.Info("#### Summarizing RequestWithdrawal events ####")

	r.Tracker.mu.Lock()
	defer r.Tracker.mu.Unlock()

	count := r.Total()

	tokenAmounts := make(map[common.Address]*big.Int)
	txCount := make(map[common.Address]int)
	erc20Count := 0
	erc721Count := 0
	erc1155Count := 0

	for receiptHash, receipt := range r.receiptHash2Info {
		if receipt.Info.Quantity == big.NewInt(0) || receipt.Info.Erc != 0 {
			if receipt.Info.Erc == 1 {
				erc721Count++
			}
			if receipt.Info.Erc == 2 {
				erc1155Count++
			}

			continue
		}

		if receipt.Info.Quantity.Cmp(big.NewInt(0)) == 0 {
			log.Warn("Quantity is zero", "receiptHash", receiptHash.Hex(), "txHash", r.receiptHash2TxHash[receiptHash].Hex())
			continue
		}
		if receipt.Info.Erc != 0 {
			log.Warn("Erc is not zero", "receiptHash", receiptHash.Hex(), "txHash", r.receiptHash2TxHash[receiptHash].Hex())
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

	log.Info("Total RequestWithdrawal events", "total", count, "erc20", erc20Count, "erc721", erc721Count, "erc1155", erc1155Count)

	for token, amount := range tokenAmounts {
		log.Info("Token", "address", token.Hex(), "txCount", txCount[token], "total", amount.String())
	}
}

func (r *RequestWithdrawalTracker) Record(e *types.Log) error {
	log.Trace("Processing RequestWithdrawal event", "txHash", e.TxHash.Hex())
	event := new(ronin_gateway.RoninGatewayWithdrawalRequested)
	parseEvent(r.iface, event, "WithdrawalRequested", e)
	event.Raw = *e

	if !r.validate(event) {
		log.Trace("Invalid event", "txHash", e.TxHash.Hex(), "blockNumber", e.BlockNumber)
		return nil
	}

	if err := r.Tracker.Record(e); err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	r.receiptHash2Info[event.ReceiptHash] = &event.Arg1
	r.receiptHash2TxHash[event.ReceiptHash] = e.TxHash
	return nil
}

func (r *RequestWithdrawalTracker) validate(e *ronin_gateway.RoninGatewayWithdrawalRequested) bool {
	if _, ok := r.excludeTxHashes[e.Raw.TxHash]; ok {
		log.Warn("Transaction hash is excluded", "TxHash", e.Raw.TxHash.String())
		return false
	}

	if e.Arg1.Kind != 1 {
		return false
	}
	// if e.Arg1.Info.Erc != 0 {
	// 	return false
	// }
	// if e.Arg1.Info.Quantity.Int64() == 0 {
	// 	return false
	// }

	return true
}
