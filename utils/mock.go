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

type PackageUtilsMock struct{}

var SuggestGasPriceWithRetryMock func(*ethclient.Client, Utils) (*big.Int, error)

var MultiplyFloatAndBigIntMock func(*big.Int, float64) *big.Int

var parseMock func(io.Reader) (abi.ABI, error)

var PackMock func(abi.ABI, string, ...interface{}) ([]byte, error)

var EstimateGasWithRetryMock func(*ethclient.Client, ethereum.CallMsg, Utils) (uint64, error)

var increaseGasLimitValueMock func(*ethclient.Client, uint64, float32, Utils) (uint64, error)

var GetLatestBlockWithRetryMock func(*ethclient.Client, Utils) (*types.Header, error)

var GetDefaultPathMock func() (string, error)

var GetPrivateKeyMock func(string, string, string, accounts.AccountInterface) *ecdsa.PrivateKey

var GetPendingNonceAtWithRetryMock func(*ethclient.Client, common.Address, Utils) (uint64, error)

var getGasPriceMock func(*ethclient.Client, Types.Configurations, Utils) *big.Int

var getGasLimitMock func(Types.TransactionOptions, *bind.TransactOpts, Utils) (uint64, error)

var NewKeyedTransactorWithChainIDMock func(*ecdsa.PrivateKey, *big.Int) (*bind.TransactOpts, error)

var PendingNonceAtMock func(*ethclient.Client, context.Context, common.Address) (uint64, error)

var HeaderByNumberMock func(*ethclient.Client, context.Context, *big.Int) (*types.Header, error)

var SuggestGasPriceMock func(*ethclient.Client, context.Context) (*big.Int, error)

var EstimateGasMock func(*ethclient.Client, context.Context, ethereum.CallMsg) (uint64, error)

var FilterLogsMock func(*ethclient.Client, context.Context, ethereum.FilterQuery) ([]types.Log, error)

var BalanceAtMock func(*ethclient.Client, context.Context, common.Address, *big.Int) (*big.Int, error)

var RetryAttemptsMock func(uint) retry.Option

func (utilsMock PackageUtilsMock) SuggestGasPriceWithRetry(client *ethclient.Client, razorUtils Utils) (*big.Int, error) {
	return SuggestGasPriceWithRetryMock(client, razorUtils)
}

func (utilsMock PackageUtilsMock) MultiplyFloatAndBigInt(bigIntVal *big.Int, floatingVal float64) *big.Int {
	return MultiplyFloatAndBigIntMock(bigIntVal, floatingVal)
}

func (utilsMock PackageUtilsMock) parse(reader io.Reader) (abi.ABI, error) {
	return parseMock(reader)
}

func (utilsMock PackageUtilsMock) Pack(parsedData abi.ABI, name string, args ...interface{}) ([]byte, error) {
	return PackMock(parsedData, name, args)
}

func (utilsMock PackageUtilsMock) EstimateGasWithRetry(client *ethclient.Client, message ethereum.CallMsg, razorUtils Utils) (uint64, error) {
	return EstimateGasWithRetryMock(client, message, razorUtils)
}

func (utilsMock PackageUtilsMock) increaseGasLimitValue(client *ethclient.Client, gasLimit uint64, gasLimitMultiplier float32, razorUtils Utils) (uint64, error) {
	return increaseGasLimitValueMock(client, gasLimit, gasLimitMultiplier, razorUtils)
}

func (utilsMock PackageUtilsMock) GetLatestBlockWithRetry(client *ethclient.Client, razorUtils Utils) (*types.Header, error) {
	return GetLatestBlockWithRetryMock(client, razorUtils)
}

func (utilsMock PackageUtilsMock) GetDefaultPath() (string, error) {
	return GetDefaultPathMock()
}

func (utilsMock PackageUtilsMock) GetPrivateKey(address string, password string, keystorePath string, accountUtils accounts.AccountInterface) *ecdsa.PrivateKey {
	return GetPrivateKeyMock(address, password, keystorePath, accountUtils)
}

func (utilsMock PackageUtilsMock) GetPendingNonceAtWithRetry(client *ethclient.Client, accountAddress common.Address, razorUtils Utils) (uint64, error) {
	return GetPendingNonceAtWithRetryMock(client, accountAddress, razorUtils)
}

func (utilsMock PackageUtilsMock) getGasPrice(client *ethclient.Client, config Types.Configurations, razorUtils Utils) *big.Int {
	return getGasPriceMock(client, config, razorUtils)
}

func (utilsMock PackageUtilsMock) getGasLimit(transactionData Types.TransactionOptions, txnOpts *bind.TransactOpts, razorUtils Utils) (uint64, error) {
	return getGasLimitMock(transactionData, txnOpts, razorUtils)
}

func (utilsMock PackageUtilsMock) NewKeyedTransactorWithChainID(privateKey *ecdsa.PrivateKey, chainId *big.Int) (*bind.TransactOpts, error) {
	return NewKeyedTransactorWithChainIDMock(privateKey, chainId)
}

func (utilsMock PackageUtilsMock) PendingNonceAt(client *ethclient.Client, ctx context.Context, account common.Address) (uint64, error) {
	return PendingNonceAtMock(client, ctx, account)
}

func (utilsMock PackageUtilsMock) HeaderByNumber(client *ethclient.Client, ctx context.Context, number *big.Int) (*types.Header, error) {
	return HeaderByNumberMock(client, ctx, number)
}

func (utilsMock PackageUtilsMock) SuggestGasPrice(client *ethclient.Client, ctx context.Context) (*big.Int, error) {
	return SuggestGasPriceMock(client, ctx)
}

func (utilsMock PackageUtilsMock) EstimateGas(client *ethclient.Client, ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	return EstimateGasMock(client, ctx, msg)
}

func (utilsMock PackageUtilsMock) FilterLogs(client *ethclient.Client, ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	return FilterLogsMock(client, ctx, q)
}

func (utilsMock PackageUtilsMock) BalanceAt(client *ethclient.Client, ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return BalanceAtMock(client, ctx, account, blockNumber)
}

func (utilsMock PackageUtilsMock) RetryAttempts(numberOfAttempts uint) retry.Option {
	return RetryAttemptsMock(numberOfAttempts)
}
