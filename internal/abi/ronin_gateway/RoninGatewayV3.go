// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ronin_gateway

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MappedTokenConsumerMappedToken is an auto generated low-level Go binding around an user-defined struct.
type MappedTokenConsumerMappedToken struct {
	Erc       uint8
	TokenAddr common.Address
}

// TokenInfo is an auto generated low-level Go binding around an user-defined struct.
type TokenInfo struct {
	Erc      uint8
	Id       *big.Int
	Quantity *big.Int
}

// TokenOwner is an auto generated low-level Go binding around an user-defined struct.
type TokenOwner struct {
	Addr      common.Address
	TokenAddr common.Address
	ChainId   *big.Int
}

// TransferReceipt is an auto generated low-level Go binding around an user-defined struct.
type TransferReceipt struct {
	Id        *big.Int
	Kind      uint8
	Mainchain TokenOwner
	Ronin     TokenOwner
	Info      TokenInfo
}

// TransferRequest is an auto generated low-level Go binding around an user-defined struct.
type TransferRequest struct {
	RecipientAddr common.Address
	TokenAddr     common.Address
	Info          TokenInfo
}

// RoninGatewayMetaData contains all meta data concerning the RoninGateway contract.
var RoninGatewayMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"bulkRequestWithdrawalFor\",\"inputs\":[{\"name\":\"_requests\",\"type\":\"tuple[]\",\"internalType\":\"structTransfer.Request[]\",\"components\":[{\"name\":\"recipientAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"info\",\"type\":\"tuple\",\"internalType\":\"structTokenInfo\",\"components\":[{\"name\":\"erc\",\"type\":\"uint8\",\"internalType\":\"enumTokenStandard\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"_chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bulkSubmitWithdrawalSignatures\",\"inputs\":[{\"name\":\"_withdrawals\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_signatures\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositFor\",\"inputs\":[{\"name\":\"_receipt\",\"type\":\"tuple\",\"internalType\":\"structTransfer.Receipt\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumTransfer.Kind\"},{\"name\":\"mainchain\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"ronin\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"info\",\"type\":\"tuple\",\"internalType\":\"structTokenInfo\",\"components\":[{\"name\":\"erc\",\"type\":\"uint8\",\"internalType\":\"enumTokenStandard\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositVote\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumVoteStatusConsumer.VoteStatus\"},{\"name\":\"finalHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiredAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositVoted\",\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_depositId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_voter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMainchainToken\",\"inputs\":[{\"name\":\"_roninToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"_token\",\"type\":\"tuple\",\"internalType\":\"structMappedTokenConsumer.MappedToken\",\"components\":[{\"name\":\"erc\",\"type\":\"uint8\",\"internalType\":\"enumTokenStandard\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWithdrawalSignatures\",\"inputs\":[{\"name\":\"_withdrawalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_validators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mainchainWithdrew\",\"inputs\":[{\"name\":\"_withdrawalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mainchainWithdrewVoted\",\"inputs\":[{\"name\":\"_withdrawalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_voter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mapTokens\",\"inputs\":[{\"name\":\"_roninTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_mainchainTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"chainIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_standards\",\"type\":\"uint8[]\",\"internalType\":\"enumTokenStandard[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mapTokensWithMinThresholds\",\"inputs\":[{\"name\":\"roninTokens_\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"mainchainTokens_\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"chainIds_\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"standards_\",\"type\":\"uint8[]\",\"internalType\":\"enumTokenStandard[]\"},{\"name\":\"minimumThresholds_\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestWithdrawalFor\",\"inputs\":[{\"name\":\"_request\",\"type\":\"tuple\",\"internalType\":\"structTransfer.Request\",\"components\":[{\"name\":\"recipientAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"info\",\"type\":\"tuple\",\"internalType\":\"structTokenInfo\",\"components\":[{\"name\":\"erc\",\"type\":\"uint8\",\"internalType\":\"enumTokenStandard\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"_chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestWithdrawalSignatures\",\"inputs\":[{\"name\":\"_withdrawalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tryBulkAcknowledgeMainchainWithdrew\",\"inputs\":[{\"name\":\"_withdrawalIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tryBulkDepositFor\",\"inputs\":[{\"name\":\"_receipts\",\"type\":\"tuple[]\",\"internalType\":\"structTransfer.Receipt[]\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumTransfer.Kind\"},{\"name\":\"mainchain\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"ronin\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"info\",\"type\":\"tuple\",\"internalType\":\"structTokenInfo\",\"components\":[{\"name\":\"erc\",\"type\":\"uint8\",\"internalType\":\"enumTokenStandard\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unmapTokens\",\"inputs\":[{\"name\":\"roninTokens_\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"chainIds_\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawalCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DepositVoted\",\"inputs\":[{\"name\":\"bridgeOperator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiptHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"receiptHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"receipt\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structTransfer.Receipt\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumTransfer.Kind\"},{\"name\":\"mainchain\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"ronin\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"info\",\"type\":\"tuple\",\"internalType\":\"structTokenInfo\",\"components\":[{\"name\":\"erc\",\"type\":\"uint8\",\"internalType\":\"enumTokenStandard\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MainchainWithdrew\",\"inputs\":[{\"name\":\"receiptHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"receipt\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structTransfer.Receipt\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumTransfer.Kind\"},{\"name\":\"mainchain\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"ronin\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"info\",\"type\":\"tuple\",\"internalType\":\"structTokenInfo\",\"components\":[{\"name\":\"erc\",\"type\":\"uint8\",\"internalType\":\"enumTokenStandard\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenMapped\",\"inputs\":[{\"name\":\"roninTokens\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"mainchainTokens\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"chainIds\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"standards\",\"type\":\"uint8[]\",\"indexed\":false,\"internalType\":\"enumTokenStandard[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenUnmapped\",\"inputs\":[{\"name\":\"roninTokens\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"chainIds\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TrustedThresholdUpdated\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"numerator\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"denominator\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"previousNumerator\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"previousDenominator\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalRequested\",\"inputs\":[{\"name\":\"receiptHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structTransfer.Receipt\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumTransfer.Kind\"},{\"name\":\"mainchain\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"ronin\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"info\",\"type\":\"tuple\",\"internalType\":\"structTokenInfo\",\"components\":[{\"name\":\"erc\",\"type\":\"uint8\",\"internalType\":\"enumTokenStandard\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalSignaturesRequested\",\"inputs\":[{\"name\":\"receiptHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structTransfer.Receipt\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumTransfer.Kind\"},{\"name\":\"mainchain\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"ronin\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"info\",\"type\":\"tuple\",\"internalType\":\"structTokenInfo\",\"components\":[{\"name\":\"erc\",\"type\":\"uint8\",\"internalType\":\"enumTokenStandard\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ErrInvalidTrustedThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrWithdrawalsMigrated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrWithdrawnOnMainchainAlready\",\"inputs\":[]}]",
}

// RoninGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use RoninGatewayMetaData.ABI instead.
var RoninGatewayABI = RoninGatewayMetaData.ABI

// RoninGateway is an auto generated Go binding around an Ethereum contract.
type RoninGateway struct {
	RoninGatewayCaller     // Read-only binding to the contract
	RoninGatewayTransactor // Write-only binding to the contract
	RoninGatewayFilterer   // Log filterer for contract events
}

// RoninGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type RoninGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoninGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RoninGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoninGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RoninGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoninGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RoninGatewaySession struct {
	Contract     *RoninGateway     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RoninGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RoninGatewayCallerSession struct {
	Contract *RoninGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RoninGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RoninGatewayTransactorSession struct {
	Contract     *RoninGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RoninGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type RoninGatewayRaw struct {
	Contract *RoninGateway // Generic contract binding to access the raw methods on
}

// RoninGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RoninGatewayCallerRaw struct {
	Contract *RoninGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// RoninGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RoninGatewayTransactorRaw struct {
	Contract *RoninGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoninGateway creates a new instance of RoninGateway, bound to a specific deployed contract.
func NewRoninGateway(address common.Address, backend bind.ContractBackend) (*RoninGateway, error) {
	contract, err := bindRoninGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RoninGateway{RoninGatewayCaller: RoninGatewayCaller{contract: contract}, RoninGatewayTransactor: RoninGatewayTransactor{contract: contract}, RoninGatewayFilterer: RoninGatewayFilterer{contract: contract}}, nil
}

// NewRoninGatewayCaller creates a new read-only instance of RoninGateway, bound to a specific deployed contract.
func NewRoninGatewayCaller(address common.Address, caller bind.ContractCaller) (*RoninGatewayCaller, error) {
	contract, err := bindRoninGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RoninGatewayCaller{contract: contract}, nil
}

// NewRoninGatewayTransactor creates a new write-only instance of RoninGateway, bound to a specific deployed contract.
func NewRoninGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*RoninGatewayTransactor, error) {
	contract, err := bindRoninGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RoninGatewayTransactor{contract: contract}, nil
}

// NewRoninGatewayFilterer creates a new log filterer instance of RoninGateway, bound to a specific deployed contract.
func NewRoninGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*RoninGatewayFilterer, error) {
	contract, err := bindRoninGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RoninGatewayFilterer{contract: contract}, nil
}

// bindRoninGateway binds a generic wrapper to an already deployed contract.
func bindRoninGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RoninGatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoninGateway *RoninGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoninGateway.Contract.RoninGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoninGateway *RoninGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoninGateway.Contract.RoninGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoninGateway *RoninGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoninGateway.Contract.RoninGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoninGateway *RoninGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoninGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoninGateway *RoninGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoninGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoninGateway *RoninGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoninGateway.Contract.contract.Transact(opts, method, params...)
}

