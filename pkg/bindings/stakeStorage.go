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

// StakeStorageABI is the input ABI used to generate the binding from.
const StakeStorageABI = "[{\"inputs\":[],\"name\":\"blockReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStakers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeGettingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakerIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochStaked\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochLastCommitted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochLastRevealed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeAfter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAfter\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// StakeStorage is an auto generated Go binding around an Ethereum contract.
type StakeStorage struct {
	StakeStorageCaller     // Read-only binding to the contract
	StakeStorageTransactor // Write-only binding to the contract
	StakeStorageFilterer   // Log filterer for contract events
}

// StakeStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakeStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakeStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakeStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakeStorageSession struct {
	Contract     *StakeStorage     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakeStorageCallerSession struct {
	Contract *StakeStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StakeStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakeStorageTransactorSession struct {
	Contract     *StakeStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StakeStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakeStorageRaw struct {
	Contract *StakeStorage // Generic contract binding to access the raw methods on
}

// StakeStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakeStorageCallerRaw struct {
	Contract *StakeStorageCaller // Generic read-only contract binding to access the raw methods on
}

// StakeStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakeStorageTransactorRaw struct {
	Contract *StakeStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakeStorage creates a new instance of StakeStorage, bound to a specific deployed contract.
func NewStakeStorage(address common.Address, backend bind.ContractBackend) (*StakeStorage, error) {
	contract, err := bindStakeStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakeStorage{StakeStorageCaller: StakeStorageCaller{contract: contract}, StakeStorageTransactor: StakeStorageTransactor{contract: contract}, StakeStorageFilterer: StakeStorageFilterer{contract: contract}}, nil
}

// NewStakeStorageCaller creates a new read-only instance of StakeStorage, bound to a specific deployed contract.
func NewStakeStorageCaller(address common.Address, caller bind.ContractCaller) (*StakeStorageCaller, error) {
	contract, err := bindStakeStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakeStorageCaller{contract: contract}, nil
}

// NewStakeStorageTransactor creates a new write-only instance of StakeStorage, bound to a specific deployed contract.
func NewStakeStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StakeStorageTransactor, error) {
	contract, err := bindStakeStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakeStorageTransactor{contract: contract}, nil
}

// NewStakeStorageFilterer creates a new log filterer instance of StakeStorage, bound to a specific deployed contract.
func NewStakeStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StakeStorageFilterer, error) {
	contract, err := bindStakeStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakeStorageFilterer{contract: contract}, nil
}

// bindStakeStorage binds a generic wrapper to an already deployed contract.
func bindStakeStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakeStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeStorage *StakeStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeStorage.Contract.StakeStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeStorage *StakeStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeStorage.Contract.StakeStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeStorage *StakeStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeStorage.Contract.StakeStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeStorage *StakeStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeStorage *StakeStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeStorage *StakeStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeStorage.Contract.contract.Transact(opts, method, params...)
}

