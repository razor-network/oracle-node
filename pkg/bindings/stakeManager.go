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
	UnstakeAfter       *big.Int
	WithdrawAfter      *big.Int
}

// StakeManagerABI is the input ABI used to generate the binding from.
const StakeManagerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"DebugUint256\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevRewardPool\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardPool\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"RewardPoolChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousStake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newStake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"StakeChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevStakeGettingReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakeGettingReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"StakeGettingRewardChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousStake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newStake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newStake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newStake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Withdrew\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockManager\",\"outputs\":[{\"internalType\":\"contractIBlockManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeValue\",\"type\":\"uint256\"}],\"name\":\"calculateInactivityPenalties\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"genesisBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumStakers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakeGettingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getStaker\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochStaked\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochLastCommitted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochLastRevealed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeAfter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAfter\",\"type\":\"uint256\"}],\"internalType\":\"structStructs.Staker\",\"name\":\"staker\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getStakerId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"giveBlockReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"givePenalties\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"giveRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_schAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_voteManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_blockManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stateManagerAddress\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBlockRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastHalvings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStakers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sch\",\"outputs\":[{\"internalType\":\"contractSchellingCoin\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_epochLastRevealed\",\"type\":\"uint256\"}],\"name\":\"setStakerEpochLastRevealed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bountyHunter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeGettingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakerIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochStaked\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochLastCommitted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochLastRevealed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeAfter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAfter\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stateManager\",\"outputs\":[{\"internalType\":\"contractIStateManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerId\",\"type\":\"uint256\"}],\"name\":\"updateCommitmentEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteManager\",\"outputs\":[{\"internalType\":\"contractIVoteManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// BlockManager is a free data retrieval call binding the contract method 0xd9169b32.
//
// Solidity: function blockManager() view returns(address)
func (_StakeManager *StakeManagerCaller) BlockManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "blockManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlockManager is a free data retrieval call binding the contract method 0xd9169b32.
//
// Solidity: function blockManager() view returns(address)
func (_StakeManager *StakeManagerSession) BlockManager() (common.Address, error) {
	return _StakeManager.Contract.BlockManager(&_StakeManager.CallOpts)
}

// BlockManager is a free data retrieval call binding the contract method 0xd9169b32.
//
// Solidity: function blockManager() view returns(address)
func (_StakeManager *StakeManagerCallerSession) BlockManager() (common.Address, error) {
	return _StakeManager.Contract.BlockManager(&_StakeManager.CallOpts)
}

