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

// AssetManagerABI is the input ABI used to generate the binding from.
const AssetManagerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"parametersAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"aggregationMethod\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"jobIDs\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"credit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumAssetStorage.assetTypes\",\"name\":\"assetType\",\"type\":\"uint8\"}],\"name\":\"CollectionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"aggregationMethod\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"jobIDs\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"credit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CollectionReported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"updatedJobIDs\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CollectionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"selector\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"repeat\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"credit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumAssetStorage.assetTypes\",\"name\":\"assetType\",\"type\":\"uint8\"}],\"name\":\"JobCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"selector\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"repeat\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"credit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"JobReported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collectionID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobID\",\"type\":\"uint256\"}],\"name\":\"addJobToCollection\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"collections\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"aggregationMethod\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"credit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"jobIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint32\",\"name\":\"aggregationMethod\",\"type\":\"uint32\"}],\"name\":\"createCollection\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"selector\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"repeat\",\"type\":\"bool\"}],\"name\":\"createJob\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"fulfillAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getAssetType\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getCollection\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"aggregationMethod\",\"type\":\"uint32\"},{\"internalType\":\"uint256[]\",\"name\":\"jobIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getJob\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"selector\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"repeat\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getResult\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"jobs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"selector\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"repeat\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"credit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetType\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"parameters\",\"outputs\":[{\"internalType\":\"contractIParameters\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AssetManager is an auto generated Go binding around an Ethereum contract.
type AssetManager struct {
	AssetManagerCaller     // Read-only binding to the contract
	AssetManagerTransactor // Write-only binding to the contract
	AssetManagerFilterer   // Log filterer for contract events
}

// AssetManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssetManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssetManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssetManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssetManagerSession struct {
	Contract     *AssetManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssetManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssetManagerCallerSession struct {
	Contract *AssetManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AssetManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssetManagerTransactorSession struct {
	Contract     *AssetManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AssetManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssetManagerRaw struct {
	Contract *AssetManager // Generic contract binding to access the raw methods on
}

// AssetManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssetManagerCallerRaw struct {
	Contract *AssetManagerCaller // Generic read-only contract binding to access the raw methods on
}

// AssetManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssetManagerTransactorRaw struct {
	Contract *AssetManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssetManager creates a new instance of AssetManager, bound to a specific deployed contract.
func NewAssetManager(address common.Address, backend bind.ContractBackend) (*AssetManager, error) {
	contract, err := bindAssetManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AssetManager{AssetManagerCaller: AssetManagerCaller{contract: contract}, AssetManagerTransactor: AssetManagerTransactor{contract: contract}, AssetManagerFilterer: AssetManagerFilterer{contract: contract}}, nil
}

// NewAssetManagerCaller creates a new read-only instance of AssetManager, bound to a specific deployed contract.
func NewAssetManagerCaller(address common.Address, caller bind.ContractCaller) (*AssetManagerCaller, error) {
	contract, err := bindAssetManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssetManagerCaller{contract: contract}, nil
}

// NewAssetManagerTransactor creates a new write-only instance of AssetManager, bound to a specific deployed contract.
func NewAssetManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*AssetManagerTransactor, error) {
	contract, err := bindAssetManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssetManagerTransactor{contract: contract}, nil
}

// NewAssetManagerFilterer creates a new log filterer instance of AssetManager, bound to a specific deployed contract.
func NewAssetManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*AssetManagerFilterer, error) {
	contract, err := bindAssetManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssetManagerFilterer{contract: contract}, nil
}

// bindAssetManager binds a generic wrapper to an already deployed contract.
func bindAssetManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AssetManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetManager *AssetManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssetManager.Contract.AssetManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetManager *AssetManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetManager.Contract.AssetManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetManager *AssetManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetManager.Contract.AssetManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetManager *AssetManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssetManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetManager *AssetManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetManager *AssetManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetManager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AssetManager *AssetManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AssetManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AssetManager *AssetManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AssetManager.Contract.DEFAULTADMINROLE(&_AssetManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AssetManager *AssetManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AssetManager.Contract.DEFAULTADMINROLE(&_AssetManager.CallOpts)
}

