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
	"razor/path"
)

type PackageUtils struct{}

func (u PackageUtils) SuggestGasPriceWithRetry(client *ethclient.Client, razorUtils Utils) (*big.Int, error) {
	return SuggestGasPriceWithRetry(client, razorUtils)
}

func (u PackageUtils) MultiplyFloatAndBigInt(bigIntVal *big.Int, floatingVal float64) *big.Int {
	return MultiplyFloatAndBigInt(bigIntVal, floatingVal)
}

func (u PackageUtils) parse(reader io.Reader) (abi.ABI, error) {
	return abi.JSON(reader)
}

func (u PackageUtils) Pack(parsedData abi.ABI, name string, args ...interface{}) ([]byte, error) {
	return parsedData.Pack(name, args...)
}

func (u PackageUtils) EstimateGasWithRetry(client *ethclient.Client, message ethereum.CallMsg, razorUtils Utils) (uint64, error) {
	return EstimateGasWithRetry(client, message, razorUtils)
}

func (u PackageUtils) increaseGasLimitValue(client *ethclient.Client, gasLimit uint64, gasLimitMultiplier float32, razorUtils Utils) (uint64, error) {
	return increaseGasLimitValue(client, gasLimit, gasLimitMultiplier, razorUtils)
}

func (u PackageUtils) GetLatestBlockWithRetry(client *ethclient.Client, razorUtils Utils) (*types.Header, error) {
	return GetLatestBlockWithRetry(client, razorUtils)
}

func (u PackageUtils) GetDefaultPath() (string, error) {
	return path.GetDefaultPath()
}

func (u PackageUtils) GetPrivateKey(address string, password string, keystorePath string, accountUtils accounts.AccountInterface) *ecdsa.PrivateKey {
	return accounts.GetPrivateKey(address, password, keystorePath, accountUtils)
}

func (u PackageUtils) GetPendingNonceAtWithRetry(client *ethclient.Client, accountAddress common.Address, razorUtils Utils) (uint64, error) {
	return GetPendingNonceAtWithRetry(client, accountAddress, razorUtils)
}

func (u PackageUtils) getGasPrice(client *ethclient.Client, config Types.Configurations, razorUtils Utils) *big.Int {
	return getGasPrice(client, config, razorUtils)
}

func (u PackageUtils) getGasLimit(transactionData Types.TransactionOptions, txnOpts *bind.TransactOpts, razorUtils Utils) (uint64, error) {
	return getGasLimit(transactionData, txnOpts, razorUtils)
}

func (u PackageUtils) NewKeyedTransactorWithChainID(privateKey *ecdsa.PrivateKey, chainId *big.Int) (*bind.TransactOpts, error) {
	return bind.NewKeyedTransactorWithChainID(privateKey, chainId)
}

func (u PackageUtils) PendingNonceAt(client *ethclient.Client, ctx context.Context, account common.Address) (uint64, error) {
	return client.PendingNonceAt(ctx, account)
}

func (u PackageUtils) HeaderByNumber(client *ethclient.Client, ctx context.Context, number *big.Int) (*types.Header, error) {
	return client.HeaderByNumber(ctx, number)
}

func (u PackageUtils) SuggestGasPrice(client *ethclient.Client, ctx context.Context) (*big.Int, error) {
	return client.SuggestGasPrice(ctx)
}

func (u PackageUtils) EstimateGas(client *ethclient.Client, ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	return client.EstimateGas(ctx, msg)
}

func (u PackageUtils) FilterLogs(client *ethclient.Client, ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	return client.FilterLogs(ctx, q)
}

func (u PackageUtils) BalanceAt(client *ethclient.Client, ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return client.BalanceAt(ctx, account, blockNumber)
}

func (u PackageUtils) RetryAttempts(numberOfAttempts uint) retry.Option {
	return retry.Attempts(numberOfAttempts)
}