// CalculateInactivityPenalties is a free data retrieval call binding the contract method 0x1a69f008.
//
// Solidity: function calculateInactivityPenalties(uint256 epochs, uint256 stakeValue) pure returns(uint256)
func (_StakeManager *StakeManagerCaller) CalculateInactivityPenalties(opts *bind.CallOpts, epochs *big.Int, stakeValue *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "calculateInactivityPenalties", epochs, stakeValue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateInactivityPenalties is a free data retrieval call binding the contract method 0x1a69f008.
//
// Solidity: function calculateInactivityPenalties(uint256 epochs, uint256 stakeValue) pure returns(uint256)
func (_StakeManager *StakeManagerSession) CalculateInactivityPenalties(epochs *big.Int, stakeValue *big.Int) (*big.Int, error) {
	return _StakeManager.Contract.CalculateInactivityPenalties(&_StakeManager.CallOpts, epochs, stakeValue)
}

// CalculateInactivityPenalties is a free data retrieval call binding the contract method 0x1a69f008.
//
// Solidity: function calculateInactivityPenalties(uint256 epochs, uint256 stakeValue) pure returns(uint256)
func (_StakeManager *StakeManagerCallerSession) CalculateInactivityPenalties(epochs *big.Int, stakeValue *big.Int) (*big.Int, error) {
	return _StakeManager.Contract.CalculateInactivityPenalties(&_StakeManager.CallOpts, epochs, stakeValue)
}

// GenesisBlock is a free data retrieval call binding the contract method 0x4cdc9c63.
//
// Solidity: function genesisBlock() view returns(uint256)
func (_StakeManager *StakeManagerCaller) GenesisBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "genesisBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GenesisBlock is a free data retrieval call binding the contract method 0x4cdc9c63.
//
// Solidity: function genesisBlock() view returns(uint256)
func (_StakeManager *StakeManagerSession) GenesisBlock() (*big.Int, error) {
	return _StakeManager.Contract.GenesisBlock(&_StakeManager.CallOpts)
}

// GenesisBlock is a free data retrieval call binding the contract method 0x4cdc9c63.
//
// Solidity: function genesisBlock() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) GenesisBlock() (*big.Int, error) {
	return _StakeManager.Contract.GenesisBlock(&_StakeManager.CallOpts)
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

// GetRewardPool is a free data retrieval call binding the contract method 0x1b8b13a7.
//
// Solidity: function getRewardPool() view returns(uint256)
func (_StakeManager *StakeManagerCaller) GetRewardPool(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "getRewardPool")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardPool is a free data retrieval call binding the contract method 0x1b8b13a7.
//
// Solidity: function getRewardPool() view returns(uint256)
func (_StakeManager *StakeManagerSession) GetRewardPool() (*big.Int, error) {
	return _StakeManager.Contract.GetRewardPool(&_StakeManager.CallOpts)
}

// GetRewardPool is a free data retrieval call binding the contract method 0x1b8b13a7.
//
// Solidity: function getRewardPool() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) GetRewardPool() (*big.Int, error) {
	return _StakeManager.Contract.GetRewardPool(&_StakeManager.CallOpts)
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

// GetStakeGettingReward is a free data retrieval call binding the contract method 0x1ad54991.
//
// Solidity: function getStakeGettingReward() view returns(uint256)
func (_StakeManager *StakeManagerCaller) GetStakeGettingReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "getStakeGettingReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakeGettingReward is a free data retrieval call binding the contract method 0x1ad54991.
//
// Solidity: function getStakeGettingReward() view returns(uint256)
func (_StakeManager *StakeManagerSession) GetStakeGettingReward() (*big.Int, error) {
	return _StakeManager.Contract.GetStakeGettingReward(&_StakeManager.CallOpts)
}

// GetStakeGettingReward is a free data retrieval call binding the contract method 0x1ad54991.
//
// Solidity: function getStakeGettingReward() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) GetStakeGettingReward() (*big.Int, error) {
	return _StakeManager.Contract.GetStakeGettingReward(&_StakeManager.CallOpts)
}

// GetStaker is a free data retrieval call binding the contract method 0xe3c998fe.
//
// Solidity: function getStaker(uint256 _id) view returns((uint256,address,uint256,uint256,uint256,uint256,uint256,uint256) staker)
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
// Solidity: function getStaker(uint256 _id) view returns((uint256,address,uint256,uint256,uint256,uint256,uint256,uint256) staker)
func (_StakeManager *StakeManagerSession) GetStaker(_id *big.Int) (StructsStaker, error) {
	return _StakeManager.Contract.GetStaker(&_StakeManager.CallOpts, _id)
}

// GetStaker is a free data retrieval call binding the contract method 0xe3c998fe.
//
// Solidity: function getStaker(uint256 _id) view returns((uint256,address,uint256,uint256,uint256,uint256,uint256,uint256) staker)
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

// LastBlockRewards is a free data retrieval call binding the contract method 0x2f594949.
//
// Solidity: function lastBlockRewards() view returns(uint256)
func (_StakeManager *StakeManagerCaller) LastBlockRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "lastBlockRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBlockRewards is a free data retrieval call binding the contract method 0x2f594949.
//
// Solidity: function lastBlockRewards() view returns(uint256)
func (_StakeManager *StakeManagerSession) LastBlockRewards() (*big.Int, error) {
	return _StakeManager.Contract.LastBlockRewards(&_StakeManager.CallOpts)
}

// LastBlockRewards is a free data retrieval call binding the contract method 0x2f594949.
//
// Solidity: function lastBlockRewards() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) LastBlockRewards() (*big.Int, error) {
	return _StakeManager.Contract.LastBlockRewards(&_StakeManager.CallOpts)
}

