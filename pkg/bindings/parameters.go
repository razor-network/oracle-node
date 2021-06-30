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

// ParametersABI is the input ABI used to generate the binding from.
const ParametersABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggregationRange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"commit\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dispute\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exposureDenominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAssetConfirmerHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockConfirmerHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDefaultAdminHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardModifierHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakeModifierHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakerActivityUpdaterHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gracePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxAltBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"penaltyNotRevealDenom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"penaltyNotRevealNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"propose\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetLockPenalty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reveal\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_aggregationRange\",\"type\":\"uint256\"}],\"name\":\"setAggregationRange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochLength\",\"type\":\"uint256\"}],\"name\":\"setEpochLength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_exposureDenominator\",\"type\":\"uint256\"}],\"name\":\"setExposureDenominator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gracePeriod\",\"type\":\"uint256\"}],\"name\":\"setGracePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxAltBlocks\",\"type\":\"uint256\"}],\"name\":\"setMaxAltBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minStake\",\"type\":\"uint256\"}],\"name\":\"setMinStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_numStates\",\"type\":\"uint256\"}],\"name\":\"setNumStates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_penaltyNotRevealDenom\",\"type\":\"uint256\"}],\"name\":\"setPenaltyNotRevealDeom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_penaltyNotRevealNumerator\",\"type\":\"uint256\"}],\"name\":\"setPenaltyNotRevealNum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_resetLockPenalty\",\"type\":\"uint256\"}],\"name\":\"setResetLockPenalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawLockPeriod\",\"type\":\"uint256\"}],\"name\":\"setWithdrawLockPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawReleasePeriod\",\"type\":\"uint256\"}],\"name\":\"setWithdrawReleasePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawLockPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawReleasePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Parameters is an auto generated Go binding around an Ethereum contract.
type Parameters struct {
	ParametersCaller     // Read-only binding to the contract
	ParametersTransactor // Write-only binding to the contract
	ParametersFilterer   // Log filterer for contract events
}

