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

// StructsBlock is an auto generated low-level Go binding around an user-defined struct.
type StructsBlock struct {
	ProposerId    *big.Int
	JobIds        []*big.Int
	Medians       []*big.Int
	LowerCutoffs  []*big.Int
	HigherCutoffs []*big.Int
	Iteration     *big.Int
	BiggestStake  *big.Int
	Valid         bool
}

// IBlockManagerABI is the input ABI used to generate the binding from.
const IBlockManagerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"medians\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"iteration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"biggestStakerId\",\"type\":\"uint256\"}],\"name\":\"Proposed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"confirmBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getBlock\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"proposerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"jobIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"medians\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"lowerCutoffs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"higherCutoffs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"iteration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"biggestStake\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"internalType\":\"structStructs.Block\",\"name\":\"_block\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getBlockMedians\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_blockMedians\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getHigherCutoffs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_higherCutoffs\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getLowerCutoffs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_lowerCutoffs\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getNumProposedBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposedBlock\",\"type\":\"uint256\"}],\"name\":\"getProposedBlockMedians\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_blockMedians\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"sorted\",\"type\":\"uint256[]\"}],\"name\":\"giveSorted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakeManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_voteManagerAddress\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"iteration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"biggestStakerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakerId\",\"type\":\"uint256\"}],\"name\":\"isElectedProposer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"jobIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"medians\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"lowerCutoffs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"higherCutoffs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"iteration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"biggestStakerId\",\"type\":\"uint256\"}],\"name\":\"propose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"resetDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IBlockManager is an auto generated Go binding around an Ethereum contract.
type IBlockManager struct {
	IBlockManagerCaller     // Read-only binding to the contract
	IBlockManagerTransactor // Write-only binding to the contract
	IBlockManagerFilterer   // Log filterer for contract events
}

// IBlockManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBlockManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBlockManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBlockManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBlockManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBlockManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBlockManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBlockManagerSession struct {
	Contract     *IBlockManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBlockManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBlockManagerCallerSession struct {
	Contract *IBlockManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IBlockManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBlockManagerTransactorSession struct {
	Contract     *IBlockManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IBlockManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBlockManagerRaw struct {
	Contract *IBlockManager // Generic contract binding to access the raw methods on
}

// IBlockManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBlockManagerCallerRaw struct {
	Contract *IBlockManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IBlockManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBlockManagerTransactorRaw struct {
	Contract *IBlockManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBlockManager creates a new instance of IBlockManager, bound to a specific deployed contract.
func NewIBlockManager(address common.Address, backend bind.ContractBackend) (*IBlockManager, error) {
	contract, err := bindIBlockManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBlockManager{IBlockManagerCaller: IBlockManagerCaller{contract: contract}, IBlockManagerTransactor: IBlockManagerTransactor{contract: contract}, IBlockManagerFilterer: IBlockManagerFilterer{contract: contract}}, nil
}

// NewIBlockManagerCaller creates a new read-only instance of IBlockManager, bound to a specific deployed contract.
func NewIBlockManagerCaller(address common.Address, caller bind.ContractCaller) (*IBlockManagerCaller, error) {
	contract, err := bindIBlockManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBlockManagerCaller{contract: contract}, nil
}

// NewIBlockManagerTransactor creates a new write-only instance of IBlockManager, bound to a specific deployed contract.
func NewIBlockManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IBlockManagerTransactor, error) {
	contract, err := bindIBlockManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBlockManagerTransactor{contract: contract}, nil
}

// NewIBlockManagerFilterer creates a new log filterer instance of IBlockManager, bound to a specific deployed contract.
func NewIBlockManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IBlockManagerFilterer, error) {
	contract, err := bindIBlockManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBlockManagerFilterer{contract: contract}, nil
}

// bindIBlockManager binds a generic wrapper to an already deployed contract.
func bindIBlockManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBlockManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBlockManager *IBlockManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBlockManager.Contract.IBlockManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBlockManager *IBlockManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBlockManager.Contract.IBlockManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBlockManager *IBlockManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBlockManager.Contract.IBlockManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBlockManager *IBlockManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBlockManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBlockManager *IBlockManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBlockManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBlockManager *IBlockManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBlockManager.Contract.contract.Transact(opts, method, params...)
}

