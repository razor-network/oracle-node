package utils

import (
	"crypto/ecdsa"
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
	SuggestGasPriceWithRetry(*ethclient.Client) (*big.Int, error)
	MultiplyFloatAndBigInt(*big.Int, float64) *big.Int
	parse(io.Reader) (abi.ABI, error)
	Pack(abi.ABI, string, ...interface{}) ([]byte, error)
	EstimateGasWithRetry(*ethclient.Client, ethereum.CallMsg) (uint64, error)
	increaseGasLimitValue(*ethclient.Client, uint64, float32, Utils) (uint64, error)
	GetLatestBlockWithRetry(*ethclient.Client) (*types.Header, error)
	GetDefaultPath() (string, error)
	GetPrivateKey(string, string, string, accounts.AccountInterface) *ecdsa.PrivateKey
	GetPendingNonceAtWithRetry(*ethclient.Client, common.Address) (uint64, error)
	getGasPrice(*ethclient.Client, Types.Configurations, Utils) *big.Int
	getGasLimit(Types.TransactionOptions, *bind.TransactOpts, Utils) (uint64, error)
	NewKeyedTransactorWithChainID(*ecdsa.PrivateKey, *big.Int) (*bind.TransactOpts, error)
}
