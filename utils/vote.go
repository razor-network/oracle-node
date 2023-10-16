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

func (*UtilsStruct) GetCommitment(client *ethclient.Client, address string) (types.Commitment, error) {
	stakerId, err := UtilsInterface.GetStakerId(client, address)
	if err != nil {
		return types.Commitment{}, err
	}
	returnedValues, err := InvokeFunctionWithRetryAttempts(VoteManagerInterface, "GetCommitment", client, stakerId)
	if err != nil {
		return types.Commitment{}, err
	}
	commitment := returnedValues[0].Interface().(types.Commitment)
	return commitment, nil
}

func (*UtilsStruct) GetVoteValue(client *ethclient.Client, epoch uint32, stakerId uint32, medianIndex uint16) (*big.Int, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(VoteManagerInterface, "GetVoteValue", client, epoch, stakerId, medianIndex)
	if err != nil {
		return big.NewInt(0), err
	}
	return returnedValues[0].Interface().(*big.Int), nil
}

func (*UtilsStruct) GetInfluenceSnapshot(client *ethclient.Client, stakerId uint32, epoch uint32) (*big.Int, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(VoteManagerInterface, "GetInfluenceSnapshot", client, epoch, stakerId)
	if err != nil {
		return nil, err
	}
	return returnedValues[0].Interface().(*big.Int), nil
}

func (*UtilsStruct) GetStakeSnapshot(client *ethclient.Client, stakerId uint32, epoch uint32) (*big.Int, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(VoteManagerInterface, "GetStakeSnapshot", client, epoch, stakerId)
	if err != nil {
		return nil, err
	}
	return returnedValues[0].Interface().(*big.Int), nil
}

func (*UtilsStruct) GetTotalInfluenceRevealed(client *ethclient.Client, epoch uint32, medianIndex uint16) (*big.Int, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(VoteManagerInterface, "GetTotalInfluenceRevealed", client, epoch, medianIndex)
	if err != nil {
		return nil, err
	}
	return returnedValues[0].Interface().(*big.Int), nil
}

func (*UtilsStruct) GetEpochLastCommitted(client *ethclient.Client, stakerId uint32) (uint32, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(VoteManagerInterface, "GetEpochLastCommitted", client, stakerId)
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint32), nil
}

func (*UtilsStruct) GetEpochLastRevealed(client *ethclient.Client, stakerId uint32) (uint32, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(VoteManagerInterface, "GetEpochLastRevealed", client, stakerId)
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint32), nil
}

func (*UtilsStruct) ToAssign(client *ethclient.Client) (uint16, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(VoteManagerInterface, "ToAssign", client)
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint16), nil
}
