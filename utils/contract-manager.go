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

func GetConstantsManager(client *ethclient.Client) *bindings.Constants {
	constantsManager, err := bindings.NewConstants(common.HexToAddress(core.ConstantsAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return constantsManager
}

func GetJobManager(client *ethclient.Client) *bindings.JobManager {
	jobManager, err := bindings.NewJobManager(common.HexToAddress(core.JobManagerAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return jobManager
}

func GetVoteManager(client *ethclient.Client) *bindings.VoteManager {
	voteManager, err := bindings.NewVoteManager(common.HexToAddress(core.VoteManagerAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return voteManager
}

func GetRandomClient(client *ethclient.Client) *bindings.Random {
	randomClient, err := bindings.NewRandom(common.HexToAddress(core.RandomClientAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return randomClient
}

func GetBlockManager(client *ethclient.Client) *bindings.BlockManager {
	blockManager, err := bindings.NewBlockManager(common.HexToAddress(core.BlockManagerAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return blockManager
}