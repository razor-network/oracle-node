package utils

import (
	"razor/core"

	"razor/pkg/bindings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetTokenManager(client *ethclient.Client) *bindings.RAZOR {
	coinContract, err := bindings.NewRAZOR(common.HexToAddress(core.RAZORAddress), client)
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

func (*UtilsStruct) GetAssetManager(client *ethclient.Client) *bindings.AssetManager {
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

func (*UtilsStruct) GetBlockManager(client *ethclient.Client) *bindings.BlockManager {
	blockManager, err := bindings.NewBlockManager(common.HexToAddress(core.BlockManagerAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return blockManager
}

func GetStakedToken(client *ethclient.Client, tokenAddress common.Address) *bindings.StakedToken {
	stakedTokenContract, err := bindings.NewStakedToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	return stakedTokenContract
}
