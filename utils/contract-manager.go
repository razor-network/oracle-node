package utils

import (
	"razor/core"

	"razor/pkg/bindings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func (*UtilsStruct) GetTokenManager(client *ethclient.Client) *bindings.RAZOR {
	coinContract, err := Options.NewRAZOR(common.HexToAddress(core.RAZORAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return coinContract
}

func (*UtilsStruct) GetStakeManager(client *ethclient.Client) *bindings.StakeManager {
	stakeManagerContract, err := Options.NewStakeManager(common.HexToAddress(core.StakeManagerAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return stakeManagerContract
}

func (*UtilsStruct) GetAssetManager(client *ethclient.Client) *bindings.AssetManager {
	assetManager, err := Options.NewAssetManager(common.HexToAddress(core.AssetManagerAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return assetManager
}

func (*UtilsStruct) GetVoteManager(client *ethclient.Client) *bindings.VoteManager {
	voteManager, err := Options.NewVoteManager(common.HexToAddress(core.VoteManagerAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return voteManager
}

func (*UtilsStruct) GetBlockManager(client *ethclient.Client) *bindings.BlockManager {
	blockManager, err := Options.NewBlockManager(common.HexToAddress(core.BlockManagerAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return blockManager
}

func (*UtilsStruct) GetStakedToken(client *ethclient.Client, tokenAddress common.Address) *bindings.StakedToken {
	stakedTokenContract, err := Options.NewStakedToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	return stakedTokenContract
}
