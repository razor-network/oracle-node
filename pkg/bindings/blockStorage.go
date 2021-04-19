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

// BlockStorageABI is the input ABI used to generate the binding from.
const BlockStorageABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"blocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"proposerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iteration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"biggestStake\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"disputes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"accWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"median\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lowerCutoff\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"higherCutoff\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastVisited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposedBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"proposerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iteration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"biggestStake\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BlockStorage is an auto generated Go binding around an Ethereum contract.
type BlockStorage struct {
	BlockStorageCaller     // Read-only binding to the contract
	BlockStorageTransactor // Write-only binding to the contract
	BlockStorageFilterer   // Log filterer for contract events
}

// BlockStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlockStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockStorageSession struct {
	Contract     *BlockStorage     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlockStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockStorageCallerSession struct {
	Contract *BlockStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BlockStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockStorageTransactorSession struct {
	Contract     *BlockStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BlockStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlockStorageRaw struct {
	Contract *BlockStorage // Generic contract binding to access the raw methods on
}

// BlockStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockStorageCallerRaw struct {
	Contract *BlockStorageCaller // Generic read-only contract binding to access the raw methods on
}

// BlockStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockStorageTransactorRaw struct {
	Contract *BlockStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockStorage creates a new instance of BlockStorage, bound to a specific deployed contract.
func NewBlockStorage(address common.Address, backend bind.ContractBackend) (*BlockStorage, error) {
	contract, err := bindBlockStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlockStorage{BlockStorageCaller: BlockStorageCaller{contract: contract}, BlockStorageTransactor: BlockStorageTransactor{contract: contract}, BlockStorageFilterer: BlockStorageFilterer{contract: contract}}, nil
}

// NewBlockStorageCaller creates a new read-only instance of BlockStorage, bound to a specific deployed contract.
func NewBlockStorageCaller(address common.Address, caller bind.ContractCaller) (*BlockStorageCaller, error) {
	contract, err := bindBlockStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockStorageCaller{contract: contract}, nil
}

// NewBlockStorageTransactor creates a new write-only instance of BlockStorage, bound to a specific deployed contract.
func NewBlockStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockStorageTransactor, error) {
	contract, err := bindBlockStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockStorageTransactor{contract: contract}, nil
}

// NewBlockStorageFilterer creates a new log filterer instance of BlockStorage, bound to a specific deployed contract.
func NewBlockStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockStorageFilterer, error) {
	contract, err := bindBlockStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockStorageFilterer{contract: contract}, nil
}

