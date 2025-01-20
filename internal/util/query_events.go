package util

import (
	"context"
	"math/big"
	"ronin-bridge-snapshot/internal/abi/erc20"
	"ronin-bridge-snapshot/internal/abi/mainchain_gateway"
	"ronin-bridge-snapshot/internal/abi/ronin_gateway"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

func FilterRequestWithdraw(ctx context.Context, done <-chan struct{}, clients []*ethclient.Client, gatewayAddr common.Address, startBlock, endBlock uint64, rawLogChan chan []types.Log) {
	startTime := time.Now()

	for i, client := range clients {

		gwFilterer, err := ronin_gateway.NewRoninGatewayFilterer(gatewayAddr, client)
		if err != nil {
			log.Crit("failed to create gateway contract", err)
		}

		iterator, err := gwFilterer.FilterWithdrawalRequested(&bind.FilterOpts{
			Start:   startBlock,
			End:     &[]uint64{endBlock}[0],
			Context: ctx,
		})
		if err != nil {
			if i == len(clients)-1 {
				log.Crit("failed to filter request withdrawals logs", err)
			}

			log.Debug("Retrying", "From", startBlock, "To", endBlock, "Using Idx", i+1)
			continue
		}

		var rawLogs []types.Log
		for iterator.Next() {
			rawLogs = append(rawLogs, iterator.Event.Raw)
		}

		responseTime := time.Since(startTime)
		if responseTime > 10*time.Second {
			log.Info("Low Response Time", "From", startBlock, "To", endBlock, "Time", responseTime)
			time.Sleep(60 * time.Second)
		}

		log.Debug("Raw Filter", "Count", len(rawLogs), "From", startBlock, "To", endBlock, "Size", endBlock-startBlock, "Time", responseTime)

		rawLogChan <- rawLogs

		break
	}
}

func SanitizeRequestWithdrawalsLogs(ctx context.Context, done <-chan struct{}, rawLogChan chan []types.Log, sanitizedLogChan chan map[common.Address][]*ronin_gateway.RoninGatewayWithdrawalRequested, sanitizedTxHashesChan chan map[common.Address][]common.Hash, receiptHashesChan chan map[common.Address][]common.Hash, receiptIdsChan chan map[common.Address][]*big.Int, quantitiesChan chan map[common.Address][]*big.Int) {
	// Load ABI once to avoid unnecessary recomputation
	gatewayABI, err := ronin_gateway.RoninGatewayMetaData.GetAbi()
	if err != nil {
		log.Error("failed to parse ABI: %v", err)
		return
	}

	for {
		select {
		case <-ctx.Done():
			log.Trace("Sanitizer routine stopped.")
			return

		case <-done:
			log.Trace("Sanitizer routine received done signal. Exiting.")
			return

		case rawLogs := <-rawLogChan:
			sanitizedLogs := make(map[common.Address][]*ronin_gateway.RoninGatewayWithdrawalRequested)
			sanitizedTxHashes := make(map[common.Address][]common.Hash)
			receiptHashes := make(map[common.Address][]common.Hash)
			receiptIds := make(map[common.Address][]*big.Int)
			quantities := make(map[common.Address][]*big.Int)

			for _, vLog := range rawLogs {
				reqWithdrawEvent := new(ronin_gateway.RoninGatewayWithdrawalRequested)
				err := gatewayABI.UnpackIntoInterface(reqWithdrawEvent, "WithdrawalRequested", vLog.Data)
				if err != nil {
					log.Error("failed to unpack withdrew log", err)
					continue
				}

				tokenAddr := reqWithdrawEvent.Arg1.Ronin.TokenAddr

				if reqWithdrawEvent.Arg1.Kind != 1 {
					log.Crit("Invalid Withdrawal", "TxHash", vLog.TxHash.String(), "Kind", reqWithdrawEvent.Arg1.Kind)
				}
				if reqWithdrawEvent.Arg1.Info.Id.Cmp(big.NewInt(70620)) == 0 {
					log.Crit("Invalid Withdrawal", "TxHash", vLog.TxHash.String(), "Id", reqWithdrawEvent.Arg1.Info.Id)
				}
				if reqWithdrawEvent.Arg1.Info.Erc != 0 {
					log.Warn("Non-ERC20 Withdrawal", "TxHash", vLog.TxHash.String(), "ERC", reqWithdrawEvent.Arg1.Info.Erc)
					continue
				}

				if vLog.Removed {
					log.Warn("Log Removed", "TxHash", vLog.TxHash)
					continue
				}

				if reqWithdrawEvent.Arg1.Info.Quantity.Cmp(big.NewInt(0)) == 0 {
					log.Error("Null Value Withdrawal", "TxHash", vLog.TxHash.String())
				}

				if vLog.TxHash == common.HexToHash("0xa53c7c4402b40a63e485a1a8798d7b7ac6a151dce19c3486d9928575476ac3ce") ||
					vLog.TxHash == common.HexToHash("0x6af887fda63d7ec5538c43fdd3b1fa87ce5e2aaf824c5332f3bd8330d45f4c0b") ||
					vLog.TxHash == common.HexToHash("0x2392f70de576a6e90d95eb4686561ab6232fda7d972ca56fa94df20840f289c9") ||
					vLog.TxHash == common.HexToHash("0xc27bbc806abf466645a24f2be8bb7f082cba007f52767739eee0815153ae6fea") ||
					vLog.TxHash == common.HexToHash("0x78c3e08640228ff072749632f407617695a731f2dd377fa3b1a87f7697282897") ||
					vLog.TxHash == common.HexToHash("0x542c6ddaf6d3752d488153b212346b94bbe30d8b25eee94c80051baa4a8714b7") ||
					vLog.TxHash == common.HexToHash("0x2c4db38cafde1d399e07a754b6605bc5dd242119929dfd8382dfb9925d322379") ||
					vLog.TxHash == common.HexToHash("0x534c69d4c943b20f180774b942147e446e39e12920057e9795927defa547e2ab") ||
					vLog.TxHash == common.HexToHash("0xced4cf252ff100eb7cc96859fc0e85fcad97d97d26f96ed4ae78665301e09094") ||
					vLog.TxHash == common.HexToHash("0x28211787395779c94fca595000f3d079e86ebe12b92e4f005fad4dae5c44d2e9") ||
					vLog.TxHash == common.HexToHash("0x4887f384ba119800b85cb65e146bb559972cc15fab65e5d8ccafd32999b419c9") ||
					vLog.TxHash == common.HexToHash("0x1357e4fa4b955249724d437f634105195f1337edc2db8905a72e69b35ccbd870") {

					log.Warn("Hacked Tx Detected", "TxHash", vLog.TxHash.String(), "From", vLog.Address, "Token", tokenAddr, "Amount", reqWithdrawEvent.Arg1.Info.Quantity)
					continue
				}

				log.Trace("Request Withdrawal Event", "Token", tokenAddr, "Amount", reqWithdrawEvent.Arg1.Info.Quantity, "TxHash", vLog.TxHash.String())

				// Attach original log data
				reqWithdrawEvent.Raw = vLog

				// Push to sanitized logs
				receiptHashes[tokenAddr] = append(receiptHashes[tokenAddr], reqWithdrawEvent.ReceiptHash)
				receiptIds[tokenAddr] = append(receiptIds[tokenAddr], reqWithdrawEvent.Arg1.Id)
				quantities[tokenAddr] = append(quantities[tokenAddr], reqWithdrawEvent.Arg1.Info.Quantity)
				sanitizedLogs[tokenAddr] = append(sanitizedLogs[tokenAddr], reqWithdrawEvent)
				sanitizedTxHashes[tokenAddr] = append(sanitizedTxHashes[tokenAddr], vLog.TxHash)
			}

			// Send sanitized data to respective channels
			receiptHashesChan <- receiptHashes
			receiptIdsChan <- receiptIds
			quantitiesChan <- quantities
			sanitizedTxHashesChan <- sanitizedTxHashes
			sanitizedLogChan <- sanitizedLogs

			log.Trace("Sanitized Logs", "Count", len(sanitizedLogs))
		}
	}
}

func FilterWithdrew(ctx context.Context, done <-chan struct{}, clients []*ethclient.Client, gatewayAddr common.Address, startBlock, endBlock uint64, rawLogChan chan []types.Log) {
	startTime := time.Now()

	for i, client := range clients {

		gwFilterer, err := mainchain_gateway.NewMainchainGatewayFilterer(gatewayAddr, client)
		if err != nil {
			log.Crit("failed to create gateway contract", err)
		}

		iterator, err := gwFilterer.FilterWithdrew(&bind.FilterOpts{
			Start:   startBlock,
			End:     &[]uint64{endBlock}[0],
			Context: ctx,
		})
		if err != nil {
			if i == len(clients)-1 {
				log.Crit("failed to filter withdrew logs", err)
			}

			log.Debug("Retrying", "From", startBlock, "To", endBlock, "Using Idx", i+1)
			continue
		}

		var rawLogs []types.Log
		for iterator.Next() {
			rawLogs = append(rawLogs, iterator.Event.Raw)
		}

		responseTime := time.Since(startTime)
		if responseTime > 10*time.Second {
			log.Info("Low Response Time", "From", startBlock, "To", endBlock, "Time", responseTime)
			time.Sleep(60 * time.Second)
		}

		log.Debug("Raw Filter", "Count", len(rawLogs), "From", startBlock, "To", endBlock, "Size", endBlock-startBlock, "Time", responseTime)

		rawLogChan <- rawLogs

		break
	}
}

func SanitizeWithdrewLogs(ctx context.Context, done <-chan struct{}, rawLogChan chan []types.Log, sanitizedLogChan chan map[common.Address][]*mainchain_gateway.MainchainGatewayWithdrew, sanitizedTxHashesChan chan map[common.Address][]common.Hash, receiptHashesChan chan map[common.Address][]common.Hash, receiptIdsChan chan map[common.Address][]*big.Int, quantitiesChan chan map[common.Address][]*big.Int) {
	// Load ABI once to avoid unnecessary recomputation
	gatewayABI, err := mainchain_gateway.MainchainGatewayMetaData.GetAbi()
	if err != nil {
		log.Error("failed to parse ABI: %v", err)
		return
	}

	for {
		select {
		case <-ctx.Done():
			log.Trace("Sanitizer routine stopped.")
			return

		case <-done:
			log.Trace("Sanitizer routine received done signal. Exiting.")
			return

		case rawLogs := <-rawLogChan:
			sanitizedLogs := make(map[common.Address][]*mainchain_gateway.MainchainGatewayWithdrew)
			sanitizedTxHashes := make(map[common.Address][]common.Hash)
			receiptHashes := make(map[common.Address][]common.Hash)
			receiptIds := make(map[common.Address][]*big.Int)
			quantities := make(map[common.Address][]*big.Int)
			for _, vLog := range rawLogs {
				withdrewEvent := new(mainchain_gateway.MainchainGatewayWithdrew)
				err := gatewayABI.UnpackIntoInterface(withdrewEvent, "Withdrew", vLog.Data)
				if err != nil {
					log.Error("failed to unpack withdrew log", err)
					continue
				}

				tokenAddr := withdrewEvent.Receipt.Mainchain.TokenAddr

				if withdrewEvent.Receipt.Kind != 1 {
					log.Crit("Invalid Withdrawal", "TxHash", vLog.TxHash.String(), "Kind", withdrewEvent.Receipt.Kind)
				}
				if withdrewEvent.Receipt.Info.Erc != 0 {
					log.Warn("Non-ERC20 Withdrawal", "TxHash", vLog.TxHash.String(), "ERC", withdrewEvent.Receipt.Info.Erc)
					continue
				}

				if vLog.Removed {
					log.Warn("Log Removed", "TxHash", vLog.TxHash)
					continue
				}

				if withdrewEvent.Receipt.Info.Quantity.Cmp(big.NewInt(0)) == 0 {
					log.Crit("Null Value Withdrawal", "TxHash", vLog.TxHash.String())
				}

				if vLog.TxHash == common.HexToHash("0x2619570088683e6cc3a38d93c3d98899e5783864e15525d5f5810c11189ba6cb") ||
					vLog.TxHash == common.HexToHash("0x77bb83dd9b25a0af12c0092993ddf1a2f54b77fbcd1dbf1bf33a965ec707268b") ||
					vLog.TxHash == common.HexToHash("0xbce5b8548db486c561948e8a177c8ccaa72810f972cee3909ea50af015a60ad8") ||
					vLog.TxHash == common.HexToHash("0xb0c08ecdad943ad3830f610c287fc7289b59994f5a84e6658f3fe2582625ea06") ||
					vLog.TxHash == common.HexToHash("0xb0fc7fae9a3b9f210183f151c59fcb892faa6dbd0b7bd789f26b877147af8c41") ||
					vLog.TxHash == common.HexToHash("0x5623319a8232e6d22aa7c652ecc120889d9a3c073fa9bee0bca8d4dece7680b1") ||
					vLog.TxHash == common.HexToHash("0x368e8fc1a1fc9a1896616819b5393ff118954e5c964365f439ed66deba3cd5c8") ||
					vLog.TxHash == common.HexToHash("0xbaf540fe846da0537c49d9e97c509abd1b9f336893763a04b479c0c3a65afd3e") ||
					vLog.TxHash == common.HexToHash("0x4f519d09f6313c22864b7bec6f5f4a07f56b88a356bc3d0a06b659628766a536") ||
					vLog.TxHash == common.HexToHash("0xc8e0141694aac94999da6432f375db44f0a379c8fa6777a8e1f5c1009feab025") ||
					vLog.TxHash == common.HexToHash("0x6491f922971341c2437e8a77320ab2e258d0493c8fe7cec6b60cb30a6ff1b3f5") ||
					vLog.TxHash == common.HexToHash("0x5b99165fb07881590fab962185b5dee877f2978fac83f5f76c88639edcba827d") {

					log.Warn("MEV Detected", "TxHash", vLog.TxHash.String(), "From", vLog.Address, "Token", tokenAddr, "Amount", withdrewEvent.Receipt.Info.Quantity)
					continue
				}

				log.Trace("Withdrew Event", "Token", tokenAddr, "Amount", withdrewEvent.Receipt.Info.Quantity, "TxHash", vLog.TxHash.String())

				// Attach original log data
				withdrewEvent.Raw = vLog

				// Push to sanitized logs
				receiptHashes[tokenAddr] = append(receiptHashes[tokenAddr], withdrewEvent.ReceiptHash)
				receiptIds[tokenAddr] = append(receiptIds[tokenAddr], withdrewEvent.Receipt.Id)
				quantities[tokenAddr] = append(quantities[tokenAddr], withdrewEvent.Receipt.Info.Quantity)
				sanitizedLogs[tokenAddr] = append(sanitizedLogs[tokenAddr], withdrewEvent)
				sanitizedTxHashes[tokenAddr] = append(sanitizedTxHashes[tokenAddr], vLog.TxHash)
			}

			// Send sanitized data to respective channels
			sanitizedLogChan <- sanitizedLogs
			sanitizedTxHashesChan <- sanitizedTxHashes
			receiptHashesChan <- receiptHashes
			receiptIdsChan <- receiptIds
			quantitiesChan <- quantities

			log.Trace("Sanitized Logs", "Count", len(sanitizedLogs))
		}
	}
}

func FilterERC20TransferForTokens(ctx context.Context, done <-chan struct{}, client *ethclient.Client, tokenAddrs []common.Address, toAddr common.Address, startBlock, endBlock uint64, rawLogChan chan []types.Log) {
	if len(tokenAddrs) == 0 {
		log.Error("no token addresses provided")
		return
	}

	erc20ABI, err := erc20.Erc20MetaData.GetAbi() // or however you load your ABI
	if err != nil {
		log.Error("failed to parse ABI", err)
		return
	}

	startTime := time.Now()

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(startBlock)),
		ToBlock:   big.NewInt(int64(endBlock)),
		Addresses: tokenAddrs,
		Topics: [][]common.Hash{
			{erc20ABI.Events["Transfer"].ID},
			{},
			{common.BytesToHash(toAddr.Bytes())},
		},
	}

	rawLogs, err := client.FilterLogs(ctx, query)
	if err != nil {
		log.Error("failed to filter logs for query", query, err)
	}

	responseTime := time.Since(startTime)
	if responseTime > 10*time.Second {
		log.Info("Low Response Time", "From", startBlock, "To", endBlock, "Time", responseTime)
		time.Sleep(60 * time.Second)
	}

	log.Debug("Raw Filter", "Count", len(rawLogs), "From", startBlock, "To", endBlock, "Size", endBlock-startBlock, "Time", responseTime)
	rawLogChan <- rawLogs
}

