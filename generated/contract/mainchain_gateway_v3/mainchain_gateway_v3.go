// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mainchain_gateway_v3

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

// MainchainGatewayV3MetaData contains all meta data concerning the MainchainGatewayV3 contract.
var MainchainGatewayV3MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAWAL_UNLOCKER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_MAX_PERCENTAGE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkHighTierVoteWeightThreshold\",\"inputs\":[{\"name\":\"_voteWeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkThreshold\",\"inputs\":[{\"name\":\"_voteWeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dailyWithdrawalLimit\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"emergencyPauser\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContract\",\"inputs\":[{\"name\":\"contractType\",\"type\":\"uint8\",\"internalType\":\"enumContractType\"}],\"outputs\":[{\"name\":\"contract_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getHighTierVoteWeightThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMember\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMemberCount\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoninToken\",\"inputs\":[{\"name\":\"mainchainToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"token\",\"type\":\"tuple\",\"internalType\":\"structMappedTokenConsumer.MappedToken\",\"components\":[{\"name\":\"erc\",\"type\":\"uint8\",\"internalType\":\"enumTokenStandard\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"num_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWhitelistedAddresses\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"whitelisteds\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"highTierThreshold\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initializeV5\",\"inputs\":[{\"name\":\"migrator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newEmergencyPauser\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"recipients\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastDateSynced\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastSyncedWithdrawal\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockedThreshold\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mapTokens\",\"inputs\":[{\"name\":\"_mainchainTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_roninTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_standards\",\"type\":\"uint8[]\",\"internalType\":\"enumTokenStandard[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mapTokensAndThresholds\",\"inputs\":[{\"name\":\"_mainchainTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_roninTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_standards\",\"type\":\"uint8[]\",\"internalType\":\"enumTokenStandard[]\"},{\"name\":\"_thresholds\",\"type\":\"uint256[][4]\",\"internalType\":\"uint256[][4]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"migrateERC20\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"minimumVoteWeight\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onBridgeOperatorsAdded\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"weights\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"addeds\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onBridgeOperatorsRemoved\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removeds\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onERC1155BatchReceived\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onERC1155Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reachedWithdrawalLimit\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestDepositFor\",\"inputs\":[{\"name\":\"_request\",\"type\":\"tuple\",\"internalType\":\"structTransfer.Request\",\"components\":[{\"name\":\"recipientAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"info\",\"type\":\"tuple\",\"internalType\":\"structTokenInfo\",\"components\":[{\"name\":\"erc\",\"type\":\"uint8\",\"internalType\":\"enumTokenStandard\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"restrict\",\"inputs\":[{\"name\":\"fnSig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"enumBitmap\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"restricted\",\"inputs\":[{\"name\":\"fnSig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"standard\",\"type\":\"uint8\",\"internalType\":\"enumTokenStandard\"}],\"outputs\":[{\"name\":\"yes\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"roninChainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setContract\",\"inputs\":[{\"name\":\"contractType\",\"type\":\"uint8\",\"internalType\":\"enumContractType\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDailyWithdrawalLimits\",\"inputs\":[{\"name\":\"_tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_limits\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEmergencyPauser\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setHighTierThresholds\",\"inputs\":[{\"name\":\"_tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_thresholds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setHighTierVoteWeightThreshold\",\"inputs\":[{\"name\":\"_numerator\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_denominator\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"_previousNum\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_previousDenom\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLockedThresholds\",\"inputs\":[{\"name\":\"_tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_thresholds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setThreshold\",\"inputs\":[{\"name\":\"num\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUnlockFeePercentages\",\"inputs\":[{\"name\":\"_tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_percentages\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWrappedNativeTokenContract\",\"inputs\":[{\"name\":\"_wrappedToken\",\"type\":\"address\",\"internalType\":\"contractIWETH\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitWithdrawal\",\"inputs\":[{\"name\":\"_receipt\",\"type\":\"tuple\",\"internalType\":\"structTransfer.Receipt\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumTransfer.Kind\"},{\"name\":\"mainchain\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"ronin\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"info\",\"type\":\"tuple\",\"internalType\":\"structTokenInfo\",\"components\":[{\"name\":\"erc\",\"type\":\"uint8\",\"internalType\":\"enumTokenStandard\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"_signatures\",\"type\":\"tuple[]\",\"internalType\":\"structSignatureConsumer.Signature[]\",\"components\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"_locked\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unlockFeePercentages\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unlockWithdrawal\",\"inputs\":[{\"name\":\"receipt\",\"type\":\"tuple\",\"internalType\":\"structTransfer.Receipt\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumTransfer.Kind\"},{\"name\":\"mainchain\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"ronin\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"info\",\"type\":\"tuple\",\"internalType\":\"structTokenInfo\",\"components\":[{\"name\":\"erc\",\"type\":\"uint8\",\"internalType\":\"enumTokenStandard\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"whitelist\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"recipients\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawalHash\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawalLocked\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"wrappedNativeToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIWETH\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ContractUpdated\",\"inputs\":[{\"name\":\"contractType\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumContractType\"},{\"name\":\"addr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DailyWithdrawalLimitsUpdated\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"limits\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositRequested\",\"inputs\":[{\"name\":\"receiptHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"receipt\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structTransfer.Receipt\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumTransfer.Kind\"},{\"name\":\"mainchain\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"ronin\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"info\",\"type\":\"tuple\",\"internalType\":\"structTokenInfo\",\"components\":[{\"name\":\"erc\",\"type\":\"uint8\",\"internalType\":\"enumTokenStandard\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"HighTierThresholdsUpdated\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"thresholds\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"HighTierVoteWeightThresholdUpdated\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"numerator\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"denominator\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"previousNumerator\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"previousDenominator\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedThresholdsUpdated\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"thresholds\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Restricted\",\"inputs\":[{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fnSig\",\"type\":\"bytes4\",\"indexed\":true,\"internalType\":\"bytes4\"},{\"name\":\"enumBitmap\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ThresholdUpdated\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"numerator\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"denominator\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"previousNumerator\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"previousDenominator\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenMapped\",\"inputs\":[{\"name\":\"mainchainTokens\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"roninTokens\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"standards\",\"type\":\"uint8[]\",\"indexed\":false,\"internalType\":\"enumTokenStandard[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnRestricted\",\"inputs\":[{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fnSig\",\"type\":\"bytes4\",\"indexed\":true,\"internalType\":\"bytes4\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnlockFeePercentagesUpdated\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"percentages\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WhitelistUpdated\",\"inputs\":[{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokens\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"recipients\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"indexed\":false,\"internalType\":\"uint64[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalLocked\",\"inputs\":[{\"name\":\"receiptHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"receipt\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structTransfer.Receipt\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumTransfer.Kind\"},{\"name\":\"mainchain\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"ronin\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"info\",\"type\":\"tuple\",\"internalType\":\"structTokenInfo\",\"components\":[{\"name\":\"erc\",\"type\":\"uint8\",\"internalType\":\"enumTokenStandard\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalUnlocked\",\"inputs\":[{\"name\":\"receiptHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"receipt\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structTransfer.Receipt\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumTransfer.Kind\"},{\"name\":\"mainchain\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"ronin\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"info\",\"type\":\"tuple\",\"internalType\":\"structTokenInfo\",\"components\":[{\"name\":\"erc\",\"type\":\"uint8\",\"internalType\":\"enumTokenStandard\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrew\",\"inputs\":[{\"name\":\"receiptHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"receipt\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structTransfer.Receipt\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumTransfer.Kind\"},{\"name\":\"mainchain\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"ronin\",\"type\":\"tuple\",\"internalType\":\"structTokenOwner\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"info\",\"type\":\"tuple\",\"internalType\":\"structTokenInfo\",\"components\":[{\"name\":\"erc\",\"type\":\"uint8\",\"internalType\":\"enumTokenStandard\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WrappedNativeTokenContractUpdated\",\"inputs\":[{\"name\":\"weth\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIWETH\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ErrContractTypeNotFound\",\"inputs\":[{\"name\":\"contractType\",\"type\":\"uint8\",\"internalType\":\"enumContractType\"}]},{\"type\":\"error\",\"name\":\"ErrERC1155MintingFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrERC20MintingFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrERC721MintingFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrEmptyArray\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrInvalidChainId\",\"inputs\":[{\"name\":\"msgSig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ErrInvalidInfo\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrInvalidOrder\",\"inputs\":[{\"name\":\"msgSig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"ErrInvalidPercentage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrInvalidReceipt\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrInvalidReceiptKind\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrInvalidRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrInvalidSigner\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sig\",\"type\":\"tuple\",\"internalType\":\"structSignatureConsumer.Signature\",\"components\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]},{\"type\":\"error\",\"name\":\"ErrInvalidThreshold\",\"inputs\":[{\"name\":\"msgSig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"ErrInvalidTokenStandard\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrLengthMismatch\",\"inputs\":[{\"name\":\"msgSig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"ErrNotWhitelistedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ErrNullHighTierVoteWeightProvided\",\"inputs\":[{\"name\":\"msgSig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"ErrNullMinVoteWeightProvided\",\"inputs\":[{\"name\":\"msgSig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"ErrNullTotalWeightProvided\",\"inputs\":[{\"name\":\"msgSig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"ErrQueryForApprovedWithdrawal\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrQueryForInsufficientVoteWeight\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrQueryForProcessedWithdrawal\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrReachedDailyWithdrawalLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrRestricted\",\"inputs\":[{\"name\":\"fnSig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"standard\",\"type\":\"uint8\",\"internalType\":\"enumTokenStandard\"}]},{\"type\":\"error\",\"name\":\"ErrTokenCouldNotTransfer\",\"inputs\":[{\"name\":\"tokenInfo\",\"type\":\"tuple\",\"internalType\":\"structTokenInfo\",\"components\":[{\"name\":\"erc\",\"type\":\"uint8\",\"internalType\":\"enumTokenStandard\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ErrTokenCouldNotTransferFrom\",\"inputs\":[{\"name\":\"tokenInfo\",\"type\":\"tuple\",\"internalType\":\"structTokenInfo\",\"components\":[{\"name\":\"erc\",\"type\":\"uint8\",\"internalType\":\"enumTokenStandard\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ErrUnauthorized\",\"inputs\":[{\"name\":\"msgSig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"expectedRole\",\"type\":\"uint8\",\"internalType\":\"enumRoleAccess\"}]},{\"type\":\"error\",\"name\":\"ErrUnexpectedInternalCall\",\"inputs\":[{\"name\":\"msgSig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"expectedContractType\",\"type\":\"uint8\",\"internalType\":\"enumContractType\"},{\"name\":\"actual\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ErrUnsupportedStandard\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrUnsupportedToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrWhitelistWrappedTokenInstead\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ErrZeroCodeContract\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// MainchainGatewayV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use MainchainGatewayV3MetaData.ABI instead.
var MainchainGatewayV3ABI = MainchainGatewayV3MetaData.ABI

// MainchainGatewayV3 is an auto generated Go binding around an Ethereum contract.
type MainchainGatewayV3 struct {
	MainchainGatewayV3Caller     // Read-only binding to the contract
	MainchainGatewayV3Transactor // Write-only binding to the contract
	MainchainGatewayV3Filterer   // Log filterer for contract events
}

// MainchainGatewayV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type MainchainGatewayV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainchainGatewayV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MainchainGatewayV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainchainGatewayV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainchainGatewayV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainchainGatewayV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainchainGatewayV3Session struct {
	Contract     *MainchainGatewayV3 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MainchainGatewayV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainchainGatewayV3CallerSession struct {
	Contract *MainchainGatewayV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// MainchainGatewayV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainchainGatewayV3TransactorSession struct {
	Contract     *MainchainGatewayV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// MainchainGatewayV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type MainchainGatewayV3Raw struct {
	Contract *MainchainGatewayV3 // Generic contract binding to access the raw methods on
}

// MainchainGatewayV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainchainGatewayV3CallerRaw struct {
	Contract *MainchainGatewayV3Caller // Generic read-only contract binding to access the raw methods on
}

// MainchainGatewayV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainchainGatewayV3TransactorRaw struct {
	Contract *MainchainGatewayV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMainchainGatewayV3 creates a new instance of MainchainGatewayV3, bound to a specific deployed contract.
func NewMainchainGatewayV3(address common.Address, backend bind.ContractBackend) (*MainchainGatewayV3, error) {
	contract, err := bindMainchainGatewayV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayV3{MainchainGatewayV3Caller: MainchainGatewayV3Caller{contract: contract}, MainchainGatewayV3Transactor: MainchainGatewayV3Transactor{contract: contract}, MainchainGatewayV3Filterer: MainchainGatewayV3Filterer{contract: contract}}, nil
}

// NewMainchainGatewayV3Caller creates a new read-only instance of MainchainGatewayV3, bound to a specific deployed contract.
func NewMainchainGatewayV3Caller(address common.Address, caller bind.ContractCaller) (*MainchainGatewayV3Caller, error) {
	contract, err := bindMainchainGatewayV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayV3Caller{contract: contract}, nil
}

// NewMainchainGatewayV3Transactor creates a new write-only instance of MainchainGatewayV3, bound to a specific deployed contract.
func NewMainchainGatewayV3Transactor(address common.Address, transactor bind.ContractTransactor) (*MainchainGatewayV3Transactor, error) {
	contract, err := bindMainchainGatewayV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayV3Transactor{contract: contract}, nil
}

// NewMainchainGatewayV3Filterer creates a new log filterer instance of MainchainGatewayV3, bound to a specific deployed contract.
func NewMainchainGatewayV3Filterer(address common.Address, filterer bind.ContractFilterer) (*MainchainGatewayV3Filterer, error) {
	contract, err := bindMainchainGatewayV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayV3Filterer{contract: contract}, nil
}

