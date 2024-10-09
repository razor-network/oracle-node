package utils

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func (*ClientStruct) GetNonceAtWithRetry(ctx context.Context, client *ethclient.Client, accountAddress common.Address) (uint64, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(ctx, ClientInterface, "NonceAt", client, context.Background(), accountAddress)
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint64), nil
}

func (*ClientStruct) GetLatestBlockWithRetry(ctx context.Context, client *ethclient.Client) (*types.Header, error) {
	var blockNumberArgument *big.Int
	returnedValues, err := InvokeFunctionWithRetryAttempts(ctx, ClientInterface, "HeaderByNumber", client, context.Background(), blockNumberArgument)
	if err != nil {
		return nil, err
	}
	return returnedValues[0].Interface().(*types.Header), nil
}

func (*ClientStruct) SuggestGasPriceWithRetry(ctx context.Context, client *ethclient.Client) (*big.Int, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(ctx, ClientInterface, "SuggestGasPrice", client, context.Background())
	if err != nil {
		return nil, err
	}
	return returnedValues[0].Interface().(*big.Int), nil
}

func (*ClientStruct) EstimateGasWithRetry(ctx context.Context, client *ethclient.Client, message ethereum.CallMsg) (uint64, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(ctx, ClientInterface, "EstimateGas", client, context.Background(), message)
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint64), nil
}

func (*ClientStruct) FilterLogsWithRetry(ctx context.Context, client *ethclient.Client, query ethereum.FilterQuery) ([]types.Log, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(ctx, ClientInterface, "FilterLogs", client, context.Background(), query)
	if err != nil {
		return nil, err
	}
	return returnedValues[0].Interface().([]types.Log), nil
}

func (*ClientStruct) BalanceAtWithRetry(ctx context.Context, client *ethclient.Client, account common.Address) (*big.Int, error) {
	var blockNumberArgument *big.Int
	returnedValues, err := InvokeFunctionWithRetryAttempts(ctx, ClientInterface, "BalanceAt", client, context.Background(), account, blockNumberArgument)
	if err != nil {
		return nil, err
	}
	return returnedValues[0].Interface().(*big.Int), nil
}