// DepositVote is a free data retrieval call binding the contract method 0x4d92c4f0.
//
// Solidity: function depositVote(uint256 , uint256 ) view returns(uint8 status, bytes32 finalHash, uint256 expiredAt, uint256 createdAt)
func (_RoninGateway *RoninGatewayCaller) DepositVote(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Status    uint8
	FinalHash [32]byte
	ExpiredAt *big.Int
	CreatedAt *big.Int
}, error) {
	var out []interface{}
	err := _RoninGateway.contract.Call(opts, &out, "depositVote", arg0, arg1)

	outstruct := new(struct {
		Status    uint8
		FinalHash [32]byte
		ExpiredAt *big.Int
		CreatedAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.FinalHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.ExpiredAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CreatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DepositVote is a free data retrieval call binding the contract method 0x4d92c4f0.
//
// Solidity: function depositVote(uint256 , uint256 ) view returns(uint8 status, bytes32 finalHash, uint256 expiredAt, uint256 createdAt)
func (_RoninGateway *RoninGatewaySession) DepositVote(arg0 *big.Int, arg1 *big.Int) (struct {
	Status    uint8
	FinalHash [32]byte
	ExpiredAt *big.Int
	CreatedAt *big.Int
}, error) {
	return _RoninGateway.Contract.DepositVote(&_RoninGateway.CallOpts, arg0, arg1)
}

// DepositVote is a free data retrieval call binding the contract method 0x4d92c4f0.
//
// Solidity: function depositVote(uint256 , uint256 ) view returns(uint8 status, bytes32 finalHash, uint256 expiredAt, uint256 createdAt)
func (_RoninGateway *RoninGatewayCallerSession) DepositVote(arg0 *big.Int, arg1 *big.Int) (struct {
	Status    uint8
	FinalHash [32]byte
	ExpiredAt *big.Int
	CreatedAt *big.Int
}, error) {
	return _RoninGateway.Contract.DepositVote(&_RoninGateway.CallOpts, arg0, arg1)
}

// DepositVoted is a free data retrieval call binding the contract method 0xfc6574bc.
//
// Solidity: function depositVoted(uint256 _chainId, uint256 _depositId, address _voter) view returns(bool)
func (_RoninGateway *RoninGatewayCaller) DepositVoted(opts *bind.CallOpts, _chainId *big.Int, _depositId *big.Int, _voter common.Address) (bool, error) {
	var out []interface{}
	err := _RoninGateway.contract.Call(opts, &out, "depositVoted", _chainId, _depositId, _voter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DepositVoted is a free data retrieval call binding the contract method 0xfc6574bc.
//
// Solidity: function depositVoted(uint256 _chainId, uint256 _depositId, address _voter) view returns(bool)
func (_RoninGateway *RoninGatewaySession) DepositVoted(_chainId *big.Int, _depositId *big.Int, _voter common.Address) (bool, error) {
	return _RoninGateway.Contract.DepositVoted(&_RoninGateway.CallOpts, _chainId, _depositId, _voter)
}

// DepositVoted is a free data retrieval call binding the contract method 0xfc6574bc.
//
// Solidity: function depositVoted(uint256 _chainId, uint256 _depositId, address _voter) view returns(bool)
func (_RoninGateway *RoninGatewayCallerSession) DepositVoted(_chainId *big.Int, _depositId *big.Int, _voter common.Address) (bool, error) {
	return _RoninGateway.Contract.DepositVoted(&_RoninGateway.CallOpts, _chainId, _depositId, _voter)
}

// GetMainchainToken is a free data retrieval call binding the contract method 0x5d6a9a90.
//
// Solidity: function getMainchainToken(address _roninToken, uint256 _chainId) view returns((uint8,address) _token)
func (_RoninGateway *RoninGatewayCaller) GetMainchainToken(opts *bind.CallOpts, _roninToken common.Address, _chainId *big.Int) (MappedTokenConsumerMappedToken, error) {
	var out []interface{}
	err := _RoninGateway.contract.Call(opts, &out, "getMainchainToken", _roninToken, _chainId)

	if err != nil {
		return *new(MappedTokenConsumerMappedToken), err
	}

	out0 := *abi.ConvertType(out[0], new(MappedTokenConsumerMappedToken)).(*MappedTokenConsumerMappedToken)

	return out0, err

}

// GetMainchainToken is a free data retrieval call binding the contract method 0x5d6a9a90.
//
// Solidity: function getMainchainToken(address _roninToken, uint256 _chainId) view returns((uint8,address) _token)
func (_RoninGateway *RoninGatewaySession) GetMainchainToken(_roninToken common.Address, _chainId *big.Int) (MappedTokenConsumerMappedToken, error) {
	return _RoninGateway.Contract.GetMainchainToken(&_RoninGateway.CallOpts, _roninToken, _chainId)
}

// GetMainchainToken is a free data retrieval call binding the contract method 0x5d6a9a90.
//
// Solidity: function getMainchainToken(address _roninToken, uint256 _chainId) view returns((uint8,address) _token)
func (_RoninGateway *RoninGatewayCallerSession) GetMainchainToken(_roninToken common.Address, _chainId *big.Int) (MappedTokenConsumerMappedToken, error) {
	return _RoninGateway.Contract.GetMainchainToken(&_RoninGateway.CallOpts, _roninToken, _chainId)
}

// GetWithdrawalSignatures is a free data retrieval call binding the contract method 0xecc83649.
//
// Solidity: function getWithdrawalSignatures(uint256 _withdrawalId, address[] _validators) view returns(bytes[])
func (_RoninGateway *RoninGatewayCaller) GetWithdrawalSignatures(opts *bind.CallOpts, _withdrawalId *big.Int, _validators []common.Address) ([][]byte, error) {
	var out []interface{}
	err := _RoninGateway.contract.Call(opts, &out, "getWithdrawalSignatures", _withdrawalId, _validators)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetWithdrawalSignatures is a free data retrieval call binding the contract method 0xecc83649.
//
// Solidity: function getWithdrawalSignatures(uint256 _withdrawalId, address[] _validators) view returns(bytes[])
func (_RoninGateway *RoninGatewaySession) GetWithdrawalSignatures(_withdrawalId *big.Int, _validators []common.Address) ([][]byte, error) {
	return _RoninGateway.Contract.GetWithdrawalSignatures(&_RoninGateway.CallOpts, _withdrawalId, _validators)
}

// GetWithdrawalSignatures is a free data retrieval call binding the contract method 0xecc83649.
//
// Solidity: function getWithdrawalSignatures(uint256 _withdrawalId, address[] _validators) view returns(bytes[])
func (_RoninGateway *RoninGatewayCallerSession) GetWithdrawalSignatures(_withdrawalId *big.Int, _validators []common.Address) ([][]byte, error) {
	return _RoninGateway.Contract.GetWithdrawalSignatures(&_RoninGateway.CallOpts, _withdrawalId, _validators)
}

// MainchainWithdrew is a free data retrieval call binding the contract method 0xf668214a.
//
// Solidity: function mainchainWithdrew(uint256 _withdrawalId) view returns(bool)
func (_RoninGateway *RoninGatewayCaller) MainchainWithdrew(opts *bind.CallOpts, _withdrawalId *big.Int) (bool, error) {
	var out []interface{}
	err := _RoninGateway.contract.Call(opts, &out, "mainchainWithdrew", _withdrawalId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MainchainWithdrew is a free data retrieval call binding the contract method 0xf668214a.
//
// Solidity: function mainchainWithdrew(uint256 _withdrawalId) view returns(bool)
func (_RoninGateway *RoninGatewaySession) MainchainWithdrew(_withdrawalId *big.Int) (bool, error) {
	return _RoninGateway.Contract.MainchainWithdrew(&_RoninGateway.CallOpts, _withdrawalId)
}

// MainchainWithdrew is a free data retrieval call binding the contract method 0xf668214a.
//
// Solidity: function mainchainWithdrew(uint256 _withdrawalId) view returns(bool)
func (_RoninGateway *RoninGatewayCallerSession) MainchainWithdrew(_withdrawalId *big.Int) (bool, error) {
	return _RoninGateway.Contract.MainchainWithdrew(&_RoninGateway.CallOpts, _withdrawalId)
}

// MainchainWithdrewVoted is a free data retrieval call binding the contract method 0x3e4574ec.
//
// Solidity: function mainchainWithdrewVoted(uint256 _withdrawalId, address _voter) view returns(bool)
func (_RoninGateway *RoninGatewayCaller) MainchainWithdrewVoted(opts *bind.CallOpts, _withdrawalId *big.Int, _voter common.Address) (bool, error) {
	var out []interface{}
	err := _RoninGateway.contract.Call(opts, &out, "mainchainWithdrewVoted", _withdrawalId, _voter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MainchainWithdrewVoted is a free data retrieval call binding the contract method 0x3e4574ec.
//
// Solidity: function mainchainWithdrewVoted(uint256 _withdrawalId, address _voter) view returns(bool)
func (_RoninGateway *RoninGatewaySession) MainchainWithdrewVoted(_withdrawalId *big.Int, _voter common.Address) (bool, error) {
	return _RoninGateway.Contract.MainchainWithdrewVoted(&_RoninGateway.CallOpts, _withdrawalId, _voter)
}

// MainchainWithdrewVoted is a free data retrieval call binding the contract method 0x3e4574ec.
//
// Solidity: function mainchainWithdrewVoted(uint256 _withdrawalId, address _voter) view returns(bool)
func (_RoninGateway *RoninGatewayCallerSession) MainchainWithdrewVoted(_withdrawalId *big.Int, _voter common.Address) (bool, error) {
	return _RoninGateway.Contract.MainchainWithdrewVoted(&_RoninGateway.CallOpts, _withdrawalId, _voter)
}

// WithdrawalCount is a free data retrieval call binding the contract method 0x71706cbe.
//
// Solidity: function withdrawalCount() view returns(uint256)
func (_RoninGateway *RoninGatewayCaller) WithdrawalCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninGateway.contract.Call(opts, &out, "withdrawalCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalCount is a free data retrieval call binding the contract method 0x71706cbe.
//
// Solidity: function withdrawalCount() view returns(uint256)
func (_RoninGateway *RoninGatewaySession) WithdrawalCount() (*big.Int, error) {
	return _RoninGateway.Contract.WithdrawalCount(&_RoninGateway.CallOpts)
}

// WithdrawalCount is a free data retrieval call binding the contract method 0x71706cbe.
//
// Solidity: function withdrawalCount() view returns(uint256)
func (_RoninGateway *RoninGatewayCallerSession) WithdrawalCount() (*big.Int, error) {
	return _RoninGateway.Contract.WithdrawalCount(&_RoninGateway.CallOpts)
}

// BulkRequestWithdrawalFor is a paid mutator transaction binding the contract method 0x5a7dd06a.
//
// Solidity: function bulkRequestWithdrawalFor((address,address,(uint8,uint256,uint256))[] _requests, uint256 _chainId) returns()
func (_RoninGateway *RoninGatewayTransactor) BulkRequestWithdrawalFor(opts *bind.TransactOpts, _requests []TransferRequest, _chainId *big.Int) (*types.Transaction, error) {
	return _RoninGateway.contract.Transact(opts, "bulkRequestWithdrawalFor", _requests, _chainId)
}

// BulkRequestWithdrawalFor is a paid mutator transaction binding the contract method 0x5a7dd06a.
//
// Solidity: function bulkRequestWithdrawalFor((address,address,(uint8,uint256,uint256))[] _requests, uint256 _chainId) returns()
func (_RoninGateway *RoninGatewaySession) BulkRequestWithdrawalFor(_requests []TransferRequest, _chainId *big.Int) (*types.Transaction, error) {
	return _RoninGateway.Contract.BulkRequestWithdrawalFor(&_RoninGateway.TransactOpts, _requests, _chainId)
}

// BulkRequestWithdrawalFor is a paid mutator transaction binding the contract method 0x5a7dd06a.
//
// Solidity: function bulkRequestWithdrawalFor((address,address,(uint8,uint256,uint256))[] _requests, uint256 _chainId) returns()
func (_RoninGateway *RoninGatewayTransactorSession) BulkRequestWithdrawalFor(_requests []TransferRequest, _chainId *big.Int) (*types.Transaction, error) {
	return _RoninGateway.Contract.BulkRequestWithdrawalFor(&_RoninGateway.TransactOpts, _requests, _chainId)
}

// BulkSubmitWithdrawalSignatures is a paid mutator transaction binding the contract method 0xfa389659.
//
// Solidity: function bulkSubmitWithdrawalSignatures(uint256[] _withdrawals, bytes[] _signatures) returns()
func (_RoninGateway *RoninGatewayTransactor) BulkSubmitWithdrawalSignatures(opts *bind.TransactOpts, _withdrawals []*big.Int, _signatures [][]byte) (*types.Transaction, error) {
	return _RoninGateway.contract.Transact(opts, "bulkSubmitWithdrawalSignatures", _withdrawals, _signatures)
}

// BulkSubmitWithdrawalSignatures is a paid mutator transaction binding the contract method 0xfa389659.
//
// Solidity: function bulkSubmitWithdrawalSignatures(uint256[] _withdrawals, bytes[] _signatures) returns()
func (_RoninGateway *RoninGatewaySession) BulkSubmitWithdrawalSignatures(_withdrawals []*big.Int, _signatures [][]byte) (*types.Transaction, error) {
	return _RoninGateway.Contract.BulkSubmitWithdrawalSignatures(&_RoninGateway.TransactOpts, _withdrawals, _signatures)
}

// BulkSubmitWithdrawalSignatures is a paid mutator transaction binding the contract method 0xfa389659.
//
// Solidity: function bulkSubmitWithdrawalSignatures(uint256[] _withdrawals, bytes[] _signatures) returns()
func (_RoninGateway *RoninGatewayTransactorSession) BulkSubmitWithdrawalSignatures(_withdrawals []*big.Int, _signatures [][]byte) (*types.Transaction, error) {
	return _RoninGateway.Contract.BulkSubmitWithdrawalSignatures(&_RoninGateway.TransactOpts, _withdrawals, _signatures)
}

// DepositFor is a paid mutator transaction binding the contract method 0x109679ef.
//
// Solidity: function depositFor((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) _receipt) returns()
func (_RoninGateway *RoninGatewayTransactor) DepositFor(opts *bind.TransactOpts, _receipt TransferReceipt) (*types.Transaction, error) {
	return _RoninGateway.contract.Transact(opts, "depositFor", _receipt)
}

// DepositFor is a paid mutator transaction binding the contract method 0x109679ef.
//
// Solidity: function depositFor((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) _receipt) returns()
func (_RoninGateway *RoninGatewaySession) DepositFor(_receipt TransferReceipt) (*types.Transaction, error) {
	return _RoninGateway.Contract.DepositFor(&_RoninGateway.TransactOpts, _receipt)
}

// DepositFor is a paid mutator transaction binding the contract method 0x109679ef.
//
// Solidity: function depositFor((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) _receipt) returns()
func (_RoninGateway *RoninGatewayTransactorSession) DepositFor(_receipt TransferReceipt) (*types.Transaction, error) {
	return _RoninGateway.Contract.DepositFor(&_RoninGateway.TransactOpts, _receipt)
}

// MapTokens is a paid mutator transaction binding the contract method 0xdbd2ef6c.
//
// Solidity: function mapTokens(address[] _roninTokens, address[] _mainchainTokens, uint256[] chainIds, uint8[] _standards) returns()
func (_RoninGateway *RoninGatewayTransactor) MapTokens(opts *bind.TransactOpts, _roninTokens []common.Address, _mainchainTokens []common.Address, chainIds []*big.Int, _standards []uint8) (*types.Transaction, error) {
	return _RoninGateway.contract.Transact(opts, "mapTokens", _roninTokens, _mainchainTokens, chainIds, _standards)
}

// MapTokens is a paid mutator transaction binding the contract method 0xdbd2ef6c.
//
// Solidity: function mapTokens(address[] _roninTokens, address[] _mainchainTokens, uint256[] chainIds, uint8[] _standards) returns()
func (_RoninGateway *RoninGatewaySession) MapTokens(_roninTokens []common.Address, _mainchainTokens []common.Address, chainIds []*big.Int, _standards []uint8) (*types.Transaction, error) {
	return _RoninGateway.Contract.MapTokens(&_RoninGateway.TransactOpts, _roninTokens, _mainchainTokens, chainIds, _standards)
}

// MapTokens is a paid mutator transaction binding the contract method 0xdbd2ef6c.
//
// Solidity: function mapTokens(address[] _roninTokens, address[] _mainchainTokens, uint256[] chainIds, uint8[] _standards) returns()
func (_RoninGateway *RoninGatewayTransactorSession) MapTokens(_roninTokens []common.Address, _mainchainTokens []common.Address, chainIds []*big.Int, _standards []uint8) (*types.Transaction, error) {
	return _RoninGateway.Contract.MapTokens(&_RoninGateway.TransactOpts, _roninTokens, _mainchainTokens, chainIds, _standards)
}

// MapTokensWithMinThresholds is a paid mutator transaction binding the contract method 0x49cae6b5.
//
// Solidity: function mapTokensWithMinThresholds(address[] roninTokens_, address[] mainchainTokens_, uint256[] chainIds_, uint8[] standards_, uint256[] minimumThresholds_) returns()
func (_RoninGateway *RoninGatewayTransactor) MapTokensWithMinThresholds(opts *bind.TransactOpts, roninTokens_ []common.Address, mainchainTokens_ []common.Address, chainIds_ []*big.Int, standards_ []uint8, minimumThresholds_ []*big.Int) (*types.Transaction, error) {
	return _RoninGateway.contract.Transact(opts, "mapTokensWithMinThresholds", roninTokens_, mainchainTokens_, chainIds_, standards_, minimumThresholds_)
}

// MapTokensWithMinThresholds is a paid mutator transaction binding the contract method 0x49cae6b5.
//
// Solidity: function mapTokensWithMinThresholds(address[] roninTokens_, address[] mainchainTokens_, uint256[] chainIds_, uint8[] standards_, uint256[] minimumThresholds_) returns()
func (_RoninGateway *RoninGatewaySession) MapTokensWithMinThresholds(roninTokens_ []common.Address, mainchainTokens_ []common.Address, chainIds_ []*big.Int, standards_ []uint8, minimumThresholds_ []*big.Int) (*types.Transaction, error) {
	return _RoninGateway.Contract.MapTokensWithMinThresholds(&_RoninGateway.TransactOpts, roninTokens_, mainchainTokens_, chainIds_, standards_, minimumThresholds_)
}

// MapTokensWithMinThresholds is a paid mutator transaction binding the contract method 0x49cae6b5.
//
// Solidity: function mapTokensWithMinThresholds(address[] roninTokens_, address[] mainchainTokens_, uint256[] chainIds_, uint8[] standards_, uint256[] minimumThresholds_) returns()
func (_RoninGateway *RoninGatewayTransactorSession) MapTokensWithMinThresholds(roninTokens_ []common.Address, mainchainTokens_ []common.Address, chainIds_ []*big.Int, standards_ []uint8, minimumThresholds_ []*big.Int) (*types.Transaction, error) {
	return _RoninGateway.Contract.MapTokensWithMinThresholds(&_RoninGateway.TransactOpts, roninTokens_, mainchainTokens_, chainIds_, standards_, minimumThresholds_)
}

// RequestWithdrawalFor is a paid mutator transaction binding the contract method 0x0b1ff17f.
//
// Solidity: function requestWithdrawalFor((address,address,(uint8,uint256,uint256)) _request, uint256 _chainId) returns()
func (_RoninGateway *RoninGatewayTransactor) RequestWithdrawalFor(opts *bind.TransactOpts, _request TransferRequest, _chainId *big.Int) (*types.Transaction, error) {
	return _RoninGateway.contract.Transact(opts, "requestWithdrawalFor", _request, _chainId)
}

// RequestWithdrawalFor is a paid mutator transaction binding the contract method 0x0b1ff17f.
//
// Solidity: function requestWithdrawalFor((address,address,(uint8,uint256,uint256)) _request, uint256 _chainId) returns()
func (_RoninGateway *RoninGatewaySession) RequestWithdrawalFor(_request TransferRequest, _chainId *big.Int) (*types.Transaction, error) {
	return _RoninGateway.Contract.RequestWithdrawalFor(&_RoninGateway.TransactOpts, _request, _chainId)
}

// RequestWithdrawalFor is a paid mutator transaction binding the contract method 0x0b1ff17f.
//
// Solidity: function requestWithdrawalFor((address,address,(uint8,uint256,uint256)) _request, uint256 _chainId) returns()
func (_RoninGateway *RoninGatewayTransactorSession) RequestWithdrawalFor(_request TransferRequest, _chainId *big.Int) (*types.Transaction, error) {
	return _RoninGateway.Contract.RequestWithdrawalFor(&_RoninGateway.TransactOpts, _request, _chainId)
}

// RequestWithdrawalSignatures is a paid mutator transaction binding the contract method 0x47b56b2c.
//
// Solidity: function requestWithdrawalSignatures(uint256 _withdrawalId) returns()
func (_RoninGateway *RoninGatewayTransactor) RequestWithdrawalSignatures(opts *bind.TransactOpts, _withdrawalId *big.Int) (*types.Transaction, error) {
	return _RoninGateway.contract.Transact(opts, "requestWithdrawalSignatures", _withdrawalId)
}

// RequestWithdrawalSignatures is a paid mutator transaction binding the contract method 0x47b56b2c.
//
// Solidity: function requestWithdrawalSignatures(uint256 _withdrawalId) returns()
func (_RoninGateway *RoninGatewaySession) RequestWithdrawalSignatures(_withdrawalId *big.Int) (*types.Transaction, error) {
	return _RoninGateway.Contract.RequestWithdrawalSignatures(&_RoninGateway.TransactOpts, _withdrawalId)
}

// RequestWithdrawalSignatures is a paid mutator transaction binding the contract method 0x47b56b2c.
//
// Solidity: function requestWithdrawalSignatures(uint256 _withdrawalId) returns()
func (_RoninGateway *RoninGatewayTransactorSession) RequestWithdrawalSignatures(_withdrawalId *big.Int) (*types.Transaction, error) {
	return _RoninGateway.Contract.RequestWithdrawalSignatures(&_RoninGateway.TransactOpts, _withdrawalId)
}

// TryBulkAcknowledgeMainchainWithdrew is a paid mutator transaction binding the contract method 0x17fa2ea1.
//
// Solidity: function tryBulkAcknowledgeMainchainWithdrew(uint256[] _withdrawalIds) returns(bool[])
func (_RoninGateway *RoninGatewayTransactor) TryBulkAcknowledgeMainchainWithdrew(opts *bind.TransactOpts, _withdrawalIds []*big.Int) (*types.Transaction, error) {
	return _RoninGateway.contract.Transact(opts, "tryBulkAcknowledgeMainchainWithdrew", _withdrawalIds)
}

// TryBulkAcknowledgeMainchainWithdrew is a paid mutator transaction binding the contract method 0x17fa2ea1.
//
// Solidity: function tryBulkAcknowledgeMainchainWithdrew(uint256[] _withdrawalIds) returns(bool[])
func (_RoninGateway *RoninGatewaySession) TryBulkAcknowledgeMainchainWithdrew(_withdrawalIds []*big.Int) (*types.Transaction, error) {
	return _RoninGateway.Contract.TryBulkAcknowledgeMainchainWithdrew(&_RoninGateway.TransactOpts, _withdrawalIds)
}

// TryBulkAcknowledgeMainchainWithdrew is a paid mutator transaction binding the contract method 0x17fa2ea1.
//
// Solidity: function tryBulkAcknowledgeMainchainWithdrew(uint256[] _withdrawalIds) returns(bool[])
func (_RoninGateway *RoninGatewayTransactorSession) TryBulkAcknowledgeMainchainWithdrew(_withdrawalIds []*big.Int) (*types.Transaction, error) {
	return _RoninGateway.Contract.TryBulkAcknowledgeMainchainWithdrew(&_RoninGateway.TransactOpts, _withdrawalIds)
}

// TryBulkDepositFor is a paid mutator transaction binding the contract method 0xb9afa177.
//
// Solidity: function tryBulkDepositFor((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256))[] _receipts) returns(bool[])
func (_RoninGateway *RoninGatewayTransactor) TryBulkDepositFor(opts *bind.TransactOpts, _receipts []TransferReceipt) (*types.Transaction, error) {
	return _RoninGateway.contract.Transact(opts, "tryBulkDepositFor", _receipts)
}

// TryBulkDepositFor is a paid mutator transaction binding the contract method 0xb9afa177.
//
// Solidity: function tryBulkDepositFor((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256))[] _receipts) returns(bool[])
func (_RoninGateway *RoninGatewaySession) TryBulkDepositFor(_receipts []TransferReceipt) (*types.Transaction, error) {
	return _RoninGateway.Contract.TryBulkDepositFor(&_RoninGateway.TransactOpts, _receipts)
}

// TryBulkDepositFor is a paid mutator transaction binding the contract method 0xb9afa177.
//
// Solidity: function tryBulkDepositFor((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256))[] _receipts) returns(bool[])
func (_RoninGateway *RoninGatewayTransactorSession) TryBulkDepositFor(_receipts []TransferReceipt) (*types.Transaction, error) {
	return _RoninGateway.Contract.TryBulkDepositFor(&_RoninGateway.TransactOpts, _receipts)
}

// UnmapTokens is a paid mutator transaction binding the contract method 0x48e6a22c.
//
// Solidity: function unmapTokens(address[] roninTokens_, uint256[] chainIds_) returns()
func (_RoninGateway *RoninGatewayTransactor) UnmapTokens(opts *bind.TransactOpts, roninTokens_ []common.Address, chainIds_ []*big.Int) (*types.Transaction, error) {
	return _RoninGateway.contract.Transact(opts, "unmapTokens", roninTokens_, chainIds_)
}

// UnmapTokens is a paid mutator transaction binding the contract method 0x48e6a22c.
//
// Solidity: function unmapTokens(address[] roninTokens_, uint256[] chainIds_) returns()
func (_RoninGateway *RoninGatewaySession) UnmapTokens(roninTokens_ []common.Address, chainIds_ []*big.Int) (*types.Transaction, error) {
	return _RoninGateway.Contract.UnmapTokens(&_RoninGateway.TransactOpts, roninTokens_, chainIds_)
}

// UnmapTokens is a paid mutator transaction binding the contract method 0x48e6a22c.
//
// Solidity: function unmapTokens(address[] roninTokens_, uint256[] chainIds_) returns()
func (_RoninGateway *RoninGatewayTransactorSession) UnmapTokens(roninTokens_ []common.Address, chainIds_ []*big.Int) (*types.Transaction, error) {
	return _RoninGateway.Contract.UnmapTokens(&_RoninGateway.TransactOpts, roninTokens_, chainIds_)
}

// RoninGatewayDepositVotedIterator is returned from FilterDepositVoted and is used to iterate over the raw logs and unpacked data for DepositVoted events raised by the RoninGateway contract.
type RoninGatewayDepositVotedIterator struct {
	Event *RoninGatewayDepositVoted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RoninGatewayDepositVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninGatewayDepositVoted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RoninGatewayDepositVoted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RoninGatewayDepositVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninGatewayDepositVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninGatewayDepositVoted represents a DepositVoted event raised by the RoninGateway contract.
type RoninGatewayDepositVoted struct {
	BridgeOperator common.Address
	Id             *big.Int
	ChainId        *big.Int
	ReceiptHash    [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDepositVoted is a free log retrieval operation binding the contract event 0x48c4262ed68beb92fe5d7d48d70772e49cd50c317937dea60a99f15f794b6459.
//
// Solidity: event DepositVoted(address indexed bridgeOperator, uint256 indexed id, uint256 indexed chainId, bytes32 receiptHash)
func (_RoninGateway *RoninGatewayFilterer) FilterDepositVoted(opts *bind.FilterOpts, bridgeOperator []common.Address, id []*big.Int, chainId []*big.Int) (*RoninGatewayDepositVotedIterator, error) {

	var bridgeOperatorRule []interface{}
	for _, bridgeOperatorItem := range bridgeOperator {
		bridgeOperatorRule = append(bridgeOperatorRule, bridgeOperatorItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _RoninGateway.contract.FilterLogs(opts, "DepositVoted", bridgeOperatorRule, idRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &RoninGatewayDepositVotedIterator{contract: _RoninGateway.contract, event: "DepositVoted", logs: logs, sub: sub}, nil
}

// WatchDepositVoted is a free log subscription operation binding the contract event 0x48c4262ed68beb92fe5d7d48d70772e49cd50c317937dea60a99f15f794b6459.
//
// Solidity: event DepositVoted(address indexed bridgeOperator, uint256 indexed id, uint256 indexed chainId, bytes32 receiptHash)
func (_RoninGateway *RoninGatewayFilterer) WatchDepositVoted(opts *bind.WatchOpts, sink chan<- *RoninGatewayDepositVoted, bridgeOperator []common.Address, id []*big.Int, chainId []*big.Int) (event.Subscription, error) {

	var bridgeOperatorRule []interface{}
	for _, bridgeOperatorItem := range bridgeOperator {
		bridgeOperatorRule = append(bridgeOperatorRule, bridgeOperatorItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _RoninGateway.contract.WatchLogs(opts, "DepositVoted", bridgeOperatorRule, idRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninGatewayDepositVoted)
				if err := _RoninGateway.contract.UnpackLog(event, "DepositVoted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositVoted is a log parse operation binding the contract event 0x48c4262ed68beb92fe5d7d48d70772e49cd50c317937dea60a99f15f794b6459.
//
// Solidity: event DepositVoted(address indexed bridgeOperator, uint256 indexed id, uint256 indexed chainId, bytes32 receiptHash)
func (_RoninGateway *RoninGatewayFilterer) ParseDepositVoted(log types.Log) (*RoninGatewayDepositVoted, error) {
	event := new(RoninGatewayDepositVoted)
	if err := _RoninGateway.contract.UnpackLog(event, "DepositVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninGatewayDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the RoninGateway contract.
type RoninGatewayDepositedIterator struct {
	Event *RoninGatewayDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RoninGatewayDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninGatewayDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RoninGatewayDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RoninGatewayDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninGatewayDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninGatewayDeposited represents a Deposited event raised by the RoninGateway contract.
type RoninGatewayDeposited struct {
	ReceiptHash [32]byte
	Receipt     TransferReceipt
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x8d20d8121a34dded9035ff5b43e901c142824f7a22126392992c353c37890524.
//
// Solidity: event Deposited(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_RoninGateway *RoninGatewayFilterer) FilterDeposited(opts *bind.FilterOpts) (*RoninGatewayDepositedIterator, error) {

	logs, sub, err := _RoninGateway.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &RoninGatewayDepositedIterator{contract: _RoninGateway.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x8d20d8121a34dded9035ff5b43e901c142824f7a22126392992c353c37890524.
//
// Solidity: event Deposited(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_RoninGateway *RoninGatewayFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *RoninGatewayDeposited) (event.Subscription, error) {

	logs, sub, err := _RoninGateway.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninGatewayDeposited)
				if err := _RoninGateway.contract.UnpackLog(event, "Deposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposited is a log parse operation binding the contract event 0x8d20d8121a34dded9035ff5b43e901c142824f7a22126392992c353c37890524.
//
// Solidity: event Deposited(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_RoninGateway *RoninGatewayFilterer) ParseDeposited(log types.Log) (*RoninGatewayDeposited, error) {
	event := new(RoninGatewayDeposited)
	if err := _RoninGateway.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninGatewayMainchainWithdrewIterator is returned from FilterMainchainWithdrew and is used to iterate over the raw logs and unpacked data for MainchainWithdrew events raised by the RoninGateway contract.
type RoninGatewayMainchainWithdrewIterator struct {
	Event *RoninGatewayMainchainWithdrew // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RoninGatewayMainchainWithdrewIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninGatewayMainchainWithdrew)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RoninGatewayMainchainWithdrew)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RoninGatewayMainchainWithdrewIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninGatewayMainchainWithdrewIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninGatewayMainchainWithdrew represents a MainchainWithdrew event raised by the RoninGateway contract.
type RoninGatewayMainchainWithdrew struct {
	ReceiptHash [32]byte
	Receipt     TransferReceipt
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMainchainWithdrew is a free log retrieval operation binding the contract event 0x62520d049932cdee872e9b3c59c0f6073637147e5e9bc8b050b062430eaf5c9f.
//
// Solidity: event MainchainWithdrew(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_RoninGateway *RoninGatewayFilterer) FilterMainchainWithdrew(opts *bind.FilterOpts) (*RoninGatewayMainchainWithdrewIterator, error) {

	logs, sub, err := _RoninGateway.contract.FilterLogs(opts, "MainchainWithdrew")
	if err != nil {
		return nil, err
	}
	return &RoninGatewayMainchainWithdrewIterator{contract: _RoninGateway.contract, event: "MainchainWithdrew", logs: logs, sub: sub}, nil
}

// WatchMainchainWithdrew is a free log subscription operation binding the contract event 0x62520d049932cdee872e9b3c59c0f6073637147e5e9bc8b050b062430eaf5c9f.
//
// Solidity: event MainchainWithdrew(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_RoninGateway *RoninGatewayFilterer) WatchMainchainWithdrew(opts *bind.WatchOpts, sink chan<- *RoninGatewayMainchainWithdrew) (event.Subscription, error) {

	logs, sub, err := _RoninGateway.contract.WatchLogs(opts, "MainchainWithdrew")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninGatewayMainchainWithdrew)
				if err := _RoninGateway.contract.UnpackLog(event, "MainchainWithdrew", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMainchainWithdrew is a log parse operation binding the contract event 0x62520d049932cdee872e9b3c59c0f6073637147e5e9bc8b050b062430eaf5c9f.
//
// Solidity: event MainchainWithdrew(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_RoninGateway *RoninGatewayFilterer) ParseMainchainWithdrew(log types.Log) (*RoninGatewayMainchainWithdrew, error) {
	event := new(RoninGatewayMainchainWithdrew)
	if err := _RoninGateway.contract.UnpackLog(event, "MainchainWithdrew", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninGatewayTokenMappedIterator is returned from FilterTokenMapped and is used to iterate over the raw logs and unpacked data for TokenMapped events raised by the RoninGateway contract.
type RoninGatewayTokenMappedIterator struct {
	Event *RoninGatewayTokenMapped // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RoninGatewayTokenMappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninGatewayTokenMapped)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RoninGatewayTokenMapped)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RoninGatewayTokenMappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninGatewayTokenMappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninGatewayTokenMapped represents a TokenMapped event raised by the RoninGateway contract.
type RoninGatewayTokenMapped struct {
	RoninTokens     []common.Address
	MainchainTokens []common.Address
	ChainIds        []*big.Int
	Standards       []uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTokenMapped is a free log retrieval operation binding the contract event 0x2544bff60c6d5b84946e06804af9f84e150bbee85238dbdee79efca4e0adf401.
//
// Solidity: event TokenMapped(address[] roninTokens, address[] mainchainTokens, uint256[] chainIds, uint8[] standards)
func (_RoninGateway *RoninGatewayFilterer) FilterTokenMapped(opts *bind.FilterOpts) (*RoninGatewayTokenMappedIterator, error) {

	logs, sub, err := _RoninGateway.contract.FilterLogs(opts, "TokenMapped")
	if err != nil {
		return nil, err
	}
	return &RoninGatewayTokenMappedIterator{contract: _RoninGateway.contract, event: "TokenMapped", logs: logs, sub: sub}, nil
}

// WatchTokenMapped is a free log subscription operation binding the contract event 0x2544bff60c6d5b84946e06804af9f84e150bbee85238dbdee79efca4e0adf401.
//
// Solidity: event TokenMapped(address[] roninTokens, address[] mainchainTokens, uint256[] chainIds, uint8[] standards)
func (_RoninGateway *RoninGatewayFilterer) WatchTokenMapped(opts *bind.WatchOpts, sink chan<- *RoninGatewayTokenMapped) (event.Subscription, error) {

	logs, sub, err := _RoninGateway.contract.WatchLogs(opts, "TokenMapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninGatewayTokenMapped)
				if err := _RoninGateway.contract.UnpackLog(event, "TokenMapped", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenMapped is a log parse operation binding the contract event 0x2544bff60c6d5b84946e06804af9f84e150bbee85238dbdee79efca4e0adf401.
//
// Solidity: event TokenMapped(address[] roninTokens, address[] mainchainTokens, uint256[] chainIds, uint8[] standards)
func (_RoninGateway *RoninGatewayFilterer) ParseTokenMapped(log types.Log) (*RoninGatewayTokenMapped, error) {
	event := new(RoninGatewayTokenMapped)
	if err := _RoninGateway.contract.UnpackLog(event, "TokenMapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninGatewayTokenUnmappedIterator is returned from FilterTokenUnmapped and is used to iterate over the raw logs and unpacked data for TokenUnmapped events raised by the RoninGateway contract.
type RoninGatewayTokenUnmappedIterator struct {
	Event *RoninGatewayTokenUnmapped // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RoninGatewayTokenUnmappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninGatewayTokenUnmapped)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RoninGatewayTokenUnmapped)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RoninGatewayTokenUnmappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninGatewayTokenUnmappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninGatewayTokenUnmapped represents a TokenUnmapped event raised by the RoninGateway contract.
type RoninGatewayTokenUnmapped struct {
	RoninTokens []common.Address
	ChainIds    []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenUnmapped is a free log retrieval operation binding the contract event 0xd60250f85788f0c76ead12a4b178663fd7dcfd27c172f341d09680150ad11c49.
//
// Solidity: event TokenUnmapped(address[] roninTokens, uint256[] chainIds)
func (_RoninGateway *RoninGatewayFilterer) FilterTokenUnmapped(opts *bind.FilterOpts) (*RoninGatewayTokenUnmappedIterator, error) {

	logs, sub, err := _RoninGateway.contract.FilterLogs(opts, "TokenUnmapped")
	if err != nil {
		return nil, err
	}
	return &RoninGatewayTokenUnmappedIterator{contract: _RoninGateway.contract, event: "TokenUnmapped", logs: logs, sub: sub}, nil
}

// WatchTokenUnmapped is a free log subscription operation binding the contract event 0xd60250f85788f0c76ead12a4b178663fd7dcfd27c172f341d09680150ad11c49.
//
// Solidity: event TokenUnmapped(address[] roninTokens, uint256[] chainIds)
func (_RoninGateway *RoninGatewayFilterer) WatchTokenUnmapped(opts *bind.WatchOpts, sink chan<- *RoninGatewayTokenUnmapped) (event.Subscription, error) {

	logs, sub, err := _RoninGateway.contract.WatchLogs(opts, "TokenUnmapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninGatewayTokenUnmapped)
				if err := _RoninGateway.contract.UnpackLog(event, "TokenUnmapped", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenUnmapped is a log parse operation binding the contract event 0xd60250f85788f0c76ead12a4b178663fd7dcfd27c172f341d09680150ad11c49.
//
// Solidity: event TokenUnmapped(address[] roninTokens, uint256[] chainIds)
func (_RoninGateway *RoninGatewayFilterer) ParseTokenUnmapped(log types.Log) (*RoninGatewayTokenUnmapped, error) {
	event := new(RoninGatewayTokenUnmapped)
	if err := _RoninGateway.contract.UnpackLog(event, "TokenUnmapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninGatewayTrustedThresholdUpdatedIterator is returned from FilterTrustedThresholdUpdated and is used to iterate over the raw logs and unpacked data for TrustedThresholdUpdated events raised by the RoninGateway contract.
type RoninGatewayTrustedThresholdUpdatedIterator struct {
	Event *RoninGatewayTrustedThresholdUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RoninGatewayTrustedThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninGatewayTrustedThresholdUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RoninGatewayTrustedThresholdUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RoninGatewayTrustedThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninGatewayTrustedThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninGatewayTrustedThresholdUpdated represents a TrustedThresholdUpdated event raised by the RoninGateway contract.
type RoninGatewayTrustedThresholdUpdated struct {
	Nonce               *big.Int
	Numerator           *big.Int
	Denominator         *big.Int
	PreviousNumerator   *big.Int
	PreviousDenominator *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTrustedThresholdUpdated is a free log retrieval operation binding the contract event 0xeac82d4d949d2d4f77f96aa68ab6b1bb750da73f14e55d41a1b93f387471ecba.
//
// Solidity: event TrustedThresholdUpdated(uint256 indexed nonce, uint256 indexed numerator, uint256 indexed denominator, uint256 previousNumerator, uint256 previousDenominator)
func (_RoninGateway *RoninGatewayFilterer) FilterTrustedThresholdUpdated(opts *bind.FilterOpts, nonce []*big.Int, numerator []*big.Int, denominator []*big.Int) (*RoninGatewayTrustedThresholdUpdatedIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var numeratorRule []interface{}
	for _, numeratorItem := range numerator {
		numeratorRule = append(numeratorRule, numeratorItem)
	}
	var denominatorRule []interface{}
	for _, denominatorItem := range denominator {
		denominatorRule = append(denominatorRule, denominatorItem)
	}

	logs, sub, err := _RoninGateway.contract.FilterLogs(opts, "TrustedThresholdUpdated", nonceRule, numeratorRule, denominatorRule)
	if err != nil {
		return nil, err
	}
	return &RoninGatewayTrustedThresholdUpdatedIterator{contract: _RoninGateway.contract, event: "TrustedThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchTrustedThresholdUpdated is a free log subscription operation binding the contract event 0xeac82d4d949d2d4f77f96aa68ab6b1bb750da73f14e55d41a1b93f387471ecba.
//
// Solidity: event TrustedThresholdUpdated(uint256 indexed nonce, uint256 indexed numerator, uint256 indexed denominator, uint256 previousNumerator, uint256 previousDenominator)
func (_RoninGateway *RoninGatewayFilterer) WatchTrustedThresholdUpdated(opts *bind.WatchOpts, sink chan<- *RoninGatewayTrustedThresholdUpdated, nonce []*big.Int, numerator []*big.Int, denominator []*big.Int) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var numeratorRule []interface{}
	for _, numeratorItem := range numerator {
		numeratorRule = append(numeratorRule, numeratorItem)
	}
	var denominatorRule []interface{}
	for _, denominatorItem := range denominator {
		denominatorRule = append(denominatorRule, denominatorItem)
	}

	logs, sub, err := _RoninGateway.contract.WatchLogs(opts, "TrustedThresholdUpdated", nonceRule, numeratorRule, denominatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninGatewayTrustedThresholdUpdated)
				if err := _RoninGateway.contract.UnpackLog(event, "TrustedThresholdUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTrustedThresholdUpdated is a log parse operation binding the contract event 0xeac82d4d949d2d4f77f96aa68ab6b1bb750da73f14e55d41a1b93f387471ecba.
//
// Solidity: event TrustedThresholdUpdated(uint256 indexed nonce, uint256 indexed numerator, uint256 indexed denominator, uint256 previousNumerator, uint256 previousDenominator)
func (_RoninGateway *RoninGatewayFilterer) ParseTrustedThresholdUpdated(log types.Log) (*RoninGatewayTrustedThresholdUpdated, error) {
	event := new(RoninGatewayTrustedThresholdUpdated)
	if err := _RoninGateway.contract.UnpackLog(event, "TrustedThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninGatewayWithdrawalRequestedIterator is returned from FilterWithdrawalRequested and is used to iterate over the raw logs and unpacked data for WithdrawalRequested events raised by the RoninGateway contract.
type RoninGatewayWithdrawalRequestedIterator struct {
	Event *RoninGatewayWithdrawalRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RoninGatewayWithdrawalRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninGatewayWithdrawalRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RoninGatewayWithdrawalRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RoninGatewayWithdrawalRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninGatewayWithdrawalRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninGatewayWithdrawalRequested represents a WithdrawalRequested event raised by the RoninGateway contract.
type RoninGatewayWithdrawalRequested struct {
	ReceiptHash [32]byte
	Arg1        TransferReceipt
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalRequested is a free log retrieval operation binding the contract event 0xf313c253a5be72c29d0deb2c8768a9543744ac03d6b3cafd50cc976f1c2632fc.
//
// Solidity: event WithdrawalRequested(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) arg1)
func (_RoninGateway *RoninGatewayFilterer) FilterWithdrawalRequested(opts *bind.FilterOpts) (*RoninGatewayWithdrawalRequestedIterator, error) {

	logs, sub, err := _RoninGateway.contract.FilterLogs(opts, "WithdrawalRequested")
	if err != nil {
		return nil, err
	}
	return &RoninGatewayWithdrawalRequestedIterator{contract: _RoninGateway.contract, event: "WithdrawalRequested", logs: logs, sub: sub}, nil
}

// WatchWithdrawalRequested is a free log subscription operation binding the contract event 0xf313c253a5be72c29d0deb2c8768a9543744ac03d6b3cafd50cc976f1c2632fc.
//
// Solidity: event WithdrawalRequested(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) arg1)
func (_RoninGateway *RoninGatewayFilterer) WatchWithdrawalRequested(opts *bind.WatchOpts, sink chan<- *RoninGatewayWithdrawalRequested) (event.Subscription, error) {

	logs, sub, err := _RoninGateway.contract.WatchLogs(opts, "WithdrawalRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninGatewayWithdrawalRequested)
				if err := _RoninGateway.contract.UnpackLog(event, "WithdrawalRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawalRequested is a log parse operation binding the contract event 0xf313c253a5be72c29d0deb2c8768a9543744ac03d6b3cafd50cc976f1c2632fc.
//
// Solidity: event WithdrawalRequested(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) arg1)
func (_RoninGateway *RoninGatewayFilterer) ParseWithdrawalRequested(log types.Log) (*RoninGatewayWithdrawalRequested, error) {
	event := new(RoninGatewayWithdrawalRequested)
	if err := _RoninGateway.contract.UnpackLog(event, "WithdrawalRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninGatewayWithdrawalSignaturesRequestedIterator is returned from FilterWithdrawalSignaturesRequested and is used to iterate over the raw logs and unpacked data for WithdrawalSignaturesRequested events raised by the RoninGateway contract.
type RoninGatewayWithdrawalSignaturesRequestedIterator struct {
	Event *RoninGatewayWithdrawalSignaturesRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RoninGatewayWithdrawalSignaturesRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninGatewayWithdrawalSignaturesRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RoninGatewayWithdrawalSignaturesRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RoninGatewayWithdrawalSignaturesRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninGatewayWithdrawalSignaturesRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninGatewayWithdrawalSignaturesRequested represents a WithdrawalSignaturesRequested event raised by the RoninGateway contract.
type RoninGatewayWithdrawalSignaturesRequested struct {
	ReceiptHash [32]byte
	Arg1        TransferReceipt
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalSignaturesRequested is a free log retrieval operation binding the contract event 0x04e8cbd836dea43a2dc7eb19de345cca3a8e6978a2ef5225d924775500f67c7c.
//
// Solidity: event WithdrawalSignaturesRequested(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) arg1)
func (_RoninGateway *RoninGatewayFilterer) FilterWithdrawalSignaturesRequested(opts *bind.FilterOpts) (*RoninGatewayWithdrawalSignaturesRequestedIterator, error) {

	logs, sub, err := _RoninGateway.contract.FilterLogs(opts, "WithdrawalSignaturesRequested")
	if err != nil {
		return nil, err
	}
	return &RoninGatewayWithdrawalSignaturesRequestedIterator{contract: _RoninGateway.contract, event: "WithdrawalSignaturesRequested", logs: logs, sub: sub}, nil
}

// WatchWithdrawalSignaturesRequested is a free log subscription operation binding the contract event 0x04e8cbd836dea43a2dc7eb19de345cca3a8e6978a2ef5225d924775500f67c7c.
//
// Solidity: event WithdrawalSignaturesRequested(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) arg1)
func (_RoninGateway *RoninGatewayFilterer) WatchWithdrawalSignaturesRequested(opts *bind.WatchOpts, sink chan<- *RoninGatewayWithdrawalSignaturesRequested) (event.Subscription, error) {

	logs, sub, err := _RoninGateway.contract.WatchLogs(opts, "WithdrawalSignaturesRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninGatewayWithdrawalSignaturesRequested)
				if err := _RoninGateway.contract.UnpackLog(event, "WithdrawalSignaturesRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawalSignaturesRequested is a log parse operation binding the contract event 0x04e8cbd836dea43a2dc7eb19de345cca3a8e6978a2ef5225d924775500f67c7c.
//
// Solidity: event WithdrawalSignaturesRequested(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) arg1)
func (_RoninGateway *RoninGatewayFilterer) ParseWithdrawalSignaturesRequested(log types.Log) (*RoninGatewayWithdrawalSignaturesRequested, error) {
	event := new(RoninGatewayWithdrawalSignaturesRequested)
	if err := _RoninGateway.contract.UnpackLog(event, "WithdrawalSignaturesRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
