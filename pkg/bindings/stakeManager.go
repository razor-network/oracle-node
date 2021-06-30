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

// StructsStaker is an auto generated low-level Go binding around an user-defined struct.
type StructsStaker struct {
	Id                 *big.Int
	Address            common.Address
	Stake              *big.Int
	EpochStaked        *big.Int
	EpochLastCommitted *big.Int
	EpochLastRevealed  *big.Int
	AcceptDelegation   bool
	Commission         *big.Int
	TokenAddress       common.Address
}

// StakeManagerABI is the input ABI used to generate the binding from.
const StakeManagerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousStake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newStake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Delegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousStake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newStake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"StakeChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousStake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newStake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newStake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newStake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Withdrew\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"}],\"name\":\"decreaseCommission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakerId\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumStakers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getStaker\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochStaked\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochLastCommitted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochLastRevealed\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"acceptDelegation\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"internalType\":\"structStructs.Staker\",\"name\":\"staker\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getStakerId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"schAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"voteManagersAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"parametersAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"locks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAfter\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStakers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"parameters\",\"outputs\":[{\"internalType\":\"contractIParameters\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerId\",\"type\":\"uint256\"}],\"name\":\"resetLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardManager\",\"outputs\":[{\"internalType\":\"contractIRewardManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sch\",\"outputs\":[{\"internalType\":\"contractSchellingCoin\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"}],\"name\":\"setCommission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"setDelegationAcceptance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_epochLastRevealed\",\"type\":\"uint256\"}],\"name\":\"setStakerEpochLastRevealed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stake\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_reason\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"}],\"name\":\"setStakerStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakerIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochStaked\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochLastCommitted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochLastRevealed\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"acceptDelegation\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bountyHunter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bounty\",\"type\":\"uint256\"}],\"name\":\"transferBounty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sAmount\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerId\",\"type\":\"uint256\"}],\"name\":\"updateCommitmentEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteManager\",\"outputs\":[{\"internalType\":\"contractIVoteManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakerId\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// StakeManager is an auto generated Go binding around an Ethereum contract.
type StakeManager struct {
	StakeManagerCaller     // Read-only binding to the contract
	StakeManagerTransactor // Write-only binding to the contract
	StakeManagerFilterer   // Log filterer for contract events
}

// StakeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakeManagerSession struct {
	Contract     *StakeManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakeManagerCallerSession struct {
	Contract *StakeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StakeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakeManagerTransactorSession struct {
	Contract     *StakeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StakeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakeManagerRaw struct {
	Contract *StakeManager // Generic contract binding to access the raw methods on
}

// StakeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakeManagerCallerRaw struct {
	Contract *StakeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// StakeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakeManagerTransactorRaw struct {
	Contract *StakeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakeManager creates a new instance of StakeManager, bound to a specific deployed contract.
func NewStakeManager(address common.Address, backend bind.ContractBackend) (*StakeManager, error) {
	contract, err := bindStakeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakeManager{StakeManagerCaller: StakeManagerCaller{contract: contract}, StakeManagerTransactor: StakeManagerTransactor{contract: contract}, StakeManagerFilterer: StakeManagerFilterer{contract: contract}}, nil
}

// NewStakeManagerCaller creates a new read-only instance of StakeManager, bound to a specific deployed contract.
func NewStakeManagerCaller(address common.Address, caller bind.ContractCaller) (*StakeManagerCaller, error) {
	contract, err := bindStakeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakeManagerCaller{contract: contract}, nil
}

// NewStakeManagerTransactor creates a new write-only instance of StakeManager, bound to a specific deployed contract.
func NewStakeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*StakeManagerTransactor, error) {
	contract, err := bindStakeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakeManagerTransactor{contract: contract}, nil
}

// NewStakeManagerFilterer creates a new log filterer instance of StakeManager, bound to a specific deployed contract.
func NewStakeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*StakeManagerFilterer, error) {
	contract, err := bindStakeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakeManagerFilterer{contract: contract}, nil
}

// bindStakeManager binds a generic wrapper to an already deployed contract.
func bindStakeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakeManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeManager *StakeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeManager.Contract.StakeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeManager *StakeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeManager.Contract.StakeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeManager *StakeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeManager.Contract.StakeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeManager *StakeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeManager *StakeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeManager *StakeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeManager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakeManager *StakeManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakeManager *StakeManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StakeManager.Contract.DEFAULTADMINROLE(&_StakeManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakeManager *StakeManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StakeManager.Contract.DEFAULTADMINROLE(&_StakeManager.CallOpts)
}

// GetNumStakers is a free data retrieval call binding the contract method 0xbc788d46.
//
// Solidity: function getNumStakers() view returns(uint256)
func (_StakeManager *StakeManagerCaller) GetNumStakers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "getNumStakers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumStakers is a free data retrieval call binding the contract method 0xbc788d46.
//
// Solidity: function getNumStakers() view returns(uint256)
func (_StakeManager *StakeManagerSession) GetNumStakers() (*big.Int, error) {
	return _StakeManager.Contract.GetNumStakers(&_StakeManager.CallOpts)
}

// GetNumStakers is a free data retrieval call binding the contract method 0xbc788d46.
//
// Solidity: function getNumStakers() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) GetNumStakers() (*big.Int, error) {
	return _StakeManager.Contract.GetNumStakers(&_StakeManager.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakeManager *StakeManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakeManager *StakeManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StakeManager.Contract.GetRoleAdmin(&_StakeManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakeManager *StakeManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StakeManager.Contract.GetRoleAdmin(&_StakeManager.CallOpts, role)
}

// GetStaker is a free data retrieval call binding the contract method 0xe3c998fe.
//
// Solidity: function getStaker(uint256 _id) view returns((uint256,address,uint256,uint256,uint256,uint256,bool,uint256,address) staker)
func (_StakeManager *StakeManagerCaller) GetStaker(opts *bind.CallOpts, _id *big.Int) (StructsStaker, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "getStaker", _id)

	if err != nil {
		return *new(StructsStaker), err
	}

	out0 := *abi.ConvertType(out[0], new(StructsStaker)).(*StructsStaker)

	return out0, err

}

