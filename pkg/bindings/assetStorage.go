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

// AssetStorageABI is the input ABI used to generate the binding from.
const AssetStorageABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"collections\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"aggregationMethod\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"credit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"jobs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"selector\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"repeat\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"credit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AssetStorage is an auto generated Go binding around an Ethereum contract.
type AssetStorage struct {
	AssetStorageCaller     // Read-only binding to the contract
	AssetStorageTransactor // Write-only binding to the contract
	AssetStorageFilterer   // Log filterer for contract events
}

// AssetStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssetStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssetStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssetStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssetStorageSession struct {
	Contract     *AssetStorage     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssetStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssetStorageCallerSession struct {
	Contract *AssetStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AssetStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssetStorageTransactorSession struct {
	Contract     *AssetStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AssetStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssetStorageRaw struct {
	Contract *AssetStorage // Generic contract binding to access the raw methods on
}

// AssetStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssetStorageCallerRaw struct {
	Contract *AssetStorageCaller // Generic read-only contract binding to access the raw methods on
}

// AssetStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssetStorageTransactorRaw struct {
	Contract *AssetStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssetStorage creates a new instance of AssetStorage, bound to a specific deployed contract.
func NewAssetStorage(address common.Address, backend bind.ContractBackend) (*AssetStorage, error) {
	contract, err := bindAssetStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AssetStorage{AssetStorageCaller: AssetStorageCaller{contract: contract}, AssetStorageTransactor: AssetStorageTransactor{contract: contract}, AssetStorageFilterer: AssetStorageFilterer{contract: contract}}, nil
}

// NewAssetStorageCaller creates a new read-only instance of AssetStorage, bound to a specific deployed contract.
func NewAssetStorageCaller(address common.Address, caller bind.ContractCaller) (*AssetStorageCaller, error) {
	contract, err := bindAssetStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssetStorageCaller{contract: contract}, nil
}

// NewAssetStorageTransactor creates a new write-only instance of AssetStorage, bound to a specific deployed contract.
func NewAssetStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*AssetStorageTransactor, error) {
	contract, err := bindAssetStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssetStorageTransactor{contract: contract}, nil
}

// NewAssetStorageFilterer creates a new log filterer instance of AssetStorage, bound to a specific deployed contract.
func NewAssetStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*AssetStorageFilterer, error) {
	contract, err := bindAssetStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssetStorageFilterer{contract: contract}, nil
}

// bindAssetStorage binds a generic wrapper to an already deployed contract.
func bindAssetStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AssetStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetStorage *AssetStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssetStorage.Contract.AssetStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetStorage *AssetStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetStorage.Contract.AssetStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetStorage *AssetStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetStorage.Contract.AssetStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetStorage *AssetStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssetStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetStorage *AssetStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetStorage *AssetStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetStorage.Contract.contract.Transact(opts, method, params...)
}