func SanitizeTransferLogs(ctx context.Context, done <-chan struct{}, toAddr common.Address, rawLogChan chan []types.Log, sanitizedLogChan chan map[common.Address][]*erc20.Erc20Transfer, mintTxHashesChan chan map[common.Address][]common.Hash, sanitizedTxHashesChan chan map[common.Address][]common.Hash) {
	// Load ABI once to avoid unnecessary recomputation
	erc20ABI, err := erc20.Erc20MetaData.GetAbi()
	if err != nil {
		log.Error("failed to parse ABI: %v", err)
		return
	}

	for {
		select {
		case <-ctx.Done():
			log.Trace("Sanitizer routine stopped.")
			return

		case <-done:
			log.Trace("Sanitizer routine received done signal. Exiting.")
			return

		case rawLogs := <-rawLogChan:
			sanitizedLogs := make(map[common.Address][]*erc20.Erc20Transfer)
			sanitizedTxHashes := make(map[common.Address][]common.Hash)
			mintTxHashes := make(map[common.Address][]common.Hash)

			for _, vLog := range rawLogs {
				tokenAddr := vLog.Address

				transferEvent := new(erc20.Erc20Transfer)
				err := erc20ABI.UnpackIntoInterface(transferEvent, "Transfer", vLog.Data)
				if err != nil {
					log.Error("failed to unpack transfer log", err)
					continue
				}

				transferEvent.From = common.BytesToAddress(vLog.Topics[1].Bytes())
				transferEvent.To = common.BytesToAddress(vLog.Topics[2].Bytes())

				// If "From" address is zero, it's a mint transaction
				if transferEvent.From == (common.Address{}) {
					log.Trace("Mint Transaction", "To", transferEvent.To, "TxHash", vLog.TxHash)
					mintTxHashes[tokenAddr] = append(mintTxHashes[tokenAddr], vLog.TxHash)
					continue
				}

				// If it's not sent to `toAddr`, ignore it
				if transferEvent.To != toAddr {
					log.Trace("Not Sent To Gateway", "To", transferEvent.To, "Gateway", toAddr, "TxHash", vLog.TxHash)
					continue
				}

				// Skip zero-value transfers
				if transferEvent.Value.Cmp(big.NewInt(0)) == 0 {
					log.Trace("Null Value Transfer", "From", transferEvent.From, "To", transferEvent.To, "TxHash", vLog.TxHash)
					continue
				}

				if vLog.Removed {
					log.Warn("Log Removed", "TxHash", vLog.TxHash)
					continue
				}

				if vLog.TxHash == common.HexToHash("0xa53c7c4402b40a63e485a1a8798d7b7ac6a151dce19c3486d9928575476ac3ce") ||
					vLog.TxHash == common.HexToHash("0x6af887fda63d7ec5538c43fdd3b1fa87ce5e2aaf824c5332f3bd8330d45f4c0b") ||
					vLog.TxHash == common.HexToHash("0x2392f70de576a6e90d95eb4686561ab6232fda7d972ca56fa94df20840f289c9") ||
					vLog.TxHash == common.HexToHash("0xc27bbc806abf466645a24f2be8bb7f082cba007f52767739eee0815153ae6fea") ||
					vLog.TxHash == common.HexToHash("0x78c3e08640228ff072749632f407617695a731f2dd377fa3b1a87f7697282897") ||
					vLog.TxHash == common.HexToHash("0x542c6ddaf6d3752d488153b212346b94bbe30d8b25eee94c80051baa4a8714b7") ||
					vLog.TxHash == common.HexToHash("0x2c4db38cafde1d399e07a754b6605bc5dd242119929dfd8382dfb9925d322379") ||
					vLog.TxHash == common.HexToHash("0x534c69d4c943b20f180774b942147e446e39e12920057e9795927defa547e2ab") ||
					vLog.TxHash == common.HexToHash("0xced4cf252ff100eb7cc96859fc0e85fcad97d97d26f96ed4ae78665301e09094") ||
					vLog.TxHash == common.HexToHash("0x28211787395779c94fca595000f3d079e86ebe12b92e4f005fad4dae5c44d2e9") ||
					vLog.TxHash == common.HexToHash("0x4887f384ba119800b85cb65e146bb559972cc15fab65e5d8ccafd32999b419c9") ||
					vLog.TxHash == common.HexToHash("0x1357e4fa4b955249724d437f634105195f1337edc2db8905a72e69b35ccbd870") {

					log.Warn("Hack Txs", "TxHash", vLog.TxHash.String(), "From", vLog.Address, "Token", tokenAddr, "Amount", transferEvent.Value)
					continue
				}

				log.Trace("Transfer Event", "From", transferEvent.From, "To", transferEvent.To, "Value", transferEvent.Value, "TxHash", vLog.TxHash)
				// Attach original log data
				transferEvent.Raw = vLog

				// Push to sanitized logs
				sanitizedLogs[tokenAddr] = append(sanitizedLogs[tokenAddr], transferEvent)
				sanitizedTxHashes[tokenAddr] = append(sanitizedTxHashes[tokenAddr], vLog.TxHash)
			}

			// Send sanitized data to respective channels
			sanitizedLogChan <- sanitizedLogs
			mintTxHashesChan <- mintTxHashes
			sanitizedTxHashesChan <- sanitizedTxHashes

			log.Trace("Sanitized Logs", "Count", len(sanitizedLogs))
		}
	}
}
