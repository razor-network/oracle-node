package utils

import (
	"github.com/avast/retry-go"
	"math/big"
	"razor/core/types"
	"razor/pkg/bindings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getStakeManagerWithOpts(client *ethclient.Client) (*bindings.StakeManager, bind.CallOpts) {
	return GetStakeManager(client), UtilsInterface.GetOptions()
}

func GetStakerId(client *ethclient.Client, address string) (uint32, error) {
	stakeManager, callOpts := getStakeManagerWithOpts(client)
	var (
		stakerId  uint32
		stakerErr error
	)
	stakerErr = retry.Do(
		func() error {
			stakerId, stakerErr = stakeManager.GetStakerId(&callOpts, common.HexToAddress(address))
			if stakerErr != nil {
				log.Error("Error in fetching staker id.... Retrying")
				return stakerErr
			}
			return nil
		},
	)
	if stakerErr != nil {
		return 0, stakerErr
	}
	return stakerId, nil
}

func GetStake(client *ethclient.Client, address string, stakerId uint32) (*big.Int, error) {
	stake, err := GetStaker(client, address, stakerId)
	if err != nil {
		return nil, err
	}
	return stake.Stake, nil
}

//TODO: Remove address
func GetStaker(client *ethclient.Client, address string, stakerId uint32) (bindings.StructsStaker, error) {
	stakeManager, callOpts := getStakeManagerWithOpts(client)
	var (
		staker    bindings.StructsStaker
		stakerErr error
	)
	stakerErr = retry.Do(
		func() error {
			staker, stakerErr = stakeManager.GetStaker(&callOpts, stakerId)
			if stakerErr != nil {
				log.Error("Error in fetching staker id.... Retrying")
				return stakerErr
			}
			return nil
		},
	)
	if stakerErr != nil {
		return bindings.StructsStaker{}, stakerErr
	}
	return staker, nil
}

//TODO: Remove address
func GetNumberOfStakers(client *ethclient.Client, address string) (uint32, error) {
	stakeManager, callOpts := getStakeManagerWithOpts(client)
	var (
		numStakers uint32
		stakerErr  error
	)
	stakerErr = retry.Do(
		func() error {
			numStakers, stakerErr = stakeManager.GetNumStakers(&callOpts)
			if stakerErr != nil {
				log.Error("Error in fetching number of stakers.... Retrying")
				return stakerErr
			}
			return nil
		},
	)
	if stakerErr != nil {
		return 0, stakerErr
	}
	return numStakers, nil
}

func GetLock(client *ethclient.Client, address string, stakerId uint32) (types.Locks, error) {
	stakeManager, callOpts := getStakeManagerWithOpts(client)
	staker, err := GetStaker(client, address, stakerId)
	if err != nil {
		return types.Locks{}, err
	}
	var (
		locks   types.Locks
		lockErr error
	)
	lockErr = retry.Do(
		func() error {
			locks, lockErr = stakeManager.Locks(&callOpts, common.HexToAddress(address), staker.TokenAddress)
			if lockErr != nil {
				log.Error("Error in fetching locks.... Retrying")
				return lockErr
			}
			return nil
		},
	)
	if lockErr != nil {
		return types.Locks{}, lockErr
	}
	return locks, nil
}

func GetWithdrawReleasePeriod(client *ethclient.Client, address string) (uint8, error) {
	stakeManager, callOpts := getStakeManagerWithOpts(client)
	return stakeManager.WithdrawReleasePeriod(&callOpts)
}

func GetMaxCommission(client *ethclient.Client) (uint8, error) {
	stakeManager, callOpts := getStakeManagerWithOpts(client)
	var (
		maxCommission uint8
		err           error
	)
	err = retry.Do(func() error {
		maxCommission, err = stakeManager.MaxCommission(&callOpts)
		if err != nil {
			log.Error("Error in fetching max commission.... Retrying")
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return maxCommission, nil
}

func GetEpochLimitForUpdateCommission(client *ethclient.Client) (uint16, error) {
	stakeManager, callOpts := getStakeManagerWithOpts(client)
	var (
		epochLimitForUpdateCommission uint16
		err                           error
	)
	err = retry.Do(func() error {
		epochLimitForUpdateCommission, err = stakeManager.EpochLimitForUpdateCommission(&callOpts)
		if err != nil {
			log.Error("Error in fetching epoch limit for update commission")
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return epochLimitForUpdateCommission, nil
}
