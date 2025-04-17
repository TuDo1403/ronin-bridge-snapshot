// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package transparent_proxy_v2

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

// TransparentProxyV2MetaData contains all meta data concerning the TransparentProxyV2 contract.
var TransparentProxyV2MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"logic\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"functionDelegateCall\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeaconUpgraded\",\"inputs\":[{\"name\":\"beacon\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// TransparentProxyV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use TransparentProxyV2MetaData.ABI instead.
var TransparentProxyV2ABI = TransparentProxyV2MetaData.ABI

// TransparentProxyV2 is an auto generated Go binding around an Ethereum contract.
type TransparentProxyV2 struct {
	TransparentProxyV2Caller     // Read-only binding to the contract
	TransparentProxyV2Transactor // Write-only binding to the contract
	TransparentProxyV2Filterer   // Log filterer for contract events
}

// TransparentProxyV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type TransparentProxyV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransparentProxyV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TransparentProxyV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransparentProxyV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransparentProxyV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransparentProxyV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransparentProxyV2Session struct {
	Contract     *TransparentProxyV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TransparentProxyV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransparentProxyV2CallerSession struct {
	Contract *TransparentProxyV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// TransparentProxyV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransparentProxyV2TransactorSession struct {
	Contract     *TransparentProxyV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// TransparentProxyV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type TransparentProxyV2Raw struct {
	Contract *TransparentProxyV2 // Generic contract binding to access the raw methods on
}

// TransparentProxyV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransparentProxyV2CallerRaw struct {
	Contract *TransparentProxyV2Caller // Generic read-only contract binding to access the raw methods on
}

// TransparentProxyV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransparentProxyV2TransactorRaw struct {
	Contract *TransparentProxyV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTransparentProxyV2 creates a new instance of TransparentProxyV2, bound to a specific deployed contract.
func NewTransparentProxyV2(address common.Address, backend bind.ContractBackend) (*TransparentProxyV2, error) {
	contract, err := bindTransparentProxyV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TransparentProxyV2{TransparentProxyV2Caller: TransparentProxyV2Caller{contract: contract}, TransparentProxyV2Transactor: TransparentProxyV2Transactor{contract: contract}, TransparentProxyV2Filterer: TransparentProxyV2Filterer{contract: contract}}, nil
}

// NewTransparentProxyV2Caller creates a new read-only instance of TransparentProxyV2, bound to a specific deployed contract.
func NewTransparentProxyV2Caller(address common.Address, caller bind.ContractCaller) (*TransparentProxyV2Caller, error) {
	contract, err := bindTransparentProxyV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransparentProxyV2Caller{contract: contract}, nil
}

// NewTransparentProxyV2Transactor creates a new write-only instance of TransparentProxyV2, bound to a specific deployed contract.
func NewTransparentProxyV2Transactor(address common.Address, transactor bind.ContractTransactor) (*TransparentProxyV2Transactor, error) {
	contract, err := bindTransparentProxyV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransparentProxyV2Transactor{contract: contract}, nil
}

// NewTransparentProxyV2Filterer creates a new log filterer instance of TransparentProxyV2, bound to a specific deployed contract.
func NewTransparentProxyV2Filterer(address common.Address, filterer bind.ContractFilterer) (*TransparentProxyV2Filterer, error) {
	contract, err := bindTransparentProxyV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransparentProxyV2Filterer{contract: contract}, nil
}

// bindTransparentProxyV2 binds a generic wrapper to an already deployed contract.
func bindTransparentProxyV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TransparentProxyV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransparentProxyV2 *TransparentProxyV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransparentProxyV2.Contract.TransparentProxyV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransparentProxyV2 *TransparentProxyV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransparentProxyV2.Contract.TransparentProxyV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransparentProxyV2 *TransparentProxyV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransparentProxyV2.Contract.TransparentProxyV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransparentProxyV2 *TransparentProxyV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransparentProxyV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransparentProxyV2 *TransparentProxyV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransparentProxyV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransparentProxyV2 *TransparentProxyV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransparentProxyV2.Contract.contract.Transact(opts, method, params...)
}

// FunctionDelegateCall is a paid mutator transaction binding the contract method 0x4bb5274a.
//
// Solidity: function functionDelegateCall(bytes data) payable returns()
func (_TransparentProxyV2 *TransparentProxyV2Transactor) FunctionDelegateCall(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _TransparentProxyV2.contract.Transact(opts, "functionDelegateCall", data)
}

// FunctionDelegateCall is a paid mutator transaction binding the contract method 0x4bb5274a.
//
// Solidity: function functionDelegateCall(bytes data) payable returns()
func (_TransparentProxyV2 *TransparentProxyV2Session) FunctionDelegateCall(data []byte) (*types.Transaction, error) {
	return _TransparentProxyV2.Contract.FunctionDelegateCall(&_TransparentProxyV2.TransactOpts, data)
}

