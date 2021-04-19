// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// IStateManagerABI is the input ABI used to generate the binding from.
const IStateManagerABI = "[{\"inputs\":[],\"name\":\"getEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"setEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"}],\"name\":\"setState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IStateManager is an auto generated Go binding around an Ethereum contract.
type IStateManager struct {
	IStateManagerCaller     // Read-only binding to the contract
	IStateManagerTransactor // Write-only binding to the contract
	IStateManagerFilterer   // Log filterer for contract events
}

// IStateManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStateManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStateManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStateManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStateManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStateManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStateManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStateManagerSession struct {
	Contract     *IStateManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IStateManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStateManagerCallerSession struct {
	Contract *IStateManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IStateManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStateManagerTransactorSession struct {
	Contract     *IStateManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IStateManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStateManagerRaw struct {
	Contract *IStateManager // Generic contract binding to access the raw methods on
}

// IStateManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStateManagerCallerRaw struct {
	Contract *IStateManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IStateManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStateManagerTransactorRaw struct {
	Contract *IStateManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStateManager creates a new instance of IStateManager, bound to a specific deployed contract.
func NewIStateManager(address common.Address, backend bind.ContractBackend) (*IStateManager, error) {
	contract, err := bindIStateManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStateManager{IStateManagerCaller: IStateManagerCaller{contract: contract}, IStateManagerTransactor: IStateManagerTransactor{contract: contract}, IStateManagerFilterer: IStateManagerFilterer{contract: contract}}, nil
}

// NewIStateManagerCaller creates a new read-only instance of IStateManager, bound to a specific deployed contract.
func NewIStateManagerCaller(address common.Address, caller bind.ContractCaller) (*IStateManagerCaller, error) {
	contract, err := bindIStateManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStateManagerCaller{contract: contract}, nil
}

// NewIStateManagerTransactor creates a new write-only instance of IStateManager, bound to a specific deployed contract.
func NewIStateManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IStateManagerTransactor, error) {
	contract, err := bindIStateManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStateManagerTransactor{contract: contract}, nil
}

// NewIStateManagerFilterer creates a new log filterer instance of IStateManager, bound to a specific deployed contract.
func NewIStateManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IStateManagerFilterer, error) {
	contract, err := bindIStateManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStateManagerFilterer{contract: contract}, nil
}

// bindIStateManager binds a generic wrapper to an already deployed contract.
func bindIStateManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IStateManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStateManager *IStateManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStateManager.Contract.IStateManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStateManager *IStateManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStateManager.Contract.IStateManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStateManager *IStateManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStateManager.Contract.IStateManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStateManager *IStateManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStateManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStateManager *IStateManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStateManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStateManager *IStateManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStateManager.Contract.contract.Transact(opts, method, params...)
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() view returns(uint256)
func (_IStateManager *IStateManagerCaller) GetEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IStateManager.contract.Call(opts, &out, "getEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() view returns(uint256)
func (_IStateManager *IStateManagerSession) GetEpoch() (*big.Int, error) {
	return _IStateManager.Contract.GetEpoch(&_IStateManager.CallOpts)
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() view returns(uint256)
func (_IStateManager *IStateManagerCallerSession) GetEpoch() (*big.Int, error) {
	return _IStateManager.Contract.GetEpoch(&_IStateManager.CallOpts)
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() view returns(uint256)
func (_IStateManager *IStateManagerCaller) GetState(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IStateManager.contract.Call(opts, &out, "getState")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() view returns(uint256)
func (_IStateManager *IStateManagerSession) GetState() (*big.Int, error) {
	return _IStateManager.Contract.GetState(&_IStateManager.CallOpts)
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() view returns(uint256)
func (_IStateManager *IStateManagerCallerSession) GetState() (*big.Int, error) {
	return _IStateManager.Contract.GetState(&_IStateManager.CallOpts)
}

// SetEpoch is a paid mutator transaction binding the contract method 0x0ceb2cef.
//
// Solidity: function setEpoch(uint256 epoch) returns()
func (_IStateManager *IStateManagerTransactor) SetEpoch(opts *bind.TransactOpts, epoch *big.Int) (*types.Transaction, error) {
	return _IStateManager.contract.Transact(opts, "setEpoch", epoch)
}

// SetEpoch is a paid mutator transaction binding the contract method 0x0ceb2cef.
//
// Solidity: function setEpoch(uint256 epoch) returns()
func (_IStateManager *IStateManagerSession) SetEpoch(epoch *big.Int) (*types.Transaction, error) {
	return _IStateManager.Contract.SetEpoch(&_IStateManager.TransactOpts, epoch)
}

// SetEpoch is a paid mutator transaction binding the contract method 0x0ceb2cef.
//
// Solidity: function setEpoch(uint256 epoch) returns()
func (_IStateManager *IStateManagerTransactorSession) SetEpoch(epoch *big.Int) (*types.Transaction, error) {
	return _IStateManager.Contract.SetEpoch(&_IStateManager.TransactOpts, epoch)
}

// SetState is a paid mutator transaction binding the contract method 0xa9e966b7.
//
// Solidity: function setState(uint256 state) returns()
func (_IStateManager *IStateManagerTransactor) SetState(opts *bind.TransactOpts, state *big.Int) (*types.Transaction, error) {
	return _IStateManager.contract.Transact(opts, "setState", state)
}

// SetState is a paid mutator transaction binding the contract method 0xa9e966b7.
//
// Solidity: function setState(uint256 state) returns()
func (_IStateManager *IStateManagerSession) SetState(state *big.Int) (*types.Transaction, error) {
	return _IStateManager.Contract.SetState(&_IStateManager.TransactOpts, state)
}

// SetState is a paid mutator transaction binding the contract method 0xa9e966b7.
//
// Solidity: function setState(uint256 state) returns()
func (_IStateManager *IStateManagerTransactorSession) SetState(state *big.Int) (*types.Transaction, error) {
	return _IStateManager.Contract.SetState(&_IStateManager.TransactOpts, state)
}
