package utils

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
)

func getStakeManagerWithOpts(client *ethclient.Client, address string) (*bindings.StakeManager, bind.CallOpts) {
	return GetStakeManager(client), GetOptions(false, address, "")
}

func GetStakerId(client *ethclient.Client, address string) (uint32, error) {
	stakeManager, callOpts := getStakeManagerWithOpts(client, address)
	var (
		stakerId  uint32
		stakerErr error
	)
	for retry := 1; retry <= core.MaxRetries; retry++ {
		stakerId, stakerErr = stakeManager.GetStakerId(&callOpts, common.HexToAddress(address))
		if stakerErr != nil {
			Retry(retry, "Error in fetching staker id: ", stakerErr)
			continue
		}
		break
	}
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

func GetStaker(client *ethclient.Client, address string, stakerId uint32) (bindings.StructsStaker, error) {
	stakeManager, callOpts := getStakeManagerWithOpts(client, address)
	var (
		staker    bindings.StructsStaker
		stakerErr error
	)
	for retry := 1; retry <= core.MaxRetries; retry++ {
		staker, stakerErr = stakeManager.GetStaker(&callOpts, stakerId)
		if stakerErr != nil {
			Retry(retry, "Error in fetching staker: ", stakerErr)
			continue
		}
		break
	}
	if stakerErr != nil {
		return bindings.StructsStaker{}, stakerErr
	}
	return staker, nil
}

func GetNumberOfStakers(client *ethclient.Client, address string) (uint32, error) {
	stakeManager, callOpts := getStakeManagerWithOpts(client, address)
	var (
		numStakers   uint32
		stakerErr error
	)
	for retry := 1; retry <= core.MaxRetries; retry++ {
		numStakers, stakerErr = stakeManager.GetNumStakers(&callOpts)
		if stakerErr != nil {
			Retry(retry, "Error in fetching number of stakers: ", stakerErr)
			continue
		}
		break
	}
	if stakerErr != nil {
		return 0, stakerErr
	}
	return numStakers, nil
}

func GetInfluence(client *ethclient.Client, address string, stakerId uint32) (*big.Int, error) {
	stakeManager, callOpts := getStakeManagerWithOpts(client, address)
	var (
		influence   *big.Int
		influenceErr error
	)
	for retry := 1; retry <= core.MaxRetries; retry++ {
		influence, influenceErr = stakeManager.GetInfluence(&callOpts, stakerId)
		if influenceErr != nil {
			Retry(retry, "Error in fetching influence: ", influenceErr)
			continue
		}
		break
	}
	if influenceErr != nil {
		return big.NewInt(0), influenceErr
	}
	return influence, nil
}

func GetLock(client *ethclient.Client, address string, stakerId uint32) (types.Locks, error) {
	stakeManager, callOpts := getStakeManagerWithOpts(client, address)
	staker, err := GetStaker(client, address, stakerId)
	if err != nil {
		return types.Locks{}, err
	}
	var (
		locks types.Locks
		lockErr error
	)
	for retry := 1; retry <= core.MaxRetries; retry++ {
		locks, lockErr = stakeManager.Locks(&callOpts, common.HexToAddress(address), staker.TokenAddress)
		if lockErr != nil {
			Retry(retry, "Error in fetching locks: ", lockErr)
			continue
		}
		break
	}
	if lockErr != nil {
		return types.Locks{}, lockErr
	}
	return locks, nil
}
