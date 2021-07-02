package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
	"razor/core/types"
	"razor/pkg/bindings"
)

func GetStakerId(client *ethclient.Client, address string) (*big.Int, error) {
	stakeManager := GetStakeManager(client)
	callOpts := GetOptions(false, address, "")
	return stakeManager.GetStakerId(&callOpts, common.HexToAddress(address))
}

func GetStake(client *ethclient.Client, address string, stakerId *big.Int) (*big.Int, error) {
	stakeManager := GetStakeManager(client)
	callOpts := GetOptions(false, address, "")
	stake, err := stakeManager.Stakers(&callOpts, stakerId)
	if err != nil {
		return nil, err
	}
	return stake.Stake, nil
}

func GetStaker(client *ethclient.Client, address string, stakerId *big.Int) (bindings.StructsStaker, error) {
	stakeManager := GetStakeManager(client)
	callOpts := GetOptions(false, address, "")
	return stakeManager.GetStaker(&callOpts, stakerId)
}

func GetNumberOfStakers(client *ethclient.Client, address string) (*big.Int, error) {
	stakeManager := GetStakeManager(client)
	callOpts := GetOptions(false, address, "")
	return stakeManager.GetNumStakers(&callOpts)
}

func GetLock(client *ethclient.Client, address string, stakerId *big.Int) (types.Locks, error) {
	stakeManager := GetStakeManager(client)
	callOpts := GetOptions(false, address, "")
	staker, err := GetStaker(client, address, stakerId)
	if err != nil {
		return types.Locks{}, err
	}
	log.Info("Staker Token Address: ", staker.TokenAddress)
	return stakeManager.Locks(&callOpts, common.HexToAddress(address), staker.TokenAddress)
}