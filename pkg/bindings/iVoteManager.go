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

// StructsIVote is an auto generated low-level Go binding around an user-defined struct.
type StructsIVote struct {
	Value  *big.Int
	Weight *big.Int
}

// IVoteManagerABI is the input ABI used to generate the binding from.
const IVoteManagerABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakerId\",\"type\":\"uint256\"}],\"name\":\"getCommitment\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"getTotalStakeRevealed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voteValue\",\"type\":\"uint256\"}],\"name\":\"getTotalStakeRevealed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"getVote\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"internalType\":\"structStructs.Vote\",\"name\":\"vote\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voteValue\",\"type\":\"uint256\"}],\"name\":\"getVoteWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakeManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_blockManagerAddress\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"proofs\",\"type\":\"bytes32[][]\"},{\"internalType\":\"bytes32\",\"name\":\"secret\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"reveal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IVoteManager is an auto generated Go binding around an Ethereum contract.
type IVoteManager struct {
	IVoteManagerCaller     // Read-only binding to the contract
	IVoteManagerTransactor // Write-only binding to the contract
	IVoteManagerFilterer   // Log filterer for contract events
}

// IVoteManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IVoteManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVoteManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IVoteManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVoteManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IVoteManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVoteManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IVoteManagerSession struct {
	Contract     *IVoteManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IVoteManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IVoteManagerCallerSession struct {
	Contract *IVoteManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IVoteManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IVoteManagerTransactorSession struct {
	Contract     *IVoteManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IVoteManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IVoteManagerRaw struct {
	Contract *IVoteManager // Generic contract binding to access the raw methods on
}

// IVoteManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IVoteManagerCallerRaw struct {
	Contract *IVoteManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IVoteManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IVoteManagerTransactorRaw struct {
	Contract *IVoteManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIVoteManager creates a new instance of IVoteManager, bound to a specific deployed contract.
func NewIVoteManager(address common.Address, backend bind.ContractBackend) (*IVoteManager, error) {
	contract, err := bindIVoteManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVoteManager{IVoteManagerCaller: IVoteManagerCaller{contract: contract}, IVoteManagerTransactor: IVoteManagerTransactor{contract: contract}, IVoteManagerFilterer: IVoteManagerFilterer{contract: contract}}, nil
}

// NewIVoteManagerCaller creates a new read-only instance of IVoteManager, bound to a specific deployed contract.
func NewIVoteManagerCaller(address common.Address, caller bind.ContractCaller) (*IVoteManagerCaller, error) {
	contract, err := bindIVoteManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVoteManagerCaller{contract: contract}, nil
}

// NewIVoteManagerTransactor creates a new write-only instance of IVoteManager, bound to a specific deployed contract.
func NewIVoteManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IVoteManagerTransactor, error) {
	contract, err := bindIVoteManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVoteManagerTransactor{contract: contract}, nil
}

// NewIVoteManagerFilterer creates a new log filterer instance of IVoteManager, bound to a specific deployed contract.
func NewIVoteManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IVoteManagerFilterer, error) {
	contract, err := bindIVoteManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVoteManagerFilterer{contract: contract}, nil
}

// bindIVoteManager binds a generic wrapper to an already deployed contract.
func bindIVoteManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IVoteManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVoteManager *IVoteManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVoteManager.Contract.IVoteManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVoteManager *IVoteManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVoteManager.Contract.IVoteManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVoteManager *IVoteManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVoteManager.Contract.IVoteManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVoteManager *IVoteManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVoteManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVoteManager *IVoteManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVoteManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVoteManager *IVoteManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVoteManager.Contract.contract.Transact(opts, method, params...)
}

