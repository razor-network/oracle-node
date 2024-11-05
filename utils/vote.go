package utils

import (
	"math/big"
	"razor/core/types"
	"razor/pkg/bindings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

func (*UtilsStruct) GetVoteManagerWithOpts(client *ethclient.Client) (*bindings.VoteManager, bind.CallOpts) {
	return UtilsInterface.GetVoteManager(client), UtilsInterface.GetOptions()
}

func (*UtilsStruct) GetCommitment(rpcParameters types.RPCParameters, address string) (types.Commitment, error) {
	stakerId, err := UtilsInterface.GetStakerId(rpcParameters, address)
	if err != nil {
		return types.Commitment{}, err
	}
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, VoteManagerInterface, "GetCommitment", stakerId)
	if err != nil {
		return types.Commitment{}, err
	}
	commitment := returnedValues[0].Interface().(types.Commitment)
	return commitment, nil
}

func (*UtilsStruct) GetVoteValue(rpcParameters types.RPCParameters, epoch uint32, stakerId uint32, medianIndex uint16) (*big.Int, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, VoteManagerInterface, "GetVoteValue", epoch, stakerId, medianIndex)
	if err != nil {
		return big.NewInt(0), err
	}
	return returnedValues[0].Interface().(*big.Int), nil
}

func (*UtilsStruct) GetInfluenceSnapshot(rpcParameters types.RPCParameters, stakerId uint32, epoch uint32) (*big.Int, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, VoteManagerInterface, "GetInfluenceSnapshot", epoch, stakerId)
	if err != nil {
		return nil, err
	}
	return returnedValues[0].Interface().(*big.Int), nil
}

func (*UtilsStruct) GetStakeSnapshot(rpcParameters types.RPCParameters, stakerId uint32, epoch uint32) (*big.Int, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, VoteManagerInterface, "GetStakeSnapshot", epoch, stakerId)
	if err != nil {
		return nil, err
	}
	return returnedValues[0].Interface().(*big.Int), nil
}

func (*UtilsStruct) GetTotalInfluenceRevealed(rpcParameters types.RPCParameters, epoch uint32, medianIndex uint16) (*big.Int, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, VoteManagerInterface, "GetTotalInfluenceRevealed", epoch, medianIndex)
	if err != nil {
		return nil, err
	}
	return returnedValues[0].Interface().(*big.Int), nil
}

func (*UtilsStruct) GetEpochLastCommitted(rpcParameters types.RPCParameters, stakerId uint32) (uint32, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, VoteManagerInterface, "GetEpochLastCommitted", stakerId)
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint32), nil
}

func (*UtilsStruct) GetEpochLastRevealed(rpcParameters types.RPCParameters, stakerId uint32) (uint32, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, VoteManagerInterface, "GetEpochLastRevealed", stakerId)
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint32), nil
}

func (*UtilsStruct) ToAssign(rpcParameters types.RPCParameters) (uint16, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, VoteManagerInterface, "ToAssign")
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint16), nil
}
