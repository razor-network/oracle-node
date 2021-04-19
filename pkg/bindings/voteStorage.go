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

// VoteStorageABI is the input ABI used to generate the binding from.
const VoteStorageABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"commitments\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"totalStakeRevealed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voteWeights\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"votes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VoteStorage is an auto generated Go binding around an Ethereum contract.
type VoteStorage struct {
	VoteStorageCaller     // Read-only binding to the contract
	VoteStorageTransactor // Write-only binding to the contract
	VoteStorageFilterer   // Log filterer for contract events
}

// VoteStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoteStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoteStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoteStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoteStorageSession struct {
	Contract     *VoteStorage      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoteStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoteStorageCallerSession struct {
	Contract *VoteStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// VoteStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoteStorageTransactorSession struct {
	Contract     *VoteStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// VoteStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoteStorageRaw struct {
	Contract *VoteStorage // Generic contract binding to access the raw methods on
}

// VoteStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoteStorageCallerRaw struct {
	Contract *VoteStorageCaller // Generic read-only contract binding to access the raw methods on
}

// VoteStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoteStorageTransactorRaw struct {
	Contract *VoteStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoteStorage creates a new instance of VoteStorage, bound to a specific deployed contract.
func NewVoteStorage(address common.Address, backend bind.ContractBackend) (*VoteStorage, error) {
	contract, err := bindVoteStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VoteStorage{VoteStorageCaller: VoteStorageCaller{contract: contract}, VoteStorageTransactor: VoteStorageTransactor{contract: contract}, VoteStorageFilterer: VoteStorageFilterer{contract: contract}}, nil
}

// NewVoteStorageCaller creates a new read-only instance of VoteStorage, bound to a specific deployed contract.
func NewVoteStorageCaller(address common.Address, caller bind.ContractCaller) (*VoteStorageCaller, error) {
	contract, err := bindVoteStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoteStorageCaller{contract: contract}, nil
}

// NewVoteStorageTransactor creates a new write-only instance of VoteStorage, bound to a specific deployed contract.
func NewVoteStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*VoteStorageTransactor, error) {
	contract, err := bindVoteStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoteStorageTransactor{contract: contract}, nil
}

// NewVoteStorageFilterer creates a new log filterer instance of VoteStorage, bound to a specific deployed contract.
func NewVoteStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*VoteStorageFilterer, error) {
	contract, err := bindVoteStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoteStorageFilterer{contract: contract}, nil
}

// bindVoteStorage binds a generic wrapper to an already deployed contract.
func bindVoteStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VoteStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VoteStorage *VoteStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VoteStorage.Contract.VoteStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VoteStorage *VoteStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoteStorage.Contract.VoteStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VoteStorage *VoteStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VoteStorage.Contract.VoteStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VoteStorage *VoteStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VoteStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VoteStorage *VoteStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoteStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VoteStorage *VoteStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VoteStorage.Contract.contract.Transact(opts, method, params...)
}

// Commitments is a free data retrieval call binding the contract method 0xd13e2e60.
//
// Solidity: function commitments(uint256 , uint256 ) view returns(bytes32)
func (_VoteStorage *VoteStorageCaller) Commitments(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _VoteStorage.contract.Call(opts, &out, "commitments", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Commitments is a free data retrieval call binding the contract method 0xd13e2e60.
//
// Solidity: function commitments(uint256 , uint256 ) view returns(bytes32)
func (_VoteStorage *VoteStorageSession) Commitments(arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	return _VoteStorage.Contract.Commitments(&_VoteStorage.CallOpts, arg0, arg1)
}

// Commitments is a free data retrieval call binding the contract method 0xd13e2e60.
//
// Solidity: function commitments(uint256 , uint256 ) view returns(bytes32)
func (_VoteStorage *VoteStorageCallerSession) Commitments(arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	return _VoteStorage.Contract.Commitments(&_VoteStorage.CallOpts, arg0, arg1)
}

// TotalStakeRevealed is a free data retrieval call binding the contract method 0x8a757ecc.
//
// Solidity: function totalStakeRevealed(uint256 , uint256 ) view returns(uint256)
func (_VoteStorage *VoteStorageCaller) TotalStakeRevealed(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VoteStorage.contract.Call(opts, &out, "totalStakeRevealed", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStakeRevealed is a free data retrieval call binding the contract method 0x8a757ecc.
//
// Solidity: function totalStakeRevealed(uint256 , uint256 ) view returns(uint256)
func (_VoteStorage *VoteStorageSession) TotalStakeRevealed(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _VoteStorage.Contract.TotalStakeRevealed(&_VoteStorage.CallOpts, arg0, arg1)
}

// TotalStakeRevealed is a free data retrieval call binding the contract method 0x8a757ecc.
//
// Solidity: function totalStakeRevealed(uint256 , uint256 ) view returns(uint256)
func (_VoteStorage *VoteStorageCallerSession) TotalStakeRevealed(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _VoteStorage.Contract.TotalStakeRevealed(&_VoteStorage.CallOpts, arg0, arg1)
}

// VoteWeights is a free data retrieval call binding the contract method 0x8fd5ff00.
//
// Solidity: function voteWeights(uint256 , uint256 , uint256 ) view returns(uint256)
func (_VoteStorage *VoteStorageCaller) VoteWeights(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VoteStorage.contract.Call(opts, &out, "voteWeights", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoteWeights is a free data retrieval call binding the contract method 0x8fd5ff00.
//
// Solidity: function voteWeights(uint256 , uint256 , uint256 ) view returns(uint256)
func (_VoteStorage *VoteStorageSession) VoteWeights(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (*big.Int, error) {
	return _VoteStorage.Contract.VoteWeights(&_VoteStorage.CallOpts, arg0, arg1, arg2)
}

// VoteWeights is a free data retrieval call binding the contract method 0x8fd5ff00.
//
// Solidity: function voteWeights(uint256 , uint256 , uint256 ) view returns(uint256)
func (_VoteStorage *VoteStorageCallerSession) VoteWeights(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (*big.Int, error) {
	return _VoteStorage.Contract.VoteWeights(&_VoteStorage.CallOpts, arg0, arg1, arg2)
}

// Votes is a free data retrieval call binding the contract method 0x283e2905.
//
// Solidity: function votes(uint256 , uint256 , uint256 ) view returns(uint256 value, uint256 weight)
func (_VoteStorage *VoteStorageCaller) Votes(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (struct {
	Value  *big.Int
	Weight *big.Int
}, error) {
	var out []interface{}
	err := _VoteStorage.contract.Call(opts, &out, "votes", arg0, arg1, arg2)

	outstruct := new(struct {
		Value  *big.Int
		Weight *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Weight = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Votes is a free data retrieval call binding the contract method 0x283e2905.
//
// Solidity: function votes(uint256 , uint256 , uint256 ) view returns(uint256 value, uint256 weight)
func (_VoteStorage *VoteStorageSession) Votes(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (struct {
	Value  *big.Int
	Weight *big.Int
}, error) {
	return _VoteStorage.Contract.Votes(&_VoteStorage.CallOpts, arg0, arg1, arg2)
}

// Votes is a free data retrieval call binding the contract method 0x283e2905.
//
// Solidity: function votes(uint256 , uint256 , uint256 ) view returns(uint256 value, uint256 weight)
func (_VoteStorage *VoteStorageCallerSession) Votes(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (struct {
	Value  *big.Int
	Weight *big.Int
}, error) {
	return _VoteStorage.Contract.Votes(&_VoteStorage.CallOpts, arg0, arg1, arg2)
}
