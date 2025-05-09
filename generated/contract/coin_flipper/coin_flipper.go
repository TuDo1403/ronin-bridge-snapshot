// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package coin_flipper

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

// CoinFlipperCoinFlipConfig is an auto generated low-level Go binding around an user-defined struct.
type CoinFlipperCoinFlipConfig struct {
	TokenType    *big.Int
	TokenAddress common.Address
	TokenId      *big.Int
	TokenAmount  *big.Int
	Supply       *big.Int
}

// CoinFlipperMetaData contains all meta data concerning the CoinFlipper contract.
var CoinFlipperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"OnlyCoordinatorCanFulfill\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"player_\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"configId_\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"reqHash_\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"choice_\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"CoinFlipInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPaused_\",\"type\":\"bool\"}],\"name\":\"CoinFlipPauseToggled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"player_\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"configId_\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"reqHash_\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"playerWin_\",\"type\":\"bool\"}],\"name\":\"CoinFlipResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"configId_\",\"type\":\"uint256\"}],\"name\":\"ConfigDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"configId_\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftIndex_\",\"type\":\"uint256\"}],\"name\":\"ConfigNFTAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"configId_\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structCoinFlipper.CoinFlipConfig\",\"name\":\"config_\",\"type\":\"tuple\"}],\"name\":\"ConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"configId\",\"type\":\"uint16\"}],\"name\":\"coinFlipConfigs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reqHash\",\"type\":\"bytes32\"}],\"name\":\"coinFlipData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"choice\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"configId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_constantGas\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_transferGas\",\"type\":\"uint64\"}],\"name\":\"configureGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"constantGas\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16[]\",\"name\":\"_configIds\",\"type\":\"uint16[]\"}],\"name\":\"deleteCoinFlipConfigs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"estimateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"estFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callbackGasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_choice\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"_configId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_nftId\",\"type\":\"uint256\"}],\"name\":\"flipACoin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vrfCoordinator\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_constantGas\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_transferGas\",\"type\":\"uint64\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"}],\"name\":\"playHistory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalPlay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"countWinning\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"}],\"name\":\"playersWinStreak\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"winStreak\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_reqHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_randomSeed\",\"type\":\"uint256\"}],\"name\":\"rawFulfillRandomSeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenTypes\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_tokenAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenAmounts\",\"type\":\"uint256[]\"}],\"name\":\"rescueAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_value\",\"type\":\"bool\"}],\"name\":\"togglePauseFlipping\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"configId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfConfigByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferGas\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16[]\",\"name\":\"_configIds\",\"type\":\"uint16[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"}],\"internalType\":\"structCoinFlipper.CoinFlipConfig[]\",\"name\":\"_configs\",\"type\":\"tuple[]\"}],\"name\":\"updateCoinFlipConfigs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_configId\",\"type\":\"uint16\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"}],\"internalType\":\"structCoinFlipper.CoinFlipConfig\",\"name\":\"_config\",\"type\":\"tuple\"},{\"internalType\":\"uint256[]\",\"name\":\"_indexes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"}],\"name\":\"updateNftCoinFlipConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vrfCoordinator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// CoinFlipperABI is the input ABI used to generate the binding from.
// Deprecated: Use CoinFlipperMetaData.ABI instead.
var CoinFlipperABI = CoinFlipperMetaData.ABI

// CoinFlipper is an auto generated Go binding around an Ethereum contract.
type CoinFlipper struct {
	CoinFlipperCaller     // Read-only binding to the contract
	CoinFlipperTransactor // Write-only binding to the contract
	CoinFlipperFilterer   // Log filterer for contract events
}

