package utils

import (
	log "github.com/sirupsen/logrus"

	"razor/pkg/bindings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetTokenManager(client *ethclient.Client) *bindings.RAZOR {
	RAZORAddress := GetRAZORAddress()
	coinContract, err := bindings.NewRAZOR(common.HexToAddress(RAZORAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return coinContract
}

func GetStakeManager(client *ethclient.Client) *bindings.StakeManager {
	stakeManagerAddress := GetStakeManagerAddress()
	stakeManagerContract, err := bindings.NewStakeManager(common.HexToAddress(stakeManagerAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return stakeManagerContract
}

func GetParametersManager(client *ethclient.Client) *bindings.Parameters {
	parametersAddress := GetParametersAddress()
	parametersManager, err := bindings.NewParameters(common.HexToAddress(parametersAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return parametersManager
}

func GetAssetManager(client *ethclient.Client) *bindings.AssetManager {
	assetManagerAddress := GetAssetManagerAddress()
	assetManager, err := bindings.NewAssetManager(common.HexToAddress(assetManagerAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return assetManager
}

func GetVoteManager(client *ethclient.Client) *bindings.VoteManager {
	voteManagerAddress := GetVoteManagerAddress()
	voteManager, err := bindings.NewVoteManager(common.HexToAddress(voteManagerAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return voteManager
}

func GetRandomClient(client *ethclient.Client) *bindings.Random {
	randomClientAddress := GetRandomAddress()
	randomClient, err := bindings.NewRandom(common.HexToAddress(randomClientAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return randomClient
}

func GetBlockManager(client *ethclient.Client) *bindings.BlockManager {
	blockManagerAddress := GetBlockManagerAddress()
	blockManager, err := bindings.NewBlockManager(common.HexToAddress(blockManagerAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return blockManager
}