// GetStaker is a free data retrieval call binding the contract method 0xe3c998fe.
//
// Solidity: function getStaker(uint256 _id) view returns((uint256,address,uint256,uint256,uint256,uint256,bool,uint256,address) staker)
func (_StakeManager *StakeManagerSession) GetStaker(_id *big.Int) (StructsStaker, error) {
	return _StakeManager.Contract.GetStaker(&_StakeManager.CallOpts, _id)
}

// GetStaker is a free data retrieval call binding the contract method 0xe3c998fe.
//
// Solidity: function getStaker(uint256 _id) view returns((uint256,address,uint256,uint256,uint256,uint256,bool,uint256,address) staker)
func (_StakeManager *StakeManagerCallerSession) GetStaker(_id *big.Int) (StructsStaker, error) {
	return _StakeManager.Contract.GetStaker(&_StakeManager.CallOpts, _id)
}

// GetStakerId is a free data retrieval call binding the contract method 0x6022a485.
//
// Solidity: function getStakerId(address _address) view returns(uint256)
func (_StakeManager *StakeManagerCaller) GetStakerId(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "getStakerId", _address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakerId is a free data retrieval call binding the contract method 0x6022a485.
//
// Solidity: function getStakerId(address _address) view returns(uint256)
func (_StakeManager *StakeManagerSession) GetStakerId(_address common.Address) (*big.Int, error) {
	return _StakeManager.Contract.GetStakerId(&_StakeManager.CallOpts, _address)
}

// GetStakerId is a free data retrieval call binding the contract method 0x6022a485.
//
// Solidity: function getStakerId(address _address) view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) GetStakerId(_address common.Address) (*big.Int, error) {
	return _StakeManager.Contract.GetStakerId(&_StakeManager.CallOpts, _address)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakeManager *StakeManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakeManager *StakeManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StakeManager.Contract.HasRole(&_StakeManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakeManager *StakeManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StakeManager.Contract.HasRole(&_StakeManager.CallOpts, role, account)
}

// Locks is a free data retrieval call binding the contract method 0xc05f6155.
//
// Solidity: function locks(address , address ) view returns(uint256 amount, uint256 withdrawAfter)
func (_StakeManager *StakeManagerCaller) Locks(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	Amount        *big.Int
	WithdrawAfter *big.Int
}, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "locks", arg0, arg1)

	outstruct := new(struct {
		Amount        *big.Int
		WithdrawAfter *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.WithdrawAfter = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Locks is a free data retrieval call binding the contract method 0xc05f6155.
//
// Solidity: function locks(address , address ) view returns(uint256 amount, uint256 withdrawAfter)
func (_StakeManager *StakeManagerSession) Locks(arg0 common.Address, arg1 common.Address) (struct {
	Amount        *big.Int
	WithdrawAfter *big.Int
}, error) {
	return _StakeManager.Contract.Locks(&_StakeManager.CallOpts, arg0, arg1)
}

// Locks is a free data retrieval call binding the contract method 0xc05f6155.
//
// Solidity: function locks(address , address ) view returns(uint256 amount, uint256 withdrawAfter)
func (_StakeManager *StakeManagerCallerSession) Locks(arg0 common.Address, arg1 common.Address) (struct {
	Amount        *big.Int
	WithdrawAfter *big.Int
}, error) {
	return _StakeManager.Contract.Locks(&_StakeManager.CallOpts, arg0, arg1)
}

// NumStakers is a free data retrieval call binding the contract method 0x6c8b052a.
//
// Solidity: function numStakers() view returns(uint256)
func (_StakeManager *StakeManagerCaller) NumStakers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "numStakers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumStakers is a free data retrieval call binding the contract method 0x6c8b052a.
//
// Solidity: function numStakers() view returns(uint256)
func (_StakeManager *StakeManagerSession) NumStakers() (*big.Int, error) {
	return _StakeManager.Contract.NumStakers(&_StakeManager.CallOpts)
}

// NumStakers is a free data retrieval call binding the contract method 0x6c8b052a.
//
// Solidity: function numStakers() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) NumStakers() (*big.Int, error) {
	return _StakeManager.Contract.NumStakers(&_StakeManager.CallOpts)
}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address)
func (_StakeManager *StakeManagerCaller) Parameters(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "parameters")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address)
func (_StakeManager *StakeManagerSession) Parameters() (common.Address, error) {
	return _StakeManager.Contract.Parameters(&_StakeManager.CallOpts)
}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address)
func (_StakeManager *StakeManagerCallerSession) Parameters() (common.Address, error) {
	return _StakeManager.Contract.Parameters(&_StakeManager.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_StakeManager *StakeManagerCaller) RewardManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "rewardManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_StakeManager *StakeManagerSession) RewardManager() (common.Address, error) {
	return _StakeManager.Contract.RewardManager(&_StakeManager.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_StakeManager *StakeManagerCallerSession) RewardManager() (common.Address, error) {
	return _StakeManager.Contract.RewardManager(&_StakeManager.CallOpts)
}

// Sch is a free data retrieval call binding the contract method 0xc584bb9f.
//
// Solidity: function sch() view returns(address)
func (_StakeManager *StakeManagerCaller) Sch(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "sch")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sch is a free data retrieval call binding the contract method 0xc584bb9f.
//
// Solidity: function sch() view returns(address)
func (_StakeManager *StakeManagerSession) Sch() (common.Address, error) {
	return _StakeManager.Contract.Sch(&_StakeManager.CallOpts)
}

// Sch is a free data retrieval call binding the contract method 0xc584bb9f.
//
// Solidity: function sch() view returns(address)
func (_StakeManager *StakeManagerCallerSession) Sch() (common.Address, error) {
	return _StakeManager.Contract.Sch(&_StakeManager.CallOpts)
}

// StakerIds is a free data retrieval call binding the contract method 0xc8ae0d7d.
//
// Solidity: function stakerIds(address ) view returns(uint256)
func (_StakeManager *StakeManagerCaller) StakerIds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "stakerIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerIds is a free data retrieval call binding the contract method 0xc8ae0d7d.
//
// Solidity: function stakerIds(address ) view returns(uint256)
func (_StakeManager *StakeManagerSession) StakerIds(arg0 common.Address) (*big.Int, error) {
	return _StakeManager.Contract.StakerIds(&_StakeManager.CallOpts, arg0)
}

// StakerIds is a free data retrieval call binding the contract method 0xc8ae0d7d.
//
// Solidity: function stakerIds(address ) view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) StakerIds(arg0 common.Address) (*big.Int, error) {
	return _StakeManager.Contract.StakerIds(&_StakeManager.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) view returns(uint256 id, address _address, uint256 stake, uint256 epochStaked, uint256 epochLastCommitted, uint256 epochLastRevealed, bool acceptDelegation, uint256 commission, address tokenAddress)
func (_StakeManager *StakeManagerCaller) Stakers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id                 *big.Int
	Address            common.Address
	Stake              *big.Int
	EpochStaked        *big.Int
	EpochLastCommitted *big.Int
	EpochLastRevealed  *big.Int
	AcceptDelegation   bool
	Commission         *big.Int
	TokenAddress       common.Address
}, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "stakers", arg0)

	outstruct := new(struct {
		Id                 *big.Int
		Address            common.Address
		Stake              *big.Int
		EpochStaked        *big.Int
		EpochLastCommitted *big.Int
		EpochLastRevealed  *big.Int
		AcceptDelegation   bool
		Commission         *big.Int
		TokenAddress       common.Address
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
	outstruct.AcceptDelegation = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.Commission = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.TokenAddress = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) view returns(uint256 id, address _address, uint256 stake, uint256 epochStaked, uint256 epochLastCommitted, uint256 epochLastRevealed, bool acceptDelegation, uint256 commission, address tokenAddress)
func (_StakeManager *StakeManagerSession) Stakers(arg0 *big.Int) (struct {
	Id                 *big.Int
	Address            common.Address
	Stake              *big.Int
	EpochStaked        *big.Int
	EpochLastCommitted *big.Int
	EpochLastRevealed  *big.Int
	AcceptDelegation   bool
	Commission         *big.Int
	TokenAddress       common.Address
}, error) {
	return _StakeManager.Contract.Stakers(&_StakeManager.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) view returns(uint256 id, address _address, uint256 stake, uint256 epochStaked, uint256 epochLastCommitted, uint256 epochLastRevealed, bool acceptDelegation, uint256 commission, address tokenAddress)
func (_StakeManager *StakeManagerCallerSession) Stakers(arg0 *big.Int) (struct {
	Id                 *big.Int
	Address            common.Address
	Stake              *big.Int
	EpochStaked        *big.Int
	EpochLastCommitted *big.Int
	EpochLastRevealed  *big.Int
	AcceptDelegation   bool
	Commission         *big.Int
	TokenAddress       common.Address
}, error) {
	return _StakeManager.Contract.Stakers(&_StakeManager.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakeManager *StakeManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakeManager *StakeManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StakeManager.Contract.SupportsInterface(&_StakeManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakeManager *StakeManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StakeManager.Contract.SupportsInterface(&_StakeManager.CallOpts, interfaceId)
}

// VoteManager is a free data retrieval call binding the contract method 0x42c1e587.
//
// Solidity: function voteManager() view returns(address)
func (_StakeManager *StakeManagerCaller) VoteManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "voteManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoteManager is a free data retrieval call binding the contract method 0x42c1e587.
//
// Solidity: function voteManager() view returns(address)
func (_StakeManager *StakeManagerSession) VoteManager() (common.Address, error) {
	return _StakeManager.Contract.VoteManager(&_StakeManager.CallOpts)
}

// VoteManager is a free data retrieval call binding the contract method 0x42c1e587.
//
// Solidity: function voteManager() view returns(address)
func (_StakeManager *StakeManagerCallerSession) VoteManager() (common.Address, error) {
	return _StakeManager.Contract.VoteManager(&_StakeManager.CallOpts)
}

// DecreaseCommission is a paid mutator transaction binding the contract method 0x2c2984a1.
//
// Solidity: function decreaseCommission(uint256 commission) returns()
func (_StakeManager *StakeManagerTransactor) DecreaseCommission(opts *bind.TransactOpts, commission *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "decreaseCommission", commission)
}

// DecreaseCommission is a paid mutator transaction binding the contract method 0x2c2984a1.
//
// Solidity: function decreaseCommission(uint256 commission) returns()
func (_StakeManager *StakeManagerSession) DecreaseCommission(commission *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.DecreaseCommission(&_StakeManager.TransactOpts, commission)
}

// DecreaseCommission is a paid mutator transaction binding the contract method 0x2c2984a1.
//
// Solidity: function decreaseCommission(uint256 commission) returns()
func (_StakeManager *StakeManagerTransactorSession) DecreaseCommission(commission *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.DecreaseCommission(&_StakeManager.TransactOpts, commission)
}

// Delegate is a paid mutator transaction binding the contract method 0x66fb64d0.
//
// Solidity: function delegate(uint256 epoch, uint256 amount, uint256 stakerId) returns()
func (_StakeManager *StakeManagerTransactor) Delegate(opts *bind.TransactOpts, epoch *big.Int, amount *big.Int, stakerId *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "delegate", epoch, amount, stakerId)
}

// Delegate is a paid mutator transaction binding the contract method 0x66fb64d0.
//
// Solidity: function delegate(uint256 epoch, uint256 amount, uint256 stakerId) returns()
func (_StakeManager *StakeManagerSession) Delegate(epoch *big.Int, amount *big.Int, stakerId *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Delegate(&_StakeManager.TransactOpts, epoch, amount, stakerId)
}

// Delegate is a paid mutator transaction binding the contract method 0x66fb64d0.
//
// Solidity: function delegate(uint256 epoch, uint256 amount, uint256 stakerId) returns()
func (_StakeManager *StakeManagerTransactorSession) Delegate(epoch *big.Int, amount *big.Int, stakerId *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Delegate(&_StakeManager.TransactOpts, epoch, amount, stakerId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakeManager *StakeManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakeManager *StakeManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.GrantRole(&_StakeManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakeManager *StakeManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.GrantRole(&_StakeManager.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address schAddress, address rewardManagerAddress, address voteManagersAddress, address parametersAddress) returns()
func (_StakeManager *StakeManagerTransactor) Initialize(opts *bind.TransactOpts, schAddress common.Address, rewardManagerAddress common.Address, voteManagersAddress common.Address, parametersAddress common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "initialize", schAddress, rewardManagerAddress, voteManagersAddress, parametersAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address schAddress, address rewardManagerAddress, address voteManagersAddress, address parametersAddress) returns()
func (_StakeManager *StakeManagerSession) Initialize(schAddress common.Address, rewardManagerAddress common.Address, voteManagersAddress common.Address, parametersAddress common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.Initialize(&_StakeManager.TransactOpts, schAddress, rewardManagerAddress, voteManagersAddress, parametersAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address schAddress, address rewardManagerAddress, address voteManagersAddress, address parametersAddress) returns()
func (_StakeManager *StakeManagerTransactorSession) Initialize(schAddress common.Address, rewardManagerAddress common.Address, voteManagersAddress common.Address, parametersAddress common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.Initialize(&_StakeManager.TransactOpts, schAddress, rewardManagerAddress, voteManagersAddress, parametersAddress)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StakeManager *StakeManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StakeManager *StakeManagerSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.RenounceRole(&_StakeManager.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StakeManager *StakeManagerTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.RenounceRole(&_StakeManager.TransactOpts, role, account)
}

// ResetLock is a paid mutator transaction binding the contract method 0x421c53ca.
//
// Solidity: function resetLock(uint256 stakerId) returns()
func (_StakeManager *StakeManagerTransactor) ResetLock(opts *bind.TransactOpts, stakerId *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "resetLock", stakerId)
}

// ResetLock is a paid mutator transaction binding the contract method 0x421c53ca.
//
// Solidity: function resetLock(uint256 stakerId) returns()
func (_StakeManager *StakeManagerSession) ResetLock(stakerId *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.ResetLock(&_StakeManager.TransactOpts, stakerId)
}

// ResetLock is a paid mutator transaction binding the contract method 0x421c53ca.
//
// Solidity: function resetLock(uint256 stakerId) returns()
func (_StakeManager *StakeManagerTransactorSession) ResetLock(stakerId *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.ResetLock(&_StakeManager.TransactOpts, stakerId)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakeManager *StakeManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakeManager *StakeManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.RevokeRole(&_StakeManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakeManager *StakeManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.RevokeRole(&_StakeManager.TransactOpts, role, account)
}

// SetCommission is a paid mutator transaction binding the contract method 0x355e6b43.
//
// Solidity: function setCommission(uint256 commission) returns()
func (_StakeManager *StakeManagerTransactor) SetCommission(opts *bind.TransactOpts, commission *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "setCommission", commission)
}

// SetCommission is a paid mutator transaction binding the contract method 0x355e6b43.
//
// Solidity: function setCommission(uint256 commission) returns()
func (_StakeManager *StakeManagerSession) SetCommission(commission *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.SetCommission(&_StakeManager.TransactOpts, commission)
}

// SetCommission is a paid mutator transaction binding the contract method 0x355e6b43.
//
// Solidity: function setCommission(uint256 commission) returns()
func (_StakeManager *StakeManagerTransactorSession) SetCommission(commission *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.SetCommission(&_StakeManager.TransactOpts, commission)
}

// SetDelegationAcceptance is a paid mutator transaction binding the contract method 0x85d628fd.
//
// Solidity: function setDelegationAcceptance(bool status) returns()
func (_StakeManager *StakeManagerTransactor) SetDelegationAcceptance(opts *bind.TransactOpts, status bool) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "setDelegationAcceptance", status)
}

// SetDelegationAcceptance is a paid mutator transaction binding the contract method 0x85d628fd.
//
// Solidity: function setDelegationAcceptance(bool status) returns()
func (_StakeManager *StakeManagerSession) SetDelegationAcceptance(status bool) (*types.Transaction, error) {
	return _StakeManager.Contract.SetDelegationAcceptance(&_StakeManager.TransactOpts, status)
}

// SetDelegationAcceptance is a paid mutator transaction binding the contract method 0x85d628fd.
//
// Solidity: function setDelegationAcceptance(bool status) returns()
func (_StakeManager *StakeManagerTransactorSession) SetDelegationAcceptance(status bool) (*types.Transaction, error) {
	return _StakeManager.Contract.SetDelegationAcceptance(&_StakeManager.TransactOpts, status)
}

// SetStakerEpochLastRevealed is a paid mutator transaction binding the contract method 0x9864f70a.
//
// Solidity: function setStakerEpochLastRevealed(uint256 _id, uint256 _epochLastRevealed) returns()
func (_StakeManager *StakeManagerTransactor) SetStakerEpochLastRevealed(opts *bind.TransactOpts, _id *big.Int, _epochLastRevealed *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "setStakerEpochLastRevealed", _id, _epochLastRevealed)
}

// SetStakerEpochLastRevealed is a paid mutator transaction binding the contract method 0x9864f70a.
//
// Solidity: function setStakerEpochLastRevealed(uint256 _id, uint256 _epochLastRevealed) returns()
func (_StakeManager *StakeManagerSession) SetStakerEpochLastRevealed(_id *big.Int, _epochLastRevealed *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.SetStakerEpochLastRevealed(&_StakeManager.TransactOpts, _id, _epochLastRevealed)
}

// SetStakerEpochLastRevealed is a paid mutator transaction binding the contract method 0x9864f70a.
//
// Solidity: function setStakerEpochLastRevealed(uint256 _id, uint256 _epochLastRevealed) returns()
func (_StakeManager *StakeManagerTransactorSession) SetStakerEpochLastRevealed(_id *big.Int, _epochLastRevealed *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.SetStakerEpochLastRevealed(&_StakeManager.TransactOpts, _id, _epochLastRevealed)
}

// SetStakerStake is a paid mutator transaction binding the contract method 0xa8a3e96f.
//
// Solidity: function setStakerStake(uint256 _id, uint256 _stake, string _reason, uint256 _epoch) returns()
func (_StakeManager *StakeManagerTransactor) SetStakerStake(opts *bind.TransactOpts, _id *big.Int, _stake *big.Int, _reason string, _epoch *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "setStakerStake", _id, _stake, _reason, _epoch)
}

// SetStakerStake is a paid mutator transaction binding the contract method 0xa8a3e96f.
//
// Solidity: function setStakerStake(uint256 _id, uint256 _stake, string _reason, uint256 _epoch) returns()
func (_StakeManager *StakeManagerSession) SetStakerStake(_id *big.Int, _stake *big.Int, _reason string, _epoch *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.SetStakerStake(&_StakeManager.TransactOpts, _id, _stake, _reason, _epoch)
}

// SetStakerStake is a paid mutator transaction binding the contract method 0xa8a3e96f.
//
// Solidity: function setStakerStake(uint256 _id, uint256 _stake, string _reason, uint256 _epoch) returns()
func (_StakeManager *StakeManagerTransactorSession) SetStakerStake(_id *big.Int, _stake *big.Int, _reason string, _epoch *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.SetStakerStake(&_StakeManager.TransactOpts, _id, _stake, _reason, _epoch)
}

// Stake is a paid mutator transaction binding the contract method 0x7b0472f0.
//
// Solidity: function stake(uint256 epoch, uint256 amount) returns()
func (_StakeManager *StakeManagerTransactor) Stake(opts *bind.TransactOpts, epoch *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "stake", epoch, amount)
}

// Stake is a paid mutator transaction binding the contract method 0x7b0472f0.
//
// Solidity: function stake(uint256 epoch, uint256 amount) returns()
func (_StakeManager *StakeManagerSession) Stake(epoch *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Stake(&_StakeManager.TransactOpts, epoch, amount)
}

// Stake is a paid mutator transaction binding the contract method 0x7b0472f0.
//
// Solidity: function stake(uint256 epoch, uint256 amount) returns()
func (_StakeManager *StakeManagerTransactorSession) Stake(epoch *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Stake(&_StakeManager.TransactOpts, epoch, amount)
}

// TransferBounty is a paid mutator transaction binding the contract method 0x2f7442bb.
//
// Solidity: function transferBounty(address bountyHunter, uint256 bounty) returns()
func (_StakeManager *StakeManagerTransactor) TransferBounty(opts *bind.TransactOpts, bountyHunter common.Address, bounty *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "transferBounty", bountyHunter, bounty)
}

