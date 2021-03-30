// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package StateManager

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

// StateManagerABI is the input ABI used to generate the binding from.
const StateManagerABI = "[{\"inputs\":[],\"name\":\"getEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// StateManager is an auto generated Go binding around an Ethereum contract.
type StateManager struct {
	StateManagerCaller     // Read-only binding to the contract
	StateManagerTransactor // Write-only binding to the contract
	StateManagerFilterer   // Log filterer for contract events
}

// StateManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StateManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StateManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StateManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StateManagerSession struct {
	Contract     *StateManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StateManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StateManagerCallerSession struct {
	Contract *StateManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StateManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StateManagerTransactorSession struct {
	Contract     *StateManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StateManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StateManagerRaw struct {
	Contract *StateManager // Generic contract binding to access the raw methods on
}

// StateManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StateManagerCallerRaw struct {
	Contract *StateManagerCaller // Generic read-only contract binding to access the raw methods on
}

// StateManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StateManagerTransactorRaw struct {
	Contract *StateManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStateManager creates a new instance of StateManager, bound to a specific deployed contract.
func NewStateManager(address common.Address, backend bind.ContractBackend) (*StateManager, error) {
	contract, err := bindStateManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StateManager{StateManagerCaller: StateManagerCaller{contract: contract}, StateManagerTransactor: StateManagerTransactor{contract: contract}, StateManagerFilterer: StateManagerFilterer{contract: contract}}, nil
}

// NewStateManagerCaller creates a new read-only instance of StateManager, bound to a specific deployed contract.
func NewStateManagerCaller(address common.Address, caller bind.ContractCaller) (*StateManagerCaller, error) {
	contract, err := bindStateManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StateManagerCaller{contract: contract}, nil
}

// NewStateManagerTransactor creates a new write-only instance of StateManager, bound to a specific deployed contract.
func NewStateManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*StateManagerTransactor, error) {
	contract, err := bindStateManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StateManagerTransactor{contract: contract}, nil
}

// NewStateManagerFilterer creates a new log filterer instance of StateManager, bound to a specific deployed contract.
func NewStateManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*StateManagerFilterer, error) {
	contract, err := bindStateManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StateManagerFilterer{contract: contract}, nil
}

// bindStateManager binds a generic wrapper to an already deployed contract.
func bindStateManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StateManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateManager *StateManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateManager.Contract.StateManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateManager *StateManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateManager.Contract.StateManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateManager *StateManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateManager.Contract.StateManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateManager *StateManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateManager *StateManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateManager *StateManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateManager.Contract.contract.Transact(opts, method, params...)
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() view returns(uint256)
func (_StateManager *StateManagerCaller) GetEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StateManager.contract.Call(opts, &out, "getEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() view returns(uint256)
func (_StateManager *StateManagerSession) GetEpoch() (*big.Int, error) {
	return _StateManager.Contract.GetEpoch(&_StateManager.CallOpts)
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() view returns(uint256)
func (_StateManager *StateManagerCallerSession) GetEpoch() (*big.Int, error) {
	return _StateManager.Contract.GetEpoch(&_StateManager.CallOpts)
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() view returns(uint256)
func (_StateManager *StateManagerCaller) GetState(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StateManager.contract.Call(opts, &out, "getState")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() view returns(uint256)
func (_StateManager *StateManagerSession) GetState() (*big.Int, error) {
	return _StateManager.Contract.GetState(&_StateManager.CallOpts)
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() view returns(uint256)
func (_StateManager *StateManagerCallerSession) GetState() (*big.Int, error) {
	return _StateManager.Contract.GetState(&_StateManager.CallOpts)
}
