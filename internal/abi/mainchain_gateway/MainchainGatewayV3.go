// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mainchain_gateway

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

// SignatureConsumerSignature is an auto generated low-level Go binding around an user-defined struct.
type SignatureConsumerSignature struct {
	V uint8
	R [32]byte
	S [32]byte
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

// MainchainGatewayMetaData contains all meta data concerning the MainchainGateway contract.
var MainchainGatewayMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoninToken\",\"inputs\":[{\"name\":\"_mainchainToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"_token\",\"type\":\"tuple\",\"internalType\":\"structMappedTokenConsumer.MappedToken\",\"components\":[{\"name\":\"erc\",\"type\":\"uint8\",\"internalType\":\"enumTokenStandard\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mapTokens\",\"inputs\":[{\"name\":\"_mainchainTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_roninTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_standards\",\"type\":\"uint8[]\",\"internalType\":\"enumTokenStandard[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mapTokensAndThresholds\",\"inputs\":[{\"name\":\"_mainchainTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_roninTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_standards\",\"type\":\"uint8[]\",\"internalType\":\"enumTokenStandard[]\"},{\"name\":\"_thresholds\",\"type\":\"uint256[][4]\",\"internalType\":\"uint256[][4]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestDepositFor\",\"inputs\":[{\"name\":\"_request\",\"type\":\"tuple\",\"internalType\":\"structTransfer.Request\",\"components\":[{\"name\":\"recipientAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"info\",\"type\":\"tuple\",\"internalType\":\"structTokenInfo\",\"components\":[{\"name\":\"erc\",\"type\":\"uint8\",\"internalType\":\"enumTokenStandard\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setWrappedNativeTokenContract\",\"inputs\":[{\"name\":\"_wrappedToken\",\"type\":\"address\",\"internalType\":\"contractIWETH\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitWithdrawal\",\"inputs\":[{\"name\":\"_receipt\",\"type\":\"tuple\",\"internalType\":\"structTransfer.Receipt\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumTransfer.Kind\"},{\"name\":\"mainchain\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"ronin\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"info\",\"type\":\"tuple\",\"internalType\":\"structTokenInfo\",\"components\":[{\"name\":\"erc\",\"type\":\"uint8\",\"internalType\":\"enumTokenStandard\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"_signatures\",\"type\":\"tuple[]\",\"internalType\":\"structSignatureConsumer.Signature[]\",\"components\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"_locked\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unlockWithdrawal\",\"inputs\":[{\"name\":\"_receipt\",\"type\":\"tuple\",\"internalType\":\"structTransfer.Receipt\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumTransfer.Kind\"},{\"name\":\"mainchain\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"ronin\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"info\",\"type\":\"tuple\",\"internalType\":\"structTokenInfo\",\"components\":[{\"name\":\"erc\",\"type\":\"uint8\",\"internalType\":\"enumTokenStandard\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawalHash\",\"inputs\":[{\"name\":\"withdrawalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawalLocked\",\"inputs\":[{\"name\":\"withdrawalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"wrappedNativeToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIWETH\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DepositRequested\",\"inputs\":[{\"name\":\"receiptHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"receipt\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structTransfer.Receipt\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumTransfer.Kind\"},{\"name\":\"mainchain\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"ronin\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"info\",\"type\":\"tuple\",\"internalType\":\"structTokenInfo\",\"components\":[{\"name\":\"erc\",\"type\":\"uint8\",\"internalType\":\"enumTokenStandard\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenMapped\",\"inputs\":[{\"name\":\"mainchainTokens\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"roninTokens\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"standards\",\"type\":\"uint8[]\",\"indexed\":false,\"internalType\":\"enumTokenStandard[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalLocked\",\"inputs\":[{\"name\":\"receiptHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"receipt\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structTransfer.Receipt\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumTransfer.Kind\"},{\"name\":\"mainchain\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"ronin\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"info\",\"type\":\"tuple\",\"internalType\":\"structTokenInfo\",\"components\":[{\"name\":\"erc\",\"type\":\"uint8\",\"internalType\":\"enumTokenStandard\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalUnlocked\",\"inputs\":[{\"name\":\"receiptHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"receipt\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structTransfer.Receipt\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumTransfer.Kind\"},{\"name\":\"mainchain\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"ronin\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"info\",\"type\":\"tuple\",\"internalType\":\"structTokenInfo\",\"components\":[{\"name\":\"erc\",\"type\":\"uint8\",\"internalType\":\"enumTokenStandard\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrew\",\"inputs\":[{\"name\":\"receiptHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"receipt\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structTransfer.Receipt\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumTransfer.Kind\"},{\"name\":\"mainchain\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"ronin\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"info\",\"type\":\"tuple\",\"internalType\":\"structTokenInfo\",\"components\":[{\"name\":\"erc\",\"type\":\"uint8\",\"internalType\":\"enumTokenStandard\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WrappedNativeTokenContractUpdated\",\"inputs\":[{\"name\":\"weth\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIWETH\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ErrInvalidSigner\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sig\",\"type\":\"tuple\",\"internalType\":\"structSignatureConsumer.Signature\",\"components\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]},{\"type\":\"error\",\"name\":\"ErrNullTotalWeightProvided\",\"inputs\":[{\"name\":\"msgSig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"ErrQueryForApprovedWithdrawal\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrQueryForInsufficientVoteWeight\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrQueryForProcessedWithdrawal\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrReachedDailyWithdrawalLimit\",\"inputs\":[]}]",
}

// MainchainGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use MainchainGatewayMetaData.ABI instead.
var MainchainGatewayABI = MainchainGatewayMetaData.ABI

// MainchainGateway is an auto generated Go binding around an Ethereum contract.
type MainchainGateway struct {
	MainchainGatewayCaller     // Read-only binding to the contract
	MainchainGatewayTransactor // Write-only binding to the contract
	MainchainGatewayFilterer   // Log filterer for contract events
}

// MainchainGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainchainGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainchainGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainchainGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainchainGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainchainGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainchainGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainchainGatewaySession struct {
	Contract     *MainchainGateway // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainchainGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainchainGatewayCallerSession struct {
	Contract *MainchainGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MainchainGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainchainGatewayTransactorSession struct {
	Contract     *MainchainGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MainchainGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainchainGatewayRaw struct {
	Contract *MainchainGateway // Generic contract binding to access the raw methods on
}

// MainchainGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainchainGatewayCallerRaw struct {
	Contract *MainchainGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// MainchainGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainchainGatewayTransactorRaw struct {
	Contract *MainchainGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMainchainGateway creates a new instance of MainchainGateway, bound to a specific deployed contract.
func NewMainchainGateway(address common.Address, backend bind.ContractBackend) (*MainchainGateway, error) {
	contract, err := bindMainchainGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MainchainGateway{MainchainGatewayCaller: MainchainGatewayCaller{contract: contract}, MainchainGatewayTransactor: MainchainGatewayTransactor{contract: contract}, MainchainGatewayFilterer: MainchainGatewayFilterer{contract: contract}}, nil
}

// NewMainchainGatewayCaller creates a new read-only instance of MainchainGateway, bound to a specific deployed contract.
func NewMainchainGatewayCaller(address common.Address, caller bind.ContractCaller) (*MainchainGatewayCaller, error) {
	contract, err := bindMainchainGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayCaller{contract: contract}, nil
}

// NewMainchainGatewayTransactor creates a new write-only instance of MainchainGateway, bound to a specific deployed contract.
func NewMainchainGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*MainchainGatewayTransactor, error) {
	contract, err := bindMainchainGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayTransactor{contract: contract}, nil
}

// NewMainchainGatewayFilterer creates a new log filterer instance of MainchainGateway, bound to a specific deployed contract.
func NewMainchainGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*MainchainGatewayFilterer, error) {
	contract, err := bindMainchainGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayFilterer{contract: contract}, nil
}