// TransferBounty is a paid mutator transaction binding the contract method 0x2f7442bb.
//
// Solidity: function transferBounty(address bountyHunter, uint256 bounty) returns()
func (_StakeManager *StakeManagerSession) TransferBounty(bountyHunter common.Address, bounty *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.TransferBounty(&_StakeManager.TransactOpts, bountyHunter, bounty)
}

// TransferBounty is a paid mutator transaction binding the contract method 0x2f7442bb.
//
// Solidity: function transferBounty(address bountyHunter, uint256 bounty) returns()
func (_StakeManager *StakeManagerTransactorSession) TransferBounty(bountyHunter common.Address, bounty *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.TransferBounty(&_StakeManager.TransactOpts, bountyHunter, bounty)
}

// Unstake is a paid mutator transaction binding the contract method 0xbd97b375.
//
// Solidity: function unstake(uint256 epoch, uint256 stakerId, uint256 sAmount) returns()
func (_StakeManager *StakeManagerTransactor) Unstake(opts *bind.TransactOpts, epoch *big.Int, stakerId *big.Int, sAmount *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "unstake", epoch, stakerId, sAmount)
}

// Unstake is a paid mutator transaction binding the contract method 0xbd97b375.
//
// Solidity: function unstake(uint256 epoch, uint256 stakerId, uint256 sAmount) returns()
func (_StakeManager *StakeManagerSession) Unstake(epoch *big.Int, stakerId *big.Int, sAmount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Unstake(&_StakeManager.TransactOpts, epoch, stakerId, sAmount)
}