// bindMainchainGatewayV3 binds a generic wrapper to an already deployed contract.
func bindMainchainGatewayV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MainchainGatewayV3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MainchainGatewayV3 *MainchainGatewayV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MainchainGatewayV3.Contract.MainchainGatewayV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MainchainGatewayV3 *MainchainGatewayV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.MainchainGatewayV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MainchainGatewayV3 *MainchainGatewayV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.MainchainGatewayV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MainchainGatewayV3 *MainchainGatewayV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MainchainGatewayV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MainchainGatewayV3 *MainchainGatewayV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MainchainGatewayV3 *MainchainGatewayV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MainchainGatewayV3 *MainchainGatewayV3Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MainchainGatewayV3.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _MainchainGatewayV3.Contract.DEFAULTADMINROLE(&_MainchainGatewayV3.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MainchainGatewayV3 *MainchainGatewayV3CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MainchainGatewayV3.Contract.DEFAULTADMINROLE(&_MainchainGatewayV3.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_MainchainGatewayV3 *MainchainGatewayV3Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MainchainGatewayV3.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _MainchainGatewayV3.Contract.DOMAINSEPARATOR(&_MainchainGatewayV3.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_MainchainGatewayV3 *MainchainGatewayV3CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _MainchainGatewayV3.Contract.DOMAINSEPARATOR(&_MainchainGatewayV3.CallOpts)
}

// WITHDRAWALUNLOCKERROLE is a free data retrieval call binding the contract method 0x8f34e347.
//
// Solidity: function WITHDRAWAL_UNLOCKER_ROLE() view returns(bytes32)
func (_MainchainGatewayV3 *MainchainGatewayV3Caller) WITHDRAWALUNLOCKERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MainchainGatewayV3.contract.Call(opts, &out, "WITHDRAWAL_UNLOCKER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWALUNLOCKERROLE is a free data retrieval call binding the contract method 0x8f34e347.
//
// Solidity: function WITHDRAWAL_UNLOCKER_ROLE() view returns(bytes32)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) WITHDRAWALUNLOCKERROLE() ([32]byte, error) {
	return _MainchainGatewayV3.Contract.WITHDRAWALUNLOCKERROLE(&_MainchainGatewayV3.CallOpts)
}

// WITHDRAWALUNLOCKERROLE is a free data retrieval call binding the contract method 0x8f34e347.
//
// Solidity: function WITHDRAWAL_UNLOCKER_ROLE() view returns(bytes32)
func (_MainchainGatewayV3 *MainchainGatewayV3CallerSession) WITHDRAWALUNLOCKERROLE() ([32]byte, error) {
	return _MainchainGatewayV3.Contract.WITHDRAWALUNLOCKERROLE(&_MainchainGatewayV3.CallOpts)
}

// MAXPERCENTAGE is a free data retrieval call binding the contract method 0x302d12db.
//
// Solidity: function _MAX_PERCENTAGE() view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3Caller) MAXPERCENTAGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MainchainGatewayV3.contract.Call(opts, &out, "_MAX_PERCENTAGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPERCENTAGE is a free data retrieval call binding the contract method 0x302d12db.
//
// Solidity: function _MAX_PERCENTAGE() view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) MAXPERCENTAGE() (*big.Int, error) {
	return _MainchainGatewayV3.Contract.MAXPERCENTAGE(&_MainchainGatewayV3.CallOpts)
}

// MAXPERCENTAGE is a free data retrieval call binding the contract method 0x302d12db.
//
// Solidity: function _MAX_PERCENTAGE() view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3CallerSession) MAXPERCENTAGE() (*big.Int, error) {
	return _MainchainGatewayV3.Contract.MAXPERCENTAGE(&_MainchainGatewayV3.CallOpts)
}

// CheckHighTierVoteWeightThreshold is a free data retrieval call binding the contract method 0xac78dfe8.
//
// Solidity: function checkHighTierVoteWeightThreshold(uint256 _voteWeight) view returns(bool)
func (_MainchainGatewayV3 *MainchainGatewayV3Caller) CheckHighTierVoteWeightThreshold(opts *bind.CallOpts, _voteWeight *big.Int) (bool, error) {
	var out []interface{}
	err := _MainchainGatewayV3.contract.Call(opts, &out, "checkHighTierVoteWeightThreshold", _voteWeight)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckHighTierVoteWeightThreshold is a free data retrieval call binding the contract method 0xac78dfe8.
//
// Solidity: function checkHighTierVoteWeightThreshold(uint256 _voteWeight) view returns(bool)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) CheckHighTierVoteWeightThreshold(_voteWeight *big.Int) (bool, error) {
	return _MainchainGatewayV3.Contract.CheckHighTierVoteWeightThreshold(&_MainchainGatewayV3.CallOpts, _voteWeight)
}

// CheckHighTierVoteWeightThreshold is a free data retrieval call binding the contract method 0xac78dfe8.
//
// Solidity: function checkHighTierVoteWeightThreshold(uint256 _voteWeight) view returns(bool)
func (_MainchainGatewayV3 *MainchainGatewayV3CallerSession) CheckHighTierVoteWeightThreshold(_voteWeight *big.Int) (bool, error) {
	return _MainchainGatewayV3.Contract.CheckHighTierVoteWeightThreshold(&_MainchainGatewayV3.CallOpts, _voteWeight)
}

// CheckThreshold is a free data retrieval call binding the contract method 0xdafae408.
//
// Solidity: function checkThreshold(uint256 _voteWeight) view returns(bool)
func (_MainchainGatewayV3 *MainchainGatewayV3Caller) CheckThreshold(opts *bind.CallOpts, _voteWeight *big.Int) (bool, error) {
	var out []interface{}
	err := _MainchainGatewayV3.contract.Call(opts, &out, "checkThreshold", _voteWeight)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckThreshold is a free data retrieval call binding the contract method 0xdafae408.
//
// Solidity: function checkThreshold(uint256 _voteWeight) view returns(bool)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) CheckThreshold(_voteWeight *big.Int) (bool, error) {
	return _MainchainGatewayV3.Contract.CheckThreshold(&_MainchainGatewayV3.CallOpts, _voteWeight)
}

// CheckThreshold is a free data retrieval call binding the contract method 0xdafae408.
//
// Solidity: function checkThreshold(uint256 _voteWeight) view returns(bool)
func (_MainchainGatewayV3 *MainchainGatewayV3CallerSession) CheckThreshold(_voteWeight *big.Int) (bool, error) {
	return _MainchainGatewayV3.Contract.CheckThreshold(&_MainchainGatewayV3.CallOpts, _voteWeight)
}

// DailyWithdrawalLimit is a free data retrieval call binding the contract method 0xab796566.
//
// Solidity: function dailyWithdrawalLimit(address ) view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3Caller) DailyWithdrawalLimit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MainchainGatewayV3.contract.Call(opts, &out, "dailyWithdrawalLimit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DailyWithdrawalLimit is a free data retrieval call binding the contract method 0xab796566.
//
// Solidity: function dailyWithdrawalLimit(address ) view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) DailyWithdrawalLimit(arg0 common.Address) (*big.Int, error) {
	return _MainchainGatewayV3.Contract.DailyWithdrawalLimit(&_MainchainGatewayV3.CallOpts, arg0)
}

// DailyWithdrawalLimit is a free data retrieval call binding the contract method 0xab796566.
//
// Solidity: function dailyWithdrawalLimit(address ) view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3CallerSession) DailyWithdrawalLimit(arg0 common.Address) (*big.Int, error) {
	return _MainchainGatewayV3.Contract.DailyWithdrawalLimit(&_MainchainGatewayV3.CallOpts, arg0)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3Caller) DepositCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MainchainGatewayV3.contract.Call(opts, &out, "depositCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) DepositCount() (*big.Int, error) {
	return _MainchainGatewayV3.Contract.DepositCount(&_MainchainGatewayV3.CallOpts)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3CallerSession) DepositCount() (*big.Int, error) {
	return _MainchainGatewayV3.Contract.DepositCount(&_MainchainGatewayV3.CallOpts)
}

// EmergencyPauser is a free data retrieval call binding the contract method 0x065b3adf.
//
// Solidity: function emergencyPauser() view returns(address)
func (_MainchainGatewayV3 *MainchainGatewayV3Caller) EmergencyPauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MainchainGatewayV3.contract.Call(opts, &out, "emergencyPauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EmergencyPauser is a free data retrieval call binding the contract method 0x065b3adf.
//
// Solidity: function emergencyPauser() view returns(address)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) EmergencyPauser() (common.Address, error) {
	return _MainchainGatewayV3.Contract.EmergencyPauser(&_MainchainGatewayV3.CallOpts)
}

// EmergencyPauser is a free data retrieval call binding the contract method 0x065b3adf.
//
// Solidity: function emergencyPauser() view returns(address)
func (_MainchainGatewayV3 *MainchainGatewayV3CallerSession) EmergencyPauser() (common.Address, error) {
	return _MainchainGatewayV3.Contract.EmergencyPauser(&_MainchainGatewayV3.CallOpts)
}

// GetContract is a free data retrieval call binding the contract method 0xde981f1b.
//
// Solidity: function getContract(uint8 contractType) view returns(address contract_)
func (_MainchainGatewayV3 *MainchainGatewayV3Caller) GetContract(opts *bind.CallOpts, contractType uint8) (common.Address, error) {
	var out []interface{}
	err := _MainchainGatewayV3.contract.Call(opts, &out, "getContract", contractType)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetContract is a free data retrieval call binding the contract method 0xde981f1b.
//
// Solidity: function getContract(uint8 contractType) view returns(address contract_)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) GetContract(contractType uint8) (common.Address, error) {
	return _MainchainGatewayV3.Contract.GetContract(&_MainchainGatewayV3.CallOpts, contractType)
}

// GetContract is a free data retrieval call binding the contract method 0xde981f1b.
//
// Solidity: function getContract(uint8 contractType) view returns(address contract_)
func (_MainchainGatewayV3 *MainchainGatewayV3CallerSession) GetContract(contractType uint8) (common.Address, error) {
	return _MainchainGatewayV3.Contract.GetContract(&_MainchainGatewayV3.CallOpts, contractType)
}

// GetHighTierVoteWeightThreshold is a free data retrieval call binding the contract method 0xcdb67444.
//
// Solidity: function getHighTierVoteWeightThreshold() view returns(uint256, uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3Caller) GetHighTierVoteWeightThreshold(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _MainchainGatewayV3.contract.Call(opts, &out, "getHighTierVoteWeightThreshold")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetHighTierVoteWeightThreshold is a free data retrieval call binding the contract method 0xcdb67444.
//
// Solidity: function getHighTierVoteWeightThreshold() view returns(uint256, uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) GetHighTierVoteWeightThreshold() (*big.Int, *big.Int, error) {
	return _MainchainGatewayV3.Contract.GetHighTierVoteWeightThreshold(&_MainchainGatewayV3.CallOpts)
}

// GetHighTierVoteWeightThreshold is a free data retrieval call binding the contract method 0xcdb67444.
//
// Solidity: function getHighTierVoteWeightThreshold() view returns(uint256, uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3CallerSession) GetHighTierVoteWeightThreshold() (*big.Int, *big.Int, error) {
	return _MainchainGatewayV3.Contract.GetHighTierVoteWeightThreshold(&_MainchainGatewayV3.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MainchainGatewayV3 *MainchainGatewayV3Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MainchainGatewayV3.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MainchainGatewayV3.Contract.GetRoleAdmin(&_MainchainGatewayV3.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MainchainGatewayV3 *MainchainGatewayV3CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MainchainGatewayV3.Contract.GetRoleAdmin(&_MainchainGatewayV3.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_MainchainGatewayV3 *MainchainGatewayV3Caller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MainchainGatewayV3.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _MainchainGatewayV3.Contract.GetRoleMember(&_MainchainGatewayV3.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_MainchainGatewayV3 *MainchainGatewayV3CallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _MainchainGatewayV3.Contract.GetRoleMember(&_MainchainGatewayV3.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3Caller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _MainchainGatewayV3.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _MainchainGatewayV3.Contract.GetRoleMemberCount(&_MainchainGatewayV3.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3CallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _MainchainGatewayV3.Contract.GetRoleMemberCount(&_MainchainGatewayV3.CallOpts, role)
}

// GetRoninToken is a free data retrieval call binding the contract method 0xb2975794.
//
// Solidity: function getRoninToken(address mainchainToken) view returns((uint8,address) token)
func (_MainchainGatewayV3 *MainchainGatewayV3Caller) GetRoninToken(opts *bind.CallOpts, mainchainToken common.Address) (MappedTokenConsumerMappedToken, error) {
	var out []interface{}
	err := _MainchainGatewayV3.contract.Call(opts, &out, "getRoninToken", mainchainToken)

	if err != nil {
		return *new(MappedTokenConsumerMappedToken), err
	}

	out0 := *abi.ConvertType(out[0], new(MappedTokenConsumerMappedToken)).(*MappedTokenConsumerMappedToken)

	return out0, err

}

// GetRoninToken is a free data retrieval call binding the contract method 0xb2975794.
//
// Solidity: function getRoninToken(address mainchainToken) view returns((uint8,address) token)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) GetRoninToken(mainchainToken common.Address) (MappedTokenConsumerMappedToken, error) {
	return _MainchainGatewayV3.Contract.GetRoninToken(&_MainchainGatewayV3.CallOpts, mainchainToken)
}

// GetRoninToken is a free data retrieval call binding the contract method 0xb2975794.
//
// Solidity: function getRoninToken(address mainchainToken) view returns((uint8,address) token)
func (_MainchainGatewayV3 *MainchainGatewayV3CallerSession) GetRoninToken(mainchainToken common.Address) (MappedTokenConsumerMappedToken, error) {
	return _MainchainGatewayV3.Contract.GetRoninToken(&_MainchainGatewayV3.CallOpts, mainchainToken)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256 num_, uint256 denom_)
func (_MainchainGatewayV3 *MainchainGatewayV3Caller) GetThreshold(opts *bind.CallOpts) (struct {
	Num   *big.Int
	Denom *big.Int
}, error) {
	var out []interface{}
	err := _MainchainGatewayV3.contract.Call(opts, &out, "getThreshold")

	outstruct := new(struct {
		Num   *big.Int
		Denom *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Num = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Denom = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256 num_, uint256 denom_)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) GetThreshold() (struct {
	Num   *big.Int
	Denom *big.Int
}, error) {
	return _MainchainGatewayV3.Contract.GetThreshold(&_MainchainGatewayV3.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256 num_, uint256 denom_)
func (_MainchainGatewayV3 *MainchainGatewayV3CallerSession) GetThreshold() (struct {
	Num   *big.Int
	Denom *big.Int
}, error) {
	return _MainchainGatewayV3.Contract.GetThreshold(&_MainchainGatewayV3.CallOpts)
}

// GetWhitelistedAddresses is a free data retrieval call binding the contract method 0x6a10be12.
//
// Solidity: function getWhitelistedAddresses(address[] tokens) view returns(address[] whitelisteds, uint64[] remoteChainSelectors)
func (_MainchainGatewayV3 *MainchainGatewayV3Caller) GetWhitelistedAddresses(opts *bind.CallOpts, tokens []common.Address) (struct {
	Whitelisteds         []common.Address
	RemoteChainSelectors []uint64
}, error) {
	var out []interface{}
	err := _MainchainGatewayV3.contract.Call(opts, &out, "getWhitelistedAddresses", tokens)

	outstruct := new(struct {
		Whitelisteds         []common.Address
		RemoteChainSelectors []uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Whitelisteds = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.RemoteChainSelectors = *abi.ConvertType(out[1], new([]uint64)).(*[]uint64)

	return *outstruct, err

}

// GetWhitelistedAddresses is a free data retrieval call binding the contract method 0x6a10be12.
//
// Solidity: function getWhitelistedAddresses(address[] tokens) view returns(address[] whitelisteds, uint64[] remoteChainSelectors)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) GetWhitelistedAddresses(tokens []common.Address) (struct {
	Whitelisteds         []common.Address
	RemoteChainSelectors []uint64
}, error) {
	return _MainchainGatewayV3.Contract.GetWhitelistedAddresses(&_MainchainGatewayV3.CallOpts, tokens)
}

// GetWhitelistedAddresses is a free data retrieval call binding the contract method 0x6a10be12.
//
// Solidity: function getWhitelistedAddresses(address[] tokens) view returns(address[] whitelisteds, uint64[] remoteChainSelectors)
func (_MainchainGatewayV3 *MainchainGatewayV3CallerSession) GetWhitelistedAddresses(tokens []common.Address) (struct {
	Whitelisteds         []common.Address
	RemoteChainSelectors []uint64
}, error) {
	return _MainchainGatewayV3.Contract.GetWhitelistedAddresses(&_MainchainGatewayV3.CallOpts, tokens)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MainchainGatewayV3 *MainchainGatewayV3Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _MainchainGatewayV3.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MainchainGatewayV3.Contract.HasRole(&_MainchainGatewayV3.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MainchainGatewayV3 *MainchainGatewayV3CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MainchainGatewayV3.Contract.HasRole(&_MainchainGatewayV3.CallOpts, role, account)
}

// HighTierThreshold is a free data retrieval call binding the contract method 0xb1d08a03.
//
// Solidity: function highTierThreshold(address ) view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3Caller) HighTierThreshold(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MainchainGatewayV3.contract.Call(opts, &out, "highTierThreshold", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighTierThreshold is a free data retrieval call binding the contract method 0xb1d08a03.
//
// Solidity: function highTierThreshold(address ) view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) HighTierThreshold(arg0 common.Address) (*big.Int, error) {
	return _MainchainGatewayV3.Contract.HighTierThreshold(&_MainchainGatewayV3.CallOpts, arg0)
}

// HighTierThreshold is a free data retrieval call binding the contract method 0xb1d08a03.
//
// Solidity: function highTierThreshold(address ) view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3CallerSession) HighTierThreshold(arg0 common.Address) (*big.Int, error) {
	return _MainchainGatewayV3.Contract.HighTierThreshold(&_MainchainGatewayV3.CallOpts, arg0)
}

// LastDateSynced is a free data retrieval call binding the contract method 0x1d4a7210.
//
// Solidity: function lastDateSynced(address ) view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3Caller) LastDateSynced(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MainchainGatewayV3.contract.Call(opts, &out, "lastDateSynced", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastDateSynced is a free data retrieval call binding the contract method 0x1d4a7210.
//
// Solidity: function lastDateSynced(address ) view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) LastDateSynced(arg0 common.Address) (*big.Int, error) {
	return _MainchainGatewayV3.Contract.LastDateSynced(&_MainchainGatewayV3.CallOpts, arg0)
}

// LastDateSynced is a free data retrieval call binding the contract method 0x1d4a7210.
//
// Solidity: function lastDateSynced(address ) view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3CallerSession) LastDateSynced(arg0 common.Address) (*big.Int, error) {
	return _MainchainGatewayV3.Contract.LastDateSynced(&_MainchainGatewayV3.CallOpts, arg0)
}

// LastSyncedWithdrawal is a free data retrieval call binding the contract method 0xd55ed103.
//
// Solidity: function lastSyncedWithdrawal(address ) view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3Caller) LastSyncedWithdrawal(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MainchainGatewayV3.contract.Call(opts, &out, "lastSyncedWithdrawal", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastSyncedWithdrawal is a free data retrieval call binding the contract method 0xd55ed103.
//
// Solidity: function lastSyncedWithdrawal(address ) view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) LastSyncedWithdrawal(arg0 common.Address) (*big.Int, error) {
	return _MainchainGatewayV3.Contract.LastSyncedWithdrawal(&_MainchainGatewayV3.CallOpts, arg0)
}

// LastSyncedWithdrawal is a free data retrieval call binding the contract method 0xd55ed103.
//
// Solidity: function lastSyncedWithdrawal(address ) view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3CallerSession) LastSyncedWithdrawal(arg0 common.Address) (*big.Int, error) {
	return _MainchainGatewayV3.Contract.LastSyncedWithdrawal(&_MainchainGatewayV3.CallOpts, arg0)
}

// LockedThreshold is a free data retrieval call binding the contract method 0x59122f6b.
//
// Solidity: function lockedThreshold(address ) view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3Caller) LockedThreshold(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MainchainGatewayV3.contract.Call(opts, &out, "lockedThreshold", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedThreshold is a free data retrieval call binding the contract method 0x59122f6b.
//
// Solidity: function lockedThreshold(address ) view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) LockedThreshold(arg0 common.Address) (*big.Int, error) {
	return _MainchainGatewayV3.Contract.LockedThreshold(&_MainchainGatewayV3.CallOpts, arg0)
}

