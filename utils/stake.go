package utils

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"time"
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
			log.Error("Error in fetching staker id: ", stakerErr)
			retryingIn := math.Pow(2, float64(retry))
			log.Debugf("Retrying in %f seconds.....", retryingIn)
			time.Sleep(time.Duration(retryingIn) * time.Second)
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
	return stakeManager.GetStaker(&callOpts, stakerId)
}

func GetNumberOfStakers(client *ethclient.Client, address string) (uint32, error) {
	stakeManager, callOpts := getStakeManagerWithOpts(client, address)
	return stakeManager.GetNumStakers(&callOpts)
}

func GetInfluence(client *ethclient.Client, address string, stakerId uint32) (*big.Int, error) {
	stakeManager, callOpts := getStakeManagerWithOpts(client, address)
	return stakeManager.GetInfluence(&callOpts, stakerId)
}

func GetLock(client *ethclient.Client, address string, stakerId uint32) (types.Locks, error) {
	stakeManager, callOpts := getStakeManagerWithOpts(client, address)
	staker, err := GetStaker(client, address, stakerId)
	if err != nil {
		return types.Locks{}, err
	}
	return stakeManager.Locks(&callOpts, common.HexToAddress(address), staker.TokenAddress)
}