// Collections is a free data retrieval call binding the contract method 0xfdbda0ec.
//
// Solidity: function collections(uint256 ) view returns(uint256 id, string name, uint32 aggregationMethod, uint256 epoch, address creator, uint256 credit, uint256 result, uint256 assetType)
func (_AssetManager *AssetManagerCaller) Collections(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _AssetManager.contract.Call(opts, &out, "collections", arg0)

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
func (_AssetManager *AssetManagerSession) Collections(arg0 *big.Int) (struct {
	Id                *big.Int
	Name              string
	AggregationMethod uint32
	Epoch             *big.Int
	Creator           common.Address
	Credit            *big.Int
	Result            *big.Int
	AssetType         *big.Int
}, error) {
	return _AssetManager.Contract.Collections(&_AssetManager.CallOpts, arg0)
}

// Collections is a free data retrieval call binding the contract method 0xfdbda0ec.
//
// Solidity: function collections(uint256 ) view returns(uint256 id, string name, uint32 aggregationMethod, uint256 epoch, address creator, uint256 credit, uint256 result, uint256 assetType)
func (_AssetManager *AssetManagerCallerSession) Collections(arg0 *big.Int) (struct {
	Id                *big.Int
	Name              string
	AggregationMethod uint32
	Epoch             *big.Int
	Creator           common.Address
	Credit            *big.Int
	Result            *big.Int
	AssetType         *big.Int
}, error) {
	return _AssetManager.Contract.Collections(&_AssetManager.CallOpts, arg0)
}

// GetAssetType is a free data retrieval call binding the contract method 0x9a4b8c0d.
//
// Solidity: function getAssetType(uint256 id) view returns(uint256)
func (_AssetManager *AssetManagerCaller) GetAssetType(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AssetManager.contract.Call(opts, &out, "getAssetType", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAssetType is a free data retrieval call binding the contract method 0x9a4b8c0d.
//
// Solidity: function getAssetType(uint256 id) view returns(uint256)
func (_AssetManager *AssetManagerSession) GetAssetType(id *big.Int) (*big.Int, error) {
	return _AssetManager.Contract.GetAssetType(&_AssetManager.CallOpts, id)
}

// GetAssetType is a free data retrieval call binding the contract method 0x9a4b8c0d.
//
// Solidity: function getAssetType(uint256 id) view returns(uint256)
func (_AssetManager *AssetManagerCallerSession) GetAssetType(id *big.Int) (*big.Int, error) {
	return _AssetManager.Contract.GetAssetType(&_AssetManager.CallOpts, id)
}

// GetCollection is a free data retrieval call binding the contract method 0x5a1f3c28.
//
// Solidity: function getCollection(uint256 id) view returns(string name, uint32 aggregationMethod, uint256[] jobIDs, uint256 result)
func (_AssetManager *AssetManagerCaller) GetCollection(opts *bind.CallOpts, id *big.Int) (struct {
	Name              string
	AggregationMethod uint32
	JobIDs            []*big.Int
	Result            *big.Int
}, error) {
	var out []interface{}
	err := _AssetManager.contract.Call(opts, &out, "getCollection", id)

	outstruct := new(struct {
		Name              string
		AggregationMethod uint32
		JobIDs            []*big.Int
		Result            *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.AggregationMethod = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.JobIDs = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	outstruct.Result = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCollection is a free data retrieval call binding the contract method 0x5a1f3c28.
//
// Solidity: function getCollection(uint256 id) view returns(string name, uint32 aggregationMethod, uint256[] jobIDs, uint256 result)
func (_AssetManager *AssetManagerSession) GetCollection(id *big.Int) (struct {
	Name              string
	AggregationMethod uint32
	JobIDs            []*big.Int
	Result            *big.Int
}, error) {
	return _AssetManager.Contract.GetCollection(&_AssetManager.CallOpts, id)
}

// GetCollection is a free data retrieval call binding the contract method 0x5a1f3c28.
//
// Solidity: function getCollection(uint256 id) view returns(string name, uint32 aggregationMethod, uint256[] jobIDs, uint256 result)
func (_AssetManager *AssetManagerCallerSession) GetCollection(id *big.Int) (struct {
	Name              string
	AggregationMethod uint32
	JobIDs            []*big.Int
	Result            *big.Int
}, error) {
	return _AssetManager.Contract.GetCollection(&_AssetManager.CallOpts, id)
}

// GetJob is a free data retrieval call binding the contract method 0xbf22c457.
//
// Solidity: function getJob(uint256 id) view returns(string url, string selector, string name, bool repeat, uint256 result)
func (_AssetManager *AssetManagerCaller) GetJob(opts *bind.CallOpts, id *big.Int) (struct {
	Url      string
	Selector string
	Name     string
	Repeat   bool
	Result   *big.Int
}, error) {
	var out []interface{}
	err := _AssetManager.contract.Call(opts, &out, "getJob", id)

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
func (_AssetManager *AssetManagerSession) GetJob(id *big.Int) (struct {
	Url      string
	Selector string
	Name     string
	Repeat   bool
	Result   *big.Int
}, error) {
	return _AssetManager.Contract.GetJob(&_AssetManager.CallOpts, id)
}

// GetJob is a free data retrieval call binding the contract method 0xbf22c457.
//
// Solidity: function getJob(uint256 id) view returns(string url, string selector, string name, bool repeat, uint256 result)
func (_AssetManager *AssetManagerCallerSession) GetJob(id *big.Int) (struct {
	Url      string
	Selector string
	Name     string
	Repeat   bool
	Result   *big.Int
}, error) {
	return _AssetManager.Contract.GetJob(&_AssetManager.CallOpts, id)
}

// GetNumAssets is a free data retrieval call binding the contract method 0x7812e0ce.
//
// Solidity: function getNumAssets() view returns(uint256)
func (_AssetManager *AssetManagerCaller) GetNumAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AssetManager.contract.Call(opts, &out, "getNumAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumAssets is a free data retrieval call binding the contract method 0x7812e0ce.
//
// Solidity: function getNumAssets() view returns(uint256)
func (_AssetManager *AssetManagerSession) GetNumAssets() (*big.Int, error) {
	return _AssetManager.Contract.GetNumAssets(&_AssetManager.CallOpts)
}

// GetNumAssets is a free data retrieval call binding the contract method 0x7812e0ce.
//
// Solidity: function getNumAssets() view returns(uint256)
func (_AssetManager *AssetManagerCallerSession) GetNumAssets() (*big.Int, error) {
	return _AssetManager.Contract.GetNumAssets(&_AssetManager.CallOpts)
}

// GetResult is a free data retrieval call binding the contract method 0x995e4339.
//
// Solidity: function getResult(uint256 id) view returns(uint256 result)
func (_AssetManager *AssetManagerCaller) GetResult(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AssetManager.contract.Call(opts, &out, "getResult", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetResult is a free data retrieval call binding the contract method 0x995e4339.
//
// Solidity: function getResult(uint256 id) view returns(uint256 result)
func (_AssetManager *AssetManagerSession) GetResult(id *big.Int) (*big.Int, error) {
	return _AssetManager.Contract.GetResult(&_AssetManager.CallOpts, id)
}

// GetResult is a free data retrieval call binding the contract method 0x995e4339.
//
// Solidity: function getResult(uint256 id) view returns(uint256 result)
func (_AssetManager *AssetManagerCallerSession) GetResult(id *big.Int) (*big.Int, error) {
	return _AssetManager.Contract.GetResult(&_AssetManager.CallOpts, id)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AssetManager *AssetManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AssetManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AssetManager *AssetManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AssetManager.Contract.GetRoleAdmin(&_AssetManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AssetManager *AssetManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AssetManager.Contract.GetRoleAdmin(&_AssetManager.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AssetManager *AssetManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AssetManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AssetManager *AssetManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AssetManager.Contract.HasRole(&_AssetManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AssetManager *AssetManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AssetManager.Contract.HasRole(&_AssetManager.CallOpts, role, account)
}

// Jobs is a free data retrieval call binding the contract method 0x180aedf3.
//
// Solidity: function jobs(uint256 ) view returns(uint256 id, uint256 epoch, string url, string selector, string name, bool repeat, address creator, uint256 credit, bool fulfilled, uint256 result, uint256 assetType)
func (_AssetManager *AssetManagerCaller) Jobs(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _AssetManager.contract.Call(opts, &out, "jobs", arg0)

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
func (_AssetManager *AssetManagerSession) Jobs(arg0 *big.Int) (struct {
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
	return _AssetManager.Contract.Jobs(&_AssetManager.CallOpts, arg0)
}

// Jobs is a free data retrieval call binding the contract method 0x180aedf3.
//
// Solidity: function jobs(uint256 ) view returns(uint256 id, uint256 epoch, string url, string selector, string name, bool repeat, address creator, uint256 credit, bool fulfilled, uint256 result, uint256 assetType)
func (_AssetManager *AssetManagerCallerSession) Jobs(arg0 *big.Int) (struct {
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
	return _AssetManager.Contract.Jobs(&_AssetManager.CallOpts, arg0)
}

// NumAssets is a free data retrieval call binding the contract method 0xa46fe83b.
//
// Solidity: function numAssets() view returns(uint256)
func (_AssetManager *AssetManagerCaller) NumAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AssetManager.contract.Call(opts, &out, "numAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumAssets is a free data retrieval call binding the contract method 0xa46fe83b.
//
// Solidity: function numAssets() view returns(uint256)
func (_AssetManager *AssetManagerSession) NumAssets() (*big.Int, error) {
	return _AssetManager.Contract.NumAssets(&_AssetManager.CallOpts)
}

// NumAssets is a free data retrieval call binding the contract method 0xa46fe83b.
//
// Solidity: function numAssets() view returns(uint256)
func (_AssetManager *AssetManagerCallerSession) NumAssets() (*big.Int, error) {
	return _AssetManager.Contract.NumAssets(&_AssetManager.CallOpts)
}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address)
func (_AssetManager *AssetManagerCaller) Parameters(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AssetManager.contract.Call(opts, &out, "parameters")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address)
func (_AssetManager *AssetManagerSession) Parameters() (common.Address, error) {
	return _AssetManager.Contract.Parameters(&_AssetManager.CallOpts)
}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address)
func (_AssetManager *AssetManagerCallerSession) Parameters() (common.Address, error) {
	return _AssetManager.Contract.Parameters(&_AssetManager.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AssetManager *AssetManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AssetManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AssetManager *AssetManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AssetManager.Contract.SupportsInterface(&_AssetManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AssetManager *AssetManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AssetManager.Contract.SupportsInterface(&_AssetManager.CallOpts, interfaceId)
}

// AddJobToCollection is a paid mutator transaction binding the contract method 0x2f9fcc0a.
//
// Solidity: function addJobToCollection(uint256 collectionID, uint256 jobID) returns()
func (_AssetManager *AssetManagerTransactor) AddJobToCollection(opts *bind.TransactOpts, collectionID *big.Int, jobID *big.Int) (*types.Transaction, error) {
	return _AssetManager.contract.Transact(opts, "addJobToCollection", collectionID, jobID)
}

// AddJobToCollection is a paid mutator transaction binding the contract method 0x2f9fcc0a.
//
// Solidity: function addJobToCollection(uint256 collectionID, uint256 jobID) returns()
func (_AssetManager *AssetManagerSession) AddJobToCollection(collectionID *big.Int, jobID *big.Int) (*types.Transaction, error) {
	return _AssetManager.Contract.AddJobToCollection(&_AssetManager.TransactOpts, collectionID, jobID)
}

// AddJobToCollection is a paid mutator transaction binding the contract method 0x2f9fcc0a.
//
// Solidity: function addJobToCollection(uint256 collectionID, uint256 jobID) returns()
func (_AssetManager *AssetManagerTransactorSession) AddJobToCollection(collectionID *big.Int, jobID *big.Int) (*types.Transaction, error) {
	return _AssetManager.Contract.AddJobToCollection(&_AssetManager.TransactOpts, collectionID, jobID)
}

// CreateCollection is a paid mutator transaction binding the contract method 0xeddfa070.
//
// Solidity: function createCollection(string name, uint256[] jobIDs, uint32 aggregationMethod) payable returns()
func (_AssetManager *AssetManagerTransactor) CreateCollection(opts *bind.TransactOpts, name string, jobIDs []*big.Int, aggregationMethod uint32) (*types.Transaction, error) {
	return _AssetManager.contract.Transact(opts, "createCollection", name, jobIDs, aggregationMethod)
}

// CreateCollection is a paid mutator transaction binding the contract method 0xeddfa070.
//
// Solidity: function createCollection(string name, uint256[] jobIDs, uint32 aggregationMethod) payable returns()
func (_AssetManager *AssetManagerSession) CreateCollection(name string, jobIDs []*big.Int, aggregationMethod uint32) (*types.Transaction, error) {
	return _AssetManager.Contract.CreateCollection(&_AssetManager.TransactOpts, name, jobIDs, aggregationMethod)
}

// CreateCollection is a paid mutator transaction binding the contract method 0xeddfa070.
//
// Solidity: function createCollection(string name, uint256[] jobIDs, uint32 aggregationMethod) payable returns()
func (_AssetManager *AssetManagerTransactorSession) CreateCollection(name string, jobIDs []*big.Int, aggregationMethod uint32) (*types.Transaction, error) {
	return _AssetManager.Contract.CreateCollection(&_AssetManager.TransactOpts, name, jobIDs, aggregationMethod)
}

// CreateJob is a paid mutator transaction binding the contract method 0x628aff1d.
//
// Solidity: function createJob(string url, string selector, string name, bool repeat) payable returns()
func (_AssetManager *AssetManagerTransactor) CreateJob(opts *bind.TransactOpts, url string, selector string, name string, repeat bool) (*types.Transaction, error) {
	return _AssetManager.contract.Transact(opts, "createJob", url, selector, name, repeat)
}

// CreateJob is a paid mutator transaction binding the contract method 0x628aff1d.
//
// Solidity: function createJob(string url, string selector, string name, bool repeat) payable returns()
func (_AssetManager *AssetManagerSession) CreateJob(url string, selector string, name string, repeat bool) (*types.Transaction, error) {
	return _AssetManager.Contract.CreateJob(&_AssetManager.TransactOpts, url, selector, name, repeat)
}

// CreateJob is a paid mutator transaction binding the contract method 0x628aff1d.
//
// Solidity: function createJob(string url, string selector, string name, bool repeat) payable returns()
func (_AssetManager *AssetManagerTransactorSession) CreateJob(url string, selector string, name string, repeat bool) (*types.Transaction, error) {
	return _AssetManager.Contract.CreateJob(&_AssetManager.TransactOpts, url, selector, name, repeat)
}

// FulfillAsset is a paid mutator transaction binding the contract method 0x51be9717.
//
// Solidity: function fulfillAsset(uint256 id, uint256 value) returns()
func (_AssetManager *AssetManagerTransactor) FulfillAsset(opts *bind.TransactOpts, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _AssetManager.contract.Transact(opts, "fulfillAsset", id, value)
}

// FulfillAsset is a paid mutator transaction binding the contract method 0x51be9717.
//
// Solidity: function fulfillAsset(uint256 id, uint256 value) returns()
func (_AssetManager *AssetManagerSession) FulfillAsset(id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _AssetManager.Contract.FulfillAsset(&_AssetManager.TransactOpts, id, value)
}

// FulfillAsset is a paid mutator transaction binding the contract method 0x51be9717.
//
// Solidity: function fulfillAsset(uint256 id, uint256 value) returns()
func (_AssetManager *AssetManagerTransactorSession) FulfillAsset(id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _AssetManager.Contract.FulfillAsset(&_AssetManager.TransactOpts, id, value)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AssetManager *AssetManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AssetManager *AssetManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetManager.Contract.GrantRole(&_AssetManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AssetManager *AssetManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetManager.Contract.GrantRole(&_AssetManager.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AssetManager *AssetManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetManager.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AssetManager *AssetManagerSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetManager.Contract.RenounceRole(&_AssetManager.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AssetManager *AssetManagerTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetManager.Contract.RenounceRole(&_AssetManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AssetManager *AssetManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AssetManager *AssetManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetManager.Contract.RevokeRole(&_AssetManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AssetManager *AssetManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetManager.Contract.RevokeRole(&_AssetManager.TransactOpts, role, account)
}

// AssetManagerCollectionCreatedIterator is returned from FilterCollectionCreated and is used to iterate over the raw logs and unpacked data for CollectionCreated events raised by the AssetManager contract.
type AssetManagerCollectionCreatedIterator struct {
	Event *AssetManagerCollectionCreated // Event containing the contract specifics and raw log

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
func (it *AssetManagerCollectionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetManagerCollectionCreated)
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
		it.Event = new(AssetManagerCollectionCreated)
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
func (it *AssetManagerCollectionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetManagerCollectionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetManagerCollectionCreated represents a CollectionCreated event raised by the AssetManager contract.
type AssetManagerCollectionCreated struct {
	Id                *big.Int
	Epoch             *big.Int
	Name              string
	AggregationMethod uint32
	JobIDs            []*big.Int
	Creator           common.Address
	Credit            *big.Int
	Timestamp         *big.Int
	AssetType         uint8
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCollectionCreated is a free log retrieval operation binding the contract event 0x62702e8a1ad6ee48a1ccbf5b7b8041a174af9b18604f29a2dff0c8394e8638a7.
//
// Solidity: event CollectionCreated(uint256 id, uint256 epoch, string name, uint32 aggregationMethod, uint256[] jobIDs, address creator, uint256 credit, uint256 timestamp, uint8 assetType)
func (_AssetManager *AssetManagerFilterer) FilterCollectionCreated(opts *bind.FilterOpts) (*AssetManagerCollectionCreatedIterator, error) {

	logs, sub, err := _AssetManager.contract.FilterLogs(opts, "CollectionCreated")
	if err != nil {
		return nil, err
	}
	return &AssetManagerCollectionCreatedIterator{contract: _AssetManager.contract, event: "CollectionCreated", logs: logs, sub: sub}, nil
}

// WatchCollectionCreated is a free log subscription operation binding the contract event 0x62702e8a1ad6ee48a1ccbf5b7b8041a174af9b18604f29a2dff0c8394e8638a7.
//
// Solidity: event CollectionCreated(uint256 id, uint256 epoch, string name, uint32 aggregationMethod, uint256[] jobIDs, address creator, uint256 credit, uint256 timestamp, uint8 assetType)
func (_AssetManager *AssetManagerFilterer) WatchCollectionCreated(opts *bind.WatchOpts, sink chan<- *AssetManagerCollectionCreated) (event.Subscription, error) {

	logs, sub, err := _AssetManager.contract.WatchLogs(opts, "CollectionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetManagerCollectionCreated)
				if err := _AssetManager.contract.UnpackLog(event, "CollectionCreated", log); err != nil {
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

// ParseCollectionCreated is a log parse operation binding the contract event 0x62702e8a1ad6ee48a1ccbf5b7b8041a174af9b18604f29a2dff0c8394e8638a7.
//
// Solidity: event CollectionCreated(uint256 id, uint256 epoch, string name, uint32 aggregationMethod, uint256[] jobIDs, address creator, uint256 credit, uint256 timestamp, uint8 assetType)
func (_AssetManager *AssetManagerFilterer) ParseCollectionCreated(log types.Log) (*AssetManagerCollectionCreated, error) {
	event := new(AssetManagerCollectionCreated)
	if err := _AssetManager.contract.UnpackLog(event, "CollectionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetManagerCollectionReportedIterator is returned from FilterCollectionReported and is used to iterate over the raw logs and unpacked data for CollectionReported events raised by the AssetManager contract.
type AssetManagerCollectionReportedIterator struct {
	Event *AssetManagerCollectionReported // Event containing the contract specifics and raw log

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
func (it *AssetManagerCollectionReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetManagerCollectionReported)
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
		it.Event = new(AssetManagerCollectionReported)
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
func (it *AssetManagerCollectionReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetManagerCollectionReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetManagerCollectionReported represents a CollectionReported event raised by the AssetManager contract.
type AssetManagerCollectionReported struct {
	Id                *big.Int
	Value             *big.Int
	Epoch             *big.Int
	Name              string
	AggregationMethod uint32
	JobIDs            []*big.Int
	Creator           common.Address
	Credit            *big.Int
	Timestamp         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCollectionReported is a free log retrieval operation binding the contract event 0x66c694a0ea197f0994021a7c6132b1f5fa93687f30065309d9ae09b1ade534fc.
//
// Solidity: event CollectionReported(uint256 id, uint256 value, uint256 epoch, string name, uint32 aggregationMethod, uint256[] jobIDs, address creator, uint256 credit, uint256 timestamp)
func (_AssetManager *AssetManagerFilterer) FilterCollectionReported(opts *bind.FilterOpts) (*AssetManagerCollectionReportedIterator, error) {

	logs, sub, err := _AssetManager.contract.FilterLogs(opts, "CollectionReported")
	if err != nil {
		return nil, err
	}
	return &AssetManagerCollectionReportedIterator{contract: _AssetManager.contract, event: "CollectionReported", logs: logs, sub: sub}, nil
}

// WatchCollectionReported is a free log subscription operation binding the contract event 0x66c694a0ea197f0994021a7c6132b1f5fa93687f30065309d9ae09b1ade534fc.
//
// Solidity: event CollectionReported(uint256 id, uint256 value, uint256 epoch, string name, uint32 aggregationMethod, uint256[] jobIDs, address creator, uint256 credit, uint256 timestamp)
func (_AssetManager *AssetManagerFilterer) WatchCollectionReported(opts *bind.WatchOpts, sink chan<- *AssetManagerCollectionReported) (event.Subscription, error) {

	logs, sub, err := _AssetManager.contract.WatchLogs(opts, "CollectionReported")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetManagerCollectionReported)
				if err := _AssetManager.contract.UnpackLog(event, "CollectionReported", log); err != nil {
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

// ParseCollectionReported is a log parse operation binding the contract event 0x66c694a0ea197f0994021a7c6132b1f5fa93687f30065309d9ae09b1ade534fc.
//
// Solidity: event CollectionReported(uint256 id, uint256 value, uint256 epoch, string name, uint32 aggregationMethod, uint256[] jobIDs, address creator, uint256 credit, uint256 timestamp)
func (_AssetManager *AssetManagerFilterer) ParseCollectionReported(log types.Log) (*AssetManagerCollectionReported, error) {
	event := new(AssetManagerCollectionReported)
	if err := _AssetManager.contract.UnpackLog(event, "CollectionReported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetManagerCollectionUpdatedIterator is returned from FilterCollectionUpdated and is used to iterate over the raw logs and unpacked data for CollectionUpdated events raised by the AssetManager contract.
type AssetManagerCollectionUpdatedIterator struct {
	Event *AssetManagerCollectionUpdated // Event containing the contract specifics and raw log

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
func (it *AssetManagerCollectionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetManagerCollectionUpdated)
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
		it.Event = new(AssetManagerCollectionUpdated)
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
func (it *AssetManagerCollectionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetManagerCollectionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetManagerCollectionUpdated represents a CollectionUpdated event raised by the AssetManager contract.
type AssetManagerCollectionUpdated struct {
	Id            *big.Int
	Epoch         *big.Int
	Name          string
	UpdatedJobIDs []*big.Int
	Timestamp     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCollectionUpdated is a free log retrieval operation binding the contract event 0x82e260cefb04ca78918ec0d223924a9c7893749eec0ccb5749ed7ada6145fdd5.
//
// Solidity: event CollectionUpdated(uint256 id, uint256 epoch, string name, uint256[] updatedJobIDs, uint256 timestamp)
func (_AssetManager *AssetManagerFilterer) FilterCollectionUpdated(opts *bind.FilterOpts) (*AssetManagerCollectionUpdatedIterator, error) {

	logs, sub, err := _AssetManager.contract.FilterLogs(opts, "CollectionUpdated")
	if err != nil {
		return nil, err
	}
	return &AssetManagerCollectionUpdatedIterator{contract: _AssetManager.contract, event: "CollectionUpdated", logs: logs, sub: sub}, nil
}

// WatchCollectionUpdated is a free log subscription operation binding the contract event 0x82e260cefb04ca78918ec0d223924a9c7893749eec0ccb5749ed7ada6145fdd5.
//
// Solidity: event CollectionUpdated(uint256 id, uint256 epoch, string name, uint256[] updatedJobIDs, uint256 timestamp)
func (_AssetManager *AssetManagerFilterer) WatchCollectionUpdated(opts *bind.WatchOpts, sink chan<- *AssetManagerCollectionUpdated) (event.Subscription, error) {

	logs, sub, err := _AssetManager.contract.WatchLogs(opts, "CollectionUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetManagerCollectionUpdated)
				if err := _AssetManager.contract.UnpackLog(event, "CollectionUpdated", log); err != nil {
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

// ParseCollectionUpdated is a log parse operation binding the contract event 0x82e260cefb04ca78918ec0d223924a9c7893749eec0ccb5749ed7ada6145fdd5.
//
// Solidity: event CollectionUpdated(uint256 id, uint256 epoch, string name, uint256[] updatedJobIDs, uint256 timestamp)
func (_AssetManager *AssetManagerFilterer) ParseCollectionUpdated(log types.Log) (*AssetManagerCollectionUpdated, error) {
	event := new(AssetManagerCollectionUpdated)
	if err := _AssetManager.contract.UnpackLog(event, "CollectionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetManagerJobCreatedIterator is returned from FilterJobCreated and is used to iterate over the raw logs and unpacked data for JobCreated events raised by the AssetManager contract.
type AssetManagerJobCreatedIterator struct {
	Event *AssetManagerJobCreated // Event containing the contract specifics and raw log

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
func (it *AssetManagerJobCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetManagerJobCreated)
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
		it.Event = new(AssetManagerJobCreated)
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
func (it *AssetManagerJobCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetManagerJobCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetManagerJobCreated represents a JobCreated event raised by the AssetManager contract.
type AssetManagerJobCreated struct {
	Id        *big.Int
	Epoch     *big.Int
	Url       string
	Selector  string
	Name      string
	Repeat    bool
	Creator   common.Address
	Credit    *big.Int
	Timestamp *big.Int
	AssetType uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterJobCreated is a free log retrieval operation binding the contract event 0x9aaa4f7b91a3f010f59eee49d2c35ac432427b72855a2f7d77f258ebd292ed83.
//
// Solidity: event JobCreated(uint256 id, uint256 epoch, string url, string selector, string name, bool repeat, address creator, uint256 credit, uint256 timestamp, uint8 assetType)
func (_AssetManager *AssetManagerFilterer) FilterJobCreated(opts *bind.FilterOpts) (*AssetManagerJobCreatedIterator, error) {

	logs, sub, err := _AssetManager.contract.FilterLogs(opts, "JobCreated")
	if err != nil {
		return nil, err
	}
	return &AssetManagerJobCreatedIterator{contract: _AssetManager.contract, event: "JobCreated", logs: logs, sub: sub}, nil
}

// WatchJobCreated is a free log subscription operation binding the contract event 0x9aaa4f7b91a3f010f59eee49d2c35ac432427b72855a2f7d77f258ebd292ed83.
//
// Solidity: event JobCreated(uint256 id, uint256 epoch, string url, string selector, string name, bool repeat, address creator, uint256 credit, uint256 timestamp, uint8 assetType)
func (_AssetManager *AssetManagerFilterer) WatchJobCreated(opts *bind.WatchOpts, sink chan<- *AssetManagerJobCreated) (event.Subscription, error) {

	logs, sub, err := _AssetManager.contract.WatchLogs(opts, "JobCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetManagerJobCreated)
				if err := _AssetManager.contract.UnpackLog(event, "JobCreated", log); err != nil {
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

// ParseJobCreated is a log parse operation binding the contract event 0x9aaa4f7b91a3f010f59eee49d2c35ac432427b72855a2f7d77f258ebd292ed83.
//
// Solidity: event JobCreated(uint256 id, uint256 epoch, string url, string selector, string name, bool repeat, address creator, uint256 credit, uint256 timestamp, uint8 assetType)
func (_AssetManager *AssetManagerFilterer) ParseJobCreated(log types.Log) (*AssetManagerJobCreated, error) {
	event := new(AssetManagerJobCreated)
	if err := _AssetManager.contract.UnpackLog(event, "JobCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetManagerJobReportedIterator is returned from FilterJobReported and is used to iterate over the raw logs and unpacked data for JobReported events raised by the AssetManager contract.
type AssetManagerJobReportedIterator struct {
	Event *AssetManagerJobReported // Event containing the contract specifics and raw log

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
func (it *AssetManagerJobReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetManagerJobReported)
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
		it.Event = new(AssetManagerJobReported)
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
func (it *AssetManagerJobReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetManagerJobReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetManagerJobReported represents a JobReported event raised by the AssetManager contract.
type AssetManagerJobReported struct {
	Id        *big.Int
	Value     *big.Int
	Epoch     *big.Int
	Url       string
	Selector  string
	Name      string
	Repeat    bool
	Creator   common.Address
	Credit    *big.Int
	Fulfilled bool
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterJobReported is a free log retrieval operation binding the contract event 0x9028bea5bfa7ed26c48df402d89085a995447dc8c1fb167cb92a3c7411b54480.
//
// Solidity: event JobReported(uint256 id, uint256 value, uint256 epoch, string url, string selector, string name, bool repeat, address creator, uint256 credit, bool fulfilled, uint256 timestamp)
func (_AssetManager *AssetManagerFilterer) FilterJobReported(opts *bind.FilterOpts) (*AssetManagerJobReportedIterator, error) {

	logs, sub, err := _AssetManager.contract.FilterLogs(opts, "JobReported")
	if err != nil {
		return nil, err
	}
	return &AssetManagerJobReportedIterator{contract: _AssetManager.contract, event: "JobReported", logs: logs, sub: sub}, nil
}

// WatchJobReported is a free log subscription operation binding the contract event 0x9028bea5bfa7ed26c48df402d89085a995447dc8c1fb167cb92a3c7411b54480.
//
// Solidity: event JobReported(uint256 id, uint256 value, uint256 epoch, string url, string selector, string name, bool repeat, address creator, uint256 credit, bool fulfilled, uint256 timestamp)
func (_AssetManager *AssetManagerFilterer) WatchJobReported(opts *bind.WatchOpts, sink chan<- *AssetManagerJobReported) (event.Subscription, error) {

	logs, sub, err := _AssetManager.contract.WatchLogs(opts, "JobReported")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetManagerJobReported)
				if err := _AssetManager.contract.UnpackLog(event, "JobReported", log); err != nil {
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

// ParseJobReported is a log parse operation binding the contract event 0x9028bea5bfa7ed26c48df402d89085a995447dc8c1fb167cb92a3c7411b54480.
//
// Solidity: event JobReported(uint256 id, uint256 value, uint256 epoch, string url, string selector, string name, bool repeat, address creator, uint256 credit, bool fulfilled, uint256 timestamp)
func (_AssetManager *AssetManagerFilterer) ParseJobReported(log types.Log) (*AssetManagerJobReported, error) {
	event := new(AssetManagerJobReported)
	if err := _AssetManager.contract.UnpackLog(event, "JobReported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AssetManager contract.
type AssetManagerRoleAdminChangedIterator struct {
	Event *AssetManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AssetManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetManagerRoleAdminChanged)
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
		it.Event = new(AssetManagerRoleAdminChanged)
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
func (it *AssetManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetManagerRoleAdminChanged represents a RoleAdminChanged event raised by the AssetManager contract.
type AssetManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AssetManager *AssetManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AssetManagerRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AssetManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AssetManagerRoleAdminChangedIterator{contract: _AssetManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AssetManager *AssetManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AssetManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AssetManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetManagerRoleAdminChanged)
				if err := _AssetManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AssetManager *AssetManagerFilterer) ParseRoleAdminChanged(log types.Log) (*AssetManagerRoleAdminChanged, error) {
	event := new(AssetManagerRoleAdminChanged)
	if err := _AssetManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AssetManager contract.
type AssetManagerRoleGrantedIterator struct {
	Event *AssetManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *AssetManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetManagerRoleGranted)
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
		it.Event = new(AssetManagerRoleGranted)
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
func (it *AssetManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetManagerRoleGranted represents a RoleGranted event raised by the AssetManager contract.
type AssetManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AssetManager *AssetManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AssetManagerRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AssetManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AssetManagerRoleGrantedIterator{contract: _AssetManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AssetManager *AssetManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AssetManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AssetManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetManagerRoleGranted)
				if err := _AssetManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AssetManager *AssetManagerFilterer) ParseRoleGranted(log types.Log) (*AssetManagerRoleGranted, error) {
	event := new(AssetManagerRoleGranted)
	if err := _AssetManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AssetManager contract.
type AssetManagerRoleRevokedIterator struct {
	Event *AssetManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AssetManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetManagerRoleRevoked)
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
		it.Event = new(AssetManagerRoleRevoked)
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
func (it *AssetManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetManagerRoleRevoked represents a RoleRevoked event raised by the AssetManager contract.
type AssetManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AssetManager *AssetManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AssetManagerRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AssetManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AssetManagerRoleRevokedIterator{contract: _AssetManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AssetManager *AssetManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AssetManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AssetManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetManagerRoleRevoked)
				if err := _AssetManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AssetManager *AssetManagerFilterer) ParseRoleRevoked(log types.Log) (*AssetManagerRoleRevoked, error) {
	event := new(AssetManagerRoleRevoked)
	if err := _AssetManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