// bindBlockStorage binds a generic wrapper to an already deployed contract.
func bindBlockStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlockStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockStorage *BlockStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockStorage.Contract.BlockStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockStorage *BlockStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockStorage.Contract.BlockStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockStorage *BlockStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockStorage.Contract.BlockStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockStorage *BlockStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockStorage *BlockStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockStorage *BlockStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockStorage.Contract.contract.Transact(opts, method, params...)
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks(uint256 ) view returns(uint256 proposerId, uint256 iteration, uint256 biggestStake, bool valid)
func (_BlockStorage *BlockStorageCaller) Blocks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ProposerId   *big.Int
	Iteration    *big.Int
	BiggestStake *big.Int
	Valid        bool
}, error) {
	var out []interface{}
	err := _BlockStorage.contract.Call(opts, &out, "blocks", arg0)

	outstruct := new(struct {
		ProposerId   *big.Int
		Iteration    *big.Int
		BiggestStake *big.Int
		Valid        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProposerId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Iteration = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BiggestStake = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Valid = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks(uint256 ) view returns(uint256 proposerId, uint256 iteration, uint256 biggestStake, bool valid)
func (_BlockStorage *BlockStorageSession) Blocks(arg0 *big.Int) (struct {
	ProposerId   *big.Int
	Iteration    *big.Int
	BiggestStake *big.Int
	Valid        bool
}, error) {
	return _BlockStorage.Contract.Blocks(&_BlockStorage.CallOpts, arg0)
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks(uint256 ) view returns(uint256 proposerId, uint256 iteration, uint256 biggestStake, bool valid)
func (_BlockStorage *BlockStorageCallerSession) Blocks(arg0 *big.Int) (struct {
	ProposerId   *big.Int
	Iteration    *big.Int
	BiggestStake *big.Int
	Valid        bool
}, error) {
	return _BlockStorage.Contract.Blocks(&_BlockStorage.CallOpts, arg0)
}

// Disputes is a free data retrieval call binding the contract method 0x828496d6.
//
// Solidity: function disputes(uint256 , address ) view returns(uint256 accWeight, uint256 median, uint256 lowerCutoff, uint256 higherCutoff, uint256 lastVisited, uint256 assetId)
func (_BlockStorage *BlockStorageCaller) Disputes(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	AccWeight    *big.Int
	Median       *big.Int
	LowerCutoff  *big.Int
	HigherCutoff *big.Int
	LastVisited  *big.Int
	AssetId      *big.Int
}, error) {
	var out []interface{}
	err := _BlockStorage.contract.Call(opts, &out, "disputes", arg0, arg1)

	outstruct := new(struct {
		AccWeight    *big.Int
		Median       *big.Int
		LowerCutoff  *big.Int
		HigherCutoff *big.Int
		LastVisited  *big.Int
		AssetId      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AccWeight = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Median = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LowerCutoff = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.HigherCutoff = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LastVisited = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.AssetId = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Disputes is a free data retrieval call binding the contract method 0x828496d6.
//
// Solidity: function disputes(uint256 , address ) view returns(uint256 accWeight, uint256 median, uint256 lowerCutoff, uint256 higherCutoff, uint256 lastVisited, uint256 assetId)
func (_BlockStorage *BlockStorageSession) Disputes(arg0 *big.Int, arg1 common.Address) (struct {
	AccWeight    *big.Int
	Median       *big.Int
	LowerCutoff  *big.Int
	HigherCutoff *big.Int
	LastVisited  *big.Int
	AssetId      *big.Int
}, error) {
	return _BlockStorage.Contract.Disputes(&_BlockStorage.CallOpts, arg0, arg1)
}

// Disputes is a free data retrieval call binding the contract method 0x828496d6.
//
// Solidity: function disputes(uint256 , address ) view returns(uint256 accWeight, uint256 median, uint256 lowerCutoff, uint256 higherCutoff, uint256 lastVisited, uint256 assetId)
func (_BlockStorage *BlockStorageCallerSession) Disputes(arg0 *big.Int, arg1 common.Address) (struct {
	AccWeight    *big.Int
	Median       *big.Int
	LowerCutoff  *big.Int
	HigherCutoff *big.Int
	LastVisited  *big.Int
	AssetId      *big.Int
}, error) {
	return _BlockStorage.Contract.Disputes(&_BlockStorage.CallOpts, arg0, arg1)
}

// ProposedBlocks is a free data retrieval call binding the contract method 0x92b48411.
//
// Solidity: function proposedBlocks(uint256 , uint256 ) view returns(uint256 proposerId, uint256 iteration, uint256 biggestStake, bool valid)
func (_BlockStorage *BlockStorageCaller) ProposedBlocks(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	ProposerId   *big.Int
	Iteration    *big.Int
	BiggestStake *big.Int
	Valid        bool
}, error) {
	var out []interface{}
	err := _BlockStorage.contract.Call(opts, &out, "proposedBlocks", arg0, arg1)

	outstruct := new(struct {
		ProposerId   *big.Int
		Iteration    *big.Int
		BiggestStake *big.Int
		Valid        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProposerId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Iteration = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BiggestStake = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Valid = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// ProposedBlocks is a free data retrieval call binding the contract method 0x92b48411.
//
// Solidity: function proposedBlocks(uint256 , uint256 ) view returns(uint256 proposerId, uint256 iteration, uint256 biggestStake, bool valid)
func (_BlockStorage *BlockStorageSession) ProposedBlocks(arg0 *big.Int, arg1 *big.Int) (struct {
	ProposerId   *big.Int
	Iteration    *big.Int
	BiggestStake *big.Int
	Valid        bool
}, error) {
	return _BlockStorage.Contract.ProposedBlocks(&_BlockStorage.CallOpts, arg0, arg1)
}

// ProposedBlocks is a free data retrieval call binding the contract method 0x92b48411.
//
// Solidity: function proposedBlocks(uint256 , uint256 ) view returns(uint256 proposerId, uint256 iteration, uint256 biggestStake, bool valid)
func (_BlockStorage *BlockStorageCallerSession) ProposedBlocks(arg0 *big.Int, arg1 *big.Int) (struct {
	ProposerId   *big.Int
	Iteration    *big.Int
	BiggestStake *big.Int
	Valid        bool
}, error) {
	return _BlockStorage.Contract.ProposedBlocks(&_BlockStorage.CallOpts, arg0, arg1)
}
