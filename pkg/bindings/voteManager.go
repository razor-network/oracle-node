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

// StructsVote is an auto generated low-level Go binding around an user-defined struct.
type StructsVote struct {
	Value  *big.Int
	Weight *big.Int
}

// VoteManagerABI is the input ABI used to generate the binding from.
const VoteManagerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Committed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Revealed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockManager\",\"outputs\":[{\"internalType\":\"contractIBlockManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"commitments\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakerId\",\"type\":\"uint256\"}],\"name\":\"getCommitment\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"getTotalStakeRevealed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voteValue\",\"type\":\"uint256\"}],\"name\":\"getTotalStakeRevealed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"getVote\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"internalType\":\"structStructs.Vote\",\"name\":\"vote\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voteValue\",\"type\":\"uint256\"}],\"name\":\"getVoteWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakeManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"blockManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"parametersAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"parameters\",\"outputs\":[{\"internalType\":\"contractIParameters\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"proofs\",\"type\":\"bytes32[][]\"},{\"internalType\":\"bytes32\",\"name\":\"secret\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"name\":\"reveal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardManager\",\"outputs\":[{\"internalType\":\"contractIRewardManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeManager\",\"outputs\":[{\"internalType\":\"contractIStakeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"totalStakeRevealed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voteWeights\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"votes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VoteManager is an auto generated Go binding around an Ethereum contract.
type VoteManager struct {
	VoteManagerCaller     // Read-only binding to the contract
	VoteManagerTransactor // Write-only binding to the contract
	VoteManagerFilterer   // Log filterer for contract events
}

// VoteManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoteManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoteManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoteManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoteManagerSession struct {
	Contract     *VoteManager      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoteManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoteManagerCallerSession struct {
	Contract *VoteManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// VoteManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoteManagerTransactorSession struct {
	Contract     *VoteManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// VoteManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoteManagerRaw struct {
	Contract *VoteManager // Generic contract binding to access the raw methods on
}

// VoteManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoteManagerCallerRaw struct {
	Contract *VoteManagerCaller // Generic read-only contract binding to access the raw methods on
}

// VoteManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoteManagerTransactorRaw struct {
	Contract *VoteManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoteManager creates a new instance of VoteManager, bound to a specific deployed contract.
func NewVoteManager(address common.Address, backend bind.ContractBackend) (*VoteManager, error) {
	contract, err := bindVoteManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VoteManager{VoteManagerCaller: VoteManagerCaller{contract: contract}, VoteManagerTransactor: VoteManagerTransactor{contract: contract}, VoteManagerFilterer: VoteManagerFilterer{contract: contract}}, nil
}

// NewVoteManagerCaller creates a new read-only instance of VoteManager, bound to a specific deployed contract.
func NewVoteManagerCaller(address common.Address, caller bind.ContractCaller) (*VoteManagerCaller, error) {
	contract, err := bindVoteManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoteManagerCaller{contract: contract}, nil
}

// NewVoteManagerTransactor creates a new write-only instance of VoteManager, bound to a specific deployed contract.
func NewVoteManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*VoteManagerTransactor, error) {
	contract, err := bindVoteManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoteManagerTransactor{contract: contract}, nil
}

// NewVoteManagerFilterer creates a new log filterer instance of VoteManager, bound to a specific deployed contract.
func NewVoteManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*VoteManagerFilterer, error) {
	contract, err := bindVoteManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoteManagerFilterer{contract: contract}, nil
}

// bindVoteManager binds a generic wrapper to an already deployed contract.
func bindVoteManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VoteManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VoteManager *VoteManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VoteManager.Contract.VoteManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VoteManager *VoteManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoteManager.Contract.VoteManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VoteManager *VoteManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VoteManager.Contract.VoteManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VoteManager *VoteManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VoteManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VoteManager *VoteManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoteManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VoteManager *VoteManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VoteManager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VoteManager *VoteManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VoteManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VoteManager *VoteManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _VoteManager.Contract.DEFAULTADMINROLE(&_VoteManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VoteManager *VoteManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _VoteManager.Contract.DEFAULTADMINROLE(&_VoteManager.CallOpts)
}

