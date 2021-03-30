// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IStakeManager

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

// IStakeManagerABI is the input ABI used to generate the binding from.
const IStakeManagerABI = "[{\"inputs\":[],\"name\":\"getNumStakers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakeGettingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getStaker\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochStaked\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochLastCommitted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochLastRevealed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeAfter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAfter\",\"type\":\"uint256\"}],\"internalType\":\"structStructs.Staker\",\"name\":\"staker\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getStakerId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"giveBlockReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"givePenalties\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"giveRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_schAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_voteManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_blockManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stateManagerAddress\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_epochLastRevealed\",\"type\":\"uint256\"}],\"name\":\"setStakerEpochLastRevealed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bountyHunter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerId\",\"type\":\"uint256\"}],\"name\":\"updateCommitmentEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IStakeManager is an auto generated Go binding around an Ethereum contract.
type IStakeManager struct {
	IStakeManagerCaller     // Read-only binding to the contract
	IStakeManagerTransactor // Write-only binding to the contract
	IStakeManagerFilterer   // Log filterer for contract events
}

// IStakeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStakeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStakeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStakeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStakeManagerSession struct {
	Contract     *IStakeManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IStakeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStakeManagerCallerSession struct {
	Contract *IStakeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IStakeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStakeManagerTransactorSession struct {
	Contract     *IStakeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IStakeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStakeManagerRaw struct {
	Contract *IStakeManager // Generic contract binding to access the raw methods on
}

// IStakeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStakeManagerCallerRaw struct {
	Contract *IStakeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IStakeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStakeManagerTransactorRaw struct {
	Contract *IStakeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStakeManager creates a new instance of IStakeManager, bound to a specific deployed contract.
func NewIStakeManager(address common.Address, backend bind.ContractBackend) (*IStakeManager, error) {
	contract, err := bindIStakeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStakeManager{IStakeManagerCaller: IStakeManagerCaller{contract: contract}, IStakeManagerTransactor: IStakeManagerTransactor{contract: contract}, IStakeManagerFilterer: IStakeManagerFilterer{contract: contract}}, nil
}

// NewIStakeManagerCaller creates a new read-only instance of IStakeManager, bound to a specific deployed contract.
func NewIStakeManagerCaller(address common.Address, caller bind.ContractCaller) (*IStakeManagerCaller, error) {
	contract, err := bindIStakeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStakeManagerCaller{contract: contract}, nil
}

// NewIStakeManagerTransactor creates a new write-only instance of IStakeManager, bound to a specific deployed contract.
func NewIStakeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IStakeManagerTransactor, error) {
	contract, err := bindIStakeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStakeManagerTransactor{contract: contract}, nil
}

// NewIStakeManagerFilterer creates a new log filterer instance of IStakeManager, bound to a specific deployed contract.
func NewIStakeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IStakeManagerFilterer, error) {
	contract, err := bindIStakeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStakeManagerFilterer{contract: contract}, nil
}

// bindIStakeManager binds a generic wrapper to an already deployed contract.
func bindIStakeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IStakeManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStakeManager *IStakeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStakeManager.Contract.IStakeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStakeManager *IStakeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakeManager.Contract.IStakeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStakeManager *IStakeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStakeManager.Contract.IStakeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStakeManager *IStakeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStakeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStakeManager *IStakeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStakeManager *IStakeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStakeManager.Contract.contract.Transact(opts, method, params...)
}

