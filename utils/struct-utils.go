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
	"razor/path"
)

type PackageUtils struct{}

func (u PackageUtils) SuggestGasPriceWithRetry(client *ethclient.Client) (*big.Int, error) {
	return SuggestGasPriceWithRetry(client)
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

func (u PackageUtils) EstimateGasWithRetry(client *ethclient.Client, message ethereum.CallMsg) (uint64, error) {
	return EstimateGasWithRetry(client, message)
}

func (u PackageUtils) increaseGasLimitValue(client *ethclient.Client, gasLimit uint64, gasLimitMultiplier float32, razorUtils Utils) (uint64, error) {
	return increaseGasLimitValue(client, gasLimit, gasLimitMultiplier, razorUtils)
}

func (u PackageUtils) GetLatestBlockWithRetry(client *ethclient.Client) (*types.Header, error) {
	return GetLatestBlockWithRetry(client)
}

func (u PackageUtils) GetDefaultPath() (string, error) {
	return path.GetDefaultPath()
}

func (u PackageUtils) GetPrivateKey(address string, password string, keystorePath string, accountUtils accounts.AccountInterface) *ecdsa.PrivateKey {
	return accounts.GetPrivateKey(address, password, keystorePath, accountUtils)
}

func (u PackageUtils) GetPendingNonceAtWithRetry(client *ethclient.Client, accountAddress common.Address) (uint64, error) {
	return GetPendingNonceAtWithRetry(client, accountAddress)
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
