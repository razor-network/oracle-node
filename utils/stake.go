package utils

import (
	"context"
	"math/big"
	"razor/core/types"
	"razor/pkg/bindings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func (*UtilsStruct) GetStakeManagerWithOpts(client *ethclient.Client) (*bindings.StakeManager, bind.CallOpts) {
	return UtilsInterface.GetStakeManager(client), UtilsInterface.GetOptions()
}

func (*UtilsStruct) GetStakerId(ctx context.Context, client *ethclient.Client, address string) (uint32, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(ctx, StakeManagerInterface, "GetStakerId", client, common.HexToAddress(address))
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint32), nil
}

func (*UtilsStruct) GetStake(ctx context.Context, client *ethclient.Client, stakerId uint32) (*big.Int, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(ctx, StakeManagerInterface, "GetStaker", client, stakerId)
	if err != nil {
		return nil, err
	}
	staker := returnedValues[0].Interface().(bindings.StructsStaker)
	return staker.Stake, nil
}

func (*UtilsStruct) GetStaker(ctx context.Context, client *ethclient.Client, stakerId uint32) (bindings.StructsStaker, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(ctx, StakeManagerInterface, "GetStaker", client, stakerId)
	if err != nil {
		return bindings.StructsStaker{}, err
	}
	return returnedValues[0].Interface().(bindings.StructsStaker), nil
}

func (*UtilsStruct) GetNumberOfStakers(ctx context.Context, client *ethclient.Client) (uint32, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(ctx, StakeManagerInterface, "GetNumStakers", client)
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint32), nil
}

func (*UtilsStruct) GetLock(ctx context.Context, client *ethclient.Client, address string, stakerId uint32, lockType uint8) (types.Locks, error) {
	staker, err := UtilsInterface.GetStaker(ctx, client, stakerId)
	if err != nil {
		return types.Locks{}, err
	}
	returnedValues, err := InvokeFunctionWithRetryAttempts(ctx, StakeManagerInterface, "Locks", client, common.HexToAddress(address), staker.TokenAddress, lockType)
	if err != nil {
		return types.Locks{}, err
	}
	return returnedValues[0].Interface().(types.Locks), nil
}

func (*UtilsStruct) GetWithdrawInitiationPeriod(ctx context.Context, client *ethclient.Client) (uint16, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(ctx, StakeManagerInterface, "WithdrawInitiationPeriod", client)
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint16), nil
}

func (*UtilsStruct) GetMaxCommission(ctx context.Context, client *ethclient.Client) (uint8, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(ctx, StakeManagerInterface, "MaxCommission", client)
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint8), nil
}

func (*UtilsStruct) GetEpochLimitForUpdateCommission(ctx context.Context, client *ethclient.Client) (uint16, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(ctx, StakeManagerInterface, "EpochLimitForUpdateCommission", client)
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint16), nil
}

func (*UtilsStruct) GetMinSafeRazor(ctx context.Context, client *ethclient.Client) (*big.Int, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(ctx, StakeManagerInterface, "MinSafeRazor", client)
	if err != nil {
		return nil, err
	}
	return returnedValues[0].Interface().(*big.Int), nil
}
