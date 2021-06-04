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

// BlockManagerABI is the input ABI used to generate the binding from.
const BlockManagerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"medians\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"lowerCutoffs\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"higherCutoffs\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"jobIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BlockConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"jobIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"medians\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"lowerCutoffs\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"higherCutoffs\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"iteration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"biggestStakerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Proposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"blocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"proposerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iteration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"biggestStake\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"disputes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"accWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"median\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lowerCutoff\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"higherCutoff\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastVisited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockId\",\"type\":\"uint256\"}],\"name\":\"finalizeDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getBlock\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"proposerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"jobIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"medians\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"lowerCutoffs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"higherCutoffs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"iteration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"biggestStake\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"internalType\":\"structStructs.Block\",\"name\":\"_block\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getBlockMedians\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_blockMedians\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getHigherCutoffs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_higherCutoffs\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getLowerCutoffs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_lowerCutoffs\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getNumProposedBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposedBlock\",\"type\":\"uint256\"}],\"name\":\"getProposedBlock\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"proposerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"jobIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"medians\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"lowerCutoffs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"higherCutoffs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"iteration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"biggestStake\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"internalType\":\"structStructs.Block\",\"name\":\"_block\",\"type\":\"tuple\"},{\"internalType\":\"uint256[]\",\"name\":\"_blockMedians\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_lowerCutoffs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_higherCutoffs\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposedBlock\",\"type\":\"uint256\"}],\"name\":\"getProposedBlockMedians\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_blockMedians\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"sorted\",\"type\":\"uint256[]\"}],\"name\":\"giveSorted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakeManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stateManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_voteManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_jobManagerAddress\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"iteration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"biggestStakerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakerId\",\"type\":\"uint256\"}],\"name\":\"isElectedProposer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"jobManager\",\"outputs\":[{\"internalType\":\"contractIJobManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"jobIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"medians\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"lowerCutoffs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"higherCutoffs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"iteration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"biggestStakerId\",\"type\":\"uint256\"}],\"name\":\"propose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposedBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"proposerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iteration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"biggestStake\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"resetDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeManager\",\"outputs\":[{\"internalType\":\"contractIStakeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stateManager\",\"outputs\":[{\"internalType\":\"contractIStateManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteManager\",\"outputs\":[{\"internalType\":\"contractIVoteManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BlockManager is an auto generated Go binding around an Ethereum contract.
type BlockManager struct {
	BlockManagerCaller     // Read-only binding to the contract
	BlockManagerTransactor // Write-only binding to the contract
	BlockManagerFilterer   // Log filterer for contract events
}

// BlockManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlockManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockManagerSession struct {
	Contract     *BlockManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlockManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockManagerCallerSession struct {
	Contract *BlockManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BlockManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockManagerTransactorSession struct {
	Contract     *BlockManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BlockManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlockManagerRaw struct {
	Contract *BlockManager // Generic contract binding to access the raw methods on
}

// BlockManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockManagerCallerRaw struct {
	Contract *BlockManagerCaller // Generic read-only contract binding to access the raw methods on
}

// BlockManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockManagerTransactorRaw struct {
	Contract *BlockManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockManager creates a new instance of BlockManager, bound to a specific deployed contract.
func NewBlockManager(address common.Address, backend bind.ContractBackend) (*BlockManager, error) {
	contract, err := bindBlockManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlockManager{BlockManagerCaller: BlockManagerCaller{contract: contract}, BlockManagerTransactor: BlockManagerTransactor{contract: contract}, BlockManagerFilterer: BlockManagerFilterer{contract: contract}}, nil
}

// NewBlockManagerCaller creates a new read-only instance of BlockManager, bound to a specific deployed contract.
func NewBlockManagerCaller(address common.Address, caller bind.ContractCaller) (*BlockManagerCaller, error) {
	contract, err := bindBlockManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockManagerCaller{contract: contract}, nil
}

// NewBlockManagerTransactor creates a new write-only instance of BlockManager, bound to a specific deployed contract.
func NewBlockManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockManagerTransactor, error) {
	contract, err := bindBlockManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockManagerTransactor{contract: contract}, nil
}

// NewBlockManagerFilterer creates a new log filterer instance of BlockManager, bound to a specific deployed contract.
func NewBlockManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockManagerFilterer, error) {
	contract, err := bindBlockManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockManagerFilterer{contract: contract}, nil
}

