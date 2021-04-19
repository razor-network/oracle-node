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

// IDelegatorABI is the input ABI used to generate the binding from.
const IDelegatorABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getJob\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"selector\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"repeat\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getResult\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDelegateAddress\",\"type\":\"address\"}],\"name\":\"upgradeDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IDelegator is an auto generated Go binding around an Ethereum contract.
type IDelegator struct {
	IDelegatorCaller     // Read-only binding to the contract
	IDelegatorTransactor // Write-only binding to the contract
	IDelegatorFilterer   // Log filterer for contract events
}

// IDelegatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDelegatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDelegatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDelegatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDelegatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDelegatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDelegatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDelegatorSession struct {
	Contract     *IDelegator       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDelegatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDelegatorCallerSession struct {
	Contract *IDelegatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IDelegatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDelegatorTransactorSession struct {
	Contract     *IDelegatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IDelegatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDelegatorRaw struct {
	Contract *IDelegator // Generic contract binding to access the raw methods on
}

// IDelegatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDelegatorCallerRaw struct {
	Contract *IDelegatorCaller // Generic read-only contract binding to access the raw methods on
}

// IDelegatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDelegatorTransactorRaw struct {
	Contract *IDelegatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDelegator creates a new instance of IDelegator, bound to a specific deployed contract.
func NewIDelegator(address common.Address, backend bind.ContractBackend) (*IDelegator, error) {
	contract, err := bindIDelegator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDelegator{IDelegatorCaller: IDelegatorCaller{contract: contract}, IDelegatorTransactor: IDelegatorTransactor{contract: contract}, IDelegatorFilterer: IDelegatorFilterer{contract: contract}}, nil
}

// NewIDelegatorCaller creates a new read-only instance of IDelegator, bound to a specific deployed contract.
func NewIDelegatorCaller(address common.Address, caller bind.ContractCaller) (*IDelegatorCaller, error) {
	contract, err := bindIDelegator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDelegatorCaller{contract: contract}, nil
}

// NewIDelegatorTransactor creates a new write-only instance of IDelegator, bound to a specific deployed contract.
func NewIDelegatorTransactor(address common.Address, transactor bind.ContractTransactor) (*IDelegatorTransactor, error) {
	contract, err := bindIDelegator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDelegatorTransactor{contract: contract}, nil
}

// NewIDelegatorFilterer creates a new log filterer instance of IDelegator, bound to a specific deployed contract.
func NewIDelegatorFilterer(address common.Address, filterer bind.ContractFilterer) (*IDelegatorFilterer, error) {
	contract, err := bindIDelegator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDelegatorFilterer{contract: contract}, nil
}

// bindIDelegator binds a generic wrapper to an already deployed contract.
func bindIDelegator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDelegatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDelegator *IDelegatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDelegator.Contract.IDelegatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDelegator *IDelegatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDelegator.Contract.IDelegatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDelegator *IDelegatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDelegator.Contract.IDelegatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDelegator *IDelegatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDelegator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDelegator *IDelegatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDelegator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDelegator *IDelegatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDelegator.Contract.contract.Transact(opts, method, params...)
}

// GetJob is a free data retrieval call binding the contract method 0xbf22c457.
//
// Solidity: function getJob(uint256 id) view returns(string url, string selector, string name, bool repeat, uint256 result)
func (_IDelegator *IDelegatorCaller) GetJob(opts *bind.CallOpts, id *big.Int) (struct {
	Url      string
	Selector string
	Name     string
	Repeat   bool
	Result   *big.Int
}, error) {
	var out []interface{}
	err := _IDelegator.contract.Call(opts, &out, "getJob", id)

	outstruct := new(struct {
		Url      string
		Selector string
		Name     string
		Repeat   bool
		Result   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Url = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Selector = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Name = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Repeat = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.Result = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetJob is a free data retrieval call binding the contract method 0xbf22c457.
//
// Solidity: function getJob(uint256 id) view returns(string url, string selector, string name, bool repeat, uint256 result)
func (_IDelegator *IDelegatorSession) GetJob(id *big.Int) (struct {
	Url      string
	Selector string
	Name     string
	Repeat   bool
	Result   *big.Int
}, error) {
	return _IDelegator.Contract.GetJob(&_IDelegator.CallOpts, id)
}

// GetJob is a free data retrieval call binding the contract method 0xbf22c457.
//
// Solidity: function getJob(uint256 id) view returns(string url, string selector, string name, bool repeat, uint256 result)
func (_IDelegator *IDelegatorCallerSession) GetJob(id *big.Int) (struct {
	Url      string
	Selector string
	Name     string
	Repeat   bool
	Result   *big.Int
}, error) {
	return _IDelegator.Contract.GetJob(&_IDelegator.CallOpts, id)
}

// GetResult is a free data retrieval call binding the contract method 0x995e4339.
//
// Solidity: function getResult(uint256 id) view returns(uint256)
func (_IDelegator *IDelegatorCaller) GetResult(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IDelegator.contract.Call(opts, &out, "getResult", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetResult is a free data retrieval call binding the contract method 0x995e4339.
//
// Solidity: function getResult(uint256 id) view returns(uint256)
func (_IDelegator *IDelegatorSession) GetResult(id *big.Int) (*big.Int, error) {
	return _IDelegator.Contract.GetResult(&_IDelegator.CallOpts, id)
}

// GetResult is a free data retrieval call binding the contract method 0x995e4339.
//
// Solidity: function getResult(uint256 id) view returns(uint256)
func (_IDelegator *IDelegatorCallerSession) GetResult(id *big.Int) (*big.Int, error) {
	return _IDelegator.Contract.GetResult(&_IDelegator.CallOpts, id)
}

// UpgradeDelegate is a paid mutator transaction binding the contract method 0x2da4e75c.
//
// Solidity: function upgradeDelegate(address newDelegateAddress) returns()
func (_IDelegator *IDelegatorTransactor) UpgradeDelegate(opts *bind.TransactOpts, newDelegateAddress common.Address) (*types.Transaction, error) {
	return _IDelegator.contract.Transact(opts, "upgradeDelegate", newDelegateAddress)
}

// UpgradeDelegate is a paid mutator transaction binding the contract method 0x2da4e75c.
//
// Solidity: function upgradeDelegate(address newDelegateAddress) returns()
func (_IDelegator *IDelegatorSession) UpgradeDelegate(newDelegateAddress common.Address) (*types.Transaction, error) {
	return _IDelegator.Contract.UpgradeDelegate(&_IDelegator.TransactOpts, newDelegateAddress)
}

// UpgradeDelegate is a paid mutator transaction binding the contract method 0x2da4e75c.
//
// Solidity: function upgradeDelegate(address newDelegateAddress) returns()
func (_IDelegator *IDelegatorTransactorSession) UpgradeDelegate(newDelegateAddress common.Address) (*types.Transaction, error) {
	return _IDelegator.Contract.UpgradeDelegate(&_IDelegator.TransactOpts, newDelegateAddress)
}
