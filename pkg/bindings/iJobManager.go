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

// IJobManagerABI is the input ABI used to generate the binding from.
const IJobManagerABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"selector\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"repeat\",\"type\":\"bool\"}],\"name\":\"createJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"fulfillJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getJob\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"selector\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"repeat\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getResult\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IJobManager is an auto generated Go binding around an Ethereum contract.
type IJobManager struct {
	IJobManagerCaller     // Read-only binding to the contract
	IJobManagerTransactor // Write-only binding to the contract
	IJobManagerFilterer   // Log filterer for contract events
}

// IJobManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IJobManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IJobManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IJobManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IJobManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IJobManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IJobManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IJobManagerSession struct {
	Contract     *IJobManager      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IJobManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IJobManagerCallerSession struct {
	Contract *IJobManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IJobManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IJobManagerTransactorSession struct {
	Contract     *IJobManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IJobManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IJobManagerRaw struct {
	Contract *IJobManager // Generic contract binding to access the raw methods on
}

// IJobManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IJobManagerCallerRaw struct {
	Contract *IJobManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IJobManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IJobManagerTransactorRaw struct {
	Contract *IJobManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIJobManager creates a new instance of IJobManager, bound to a specific deployed contract.
func NewIJobManager(address common.Address, backend bind.ContractBackend) (*IJobManager, error) {
	contract, err := bindIJobManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IJobManager{IJobManagerCaller: IJobManagerCaller{contract: contract}, IJobManagerTransactor: IJobManagerTransactor{contract: contract}, IJobManagerFilterer: IJobManagerFilterer{contract: contract}}, nil
}

// NewIJobManagerCaller creates a new read-only instance of IJobManager, bound to a specific deployed contract.
func NewIJobManagerCaller(address common.Address, caller bind.ContractCaller) (*IJobManagerCaller, error) {
	contract, err := bindIJobManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IJobManagerCaller{contract: contract}, nil
}

// NewIJobManagerTransactor creates a new write-only instance of IJobManager, bound to a specific deployed contract.
func NewIJobManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IJobManagerTransactor, error) {
	contract, err := bindIJobManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IJobManagerTransactor{contract: contract}, nil
}

// NewIJobManagerFilterer creates a new log filterer instance of IJobManager, bound to a specific deployed contract.
func NewIJobManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IJobManagerFilterer, error) {
	contract, err := bindIJobManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IJobManagerFilterer{contract: contract}, nil
}

// bindIJobManager binds a generic wrapper to an already deployed contract.
func bindIJobManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IJobManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IJobManager *IJobManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IJobManager.Contract.IJobManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IJobManager *IJobManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IJobManager.Contract.IJobManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IJobManager *IJobManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IJobManager.Contract.IJobManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IJobManager *IJobManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IJobManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IJobManager *IJobManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IJobManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IJobManager *IJobManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IJobManager.Contract.contract.Transact(opts, method, params...)
}

// GetJob is a free data retrieval call binding the contract method 0xbf22c457.
//
// Solidity: function getJob(uint256 id) view returns(string url, string selector, string name, bool repeat, uint256 result)
func (_IJobManager *IJobManagerCaller) GetJob(opts *bind.CallOpts, id *big.Int) (struct {
	Url      string
	Selector string
	Name     string
	Repeat   bool
	Result   *big.Int
}, error) {
	var out []interface{}
	err := _IJobManager.contract.Call(opts, &out, "getJob", id)

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
func (_IJobManager *IJobManagerSession) GetJob(id *big.Int) (struct {
	Url      string
	Selector string
	Name     string
	Repeat   bool
	Result   *big.Int
}, error) {
	return _IJobManager.Contract.GetJob(&_IJobManager.CallOpts, id)
}

// GetJob is a free data retrieval call binding the contract method 0xbf22c457.
//
// Solidity: function getJob(uint256 id) view returns(string url, string selector, string name, bool repeat, uint256 result)
func (_IJobManager *IJobManagerCallerSession) GetJob(id *big.Int) (struct {
	Url      string
	Selector string
	Name     string
	Repeat   bool
	Result   *big.Int
}, error) {
	return _IJobManager.Contract.GetJob(&_IJobManager.CallOpts, id)
}

// GetResult is a free data retrieval call binding the contract method 0x995e4339.
//
// Solidity: function getResult(uint256 id) view returns(uint256)
func (_IJobManager *IJobManagerCaller) GetResult(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IJobManager.contract.Call(opts, &out, "getResult", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetResult is a free data retrieval call binding the contract method 0x995e4339.
//
// Solidity: function getResult(uint256 id) view returns(uint256)
func (_IJobManager *IJobManagerSession) GetResult(id *big.Int) (*big.Int, error) {
	return _IJobManager.Contract.GetResult(&_IJobManager.CallOpts, id)
}

// GetResult is a free data retrieval call binding the contract method 0x995e4339.
//
// Solidity: function getResult(uint256 id) view returns(uint256)
func (_IJobManager *IJobManagerCallerSession) GetResult(id *big.Int) (*big.Int, error) {
	return _IJobManager.Contract.GetResult(&_IJobManager.CallOpts, id)
}

// CreateJob is a paid mutator transaction binding the contract method 0x25d10c3f.
//
// Solidity: function createJob(string url, string selector, bool repeat) returns()
func (_IJobManager *IJobManagerTransactor) CreateJob(opts *bind.TransactOpts, url string, selector string, repeat bool) (*types.Transaction, error) {
	return _IJobManager.contract.Transact(opts, "createJob", url, selector, repeat)
}

// CreateJob is a paid mutator transaction binding the contract method 0x25d10c3f.
//
// Solidity: function createJob(string url, string selector, bool repeat) returns()
func (_IJobManager *IJobManagerSession) CreateJob(url string, selector string, repeat bool) (*types.Transaction, error) {
	return _IJobManager.Contract.CreateJob(&_IJobManager.TransactOpts, url, selector, repeat)
}

// CreateJob is a paid mutator transaction binding the contract method 0x25d10c3f.
//
// Solidity: function createJob(string url, string selector, bool repeat) returns()
func (_IJobManager *IJobManagerTransactorSession) CreateJob(url string, selector string, repeat bool) (*types.Transaction, error) {
	return _IJobManager.Contract.CreateJob(&_IJobManager.TransactOpts, url, selector, repeat)
}

// FulfillJob is a paid mutator transaction binding the contract method 0x56350bdf.
//
// Solidity: function fulfillJob(uint256 jobId, uint256 value) returns()
func (_IJobManager *IJobManagerTransactor) FulfillJob(opts *bind.TransactOpts, jobId *big.Int, value *big.Int) (*types.Transaction, error) {
	return _IJobManager.contract.Transact(opts, "fulfillJob", jobId, value)
}

// FulfillJob is a paid mutator transaction binding the contract method 0x56350bdf.
//
// Solidity: function fulfillJob(uint256 jobId, uint256 value) returns()
func (_IJobManager *IJobManagerSession) FulfillJob(jobId *big.Int, value *big.Int) (*types.Transaction, error) {
	return _IJobManager.Contract.FulfillJob(&_IJobManager.TransactOpts, jobId, value)
}

// FulfillJob is a paid mutator transaction binding the contract method 0x56350bdf.
//
// Solidity: function fulfillJob(uint256 jobId, uint256 value) returns()
func (_IJobManager *IJobManagerTransactorSession) FulfillJob(jobId *big.Int, value *big.Int) (*types.Transaction, error) {
	return _IJobManager.Contract.FulfillJob(&_IJobManager.TransactOpts, jobId, value)
}
