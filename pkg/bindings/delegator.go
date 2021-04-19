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

// DelegatorABI is the input ABI used to generate the binding from.
const DelegatorABI = "[{\"inputs\":[],\"name\":\"delegate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getJob\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"selector\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"repeat\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getResult\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"jobManager\",\"outputs\":[{\"internalType\":\"contractIJobManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDelegateAddress\",\"type\":\"address\"}],\"name\":\"upgradeDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Delegator is an auto generated Go binding around an Ethereum contract.
type Delegator struct {
	DelegatorCaller     // Read-only binding to the contract
	DelegatorTransactor // Write-only binding to the contract
	DelegatorFilterer   // Log filterer for contract events
}

// DelegatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type DelegatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DelegatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DelegatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DelegatorSession struct {
	Contract     *Delegator        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DelegatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DelegatorCallerSession struct {
	Contract *DelegatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DelegatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DelegatorTransactorSession struct {
	Contract     *DelegatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DelegatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type DelegatorRaw struct {
	Contract *Delegator // Generic contract binding to access the raw methods on
}

// DelegatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DelegatorCallerRaw struct {
	Contract *DelegatorCaller // Generic read-only contract binding to access the raw methods on
}

// DelegatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DelegatorTransactorRaw struct {
	Contract *DelegatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDelegator creates a new instance of Delegator, bound to a specific deployed contract.
func NewDelegator(address common.Address, backend bind.ContractBackend) (*Delegator, error) {
	contract, err := bindDelegator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Delegator{DelegatorCaller: DelegatorCaller{contract: contract}, DelegatorTransactor: DelegatorTransactor{contract: contract}, DelegatorFilterer: DelegatorFilterer{contract: contract}}, nil
}

// NewDelegatorCaller creates a new read-only instance of Delegator, bound to a specific deployed contract.
func NewDelegatorCaller(address common.Address, caller bind.ContractCaller) (*DelegatorCaller, error) {
	contract, err := bindDelegator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DelegatorCaller{contract: contract}, nil
}

// NewDelegatorTransactor creates a new write-only instance of Delegator, bound to a specific deployed contract.
func NewDelegatorTransactor(address common.Address, transactor bind.ContractTransactor) (*DelegatorTransactor, error) {
	contract, err := bindDelegator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DelegatorTransactor{contract: contract}, nil
}

// NewDelegatorFilterer creates a new log filterer instance of Delegator, bound to a specific deployed contract.
func NewDelegatorFilterer(address common.Address, filterer bind.ContractFilterer) (*DelegatorFilterer, error) {
	contract, err := bindDelegator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DelegatorFilterer{contract: contract}, nil
}

// bindDelegator binds a generic wrapper to an already deployed contract.
func bindDelegator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DelegatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Delegator *DelegatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Delegator.Contract.DelegatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Delegator *DelegatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Delegator.Contract.DelegatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Delegator *DelegatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Delegator.Contract.DelegatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Delegator *DelegatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Delegator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Delegator *DelegatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Delegator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Delegator *DelegatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Delegator.Contract.contract.Transact(opts, method, params...)
}