// Unstake is a paid mutator transaction binding the contract method 0xbd97b375.
//
// Solidity: function unstake(uint256 epoch, uint256 stakerId, uint256 sAmount) returns()
func (_StakeManager *StakeManagerTransactorSession) Unstake(epoch *big.Int, stakerId *big.Int, sAmount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Unstake(&_StakeManager.TransactOpts, epoch, stakerId, sAmount)
}

// UpdateBlockReward is a paid mutator transaction binding the contract method 0xf580ffcb.
//
// Solidity: function updateBlockReward(uint256 _blockReward) returns()
func (_StakeManager *StakeManagerTransactor) UpdateBlockReward(opts *bind.TransactOpts, _blockReward *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "updateBlockReward", _blockReward)
}

// UpdateBlockReward is a paid mutator transaction binding the contract method 0xf580ffcb.
//
// Solidity: function updateBlockReward(uint256 _blockReward) returns()
func (_StakeManager *StakeManagerSession) UpdateBlockReward(_blockReward *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UpdateBlockReward(&_StakeManager.TransactOpts, _blockReward)
}

// UpdateBlockReward is a paid mutator transaction binding the contract method 0xf580ffcb.
//
// Solidity: function updateBlockReward(uint256 _blockReward) returns()
func (_StakeManager *StakeManagerTransactorSession) UpdateBlockReward(_blockReward *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UpdateBlockReward(&_StakeManager.TransactOpts, _blockReward)
}

