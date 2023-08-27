// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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
)

// BotMetaData contains all meta data concerning the Bot contract.
var BotMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_WETH\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BaseTokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"BaseTokenRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"addBaseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"baseTokensContains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool1\",\"type\":\"address\"}],\"name\":\"flashArbitrage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool1\",\"type\":\"address\"}],\"name\":\"getProfit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeBaseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV2Call\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// BotABI is the input ABI used to generate the binding from.
// Deprecated: Use BotMetaData.ABI instead.
var BotABI = BotMetaData.ABI

// Bot is an auto generated Go binding around an Ethereum contract.
type Bot struct {
	BotCaller     // Read-only binding to the contract
	BotTransactor // Write-only binding to the contract
	BotFilterer   // Log filterer for contract events
}

// BotCaller is an auto generated read-only Go binding around an Ethereum contract.
type BotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BotSession struct {
	Contract     *Bot              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BotCallerSession struct {
	Contract *BotCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BotTransactorSession struct {
	Contract     *BotTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BotRaw is an auto generated low-level Go binding around an Ethereum contract.
type BotRaw struct {
	Contract *Bot // Generic contract binding to access the raw methods on
}

// BotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BotCallerRaw struct {
	Contract *BotCaller // Generic read-only contract binding to access the raw methods on
}

// BotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BotTransactorRaw struct {
	Contract *BotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBot creates a new instance of Bot, bound to a specific deployed contract.
func NewBot(address common.Address, backend bind.ContractBackend) (*Bot, error) {
	contract, err := bindBot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bot{BotCaller: BotCaller{contract: contract}, BotTransactor: BotTransactor{contract: contract}, BotFilterer: BotFilterer{contract: contract}}, nil
}

// NewBotCaller creates a new read-only instance of Bot, bound to a specific deployed contract.
func NewBotCaller(address common.Address, caller bind.ContractCaller) (*BotCaller, error) {
	contract, err := bindBot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BotCaller{contract: contract}, nil
}

// NewBotTransactor creates a new write-only instance of Bot, bound to a specific deployed contract.
func NewBotTransactor(address common.Address, transactor bind.ContractTransactor) (*BotTransactor, error) {
	contract, err := bindBot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BotTransactor{contract: contract}, nil
}

// NewBotFilterer creates a new log filterer instance of Bot, bound to a specific deployed contract.
func NewBotFilterer(address common.Address, filterer bind.ContractFilterer) (*BotFilterer, error) {
	contract, err := bindBot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BotFilterer{contract: contract}, nil
}

// bindBot binds a generic wrapper to an already deployed contract.
func bindBot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BotABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bot *BotRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bot.Contract.BotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bot *BotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bot.Contract.BotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bot *BotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bot.Contract.BotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bot *BotCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bot *BotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bot *BotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bot.Contract.contract.Transact(opts, method, params...)
}

// BaseTokensContains is a free data retrieval call binding the contract method 0x21d09426.
//
// Solidity: function baseTokensContains(address token) view returns(bool)
func (_Bot *BotCaller) BaseTokensContains(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _Bot.contract.Call(opts, &out, "baseTokensContains", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BaseTokensContains is a free data retrieval call binding the contract method 0x21d09426.
//
// Solidity: function baseTokensContains(address token) view returns(bool)
func (_Bot *BotSession) BaseTokensContains(token common.Address) (bool, error) {
	return _Bot.Contract.BaseTokensContains(&_Bot.CallOpts, token)
}

// BaseTokensContains is a free data retrieval call binding the contract method 0x21d09426.
//
// Solidity: function baseTokensContains(address token) view returns(bool)
func (_Bot *BotCallerSession) BaseTokensContains(token common.Address) (bool, error) {
	return _Bot.Contract.BaseTokensContains(&_Bot.CallOpts, token)
}

// GetBaseTokens is a free data retrieval call binding the contract method 0xbed64c2f.
//
// Solidity: function getBaseTokens() view returns(address[] tokens)
func (_Bot *BotCaller) GetBaseTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Bot.contract.Call(opts, &out, "getBaseTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBaseTokens is a free data retrieval call binding the contract method 0xbed64c2f.
//
// Solidity: function getBaseTokens() view returns(address[] tokens)
func (_Bot *BotSession) GetBaseTokens() ([]common.Address, error) {
	return _Bot.Contract.GetBaseTokens(&_Bot.CallOpts)
}

// GetBaseTokens is a free data retrieval call binding the contract method 0xbed64c2f.
//
// Solidity: function getBaseTokens() view returns(address[] tokens)
func (_Bot *BotCallerSession) GetBaseTokens() ([]common.Address, error) {
	return _Bot.Contract.GetBaseTokens(&_Bot.CallOpts)
}

// GetProfit is a free data retrieval call binding the contract method 0x759eee10.
//
// Solidity: function getProfit(address pool0, address pool1) view returns(uint256 profit, address baseToken)
func (_Bot *BotCaller) GetProfit(opts *bind.CallOpts, pool0 common.Address, pool1 common.Address) (struct {
	Profit    *big.Int
	BaseToken common.Address
}, error) {
	var out []interface{}
	err := _Bot.contract.Call(opts, &out, "getProfit", pool0, pool1)

	outstruct := new(struct {
		Profit    *big.Int
		BaseToken common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Profit = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BaseToken = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetProfit is a free data retrieval call binding the contract method 0x759eee10.
//
// Solidity: function getProfit(address pool0, address pool1) view returns(uint256 profit, address baseToken)
func (_Bot *BotSession) GetProfit(pool0 common.Address, pool1 common.Address) (struct {
	Profit    *big.Int
	BaseToken common.Address
}, error) {
	return _Bot.Contract.GetProfit(&_Bot.CallOpts, pool0, pool1)
}

// GetProfit is a free data retrieval call binding the contract method 0x759eee10.
//
// Solidity: function getProfit(address pool0, address pool1) view returns(uint256 profit, address baseToken)
func (_Bot *BotCallerSession) GetProfit(pool0 common.Address, pool1 common.Address) (struct {
	Profit    *big.Int
	BaseToken common.Address
}, error) {
	return _Bot.Contract.GetProfit(&_Bot.CallOpts, pool0, pool1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bot *BotCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bot.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bot *BotSession) Owner() (common.Address, error) {
	return _Bot.Contract.Owner(&_Bot.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bot *BotCallerSession) Owner() (common.Address, error) {
	return _Bot.Contract.Owner(&_Bot.CallOpts)
}

// AddBaseToken is a paid mutator transaction binding the contract method 0x83e280d9.
//
// Solidity: function addBaseToken(address token) returns()
func (_Bot *BotTransactor) AddBaseToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Bot.contract.Transact(opts, "addBaseToken", token)
}

// AddBaseToken is a paid mutator transaction binding the contract method 0x83e280d9.
//
// Solidity: function addBaseToken(address token) returns()
func (_Bot *BotSession) AddBaseToken(token common.Address) (*types.Transaction, error) {
	return _Bot.Contract.AddBaseToken(&_Bot.TransactOpts, token)
}

// AddBaseToken is a paid mutator transaction binding the contract method 0x83e280d9.
//
// Solidity: function addBaseToken(address token) returns()
func (_Bot *BotTransactorSession) AddBaseToken(token common.Address) (*types.Transaction, error) {
	return _Bot.Contract.AddBaseToken(&_Bot.TransactOpts, token)
}

// FlashArbitrage is a paid mutator transaction binding the contract method 0xbaee64f1.
//
// Solidity: function flashArbitrage(address pool0, address pool1) returns()
func (_Bot *BotTransactor) FlashArbitrage(opts *bind.TransactOpts, pool0 common.Address, pool1 common.Address) (*types.Transaction, error) {
	return _Bot.contract.Transact(opts, "flashArbitrage", pool0, pool1)
}

// FlashArbitrage is a paid mutator transaction binding the contract method 0xbaee64f1.
//
// Solidity: function flashArbitrage(address pool0, address pool1) returns()
func (_Bot *BotSession) FlashArbitrage(pool0 common.Address, pool1 common.Address) (*types.Transaction, error) {
	return _Bot.Contract.FlashArbitrage(&_Bot.TransactOpts, pool0, pool1)
}

// FlashArbitrage is a paid mutator transaction binding the contract method 0xbaee64f1.
//
// Solidity: function flashArbitrage(address pool0, address pool1) returns()
func (_Bot *BotTransactorSession) FlashArbitrage(pool0 common.Address, pool1 common.Address) (*types.Transaction, error) {
	return _Bot.Contract.FlashArbitrage(&_Bot.TransactOpts, pool0, pool1)
}

// RemoveBaseToken is a paid mutator transaction binding the contract method 0xbbd1e122.
//
// Solidity: function removeBaseToken(address token) returns()
func (_Bot *BotTransactor) RemoveBaseToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Bot.contract.Transact(opts, "removeBaseToken", token)
}

// RemoveBaseToken is a paid mutator transaction binding the contract method 0xbbd1e122.
//
// Solidity: function removeBaseToken(address token) returns()
func (_Bot *BotSession) RemoveBaseToken(token common.Address) (*types.Transaction, error) {
	return _Bot.Contract.RemoveBaseToken(&_Bot.TransactOpts, token)
}

// RemoveBaseToken is a paid mutator transaction binding the contract method 0xbbd1e122.
//
// Solidity: function removeBaseToken(address token) returns()
func (_Bot *BotTransactorSession) RemoveBaseToken(token common.Address) (*types.Transaction, error) {
	return _Bot.Contract.RemoveBaseToken(&_Bot.TransactOpts, token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bot *BotTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bot.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bot *BotSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bot.Contract.RenounceOwnership(&_Bot.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bot *BotTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bot.Contract.RenounceOwnership(&_Bot.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bot *BotTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bot.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bot *BotSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bot.Contract.TransferOwnership(&_Bot.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bot *BotTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bot.Contract.TransferOwnership(&_Bot.TransactOpts, newOwner)
}

// UniswapV2Call is a paid mutator transaction binding the contract method 0x10d1e85c.
//
// Solidity: function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_Bot *BotTransactor) UniswapV2Call(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Bot.contract.Transact(opts, "uniswapV2Call", sender, amount0, amount1, data)
}

// UniswapV2Call is a paid mutator transaction binding the contract method 0x10d1e85c.
//
// Solidity: function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_Bot *BotSession) UniswapV2Call(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Bot.Contract.UniswapV2Call(&_Bot.TransactOpts, sender, amount0, amount1, data)
}

// UniswapV2Call is a paid mutator transaction binding the contract method 0x10d1e85c.
//
// Solidity: function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_Bot *BotTransactorSession) UniswapV2Call(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Bot.Contract.UniswapV2Call(&_Bot.TransactOpts, sender, amount0, amount1, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Bot *BotTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bot.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Bot *BotSession) Withdraw() (*types.Transaction, error) {
	return _Bot.Contract.Withdraw(&_Bot.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Bot *BotTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Bot.Contract.Withdraw(&_Bot.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Bot *BotTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Bot.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Bot *BotSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Bot.Contract.Fallback(&_Bot.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Bot *BotTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Bot.Contract.Fallback(&_Bot.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bot *BotTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bot.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bot *BotSession) Receive() (*types.Transaction, error) {
	return _Bot.Contract.Receive(&_Bot.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bot *BotTransactorSession) Receive() (*types.Transaction, error) {
	return _Bot.Contract.Receive(&_Bot.TransactOpts)
}

// BotBaseTokenAddedIterator is returned from FilterBaseTokenAdded and is used to iterate over the raw logs and unpacked data for BaseTokenAdded events raised by the Bot contract.
type BotBaseTokenAddedIterator struct {
	Event *BotBaseTokenAdded // Event containing the contract specifics and raw log

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
func (it *BotBaseTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BotBaseTokenAdded)
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
		it.Event = new(BotBaseTokenAdded)
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
func (it *BotBaseTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BotBaseTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BotBaseTokenAdded represents a BaseTokenAdded event raised by the Bot contract.
type BotBaseTokenAdded struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBaseTokenAdded is a free log retrieval operation binding the contract event 0xfa1388d6e7328e9c711a539b0addfc27de8bfb6f5924cce26f80f41023b15253.
//
// Solidity: event BaseTokenAdded(address indexed token)
func (_Bot *BotFilterer) FilterBaseTokenAdded(opts *bind.FilterOpts, token []common.Address) (*BotBaseTokenAddedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Bot.contract.FilterLogs(opts, "BaseTokenAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BotBaseTokenAddedIterator{contract: _Bot.contract, event: "BaseTokenAdded", logs: logs, sub: sub}, nil
}

// WatchBaseTokenAdded is a free log subscription operation binding the contract event 0xfa1388d6e7328e9c711a539b0addfc27de8bfb6f5924cce26f80f41023b15253.
//
// Solidity: event BaseTokenAdded(address indexed token)
func (_Bot *BotFilterer) WatchBaseTokenAdded(opts *bind.WatchOpts, sink chan<- *BotBaseTokenAdded, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Bot.contract.WatchLogs(opts, "BaseTokenAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BotBaseTokenAdded)
				if err := _Bot.contract.UnpackLog(event, "BaseTokenAdded", log); err != nil {
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

// ParseBaseTokenAdded is a log parse operation binding the contract event 0xfa1388d6e7328e9c711a539b0addfc27de8bfb6f5924cce26f80f41023b15253.
//
// Solidity: event BaseTokenAdded(address indexed token)
func (_Bot *BotFilterer) ParseBaseTokenAdded(log types.Log) (*BotBaseTokenAdded, error) {
	event := new(BotBaseTokenAdded)
	if err := _Bot.contract.UnpackLog(event, "BaseTokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BotBaseTokenRemovedIterator is returned from FilterBaseTokenRemoved and is used to iterate over the raw logs and unpacked data for BaseTokenRemoved events raised by the Bot contract.
type BotBaseTokenRemovedIterator struct {
	Event *BotBaseTokenRemoved // Event containing the contract specifics and raw log

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
func (it *BotBaseTokenRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BotBaseTokenRemoved)
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
		it.Event = new(BotBaseTokenRemoved)
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
func (it *BotBaseTokenRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BotBaseTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BotBaseTokenRemoved represents a BaseTokenRemoved event raised by the Bot contract.
type BotBaseTokenRemoved struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBaseTokenRemoved is a free log retrieval operation binding the contract event 0xdc23a849435922f20a9732eb85192a9d0c1cb34725ebe6d7de0be10212ba02fb.
//
// Solidity: event BaseTokenRemoved(address indexed token)
func (_Bot *BotFilterer) FilterBaseTokenRemoved(opts *bind.FilterOpts, token []common.Address) (*BotBaseTokenRemovedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Bot.contract.FilterLogs(opts, "BaseTokenRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BotBaseTokenRemovedIterator{contract: _Bot.contract, event: "BaseTokenRemoved", logs: logs, sub: sub}, nil
}

// WatchBaseTokenRemoved is a free log subscription operation binding the contract event 0xdc23a849435922f20a9732eb85192a9d0c1cb34725ebe6d7de0be10212ba02fb.
//
// Solidity: event BaseTokenRemoved(address indexed token)
func (_Bot *BotFilterer) WatchBaseTokenRemoved(opts *bind.WatchOpts, sink chan<- *BotBaseTokenRemoved, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Bot.contract.WatchLogs(opts, "BaseTokenRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BotBaseTokenRemoved)
				if err := _Bot.contract.UnpackLog(event, "BaseTokenRemoved", log); err != nil {
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

// ParseBaseTokenRemoved is a log parse operation binding the contract event 0xdc23a849435922f20a9732eb85192a9d0c1cb34725ebe6d7de0be10212ba02fb.
//
// Solidity: event BaseTokenRemoved(address indexed token)
func (_Bot *BotFilterer) ParseBaseTokenRemoved(log types.Log) (*BotBaseTokenRemoved, error) {
	event := new(BotBaseTokenRemoved)
	if err := _Bot.contract.UnpackLog(event, "BaseTokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BotOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bot contract.
type BotOwnershipTransferredIterator struct {
	Event *BotOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BotOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BotOwnershipTransferred)
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
		it.Event = new(BotOwnershipTransferred)
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
func (it *BotOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BotOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BotOwnershipTransferred represents a OwnershipTransferred event raised by the Bot contract.
type BotOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bot *BotFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BotOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bot.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BotOwnershipTransferredIterator{contract: _Bot.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bot *BotFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BotOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bot.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BotOwnershipTransferred)
				if err := _Bot.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Bot *BotFilterer) ParseOwnershipTransferred(log types.Log) (*BotOwnershipTransferred, error) {
	event := new(BotOwnershipTransferred)
	if err := _Bot.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BotWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Bot contract.
type BotWithdrawnIterator struct {
	Event *BotWithdrawn // Event containing the contract specifics and raw log

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
func (it *BotWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BotWithdrawn)
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
		it.Event = new(BotWithdrawn)
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
func (it *BotWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BotWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BotWithdrawn represents a Withdrawn event raised by the Bot contract.
type BotWithdrawn struct {
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed to, uint256 indexed value)
func (_Bot *BotFilterer) FilterWithdrawn(opts *bind.FilterOpts, to []common.Address, value []*big.Int) (*BotWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Bot.contract.FilterLogs(opts, "Withdrawn", toRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &BotWithdrawnIterator{contract: _Bot.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed to, uint256 indexed value)
func (_Bot *BotFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *BotWithdrawn, to []common.Address, value []*big.Int) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Bot.contract.WatchLogs(opts, "Withdrawn", toRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BotWithdrawn)
				if err := _Bot.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed to, uint256 indexed value)
func (_Bot *BotFilterer) ParseWithdrawn(log types.Log) (*BotWithdrawn, error) {
	event := new(BotWithdrawn)
	if err := _Bot.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