// LockedThreshold is a free data retrieval call binding the contract method 0x59122f6b.
//
// Solidity: function lockedThreshold(address ) view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3CallerSession) LockedThreshold(arg0 common.Address) (*big.Int, error) {
	return _MainchainGatewayV3.Contract.LockedThreshold(&_MainchainGatewayV3.CallOpts, arg0)
}

// MinimumVoteWeight is a free data retrieval call binding the contract method 0x7de5dedd.
//
// Solidity: function minimumVoteWeight() view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3Caller) MinimumVoteWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MainchainGatewayV3.contract.Call(opts, &out, "minimumVoteWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumVoteWeight is a free data retrieval call binding the contract method 0x7de5dedd.
//
// Solidity: function minimumVoteWeight() view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) MinimumVoteWeight() (*big.Int, error) {
	return _MainchainGatewayV3.Contract.MinimumVoteWeight(&_MainchainGatewayV3.CallOpts)
}

// MinimumVoteWeight is a free data retrieval call binding the contract method 0x7de5dedd.
//
// Solidity: function minimumVoteWeight() view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3CallerSession) MinimumVoteWeight() (*big.Int, error) {
	return _MainchainGatewayV3.Contract.MinimumVoteWeight(&_MainchainGatewayV3.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3Caller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MainchainGatewayV3.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) Nonce() (*big.Int, error) {
	return _MainchainGatewayV3.Contract.Nonce(&_MainchainGatewayV3.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3CallerSession) Nonce() (*big.Int, error) {
	return _MainchainGatewayV3.Contract.Nonce(&_MainchainGatewayV3.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MainchainGatewayV3 *MainchainGatewayV3Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MainchainGatewayV3.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) Paused() (bool, error) {
	return _MainchainGatewayV3.Contract.Paused(&_MainchainGatewayV3.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MainchainGatewayV3 *MainchainGatewayV3CallerSession) Paused() (bool, error) {
	return _MainchainGatewayV3.Contract.Paused(&_MainchainGatewayV3.CallOpts)
}

// ReachedWithdrawalLimit is a free data retrieval call binding the contract method 0x6c1ce670.
//
// Solidity: function reachedWithdrawalLimit(address _token, uint256 _quantity) view returns(bool)
func (_MainchainGatewayV3 *MainchainGatewayV3Caller) ReachedWithdrawalLimit(opts *bind.CallOpts, _token common.Address, _quantity *big.Int) (bool, error) {
	var out []interface{}
	err := _MainchainGatewayV3.contract.Call(opts, &out, "reachedWithdrawalLimit", _token, _quantity)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ReachedWithdrawalLimit is a free data retrieval call binding the contract method 0x6c1ce670.
//
// Solidity: function reachedWithdrawalLimit(address _token, uint256 _quantity) view returns(bool)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) ReachedWithdrawalLimit(_token common.Address, _quantity *big.Int) (bool, error) {
	return _MainchainGatewayV3.Contract.ReachedWithdrawalLimit(&_MainchainGatewayV3.CallOpts, _token, _quantity)
}

// ReachedWithdrawalLimit is a free data retrieval call binding the contract method 0x6c1ce670.
//
// Solidity: function reachedWithdrawalLimit(address _token, uint256 _quantity) view returns(bool)
func (_MainchainGatewayV3 *MainchainGatewayV3CallerSession) ReachedWithdrawalLimit(_token common.Address, _quantity *big.Int) (bool, error) {
	return _MainchainGatewayV3.Contract.ReachedWithdrawalLimit(&_MainchainGatewayV3.CallOpts, _token, _quantity)
}

// Restricted is a free data retrieval call binding the contract method 0x8e77669a.
//
// Solidity: function restricted(bytes4 fnSig, uint8 standard) view returns(bool yes)
func (_MainchainGatewayV3 *MainchainGatewayV3Caller) Restricted(opts *bind.CallOpts, fnSig [4]byte, standard uint8) (bool, error) {
	var out []interface{}
	err := _MainchainGatewayV3.contract.Call(opts, &out, "restricted", fnSig, standard)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Restricted is a free data retrieval call binding the contract method 0x8e77669a.
//
// Solidity: function restricted(bytes4 fnSig, uint8 standard) view returns(bool yes)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) Restricted(fnSig [4]byte, standard uint8) (bool, error) {
	return _MainchainGatewayV3.Contract.Restricted(&_MainchainGatewayV3.CallOpts, fnSig, standard)
}

// Restricted is a free data retrieval call binding the contract method 0x8e77669a.
//
// Solidity: function restricted(bytes4 fnSig, uint8 standard) view returns(bool yes)
func (_MainchainGatewayV3 *MainchainGatewayV3CallerSession) Restricted(fnSig [4]byte, standard uint8) (bool, error) {
	return _MainchainGatewayV3.Contract.Restricted(&_MainchainGatewayV3.CallOpts, fnSig, standard)
}

// RoninChainId is a free data retrieval call binding the contract method 0x17ce2dd4.
//
// Solidity: function roninChainId() view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3Caller) RoninChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MainchainGatewayV3.contract.Call(opts, &out, "roninChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoninChainId is a free data retrieval call binding the contract method 0x17ce2dd4.
//
// Solidity: function roninChainId() view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) RoninChainId() (*big.Int, error) {
	return _MainchainGatewayV3.Contract.RoninChainId(&_MainchainGatewayV3.CallOpts)
}

// RoninChainId is a free data retrieval call binding the contract method 0x17ce2dd4.
//
// Solidity: function roninChainId() view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3CallerSession) RoninChainId() (*big.Int, error) {
	return _MainchainGatewayV3.Contract.RoninChainId(&_MainchainGatewayV3.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MainchainGatewayV3 *MainchainGatewayV3Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MainchainGatewayV3.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MainchainGatewayV3.Contract.SupportsInterface(&_MainchainGatewayV3.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MainchainGatewayV3 *MainchainGatewayV3CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MainchainGatewayV3.Contract.SupportsInterface(&_MainchainGatewayV3.CallOpts, interfaceId)
}

// UnlockFeePercentages is a free data retrieval call binding the contract method 0xd19773d2.
//
// Solidity: function unlockFeePercentages(address ) view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3Caller) UnlockFeePercentages(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MainchainGatewayV3.contract.Call(opts, &out, "unlockFeePercentages", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnlockFeePercentages is a free data retrieval call binding the contract method 0xd19773d2.
//
// Solidity: function unlockFeePercentages(address ) view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) UnlockFeePercentages(arg0 common.Address) (*big.Int, error) {
	return _MainchainGatewayV3.Contract.UnlockFeePercentages(&_MainchainGatewayV3.CallOpts, arg0)
}

// UnlockFeePercentages is a free data retrieval call binding the contract method 0xd19773d2.
//
// Solidity: function unlockFeePercentages(address ) view returns(uint256)
func (_MainchainGatewayV3 *MainchainGatewayV3CallerSession) UnlockFeePercentages(arg0 common.Address) (*big.Int, error) {
	return _MainchainGatewayV3.Contract.UnlockFeePercentages(&_MainchainGatewayV3.CallOpts, arg0)
}

// WithdrawalHash is a free data retrieval call binding the contract method 0x6932be98.
//
// Solidity: function withdrawalHash(uint256 ) view returns(bytes32)
func (_MainchainGatewayV3 *MainchainGatewayV3Caller) WithdrawalHash(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MainchainGatewayV3.contract.Call(opts, &out, "withdrawalHash", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WithdrawalHash is a free data retrieval call binding the contract method 0x6932be98.
//
// Solidity: function withdrawalHash(uint256 ) view returns(bytes32)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) WithdrawalHash(arg0 *big.Int) ([32]byte, error) {
	return _MainchainGatewayV3.Contract.WithdrawalHash(&_MainchainGatewayV3.CallOpts, arg0)
}

// WithdrawalHash is a free data retrieval call binding the contract method 0x6932be98.
//
// Solidity: function withdrawalHash(uint256 ) view returns(bytes32)
func (_MainchainGatewayV3 *MainchainGatewayV3CallerSession) WithdrawalHash(arg0 *big.Int) ([32]byte, error) {
	return _MainchainGatewayV3.Contract.WithdrawalHash(&_MainchainGatewayV3.CallOpts, arg0)
}

// WithdrawalLocked is a free data retrieval call binding the contract method 0x4d493f4e.
//
// Solidity: function withdrawalLocked(uint256 ) view returns(bool)
func (_MainchainGatewayV3 *MainchainGatewayV3Caller) WithdrawalLocked(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _MainchainGatewayV3.contract.Call(opts, &out, "withdrawalLocked", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawalLocked is a free data retrieval call binding the contract method 0x4d493f4e.
//
// Solidity: function withdrawalLocked(uint256 ) view returns(bool)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) WithdrawalLocked(arg0 *big.Int) (bool, error) {
	return _MainchainGatewayV3.Contract.WithdrawalLocked(&_MainchainGatewayV3.CallOpts, arg0)
}

// WithdrawalLocked is a free data retrieval call binding the contract method 0x4d493f4e.
//
// Solidity: function withdrawalLocked(uint256 ) view returns(bool)
func (_MainchainGatewayV3 *MainchainGatewayV3CallerSession) WithdrawalLocked(arg0 *big.Int) (bool, error) {
	return _MainchainGatewayV3.Contract.WithdrawalLocked(&_MainchainGatewayV3.CallOpts, arg0)
}

// WrappedNativeToken is a free data retrieval call binding the contract method 0x17fcb39b.
//
// Solidity: function wrappedNativeToken() view returns(address)
func (_MainchainGatewayV3 *MainchainGatewayV3Caller) WrappedNativeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MainchainGatewayV3.contract.Call(opts, &out, "wrappedNativeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedNativeToken is a free data retrieval call binding the contract method 0x17fcb39b.
//
// Solidity: function wrappedNativeToken() view returns(address)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) WrappedNativeToken() (common.Address, error) {
	return _MainchainGatewayV3.Contract.WrappedNativeToken(&_MainchainGatewayV3.CallOpts)
}

// WrappedNativeToken is a free data retrieval call binding the contract method 0x17fcb39b.
//
// Solidity: function wrappedNativeToken() view returns(address)
func (_MainchainGatewayV3 *MainchainGatewayV3CallerSession) WrappedNativeToken() (common.Address, error) {
	return _MainchainGatewayV3.Contract.WrappedNativeToken(&_MainchainGatewayV3.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MainchainGatewayV3.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.GrantRole(&_MainchainGatewayV3.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.GrantRole(&_MainchainGatewayV3.TransactOpts, role, account)
}

// InitializeV5 is a paid mutator transaction binding the contract method 0xa40727bc.
//
// Solidity: function initializeV5(address migrator, address newEmergencyPauser, address[] tokens, address[] recipients, uint64[] remoteChainSelectors) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Transactor) InitializeV5(opts *bind.TransactOpts, migrator common.Address, newEmergencyPauser common.Address, tokens []common.Address, recipients []common.Address, remoteChainSelectors []uint64) (*types.Transaction, error) {
	return _MainchainGatewayV3.contract.Transact(opts, "initializeV5", migrator, newEmergencyPauser, tokens, recipients, remoteChainSelectors)
}

// InitializeV5 is a paid mutator transaction binding the contract method 0xa40727bc.
//
// Solidity: function initializeV5(address migrator, address newEmergencyPauser, address[] tokens, address[] recipients, uint64[] remoteChainSelectors) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Session) InitializeV5(migrator common.Address, newEmergencyPauser common.Address, tokens []common.Address, recipients []common.Address, remoteChainSelectors []uint64) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.InitializeV5(&_MainchainGatewayV3.TransactOpts, migrator, newEmergencyPauser, tokens, recipients, remoteChainSelectors)
}

// InitializeV5 is a paid mutator transaction binding the contract method 0xa40727bc.
//
// Solidity: function initializeV5(address migrator, address newEmergencyPauser, address[] tokens, address[] recipients, uint64[] remoteChainSelectors) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3TransactorSession) InitializeV5(migrator common.Address, newEmergencyPauser common.Address, tokens []common.Address, recipients []common.Address, remoteChainSelectors []uint64) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.InitializeV5(&_MainchainGatewayV3.TransactOpts, migrator, newEmergencyPauser, tokens, recipients, remoteChainSelectors)
}

// MapTokens is a paid mutator transaction binding the contract method 0x1b6e7594.
//
// Solidity: function mapTokens(address[] _mainchainTokens, address[] _roninTokens, uint8[] _standards) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Transactor) MapTokens(opts *bind.TransactOpts, _mainchainTokens []common.Address, _roninTokens []common.Address, _standards []uint8) (*types.Transaction, error) {
	return _MainchainGatewayV3.contract.Transact(opts, "mapTokens", _mainchainTokens, _roninTokens, _standards)
}

// MapTokens is a paid mutator transaction binding the contract method 0x1b6e7594.
//
// Solidity: function mapTokens(address[] _mainchainTokens, address[] _roninTokens, uint8[] _standards) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Session) MapTokens(_mainchainTokens []common.Address, _roninTokens []common.Address, _standards []uint8) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.MapTokens(&_MainchainGatewayV3.TransactOpts, _mainchainTokens, _roninTokens, _standards)
}

// MapTokens is a paid mutator transaction binding the contract method 0x1b6e7594.
//
// Solidity: function mapTokens(address[] _mainchainTokens, address[] _roninTokens, uint8[] _standards) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3TransactorSession) MapTokens(_mainchainTokens []common.Address, _roninTokens []common.Address, _standards []uint8) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.MapTokens(&_MainchainGatewayV3.TransactOpts, _mainchainTokens, _roninTokens, _standards)
}

