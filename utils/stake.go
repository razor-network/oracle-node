package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
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
