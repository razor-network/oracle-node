package utils

import (
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

func (*UtilsStruct) GetStakerId(client *ethclient.Client, address string) (uint32, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(StakeManagerInterface, "GetStakerId", client, common.HexToAddress(address))
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint32), nil
}

func (*UtilsStruct) GetStake(client *ethclient.Client, stakerId uint32) (*big.Int, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(StakeManagerInterface, "GetStaker", client, stakerId)
	if err != nil {
		return nil, err
	}
	staker := returnedValues[0].Interface().(bindings.StructsStaker)
	return staker.Stake, nil
}

func (*UtilsStruct) GetStaker(client *ethclient.Client, stakerId uint32) (bindings.StructsStaker, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(StakeManagerInterface, "GetStaker", client, stakerId)
	if err != nil {
		return bindings.StructsStaker{}, err
	}
	return returnedValues[0].Interface().(bindings.StructsStaker), nil
}

func (*UtilsStruct) GetNumberOfStakers(client *ethclient.Client) (uint32, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(StakeManagerInterface, "GetNumStakers", client)
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint32), nil
}

func (*UtilsStruct) GetLock(client *ethclient.Client, address string, stakerId uint32, lockType uint8) (types.Locks, error) {
	staker, err := UtilsInterface.GetStaker(client, stakerId)
	if err != nil {
		return types.Locks{}, err
	}
	returnedValues, err := InvokeFunctionWithRetryAttempts(StakeManagerInterface, "Locks", client, common.HexToAddress(address), staker.TokenAddress, lockType)
	if err != nil {
		return types.Locks{}, err
	}
	return returnedValues[0].Interface().(types.Locks), nil
}

func (*UtilsStruct) GetWithdrawInitiationPeriod(client *ethclient.Client) (uint16, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(StakeManagerInterface, "WithdrawInitiationPeriod", client)
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint16), nil
}

func (*UtilsStruct) GetMaxCommission(client *ethclient.Client) (uint8, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(StakeManagerInterface, "MaxCommission", client)
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint8), nil
}

func (*UtilsStruct) GetEpochLimitForUpdateCommission(client *ethclient.Client) (uint16, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(StakeManagerInterface, "EpochLimitForUpdateCommission", client)
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint16), nil
}

func (*UtilsStruct) GetMinSafeRazor(client *ethclient.Client) (*big.Int, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(StakeManagerInterface, "MinSafeRazor", client)
	if err != nil {
		return nil, err
	}
	return returnedValues[0].Interface().(*big.Int), nil
}
