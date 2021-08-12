package utils

import (
	log "github.com/sirupsen/logrus"

	"razor/pkg/bindings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetTokenManager(client *ethclient.Client) *bindings.RAZOR {
	coinContract, err := bindings.NewRAZOR(common.HexToAddress(GetRAZORAddress()), client)
	if err != nil {
		log.Fatal(err)
	}
	return coinContract
}

func GetStakeManager(client *ethclient.Client) *bindings.StakeManager {
	stakeManagerContract, err := bindings.NewStakeManager(common.HexToAddress(GetStakeManagerAddress()), client)
	if err != nil {
		log.Fatal(err)
	}
	return stakeManagerContract
}

func GetParametersManager(client *ethclient.Client) *bindings.Parameters {
	parametersManager, err := bindings.NewParameters(common.HexToAddress(GetParametersAddress()), client)
	if err != nil {
		log.Fatal(err)
	}
	return parametersManager
}

func GetAssetManager(client *ethclient.Client) *bindings.AssetManager {
	assetManager, err := bindings.NewAssetManager(common.HexToAddress(GetAssetManagerAddress()), client)
	if err != nil {
		log.Fatal(err)
	}
	return assetManager
}

func GetVoteManager(client *ethclient.Client) *bindings.VoteManager {
	voteManager, err := bindings.NewVoteManager(common.HexToAddress(GetVoteManagerAddress()), client)
	if err != nil {
		log.Fatal(err)
	}
	return voteManager
}

func GetRandomClient(client *ethclient.Client) *bindings.Random {
	randomClient, err := bindings.NewRandom(common.HexToAddress(GetRandomAddress()), client)
	if err != nil {
		log.Fatal(err)
	}
	return randomClient
}

func GetBlockManager(client *ethclient.Client) *bindings.BlockManager {
	blockManager, err := bindings.NewBlockManager(common.HexToAddress(GetBlockManagerAddress()), client)
	if err != nil {
		log.Fatal(err)
	}
	return blockManager
}
