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

// RandomABI is the input ABI used to generate the binding from.
const RandomABI = "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"numBlocks\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"epochLength\",\"type\":\"uint256\"}],\"name\":\"blockHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"numBlocks\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"seed\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"epochLength\",\"type\":\"uint256\"}],\"name\":\"prng\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"numBlocks\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"seed\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"epochLength\",\"type\":\"uint256\"}],\"name\":\"prngHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Random is an auto generated Go binding around an Ethereum contract.
type Random struct {
	RandomCaller     // Read-only binding to the contract
	RandomTransactor // Write-only binding to the contract
	RandomFilterer   // Log filterer for contract events
}

// RandomCaller is an auto generated read-only Go binding around an Ethereum contract.
type RandomCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RandomTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RandomFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RandomSession struct {
	Contract     *Random           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RandomCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RandomCallerSession struct {
	Contract *RandomCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RandomTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RandomTransactorSession struct {
	Contract     *RandomTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RandomRaw is an auto generated low-level Go binding around an Ethereum contract.
type RandomRaw struct {
	Contract *Random // Generic contract binding to access the raw methods on
}

// RandomCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RandomCallerRaw struct {
	Contract *RandomCaller // Generic read-only contract binding to access the raw methods on
}

// RandomTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RandomTransactorRaw struct {
	Contract *RandomTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRandom creates a new instance of Random, bound to a specific deployed contract.
func NewRandom(address common.Address, backend bind.ContractBackend) (*Random, error) {
	contract, err := bindRandom(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Random{RandomCaller: RandomCaller{contract: contract}, RandomTransactor: RandomTransactor{contract: contract}, RandomFilterer: RandomFilterer{contract: contract}}, nil
}

// NewRandomCaller creates a new read-only instance of Random, bound to a specific deployed contract.
func NewRandomCaller(address common.Address, caller bind.ContractCaller) (*RandomCaller, error) {
	contract, err := bindRandom(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RandomCaller{contract: contract}, nil
}

// NewRandomTransactor creates a new write-only instance of Random, bound to a specific deployed contract.
func NewRandomTransactor(address common.Address, transactor bind.ContractTransactor) (*RandomTransactor, error) {
	contract, err := bindRandom(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RandomTransactor{contract: contract}, nil
}

// NewRandomFilterer creates a new log filterer instance of Random, bound to a specific deployed contract.
func NewRandomFilterer(address common.Address, filterer bind.ContractFilterer) (*RandomFilterer, error) {
	contract, err := bindRandom(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RandomFilterer{contract: contract}, nil
}

// bindRandom binds a generic wrapper to an already deployed contract.
func bindRandom(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RandomABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Random *RandomRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Random.Contract.RandomCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Random *RandomRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Random.Contract.RandomTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Random *RandomRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Random.Contract.RandomTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Random *RandomCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Random.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Random *RandomTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Random.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Random *RandomTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Random.Contract.contract.Transact(opts, method, params...)
}

// BlockHashes is a free data retrieval call binding the contract method 0x7cf6c43a.
//
// Solidity: function blockHashes(uint8 numBlocks, uint256 epochLength) view returns(bytes32)
func (_Random *RandomCaller) BlockHashes(opts *bind.CallOpts, numBlocks uint8, epochLength *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Random.contract.Call(opts, &out, "blockHashes", numBlocks, epochLength)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BlockHashes is a free data retrieval call binding the contract method 0x7cf6c43a.
//
// Solidity: function blockHashes(uint8 numBlocks, uint256 epochLength) view returns(bytes32)
func (_Random *RandomSession) BlockHashes(numBlocks uint8, epochLength *big.Int) ([32]byte, error) {
	return _Random.Contract.BlockHashes(&_Random.CallOpts, numBlocks, epochLength)
}

// BlockHashes is a free data retrieval call binding the contract method 0x7cf6c43a.
//
// Solidity: function blockHashes(uint8 numBlocks, uint256 epochLength) view returns(bytes32)
func (_Random *RandomCallerSession) BlockHashes(numBlocks uint8, epochLength *big.Int) ([32]byte, error) {
	return _Random.Contract.BlockHashes(&_Random.CallOpts, numBlocks, epochLength)
}

// Prng is a free data retrieval call binding the contract method 0x320e985c.
//
// Solidity: function prng(uint8 numBlocks, uint256 max, bytes32 seed, uint256 epochLength) view returns(uint256)
func (_Random *RandomCaller) Prng(opts *bind.CallOpts, numBlocks uint8, max *big.Int, seed [32]byte, epochLength *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Random.contract.Call(opts, &out, "prng", numBlocks, max, seed, epochLength)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Prng is a free data retrieval call binding the contract method 0x320e985c.
//
// Solidity: function prng(uint8 numBlocks, uint256 max, bytes32 seed, uint256 epochLength) view returns(uint256)
func (_Random *RandomSession) Prng(numBlocks uint8, max *big.Int, seed [32]byte, epochLength *big.Int) (*big.Int, error) {
	return _Random.Contract.Prng(&_Random.CallOpts, numBlocks, max, seed, epochLength)
}

// Prng is a free data retrieval call binding the contract method 0x320e985c.
//
// Solidity: function prng(uint8 numBlocks, uint256 max, bytes32 seed, uint256 epochLength) view returns(uint256)
func (_Random *RandomCallerSession) Prng(numBlocks uint8, max *big.Int, seed [32]byte, epochLength *big.Int) (*big.Int, error) {
	return _Random.Contract.Prng(&_Random.CallOpts, numBlocks, max, seed, epochLength)
}

// PrngHash is a free data retrieval call binding the contract method 0xa874132b.
//
// Solidity: function prngHash(uint8 numBlocks, bytes32 seed, uint256 epochLength) view returns(bytes32)
func (_Random *RandomCaller) PrngHash(opts *bind.CallOpts, numBlocks uint8, seed [32]byte, epochLength *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Random.contract.Call(opts, &out, "prngHash", numBlocks, seed, epochLength)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PrngHash is a free data retrieval call binding the contract method 0xa874132b.
//
// Solidity: function prngHash(uint8 numBlocks, bytes32 seed, uint256 epochLength) view returns(bytes32)
func (_Random *RandomSession) PrngHash(numBlocks uint8, seed [32]byte, epochLength *big.Int) ([32]byte, error) {
	return _Random.Contract.PrngHash(&_Random.CallOpts, numBlocks, seed, epochLength)
}

// PrngHash is a free data retrieval call binding the contract method 0xa874132b.
//
// Solidity: function prngHash(uint8 numBlocks, bytes32 seed, uint256 epochLength) view returns(bytes32)
func (_Random *RandomCallerSession) PrngHash(numBlocks uint8, seed [32]byte, epochLength *big.Int) ([32]byte, error) {
	return _Random.Contract.PrngHash(&_Random.CallOpts, numBlocks, seed, epochLength)
}