// bindMainchainGateway binds a generic wrapper to an already deployed contract.
func bindMainchainGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MainchainGatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MainchainGateway *MainchainGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MainchainGateway.Contract.MainchainGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MainchainGateway *MainchainGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MainchainGateway.Contract.MainchainGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MainchainGateway *MainchainGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MainchainGateway.Contract.MainchainGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MainchainGateway *MainchainGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MainchainGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MainchainGateway *MainchainGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MainchainGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MainchainGateway *MainchainGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MainchainGateway.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_MainchainGateway *MainchainGatewayCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MainchainGateway.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_MainchainGateway *MainchainGatewaySession) DOMAINSEPARATOR() ([32]byte, error) {
	return _MainchainGateway.Contract.DOMAINSEPARATOR(&_MainchainGateway.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_MainchainGateway *MainchainGatewayCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _MainchainGateway.Contract.DOMAINSEPARATOR(&_MainchainGateway.CallOpts)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_MainchainGateway *MainchainGatewayCaller) DepositCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MainchainGateway.contract.Call(opts, &out, "depositCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_MainchainGateway *MainchainGatewaySession) DepositCount() (*big.Int, error) {
	return _MainchainGateway.Contract.DepositCount(&_MainchainGateway.CallOpts)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_MainchainGateway *MainchainGatewayCallerSession) DepositCount() (*big.Int, error) {
	return _MainchainGateway.Contract.DepositCount(&_MainchainGateway.CallOpts)
}

// GetRoninToken is a free data retrieval call binding the contract method 0xb2975794.
//
// Solidity: function getRoninToken(address _mainchainToken) view returns((uint8,address) _token)
func (_MainchainGateway *MainchainGatewayCaller) GetRoninToken(opts *bind.CallOpts, _mainchainToken common.Address) (MappedTokenConsumerMappedToken, error) {
	var out []interface{}
	err := _MainchainGateway.contract.Call(opts, &out, "getRoninToken", _mainchainToken)

	if err != nil {
		return *new(MappedTokenConsumerMappedToken), err
	}

	out0 := *abi.ConvertType(out[0], new(MappedTokenConsumerMappedToken)).(*MappedTokenConsumerMappedToken)

	return out0, err

}

// GetRoninToken is a free data retrieval call binding the contract method 0xb2975794.
//
// Solidity: function getRoninToken(address _mainchainToken) view returns((uint8,address) _token)
func (_MainchainGateway *MainchainGatewaySession) GetRoninToken(_mainchainToken common.Address) (MappedTokenConsumerMappedToken, error) {
	return _MainchainGateway.Contract.GetRoninToken(&_MainchainGateway.CallOpts, _mainchainToken)
}

// GetRoninToken is a free data retrieval call binding the contract method 0xb2975794.
//
// Solidity: function getRoninToken(address _mainchainToken) view returns((uint8,address) _token)
func (_MainchainGateway *MainchainGatewayCallerSession) GetRoninToken(_mainchainToken common.Address) (MappedTokenConsumerMappedToken, error) {
	return _MainchainGateway.Contract.GetRoninToken(&_MainchainGateway.CallOpts, _mainchainToken)
}

// WithdrawalHash is a free data retrieval call binding the contract method 0x6932be98.
//
// Solidity: function withdrawalHash(uint256 withdrawalId) view returns(bytes32)
func (_MainchainGateway *MainchainGatewayCaller) WithdrawalHash(opts *bind.CallOpts, withdrawalId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MainchainGateway.contract.Call(opts, &out, "withdrawalHash", withdrawalId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WithdrawalHash is a free data retrieval call binding the contract method 0x6932be98.
//
// Solidity: function withdrawalHash(uint256 withdrawalId) view returns(bytes32)
func (_MainchainGateway *MainchainGatewaySession) WithdrawalHash(withdrawalId *big.Int) ([32]byte, error) {
	return _MainchainGateway.Contract.WithdrawalHash(&_MainchainGateway.CallOpts, withdrawalId)
}

// WithdrawalHash is a free data retrieval call binding the contract method 0x6932be98.
//
// Solidity: function withdrawalHash(uint256 withdrawalId) view returns(bytes32)
func (_MainchainGateway *MainchainGatewayCallerSession) WithdrawalHash(withdrawalId *big.Int) ([32]byte, error) {
	return _MainchainGateway.Contract.WithdrawalHash(&_MainchainGateway.CallOpts, withdrawalId)
}

// WithdrawalLocked is a free data retrieval call binding the contract method 0x4d493f4e.
//
// Solidity: function withdrawalLocked(uint256 withdrawalId) view returns(bool)
func (_MainchainGateway *MainchainGatewayCaller) WithdrawalLocked(opts *bind.CallOpts, withdrawalId *big.Int) (bool, error) {
	var out []interface{}
	err := _MainchainGateway.contract.Call(opts, &out, "withdrawalLocked", withdrawalId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawalLocked is a free data retrieval call binding the contract method 0x4d493f4e.
//
// Solidity: function withdrawalLocked(uint256 withdrawalId) view returns(bool)
func (_MainchainGateway *MainchainGatewaySession) WithdrawalLocked(withdrawalId *big.Int) (bool, error) {
	return _MainchainGateway.Contract.WithdrawalLocked(&_MainchainGateway.CallOpts, withdrawalId)
}

// WithdrawalLocked is a free data retrieval call binding the contract method 0x4d493f4e.
//
// Solidity: function withdrawalLocked(uint256 withdrawalId) view returns(bool)
func (_MainchainGateway *MainchainGatewayCallerSession) WithdrawalLocked(withdrawalId *big.Int) (bool, error) {
	return _MainchainGateway.Contract.WithdrawalLocked(&_MainchainGateway.CallOpts, withdrawalId)
}

// WrappedNativeToken is a free data retrieval call binding the contract method 0x17fcb39b.
//
// Solidity: function wrappedNativeToken() view returns(address)
func (_MainchainGateway *MainchainGatewayCaller) WrappedNativeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MainchainGateway.contract.Call(opts, &out, "wrappedNativeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedNativeToken is a free data retrieval call binding the contract method 0x17fcb39b.
//
// Solidity: function wrappedNativeToken() view returns(address)
func (_MainchainGateway *MainchainGatewaySession) WrappedNativeToken() (common.Address, error) {
	return _MainchainGateway.Contract.WrappedNativeToken(&_MainchainGateway.CallOpts)
}

// WrappedNativeToken is a free data retrieval call binding the contract method 0x17fcb39b.
//
// Solidity: function wrappedNativeToken() view returns(address)
func (_MainchainGateway *MainchainGatewayCallerSession) WrappedNativeToken() (common.Address, error) {
	return _MainchainGateway.Contract.WrappedNativeToken(&_MainchainGateway.CallOpts)
}

// MapTokens is a paid mutator transaction binding the contract method 0x1b6e7594.
//
// Solidity: function mapTokens(address[] _mainchainTokens, address[] _roninTokens, uint8[] _standards) returns()
func (_MainchainGateway *MainchainGatewayTransactor) MapTokens(opts *bind.TransactOpts, _mainchainTokens []common.Address, _roninTokens []common.Address, _standards []uint8) (*types.Transaction, error) {
	return _MainchainGateway.contract.Transact(opts, "mapTokens", _mainchainTokens, _roninTokens, _standards)
}

// MapTokens is a paid mutator transaction binding the contract method 0x1b6e7594.
//
// Solidity: function mapTokens(address[] _mainchainTokens, address[] _roninTokens, uint8[] _standards) returns()
func (_MainchainGateway *MainchainGatewaySession) MapTokens(_mainchainTokens []common.Address, _roninTokens []common.Address, _standards []uint8) (*types.Transaction, error) {
	return _MainchainGateway.Contract.MapTokens(&_MainchainGateway.TransactOpts, _mainchainTokens, _roninTokens, _standards)
}

// MapTokens is a paid mutator transaction binding the contract method 0x1b6e7594.
//
// Solidity: function mapTokens(address[] _mainchainTokens, address[] _roninTokens, uint8[] _standards) returns()
func (_MainchainGateway *MainchainGatewayTransactorSession) MapTokens(_mainchainTokens []common.Address, _roninTokens []common.Address, _standards []uint8) (*types.Transaction, error) {
	return _MainchainGateway.Contract.MapTokens(&_MainchainGateway.TransactOpts, _mainchainTokens, _roninTokens, _standards)
}

// MapTokensAndThresholds is a paid mutator transaction binding the contract method 0xdff525e1.
//
// Solidity: function mapTokensAndThresholds(address[] _mainchainTokens, address[] _roninTokens, uint8[] _standards, uint256[][4] _thresholds) returns()
func (_MainchainGateway *MainchainGatewayTransactor) MapTokensAndThresholds(opts *bind.TransactOpts, _mainchainTokens []common.Address, _roninTokens []common.Address, _standards []uint8, _thresholds [4][]*big.Int) (*types.Transaction, error) {
	return _MainchainGateway.contract.Transact(opts, "mapTokensAndThresholds", _mainchainTokens, _roninTokens, _standards, _thresholds)
}

// MapTokensAndThresholds is a paid mutator transaction binding the contract method 0xdff525e1.
//
// Solidity: function mapTokensAndThresholds(address[] _mainchainTokens, address[] _roninTokens, uint8[] _standards, uint256[][4] _thresholds) returns()
func (_MainchainGateway *MainchainGatewaySession) MapTokensAndThresholds(_mainchainTokens []common.Address, _roninTokens []common.Address, _standards []uint8, _thresholds [4][]*big.Int) (*types.Transaction, error) {
	return _MainchainGateway.Contract.MapTokensAndThresholds(&_MainchainGateway.TransactOpts, _mainchainTokens, _roninTokens, _standards, _thresholds)
}

// MapTokensAndThresholds is a paid mutator transaction binding the contract method 0xdff525e1.
//
// Solidity: function mapTokensAndThresholds(address[] _mainchainTokens, address[] _roninTokens, uint8[] _standards, uint256[][4] _thresholds) returns()
func (_MainchainGateway *MainchainGatewayTransactorSession) MapTokensAndThresholds(_mainchainTokens []common.Address, _roninTokens []common.Address, _standards []uint8, _thresholds [4][]*big.Int) (*types.Transaction, error) {
	return _MainchainGateway.Contract.MapTokensAndThresholds(&_MainchainGateway.TransactOpts, _mainchainTokens, _roninTokens, _standards, _thresholds)
}

// RequestDepositFor is a paid mutator transaction binding the contract method 0x4b14557e.
//
// Solidity: function requestDepositFor((address,address,(uint8,uint256,uint256)) _request) payable returns()
func (_MainchainGateway *MainchainGatewayTransactor) RequestDepositFor(opts *bind.TransactOpts, _request TransferRequest) (*types.Transaction, error) {
	return _MainchainGateway.contract.Transact(opts, "requestDepositFor", _request)
}

// RequestDepositFor is a paid mutator transaction binding the contract method 0x4b14557e.
//
// Solidity: function requestDepositFor((address,address,(uint8,uint256,uint256)) _request) payable returns()
func (_MainchainGateway *MainchainGatewaySession) RequestDepositFor(_request TransferRequest) (*types.Transaction, error) {
	return _MainchainGateway.Contract.RequestDepositFor(&_MainchainGateway.TransactOpts, _request)
}

// RequestDepositFor is a paid mutator transaction binding the contract method 0x4b14557e.
//
// Solidity: function requestDepositFor((address,address,(uint8,uint256,uint256)) _request) payable returns()
func (_MainchainGateway *MainchainGatewayTransactorSession) RequestDepositFor(_request TransferRequest) (*types.Transaction, error) {
	return _MainchainGateway.Contract.RequestDepositFor(&_MainchainGateway.TransactOpts, _request)
}

// SetWrappedNativeTokenContract is a paid mutator transaction binding the contract method 0xd64af2a6.
//
// Solidity: function setWrappedNativeTokenContract(address _wrappedToken) returns()
func (_MainchainGateway *MainchainGatewayTransactor) SetWrappedNativeTokenContract(opts *bind.TransactOpts, _wrappedToken common.Address) (*types.Transaction, error) {
	return _MainchainGateway.contract.Transact(opts, "setWrappedNativeTokenContract", _wrappedToken)
}

// SetWrappedNativeTokenContract is a paid mutator transaction binding the contract method 0xd64af2a6.
//
// Solidity: function setWrappedNativeTokenContract(address _wrappedToken) returns()
func (_MainchainGateway *MainchainGatewaySession) SetWrappedNativeTokenContract(_wrappedToken common.Address) (*types.Transaction, error) {
	return _MainchainGateway.Contract.SetWrappedNativeTokenContract(&_MainchainGateway.TransactOpts, _wrappedToken)
}

// SetWrappedNativeTokenContract is a paid mutator transaction binding the contract method 0xd64af2a6.
//
// Solidity: function setWrappedNativeTokenContract(address _wrappedToken) returns()
func (_MainchainGateway *MainchainGatewayTransactorSession) SetWrappedNativeTokenContract(_wrappedToken common.Address) (*types.Transaction, error) {
	return _MainchainGateway.Contract.SetWrappedNativeTokenContract(&_MainchainGateway.TransactOpts, _wrappedToken)
}

// SubmitWithdrawal is a paid mutator transaction binding the contract method 0x4d0d6673.
//
// Solidity: function submitWithdrawal((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) _receipt, (uint8,bytes32,bytes32)[] _signatures) returns(bool _locked)
func (_MainchainGateway *MainchainGatewayTransactor) SubmitWithdrawal(opts *bind.TransactOpts, _receipt TransferReceipt, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _MainchainGateway.contract.Transact(opts, "submitWithdrawal", _receipt, _signatures)
}

// SubmitWithdrawal is a paid mutator transaction binding the contract method 0x4d0d6673.
//
// Solidity: function submitWithdrawal((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) _receipt, (uint8,bytes32,bytes32)[] _signatures) returns(bool _locked)
func (_MainchainGateway *MainchainGatewaySession) SubmitWithdrawal(_receipt TransferReceipt, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _MainchainGateway.Contract.SubmitWithdrawal(&_MainchainGateway.TransactOpts, _receipt, _signatures)
}

// SubmitWithdrawal is a paid mutator transaction binding the contract method 0x4d0d6673.
//
// Solidity: function submitWithdrawal((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) _receipt, (uint8,bytes32,bytes32)[] _signatures) returns(bool _locked)
func (_MainchainGateway *MainchainGatewayTransactorSession) SubmitWithdrawal(_receipt TransferReceipt, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _MainchainGateway.Contract.SubmitWithdrawal(&_MainchainGateway.TransactOpts, _receipt, _signatures)
}

// UnlockWithdrawal is a paid mutator transaction binding the contract method 0x9157921c.
//
// Solidity: function unlockWithdrawal((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) _receipt) returns()
func (_MainchainGateway *MainchainGatewayTransactor) UnlockWithdrawal(opts *bind.TransactOpts, _receipt TransferReceipt) (*types.Transaction, error) {
	return _MainchainGateway.contract.Transact(opts, "unlockWithdrawal", _receipt)
}

// UnlockWithdrawal is a paid mutator transaction binding the contract method 0x9157921c.
//
// Solidity: function unlockWithdrawal((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) _receipt) returns()
func (_MainchainGateway *MainchainGatewaySession) UnlockWithdrawal(_receipt TransferReceipt) (*types.Transaction, error) {
	return _MainchainGateway.Contract.UnlockWithdrawal(&_MainchainGateway.TransactOpts, _receipt)
}

// UnlockWithdrawal is a paid mutator transaction binding the contract method 0x9157921c.
//
// Solidity: function unlockWithdrawal((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) _receipt) returns()
func (_MainchainGateway *MainchainGatewayTransactorSession) UnlockWithdrawal(_receipt TransferReceipt) (*types.Transaction, error) {
	return _MainchainGateway.Contract.UnlockWithdrawal(&_MainchainGateway.TransactOpts, _receipt)
}

// MainchainGatewayDepositRequestedIterator is returned from FilterDepositRequested and is used to iterate over the raw logs and unpacked data for DepositRequested events raised by the MainchainGateway contract.
type MainchainGatewayDepositRequestedIterator struct {
	Event *MainchainGatewayDepositRequested // Event containing the contract specifics and raw log

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
func (it *MainchainGatewayDepositRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainGatewayDepositRequested)
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
		it.Event = new(MainchainGatewayDepositRequested)
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
func (it *MainchainGatewayDepositRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainGatewayDepositRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainGatewayDepositRequested represents a DepositRequested event raised by the MainchainGateway contract.
type MainchainGatewayDepositRequested struct {
	ReceiptHash [32]byte
	Receipt     TransferReceipt
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDepositRequested is a free log retrieval operation binding the contract event 0xd7b25068d9dc8d00765254cfb7f5070f98d263c8d68931d937c7362fa738048b.
//
// Solidity: event DepositRequested(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_MainchainGateway *MainchainGatewayFilterer) FilterDepositRequested(opts *bind.FilterOpts) (*MainchainGatewayDepositRequestedIterator, error) {

	logs, sub, err := _MainchainGateway.contract.FilterLogs(opts, "DepositRequested")
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayDepositRequestedIterator{contract: _MainchainGateway.contract, event: "DepositRequested", logs: logs, sub: sub}, nil
}

// WatchDepositRequested is a free log subscription operation binding the contract event 0xd7b25068d9dc8d00765254cfb7f5070f98d263c8d68931d937c7362fa738048b.
//
// Solidity: event DepositRequested(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_MainchainGateway *MainchainGatewayFilterer) WatchDepositRequested(opts *bind.WatchOpts, sink chan<- *MainchainGatewayDepositRequested) (event.Subscription, error) {

	logs, sub, err := _MainchainGateway.contract.WatchLogs(opts, "DepositRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainGatewayDepositRequested)
				if err := _MainchainGateway.contract.UnpackLog(event, "DepositRequested", log); err != nil {
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

// ParseDepositRequested is a log parse operation binding the contract event 0xd7b25068d9dc8d00765254cfb7f5070f98d263c8d68931d937c7362fa738048b.
//
// Solidity: event DepositRequested(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_MainchainGateway *MainchainGatewayFilterer) ParseDepositRequested(log types.Log) (*MainchainGatewayDepositRequested, error) {
	event := new(MainchainGatewayDepositRequested)
	if err := _MainchainGateway.contract.UnpackLog(event, "DepositRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainchainGatewayTokenMappedIterator is returned from FilterTokenMapped and is used to iterate over the raw logs and unpacked data for TokenMapped events raised by the MainchainGateway contract.
type MainchainGatewayTokenMappedIterator struct {
	Event *MainchainGatewayTokenMapped // Event containing the contract specifics and raw log

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
func (it *MainchainGatewayTokenMappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainGatewayTokenMapped)
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
		it.Event = new(MainchainGatewayTokenMapped)
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
func (it *MainchainGatewayTokenMappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainGatewayTokenMappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainGatewayTokenMapped represents a TokenMapped event raised by the MainchainGateway contract.
type MainchainGatewayTokenMapped struct {
	MainchainTokens []common.Address
	RoninTokens     []common.Address
	Standards       []uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTokenMapped is a free log retrieval operation binding the contract event 0xa4f03cc9c0e0aeb5b71b4ec800702753f65748c2cf3064695ba8e8b46be70444.
//
// Solidity: event TokenMapped(address[] mainchainTokens, address[] roninTokens, uint8[] standards)
func (_MainchainGateway *MainchainGatewayFilterer) FilterTokenMapped(opts *bind.FilterOpts) (*MainchainGatewayTokenMappedIterator, error) {

	logs, sub, err := _MainchainGateway.contract.FilterLogs(opts, "TokenMapped")
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayTokenMappedIterator{contract: _MainchainGateway.contract, event: "TokenMapped", logs: logs, sub: sub}, nil
}

// WatchTokenMapped is a free log subscription operation binding the contract event 0xa4f03cc9c0e0aeb5b71b4ec800702753f65748c2cf3064695ba8e8b46be70444.
//
// Solidity: event TokenMapped(address[] mainchainTokens, address[] roninTokens, uint8[] standards)
func (_MainchainGateway *MainchainGatewayFilterer) WatchTokenMapped(opts *bind.WatchOpts, sink chan<- *MainchainGatewayTokenMapped) (event.Subscription, error) {

	logs, sub, err := _MainchainGateway.contract.WatchLogs(opts, "TokenMapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainGatewayTokenMapped)
				if err := _MainchainGateway.contract.UnpackLog(event, "TokenMapped", log); err != nil {
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

// ParseTokenMapped is a log parse operation binding the contract event 0xa4f03cc9c0e0aeb5b71b4ec800702753f65748c2cf3064695ba8e8b46be70444.
//
// Solidity: event TokenMapped(address[] mainchainTokens, address[] roninTokens, uint8[] standards)
func (_MainchainGateway *MainchainGatewayFilterer) ParseTokenMapped(log types.Log) (*MainchainGatewayTokenMapped, error) {
	event := new(MainchainGatewayTokenMapped)
	if err := _MainchainGateway.contract.UnpackLog(event, "TokenMapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainchainGatewayWithdrawalLockedIterator is returned from FilterWithdrawalLocked and is used to iterate over the raw logs and unpacked data for WithdrawalLocked events raised by the MainchainGateway contract.
type MainchainGatewayWithdrawalLockedIterator struct {
	Event *MainchainGatewayWithdrawalLocked // Event containing the contract specifics and raw log

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
func (it *MainchainGatewayWithdrawalLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainGatewayWithdrawalLocked)
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
		it.Event = new(MainchainGatewayWithdrawalLocked)
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
func (it *MainchainGatewayWithdrawalLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainGatewayWithdrawalLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainGatewayWithdrawalLocked represents a WithdrawalLocked event raised by the MainchainGateway contract.
type MainchainGatewayWithdrawalLocked struct {
	ReceiptHash [32]byte
	Receipt     TransferReceipt
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalLocked is a free log retrieval operation binding the contract event 0x89e52969465b1f1866fc5d46fd62de953962e9cb33552443cd999eba05bd20dc.
//
// Solidity: event WithdrawalLocked(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_MainchainGateway *MainchainGatewayFilterer) FilterWithdrawalLocked(opts *bind.FilterOpts) (*MainchainGatewayWithdrawalLockedIterator, error) {

	logs, sub, err := _MainchainGateway.contract.FilterLogs(opts, "WithdrawalLocked")
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayWithdrawalLockedIterator{contract: _MainchainGateway.contract, event: "WithdrawalLocked", logs: logs, sub: sub}, nil
}

// WatchWithdrawalLocked is a free log subscription operation binding the contract event 0x89e52969465b1f1866fc5d46fd62de953962e9cb33552443cd999eba05bd20dc.
//
// Solidity: event WithdrawalLocked(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_MainchainGateway *MainchainGatewayFilterer) WatchWithdrawalLocked(opts *bind.WatchOpts, sink chan<- *MainchainGatewayWithdrawalLocked) (event.Subscription, error) {

	logs, sub, err := _MainchainGateway.contract.WatchLogs(opts, "WithdrawalLocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainGatewayWithdrawalLocked)
				if err := _MainchainGateway.contract.UnpackLog(event, "WithdrawalLocked", log); err != nil {
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

// ParseWithdrawalLocked is a log parse operation binding the contract event 0x89e52969465b1f1866fc5d46fd62de953962e9cb33552443cd999eba05bd20dc.
//
// Solidity: event WithdrawalLocked(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_MainchainGateway *MainchainGatewayFilterer) ParseWithdrawalLocked(log types.Log) (*MainchainGatewayWithdrawalLocked, error) {
	event := new(MainchainGatewayWithdrawalLocked)
	if err := _MainchainGateway.contract.UnpackLog(event, "WithdrawalLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainchainGatewayWithdrawalUnlockedIterator is returned from FilterWithdrawalUnlocked and is used to iterate over the raw logs and unpacked data for WithdrawalUnlocked events raised by the MainchainGateway contract.
type MainchainGatewayWithdrawalUnlockedIterator struct {
	Event *MainchainGatewayWithdrawalUnlocked // Event containing the contract specifics and raw log

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
func (it *MainchainGatewayWithdrawalUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainGatewayWithdrawalUnlocked)
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
		it.Event = new(MainchainGatewayWithdrawalUnlocked)
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
func (it *MainchainGatewayWithdrawalUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainGatewayWithdrawalUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainGatewayWithdrawalUnlocked represents a WithdrawalUnlocked event raised by the MainchainGateway contract.
type MainchainGatewayWithdrawalUnlocked struct {
	ReceiptHash [32]byte
	Receipt     TransferReceipt
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalUnlocked is a free log retrieval operation binding the contract event 0xd639511b37b3b002cca6cfe6bca0d833945a5af5a045578a0627fc43b79b2630.
//
// Solidity: event WithdrawalUnlocked(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_MainchainGateway *MainchainGatewayFilterer) FilterWithdrawalUnlocked(opts *bind.FilterOpts) (*MainchainGatewayWithdrawalUnlockedIterator, error) {

	logs, sub, err := _MainchainGateway.contract.FilterLogs(opts, "WithdrawalUnlocked")
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayWithdrawalUnlockedIterator{contract: _MainchainGateway.contract, event: "WithdrawalUnlocked", logs: logs, sub: sub}, nil
}

// WatchWithdrawalUnlocked is a free log subscription operation binding the contract event 0xd639511b37b3b002cca6cfe6bca0d833945a5af5a045578a0627fc43b79b2630.
//
// Solidity: event WithdrawalUnlocked(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_MainchainGateway *MainchainGatewayFilterer) WatchWithdrawalUnlocked(opts *bind.WatchOpts, sink chan<- *MainchainGatewayWithdrawalUnlocked) (event.Subscription, error) {

	logs, sub, err := _MainchainGateway.contract.WatchLogs(opts, "WithdrawalUnlocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainGatewayWithdrawalUnlocked)
				if err := _MainchainGateway.contract.UnpackLog(event, "WithdrawalUnlocked", log); err != nil {
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

// ParseWithdrawalUnlocked is a log parse operation binding the contract event 0xd639511b37b3b002cca6cfe6bca0d833945a5af5a045578a0627fc43b79b2630.
//
// Solidity: event WithdrawalUnlocked(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_MainchainGateway *MainchainGatewayFilterer) ParseWithdrawalUnlocked(log types.Log) (*MainchainGatewayWithdrawalUnlocked, error) {
	event := new(MainchainGatewayWithdrawalUnlocked)
	if err := _MainchainGateway.contract.UnpackLog(event, "WithdrawalUnlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainchainGatewayWithdrewIterator is returned from FilterWithdrew and is used to iterate over the raw logs and unpacked data for Withdrew events raised by the MainchainGateway contract.
type MainchainGatewayWithdrewIterator struct {
	Event *MainchainGatewayWithdrew // Event containing the contract specifics and raw log

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
func (it *MainchainGatewayWithdrewIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainGatewayWithdrew)
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
		it.Event = new(MainchainGatewayWithdrew)
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
func (it *MainchainGatewayWithdrewIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainGatewayWithdrewIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainGatewayWithdrew represents a Withdrew event raised by the MainchainGateway contract.
type MainchainGatewayWithdrew struct {
	ReceiptHash [32]byte
	Receipt     TransferReceipt
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrew is a free log retrieval operation binding the contract event 0x21e88e956aa3e086f6388e899965cef814688f99ad8bb29b08d396571016372d.
//
// Solidity: event Withdrew(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_MainchainGateway *MainchainGatewayFilterer) FilterWithdrew(opts *bind.FilterOpts) (*MainchainGatewayWithdrewIterator, error) {

	logs, sub, err := _MainchainGateway.contract.FilterLogs(opts, "Withdrew")
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayWithdrewIterator{contract: _MainchainGateway.contract, event: "Withdrew", logs: logs, sub: sub}, nil
}

// WatchWithdrew is a free log subscription operation binding the contract event 0x21e88e956aa3e086f6388e899965cef814688f99ad8bb29b08d396571016372d.
//
// Solidity: event Withdrew(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_MainchainGateway *MainchainGatewayFilterer) WatchWithdrew(opts *bind.WatchOpts, sink chan<- *MainchainGatewayWithdrew) (event.Subscription, error) {

	logs, sub, err := _MainchainGateway.contract.WatchLogs(opts, "Withdrew")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainGatewayWithdrew)
				if err := _MainchainGateway.contract.UnpackLog(event, "Withdrew", log); err != nil {
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

// ParseWithdrew is a log parse operation binding the contract event 0x21e88e956aa3e086f6388e899965cef814688f99ad8bb29b08d396571016372d.
//
// Solidity: event Withdrew(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_MainchainGateway *MainchainGatewayFilterer) ParseWithdrew(log types.Log) (*MainchainGatewayWithdrew, error) {
	event := new(MainchainGatewayWithdrew)
	if err := _MainchainGateway.contract.UnpackLog(event, "Withdrew", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainchainGatewayWrappedNativeTokenContractUpdatedIterator is returned from FilterWrappedNativeTokenContractUpdated and is used to iterate over the raw logs and unpacked data for WrappedNativeTokenContractUpdated events raised by the MainchainGateway contract.
type MainchainGatewayWrappedNativeTokenContractUpdatedIterator struct {
	Event *MainchainGatewayWrappedNativeTokenContractUpdated // Event containing the contract specifics and raw log

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
func (it *MainchainGatewayWrappedNativeTokenContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainGatewayWrappedNativeTokenContractUpdated)
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
		it.Event = new(MainchainGatewayWrappedNativeTokenContractUpdated)
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
func (it *MainchainGatewayWrappedNativeTokenContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainGatewayWrappedNativeTokenContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainGatewayWrappedNativeTokenContractUpdated represents a WrappedNativeTokenContractUpdated event raised by the MainchainGateway contract.
type MainchainGatewayWrappedNativeTokenContractUpdated struct {
	Weth common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWrappedNativeTokenContractUpdated is a free log retrieval operation binding the contract event 0x9d2334c23be647e994f27a72c5eee42a43d5bdcfe15bb88e939103c2b114cbaf.
//
// Solidity: event WrappedNativeTokenContractUpdated(address weth)
func (_MainchainGateway *MainchainGatewayFilterer) FilterWrappedNativeTokenContractUpdated(opts *bind.FilterOpts) (*MainchainGatewayWrappedNativeTokenContractUpdatedIterator, error) {

	logs, sub, err := _MainchainGateway.contract.FilterLogs(opts, "WrappedNativeTokenContractUpdated")
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayWrappedNativeTokenContractUpdatedIterator{contract: _MainchainGateway.contract, event: "WrappedNativeTokenContractUpdated", logs: logs, sub: sub}, nil
}

// WatchWrappedNativeTokenContractUpdated is a free log subscription operation binding the contract event 0x9d2334c23be647e994f27a72c5eee42a43d5bdcfe15bb88e939103c2b114cbaf.
//
// Solidity: event WrappedNativeTokenContractUpdated(address weth)
func (_MainchainGateway *MainchainGatewayFilterer) WatchWrappedNativeTokenContractUpdated(opts *bind.WatchOpts, sink chan<- *MainchainGatewayWrappedNativeTokenContractUpdated) (event.Subscription, error) {

	logs, sub, err := _MainchainGateway.contract.WatchLogs(opts, "WrappedNativeTokenContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainGatewayWrappedNativeTokenContractUpdated)
				if err := _MainchainGateway.contract.UnpackLog(event, "WrappedNativeTokenContractUpdated", log); err != nil {
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

// ParseWrappedNativeTokenContractUpdated is a log parse operation binding the contract event 0x9d2334c23be647e994f27a72c5eee42a43d5bdcfe15bb88e939103c2b114cbaf.
//
// Solidity: event WrappedNativeTokenContractUpdated(address weth)
func (_MainchainGateway *MainchainGatewayFilterer) ParseWrappedNativeTokenContractUpdated(log types.Log) (*MainchainGatewayWrappedNativeTokenContractUpdated, error) {
	event := new(MainchainGatewayWrappedNativeTokenContractUpdated)
	if err := _MainchainGateway.contract.UnpackLog(event, "WrappedNativeTokenContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
