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
	callOpts := UtilsInterface.GetOptions()
	var (
		stakerId  uint32
		stakerErr error
	)
	stakerErr = retry.Do(
		func() error {
			stakerId, stakerErr = StakeManagerInterface.GetStakerId(client, &callOpts, common.HexToAddress(address))
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
	callOpts := UtilsInterface.GetOptions()
	var (
		staker    bindings.StructsStaker
		stakerErr error
	)
	stakerErr = retry.Do(
		func() error {
			staker, stakerErr = StakeManagerInterface.GetStaker(client, &callOpts, stakerId)
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
	callOpts := UtilsInterface.GetOptions()
	var (
		numStakers uint32
		stakerErr  error
	)
	stakerErr = retry.Do(
		func() error {
			numStakers, stakerErr = StakeManagerInterface.GetNumStakers(client, &callOpts)
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

func (*UtilsStruct) GetLock(client *ethclient.Client, address string, stakerId uint32) (types.Locks, error) {
	callOpts := UtilsInterface.GetOptions()
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
			locks, lockErr = StakeManagerInterface.Locks(client, &callOpts, common.HexToAddress(address), staker.TokenAddress)
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

func (*UtilsStruct) GetWithdrawReleasePeriod(client *ethclient.Client) (uint8, error) {
	callOpts := UtilsInterface.GetOptions()
	var (
		withdrawReleasePeriod uint8
		err                   error
	)
	err = retry.Do(
		func() error {
			withdrawReleasePeriod, err = StakeManagerInterface.WithdrawReleasePeriod(client, &callOpts)
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
	callOpts := UtilsInterface.GetOptions()
	var (
		maxCommission uint8
		err           error
	)
	err = retry.Do(func() error {
		maxCommission, err = StakeManagerInterface.MaxCommission(client, &callOpts)
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
	callOpts := UtilsInterface.GetOptions()
	var (
		epochLimitForUpdateCommission uint16
		err                           error
	)
	err = retry.Do(func() error {
		epochLimitForUpdateCommission, err = StakeManagerInterface.EpochLimitForUpdateCommission(client, &callOpts)
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

func (*UtilsStruct) GetStakerSRZRBalance(client *ethclient.Client, staker bindings.StructsStaker) (*big.Int, error) {
	stakedToken := UtilsInterface.GetStakedToken(client, staker.TokenAddress)
	callOpts := UtilsInterface.GetOptions()

	sRZRBalance, err := StakedTokenInterface.BalanceOf(stakedToken, &callOpts, staker.Address)
	if err != nil {
		log.Error("Error in getting sRZRBalance: ", err)
		return nil, err
	}
	return sRZRBalance, nil
}
