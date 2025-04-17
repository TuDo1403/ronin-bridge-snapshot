package event_tracker

import (
	"context"
	"math/big"
	"ronin-bridge-snapshot/internal/abi/mainchain_gateway"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

type WithdrewInfo struct {
	receiptId   int64
	receiptHash common.Hash
	quantity    *big.Int
	localAddr   common.Address
	remoteAddr  common.Address
}

type WithdrewTracker struct {
	*Tracker
	txHashes2Info   map[common.Hash]*WithdrewInfo
	excludeTxHashes map[common.Hash]bool
}

func NewWithdrewTracker(wg *sync.WaitGroup, ctx context.Context, nWorker int, in <-chan *types.Log) *WithdrewTracker {
	w := &WithdrewTracker{
		Tracker:         NewTracker(wg, ctx, nWorker, in, nil),
		txHashes2Info:   make(map[common.Hash]*WithdrewInfo),
		excludeTxHashes: make(map[common.Hash]bool),
	}

	w.Tracker.callback = w.Record
	return w
}

func (w *WithdrewTracker) Record(e *types.Log) error {

	event := new(mainchain_gateway.MainchainGatewayWithdrew)
	parseEvent(w.iface, event, "Withdrew", e)

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
	w.txHashes2Info[e.TxHash] = &WithdrewInfo{
		receiptId:   event.Receipt.Id.Int64(),
		receiptHash: event.ReceiptHash,
		quantity:    event.Receipt.Info.Quantity,
		localAddr:   event.Receipt.Mainchain.Addr,
		remoteAddr:  event.Receipt.Ronin.Addr,
	}

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
	if e.Receipt.Info.Erc != 0 {
		return false
	}
	if e.Receipt.Info.Quantity.Int64() == 0 {
		log.Warn("Quantity is zero", "txHash", e.Raw.TxHash.String(), "blockNumber", e.Raw.BlockNumber)
		return false
	}

	return true
}
