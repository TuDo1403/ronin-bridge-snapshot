// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ronin_gateway_v3

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

// RoninGatewayV3MetaData contains all meta data concerning the RoninGatewayV3 contract.
var RoninGatewayV3MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"receiptHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTransfer.Kind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"mainchain\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"ronin\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structTransfer.Receipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"receiptHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTransfer.Kind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"mainchain\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"ronin\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structTransfer.Receipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"MainchainWithdrew\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"threshold\",\"type\":\"uint256[]\"}],\"name\":\"MinimumThresholdsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"numerator\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"denominator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousNumerator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousDenominator\",\"type\":\"uint256\"}],\"name\":\"ThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"roninTokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"mainchainTokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"chainIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"enumToken.Standard[]\",\"name\":\"standards\",\"type\":\"uint8[]\"}],\"name\":\"TokenMapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIWeightedValidator\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ValidatorContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"receiptHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTransfer.Kind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"mainchain\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"ronin\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structTransfer.Receipt\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"WithdrawalRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"receiptHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTransfer.Kind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"mainchain\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"ronin\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structTransfer.Receipt\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"WithdrawalSignaturesRequested\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWAL_MIGRATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipientAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"internalType\":\"structTransfer.Request[]\",\"name\":\"_requests\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"bulkRequestWithdrawalFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_withdrawals\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_signatures\",\"type\":\"bytes[]\"}],\"name\":\"bulkSubmitWithdrawalSignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_voteWeight\",\"type\":\"uint256\"}],\"name\":\"checkThreshold\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTransfer.Kind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"mainchain\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"ronin\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"internalType\":\"structTransfer.Receipt\",\"name\":\"_receipt\",\"type\":\"tuple\"}],\"name\":\"depositFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"depositVote\",\"outputs\":[{\"internalType\":\"enumGatewayGovernance.VoteStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"finalHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_depositId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"depositVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_roninToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"getMainchainToken\",\"outputs\":[{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"}],\"internalType\":\"structMappedTokenConsumer.MappedToken\",\"name\":\"_token\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawalId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"}],\"name\":\"getWithdrawalSignatures\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_signatures\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_roleSetter\",\"type\":\"address\"},{\"internalType\":\"contractIWeightedValidator\",\"name\":\"_validatorContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_numerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_denominator\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_withdrawalMigrators\",\"type\":\"address[]\"},{\"internalType\":\"address[][2]\",\"name\":\"_packedAddresses\",\"type\":\"address[][2]\"},{\"internalType\":\"uint256[][2]\",\"name\":\"_packedNumbers\",\"type\":\"uint256[][2]\"},{\"internalType\":\"enumToken.Standard[]\",\"name\":\"_standards\",\"type\":\"uint8[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawalId\",\"type\":\"uint256\"}],\"name\":\"mainchainWithdrew\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mainchainWithdrewVote\",\"outputs\":[{\"internalType\":\"enumGatewayGovernance.VoteStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"finalHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_roninTokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_mainchainTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_chainIds\",\"type\":\"uint256[]\"},{\"internalType\":\"enumToken.Standard[]\",\"name\":\"_standards\",\"type\":\"uint8[]\"}],\"name\":\"mapTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"markWithdrawalMigrated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipientAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"internalType\":\"structTransfer.Request[]\",\"name\":\"_requests\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"_requesters\",\"type\":\"address[]\"}],\"name\":\"migrateWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minimumThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumVoteWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipientAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"internalType\":\"structTransfer.Request\",\"name\":\"_request\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"requestWithdrawalFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawalId\",\"type\":\"uint256\"}],\"name\":\"requestWithdrawalSignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_thresholds\",\"type\":\"uint256[]\"}],\"name\":\"setMinimumThresholds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_numerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_denominator\",\"type\":\"uint256\"}],\"name\":\"setThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIWeightedValidator\",\"name\":\"_validatorContract\",\"type\":\"address\"}],\"name\":\"setValidatorContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_withdrawalIds\",\"type\":\"uint256[]\"}],\"name\":\"tryBulkAcknowledgeMainchainWithdrew\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"_executedReceipts\",\"type\":\"bool[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTransfer.Kind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"mainchain\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"ronin\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"internalType\":\"structTransfer.Receipt[]\",\"name\":\"_receipts\",\"type\":\"tuple[]\"}],\"name\":\"tryBulkDepositFor\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"_executedReceipts\",\"type\":\"bool[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorContract\",\"outputs\":[{\"internalType\":\"contractIWeightedValidator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTransfer.Kind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"mainchain\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"ronin\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalMigrated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// RoninGatewayV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use RoninGatewayV3MetaData.ABI instead.
var RoninGatewayV3ABI = RoninGatewayV3MetaData.ABI

// RoninGatewayV3 is an auto generated Go binding around an Ethereum contract.
type RoninGatewayV3 struct {
	RoninGatewayV3Caller     // Read-only binding to the contract
	RoninGatewayV3Transactor // Write-only binding to the contract
	RoninGatewayV3Filterer   // Log filterer for contract events
}

// RoninGatewayV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type RoninGatewayV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoninGatewayV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type RoninGatewayV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoninGatewayV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RoninGatewayV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoninGatewayV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RoninGatewayV3Session struct {
	Contract     *RoninGatewayV3   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RoninGatewayV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RoninGatewayV3CallerSession struct {
	Contract *RoninGatewayV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// RoninGatewayV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RoninGatewayV3TransactorSession struct {
	Contract     *RoninGatewayV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// RoninGatewayV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type RoninGatewayV3Raw struct {
	Contract *RoninGatewayV3 // Generic contract binding to access the raw methods on
}

// RoninGatewayV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RoninGatewayV3CallerRaw struct {
	Contract *RoninGatewayV3Caller // Generic read-only contract binding to access the raw methods on
}

// RoninGatewayV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RoninGatewayV3TransactorRaw struct {
	Contract *RoninGatewayV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewRoninGatewayV3 creates a new instance of RoninGatewayV3, bound to a specific deployed contract.
func NewRoninGatewayV3(address common.Address, backend bind.ContractBackend) (*RoninGatewayV3, error) {
	contract, err := bindRoninGatewayV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RoninGatewayV3{RoninGatewayV3Caller: RoninGatewayV3Caller{contract: contract}, RoninGatewayV3Transactor: RoninGatewayV3Transactor{contract: contract}, RoninGatewayV3Filterer: RoninGatewayV3Filterer{contract: contract}}, nil
}

// NewRoninGatewayV3Caller creates a new read-only instance of RoninGatewayV3, bound to a specific deployed contract.
func NewRoninGatewayV3Caller(address common.Address, caller bind.ContractCaller) (*RoninGatewayV3Caller, error) {
	contract, err := bindRoninGatewayV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RoninGatewayV3Caller{contract: contract}, nil
}

// NewRoninGatewayV3Transactor creates a new write-only instance of RoninGatewayV3, bound to a specific deployed contract.
func NewRoninGatewayV3Transactor(address common.Address, transactor bind.ContractTransactor) (*RoninGatewayV3Transactor, error) {
	contract, err := bindRoninGatewayV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RoninGatewayV3Transactor{contract: contract}, nil
}

// NewRoninGatewayV3Filterer creates a new log filterer instance of RoninGatewayV3, bound to a specific deployed contract.
func NewRoninGatewayV3Filterer(address common.Address, filterer bind.ContractFilterer) (*RoninGatewayV3Filterer, error) {
	contract, err := bindRoninGatewayV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RoninGatewayV3Filterer{contract: contract}, nil
}

// bindRoninGatewayV3 binds a generic wrapper to an already deployed contract.
func bindRoninGatewayV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RoninGatewayV3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoninGatewayV3 *RoninGatewayV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoninGatewayV3.Contract.RoninGatewayV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoninGatewayV3 *RoninGatewayV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.RoninGatewayV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoninGatewayV3 *RoninGatewayV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.RoninGatewayV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoninGatewayV3 *RoninGatewayV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoninGatewayV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoninGatewayV3 *RoninGatewayV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoninGatewayV3 *RoninGatewayV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RoninGatewayV3 *RoninGatewayV3Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RoninGatewayV3.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RoninGatewayV3 *RoninGatewayV3Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _RoninGatewayV3.Contract.DEFAULTADMINROLE(&_RoninGatewayV3.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RoninGatewayV3 *RoninGatewayV3CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _RoninGatewayV3.Contract.DEFAULTADMINROLE(&_RoninGatewayV3.CallOpts)
}

