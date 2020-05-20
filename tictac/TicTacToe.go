// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tictac

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TictacABI is the input ABI used to generate the binding from.
const TictacABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ppw\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"circle\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"firstMove\",\"type\":\"uint256\"}],\"name\":\"startGame\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// Tictac is an auto generated Go binding around an Ethereum contract.
type Tictac struct {
	TictacCaller     // Read-only binding to the contract
	TictacTransactor // Write-only binding to the contract
	TictacFilterer   // Log filterer for contract events
}

// TictacCaller is an auto generated read-only Go binding around an Ethereum contract.
type TictacCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TictacTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TictacTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TictacFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TictacFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TictacSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TictacSession struct {
	Contract     *Tictac           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TictacCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TictacCallerSession struct {
	Contract *TictacCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TictacTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TictacTransactorSession struct {
	Contract     *TictacTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TictacRaw is an auto generated low-level Go binding around an Ethereum contract.
type TictacRaw struct {
	Contract *Tictac // Generic contract binding to access the raw methods on
}

// TictacCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TictacCallerRaw struct {
	Contract *TictacCaller // Generic read-only contract binding to access the raw methods on
}

// TictacTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TictacTransactorRaw struct {
	Contract *TictacTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTictac creates a new instance of Tictac, bound to a specific deployed contract.
func NewTictac(address common.Address, backend bind.ContractBackend) (*Tictac, error) {
	contract, err := bindTictac(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tictac{TictacCaller: TictacCaller{contract: contract}, TictacTransactor: TictacTransactor{contract: contract}, TictacFilterer: TictacFilterer{contract: contract}}, nil
}

// NewTictacCaller creates a new read-only instance of Tictac, bound to a specific deployed contract.
func NewTictacCaller(address common.Address, caller bind.ContractCaller) (*TictacCaller, error) {
	contract, err := bindTictac(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TictacCaller{contract: contract}, nil
}

// NewTictacTransactor creates a new write-only instance of Tictac, bound to a specific deployed contract.
func NewTictacTransactor(address common.Address, transactor bind.ContractTransactor) (*TictacTransactor, error) {
	contract, err := bindTictac(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TictacTransactor{contract: contract}, nil
}

// NewTictacFilterer creates a new log filterer instance of Tictac, bound to a specific deployed contract.
func NewTictacFilterer(address common.Address, filterer bind.ContractFilterer) (*TictacFilterer, error) {
	contract, err := bindTictac(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TictacFilterer{contract: contract}, nil
}

// bindTictac binds a generic wrapper to an already deployed contract.
func bindTictac(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TictacABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tictac *TictacRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Tictac.Contract.TictacCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tictac *TictacRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tictac.Contract.TictacTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tictac *TictacRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tictac.Contract.TictacTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tictac *TictacCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Tictac.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tictac *TictacTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tictac.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tictac *TictacTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tictac.Contract.contract.Transact(opts, method, params...)
}

// StartGame is a paid mutator transaction binding the contract method 0xcd45bf1d.
//
// Solidity: function startGame(string name, bool circle, uint256 firstMove) payable returns()
func (_Tictac *TictacTransactor) StartGame(opts *bind.TransactOpts, name string, circle bool, firstMove *big.Int) (*types.Transaction, error) {
	return _Tictac.contract.Transact(opts, "startGame", name, circle, firstMove)
}

// StartGame is a paid mutator transaction binding the contract method 0xcd45bf1d.
//
// Solidity: function startGame(string name, bool circle, uint256 firstMove) payable returns()
func (_Tictac *TictacSession) StartGame(name string, circle bool, firstMove *big.Int) (*types.Transaction, error) {
	return _Tictac.Contract.StartGame(&_Tictac.TransactOpts, name, circle, firstMove)
}

// StartGame is a paid mutator transaction binding the contract method 0xcd45bf1d.
//
// Solidity: function startGame(string name, bool circle, uint256 firstMove) payable returns()
func (_Tictac *TictacTransactorSession) StartGame(name string, circle bool, firstMove *big.Int) (*types.Transaction, error) {
	return _Tictac.Contract.StartGame(&_Tictac.TransactOpts, name, circle, firstMove)
}