// Collections is a free data retrieval call binding the contract method 0xfdbda0ec.
//
// Solidity: function collections(uint256 ) view returns(uint256 id, string name, uint32 aggregationMethod, uint256 epoch, address creator, uint256 credit, uint256 result, uint256 assetType)
func (_AssetStorage *AssetStorageCaller) Collections(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id                *big.Int
	Name              string
	AggregationMethod uint32
	Epoch             *big.Int
	Creator           common.Address
	Credit            *big.Int
	Result            *big.Int
	AssetType         *big.Int
}, error) {
	var out []interface{}
	err := _AssetStorage.contract.Call(opts, &out, "collections", arg0)

	outstruct := new(struct {
		Id                *big.Int
		Name              string
		AggregationMethod uint32
		Epoch             *big.Int
		Creator           common.Address
		Credit            *big.Int
		Result            *big.Int
		AssetType         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.AggregationMethod = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.Epoch = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Creator = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Credit = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Result = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.AssetType = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Collections is a free data retrieval call binding the contract method 0xfdbda0ec.
//
// Solidity: function collections(uint256 ) view returns(uint256 id, string name, uint32 aggregationMethod, uint256 epoch, address creator, uint256 credit, uint256 result, uint256 assetType)
func (_AssetStorage *AssetStorageSession) Collections(arg0 *big.Int) (struct {
	Id                *big.Int
	Name              string
	AggregationMethod uint32
	Epoch             *big.Int
	Creator           common.Address
	Credit            *big.Int
	Result            *big.Int
	AssetType         *big.Int
}, error) {
	return _AssetStorage.Contract.Collections(&_AssetStorage.CallOpts, arg0)
}

// Collections is a free data retrieval call binding the contract method 0xfdbda0ec.
//
// Solidity: function collections(uint256 ) view returns(uint256 id, string name, uint32 aggregationMethod, uint256 epoch, address creator, uint256 credit, uint256 result, uint256 assetType)
func (_AssetStorage *AssetStorageCallerSession) Collections(arg0 *big.Int) (struct {
	Id                *big.Int
	Name              string
	AggregationMethod uint32
	Epoch             *big.Int
	Creator           common.Address
	Credit            *big.Int
	Result            *big.Int
	AssetType         *big.Int
}, error) {
	return _AssetStorage.Contract.Collections(&_AssetStorage.CallOpts, arg0)
}

// Jobs is a free data retrieval call binding the contract method 0x180aedf3.
//
// Solidity: function jobs(uint256 ) view returns(uint256 id, uint256 epoch, string url, string selector, string name, bool repeat, address creator, uint256 credit, bool fulfilled, uint256 result, uint256 assetType)
func (_AssetStorage *AssetStorageCaller) Jobs(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	AssetType *big.Int
}, error) {
	var out []interface{}
	err := _AssetStorage.contract.Call(opts, &out, "jobs", arg0)

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
		AssetType *big.Int
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
	outstruct.AssetType = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Jobs is a free data retrieval call binding the contract method 0x180aedf3.
//
// Solidity: function jobs(uint256 ) view returns(uint256 id, uint256 epoch, string url, string selector, string name, bool repeat, address creator, uint256 credit, bool fulfilled, uint256 result, uint256 assetType)
func (_AssetStorage *AssetStorageSession) Jobs(arg0 *big.Int) (struct {
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
	AssetType *big.Int
}, error) {
	return _AssetStorage.Contract.Jobs(&_AssetStorage.CallOpts, arg0)
}

// Jobs is a free data retrieval call binding the contract method 0x180aedf3.
//
// Solidity: function jobs(uint256 ) view returns(uint256 id, uint256 epoch, string url, string selector, string name, bool repeat, address creator, uint256 credit, bool fulfilled, uint256 result, uint256 assetType)
func (_AssetStorage *AssetStorageCallerSession) Jobs(arg0 *big.Int) (struct {
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
	AssetType *big.Int
}, error) {
	return _AssetStorage.Contract.Jobs(&_AssetStorage.CallOpts, arg0)
}

// NumAssets is a free data retrieval call binding the contract method 0xa46fe83b.
//
// Solidity: function numAssets() view returns(uint256)
func (_AssetStorage *AssetStorageCaller) NumAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AssetStorage.contract.Call(opts, &out, "numAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumAssets is a free data retrieval call binding the contract method 0xa46fe83b.
//
// Solidity: function numAssets() view returns(uint256)
func (_AssetStorage *AssetStorageSession) NumAssets() (*big.Int, error) {
	return _AssetStorage.Contract.NumAssets(&_AssetStorage.CallOpts)
}

// NumAssets is a free data retrieval call binding the contract method 0xa46fe83b.
//
// Solidity: function numAssets() view returns(uint256)
func (_AssetStorage *AssetStorageCallerSession) NumAssets() (*big.Int, error) {
	return _AssetStorage.Contract.NumAssets(&_AssetStorage.CallOpts)
}