// bindBlockManager binds a generic wrapper to an already deployed contract.
func bindBlockManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlockManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockManager *BlockManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockManager.Contract.BlockManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockManager *BlockManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockManager.Contract.BlockManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockManager *BlockManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockManager.Contract.BlockManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockManager *BlockManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockManager *BlockManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockManager *BlockManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockManager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BlockManager *BlockManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BlockManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BlockManager *BlockManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BlockManager.Contract.DEFAULTADMINROLE(&_BlockManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BlockManager *BlockManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BlockManager.Contract.DEFAULTADMINROLE(&_BlockManager.CallOpts)
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks(uint256 ) view returns(uint256 proposerId, uint256 iteration, uint256 biggestStake, bool valid)
func (_BlockManager *BlockManagerCaller) Blocks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ProposerId   *big.Int
	Iteration    *big.Int
	BiggestStake *big.Int
	Valid        bool
}, error) {
	var out []interface{}
	err := _BlockManager.contract.Call(opts, &out, "blocks", arg0)

	outstruct := new(struct {
		ProposerId   *big.Int
		Iteration    *big.Int
		BiggestStake *big.Int
		Valid        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProposerId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Iteration = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BiggestStake = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Valid = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks(uint256 ) view returns(uint256 proposerId, uint256 iteration, uint256 biggestStake, bool valid)
func (_BlockManager *BlockManagerSession) Blocks(arg0 *big.Int) (struct {
	ProposerId   *big.Int
	Iteration    *big.Int
	BiggestStake *big.Int
	Valid        bool
}, error) {
	return _BlockManager.Contract.Blocks(&_BlockManager.CallOpts, arg0)
}

// Blocks is a free data retrieval call binding the contract method 0xf25b3f99.
//
// Solidity: function blocks(uint256 ) view returns(uint256 proposerId, uint256 iteration, uint256 biggestStake, bool valid)
func (_BlockManager *BlockManagerCallerSession) Blocks(arg0 *big.Int) (struct {
	ProposerId   *big.Int
	Iteration    *big.Int
	BiggestStake *big.Int
	Valid        bool
}, error) {
	return _BlockManager.Contract.Blocks(&_BlockManager.CallOpts, arg0)
}

// Disputes is a free data retrieval call binding the contract method 0x828496d6.
//
// Solidity: function disputes(uint256 , address ) view returns(uint256 accWeight, uint256 median, uint256 lowerCutoff, uint256 higherCutoff, uint256 lastVisited, uint256 assetId)
func (_BlockManager *BlockManagerCaller) Disputes(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	AccWeight    *big.Int
	Median       *big.Int
	LowerCutoff  *big.Int
	HigherCutoff *big.Int
	LastVisited  *big.Int
	AssetId      *big.Int
}, error) {
	var out []interface{}
	err := _BlockManager.contract.Call(opts, &out, "disputes", arg0, arg1)

	outstruct := new(struct {
		AccWeight    *big.Int
		Median       *big.Int
		LowerCutoff  *big.Int
		HigherCutoff *big.Int
		LastVisited  *big.Int
		AssetId      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AccWeight = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Median = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LowerCutoff = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.HigherCutoff = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LastVisited = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.AssetId = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Disputes is a free data retrieval call binding the contract method 0x828496d6.
//
// Solidity: function disputes(uint256 , address ) view returns(uint256 accWeight, uint256 median, uint256 lowerCutoff, uint256 higherCutoff, uint256 lastVisited, uint256 assetId)
func (_BlockManager *BlockManagerSession) Disputes(arg0 *big.Int, arg1 common.Address) (struct {
	AccWeight    *big.Int
	Median       *big.Int
	LowerCutoff  *big.Int
	HigherCutoff *big.Int
	LastVisited  *big.Int
	AssetId      *big.Int
}, error) {
	return _BlockManager.Contract.Disputes(&_BlockManager.CallOpts, arg0, arg1)
}

// Disputes is a free data retrieval call binding the contract method 0x828496d6.
//
// Solidity: function disputes(uint256 , address ) view returns(uint256 accWeight, uint256 median, uint256 lowerCutoff, uint256 higherCutoff, uint256 lastVisited, uint256 assetId)
func (_BlockManager *BlockManagerCallerSession) Disputes(arg0 *big.Int, arg1 common.Address) (struct {
	AccWeight    *big.Int
	Median       *big.Int
	LowerCutoff  *big.Int
	HigherCutoff *big.Int
	LastVisited  *big.Int
	AssetId      *big.Int
}, error) {
	return _BlockManager.Contract.Disputes(&_BlockManager.CallOpts, arg0, arg1)
}

// GetBlock is a free data retrieval call binding the contract method 0x04c07569.
//
// Solidity: function getBlock(uint256 epoch) view returns((uint256,uint256[],uint256[],uint256[],uint256[],uint256,uint256,bool) _block)
func (_BlockManager *BlockManagerCaller) GetBlock(opts *bind.CallOpts, epoch *big.Int) (StructsBlock, error) {
	var out []interface{}
	err := _BlockManager.contract.Call(opts, &out, "getBlock", epoch)

	if err != nil {
		return *new(StructsBlock), err
	}

	out0 := *abi.ConvertType(out[0], new(StructsBlock)).(*StructsBlock)

	return out0, err

}

// GetBlock is a free data retrieval call binding the contract method 0x04c07569.
//
// Solidity: function getBlock(uint256 epoch) view returns((uint256,uint256[],uint256[],uint256[],uint256[],uint256,uint256,bool) _block)
func (_BlockManager *BlockManagerSession) GetBlock(epoch *big.Int) (StructsBlock, error) {
	return _BlockManager.Contract.GetBlock(&_BlockManager.CallOpts, epoch)
}

// GetBlock is a free data retrieval call binding the contract method 0x04c07569.
//
// Solidity: function getBlock(uint256 epoch) view returns((uint256,uint256[],uint256[],uint256[],uint256[],uint256,uint256,bool) _block)
func (_BlockManager *BlockManagerCallerSession) GetBlock(epoch *big.Int) (StructsBlock, error) {
	return _BlockManager.Contract.GetBlock(&_BlockManager.CallOpts, epoch)
}

// GetBlockMedians is a free data retrieval call binding the contract method 0x378ab9a9.
//
// Solidity: function getBlockMedians(uint256 epoch) view returns(uint256[] _blockMedians)
func (_BlockManager *BlockManagerCaller) GetBlockMedians(opts *bind.CallOpts, epoch *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _BlockManager.contract.Call(opts, &out, "getBlockMedians", epoch)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetBlockMedians is a free data retrieval call binding the contract method 0x378ab9a9.
//
// Solidity: function getBlockMedians(uint256 epoch) view returns(uint256[] _blockMedians)
func (_BlockManager *BlockManagerSession) GetBlockMedians(epoch *big.Int) ([]*big.Int, error) {
	return _BlockManager.Contract.GetBlockMedians(&_BlockManager.CallOpts, epoch)
}

// GetBlockMedians is a free data retrieval call binding the contract method 0x378ab9a9.
//
// Solidity: function getBlockMedians(uint256 epoch) view returns(uint256[] _blockMedians)
func (_BlockManager *BlockManagerCallerSession) GetBlockMedians(epoch *big.Int) ([]*big.Int, error) {
	return _BlockManager.Contract.GetBlockMedians(&_BlockManager.CallOpts, epoch)
}

// GetHigherCutoffs is a free data retrieval call binding the contract method 0xfae4425d.
//
// Solidity: function getHigherCutoffs(uint256 epoch) view returns(uint256[] _higherCutoffs)
func (_BlockManager *BlockManagerCaller) GetHigherCutoffs(opts *bind.CallOpts, epoch *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _BlockManager.contract.Call(opts, &out, "getHigherCutoffs", epoch)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetHigherCutoffs is a free data retrieval call binding the contract method 0xfae4425d.
//
// Solidity: function getHigherCutoffs(uint256 epoch) view returns(uint256[] _higherCutoffs)
func (_BlockManager *BlockManagerSession) GetHigherCutoffs(epoch *big.Int) ([]*big.Int, error) {
	return _BlockManager.Contract.GetHigherCutoffs(&_BlockManager.CallOpts, epoch)
}

// GetHigherCutoffs is a free data retrieval call binding the contract method 0xfae4425d.
//
// Solidity: function getHigherCutoffs(uint256 epoch) view returns(uint256[] _higherCutoffs)
func (_BlockManager *BlockManagerCallerSession) GetHigherCutoffs(epoch *big.Int) ([]*big.Int, error) {
	return _BlockManager.Contract.GetHigherCutoffs(&_BlockManager.CallOpts, epoch)
}

// GetLowerCutoffs is a free data retrieval call binding the contract method 0xd2a4669a.
//
// Solidity: function getLowerCutoffs(uint256 epoch) view returns(uint256[] _lowerCutoffs)
func (_BlockManager *BlockManagerCaller) GetLowerCutoffs(opts *bind.CallOpts, epoch *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _BlockManager.contract.Call(opts, &out, "getLowerCutoffs", epoch)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetLowerCutoffs is a free data retrieval call binding the contract method 0xd2a4669a.
//
// Solidity: function getLowerCutoffs(uint256 epoch) view returns(uint256[] _lowerCutoffs)
func (_BlockManager *BlockManagerSession) GetLowerCutoffs(epoch *big.Int) ([]*big.Int, error) {
	return _BlockManager.Contract.GetLowerCutoffs(&_BlockManager.CallOpts, epoch)
}

// GetLowerCutoffs is a free data retrieval call binding the contract method 0xd2a4669a.
//
// Solidity: function getLowerCutoffs(uint256 epoch) view returns(uint256[] _lowerCutoffs)
func (_BlockManager *BlockManagerCallerSession) GetLowerCutoffs(epoch *big.Int) ([]*big.Int, error) {
	return _BlockManager.Contract.GetLowerCutoffs(&_BlockManager.CallOpts, epoch)
}

// GetNumProposedBlocks is a free data retrieval call binding the contract method 0xe38c7c42.
//
// Solidity: function getNumProposedBlocks(uint256 epoch) view returns(uint256)
func (_BlockManager *BlockManagerCaller) GetNumProposedBlocks(opts *bind.CallOpts, epoch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BlockManager.contract.Call(opts, &out, "getNumProposedBlocks", epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumProposedBlocks is a free data retrieval call binding the contract method 0xe38c7c42.
//
// Solidity: function getNumProposedBlocks(uint256 epoch) view returns(uint256)
func (_BlockManager *BlockManagerSession) GetNumProposedBlocks(epoch *big.Int) (*big.Int, error) {
	return _BlockManager.Contract.GetNumProposedBlocks(&_BlockManager.CallOpts, epoch)
}

// GetNumProposedBlocks is a free data retrieval call binding the contract method 0xe38c7c42.
//
// Solidity: function getNumProposedBlocks(uint256 epoch) view returns(uint256)
func (_BlockManager *BlockManagerCallerSession) GetNumProposedBlocks(epoch *big.Int) (*big.Int, error) {
	return _BlockManager.Contract.GetNumProposedBlocks(&_BlockManager.CallOpts, epoch)
}

// GetProposedBlock is a free data retrieval call binding the contract method 0xa27ce1ef.
//
// Solidity: function getProposedBlock(uint256 epoch, uint256 proposedBlock) view returns((uint256,uint256[],uint256[],uint256[],uint256[],uint256,uint256,bool) _block, uint256[] _blockMedians, uint256[] _lowerCutoffs, uint256[] _higherCutoffs)
func (_BlockManager *BlockManagerCaller) GetProposedBlock(opts *bind.CallOpts, epoch *big.Int, proposedBlock *big.Int) (struct {
	Block         StructsBlock
	BlockMedians  []*big.Int
	LowerCutoffs  []*big.Int
	HigherCutoffs []*big.Int
}, error) {
	var out []interface{}
	err := _BlockManager.contract.Call(opts, &out, "getProposedBlock", epoch, proposedBlock)

	outstruct := new(struct {
		Block         StructsBlock
		BlockMedians  []*big.Int
		LowerCutoffs  []*big.Int
		HigherCutoffs []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Block = *abi.ConvertType(out[0], new(StructsBlock)).(*StructsBlock)
	outstruct.BlockMedians = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.LowerCutoffs = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	outstruct.HigherCutoffs = *abi.ConvertType(out[3], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetProposedBlock is a free data retrieval call binding the contract method 0xa27ce1ef.
//
// Solidity: function getProposedBlock(uint256 epoch, uint256 proposedBlock) view returns((uint256,uint256[],uint256[],uint256[],uint256[],uint256,uint256,bool) _block, uint256[] _blockMedians, uint256[] _lowerCutoffs, uint256[] _higherCutoffs)
func (_BlockManager *BlockManagerSession) GetProposedBlock(epoch *big.Int, proposedBlock *big.Int) (struct {
	Block         StructsBlock
	BlockMedians  []*big.Int
	LowerCutoffs  []*big.Int
	HigherCutoffs []*big.Int
}, error) {
	return _BlockManager.Contract.GetProposedBlock(&_BlockManager.CallOpts, epoch, proposedBlock)
}

// GetProposedBlock is a free data retrieval call binding the contract method 0xa27ce1ef.
//
// Solidity: function getProposedBlock(uint256 epoch, uint256 proposedBlock) view returns((uint256,uint256[],uint256[],uint256[],uint256[],uint256,uint256,bool) _block, uint256[] _blockMedians, uint256[] _lowerCutoffs, uint256[] _higherCutoffs)
func (_BlockManager *BlockManagerCallerSession) GetProposedBlock(epoch *big.Int, proposedBlock *big.Int) (struct {
	Block         StructsBlock
	BlockMedians  []*big.Int
	LowerCutoffs  []*big.Int
	HigherCutoffs []*big.Int
}, error) {
	return _BlockManager.Contract.GetProposedBlock(&_BlockManager.CallOpts, epoch, proposedBlock)
}

// GetProposedBlockMedians is a free data retrieval call binding the contract method 0xd1a4a43d.
//
// Solidity: function getProposedBlockMedians(uint256 epoch, uint256 proposedBlock) view returns(uint256[] _blockMedians)
func (_BlockManager *BlockManagerCaller) GetProposedBlockMedians(opts *bind.CallOpts, epoch *big.Int, proposedBlock *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _BlockManager.contract.Call(opts, &out, "getProposedBlockMedians", epoch, proposedBlock)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetProposedBlockMedians is a free data retrieval call binding the contract method 0xd1a4a43d.
//
// Solidity: function getProposedBlockMedians(uint256 epoch, uint256 proposedBlock) view returns(uint256[] _blockMedians)
func (_BlockManager *BlockManagerSession) GetProposedBlockMedians(epoch *big.Int, proposedBlock *big.Int) ([]*big.Int, error) {
	return _BlockManager.Contract.GetProposedBlockMedians(&_BlockManager.CallOpts, epoch, proposedBlock)
}

// GetProposedBlockMedians is a free data retrieval call binding the contract method 0xd1a4a43d.
//
// Solidity: function getProposedBlockMedians(uint256 epoch, uint256 proposedBlock) view returns(uint256[] _blockMedians)
func (_BlockManager *BlockManagerCallerSession) GetProposedBlockMedians(epoch *big.Int, proposedBlock *big.Int) ([]*big.Int, error) {
	return _BlockManager.Contract.GetProposedBlockMedians(&_BlockManager.CallOpts, epoch, proposedBlock)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BlockManager *BlockManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BlockManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BlockManager *BlockManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BlockManager.Contract.GetRoleAdmin(&_BlockManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BlockManager *BlockManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BlockManager.Contract.GetRoleAdmin(&_BlockManager.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BlockManager *BlockManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BlockManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BlockManager *BlockManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BlockManager.Contract.HasRole(&_BlockManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BlockManager *BlockManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BlockManager.Contract.HasRole(&_BlockManager.CallOpts, role, account)
}

// IsElectedProposer is a free data retrieval call binding the contract method 0x1d69ff9b.
//
// Solidity: function isElectedProposer(uint256 iteration, uint256 biggestStakerId, uint256 stakerId) view returns(bool)
func (_BlockManager *BlockManagerCaller) IsElectedProposer(opts *bind.CallOpts, iteration *big.Int, biggestStakerId *big.Int, stakerId *big.Int) (bool, error) {
	var out []interface{}
	err := _BlockManager.contract.Call(opts, &out, "isElectedProposer", iteration, biggestStakerId, stakerId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsElectedProposer is a free data retrieval call binding the contract method 0x1d69ff9b.
//
// Solidity: function isElectedProposer(uint256 iteration, uint256 biggestStakerId, uint256 stakerId) view returns(bool)
func (_BlockManager *BlockManagerSession) IsElectedProposer(iteration *big.Int, biggestStakerId *big.Int, stakerId *big.Int) (bool, error) {
	return _BlockManager.Contract.IsElectedProposer(&_BlockManager.CallOpts, iteration, biggestStakerId, stakerId)
}

// IsElectedProposer is a free data retrieval call binding the contract method 0x1d69ff9b.
//
// Solidity: function isElectedProposer(uint256 iteration, uint256 biggestStakerId, uint256 stakerId) view returns(bool)
func (_BlockManager *BlockManagerCallerSession) IsElectedProposer(iteration *big.Int, biggestStakerId *big.Int, stakerId *big.Int) (bool, error) {
	return _BlockManager.Contract.IsElectedProposer(&_BlockManager.CallOpts, iteration, biggestStakerId, stakerId)
}

// JobManager is a free data retrieval call binding the contract method 0x3df395a3.
//
// Solidity: function jobManager() view returns(address)
func (_BlockManager *BlockManagerCaller) JobManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlockManager.contract.Call(opts, &out, "jobManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// JobManager is a free data retrieval call binding the contract method 0x3df395a3.
//
// Solidity: function jobManager() view returns(address)
func (_BlockManager *BlockManagerSession) JobManager() (common.Address, error) {
	return _BlockManager.Contract.JobManager(&_BlockManager.CallOpts)
}

// JobManager is a free data retrieval call binding the contract method 0x3df395a3.
//
// Solidity: function jobManager() view returns(address)
func (_BlockManager *BlockManagerCallerSession) JobManager() (common.Address, error) {
	return _BlockManager.Contract.JobManager(&_BlockManager.CallOpts)
}

// ProposedBlocks is a free data retrieval call binding the contract method 0x92b48411.
//
// Solidity: function proposedBlocks(uint256 , uint256 ) view returns(uint256 proposerId, uint256 iteration, uint256 biggestStake, bool valid)
func (_BlockManager *BlockManagerCaller) ProposedBlocks(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	ProposerId   *big.Int
	Iteration    *big.Int
	BiggestStake *big.Int
	Valid        bool
}, error) {
	var out []interface{}
	err := _BlockManager.contract.Call(opts, &out, "proposedBlocks", arg0, arg1)

	outstruct := new(struct {
		ProposerId   *big.Int
		Iteration    *big.Int
		BiggestStake *big.Int
		Valid        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProposerId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Iteration = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BiggestStake = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Valid = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// ProposedBlocks is a free data retrieval call binding the contract method 0x92b48411.
//
// Solidity: function proposedBlocks(uint256 , uint256 ) view returns(uint256 proposerId, uint256 iteration, uint256 biggestStake, bool valid)
func (_BlockManager *BlockManagerSession) ProposedBlocks(arg0 *big.Int, arg1 *big.Int) (struct {
	ProposerId   *big.Int
	Iteration    *big.Int
	BiggestStake *big.Int
	Valid        bool
}, error) {
	return _BlockManager.Contract.ProposedBlocks(&_BlockManager.CallOpts, arg0, arg1)
}

// ProposedBlocks is a free data retrieval call binding the contract method 0x92b48411.
//
// Solidity: function proposedBlocks(uint256 , uint256 ) view returns(uint256 proposerId, uint256 iteration, uint256 biggestStake, bool valid)
func (_BlockManager *BlockManagerCallerSession) ProposedBlocks(arg0 *big.Int, arg1 *big.Int) (struct {
	ProposerId   *big.Int
	Iteration    *big.Int
	BiggestStake *big.Int
	Valid        bool
}, error) {
	return _BlockManager.Contract.ProposedBlocks(&_BlockManager.CallOpts, arg0, arg1)
}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_BlockManager *BlockManagerCaller) StakeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlockManager.contract.Call(opts, &out, "stakeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_BlockManager *BlockManagerSession) StakeManager() (common.Address, error) {
	return _BlockManager.Contract.StakeManager(&_BlockManager.CallOpts)
}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_BlockManager *BlockManagerCallerSession) StakeManager() (common.Address, error) {
	return _BlockManager.Contract.StakeManager(&_BlockManager.CallOpts)
}

// StateManager is a free data retrieval call binding the contract method 0x2e716fb1.
//
// Solidity: function stateManager() view returns(address)
func (_BlockManager *BlockManagerCaller) StateManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlockManager.contract.Call(opts, &out, "stateManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StateManager is a free data retrieval call binding the contract method 0x2e716fb1.
//
// Solidity: function stateManager() view returns(address)
func (_BlockManager *BlockManagerSession) StateManager() (common.Address, error) {
	return _BlockManager.Contract.StateManager(&_BlockManager.CallOpts)
}

// StateManager is a free data retrieval call binding the contract method 0x2e716fb1.
//
// Solidity: function stateManager() view returns(address)
func (_BlockManager *BlockManagerCallerSession) StateManager() (common.Address, error) {
	return _BlockManager.Contract.StateManager(&_BlockManager.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BlockManager *BlockManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BlockManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BlockManager *BlockManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BlockManager.Contract.SupportsInterface(&_BlockManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BlockManager *BlockManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BlockManager.Contract.SupportsInterface(&_BlockManager.CallOpts, interfaceId)
}

// VoteManager is a free data retrieval call binding the contract method 0x42c1e587.
//
// Solidity: function voteManager() view returns(address)
func (_BlockManager *BlockManagerCaller) VoteManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlockManager.contract.Call(opts, &out, "voteManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoteManager is a free data retrieval call binding the contract method 0x42c1e587.
//
// Solidity: function voteManager() view returns(address)
func (_BlockManager *BlockManagerSession) VoteManager() (common.Address, error) {
	return _BlockManager.Contract.VoteManager(&_BlockManager.CallOpts)
}

// VoteManager is a free data retrieval call binding the contract method 0x42c1e587.
//
// Solidity: function voteManager() view returns(address)
func (_BlockManager *BlockManagerCallerSession) VoteManager() (common.Address, error) {
	return _BlockManager.Contract.VoteManager(&_BlockManager.CallOpts)
}

// ConfirmBlock is a paid mutator transaction binding the contract method 0x9b87f644.
//
// Solidity: function confirmBlock() returns()
func (_BlockManager *BlockManagerTransactor) ConfirmBlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockManager.contract.Transact(opts, "confirmBlock")
}

// ConfirmBlock is a paid mutator transaction binding the contract method 0x9b87f644.
//
// Solidity: function confirmBlock() returns()
func (_BlockManager *BlockManagerSession) ConfirmBlock() (*types.Transaction, error) {
	return _BlockManager.Contract.ConfirmBlock(&_BlockManager.TransactOpts)
}

// ConfirmBlock is a paid mutator transaction binding the contract method 0x9b87f644.
//
// Solidity: function confirmBlock() returns()
func (_BlockManager *BlockManagerTransactorSession) ConfirmBlock() (*types.Transaction, error) {
	return _BlockManager.Contract.ConfirmBlock(&_BlockManager.TransactOpts)
}

// FinalizeDispute is a paid mutator transaction binding the contract method 0x3cd32fd4.
//
// Solidity: function finalizeDispute(uint256 epoch, uint256 blockId) returns()
func (_BlockManager *BlockManagerTransactor) FinalizeDispute(opts *bind.TransactOpts, epoch *big.Int, blockId *big.Int) (*types.Transaction, error) {
	return _BlockManager.contract.Transact(opts, "finalizeDispute", epoch, blockId)
}

// FinalizeDispute is a paid mutator transaction binding the contract method 0x3cd32fd4.
//
// Solidity: function finalizeDispute(uint256 epoch, uint256 blockId) returns()
func (_BlockManager *BlockManagerSession) FinalizeDispute(epoch *big.Int, blockId *big.Int) (*types.Transaction, error) {
	return _BlockManager.Contract.FinalizeDispute(&_BlockManager.TransactOpts, epoch, blockId)
}

// FinalizeDispute is a paid mutator transaction binding the contract method 0x3cd32fd4.
//
// Solidity: function finalizeDispute(uint256 epoch, uint256 blockId) returns()
func (_BlockManager *BlockManagerTransactorSession) FinalizeDispute(epoch *big.Int, blockId *big.Int) (*types.Transaction, error) {
	return _BlockManager.Contract.FinalizeDispute(&_BlockManager.TransactOpts, epoch, blockId)
}

// GiveSorted is a paid mutator transaction binding the contract method 0x4e6753b7.
//
// Solidity: function giveSorted(uint256 epoch, uint256 assetId, uint256[] sorted) returns()
func (_BlockManager *BlockManagerTransactor) GiveSorted(opts *bind.TransactOpts, epoch *big.Int, assetId *big.Int, sorted []*big.Int) (*types.Transaction, error) {
	return _BlockManager.contract.Transact(opts, "giveSorted", epoch, assetId, sorted)
}

// GiveSorted is a paid mutator transaction binding the contract method 0x4e6753b7.
//
// Solidity: function giveSorted(uint256 epoch, uint256 assetId, uint256[] sorted) returns()
func (_BlockManager *BlockManagerSession) GiveSorted(epoch *big.Int, assetId *big.Int, sorted []*big.Int) (*types.Transaction, error) {
	return _BlockManager.Contract.GiveSorted(&_BlockManager.TransactOpts, epoch, assetId, sorted)
}

// GiveSorted is a paid mutator transaction binding the contract method 0x4e6753b7.
//
// Solidity: function giveSorted(uint256 epoch, uint256 assetId, uint256[] sorted) returns()
func (_BlockManager *BlockManagerTransactorSession) GiveSorted(epoch *big.Int, assetId *big.Int, sorted []*big.Int) (*types.Transaction, error) {
	return _BlockManager.Contract.GiveSorted(&_BlockManager.TransactOpts, epoch, assetId, sorted)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BlockManager *BlockManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BlockManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BlockManager *BlockManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BlockManager.Contract.GrantRole(&_BlockManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BlockManager *BlockManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BlockManager.Contract.GrantRole(&_BlockManager.TransactOpts, role, account)
}

// Init is a paid mutator transaction binding the contract method 0x06552ff3.
//
// Solidity: function init(address _stakeManagerAddress, address _stateManagerAddress, address _voteManagerAddress, address _jobManagerAddress) returns()
func (_BlockManager *BlockManagerTransactor) Init(opts *bind.TransactOpts, _stakeManagerAddress common.Address, _stateManagerAddress common.Address, _voteManagerAddress common.Address, _jobManagerAddress common.Address) (*types.Transaction, error) {
	return _BlockManager.contract.Transact(opts, "init", _stakeManagerAddress, _stateManagerAddress, _voteManagerAddress, _jobManagerAddress)
}

// Init is a paid mutator transaction binding the contract method 0x06552ff3.
//
// Solidity: function init(address _stakeManagerAddress, address _stateManagerAddress, address _voteManagerAddress, address _jobManagerAddress) returns()
func (_BlockManager *BlockManagerSession) Init(_stakeManagerAddress common.Address, _stateManagerAddress common.Address, _voteManagerAddress common.Address, _jobManagerAddress common.Address) (*types.Transaction, error) {
	return _BlockManager.Contract.Init(&_BlockManager.TransactOpts, _stakeManagerAddress, _stateManagerAddress, _voteManagerAddress, _jobManagerAddress)
}

// Init is a paid mutator transaction binding the contract method 0x06552ff3.
//
// Solidity: function init(address _stakeManagerAddress, address _stateManagerAddress, address _voteManagerAddress, address _jobManagerAddress) returns()
func (_BlockManager *BlockManagerTransactorSession) Init(_stakeManagerAddress common.Address, _stateManagerAddress common.Address, _voteManagerAddress common.Address, _jobManagerAddress common.Address) (*types.Transaction, error) {
	return _BlockManager.Contract.Init(&_BlockManager.TransactOpts, _stakeManagerAddress, _stateManagerAddress, _voteManagerAddress, _jobManagerAddress)
}

// Propose is a paid mutator transaction binding the contract method 0x17d99c04.
//
// Solidity: function propose(uint256 epoch, uint256[] jobIds, uint256[] medians, uint256[] lowerCutoffs, uint256[] higherCutoffs, uint256 iteration, uint256 biggestStakerId) returns()
func (_BlockManager *BlockManagerTransactor) Propose(opts *bind.TransactOpts, epoch *big.Int, jobIds []*big.Int, medians []*big.Int, lowerCutoffs []*big.Int, higherCutoffs []*big.Int, iteration *big.Int, biggestStakerId *big.Int) (*types.Transaction, error) {
	return _BlockManager.contract.Transact(opts, "propose", epoch, jobIds, medians, lowerCutoffs, higherCutoffs, iteration, biggestStakerId)
}

// Propose is a paid mutator transaction binding the contract method 0x17d99c04.
//
// Solidity: function propose(uint256 epoch, uint256[] jobIds, uint256[] medians, uint256[] lowerCutoffs, uint256[] higherCutoffs, uint256 iteration, uint256 biggestStakerId) returns()
func (_BlockManager *BlockManagerSession) Propose(epoch *big.Int, jobIds []*big.Int, medians []*big.Int, lowerCutoffs []*big.Int, higherCutoffs []*big.Int, iteration *big.Int, biggestStakerId *big.Int) (*types.Transaction, error) {
	return _BlockManager.Contract.Propose(&_BlockManager.TransactOpts, epoch, jobIds, medians, lowerCutoffs, higherCutoffs, iteration, biggestStakerId)
}

// Propose is a paid mutator transaction binding the contract method 0x17d99c04.
//
// Solidity: function propose(uint256 epoch, uint256[] jobIds, uint256[] medians, uint256[] lowerCutoffs, uint256[] higherCutoffs, uint256 iteration, uint256 biggestStakerId) returns()
func (_BlockManager *BlockManagerTransactorSession) Propose(epoch *big.Int, jobIds []*big.Int, medians []*big.Int, lowerCutoffs []*big.Int, higherCutoffs []*big.Int, iteration *big.Int, biggestStakerId *big.Int) (*types.Transaction, error) {
	return _BlockManager.Contract.Propose(&_BlockManager.TransactOpts, epoch, jobIds, medians, lowerCutoffs, higherCutoffs, iteration, biggestStakerId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_BlockManager *BlockManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BlockManager.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_BlockManager *BlockManagerSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BlockManager.Contract.RenounceRole(&_BlockManager.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_BlockManager *BlockManagerTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BlockManager.Contract.RenounceRole(&_BlockManager.TransactOpts, role, account)
}

// ResetDispute is a paid mutator transaction binding the contract method 0x5ce8772f.
//
// Solidity: function resetDispute(uint256 epoch) returns()
func (_BlockManager *BlockManagerTransactor) ResetDispute(opts *bind.TransactOpts, epoch *big.Int) (*types.Transaction, error) {
	return _BlockManager.contract.Transact(opts, "resetDispute", epoch)
}

// ResetDispute is a paid mutator transaction binding the contract method 0x5ce8772f.
//
// Solidity: function resetDispute(uint256 epoch) returns()
func (_BlockManager *BlockManagerSession) ResetDispute(epoch *big.Int) (*types.Transaction, error) {
	return _BlockManager.Contract.ResetDispute(&_BlockManager.TransactOpts, epoch)
}

// ResetDispute is a paid mutator transaction binding the contract method 0x5ce8772f.
//
// Solidity: function resetDispute(uint256 epoch) returns()
func (_BlockManager *BlockManagerTransactorSession) ResetDispute(epoch *big.Int) (*types.Transaction, error) {
	return _BlockManager.Contract.ResetDispute(&_BlockManager.TransactOpts, epoch)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BlockManager *BlockManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BlockManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BlockManager *BlockManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BlockManager.Contract.RevokeRole(&_BlockManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BlockManager *BlockManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BlockManager.Contract.RevokeRole(&_BlockManager.TransactOpts, role, account)
}

// BlockManagerBlockConfirmedIterator is returned from FilterBlockConfirmed and is used to iterate over the raw logs and unpacked data for BlockConfirmed events raised by the BlockManager contract.
type BlockManagerBlockConfirmedIterator struct {
	Event *BlockManagerBlockConfirmed // Event containing the contract specifics and raw log

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
func (it *BlockManagerBlockConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockManagerBlockConfirmed)
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
		it.Event = new(BlockManagerBlockConfirmed)
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
func (it *BlockManagerBlockConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockManagerBlockConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockManagerBlockConfirmed represents a BlockConfirmed event raised by the BlockManager contract.
type BlockManagerBlockConfirmed struct {
	Epoch         *big.Int
	StakerId      *big.Int
	Medians       []*big.Int
	LowerCutoffs  []*big.Int
	HigherCutoffs []*big.Int
	JobIds        []*big.Int
	Timestamp     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBlockConfirmed is a free log retrieval operation binding the contract event 0x5ef264ae9101b91bd7117350c5e85901297fa0569e7a3e07eb7ec0a6281529f0.
//
// Solidity: event BlockConfirmed(uint256 epoch, uint256 stakerId, uint256[] medians, uint256[] lowerCutoffs, uint256[] higherCutoffs, uint256[] jobIds, uint256 timestamp)
func (_BlockManager *BlockManagerFilterer) FilterBlockConfirmed(opts *bind.FilterOpts) (*BlockManagerBlockConfirmedIterator, error) {

	logs, sub, err := _BlockManager.contract.FilterLogs(opts, "BlockConfirmed")
	if err != nil {
		return nil, err
	}
	return &BlockManagerBlockConfirmedIterator{contract: _BlockManager.contract, event: "BlockConfirmed", logs: logs, sub: sub}, nil
}

// WatchBlockConfirmed is a free log subscription operation binding the contract event 0x5ef264ae9101b91bd7117350c5e85901297fa0569e7a3e07eb7ec0a6281529f0.
//
// Solidity: event BlockConfirmed(uint256 epoch, uint256 stakerId, uint256[] medians, uint256[] lowerCutoffs, uint256[] higherCutoffs, uint256[] jobIds, uint256 timestamp)
func (_BlockManager *BlockManagerFilterer) WatchBlockConfirmed(opts *bind.WatchOpts, sink chan<- *BlockManagerBlockConfirmed) (event.Subscription, error) {

	logs, sub, err := _BlockManager.contract.WatchLogs(opts, "BlockConfirmed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockManagerBlockConfirmed)
				if err := _BlockManager.contract.UnpackLog(event, "BlockConfirmed", log); err != nil {
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

// ParseBlockConfirmed is a log parse operation binding the contract event 0x5ef264ae9101b91bd7117350c5e85901297fa0569e7a3e07eb7ec0a6281529f0.
//
// Solidity: event BlockConfirmed(uint256 epoch, uint256 stakerId, uint256[] medians, uint256[] lowerCutoffs, uint256[] higherCutoffs, uint256[] jobIds, uint256 timestamp)
func (_BlockManager *BlockManagerFilterer) ParseBlockConfirmed(log types.Log) (*BlockManagerBlockConfirmed, error) {
	event := new(BlockManagerBlockConfirmed)
	if err := _BlockManager.contract.UnpackLog(event, "BlockConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockManagerProposedIterator is returned from FilterProposed and is used to iterate over the raw logs and unpacked data for Proposed events raised by the BlockManager contract.
type BlockManagerProposedIterator struct {
	Event *BlockManagerProposed // Event containing the contract specifics and raw log

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
func (it *BlockManagerProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockManagerProposed)
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
		it.Event = new(BlockManagerProposed)
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
func (it *BlockManagerProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockManagerProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockManagerProposed represents a Proposed event raised by the BlockManager contract.
type BlockManagerProposed struct {
	Epoch           *big.Int
	StakerId        *big.Int
	JobIds          []*big.Int
	Medians         []*big.Int
	LowerCutoffs    []*big.Int
	HigherCutoffs   []*big.Int
	Iteration       *big.Int
	BiggestStakerId *big.Int
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterProposed is a free log retrieval operation binding the contract event 0xee036cc96c88163f353feaa4d497e88baaebeb631f40ad9b8a6d51bb6fad4076.
//
// Solidity: event Proposed(uint256 epoch, uint256 stakerId, uint256[] jobIds, uint256[] medians, uint256[] lowerCutoffs, uint256[] higherCutoffs, uint256 iteration, uint256 biggestStakerId, uint256 timestamp)
func (_BlockManager *BlockManagerFilterer) FilterProposed(opts *bind.FilterOpts) (*BlockManagerProposedIterator, error) {

	logs, sub, err := _BlockManager.contract.FilterLogs(opts, "Proposed")
	if err != nil {
		return nil, err
	}
	return &BlockManagerProposedIterator{contract: _BlockManager.contract, event: "Proposed", logs: logs, sub: sub}, nil
}

// WatchProposed is a free log subscription operation binding the contract event 0xee036cc96c88163f353feaa4d497e88baaebeb631f40ad9b8a6d51bb6fad4076.
//
// Solidity: event Proposed(uint256 epoch, uint256 stakerId, uint256[] jobIds, uint256[] medians, uint256[] lowerCutoffs, uint256[] higherCutoffs, uint256 iteration, uint256 biggestStakerId, uint256 timestamp)
func (_BlockManager *BlockManagerFilterer) WatchProposed(opts *bind.WatchOpts, sink chan<- *BlockManagerProposed) (event.Subscription, error) {

	logs, sub, err := _BlockManager.contract.WatchLogs(opts, "Proposed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockManagerProposed)
				if err := _BlockManager.contract.UnpackLog(event, "Proposed", log); err != nil {
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

// ParseProposed is a log parse operation binding the contract event 0xee036cc96c88163f353feaa4d497e88baaebeb631f40ad9b8a6d51bb6fad4076.
//
// Solidity: event Proposed(uint256 epoch, uint256 stakerId, uint256[] jobIds, uint256[] medians, uint256[] lowerCutoffs, uint256[] higherCutoffs, uint256 iteration, uint256 biggestStakerId, uint256 timestamp)
func (_BlockManager *BlockManagerFilterer) ParseProposed(log types.Log) (*BlockManagerProposed, error) {
	event := new(BlockManagerProposed)
	if err := _BlockManager.contract.UnpackLog(event, "Proposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the BlockManager contract.
type BlockManagerRoleAdminChangedIterator struct {
	Event *BlockManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *BlockManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockManagerRoleAdminChanged)
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
		it.Event = new(BlockManagerRoleAdminChanged)
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
func (it *BlockManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockManagerRoleAdminChanged represents a RoleAdminChanged event raised by the BlockManager contract.
type BlockManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BlockManager *BlockManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BlockManagerRoleAdminChangedIterator, error) {

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

	logs, sub, err := _BlockManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BlockManagerRoleAdminChangedIterator{contract: _BlockManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BlockManager *BlockManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BlockManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _BlockManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockManagerRoleAdminChanged)
				if err := _BlockManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_BlockManager *BlockManagerFilterer) ParseRoleAdminChanged(log types.Log) (*BlockManagerRoleAdminChanged, error) {
	event := new(BlockManagerRoleAdminChanged)
	if err := _BlockManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the BlockManager contract.
type BlockManagerRoleGrantedIterator struct {
	Event *BlockManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *BlockManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockManagerRoleGranted)
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
		it.Event = new(BlockManagerRoleGranted)
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
func (it *BlockManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockManagerRoleGranted represents a RoleGranted event raised by the BlockManager contract.
type BlockManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BlockManager *BlockManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BlockManagerRoleGrantedIterator, error) {

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

	logs, sub, err := _BlockManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BlockManagerRoleGrantedIterator{contract: _BlockManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BlockManager *BlockManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BlockManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BlockManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockManagerRoleGranted)
				if err := _BlockManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_BlockManager *BlockManagerFilterer) ParseRoleGranted(log types.Log) (*BlockManagerRoleGranted, error) {
	event := new(BlockManagerRoleGranted)
	if err := _BlockManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the BlockManager contract.
type BlockManagerRoleRevokedIterator struct {
	Event *BlockManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BlockManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockManagerRoleRevoked)
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
		it.Event = new(BlockManagerRoleRevoked)
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
func (it *BlockManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockManagerRoleRevoked represents a RoleRevoked event raised by the BlockManager contract.
type BlockManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BlockManager *BlockManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BlockManagerRoleRevokedIterator, error) {

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

	logs, sub, err := _BlockManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BlockManagerRoleRevokedIterator{contract: _BlockManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BlockManager *BlockManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BlockManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BlockManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockManagerRoleRevoked)
				if err := _BlockManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_BlockManager *BlockManagerFilterer) ParseRoleRevoked(log types.Log) (*BlockManagerRoleRevoked, error) {
	event := new(BlockManagerRoleRevoked)
	if err := _BlockManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
