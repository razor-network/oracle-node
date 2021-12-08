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

type RazorUtils struct{}

func (u RazorUtils) SuggestGasPriceWithRetry(client *ethclient.Client) (*big.Int, error) {
	return SuggestGasPriceWithRetry(client)
}

func (u RazorUtils) MultiplyFloatAndBigInt(bigIntVal *big.Int, floatingVal float64) *big.Int {
	return MultiplyFloatAndBigInt(bigIntVal, floatingVal)
}

func (u RazorUtils) parse(reader io.Reader) (abi.ABI, error) {
	return abi.JSON(reader)
}

func (u RazorUtils) Pack(parsedData abi.ABI, name string, args ...interface{}) ([]byte, error) {
	return parsedData.Pack(name, args...)
}

func (u RazorUtils) EstimateGasWithRetry(client *ethclient.Client, message ethereum.CallMsg) (uint64, error) {
	return EstimateGasWithRetry(client, message)
}

func (u RazorUtils) increaseGasLimitValue(client *ethclient.Client, gasLimit uint64, gasLimitMultiplier float32, razorUtils RazorUtilsInterface) (uint64, error) {
	return increaseGasLimitValue(client, gasLimit, gasLimitMultiplier, razorUtils)
}

func (u RazorUtils) GetLatestBlockWithRetry(client *ethclient.Client) (*types.Header, error) {
	return GetLatestBlockWithRetry(client)
}

func (u RazorUtils) GetDefaultPath() (string, error) {
	return path.GetDefaultPath()
}

func (u RazorUtils) GetPrivateKey(address string, password string, keystorePath string, accountUtils accounts.AccountInterface) *ecdsa.PrivateKey {
	return accounts.GetPrivateKey(address, password, keystorePath, accountUtils)
}

func (u RazorUtils) GetPendingNonceAtWithRetry(client *ethclient.Client, accountAddress common.Address) (uint64, error) {
	return GetPendingNonceAtWithRetry(client, accountAddress)
}

func (u RazorUtils) getGasPrice(client *ethclient.Client, config Types.Configurations, razorUtils RazorUtilsInterface) *big.Int {
	return getGasPrice(client, config, razorUtils)
}

func (u RazorUtils) getGasLimit(transactionData Types.TransactionOptions, txnOpts *bind.TransactOpts, razorUtils RazorUtilsInterface) (uint64, error) {
	return getGasLimit(transactionData, txnOpts, razorUtils)
}

func (u RazorUtils) NewKeyedTransactorWithChainID(privateKey *ecdsa.PrivateKey, chainId *big.Int) (*bind.TransactOpts, error) {
	return bind.NewKeyedTransactorWithChainID(privateKey, chainId)
}