// WITHDRAWALMIGRATOR is a free data retrieval call binding the contract method 0xfe90d9c2.
//
// Solidity: function WITHDRAWAL_MIGRATOR() view returns(bytes32)
func (_RoninGatewayV3 *RoninGatewayV3Caller) WITHDRAWALMIGRATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RoninGatewayV3.contract.Call(opts, &out, "WITHDRAWAL_MIGRATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWALMIGRATOR is a free data retrieval call binding the contract method 0xfe90d9c2.
//
// Solidity: function WITHDRAWAL_MIGRATOR() view returns(bytes32)
func (_RoninGatewayV3 *RoninGatewayV3Session) WITHDRAWALMIGRATOR() ([32]byte, error) {
	return _RoninGatewayV3.Contract.WITHDRAWALMIGRATOR(&_RoninGatewayV3.CallOpts)
}

// WITHDRAWALMIGRATOR is a free data retrieval call binding the contract method 0xfe90d9c2.
//
// Solidity: function WITHDRAWAL_MIGRATOR() view returns(bytes32)
func (_RoninGatewayV3 *RoninGatewayV3CallerSession) WITHDRAWALMIGRATOR() ([32]byte, error) {
	return _RoninGatewayV3.Contract.WITHDRAWALMIGRATOR(&_RoninGatewayV3.CallOpts)
}

// CheckThreshold is a free data retrieval call binding the contract method 0xdafae408.
//
// Solidity: function checkThreshold(uint256 _voteWeight) view returns(bool)
func (_RoninGatewayV3 *RoninGatewayV3Caller) CheckThreshold(opts *bind.CallOpts, _voteWeight *big.Int) (bool, error) {
	var out []interface{}
	err := _RoninGatewayV3.contract.Call(opts, &out, "checkThreshold", _voteWeight)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckThreshold is a free data retrieval call binding the contract method 0xdafae408.
//
// Solidity: function checkThreshold(uint256 _voteWeight) view returns(bool)
func (_RoninGatewayV3 *RoninGatewayV3Session) CheckThreshold(_voteWeight *big.Int) (bool, error) {
	return _RoninGatewayV3.Contract.CheckThreshold(&_RoninGatewayV3.CallOpts, _voteWeight)
}

// CheckThreshold is a free data retrieval call binding the contract method 0xdafae408.
//
// Solidity: function checkThreshold(uint256 _voteWeight) view returns(bool)
func (_RoninGatewayV3 *RoninGatewayV3CallerSession) CheckThreshold(_voteWeight *big.Int) (bool, error) {
	return _RoninGatewayV3.Contract.CheckThreshold(&_RoninGatewayV3.CallOpts, _voteWeight)
}

// DepositVote is a free data retrieval call binding the contract method 0x4d92c4f0.
//
// Solidity: function depositVote(uint256 , uint256 ) view returns(uint8 status, bytes32 finalHash)
func (_RoninGatewayV3 *RoninGatewayV3Caller) DepositVote(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Status    uint8
	FinalHash [32]byte
}, error) {
	var out []interface{}
	err := _RoninGatewayV3.contract.Call(opts, &out, "depositVote", arg0, arg1)

	outstruct := new(struct {
		Status    uint8
		FinalHash [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.FinalHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// DepositVote is a free data retrieval call binding the contract method 0x4d92c4f0.
//
// Solidity: function depositVote(uint256 , uint256 ) view returns(uint8 status, bytes32 finalHash)
func (_RoninGatewayV3 *RoninGatewayV3Session) DepositVote(arg0 *big.Int, arg1 *big.Int) (struct {
	Status    uint8
	FinalHash [32]byte
}, error) {
	return _RoninGatewayV3.Contract.DepositVote(&_RoninGatewayV3.CallOpts, arg0, arg1)
}

// DepositVote is a free data retrieval call binding the contract method 0x4d92c4f0.
//
// Solidity: function depositVote(uint256 , uint256 ) view returns(uint8 status, bytes32 finalHash)
func (_RoninGatewayV3 *RoninGatewayV3CallerSession) DepositVote(arg0 *big.Int, arg1 *big.Int) (struct {
	Status    uint8
	FinalHash [32]byte
}, error) {
	return _RoninGatewayV3.Contract.DepositVote(&_RoninGatewayV3.CallOpts, arg0, arg1)
}

// DepositVoted is a free data retrieval call binding the contract method 0xfc6574bc.
//
// Solidity: function depositVoted(uint256 _chainId, uint256 _depositId, address _voter) view returns(bool)
func (_RoninGatewayV3 *RoninGatewayV3Caller) DepositVoted(opts *bind.CallOpts, _chainId *big.Int, _depositId *big.Int, _voter common.Address) (bool, error) {
	var out []interface{}
	err := _RoninGatewayV3.contract.Call(opts, &out, "depositVoted", _chainId, _depositId, _voter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DepositVoted is a free data retrieval call binding the contract method 0xfc6574bc.
//
// Solidity: function depositVoted(uint256 _chainId, uint256 _depositId, address _voter) view returns(bool)
func (_RoninGatewayV3 *RoninGatewayV3Session) DepositVoted(_chainId *big.Int, _depositId *big.Int, _voter common.Address) (bool, error) {
	return _RoninGatewayV3.Contract.DepositVoted(&_RoninGatewayV3.CallOpts, _chainId, _depositId, _voter)
}

// DepositVoted is a free data retrieval call binding the contract method 0xfc6574bc.
//
// Solidity: function depositVoted(uint256 _chainId, uint256 _depositId, address _voter) view returns(bool)
func (_RoninGatewayV3 *RoninGatewayV3CallerSession) DepositVoted(_chainId *big.Int, _depositId *big.Int, _voter common.Address) (bool, error) {
	return _RoninGatewayV3.Contract.DepositVoted(&_RoninGatewayV3.CallOpts, _chainId, _depositId, _voter)
}

// GetMainchainToken is a free data retrieval call binding the contract method 0x5d6a9a90.
//
// Solidity: function getMainchainToken(address _roninToken, uint256 _chainId) view returns((uint8,address) _token)
func (_RoninGatewayV3 *RoninGatewayV3Caller) GetMainchainToken(opts *bind.CallOpts, _roninToken common.Address, _chainId *big.Int) (MappedTokenConsumerMappedToken, error) {
	var out []interface{}
	err := _RoninGatewayV3.contract.Call(opts, &out, "getMainchainToken", _roninToken, _chainId)

	if err != nil {
		return *new(MappedTokenConsumerMappedToken), err
	}

	out0 := *abi.ConvertType(out[0], new(MappedTokenConsumerMappedToken)).(*MappedTokenConsumerMappedToken)

	return out0, err

}

// GetMainchainToken is a free data retrieval call binding the contract method 0x5d6a9a90.
//
// Solidity: function getMainchainToken(address _roninToken, uint256 _chainId) view returns((uint8,address) _token)
func (_RoninGatewayV3 *RoninGatewayV3Session) GetMainchainToken(_roninToken common.Address, _chainId *big.Int) (MappedTokenConsumerMappedToken, error) {
	return _RoninGatewayV3.Contract.GetMainchainToken(&_RoninGatewayV3.CallOpts, _roninToken, _chainId)
}

// GetMainchainToken is a free data retrieval call binding the contract method 0x5d6a9a90.
//
// Solidity: function getMainchainToken(address _roninToken, uint256 _chainId) view returns((uint8,address) _token)
func (_RoninGatewayV3 *RoninGatewayV3CallerSession) GetMainchainToken(_roninToken common.Address, _chainId *big.Int) (MappedTokenConsumerMappedToken, error) {
	return _RoninGatewayV3.Contract.GetMainchainToken(&_RoninGatewayV3.CallOpts, _roninToken, _chainId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RoninGatewayV3 *RoninGatewayV3Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _RoninGatewayV3.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RoninGatewayV3 *RoninGatewayV3Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RoninGatewayV3.Contract.GetRoleAdmin(&_RoninGatewayV3.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RoninGatewayV3 *RoninGatewayV3CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RoninGatewayV3.Contract.GetRoleAdmin(&_RoninGatewayV3.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_RoninGatewayV3 *RoninGatewayV3Caller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RoninGatewayV3.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_RoninGatewayV3 *RoninGatewayV3Session) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _RoninGatewayV3.Contract.GetRoleMember(&_RoninGatewayV3.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_RoninGatewayV3 *RoninGatewayV3CallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _RoninGatewayV3.Contract.GetRoleMember(&_RoninGatewayV3.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_RoninGatewayV3 *RoninGatewayV3Caller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _RoninGatewayV3.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_RoninGatewayV3 *RoninGatewayV3Session) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _RoninGatewayV3.Contract.GetRoleMemberCount(&_RoninGatewayV3.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_RoninGatewayV3 *RoninGatewayV3CallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _RoninGatewayV3.Contract.GetRoleMemberCount(&_RoninGatewayV3.CallOpts, role)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256, uint256)
func (_RoninGatewayV3 *RoninGatewayV3Caller) GetThreshold(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _RoninGatewayV3.contract.Call(opts, &out, "getThreshold")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256, uint256)
func (_RoninGatewayV3 *RoninGatewayV3Session) GetThreshold() (*big.Int, *big.Int, error) {
	return _RoninGatewayV3.Contract.GetThreshold(&_RoninGatewayV3.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256, uint256)
func (_RoninGatewayV3 *RoninGatewayV3CallerSession) GetThreshold() (*big.Int, *big.Int, error) {
	return _RoninGatewayV3.Contract.GetThreshold(&_RoninGatewayV3.CallOpts)
}

// GetWithdrawalSignatures is a free data retrieval call binding the contract method 0xecc83649.
//
// Solidity: function getWithdrawalSignatures(uint256 _withdrawalId, address[] _validators) view returns(bytes[] _signatures)
func (_RoninGatewayV3 *RoninGatewayV3Caller) GetWithdrawalSignatures(opts *bind.CallOpts, _withdrawalId *big.Int, _validators []common.Address) ([][]byte, error) {
	var out []interface{}
	err := _RoninGatewayV3.contract.Call(opts, &out, "getWithdrawalSignatures", _withdrawalId, _validators)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetWithdrawalSignatures is a free data retrieval call binding the contract method 0xecc83649.
//
// Solidity: function getWithdrawalSignatures(uint256 _withdrawalId, address[] _validators) view returns(bytes[] _signatures)
func (_RoninGatewayV3 *RoninGatewayV3Session) GetWithdrawalSignatures(_withdrawalId *big.Int, _validators []common.Address) ([][]byte, error) {
	return _RoninGatewayV3.Contract.GetWithdrawalSignatures(&_RoninGatewayV3.CallOpts, _withdrawalId, _validators)
}

// GetWithdrawalSignatures is a free data retrieval call binding the contract method 0xecc83649.
//
// Solidity: function getWithdrawalSignatures(uint256 _withdrawalId, address[] _validators) view returns(bytes[] _signatures)
func (_RoninGatewayV3 *RoninGatewayV3CallerSession) GetWithdrawalSignatures(_withdrawalId *big.Int, _validators []common.Address) ([][]byte, error) {
	return _RoninGatewayV3.Contract.GetWithdrawalSignatures(&_RoninGatewayV3.CallOpts, _withdrawalId, _validators)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RoninGatewayV3 *RoninGatewayV3Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _RoninGatewayV3.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RoninGatewayV3 *RoninGatewayV3Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RoninGatewayV3.Contract.HasRole(&_RoninGatewayV3.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RoninGatewayV3 *RoninGatewayV3CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RoninGatewayV3.Contract.HasRole(&_RoninGatewayV3.CallOpts, role, account)
}

// MainchainWithdrew is a free data retrieval call binding the contract method 0xf668214a.
//
// Solidity: function mainchainWithdrew(uint256 _withdrawalId) view returns(bool)
func (_RoninGatewayV3 *RoninGatewayV3Caller) MainchainWithdrew(opts *bind.CallOpts, _withdrawalId *big.Int) (bool, error) {
	var out []interface{}
	err := _RoninGatewayV3.contract.Call(opts, &out, "mainchainWithdrew", _withdrawalId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MainchainWithdrew is a free data retrieval call binding the contract method 0xf668214a.
//
// Solidity: function mainchainWithdrew(uint256 _withdrawalId) view returns(bool)
func (_RoninGatewayV3 *RoninGatewayV3Session) MainchainWithdrew(_withdrawalId *big.Int) (bool, error) {
	return _RoninGatewayV3.Contract.MainchainWithdrew(&_RoninGatewayV3.CallOpts, _withdrawalId)
}

// MainchainWithdrew is a free data retrieval call binding the contract method 0xf668214a.
//
// Solidity: function mainchainWithdrew(uint256 _withdrawalId) view returns(bool)
func (_RoninGatewayV3 *RoninGatewayV3CallerSession) MainchainWithdrew(_withdrawalId *big.Int) (bool, error) {
	return _RoninGatewayV3.Contract.MainchainWithdrew(&_RoninGatewayV3.CallOpts, _withdrawalId)
}

// MainchainWithdrewVote is a free data retrieval call binding the contract method 0xf0ce418e.
//
// Solidity: function mainchainWithdrewVote(uint256 ) view returns(uint8 status, bytes32 finalHash)
func (_RoninGatewayV3 *RoninGatewayV3Caller) MainchainWithdrewVote(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Status    uint8
	FinalHash [32]byte
}, error) {
	var out []interface{}
	err := _RoninGatewayV3.contract.Call(opts, &out, "mainchainWithdrewVote", arg0)

	outstruct := new(struct {
		Status    uint8
		FinalHash [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.FinalHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// MainchainWithdrewVote is a free data retrieval call binding the contract method 0xf0ce418e.
//
// Solidity: function mainchainWithdrewVote(uint256 ) view returns(uint8 status, bytes32 finalHash)
func (_RoninGatewayV3 *RoninGatewayV3Session) MainchainWithdrewVote(arg0 *big.Int) (struct {
	Status    uint8
	FinalHash [32]byte
}, error) {
	return _RoninGatewayV3.Contract.MainchainWithdrewVote(&_RoninGatewayV3.CallOpts, arg0)
}

// MainchainWithdrewVote is a free data retrieval call binding the contract method 0xf0ce418e.
//
// Solidity: function mainchainWithdrewVote(uint256 ) view returns(uint8 status, bytes32 finalHash)
func (_RoninGatewayV3 *RoninGatewayV3CallerSession) MainchainWithdrewVote(arg0 *big.Int) (struct {
	Status    uint8
	FinalHash [32]byte
}, error) {
	return _RoninGatewayV3.Contract.MainchainWithdrewVote(&_RoninGatewayV3.CallOpts, arg0)
}

// MinimumThreshold is a free data retrieval call binding the contract method 0xbc7f0386.
//
// Solidity: function minimumThreshold(address ) view returns(uint256)
func (_RoninGatewayV3 *RoninGatewayV3Caller) MinimumThreshold(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RoninGatewayV3.contract.Call(opts, &out, "minimumThreshold", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumThreshold is a free data retrieval call binding the contract method 0xbc7f0386.
//
// Solidity: function minimumThreshold(address ) view returns(uint256)
func (_RoninGatewayV3 *RoninGatewayV3Session) MinimumThreshold(arg0 common.Address) (*big.Int, error) {
	return _RoninGatewayV3.Contract.MinimumThreshold(&_RoninGatewayV3.CallOpts, arg0)
}

// MinimumThreshold is a free data retrieval call binding the contract method 0xbc7f0386.
//
// Solidity: function minimumThreshold(address ) view returns(uint256)
func (_RoninGatewayV3 *RoninGatewayV3CallerSession) MinimumThreshold(arg0 common.Address) (*big.Int, error) {
	return _RoninGatewayV3.Contract.MinimumThreshold(&_RoninGatewayV3.CallOpts, arg0)
}

// MinimumVoteWeight is a free data retrieval call binding the contract method 0x7de5dedd.
//
// Solidity: function minimumVoteWeight() view returns(uint256)
func (_RoninGatewayV3 *RoninGatewayV3Caller) MinimumVoteWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninGatewayV3.contract.Call(opts, &out, "minimumVoteWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumVoteWeight is a free data retrieval call binding the contract method 0x7de5dedd.
//
// Solidity: function minimumVoteWeight() view returns(uint256)
func (_RoninGatewayV3 *RoninGatewayV3Session) MinimumVoteWeight() (*big.Int, error) {
	return _RoninGatewayV3.Contract.MinimumVoteWeight(&_RoninGatewayV3.CallOpts)
}

// MinimumVoteWeight is a free data retrieval call binding the contract method 0x7de5dedd.
//
// Solidity: function minimumVoteWeight() view returns(uint256)
func (_RoninGatewayV3 *RoninGatewayV3CallerSession) MinimumVoteWeight() (*big.Int, error) {
	return _RoninGatewayV3.Contract.MinimumVoteWeight(&_RoninGatewayV3.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_RoninGatewayV3 *RoninGatewayV3Caller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninGatewayV3.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_RoninGatewayV3 *RoninGatewayV3Session) Nonce() (*big.Int, error) {
	return _RoninGatewayV3.Contract.Nonce(&_RoninGatewayV3.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_RoninGatewayV3 *RoninGatewayV3CallerSession) Nonce() (*big.Int, error) {
	return _RoninGatewayV3.Contract.Nonce(&_RoninGatewayV3.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_RoninGatewayV3 *RoninGatewayV3Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RoninGatewayV3.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_RoninGatewayV3 *RoninGatewayV3Session) Paused() (bool, error) {
	return _RoninGatewayV3.Contract.Paused(&_RoninGatewayV3.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_RoninGatewayV3 *RoninGatewayV3CallerSession) Paused() (bool, error) {
	return _RoninGatewayV3.Contract.Paused(&_RoninGatewayV3.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RoninGatewayV3 *RoninGatewayV3Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _RoninGatewayV3.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RoninGatewayV3 *RoninGatewayV3Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RoninGatewayV3.Contract.SupportsInterface(&_RoninGatewayV3.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RoninGatewayV3 *RoninGatewayV3CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RoninGatewayV3.Contract.SupportsInterface(&_RoninGatewayV3.CallOpts, interfaceId)
}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_RoninGatewayV3 *RoninGatewayV3Caller) ValidatorContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RoninGatewayV3.contract.Call(opts, &out, "validatorContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_RoninGatewayV3 *RoninGatewayV3Session) ValidatorContract() (common.Address, error) {
	return _RoninGatewayV3.Contract.ValidatorContract(&_RoninGatewayV3.CallOpts)
}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_RoninGatewayV3 *RoninGatewayV3CallerSession) ValidatorContract() (common.Address, error) {
	return _RoninGatewayV3.Contract.ValidatorContract(&_RoninGatewayV3.CallOpts)
}

// Withdrawal is a free data retrieval call binding the contract method 0x835fc6ca.
//
// Solidity: function withdrawal(uint256 ) view returns(uint256 id, uint8 kind, (address,address,uint256) mainchain, (address,address,uint256) ronin, (uint8,uint256,uint256) info)
func (_RoninGatewayV3 *RoninGatewayV3Caller) Withdrawal(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id        *big.Int
	Kind      uint8
	Mainchain TokenOwner
	Ronin     TokenOwner
	Info      TokenInfo
}, error) {
	var out []interface{}
	err := _RoninGatewayV3.contract.Call(opts, &out, "withdrawal", arg0)

	outstruct := new(struct {
		Id        *big.Int
		Kind      uint8
		Mainchain TokenOwner
		Ronin     TokenOwner
		Info      TokenInfo
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Kind = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Mainchain = *abi.ConvertType(out[2], new(TokenOwner)).(*TokenOwner)
	outstruct.Ronin = *abi.ConvertType(out[3], new(TokenOwner)).(*TokenOwner)
	outstruct.Info = *abi.ConvertType(out[4], new(TokenInfo)).(*TokenInfo)

	return *outstruct, err

}

// Withdrawal is a free data retrieval call binding the contract method 0x835fc6ca.
//
// Solidity: function withdrawal(uint256 ) view returns(uint256 id, uint8 kind, (address,address,uint256) mainchain, (address,address,uint256) ronin, (uint8,uint256,uint256) info)
func (_RoninGatewayV3 *RoninGatewayV3Session) Withdrawal(arg0 *big.Int) (struct {
	Id        *big.Int
	Kind      uint8
	Mainchain TokenOwner
	Ronin     TokenOwner
	Info      TokenInfo
}, error) {
	return _RoninGatewayV3.Contract.Withdrawal(&_RoninGatewayV3.CallOpts, arg0)
}

// Withdrawal is a free data retrieval call binding the contract method 0x835fc6ca.
//
// Solidity: function withdrawal(uint256 ) view returns(uint256 id, uint8 kind, (address,address,uint256) mainchain, (address,address,uint256) ronin, (uint8,uint256,uint256) info)
func (_RoninGatewayV3 *RoninGatewayV3CallerSession) Withdrawal(arg0 *big.Int) (struct {
	Id        *big.Int
	Kind      uint8
	Mainchain TokenOwner
	Ronin     TokenOwner
	Info      TokenInfo
}, error) {
	return _RoninGatewayV3.Contract.Withdrawal(&_RoninGatewayV3.CallOpts, arg0)
}

// WithdrawalCount is a free data retrieval call binding the contract method 0x71706cbe.
//
// Solidity: function withdrawalCount() view returns(uint256)
func (_RoninGatewayV3 *RoninGatewayV3Caller) WithdrawalCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninGatewayV3.contract.Call(opts, &out, "withdrawalCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalCount is a free data retrieval call binding the contract method 0x71706cbe.
//
// Solidity: function withdrawalCount() view returns(uint256)
func (_RoninGatewayV3 *RoninGatewayV3Session) WithdrawalCount() (*big.Int, error) {
	return _RoninGatewayV3.Contract.WithdrawalCount(&_RoninGatewayV3.CallOpts)
}

// WithdrawalCount is a free data retrieval call binding the contract method 0x71706cbe.
//
// Solidity: function withdrawalCount() view returns(uint256)
func (_RoninGatewayV3 *RoninGatewayV3CallerSession) WithdrawalCount() (*big.Int, error) {
	return _RoninGatewayV3.Contract.WithdrawalCount(&_RoninGatewayV3.CallOpts)
}

// WithdrawalMigrated is a free data retrieval call binding the contract method 0x4f2717c7.
//
// Solidity: function withdrawalMigrated() view returns(bool)
func (_RoninGatewayV3 *RoninGatewayV3Caller) WithdrawalMigrated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RoninGatewayV3.contract.Call(opts, &out, "withdrawalMigrated")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawalMigrated is a free data retrieval call binding the contract method 0x4f2717c7.
//
// Solidity: function withdrawalMigrated() view returns(bool)
func (_RoninGatewayV3 *RoninGatewayV3Session) WithdrawalMigrated() (bool, error) {
	return _RoninGatewayV3.Contract.WithdrawalMigrated(&_RoninGatewayV3.CallOpts)
}

// WithdrawalMigrated is a free data retrieval call binding the contract method 0x4f2717c7.
//
// Solidity: function withdrawalMigrated() view returns(bool)
func (_RoninGatewayV3 *RoninGatewayV3CallerSession) WithdrawalMigrated() (bool, error) {
	return _RoninGatewayV3.Contract.WithdrawalMigrated(&_RoninGatewayV3.CallOpts)
}

// BulkRequestWithdrawalFor is a paid mutator transaction binding the contract method 0x5a7dd06a.
//
// Solidity: function bulkRequestWithdrawalFor((address,address,(uint8,uint256,uint256))[] _requests, uint256 _chainId) returns()
func (_RoninGatewayV3 *RoninGatewayV3Transactor) BulkRequestWithdrawalFor(opts *bind.TransactOpts, _requests []TransferRequest, _chainId *big.Int) (*types.Transaction, error) {
	return _RoninGatewayV3.contract.Transact(opts, "bulkRequestWithdrawalFor", _requests, _chainId)
}

// BulkRequestWithdrawalFor is a paid mutator transaction binding the contract method 0x5a7dd06a.
//
// Solidity: function bulkRequestWithdrawalFor((address,address,(uint8,uint256,uint256))[] _requests, uint256 _chainId) returns()
func (_RoninGatewayV3 *RoninGatewayV3Session) BulkRequestWithdrawalFor(_requests []TransferRequest, _chainId *big.Int) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.BulkRequestWithdrawalFor(&_RoninGatewayV3.TransactOpts, _requests, _chainId)
}

// BulkRequestWithdrawalFor is a paid mutator transaction binding the contract method 0x5a7dd06a.
//
// Solidity: function bulkRequestWithdrawalFor((address,address,(uint8,uint256,uint256))[] _requests, uint256 _chainId) returns()
func (_RoninGatewayV3 *RoninGatewayV3TransactorSession) BulkRequestWithdrawalFor(_requests []TransferRequest, _chainId *big.Int) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.BulkRequestWithdrawalFor(&_RoninGatewayV3.TransactOpts, _requests, _chainId)
}

// BulkSubmitWithdrawalSignatures is a paid mutator transaction binding the contract method 0xfa389659.
//
// Solidity: function bulkSubmitWithdrawalSignatures(uint256[] _withdrawals, bytes[] _signatures) returns()
func (_RoninGatewayV3 *RoninGatewayV3Transactor) BulkSubmitWithdrawalSignatures(opts *bind.TransactOpts, _withdrawals []*big.Int, _signatures [][]byte) (*types.Transaction, error) {
	return _RoninGatewayV3.contract.Transact(opts, "bulkSubmitWithdrawalSignatures", _withdrawals, _signatures)
}

// BulkSubmitWithdrawalSignatures is a paid mutator transaction binding the contract method 0xfa389659.
//
// Solidity: function bulkSubmitWithdrawalSignatures(uint256[] _withdrawals, bytes[] _signatures) returns()
func (_RoninGatewayV3 *RoninGatewayV3Session) BulkSubmitWithdrawalSignatures(_withdrawals []*big.Int, _signatures [][]byte) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.BulkSubmitWithdrawalSignatures(&_RoninGatewayV3.TransactOpts, _withdrawals, _signatures)
}

// BulkSubmitWithdrawalSignatures is a paid mutator transaction binding the contract method 0xfa389659.
//
// Solidity: function bulkSubmitWithdrawalSignatures(uint256[] _withdrawals, bytes[] _signatures) returns()
func (_RoninGatewayV3 *RoninGatewayV3TransactorSession) BulkSubmitWithdrawalSignatures(_withdrawals []*big.Int, _signatures [][]byte) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.BulkSubmitWithdrawalSignatures(&_RoninGatewayV3.TransactOpts, _withdrawals, _signatures)
}

// DepositFor is a paid mutator transaction binding the contract method 0x109679ef.
//
// Solidity: function depositFor((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) _receipt) returns()
func (_RoninGatewayV3 *RoninGatewayV3Transactor) DepositFor(opts *bind.TransactOpts, _receipt TransferReceipt) (*types.Transaction, error) {
	return _RoninGatewayV3.contract.Transact(opts, "depositFor", _receipt)
}

// DepositFor is a paid mutator transaction binding the contract method 0x109679ef.
//
// Solidity: function depositFor((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) _receipt) returns()
func (_RoninGatewayV3 *RoninGatewayV3Session) DepositFor(_receipt TransferReceipt) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.DepositFor(&_RoninGatewayV3.TransactOpts, _receipt)
}

// DepositFor is a paid mutator transaction binding the contract method 0x109679ef.
//
// Solidity: function depositFor((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) _receipt) returns()
func (_RoninGatewayV3 *RoninGatewayV3TransactorSession) DepositFor(_receipt TransferReceipt) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.DepositFor(&_RoninGatewayV3.TransactOpts, _receipt)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RoninGatewayV3 *RoninGatewayV3Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RoninGatewayV3.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RoninGatewayV3 *RoninGatewayV3Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.GrantRole(&_RoninGatewayV3.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RoninGatewayV3 *RoninGatewayV3TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.GrantRole(&_RoninGatewayV3.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xe22fe336.
//
// Solidity: function initialize(address _roleSetter, address _validatorContract, uint256 _numerator, uint256 _denominator, address[] _withdrawalMigrators, address[][2] _packedAddresses, uint256[][2] _packedNumbers, uint8[] _standards) returns()
func (_RoninGatewayV3 *RoninGatewayV3Transactor) Initialize(opts *bind.TransactOpts, _roleSetter common.Address, _validatorContract common.Address, _numerator *big.Int, _denominator *big.Int, _withdrawalMigrators []common.Address, _packedAddresses [2][]common.Address, _packedNumbers [2][]*big.Int, _standards []uint8) (*types.Transaction, error) {
	return _RoninGatewayV3.contract.Transact(opts, "initialize", _roleSetter, _validatorContract, _numerator, _denominator, _withdrawalMigrators, _packedAddresses, _packedNumbers, _standards)
}

// Initialize is a paid mutator transaction binding the contract method 0xe22fe336.
//
// Solidity: function initialize(address _roleSetter, address _validatorContract, uint256 _numerator, uint256 _denominator, address[] _withdrawalMigrators, address[][2] _packedAddresses, uint256[][2] _packedNumbers, uint8[] _standards) returns()
func (_RoninGatewayV3 *RoninGatewayV3Session) Initialize(_roleSetter common.Address, _validatorContract common.Address, _numerator *big.Int, _denominator *big.Int, _withdrawalMigrators []common.Address, _packedAddresses [2][]common.Address, _packedNumbers [2][]*big.Int, _standards []uint8) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.Initialize(&_RoninGatewayV3.TransactOpts, _roleSetter, _validatorContract, _numerator, _denominator, _withdrawalMigrators, _packedAddresses, _packedNumbers, _standards)
}

// Initialize is a paid mutator transaction binding the contract method 0xe22fe336.
//
// Solidity: function initialize(address _roleSetter, address _validatorContract, uint256 _numerator, uint256 _denominator, address[] _withdrawalMigrators, address[][2] _packedAddresses, uint256[][2] _packedNumbers, uint8[] _standards) returns()
func (_RoninGatewayV3 *RoninGatewayV3TransactorSession) Initialize(_roleSetter common.Address, _validatorContract common.Address, _numerator *big.Int, _denominator *big.Int, _withdrawalMigrators []common.Address, _packedAddresses [2][]common.Address, _packedNumbers [2][]*big.Int, _standards []uint8) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.Initialize(&_RoninGatewayV3.TransactOpts, _roleSetter, _validatorContract, _numerator, _denominator, _withdrawalMigrators, _packedAddresses, _packedNumbers, _standards)
}

// MapTokens is a paid mutator transaction binding the contract method 0xdbd2ef6c.
//
// Solidity: function mapTokens(address[] _roninTokens, address[] _mainchainTokens, uint256[] _chainIds, uint8[] _standards) returns()
func (_RoninGatewayV3 *RoninGatewayV3Transactor) MapTokens(opts *bind.TransactOpts, _roninTokens []common.Address, _mainchainTokens []common.Address, _chainIds []*big.Int, _standards []uint8) (*types.Transaction, error) {
	return _RoninGatewayV3.contract.Transact(opts, "mapTokens", _roninTokens, _mainchainTokens, _chainIds, _standards)
}

// MapTokens is a paid mutator transaction binding the contract method 0xdbd2ef6c.
//
// Solidity: function mapTokens(address[] _roninTokens, address[] _mainchainTokens, uint256[] _chainIds, uint8[] _standards) returns()
func (_RoninGatewayV3 *RoninGatewayV3Session) MapTokens(_roninTokens []common.Address, _mainchainTokens []common.Address, _chainIds []*big.Int, _standards []uint8) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.MapTokens(&_RoninGatewayV3.TransactOpts, _roninTokens, _mainchainTokens, _chainIds, _standards)
}

// MapTokens is a paid mutator transaction binding the contract method 0xdbd2ef6c.
//
// Solidity: function mapTokens(address[] _roninTokens, address[] _mainchainTokens, uint256[] _chainIds, uint8[] _standards) returns()
func (_RoninGatewayV3 *RoninGatewayV3TransactorSession) MapTokens(_roninTokens []common.Address, _mainchainTokens []common.Address, _chainIds []*big.Int, _standards []uint8) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.MapTokens(&_RoninGatewayV3.TransactOpts, _roninTokens, _mainchainTokens, _chainIds, _standards)
}

// MarkWithdrawalMigrated is a paid mutator transaction binding the contract method 0x3b5afc22.
//
// Solidity: function markWithdrawalMigrated() returns()
func (_RoninGatewayV3 *RoninGatewayV3Transactor) MarkWithdrawalMigrated(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoninGatewayV3.contract.Transact(opts, "markWithdrawalMigrated")
}

// MarkWithdrawalMigrated is a paid mutator transaction binding the contract method 0x3b5afc22.
//
// Solidity: function markWithdrawalMigrated() returns()
func (_RoninGatewayV3 *RoninGatewayV3Session) MarkWithdrawalMigrated() (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.MarkWithdrawalMigrated(&_RoninGatewayV3.TransactOpts)
}

// MarkWithdrawalMigrated is a paid mutator transaction binding the contract method 0x3b5afc22.
//
// Solidity: function markWithdrawalMigrated() returns()
func (_RoninGatewayV3 *RoninGatewayV3TransactorSession) MarkWithdrawalMigrated() (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.MarkWithdrawalMigrated(&_RoninGatewayV3.TransactOpts)
}

// MigrateWithdrawals is a paid mutator transaction binding the contract method 0x931ec987.
//
// Solidity: function migrateWithdrawals((address,address,(uint8,uint256,uint256))[] _requests, address[] _requesters) returns()
func (_RoninGatewayV3 *RoninGatewayV3Transactor) MigrateWithdrawals(opts *bind.TransactOpts, _requests []TransferRequest, _requesters []common.Address) (*types.Transaction, error) {
	return _RoninGatewayV3.contract.Transact(opts, "migrateWithdrawals", _requests, _requesters)
}

// MigrateWithdrawals is a paid mutator transaction binding the contract method 0x931ec987.
//
// Solidity: function migrateWithdrawals((address,address,(uint8,uint256,uint256))[] _requests, address[] _requesters) returns()
func (_RoninGatewayV3 *RoninGatewayV3Session) MigrateWithdrawals(_requests []TransferRequest, _requesters []common.Address) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.MigrateWithdrawals(&_RoninGatewayV3.TransactOpts, _requests, _requesters)
}

// MigrateWithdrawals is a paid mutator transaction binding the contract method 0x931ec987.
//
// Solidity: function migrateWithdrawals((address,address,(uint8,uint256,uint256))[] _requests, address[] _requesters) returns()
func (_RoninGatewayV3 *RoninGatewayV3TransactorSession) MigrateWithdrawals(_requests []TransferRequest, _requesters []common.Address) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.MigrateWithdrawals(&_RoninGatewayV3.TransactOpts, _requests, _requesters)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RoninGatewayV3 *RoninGatewayV3Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoninGatewayV3.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RoninGatewayV3 *RoninGatewayV3Session) Pause() (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.Pause(&_RoninGatewayV3.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RoninGatewayV3 *RoninGatewayV3TransactorSession) Pause() (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.Pause(&_RoninGatewayV3.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RoninGatewayV3 *RoninGatewayV3Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RoninGatewayV3.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RoninGatewayV3 *RoninGatewayV3Session) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.RenounceRole(&_RoninGatewayV3.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RoninGatewayV3 *RoninGatewayV3TransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.RenounceRole(&_RoninGatewayV3.TransactOpts, role, account)
}

// RequestWithdrawalFor is a paid mutator transaction binding the contract method 0x0b1ff17f.
//
// Solidity: function requestWithdrawalFor((address,address,(uint8,uint256,uint256)) _request, uint256 _chainId) returns()
func (_RoninGatewayV3 *RoninGatewayV3Transactor) RequestWithdrawalFor(opts *bind.TransactOpts, _request TransferRequest, _chainId *big.Int) (*types.Transaction, error) {
	return _RoninGatewayV3.contract.Transact(opts, "requestWithdrawalFor", _request, _chainId)
}

// RequestWithdrawalFor is a paid mutator transaction binding the contract method 0x0b1ff17f.
//
// Solidity: function requestWithdrawalFor((address,address,(uint8,uint256,uint256)) _request, uint256 _chainId) returns()
func (_RoninGatewayV3 *RoninGatewayV3Session) RequestWithdrawalFor(_request TransferRequest, _chainId *big.Int) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.RequestWithdrawalFor(&_RoninGatewayV3.TransactOpts, _request, _chainId)
}

// RequestWithdrawalFor is a paid mutator transaction binding the contract method 0x0b1ff17f.
//
// Solidity: function requestWithdrawalFor((address,address,(uint8,uint256,uint256)) _request, uint256 _chainId) returns()
func (_RoninGatewayV3 *RoninGatewayV3TransactorSession) RequestWithdrawalFor(_request TransferRequest, _chainId *big.Int) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.RequestWithdrawalFor(&_RoninGatewayV3.TransactOpts, _request, _chainId)
}

// RequestWithdrawalSignatures is a paid mutator transaction binding the contract method 0x47b56b2c.
//
// Solidity: function requestWithdrawalSignatures(uint256 _withdrawalId) returns()
func (_RoninGatewayV3 *RoninGatewayV3Transactor) RequestWithdrawalSignatures(opts *bind.TransactOpts, _withdrawalId *big.Int) (*types.Transaction, error) {
	return _RoninGatewayV3.contract.Transact(opts, "requestWithdrawalSignatures", _withdrawalId)
}

// RequestWithdrawalSignatures is a paid mutator transaction binding the contract method 0x47b56b2c.
//
// Solidity: function requestWithdrawalSignatures(uint256 _withdrawalId) returns()
func (_RoninGatewayV3 *RoninGatewayV3Session) RequestWithdrawalSignatures(_withdrawalId *big.Int) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.RequestWithdrawalSignatures(&_RoninGatewayV3.TransactOpts, _withdrawalId)
}

// RequestWithdrawalSignatures is a paid mutator transaction binding the contract method 0x47b56b2c.
//
// Solidity: function requestWithdrawalSignatures(uint256 _withdrawalId) returns()
func (_RoninGatewayV3 *RoninGatewayV3TransactorSession) RequestWithdrawalSignatures(_withdrawalId *big.Int) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.RequestWithdrawalSignatures(&_RoninGatewayV3.TransactOpts, _withdrawalId)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RoninGatewayV3 *RoninGatewayV3Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RoninGatewayV3.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RoninGatewayV3 *RoninGatewayV3Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.RevokeRole(&_RoninGatewayV3.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RoninGatewayV3 *RoninGatewayV3TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.RevokeRole(&_RoninGatewayV3.TransactOpts, role, account)
}

// SetMinimumThresholds is a paid mutator transaction binding the contract method 0x64363f78.
//
// Solidity: function setMinimumThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_RoninGatewayV3 *RoninGatewayV3Transactor) SetMinimumThresholds(opts *bind.TransactOpts, _tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _RoninGatewayV3.contract.Transact(opts, "setMinimumThresholds", _tokens, _thresholds)
}

// SetMinimumThresholds is a paid mutator transaction binding the contract method 0x64363f78.
//
// Solidity: function setMinimumThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_RoninGatewayV3 *RoninGatewayV3Session) SetMinimumThresholds(_tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.SetMinimumThresholds(&_RoninGatewayV3.TransactOpts, _tokens, _thresholds)
}

// SetMinimumThresholds is a paid mutator transaction binding the contract method 0x64363f78.
//
// Solidity: function setMinimumThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_RoninGatewayV3 *RoninGatewayV3TransactorSession) SetMinimumThresholds(_tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.SetMinimumThresholds(&_RoninGatewayV3.TransactOpts, _tokens, _thresholds)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xb9c36209.
//
// Solidity: function setThreshold(uint256 _numerator, uint256 _denominator) returns(uint256, uint256)
func (_RoninGatewayV3 *RoninGatewayV3Transactor) SetThreshold(opts *bind.TransactOpts, _numerator *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _RoninGatewayV3.contract.Transact(opts, "setThreshold", _numerator, _denominator)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xb9c36209.
//
// Solidity: function setThreshold(uint256 _numerator, uint256 _denominator) returns(uint256, uint256)
func (_RoninGatewayV3 *RoninGatewayV3Session) SetThreshold(_numerator *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.SetThreshold(&_RoninGatewayV3.TransactOpts, _numerator, _denominator)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xb9c36209.
//
// Solidity: function setThreshold(uint256 _numerator, uint256 _denominator) returns(uint256, uint256)
func (_RoninGatewayV3 *RoninGatewayV3TransactorSession) SetThreshold(_numerator *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.SetThreshold(&_RoninGatewayV3.TransactOpts, _numerator, _denominator)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _validatorContract) returns()
func (_RoninGatewayV3 *RoninGatewayV3Transactor) SetValidatorContract(opts *bind.TransactOpts, _validatorContract common.Address) (*types.Transaction, error) {
	return _RoninGatewayV3.contract.Transact(opts, "setValidatorContract", _validatorContract)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _validatorContract) returns()
func (_RoninGatewayV3 *RoninGatewayV3Session) SetValidatorContract(_validatorContract common.Address) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.SetValidatorContract(&_RoninGatewayV3.TransactOpts, _validatorContract)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _validatorContract) returns()
func (_RoninGatewayV3 *RoninGatewayV3TransactorSession) SetValidatorContract(_validatorContract common.Address) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.SetValidatorContract(&_RoninGatewayV3.TransactOpts, _validatorContract)
}

// TryBulkAcknowledgeMainchainWithdrew is a paid mutator transaction binding the contract method 0x17fa2ea1.
//
// Solidity: function tryBulkAcknowledgeMainchainWithdrew(uint256[] _withdrawalIds) returns(bool[] _executedReceipts)
func (_RoninGatewayV3 *RoninGatewayV3Transactor) TryBulkAcknowledgeMainchainWithdrew(opts *bind.TransactOpts, _withdrawalIds []*big.Int) (*types.Transaction, error) {
	return _RoninGatewayV3.contract.Transact(opts, "tryBulkAcknowledgeMainchainWithdrew", _withdrawalIds)
}

// TryBulkAcknowledgeMainchainWithdrew is a paid mutator transaction binding the contract method 0x17fa2ea1.
//
// Solidity: function tryBulkAcknowledgeMainchainWithdrew(uint256[] _withdrawalIds) returns(bool[] _executedReceipts)
func (_RoninGatewayV3 *RoninGatewayV3Session) TryBulkAcknowledgeMainchainWithdrew(_withdrawalIds []*big.Int) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.TryBulkAcknowledgeMainchainWithdrew(&_RoninGatewayV3.TransactOpts, _withdrawalIds)
}

// TryBulkAcknowledgeMainchainWithdrew is a paid mutator transaction binding the contract method 0x17fa2ea1.
//
// Solidity: function tryBulkAcknowledgeMainchainWithdrew(uint256[] _withdrawalIds) returns(bool[] _executedReceipts)
func (_RoninGatewayV3 *RoninGatewayV3TransactorSession) TryBulkAcknowledgeMainchainWithdrew(_withdrawalIds []*big.Int) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.TryBulkAcknowledgeMainchainWithdrew(&_RoninGatewayV3.TransactOpts, _withdrawalIds)
}

// TryBulkDepositFor is a paid mutator transaction binding the contract method 0xb9afa177.
//
// Solidity: function tryBulkDepositFor((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256))[] _receipts) returns(bool[] _executedReceipts)
func (_RoninGatewayV3 *RoninGatewayV3Transactor) TryBulkDepositFor(opts *bind.TransactOpts, _receipts []TransferReceipt) (*types.Transaction, error) {
	return _RoninGatewayV3.contract.Transact(opts, "tryBulkDepositFor", _receipts)
}

// TryBulkDepositFor is a paid mutator transaction binding the contract method 0xb9afa177.
//
// Solidity: function tryBulkDepositFor((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256))[] _receipts) returns(bool[] _executedReceipts)
func (_RoninGatewayV3 *RoninGatewayV3Session) TryBulkDepositFor(_receipts []TransferReceipt) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.TryBulkDepositFor(&_RoninGatewayV3.TransactOpts, _receipts)
}

// TryBulkDepositFor is a paid mutator transaction binding the contract method 0xb9afa177.
//
// Solidity: function tryBulkDepositFor((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256))[] _receipts) returns(bool[] _executedReceipts)
func (_RoninGatewayV3 *RoninGatewayV3TransactorSession) TryBulkDepositFor(_receipts []TransferReceipt) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.TryBulkDepositFor(&_RoninGatewayV3.TransactOpts, _receipts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RoninGatewayV3 *RoninGatewayV3Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoninGatewayV3.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RoninGatewayV3 *RoninGatewayV3Session) Unpause() (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.Unpause(&_RoninGatewayV3.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RoninGatewayV3 *RoninGatewayV3TransactorSession) Unpause() (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.Unpause(&_RoninGatewayV3.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_RoninGatewayV3 *RoninGatewayV3Transactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _RoninGatewayV3.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_RoninGatewayV3 *RoninGatewayV3Session) Fallback(calldata []byte) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.Fallback(&_RoninGatewayV3.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_RoninGatewayV3 *RoninGatewayV3TransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.Fallback(&_RoninGatewayV3.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RoninGatewayV3 *RoninGatewayV3Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoninGatewayV3.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RoninGatewayV3 *RoninGatewayV3Session) Receive() (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.Receive(&_RoninGatewayV3.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RoninGatewayV3 *RoninGatewayV3TransactorSession) Receive() (*types.Transaction, error) {
	return _RoninGatewayV3.Contract.Receive(&_RoninGatewayV3.TransactOpts)
}

// RoninGatewayV3DepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the RoninGatewayV3 contract.
type RoninGatewayV3DepositedIterator struct {
	Event *RoninGatewayV3Deposited // Event containing the contract specifics and raw log

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
func (it *RoninGatewayV3DepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninGatewayV3Deposited)
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
		it.Event = new(RoninGatewayV3Deposited)
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
func (it *RoninGatewayV3DepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninGatewayV3DepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninGatewayV3Deposited represents a Deposited event raised by the RoninGatewayV3 contract.
type RoninGatewayV3Deposited struct {
	ReceiptHash [32]byte
	Receipt     TransferReceipt
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x8d20d8121a34dded9035ff5b43e901c142824f7a22126392992c353c37890524.
//
// Solidity: event Deposited(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_RoninGatewayV3 *RoninGatewayV3Filterer) FilterDeposited(opts *bind.FilterOpts) (*RoninGatewayV3DepositedIterator, error) {

	logs, sub, err := _RoninGatewayV3.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &RoninGatewayV3DepositedIterator{contract: _RoninGatewayV3.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x8d20d8121a34dded9035ff5b43e901c142824f7a22126392992c353c37890524.
//
// Solidity: event Deposited(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_RoninGatewayV3 *RoninGatewayV3Filterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *RoninGatewayV3Deposited) (event.Subscription, error) {

	logs, sub, err := _RoninGatewayV3.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninGatewayV3Deposited)
				if err := _RoninGatewayV3.contract.UnpackLog(event, "Deposited", log); err != nil {
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
func (_RoninGatewayV3 *RoninGatewayV3Filterer) ParseDeposited(log types.Log) (*RoninGatewayV3Deposited, error) {
	event := new(RoninGatewayV3Deposited)
	if err := _RoninGatewayV3.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninGatewayV3MainchainWithdrewIterator is returned from FilterMainchainWithdrew and is used to iterate over the raw logs and unpacked data for MainchainWithdrew events raised by the RoninGatewayV3 contract.
type RoninGatewayV3MainchainWithdrewIterator struct {
	Event *RoninGatewayV3MainchainWithdrew // Event containing the contract specifics and raw log

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
func (it *RoninGatewayV3MainchainWithdrewIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninGatewayV3MainchainWithdrew)
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
		it.Event = new(RoninGatewayV3MainchainWithdrew)
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
func (it *RoninGatewayV3MainchainWithdrewIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninGatewayV3MainchainWithdrewIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninGatewayV3MainchainWithdrew represents a MainchainWithdrew event raised by the RoninGatewayV3 contract.
type RoninGatewayV3MainchainWithdrew struct {
	ReceiptHash [32]byte
	Receipt     TransferReceipt
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMainchainWithdrew is a free log retrieval operation binding the contract event 0x62520d049932cdee872e9b3c59c0f6073637147e5e9bc8b050b062430eaf5c9f.
//
// Solidity: event MainchainWithdrew(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_RoninGatewayV3 *RoninGatewayV3Filterer) FilterMainchainWithdrew(opts *bind.FilterOpts) (*RoninGatewayV3MainchainWithdrewIterator, error) {

	logs, sub, err := _RoninGatewayV3.contract.FilterLogs(opts, "MainchainWithdrew")
	if err != nil {
		return nil, err
	}
	return &RoninGatewayV3MainchainWithdrewIterator{contract: _RoninGatewayV3.contract, event: "MainchainWithdrew", logs: logs, sub: sub}, nil
}

// WatchMainchainWithdrew is a free log subscription operation binding the contract event 0x62520d049932cdee872e9b3c59c0f6073637147e5e9bc8b050b062430eaf5c9f.
//
// Solidity: event MainchainWithdrew(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_RoninGatewayV3 *RoninGatewayV3Filterer) WatchMainchainWithdrew(opts *bind.WatchOpts, sink chan<- *RoninGatewayV3MainchainWithdrew) (event.Subscription, error) {

	logs, sub, err := _RoninGatewayV3.contract.WatchLogs(opts, "MainchainWithdrew")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninGatewayV3MainchainWithdrew)
				if err := _RoninGatewayV3.contract.UnpackLog(event, "MainchainWithdrew", log); err != nil {
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
func (_RoninGatewayV3 *RoninGatewayV3Filterer) ParseMainchainWithdrew(log types.Log) (*RoninGatewayV3MainchainWithdrew, error) {
	event := new(RoninGatewayV3MainchainWithdrew)
	if err := _RoninGatewayV3.contract.UnpackLog(event, "MainchainWithdrew", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninGatewayV3MinimumThresholdsUpdatedIterator is returned from FilterMinimumThresholdsUpdated and is used to iterate over the raw logs and unpacked data for MinimumThresholdsUpdated events raised by the RoninGatewayV3 contract.
type RoninGatewayV3MinimumThresholdsUpdatedIterator struct {
	Event *RoninGatewayV3MinimumThresholdsUpdated // Event containing the contract specifics and raw log

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
func (it *RoninGatewayV3MinimumThresholdsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninGatewayV3MinimumThresholdsUpdated)
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
		it.Event = new(RoninGatewayV3MinimumThresholdsUpdated)
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
func (it *RoninGatewayV3MinimumThresholdsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninGatewayV3MinimumThresholdsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninGatewayV3MinimumThresholdsUpdated represents a MinimumThresholdsUpdated event raised by the RoninGatewayV3 contract.
type RoninGatewayV3MinimumThresholdsUpdated struct {
	Tokens    []common.Address
	Threshold []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinimumThresholdsUpdated is a free log retrieval operation binding the contract event 0x6f52f53a938df83439fa4c6055c7df0a6906d621aa6dfa4708187037fdfc41da.
//
// Solidity: event MinimumThresholdsUpdated(address[] tokens, uint256[] threshold)
func (_RoninGatewayV3 *RoninGatewayV3Filterer) FilterMinimumThresholdsUpdated(opts *bind.FilterOpts) (*RoninGatewayV3MinimumThresholdsUpdatedIterator, error) {

	logs, sub, err := _RoninGatewayV3.contract.FilterLogs(opts, "MinimumThresholdsUpdated")
	if err != nil {
		return nil, err
	}
	return &RoninGatewayV3MinimumThresholdsUpdatedIterator{contract: _RoninGatewayV3.contract, event: "MinimumThresholdsUpdated", logs: logs, sub: sub}, nil
}

// WatchMinimumThresholdsUpdated is a free log subscription operation binding the contract event 0x6f52f53a938df83439fa4c6055c7df0a6906d621aa6dfa4708187037fdfc41da.
//
// Solidity: event MinimumThresholdsUpdated(address[] tokens, uint256[] threshold)
func (_RoninGatewayV3 *RoninGatewayV3Filterer) WatchMinimumThresholdsUpdated(opts *bind.WatchOpts, sink chan<- *RoninGatewayV3MinimumThresholdsUpdated) (event.Subscription, error) {

	logs, sub, err := _RoninGatewayV3.contract.WatchLogs(opts, "MinimumThresholdsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninGatewayV3MinimumThresholdsUpdated)
				if err := _RoninGatewayV3.contract.UnpackLog(event, "MinimumThresholdsUpdated", log); err != nil {
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

// ParseMinimumThresholdsUpdated is a log parse operation binding the contract event 0x6f52f53a938df83439fa4c6055c7df0a6906d621aa6dfa4708187037fdfc41da.
//
// Solidity: event MinimumThresholdsUpdated(address[] tokens, uint256[] threshold)
func (_RoninGatewayV3 *RoninGatewayV3Filterer) ParseMinimumThresholdsUpdated(log types.Log) (*RoninGatewayV3MinimumThresholdsUpdated, error) {
	event := new(RoninGatewayV3MinimumThresholdsUpdated)
	if err := _RoninGatewayV3.contract.UnpackLog(event, "MinimumThresholdsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninGatewayV3PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the RoninGatewayV3 contract.
type RoninGatewayV3PausedIterator struct {
	Event *RoninGatewayV3Paused // Event containing the contract specifics and raw log

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
func (it *RoninGatewayV3PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninGatewayV3Paused)
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
		it.Event = new(RoninGatewayV3Paused)
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
func (it *RoninGatewayV3PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninGatewayV3PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninGatewayV3Paused represents a Paused event raised by the RoninGatewayV3 contract.
type RoninGatewayV3Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_RoninGatewayV3 *RoninGatewayV3Filterer) FilterPaused(opts *bind.FilterOpts) (*RoninGatewayV3PausedIterator, error) {

	logs, sub, err := _RoninGatewayV3.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &RoninGatewayV3PausedIterator{contract: _RoninGatewayV3.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_RoninGatewayV3 *RoninGatewayV3Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *RoninGatewayV3Paused) (event.Subscription, error) {

	logs, sub, err := _RoninGatewayV3.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninGatewayV3Paused)
				if err := _RoninGatewayV3.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_RoninGatewayV3 *RoninGatewayV3Filterer) ParsePaused(log types.Log) (*RoninGatewayV3Paused, error) {
	event := new(RoninGatewayV3Paused)
	if err := _RoninGatewayV3.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninGatewayV3RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the RoninGatewayV3 contract.
type RoninGatewayV3RoleAdminChangedIterator struct {
	Event *RoninGatewayV3RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *RoninGatewayV3RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninGatewayV3RoleAdminChanged)
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
		it.Event = new(RoninGatewayV3RoleAdminChanged)
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
func (it *RoninGatewayV3RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninGatewayV3RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninGatewayV3RoleAdminChanged represents a RoleAdminChanged event raised by the RoninGatewayV3 contract.
type RoninGatewayV3RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RoninGatewayV3 *RoninGatewayV3Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*RoninGatewayV3RoleAdminChangedIterator, error) {

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

	logs, sub, err := _RoninGatewayV3.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &RoninGatewayV3RoleAdminChangedIterator{contract: _RoninGatewayV3.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RoninGatewayV3 *RoninGatewayV3Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *RoninGatewayV3RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _RoninGatewayV3.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninGatewayV3RoleAdminChanged)
				if err := _RoninGatewayV3.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_RoninGatewayV3 *RoninGatewayV3Filterer) ParseRoleAdminChanged(log types.Log) (*RoninGatewayV3RoleAdminChanged, error) {
	event := new(RoninGatewayV3RoleAdminChanged)
	if err := _RoninGatewayV3.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninGatewayV3RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the RoninGatewayV3 contract.
type RoninGatewayV3RoleGrantedIterator struct {
	Event *RoninGatewayV3RoleGranted // Event containing the contract specifics and raw log

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
func (it *RoninGatewayV3RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninGatewayV3RoleGranted)
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
		it.Event = new(RoninGatewayV3RoleGranted)
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
func (it *RoninGatewayV3RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninGatewayV3RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninGatewayV3RoleGranted represents a RoleGranted event raised by the RoninGatewayV3 contract.
type RoninGatewayV3RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RoninGatewayV3 *RoninGatewayV3Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RoninGatewayV3RoleGrantedIterator, error) {

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

	logs, sub, err := _RoninGatewayV3.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RoninGatewayV3RoleGrantedIterator{contract: _RoninGatewayV3.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RoninGatewayV3 *RoninGatewayV3Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *RoninGatewayV3RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _RoninGatewayV3.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninGatewayV3RoleGranted)
				if err := _RoninGatewayV3.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_RoninGatewayV3 *RoninGatewayV3Filterer) ParseRoleGranted(log types.Log) (*RoninGatewayV3RoleGranted, error) {
	event := new(RoninGatewayV3RoleGranted)
	if err := _RoninGatewayV3.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninGatewayV3RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the RoninGatewayV3 contract.
type RoninGatewayV3RoleRevokedIterator struct {
	Event *RoninGatewayV3RoleRevoked // Event containing the contract specifics and raw log

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
func (it *RoninGatewayV3RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninGatewayV3RoleRevoked)
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
		it.Event = new(RoninGatewayV3RoleRevoked)
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
func (it *RoninGatewayV3RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninGatewayV3RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninGatewayV3RoleRevoked represents a RoleRevoked event raised by the RoninGatewayV3 contract.
type RoninGatewayV3RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RoninGatewayV3 *RoninGatewayV3Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RoninGatewayV3RoleRevokedIterator, error) {

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

	logs, sub, err := _RoninGatewayV3.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RoninGatewayV3RoleRevokedIterator{contract: _RoninGatewayV3.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RoninGatewayV3 *RoninGatewayV3Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RoninGatewayV3RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _RoninGatewayV3.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninGatewayV3RoleRevoked)
				if err := _RoninGatewayV3.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_RoninGatewayV3 *RoninGatewayV3Filterer) ParseRoleRevoked(log types.Log) (*RoninGatewayV3RoleRevoked, error) {
	event := new(RoninGatewayV3RoleRevoked)
	if err := _RoninGatewayV3.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninGatewayV3ThresholdUpdatedIterator is returned from FilterThresholdUpdated and is used to iterate over the raw logs and unpacked data for ThresholdUpdated events raised by the RoninGatewayV3 contract.
type RoninGatewayV3ThresholdUpdatedIterator struct {
	Event *RoninGatewayV3ThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *RoninGatewayV3ThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninGatewayV3ThresholdUpdated)
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
		it.Event = new(RoninGatewayV3ThresholdUpdated)
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
func (it *RoninGatewayV3ThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninGatewayV3ThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninGatewayV3ThresholdUpdated represents a ThresholdUpdated event raised by the RoninGatewayV3 contract.
type RoninGatewayV3ThresholdUpdated struct {
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
func (_RoninGatewayV3 *RoninGatewayV3Filterer) FilterThresholdUpdated(opts *bind.FilterOpts, nonce []*big.Int, numerator []*big.Int, denominator []*big.Int) (*RoninGatewayV3ThresholdUpdatedIterator, error) {

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

	logs, sub, err := _RoninGatewayV3.contract.FilterLogs(opts, "ThresholdUpdated", nonceRule, numeratorRule, denominatorRule)
	if err != nil {
		return nil, err
	}
	return &RoninGatewayV3ThresholdUpdatedIterator{contract: _RoninGatewayV3.contract, event: "ThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchThresholdUpdated is a free log subscription operation binding the contract event 0x976f8a9c5bdf8248dec172376d6e2b80a8e3df2f0328e381c6db8e1cf138c0f8.
//
// Solidity: event ThresholdUpdated(uint256 indexed nonce, uint256 indexed numerator, uint256 indexed denominator, uint256 previousNumerator, uint256 previousDenominator)
func (_RoninGatewayV3 *RoninGatewayV3Filterer) WatchThresholdUpdated(opts *bind.WatchOpts, sink chan<- *RoninGatewayV3ThresholdUpdated, nonce []*big.Int, numerator []*big.Int, denominator []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _RoninGatewayV3.contract.WatchLogs(opts, "ThresholdUpdated", nonceRule, numeratorRule, denominatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninGatewayV3ThresholdUpdated)
				if err := _RoninGatewayV3.contract.UnpackLog(event, "ThresholdUpdated", log); err != nil {
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
func (_RoninGatewayV3 *RoninGatewayV3Filterer) ParseThresholdUpdated(log types.Log) (*RoninGatewayV3ThresholdUpdated, error) {
	event := new(RoninGatewayV3ThresholdUpdated)
	if err := _RoninGatewayV3.contract.UnpackLog(event, "ThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninGatewayV3TokenMappedIterator is returned from FilterTokenMapped and is used to iterate over the raw logs and unpacked data for TokenMapped events raised by the RoninGatewayV3 contract.
type RoninGatewayV3TokenMappedIterator struct {
	Event *RoninGatewayV3TokenMapped // Event containing the contract specifics and raw log

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
func (it *RoninGatewayV3TokenMappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninGatewayV3TokenMapped)
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
		it.Event = new(RoninGatewayV3TokenMapped)
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
func (it *RoninGatewayV3TokenMappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninGatewayV3TokenMappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninGatewayV3TokenMapped represents a TokenMapped event raised by the RoninGatewayV3 contract.
type RoninGatewayV3TokenMapped struct {
	RoninTokens     []common.Address
	MainchainTokens []common.Address
	ChainIds        []*big.Int
	Standards       []uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTokenMapped is a free log retrieval operation binding the contract event 0x2544bff60c6d5b84946e06804af9f84e150bbee85238dbdee79efca4e0adf401.
//
// Solidity: event TokenMapped(address[] roninTokens, address[] mainchainTokens, uint256[] chainIds, uint8[] standards)
func (_RoninGatewayV3 *RoninGatewayV3Filterer) FilterTokenMapped(opts *bind.FilterOpts) (*RoninGatewayV3TokenMappedIterator, error) {

	logs, sub, err := _RoninGatewayV3.contract.FilterLogs(opts, "TokenMapped")
	if err != nil {
		return nil, err
	}
	return &RoninGatewayV3TokenMappedIterator{contract: _RoninGatewayV3.contract, event: "TokenMapped", logs: logs, sub: sub}, nil
}

// WatchTokenMapped is a free log subscription operation binding the contract event 0x2544bff60c6d5b84946e06804af9f84e150bbee85238dbdee79efca4e0adf401.
//
// Solidity: event TokenMapped(address[] roninTokens, address[] mainchainTokens, uint256[] chainIds, uint8[] standards)
func (_RoninGatewayV3 *RoninGatewayV3Filterer) WatchTokenMapped(opts *bind.WatchOpts, sink chan<- *RoninGatewayV3TokenMapped) (event.Subscription, error) {

	logs, sub, err := _RoninGatewayV3.contract.WatchLogs(opts, "TokenMapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninGatewayV3TokenMapped)
				if err := _RoninGatewayV3.contract.UnpackLog(event, "TokenMapped", log); err != nil {
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
func (_RoninGatewayV3 *RoninGatewayV3Filterer) ParseTokenMapped(log types.Log) (*RoninGatewayV3TokenMapped, error) {
	event := new(RoninGatewayV3TokenMapped)
	if err := _RoninGatewayV3.contract.UnpackLog(event, "TokenMapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninGatewayV3UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the RoninGatewayV3 contract.
type RoninGatewayV3UnpausedIterator struct {
	Event *RoninGatewayV3Unpaused // Event containing the contract specifics and raw log

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
func (it *RoninGatewayV3UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninGatewayV3Unpaused)
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
		it.Event = new(RoninGatewayV3Unpaused)
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
func (it *RoninGatewayV3UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninGatewayV3UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninGatewayV3Unpaused represents a Unpaused event raised by the RoninGatewayV3 contract.
type RoninGatewayV3Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_RoninGatewayV3 *RoninGatewayV3Filterer) FilterUnpaused(opts *bind.FilterOpts) (*RoninGatewayV3UnpausedIterator, error) {

	logs, sub, err := _RoninGatewayV3.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &RoninGatewayV3UnpausedIterator{contract: _RoninGatewayV3.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_RoninGatewayV3 *RoninGatewayV3Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *RoninGatewayV3Unpaused) (event.Subscription, error) {

	logs, sub, err := _RoninGatewayV3.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninGatewayV3Unpaused)
				if err := _RoninGatewayV3.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_RoninGatewayV3 *RoninGatewayV3Filterer) ParseUnpaused(log types.Log) (*RoninGatewayV3Unpaused, error) {
	event := new(RoninGatewayV3Unpaused)
	if err := _RoninGatewayV3.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninGatewayV3ValidatorContractUpdatedIterator is returned from FilterValidatorContractUpdated and is used to iterate over the raw logs and unpacked data for ValidatorContractUpdated events raised by the RoninGatewayV3 contract.
type RoninGatewayV3ValidatorContractUpdatedIterator struct {
	Event *RoninGatewayV3ValidatorContractUpdated // Event containing the contract specifics and raw log

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
func (it *RoninGatewayV3ValidatorContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninGatewayV3ValidatorContractUpdated)
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
		it.Event = new(RoninGatewayV3ValidatorContractUpdated)
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
func (it *RoninGatewayV3ValidatorContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninGatewayV3ValidatorContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninGatewayV3ValidatorContractUpdated represents a ValidatorContractUpdated event raised by the RoninGatewayV3 contract.
type RoninGatewayV3ValidatorContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterValidatorContractUpdated is a free log retrieval operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_RoninGatewayV3 *RoninGatewayV3Filterer) FilterValidatorContractUpdated(opts *bind.FilterOpts) (*RoninGatewayV3ValidatorContractUpdatedIterator, error) {

	logs, sub, err := _RoninGatewayV3.contract.FilterLogs(opts, "ValidatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return &RoninGatewayV3ValidatorContractUpdatedIterator{contract: _RoninGatewayV3.contract, event: "ValidatorContractUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorContractUpdated is a free log subscription operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_RoninGatewayV3 *RoninGatewayV3Filterer) WatchValidatorContractUpdated(opts *bind.WatchOpts, sink chan<- *RoninGatewayV3ValidatorContractUpdated) (event.Subscription, error) {

	logs, sub, err := _RoninGatewayV3.contract.WatchLogs(opts, "ValidatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninGatewayV3ValidatorContractUpdated)
				if err := _RoninGatewayV3.contract.UnpackLog(event, "ValidatorContractUpdated", log); err != nil {
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

// ParseValidatorContractUpdated is a log parse operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_RoninGatewayV3 *RoninGatewayV3Filterer) ParseValidatorContractUpdated(log types.Log) (*RoninGatewayV3ValidatorContractUpdated, error) {
	event := new(RoninGatewayV3ValidatorContractUpdated)
	if err := _RoninGatewayV3.contract.UnpackLog(event, "ValidatorContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninGatewayV3WithdrawalRequestedIterator is returned from FilterWithdrawalRequested and is used to iterate over the raw logs and unpacked data for WithdrawalRequested events raised by the RoninGatewayV3 contract.
type RoninGatewayV3WithdrawalRequestedIterator struct {
	Event *RoninGatewayV3WithdrawalRequested // Event containing the contract specifics and raw log

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
func (it *RoninGatewayV3WithdrawalRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninGatewayV3WithdrawalRequested)
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
		it.Event = new(RoninGatewayV3WithdrawalRequested)
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
func (it *RoninGatewayV3WithdrawalRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninGatewayV3WithdrawalRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninGatewayV3WithdrawalRequested represents a WithdrawalRequested event raised by the RoninGatewayV3 contract.
type RoninGatewayV3WithdrawalRequested struct {
	ReceiptHash [32]byte
	Arg1        TransferReceipt
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalRequested is a free log retrieval operation binding the contract event 0xf313c253a5be72c29d0deb2c8768a9543744ac03d6b3cafd50cc976f1c2632fc.
//
// Solidity: event WithdrawalRequested(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) arg1)
func (_RoninGatewayV3 *RoninGatewayV3Filterer) FilterWithdrawalRequested(opts *bind.FilterOpts) (*RoninGatewayV3WithdrawalRequestedIterator, error) {

	logs, sub, err := _RoninGatewayV3.contract.FilterLogs(opts, "WithdrawalRequested")
	if err != nil {
		return nil, err
	}
	return &RoninGatewayV3WithdrawalRequestedIterator{contract: _RoninGatewayV3.contract, event: "WithdrawalRequested", logs: logs, sub: sub}, nil
}

// WatchWithdrawalRequested is a free log subscription operation binding the contract event 0xf313c253a5be72c29d0deb2c8768a9543744ac03d6b3cafd50cc976f1c2632fc.
//
// Solidity: event WithdrawalRequested(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) arg1)
func (_RoninGatewayV3 *RoninGatewayV3Filterer) WatchWithdrawalRequested(opts *bind.WatchOpts, sink chan<- *RoninGatewayV3WithdrawalRequested) (event.Subscription, error) {

	logs, sub, err := _RoninGatewayV3.contract.WatchLogs(opts, "WithdrawalRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninGatewayV3WithdrawalRequested)
				if err := _RoninGatewayV3.contract.UnpackLog(event, "WithdrawalRequested", log); err != nil {
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
func (_RoninGatewayV3 *RoninGatewayV3Filterer) ParseWithdrawalRequested(log types.Log) (*RoninGatewayV3WithdrawalRequested, error) {
	event := new(RoninGatewayV3WithdrawalRequested)
	if err := _RoninGatewayV3.contract.UnpackLog(event, "WithdrawalRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninGatewayV3WithdrawalSignaturesRequestedIterator is returned from FilterWithdrawalSignaturesRequested and is used to iterate over the raw logs and unpacked data for WithdrawalSignaturesRequested events raised by the RoninGatewayV3 contract.
type RoninGatewayV3WithdrawalSignaturesRequestedIterator struct {
	Event *RoninGatewayV3WithdrawalSignaturesRequested // Event containing the contract specifics and raw log

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
func (it *RoninGatewayV3WithdrawalSignaturesRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninGatewayV3WithdrawalSignaturesRequested)
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
		it.Event = new(RoninGatewayV3WithdrawalSignaturesRequested)
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
func (it *RoninGatewayV3WithdrawalSignaturesRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninGatewayV3WithdrawalSignaturesRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninGatewayV3WithdrawalSignaturesRequested represents a WithdrawalSignaturesRequested event raised by the RoninGatewayV3 contract.
type RoninGatewayV3WithdrawalSignaturesRequested struct {
	ReceiptHash [32]byte
	Arg1        TransferReceipt
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalSignaturesRequested is a free log retrieval operation binding the contract event 0x04e8cbd836dea43a2dc7eb19de345cca3a8e6978a2ef5225d924775500f67c7c.
//
// Solidity: event WithdrawalSignaturesRequested(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) arg1)
func (_RoninGatewayV3 *RoninGatewayV3Filterer) FilterWithdrawalSignaturesRequested(opts *bind.FilterOpts) (*RoninGatewayV3WithdrawalSignaturesRequestedIterator, error) {

	logs, sub, err := _RoninGatewayV3.contract.FilterLogs(opts, "WithdrawalSignaturesRequested")
	if err != nil {
		return nil, err
	}
	return &RoninGatewayV3WithdrawalSignaturesRequestedIterator{contract: _RoninGatewayV3.contract, event: "WithdrawalSignaturesRequested", logs: logs, sub: sub}, nil
}

// WatchWithdrawalSignaturesRequested is a free log subscription operation binding the contract event 0x04e8cbd836dea43a2dc7eb19de345cca3a8e6978a2ef5225d924775500f67c7c.
//
// Solidity: event WithdrawalSignaturesRequested(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) arg1)
func (_RoninGatewayV3 *RoninGatewayV3Filterer) WatchWithdrawalSignaturesRequested(opts *bind.WatchOpts, sink chan<- *RoninGatewayV3WithdrawalSignaturesRequested) (event.Subscription, error) {

	logs, sub, err := _RoninGatewayV3.contract.WatchLogs(opts, "WithdrawalSignaturesRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninGatewayV3WithdrawalSignaturesRequested)
				if err := _RoninGatewayV3.contract.UnpackLog(event, "WithdrawalSignaturesRequested", log); err != nil {
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
func (_RoninGatewayV3 *RoninGatewayV3Filterer) ParseWithdrawalSignaturesRequested(log types.Log) (*RoninGatewayV3WithdrawalSignaturesRequested, error) {
	event := new(RoninGatewayV3WithdrawalSignaturesRequested)
	if err := _RoninGatewayV3.contract.UnpackLog(event, "WithdrawalSignaturesRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
