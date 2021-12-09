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

type Utils interface {
	SuggestGasPriceWithRetry(*ethclient.Client, Utils) (*big.Int, error)
	MultiplyFloatAndBigInt(*big.Int, float64) *big.Int
	parse(io.Reader) (abi.ABI, error)
	Pack(abi.ABI, string, ...interface{}) ([]byte, error)
	EstimateGasWithRetry(*ethclient.Client, ethereum.CallMsg, Utils) (uint64, error)
	increaseGasLimitValue(*ethclient.Client, uint64, float32, Utils) (uint64, error)
	GetLatestBlockWithRetry(*ethclient.Client, Utils) (*types.Header, error)
	GetDefaultPath() (string, error)
	GetPrivateKey(string, string, string, accounts.AccountInterface) *ecdsa.PrivateKey
	GetPendingNonceAtWithRetry(*ethclient.Client, common.Address, Utils) (uint64, error)
	getGasPrice(*ethclient.Client, Types.Configurations, Utils) *big.Int
	getGasLimit(Types.TransactionOptions, *bind.TransactOpts, Utils) (uint64, error)
	NewKeyedTransactorWithChainID(*ecdsa.PrivateKey, *big.Int) (*bind.TransactOpts, error)
	PendingNonceAt(*ethclient.Client, context.Context, common.Address) (uint64, error)
	HeaderByNumber(*ethclient.Client, context.Context, *big.Int) (*types.Header, error)
	SuggestGasPrice(*ethclient.Client, context.Context) (*big.Int, error)
	EstimateGas(*ethclient.Client, context.Context, ethereum.CallMsg) (uint64, error)
	FilterLogs(*ethclient.Client, context.Context, ethereum.FilterQuery) ([]types.Log, error)
	BalanceAt(*ethclient.Client, context.Context, common.Address, *big.Int) (*big.Int, error)
	RetryAttempts(uint) retry.Option
}