// GetBlock is a free data retrieval call binding the contract method 0x04c07569.
//
// Solidity: function getBlock(uint256 epoch) view returns((uint256,uint256[],uint256[],uint256[],uint256[],uint256,uint256,bool) _block)
func (_IBlockManager *IBlockManagerCaller) GetBlock(opts *bind.CallOpts, epoch *big.Int) (StructsBlock, error) {
	var out []interface{}
	err := _IBlockManager.contract.Call(opts, &out, "getBlock", epoch)

	if err != nil {
		return *new(StructsBlock), err
	}

	out0 := *abi.ConvertType(out[0], new(StructsBlock)).(*StructsBlock)

	return out0, err

}

// GetBlock is a free data retrieval call binding the contract method 0x04c07569.
//
// Solidity: function getBlock(uint256 epoch) view returns((uint256,uint256[],uint256[],uint256[],uint256[],uint256,uint256,bool) _block)
func (_IBlockManager *IBlockManagerSession) GetBlock(epoch *big.Int) (StructsBlock, error) {
	return _IBlockManager.Contract.GetBlock(&_IBlockManager.CallOpts, epoch)
}

// GetBlock is a free data retrieval call binding the contract method 0x04c07569.
//
// Solidity: function getBlock(uint256 epoch) view returns((uint256,uint256[],uint256[],uint256[],uint256[],uint256,uint256,bool) _block)
func (_IBlockManager *IBlockManagerCallerSession) GetBlock(epoch *big.Int) (StructsBlock, error) {
	return _IBlockManager.Contract.GetBlock(&_IBlockManager.CallOpts, epoch)
}