// BlockReward is a free data retrieval call binding the contract method 0x0ac168a1.
//
// Solidity: function blockReward() view returns(uint256)
func (_StakeStorage *StakeStorageCaller) BlockReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeStorage.contract.Call(opts, &out, "blockReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockReward is a free data retrieval call binding the contract method 0x0ac168a1.
//
// Solidity: function blockReward() view returns(uint256)
func (_StakeStorage *StakeStorageSession) BlockReward() (*big.Int, error) {
	return _StakeStorage.Contract.BlockReward(&_StakeStorage.CallOpts)
}

// BlockReward is a free data retrieval call binding the contract method 0x0ac168a1.
//
// Solidity: function blockReward() view returns(uint256)
func (_StakeStorage *StakeStorageCallerSession) BlockReward() (*big.Int, error) {
	return _StakeStorage.Contract.BlockReward(&_StakeStorage.CallOpts)
}

// NumStakers is a free data retrieval call binding the contract method 0x6c8b052a.
//
// Solidity: function numStakers() view returns(uint256)
func (_StakeStorage *StakeStorageCaller) NumStakers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeStorage.contract.Call(opts, &out, "numStakers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumStakers is a free data retrieval call binding the contract method 0x6c8b052a.
//
// Solidity: function numStakers() view returns(uint256)
func (_StakeStorage *StakeStorageSession) NumStakers() (*big.Int, error) {
	return _StakeStorage.Contract.NumStakers(&_StakeStorage.CallOpts)
}

// NumStakers is a free data retrieval call binding the contract method 0x6c8b052a.
//
// Solidity: function numStakers() view returns(uint256)
func (_StakeStorage *StakeStorageCallerSession) NumStakers() (*big.Int, error) {
	return _StakeStorage.Contract.NumStakers(&_StakeStorage.CallOpts)
}

// RewardPool is a free data retrieval call binding the contract method 0x66666aa9.
//
// Solidity: function rewardPool() view returns(uint256)
func (_StakeStorage *StakeStorageCaller) RewardPool(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeStorage.contract.Call(opts, &out, "rewardPool")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPool is a free data retrieval call binding the contract method 0x66666aa9.
//
// Solidity: function rewardPool() view returns(uint256)
func (_StakeStorage *StakeStorageSession) RewardPool() (*big.Int, error) {
	return _StakeStorage.Contract.RewardPool(&_StakeStorage.CallOpts)
}

// RewardPool is a free data retrieval call binding the contract method 0x66666aa9.
//
// Solidity: function rewardPool() view returns(uint256)
func (_StakeStorage *StakeStorageCallerSession) RewardPool() (*big.Int, error) {
	return _StakeStorage.Contract.RewardPool(&_StakeStorage.CallOpts)
}

// StakeGettingReward is a free data retrieval call binding the contract method 0x0ec88d3f.
//
// Solidity: function stakeGettingReward() view returns(uint256)
func (_StakeStorage *StakeStorageCaller) StakeGettingReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeStorage.contract.Call(opts, &out, "stakeGettingReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeGettingReward is a free data retrieval call binding the contract method 0x0ec88d3f.
//
// Solidity: function stakeGettingReward() view returns(uint256)
func (_StakeStorage *StakeStorageSession) StakeGettingReward() (*big.Int, error) {
	return _StakeStorage.Contract.StakeGettingReward(&_StakeStorage.CallOpts)
}

// StakeGettingReward is a free data retrieval call binding the contract method 0x0ec88d3f.
//
// Solidity: function stakeGettingReward() view returns(uint256)
func (_StakeStorage *StakeStorageCallerSession) StakeGettingReward() (*big.Int, error) {
	return _StakeStorage.Contract.StakeGettingReward(&_StakeStorage.CallOpts)
}

// StakerIds is a free data retrieval call binding the contract method 0xc8ae0d7d.
//
// Solidity: function stakerIds(address ) view returns(uint256)
func (_StakeStorage *StakeStorageCaller) StakerIds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeStorage.contract.Call(opts, &out, "stakerIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerIds is a free data retrieval call binding the contract method 0xc8ae0d7d.
//
// Solidity: function stakerIds(address ) view returns(uint256)
func (_StakeStorage *StakeStorageSession) StakerIds(arg0 common.Address) (*big.Int, error) {
	return _StakeStorage.Contract.StakerIds(&_StakeStorage.CallOpts, arg0)
}

// StakerIds is a free data retrieval call binding the contract method 0xc8ae0d7d.
//
// Solidity: function stakerIds(address ) view returns(uint256)
func (_StakeStorage *StakeStorageCallerSession) StakerIds(arg0 common.Address) (*big.Int, error) {
	return _StakeStorage.Contract.StakerIds(&_StakeStorage.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) view returns(uint256 id, address _address, uint256 stake, uint256 epochStaked, uint256 epochLastCommitted, uint256 epochLastRevealed, uint256 unstakeAfter, uint256 withdrawAfter)
func (_StakeStorage *StakeStorageCaller) Stakers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id                 *big.Int
	Address            common.Address
	Stake              *big.Int
	EpochStaked        *big.Int
	EpochLastCommitted *big.Int
	EpochLastRevealed  *big.Int
	UnstakeAfter       *big.Int
	WithdrawAfter      *big.Int
}, error) {
	var out []interface{}
	err := _StakeStorage.contract.Call(opts, &out, "stakers", arg0)

	outstruct := new(struct {
		Id                 *big.Int
		Address            common.Address
		Stake              *big.Int
		EpochStaked        *big.Int
		EpochLastCommitted *big.Int
		EpochLastRevealed  *big.Int
		UnstakeAfter       *big.Int
		WithdrawAfter      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Address = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Stake = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EpochStaked = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.EpochLastCommitted = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.EpochLastRevealed = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.UnstakeAfter = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.WithdrawAfter = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) view returns(uint256 id, address _address, uint256 stake, uint256 epochStaked, uint256 epochLastCommitted, uint256 epochLastRevealed, uint256 unstakeAfter, uint256 withdrawAfter)
func (_StakeStorage *StakeStorageSession) Stakers(arg0 *big.Int) (struct {
	Id                 *big.Int
	Address            common.Address
	Stake              *big.Int
	EpochStaked        *big.Int
	EpochLastCommitted *big.Int
	EpochLastRevealed  *big.Int
	UnstakeAfter       *big.Int
	WithdrawAfter      *big.Int
}, error) {
	return _StakeStorage.Contract.Stakers(&_StakeStorage.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) view returns(uint256 id, address _address, uint256 stake, uint256 epochStaked, uint256 epochLastCommitted, uint256 epochLastRevealed, uint256 unstakeAfter, uint256 withdrawAfter)
func (_StakeStorage *StakeStorageCallerSession) Stakers(arg0 *big.Int) (struct {
	Id                 *big.Int
	Address            common.Address
	Stake              *big.Int
	EpochStaked        *big.Int
	EpochLastCommitted *big.Int
	EpochLastRevealed  *big.Int
	UnstakeAfter       *big.Int
	WithdrawAfter      *big.Int
}, error) {
	return _StakeStorage.Contract.Stakers(&_StakeStorage.CallOpts, arg0)
}
