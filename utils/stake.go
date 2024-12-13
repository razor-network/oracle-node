package utils

import (
	"math/big"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/rpc"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func (*UtilsStruct) GetStakeManagerWithOpts(client *ethclient.Client) (*bindings.StakeManager, bind.CallOpts) {
	return UtilsInterface.GetStakeManager(client), UtilsInterface.GetOptions()
}

func (*UtilsStruct) GetStakerId(rpcParameters rpc.RPCParameters, address string) (uint32, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, StakeManagerInterface, "GetStakerId", common.HexToAddress(address))
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint32), nil
}

func (*UtilsStruct) GetStake(rpcParameters rpc.RPCParameters, stakerId uint32) (*big.Int, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, StakeManagerInterface, "GetStaker", stakerId)
	if err != nil {
		return nil, err
	}
	staker := returnedValues[0].Interface().(bindings.StructsStaker)
	return staker.Stake, nil
}

func (*UtilsStruct) GetStaker(rpcParameters rpc.RPCParameters, stakerId uint32) (bindings.StructsStaker, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, StakeManagerInterface, "GetStaker", stakerId)
	if err != nil {
		return bindings.StructsStaker{}, err
	}
	return returnedValues[0].Interface().(bindings.StructsStaker), nil
}

func (*UtilsStruct) GetNumberOfStakers(rpcParameters rpc.RPCParameters) (uint32, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, StakeManagerInterface, "GetNumStakers")
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint32), nil
}

func (*UtilsStruct) GetLock(rpcParameters rpc.RPCParameters, address string, stakerId uint32, lockType uint8) (types.Locks, error) {
	staker, err := UtilsInterface.GetStaker(rpcParameters, stakerId)
	if err != nil {
		return types.Locks{}, err
	}
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, StakeManagerInterface, "Locks", common.HexToAddress(address), staker.TokenAddress, lockType)
	if err != nil {
		return types.Locks{}, err
	}
	return returnedValues[0].Interface().(types.Locks), nil
}

func (*UtilsStruct) GetWithdrawInitiationPeriod(rpcParameters rpc.RPCParameters) (uint16, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, StakeManagerInterface, "WithdrawInitiationPeriod")
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint16), nil
}

func (*UtilsStruct) GetMaxCommission(rpcParameters rpc.RPCParameters) (uint8, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, StakeManagerInterface, "MaxCommission")
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint8), nil
}

func (*UtilsStruct) GetEpochLimitForUpdateCommission(rpcParameters rpc.RPCParameters) (uint16, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, StakeManagerInterface, "EpochLimitForUpdateCommission")
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint16), nil
}

func (*UtilsStruct) GetMinSafeRazor(rpcParameters rpc.RPCParameters) (*big.Int, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, StakeManagerInterface, "MinSafeRazor")
	if err != nil {
		return nil, err
	}
	return returnedValues[0].Interface().(*big.Int), nil
}

func (*UtilsStruct) StakerInfo(rpcParameters rpc.RPCParameters, stakerId uint32) (types.Staker, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, StakeManagerInterface, "StakerInfo", stakerId)
	if err != nil {
		return types.Staker{}, err
	}
	return returnedValues[0].Interface().(types.Staker), nil
}

func (*UtilsStruct) GetMaturity(rpcParameters rpc.RPCParameters, age uint32) (uint16, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, StakeManagerInterface, "GetMaturity", age)
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint16), nil
}

func (*UtilsStruct) GetBountyLock(rpcParameters rpc.RPCParameters, bountyId uint32) (types.BountyLock, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, StakeManagerInterface, "GetBountyLock", bountyId)
	if err != nil {
		return types.BountyLock{}, err
	}
	return returnedValues[0].Interface().(types.BountyLock), nil
}
