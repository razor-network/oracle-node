package utils

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func (*ClientStruct) GetNonceAtWithRetry(client *ethclient.Client, accountAddress common.Address) (uint64, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(ClientInterface, "NonceAt", client, context.Background(), accountAddress)
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint64), nil
}

func (*ClientStruct) GetLatestBlockWithRetry(client *ethclient.Client) (*types.Header, error) {
	var blockNumberArgument *big.Int = nil
	returnedValues, err := InvokeFunctionWithRetryAttempts(ClientInterface, "HeaderByNumber", client, context.Background(), blockNumberArgument)
	if err != nil {
		return nil, err
	}
	return returnedValues[0].Interface().(*types.Header), nil
}

func (*ClientStruct) SuggestGasPriceWithRetry(client *ethclient.Client) (*big.Int, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(ClientInterface, "SuggestGasPrice", client, context.Background())
	if err != nil {
		return nil, err
	}
	return returnedValues[0].Interface().(*big.Int), nil
}

func (*ClientStruct) EstimateGasWithRetry(client *ethclient.Client, message ethereum.CallMsg) (uint64, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(ClientInterface, "EstimateGas", client, context.Background(), message)
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint64), nil
}

func (*ClientStruct) FilterLogsWithRetry(client *ethclient.Client, query ethereum.FilterQuery) ([]types.Log, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(ClientInterface, "FilterLogs", client, context.Background(), query)
	if err != nil {
		return nil, err
	}
	return returnedValues[0].Interface().([]types.Log), nil
}

func (*ClientStruct) BalanceAtWithRetry(client *ethclient.Client, account common.Address) (*big.Int, error) {
	var blockNumberArgument *big.Int = nil
	returnedValues, err := InvokeFunctionWithRetryAttempts(ClientInterface, "BalanceAt", client, context.Background(), account, blockNumberArgument)
	if err != nil {
		return nil, err
	}
	return returnedValues[0].Interface().(*big.Int), nil
}