// LastHalvings is a free data retrieval call binding the contract method 0xa8264f99.
//
// Solidity: function lastHalvings() view returns(uint256)
func (_StakeManager *StakeManagerCaller) LastHalvings(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "lastHalvings")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastHalvings is a free data retrieval call binding the contract method 0xa8264f99.
//
// Solidity: function lastHalvings() view returns(uint256)
func (_StakeManager *StakeManagerSession) LastHalvings() (*big.Int, error) {
	return _StakeManager.Contract.LastHalvings(&_StakeManager.CallOpts)
}

// LastHalvings is a free data retrieval call binding the contract method 0xa8264f99.
//
// Solidity: function lastHalvings() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) LastHalvings() (*big.Int, error) {
	return _StakeManager.Contract.LastHalvings(&_StakeManager.CallOpts)
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

// RewardPool is a free data retrieval call binding the contract method 0x66666aa9.
//
// Solidity: function rewardPool() view returns(uint256)
func (_StakeManager *StakeManagerCaller) RewardPool(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "rewardPool")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPool is a free data retrieval call binding the contract method 0x66666aa9.
//
// Solidity: function rewardPool() view returns(uint256)
func (_StakeManager *StakeManagerSession) RewardPool() (*big.Int, error) {
	return _StakeManager.Contract.RewardPool(&_StakeManager.CallOpts)
}

// RewardPool is a free data retrieval call binding the contract method 0x66666aa9.
//
// Solidity: function rewardPool() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) RewardPool() (*big.Int, error) {
	return _StakeManager.Contract.RewardPool(&_StakeManager.CallOpts)
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

// StakeGettingReward is a free data retrieval call binding the contract method 0x0ec88d3f.
//
// Solidity: function stakeGettingReward() view returns(uint256)
func (_StakeManager *StakeManagerCaller) StakeGettingReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "stakeGettingReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeGettingReward is a free data retrieval call binding the contract method 0x0ec88d3f.
//
// Solidity: function stakeGettingReward() view returns(uint256)
func (_StakeManager *StakeManagerSession) StakeGettingReward() (*big.Int, error) {
	return _StakeManager.Contract.StakeGettingReward(&_StakeManager.CallOpts)
}