// GetBlockMedians is a free data retrieval call binding the contract method 0x378ab9a9.
//
// Solidity: function getBlockMedians(uint256 epoch) view returns(uint256[] _blockMedians)
func (_IBlockManager *IBlockManagerCaller) GetBlockMedians(opts *bind.CallOpts, epoch *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _IBlockManager.contract.Call(opts, &out, "getBlockMedians", epoch)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetBlockMedians is a free data retrieval call binding the contract method 0x378ab9a9.
//
// Solidity: function getBlockMedians(uint256 epoch) view returns(uint256[] _blockMedians)
func (_IBlockManager *IBlockManagerSession) GetBlockMedians(epoch *big.Int) ([]*big.Int, error) {
	return _IBlockManager.Contract.GetBlockMedians(&_IBlockManager.CallOpts, epoch)
}

// GetBlockMedians is a free data retrieval call binding the contract method 0x378ab9a9.
//
// Solidity: function getBlockMedians(uint256 epoch) view returns(uint256[] _blockMedians)
func (_IBlockManager *IBlockManagerCallerSession) GetBlockMedians(epoch *big.Int) ([]*big.Int, error) {
	return _IBlockManager.Contract.GetBlockMedians(&_IBlockManager.CallOpts, epoch)
}

// GetHigherCutoffs is a free data retrieval call binding the contract method 0xfae4425d.
//
// Solidity: function getHigherCutoffs(uint256 epoch) view returns(uint256[] _higherCutoffs)
func (_IBlockManager *IBlockManagerCaller) GetHigherCutoffs(opts *bind.CallOpts, epoch *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _IBlockManager.contract.Call(opts, &out, "getHigherCutoffs", epoch)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetHigherCutoffs is a free data retrieval call binding the contract method 0xfae4425d.
//
// Solidity: function getHigherCutoffs(uint256 epoch) view returns(uint256[] _higherCutoffs)
func (_IBlockManager *IBlockManagerSession) GetHigherCutoffs(epoch *big.Int) ([]*big.Int, error) {
	return _IBlockManager.Contract.GetHigherCutoffs(&_IBlockManager.CallOpts, epoch)
}

// GetHigherCutoffs is a free data retrieval call binding the contract method 0xfae4425d.
//
// Solidity: function getHigherCutoffs(uint256 epoch) view returns(uint256[] _higherCutoffs)
func (_IBlockManager *IBlockManagerCallerSession) GetHigherCutoffs(epoch *big.Int) ([]*big.Int, error) {
	return _IBlockManager.Contract.GetHigherCutoffs(&_IBlockManager.CallOpts, epoch)
}

// GetLowerCutoffs is a free data retrieval call binding the contract method 0xd2a4669a.
//
// Solidity: function getLowerCutoffs(uint256 epoch) view returns(uint256[] _lowerCutoffs)
func (_IBlockManager *IBlockManagerCaller) GetLowerCutoffs(opts *bind.CallOpts, epoch *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _IBlockManager.contract.Call(opts, &out, "getLowerCutoffs", epoch)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetLowerCutoffs is a free data retrieval call binding the contract method 0xd2a4669a.
//
// Solidity: function getLowerCutoffs(uint256 epoch) view returns(uint256[] _lowerCutoffs)
func (_IBlockManager *IBlockManagerSession) GetLowerCutoffs(epoch *big.Int) ([]*big.Int, error) {
	return _IBlockManager.Contract.GetLowerCutoffs(&_IBlockManager.CallOpts, epoch)
}

// GetLowerCutoffs is a free data retrieval call binding the contract method 0xd2a4669a.
//
// Solidity: function getLowerCutoffs(uint256 epoch) view returns(uint256[] _lowerCutoffs)
func (_IBlockManager *IBlockManagerCallerSession) GetLowerCutoffs(epoch *big.Int) ([]*big.Int, error) {
	return _IBlockManager.Contract.GetLowerCutoffs(&_IBlockManager.CallOpts, epoch)
}

// GetNumProposedBlocks is a free data retrieval call binding the contract method 0xe38c7c42.
//
// Solidity: function getNumProposedBlocks(uint256 epoch) view returns(uint256)
func (_IBlockManager *IBlockManagerCaller) GetNumProposedBlocks(opts *bind.CallOpts, epoch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IBlockManager.contract.Call(opts, &out, "getNumProposedBlocks", epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumProposedBlocks is a free data retrieval call binding the contract method 0xe38c7c42.
//
// Solidity: function getNumProposedBlocks(uint256 epoch) view returns(uint256)
func (_IBlockManager *IBlockManagerSession) GetNumProposedBlocks(epoch *big.Int) (*big.Int, error) {
	return _IBlockManager.Contract.GetNumProposedBlocks(&_IBlockManager.CallOpts, epoch)
}

// GetNumProposedBlocks is a free data retrieval call binding the contract method 0xe38c7c42.
//
// Solidity: function getNumProposedBlocks(uint256 epoch) view returns(uint256)
func (_IBlockManager *IBlockManagerCallerSession) GetNumProposedBlocks(epoch *big.Int) (*big.Int, error) {
	return _IBlockManager.Contract.GetNumProposedBlocks(&_IBlockManager.CallOpts, epoch)
}

// GetProposedBlockMedians is a free data retrieval call binding the contract method 0xd1a4a43d.
//
// Solidity: function getProposedBlockMedians(uint256 epoch, uint256 proposedBlock) view returns(uint256[] _blockMedians)
func (_IBlockManager *IBlockManagerCaller) GetProposedBlockMedians(opts *bind.CallOpts, epoch *big.Int, proposedBlock *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _IBlockManager.contract.Call(opts, &out, "getProposedBlockMedians", epoch, proposedBlock)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetProposedBlockMedians is a free data retrieval call binding the contract method 0xd1a4a43d.
//
// Solidity: function getProposedBlockMedians(uint256 epoch, uint256 proposedBlock) view returns(uint256[] _blockMedians)
func (_IBlockManager *IBlockManagerSession) GetProposedBlockMedians(epoch *big.Int, proposedBlock *big.Int) ([]*big.Int, error) {
	return _IBlockManager.Contract.GetProposedBlockMedians(&_IBlockManager.CallOpts, epoch, proposedBlock)
}

// GetProposedBlockMedians is a free data retrieval call binding the contract method 0xd1a4a43d.
//
// Solidity: function getProposedBlockMedians(uint256 epoch, uint256 proposedBlock) view returns(uint256[] _blockMedians)
func (_IBlockManager *IBlockManagerCallerSession) GetProposedBlockMedians(epoch *big.Int, proposedBlock *big.Int) ([]*big.Int, error) {
	return _IBlockManager.Contract.GetProposedBlockMedians(&_IBlockManager.CallOpts, epoch, proposedBlock)
}

// ConfirmBlock is a paid mutator transaction binding the contract method 0x9b87f644.
//
// Solidity: function confirmBlock() returns()
func (_IBlockManager *IBlockManagerTransactor) ConfirmBlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBlockManager.contract.Transact(opts, "confirmBlock")
}

// ConfirmBlock is a paid mutator transaction binding the contract method 0x9b87f644.
//
// Solidity: function confirmBlock() returns()
func (_IBlockManager *IBlockManagerSession) ConfirmBlock() (*types.Transaction, error) {
	return _IBlockManager.Contract.ConfirmBlock(&_IBlockManager.TransactOpts)
}

// ConfirmBlock is a paid mutator transaction binding the contract method 0x9b87f644.
//
// Solidity: function confirmBlock() returns()
func (_IBlockManager *IBlockManagerTransactorSession) ConfirmBlock() (*types.Transaction, error) {
	return _IBlockManager.Contract.ConfirmBlock(&_IBlockManager.TransactOpts)
}

// GiveSorted is a paid mutator transaction binding the contract method 0x4e6753b7.
//
// Solidity: function giveSorted(uint256 epoch, uint256 assetId, uint256[] sorted) returns()
func (_IBlockManager *IBlockManagerTransactor) GiveSorted(opts *bind.TransactOpts, epoch *big.Int, assetId *big.Int, sorted []*big.Int) (*types.Transaction, error) {
	return _IBlockManager.contract.Transact(opts, "giveSorted", epoch, assetId, sorted)
}

// GiveSorted is a paid mutator transaction binding the contract method 0x4e6753b7.
//
// Solidity: function giveSorted(uint256 epoch, uint256 assetId, uint256[] sorted) returns()
func (_IBlockManager *IBlockManagerSession) GiveSorted(epoch *big.Int, assetId *big.Int, sorted []*big.Int) (*types.Transaction, error) {
	return _IBlockManager.Contract.GiveSorted(&_IBlockManager.TransactOpts, epoch, assetId, sorted)
}

// GiveSorted is a paid mutator transaction binding the contract method 0x4e6753b7.
//
// Solidity: function giveSorted(uint256 epoch, uint256 assetId, uint256[] sorted) returns()
func (_IBlockManager *IBlockManagerTransactorSession) GiveSorted(epoch *big.Int, assetId *big.Int, sorted []*big.Int) (*types.Transaction, error) {
	return _IBlockManager.Contract.GiveSorted(&_IBlockManager.TransactOpts, epoch, assetId, sorted)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address _stakeManagerAddress, address _voteManagerAddress) returns()
func (_IBlockManager *IBlockManagerTransactor) Init(opts *bind.TransactOpts, _stakeManagerAddress common.Address, _voteManagerAddress common.Address) (*types.Transaction, error) {
	return _IBlockManager.contract.Transact(opts, "init", _stakeManagerAddress, _voteManagerAddress)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address _stakeManagerAddress, address _voteManagerAddress) returns()
func (_IBlockManager *IBlockManagerSession) Init(_stakeManagerAddress common.Address, _voteManagerAddress common.Address) (*types.Transaction, error) {
	return _IBlockManager.Contract.Init(&_IBlockManager.TransactOpts, _stakeManagerAddress, _voteManagerAddress)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address _stakeManagerAddress, address _voteManagerAddress) returns()
func (_IBlockManager *IBlockManagerTransactorSession) Init(_stakeManagerAddress common.Address, _voteManagerAddress common.Address) (*types.Transaction, error) {
	return _IBlockManager.Contract.Init(&_IBlockManager.TransactOpts, _stakeManagerAddress, _voteManagerAddress)
}

// IsElectedProposer is a paid mutator transaction binding the contract method 0x1d69ff9b.
//
// Solidity: function isElectedProposer(uint256 iteration, uint256 biggestStakerId, uint256 stakerId) returns()
func (_IBlockManager *IBlockManagerTransactor) IsElectedProposer(opts *bind.TransactOpts, iteration *big.Int, biggestStakerId *big.Int, stakerId *big.Int) (*types.Transaction, error) {
	return _IBlockManager.contract.Transact(opts, "isElectedProposer", iteration, biggestStakerId, stakerId)
}

// IsElectedProposer is a paid mutator transaction binding the contract method 0x1d69ff9b.
//
// Solidity: function isElectedProposer(uint256 iteration, uint256 biggestStakerId, uint256 stakerId) returns()
func (_IBlockManager *IBlockManagerSession) IsElectedProposer(iteration *big.Int, biggestStakerId *big.Int, stakerId *big.Int) (*types.Transaction, error) {
	return _IBlockManager.Contract.IsElectedProposer(&_IBlockManager.TransactOpts, iteration, biggestStakerId, stakerId)
}

// IsElectedProposer is a paid mutator transaction binding the contract method 0x1d69ff9b.
//
// Solidity: function isElectedProposer(uint256 iteration, uint256 biggestStakerId, uint256 stakerId) returns()
func (_IBlockManager *IBlockManagerTransactorSession) IsElectedProposer(iteration *big.Int, biggestStakerId *big.Int, stakerId *big.Int) (*types.Transaction, error) {
	return _IBlockManager.Contract.IsElectedProposer(&_IBlockManager.TransactOpts, iteration, biggestStakerId, stakerId)
}

// Propose is a paid mutator transaction binding the contract method 0x17d99c04.
//
// Solidity: function propose(uint256 epoch, uint256[] jobIds, uint256[] medians, uint256[] lowerCutoffs, uint256[] higherCutoffs, uint256 iteration, uint256 biggestStakerId) returns()
func (_IBlockManager *IBlockManagerTransactor) Propose(opts *bind.TransactOpts, epoch *big.Int, jobIds []*big.Int, medians []*big.Int, lowerCutoffs []*big.Int, higherCutoffs []*big.Int, iteration *big.Int, biggestStakerId *big.Int) (*types.Transaction, error) {
	return _IBlockManager.contract.Transact(opts, "propose", epoch, jobIds, medians, lowerCutoffs, higherCutoffs, iteration, biggestStakerId)
}

// Propose is a paid mutator transaction binding the contract method 0x17d99c04.
//
// Solidity: function propose(uint256 epoch, uint256[] jobIds, uint256[] medians, uint256[] lowerCutoffs, uint256[] higherCutoffs, uint256 iteration, uint256 biggestStakerId) returns()
func (_IBlockManager *IBlockManagerSession) Propose(epoch *big.Int, jobIds []*big.Int, medians []*big.Int, lowerCutoffs []*big.Int, higherCutoffs []*big.Int, iteration *big.Int, biggestStakerId *big.Int) (*types.Transaction, error) {
	return _IBlockManager.Contract.Propose(&_IBlockManager.TransactOpts, epoch, jobIds, medians, lowerCutoffs, higherCutoffs, iteration, biggestStakerId)
}

// Propose is a paid mutator transaction binding the contract method 0x17d99c04.
//
// Solidity: function propose(uint256 epoch, uint256[] jobIds, uint256[] medians, uint256[] lowerCutoffs, uint256[] higherCutoffs, uint256 iteration, uint256 biggestStakerId) returns()
func (_IBlockManager *IBlockManagerTransactorSession) Propose(epoch *big.Int, jobIds []*big.Int, medians []*big.Int, lowerCutoffs []*big.Int, higherCutoffs []*big.Int, iteration *big.Int, biggestStakerId *big.Int) (*types.Transaction, error) {
	return _IBlockManager.Contract.Propose(&_IBlockManager.TransactOpts, epoch, jobIds, medians, lowerCutoffs, higherCutoffs, iteration, biggestStakerId)
}

// ResetDispute is a paid mutator transaction binding the contract method 0x5ce8772f.
//
// Solidity: function resetDispute(uint256 epoch) returns()
func (_IBlockManager *IBlockManagerTransactor) ResetDispute(opts *bind.TransactOpts, epoch *big.Int) (*types.Transaction, error) {
	return _IBlockManager.contract.Transact(opts, "resetDispute", epoch)
}

// ResetDispute is a paid mutator transaction binding the contract method 0x5ce8772f.
//
// Solidity: function resetDispute(uint256 epoch) returns()
func (_IBlockManager *IBlockManagerSession) ResetDispute(epoch *big.Int) (*types.Transaction, error) {
	return _IBlockManager.Contract.ResetDispute(&_IBlockManager.TransactOpts, epoch)
}

// ResetDispute is a paid mutator transaction binding the contract method 0x5ce8772f.
//
// Solidity: function resetDispute(uint256 epoch) returns()
func (_IBlockManager *IBlockManagerTransactorSession) ResetDispute(epoch *big.Int) (*types.Transaction, error) {
	return _IBlockManager.Contract.ResetDispute(&_IBlockManager.TransactOpts, epoch)
}

// IBlockManagerProposedIterator is returned from FilterProposed and is used to iterate over the raw logs and unpacked data for Proposed events raised by the IBlockManager contract.
type IBlockManagerProposedIterator struct {
	Event *IBlockManagerProposed // Event containing the contract specifics and raw log

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
func (it *IBlockManagerProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBlockManagerProposed)
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
		it.Event = new(IBlockManagerProposed)
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
func (it *IBlockManagerProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBlockManagerProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBlockManagerProposed represents a Proposed event raised by the IBlockManager contract.
type IBlockManagerProposed struct {
	Epoch           *big.Int
	StakerId        *big.Int
	Medians         []*big.Int
	Iteration       *big.Int
	BiggestStakerId *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterProposed is a free log retrieval operation binding the contract event 0xfb750a2a8aed92f932bb76b78077f486a1d583d286b29c3afb9ae793a69c63aa.
//
// Solidity: event Proposed(uint256 epoch, uint256 stakerId, uint256[] medians, uint256 iteration, uint256 biggestStakerId)
func (_IBlockManager *IBlockManagerFilterer) FilterProposed(opts *bind.FilterOpts) (*IBlockManagerProposedIterator, error) {

	logs, sub, err := _IBlockManager.contract.FilterLogs(opts, "Proposed")
	if err != nil {
		return nil, err
	}
	return &IBlockManagerProposedIterator{contract: _IBlockManager.contract, event: "Proposed", logs: logs, sub: sub}, nil
}

// WatchProposed is a free log subscription operation binding the contract event 0xfb750a2a8aed92f932bb76b78077f486a1d583d286b29c3afb9ae793a69c63aa.
//
// Solidity: event Proposed(uint256 epoch, uint256 stakerId, uint256[] medians, uint256 iteration, uint256 biggestStakerId)
func (_IBlockManager *IBlockManagerFilterer) WatchProposed(opts *bind.WatchOpts, sink chan<- *IBlockManagerProposed) (event.Subscription, error) {

	logs, sub, err := _IBlockManager.contract.WatchLogs(opts, "Proposed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBlockManagerProposed)
				if err := _IBlockManager.contract.UnpackLog(event, "Proposed", log); err != nil {
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

// ParseProposed is a log parse operation binding the contract event 0xfb750a2a8aed92f932bb76b78077f486a1d583d286b29c3afb9ae793a69c63aa.
//
// Solidity: event Proposed(uint256 epoch, uint256 stakerId, uint256[] medians, uint256 iteration, uint256 biggestStakerId)
func (_IBlockManager *IBlockManagerFilterer) ParseProposed(log types.Log) (*IBlockManagerProposed, error) {
	event := new(IBlockManagerProposed)
	if err := _IBlockManager.contract.UnpackLog(event, "Proposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
