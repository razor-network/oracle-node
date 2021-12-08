package utils

import (
	"context"
	"crypto/ecdsa"
	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"io"
	"math/big"
	"razor/accounts"
	Types "razor/core/types"
)

type RazorUtilsInterface interface {
	SuggestGasPriceWithRetry(*ethclient.Client, RazorUtilsInterface) (*big.Int, error)
	MultiplyFloatAndBigInt(*big.Int, float64) *big.Int
	parse(io.Reader) (abi.ABI, error)
	Pack(abi.ABI, string, ...interface{}) ([]byte, error)
	EstimateGasWithRetry(*ethclient.Client, ethereum.CallMsg, RazorUtilsInterface) (uint64, error)
	increaseGasLimitValue(*ethclient.Client, uint64, float32, RazorUtilsInterface) (uint64, error)
	GetLatestBlockWithRetry(*ethclient.Client, RazorUtilsInterface) (*types.Header, error)
	GetDefaultPath() (string, error)
	GetPrivateKey(string, string, string, accounts.AccountInterface) *ecdsa.PrivateKey
	GetPendingNonceAtWithRetry(*ethclient.Client, common.Address, RazorUtilsInterface) (uint64, error)
	getGasPrice(*ethclient.Client, Types.Configurations, RazorUtilsInterface) *big.Int
	getGasLimit(Types.TransactionOptions, *bind.TransactOpts, RazorUtilsInterface) (uint64, error)
	NewKeyedTransactorWithChainID(*ecdsa.PrivateKey, *big.Int) (*bind.TransactOpts, error)
	PendingNonceAt(client *ethclient.Client, ctx context.Context, account common.Address) (uint64, error)
	HeaderByNumber(client *ethclient.Client, ctx context.Context, number *big.Int) (*types.Header, error)
	SuggestGasPrice(client *ethclient.Client, ctx context.Context) (*big.Int, error)
	EstimateGas(client *ethclient.Client, ctx context.Context, msg ethereum.CallMsg) (uint64, error)
	FilterLogs(client *ethclient.Client, ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error)
	BalanceAt(client *ethclient.Client, ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error)
	RetryAttempts(uint) retry.Option
}
