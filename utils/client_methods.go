package utils

import (
	"context"
	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"razor/core"
)

func (*ClientStruct) GetNonceAtWithRetry(client *ethclient.Client, accountAddress common.Address) (uint64, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(ClientInterface, "NonceAt", client, context.Background(), accountAddress)
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint64), nil
}

func (*ClientStruct) GetLatestBlockWithRetry(client *ethclient.Client) (*types.Header, error) {
	var blockNumberArgument *big.Int
	returnedValues, err := InvokeFunctionWithRetryAttempts(ClientInterface, "HeaderByNumber", client, context.Background(), blockNumberArgument)
	if err != nil {
		return nil, err
	}
	return returnedValues[0].Interface().(*types.Header), nil
}

func (*ClientStruct) SuggestGasPriceWithRetry(client *ethclient.Client) (*big.Int, error) {
	var (
		gasPrice *big.Int
		err      error
	)
	err = retry.Do(
		func() error {
			gasPrice, err = ClientInterface.SuggestGasPrice(client, context.Background())
			if err != nil {
				log.Error("Error in fetching gas price.... Retrying")
				return err
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if err != nil {
		return nil, err
	}
	return gasPrice, nil
}

func (*ClientStruct) EstimateGasWithRetry(client *ethclient.Client, message ethereum.CallMsg) (uint64, error) {
	var (
		gasLimit uint64
		err      error
	)
	err = retry.Do(
		func() error {
			gasLimit, err = ClientInterface.EstimateGas(client, context.Background(), message)
			if err != nil {
				log.Error("Error in estimating gas limit.... Retrying")
				return err
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if err != nil {
		return 0, err
	}
	return gasLimit, nil
}

func (*ClientStruct) FilterLogsWithRetry(client *ethclient.Client, query ethereum.FilterQuery) ([]types.Log, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(ClientInterface, "FilterLogs", client, context.Background(), query)
	if err != nil {
		return nil, err
	}
	return returnedValues[0].Interface().([]types.Log), nil
}

func (*ClientStruct) BalanceAtWithRetry(client *ethclient.Client, account common.Address) (*big.Int, error) {
	var blockNumberArgument *big.Int
	returnedValues, err := InvokeFunctionWithRetryAttempts(ClientInterface, "BalanceAt", client, context.Background(), account, blockNumberArgument)
	if err != nil {
		return nil, err
	}
	return returnedValues[0].Interface().(*big.Int), nil
}
