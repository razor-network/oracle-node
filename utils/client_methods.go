package utils

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"razor/RPC"
)

func (*ClientStruct) GetNonceAtWithRetry(rpcParameters RPC.RPCParameters, accountAddress common.Address) (uint64, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, ClientInterface, "NonceAt", context.Background(), accountAddress)
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint64), nil
}

func (*ClientStruct) GetLatestBlockWithRetry(rpcParameters RPC.RPCParameters) (*types.Header, error) {
	var blockNumberArgument *big.Int
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, ClientInterface, "HeaderByNumber", context.Background(), blockNumberArgument)
	if err != nil {
		return nil, err
	}
	return returnedValues[0].Interface().(*types.Header), nil
}

func (*ClientStruct) GetBlockByNumberWithRetry(rpcParameters RPC.RPCParameters, blockNumber *big.Int) (*types.Header, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, ClientInterface, "HeaderByNumber", context.Background(), blockNumber)
	if err != nil {
		return nil, err
	}
	return returnedValues[0].Interface().(*types.Header), nil
}

func (*ClientStruct) SuggestGasPriceWithRetry(rpcParameters RPC.RPCParameters) (*big.Int, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, ClientInterface, "SuggestGasPrice", context.Background())
	if err != nil {
		return nil, err
	}
	return returnedValues[0].Interface().(*big.Int), nil
}

func (*ClientStruct) EstimateGasWithRetry(rpcParameters RPC.RPCParameters, message ethereum.CallMsg) (uint64, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, ClientInterface, "EstimateGas", context.Background(), message)
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint64), nil
}

func (*ClientStruct) FilterLogsWithRetry(rpcParameters RPC.RPCParameters, query ethereum.FilterQuery) ([]types.Log, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, ClientInterface, "FilterLogs", context.Background(), query)
	if err != nil {
		return nil, err
	}
	return returnedValues[0].Interface().([]types.Log), nil
}

func (*ClientStruct) BalanceAtWithRetry(rpcParameters RPC.RPCParameters, account common.Address) (*big.Int, error) {
	var blockNumberArgument *big.Int
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, ClientInterface, "BalanceAt", context.Background(), account, blockNumberArgument)
	if err != nil {
		return nil, err
	}
	return returnedValues[0].Interface().(*big.Int), nil
}
