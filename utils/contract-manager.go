package utils

import (
	"razor/core"

	log "github.com/sirupsen/logrus"

	"razor/pkg/bindings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetTokenManager(client *ethclient.Client) *bindings.SchellingCoin {
	coinContract, err := bindings.NewSchellingCoin(common.HexToAddress(core.SchellingCoinAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return coinContract
}

func GetStakeManager(client *ethclient.Client) *bindings.StakeManager {
	stakeManagerContract, err := bindings.NewStakeManager(common.HexToAddress(core.StakeManagerAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return stakeManagerContract
}

func GetParametersManager(client *ethclient.Client) *bindings.Parameters {
	parametersManager, err := bindings.NewParameters(common.HexToAddress(core.ParametersAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return parametersManager
}

func GetAssetManager(client *ethclient.Client) *bindings.AssetManager {
	assetManager, err := bindings.NewAssetManager(common.HexToAddress(core.AssetManagerAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return assetManager
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
