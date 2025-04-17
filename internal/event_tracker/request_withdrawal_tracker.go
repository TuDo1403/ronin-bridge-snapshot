package event_tracker

import (
	"context"
	"math/big"
	"ronin-bridge-snapshot/internal/abi/ronin_gateway"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

type RequestWithdrawalInfo struct {
	receiptId   int64
	receiptHash common.Hash
	quantity    *big.Int
	localAddr   common.Address
	remoteAddr  common.Address
}

type RequestWithdrawalTracker struct {
	*Tracker
	txHashes2Info   map[common.Hash]*RequestWithdrawalInfo
	excludeTxHashes map[common.Hash]bool
}

func NewRequestWithdrawalTracker(wg *sync.WaitGroup, ctx context.Context, nWorker int, in <-chan *types.Log) *RequestWithdrawalTracker {
	r := &RequestWithdrawalTracker{
		Tracker:         NewTracker(wg, ctx, nWorker, in, nil),
		txHashes2Info:   make(map[common.Hash]*RequestWithdrawalInfo),
		excludeTxHashes: make(map[common.Hash]bool),
	}

	r.Tracker.iface, _ = ronin_gateway.RoninGatewayMetaData.GetAbi()
	r.Tracker.callback = r.Record
	return r
}

func (r *RequestWithdrawalTracker) Summarize() {
	log.Info("\nSummarizing RequestWithdrawal events")
	count := r.Total()
	if count > 0 {
		log.Info("Total RequestWithdrawal events recorded", "count", count)
	} else {
		log.Info("No RequestWithdrawal events recorded")
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
	r.txHashes2Info[e.TxHash] = &RequestWithdrawalInfo{
		receiptId:   event.Arg1.Id.Int64(),
		receiptHash: event.ReceiptHash,
		quantity:    event.Arg1.Info.Quantity,
		localAddr:   event.Arg1.Ronin.Addr,
		remoteAddr:  event.Arg1.Mainchain.Addr,
	}

	return nil
}

func (r *RequestWithdrawalTracker) validate(e *ronin_gateway.RoninGatewayWithdrawalRequested) bool {
	if r.excludeTxHashes[e.Raw.TxHash] {
		log.Warn("Transaction hash is excluded", "TxHash", e.Raw.TxHash.String())
		return false
	}
	if e.Arg1.Kind != 1 {
		return false
	}
	if e.Arg1.Info.Erc != 0 {
		return false
	}
	if e.Arg1.Info.Quantity.Int64() == 0 {
		log.Warn("Quantity is zero", "txHash", e.Raw.TxHash.String(), "blockNumber", e.Raw.BlockNumber)
		return false
	}

	return true
}