// StakeGettingReward is a free data retrieval call binding the contract method 0x0ec88d3f.
//
// Solidity: function stakeGettingReward() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) StakeGettingReward() (*big.Int, error) {
	return _StakeManager.Contract.StakeGettingReward(&_StakeManager.CallOpts)
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
// Solidity: function stakers(uint256 ) view returns(uint256 id, address _address, uint256 stake, uint256 epochStaked, uint256 epochLastCommitted, uint256 epochLastRevealed, uint256 unstakeAfter, uint256 withdrawAfter)
func (_StakeManager *StakeManagerCaller) Stakers(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _StakeManager.contract.Call(opts, &out, "stakers", arg0)

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
func (_StakeManager *StakeManagerSession) Stakers(arg0 *big.Int) (struct {
	Id                 *big.Int
	Address            common.Address
	Stake              *big.Int
	EpochStaked        *big.Int
	EpochLastCommitted *big.Int
	EpochLastRevealed  *big.Int
	UnstakeAfter       *big.Int
	WithdrawAfter      *big.Int
}, error) {
	return _StakeManager.Contract.Stakers(&_StakeManager.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) view returns(uint256 id, address _address, uint256 stake, uint256 epochStaked, uint256 epochLastCommitted, uint256 epochLastRevealed, uint256 unstakeAfter, uint256 withdrawAfter)
func (_StakeManager *StakeManagerCallerSession) Stakers(arg0 *big.Int) (struct {
	Id                 *big.Int
	Address            common.Address
	Stake              *big.Int
	EpochStaked        *big.Int
	EpochLastCommitted *big.Int
	EpochLastRevealed  *big.Int
	UnstakeAfter       *big.Int
	WithdrawAfter      *big.Int
}, error) {
	return _StakeManager.Contract.Stakers(&_StakeManager.CallOpts, arg0)
}

// StateManager is a free data retrieval call binding the contract method 0x2e716fb1.
//
// Solidity: function stateManager() view returns(address)
func (_StakeManager *StakeManagerCaller) StateManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "stateManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StateManager is a free data retrieval call binding the contract method 0x2e716fb1.
//
// Solidity: function stateManager() view returns(address)
func (_StakeManager *StakeManagerSession) StateManager() (common.Address, error) {
	return _StakeManager.Contract.StateManager(&_StakeManager.CallOpts)
}

// StateManager is a free data retrieval call binding the contract method 0x2e716fb1.
//
// Solidity: function stateManager() view returns(address)
func (_StakeManager *StakeManagerCallerSession) StateManager() (common.Address, error) {
	return _StakeManager.Contract.StateManager(&_StakeManager.CallOpts)
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

// GiveBlockReward is a paid mutator transaction binding the contract method 0x746c8b65.
//
// Solidity: function giveBlockReward(uint256 stakerId, uint256 epoch) returns()
func (_StakeManager *StakeManagerTransactor) GiveBlockReward(opts *bind.TransactOpts, stakerId *big.Int, epoch *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "giveBlockReward", stakerId, epoch)
}

// GiveBlockReward is a paid mutator transaction binding the contract method 0x746c8b65.
//
// Solidity: function giveBlockReward(uint256 stakerId, uint256 epoch) returns()
func (_StakeManager *StakeManagerSession) GiveBlockReward(stakerId *big.Int, epoch *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.GiveBlockReward(&_StakeManager.TransactOpts, stakerId, epoch)
}

// GiveBlockReward is a paid mutator transaction binding the contract method 0x746c8b65.
//
// Solidity: function giveBlockReward(uint256 stakerId, uint256 epoch) returns()
func (_StakeManager *StakeManagerTransactorSession) GiveBlockReward(stakerId *big.Int, epoch *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.GiveBlockReward(&_StakeManager.TransactOpts, stakerId, epoch)
}

// GivePenalties is a paid mutator transaction binding the contract method 0x54eae15e.
//
// Solidity: function givePenalties(uint256 stakerId, uint256 epoch) returns()
func (_StakeManager *StakeManagerTransactor) GivePenalties(opts *bind.TransactOpts, stakerId *big.Int, epoch *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "givePenalties", stakerId, epoch)
}

// GivePenalties is a paid mutator transaction binding the contract method 0x54eae15e.
//
// Solidity: function givePenalties(uint256 stakerId, uint256 epoch) returns()
func (_StakeManager *StakeManagerSession) GivePenalties(stakerId *big.Int, epoch *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.GivePenalties(&_StakeManager.TransactOpts, stakerId, epoch)
}

// GivePenalties is a paid mutator transaction binding the contract method 0x54eae15e.
//
// Solidity: function givePenalties(uint256 stakerId, uint256 epoch) returns()
func (_StakeManager *StakeManagerTransactorSession) GivePenalties(stakerId *big.Int, epoch *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.GivePenalties(&_StakeManager.TransactOpts, stakerId, epoch)
}

// GiveRewards is a paid mutator transaction binding the contract method 0xfabb9890.
//
// Solidity: function giveRewards(uint256 stakerId, uint256 epoch) returns()
func (_StakeManager *StakeManagerTransactor) GiveRewards(opts *bind.TransactOpts, stakerId *big.Int, epoch *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "giveRewards", stakerId, epoch)
}

// GiveRewards is a paid mutator transaction binding the contract method 0xfabb9890.
//
// Solidity: function giveRewards(uint256 stakerId, uint256 epoch) returns()
func (_StakeManager *StakeManagerSession) GiveRewards(stakerId *big.Int, epoch *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.GiveRewards(&_StakeManager.TransactOpts, stakerId, epoch)
}