// GetNumStakers is a free data retrieval call binding the contract method 0xbc788d46.
//
// Solidity: function getNumStakers() view returns(uint256)
func (_IStakeManager *IStakeManagerCaller) GetNumStakers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IStakeManager.contract.Call(opts, &out, "getNumStakers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumStakers is a free data retrieval call binding the contract method 0xbc788d46.
//
// Solidity: function getNumStakers() view returns(uint256)
func (_IStakeManager *IStakeManagerSession) GetNumStakers() (*big.Int, error) {
	return _IStakeManager.Contract.GetNumStakers(&_IStakeManager.CallOpts)
}

// GetNumStakers is a free data retrieval call binding the contract method 0xbc788d46.
//
// Solidity: function getNumStakers() view returns(uint256)
func (_IStakeManager *IStakeManagerCallerSession) GetNumStakers() (*big.Int, error) {
	return _IStakeManager.Contract.GetNumStakers(&_IStakeManager.CallOpts)
}

// GetRewardPool is a free data retrieval call binding the contract method 0x1b8b13a7.
//
// Solidity: function getRewardPool() view returns(uint256)
func (_IStakeManager *IStakeManagerCaller) GetRewardPool(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IStakeManager.contract.Call(opts, &out, "getRewardPool")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardPool is a free data retrieval call binding the contract method 0x1b8b13a7.
//
// Solidity: function getRewardPool() view returns(uint256)
func (_IStakeManager *IStakeManagerSession) GetRewardPool() (*big.Int, error) {
	return _IStakeManager.Contract.GetRewardPool(&_IStakeManager.CallOpts)
}

// GetRewardPool is a free data retrieval call binding the contract method 0x1b8b13a7.
//
// Solidity: function getRewardPool() view returns(uint256)
func (_IStakeManager *IStakeManagerCallerSession) GetRewardPool() (*big.Int, error) {
	return _IStakeManager.Contract.GetRewardPool(&_IStakeManager.CallOpts)
}

// GetStakeGettingReward is a free data retrieval call binding the contract method 0x1ad54991.
//
// Solidity: function getStakeGettingReward() view returns(uint256)
func (_IStakeManager *IStakeManagerCaller) GetStakeGettingReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IStakeManager.contract.Call(opts, &out, "getStakeGettingReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakeGettingReward is a free data retrieval call binding the contract method 0x1ad54991.
//
// Solidity: function getStakeGettingReward() view returns(uint256)
func (_IStakeManager *IStakeManagerSession) GetStakeGettingReward() (*big.Int, error) {
	return _IStakeManager.Contract.GetStakeGettingReward(&_IStakeManager.CallOpts)
}

// GetStakeGettingReward is a free data retrieval call binding the contract method 0x1ad54991.
//
// Solidity: function getStakeGettingReward() view returns(uint256)
func (_IStakeManager *IStakeManagerCallerSession) GetStakeGettingReward() (*big.Int, error) {
	return _IStakeManager.Contract.GetStakeGettingReward(&_IStakeManager.CallOpts)
}

// GetStaker is a free data retrieval call binding the contract method 0xe3c998fe.
//
// Solidity: function getStaker(uint256 _id) view returns((uint256,address,uint256,uint256,uint256,uint256,uint256,uint256) staker)
func (_IStakeManager *IStakeManagerCaller) GetStaker(opts *bind.CallOpts, _id *big.Int) (StructsStaker, error) {
	var out []interface{}
	err := _IStakeManager.contract.Call(opts, &out, "getStaker", _id)

	if err != nil {
		return *new(StructsStaker), err
	}

	out0 := *abi.ConvertType(out[0], new(StructsStaker)).(*StructsStaker)

	return out0, err

}

// GetStaker is a free data retrieval call binding the contract method 0xe3c998fe.
//
// Solidity: function getStaker(uint256 _id) view returns((uint256,address,uint256,uint256,uint256,uint256,uint256,uint256) staker)
func (_IStakeManager *IStakeManagerSession) GetStaker(_id *big.Int) (StructsStaker, error) {
	return _IStakeManager.Contract.GetStaker(&_IStakeManager.CallOpts, _id)
}

// GetStaker is a free data retrieval call binding the contract method 0xe3c998fe.
//
// Solidity: function getStaker(uint256 _id) view returns((uint256,address,uint256,uint256,uint256,uint256,uint256,uint256) staker)
func (_IStakeManager *IStakeManagerCallerSession) GetStaker(_id *big.Int) (StructsStaker, error) {
	return _IStakeManager.Contract.GetStaker(&_IStakeManager.CallOpts, _id)
}

// GetStakerId is a free data retrieval call binding the contract method 0x6022a485.
//
// Solidity: function getStakerId(address _address) view returns(uint256)
func (_IStakeManager *IStakeManagerCaller) GetStakerId(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IStakeManager.contract.Call(opts, &out, "getStakerId", _address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakerId is a free data retrieval call binding the contract method 0x6022a485.
//
// Solidity: function getStakerId(address _address) view returns(uint256)
func (_IStakeManager *IStakeManagerSession) GetStakerId(_address common.Address) (*big.Int, error) {
	return _IStakeManager.Contract.GetStakerId(&_IStakeManager.CallOpts, _address)
}

// GetStakerId is a free data retrieval call binding the contract method 0x6022a485.
//
// Solidity: function getStakerId(address _address) view returns(uint256)
func (_IStakeManager *IStakeManagerCallerSession) GetStakerId(_address common.Address) (*big.Int, error) {
	return _IStakeManager.Contract.GetStakerId(&_IStakeManager.CallOpts, _address)
}

// GiveBlockReward is a paid mutator transaction binding the contract method 0x746c8b65.
//
// Solidity: function giveBlockReward(uint256 stakerId, uint256 epoch) returns()
func (_IStakeManager *IStakeManagerTransactor) GiveBlockReward(opts *bind.TransactOpts, stakerId *big.Int, epoch *big.Int) (*types.Transaction, error) {
	return _IStakeManager.contract.Transact(opts, "giveBlockReward", stakerId, epoch)
}

// GiveBlockReward is a paid mutator transaction binding the contract method 0x746c8b65.
//
// Solidity: function giveBlockReward(uint256 stakerId, uint256 epoch) returns()
func (_IStakeManager *IStakeManagerSession) GiveBlockReward(stakerId *big.Int, epoch *big.Int) (*types.Transaction, error) {
	return _IStakeManager.Contract.GiveBlockReward(&_IStakeManager.TransactOpts, stakerId, epoch)
}

// GiveBlockReward is a paid mutator transaction binding the contract method 0x746c8b65.
//
// Solidity: function giveBlockReward(uint256 stakerId, uint256 epoch) returns()
func (_IStakeManager *IStakeManagerTransactorSession) GiveBlockReward(stakerId *big.Int, epoch *big.Int) (*types.Transaction, error) {
	return _IStakeManager.Contract.GiveBlockReward(&_IStakeManager.TransactOpts, stakerId, epoch)
}

// GivePenalties is a paid mutator transaction binding the contract method 0x54eae15e.
//
// Solidity: function givePenalties(uint256 stakerId, uint256 epoch) returns()
func (_IStakeManager *IStakeManagerTransactor) GivePenalties(opts *bind.TransactOpts, stakerId *big.Int, epoch *big.Int) (*types.Transaction, error) {
	return _IStakeManager.contract.Transact(opts, "givePenalties", stakerId, epoch)
}

// GivePenalties is a paid mutator transaction binding the contract method 0x54eae15e.
//
// Solidity: function givePenalties(uint256 stakerId, uint256 epoch) returns()
func (_IStakeManager *IStakeManagerSession) GivePenalties(stakerId *big.Int, epoch *big.Int) (*types.Transaction, error) {
	return _IStakeManager.Contract.GivePenalties(&_IStakeManager.TransactOpts, stakerId, epoch)
}

// GivePenalties is a paid mutator transaction binding the contract method 0x54eae15e.
//
// Solidity: function givePenalties(uint256 stakerId, uint256 epoch) returns()
func (_IStakeManager *IStakeManagerTransactorSession) GivePenalties(stakerId *big.Int, epoch *big.Int) (*types.Transaction, error) {
	return _IStakeManager.Contract.GivePenalties(&_IStakeManager.TransactOpts, stakerId, epoch)
}

// GiveRewards is a paid mutator transaction binding the contract method 0xfabb9890.
//
// Solidity: function giveRewards(uint256 stakerId, uint256 epoch) returns()
func (_IStakeManager *IStakeManagerTransactor) GiveRewards(opts *bind.TransactOpts, stakerId *big.Int, epoch *big.Int) (*types.Transaction, error) {
	return _IStakeManager.contract.Transact(opts, "giveRewards", stakerId, epoch)
}

// GiveRewards is a paid mutator transaction binding the contract method 0xfabb9890.
//
// Solidity: function giveRewards(uint256 stakerId, uint256 epoch) returns()
func (_IStakeManager *IStakeManagerSession) GiveRewards(stakerId *big.Int, epoch *big.Int) (*types.Transaction, error) {
	return _IStakeManager.Contract.GiveRewards(&_IStakeManager.TransactOpts, stakerId, epoch)
}

// GiveRewards is a paid mutator transaction binding the contract method 0xfabb9890.
//
// Solidity: function giveRewards(uint256 stakerId, uint256 epoch) returns()
func (_IStakeManager *IStakeManagerTransactorSession) GiveRewards(stakerId *big.Int, epoch *big.Int) (*types.Transaction, error) {
	return _IStakeManager.Contract.GiveRewards(&_IStakeManager.TransactOpts, stakerId, epoch)
}

// Init is a paid mutator transaction binding the contract method 0x06552ff3.
//
// Solidity: function init(address _schAddress, address _voteManagerAddress, address _blockManagerAddress, address _stateManagerAddress) returns()
func (_IStakeManager *IStakeManagerTransactor) Init(opts *bind.TransactOpts, _schAddress common.Address, _voteManagerAddress common.Address, _blockManagerAddress common.Address, _stateManagerAddress common.Address) (*types.Transaction, error) {
	return _IStakeManager.contract.Transact(opts, "init", _schAddress, _voteManagerAddress, _blockManagerAddress, _stateManagerAddress)
}

// Init is a paid mutator transaction binding the contract method 0x06552ff3.
//
// Solidity: function init(address _schAddress, address _voteManagerAddress, address _blockManagerAddress, address _stateManagerAddress) returns()
func (_IStakeManager *IStakeManagerSession) Init(_schAddress common.Address, _voteManagerAddress common.Address, _blockManagerAddress common.Address, _stateManagerAddress common.Address) (*types.Transaction, error) {
	return _IStakeManager.Contract.Init(&_IStakeManager.TransactOpts, _schAddress, _voteManagerAddress, _blockManagerAddress, _stateManagerAddress)
}

// Init is a paid mutator transaction binding the contract method 0x06552ff3.
//
// Solidity: function init(address _schAddress, address _voteManagerAddress, address _blockManagerAddress, address _stateManagerAddress) returns()
func (_IStakeManager *IStakeManagerTransactorSession) Init(_schAddress common.Address, _voteManagerAddress common.Address, _blockManagerAddress common.Address, _stateManagerAddress common.Address) (*types.Transaction, error) {
	return _IStakeManager.Contract.Init(&_IStakeManager.TransactOpts, _schAddress, _voteManagerAddress, _blockManagerAddress, _stateManagerAddress)
}

// SetStakerEpochLastRevealed is a paid mutator transaction binding the contract method 0x9864f70a.
//
// Solidity: function setStakerEpochLastRevealed(uint256 _id, uint256 _epochLastRevealed) returns()
func (_IStakeManager *IStakeManagerTransactor) SetStakerEpochLastRevealed(opts *bind.TransactOpts, _id *big.Int, _epochLastRevealed *big.Int) (*types.Transaction, error) {
	return _IStakeManager.contract.Transact(opts, "setStakerEpochLastRevealed", _id, _epochLastRevealed)
}

// SetStakerEpochLastRevealed is a paid mutator transaction binding the contract method 0x9864f70a.
//
// Solidity: function setStakerEpochLastRevealed(uint256 _id, uint256 _epochLastRevealed) returns()
func (_IStakeManager *IStakeManagerSession) SetStakerEpochLastRevealed(_id *big.Int, _epochLastRevealed *big.Int) (*types.Transaction, error) {
	return _IStakeManager.Contract.SetStakerEpochLastRevealed(&_IStakeManager.TransactOpts, _id, _epochLastRevealed)
}

// SetStakerEpochLastRevealed is a paid mutator transaction binding the contract method 0x9864f70a.
//
// Solidity: function setStakerEpochLastRevealed(uint256 _id, uint256 _epochLastRevealed) returns()
func (_IStakeManager *IStakeManagerTransactorSession) SetStakerEpochLastRevealed(_id *big.Int, _epochLastRevealed *big.Int) (*types.Transaction, error) {
	return _IStakeManager.Contract.SetStakerEpochLastRevealed(&_IStakeManager.TransactOpts, _id, _epochLastRevealed)
}

// Slash is a paid mutator transaction binding the contract method 0x0f91ce19.
//
// Solidity: function slash(uint256 id, address bountyHunter, uint256 epoch) returns()
func (_IStakeManager *IStakeManagerTransactor) Slash(opts *bind.TransactOpts, id *big.Int, bountyHunter common.Address, epoch *big.Int) (*types.Transaction, error) {
	return _IStakeManager.contract.Transact(opts, "slash", id, bountyHunter, epoch)
}

// Slash is a paid mutator transaction binding the contract method 0x0f91ce19.
//
// Solidity: function slash(uint256 id, address bountyHunter, uint256 epoch) returns()
func (_IStakeManager *IStakeManagerSession) Slash(id *big.Int, bountyHunter common.Address, epoch *big.Int) (*types.Transaction, error) {
	return _IStakeManager.Contract.Slash(&_IStakeManager.TransactOpts, id, bountyHunter, epoch)
}

// Slash is a paid mutator transaction binding the contract method 0x0f91ce19.
//
// Solidity: function slash(uint256 id, address bountyHunter, uint256 epoch) returns()
func (_IStakeManager *IStakeManagerTransactorSession) Slash(id *big.Int, bountyHunter common.Address, epoch *big.Int) (*types.Transaction, error) {
	return _IStakeManager.Contract.Slash(&_IStakeManager.TransactOpts, id, bountyHunter, epoch)
}

// Stake is a paid mutator transaction binding the contract method 0x7b0472f0.
//
// Solidity: function stake(uint256 epoch, uint256 amount) returns()
func (_IStakeManager *IStakeManagerTransactor) Stake(opts *bind.TransactOpts, epoch *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _IStakeManager.contract.Transact(opts, "stake", epoch, amount)
}

// Stake is a paid mutator transaction binding the contract method 0x7b0472f0.
//
// Solidity: function stake(uint256 epoch, uint256 amount) returns()
func (_IStakeManager *IStakeManagerSession) Stake(epoch *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _IStakeManager.Contract.Stake(&_IStakeManager.TransactOpts, epoch, amount)
}

// Stake is a paid mutator transaction binding the contract method 0x7b0472f0.
//
// Solidity: function stake(uint256 epoch, uint256 amount) returns()
func (_IStakeManager *IStakeManagerTransactorSession) Stake(epoch *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _IStakeManager.Contract.Stake(&_IStakeManager.TransactOpts, epoch, amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 epoch) returns()
func (_IStakeManager *IStakeManagerTransactor) Unstake(opts *bind.TransactOpts, epoch *big.Int) (*types.Transaction, error) {
	return _IStakeManager.contract.Transact(opts, "unstake", epoch)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 epoch) returns()
func (_IStakeManager *IStakeManagerSession) Unstake(epoch *big.Int) (*types.Transaction, error) {
	return _IStakeManager.Contract.Unstake(&_IStakeManager.TransactOpts, epoch)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 epoch) returns()
func (_IStakeManager *IStakeManagerTransactorSession) Unstake(epoch *big.Int) (*types.Transaction, error) {
	return _IStakeManager.Contract.Unstake(&_IStakeManager.TransactOpts, epoch)
}

// UpdateCommitmentEpoch is a paid mutator transaction binding the contract method 0x188dc83b.
//
// Solidity: function updateCommitmentEpoch(uint256 stakerId) returns()
func (_IStakeManager *IStakeManagerTransactor) UpdateCommitmentEpoch(opts *bind.TransactOpts, stakerId *big.Int) (*types.Transaction, error) {
	return _IStakeManager.contract.Transact(opts, "updateCommitmentEpoch", stakerId)
}

// UpdateCommitmentEpoch is a paid mutator transaction binding the contract method 0x188dc83b.
//
// Solidity: function updateCommitmentEpoch(uint256 stakerId) returns()
func (_IStakeManager *IStakeManagerSession) UpdateCommitmentEpoch(stakerId *big.Int) (*types.Transaction, error) {
	return _IStakeManager.Contract.UpdateCommitmentEpoch(&_IStakeManager.TransactOpts, stakerId)
}

// UpdateCommitmentEpoch is a paid mutator transaction binding the contract method 0x188dc83b.
//
// Solidity: function updateCommitmentEpoch(uint256 stakerId) returns()
func (_IStakeManager *IStakeManagerTransactorSession) UpdateCommitmentEpoch(stakerId *big.Int) (*types.Transaction, error) {
	return _IStakeManager.Contract.UpdateCommitmentEpoch(&_IStakeManager.TransactOpts, stakerId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 epoch) returns()
func (_IStakeManager *IStakeManagerTransactor) Withdraw(opts *bind.TransactOpts, epoch *big.Int) (*types.Transaction, error) {
	return _IStakeManager.contract.Transact(opts, "withdraw", epoch)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 epoch) returns()
func (_IStakeManager *IStakeManagerSession) Withdraw(epoch *big.Int) (*types.Transaction, error) {
	return _IStakeManager.Contract.Withdraw(&_IStakeManager.TransactOpts, epoch)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 epoch) returns()
func (_IStakeManager *IStakeManagerTransactorSession) Withdraw(epoch *big.Int) (*types.Transaction, error) {
	return _IStakeManager.Contract.Withdraw(&_IStakeManager.TransactOpts, epoch)
}