// MapTokensAndThresholds is a paid mutator transaction binding the contract method 0xdff525e1.
//
// Solidity: function mapTokensAndThresholds(address[] _mainchainTokens, address[] _roninTokens, uint8[] _standards, uint256[][4] _thresholds) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Transactor) MapTokensAndThresholds(opts *bind.TransactOpts, _mainchainTokens []common.Address, _roninTokens []common.Address, _standards []uint8, _thresholds [4][]*big.Int) (*types.Transaction, error) {
	return _MainchainGatewayV3.contract.Transact(opts, "mapTokensAndThresholds", _mainchainTokens, _roninTokens, _standards, _thresholds)
}

// MapTokensAndThresholds is a paid mutator transaction binding the contract method 0xdff525e1.
//
// Solidity: function mapTokensAndThresholds(address[] _mainchainTokens, address[] _roninTokens, uint8[] _standards, uint256[][4] _thresholds) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Session) MapTokensAndThresholds(_mainchainTokens []common.Address, _roninTokens []common.Address, _standards []uint8, _thresholds [4][]*big.Int) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.MapTokensAndThresholds(&_MainchainGatewayV3.TransactOpts, _mainchainTokens, _roninTokens, _standards, _thresholds)
}

// MapTokensAndThresholds is a paid mutator transaction binding the contract method 0xdff525e1.
//
// Solidity: function mapTokensAndThresholds(address[] _mainchainTokens, address[] _roninTokens, uint8[] _standards, uint256[][4] _thresholds) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3TransactorSession) MapTokensAndThresholds(_mainchainTokens []common.Address, _roninTokens []common.Address, _standards []uint8, _thresholds [4][]*big.Int) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.MapTokensAndThresholds(&_MainchainGatewayV3.TransactOpts, _mainchainTokens, _roninTokens, _standards, _thresholds)
}

// MigrateERC20 is a paid mutator transaction binding the contract method 0xddd3602f.
//
// Solidity: function migrateERC20(address[] tokens, uint256[] amounts) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Transactor) MigrateERC20(opts *bind.TransactOpts, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _MainchainGatewayV3.contract.Transact(opts, "migrateERC20", tokens, amounts)
}

// MigrateERC20 is a paid mutator transaction binding the contract method 0xddd3602f.
//
// Solidity: function migrateERC20(address[] tokens, uint256[] amounts) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Session) MigrateERC20(tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.MigrateERC20(&_MainchainGatewayV3.TransactOpts, tokens, amounts)
}

// MigrateERC20 is a paid mutator transaction binding the contract method 0xddd3602f.
//
// Solidity: function migrateERC20(address[] tokens, uint256[] amounts) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3TransactorSession) MigrateERC20(tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.MigrateERC20(&_MainchainGatewayV3.TransactOpts, tokens, amounts)
}

// OnBridgeOperatorsAdded is a paid mutator transaction binding the contract method 0x8f851d8a.
//
// Solidity: function onBridgeOperatorsAdded(address[] operators, uint96[] weights, bool[] addeds) returns(bytes4)
func (_MainchainGatewayV3 *MainchainGatewayV3Transactor) OnBridgeOperatorsAdded(opts *bind.TransactOpts, operators []common.Address, weights []*big.Int, addeds []bool) (*types.Transaction, error) {
	return _MainchainGatewayV3.contract.Transact(opts, "onBridgeOperatorsAdded", operators, weights, addeds)
}

// OnBridgeOperatorsAdded is a paid mutator transaction binding the contract method 0x8f851d8a.
//
// Solidity: function onBridgeOperatorsAdded(address[] operators, uint96[] weights, bool[] addeds) returns(bytes4)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) OnBridgeOperatorsAdded(operators []common.Address, weights []*big.Int, addeds []bool) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.OnBridgeOperatorsAdded(&_MainchainGatewayV3.TransactOpts, operators, weights, addeds)
}

// OnBridgeOperatorsAdded is a paid mutator transaction binding the contract method 0x8f851d8a.
//
// Solidity: function onBridgeOperatorsAdded(address[] operators, uint96[] weights, bool[] addeds) returns(bytes4)
func (_MainchainGatewayV3 *MainchainGatewayV3TransactorSession) OnBridgeOperatorsAdded(operators []common.Address, weights []*big.Int, addeds []bool) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.OnBridgeOperatorsAdded(&_MainchainGatewayV3.TransactOpts, operators, weights, addeds)
}

// OnBridgeOperatorsRemoved is a paid mutator transaction binding the contract method 0xc48549de.
//
// Solidity: function onBridgeOperatorsRemoved(address[] operators, bool[] removeds) returns(bytes4)
func (_MainchainGatewayV3 *MainchainGatewayV3Transactor) OnBridgeOperatorsRemoved(opts *bind.TransactOpts, operators []common.Address, removeds []bool) (*types.Transaction, error) {
	return _MainchainGatewayV3.contract.Transact(opts, "onBridgeOperatorsRemoved", operators, removeds)
}

// OnBridgeOperatorsRemoved is a paid mutator transaction binding the contract method 0xc48549de.
//
// Solidity: function onBridgeOperatorsRemoved(address[] operators, bool[] removeds) returns(bytes4)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) OnBridgeOperatorsRemoved(operators []common.Address, removeds []bool) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.OnBridgeOperatorsRemoved(&_MainchainGatewayV3.TransactOpts, operators, removeds)
}

// OnBridgeOperatorsRemoved is a paid mutator transaction binding the contract method 0xc48549de.
//
// Solidity: function onBridgeOperatorsRemoved(address[] operators, bool[] removeds) returns(bytes4)
func (_MainchainGatewayV3 *MainchainGatewayV3TransactorSession) OnBridgeOperatorsRemoved(operators []common.Address, removeds []bool) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.OnBridgeOperatorsRemoved(&_MainchainGatewayV3.TransactOpts, operators, removeds)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_MainchainGatewayV3 *MainchainGatewayV3Transactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _MainchainGatewayV3.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.OnERC1155BatchReceived(&_MainchainGatewayV3.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_MainchainGatewayV3 *MainchainGatewayV3TransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.OnERC1155BatchReceived(&_MainchainGatewayV3.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_MainchainGatewayV3 *MainchainGatewayV3Transactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _MainchainGatewayV3.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.OnERC1155Received(&_MainchainGatewayV3.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_MainchainGatewayV3 *MainchainGatewayV3TransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.OnERC1155Received(&_MainchainGatewayV3.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MainchainGatewayV3.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Session) Pause() (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.Pause(&_MainchainGatewayV3.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MainchainGatewayV3 *MainchainGatewayV3TransactorSession) Pause() (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.Pause(&_MainchainGatewayV3.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MainchainGatewayV3.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Session) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.RenounceRole(&_MainchainGatewayV3.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3TransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.RenounceRole(&_MainchainGatewayV3.TransactOpts, role, account)
}

// RequestDepositFor is a paid mutator transaction binding the contract method 0x4b14557e.
//
// Solidity: function requestDepositFor((address,address,(uint8,uint256,uint256)) _request) payable returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Transactor) RequestDepositFor(opts *bind.TransactOpts, _request TransferRequest) (*types.Transaction, error) {
	return _MainchainGatewayV3.contract.Transact(opts, "requestDepositFor", _request)
}

// RequestDepositFor is a paid mutator transaction binding the contract method 0x4b14557e.
//
// Solidity: function requestDepositFor((address,address,(uint8,uint256,uint256)) _request) payable returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Session) RequestDepositFor(_request TransferRequest) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.RequestDepositFor(&_MainchainGatewayV3.TransactOpts, _request)
}

// RequestDepositFor is a paid mutator transaction binding the contract method 0x4b14557e.
//
// Solidity: function requestDepositFor((address,address,(uint8,uint256,uint256)) _request) payable returns()
func (_MainchainGatewayV3 *MainchainGatewayV3TransactorSession) RequestDepositFor(_request TransferRequest) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.RequestDepositFor(&_MainchainGatewayV3.TransactOpts, _request)
}

// Restrict is a paid mutator transaction binding the contract method 0x0e898511.
//
// Solidity: function restrict(bytes4 fnSig, uint8 enumBitmap) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Transactor) Restrict(opts *bind.TransactOpts, fnSig [4]byte, enumBitmap uint8) (*types.Transaction, error) {
	return _MainchainGatewayV3.contract.Transact(opts, "restrict", fnSig, enumBitmap)
}

// Restrict is a paid mutator transaction binding the contract method 0x0e898511.
//
// Solidity: function restrict(bytes4 fnSig, uint8 enumBitmap) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Session) Restrict(fnSig [4]byte, enumBitmap uint8) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.Restrict(&_MainchainGatewayV3.TransactOpts, fnSig, enumBitmap)
}

// Restrict is a paid mutator transaction binding the contract method 0x0e898511.
//
// Solidity: function restrict(bytes4 fnSig, uint8 enumBitmap) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3TransactorSession) Restrict(fnSig [4]byte, enumBitmap uint8) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.Restrict(&_MainchainGatewayV3.TransactOpts, fnSig, enumBitmap)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MainchainGatewayV3.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.RevokeRole(&_MainchainGatewayV3.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.RevokeRole(&_MainchainGatewayV3.TransactOpts, role, account)
}

// SetContract is a paid mutator transaction binding the contract method 0x865e6fd3.
//
// Solidity: function setContract(uint8 contractType, address addr) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Transactor) SetContract(opts *bind.TransactOpts, contractType uint8, addr common.Address) (*types.Transaction, error) {
	return _MainchainGatewayV3.contract.Transact(opts, "setContract", contractType, addr)
}

// SetContract is a paid mutator transaction binding the contract method 0x865e6fd3.
//
// Solidity: function setContract(uint8 contractType, address addr) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Session) SetContract(contractType uint8, addr common.Address) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.SetContract(&_MainchainGatewayV3.TransactOpts, contractType, addr)
}

// SetContract is a paid mutator transaction binding the contract method 0x865e6fd3.
//
// Solidity: function setContract(uint8 contractType, address addr) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3TransactorSession) SetContract(contractType uint8, addr common.Address) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.SetContract(&_MainchainGatewayV3.TransactOpts, contractType, addr)
}

// SetDailyWithdrawalLimits is a paid mutator transaction binding the contract method 0xe400327c.
//
// Solidity: function setDailyWithdrawalLimits(address[] _tokens, uint256[] _limits) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Transactor) SetDailyWithdrawalLimits(opts *bind.TransactOpts, _tokens []common.Address, _limits []*big.Int) (*types.Transaction, error) {
	return _MainchainGatewayV3.contract.Transact(opts, "setDailyWithdrawalLimits", _tokens, _limits)
}

// SetDailyWithdrawalLimits is a paid mutator transaction binding the contract method 0xe400327c.
//
// Solidity: function setDailyWithdrawalLimits(address[] _tokens, uint256[] _limits) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Session) SetDailyWithdrawalLimits(_tokens []common.Address, _limits []*big.Int) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.SetDailyWithdrawalLimits(&_MainchainGatewayV3.TransactOpts, _tokens, _limits)
}

// SetDailyWithdrawalLimits is a paid mutator transaction binding the contract method 0xe400327c.
//
// Solidity: function setDailyWithdrawalLimits(address[] _tokens, uint256[] _limits) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3TransactorSession) SetDailyWithdrawalLimits(_tokens []common.Address, _limits []*big.Int) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.SetDailyWithdrawalLimits(&_MainchainGatewayV3.TransactOpts, _tokens, _limits)
}

// SetEmergencyPauser is a paid mutator transaction binding the contract method 0x3e70838b.
//
// Solidity: function setEmergencyPauser(address _addr) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Transactor) SetEmergencyPauser(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MainchainGatewayV3.contract.Transact(opts, "setEmergencyPauser", _addr)
}

// SetEmergencyPauser is a paid mutator transaction binding the contract method 0x3e70838b.
//
// Solidity: function setEmergencyPauser(address _addr) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Session) SetEmergencyPauser(_addr common.Address) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.SetEmergencyPauser(&_MainchainGatewayV3.TransactOpts, _addr)
}

// SetEmergencyPauser is a paid mutator transaction binding the contract method 0x3e70838b.
//
// Solidity: function setEmergencyPauser(address _addr) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3TransactorSession) SetEmergencyPauser(_addr common.Address) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.SetEmergencyPauser(&_MainchainGatewayV3.TransactOpts, _addr)
}

// SetHighTierThresholds is a paid mutator transaction binding the contract method 0x93c5678f.
//
// Solidity: function setHighTierThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Transactor) SetHighTierThresholds(opts *bind.TransactOpts, _tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _MainchainGatewayV3.contract.Transact(opts, "setHighTierThresholds", _tokens, _thresholds)
}

// SetHighTierThresholds is a paid mutator transaction binding the contract method 0x93c5678f.
//
// Solidity: function setHighTierThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Session) SetHighTierThresholds(_tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.SetHighTierThresholds(&_MainchainGatewayV3.TransactOpts, _tokens, _thresholds)
}

// SetHighTierThresholds is a paid mutator transaction binding the contract method 0x93c5678f.
//
// Solidity: function setHighTierThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3TransactorSession) SetHighTierThresholds(_tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.SetHighTierThresholds(&_MainchainGatewayV3.TransactOpts, _tokens, _thresholds)
}

