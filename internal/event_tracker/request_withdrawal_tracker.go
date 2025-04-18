package event_tracker

import (
	"context"
	"math/big"
	"ronin-bridge-snapshot/generated/contract/ronin_gateway_v3"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

type RequestWithdrawalTracker struct {
	*Tracker
	receiptHash2Info   map[common.Hash]*ronin_gateway_v3.TransferReceipt
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
		receiptHash2Info:   make(map[common.Hash]*ronin_gateway_v3.TransferReceipt),
		receiptHash2TxHash: make(map[common.Hash]common.Hash),
		excludeTxHashes:    excludeTxHashesMap,
	}

	r.Tracker.iface, _ = ronin_gateway_v3.RoninGatewayV3MetaData.GetAbi()
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

func (r *RequestWithdrawalTracker) GetReceiptHashInfo(receiptHash common.Hash) *ronin_gateway_v3.TransferReceipt {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.receiptHash2Info[receiptHash]
}

func (r *RequestWithdrawalTracker) Summarize() (erc20ReceiptHashes, erc721ReceiptHashes, erc1155ReceiptHashes, nullERC20ReceiptHashes map[common.Hash]*ronin_gateway_v3.TransferReceipt, tokenAmounts map[common.Address]*big.Int) {
	log.Info("#### Summarizing RequestWithdrawal events ####")

	r.Tracker.mu.Lock()
	defer r.Tracker.mu.Unlock()

	count := r.Total()

	txCount := make(map[common.Address]int)

	tokenAmounts = make(map[common.Address]*big.Int)
	erc721ReceiptHashes = make(map[common.Hash]*ronin_gateway_v3.TransferReceipt)
	erc1155ReceiptHashes = make(map[common.Hash]*ronin_gateway_v3.TransferReceipt)
	erc20ReceiptHashes = make(map[common.Hash]*ronin_gateway_v3.TransferReceipt)
	nullERC20ReceiptHashes = make(map[common.Hash]*ronin_gateway_v3.TransferReceipt)

	for receiptHash, receipt := range r.receiptHash2Info {
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
			log.Debug("Quantity is zero", "receiptHash", receiptHash.Hex(), "txHash", r.receiptHash2TxHash[receiptHash].Hex())
			nullERC20ReceiptHashes[receiptHash] = receipt
			continue
		}
		if receipt.Info.Erc != 0 {
			log.Debug("Not ERC20", "receiptHash", receiptHash.Hex(), "txHash", r.receiptHash2TxHash[receiptHash].Hex())
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

	log.Info("Total RequestWithdrawal events", "total", count, "erc20", len(erc20ReceiptHashes), "erc721", len(erc721ReceiptHashes), "erc1155", len(erc1155ReceiptHashes))

	for token, amount := range tokenAmounts {
		log.Info("Token", "address", token.Hex(), "txCount", txCount[token], "total", amount.String())
	}

	return erc20ReceiptHashes, erc721ReceiptHashes, erc1155ReceiptHashes, nullERC20ReceiptHashes, tokenAmounts
}

func (r *RequestWithdrawalTracker) Record(e *types.Log) error {
	log.Trace("Processing RequestWithdrawal event", "txHash", e.TxHash.Hex())
	event := new(ronin_gateway_v3.RoninGatewayV3WithdrawalRequested)
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

func (r *RequestWithdrawalTracker) validate(e *ronin_gateway_v3.RoninGatewayV3WithdrawalRequested) bool {
	if _, ok := r.excludeTxHashes[e.Raw.TxHash]; ok {
		log.Debug("Transaction hash is excluded", "TxHash", e.Raw.TxHash.String())
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