// GiveRewards is a paid mutator transaction binding the contract method 0xfabb9890.
//
// Solidity: function giveRewards(uint256 stakerId, uint256 epoch) returns()
func (_StakeManager *StakeManagerTransactorSession) GiveRewards(stakerId *big.Int, epoch *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.GiveRewards(&_StakeManager.TransactOpts, stakerId, epoch)
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

// Init is a paid mutator transaction binding the contract method 0x06552ff3.
//
// Solidity: function init(address _schAddress, address _voteManagerAddress, address _blockManagerAddress, address _stateManagerAddress) returns()
func (_StakeManager *StakeManagerTransactor) Init(opts *bind.TransactOpts, _schAddress common.Address, _voteManagerAddress common.Address, _blockManagerAddress common.Address, _stateManagerAddress common.Address) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "init", _schAddress, _voteManagerAddress, _blockManagerAddress, _stateManagerAddress)
}

// Init is a paid mutator transaction binding the contract method 0x06552ff3.
//
// Solidity: function init(address _schAddress, address _voteManagerAddress, address _blockManagerAddress, address _stateManagerAddress) returns()
func (_StakeManager *StakeManagerSession) Init(_schAddress common.Address, _voteManagerAddress common.Address, _blockManagerAddress common.Address, _stateManagerAddress common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.Init(&_StakeManager.TransactOpts, _schAddress, _voteManagerAddress, _blockManagerAddress, _stateManagerAddress)
}

// Init is a paid mutator transaction binding the contract method 0x06552ff3.
//
// Solidity: function init(address _schAddress, address _voteManagerAddress, address _blockManagerAddress, address _stateManagerAddress) returns()
func (_StakeManager *StakeManagerTransactorSession) Init(_schAddress common.Address, _voteManagerAddress common.Address, _blockManagerAddress common.Address, _stateManagerAddress common.Address) (*types.Transaction, error) {
	return _StakeManager.Contract.Init(&_StakeManager.TransactOpts, _schAddress, _voteManagerAddress, _blockManagerAddress, _stateManagerAddress)
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

// Slash is a paid mutator transaction binding the contract method 0x0f91ce19.
//
// Solidity: function slash(uint256 id, address bountyHunter, uint256 epoch) returns()
func (_StakeManager *StakeManagerTransactor) Slash(opts *bind.TransactOpts, id *big.Int, bountyHunter common.Address, epoch *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "slash", id, bountyHunter, epoch)
}

// Slash is a paid mutator transaction binding the contract method 0x0f91ce19.
//
// Solidity: function slash(uint256 id, address bountyHunter, uint256 epoch) returns()
func (_StakeManager *StakeManagerSession) Slash(id *big.Int, bountyHunter common.Address, epoch *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Slash(&_StakeManager.TransactOpts, id, bountyHunter, epoch)
}

// Slash is a paid mutator transaction binding the contract method 0x0f91ce19.
//
// Solidity: function slash(uint256 id, address bountyHunter, uint256 epoch) returns()
func (_StakeManager *StakeManagerTransactorSession) Slash(id *big.Int, bountyHunter common.Address, epoch *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Slash(&_StakeManager.TransactOpts, id, bountyHunter, epoch)
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

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 epoch) returns()
func (_StakeManager *StakeManagerTransactor) Unstake(opts *bind.TransactOpts, epoch *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "unstake", epoch)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 epoch) returns()
func (_StakeManager *StakeManagerSession) Unstake(epoch *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Unstake(&_StakeManager.TransactOpts, epoch)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 epoch) returns()
func (_StakeManager *StakeManagerTransactorSession) Unstake(epoch *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Unstake(&_StakeManager.TransactOpts, epoch)
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

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 epoch) returns()
func (_StakeManager *StakeManagerTransactor) Withdraw(opts *bind.TransactOpts, epoch *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "withdraw", epoch)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 epoch) returns()
func (_StakeManager *StakeManagerSession) Withdraw(epoch *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Withdraw(&_StakeManager.TransactOpts, epoch)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 epoch) returns()
func (_StakeManager *StakeManagerTransactorSession) Withdraw(epoch *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Withdraw(&_StakeManager.TransactOpts, epoch)
}