// SetHighTierVoteWeightThreshold is a paid mutator transaction binding the contract method 0x9dcc4da3.
//
// Solidity: function setHighTierVoteWeightThreshold(uint256 _numerator, uint256 _denominator) returns(uint256 _previousNum, uint256 _previousDenom)
func (_MainchainGatewayV3 *MainchainGatewayV3Transactor) SetHighTierVoteWeightThreshold(opts *bind.TransactOpts, _numerator *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _MainchainGatewayV3.contract.Transact(opts, "setHighTierVoteWeightThreshold", _numerator, _denominator)
}

// SetHighTierVoteWeightThreshold is a paid mutator transaction binding the contract method 0x9dcc4da3.
//
// Solidity: function setHighTierVoteWeightThreshold(uint256 _numerator, uint256 _denominator) returns(uint256 _previousNum, uint256 _previousDenom)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) SetHighTierVoteWeightThreshold(_numerator *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.SetHighTierVoteWeightThreshold(&_MainchainGatewayV3.TransactOpts, _numerator, _denominator)
}

// SetHighTierVoteWeightThreshold is a paid mutator transaction binding the contract method 0x9dcc4da3.
//
// Solidity: function setHighTierVoteWeightThreshold(uint256 _numerator, uint256 _denominator) returns(uint256 _previousNum, uint256 _previousDenom)
func (_MainchainGatewayV3 *MainchainGatewayV3TransactorSession) SetHighTierVoteWeightThreshold(_numerator *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.SetHighTierVoteWeightThreshold(&_MainchainGatewayV3.TransactOpts, _numerator, _denominator)
}

// SetLockedThresholds is a paid mutator transaction binding the contract method 0x1a8e55b0.
//
// Solidity: function setLockedThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Transactor) SetLockedThresholds(opts *bind.TransactOpts, _tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _MainchainGatewayV3.contract.Transact(opts, "setLockedThresholds", _tokens, _thresholds)
}

// SetLockedThresholds is a paid mutator transaction binding the contract method 0x1a8e55b0.
//
// Solidity: function setLockedThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Session) SetLockedThresholds(_tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.SetLockedThresholds(&_MainchainGatewayV3.TransactOpts, _tokens, _thresholds)
}

// SetLockedThresholds is a paid mutator transaction binding the contract method 0x1a8e55b0.
//
// Solidity: function setLockedThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3TransactorSession) SetLockedThresholds(_tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.SetLockedThresholds(&_MainchainGatewayV3.TransactOpts, _tokens, _thresholds)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xb9c36209.
//
// Solidity: function setThreshold(uint256 num, uint256 denom) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Transactor) SetThreshold(opts *bind.TransactOpts, num *big.Int, denom *big.Int) (*types.Transaction, error) {
	return _MainchainGatewayV3.contract.Transact(opts, "setThreshold", num, denom)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xb9c36209.
//
// Solidity: function setThreshold(uint256 num, uint256 denom) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Session) SetThreshold(num *big.Int, denom *big.Int) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.SetThreshold(&_MainchainGatewayV3.TransactOpts, num, denom)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xb9c36209.
//
// Solidity: function setThreshold(uint256 num, uint256 denom) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3TransactorSession) SetThreshold(num *big.Int, denom *big.Int) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.SetThreshold(&_MainchainGatewayV3.TransactOpts, num, denom)
}

// SetUnlockFeePercentages is a paid mutator transaction binding the contract method 0xb1a2567e.
//
// Solidity: function setUnlockFeePercentages(address[] _tokens, uint256[] _percentages) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Transactor) SetUnlockFeePercentages(opts *bind.TransactOpts, _tokens []common.Address, _percentages []*big.Int) (*types.Transaction, error) {
	return _MainchainGatewayV3.contract.Transact(opts, "setUnlockFeePercentages", _tokens, _percentages)
}

// SetUnlockFeePercentages is a paid mutator transaction binding the contract method 0xb1a2567e.
//
// Solidity: function setUnlockFeePercentages(address[] _tokens, uint256[] _percentages) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Session) SetUnlockFeePercentages(_tokens []common.Address, _percentages []*big.Int) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.SetUnlockFeePercentages(&_MainchainGatewayV3.TransactOpts, _tokens, _percentages)
}

// SetUnlockFeePercentages is a paid mutator transaction binding the contract method 0xb1a2567e.
//
// Solidity: function setUnlockFeePercentages(address[] _tokens, uint256[] _percentages) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3TransactorSession) SetUnlockFeePercentages(_tokens []common.Address, _percentages []*big.Int) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.SetUnlockFeePercentages(&_MainchainGatewayV3.TransactOpts, _tokens, _percentages)
}

// SetWrappedNativeTokenContract is a paid mutator transaction binding the contract method 0xd64af2a6.
//
// Solidity: function setWrappedNativeTokenContract(address _wrappedToken) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Transactor) SetWrappedNativeTokenContract(opts *bind.TransactOpts, _wrappedToken common.Address) (*types.Transaction, error) {
	return _MainchainGatewayV3.contract.Transact(opts, "setWrappedNativeTokenContract", _wrappedToken)
}

// SetWrappedNativeTokenContract is a paid mutator transaction binding the contract method 0xd64af2a6.
//
// Solidity: function setWrappedNativeTokenContract(address _wrappedToken) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Session) SetWrappedNativeTokenContract(_wrappedToken common.Address) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.SetWrappedNativeTokenContract(&_MainchainGatewayV3.TransactOpts, _wrappedToken)
}

// SetWrappedNativeTokenContract is a paid mutator transaction binding the contract method 0xd64af2a6.
//
// Solidity: function setWrappedNativeTokenContract(address _wrappedToken) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3TransactorSession) SetWrappedNativeTokenContract(_wrappedToken common.Address) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.SetWrappedNativeTokenContract(&_MainchainGatewayV3.TransactOpts, _wrappedToken)
}

// SubmitWithdrawal is a paid mutator transaction binding the contract method 0x4d0d6673.
//
// Solidity: function submitWithdrawal((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) _receipt, (uint8,bytes32,bytes32)[] _signatures) returns(bool _locked)
func (_MainchainGatewayV3 *MainchainGatewayV3Transactor) SubmitWithdrawal(opts *bind.TransactOpts, _receipt TransferReceipt, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _MainchainGatewayV3.contract.Transact(opts, "submitWithdrawal", _receipt, _signatures)
}

// SubmitWithdrawal is a paid mutator transaction binding the contract method 0x4d0d6673.
//
// Solidity: function submitWithdrawal((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) _receipt, (uint8,bytes32,bytes32)[] _signatures) returns(bool _locked)
func (_MainchainGatewayV3 *MainchainGatewayV3Session) SubmitWithdrawal(_receipt TransferReceipt, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.SubmitWithdrawal(&_MainchainGatewayV3.TransactOpts, _receipt, _signatures)
}

// SubmitWithdrawal is a paid mutator transaction binding the contract method 0x4d0d6673.
//
// Solidity: function submitWithdrawal((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) _receipt, (uint8,bytes32,bytes32)[] _signatures) returns(bool _locked)
func (_MainchainGatewayV3 *MainchainGatewayV3TransactorSession) SubmitWithdrawal(_receipt TransferReceipt, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.SubmitWithdrawal(&_MainchainGatewayV3.TransactOpts, _receipt, _signatures)
}

// UnlockWithdrawal is a paid mutator transaction binding the contract method 0x9157921c.
//
// Solidity: function unlockWithdrawal((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Transactor) UnlockWithdrawal(opts *bind.TransactOpts, receipt TransferReceipt) (*types.Transaction, error) {
	return _MainchainGatewayV3.contract.Transact(opts, "unlockWithdrawal", receipt)
}

// UnlockWithdrawal is a paid mutator transaction binding the contract method 0x9157921c.
//
// Solidity: function unlockWithdrawal((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Session) UnlockWithdrawal(receipt TransferReceipt) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.UnlockWithdrawal(&_MainchainGatewayV3.TransactOpts, receipt)
}

// UnlockWithdrawal is a paid mutator transaction binding the contract method 0x9157921c.
//
// Solidity: function unlockWithdrawal((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3TransactorSession) UnlockWithdrawal(receipt TransferReceipt) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.UnlockWithdrawal(&_MainchainGatewayV3.TransactOpts, receipt)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MainchainGatewayV3.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Session) Unpause() (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.Unpause(&_MainchainGatewayV3.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MainchainGatewayV3 *MainchainGatewayV3TransactorSession) Unpause() (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.Unpause(&_MainchainGatewayV3.TransactOpts)
}

// Whitelist is a paid mutator transaction binding the contract method 0x87a43fa5.
//
// Solidity: function whitelist(address[] tokens, address[] recipients, uint64[] remoteChainSelectors) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Transactor) Whitelist(opts *bind.TransactOpts, tokens []common.Address, recipients []common.Address, remoteChainSelectors []uint64) (*types.Transaction, error) {
	return _MainchainGatewayV3.contract.Transact(opts, "whitelist", tokens, recipients, remoteChainSelectors)
}

// Whitelist is a paid mutator transaction binding the contract method 0x87a43fa5.
//
// Solidity: function whitelist(address[] tokens, address[] recipients, uint64[] remoteChainSelectors) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Session) Whitelist(tokens []common.Address, recipients []common.Address, remoteChainSelectors []uint64) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.Whitelist(&_MainchainGatewayV3.TransactOpts, tokens, recipients, remoteChainSelectors)
}