// GetCommitment is a free data retrieval call binding the contract method 0x7164f0d6.
//
// Solidity: function getCommitment(uint256 epoch, uint256 stakerId) view returns(bytes32)
func (_IVoteManager *IVoteManagerCaller) GetCommitment(opts *bind.CallOpts, epoch *big.Int, stakerId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IVoteManager.contract.Call(opts, &out, "getCommitment", epoch, stakerId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCommitment is a free data retrieval call binding the contract method 0x7164f0d6.
//
// Solidity: function getCommitment(uint256 epoch, uint256 stakerId) view returns(bytes32)
func (_IVoteManager *IVoteManagerSession) GetCommitment(epoch *big.Int, stakerId *big.Int) ([32]byte, error) {
	return _IVoteManager.Contract.GetCommitment(&_IVoteManager.CallOpts, epoch, stakerId)
}

// GetCommitment is a free data retrieval call binding the contract method 0x7164f0d6.
//
// Solidity: function getCommitment(uint256 epoch, uint256 stakerId) view returns(bytes32)
func (_IVoteManager *IVoteManagerCallerSession) GetCommitment(epoch *big.Int, stakerId *big.Int) ([32]byte, error) {
	return _IVoteManager.Contract.GetCommitment(&_IVoteManager.CallOpts, epoch, stakerId)
}

// GetTotalStakeRevealed is a free data retrieval call binding the contract method 0xa6a145d9.
//
// Solidity: function getTotalStakeRevealed(uint256 epoch, uint256 assetId) view returns(uint256)
func (_IVoteManager *IVoteManagerCaller) GetTotalStakeRevealed(opts *bind.CallOpts, epoch *big.Int, assetId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IVoteManager.contract.Call(opts, &out, "getTotalStakeRevealed", epoch, assetId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalStakeRevealed is a free data retrieval call binding the contract method 0xa6a145d9.
//
// Solidity: function getTotalStakeRevealed(uint256 epoch, uint256 assetId) view returns(uint256)
func (_IVoteManager *IVoteManagerSession) GetTotalStakeRevealed(epoch *big.Int, assetId *big.Int) (*big.Int, error) {
	return _IVoteManager.Contract.GetTotalStakeRevealed(&_IVoteManager.CallOpts, epoch, assetId)
}

// GetTotalStakeRevealed is a free data retrieval call binding the contract method 0xa6a145d9.
//
// Solidity: function getTotalStakeRevealed(uint256 epoch, uint256 assetId) view returns(uint256)
func (_IVoteManager *IVoteManagerCallerSession) GetTotalStakeRevealed(epoch *big.Int, assetId *big.Int) (*big.Int, error) {
	return _IVoteManager.Contract.GetTotalStakeRevealed(&_IVoteManager.CallOpts, epoch, assetId)
}

// GetTotalStakeRevealed0 is a free data retrieval call binding the contract method 0xf322bd7d.
//
// Solidity: function getTotalStakeRevealed(uint256 epoch, uint256 assetId, uint256 voteValue) view returns(uint256)
func (_IVoteManager *IVoteManagerCaller) GetTotalStakeRevealed0(opts *bind.CallOpts, epoch *big.Int, assetId *big.Int, voteValue *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IVoteManager.contract.Call(opts, &out, "getTotalStakeRevealed0", epoch, assetId, voteValue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalStakeRevealed0 is a free data retrieval call binding the contract method 0xf322bd7d.
//
// Solidity: function getTotalStakeRevealed(uint256 epoch, uint256 assetId, uint256 voteValue) view returns(uint256)
func (_IVoteManager *IVoteManagerSession) GetTotalStakeRevealed0(epoch *big.Int, assetId *big.Int, voteValue *big.Int) (*big.Int, error) {
	return _IVoteManager.Contract.GetTotalStakeRevealed0(&_IVoteManager.CallOpts, epoch, assetId, voteValue)
}

// GetTotalStakeRevealed0 is a free data retrieval call binding the contract method 0xf322bd7d.
//
// Solidity: function getTotalStakeRevealed(uint256 epoch, uint256 assetId, uint256 voteValue) view returns(uint256)
func (_IVoteManager *IVoteManagerCallerSession) GetTotalStakeRevealed0(epoch *big.Int, assetId *big.Int, voteValue *big.Int) (*big.Int, error) {
	return _IVoteManager.Contract.GetTotalStakeRevealed0(&_IVoteManager.CallOpts, epoch, assetId, voteValue)
}

// GetVote is a free data retrieval call binding the contract method 0x8ce7ff4a.
//
// Solidity: function getVote(uint256 epoch, uint256 stakerId, uint256 assetId) view returns((uint256,uint256) vote)
func (_IVoteManager *IVoteManagerCaller) GetVote(opts *bind.CallOpts, epoch *big.Int, stakerId *big.Int, assetId *big.Int) (StructsIVote, error) {
	var out []interface{}
	err := _IVoteManager.contract.Call(opts, &out, "getVote", epoch, stakerId, assetId)

	if err != nil {
		return *new(StructsIVote), err
	}

	out0 := *abi.ConvertType(out[0], new(StructsIVote)).(*StructsIVote)

	return out0, err

}

// GetVote is a free data retrieval call binding the contract method 0x8ce7ff4a.
//
// Solidity: function getVote(uint256 epoch, uint256 stakerId, uint256 assetId) view returns((uint256,uint256) vote)
func (_IVoteManager *IVoteManagerSession) GetVote(epoch *big.Int, stakerId *big.Int, assetId *big.Int) (StructsIVote, error) {
	return _IVoteManager.Contract.GetVote(&_IVoteManager.CallOpts, epoch, stakerId, assetId)
}

// GetVote is a free data retrieval call binding the contract method 0x8ce7ff4a.
//
// Solidity: function getVote(uint256 epoch, uint256 stakerId, uint256 assetId) view returns((uint256,uint256) vote)
func (_IVoteManager *IVoteManagerCallerSession) GetVote(epoch *big.Int, stakerId *big.Int, assetId *big.Int) (StructsIVote, error) {
	return _IVoteManager.Contract.GetVote(&_IVoteManager.CallOpts, epoch, stakerId, assetId)
}

// GetVoteWeight is a free data retrieval call binding the contract method 0x9c66556f.
//
// Solidity: function getVoteWeight(uint256 epoch, uint256 assetId, uint256 voteValue) view returns(uint256)
func (_IVoteManager *IVoteManagerCaller) GetVoteWeight(opts *bind.CallOpts, epoch *big.Int, assetId *big.Int, voteValue *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IVoteManager.contract.Call(opts, &out, "getVoteWeight", epoch, assetId, voteValue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVoteWeight is a free data retrieval call binding the contract method 0x9c66556f.
//
// Solidity: function getVoteWeight(uint256 epoch, uint256 assetId, uint256 voteValue) view returns(uint256)
func (_IVoteManager *IVoteManagerSession) GetVoteWeight(epoch *big.Int, assetId *big.Int, voteValue *big.Int) (*big.Int, error) {
	return _IVoteManager.Contract.GetVoteWeight(&_IVoteManager.CallOpts, epoch, assetId, voteValue)
}

// GetVoteWeight is a free data retrieval call binding the contract method 0x9c66556f.
//
// Solidity: function getVoteWeight(uint256 epoch, uint256 assetId, uint256 voteValue) view returns(uint256)
func (_IVoteManager *IVoteManagerCallerSession) GetVoteWeight(epoch *big.Int, assetId *big.Int, voteValue *big.Int) (*big.Int, error) {
	return _IVoteManager.Contract.GetVoteWeight(&_IVoteManager.CallOpts, epoch, assetId, voteValue)
}

// Commit is a paid mutator transaction binding the contract method 0xf2f03877.
//
// Solidity: function commit(uint256 epoch, bytes32 commitment) returns()
func (_IVoteManager *IVoteManagerTransactor) Commit(opts *bind.TransactOpts, epoch *big.Int, commitment [32]byte) (*types.Transaction, error) {
	return _IVoteManager.contract.Transact(opts, "commit", epoch, commitment)
}

// Commit is a paid mutator transaction binding the contract method 0xf2f03877.
//
// Solidity: function commit(uint256 epoch, bytes32 commitment) returns()
func (_IVoteManager *IVoteManagerSession) Commit(epoch *big.Int, commitment [32]byte) (*types.Transaction, error) {
	return _IVoteManager.Contract.Commit(&_IVoteManager.TransactOpts, epoch, commitment)
}

// Commit is a paid mutator transaction binding the contract method 0xf2f03877.
//
// Solidity: function commit(uint256 epoch, bytes32 commitment) returns()
func (_IVoteManager *IVoteManagerTransactorSession) Commit(epoch *big.Int, commitment [32]byte) (*types.Transaction, error) {
	return _IVoteManager.Contract.Commit(&_IVoteManager.TransactOpts, epoch, commitment)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address _stakeManagerAddress, address _blockManagerAddress) returns()
func (_IVoteManager *IVoteManagerTransactor) Init(opts *bind.TransactOpts, _stakeManagerAddress common.Address, _blockManagerAddress common.Address) (*types.Transaction, error) {
	return _IVoteManager.contract.Transact(opts, "init", _stakeManagerAddress, _blockManagerAddress)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address _stakeManagerAddress, address _blockManagerAddress) returns()
func (_IVoteManager *IVoteManagerSession) Init(_stakeManagerAddress common.Address, _blockManagerAddress common.Address) (*types.Transaction, error) {
	return _IVoteManager.Contract.Init(&_IVoteManager.TransactOpts, _stakeManagerAddress, _blockManagerAddress)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address _stakeManagerAddress, address _blockManagerAddress) returns()
func (_IVoteManager *IVoteManagerTransactorSession) Init(_stakeManagerAddress common.Address, _blockManagerAddress common.Address) (*types.Transaction, error) {
	return _IVoteManager.Contract.Init(&_IVoteManager.TransactOpts, _stakeManagerAddress, _blockManagerAddress)
}

// Reveal is a paid mutator transaction binding the contract method 0x125f68f1.
//
// Solidity: function reveal(uint256 epoch, bytes32 root, uint256[] values, bytes32[][] proofs, bytes32 secret, address stakerAddress) returns()
func (_IVoteManager *IVoteManagerTransactor) Reveal(opts *bind.TransactOpts, epoch *big.Int, root [32]byte, values []*big.Int, proofs [][][32]byte, secret [32]byte, stakerAddress common.Address) (*types.Transaction, error) {
	return _IVoteManager.contract.Transact(opts, "reveal", epoch, root, values, proofs, secret, stakerAddress)
}

// Reveal is a paid mutator transaction binding the contract method 0x125f68f1.
//
// Solidity: function reveal(uint256 epoch, bytes32 root, uint256[] values, bytes32[][] proofs, bytes32 secret, address stakerAddress) returns()
func (_IVoteManager *IVoteManagerSession) Reveal(epoch *big.Int, root [32]byte, values []*big.Int, proofs [][][32]byte, secret [32]byte, stakerAddress common.Address) (*types.Transaction, error) {
	return _IVoteManager.Contract.Reveal(&_IVoteManager.TransactOpts, epoch, root, values, proofs, secret, stakerAddress)
}

// Reveal is a paid mutator transaction binding the contract method 0x125f68f1.
//
// Solidity: function reveal(uint256 epoch, bytes32 root, uint256[] values, bytes32[][] proofs, bytes32 secret, address stakerAddress) returns()
func (_IVoteManager *IVoteManagerTransactorSession) Reveal(epoch *big.Int, root [32]byte, values []*big.Int, proofs [][][32]byte, secret [32]byte, stakerAddress common.Address) (*types.Transaction, error) {
	return _IVoteManager.Contract.Reveal(&_IVoteManager.TransactOpts, epoch, root, values, proofs, secret, stakerAddress)
}