// CoinFlipperCaller is an auto generated read-only Go binding around an Ethereum contract.
type CoinFlipperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoinFlipperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CoinFlipperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoinFlipperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CoinFlipperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoinFlipperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CoinFlipperSession struct {
	Contract     *CoinFlipper      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoinFlipperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CoinFlipperCallerSession struct {
	Contract *CoinFlipperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CoinFlipperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CoinFlipperTransactorSession struct {
	Contract     *CoinFlipperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CoinFlipperRaw is an auto generated low-level Go binding around an Ethereum contract.
type CoinFlipperRaw struct {
	Contract *CoinFlipper // Generic contract binding to access the raw methods on
}

// CoinFlipperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CoinFlipperCallerRaw struct {
	Contract *CoinFlipperCaller // Generic read-only contract binding to access the raw methods on
}

// CoinFlipperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CoinFlipperTransactorRaw struct {
	Contract *CoinFlipperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCoinFlipper creates a new instance of CoinFlipper, bound to a specific deployed contract.
func NewCoinFlipper(address common.Address, backend bind.ContractBackend) (*CoinFlipper, error) {
	contract, err := bindCoinFlipper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CoinFlipper{CoinFlipperCaller: CoinFlipperCaller{contract: contract}, CoinFlipperTransactor: CoinFlipperTransactor{contract: contract}, CoinFlipperFilterer: CoinFlipperFilterer{contract: contract}}, nil
}

// NewCoinFlipperCaller creates a new read-only instance of CoinFlipper, bound to a specific deployed contract.
func NewCoinFlipperCaller(address common.Address, caller bind.ContractCaller) (*CoinFlipperCaller, error) {
	contract, err := bindCoinFlipper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CoinFlipperCaller{contract: contract}, nil
}

// NewCoinFlipperTransactor creates a new write-only instance of CoinFlipper, bound to a specific deployed contract.
func NewCoinFlipperTransactor(address common.Address, transactor bind.ContractTransactor) (*CoinFlipperTransactor, error) {
	contract, err := bindCoinFlipper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CoinFlipperTransactor{contract: contract}, nil
}

// NewCoinFlipperFilterer creates a new log filterer instance of CoinFlipper, bound to a specific deployed contract.
func NewCoinFlipperFilterer(address common.Address, filterer bind.ContractFilterer) (*CoinFlipperFilterer, error) {
	contract, err := bindCoinFlipper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CoinFlipperFilterer{contract: contract}, nil
}

// bindCoinFlipper binds a generic wrapper to an already deployed contract.
func bindCoinFlipper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CoinFlipperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CoinFlipper *CoinFlipperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CoinFlipper.Contract.CoinFlipperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CoinFlipper *CoinFlipperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoinFlipper.Contract.CoinFlipperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CoinFlipper *CoinFlipperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CoinFlipper.Contract.CoinFlipperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CoinFlipper *CoinFlipperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CoinFlipper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CoinFlipper *CoinFlipperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoinFlipper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CoinFlipper *CoinFlipperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CoinFlipper.Contract.contract.Transact(opts, method, params...)
}

// CoinFlipConfigs is a free data retrieval call binding the contract method 0xeef81d62.
//
// Solidity: function coinFlipConfigs(uint16 configId) view returns(uint256 tokenType, address tokenAddress, uint256 tokenId, uint256 tokenAmount, uint256 supply)
func (_CoinFlipper *CoinFlipperCaller) CoinFlipConfigs(opts *bind.CallOpts, configId uint16) (struct {
	TokenType    *big.Int
	TokenAddress common.Address
	TokenId      *big.Int
	TokenAmount  *big.Int
	Supply       *big.Int
}, error) {
	var out []interface{}
	err := _CoinFlipper.contract.Call(opts, &out, "coinFlipConfigs", configId)

	outstruct := new(struct {
		TokenType    *big.Int
		TokenAddress common.Address
		TokenId      *big.Int
		TokenAmount  *big.Int
		Supply       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenType = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TokenAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TokenAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Supply = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CoinFlipConfigs is a free data retrieval call binding the contract method 0xeef81d62.
//
// Solidity: function coinFlipConfigs(uint16 configId) view returns(uint256 tokenType, address tokenAddress, uint256 tokenId, uint256 tokenAmount, uint256 supply)
func (_CoinFlipper *CoinFlipperSession) CoinFlipConfigs(configId uint16) (struct {
	TokenType    *big.Int
	TokenAddress common.Address
	TokenId      *big.Int
	TokenAmount  *big.Int
	Supply       *big.Int
}, error) {
	return _CoinFlipper.Contract.CoinFlipConfigs(&_CoinFlipper.CallOpts, configId)
}

// CoinFlipConfigs is a free data retrieval call binding the contract method 0xeef81d62.
//
// Solidity: function coinFlipConfigs(uint16 configId) view returns(uint256 tokenType, address tokenAddress, uint256 tokenId, uint256 tokenAmount, uint256 supply)
func (_CoinFlipper *CoinFlipperCallerSession) CoinFlipConfigs(configId uint16) (struct {
	TokenType    *big.Int
	TokenAddress common.Address
	TokenId      *big.Int
	TokenAmount  *big.Int
	Supply       *big.Int
}, error) {
	return _CoinFlipper.Contract.CoinFlipConfigs(&_CoinFlipper.CallOpts, configId)
}

// CoinFlipData is a free data retrieval call binding the contract method 0x381e77eb.
//
// Solidity: function coinFlipData(bytes32 reqHash) view returns(bool fulfilled, bool choice, uint16 configId, address player, uint256 nftId)
func (_CoinFlipper *CoinFlipperCaller) CoinFlipData(opts *bind.CallOpts, reqHash [32]byte) (struct {
	Fulfilled bool
	Choice    bool
	ConfigId  uint16
	Player    common.Address
	NftId     *big.Int
}, error) {
	var out []interface{}
	err := _CoinFlipper.contract.Call(opts, &out, "coinFlipData", reqHash)

	outstruct := new(struct {
		Fulfilled bool
		Choice    bool
		ConfigId  uint16
		Player    common.Address
		NftId     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fulfilled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Choice = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.ConfigId = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.Player = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.NftId = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CoinFlipData is a free data retrieval call binding the contract method 0x381e77eb.
//
// Solidity: function coinFlipData(bytes32 reqHash) view returns(bool fulfilled, bool choice, uint16 configId, address player, uint256 nftId)
func (_CoinFlipper *CoinFlipperSession) CoinFlipData(reqHash [32]byte) (struct {
	Fulfilled bool
	Choice    bool
	ConfigId  uint16
	Player    common.Address
	NftId     *big.Int
}, error) {
	return _CoinFlipper.Contract.CoinFlipData(&_CoinFlipper.CallOpts, reqHash)
}

// CoinFlipData is a free data retrieval call binding the contract method 0x381e77eb.
//
// Solidity: function coinFlipData(bytes32 reqHash) view returns(bool fulfilled, bool choice, uint16 configId, address player, uint256 nftId)
func (_CoinFlipper *CoinFlipperCallerSession) CoinFlipData(reqHash [32]byte) (struct {
	Fulfilled bool
	Choice    bool
	ConfigId  uint16
	Player    common.Address
	NftId     *big.Int
}, error) {
	return _CoinFlipper.Contract.CoinFlipData(&_CoinFlipper.CallOpts, reqHash)
}

// ConstantGas is a free data retrieval call binding the contract method 0x573fc055.
//
// Solidity: function constantGas() view returns(uint64)
func (_CoinFlipper *CoinFlipperCaller) ConstantGas(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _CoinFlipper.contract.Call(opts, &out, "constantGas")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ConstantGas is a free data retrieval call binding the contract method 0x573fc055.
//
// Solidity: function constantGas() view returns(uint64)
func (_CoinFlipper *CoinFlipperSession) ConstantGas() (uint64, error) {
	return _CoinFlipper.Contract.ConstantGas(&_CoinFlipper.CallOpts)
}

// ConstantGas is a free data retrieval call binding the contract method 0x573fc055.
//
// Solidity: function constantGas() view returns(uint64)
func (_CoinFlipper *CoinFlipperCallerSession) ConstantGas() (uint64, error) {
	return _CoinFlipper.Contract.ConstantGas(&_CoinFlipper.CallOpts)
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(uint256 gasPrice) view returns(uint256 estFee, uint256 callbackGasLimit)
func (_CoinFlipper *CoinFlipperCaller) EstimateFee(opts *bind.CallOpts, gasPrice *big.Int) (struct {
	EstFee           *big.Int
	CallbackGasLimit *big.Int
}, error) {
	var out []interface{}
	err := _CoinFlipper.contract.Call(opts, &out, "estimateFee", gasPrice)

	outstruct := new(struct {
		EstFee           *big.Int
		CallbackGasLimit *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EstFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CallbackGasLimit = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(uint256 gasPrice) view returns(uint256 estFee, uint256 callbackGasLimit)
func (_CoinFlipper *CoinFlipperSession) EstimateFee(gasPrice *big.Int) (struct {
	EstFee           *big.Int
	CallbackGasLimit *big.Int
}, error) {
	return _CoinFlipper.Contract.EstimateFee(&_CoinFlipper.CallOpts, gasPrice)
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(uint256 gasPrice) view returns(uint256 estFee, uint256 callbackGasLimit)
func (_CoinFlipper *CoinFlipperCallerSession) EstimateFee(gasPrice *big.Int) (struct {
	EstFee           *big.Int
	CallbackGasLimit *big.Int
}, error) {
	return _CoinFlipper.Contract.EstimateFee(&_CoinFlipper.CallOpts, gasPrice)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CoinFlipper *CoinFlipperCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CoinFlipper.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CoinFlipper *CoinFlipperSession) Owner() (common.Address, error) {
	return _CoinFlipper.Contract.Owner(&_CoinFlipper.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CoinFlipper *CoinFlipperCallerSession) Owner() (common.Address, error) {
	return _CoinFlipper.Contract.Owner(&_CoinFlipper.CallOpts)
}

// PlayHistory is a free data retrieval call binding the contract method 0xff68e848.
//
// Solidity: function playHistory(address player) view returns(uint256 totalPlay, uint256 countWinning)
func (_CoinFlipper *CoinFlipperCaller) PlayHistory(opts *bind.CallOpts, player common.Address) (struct {
	TotalPlay    *big.Int
	CountWinning *big.Int
}, error) {
	var out []interface{}
	err := _CoinFlipper.contract.Call(opts, &out, "playHistory", player)

	outstruct := new(struct {
		TotalPlay    *big.Int
		CountWinning *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalPlay = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CountWinning = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PlayHistory is a free data retrieval call binding the contract method 0xff68e848.
//
// Solidity: function playHistory(address player) view returns(uint256 totalPlay, uint256 countWinning)
func (_CoinFlipper *CoinFlipperSession) PlayHistory(player common.Address) (struct {
	TotalPlay    *big.Int
	CountWinning *big.Int
}, error) {
	return _CoinFlipper.Contract.PlayHistory(&_CoinFlipper.CallOpts, player)
}

// PlayHistory is a free data retrieval call binding the contract method 0xff68e848.
//
// Solidity: function playHistory(address player) view returns(uint256 totalPlay, uint256 countWinning)
func (_CoinFlipper *CoinFlipperCallerSession) PlayHistory(player common.Address) (struct {
	TotalPlay    *big.Int
	CountWinning *big.Int
}, error) {
	return _CoinFlipper.Contract.PlayHistory(&_CoinFlipper.CallOpts, player)
}

// PlayersWinStreak is a free data retrieval call binding the contract method 0x6ee5572b.
//
// Solidity: function playersWinStreak(address player) view returns(uint256 winStreak)
func (_CoinFlipper *CoinFlipperCaller) PlayersWinStreak(opts *bind.CallOpts, player common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CoinFlipper.contract.Call(opts, &out, "playersWinStreak", player)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlayersWinStreak is a free data retrieval call binding the contract method 0x6ee5572b.
//
// Solidity: function playersWinStreak(address player) view returns(uint256 winStreak)
func (_CoinFlipper *CoinFlipperSession) PlayersWinStreak(player common.Address) (*big.Int, error) {
	return _CoinFlipper.Contract.PlayersWinStreak(&_CoinFlipper.CallOpts, player)
}

// PlayersWinStreak is a free data retrieval call binding the contract method 0x6ee5572b.
//
// Solidity: function playersWinStreak(address player) view returns(uint256 winStreak)
func (_CoinFlipper *CoinFlipperCallerSession) PlayersWinStreak(player common.Address) (*big.Int, error) {
	return _CoinFlipper.Contract.PlayersWinStreak(&_CoinFlipper.CallOpts, player)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CoinFlipper *CoinFlipperCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CoinFlipper.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CoinFlipper *CoinFlipperSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CoinFlipper.Contract.SupportsInterface(&_CoinFlipper.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CoinFlipper *CoinFlipperCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CoinFlipper.Contract.SupportsInterface(&_CoinFlipper.CallOpts, interfaceId)
}

// TokenOfConfigByIndex is a free data retrieval call binding the contract method 0xf9a52fbc.
//
// Solidity: function tokenOfConfigByIndex(uint16 configId, uint256 index) view returns(uint256 id)
func (_CoinFlipper *CoinFlipperCaller) TokenOfConfigByIndex(opts *bind.CallOpts, configId uint16, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CoinFlipper.contract.Call(opts, &out, "tokenOfConfigByIndex", configId, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfConfigByIndex is a free data retrieval call binding the contract method 0xf9a52fbc.
//
// Solidity: function tokenOfConfigByIndex(uint16 configId, uint256 index) view returns(uint256 id)
func (_CoinFlipper *CoinFlipperSession) TokenOfConfigByIndex(configId uint16, index *big.Int) (*big.Int, error) {
	return _CoinFlipper.Contract.TokenOfConfigByIndex(&_CoinFlipper.CallOpts, configId, index)
}

// TokenOfConfigByIndex is a free data retrieval call binding the contract method 0xf9a52fbc.
//
// Solidity: function tokenOfConfigByIndex(uint16 configId, uint256 index) view returns(uint256 id)
func (_CoinFlipper *CoinFlipperCallerSession) TokenOfConfigByIndex(configId uint16, index *big.Int) (*big.Int, error) {
	return _CoinFlipper.Contract.TokenOfConfigByIndex(&_CoinFlipper.CallOpts, configId, index)
}

// TransferGas is a free data retrieval call binding the contract method 0xfa03f797.
//
// Solidity: function transferGas() view returns(uint64)
func (_CoinFlipper *CoinFlipperCaller) TransferGas(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _CoinFlipper.contract.Call(opts, &out, "transferGas")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TransferGas is a free data retrieval call binding the contract method 0xfa03f797.
//
// Solidity: function transferGas() view returns(uint64)
func (_CoinFlipper *CoinFlipperSession) TransferGas() (uint64, error) {
	return _CoinFlipper.Contract.TransferGas(&_CoinFlipper.CallOpts)
}

// TransferGas is a free data retrieval call binding the contract method 0xfa03f797.
//
// Solidity: function transferGas() view returns(uint64)
func (_CoinFlipper *CoinFlipperCallerSession) TransferGas() (uint64, error) {
	return _CoinFlipper.Contract.TransferGas(&_CoinFlipper.CallOpts)
}

// VrfCoordinator is a free data retrieval call binding the contract method 0xa3e56fa8.
//
// Solidity: function vrfCoordinator() view returns(address)
func (_CoinFlipper *CoinFlipperCaller) VrfCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CoinFlipper.contract.Call(opts, &out, "vrfCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VrfCoordinator is a free data retrieval call binding the contract method 0xa3e56fa8.
//
// Solidity: function vrfCoordinator() view returns(address)
func (_CoinFlipper *CoinFlipperSession) VrfCoordinator() (common.Address, error) {
	return _CoinFlipper.Contract.VrfCoordinator(&_CoinFlipper.CallOpts)
}

// VrfCoordinator is a free data retrieval call binding the contract method 0xa3e56fa8.
//
// Solidity: function vrfCoordinator() view returns(address)
func (_CoinFlipper *CoinFlipperCallerSession) VrfCoordinator() (common.Address, error) {
	return _CoinFlipper.Contract.VrfCoordinator(&_CoinFlipper.CallOpts)
}

// ConfigureGas is a paid mutator transaction binding the contract method 0x89534a7d.
//
// Solidity: function configureGas(uint64 _constantGas, uint64 _transferGas) returns()
func (_CoinFlipper *CoinFlipperTransactor) ConfigureGas(opts *bind.TransactOpts, _constantGas uint64, _transferGas uint64) (*types.Transaction, error) {
	return _CoinFlipper.contract.Transact(opts, "configureGas", _constantGas, _transferGas)
}

// ConfigureGas is a paid mutator transaction binding the contract method 0x89534a7d.
//
// Solidity: function configureGas(uint64 _constantGas, uint64 _transferGas) returns()
func (_CoinFlipper *CoinFlipperSession) ConfigureGas(_constantGas uint64, _transferGas uint64) (*types.Transaction, error) {
	return _CoinFlipper.Contract.ConfigureGas(&_CoinFlipper.TransactOpts, _constantGas, _transferGas)
}

// ConfigureGas is a paid mutator transaction binding the contract method 0x89534a7d.
//
// Solidity: function configureGas(uint64 _constantGas, uint64 _transferGas) returns()
func (_CoinFlipper *CoinFlipperTransactorSession) ConfigureGas(_constantGas uint64, _transferGas uint64) (*types.Transaction, error) {
	return _CoinFlipper.Contract.ConfigureGas(&_CoinFlipper.TransactOpts, _constantGas, _transferGas)
}

// DeleteCoinFlipConfigs is a paid mutator transaction binding the contract method 0x2265a6bc.
//
// Solidity: function deleteCoinFlipConfigs(uint16[] _configIds) returns()
func (_CoinFlipper *CoinFlipperTransactor) DeleteCoinFlipConfigs(opts *bind.TransactOpts, _configIds []uint16) (*types.Transaction, error) {
	return _CoinFlipper.contract.Transact(opts, "deleteCoinFlipConfigs", _configIds)
}

// DeleteCoinFlipConfigs is a paid mutator transaction binding the contract method 0x2265a6bc.
//
// Solidity: function deleteCoinFlipConfigs(uint16[] _configIds) returns()
func (_CoinFlipper *CoinFlipperSession) DeleteCoinFlipConfigs(_configIds []uint16) (*types.Transaction, error) {
	return _CoinFlipper.Contract.DeleteCoinFlipConfigs(&_CoinFlipper.TransactOpts, _configIds)
}

// DeleteCoinFlipConfigs is a paid mutator transaction binding the contract method 0x2265a6bc.
//
// Solidity: function deleteCoinFlipConfigs(uint16[] _configIds) returns()
func (_CoinFlipper *CoinFlipperTransactorSession) DeleteCoinFlipConfigs(_configIds []uint16) (*types.Transaction, error) {
	return _CoinFlipper.Contract.DeleteCoinFlipConfigs(&_CoinFlipper.TransactOpts, _configIds)
}

// FlipACoin is a paid mutator transaction binding the contract method 0xf2a45f9e.
//
// Solidity: function flipACoin(bool _choice, uint16 _configId, uint256 _nftId) payable returns(bytes32)
func (_CoinFlipper *CoinFlipperTransactor) FlipACoin(opts *bind.TransactOpts, _choice bool, _configId uint16, _nftId *big.Int) (*types.Transaction, error) {
	return _CoinFlipper.contract.Transact(opts, "flipACoin", _choice, _configId, _nftId)
}

// FlipACoin is a paid mutator transaction binding the contract method 0xf2a45f9e.
//
// Solidity: function flipACoin(bool _choice, uint16 _configId, uint256 _nftId) payable returns(bytes32)
func (_CoinFlipper *CoinFlipperSession) FlipACoin(_choice bool, _configId uint16, _nftId *big.Int) (*types.Transaction, error) {
	return _CoinFlipper.Contract.FlipACoin(&_CoinFlipper.TransactOpts, _choice, _configId, _nftId)
}

// FlipACoin is a paid mutator transaction binding the contract method 0xf2a45f9e.
//
// Solidity: function flipACoin(bool _choice, uint16 _configId, uint256 _nftId) payable returns(bytes32)
func (_CoinFlipper *CoinFlipperTransactorSession) FlipACoin(_choice bool, _configId uint16, _nftId *big.Int) (*types.Transaction, error) {
	return _CoinFlipper.Contract.FlipACoin(&_CoinFlipper.TransactOpts, _choice, _configId, _nftId)
}

// Initialize is a paid mutator transaction binding the contract method 0x989a8366.
//
// Solidity: function initialize(address _vrfCoordinator, uint64 _constantGas, uint64 _transferGas) returns()
func (_CoinFlipper *CoinFlipperTransactor) Initialize(opts *bind.TransactOpts, _vrfCoordinator common.Address, _constantGas uint64, _transferGas uint64) (*types.Transaction, error) {
	return _CoinFlipper.contract.Transact(opts, "initialize", _vrfCoordinator, _constantGas, _transferGas)
}

// Initialize is a paid mutator transaction binding the contract method 0x989a8366.
//
// Solidity: function initialize(address _vrfCoordinator, uint64 _constantGas, uint64 _transferGas) returns()
func (_CoinFlipper *CoinFlipperSession) Initialize(_vrfCoordinator common.Address, _constantGas uint64, _transferGas uint64) (*types.Transaction, error) {
	return _CoinFlipper.Contract.Initialize(&_CoinFlipper.TransactOpts, _vrfCoordinator, _constantGas, _transferGas)
}

// Initialize is a paid mutator transaction binding the contract method 0x989a8366.
//
// Solidity: function initialize(address _vrfCoordinator, uint64 _constantGas, uint64 _transferGas) returns()
func (_CoinFlipper *CoinFlipperTransactorSession) Initialize(_vrfCoordinator common.Address, _constantGas uint64, _transferGas uint64) (*types.Transaction, error) {
	return _CoinFlipper.Contract.Initialize(&_CoinFlipper.TransactOpts, _vrfCoordinator, _constantGas, _transferGas)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_CoinFlipper *CoinFlipperTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _CoinFlipper.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_CoinFlipper *CoinFlipperSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _CoinFlipper.Contract.OnERC1155BatchReceived(&_CoinFlipper.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_CoinFlipper *CoinFlipperTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _CoinFlipper.Contract.OnERC1155BatchReceived(&_CoinFlipper.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_CoinFlipper *CoinFlipperTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _CoinFlipper.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_CoinFlipper *CoinFlipperSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _CoinFlipper.Contract.OnERC1155Received(&_CoinFlipper.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_CoinFlipper *CoinFlipperTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _CoinFlipper.Contract.OnERC1155Received(&_CoinFlipper.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_CoinFlipper *CoinFlipperTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _CoinFlipper.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_CoinFlipper *CoinFlipperSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _CoinFlipper.Contract.OnERC721Received(&_CoinFlipper.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_CoinFlipper *CoinFlipperTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _CoinFlipper.Contract.OnERC721Received(&_CoinFlipper.TransactOpts, arg0, arg1, arg2, arg3)
}

// RawFulfillRandomSeed is a paid mutator transaction binding the contract method 0x542dc2a6.
//
// Solidity: function rawFulfillRandomSeed(bytes32 _reqHash, uint256 _randomSeed) returns()
func (_CoinFlipper *CoinFlipperTransactor) RawFulfillRandomSeed(opts *bind.TransactOpts, _reqHash [32]byte, _randomSeed *big.Int) (*types.Transaction, error) {
	return _CoinFlipper.contract.Transact(opts, "rawFulfillRandomSeed", _reqHash, _randomSeed)
}

// RawFulfillRandomSeed is a paid mutator transaction binding the contract method 0x542dc2a6.
//
// Solidity: function rawFulfillRandomSeed(bytes32 _reqHash, uint256 _randomSeed) returns()
func (_CoinFlipper *CoinFlipperSession) RawFulfillRandomSeed(_reqHash [32]byte, _randomSeed *big.Int) (*types.Transaction, error) {
	return _CoinFlipper.Contract.RawFulfillRandomSeed(&_CoinFlipper.TransactOpts, _reqHash, _randomSeed)
}

// RawFulfillRandomSeed is a paid mutator transaction binding the contract method 0x542dc2a6.
//
// Solidity: function rawFulfillRandomSeed(bytes32 _reqHash, uint256 _randomSeed) returns()
func (_CoinFlipper *CoinFlipperTransactorSession) RawFulfillRandomSeed(_reqHash [32]byte, _randomSeed *big.Int) (*types.Transaction, error) {
	return _CoinFlipper.Contract.RawFulfillRandomSeed(&_CoinFlipper.TransactOpts, _reqHash, _randomSeed)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CoinFlipper *CoinFlipperTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoinFlipper.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CoinFlipper *CoinFlipperSession) RenounceOwnership() (*types.Transaction, error) {
	return _CoinFlipper.Contract.RenounceOwnership(&_CoinFlipper.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CoinFlipper *CoinFlipperTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CoinFlipper.Contract.RenounceOwnership(&_CoinFlipper.TransactOpts)
}

// RescueAssets is a paid mutator transaction binding the contract method 0x3fba6650.
//
// Solidity: function rescueAssets(uint256[] _tokenTypes, address[] _tokenAddresses, uint256[] _tokenIds, uint256[] _tokenAmounts) returns()
func (_CoinFlipper *CoinFlipperTransactor) RescueAssets(opts *bind.TransactOpts, _tokenTypes []*big.Int, _tokenAddresses []common.Address, _tokenIds []*big.Int, _tokenAmounts []*big.Int) (*types.Transaction, error) {
	return _CoinFlipper.contract.Transact(opts, "rescueAssets", _tokenTypes, _tokenAddresses, _tokenIds, _tokenAmounts)
}

// RescueAssets is a paid mutator transaction binding the contract method 0x3fba6650.
//
// Solidity: function rescueAssets(uint256[] _tokenTypes, address[] _tokenAddresses, uint256[] _tokenIds, uint256[] _tokenAmounts) returns()
func (_CoinFlipper *CoinFlipperSession) RescueAssets(_tokenTypes []*big.Int, _tokenAddresses []common.Address, _tokenIds []*big.Int, _tokenAmounts []*big.Int) (*types.Transaction, error) {
	return _CoinFlipper.Contract.RescueAssets(&_CoinFlipper.TransactOpts, _tokenTypes, _tokenAddresses, _tokenIds, _tokenAmounts)
}

// RescueAssets is a paid mutator transaction binding the contract method 0x3fba6650.
//
// Solidity: function rescueAssets(uint256[] _tokenTypes, address[] _tokenAddresses, uint256[] _tokenIds, uint256[] _tokenAmounts) returns()
func (_CoinFlipper *CoinFlipperTransactorSession) RescueAssets(_tokenTypes []*big.Int, _tokenAddresses []common.Address, _tokenIds []*big.Int, _tokenAmounts []*big.Int) (*types.Transaction, error) {
	return _CoinFlipper.Contract.RescueAssets(&_CoinFlipper.TransactOpts, _tokenTypes, _tokenAddresses, _tokenIds, _tokenAmounts)
}

// TogglePauseFlipping is a paid mutator transaction binding the contract method 0x3af3d5c1.
//
// Solidity: function togglePauseFlipping(bool _value) returns()
func (_CoinFlipper *CoinFlipperTransactor) TogglePauseFlipping(opts *bind.TransactOpts, _value bool) (*types.Transaction, error) {
	return _CoinFlipper.contract.Transact(opts, "togglePauseFlipping", _value)
}

// TogglePauseFlipping is a paid mutator transaction binding the contract method 0x3af3d5c1.
//
// Solidity: function togglePauseFlipping(bool _value) returns()
func (_CoinFlipper *CoinFlipperSession) TogglePauseFlipping(_value bool) (*types.Transaction, error) {
	return _CoinFlipper.Contract.TogglePauseFlipping(&_CoinFlipper.TransactOpts, _value)
}

// TogglePauseFlipping is a paid mutator transaction binding the contract method 0x3af3d5c1.
//
// Solidity: function togglePauseFlipping(bool _value) returns()
func (_CoinFlipper *CoinFlipperTransactorSession) TogglePauseFlipping(_value bool) (*types.Transaction, error) {
	return _CoinFlipper.Contract.TogglePauseFlipping(&_CoinFlipper.TransactOpts, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CoinFlipper *CoinFlipperTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CoinFlipper.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CoinFlipper *CoinFlipperSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CoinFlipper.Contract.TransferOwnership(&_CoinFlipper.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CoinFlipper *CoinFlipperTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CoinFlipper.Contract.TransferOwnership(&_CoinFlipper.TransactOpts, newOwner)
}

// UpdateCoinFlipConfigs is a paid mutator transaction binding the contract method 0xd8516a2e.
//
// Solidity: function updateCoinFlipConfigs(uint16[] _configIds, (uint256,address,uint256,uint256,uint256)[] _configs) returns()
func (_CoinFlipper *CoinFlipperTransactor) UpdateCoinFlipConfigs(opts *bind.TransactOpts, _configIds []uint16, _configs []CoinFlipperCoinFlipConfig) (*types.Transaction, error) {
	return _CoinFlipper.contract.Transact(opts, "updateCoinFlipConfigs", _configIds, _configs)
}

// UpdateCoinFlipConfigs is a paid mutator transaction binding the contract method 0xd8516a2e.
//
// Solidity: function updateCoinFlipConfigs(uint16[] _configIds, (uint256,address,uint256,uint256,uint256)[] _configs) returns()
func (_CoinFlipper *CoinFlipperSession) UpdateCoinFlipConfigs(_configIds []uint16, _configs []CoinFlipperCoinFlipConfig) (*types.Transaction, error) {
	return _CoinFlipper.Contract.UpdateCoinFlipConfigs(&_CoinFlipper.TransactOpts, _configIds, _configs)
}

// UpdateCoinFlipConfigs is a paid mutator transaction binding the contract method 0xd8516a2e.
//
// Solidity: function updateCoinFlipConfigs(uint16[] _configIds, (uint256,address,uint256,uint256,uint256)[] _configs) returns()
func (_CoinFlipper *CoinFlipperTransactorSession) UpdateCoinFlipConfigs(_configIds []uint16, _configs []CoinFlipperCoinFlipConfig) (*types.Transaction, error) {
	return _CoinFlipper.Contract.UpdateCoinFlipConfigs(&_CoinFlipper.TransactOpts, _configIds, _configs)
}

// UpdateNftCoinFlipConfig is a paid mutator transaction binding the contract method 0x30623604.
//
// Solidity: function updateNftCoinFlipConfig(uint16 _configId, (uint256,address,uint256,uint256,uint256) _config, uint256[] _indexes, uint256[] _ids) returns()
func (_CoinFlipper *CoinFlipperTransactor) UpdateNftCoinFlipConfig(opts *bind.TransactOpts, _configId uint16, _config CoinFlipperCoinFlipConfig, _indexes []*big.Int, _ids []*big.Int) (*types.Transaction, error) {
	return _CoinFlipper.contract.Transact(opts, "updateNftCoinFlipConfig", _configId, _config, _indexes, _ids)
}

// UpdateNftCoinFlipConfig is a paid mutator transaction binding the contract method 0x30623604.
//
// Solidity: function updateNftCoinFlipConfig(uint16 _configId, (uint256,address,uint256,uint256,uint256) _config, uint256[] _indexes, uint256[] _ids) returns()
func (_CoinFlipper *CoinFlipperSession) UpdateNftCoinFlipConfig(_configId uint16, _config CoinFlipperCoinFlipConfig, _indexes []*big.Int, _ids []*big.Int) (*types.Transaction, error) {
	return _CoinFlipper.Contract.UpdateNftCoinFlipConfig(&_CoinFlipper.TransactOpts, _configId, _config, _indexes, _ids)
}

// UpdateNftCoinFlipConfig is a paid mutator transaction binding the contract method 0x30623604.
//
// Solidity: function updateNftCoinFlipConfig(uint16 _configId, (uint256,address,uint256,uint256,uint256) _config, uint256[] _indexes, uint256[] _ids) returns()
func (_CoinFlipper *CoinFlipperTransactorSession) UpdateNftCoinFlipConfig(_configId uint16, _config CoinFlipperCoinFlipConfig, _indexes []*big.Int, _ids []*big.Int) (*types.Transaction, error) {
	return _CoinFlipper.Contract.UpdateNftCoinFlipConfig(&_CoinFlipper.TransactOpts, _configId, _config, _indexes, _ids)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CoinFlipper *CoinFlipperTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoinFlipper.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CoinFlipper *CoinFlipperSession) Receive() (*types.Transaction, error) {
	return _CoinFlipper.Contract.Receive(&_CoinFlipper.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CoinFlipper *CoinFlipperTransactorSession) Receive() (*types.Transaction, error) {
	return _CoinFlipper.Contract.Receive(&_CoinFlipper.TransactOpts)
}

// CoinFlipperCoinFlipInitiatedIterator is returned from FilterCoinFlipInitiated and is used to iterate over the raw logs and unpacked data for CoinFlipInitiated events raised by the CoinFlipper contract.
type CoinFlipperCoinFlipInitiatedIterator struct {
	Event *CoinFlipperCoinFlipInitiated // Event containing the contract specifics and raw log

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
func (it *CoinFlipperCoinFlipInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinFlipperCoinFlipInitiated)
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
		it.Event = new(CoinFlipperCoinFlipInitiated)
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
func (it *CoinFlipperCoinFlipInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinFlipperCoinFlipInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinFlipperCoinFlipInitiated represents a CoinFlipInitiated event raised by the CoinFlipper contract.
type CoinFlipperCoinFlipInitiated struct {
	Player   common.Address
	ConfigId *big.Int
	ReqHash  [32]byte
	Choice   bool
	NftId    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCoinFlipInitiated is a free log retrieval operation binding the contract event 0xa1c0a23d2d8081d5731b50568b1f0a3a9e6733f775892c0d3eecec3baa7c3918.
//
// Solidity: event CoinFlipInitiated(address indexed player_, uint256 indexed configId_, bytes32 indexed reqHash_, bool choice_, uint256 nftId_)
func (_CoinFlipper *CoinFlipperFilterer) FilterCoinFlipInitiated(opts *bind.FilterOpts, player_ []common.Address, configId_ []*big.Int, reqHash_ [][32]byte) (*CoinFlipperCoinFlipInitiatedIterator, error) {

	var player_Rule []interface{}
	for _, player_Item := range player_ {
		player_Rule = append(player_Rule, player_Item)
	}
	var configId_Rule []interface{}
	for _, configId_Item := range configId_ {
		configId_Rule = append(configId_Rule, configId_Item)
	}
	var reqHash_Rule []interface{}
	for _, reqHash_Item := range reqHash_ {
		reqHash_Rule = append(reqHash_Rule, reqHash_Item)
	}

	logs, sub, err := _CoinFlipper.contract.FilterLogs(opts, "CoinFlipInitiated", player_Rule, configId_Rule, reqHash_Rule)
	if err != nil {
		return nil, err
	}
	return &CoinFlipperCoinFlipInitiatedIterator{contract: _CoinFlipper.contract, event: "CoinFlipInitiated", logs: logs, sub: sub}, nil
}

// WatchCoinFlipInitiated is a free log subscription operation binding the contract event 0xa1c0a23d2d8081d5731b50568b1f0a3a9e6733f775892c0d3eecec3baa7c3918.
//
// Solidity: event CoinFlipInitiated(address indexed player_, uint256 indexed configId_, bytes32 indexed reqHash_, bool choice_, uint256 nftId_)
func (_CoinFlipper *CoinFlipperFilterer) WatchCoinFlipInitiated(opts *bind.WatchOpts, sink chan<- *CoinFlipperCoinFlipInitiated, player_ []common.Address, configId_ []*big.Int, reqHash_ [][32]byte) (event.Subscription, error) {

	var player_Rule []interface{}
	for _, player_Item := range player_ {
		player_Rule = append(player_Rule, player_Item)
	}
	var configId_Rule []interface{}
	for _, configId_Item := range configId_ {
		configId_Rule = append(configId_Rule, configId_Item)
	}
	var reqHash_Rule []interface{}
	for _, reqHash_Item := range reqHash_ {
		reqHash_Rule = append(reqHash_Rule, reqHash_Item)
	}

	logs, sub, err := _CoinFlipper.contract.WatchLogs(opts, "CoinFlipInitiated", player_Rule, configId_Rule, reqHash_Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinFlipperCoinFlipInitiated)
				if err := _CoinFlipper.contract.UnpackLog(event, "CoinFlipInitiated", log); err != nil {
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

// ParseCoinFlipInitiated is a log parse operation binding the contract event 0xa1c0a23d2d8081d5731b50568b1f0a3a9e6733f775892c0d3eecec3baa7c3918.
//
// Solidity: event CoinFlipInitiated(address indexed player_, uint256 indexed configId_, bytes32 indexed reqHash_, bool choice_, uint256 nftId_)
func (_CoinFlipper *CoinFlipperFilterer) ParseCoinFlipInitiated(log types.Log) (*CoinFlipperCoinFlipInitiated, error) {
	event := new(CoinFlipperCoinFlipInitiated)
	if err := _CoinFlipper.contract.UnpackLog(event, "CoinFlipInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoinFlipperCoinFlipPauseToggledIterator is returned from FilterCoinFlipPauseToggled and is used to iterate over the raw logs and unpacked data for CoinFlipPauseToggled events raised by the CoinFlipper contract.
type CoinFlipperCoinFlipPauseToggledIterator struct {
	Event *CoinFlipperCoinFlipPauseToggled // Event containing the contract specifics and raw log

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
func (it *CoinFlipperCoinFlipPauseToggledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinFlipperCoinFlipPauseToggled)
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
		it.Event = new(CoinFlipperCoinFlipPauseToggled)
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
func (it *CoinFlipperCoinFlipPauseToggledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinFlipperCoinFlipPauseToggledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinFlipperCoinFlipPauseToggled represents a CoinFlipPauseToggled event raised by the CoinFlipper contract.
type CoinFlipperCoinFlipPauseToggled struct {
	IsPaused bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCoinFlipPauseToggled is a free log retrieval operation binding the contract event 0x2e9c8539d5aa626c4516fafbb699d424ffa5718914caec164b604035a477cf5b.
//
// Solidity: event CoinFlipPauseToggled(bool isPaused_)
func (_CoinFlipper *CoinFlipperFilterer) FilterCoinFlipPauseToggled(opts *bind.FilterOpts) (*CoinFlipperCoinFlipPauseToggledIterator, error) {

	logs, sub, err := _CoinFlipper.contract.FilterLogs(opts, "CoinFlipPauseToggled")
	if err != nil {
		return nil, err
	}
	return &CoinFlipperCoinFlipPauseToggledIterator{contract: _CoinFlipper.contract, event: "CoinFlipPauseToggled", logs: logs, sub: sub}, nil
}

// WatchCoinFlipPauseToggled is a free log subscription operation binding the contract event 0x2e9c8539d5aa626c4516fafbb699d424ffa5718914caec164b604035a477cf5b.
//
// Solidity: event CoinFlipPauseToggled(bool isPaused_)
func (_CoinFlipper *CoinFlipperFilterer) WatchCoinFlipPauseToggled(opts *bind.WatchOpts, sink chan<- *CoinFlipperCoinFlipPauseToggled) (event.Subscription, error) {

	logs, sub, err := _CoinFlipper.contract.WatchLogs(opts, "CoinFlipPauseToggled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinFlipperCoinFlipPauseToggled)
				if err := _CoinFlipper.contract.UnpackLog(event, "CoinFlipPauseToggled", log); err != nil {
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

// ParseCoinFlipPauseToggled is a log parse operation binding the contract event 0x2e9c8539d5aa626c4516fafbb699d424ffa5718914caec164b604035a477cf5b.
//
// Solidity: event CoinFlipPauseToggled(bool isPaused_)
func (_CoinFlipper *CoinFlipperFilterer) ParseCoinFlipPauseToggled(log types.Log) (*CoinFlipperCoinFlipPauseToggled, error) {
	event := new(CoinFlipperCoinFlipPauseToggled)
	if err := _CoinFlipper.contract.UnpackLog(event, "CoinFlipPauseToggled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoinFlipperCoinFlipResolvedIterator is returned from FilterCoinFlipResolved and is used to iterate over the raw logs and unpacked data for CoinFlipResolved events raised by the CoinFlipper contract.
type CoinFlipperCoinFlipResolvedIterator struct {
	Event *CoinFlipperCoinFlipResolved // Event containing the contract specifics and raw log

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
func (it *CoinFlipperCoinFlipResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinFlipperCoinFlipResolved)
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
		it.Event = new(CoinFlipperCoinFlipResolved)
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
func (it *CoinFlipperCoinFlipResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinFlipperCoinFlipResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinFlipperCoinFlipResolved represents a CoinFlipResolved event raised by the CoinFlipper contract.
type CoinFlipperCoinFlipResolved struct {
	Player    common.Address
	ConfigId  *big.Int
	ReqHash   [32]byte
	PlayerWin bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCoinFlipResolved is a free log retrieval operation binding the contract event 0x6dd584b4b41fce62e13224deae7109f49c59c133c9bc8e52e08da35bfb45a343.
//
// Solidity: event CoinFlipResolved(address indexed player_, uint256 indexed configId_, bytes32 indexed reqHash_, bool playerWin_)
func (_CoinFlipper *CoinFlipperFilterer) FilterCoinFlipResolved(opts *bind.FilterOpts, player_ []common.Address, configId_ []*big.Int, reqHash_ [][32]byte) (*CoinFlipperCoinFlipResolvedIterator, error) {

	var player_Rule []interface{}
	for _, player_Item := range player_ {
		player_Rule = append(player_Rule, player_Item)
	}
	var configId_Rule []interface{}
	for _, configId_Item := range configId_ {
		configId_Rule = append(configId_Rule, configId_Item)
	}
	var reqHash_Rule []interface{}
	for _, reqHash_Item := range reqHash_ {
		reqHash_Rule = append(reqHash_Rule, reqHash_Item)
	}

	logs, sub, err := _CoinFlipper.contract.FilterLogs(opts, "CoinFlipResolved", player_Rule, configId_Rule, reqHash_Rule)
	if err != nil {
		return nil, err
	}
	return &CoinFlipperCoinFlipResolvedIterator{contract: _CoinFlipper.contract, event: "CoinFlipResolved", logs: logs, sub: sub}, nil
}

// WatchCoinFlipResolved is a free log subscription operation binding the contract event 0x6dd584b4b41fce62e13224deae7109f49c59c133c9bc8e52e08da35bfb45a343.
//
// Solidity: event CoinFlipResolved(address indexed player_, uint256 indexed configId_, bytes32 indexed reqHash_, bool playerWin_)
func (_CoinFlipper *CoinFlipperFilterer) WatchCoinFlipResolved(opts *bind.WatchOpts, sink chan<- *CoinFlipperCoinFlipResolved, player_ []common.Address, configId_ []*big.Int, reqHash_ [][32]byte) (event.Subscription, error) {

	var player_Rule []interface{}
	for _, player_Item := range player_ {
		player_Rule = append(player_Rule, player_Item)
	}
	var configId_Rule []interface{}
	for _, configId_Item := range configId_ {
		configId_Rule = append(configId_Rule, configId_Item)
	}
	var reqHash_Rule []interface{}
	for _, reqHash_Item := range reqHash_ {
		reqHash_Rule = append(reqHash_Rule, reqHash_Item)
	}

	logs, sub, err := _CoinFlipper.contract.WatchLogs(opts, "CoinFlipResolved", player_Rule, configId_Rule, reqHash_Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinFlipperCoinFlipResolved)
				if err := _CoinFlipper.contract.UnpackLog(event, "CoinFlipResolved", log); err != nil {
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

// ParseCoinFlipResolved is a log parse operation binding the contract event 0x6dd584b4b41fce62e13224deae7109f49c59c133c9bc8e52e08da35bfb45a343.
//
// Solidity: event CoinFlipResolved(address indexed player_, uint256 indexed configId_, bytes32 indexed reqHash_, bool playerWin_)
func (_CoinFlipper *CoinFlipperFilterer) ParseCoinFlipResolved(log types.Log) (*CoinFlipperCoinFlipResolved, error) {
	event := new(CoinFlipperCoinFlipResolved)
	if err := _CoinFlipper.contract.UnpackLog(event, "CoinFlipResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoinFlipperConfigDeletedIterator is returned from FilterConfigDeleted and is used to iterate over the raw logs and unpacked data for ConfigDeleted events raised by the CoinFlipper contract.
type CoinFlipperConfigDeletedIterator struct {
	Event *CoinFlipperConfigDeleted // Event containing the contract specifics and raw log

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
func (it *CoinFlipperConfigDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinFlipperConfigDeleted)
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
		it.Event = new(CoinFlipperConfigDeleted)
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
func (it *CoinFlipperConfigDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinFlipperConfigDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinFlipperConfigDeleted represents a ConfigDeleted event raised by the CoinFlipper contract.
type CoinFlipperConfigDeleted struct {
	ConfigId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterConfigDeleted is a free log retrieval operation binding the contract event 0x324f658e47dac48cd6c9d447b62fd8060669c35e7e570a80d53d6660f0363992.
//
// Solidity: event ConfigDeleted(uint256 indexed configId_)
func (_CoinFlipper *CoinFlipperFilterer) FilterConfigDeleted(opts *bind.FilterOpts, configId_ []*big.Int) (*CoinFlipperConfigDeletedIterator, error) {

	var configId_Rule []interface{}
	for _, configId_Item := range configId_ {
		configId_Rule = append(configId_Rule, configId_Item)
	}

	logs, sub, err := _CoinFlipper.contract.FilterLogs(opts, "ConfigDeleted", configId_Rule)
	if err != nil {
		return nil, err
	}
	return &CoinFlipperConfigDeletedIterator{contract: _CoinFlipper.contract, event: "ConfigDeleted", logs: logs, sub: sub}, nil
}

// WatchConfigDeleted is a free log subscription operation binding the contract event 0x324f658e47dac48cd6c9d447b62fd8060669c35e7e570a80d53d6660f0363992.
//
// Solidity: event ConfigDeleted(uint256 indexed configId_)
func (_CoinFlipper *CoinFlipperFilterer) WatchConfigDeleted(opts *bind.WatchOpts, sink chan<- *CoinFlipperConfigDeleted, configId_ []*big.Int) (event.Subscription, error) {

	var configId_Rule []interface{}
	for _, configId_Item := range configId_ {
		configId_Rule = append(configId_Rule, configId_Item)
	}

	logs, sub, err := _CoinFlipper.contract.WatchLogs(opts, "ConfigDeleted", configId_Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinFlipperConfigDeleted)
				if err := _CoinFlipper.contract.UnpackLog(event, "ConfigDeleted", log); err != nil {
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

// ParseConfigDeleted is a log parse operation binding the contract event 0x324f658e47dac48cd6c9d447b62fd8060669c35e7e570a80d53d6660f0363992.
//
// Solidity: event ConfigDeleted(uint256 indexed configId_)
func (_CoinFlipper *CoinFlipperFilterer) ParseConfigDeleted(log types.Log) (*CoinFlipperConfigDeleted, error) {
	event := new(CoinFlipperConfigDeleted)
	if err := _CoinFlipper.contract.UnpackLog(event, "ConfigDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoinFlipperConfigNFTAddedIterator is returned from FilterConfigNFTAdded and is used to iterate over the raw logs and unpacked data for ConfigNFTAdded events raised by the CoinFlipper contract.
type CoinFlipperConfigNFTAddedIterator struct {
	Event *CoinFlipperConfigNFTAdded // Event containing the contract specifics and raw log

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
func (it *CoinFlipperConfigNFTAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinFlipperConfigNFTAdded)
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
		it.Event = new(CoinFlipperConfigNFTAdded)
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
func (it *CoinFlipperConfigNFTAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinFlipperConfigNFTAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinFlipperConfigNFTAdded represents a ConfigNFTAdded event raised by the CoinFlipper contract.
type CoinFlipperConfigNFTAdded struct {
	ConfigId *big.Int
	NftId    *big.Int
	NftIndex *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterConfigNFTAdded is a free log retrieval operation binding the contract event 0xc32dd4b4f69498c52fd4a4a67ab5ec826c366b1666475b30c2a6bb46884a34cc.
//
// Solidity: event ConfigNFTAdded(uint256 indexed configId_, uint256 nftId_, uint256 nftIndex_)
func (_CoinFlipper *CoinFlipperFilterer) FilterConfigNFTAdded(opts *bind.FilterOpts, configId_ []*big.Int) (*CoinFlipperConfigNFTAddedIterator, error) {

	var configId_Rule []interface{}
	for _, configId_Item := range configId_ {
		configId_Rule = append(configId_Rule, configId_Item)
	}

	logs, sub, err := _CoinFlipper.contract.FilterLogs(opts, "ConfigNFTAdded", configId_Rule)
	if err != nil {
		return nil, err
	}
	return &CoinFlipperConfigNFTAddedIterator{contract: _CoinFlipper.contract, event: "ConfigNFTAdded", logs: logs, sub: sub}, nil
}

// WatchConfigNFTAdded is a free log subscription operation binding the contract event 0xc32dd4b4f69498c52fd4a4a67ab5ec826c366b1666475b30c2a6bb46884a34cc.
//
// Solidity: event ConfigNFTAdded(uint256 indexed configId_, uint256 nftId_, uint256 nftIndex_)
func (_CoinFlipper *CoinFlipperFilterer) WatchConfigNFTAdded(opts *bind.WatchOpts, sink chan<- *CoinFlipperConfigNFTAdded, configId_ []*big.Int) (event.Subscription, error) {

	var configId_Rule []interface{}
	for _, configId_Item := range configId_ {
		configId_Rule = append(configId_Rule, configId_Item)
	}

	logs, sub, err := _CoinFlipper.contract.WatchLogs(opts, "ConfigNFTAdded", configId_Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinFlipperConfigNFTAdded)
				if err := _CoinFlipper.contract.UnpackLog(event, "ConfigNFTAdded", log); err != nil {
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

// ParseConfigNFTAdded is a log parse operation binding the contract event 0xc32dd4b4f69498c52fd4a4a67ab5ec826c366b1666475b30c2a6bb46884a34cc.
//
// Solidity: event ConfigNFTAdded(uint256 indexed configId_, uint256 nftId_, uint256 nftIndex_)
func (_CoinFlipper *CoinFlipperFilterer) ParseConfigNFTAdded(log types.Log) (*CoinFlipperConfigNFTAdded, error) {
	event := new(CoinFlipperConfigNFTAdded)
	if err := _CoinFlipper.contract.UnpackLog(event, "ConfigNFTAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoinFlipperConfigUpdatedIterator is returned from FilterConfigUpdated and is used to iterate over the raw logs and unpacked data for ConfigUpdated events raised by the CoinFlipper contract.
type CoinFlipperConfigUpdatedIterator struct {
	Event *CoinFlipperConfigUpdated // Event containing the contract specifics and raw log

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
func (it *CoinFlipperConfigUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinFlipperConfigUpdated)
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
		it.Event = new(CoinFlipperConfigUpdated)
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
func (it *CoinFlipperConfigUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinFlipperConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinFlipperConfigUpdated represents a ConfigUpdated event raised by the CoinFlipper contract.
type CoinFlipperConfigUpdated struct {
	ConfigId *big.Int
	Config   CoinFlipperCoinFlipConfig
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterConfigUpdated is a free log retrieval operation binding the contract event 0x606bf8e18a5e25fc2918dad277ffe9ec4b87e586deb4a7745b6df95b46e1b134.
//
// Solidity: event ConfigUpdated(uint256 indexed configId_, (uint256,address,uint256,uint256,uint256) config_)
func (_CoinFlipper *CoinFlipperFilterer) FilterConfigUpdated(opts *bind.FilterOpts, configId_ []*big.Int) (*CoinFlipperConfigUpdatedIterator, error) {

	var configId_Rule []interface{}
	for _, configId_Item := range configId_ {
		configId_Rule = append(configId_Rule, configId_Item)
	}

	logs, sub, err := _CoinFlipper.contract.FilterLogs(opts, "ConfigUpdated", configId_Rule)
	if err != nil {
		return nil, err
	}
	return &CoinFlipperConfigUpdatedIterator{contract: _CoinFlipper.contract, event: "ConfigUpdated", logs: logs, sub: sub}, nil
}

// WatchConfigUpdated is a free log subscription operation binding the contract event 0x606bf8e18a5e25fc2918dad277ffe9ec4b87e586deb4a7745b6df95b46e1b134.
//
// Solidity: event ConfigUpdated(uint256 indexed configId_, (uint256,address,uint256,uint256,uint256) config_)
func (_CoinFlipper *CoinFlipperFilterer) WatchConfigUpdated(opts *bind.WatchOpts, sink chan<- *CoinFlipperConfigUpdated, configId_ []*big.Int) (event.Subscription, error) {

	var configId_Rule []interface{}
	for _, configId_Item := range configId_ {
		configId_Rule = append(configId_Rule, configId_Item)
	}

	logs, sub, err := _CoinFlipper.contract.WatchLogs(opts, "ConfigUpdated", configId_Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinFlipperConfigUpdated)
				if err := _CoinFlipper.contract.UnpackLog(event, "ConfigUpdated", log); err != nil {
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

// ParseConfigUpdated is a log parse operation binding the contract event 0x606bf8e18a5e25fc2918dad277ffe9ec4b87e586deb4a7745b6df95b46e1b134.
//
// Solidity: event ConfigUpdated(uint256 indexed configId_, (uint256,address,uint256,uint256,uint256) config_)
func (_CoinFlipper *CoinFlipperFilterer) ParseConfigUpdated(log types.Log) (*CoinFlipperConfigUpdated, error) {
	event := new(CoinFlipperConfigUpdated)
	if err := _CoinFlipper.contract.UnpackLog(event, "ConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoinFlipperInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CoinFlipper contract.
type CoinFlipperInitializedIterator struct {
	Event *CoinFlipperInitialized // Event containing the contract specifics and raw log

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
func (it *CoinFlipperInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinFlipperInitialized)
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
		it.Event = new(CoinFlipperInitialized)
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
func (it *CoinFlipperInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinFlipperInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinFlipperInitialized represents a Initialized event raised by the CoinFlipper contract.
type CoinFlipperInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CoinFlipper *CoinFlipperFilterer) FilterInitialized(opts *bind.FilterOpts) (*CoinFlipperInitializedIterator, error) {

	logs, sub, err := _CoinFlipper.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CoinFlipperInitializedIterator{contract: _CoinFlipper.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CoinFlipper *CoinFlipperFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CoinFlipperInitialized) (event.Subscription, error) {

	logs, sub, err := _CoinFlipper.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinFlipperInitialized)
				if err := _CoinFlipper.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CoinFlipper *CoinFlipperFilterer) ParseInitialized(log types.Log) (*CoinFlipperInitialized, error) {
	event := new(CoinFlipperInitialized)
	if err := _CoinFlipper.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoinFlipperOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CoinFlipper contract.
type CoinFlipperOwnershipTransferredIterator struct {
	Event *CoinFlipperOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CoinFlipperOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinFlipperOwnershipTransferred)
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
		it.Event = new(CoinFlipperOwnershipTransferred)
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
func (it *CoinFlipperOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinFlipperOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinFlipperOwnershipTransferred represents a OwnershipTransferred event raised by the CoinFlipper contract.
type CoinFlipperOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CoinFlipper *CoinFlipperFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CoinFlipperOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CoinFlipper.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CoinFlipperOwnershipTransferredIterator{contract: _CoinFlipper.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CoinFlipper *CoinFlipperFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CoinFlipperOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CoinFlipper.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinFlipperOwnershipTransferred)
				if err := _CoinFlipper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CoinFlipper *CoinFlipperFilterer) ParseOwnershipTransferred(log types.Log) (*CoinFlipperOwnershipTransferred, error) {
	event := new(CoinFlipperOwnershipTransferred)
	if err := _CoinFlipper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
