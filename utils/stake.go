package utils

import (
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"

	"github.com/avast/retry-go"

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
			stakerId, stakerErr = StakeManagerInterface.GetStakerId(client, common.HexToAddress(address))
			if stakerErr != nil {
				log.Error("Error in fetching staker id.... Retrying")
				return stakerErr
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
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
		}, RetryInterface.RetryAttempts(core.MaxRetries))
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
			staker, stakerErr = StakeManagerInterface.GetStaker(client, stakerId)
			if stakerErr != nil {
				log.Error("Error in fetching staker id.... Retrying")
				return stakerErr
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
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
			numStakers, stakerErr = StakeManagerInterface.GetNumStakers(client)
			if stakerErr != nil {
				log.Error("Error in fetching number of stakers.... Retrying")
				return stakerErr
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
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
			locks, lockErr = StakeManagerInterface.Locks(client, common.HexToAddress(address), staker.TokenAddress, lockType)
			if lockErr != nil {
				log.Error("Error in fetching locks.... Retrying")
				return lockErr
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if lockErr != nil {
		return types.Locks{}, lockErr
	}
	return locks, nil
}

func (*UtilsStruct) GetWithdrawInitiationPeriod(client *ethclient.Client) (uint16, error) {
	var (
		withdrawReleasePeriod uint16
		err                   error
	)
	err = retry.Do(
		func() error {
			withdrawReleasePeriod, err = StakeManagerInterface.WithdrawInitiationPeriod(client)
			if err != nil {
				log.Error("Error in fetching withdraw release period.... Retrying")
				return err
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
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
		maxCommission, err = StakeManagerInterface.MaxCommission(client)
		if err != nil {
			log.Error("Error in fetching max commission.... Retrying")
			return err
		}
		return nil
	}, RetryInterface.RetryAttempts(core.MaxRetries))
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
		epochLimitForUpdateCommission, err = StakeManagerInterface.EpochLimitForUpdateCommission(client)
		if err != nil {
			log.Error("Error in fetching epoch limit for update commission")
			return err
		}
		return nil
	}, RetryInterface.RetryAttempts(core.MaxRetries))
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

func (*UtilsStruct) GetMinSafeRazor(client *ethclient.Client) (*big.Int, error) {
	var (
		minSafeRazor *big.Int
		err          error
	)
	err = retry.Do(
		func() error {
			minSafeRazor, err = StakeManagerInterface.MinSafeRazor(client)
			if err != nil {
				log.Error("Error in fetching minimum safe razor.... Retrying")
				return err
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if err != nil {
		return nil, err
	}
	return minSafeRazor, nil
}
