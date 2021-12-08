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

type RazorUtilsMock struct{}

var SuggestGasPriceWithRetryMock func(*ethclient.Client, RazorUtilsInterface) (*big.Int, error)

var MultiplyFloatAndBigIntMock func(*big.Int, float64) *big.Int

var parseMock func(io.Reader) (abi.ABI, error)

var PackMock func(abi.ABI, string, ...interface{}) ([]byte, error)

var EstimateGasWithRetryMock func(*ethclient.Client, ethereum.CallMsg, RazorUtilsInterface) (uint64, error)

var increaseGasLimitValueMock func(*ethclient.Client, uint64, float32, RazorUtilsInterface) (uint64, error)

var GetLatestBlockWithRetryMock func(*ethclient.Client, RazorUtilsInterface) (*types.Header, error)

var GetDefaultPathMock func() (string, error)

var GetPrivateKeyMock func(string, string, string, accounts.AccountInterface) *ecdsa.PrivateKey

var GetPendingNonceAtWithRetryMock func(*ethclient.Client, common.Address, RazorUtilsInterface) (uint64, error)

var getGasPriceMock func(*ethclient.Client, Types.Configurations, RazorUtilsInterface) *big.Int

var getGasLimitMock func(Types.TransactionOptions, *bind.TransactOpts, RazorUtilsInterface) (uint64, error)

var NewKeyedTransactorWithChainIDMock func(*ecdsa.PrivateKey, *big.Int) (*bind.TransactOpts, error)

var PendingNonceAtMock func(*ethclient.Client, context.Context, common.Address) (uint64, error)

var HeaderByNumberMock func(*ethclient.Client, context.Context, *big.Int) (*types.Header, error)

var SuggestGasPriceMock func(*ethclient.Client, context.Context) (*big.Int, error)

var EstimateGasMock func(*ethclient.Client, context.Context, ethereum.CallMsg) (uint64, error)

var FilterLogsMock func(*ethclient.Client, context.Context, ethereum.FilterQuery) ([]types.Log, error)

var BalanceAtMock func(*ethclient.Client, context.Context, common.Address, *big.Int) (*big.Int, error)

var RetryAttemptsMock func(uint) retry.Option

func (utilsMock RazorUtilsMock) SuggestGasPriceWithRetry(client *ethclient.Client, razorUtils RazorUtilsInterface) (*big.Int, error) {
	return SuggestGasPriceWithRetryMock(client, razorUtils)
}

func (utilsMock RazorUtilsMock) MultiplyFloatAndBigInt(bigIntVal *big.Int, floatingVal float64) *big.Int {
	return MultiplyFloatAndBigIntMock(bigIntVal, floatingVal)
}

func (utilsMock RazorUtilsMock) parse(reader io.Reader) (abi.ABI, error) {
	return parseMock(reader)
}

func (utilsMock RazorUtilsMock) Pack(parsedData abi.ABI, name string, args ...interface{}) ([]byte, error) {
	return PackMock(parsedData, name, args)
}

func (utilsMock RazorUtilsMock) EstimateGasWithRetry(client *ethclient.Client, message ethereum.CallMsg, razorUtils RazorUtilsInterface) (uint64, error) {
	return EstimateGasWithRetryMock(client, message, razorUtils)
}

func (utilsMock RazorUtilsMock) increaseGasLimitValue(client *ethclient.Client, gasLimit uint64, gasLimitMultiplier float32, razorUtils RazorUtilsInterface) (uint64, error) {
	return increaseGasLimitValueMock(client, gasLimit, gasLimitMultiplier, razorUtils)
}

func (utilsMock RazorUtilsMock) GetLatestBlockWithRetry(client *ethclient.Client, razorUtils RazorUtilsInterface) (*types.Header, error) {
	return GetLatestBlockWithRetryMock(client, razorUtils)
}

func (utilsMock RazorUtilsMock) GetDefaultPath() (string, error) {
	return GetDefaultPathMock()
}

func (utilsMock RazorUtilsMock) GetPrivateKey(address string, password string, keystorePath string, accountUtils accounts.AccountInterface) *ecdsa.PrivateKey {
	return GetPrivateKeyMock(address, password, keystorePath, accountUtils)
}

func (utilsMock RazorUtilsMock) GetPendingNonceAtWithRetry(client *ethclient.Client, accountAddress common.Address, razorUtils RazorUtilsInterface) (uint64, error) {
	return GetPendingNonceAtWithRetryMock(client, accountAddress, razorUtils)
}

func (utilsMock RazorUtilsMock) getGasPrice(client *ethclient.Client, config Types.Configurations, razorUtils RazorUtilsInterface) *big.Int {
	return getGasPriceMock(client, config, razorUtils)
}

func (utilsMock RazorUtilsMock) getGasLimit(transactionData Types.TransactionOptions, txnOpts *bind.TransactOpts, razorUtils RazorUtilsInterface) (uint64, error) {
	return getGasLimitMock(transactionData, txnOpts, razorUtils)
}

func (utilsMock RazorUtilsMock) NewKeyedTransactorWithChainID(privateKey *ecdsa.PrivateKey, chainId *big.Int) (*bind.TransactOpts, error) {
	return NewKeyedTransactorWithChainIDMock(privateKey, chainId)
}

func (utilsMock RazorUtilsMock) PendingNonceAt(client *ethclient.Client, ctx context.Context, account common.Address) (uint64, error) {
	return PendingNonceAtMock(client, ctx, account)
}

func (utilsMock RazorUtilsMock) HeaderByNumber(client *ethclient.Client, ctx context.Context, number *big.Int) (*types.Header, error) {
	return HeaderByNumberMock(client, ctx, number)
}

func (utilsMock RazorUtilsMock) SuggestGasPrice(client *ethclient.Client, ctx context.Context) (*big.Int, error) {
	return SuggestGasPriceMock(client, ctx)
}

func (utilsMock RazorUtilsMock) EstimateGas(client *ethclient.Client, ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	return EstimateGasMock(client, ctx, msg)
}

func (utilsMock RazorUtilsMock) FilterLogs(client *ethclient.Client, ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	return FilterLogsMock(client, ctx, q)
}

func (utilsMock RazorUtilsMock) BalanceAt(client *ethclient.Client, ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return BalanceAtMock(client, ctx, account, blockNumber)
}

func (utilsMock RazorUtilsMock) RetryAttempts(numberOfAttempts uint) retry.Option {
	return RetryAttemptsMock(numberOfAttempts)
}
