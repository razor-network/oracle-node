package cmd

import (
	log "github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"razor/pkg/bindings"
)

func getCoinContract(client *ethclient.Client) *bindings.SchellingCoin {
	coinContract, err := bindings.NewSchellingCoin(common.HexToAddress("0x3a3dC05fB85e44b97A358850F117d3B5df7D1d48"), client)
	if err != nil {
		log.Fatal(err)
	}
	return coinContract
}

func getStateManager(client *ethclient.Client) *bindings.StateManager {
	stateManagerContract, err := bindings.NewStateManager(common.HexToAddress("0x7568FC1928Ac51cE0bd4E417650b0593139842C8"), client)
	if err != nil {
		log.Fatal(err)
	}
	return stateManagerContract
}

func getStakeManager(client *ethclient.Client) *bindings.StakeManager {
	stakeManagerContract, err := bindings.NewStakeManager(common.HexToAddress("0x1C7Ccf3054bA60bA8Ec1fecC7E4E722b59bDD90b"), client)
	if err != nil {
		log.Fatal(err)
	}
	return stakeManagerContract
}