// BlockManager is a free data retrieval call binding the contract method 0xd9169b32.
//
// Solidity: function blockManager() view returns(address)
func (_VoteManager *VoteManagerCaller) BlockManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VoteManager.contract.Call(opts, &out, "blockManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlockManager is a free data retrieval call binding the contract method 0xd9169b32.
//
// Solidity: function blockManager() view returns(address)
func (_VoteManager *VoteManagerSession) BlockManager() (common.Address, error) {
	return _VoteManager.Contract.BlockManager(&_VoteManager.CallOpts)
}

// BlockManager is a free data retrieval call binding the contract method 0xd9169b32.
//
// Solidity: function blockManager() view returns(address)
func (_VoteManager *VoteManagerCallerSession) BlockManager() (common.Address, error) {
	return _VoteManager.Contract.BlockManager(&_VoteManager.CallOpts)
}

// Commitments is a free data retrieval call binding the contract method 0xd13e2e60.
//
// Solidity: function commitments(uint256 , uint256 ) view returns(bytes32)
func (_VoteManager *VoteManagerCaller) Commitments(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _VoteManager.contract.Call(opts, &out, "commitments", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Commitments is a free data retrieval call binding the contract method 0xd13e2e60.
//
// Solidity: function commitments(uint256 , uint256 ) view returns(bytes32)
func (_VoteManager *VoteManagerSession) Commitments(arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	return _VoteManager.Contract.Commitments(&_VoteManager.CallOpts, arg0, arg1)
}

// Commitments is a free data retrieval call binding the contract method 0xd13e2e60.
//
// Solidity: function commitments(uint256 , uint256 ) view returns(bytes32)
func (_VoteManager *VoteManagerCallerSession) Commitments(arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	return _VoteManager.Contract.Commitments(&_VoteManager.CallOpts, arg0, arg1)
}

// GetCommitment is a free data retrieval call binding the contract method 0x7164f0d6.
//
// Solidity: function getCommitment(uint256 epoch, uint256 stakerId) view returns(bytes32)
func (_VoteManager *VoteManagerCaller) GetCommitment(opts *bind.CallOpts, epoch *big.Int, stakerId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _VoteManager.contract.Call(opts, &out, "getCommitment", epoch, stakerId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCommitment is a free data retrieval call binding the contract method 0x7164f0d6.
//
// Solidity: function getCommitment(uint256 epoch, uint256 stakerId) view returns(bytes32)
func (_VoteManager *VoteManagerSession) GetCommitment(epoch *big.Int, stakerId *big.Int) ([32]byte, error) {
	return _VoteManager.Contract.GetCommitment(&_VoteManager.CallOpts, epoch, stakerId)
}

// GetCommitment is a free data retrieval call binding the contract method 0x7164f0d6.
//
// Solidity: function getCommitment(uint256 epoch, uint256 stakerId) view returns(bytes32)
func (_VoteManager *VoteManagerCallerSession) GetCommitment(epoch *big.Int, stakerId *big.Int) ([32]byte, error) {
	return _VoteManager.Contract.GetCommitment(&_VoteManager.CallOpts, epoch, stakerId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VoteManager *VoteManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _VoteManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VoteManager *VoteManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _VoteManager.Contract.GetRoleAdmin(&_VoteManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VoteManager *VoteManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _VoteManager.Contract.GetRoleAdmin(&_VoteManager.CallOpts, role)
}

// GetTotalStakeRevealed is a free data retrieval call binding the contract method 0xa6a145d9.
//
// Solidity: function getTotalStakeRevealed(uint256 epoch, uint256 assetId) view returns(uint256)
func (_VoteManager *VoteManagerCaller) GetTotalStakeRevealed(opts *bind.CallOpts, epoch *big.Int, assetId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VoteManager.contract.Call(opts, &out, "getTotalStakeRevealed", epoch, assetId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalStakeRevealed is a free data retrieval call binding the contract method 0xa6a145d9.
//
// Solidity: function getTotalStakeRevealed(uint256 epoch, uint256 assetId) view returns(uint256)
func (_VoteManager *VoteManagerSession) GetTotalStakeRevealed(epoch *big.Int, assetId *big.Int) (*big.Int, error) {
	return _VoteManager.Contract.GetTotalStakeRevealed(&_VoteManager.CallOpts, epoch, assetId)
}

// GetTotalStakeRevealed is a free data retrieval call binding the contract method 0xa6a145d9.
//
// Solidity: function getTotalStakeRevealed(uint256 epoch, uint256 assetId) view returns(uint256)
func (_VoteManager *VoteManagerCallerSession) GetTotalStakeRevealed(epoch *big.Int, assetId *big.Int) (*big.Int, error) {
	return _VoteManager.Contract.GetTotalStakeRevealed(&_VoteManager.CallOpts, epoch, assetId)
}

// GetTotalStakeRevealed0 is a free data retrieval call binding the contract method 0xf322bd7d.
//
// Solidity: function getTotalStakeRevealed(uint256 epoch, uint256 assetId, uint256 voteValue) view returns(uint256)
func (_VoteManager *VoteManagerCaller) GetTotalStakeRevealed0(opts *bind.CallOpts, epoch *big.Int, assetId *big.Int, voteValue *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VoteManager.contract.Call(opts, &out, "getTotalStakeRevealed0", epoch, assetId, voteValue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalStakeRevealed0 is a free data retrieval call binding the contract method 0xf322bd7d.
//
// Solidity: function getTotalStakeRevealed(uint256 epoch, uint256 assetId, uint256 voteValue) view returns(uint256)
func (_VoteManager *VoteManagerSession) GetTotalStakeRevealed0(epoch *big.Int, assetId *big.Int, voteValue *big.Int) (*big.Int, error) {
	return _VoteManager.Contract.GetTotalStakeRevealed0(&_VoteManager.CallOpts, epoch, assetId, voteValue)
}

// GetTotalStakeRevealed0 is a free data retrieval call binding the contract method 0xf322bd7d.
//
// Solidity: function getTotalStakeRevealed(uint256 epoch, uint256 assetId, uint256 voteValue) view returns(uint256)
func (_VoteManager *VoteManagerCallerSession) GetTotalStakeRevealed0(epoch *big.Int, assetId *big.Int, voteValue *big.Int) (*big.Int, error) {
	return _VoteManager.Contract.GetTotalStakeRevealed0(&_VoteManager.CallOpts, epoch, assetId, voteValue)
}

// GetVote is a free data retrieval call binding the contract method 0x8ce7ff4a.
//
// Solidity: function getVote(uint256 epoch, uint256 stakerId, uint256 assetId) view returns((uint256,uint256) vote)
func (_VoteManager *VoteManagerCaller) GetVote(opts *bind.CallOpts, epoch *big.Int, stakerId *big.Int, assetId *big.Int) (StructsVote, error) {
	var out []interface{}
	err := _VoteManager.contract.Call(opts, &out, "getVote", epoch, stakerId, assetId)

	if err != nil {
		return *new(StructsVote), err
	}

	out0 := *abi.ConvertType(out[0], new(StructsVote)).(*StructsVote)

	return out0, err

}

// GetVote is a free data retrieval call binding the contract method 0x8ce7ff4a.
//
// Solidity: function getVote(uint256 epoch, uint256 stakerId, uint256 assetId) view returns((uint256,uint256) vote)
func (_VoteManager *VoteManagerSession) GetVote(epoch *big.Int, stakerId *big.Int, assetId *big.Int) (StructsVote, error) {
	return _VoteManager.Contract.GetVote(&_VoteManager.CallOpts, epoch, stakerId, assetId)
}

// GetVote is a free data retrieval call binding the contract method 0x8ce7ff4a.
//
// Solidity: function getVote(uint256 epoch, uint256 stakerId, uint256 assetId) view returns((uint256,uint256) vote)
func (_VoteManager *VoteManagerCallerSession) GetVote(epoch *big.Int, stakerId *big.Int, assetId *big.Int) (StructsVote, error) {
	return _VoteManager.Contract.GetVote(&_VoteManager.CallOpts, epoch, stakerId, assetId)
}

// GetVoteWeight is a free data retrieval call binding the contract method 0x9c66556f.
//
// Solidity: function getVoteWeight(uint256 epoch, uint256 assetId, uint256 voteValue) view returns(uint256)
func (_VoteManager *VoteManagerCaller) GetVoteWeight(opts *bind.CallOpts, epoch *big.Int, assetId *big.Int, voteValue *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VoteManager.contract.Call(opts, &out, "getVoteWeight", epoch, assetId, voteValue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVoteWeight is a free data retrieval call binding the contract method 0x9c66556f.
//
// Solidity: function getVoteWeight(uint256 epoch, uint256 assetId, uint256 voteValue) view returns(uint256)
func (_VoteManager *VoteManagerSession) GetVoteWeight(epoch *big.Int, assetId *big.Int, voteValue *big.Int) (*big.Int, error) {
	return _VoteManager.Contract.GetVoteWeight(&_VoteManager.CallOpts, epoch, assetId, voteValue)
}

// GetVoteWeight is a free data retrieval call binding the contract method 0x9c66556f.
//
// Solidity: function getVoteWeight(uint256 epoch, uint256 assetId, uint256 voteValue) view returns(uint256)
func (_VoteManager *VoteManagerCallerSession) GetVoteWeight(epoch *big.Int, assetId *big.Int, voteValue *big.Int) (*big.Int, error) {
	return _VoteManager.Contract.GetVoteWeight(&_VoteManager.CallOpts, epoch, assetId, voteValue)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VoteManager *VoteManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _VoteManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VoteManager *VoteManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _VoteManager.Contract.HasRole(&_VoteManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VoteManager *VoteManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _VoteManager.Contract.HasRole(&_VoteManager.CallOpts, role, account)
}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address)
func (_VoteManager *VoteManagerCaller) Parameters(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VoteManager.contract.Call(opts, &out, "parameters")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address)
func (_VoteManager *VoteManagerSession) Parameters() (common.Address, error) {
	return _VoteManager.Contract.Parameters(&_VoteManager.CallOpts)
}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address)
func (_VoteManager *VoteManagerCallerSession) Parameters() (common.Address, error) {
	return _VoteManager.Contract.Parameters(&_VoteManager.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_VoteManager *VoteManagerCaller) RewardManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VoteManager.contract.Call(opts, &out, "rewardManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_VoteManager *VoteManagerSession) RewardManager() (common.Address, error) {
	return _VoteManager.Contract.RewardManager(&_VoteManager.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_VoteManager *VoteManagerCallerSession) RewardManager() (common.Address, error) {
	return _VoteManager.Contract.RewardManager(&_VoteManager.CallOpts)
}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_VoteManager *VoteManagerCaller) StakeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VoteManager.contract.Call(opts, &out, "stakeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_VoteManager *VoteManagerSession) StakeManager() (common.Address, error) {
	return _VoteManager.Contract.StakeManager(&_VoteManager.CallOpts)
}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_VoteManager *VoteManagerCallerSession) StakeManager() (common.Address, error) {
	return _VoteManager.Contract.StakeManager(&_VoteManager.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VoteManager *VoteManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _VoteManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VoteManager *VoteManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VoteManager.Contract.SupportsInterface(&_VoteManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VoteManager *VoteManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VoteManager.Contract.SupportsInterface(&_VoteManager.CallOpts, interfaceId)
}

// TotalStakeRevealed is a free data retrieval call binding the contract method 0x8a757ecc.
//
// Solidity: function totalStakeRevealed(uint256 , uint256 ) view returns(uint256)
func (_VoteManager *VoteManagerCaller) TotalStakeRevealed(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VoteManager.contract.Call(opts, &out, "totalStakeRevealed", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStakeRevealed is a free data retrieval call binding the contract method 0x8a757ecc.
//
// Solidity: function totalStakeRevealed(uint256 , uint256 ) view returns(uint256)
func (_VoteManager *VoteManagerSession) TotalStakeRevealed(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _VoteManager.Contract.TotalStakeRevealed(&_VoteManager.CallOpts, arg0, arg1)
}

// TotalStakeRevealed is a free data retrieval call binding the contract method 0x8a757ecc.
//
// Solidity: function totalStakeRevealed(uint256 , uint256 ) view returns(uint256)
func (_VoteManager *VoteManagerCallerSession) TotalStakeRevealed(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _VoteManager.Contract.TotalStakeRevealed(&_VoteManager.CallOpts, arg0, arg1)
}

// VoteWeights is a free data retrieval call binding the contract method 0x8fd5ff00.
//
// Solidity: function voteWeights(uint256 , uint256 , uint256 ) view returns(uint256)
func (_VoteManager *VoteManagerCaller) VoteWeights(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VoteManager.contract.Call(opts, &out, "voteWeights", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoteWeights is a free data retrieval call binding the contract method 0x8fd5ff00.
//
// Solidity: function voteWeights(uint256 , uint256 , uint256 ) view returns(uint256)
func (_VoteManager *VoteManagerSession) VoteWeights(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (*big.Int, error) {
	return _VoteManager.Contract.VoteWeights(&_VoteManager.CallOpts, arg0, arg1, arg2)
}

// VoteWeights is a free data retrieval call binding the contract method 0x8fd5ff00.
//
// Solidity: function voteWeights(uint256 , uint256 , uint256 ) view returns(uint256)
func (_VoteManager *VoteManagerCallerSession) VoteWeights(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (*big.Int, error) {
	return _VoteManager.Contract.VoteWeights(&_VoteManager.CallOpts, arg0, arg1, arg2)
}

// Votes is a free data retrieval call binding the contract method 0x283e2905.
//
// Solidity: function votes(uint256 , uint256 , uint256 ) view returns(uint256 value, uint256 weight)
func (_VoteManager *VoteManagerCaller) Votes(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (struct {
	Value  *big.Int
	Weight *big.Int
}, error) {
	var out []interface{}
	err := _VoteManager.contract.Call(opts, &out, "votes", arg0, arg1, arg2)

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
func (_VoteManager *VoteManagerSession) Votes(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (struct {
	Value  *big.Int
	Weight *big.Int
}, error) {
	return _VoteManager.Contract.Votes(&_VoteManager.CallOpts, arg0, arg1, arg2)
}

// Votes is a free data retrieval call binding the contract method 0x283e2905.
//
// Solidity: function votes(uint256 , uint256 , uint256 ) view returns(uint256 value, uint256 weight)
func (_VoteManager *VoteManagerCallerSession) Votes(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (struct {
	Value  *big.Int
	Weight *big.Int
}, error) {
	return _VoteManager.Contract.Votes(&_VoteManager.CallOpts, arg0, arg1, arg2)
}

// Commit is a paid mutator transaction binding the contract method 0xf2f03877.
//
// Solidity: function commit(uint256 epoch, bytes32 commitment) returns()
func (_VoteManager *VoteManagerTransactor) Commit(opts *bind.TransactOpts, epoch *big.Int, commitment [32]byte) (*types.Transaction, error) {
	return _VoteManager.contract.Transact(opts, "commit", epoch, commitment)
}

// Commit is a paid mutator transaction binding the contract method 0xf2f03877.
//
// Solidity: function commit(uint256 epoch, bytes32 commitment) returns()
func (_VoteManager *VoteManagerSession) Commit(epoch *big.Int, commitment [32]byte) (*types.Transaction, error) {
	return _VoteManager.Contract.Commit(&_VoteManager.TransactOpts, epoch, commitment)
}

// Commit is a paid mutator transaction binding the contract method 0xf2f03877.
//
// Solidity: function commit(uint256 epoch, bytes32 commitment) returns()
func (_VoteManager *VoteManagerTransactorSession) Commit(epoch *big.Int, commitment [32]byte) (*types.Transaction, error) {
	return _VoteManager.Contract.Commit(&_VoteManager.TransactOpts, epoch, commitment)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VoteManager *VoteManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VoteManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VoteManager *VoteManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VoteManager.Contract.GrantRole(&_VoteManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VoteManager *VoteManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VoteManager.Contract.GrantRole(&_VoteManager.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address stakeManagerAddress, address rewardManagerAddress, address blockManagerAddress, address parametersAddress) returns()
func (_VoteManager *VoteManagerTransactor) Initialize(opts *bind.TransactOpts, stakeManagerAddress common.Address, rewardManagerAddress common.Address, blockManagerAddress common.Address, parametersAddress common.Address) (*types.Transaction, error) {
	return _VoteManager.contract.Transact(opts, "initialize", stakeManagerAddress, rewardManagerAddress, blockManagerAddress, parametersAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address stakeManagerAddress, address rewardManagerAddress, address blockManagerAddress, address parametersAddress) returns()
func (_VoteManager *VoteManagerSession) Initialize(stakeManagerAddress common.Address, rewardManagerAddress common.Address, blockManagerAddress common.Address, parametersAddress common.Address) (*types.Transaction, error) {
	return _VoteManager.Contract.Initialize(&_VoteManager.TransactOpts, stakeManagerAddress, rewardManagerAddress, blockManagerAddress, parametersAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address stakeManagerAddress, address rewardManagerAddress, address blockManagerAddress, address parametersAddress) returns()
func (_VoteManager *VoteManagerTransactorSession) Initialize(stakeManagerAddress common.Address, rewardManagerAddress common.Address, blockManagerAddress common.Address, parametersAddress common.Address) (*types.Transaction, error) {
	return _VoteManager.Contract.Initialize(&_VoteManager.TransactOpts, stakeManagerAddress, rewardManagerAddress, blockManagerAddress, parametersAddress)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_VoteManager *VoteManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VoteManager.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_VoteManager *VoteManagerSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VoteManager.Contract.RenounceRole(&_VoteManager.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_VoteManager *VoteManagerTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VoteManager.Contract.RenounceRole(&_VoteManager.TransactOpts, role, account)
}

// Reveal is a paid mutator transaction binding the contract method 0x125f68f1.
//
// Solidity: function reveal(uint256 epoch, bytes32 root, uint256[] values, bytes32[][] proofs, bytes32 secret, address stakerAddress) returns()
func (_VoteManager *VoteManagerTransactor) Reveal(opts *bind.TransactOpts, epoch *big.Int, root [32]byte, values []*big.Int, proofs [][][32]byte, secret [32]byte, stakerAddress common.Address) (*types.Transaction, error) {
	return _VoteManager.contract.Transact(opts, "reveal", epoch, root, values, proofs, secret, stakerAddress)
}

// Reveal is a paid mutator transaction binding the contract method 0x125f68f1.
//
// Solidity: function reveal(uint256 epoch, bytes32 root, uint256[] values, bytes32[][] proofs, bytes32 secret, address stakerAddress) returns()
func (_VoteManager *VoteManagerSession) Reveal(epoch *big.Int, root [32]byte, values []*big.Int, proofs [][][32]byte, secret [32]byte, stakerAddress common.Address) (*types.Transaction, error) {
	return _VoteManager.Contract.Reveal(&_VoteManager.TransactOpts, epoch, root, values, proofs, secret, stakerAddress)
}

// Reveal is a paid mutator transaction binding the contract method 0x125f68f1.
//
// Solidity: function reveal(uint256 epoch, bytes32 root, uint256[] values, bytes32[][] proofs, bytes32 secret, address stakerAddress) returns()
func (_VoteManager *VoteManagerTransactorSession) Reveal(epoch *big.Int, root [32]byte, values []*big.Int, proofs [][][32]byte, secret [32]byte, stakerAddress common.Address) (*types.Transaction, error) {
	return _VoteManager.Contract.Reveal(&_VoteManager.TransactOpts, epoch, root, values, proofs, secret, stakerAddress)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VoteManager *VoteManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VoteManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VoteManager *VoteManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VoteManager.Contract.RevokeRole(&_VoteManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VoteManager *VoteManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VoteManager.Contract.RevokeRole(&_VoteManager.TransactOpts, role, account)
}

// VoteManagerCommittedIterator is returned from FilterCommitted and is used to iterate over the raw logs and unpacked data for Committed events raised by the VoteManager contract.
type VoteManagerCommittedIterator struct {
	Event *VoteManagerCommitted // Event containing the contract specifics and raw log

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
func (it *VoteManagerCommittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteManagerCommitted)
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
		it.Event = new(VoteManagerCommitted)
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
func (it *VoteManagerCommittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteManagerCommittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteManagerCommitted represents a Committed event raised by the VoteManager contract.
type VoteManagerCommitted struct {
	Epoch      *big.Int
	StakerId   *big.Int
	Commitment [32]byte
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCommitted is a free log retrieval operation binding the contract event 0x6ad04b07dbe80ee3971460ebc13808fb6dd0fa67fbf7d7ebc4de85811d2c9928.
//
// Solidity: event Committed(uint256 epoch, uint256 stakerId, bytes32 commitment, uint256 timestamp)
func (_VoteManager *VoteManagerFilterer) FilterCommitted(opts *bind.FilterOpts) (*VoteManagerCommittedIterator, error) {

	logs, sub, err := _VoteManager.contract.FilterLogs(opts, "Committed")
	if err != nil {
		return nil, err
	}
	return &VoteManagerCommittedIterator{contract: _VoteManager.contract, event: "Committed", logs: logs, sub: sub}, nil
}

// WatchCommitted is a free log subscription operation binding the contract event 0x6ad04b07dbe80ee3971460ebc13808fb6dd0fa67fbf7d7ebc4de85811d2c9928.
//
// Solidity: event Committed(uint256 epoch, uint256 stakerId, bytes32 commitment, uint256 timestamp)
func (_VoteManager *VoteManagerFilterer) WatchCommitted(opts *bind.WatchOpts, sink chan<- *VoteManagerCommitted) (event.Subscription, error) {

	logs, sub, err := _VoteManager.contract.WatchLogs(opts, "Committed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteManagerCommitted)
				if err := _VoteManager.contract.UnpackLog(event, "Committed", log); err != nil {
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

// ParseCommitted is a log parse operation binding the contract event 0x6ad04b07dbe80ee3971460ebc13808fb6dd0fa67fbf7d7ebc4de85811d2c9928.
//
// Solidity: event Committed(uint256 epoch, uint256 stakerId, bytes32 commitment, uint256 timestamp)
func (_VoteManager *VoteManagerFilterer) ParseCommitted(log types.Log) (*VoteManagerCommitted, error) {
	event := new(VoteManagerCommitted)
	if err := _VoteManager.contract.UnpackLog(event, "Committed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteManagerRevealedIterator is returned from FilterRevealed and is used to iterate over the raw logs and unpacked data for Revealed events raised by the VoteManager contract.
type VoteManagerRevealedIterator struct {
	Event *VoteManagerRevealed // Event containing the contract specifics and raw log

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
func (it *VoteManagerRevealedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteManagerRevealed)
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
		it.Event = new(VoteManagerRevealed)
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
func (it *VoteManagerRevealedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteManagerRevealedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteManagerRevealed represents a Revealed event raised by the VoteManager contract.
type VoteManagerRevealed struct {
	Epoch     *big.Int
	StakerId  *big.Int
	Stake     *big.Int
	Values    []*big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRevealed is a free log retrieval operation binding the contract event 0x17f4759c403dfc3fa26ab75ce82d3d090a0bdbfe9977d8ee87445d6dd83c68a1.
//
// Solidity: event Revealed(uint256 epoch, uint256 stakerId, uint256 stake, uint256[] values, uint256 timestamp)
func (_VoteManager *VoteManagerFilterer) FilterRevealed(opts *bind.FilterOpts) (*VoteManagerRevealedIterator, error) {

	logs, sub, err := _VoteManager.contract.FilterLogs(opts, "Revealed")
	if err != nil {
		return nil, err
	}
	return &VoteManagerRevealedIterator{contract: _VoteManager.contract, event: "Revealed", logs: logs, sub: sub}, nil
}

// WatchRevealed is a free log subscription operation binding the contract event 0x17f4759c403dfc3fa26ab75ce82d3d090a0bdbfe9977d8ee87445d6dd83c68a1.
//
// Solidity: event Revealed(uint256 epoch, uint256 stakerId, uint256 stake, uint256[] values, uint256 timestamp)
func (_VoteManager *VoteManagerFilterer) WatchRevealed(opts *bind.WatchOpts, sink chan<- *VoteManagerRevealed) (event.Subscription, error) {

	logs, sub, err := _VoteManager.contract.WatchLogs(opts, "Revealed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteManagerRevealed)
				if err := _VoteManager.contract.UnpackLog(event, "Revealed", log); err != nil {
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

// ParseRevealed is a log parse operation binding the contract event 0x17f4759c403dfc3fa26ab75ce82d3d090a0bdbfe9977d8ee87445d6dd83c68a1.
//
// Solidity: event Revealed(uint256 epoch, uint256 stakerId, uint256 stake, uint256[] values, uint256 timestamp)
func (_VoteManager *VoteManagerFilterer) ParseRevealed(log types.Log) (*VoteManagerRevealed, error) {
	event := new(VoteManagerRevealed)
	if err := _VoteManager.contract.UnpackLog(event, "Revealed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the VoteManager contract.
type VoteManagerRoleAdminChangedIterator struct {
	Event *VoteManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *VoteManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteManagerRoleAdminChanged)
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
		it.Event = new(VoteManagerRoleAdminChanged)
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
func (it *VoteManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteManagerRoleAdminChanged represents a RoleAdminChanged event raised by the VoteManager contract.
type VoteManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_VoteManager *VoteManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*VoteManagerRoleAdminChangedIterator, error) {

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

	logs, sub, err := _VoteManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &VoteManagerRoleAdminChangedIterator{contract: _VoteManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_VoteManager *VoteManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *VoteManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _VoteManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteManagerRoleAdminChanged)
				if err := _VoteManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_VoteManager *VoteManagerFilterer) ParseRoleAdminChanged(log types.Log) (*VoteManagerRoleAdminChanged, error) {
	event := new(VoteManagerRoleAdminChanged)
	if err := _VoteManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the VoteManager contract.
type VoteManagerRoleGrantedIterator struct {
	Event *VoteManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *VoteManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteManagerRoleGranted)
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
		it.Event = new(VoteManagerRoleGranted)
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
func (it *VoteManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteManagerRoleGranted represents a RoleGranted event raised by the VoteManager contract.
type VoteManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_VoteManager *VoteManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VoteManagerRoleGrantedIterator, error) {

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

	logs, sub, err := _VoteManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VoteManagerRoleGrantedIterator{contract: _VoteManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_VoteManager *VoteManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *VoteManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _VoteManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteManagerRoleGranted)
				if err := _VoteManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_VoteManager *VoteManagerFilterer) ParseRoleGranted(log types.Log) (*VoteManagerRoleGranted, error) {
	event := new(VoteManagerRoleGranted)
	if err := _VoteManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the VoteManager contract.
type VoteManagerRoleRevokedIterator struct {
	Event *VoteManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *VoteManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteManagerRoleRevoked)
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
		it.Event = new(VoteManagerRoleRevoked)
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
func (it *VoteManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteManagerRoleRevoked represents a RoleRevoked event raised by the VoteManager contract.
type VoteManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_VoteManager *VoteManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VoteManagerRoleRevokedIterator, error) {

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

	logs, sub, err := _VoteManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VoteManagerRoleRevokedIterator{contract: _VoteManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_VoteManager *VoteManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *VoteManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _VoteManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteManagerRoleRevoked)
				if err := _VoteManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_VoteManager *VoteManagerFilterer) ParseRoleRevoked(log types.Log) (*VoteManagerRoleRevoked, error) {
	event := new(VoteManagerRoleRevoked)
	if err := _VoteManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
