//Package utils provides the utils functions
package utils

import (
	"razor/core"

	"razor/pkg/bindings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

//This function returns the token manager
func (*UtilsStruct) GetTokenManager(client *ethclient.Client) *bindings.RAZOR {
	coinContract, err := BindingsInterface.NewRAZOR(common.HexToAddress(core.RAZORAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return coinContract
}

//This function returns the stake manager
func (*UtilsStruct) GetStakeManager(client *ethclient.Client) *bindings.StakeManager {
	stakeManagerContract, err := BindingsInterface.NewStakeManager(common.HexToAddress(core.StakeManagerAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return stakeManagerContract
}

//This function returns the collection manager
func (*UtilsStruct) GetCollectionManager(client *ethclient.Client) *bindings.CollectionManager {
	collectionManager, err := BindingsInterface.NewCollectionManager(common.HexToAddress(core.CollectionManagerAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return collectionManager
}

//This function returns the vote manager
func (*UtilsStruct) GetVoteManager(client *ethclient.Client) *bindings.VoteManager {
	voteManager, err := BindingsInterface.NewVoteManager(common.HexToAddress(core.VoteManagerAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return voteManager
}

//This function returns the block manager
func (*UtilsStruct) GetBlockManager(client *ethclient.Client) *bindings.BlockManager {
	blockManager, err := BindingsInterface.NewBlockManager(common.HexToAddress(core.BlockManagerAddress), client)
	if err != nil {
		log.Fatal(err)
	}
	return blockManager
}

//This function returns the staked token
func (*UtilsStruct) GetStakedToken(client *ethclient.Client, tokenAddress common.Address) *bindings.StakedToken {
	stakedTokenContract, err := BindingsInterface.NewStakedToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	return stakedTokenContract
}