// UpdateCommitmentEpoch is a paid mutator transaction binding the contract method 0x188dc83b.
//
// Solidity: function updateCommitmentEpoch(uint256 stakerId) returns()
func (_StakeManager *StakeManagerTransactor) UpdateCommitmentEpoch(opts *bind.TransactOpts, stakerId *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "updateCommitmentEpoch", stakerId)
}

// UpdateCommitmentEpoch is a paid mutator transaction binding the contract method 0x188dc83b.
//
// Solidity: function updateCommitmentEpoch(uint256 stakerId) returns()
func (_StakeManager *StakeManagerSession) UpdateCommitmentEpoch(stakerId *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UpdateCommitmentEpoch(&_StakeManager.TransactOpts, stakerId)
}

// UpdateCommitmentEpoch is a paid mutator transaction binding the contract method 0x188dc83b.
//
// Solidity: function updateCommitmentEpoch(uint256 stakerId) returns()
func (_StakeManager *StakeManagerTransactorSession) UpdateCommitmentEpoch(stakerId *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.UpdateCommitmentEpoch(&_StakeManager.TransactOpts, stakerId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 epoch, uint256 stakerId) returns()
func (_StakeManager *StakeManagerTransactor) Withdraw(opts *bind.TransactOpts, epoch *big.Int, stakerId *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "withdraw", epoch, stakerId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 epoch, uint256 stakerId) returns()
func (_StakeManager *StakeManagerSession) Withdraw(epoch *big.Int, stakerId *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Withdraw(&_StakeManager.TransactOpts, epoch, stakerId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 epoch, uint256 stakerId) returns()
func (_StakeManager *StakeManagerTransactorSession) Withdraw(epoch *big.Int, stakerId *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Withdraw(&_StakeManager.TransactOpts, epoch, stakerId)
}

// StakeManagerDelegatedIterator is returned from FilterDelegated and is used to iterate over the raw logs and unpacked data for Delegated events raised by the StakeManager contract.
type StakeManagerDelegatedIterator struct {
	Event *StakeManagerDelegated // Event containing the contract specifics and raw log

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
func (it *StakeManagerDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerDelegated)
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
		it.Event = new(StakeManagerDelegated)
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
func (it *StakeManagerDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerDelegated represents a Delegated event raised by the StakeManager contract.
type StakeManagerDelegated struct {
	Epoch         *big.Int
	StakerId      *big.Int
	Delegator     common.Address
	PreviousStake *big.Int
	NewStake      *big.Int
	Timestamp     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDelegated is a free log retrieval operation binding the contract event 0x62c8696fe18e088d2a304fd3b56ab661b0e33cd63749ede69d4986b79d0e3f7f.
//
// Solidity: event Delegated(uint256 epoch, uint256 indexed stakerId, address delegator, uint256 previousStake, uint256 newStake, uint256 timestamp)
func (_StakeManager *StakeManagerFilterer) FilterDelegated(opts *bind.FilterOpts, stakerId []*big.Int) (*StakeManagerDelegatedIterator, error) {

	var stakerIdRule []interface{}
	for _, stakerIdItem := range stakerId {
		stakerIdRule = append(stakerIdRule, stakerIdItem)
	}

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "Delegated", stakerIdRule)
	if err != nil {
		return nil, err
	}
	return &StakeManagerDelegatedIterator{contract: _StakeManager.contract, event: "Delegated", logs: logs, sub: sub}, nil
}

// WatchDelegated is a free log subscription operation binding the contract event 0x62c8696fe18e088d2a304fd3b56ab661b0e33cd63749ede69d4986b79d0e3f7f.
//
// Solidity: event Delegated(uint256 epoch, uint256 indexed stakerId, address delegator, uint256 previousStake, uint256 newStake, uint256 timestamp)
func (_StakeManager *StakeManagerFilterer) WatchDelegated(opts *bind.WatchOpts, sink chan<- *StakeManagerDelegated, stakerId []*big.Int) (event.Subscription, error) {

	var stakerIdRule []interface{}
	for _, stakerIdItem := range stakerId {
		stakerIdRule = append(stakerIdRule, stakerIdItem)
	}

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "Delegated", stakerIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerDelegated)
				if err := _StakeManager.contract.UnpackLog(event, "Delegated", log); err != nil {
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

// ParseDelegated is a log parse operation binding the contract event 0x62c8696fe18e088d2a304fd3b56ab661b0e33cd63749ede69d4986b79d0e3f7f.
//
// Solidity: event Delegated(uint256 epoch, uint256 indexed stakerId, address delegator, uint256 previousStake, uint256 newStake, uint256 timestamp)
func (_StakeManager *StakeManagerFilterer) ParseDelegated(log types.Log) (*StakeManagerDelegated, error) {
	event := new(StakeManagerDelegated)
	if err := _StakeManager.contract.UnpackLog(event, "Delegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the StakeManager contract.
type StakeManagerRoleAdminChangedIterator struct {
	Event *StakeManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *StakeManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerRoleAdminChanged)
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
		it.Event = new(StakeManagerRoleAdminChanged)
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
func (it *StakeManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerRoleAdminChanged represents a RoleAdminChanged event raised by the StakeManager contract.
type StakeManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StakeManager *StakeManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StakeManagerRoleAdminChangedIterator, error) {

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

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StakeManagerRoleAdminChangedIterator{contract: _StakeManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StakeManager *StakeManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StakeManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerRoleAdminChanged)
				if err := _StakeManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_StakeManager *StakeManagerFilterer) ParseRoleAdminChanged(log types.Log) (*StakeManagerRoleAdminChanged, error) {
	event := new(StakeManagerRoleAdminChanged)
	if err := _StakeManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the StakeManager contract.
type StakeManagerRoleGrantedIterator struct {
	Event *StakeManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *StakeManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerRoleGranted)
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
		it.Event = new(StakeManagerRoleGranted)
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
func (it *StakeManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerRoleGranted represents a RoleGranted event raised by the StakeManager contract.
type StakeManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakeManager *StakeManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakeManagerRoleGrantedIterator, error) {

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

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakeManagerRoleGrantedIterator{contract: _StakeManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakeManager *StakeManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StakeManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerRoleGranted)
				if err := _StakeManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_StakeManager *StakeManagerFilterer) ParseRoleGranted(log types.Log) (*StakeManagerRoleGranted, error) {
	event := new(StakeManagerRoleGranted)
	if err := _StakeManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the StakeManager contract.
type StakeManagerRoleRevokedIterator struct {
	Event *StakeManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *StakeManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerRoleRevoked)
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
		it.Event = new(StakeManagerRoleRevoked)
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
func (it *StakeManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerRoleRevoked represents a RoleRevoked event raised by the StakeManager contract.
type StakeManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakeManager *StakeManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakeManagerRoleRevokedIterator, error) {

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

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakeManagerRoleRevokedIterator{contract: _StakeManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakeManager *StakeManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StakeManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerRoleRevoked)
				if err := _StakeManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_StakeManager *StakeManagerFilterer) ParseRoleRevoked(log types.Log) (*StakeManagerRoleRevoked, error) {
	event := new(StakeManagerRoleRevoked)
	if err := _StakeManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerStakeChangeIterator is returned from FilterStakeChange and is used to iterate over the raw logs and unpacked data for StakeChange events raised by the StakeManager contract.
type StakeManagerStakeChangeIterator struct {
	Event *StakeManagerStakeChange // Event containing the contract specifics and raw log

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
func (it *StakeManagerStakeChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerStakeChange)
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
		it.Event = new(StakeManagerStakeChange)
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
func (it *StakeManagerStakeChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerStakeChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerStakeChange represents a StakeChange event raised by the StakeManager contract.
type StakeManagerStakeChange struct {
	StakerId      *big.Int
	PreviousStake *big.Int
	NewStake      *big.Int
	Reason        string
	Epoch         *big.Int
	Timestamp     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterStakeChange is a free log retrieval operation binding the contract event 0xab9982d99de79485b3cd6d1e72c881fd0342f42f1cbc7d220d4f373b97c4f363.
//
// Solidity: event StakeChange(uint256 indexed stakerId, uint256 previousStake, uint256 newStake, string reason, uint256 epoch, uint256 timestamp)
func (_StakeManager *StakeManagerFilterer) FilterStakeChange(opts *bind.FilterOpts, stakerId []*big.Int) (*StakeManagerStakeChangeIterator, error) {

	var stakerIdRule []interface{}
	for _, stakerIdItem := range stakerId {
		stakerIdRule = append(stakerIdRule, stakerIdItem)
	}

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "StakeChange", stakerIdRule)
	if err != nil {
		return nil, err
	}
	return &StakeManagerStakeChangeIterator{contract: _StakeManager.contract, event: "StakeChange", logs: logs, sub: sub}, nil
}

// WatchStakeChange is a free log subscription operation binding the contract event 0xab9982d99de79485b3cd6d1e72c881fd0342f42f1cbc7d220d4f373b97c4f363.
//
// Solidity: event StakeChange(uint256 indexed stakerId, uint256 previousStake, uint256 newStake, string reason, uint256 epoch, uint256 timestamp)
func (_StakeManager *StakeManagerFilterer) WatchStakeChange(opts *bind.WatchOpts, sink chan<- *StakeManagerStakeChange, stakerId []*big.Int) (event.Subscription, error) {

	var stakerIdRule []interface{}
	for _, stakerIdItem := range stakerId {
		stakerIdRule = append(stakerIdRule, stakerIdItem)
	}

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "StakeChange", stakerIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerStakeChange)
				if err := _StakeManager.contract.UnpackLog(event, "StakeChange", log); err != nil {
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

// ParseStakeChange is a log parse operation binding the contract event 0xab9982d99de79485b3cd6d1e72c881fd0342f42f1cbc7d220d4f373b97c4f363.
//
// Solidity: event StakeChange(uint256 indexed stakerId, uint256 previousStake, uint256 newStake, string reason, uint256 epoch, uint256 timestamp)
func (_StakeManager *StakeManagerFilterer) ParseStakeChange(log types.Log) (*StakeManagerStakeChange, error) {
	event := new(StakeManagerStakeChange)
	if err := _StakeManager.contract.UnpackLog(event, "StakeChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the StakeManager contract.
type StakeManagerStakedIterator struct {
	Event *StakeManagerStaked // Event containing the contract specifics and raw log

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
func (it *StakeManagerStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerStaked)
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
		it.Event = new(StakeManagerStaked)
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
func (it *StakeManagerStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerStaked represents a Staked event raised by the StakeManager contract.
type StakeManagerStaked struct {
	Epoch         *big.Int
	StakerId      *big.Int
	PreviousStake *big.Int
	NewStake      *big.Int
	Timestamp     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x52d3b6bf695a499d39fcdb47e3b450c09f2f6aa091ca8809cc51c00e705996cc.
//
// Solidity: event Staked(uint256 epoch, uint256 indexed stakerId, uint256 previousStake, uint256 newStake, uint256 timestamp)
func (_StakeManager *StakeManagerFilterer) FilterStaked(opts *bind.FilterOpts, stakerId []*big.Int) (*StakeManagerStakedIterator, error) {

	var stakerIdRule []interface{}
	for _, stakerIdItem := range stakerId {
		stakerIdRule = append(stakerIdRule, stakerIdItem)
	}

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "Staked", stakerIdRule)
	if err != nil {
		return nil, err
	}
	return &StakeManagerStakedIterator{contract: _StakeManager.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x52d3b6bf695a499d39fcdb47e3b450c09f2f6aa091ca8809cc51c00e705996cc.
//
// Solidity: event Staked(uint256 epoch, uint256 indexed stakerId, uint256 previousStake, uint256 newStake, uint256 timestamp)
func (_StakeManager *StakeManagerFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *StakeManagerStaked, stakerId []*big.Int) (event.Subscription, error) {

	var stakerIdRule []interface{}
	for _, stakerIdItem := range stakerId {
		stakerIdRule = append(stakerIdRule, stakerIdItem)
	}

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "Staked", stakerIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerStaked)
				if err := _StakeManager.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x52d3b6bf695a499d39fcdb47e3b450c09f2f6aa091ca8809cc51c00e705996cc.
//
// Solidity: event Staked(uint256 epoch, uint256 indexed stakerId, uint256 previousStake, uint256 newStake, uint256 timestamp)
func (_StakeManager *StakeManagerFilterer) ParseStaked(log types.Log) (*StakeManagerStaked, error) {
	event := new(StakeManagerStaked)
	if err := _StakeManager.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the StakeManager contract.
type StakeManagerUnstakedIterator struct {
	Event *StakeManagerUnstaked // Event containing the contract specifics and raw log

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
func (it *StakeManagerUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerUnstaked)
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
		it.Event = new(StakeManagerUnstaked)
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
func (it *StakeManagerUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerUnstaked represents a Unstaked event raised by the StakeManager contract.
type StakeManagerUnstaked struct {
	Epoch     *big.Int
	StakerId  *big.Int
	Amount    *big.Int
	NewStake  *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0x0cfbc484edc798d2584502ca0d64e7e9514b8dd091d96a2a5b4deb58478da19e.
//
// Solidity: event Unstaked(uint256 epoch, uint256 indexed stakerId, uint256 amount, uint256 newStake, uint256 timestamp)
func (_StakeManager *StakeManagerFilterer) FilterUnstaked(opts *bind.FilterOpts, stakerId []*big.Int) (*StakeManagerUnstakedIterator, error) {

	var stakerIdRule []interface{}
	for _, stakerIdItem := range stakerId {
		stakerIdRule = append(stakerIdRule, stakerIdItem)
	}

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "Unstaked", stakerIdRule)
	if err != nil {
		return nil, err
	}
	return &StakeManagerUnstakedIterator{contract: _StakeManager.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0x0cfbc484edc798d2584502ca0d64e7e9514b8dd091d96a2a5b4deb58478da19e.
//
// Solidity: event Unstaked(uint256 epoch, uint256 indexed stakerId, uint256 amount, uint256 newStake, uint256 timestamp)
func (_StakeManager *StakeManagerFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *StakeManagerUnstaked, stakerId []*big.Int) (event.Subscription, error) {

	var stakerIdRule []interface{}
	for _, stakerIdItem := range stakerId {
		stakerIdRule = append(stakerIdRule, stakerIdItem)
	}

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "Unstaked", stakerIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerUnstaked)
				if err := _StakeManager.contract.UnpackLog(event, "Unstaked", log); err != nil {
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

// ParseUnstaked is a log parse operation binding the contract event 0x0cfbc484edc798d2584502ca0d64e7e9514b8dd091d96a2a5b4deb58478da19e.
//
// Solidity: event Unstaked(uint256 epoch, uint256 indexed stakerId, uint256 amount, uint256 newStake, uint256 timestamp)
func (_StakeManager *StakeManagerFilterer) ParseUnstaked(log types.Log) (*StakeManagerUnstaked, error) {
	event := new(StakeManagerUnstaked)
	if err := _StakeManager.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerWithdrewIterator is returned from FilterWithdrew and is used to iterate over the raw logs and unpacked data for Withdrew events raised by the StakeManager contract.
type StakeManagerWithdrewIterator struct {
	Event *StakeManagerWithdrew // Event containing the contract specifics and raw log

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
func (it *StakeManagerWithdrewIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerWithdrew)
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
		it.Event = new(StakeManagerWithdrew)
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
func (it *StakeManagerWithdrewIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerWithdrewIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerWithdrew represents a Withdrew event raised by the StakeManager contract.
type StakeManagerWithdrew struct {
	Epoch     *big.Int
	StakerId  *big.Int
	Amount    *big.Int
	NewStake  *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrew is a free log retrieval operation binding the contract event 0x0a6fecee2c95fc6b4c7f291e3435a3a832bcd700c2a7ffdedd8c909d56dfa49f.
//
// Solidity: event Withdrew(uint256 epoch, uint256 indexed stakerId, uint256 amount, uint256 newStake, uint256 timestamp)
func (_StakeManager *StakeManagerFilterer) FilterWithdrew(opts *bind.FilterOpts, stakerId []*big.Int) (*StakeManagerWithdrewIterator, error) {

	var stakerIdRule []interface{}
	for _, stakerIdItem := range stakerId {
		stakerIdRule = append(stakerIdRule, stakerIdItem)
	}

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "Withdrew", stakerIdRule)
	if err != nil {
		return nil, err
	}
	return &StakeManagerWithdrewIterator{contract: _StakeManager.contract, event: "Withdrew", logs: logs, sub: sub}, nil
}

// WatchWithdrew is a free log subscription operation binding the contract event 0x0a6fecee2c95fc6b4c7f291e3435a3a832bcd700c2a7ffdedd8c909d56dfa49f.
//
// Solidity: event Withdrew(uint256 epoch, uint256 indexed stakerId, uint256 amount, uint256 newStake, uint256 timestamp)
func (_StakeManager *StakeManagerFilterer) WatchWithdrew(opts *bind.WatchOpts, sink chan<- *StakeManagerWithdrew, stakerId []*big.Int) (event.Subscription, error) {

	var stakerIdRule []interface{}
	for _, stakerIdItem := range stakerId {
		stakerIdRule = append(stakerIdRule, stakerIdItem)
	}

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "Withdrew", stakerIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerWithdrew)
				if err := _StakeManager.contract.UnpackLog(event, "Withdrew", log); err != nil {
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

// ParseWithdrew is a log parse operation binding the contract event 0x0a6fecee2c95fc6b4c7f291e3435a3a832bcd700c2a7ffdedd8c909d56dfa49f.
//
// Solidity: event Withdrew(uint256 epoch, uint256 indexed stakerId, uint256 amount, uint256 newStake, uint256 timestamp)
func (_StakeManager *StakeManagerFilterer) ParseWithdrew(log types.Log) (*StakeManagerWithdrew, error) {
	event := new(StakeManagerWithdrew)
	if err := _StakeManager.contract.UnpackLog(event, "Withdrew", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
