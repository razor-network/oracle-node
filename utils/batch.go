package utils

import (
	"context"
	"errors"
	"fmt"
	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	RPC "github.com/ethereum/go-ethereum/rpc"
	"razor/core"
	"razor/rpc"
)

//Each batch call may require multiple arguments therefore defining args as [][]interface{}

// BatchCall performs a batch call to the Ethereum client, using the provided contract ABI, address, method name, and arguments.
func (c ClientStruct) BatchCall(rpcParameters rpc.RPCParameters, contractABI *abi.ABI, contractAddress, methodName string, args [][]interface{}) ([][]interface{}, error) {
	calls, err := ClientInterface.CreateBatchCalls(contractABI, contractAddress, methodName, args)
	if err != nil {
		log.Errorf("Error in creating batch calls: %v", err)
		return nil, err
	}

	err = performBatchCallWithRetry(rpcParameters, calls)
	if err != nil {
		log.Errorf("Error in performing batch call: %v", err)
		return nil, err
	}

	results, err := processBatchResults(contractABI, methodName, calls)
	if err != nil {
		log.Errorf("Error in processing batch call result: %v", err)
		return nil, err
	}

	return results, nil
}

// CreateBatchCalls creates a slice of rpc.BatchElem, each representing an Ethereum call, using the provided ABI, contract address, method name, and arguments.
func (c ClientStruct) CreateBatchCalls(contractABI *abi.ABI, contractAddress, methodName string, args [][]interface{}) ([]RPC.BatchElem, error) {
	var calls []RPC.BatchElem

	for _, arg := range args {
		data, err := contractABI.Pack(methodName, arg...)
		if err != nil {
			log.Errorf("Failed to pack data for method %s: %v", methodName, err)
			return nil, err
		}

		calls = append(calls, RPC.BatchElem{
			Method: "eth_call",
			Args: []interface{}{
				map[string]interface{}{
					"to":   contractAddress,
					"data": fmt.Sprintf("0x%x", data),
				},
				"latest",
			},
			Result: new(string),
		})
	}
	return calls, nil
}

func (c ClientStruct) PerformBatchCall(rpcParameters rpc.RPCParameters, calls []RPC.BatchElem) error {
	client, err := rpcParameters.RPCManager.GetBestRPCClient()
	if err != nil {
		return err
	}

	err = client.Client().BatchCallContext(context.Background(), calls)
	if err != nil {
		return err
	}
	return nil
}

// performBatchCallWithRetry performs the batch call to the Ethereum client with retry logic.
func performBatchCallWithRetry(rpcParameters rpc.RPCParameters, calls []RPC.BatchElem) error {
	err := retry.Do(func() error {
		err := ClientInterface.PerformBatchCall(rpcParameters, calls)
		if err != nil {
			log.Errorf("Error in performing batch call, retrying: %v", err)
			return err
		}
		for _, call := range calls {
			if call.Error != nil {
				log.Errorf("Error in call result: %v", call.Error)
				return call.Error
			}
		}
		return nil
	}, retry.Attempts(core.MaxRetries))

	if err != nil {
		log.Errorf("All attempts failed to perform batch call: %v", err)
		return err
	}

	return nil
}

// processBatchResults processes the results of the batch call, unpacking the data using the provided ABI and method name.
func processBatchResults(contractABI *abi.ABI, methodName string, calls []RPC.BatchElem) ([][]interface{}, error) {
	var results [][]interface{}

	for _, call := range calls {
		if call.Error != nil {
			log.Errorf("Error in call result: %v", call.Error)
			return nil, call.Error
		}

		result, ok := call.Result.(*string)
		if !ok {
			log.Error("Failed to type assert call result to *string")
			return nil, errors.New("type asserting of batch call result error")
		}

		if result == nil || *result == "" {
			return nil, errors.New("empty batch call result")
		}

		data := common.FromHex(*result)
		if len(data) == 0 {
			return nil, errors.New("empty hex data")
		}

		unpackedData, err := contractABI.Unpack(methodName, data)
		if err != nil {
			return nil, errors.New("unpacking data error")
		}

		results = append(results, unpackedData)
	}
	return results, nil
}
