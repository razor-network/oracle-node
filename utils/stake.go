package utils

import (
	"github.com/avast/retry-go"
	"math/big"
	"razor/core"
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
	var (
		stakerId  uint32
		stakerErr error
	)
	stakerErr = retry.Do(
		func() error {
			stakerId, stakerErr = Options.GetStakerId(client, common.HexToAddress(address))
			if stakerErr != nil {
				log.Error("Error in fetching staker id.... Retrying")
				return stakerErr
			}
			return nil
		}, Options.RetryAttempts(core.MaxRetries))
	if stakerErr != nil {
		return 0, stakerErr
	}
	return stakerId, nil
}

func (*UtilsStruct) GetStake(client *ethclient.Client, stakerId uint32) (*big.Int, error) {
	var (
		staker    bindings.StructsStaker
		stakerErr error
	)
	stakerErr = retry.Do(
		func() error {
			staker, stakerErr = UtilsInterface.GetStaker(client, stakerId)
			if stakerErr != nil {
				log.Error("Error in fetching staker ... Retrying")
				return stakerErr
			}
			return nil
		}, Options.RetryAttempts(core.MaxRetries))
	if stakerErr != nil {
		return nil, stakerErr
	}
	return staker.Stake, nil
}

func (*UtilsStruct) GetStaker(client *ethclient.Client, stakerId uint32) (bindings.StructsStaker, error) {
	var (
		staker    bindings.StructsStaker
		stakerErr error
	)
	stakerErr = retry.Do(
		func() error {
			staker, stakerErr = Options.GetStaker(client, stakerId)
			if stakerErr != nil {
				log.Error("Error in fetching staker id.... Retrying")
				return stakerErr
			}
			return nil
		}, Options.RetryAttempts(core.MaxRetries))
	if stakerErr != nil {
		return bindings.StructsStaker{}, stakerErr
	}
	return staker, nil
}

func (*UtilsStruct) GetNumberOfStakers(client *ethclient.Client) (uint32, error) {
	var (
		numStakers uint32
		stakerErr  error
	)
	stakerErr = retry.Do(
		func() error {
			numStakers, stakerErr = Options.GetNumStakers(client)
			if stakerErr != nil {
				log.Error("Error in fetching number of stakers.... Retrying")
				return stakerErr
			}
			return nil
		}, Options.RetryAttempts(core.MaxRetries))
	if stakerErr != nil {
		return 0, stakerErr
	}
	return numStakers, nil
}

func (*UtilsStruct) GetLock(client *ethclient.Client, address string, stakerId uint32, lockType uint8) (types.Locks, error) {
	staker, err := UtilsInterface.GetStaker(client, stakerId)
	if err != nil {
		return types.Locks{}, err
	}
	var (
		locks   types.Locks
		lockErr error
	)
	lockErr = retry.Do(
		func() error {
			locks, lockErr = Options.Locks(client, common.HexToAddress(address), staker.TokenAddress, lockType)
			if lockErr != nil {
				log.Error("Error in fetching locks.... Retrying")
				return lockErr
			}
			return nil
		}, Options.RetryAttempts(core.MaxRetries))
	if lockErr != nil {
		return types.Locks{}, lockErr
	}
	return locks, nil
}

func (*UtilsStruct) GetWithdrawInitiationPeriod(client *ethclient.Client) (uint8, error) {
	var (
		withdrawReleasePeriod uint8
		err                   error
	)
	err = retry.Do(
		func() error {
			withdrawReleasePeriod, err = Options.WithdrawInitiationPeriod(client)
			if err != nil {
				log.Error("Error in fetching withdraw release period.... Retrying")
				return err
			}
			return nil
		}, Options.RetryAttempts(core.MaxRetries))
	if err != nil {
		return 0, err
	}
	return withdrawReleasePeriod, nil
}

func (*UtilsStruct) GetMaxCommission(client *ethclient.Client) (uint8, error) {
	var (
		maxCommission uint8
		err           error
	)
	err = retry.Do(func() error {
		maxCommission, err = Options.MaxCommission(client)
		if err != nil {
			log.Error("Error in fetching max commission.... Retrying")
			return err
		}
		return nil
	}, Options.RetryAttempts(core.MaxRetries))
	if err != nil {
		return 0, err
	}
	return maxCommission, nil
}

func (*UtilsStruct) GetEpochLimitForUpdateCommission(client *ethclient.Client) (uint16, error) {
	var (
		epochLimitForUpdateCommission uint16
		err                           error
	)
	err = retry.Do(func() error {
		epochLimitForUpdateCommission, err = Options.EpochLimitForUpdateCommission(client)
		if err != nil {
			log.Error("Error in fetching epoch limit for update commission")
			return err
		}
		return nil
	}, Options.RetryAttempts(core.MaxRetries))
	if err != nil {
		return 0, err
	}
	return epochLimitForUpdateCommission, nil
}