// ParametersCaller is an auto generated read-only Go binding around an Ethereum contract.
type ParametersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParametersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ParametersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParametersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ParametersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParametersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ParametersSession struct {
	Contract     *Parameters       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ParametersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ParametersCallerSession struct {
	Contract *ParametersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ParametersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ParametersTransactorSession struct {
	Contract     *ParametersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ParametersRaw is an auto generated low-level Go binding around an Ethereum contract.
type ParametersRaw struct {
	Contract *Parameters // Generic contract binding to access the raw methods on
}

// ParametersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ParametersCallerRaw struct {
	Contract *ParametersCaller // Generic read-only contract binding to access the raw methods on
}

// ParametersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ParametersTransactorRaw struct {
	Contract *ParametersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewParameters creates a new instance of Parameters, bound to a specific deployed contract.
func NewParameters(address common.Address, backend bind.ContractBackend) (*Parameters, error) {
	contract, err := bindParameters(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Parameters{ParametersCaller: ParametersCaller{contract: contract}, ParametersTransactor: ParametersTransactor{contract: contract}, ParametersFilterer: ParametersFilterer{contract: contract}}, nil
}

// NewParametersCaller creates a new read-only instance of Parameters, bound to a specific deployed contract.
func NewParametersCaller(address common.Address, caller bind.ContractCaller) (*ParametersCaller, error) {
	contract, err := bindParameters(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ParametersCaller{contract: contract}, nil
}

// NewParametersTransactor creates a new write-only instance of Parameters, bound to a specific deployed contract.
func NewParametersTransactor(address common.Address, transactor bind.ContractTransactor) (*ParametersTransactor, error) {
	contract, err := bindParameters(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ParametersTransactor{contract: contract}, nil
}

// NewParametersFilterer creates a new log filterer instance of Parameters, bound to a specific deployed contract.
func NewParametersFilterer(address common.Address, filterer bind.ContractFilterer) (*ParametersFilterer, error) {
	contract, err := bindParameters(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ParametersFilterer{contract: contract}, nil
}

// bindParameters binds a generic wrapper to an already deployed contract.
func bindParameters(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ParametersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Parameters *ParametersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Parameters.Contract.ParametersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Parameters *ParametersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Parameters.Contract.ParametersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Parameters *ParametersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Parameters.Contract.ParametersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Parameters *ParametersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Parameters.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Parameters *ParametersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Parameters.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Parameters *ParametersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Parameters.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Parameters *ParametersCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Parameters.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Parameters *ParametersSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Parameters.Contract.DEFAULTADMINROLE(&_Parameters.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Parameters *ParametersCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Parameters.Contract.DEFAULTADMINROLE(&_Parameters.CallOpts)
}

// AggregationRange is a free data retrieval call binding the contract method 0x3c5dd62f.
//
// Solidity: function aggregationRange() view returns(uint256)
func (_Parameters *ParametersCaller) AggregationRange(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Parameters.contract.Call(opts, &out, "aggregationRange")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AggregationRange is a free data retrieval call binding the contract method 0x3c5dd62f.
//
// Solidity: function aggregationRange() view returns(uint256)
func (_Parameters *ParametersSession) AggregationRange() (*big.Int, error) {
	return _Parameters.Contract.AggregationRange(&_Parameters.CallOpts)
}

// AggregationRange is a free data retrieval call binding the contract method 0x3c5dd62f.
//
// Solidity: function aggregationRange() view returns(uint256)
func (_Parameters *ParametersCallerSession) AggregationRange() (*big.Int, error) {
	return _Parameters.Contract.AggregationRange(&_Parameters.CallOpts)
}

// Commit is a free data retrieval call binding the contract method 0x3c7a3aff.
//
// Solidity: function commit() pure returns(uint32)
func (_Parameters *ParametersCaller) Commit(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Parameters.contract.Call(opts, &out, "commit")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Commit is a free data retrieval call binding the contract method 0x3c7a3aff.
//
// Solidity: function commit() pure returns(uint32)
func (_Parameters *ParametersSession) Commit() (uint32, error) {
	return _Parameters.Contract.Commit(&_Parameters.CallOpts)
}

// Commit is a free data retrieval call binding the contract method 0x3c7a3aff.
//
// Solidity: function commit() pure returns(uint32)
func (_Parameters *ParametersCallerSession) Commit() (uint32, error) {
	return _Parameters.Contract.Commit(&_Parameters.CallOpts)
}

// Dispute is a free data retrieval call binding the contract method 0xf240f7c3.
//
// Solidity: function dispute() pure returns(uint32)
func (_Parameters *ParametersCaller) Dispute(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Parameters.contract.Call(opts, &out, "dispute")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Dispute is a free data retrieval call binding the contract method 0xf240f7c3.
//
// Solidity: function dispute() pure returns(uint32)
func (_Parameters *ParametersSession) Dispute() (uint32, error) {
	return _Parameters.Contract.Dispute(&_Parameters.CallOpts)
}

// Dispute is a free data retrieval call binding the contract method 0xf240f7c3.
//
// Solidity: function dispute() pure returns(uint32)
func (_Parameters *ParametersCallerSession) Dispute() (uint32, error) {
	return _Parameters.Contract.Dispute(&_Parameters.CallOpts)
}

// EpochLength is a free data retrieval call binding the contract method 0x57d775f8.
//
// Solidity: function epochLength() view returns(uint256)
func (_Parameters *ParametersCaller) EpochLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Parameters.contract.Call(opts, &out, "epochLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochLength is a free data retrieval call binding the contract method 0x57d775f8.
//
// Solidity: function epochLength() view returns(uint256)
func (_Parameters *ParametersSession) EpochLength() (*big.Int, error) {
	return _Parameters.Contract.EpochLength(&_Parameters.CallOpts)
}

// EpochLength is a free data retrieval call binding the contract method 0x57d775f8.
//
// Solidity: function epochLength() view returns(uint256)
func (_Parameters *ParametersCallerSession) EpochLength() (*big.Int, error) {
	return _Parameters.Contract.EpochLength(&_Parameters.CallOpts)
}

// ExposureDenominator is a free data retrieval call binding the contract method 0x3002c9ac.
//
// Solidity: function exposureDenominator() view returns(uint256)
func (_Parameters *ParametersCaller) ExposureDenominator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Parameters.contract.Call(opts, &out, "exposureDenominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExposureDenominator is a free data retrieval call binding the contract method 0x3002c9ac.
//
// Solidity: function exposureDenominator() view returns(uint256)
func (_Parameters *ParametersSession) ExposureDenominator() (*big.Int, error) {
	return _Parameters.Contract.ExposureDenominator(&_Parameters.CallOpts)
}

// ExposureDenominator is a free data retrieval call binding the contract method 0x3002c9ac.
//
// Solidity: function exposureDenominator() view returns(uint256)
func (_Parameters *ParametersCallerSession) ExposureDenominator() (*big.Int, error) {
	return _Parameters.Contract.ExposureDenominator(&_Parameters.CallOpts)
}

// GetAssetConfirmerHash is a free data retrieval call binding the contract method 0x50f62021.
//
// Solidity: function getAssetConfirmerHash() pure returns(bytes32)
func (_Parameters *ParametersCaller) GetAssetConfirmerHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Parameters.contract.Call(opts, &out, "getAssetConfirmerHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAssetConfirmerHash is a free data retrieval call binding the contract method 0x50f62021.
//
// Solidity: function getAssetConfirmerHash() pure returns(bytes32)
func (_Parameters *ParametersSession) GetAssetConfirmerHash() ([32]byte, error) {
	return _Parameters.Contract.GetAssetConfirmerHash(&_Parameters.CallOpts)
}

// GetAssetConfirmerHash is a free data retrieval call binding the contract method 0x50f62021.
//
// Solidity: function getAssetConfirmerHash() pure returns(bytes32)
func (_Parameters *ParametersCallerSession) GetAssetConfirmerHash() ([32]byte, error) {
	return _Parameters.Contract.GetAssetConfirmerHash(&_Parameters.CallOpts)
}

// GetBlockConfirmerHash is a free data retrieval call binding the contract method 0xd98b3ced.
//
// Solidity: function getBlockConfirmerHash() pure returns(bytes32)
func (_Parameters *ParametersCaller) GetBlockConfirmerHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Parameters.contract.Call(opts, &out, "getBlockConfirmerHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockConfirmerHash is a free data retrieval call binding the contract method 0xd98b3ced.
//
// Solidity: function getBlockConfirmerHash() pure returns(bytes32)
func (_Parameters *ParametersSession) GetBlockConfirmerHash() ([32]byte, error) {
	return _Parameters.Contract.GetBlockConfirmerHash(&_Parameters.CallOpts)
}

// GetBlockConfirmerHash is a free data retrieval call binding the contract method 0xd98b3ced.
//
// Solidity: function getBlockConfirmerHash() pure returns(bytes32)
func (_Parameters *ParametersCallerSession) GetBlockConfirmerHash() ([32]byte, error) {
	return _Parameters.Contract.GetBlockConfirmerHash(&_Parameters.CallOpts)
}

// GetDefaultAdminHash is a free data retrieval call binding the contract method 0xb1a62781.
//
// Solidity: function getDefaultAdminHash() pure returns(bytes32)
func (_Parameters *ParametersCaller) GetDefaultAdminHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Parameters.contract.Call(opts, &out, "getDefaultAdminHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDefaultAdminHash is a free data retrieval call binding the contract method 0xb1a62781.
//
// Solidity: function getDefaultAdminHash() pure returns(bytes32)
func (_Parameters *ParametersSession) GetDefaultAdminHash() ([32]byte, error) {
	return _Parameters.Contract.GetDefaultAdminHash(&_Parameters.CallOpts)
}

// GetDefaultAdminHash is a free data retrieval call binding the contract method 0xb1a62781.
//
// Solidity: function getDefaultAdminHash() pure returns(bytes32)
func (_Parameters *ParametersCallerSession) GetDefaultAdminHash() ([32]byte, error) {
	return _Parameters.Contract.GetDefaultAdminHash(&_Parameters.CallOpts)
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() view returns(uint256)
func (_Parameters *ParametersCaller) GetEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Parameters.contract.Call(opts, &out, "getEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() view returns(uint256)
func (_Parameters *ParametersSession) GetEpoch() (*big.Int, error) {
	return _Parameters.Contract.GetEpoch(&_Parameters.CallOpts)
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() view returns(uint256)
func (_Parameters *ParametersCallerSession) GetEpoch() (*big.Int, error) {
	return _Parameters.Contract.GetEpoch(&_Parameters.CallOpts)
}

// GetRewardModifierHash is a free data retrieval call binding the contract method 0x66061b17.
//
// Solidity: function getRewardModifierHash() pure returns(bytes32)
func (_Parameters *ParametersCaller) GetRewardModifierHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Parameters.contract.Call(opts, &out, "getRewardModifierHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRewardModifierHash is a free data retrieval call binding the contract method 0x66061b17.
//
// Solidity: function getRewardModifierHash() pure returns(bytes32)
func (_Parameters *ParametersSession) GetRewardModifierHash() ([32]byte, error) {
	return _Parameters.Contract.GetRewardModifierHash(&_Parameters.CallOpts)
}

// GetRewardModifierHash is a free data retrieval call binding the contract method 0x66061b17.
//
// Solidity: function getRewardModifierHash() pure returns(bytes32)
func (_Parameters *ParametersCallerSession) GetRewardModifierHash() ([32]byte, error) {
	return _Parameters.Contract.GetRewardModifierHash(&_Parameters.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Parameters *ParametersCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Parameters.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Parameters *ParametersSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Parameters.Contract.GetRoleAdmin(&_Parameters.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Parameters *ParametersCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Parameters.Contract.GetRoleAdmin(&_Parameters.CallOpts, role)
}

// GetStakeModifierHash is a free data retrieval call binding the contract method 0x80638d42.
//
// Solidity: function getStakeModifierHash() pure returns(bytes32)
func (_Parameters *ParametersCaller) GetStakeModifierHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Parameters.contract.Call(opts, &out, "getStakeModifierHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetStakeModifierHash is a free data retrieval call binding the contract method 0x80638d42.
//
// Solidity: function getStakeModifierHash() pure returns(bytes32)
func (_Parameters *ParametersSession) GetStakeModifierHash() ([32]byte, error) {
	return _Parameters.Contract.GetStakeModifierHash(&_Parameters.CallOpts)
}

// GetStakeModifierHash is a free data retrieval call binding the contract method 0x80638d42.
//
// Solidity: function getStakeModifierHash() pure returns(bytes32)
func (_Parameters *ParametersCallerSession) GetStakeModifierHash() ([32]byte, error) {
	return _Parameters.Contract.GetStakeModifierHash(&_Parameters.CallOpts)
}

// GetStakerActivityUpdaterHash is a free data retrieval call binding the contract method 0x9cec6f8a.
//
// Solidity: function getStakerActivityUpdaterHash() pure returns(bytes32)
func (_Parameters *ParametersCaller) GetStakerActivityUpdaterHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Parameters.contract.Call(opts, &out, "getStakerActivityUpdaterHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetStakerActivityUpdaterHash is a free data retrieval call binding the contract method 0x9cec6f8a.
//
// Solidity: function getStakerActivityUpdaterHash() pure returns(bytes32)
func (_Parameters *ParametersSession) GetStakerActivityUpdaterHash() ([32]byte, error) {
	return _Parameters.Contract.GetStakerActivityUpdaterHash(&_Parameters.CallOpts)
}

// GetStakerActivityUpdaterHash is a free data retrieval call binding the contract method 0x9cec6f8a.
//
// Solidity: function getStakerActivityUpdaterHash() pure returns(bytes32)
func (_Parameters *ParametersCallerSession) GetStakerActivityUpdaterHash() ([32]byte, error) {
	return _Parameters.Contract.GetStakerActivityUpdaterHash(&_Parameters.CallOpts)
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() view returns(uint256)
func (_Parameters *ParametersCaller) GetState(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Parameters.contract.Call(opts, &out, "getState")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() view returns(uint256)
func (_Parameters *ParametersSession) GetState() (*big.Int, error) {
	return _Parameters.Contract.GetState(&_Parameters.CallOpts)
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() view returns(uint256)
func (_Parameters *ParametersCallerSession) GetState() (*big.Int, error) {
	return _Parameters.Contract.GetState(&_Parameters.CallOpts)
}

// GracePeriod is a free data retrieval call binding the contract method 0xa06db7dc.
//
// Solidity: function gracePeriod() view returns(uint256)
func (_Parameters *ParametersCaller) GracePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Parameters.contract.Call(opts, &out, "gracePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GracePeriod is a free data retrieval call binding the contract method 0xa06db7dc.
//
// Solidity: function gracePeriod() view returns(uint256)
func (_Parameters *ParametersSession) GracePeriod() (*big.Int, error) {
	return _Parameters.Contract.GracePeriod(&_Parameters.CallOpts)
}

// GracePeriod is a free data retrieval call binding the contract method 0xa06db7dc.
//
// Solidity: function gracePeriod() view returns(uint256)
func (_Parameters *ParametersCallerSession) GracePeriod() (*big.Int, error) {
	return _Parameters.Contract.GracePeriod(&_Parameters.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Parameters *ParametersCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Parameters.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Parameters *ParametersSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Parameters.Contract.HasRole(&_Parameters.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Parameters *ParametersCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Parameters.Contract.HasRole(&_Parameters.CallOpts, role, account)
}

// MaxAltBlocks is a free data retrieval call binding the contract method 0x379597e0.
//
// Solidity: function maxAltBlocks() view returns(uint256)
func (_Parameters *ParametersCaller) MaxAltBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Parameters.contract.Call(opts, &out, "maxAltBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAltBlocks is a free data retrieval call binding the contract method 0x379597e0.
//
// Solidity: function maxAltBlocks() view returns(uint256)
func (_Parameters *ParametersSession) MaxAltBlocks() (*big.Int, error) {
	return _Parameters.Contract.MaxAltBlocks(&_Parameters.CallOpts)
}

// MaxAltBlocks is a free data retrieval call binding the contract method 0x379597e0.
//
// Solidity: function maxAltBlocks() view returns(uint256)
func (_Parameters *ParametersCallerSession) MaxAltBlocks() (*big.Int, error) {
	return _Parameters.Contract.MaxAltBlocks(&_Parameters.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() view returns(uint256)
func (_Parameters *ParametersCaller) MinStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Parameters.contract.Call(opts, &out, "minStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() view returns(uint256)
func (_Parameters *ParametersSession) MinStake() (*big.Int, error) {
	return _Parameters.Contract.MinStake(&_Parameters.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() view returns(uint256)
func (_Parameters *ParametersCallerSession) MinStake() (*big.Int, error) {
	return _Parameters.Contract.MinStake(&_Parameters.CallOpts)
}

// NumStates is a free data retrieval call binding the contract method 0xf4f29c5b.
//
// Solidity: function numStates() view returns(uint256)
func (_Parameters *ParametersCaller) NumStates(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Parameters.contract.Call(opts, &out, "numStates")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumStates is a free data retrieval call binding the contract method 0xf4f29c5b.
//
// Solidity: function numStates() view returns(uint256)
func (_Parameters *ParametersSession) NumStates() (*big.Int, error) {
	return _Parameters.Contract.NumStates(&_Parameters.CallOpts)
}

// NumStates is a free data retrieval call binding the contract method 0xf4f29c5b.
//
// Solidity: function numStates() view returns(uint256)
func (_Parameters *ParametersCallerSession) NumStates() (*big.Int, error) {
	return _Parameters.Contract.NumStates(&_Parameters.CallOpts)
}

// PenaltyNotRevealDenom is a free data retrieval call binding the contract method 0xb1cc4500.
//
// Solidity: function penaltyNotRevealDenom() view returns(uint256)
func (_Parameters *ParametersCaller) PenaltyNotRevealDenom(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Parameters.contract.Call(opts, &out, "penaltyNotRevealDenom")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PenaltyNotRevealDenom is a free data retrieval call binding the contract method 0xb1cc4500.
//
// Solidity: function penaltyNotRevealDenom() view returns(uint256)
func (_Parameters *ParametersSession) PenaltyNotRevealDenom() (*big.Int, error) {
	return _Parameters.Contract.PenaltyNotRevealDenom(&_Parameters.CallOpts)
}

// PenaltyNotRevealDenom is a free data retrieval call binding the contract method 0xb1cc4500.
//
// Solidity: function penaltyNotRevealDenom() view returns(uint256)
func (_Parameters *ParametersCallerSession) PenaltyNotRevealDenom() (*big.Int, error) {
	return _Parameters.Contract.PenaltyNotRevealDenom(&_Parameters.CallOpts)
}

// PenaltyNotRevealNum is a free data retrieval call binding the contract method 0xa86f5a3f.
//
// Solidity: function penaltyNotRevealNum() view returns(uint256)
func (_Parameters *ParametersCaller) PenaltyNotRevealNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Parameters.contract.Call(opts, &out, "penaltyNotRevealNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PenaltyNotRevealNum is a free data retrieval call binding the contract method 0xa86f5a3f.
//
// Solidity: function penaltyNotRevealNum() view returns(uint256)
func (_Parameters *ParametersSession) PenaltyNotRevealNum() (*big.Int, error) {
	return _Parameters.Contract.PenaltyNotRevealNum(&_Parameters.CallOpts)
}

// PenaltyNotRevealNum is a free data retrieval call binding the contract method 0xa86f5a3f.
//
// Solidity: function penaltyNotRevealNum() view returns(uint256)
func (_Parameters *ParametersCallerSession) PenaltyNotRevealNum() (*big.Int, error) {
	return _Parameters.Contract.PenaltyNotRevealNum(&_Parameters.CallOpts)
}

// Propose is a free data retrieval call binding the contract method 0xc198f8ba.
//
// Solidity: function propose() pure returns(uint32)
func (_Parameters *ParametersCaller) Propose(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Parameters.contract.Call(opts, &out, "propose")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Propose is a free data retrieval call binding the contract method 0xc198f8ba.
//
// Solidity: function propose() pure returns(uint32)
func (_Parameters *ParametersSession) Propose() (uint32, error) {
	return _Parameters.Contract.Propose(&_Parameters.CallOpts)
}

// Propose is a free data retrieval call binding the contract method 0xc198f8ba.
//
// Solidity: function propose() pure returns(uint32)
func (_Parameters *ParametersCallerSession) Propose() (uint32, error) {
	return _Parameters.Contract.Propose(&_Parameters.CallOpts)
}

// ResetLockPenalty is a free data retrieval call binding the contract method 0x86d51626.
//
// Solidity: function resetLockPenalty() view returns(uint256)
func (_Parameters *ParametersCaller) ResetLockPenalty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Parameters.contract.Call(opts, &out, "resetLockPenalty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ResetLockPenalty is a free data retrieval call binding the contract method 0x86d51626.
//
// Solidity: function resetLockPenalty() view returns(uint256)
func (_Parameters *ParametersSession) ResetLockPenalty() (*big.Int, error) {
	return _Parameters.Contract.ResetLockPenalty(&_Parameters.CallOpts)
}

// ResetLockPenalty is a free data retrieval call binding the contract method 0x86d51626.
//
// Solidity: function resetLockPenalty() view returns(uint256)
func (_Parameters *ParametersCallerSession) ResetLockPenalty() (*big.Int, error) {
	return _Parameters.Contract.ResetLockPenalty(&_Parameters.CallOpts)
}

// Reveal is a free data retrieval call binding the contract method 0xa475b5dd.
//
// Solidity: function reveal() pure returns(uint32)
func (_Parameters *ParametersCaller) Reveal(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Parameters.contract.Call(opts, &out, "reveal")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Reveal is a free data retrieval call binding the contract method 0xa475b5dd.
//
// Solidity: function reveal() pure returns(uint32)
func (_Parameters *ParametersSession) Reveal() (uint32, error) {
	return _Parameters.Contract.Reveal(&_Parameters.CallOpts)
}

// Reveal is a free data retrieval call binding the contract method 0xa475b5dd.
//
// Solidity: function reveal() pure returns(uint32)
func (_Parameters *ParametersCallerSession) Reveal() (uint32, error) {
	return _Parameters.Contract.Reveal(&_Parameters.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Parameters *ParametersCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Parameters.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Parameters *ParametersSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Parameters.Contract.SupportsInterface(&_Parameters.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Parameters *ParametersCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Parameters.Contract.SupportsInterface(&_Parameters.CallOpts, interfaceId)
}

// WithdrawLockPeriod is a free data retrieval call binding the contract method 0x2628490f.
//
// Solidity: function withdrawLockPeriod() view returns(uint256)
func (_Parameters *ParametersCaller) WithdrawLockPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Parameters.contract.Call(opts, &out, "withdrawLockPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawLockPeriod is a free data retrieval call binding the contract method 0x2628490f.
//
// Solidity: function withdrawLockPeriod() view returns(uint256)
func (_Parameters *ParametersSession) WithdrawLockPeriod() (*big.Int, error) {
	return _Parameters.Contract.WithdrawLockPeriod(&_Parameters.CallOpts)
}

// WithdrawLockPeriod is a free data retrieval call binding the contract method 0x2628490f.
//
// Solidity: function withdrawLockPeriod() view returns(uint256)
func (_Parameters *ParametersCallerSession) WithdrawLockPeriod() (*big.Int, error) {
	return _Parameters.Contract.WithdrawLockPeriod(&_Parameters.CallOpts)
}

// WithdrawReleasePeriod is a free data retrieval call binding the contract method 0x13072f49.
//
// Solidity: function withdrawReleasePeriod() view returns(uint256)
func (_Parameters *ParametersCaller) WithdrawReleasePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Parameters.contract.Call(opts, &out, "withdrawReleasePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawReleasePeriod is a free data retrieval call binding the contract method 0x13072f49.
//
// Solidity: function withdrawReleasePeriod() view returns(uint256)
func (_Parameters *ParametersSession) WithdrawReleasePeriod() (*big.Int, error) {
	return _Parameters.Contract.WithdrawReleasePeriod(&_Parameters.CallOpts)
}

// WithdrawReleasePeriod is a free data retrieval call binding the contract method 0x13072f49.
//
// Solidity: function withdrawReleasePeriod() view returns(uint256)
func (_Parameters *ParametersCallerSession) WithdrawReleasePeriod() (*big.Int, error) {
	return _Parameters.Contract.WithdrawReleasePeriod(&_Parameters.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Parameters *ParametersTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Parameters *ParametersSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Parameters.Contract.GrantRole(&_Parameters.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Parameters *ParametersTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Parameters.Contract.GrantRole(&_Parameters.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Parameters *ParametersTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Parameters *ParametersSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Parameters.Contract.RenounceRole(&_Parameters.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Parameters *ParametersTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Parameters.Contract.RenounceRole(&_Parameters.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Parameters *ParametersTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Parameters *ParametersSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Parameters.Contract.RevokeRole(&_Parameters.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Parameters *ParametersTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Parameters.Contract.RevokeRole(&_Parameters.TransactOpts, role, account)
}

// SetAggregationRange is a paid mutator transaction binding the contract method 0x1a6adb9b.
//
// Solidity: function setAggregationRange(uint256 _aggregationRange) returns()
func (_Parameters *ParametersTransactor) SetAggregationRange(opts *bind.TransactOpts, _aggregationRange *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setAggregationRange", _aggregationRange)
}

// SetAggregationRange is a paid mutator transaction binding the contract method 0x1a6adb9b.
//
// Solidity: function setAggregationRange(uint256 _aggregationRange) returns()
func (_Parameters *ParametersSession) SetAggregationRange(_aggregationRange *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetAggregationRange(&_Parameters.TransactOpts, _aggregationRange)
}

// SetAggregationRange is a paid mutator transaction binding the contract method 0x1a6adb9b.
//
// Solidity: function setAggregationRange(uint256 _aggregationRange) returns()
func (_Parameters *ParametersTransactorSession) SetAggregationRange(_aggregationRange *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetAggregationRange(&_Parameters.TransactOpts, _aggregationRange)
}

// SetEpochLength is a paid mutator transaction binding the contract method 0x54eea796.
//
// Solidity: function setEpochLength(uint256 _epochLength) returns()
func (_Parameters *ParametersTransactor) SetEpochLength(opts *bind.TransactOpts, _epochLength *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setEpochLength", _epochLength)
}

// SetEpochLength is a paid mutator transaction binding the contract method 0x54eea796.
//
// Solidity: function setEpochLength(uint256 _epochLength) returns()
func (_Parameters *ParametersSession) SetEpochLength(_epochLength *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetEpochLength(&_Parameters.TransactOpts, _epochLength)
}

// SetEpochLength is a paid mutator transaction binding the contract method 0x54eea796.
//
// Solidity: function setEpochLength(uint256 _epochLength) returns()
func (_Parameters *ParametersTransactorSession) SetEpochLength(_epochLength *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetEpochLength(&_Parameters.TransactOpts, _epochLength)
}

// SetExposureDenominator is a paid mutator transaction binding the contract method 0x27ad73fe.
//
// Solidity: function setExposureDenominator(uint256 _exposureDenominator) returns()
func (_Parameters *ParametersTransactor) SetExposureDenominator(opts *bind.TransactOpts, _exposureDenominator *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setExposureDenominator", _exposureDenominator)
}

// SetExposureDenominator is a paid mutator transaction binding the contract method 0x27ad73fe.
//
// Solidity: function setExposureDenominator(uint256 _exposureDenominator) returns()
func (_Parameters *ParametersSession) SetExposureDenominator(_exposureDenominator *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetExposureDenominator(&_Parameters.TransactOpts, _exposureDenominator)
}

// SetExposureDenominator is a paid mutator transaction binding the contract method 0x27ad73fe.
//
// Solidity: function setExposureDenominator(uint256 _exposureDenominator) returns()
func (_Parameters *ParametersTransactorSession) SetExposureDenominator(_exposureDenominator *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetExposureDenominator(&_Parameters.TransactOpts, _exposureDenominator)
}

// SetGracePeriod is a paid mutator transaction binding the contract method 0xf2f65960.
//
// Solidity: function setGracePeriod(uint256 _gracePeriod) returns()
func (_Parameters *ParametersTransactor) SetGracePeriod(opts *bind.TransactOpts, _gracePeriod *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setGracePeriod", _gracePeriod)
}

// SetGracePeriod is a paid mutator transaction binding the contract method 0xf2f65960.
//
// Solidity: function setGracePeriod(uint256 _gracePeriod) returns()
func (_Parameters *ParametersSession) SetGracePeriod(_gracePeriod *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetGracePeriod(&_Parameters.TransactOpts, _gracePeriod)
}

// SetGracePeriod is a paid mutator transaction binding the contract method 0xf2f65960.
//
// Solidity: function setGracePeriod(uint256 _gracePeriod) returns()
func (_Parameters *ParametersTransactorSession) SetGracePeriod(_gracePeriod *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetGracePeriod(&_Parameters.TransactOpts, _gracePeriod)
}

// SetMaxAltBlocks is a paid mutator transaction binding the contract method 0x55ab6929.
//
// Solidity: function setMaxAltBlocks(uint256 _maxAltBlocks) returns()
func (_Parameters *ParametersTransactor) SetMaxAltBlocks(opts *bind.TransactOpts, _maxAltBlocks *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setMaxAltBlocks", _maxAltBlocks)
}

// SetMaxAltBlocks is a paid mutator transaction binding the contract method 0x55ab6929.
//
// Solidity: function setMaxAltBlocks(uint256 _maxAltBlocks) returns()
func (_Parameters *ParametersSession) SetMaxAltBlocks(_maxAltBlocks *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetMaxAltBlocks(&_Parameters.TransactOpts, _maxAltBlocks)
}

// SetMaxAltBlocks is a paid mutator transaction binding the contract method 0x55ab6929.
//
// Solidity: function setMaxAltBlocks(uint256 _maxAltBlocks) returns()
func (_Parameters *ParametersTransactorSession) SetMaxAltBlocks(_maxAltBlocks *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetMaxAltBlocks(&_Parameters.TransactOpts, _maxAltBlocks)
}

// SetMinStake is a paid mutator transaction binding the contract method 0x8c80fd90.
//
// Solidity: function setMinStake(uint256 _minStake) returns()
func (_Parameters *ParametersTransactor) SetMinStake(opts *bind.TransactOpts, _minStake *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setMinStake", _minStake)
}

// SetMinStake is a paid mutator transaction binding the contract method 0x8c80fd90.
//
// Solidity: function setMinStake(uint256 _minStake) returns()
func (_Parameters *ParametersSession) SetMinStake(_minStake *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetMinStake(&_Parameters.TransactOpts, _minStake)
}

// SetMinStake is a paid mutator transaction binding the contract method 0x8c80fd90.
//
// Solidity: function setMinStake(uint256 _minStake) returns()
func (_Parameters *ParametersTransactorSession) SetMinStake(_minStake *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetMinStake(&_Parameters.TransactOpts, _minStake)
}

// SetNumStates is a paid mutator transaction binding the contract method 0xf15dc529.
//
// Solidity: function setNumStates(uint256 _numStates) returns()
func (_Parameters *ParametersTransactor) SetNumStates(opts *bind.TransactOpts, _numStates *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setNumStates", _numStates)
}

// SetNumStates is a paid mutator transaction binding the contract method 0xf15dc529.
//
// Solidity: function setNumStates(uint256 _numStates) returns()
func (_Parameters *ParametersSession) SetNumStates(_numStates *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetNumStates(&_Parameters.TransactOpts, _numStates)
}

// SetNumStates is a paid mutator transaction binding the contract method 0xf15dc529.
//
// Solidity: function setNumStates(uint256 _numStates) returns()
func (_Parameters *ParametersTransactorSession) SetNumStates(_numStates *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetNumStates(&_Parameters.TransactOpts, _numStates)
}

// SetPenaltyNotRevealDeom is a paid mutator transaction binding the contract method 0x6a7d2955.
//
// Solidity: function setPenaltyNotRevealDeom(uint256 _penaltyNotRevealDenom) returns()
func (_Parameters *ParametersTransactor) SetPenaltyNotRevealDeom(opts *bind.TransactOpts, _penaltyNotRevealDenom *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setPenaltyNotRevealDeom", _penaltyNotRevealDenom)
}

// SetPenaltyNotRevealDeom is a paid mutator transaction binding the contract method 0x6a7d2955.
//
// Solidity: function setPenaltyNotRevealDeom(uint256 _penaltyNotRevealDenom) returns()
func (_Parameters *ParametersSession) SetPenaltyNotRevealDeom(_penaltyNotRevealDenom *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetPenaltyNotRevealDeom(&_Parameters.TransactOpts, _penaltyNotRevealDenom)
}

// SetPenaltyNotRevealDeom is a paid mutator transaction binding the contract method 0x6a7d2955.
//
// Solidity: function setPenaltyNotRevealDeom(uint256 _penaltyNotRevealDenom) returns()
func (_Parameters *ParametersTransactorSession) SetPenaltyNotRevealDeom(_penaltyNotRevealDenom *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetPenaltyNotRevealDeom(&_Parameters.TransactOpts, _penaltyNotRevealDenom)
}

// SetPenaltyNotRevealNum is a paid mutator transaction binding the contract method 0x5583f92c.
//
// Solidity: function setPenaltyNotRevealNum(uint256 _penaltyNotRevealNumerator) returns()
func (_Parameters *ParametersTransactor) SetPenaltyNotRevealNum(opts *bind.TransactOpts, _penaltyNotRevealNumerator *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setPenaltyNotRevealNum", _penaltyNotRevealNumerator)
}

// SetPenaltyNotRevealNum is a paid mutator transaction binding the contract method 0x5583f92c.
//
// Solidity: function setPenaltyNotRevealNum(uint256 _penaltyNotRevealNumerator) returns()
func (_Parameters *ParametersSession) SetPenaltyNotRevealNum(_penaltyNotRevealNumerator *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetPenaltyNotRevealNum(&_Parameters.TransactOpts, _penaltyNotRevealNumerator)
}

// SetPenaltyNotRevealNum is a paid mutator transaction binding the contract method 0x5583f92c.
//
// Solidity: function setPenaltyNotRevealNum(uint256 _penaltyNotRevealNumerator) returns()
func (_Parameters *ParametersTransactorSession) SetPenaltyNotRevealNum(_penaltyNotRevealNumerator *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetPenaltyNotRevealNum(&_Parameters.TransactOpts, _penaltyNotRevealNumerator)
}

// SetResetLockPenalty is a paid mutator transaction binding the contract method 0x802e00e4.
//
// Solidity: function setResetLockPenalty(uint256 _resetLockPenalty) returns()
func (_Parameters *ParametersTransactor) SetResetLockPenalty(opts *bind.TransactOpts, _resetLockPenalty *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setResetLockPenalty", _resetLockPenalty)
}

// SetResetLockPenalty is a paid mutator transaction binding the contract method 0x802e00e4.
//
// Solidity: function setResetLockPenalty(uint256 _resetLockPenalty) returns()
func (_Parameters *ParametersSession) SetResetLockPenalty(_resetLockPenalty *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetResetLockPenalty(&_Parameters.TransactOpts, _resetLockPenalty)
}

// SetResetLockPenalty is a paid mutator transaction binding the contract method 0x802e00e4.
//
// Solidity: function setResetLockPenalty(uint256 _resetLockPenalty) returns()
func (_Parameters *ParametersTransactorSession) SetResetLockPenalty(_resetLockPenalty *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetResetLockPenalty(&_Parameters.TransactOpts, _resetLockPenalty)
}

// SetWithdrawLockPeriod is a paid mutator transaction binding the contract method 0x8aa0c7bb.
//
// Solidity: function setWithdrawLockPeriod(uint256 _withdrawLockPeriod) returns()
func (_Parameters *ParametersTransactor) SetWithdrawLockPeriod(opts *bind.TransactOpts, _withdrawLockPeriod *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setWithdrawLockPeriod", _withdrawLockPeriod)
}

// SetWithdrawLockPeriod is a paid mutator transaction binding the contract method 0x8aa0c7bb.
//
// Solidity: function setWithdrawLockPeriod(uint256 _withdrawLockPeriod) returns()
func (_Parameters *ParametersSession) SetWithdrawLockPeriod(_withdrawLockPeriod *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetWithdrawLockPeriod(&_Parameters.TransactOpts, _withdrawLockPeriod)
}

// SetWithdrawLockPeriod is a paid mutator transaction binding the contract method 0x8aa0c7bb.
//
// Solidity: function setWithdrawLockPeriod(uint256 _withdrawLockPeriod) returns()
func (_Parameters *ParametersTransactorSession) SetWithdrawLockPeriod(_withdrawLockPeriod *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetWithdrawLockPeriod(&_Parameters.TransactOpts, _withdrawLockPeriod)
}

// SetWithdrawReleasePeriod is a paid mutator transaction binding the contract method 0x1169e9bf.
//
// Solidity: function setWithdrawReleasePeriod(uint256 _withdrawReleasePeriod) returns()
func (_Parameters *ParametersTransactor) SetWithdrawReleasePeriod(opts *bind.TransactOpts, _withdrawReleasePeriod *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setWithdrawReleasePeriod", _withdrawReleasePeriod)
}

// SetWithdrawReleasePeriod is a paid mutator transaction binding the contract method 0x1169e9bf.
//
// Solidity: function setWithdrawReleasePeriod(uint256 _withdrawReleasePeriod) returns()
func (_Parameters *ParametersSession) SetWithdrawReleasePeriod(_withdrawReleasePeriod *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetWithdrawReleasePeriod(&_Parameters.TransactOpts, _withdrawReleasePeriod)
}

// SetWithdrawReleasePeriod is a paid mutator transaction binding the contract method 0x1169e9bf.
//
// Solidity: function setWithdrawReleasePeriod(uint256 _withdrawReleasePeriod) returns()
func (_Parameters *ParametersTransactorSession) SetWithdrawReleasePeriod(_withdrawReleasePeriod *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetWithdrawReleasePeriod(&_Parameters.TransactOpts, _withdrawReleasePeriod)
}

// ParametersRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Parameters contract.
type ParametersRoleAdminChangedIterator struct {
	Event *ParametersRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ParametersRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParametersRoleAdminChanged)
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
		it.Event = new(ParametersRoleAdminChanged)
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
func (it *ParametersRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParametersRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParametersRoleAdminChanged represents a RoleAdminChanged event raised by the Parameters contract.
type ParametersRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Parameters *ParametersFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ParametersRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Parameters.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ParametersRoleAdminChangedIterator{contract: _Parameters.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Parameters *ParametersFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ParametersRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Parameters.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParametersRoleAdminChanged)
				if err := _Parameters.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Parameters *ParametersFilterer) ParseRoleAdminChanged(log types.Log) (*ParametersRoleAdminChanged, error) {
	event := new(ParametersRoleAdminChanged)
	if err := _Parameters.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ParametersRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Parameters contract.
type ParametersRoleGrantedIterator struct {
	Event *ParametersRoleGranted // Event containing the contract specifics and raw log

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
func (it *ParametersRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParametersRoleGranted)
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
		it.Event = new(ParametersRoleGranted)
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
func (it *ParametersRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParametersRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParametersRoleGranted represents a RoleGranted event raised by the Parameters contract.
type ParametersRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Parameters *ParametersFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ParametersRoleGrantedIterator, error) {

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

	logs, sub, err := _Parameters.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ParametersRoleGrantedIterator{contract: _Parameters.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Parameters *ParametersFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ParametersRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Parameters.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParametersRoleGranted)
				if err := _Parameters.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Parameters *ParametersFilterer) ParseRoleGranted(log types.Log) (*ParametersRoleGranted, error) {
	event := new(ParametersRoleGranted)
	if err := _Parameters.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ParametersRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Parameters contract.
type ParametersRoleRevokedIterator struct {
	Event *ParametersRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ParametersRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParametersRoleRevoked)
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
		it.Event = new(ParametersRoleRevoked)
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
func (it *ParametersRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParametersRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParametersRoleRevoked represents a RoleRevoked event raised by the Parameters contract.
type ParametersRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Parameters *ParametersFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ParametersRoleRevokedIterator, error) {

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

	logs, sub, err := _Parameters.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ParametersRoleRevokedIterator{contract: _Parameters.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Parameters *ParametersFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ParametersRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Parameters.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParametersRoleRevoked)
				if err := _Parameters.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Parameters *ParametersFilterer) ParseRoleRevoked(log types.Log) (*ParametersRoleRevoked, error) {
	event := new(ParametersRoleRevoked)
	if err := _Parameters.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
