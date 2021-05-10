package utils

import (
	log "github.com/sirupsen/logrus"
	"razor/core"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"razor/pkg/bindings"
)

func GetCoinContract(client *ethclient.Client) *bindings.SchellingCoin {
	coinContract, err := bindings.NewSchellingCoin(common.HexToAddress(core.SchellingCoinAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return coinContract
}

func GetStateManager(client *ethclient.Client) *bindings.StateManager {
	stateManagerContract, err := bindings.NewStateManager(common.HexToAddress(core.StateManagerAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return stateManagerContract
}

func GetStakeManager(client *ethclient.Client) *bindings.StakeManager {
	stakeManagerContract, err := bindings.NewStakeManager(common.HexToAddress(core.StakeManagerAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return stakeManagerContract
}
