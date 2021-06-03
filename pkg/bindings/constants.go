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

// ConstantsABI is the input ABI used to generate the binding from.
const ConstantsABI = "[{\"inputs\":[],\"name\":\"commit\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dispute\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exposureDenominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockConfirmerHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDefaultAdminHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getJobConfirmerHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakeModifierHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakerActivityUpdaterHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxAltBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"penaltyNotRevealDenom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"penaltyNotRevealNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"propose\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reveal\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unstakeLockPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawLockPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// Constants is an auto generated Go binding around an Ethereum contract.
type Constants struct {
	ConstantsCaller     // Read-only binding to the contract
	ConstantsTransactor // Write-only binding to the contract
	ConstantsFilterer   // Log filterer for contract events
}

// ConstantsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConstantsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstantsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConstantsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstantsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConstantsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstantsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConstantsSession struct {
	Contract     *Constants        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConstantsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConstantsCallerSession struct {
	Contract *ConstantsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ConstantsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConstantsTransactorSession struct {
	Contract     *ConstantsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ConstantsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConstantsRaw struct {
	Contract *Constants // Generic contract binding to access the raw methods on
}

// ConstantsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConstantsCallerRaw struct {
	Contract *ConstantsCaller // Generic read-only contract binding to access the raw methods on
}

// ConstantsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConstantsTransactorRaw struct {
	Contract *ConstantsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConstants creates a new instance of Constants, bound to a specific deployed contract.
func NewConstants(address common.Address, backend bind.ContractBackend) (*Constants, error) {
	contract, err := bindConstants(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Constants{ConstantsCaller: ConstantsCaller{contract: contract}, ConstantsTransactor: ConstantsTransactor{contract: contract}, ConstantsFilterer: ConstantsFilterer{contract: contract}}, nil
}

// NewConstantsCaller creates a new read-only instance of Constants, bound to a specific deployed contract.
func NewConstantsCaller(address common.Address, caller bind.ContractCaller) (*ConstantsCaller, error) {
	contract, err := bindConstants(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConstantsCaller{contract: contract}, nil
}

// NewConstantsTransactor creates a new write-only instance of Constants, bound to a specific deployed contract.
func NewConstantsTransactor(address common.Address, transactor bind.ContractTransactor) (*ConstantsTransactor, error) {
	contract, err := bindConstants(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConstantsTransactor{contract: contract}, nil
}

// NewConstantsFilterer creates a new log filterer instance of Constants, bound to a specific deployed contract.
func NewConstantsFilterer(address common.Address, filterer bind.ContractFilterer) (*ConstantsFilterer, error) {
	contract, err := bindConstants(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConstantsFilterer{contract: contract}, nil
}

// bindConstants binds a generic wrapper to an already deployed contract.
func bindConstants(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConstantsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Constants *ConstantsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Constants.Contract.ConstantsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Constants *ConstantsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Constants.Contract.ConstantsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Constants *ConstantsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Constants.Contract.ConstantsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Constants *ConstantsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Constants.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Constants *ConstantsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Constants.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Constants *ConstantsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Constants.Contract.contract.Transact(opts, method, params...)
}

// Commit is a free data retrieval call binding the contract method 0x3c7a3aff.
//
// Solidity: function commit() pure returns(uint8)
func (_Constants *ConstantsCaller) Commit(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Constants.contract.Call(opts, &out, "commit")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Commit is a free data retrieval call binding the contract method 0x3c7a3aff.
//
// Solidity: function commit() pure returns(uint8)
func (_Constants *ConstantsSession) Commit() (uint8, error) {
	return _Constants.Contract.Commit(&_Constants.CallOpts)
}

// Commit is a free data retrieval call binding the contract method 0x3c7a3aff.
//
// Solidity: function commit() pure returns(uint8)
func (_Constants *ConstantsCallerSession) Commit() (uint8, error) {
	return _Constants.Contract.Commit(&_Constants.CallOpts)
}

// Dispute is a free data retrieval call binding the contract method 0xf240f7c3.
//
// Solidity: function dispute() pure returns(uint8)
func (_Constants *ConstantsCaller) Dispute(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Constants.contract.Call(opts, &out, "dispute")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Dispute is a free data retrieval call binding the contract method 0xf240f7c3.
//
// Solidity: function dispute() pure returns(uint8)
func (_Constants *ConstantsSession) Dispute() (uint8, error) {
	return _Constants.Contract.Dispute(&_Constants.CallOpts)
}

// Dispute is a free data retrieval call binding the contract method 0xf240f7c3.
//
// Solidity: function dispute() pure returns(uint8)
func (_Constants *ConstantsCallerSession) Dispute() (uint8, error) {
	return _Constants.Contract.Dispute(&_Constants.CallOpts)
}

// EpochLength is a free data retrieval call binding the contract method 0x57d775f8.
//
// Solidity: function epochLength() pure returns(uint256)
func (_Constants *ConstantsCaller) EpochLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Constants.contract.Call(opts, &out, "epochLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochLength is a free data retrieval call binding the contract method 0x57d775f8.
//
// Solidity: function epochLength() pure returns(uint256)
func (_Constants *ConstantsSession) EpochLength() (*big.Int, error) {
	return _Constants.Contract.EpochLength(&_Constants.CallOpts)
}

// EpochLength is a free data retrieval call binding the contract method 0x57d775f8.
//
// Solidity: function epochLength() pure returns(uint256)
func (_Constants *ConstantsCallerSession) EpochLength() (*big.Int, error) {
	return _Constants.Contract.EpochLength(&_Constants.CallOpts)
}

// ExposureDenominator is a free data retrieval call binding the contract method 0x3002c9ac.
//
// Solidity: function exposureDenominator() pure returns(uint256)
func (_Constants *ConstantsCaller) ExposureDenominator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Constants.contract.Call(opts, &out, "exposureDenominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExposureDenominator is a free data retrieval call binding the contract method 0x3002c9ac.
//
// Solidity: function exposureDenominator() pure returns(uint256)
func (_Constants *ConstantsSession) ExposureDenominator() (*big.Int, error) {
	return _Constants.Contract.ExposureDenominator(&_Constants.CallOpts)
}

// ExposureDenominator is a free data retrieval call binding the contract method 0x3002c9ac.
//
// Solidity: function exposureDenominator() pure returns(uint256)
func (_Constants *ConstantsCallerSession) ExposureDenominator() (*big.Int, error) {
	return _Constants.Contract.ExposureDenominator(&_Constants.CallOpts)
}

// GetBlockConfirmerHash is a free data retrieval call binding the contract method 0xd98b3ced.
//
// Solidity: function getBlockConfirmerHash() pure returns(bytes32)
func (_Constants *ConstantsCaller) GetBlockConfirmerHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Constants.contract.Call(opts, &out, "getBlockConfirmerHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockConfirmerHash is a free data retrieval call binding the contract method 0xd98b3ced.
//
// Solidity: function getBlockConfirmerHash() pure returns(bytes32)
func (_Constants *ConstantsSession) GetBlockConfirmerHash() ([32]byte, error) {
	return _Constants.Contract.GetBlockConfirmerHash(&_Constants.CallOpts)
}

// GetBlockConfirmerHash is a free data retrieval call binding the contract method 0xd98b3ced.
//
// Solidity: function getBlockConfirmerHash() pure returns(bytes32)
func (_Constants *ConstantsCallerSession) GetBlockConfirmerHash() ([32]byte, error) {
	return _Constants.Contract.GetBlockConfirmerHash(&_Constants.CallOpts)
}

// GetDefaultAdminHash is a free data retrieval call binding the contract method 0xb1a62781.
//
// Solidity: function getDefaultAdminHash() pure returns(bytes32)
func (_Constants *ConstantsCaller) GetDefaultAdminHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Constants.contract.Call(opts, &out, "getDefaultAdminHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDefaultAdminHash is a free data retrieval call binding the contract method 0xb1a62781.
//
// Solidity: function getDefaultAdminHash() pure returns(bytes32)
func (_Constants *ConstantsSession) GetDefaultAdminHash() ([32]byte, error) {
	return _Constants.Contract.GetDefaultAdminHash(&_Constants.CallOpts)
}

// GetDefaultAdminHash is a free data retrieval call binding the contract method 0xb1a62781.
//
// Solidity: function getDefaultAdminHash() pure returns(bytes32)
func (_Constants *ConstantsCallerSession) GetDefaultAdminHash() ([32]byte, error) {
	return _Constants.Contract.GetDefaultAdminHash(&_Constants.CallOpts)
}

// GetJobConfirmerHash is a free data retrieval call binding the contract method 0x7f890b11.
//
// Solidity: function getJobConfirmerHash() pure returns(bytes32)
func (_Constants *ConstantsCaller) GetJobConfirmerHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Constants.contract.Call(opts, &out, "getJobConfirmerHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetJobConfirmerHash is a free data retrieval call binding the contract method 0x7f890b11.
//
// Solidity: function getJobConfirmerHash() pure returns(bytes32)
func (_Constants *ConstantsSession) GetJobConfirmerHash() ([32]byte, error) {
	return _Constants.Contract.GetJobConfirmerHash(&_Constants.CallOpts)
}

// GetJobConfirmerHash is a free data retrieval call binding the contract method 0x7f890b11.
//
// Solidity: function getJobConfirmerHash() pure returns(bytes32)
func (_Constants *ConstantsCallerSession) GetJobConfirmerHash() ([32]byte, error) {
	return _Constants.Contract.GetJobConfirmerHash(&_Constants.CallOpts)
}

// GetStakeModifierHash is a free data retrieval call binding the contract method 0x80638d42.
//
// Solidity: function getStakeModifierHash() pure returns(bytes32)
func (_Constants *ConstantsCaller) GetStakeModifierHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Constants.contract.Call(opts, &out, "getStakeModifierHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetStakeModifierHash is a free data retrieval call binding the contract method 0x80638d42.
//
// Solidity: function getStakeModifierHash() pure returns(bytes32)
func (_Constants *ConstantsSession) GetStakeModifierHash() ([32]byte, error) {
	return _Constants.Contract.GetStakeModifierHash(&_Constants.CallOpts)
}

// GetStakeModifierHash is a free data retrieval call binding the contract method 0x80638d42.
//
// Solidity: function getStakeModifierHash() pure returns(bytes32)
func (_Constants *ConstantsCallerSession) GetStakeModifierHash() ([32]byte, error) {
	return _Constants.Contract.GetStakeModifierHash(&_Constants.CallOpts)
}

// GetStakerActivityUpdaterHash is a free data retrieval call binding the contract method 0x9cec6f8a.
//
// Solidity: function getStakerActivityUpdaterHash() pure returns(bytes32)
func (_Constants *ConstantsCaller) GetStakerActivityUpdaterHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Constants.contract.Call(opts, &out, "getStakerActivityUpdaterHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetStakerActivityUpdaterHash is a free data retrieval call binding the contract method 0x9cec6f8a.
//
// Solidity: function getStakerActivityUpdaterHash() pure returns(bytes32)
func (_Constants *ConstantsSession) GetStakerActivityUpdaterHash() ([32]byte, error) {
	return _Constants.Contract.GetStakerActivityUpdaterHash(&_Constants.CallOpts)
}

// GetStakerActivityUpdaterHash is a free data retrieval call binding the contract method 0x9cec6f8a.
//
// Solidity: function getStakerActivityUpdaterHash() pure returns(bytes32)
func (_Constants *ConstantsCallerSession) GetStakerActivityUpdaterHash() ([32]byte, error) {
	return _Constants.Contract.GetStakerActivityUpdaterHash(&_Constants.CallOpts)
}

// MaxAltBlocks is a free data retrieval call binding the contract method 0x379597e0.
//
// Solidity: function maxAltBlocks() pure returns(uint256)
func (_Constants *ConstantsCaller) MaxAltBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Constants.contract.Call(opts, &out, "maxAltBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAltBlocks is a free data retrieval call binding the contract method 0x379597e0.
//
// Solidity: function maxAltBlocks() pure returns(uint256)
func (_Constants *ConstantsSession) MaxAltBlocks() (*big.Int, error) {
	return _Constants.Contract.MaxAltBlocks(&_Constants.CallOpts)
}

// MaxAltBlocks is a free data retrieval call binding the contract method 0x379597e0.
//
// Solidity: function maxAltBlocks() pure returns(uint256)
func (_Constants *ConstantsCallerSession) MaxAltBlocks() (*big.Int, error) {
	return _Constants.Contract.MaxAltBlocks(&_Constants.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() pure returns(uint256)
func (_Constants *ConstantsCaller) MinStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Constants.contract.Call(opts, &out, "minStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() pure returns(uint256)
func (_Constants *ConstantsSession) MinStake() (*big.Int, error) {
	return _Constants.Contract.MinStake(&_Constants.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() pure returns(uint256)
func (_Constants *ConstantsCallerSession) MinStake() (*big.Int, error) {
	return _Constants.Contract.MinStake(&_Constants.CallOpts)
}

// NumStates is a free data retrieval call binding the contract method 0xf4f29c5b.
//
// Solidity: function numStates() pure returns(uint256)
func (_Constants *ConstantsCaller) NumStates(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Constants.contract.Call(opts, &out, "numStates")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumStates is a free data retrieval call binding the contract method 0xf4f29c5b.
//
// Solidity: function numStates() pure returns(uint256)
func (_Constants *ConstantsSession) NumStates() (*big.Int, error) {
	return _Constants.Contract.NumStates(&_Constants.CallOpts)
}

// NumStates is a free data retrieval call binding the contract method 0xf4f29c5b.
//
// Solidity: function numStates() pure returns(uint256)
func (_Constants *ConstantsCallerSession) NumStates() (*big.Int, error) {
	return _Constants.Contract.NumStates(&_Constants.CallOpts)
}

// PenaltyNotRevealDenom is a free data retrieval call binding the contract method 0xb1cc4500.
//
// Solidity: function penaltyNotRevealDenom() pure returns(uint256)
func (_Constants *ConstantsCaller) PenaltyNotRevealDenom(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Constants.contract.Call(opts, &out, "penaltyNotRevealDenom")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PenaltyNotRevealDenom is a free data retrieval call binding the contract method 0xb1cc4500.
//
// Solidity: function penaltyNotRevealDenom() pure returns(uint256)
func (_Constants *ConstantsSession) PenaltyNotRevealDenom() (*big.Int, error) {
	return _Constants.Contract.PenaltyNotRevealDenom(&_Constants.CallOpts)
}

// PenaltyNotRevealDenom is a free data retrieval call binding the contract method 0xb1cc4500.
//
// Solidity: function penaltyNotRevealDenom() pure returns(uint256)
func (_Constants *ConstantsCallerSession) PenaltyNotRevealDenom() (*big.Int, error) {
	return _Constants.Contract.PenaltyNotRevealDenom(&_Constants.CallOpts)
}

// PenaltyNotRevealNum is a free data retrieval call binding the contract method 0xa86f5a3f.
//
// Solidity: function penaltyNotRevealNum() pure returns(uint256)
func (_Constants *ConstantsCaller) PenaltyNotRevealNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Constants.contract.Call(opts, &out, "penaltyNotRevealNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PenaltyNotRevealNum is a free data retrieval call binding the contract method 0xa86f5a3f.
//
// Solidity: function penaltyNotRevealNum() pure returns(uint256)
func (_Constants *ConstantsSession) PenaltyNotRevealNum() (*big.Int, error) {
	return _Constants.Contract.PenaltyNotRevealNum(&_Constants.CallOpts)
}

// PenaltyNotRevealNum is a free data retrieval call binding the contract method 0xa86f5a3f.
//
// Solidity: function penaltyNotRevealNum() pure returns(uint256)
func (_Constants *ConstantsCallerSession) PenaltyNotRevealNum() (*big.Int, error) {
	return _Constants.Contract.PenaltyNotRevealNum(&_Constants.CallOpts)
}

// Propose is a free data retrieval call binding the contract method 0xc198f8ba.
//
// Solidity: function propose() pure returns(uint8)
func (_Constants *ConstantsCaller) Propose(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Constants.contract.Call(opts, &out, "propose")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Propose is a free data retrieval call binding the contract method 0xc198f8ba.
//
// Solidity: function propose() pure returns(uint8)
func (_Constants *ConstantsSession) Propose() (uint8, error) {
	return _Constants.Contract.Propose(&_Constants.CallOpts)
}

// Propose is a free data retrieval call binding the contract method 0xc198f8ba.
//
// Solidity: function propose() pure returns(uint8)
func (_Constants *ConstantsCallerSession) Propose() (uint8, error) {
	return _Constants.Contract.Propose(&_Constants.CallOpts)
}

// Reveal is a free data retrieval call binding the contract method 0xa475b5dd.
//
// Solidity: function reveal() pure returns(uint8)
func (_Constants *ConstantsCaller) Reveal(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Constants.contract.Call(opts, &out, "reveal")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Reveal is a free data retrieval call binding the contract method 0xa475b5dd.
//
// Solidity: function reveal() pure returns(uint8)
func (_Constants *ConstantsSession) Reveal() (uint8, error) {
	return _Constants.Contract.Reveal(&_Constants.CallOpts)
}

// Reveal is a free data retrieval call binding the contract method 0xa475b5dd.
//
// Solidity: function reveal() pure returns(uint8)
func (_Constants *ConstantsCallerSession) Reveal() (uint8, error) {
	return _Constants.Contract.Reveal(&_Constants.CallOpts)
}

// UnstakeLockPeriod is a free data retrieval call binding the contract method 0x26bf1c03.
//
// Solidity: function unstakeLockPeriod() pure returns(uint256)
func (_Constants *ConstantsCaller) UnstakeLockPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Constants.contract.Call(opts, &out, "unstakeLockPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnstakeLockPeriod is a free data retrieval call binding the contract method 0x26bf1c03.
//
// Solidity: function unstakeLockPeriod() pure returns(uint256)
func (_Constants *ConstantsSession) UnstakeLockPeriod() (*big.Int, error) {
	return _Constants.Contract.UnstakeLockPeriod(&_Constants.CallOpts)
}

// UnstakeLockPeriod is a free data retrieval call binding the contract method 0x26bf1c03.
//
// Solidity: function unstakeLockPeriod() pure returns(uint256)
func (_Constants *ConstantsCallerSession) UnstakeLockPeriod() (*big.Int, error) {
	return _Constants.Contract.UnstakeLockPeriod(&_Constants.CallOpts)
}

// WithdrawLockPeriod is a free data retrieval call binding the contract method 0x2628490f.
//
// Solidity: function withdrawLockPeriod() pure returns(uint256)
func (_Constants *ConstantsCaller) WithdrawLockPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Constants.contract.Call(opts, &out, "withdrawLockPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawLockPeriod is a free data retrieval call binding the contract method 0x2628490f.
//
// Solidity: function withdrawLockPeriod() pure returns(uint256)
func (_Constants *ConstantsSession) WithdrawLockPeriod() (*big.Int, error) {
	return _Constants.Contract.WithdrawLockPeriod(&_Constants.CallOpts)
}

// WithdrawLockPeriod is a free data retrieval call binding the contract method 0x2628490f.
//
// Solidity: function withdrawLockPeriod() pure returns(uint256)
func (_Constants *ConstantsCallerSession) WithdrawLockPeriod() (*big.Int, error) {
	return _Constants.Contract.WithdrawLockPeriod(&_Constants.CallOpts)
}
