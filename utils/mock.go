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

type PackageUtilsMock struct{}

var SuggestGasPriceWithRetryMock func(*ethclient.Client) (*big.Int, error)

var MultiplyFloatAndBigIntMock func(*big.Int, float64) *big.Int

var parseMock func(io.Reader) (abi.ABI, error)

var PackMock func(abi.ABI, string, ...interface{}) ([]byte, error)

var EstimateGasWithRetryMock func(*ethclient.Client, ethereum.CallMsg) (uint64, error)

var increaseGasLimitValueMock func(*ethclient.Client, uint64, float32, Utils) (uint64, error)

var GetLatestBlockWithRetryMock func(*ethclient.Client) (*types.Header, error)

var GetDefaultPathMock func() (string, error)

var GetPrivateKeyMock func(string, string, string, accounts.AccountInterface) *ecdsa.PrivateKey

var GetPendingNonceAtWithRetryMock func(*ethclient.Client, common.Address) (uint64, error)

var getGasPriceMock func(*ethclient.Client, Types.Configurations, Utils) *big.Int

var getGasLimitMock func(Types.TransactionOptions, *bind.TransactOpts, Utils) (uint64, error)

var NewKeyedTransactorWithChainIDMock func(*ecdsa.PrivateKey, *big.Int) (*bind.TransactOpts, error)

func (utilsMock PackageUtilsMock) SuggestGasPriceWithRetry(client *ethclient.Client) (*big.Int, error) {
	return SuggestGasPriceWithRetryMock(client)
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

func (utilsMock PackageUtilsMock) EstimateGasWithRetry(client *ethclient.Client, message ethereum.CallMsg) (uint64, error) {
	return EstimateGasWithRetryMock(client, message)
}

func (utilsMock PackageUtilsMock) increaseGasLimitValue(client *ethclient.Client, gasLimit uint64, gasLimitMultiplier float32, razorUtils Utils) (uint64, error) {
	return increaseGasLimitValueMock(client, gasLimit, gasLimitMultiplier, razorUtils)
}

func (utilsMock PackageUtilsMock) GetLatestBlockWithRetry(client *ethclient.Client) (*types.Header, error) {
	return GetLatestBlockWithRetryMock(client)
}

func (utilsMock PackageUtilsMock) GetDefaultPath() (string, error) {
	return GetDefaultPathMock()
}

func (utilsMock PackageUtilsMock) GetPrivateKey(address string, password string, keystorePath string, accountUtils accounts.AccountInterface) *ecdsa.PrivateKey {
	return GetPrivateKeyMock(address, password, keystorePath, accountUtils)
}

func (utilsMock PackageUtilsMock) GetPendingNonceAtWithRetry(client *ethclient.Client, accountAddress common.Address) (uint64, error) {
	return GetPendingNonceAtWithRetryMock(client, accountAddress)
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
