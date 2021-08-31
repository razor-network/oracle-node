package utils

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"razor/core/types"
	"razor/pkg/bindings"
)

func getStakeManagerWithOpts(client *ethclient.Client, address string) (*bindings.StakeManager, bind.CallOpts) {
	return GetStakeManager(client), GetOptions(false, address, "")
}

func GetStakerId(client *ethclient.Client, address string) (uint32, error) {
	stakeManager, callOpts := getStakeManagerWithOpts(client, address)
	return stakeManager.GetStakerId(&callOpts, common.HexToAddress(address))
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
	//log.Info("Staker Token Address: ", staker.TokenAddress)
	return stakeManager.Locks(&callOpts, common.HexToAddress(address), staker.TokenAddress)
}