// Whitelist is a paid mutator transaction binding the contract method 0x87a43fa5.
//
// Solidity: function whitelist(address[] tokens, address[] recipients, uint64[] remoteChainSelectors) returns()
func (_MainchainGatewayV3 *MainchainGatewayV3TransactorSession) Whitelist(tokens []common.Address, recipients []common.Address, remoteChainSelectors []uint64) (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.Whitelist(&_MainchainGatewayV3.TransactOpts, tokens, recipients, remoteChainSelectors)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MainchainGatewayV3.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MainchainGatewayV3 *MainchainGatewayV3Session) Receive() (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.Receive(&_MainchainGatewayV3.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MainchainGatewayV3 *MainchainGatewayV3TransactorSession) Receive() (*types.Transaction, error) {
	return _MainchainGatewayV3.Contract.Receive(&_MainchainGatewayV3.TransactOpts)
}

// MainchainGatewayV3ContractUpdatedIterator is returned from FilterContractUpdated and is used to iterate over the raw logs and unpacked data for ContractUpdated events raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3ContractUpdatedIterator struct {
	Event *MainchainGatewayV3ContractUpdated // Event containing the contract specifics and raw log

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
func (it *MainchainGatewayV3ContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainGatewayV3ContractUpdated)
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
		it.Event = new(MainchainGatewayV3ContractUpdated)
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
func (it *MainchainGatewayV3ContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainGatewayV3ContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainGatewayV3ContractUpdated represents a ContractUpdated event raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3ContractUpdated struct {
	ContractType uint8
	Addr         common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterContractUpdated is a free log retrieval operation binding the contract event 0x865d1c228a8ea13709cfe61f346f7ff67f1bbd4a18ff31ad3e7147350d317c59.
//
// Solidity: event ContractUpdated(uint8 indexed contractType, address indexed addr)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) FilterContractUpdated(opts *bind.FilterOpts, contractType []uint8, addr []common.Address) (*MainchainGatewayV3ContractUpdatedIterator, error) {

	var contractTypeRule []interface{}
	for _, contractTypeItem := range contractType {
		contractTypeRule = append(contractTypeRule, contractTypeItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _MainchainGatewayV3.contract.FilterLogs(opts, "ContractUpdated", contractTypeRule, addrRule)
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayV3ContractUpdatedIterator{contract: _MainchainGatewayV3.contract, event: "ContractUpdated", logs: logs, sub: sub}, nil
}

// WatchContractUpdated is a free log subscription operation binding the contract event 0x865d1c228a8ea13709cfe61f346f7ff67f1bbd4a18ff31ad3e7147350d317c59.
//
// Solidity: event ContractUpdated(uint8 indexed contractType, address indexed addr)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) WatchContractUpdated(opts *bind.WatchOpts, sink chan<- *MainchainGatewayV3ContractUpdated, contractType []uint8, addr []common.Address) (event.Subscription, error) {

	var contractTypeRule []interface{}
	for _, contractTypeItem := range contractType {
		contractTypeRule = append(contractTypeRule, contractTypeItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _MainchainGatewayV3.contract.WatchLogs(opts, "ContractUpdated", contractTypeRule, addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainGatewayV3ContractUpdated)
				if err := _MainchainGatewayV3.contract.UnpackLog(event, "ContractUpdated", log); err != nil {
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

// ParseContractUpdated is a log parse operation binding the contract event 0x865d1c228a8ea13709cfe61f346f7ff67f1bbd4a18ff31ad3e7147350d317c59.
//
// Solidity: event ContractUpdated(uint8 indexed contractType, address indexed addr)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) ParseContractUpdated(log types.Log) (*MainchainGatewayV3ContractUpdated, error) {
	event := new(MainchainGatewayV3ContractUpdated)
	if err := _MainchainGatewayV3.contract.UnpackLog(event, "ContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainchainGatewayV3DailyWithdrawalLimitsUpdatedIterator is returned from FilterDailyWithdrawalLimitsUpdated and is used to iterate over the raw logs and unpacked data for DailyWithdrawalLimitsUpdated events raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3DailyWithdrawalLimitsUpdatedIterator struct {
	Event *MainchainGatewayV3DailyWithdrawalLimitsUpdated // Event containing the contract specifics and raw log

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
func (it *MainchainGatewayV3DailyWithdrawalLimitsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainGatewayV3DailyWithdrawalLimitsUpdated)
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
		it.Event = new(MainchainGatewayV3DailyWithdrawalLimitsUpdated)
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
func (it *MainchainGatewayV3DailyWithdrawalLimitsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainGatewayV3DailyWithdrawalLimitsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainGatewayV3DailyWithdrawalLimitsUpdated represents a DailyWithdrawalLimitsUpdated event raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3DailyWithdrawalLimitsUpdated struct {
	Tokens []common.Address
	Limits []*big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDailyWithdrawalLimitsUpdated is a free log retrieval operation binding the contract event 0xb5d2963614d72181b4df1f993d45b83edf42fa19710f0204217ba1b3e183bb73.
//
// Solidity: event DailyWithdrawalLimitsUpdated(address[] tokens, uint256[] limits)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) FilterDailyWithdrawalLimitsUpdated(opts *bind.FilterOpts) (*MainchainGatewayV3DailyWithdrawalLimitsUpdatedIterator, error) {

	logs, sub, err := _MainchainGatewayV3.contract.FilterLogs(opts, "DailyWithdrawalLimitsUpdated")
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayV3DailyWithdrawalLimitsUpdatedIterator{contract: _MainchainGatewayV3.contract, event: "DailyWithdrawalLimitsUpdated", logs: logs, sub: sub}, nil
}

// WatchDailyWithdrawalLimitsUpdated is a free log subscription operation binding the contract event 0xb5d2963614d72181b4df1f993d45b83edf42fa19710f0204217ba1b3e183bb73.
//
// Solidity: event DailyWithdrawalLimitsUpdated(address[] tokens, uint256[] limits)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) WatchDailyWithdrawalLimitsUpdated(opts *bind.WatchOpts, sink chan<- *MainchainGatewayV3DailyWithdrawalLimitsUpdated) (event.Subscription, error) {

	logs, sub, err := _MainchainGatewayV3.contract.WatchLogs(opts, "DailyWithdrawalLimitsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainGatewayV3DailyWithdrawalLimitsUpdated)
				if err := _MainchainGatewayV3.contract.UnpackLog(event, "DailyWithdrawalLimitsUpdated", log); err != nil {
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

// ParseDailyWithdrawalLimitsUpdated is a log parse operation binding the contract event 0xb5d2963614d72181b4df1f993d45b83edf42fa19710f0204217ba1b3e183bb73.
//
// Solidity: event DailyWithdrawalLimitsUpdated(address[] tokens, uint256[] limits)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) ParseDailyWithdrawalLimitsUpdated(log types.Log) (*MainchainGatewayV3DailyWithdrawalLimitsUpdated, error) {
	event := new(MainchainGatewayV3DailyWithdrawalLimitsUpdated)
	if err := _MainchainGatewayV3.contract.UnpackLog(event, "DailyWithdrawalLimitsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainchainGatewayV3DepositRequestedIterator is returned from FilterDepositRequested and is used to iterate over the raw logs and unpacked data for DepositRequested events raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3DepositRequestedIterator struct {
	Event *MainchainGatewayV3DepositRequested // Event containing the contract specifics and raw log

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
func (it *MainchainGatewayV3DepositRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainGatewayV3DepositRequested)
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
		it.Event = new(MainchainGatewayV3DepositRequested)
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
func (it *MainchainGatewayV3DepositRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainGatewayV3DepositRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainGatewayV3DepositRequested represents a DepositRequested event raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3DepositRequested struct {
	ReceiptHash [32]byte
	Receipt     TransferReceipt
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDepositRequested is a free log retrieval operation binding the contract event 0xd7b25068d9dc8d00765254cfb7f5070f98d263c8d68931d937c7362fa738048b.
//
// Solidity: event DepositRequested(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) FilterDepositRequested(opts *bind.FilterOpts) (*MainchainGatewayV3DepositRequestedIterator, error) {

	logs, sub, err := _MainchainGatewayV3.contract.FilterLogs(opts, "DepositRequested")
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayV3DepositRequestedIterator{contract: _MainchainGatewayV3.contract, event: "DepositRequested", logs: logs, sub: sub}, nil
}

// WatchDepositRequested is a free log subscription operation binding the contract event 0xd7b25068d9dc8d00765254cfb7f5070f98d263c8d68931d937c7362fa738048b.
//
// Solidity: event DepositRequested(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) WatchDepositRequested(opts *bind.WatchOpts, sink chan<- *MainchainGatewayV3DepositRequested) (event.Subscription, error) {

	logs, sub, err := _MainchainGatewayV3.contract.WatchLogs(opts, "DepositRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainGatewayV3DepositRequested)
				if err := _MainchainGatewayV3.contract.UnpackLog(event, "DepositRequested", log); err != nil {
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
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) ParseDepositRequested(log types.Log) (*MainchainGatewayV3DepositRequested, error) {
	event := new(MainchainGatewayV3DepositRequested)
	if err := _MainchainGatewayV3.contract.UnpackLog(event, "DepositRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainchainGatewayV3HighTierThresholdsUpdatedIterator is returned from FilterHighTierThresholdsUpdated and is used to iterate over the raw logs and unpacked data for HighTierThresholdsUpdated events raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3HighTierThresholdsUpdatedIterator struct {
	Event *MainchainGatewayV3HighTierThresholdsUpdated // Event containing the contract specifics and raw log

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
func (it *MainchainGatewayV3HighTierThresholdsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainGatewayV3HighTierThresholdsUpdated)
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
		it.Event = new(MainchainGatewayV3HighTierThresholdsUpdated)
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
func (it *MainchainGatewayV3HighTierThresholdsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainGatewayV3HighTierThresholdsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainGatewayV3HighTierThresholdsUpdated represents a HighTierThresholdsUpdated event raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3HighTierThresholdsUpdated struct {
	Tokens     []common.Address
	Thresholds []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterHighTierThresholdsUpdated is a free log retrieval operation binding the contract event 0x80bc635c452ae67f12f9b6f12ad4daa6dbbc04eeb9ebb87d354ce10c0e210dc0.
//
// Solidity: event HighTierThresholdsUpdated(address[] tokens, uint256[] thresholds)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) FilterHighTierThresholdsUpdated(opts *bind.FilterOpts) (*MainchainGatewayV3HighTierThresholdsUpdatedIterator, error) {

	logs, sub, err := _MainchainGatewayV3.contract.FilterLogs(opts, "HighTierThresholdsUpdated")
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayV3HighTierThresholdsUpdatedIterator{contract: _MainchainGatewayV3.contract, event: "HighTierThresholdsUpdated", logs: logs, sub: sub}, nil
}

// WatchHighTierThresholdsUpdated is a free log subscription operation binding the contract event 0x80bc635c452ae67f12f9b6f12ad4daa6dbbc04eeb9ebb87d354ce10c0e210dc0.
//
// Solidity: event HighTierThresholdsUpdated(address[] tokens, uint256[] thresholds)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) WatchHighTierThresholdsUpdated(opts *bind.WatchOpts, sink chan<- *MainchainGatewayV3HighTierThresholdsUpdated) (event.Subscription, error) {

	logs, sub, err := _MainchainGatewayV3.contract.WatchLogs(opts, "HighTierThresholdsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainGatewayV3HighTierThresholdsUpdated)
				if err := _MainchainGatewayV3.contract.UnpackLog(event, "HighTierThresholdsUpdated", log); err != nil {
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

// ParseHighTierThresholdsUpdated is a log parse operation binding the contract event 0x80bc635c452ae67f12f9b6f12ad4daa6dbbc04eeb9ebb87d354ce10c0e210dc0.
//
// Solidity: event HighTierThresholdsUpdated(address[] tokens, uint256[] thresholds)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) ParseHighTierThresholdsUpdated(log types.Log) (*MainchainGatewayV3HighTierThresholdsUpdated, error) {
	event := new(MainchainGatewayV3HighTierThresholdsUpdated)
	if err := _MainchainGatewayV3.contract.UnpackLog(event, "HighTierThresholdsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainchainGatewayV3HighTierVoteWeightThresholdUpdatedIterator is returned from FilterHighTierVoteWeightThresholdUpdated and is used to iterate over the raw logs and unpacked data for HighTierVoteWeightThresholdUpdated events raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3HighTierVoteWeightThresholdUpdatedIterator struct {
	Event *MainchainGatewayV3HighTierVoteWeightThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *MainchainGatewayV3HighTierVoteWeightThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainGatewayV3HighTierVoteWeightThresholdUpdated)
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
		it.Event = new(MainchainGatewayV3HighTierVoteWeightThresholdUpdated)
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
func (it *MainchainGatewayV3HighTierVoteWeightThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainGatewayV3HighTierVoteWeightThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainGatewayV3HighTierVoteWeightThresholdUpdated represents a HighTierVoteWeightThresholdUpdated event raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3HighTierVoteWeightThresholdUpdated struct {
	Nonce               *big.Int
	Numerator           *big.Int
	Denominator         *big.Int
	PreviousNumerator   *big.Int
	PreviousDenominator *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterHighTierVoteWeightThresholdUpdated is a free log retrieval operation binding the contract event 0x31312c97b89cc751b832d98fd459b967a2c3eef3b49757d1cf5ebaa12bb6eee1.
//
// Solidity: event HighTierVoteWeightThresholdUpdated(uint256 indexed nonce, uint256 indexed numerator, uint256 indexed denominator, uint256 previousNumerator, uint256 previousDenominator)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) FilterHighTierVoteWeightThresholdUpdated(opts *bind.FilterOpts, nonce []*big.Int, numerator []*big.Int, denominator []*big.Int) (*MainchainGatewayV3HighTierVoteWeightThresholdUpdatedIterator, error) {

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

	logs, sub, err := _MainchainGatewayV3.contract.FilterLogs(opts, "HighTierVoteWeightThresholdUpdated", nonceRule, numeratorRule, denominatorRule)
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayV3HighTierVoteWeightThresholdUpdatedIterator{contract: _MainchainGatewayV3.contract, event: "HighTierVoteWeightThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchHighTierVoteWeightThresholdUpdated is a free log subscription operation binding the contract event 0x31312c97b89cc751b832d98fd459b967a2c3eef3b49757d1cf5ebaa12bb6eee1.
//
// Solidity: event HighTierVoteWeightThresholdUpdated(uint256 indexed nonce, uint256 indexed numerator, uint256 indexed denominator, uint256 previousNumerator, uint256 previousDenominator)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) WatchHighTierVoteWeightThresholdUpdated(opts *bind.WatchOpts, sink chan<- *MainchainGatewayV3HighTierVoteWeightThresholdUpdated, nonce []*big.Int, numerator []*big.Int, denominator []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _MainchainGatewayV3.contract.WatchLogs(opts, "HighTierVoteWeightThresholdUpdated", nonceRule, numeratorRule, denominatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainGatewayV3HighTierVoteWeightThresholdUpdated)
				if err := _MainchainGatewayV3.contract.UnpackLog(event, "HighTierVoteWeightThresholdUpdated", log); err != nil {
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

// ParseHighTierVoteWeightThresholdUpdated is a log parse operation binding the contract event 0x31312c97b89cc751b832d98fd459b967a2c3eef3b49757d1cf5ebaa12bb6eee1.
//
// Solidity: event HighTierVoteWeightThresholdUpdated(uint256 indexed nonce, uint256 indexed numerator, uint256 indexed denominator, uint256 previousNumerator, uint256 previousDenominator)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) ParseHighTierVoteWeightThresholdUpdated(log types.Log) (*MainchainGatewayV3HighTierVoteWeightThresholdUpdated, error) {
	event := new(MainchainGatewayV3HighTierVoteWeightThresholdUpdated)
	if err := _MainchainGatewayV3.contract.UnpackLog(event, "HighTierVoteWeightThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainchainGatewayV3InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3InitializedIterator struct {
	Event *MainchainGatewayV3Initialized // Event containing the contract specifics and raw log

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
func (it *MainchainGatewayV3InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainGatewayV3Initialized)
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
		it.Event = new(MainchainGatewayV3Initialized)
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
func (it *MainchainGatewayV3InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainGatewayV3InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainGatewayV3Initialized represents a Initialized event raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) FilterInitialized(opts *bind.FilterOpts) (*MainchainGatewayV3InitializedIterator, error) {

	logs, sub, err := _MainchainGatewayV3.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayV3InitializedIterator{contract: _MainchainGatewayV3.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MainchainGatewayV3Initialized) (event.Subscription, error) {

	logs, sub, err := _MainchainGatewayV3.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainGatewayV3Initialized)
				if err := _MainchainGatewayV3.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) ParseInitialized(log types.Log) (*MainchainGatewayV3Initialized, error) {
	event := new(MainchainGatewayV3Initialized)
	if err := _MainchainGatewayV3.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainchainGatewayV3LockedThresholdsUpdatedIterator is returned from FilterLockedThresholdsUpdated and is used to iterate over the raw logs and unpacked data for LockedThresholdsUpdated events raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3LockedThresholdsUpdatedIterator struct {
	Event *MainchainGatewayV3LockedThresholdsUpdated // Event containing the contract specifics and raw log

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
func (it *MainchainGatewayV3LockedThresholdsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainGatewayV3LockedThresholdsUpdated)
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
		it.Event = new(MainchainGatewayV3LockedThresholdsUpdated)
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
func (it *MainchainGatewayV3LockedThresholdsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainGatewayV3LockedThresholdsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainGatewayV3LockedThresholdsUpdated represents a LockedThresholdsUpdated event raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3LockedThresholdsUpdated struct {
	Tokens     []common.Address
	Thresholds []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLockedThresholdsUpdated is a free log retrieval operation binding the contract event 0x64557254143204d91ba2d95acb9fda1e5fea55f77efd028685765bc1e94dd4b5.
//
// Solidity: event LockedThresholdsUpdated(address[] tokens, uint256[] thresholds)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) FilterLockedThresholdsUpdated(opts *bind.FilterOpts) (*MainchainGatewayV3LockedThresholdsUpdatedIterator, error) {

	logs, sub, err := _MainchainGatewayV3.contract.FilterLogs(opts, "LockedThresholdsUpdated")
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayV3LockedThresholdsUpdatedIterator{contract: _MainchainGatewayV3.contract, event: "LockedThresholdsUpdated", logs: logs, sub: sub}, nil
}

// WatchLockedThresholdsUpdated is a free log subscription operation binding the contract event 0x64557254143204d91ba2d95acb9fda1e5fea55f77efd028685765bc1e94dd4b5.
//
// Solidity: event LockedThresholdsUpdated(address[] tokens, uint256[] thresholds)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) WatchLockedThresholdsUpdated(opts *bind.WatchOpts, sink chan<- *MainchainGatewayV3LockedThresholdsUpdated) (event.Subscription, error) {

	logs, sub, err := _MainchainGatewayV3.contract.WatchLogs(opts, "LockedThresholdsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainGatewayV3LockedThresholdsUpdated)
				if err := _MainchainGatewayV3.contract.UnpackLog(event, "LockedThresholdsUpdated", log); err != nil {
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

// ParseLockedThresholdsUpdated is a log parse operation binding the contract event 0x64557254143204d91ba2d95acb9fda1e5fea55f77efd028685765bc1e94dd4b5.
//
// Solidity: event LockedThresholdsUpdated(address[] tokens, uint256[] thresholds)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) ParseLockedThresholdsUpdated(log types.Log) (*MainchainGatewayV3LockedThresholdsUpdated, error) {
	event := new(MainchainGatewayV3LockedThresholdsUpdated)
	if err := _MainchainGatewayV3.contract.UnpackLog(event, "LockedThresholdsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainchainGatewayV3PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3PausedIterator struct {
	Event *MainchainGatewayV3Paused // Event containing the contract specifics and raw log

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
func (it *MainchainGatewayV3PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainGatewayV3Paused)
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
		it.Event = new(MainchainGatewayV3Paused)
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
func (it *MainchainGatewayV3PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainGatewayV3PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainGatewayV3Paused represents a Paused event raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) FilterPaused(opts *bind.FilterOpts) (*MainchainGatewayV3PausedIterator, error) {

	logs, sub, err := _MainchainGatewayV3.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayV3PausedIterator{contract: _MainchainGatewayV3.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MainchainGatewayV3Paused) (event.Subscription, error) {

	logs, sub, err := _MainchainGatewayV3.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainGatewayV3Paused)
				if err := _MainchainGatewayV3.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) ParsePaused(log types.Log) (*MainchainGatewayV3Paused, error) {
	event := new(MainchainGatewayV3Paused)
	if err := _MainchainGatewayV3.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainchainGatewayV3RestrictedIterator is returned from FilterRestricted and is used to iterate over the raw logs and unpacked data for Restricted events raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3RestrictedIterator struct {
	Event *MainchainGatewayV3Restricted // Event containing the contract specifics and raw log

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
func (it *MainchainGatewayV3RestrictedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainGatewayV3Restricted)
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
		it.Event = new(MainchainGatewayV3Restricted)
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
func (it *MainchainGatewayV3RestrictedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainGatewayV3RestrictedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainGatewayV3Restricted represents a Restricted event raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3Restricted struct {
	By         common.Address
	FnSig      [4]byte
	EnumBitmap uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRestricted is a free log retrieval operation binding the contract event 0x9562876a66c0b06c9a42651c249575b275ba92206f3daffa644c8c6a1856c8eb.
//
// Solidity: event Restricted(address indexed by, bytes4 indexed fnSig, uint8 enumBitmap)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) FilterRestricted(opts *bind.FilterOpts, by []common.Address, fnSig [][4]byte) (*MainchainGatewayV3RestrictedIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var fnSigRule []interface{}
	for _, fnSigItem := range fnSig {
		fnSigRule = append(fnSigRule, fnSigItem)
	}

	logs, sub, err := _MainchainGatewayV3.contract.FilterLogs(opts, "Restricted", byRule, fnSigRule)
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayV3RestrictedIterator{contract: _MainchainGatewayV3.contract, event: "Restricted", logs: logs, sub: sub}, nil
}

// WatchRestricted is a free log subscription operation binding the contract event 0x9562876a66c0b06c9a42651c249575b275ba92206f3daffa644c8c6a1856c8eb.
//
// Solidity: event Restricted(address indexed by, bytes4 indexed fnSig, uint8 enumBitmap)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) WatchRestricted(opts *bind.WatchOpts, sink chan<- *MainchainGatewayV3Restricted, by []common.Address, fnSig [][4]byte) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var fnSigRule []interface{}
	for _, fnSigItem := range fnSig {
		fnSigRule = append(fnSigRule, fnSigItem)
	}

	logs, sub, err := _MainchainGatewayV3.contract.WatchLogs(opts, "Restricted", byRule, fnSigRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainGatewayV3Restricted)
				if err := _MainchainGatewayV3.contract.UnpackLog(event, "Restricted", log); err != nil {
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

// ParseRestricted is a log parse operation binding the contract event 0x9562876a66c0b06c9a42651c249575b275ba92206f3daffa644c8c6a1856c8eb.
//
// Solidity: event Restricted(address indexed by, bytes4 indexed fnSig, uint8 enumBitmap)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) ParseRestricted(log types.Log) (*MainchainGatewayV3Restricted, error) {
	event := new(MainchainGatewayV3Restricted)
	if err := _MainchainGatewayV3.contract.UnpackLog(event, "Restricted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainchainGatewayV3RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3RoleAdminChangedIterator struct {
	Event *MainchainGatewayV3RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *MainchainGatewayV3RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainGatewayV3RoleAdminChanged)
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
		it.Event = new(MainchainGatewayV3RoleAdminChanged)
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
func (it *MainchainGatewayV3RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainGatewayV3RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainGatewayV3RoleAdminChanged represents a RoleAdminChanged event raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*MainchainGatewayV3RoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _MainchainGatewayV3.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayV3RoleAdminChangedIterator{contract: _MainchainGatewayV3.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *MainchainGatewayV3RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _MainchainGatewayV3.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainGatewayV3RoleAdminChanged)
				if err := _MainchainGatewayV3.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) ParseRoleAdminChanged(log types.Log) (*MainchainGatewayV3RoleAdminChanged, error) {
	event := new(MainchainGatewayV3RoleAdminChanged)
	if err := _MainchainGatewayV3.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainchainGatewayV3RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3RoleGrantedIterator struct {
	Event *MainchainGatewayV3RoleGranted // Event containing the contract specifics and raw log

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
func (it *MainchainGatewayV3RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainGatewayV3RoleGranted)
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
		it.Event = new(MainchainGatewayV3RoleGranted)
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
func (it *MainchainGatewayV3RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainGatewayV3RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainGatewayV3RoleGranted represents a RoleGranted event raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MainchainGatewayV3RoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MainchainGatewayV3.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayV3RoleGrantedIterator{contract: _MainchainGatewayV3.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *MainchainGatewayV3RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MainchainGatewayV3.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainGatewayV3RoleGranted)
				if err := _MainchainGatewayV3.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) ParseRoleGranted(log types.Log) (*MainchainGatewayV3RoleGranted, error) {
	event := new(MainchainGatewayV3RoleGranted)
	if err := _MainchainGatewayV3.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainchainGatewayV3RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3RoleRevokedIterator struct {
	Event *MainchainGatewayV3RoleRevoked // Event containing the contract specifics and raw log

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
func (it *MainchainGatewayV3RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainGatewayV3RoleRevoked)
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
		it.Event = new(MainchainGatewayV3RoleRevoked)
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
func (it *MainchainGatewayV3RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainGatewayV3RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainGatewayV3RoleRevoked represents a RoleRevoked event raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MainchainGatewayV3RoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MainchainGatewayV3.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayV3RoleRevokedIterator{contract: _MainchainGatewayV3.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *MainchainGatewayV3RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MainchainGatewayV3.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainGatewayV3RoleRevoked)
				if err := _MainchainGatewayV3.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) ParseRoleRevoked(log types.Log) (*MainchainGatewayV3RoleRevoked, error) {
	event := new(MainchainGatewayV3RoleRevoked)
	if err := _MainchainGatewayV3.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainchainGatewayV3ThresholdUpdatedIterator is returned from FilterThresholdUpdated and is used to iterate over the raw logs and unpacked data for ThresholdUpdated events raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3ThresholdUpdatedIterator struct {
	Event *MainchainGatewayV3ThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *MainchainGatewayV3ThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainGatewayV3ThresholdUpdated)
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
		it.Event = new(MainchainGatewayV3ThresholdUpdated)
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
func (it *MainchainGatewayV3ThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainGatewayV3ThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainGatewayV3ThresholdUpdated represents a ThresholdUpdated event raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3ThresholdUpdated struct {
	Nonce               *big.Int
	Numerator           *big.Int
	Denominator         *big.Int
	PreviousNumerator   *big.Int
	PreviousDenominator *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterThresholdUpdated is a free log retrieval operation binding the contract event 0x976f8a9c5bdf8248dec172376d6e2b80a8e3df2f0328e381c6db8e1cf138c0f8.
//
// Solidity: event ThresholdUpdated(uint256 indexed nonce, uint256 indexed numerator, uint256 indexed denominator, uint256 previousNumerator, uint256 previousDenominator)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) FilterThresholdUpdated(opts *bind.FilterOpts, nonce []*big.Int, numerator []*big.Int, denominator []*big.Int) (*MainchainGatewayV3ThresholdUpdatedIterator, error) {

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

	logs, sub, err := _MainchainGatewayV3.contract.FilterLogs(opts, "ThresholdUpdated", nonceRule, numeratorRule, denominatorRule)
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayV3ThresholdUpdatedIterator{contract: _MainchainGatewayV3.contract, event: "ThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchThresholdUpdated is a free log subscription operation binding the contract event 0x976f8a9c5bdf8248dec172376d6e2b80a8e3df2f0328e381c6db8e1cf138c0f8.
//
// Solidity: event ThresholdUpdated(uint256 indexed nonce, uint256 indexed numerator, uint256 indexed denominator, uint256 previousNumerator, uint256 previousDenominator)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) WatchThresholdUpdated(opts *bind.WatchOpts, sink chan<- *MainchainGatewayV3ThresholdUpdated, nonce []*big.Int, numerator []*big.Int, denominator []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _MainchainGatewayV3.contract.WatchLogs(opts, "ThresholdUpdated", nonceRule, numeratorRule, denominatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainGatewayV3ThresholdUpdated)
				if err := _MainchainGatewayV3.contract.UnpackLog(event, "ThresholdUpdated", log); err != nil {
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

// ParseThresholdUpdated is a log parse operation binding the contract event 0x976f8a9c5bdf8248dec172376d6e2b80a8e3df2f0328e381c6db8e1cf138c0f8.
//
// Solidity: event ThresholdUpdated(uint256 indexed nonce, uint256 indexed numerator, uint256 indexed denominator, uint256 previousNumerator, uint256 previousDenominator)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) ParseThresholdUpdated(log types.Log) (*MainchainGatewayV3ThresholdUpdated, error) {
	event := new(MainchainGatewayV3ThresholdUpdated)
	if err := _MainchainGatewayV3.contract.UnpackLog(event, "ThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainchainGatewayV3TokenMappedIterator is returned from FilterTokenMapped and is used to iterate over the raw logs and unpacked data for TokenMapped events raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3TokenMappedIterator struct {
	Event *MainchainGatewayV3TokenMapped // Event containing the contract specifics and raw log

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
func (it *MainchainGatewayV3TokenMappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainGatewayV3TokenMapped)
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
		it.Event = new(MainchainGatewayV3TokenMapped)
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
func (it *MainchainGatewayV3TokenMappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainGatewayV3TokenMappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainGatewayV3TokenMapped represents a TokenMapped event raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3TokenMapped struct {
	MainchainTokens []common.Address
	RoninTokens     []common.Address
	Standards       []uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTokenMapped is a free log retrieval operation binding the contract event 0xa4f03cc9c0e0aeb5b71b4ec800702753f65748c2cf3064695ba8e8b46be70444.
//
// Solidity: event TokenMapped(address[] mainchainTokens, address[] roninTokens, uint8[] standards)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) FilterTokenMapped(opts *bind.FilterOpts) (*MainchainGatewayV3TokenMappedIterator, error) {

	logs, sub, err := _MainchainGatewayV3.contract.FilterLogs(opts, "TokenMapped")
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayV3TokenMappedIterator{contract: _MainchainGatewayV3.contract, event: "TokenMapped", logs: logs, sub: sub}, nil
}

// WatchTokenMapped is a free log subscription operation binding the contract event 0xa4f03cc9c0e0aeb5b71b4ec800702753f65748c2cf3064695ba8e8b46be70444.
//
// Solidity: event TokenMapped(address[] mainchainTokens, address[] roninTokens, uint8[] standards)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) WatchTokenMapped(opts *bind.WatchOpts, sink chan<- *MainchainGatewayV3TokenMapped) (event.Subscription, error) {

	logs, sub, err := _MainchainGatewayV3.contract.WatchLogs(opts, "TokenMapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainGatewayV3TokenMapped)
				if err := _MainchainGatewayV3.contract.UnpackLog(event, "TokenMapped", log); err != nil {
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
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) ParseTokenMapped(log types.Log) (*MainchainGatewayV3TokenMapped, error) {
	event := new(MainchainGatewayV3TokenMapped)
	if err := _MainchainGatewayV3.contract.UnpackLog(event, "TokenMapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainchainGatewayV3UnRestrictedIterator is returned from FilterUnRestricted and is used to iterate over the raw logs and unpacked data for UnRestricted events raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3UnRestrictedIterator struct {
	Event *MainchainGatewayV3UnRestricted // Event containing the contract specifics and raw log

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
func (it *MainchainGatewayV3UnRestrictedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainGatewayV3UnRestricted)
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
		it.Event = new(MainchainGatewayV3UnRestricted)
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
func (it *MainchainGatewayV3UnRestrictedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainGatewayV3UnRestrictedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainGatewayV3UnRestricted represents a UnRestricted event raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3UnRestricted struct {
	By    common.Address
	FnSig [4]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnRestricted is a free log retrieval operation binding the contract event 0x4f91021730250ae6dd01e5ee7051cfde47c138483157131349b5f07a7c7c29c3.
//
// Solidity: event UnRestricted(address indexed by, bytes4 indexed fnSig)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) FilterUnRestricted(opts *bind.FilterOpts, by []common.Address, fnSig [][4]byte) (*MainchainGatewayV3UnRestrictedIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var fnSigRule []interface{}
	for _, fnSigItem := range fnSig {
		fnSigRule = append(fnSigRule, fnSigItem)
	}

	logs, sub, err := _MainchainGatewayV3.contract.FilterLogs(opts, "UnRestricted", byRule, fnSigRule)
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayV3UnRestrictedIterator{contract: _MainchainGatewayV3.contract, event: "UnRestricted", logs: logs, sub: sub}, nil
}

// WatchUnRestricted is a free log subscription operation binding the contract event 0x4f91021730250ae6dd01e5ee7051cfde47c138483157131349b5f07a7c7c29c3.
//
// Solidity: event UnRestricted(address indexed by, bytes4 indexed fnSig)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) WatchUnRestricted(opts *bind.WatchOpts, sink chan<- *MainchainGatewayV3UnRestricted, by []common.Address, fnSig [][4]byte) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var fnSigRule []interface{}
	for _, fnSigItem := range fnSig {
		fnSigRule = append(fnSigRule, fnSigItem)
	}

	logs, sub, err := _MainchainGatewayV3.contract.WatchLogs(opts, "UnRestricted", byRule, fnSigRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainGatewayV3UnRestricted)
				if err := _MainchainGatewayV3.contract.UnpackLog(event, "UnRestricted", log); err != nil {
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

// ParseUnRestricted is a log parse operation binding the contract event 0x4f91021730250ae6dd01e5ee7051cfde47c138483157131349b5f07a7c7c29c3.
//
// Solidity: event UnRestricted(address indexed by, bytes4 indexed fnSig)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) ParseUnRestricted(log types.Log) (*MainchainGatewayV3UnRestricted, error) {
	event := new(MainchainGatewayV3UnRestricted)
	if err := _MainchainGatewayV3.contract.UnpackLog(event, "UnRestricted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainchainGatewayV3UnlockFeePercentagesUpdatedIterator is returned from FilterUnlockFeePercentagesUpdated and is used to iterate over the raw logs and unpacked data for UnlockFeePercentagesUpdated events raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3UnlockFeePercentagesUpdatedIterator struct {
	Event *MainchainGatewayV3UnlockFeePercentagesUpdated // Event containing the contract specifics and raw log

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
func (it *MainchainGatewayV3UnlockFeePercentagesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainGatewayV3UnlockFeePercentagesUpdated)
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
		it.Event = new(MainchainGatewayV3UnlockFeePercentagesUpdated)
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
func (it *MainchainGatewayV3UnlockFeePercentagesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainGatewayV3UnlockFeePercentagesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainGatewayV3UnlockFeePercentagesUpdated represents a UnlockFeePercentagesUpdated event raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3UnlockFeePercentagesUpdated struct {
	Tokens      []common.Address
	Percentages []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnlockFeePercentagesUpdated is a free log retrieval operation binding the contract event 0xb05f5de88ae0294ebb6f67c5af2fcbbd593cc6bdfe543e2869794a4c8ce3ea50.
//
// Solidity: event UnlockFeePercentagesUpdated(address[] tokens, uint256[] percentages)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) FilterUnlockFeePercentagesUpdated(opts *bind.FilterOpts) (*MainchainGatewayV3UnlockFeePercentagesUpdatedIterator, error) {

	logs, sub, err := _MainchainGatewayV3.contract.FilterLogs(opts, "UnlockFeePercentagesUpdated")
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayV3UnlockFeePercentagesUpdatedIterator{contract: _MainchainGatewayV3.contract, event: "UnlockFeePercentagesUpdated", logs: logs, sub: sub}, nil
}

// WatchUnlockFeePercentagesUpdated is a free log subscription operation binding the contract event 0xb05f5de88ae0294ebb6f67c5af2fcbbd593cc6bdfe543e2869794a4c8ce3ea50.
//
// Solidity: event UnlockFeePercentagesUpdated(address[] tokens, uint256[] percentages)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) WatchUnlockFeePercentagesUpdated(opts *bind.WatchOpts, sink chan<- *MainchainGatewayV3UnlockFeePercentagesUpdated) (event.Subscription, error) {

	logs, sub, err := _MainchainGatewayV3.contract.WatchLogs(opts, "UnlockFeePercentagesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainGatewayV3UnlockFeePercentagesUpdated)
				if err := _MainchainGatewayV3.contract.UnpackLog(event, "UnlockFeePercentagesUpdated", log); err != nil {
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

// ParseUnlockFeePercentagesUpdated is a log parse operation binding the contract event 0xb05f5de88ae0294ebb6f67c5af2fcbbd593cc6bdfe543e2869794a4c8ce3ea50.
//
// Solidity: event UnlockFeePercentagesUpdated(address[] tokens, uint256[] percentages)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) ParseUnlockFeePercentagesUpdated(log types.Log) (*MainchainGatewayV3UnlockFeePercentagesUpdated, error) {
	event := new(MainchainGatewayV3UnlockFeePercentagesUpdated)
	if err := _MainchainGatewayV3.contract.UnpackLog(event, "UnlockFeePercentagesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainchainGatewayV3UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3UnpausedIterator struct {
	Event *MainchainGatewayV3Unpaused // Event containing the contract specifics and raw log

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
func (it *MainchainGatewayV3UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainGatewayV3Unpaused)
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
		it.Event = new(MainchainGatewayV3Unpaused)
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
func (it *MainchainGatewayV3UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainGatewayV3UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainGatewayV3Unpaused represents a Unpaused event raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) FilterUnpaused(opts *bind.FilterOpts) (*MainchainGatewayV3UnpausedIterator, error) {

	logs, sub, err := _MainchainGatewayV3.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayV3UnpausedIterator{contract: _MainchainGatewayV3.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MainchainGatewayV3Unpaused) (event.Subscription, error) {

	logs, sub, err := _MainchainGatewayV3.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainGatewayV3Unpaused)
				if err := _MainchainGatewayV3.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) ParseUnpaused(log types.Log) (*MainchainGatewayV3Unpaused, error) {
	event := new(MainchainGatewayV3Unpaused)
	if err := _MainchainGatewayV3.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainchainGatewayV3WhitelistUpdatedIterator is returned from FilterWhitelistUpdated and is used to iterate over the raw logs and unpacked data for WhitelistUpdated events raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3WhitelistUpdatedIterator struct {
	Event *MainchainGatewayV3WhitelistUpdated // Event containing the contract specifics and raw log

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
func (it *MainchainGatewayV3WhitelistUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainGatewayV3WhitelistUpdated)
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
		it.Event = new(MainchainGatewayV3WhitelistUpdated)
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
func (it *MainchainGatewayV3WhitelistUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainGatewayV3WhitelistUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainGatewayV3WhitelistUpdated represents a WhitelistUpdated event raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3WhitelistUpdated struct {
	By                   common.Address
	Tokens               []common.Address
	Recipients           []common.Address
	RemoteChainSelectors []uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterWhitelistUpdated is a free log retrieval operation binding the contract event 0x58881457e968e680588d4a27c22619b7f8f0d3f394c679907123547a2ae200ca.
//
// Solidity: event WhitelistUpdated(address indexed by, address[] tokens, address[] recipients, uint64[] remoteChainSelectors)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) FilterWhitelistUpdated(opts *bind.FilterOpts, by []common.Address) (*MainchainGatewayV3WhitelistUpdatedIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _MainchainGatewayV3.contract.FilterLogs(opts, "WhitelistUpdated", byRule)
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayV3WhitelistUpdatedIterator{contract: _MainchainGatewayV3.contract, event: "WhitelistUpdated", logs: logs, sub: sub}, nil
}

// WatchWhitelistUpdated is a free log subscription operation binding the contract event 0x58881457e968e680588d4a27c22619b7f8f0d3f394c679907123547a2ae200ca.
//
// Solidity: event WhitelistUpdated(address indexed by, address[] tokens, address[] recipients, uint64[] remoteChainSelectors)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) WatchWhitelistUpdated(opts *bind.WatchOpts, sink chan<- *MainchainGatewayV3WhitelistUpdated, by []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _MainchainGatewayV3.contract.WatchLogs(opts, "WhitelistUpdated", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainGatewayV3WhitelistUpdated)
				if err := _MainchainGatewayV3.contract.UnpackLog(event, "WhitelistUpdated", log); err != nil {
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

// ParseWhitelistUpdated is a log parse operation binding the contract event 0x58881457e968e680588d4a27c22619b7f8f0d3f394c679907123547a2ae200ca.
//
// Solidity: event WhitelistUpdated(address indexed by, address[] tokens, address[] recipients, uint64[] remoteChainSelectors)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) ParseWhitelistUpdated(log types.Log) (*MainchainGatewayV3WhitelistUpdated, error) {
	event := new(MainchainGatewayV3WhitelistUpdated)
	if err := _MainchainGatewayV3.contract.UnpackLog(event, "WhitelistUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainchainGatewayV3WithdrawalLockedIterator is returned from FilterWithdrawalLocked and is used to iterate over the raw logs and unpacked data for WithdrawalLocked events raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3WithdrawalLockedIterator struct {
	Event *MainchainGatewayV3WithdrawalLocked // Event containing the contract specifics and raw log

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
func (it *MainchainGatewayV3WithdrawalLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainGatewayV3WithdrawalLocked)
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
		it.Event = new(MainchainGatewayV3WithdrawalLocked)
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
func (it *MainchainGatewayV3WithdrawalLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainGatewayV3WithdrawalLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainGatewayV3WithdrawalLocked represents a WithdrawalLocked event raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3WithdrawalLocked struct {
	ReceiptHash [32]byte
	Receipt     TransferReceipt
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalLocked is a free log retrieval operation binding the contract event 0x89e52969465b1f1866fc5d46fd62de953962e9cb33552443cd999eba05bd20dc.
//
// Solidity: event WithdrawalLocked(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) FilterWithdrawalLocked(opts *bind.FilterOpts) (*MainchainGatewayV3WithdrawalLockedIterator, error) {

	logs, sub, err := _MainchainGatewayV3.contract.FilterLogs(opts, "WithdrawalLocked")
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayV3WithdrawalLockedIterator{contract: _MainchainGatewayV3.contract, event: "WithdrawalLocked", logs: logs, sub: sub}, nil
}

// WatchWithdrawalLocked is a free log subscription operation binding the contract event 0x89e52969465b1f1866fc5d46fd62de953962e9cb33552443cd999eba05bd20dc.
//
// Solidity: event WithdrawalLocked(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) WatchWithdrawalLocked(opts *bind.WatchOpts, sink chan<- *MainchainGatewayV3WithdrawalLocked) (event.Subscription, error) {

	logs, sub, err := _MainchainGatewayV3.contract.WatchLogs(opts, "WithdrawalLocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainGatewayV3WithdrawalLocked)
				if err := _MainchainGatewayV3.contract.UnpackLog(event, "WithdrawalLocked", log); err != nil {
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
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) ParseWithdrawalLocked(log types.Log) (*MainchainGatewayV3WithdrawalLocked, error) {
	event := new(MainchainGatewayV3WithdrawalLocked)
	if err := _MainchainGatewayV3.contract.UnpackLog(event, "WithdrawalLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainchainGatewayV3WithdrawalUnlockedIterator is returned from FilterWithdrawalUnlocked and is used to iterate over the raw logs and unpacked data for WithdrawalUnlocked events raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3WithdrawalUnlockedIterator struct {
	Event *MainchainGatewayV3WithdrawalUnlocked // Event containing the contract specifics and raw log

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
func (it *MainchainGatewayV3WithdrawalUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainGatewayV3WithdrawalUnlocked)
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
		it.Event = new(MainchainGatewayV3WithdrawalUnlocked)
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
func (it *MainchainGatewayV3WithdrawalUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainGatewayV3WithdrawalUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainGatewayV3WithdrawalUnlocked represents a WithdrawalUnlocked event raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3WithdrawalUnlocked struct {
	ReceiptHash [32]byte
	Receipt     TransferReceipt
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalUnlocked is a free log retrieval operation binding the contract event 0xd639511b37b3b002cca6cfe6bca0d833945a5af5a045578a0627fc43b79b2630.
//
// Solidity: event WithdrawalUnlocked(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) FilterWithdrawalUnlocked(opts *bind.FilterOpts) (*MainchainGatewayV3WithdrawalUnlockedIterator, error) {

	logs, sub, err := _MainchainGatewayV3.contract.FilterLogs(opts, "WithdrawalUnlocked")
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayV3WithdrawalUnlockedIterator{contract: _MainchainGatewayV3.contract, event: "WithdrawalUnlocked", logs: logs, sub: sub}, nil
}

// WatchWithdrawalUnlocked is a free log subscription operation binding the contract event 0xd639511b37b3b002cca6cfe6bca0d833945a5af5a045578a0627fc43b79b2630.
//
// Solidity: event WithdrawalUnlocked(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) WatchWithdrawalUnlocked(opts *bind.WatchOpts, sink chan<- *MainchainGatewayV3WithdrawalUnlocked) (event.Subscription, error) {

	logs, sub, err := _MainchainGatewayV3.contract.WatchLogs(opts, "WithdrawalUnlocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainGatewayV3WithdrawalUnlocked)
				if err := _MainchainGatewayV3.contract.UnpackLog(event, "WithdrawalUnlocked", log); err != nil {
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
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) ParseWithdrawalUnlocked(log types.Log) (*MainchainGatewayV3WithdrawalUnlocked, error) {
	event := new(MainchainGatewayV3WithdrawalUnlocked)
	if err := _MainchainGatewayV3.contract.UnpackLog(event, "WithdrawalUnlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainchainGatewayV3WithdrewIterator is returned from FilterWithdrew and is used to iterate over the raw logs and unpacked data for Withdrew events raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3WithdrewIterator struct {
	Event *MainchainGatewayV3Withdrew // Event containing the contract specifics and raw log

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
func (it *MainchainGatewayV3WithdrewIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainGatewayV3Withdrew)
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
		it.Event = new(MainchainGatewayV3Withdrew)
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
func (it *MainchainGatewayV3WithdrewIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainGatewayV3WithdrewIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainGatewayV3Withdrew represents a Withdrew event raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3Withdrew struct {
	ReceiptHash [32]byte
	Receipt     TransferReceipt
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrew is a free log retrieval operation binding the contract event 0x21e88e956aa3e086f6388e899965cef814688f99ad8bb29b08d396571016372d.
//
// Solidity: event Withdrew(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) FilterWithdrew(opts *bind.FilterOpts) (*MainchainGatewayV3WithdrewIterator, error) {

	logs, sub, err := _MainchainGatewayV3.contract.FilterLogs(opts, "Withdrew")
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayV3WithdrewIterator{contract: _MainchainGatewayV3.contract, event: "Withdrew", logs: logs, sub: sub}, nil
}

// WatchWithdrew is a free log subscription operation binding the contract event 0x21e88e956aa3e086f6388e899965cef814688f99ad8bb29b08d396571016372d.
//
// Solidity: event Withdrew(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) WatchWithdrew(opts *bind.WatchOpts, sink chan<- *MainchainGatewayV3Withdrew) (event.Subscription, error) {

	logs, sub, err := _MainchainGatewayV3.contract.WatchLogs(opts, "Withdrew")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainGatewayV3Withdrew)
				if err := _MainchainGatewayV3.contract.UnpackLog(event, "Withdrew", log); err != nil {
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
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) ParseWithdrew(log types.Log) (*MainchainGatewayV3Withdrew, error) {
	event := new(MainchainGatewayV3Withdrew)
	if err := _MainchainGatewayV3.contract.UnpackLog(event, "Withdrew", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainchainGatewayV3WrappedNativeTokenContractUpdatedIterator is returned from FilterWrappedNativeTokenContractUpdated and is used to iterate over the raw logs and unpacked data for WrappedNativeTokenContractUpdated events raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3WrappedNativeTokenContractUpdatedIterator struct {
	Event *MainchainGatewayV3WrappedNativeTokenContractUpdated // Event containing the contract specifics and raw log

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
func (it *MainchainGatewayV3WrappedNativeTokenContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainGatewayV3WrappedNativeTokenContractUpdated)
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
		it.Event = new(MainchainGatewayV3WrappedNativeTokenContractUpdated)
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
func (it *MainchainGatewayV3WrappedNativeTokenContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainGatewayV3WrappedNativeTokenContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainGatewayV3WrappedNativeTokenContractUpdated represents a WrappedNativeTokenContractUpdated event raised by the MainchainGatewayV3 contract.
type MainchainGatewayV3WrappedNativeTokenContractUpdated struct {
	Weth common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWrappedNativeTokenContractUpdated is a free log retrieval operation binding the contract event 0x9d2334c23be647e994f27a72c5eee42a43d5bdcfe15bb88e939103c2b114cbaf.
//
// Solidity: event WrappedNativeTokenContractUpdated(address weth)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) FilterWrappedNativeTokenContractUpdated(opts *bind.FilterOpts) (*MainchainGatewayV3WrappedNativeTokenContractUpdatedIterator, error) {

	logs, sub, err := _MainchainGatewayV3.contract.FilterLogs(opts, "WrappedNativeTokenContractUpdated")
	if err != nil {
		return nil, err
	}
	return &MainchainGatewayV3WrappedNativeTokenContractUpdatedIterator{contract: _MainchainGatewayV3.contract, event: "WrappedNativeTokenContractUpdated", logs: logs, sub: sub}, nil
}

// WatchWrappedNativeTokenContractUpdated is a free log subscription operation binding the contract event 0x9d2334c23be647e994f27a72c5eee42a43d5bdcfe15bb88e939103c2b114cbaf.
//
// Solidity: event WrappedNativeTokenContractUpdated(address weth)
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) WatchWrappedNativeTokenContractUpdated(opts *bind.WatchOpts, sink chan<- *MainchainGatewayV3WrappedNativeTokenContractUpdated) (event.Subscription, error) {

	logs, sub, err := _MainchainGatewayV3.contract.WatchLogs(opts, "WrappedNativeTokenContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainGatewayV3WrappedNativeTokenContractUpdated)
				if err := _MainchainGatewayV3.contract.UnpackLog(event, "WrappedNativeTokenContractUpdated", log); err != nil {
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
func (_MainchainGatewayV3 *MainchainGatewayV3Filterer) ParseWrappedNativeTokenContractUpdated(log types.Log) (*MainchainGatewayV3WrappedNativeTokenContractUpdated, error) {
	event := new(MainchainGatewayV3WrappedNativeTokenContractUpdated)
	if err := _MainchainGatewayV3.contract.UnpackLog(event, "WrappedNativeTokenContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
