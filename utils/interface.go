package utils

import (
	"context"
	"crypto/ecdsa"
	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"io"
	"math/big"
	"razor/accounts"
	"razor/core/types"
)

//go:generate mockery --name OptionUtils --output ./mocks/ --case=underscore
//go:generate mockery --name Utils --output ./mocks/ --case=underscore

var Options OptionUtils
var UtilsInterface Utils

type OptionUtils interface {
	Parse(io.Reader) (abi.ABI, error)
	Pack(abi.ABI, string, ...interface{}) ([]byte, error)
	GetDefaultPath() (string, error)
	GetPrivateKey(address string, password string, keystorePath string, accountUtils accounts.AccountInterface) *ecdsa.PrivateKey
	NewKeyedTransactorWithChainID(key *ecdsa.PrivateKey, chainID *big.Int) (*bind.TransactOpts, error)
	RetryAttempts(uint) retry.Option
	PendingNonceAt(*ethclient.Client, context.Context, common.Address) (uint64, error)
	HeaderByNumber(*ethclient.Client, context.Context, *big.Int) (*Types.Header, error)
	SuggestGasPrice(*ethclient.Client, context.Context) (*big.Int, error)
	EstimateGas(*ethclient.Client, context.Context, ethereum.CallMsg) (uint64, error)
	FilterLogs(*ethclient.Client, context.Context, ethereum.FilterQuery) ([]Types.Log, error)
	BalanceAt(*ethclient.Client, context.Context, common.Address, *big.Int) (*big.Int, error)
}

type Utils interface {
	SuggestGasPriceWithRetry(*ethclient.Client) (*big.Int, error)
	MultiplyFloatAndBigInt(*big.Int, float64) *big.Int
	GetPendingNonceAtWithRetry(*ethclient.Client, common.Address) (uint64, error)
	GetGasPrice(*ethclient.Client, types.Configurations) *big.Int
	GetTxnOpts(types.TransactionOptions) *bind.TransactOpts
	GetGasLimit(transactionData types.TransactionOptions, txnOpts *bind.TransactOpts) (uint64, error)
	EstimateGasWithRetry(client *ethclient.Client, message ethereum.CallMsg) (uint64, error)
	IncreaseGasLimitValue(client *ethclient.Client, gasLimit uint64, gasLimitMultiplier float32) (uint64, error)
	GetLatestBlockWithRetry(client *ethclient.Client) (*Types.Header, error)
	FilterLogsWithRetry(client *ethclient.Client, query ethereum.FilterQuery) ([]Types.Log, error)
	BalanceAtWithRetry(client *ethclient.Client, account common.Address) (*big.Int, error)
}

type OptionsStruct struct{}
type ABIUtilsStruct struct{}
type PathStruct struct{}
type AccountStruct struct{}
type BindStruct struct{}

type UtilsStruct struct{}

type OptionsPackageStruct struct {
	Options        OptionUtils
	UtilsInterface Utils
}