// Delegate is a free data retrieval call binding the contract method 0xc89e4361.
//
// Solidity: function delegate() view returns(address)
func (_Delegator *DelegatorCaller) Delegate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Delegator.contract.Call(opts, &out, "delegate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegate is a free data retrieval call binding the contract method 0xc89e4361.
//
// Solidity: function delegate() view returns(address)
func (_Delegator *DelegatorSession) Delegate() (common.Address, error) {
	return _Delegator.Contract.Delegate(&_Delegator.CallOpts)
}

// Delegate is a free data retrieval call binding the contract method 0xc89e4361.
//
// Solidity: function delegate() view returns(address)
func (_Delegator *DelegatorCallerSession) Delegate() (common.Address, error) {
	return _Delegator.Contract.Delegate(&_Delegator.CallOpts)
}

// GetJob is a free data retrieval call binding the contract method 0xbf22c457.
//
// Solidity: function getJob(uint256 id) view returns(string url, string selector, string name, bool repeat, uint256 result)
func (_Delegator *DelegatorCaller) GetJob(opts *bind.CallOpts, id *big.Int) (struct {
	Url      string
	Selector string
	Name     string
	Repeat   bool
	Result   *big.Int
}, error) {
	var out []interface{}
	err := _Delegator.contract.Call(opts, &out, "getJob", id)

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
func (_Delegator *DelegatorSession) GetJob(id *big.Int) (struct {
	Url      string
	Selector string
	Name     string
	Repeat   bool
	Result   *big.Int
}, error) {
	return _Delegator.Contract.GetJob(&_Delegator.CallOpts, id)
}

// GetJob is a free data retrieval call binding the contract method 0xbf22c457.
//
// Solidity: function getJob(uint256 id) view returns(string url, string selector, string name, bool repeat, uint256 result)
func (_Delegator *DelegatorCallerSession) GetJob(id *big.Int) (struct {
	Url      string
	Selector string
	Name     string
	Repeat   bool
	Result   *big.Int
}, error) {
	return _Delegator.Contract.GetJob(&_Delegator.CallOpts, id)
}

// GetResult is a free data retrieval call binding the contract method 0x995e4339.
//
// Solidity: function getResult(uint256 id) view returns(uint256)
func (_Delegator *DelegatorCaller) GetResult(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Delegator.contract.Call(opts, &out, "getResult", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetResult is a free data retrieval call binding the contract method 0x995e4339.
//
// Solidity: function getResult(uint256 id) view returns(uint256)
func (_Delegator *DelegatorSession) GetResult(id *big.Int) (*big.Int, error) {
	return _Delegator.Contract.GetResult(&_Delegator.CallOpts, id)
}

// GetResult is a free data retrieval call binding the contract method 0x995e4339.
//
// Solidity: function getResult(uint256 id) view returns(uint256)
func (_Delegator *DelegatorCallerSession) GetResult(id *big.Int) (*big.Int, error) {
	return _Delegator.Contract.GetResult(&_Delegator.CallOpts, id)
}

// JobManager is a free data retrieval call binding the contract method 0x3df395a3.
//
// Solidity: function jobManager() view returns(address)
func (_Delegator *DelegatorCaller) JobManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Delegator.contract.Call(opts, &out, "jobManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// JobManager is a free data retrieval call binding the contract method 0x3df395a3.
//
// Solidity: function jobManager() view returns(address)
func (_Delegator *DelegatorSession) JobManager() (common.Address, error) {
	return _Delegator.Contract.JobManager(&_Delegator.CallOpts)
}

// JobManager is a free data retrieval call binding the contract method 0x3df395a3.
//
// Solidity: function jobManager() view returns(address)
func (_Delegator *DelegatorCallerSession) JobManager() (common.Address, error) {
	return _Delegator.Contract.JobManager(&_Delegator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Delegator *DelegatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Delegator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Delegator *DelegatorSession) Owner() (common.Address, error) {
	return _Delegator.Contract.Owner(&_Delegator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Delegator *DelegatorCallerSession) Owner() (common.Address, error) {
	return _Delegator.Contract.Owner(&_Delegator.CallOpts)
}

// UpgradeDelegate is a paid mutator transaction binding the contract method 0x2da4e75c.
//
// Solidity: function upgradeDelegate(address newDelegateAddress) returns()
func (_Delegator *DelegatorTransactor) UpgradeDelegate(opts *bind.TransactOpts, newDelegateAddress common.Address) (*types.Transaction, error) {
	return _Delegator.contract.Transact(opts, "upgradeDelegate", newDelegateAddress)
}

// UpgradeDelegate is a paid mutator transaction binding the contract method 0x2da4e75c.
//
// Solidity: function upgradeDelegate(address newDelegateAddress) returns()
func (_Delegator *DelegatorSession) UpgradeDelegate(newDelegateAddress common.Address) (*types.Transaction, error) {
	return _Delegator.Contract.UpgradeDelegate(&_Delegator.TransactOpts, newDelegateAddress)
}

// UpgradeDelegate is a paid mutator transaction binding the contract method 0x2da4e75c.
//
// Solidity: function upgradeDelegate(address newDelegateAddress) returns()
func (_Delegator *DelegatorTransactorSession) UpgradeDelegate(newDelegateAddress common.Address) (*types.Transaction, error) {
	return _Delegator.Contract.UpgradeDelegate(&_Delegator.TransactOpts, newDelegateAddress)
}
