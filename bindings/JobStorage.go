// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package JobStorage

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

// JobStorageABI is the input ABI used to generate the binding from.
const JobStorageABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"jobs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"selector\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"repeat\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"credit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numJobs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// JobStorage is an auto generated Go binding around an Ethereum contract.
type JobStorage struct {
	JobStorageCaller     // Read-only binding to the contract
	JobStorageTransactor // Write-only binding to the contract
	JobStorageFilterer   // Log filterer for contract events
}

// JobStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type JobStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JobStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JobStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JobStorageSession struct {
	Contract     *JobStorage       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JobStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JobStorageCallerSession struct {
	Contract *JobStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// JobStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JobStorageTransactorSession struct {
	Contract     *JobStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// JobStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type JobStorageRaw struct {
	Contract *JobStorage // Generic contract binding to access the raw methods on
}

// JobStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JobStorageCallerRaw struct {
	Contract *JobStorageCaller // Generic read-only contract binding to access the raw methods on
}

// JobStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JobStorageTransactorRaw struct {
	Contract *JobStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJobStorage creates a new instance of JobStorage, bound to a specific deployed contract.
func NewJobStorage(address common.Address, backend bind.ContractBackend) (*JobStorage, error) {
	contract, err := bindJobStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &JobStorage{JobStorageCaller: JobStorageCaller{contract: contract}, JobStorageTransactor: JobStorageTransactor{contract: contract}, JobStorageFilterer: JobStorageFilterer{contract: contract}}, nil
}

// NewJobStorageCaller creates a new read-only instance of JobStorage, bound to a specific deployed contract.
func NewJobStorageCaller(address common.Address, caller bind.ContractCaller) (*JobStorageCaller, error) {
	contract, err := bindJobStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JobStorageCaller{contract: contract}, nil
}

// NewJobStorageTransactor creates a new write-only instance of JobStorage, bound to a specific deployed contract.
func NewJobStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*JobStorageTransactor, error) {
	contract, err := bindJobStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JobStorageTransactor{contract: contract}, nil
}

// NewJobStorageFilterer creates a new log filterer instance of JobStorage, bound to a specific deployed contract.
func NewJobStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*JobStorageFilterer, error) {
	contract, err := bindJobStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JobStorageFilterer{contract: contract}, nil
}

// bindJobStorage binds a generic wrapper to an already deployed contract.
func bindJobStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(JobStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JobStorage *JobStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JobStorage.Contract.JobStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JobStorage *JobStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JobStorage.Contract.JobStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JobStorage *JobStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JobStorage.Contract.JobStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JobStorage *JobStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JobStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JobStorage *JobStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JobStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JobStorage *JobStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JobStorage.Contract.contract.Transact(opts, method, params...)
}

// Jobs is a free data retrieval call binding the contract method 0x180aedf3.
//
// Solidity: function jobs(uint256 ) view returns(uint256 id, uint256 epoch, string url, string selector, string name, bool repeat, address creator, uint256 credit, bool fulfilled, uint256 result)
func (_JobStorage *JobStorageCaller) Jobs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id        *big.Int
	Epoch     *big.Int
	Url       string
	Selector  string
	Name      string
	Repeat    bool
	Creator   common.Address
	Credit    *big.Int
	Fulfilled bool
	Result    *big.Int
}, error) {
	var out []interface{}
	err := _JobStorage.contract.Call(opts, &out, "jobs", arg0)

	outstruct := new(struct {
		Id        *big.Int
		Epoch     *big.Int
		Url       string
		Selector  string
		Name      string
		Repeat    bool
		Creator   common.Address
		Credit    *big.Int
		Fulfilled bool
		Result    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Epoch = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Url = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Selector = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Name = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Repeat = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.Creator = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.Credit = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Fulfilled = *abi.ConvertType(out[8], new(bool)).(*bool)
	outstruct.Result = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Jobs is a free data retrieval call binding the contract method 0x180aedf3.
//
// Solidity: function jobs(uint256 ) view returns(uint256 id, uint256 epoch, string url, string selector, string name, bool repeat, address creator, uint256 credit, bool fulfilled, uint256 result)
func (_JobStorage *JobStorageSession) Jobs(arg0 *big.Int) (struct {
	Id        *big.Int
	Epoch     *big.Int
	Url       string
	Selector  string
	Name      string
	Repeat    bool
	Creator   common.Address
	Credit    *big.Int
	Fulfilled bool
	Result    *big.Int
}, error) {
	return _JobStorage.Contract.Jobs(&_JobStorage.CallOpts, arg0)
}

// Jobs is a free data retrieval call binding the contract method 0x180aedf3.
//
// Solidity: function jobs(uint256 ) view returns(uint256 id, uint256 epoch, string url, string selector, string name, bool repeat, address creator, uint256 credit, bool fulfilled, uint256 result)
func (_JobStorage *JobStorageCallerSession) Jobs(arg0 *big.Int) (struct {
	Id        *big.Int
	Epoch     *big.Int
	Url       string
	Selector  string
	Name      string
	Repeat    bool
	Creator   common.Address
	Credit    *big.Int
	Fulfilled bool
	Result    *big.Int
}, error) {
	return _JobStorage.Contract.Jobs(&_JobStorage.CallOpts, arg0)
}

// NumJobs is a free data retrieval call binding the contract method 0x9212051c.
//
// Solidity: function numJobs() view returns(uint256)
func (_JobStorage *JobStorageCaller) NumJobs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _JobStorage.contract.Call(opts, &out, "numJobs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumJobs is a free data retrieval call binding the contract method 0x9212051c.
//
// Solidity: function numJobs() view returns(uint256)
func (_JobStorage *JobStorageSession) NumJobs() (*big.Int, error) {
	return _JobStorage.Contract.NumJobs(&_JobStorage.CallOpts)
}

// NumJobs is a free data retrieval call binding the contract method 0x9212051c.
//
// Solidity: function numJobs() view returns(uint256)
func (_JobStorage *JobStorageCallerSession) NumJobs() (*big.Int, error) {
	return _JobStorage.Contract.NumJobs(&_JobStorage.CallOpts)
}