// StakeManagerDebugUint256Iterator is returned from FilterDebugUint256 and is used to iterate over the raw logs and unpacked data for DebugUint256 events raised by the StakeManager contract.
type StakeManagerDebugUint256Iterator struct {
	Event *StakeManagerDebugUint256 // Event containing the contract specifics and raw log

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
func (it *StakeManagerDebugUint256Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerDebugUint256)
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
		it.Event = new(StakeManagerDebugUint256)
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
func (it *StakeManagerDebugUint256Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerDebugUint256Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerDebugUint256 represents a DebugUint256 event raised by the StakeManager contract.
type StakeManagerDebugUint256 struct {
	A   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDebugUint256 is a free log retrieval operation binding the contract event 0x43d4b4706539f9e22baf8767ebea21ad24f723f14b6981664ac4d0af596dddbe.
//
// Solidity: event DebugUint256(uint256 a)
func (_StakeManager *StakeManagerFilterer) FilterDebugUint256(opts *bind.FilterOpts) (*StakeManagerDebugUint256Iterator, error) {

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "DebugUint256")
	if err != nil {
		return nil, err
	}
	return &StakeManagerDebugUint256Iterator{contract: _StakeManager.contract, event: "DebugUint256", logs: logs, sub: sub}, nil
}

// WatchDebugUint256 is a free log subscription operation binding the contract event 0x43d4b4706539f9e22baf8767ebea21ad24f723f14b6981664ac4d0af596dddbe.
//
// Solidity: event DebugUint256(uint256 a)
func (_StakeManager *StakeManagerFilterer) WatchDebugUint256(opts *bind.WatchOpts, sink chan<- *StakeManagerDebugUint256) (event.Subscription, error) {

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "DebugUint256")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerDebugUint256)
				if err := _StakeManager.contract.UnpackLog(event, "DebugUint256", log); err != nil {
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

// ParseDebugUint256 is a log parse operation binding the contract event 0x43d4b4706539f9e22baf8767ebea21ad24f723f14b6981664ac4d0af596dddbe.
//
// Solidity: event DebugUint256(uint256 a)
func (_StakeManager *StakeManagerFilterer) ParseDebugUint256(log types.Log) (*StakeManagerDebugUint256, error) {
	event := new(StakeManagerDebugUint256)
	if err := _StakeManager.contract.UnpackLog(event, "DebugUint256", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerRewardPoolChangeIterator is returned from FilterRewardPoolChange and is used to iterate over the raw logs and unpacked data for RewardPoolChange events raised by the StakeManager contract.
type StakeManagerRewardPoolChangeIterator struct {
	Event *StakeManagerRewardPoolChange // Event containing the contract specifics and raw log

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
func (it *StakeManagerRewardPoolChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerRewardPoolChange)
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
		it.Event = new(StakeManagerRewardPoolChange)
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
func (it *StakeManagerRewardPoolChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerRewardPoolChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerRewardPoolChange represents a RewardPoolChange event raised by the StakeManager contract.
type StakeManagerRewardPoolChange struct {
	Epoch          *big.Int
	PrevRewardPool *big.Int
	RewardPool     *big.Int
	Timestamp      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRewardPoolChange is a free log retrieval operation binding the contract event 0xf7157643fd549e213a0105625da6b3bf58c86068ccb954a4449e18bf9427bff4.
//
// Solidity: event RewardPoolChange(uint256 epoch, uint256 prevRewardPool, uint256 rewardPool, uint256 timestamp)
func (_StakeManager *StakeManagerFilterer) FilterRewardPoolChange(opts *bind.FilterOpts) (*StakeManagerRewardPoolChangeIterator, error) {

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "RewardPoolChange")
	if err != nil {
		return nil, err
	}
	return &StakeManagerRewardPoolChangeIterator{contract: _StakeManager.contract, event: "RewardPoolChange", logs: logs, sub: sub}, nil
}

// WatchRewardPoolChange is a free log subscription operation binding the contract event 0xf7157643fd549e213a0105625da6b3bf58c86068ccb954a4449e18bf9427bff4.
//
// Solidity: event RewardPoolChange(uint256 epoch, uint256 prevRewardPool, uint256 rewardPool, uint256 timestamp)
func (_StakeManager *StakeManagerFilterer) WatchRewardPoolChange(opts *bind.WatchOpts, sink chan<- *StakeManagerRewardPoolChange) (event.Subscription, error) {

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "RewardPoolChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerRewardPoolChange)
				if err := _StakeManager.contract.UnpackLog(event, "RewardPoolChange", log); err != nil {
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

// ParseRewardPoolChange is a log parse operation binding the contract event 0xf7157643fd549e213a0105625da6b3bf58c86068ccb954a4449e18bf9427bff4.
//
// Solidity: event RewardPoolChange(uint256 epoch, uint256 prevRewardPool, uint256 rewardPool, uint256 timestamp)
func (_StakeManager *StakeManagerFilterer) ParseRewardPoolChange(log types.Log) (*StakeManagerRewardPoolChange, error) {
	event := new(StakeManagerRewardPoolChange)
	if err := _StakeManager.contract.UnpackLog(event, "RewardPoolChange", log); err != nil {
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

// StakeManagerStakeGettingRewardChangeIterator is returned from FilterStakeGettingRewardChange and is used to iterate over the raw logs and unpacked data for StakeGettingRewardChange events raised by the StakeManager contract.
type StakeManagerStakeGettingRewardChangeIterator struct {
	Event *StakeManagerStakeGettingRewardChange // Event containing the contract specifics and raw log

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
func (it *StakeManagerStakeGettingRewardChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerStakeGettingRewardChange)
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
		it.Event = new(StakeManagerStakeGettingRewardChange)
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
func (it *StakeManagerStakeGettingRewardChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerStakeGettingRewardChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerStakeGettingRewardChange represents a StakeGettingRewardChange event raised by the StakeManager contract.
type StakeManagerStakeGettingRewardChange struct {
	Epoch                  *big.Int
	PrevStakeGettingReward *big.Int
	StakeGettingReward     *big.Int
	Timestamp              *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterStakeGettingRewardChange is a free log retrieval operation binding the contract event 0xeee26d3c1e406e24904d3748e49da85096c92ac1a6ad38f3fe8a404f85212dd1.
//
// Solidity: event StakeGettingRewardChange(uint256 epoch, uint256 prevStakeGettingReward, uint256 stakeGettingReward, uint256 timestamp)
func (_StakeManager *StakeManagerFilterer) FilterStakeGettingRewardChange(opts *bind.FilterOpts) (*StakeManagerStakeGettingRewardChangeIterator, error) {

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "StakeGettingRewardChange")
	if err != nil {
		return nil, err
	}
	return &StakeManagerStakeGettingRewardChangeIterator{contract: _StakeManager.contract, event: "StakeGettingRewardChange", logs: logs, sub: sub}, nil
}

// WatchStakeGettingRewardChange is a free log subscription operation binding the contract event 0xeee26d3c1e406e24904d3748e49da85096c92ac1a6ad38f3fe8a404f85212dd1.
//
// Solidity: event StakeGettingRewardChange(uint256 epoch, uint256 prevStakeGettingReward, uint256 stakeGettingReward, uint256 timestamp)
func (_StakeManager *StakeManagerFilterer) WatchStakeGettingRewardChange(opts *bind.WatchOpts, sink chan<- *StakeManagerStakeGettingRewardChange) (event.Subscription, error) {

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "StakeGettingRewardChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerStakeGettingRewardChange)
				if err := _StakeManager.contract.UnpackLog(event, "StakeGettingRewardChange", log); err != nil {
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

// ParseStakeGettingRewardChange is a log parse operation binding the contract event 0xeee26d3c1e406e24904d3748e49da85096c92ac1a6ad38f3fe8a404f85212dd1.
//
// Solidity: event StakeGettingRewardChange(uint256 epoch, uint256 prevStakeGettingReward, uint256 stakeGettingReward, uint256 timestamp)
func (_StakeManager *StakeManagerFilterer) ParseStakeGettingRewardChange(log types.Log) (*StakeManagerStakeGettingRewardChange, error) {
	event := new(StakeManagerStakeGettingRewardChange)
	if err := _StakeManager.contract.UnpackLog(event, "StakeGettingRewardChange", log); err != nil {
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