// FunctionDelegateCall is a paid mutator transaction binding the contract method 0x4bb5274a.
//
// Solidity: function functionDelegateCall(bytes data) payable returns()
func (_TransparentProxyV2 *TransparentProxyV2TransactorSession) FunctionDelegateCall(data []byte) (*types.Transaction, error) {
	return _TransparentProxyV2.Contract.FunctionDelegateCall(&_TransparentProxyV2.TransactOpts, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TransparentProxyV2 *TransparentProxyV2Transactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _TransparentProxyV2.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TransparentProxyV2 *TransparentProxyV2Session) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TransparentProxyV2.Contract.Fallback(&_TransparentProxyV2.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TransparentProxyV2 *TransparentProxyV2TransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TransparentProxyV2.Contract.Fallback(&_TransparentProxyV2.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TransparentProxyV2 *TransparentProxyV2Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransparentProxyV2.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TransparentProxyV2 *TransparentProxyV2Session) Receive() (*types.Transaction, error) {
	return _TransparentProxyV2.Contract.Receive(&_TransparentProxyV2.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TransparentProxyV2 *TransparentProxyV2TransactorSession) Receive() (*types.Transaction, error) {
	return _TransparentProxyV2.Contract.Receive(&_TransparentProxyV2.TransactOpts)
}

// TransparentProxyV2AdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the TransparentProxyV2 contract.
type TransparentProxyV2AdminChangedIterator struct {
	Event *TransparentProxyV2AdminChanged // Event containing the contract specifics and raw log

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
func (it *TransparentProxyV2AdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransparentProxyV2AdminChanged)
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
		it.Event = new(TransparentProxyV2AdminChanged)
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
func (it *TransparentProxyV2AdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransparentProxyV2AdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransparentProxyV2AdminChanged represents a AdminChanged event raised by the TransparentProxyV2 contract.
type TransparentProxyV2AdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_TransparentProxyV2 *TransparentProxyV2Filterer) FilterAdminChanged(opts *bind.FilterOpts) (*TransparentProxyV2AdminChangedIterator, error) {

	logs, sub, err := _TransparentProxyV2.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &TransparentProxyV2AdminChangedIterator{contract: _TransparentProxyV2.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_TransparentProxyV2 *TransparentProxyV2Filterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *TransparentProxyV2AdminChanged) (event.Subscription, error) {

	logs, sub, err := _TransparentProxyV2.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransparentProxyV2AdminChanged)
				if err := _TransparentProxyV2.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_TransparentProxyV2 *TransparentProxyV2Filterer) ParseAdminChanged(log types.Log) (*TransparentProxyV2AdminChanged, error) {
	event := new(TransparentProxyV2AdminChanged)
	if err := _TransparentProxyV2.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransparentProxyV2BeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the TransparentProxyV2 contract.
type TransparentProxyV2BeaconUpgradedIterator struct {
	Event *TransparentProxyV2BeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *TransparentProxyV2BeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransparentProxyV2BeaconUpgraded)
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
		it.Event = new(TransparentProxyV2BeaconUpgraded)
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
func (it *TransparentProxyV2BeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransparentProxyV2BeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransparentProxyV2BeaconUpgraded represents a BeaconUpgraded event raised by the TransparentProxyV2 contract.
type TransparentProxyV2BeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_TransparentProxyV2 *TransparentProxyV2Filterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*TransparentProxyV2BeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _TransparentProxyV2.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &TransparentProxyV2BeaconUpgradedIterator{contract: _TransparentProxyV2.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_TransparentProxyV2 *TransparentProxyV2Filterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *TransparentProxyV2BeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _TransparentProxyV2.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransparentProxyV2BeaconUpgraded)
				if err := _TransparentProxyV2.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_TransparentProxyV2 *TransparentProxyV2Filterer) ParseBeaconUpgraded(log types.Log) (*TransparentProxyV2BeaconUpgraded, error) {
	event := new(TransparentProxyV2BeaconUpgraded)
	if err := _TransparentProxyV2.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransparentProxyV2UpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TransparentProxyV2 contract.
type TransparentProxyV2UpgradedIterator struct {
	Event *TransparentProxyV2Upgraded // Event containing the contract specifics and raw log

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
func (it *TransparentProxyV2UpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransparentProxyV2Upgraded)
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
		it.Event = new(TransparentProxyV2Upgraded)
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
func (it *TransparentProxyV2UpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransparentProxyV2UpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransparentProxyV2Upgraded represents a Upgraded event raised by the TransparentProxyV2 contract.
type TransparentProxyV2Upgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TransparentProxyV2 *TransparentProxyV2Filterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TransparentProxyV2UpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TransparentProxyV2.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TransparentProxyV2UpgradedIterator{contract: _TransparentProxyV2.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TransparentProxyV2 *TransparentProxyV2Filterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TransparentProxyV2Upgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TransparentProxyV2.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransparentProxyV2Upgraded)
				if err := _TransparentProxyV2.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TransparentProxyV2 *TransparentProxyV2Filterer) ParseUpgraded(log types.Log) (*TransparentProxyV2Upgraded, error) {
	event := new(TransparentProxyV2Upgraded)
	if err := _TransparentProxyV2.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
