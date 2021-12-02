package utils

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"razor/core"
	"razor/pkg/bindings"
)

func getBlockManagerWithOpts(client *ethclient.Client, address string) (*bindings.BlockManager, bind.CallOpts) {
	return GetBlockManager(client), GetOptions()
}

func GetNumberOfProposedBlocks(client *ethclient.Client, address string, epoch uint32) (uint8, error) {
	blockManager := GetBlockManager(client)
	callOpts := GetOptions()
	var (
		numProposedBlocks uint8
		err               error
	)
	for retry := 1; retry <= core.MaxRetries; retry++ {
		numProposedBlocks, err = blockManager.GetNumProposedBlocks(&callOpts, epoch)
		if err != nil {
			Retry(retry, "Error in fetching numProposedBlocks: ", err)
			continue
		}
		break
	}
	if err != nil {
		return 0, err
	}
	return numProposedBlocks, nil
}

func GetProposedBlock(client *ethclient.Client, address string, epoch uint32, proposedBlockId uint8) (bindings.StructsBlock, error) {
	blockManager := GetBlockManager(client)
	callOpts := GetOptions()
	var (
		proposedBlock bindings.StructsBlock
		err           error
	)
	for retry := 1; retry <= core.MaxRetries; retry++ {
		proposedBlock, err = blockManager.GetProposedBlock(&callOpts, epoch, proposedBlockId)
		if err != nil {
			Retry(retry, "Error in fetching proposed block: ", err)
			continue
		}
		break
	}
	if err != nil {
		return bindings.StructsBlock{}, err
	}
	return proposedBlock, nil
}

func FetchPreviousValue(client *ethclient.Client, address string, epoch uint32, assetId uint8) (uint32, error) {
	blockManager := GetBlockManager(client)
	callOpts := GetOptions()
	var (
		block bindings.StructsBlock
		err   error
	)
	for retry := 1; retry <= core.MaxRetries; retry++ {
		block, err = blockManager.GetBlock(&callOpts, epoch)
		if err != nil {
			Retry(retry, "Error in fetching proposed block: ", err)
			continue
		}
		break
	}
	if err != nil {
		return 0, err
	}
	return block.Medians[assetId-1], nil
}

func GetMinStakeAmount(client *ethclient.Client, address string) (*big.Int, error) {
	blockManager, callOpts := getBlockManagerWithOpts(client, address)
	return blockManager.MinStake(&callOpts)
}

func GetMaxAltBlocks(client *ethclient.Client, address string) (uint8, error) {
	blockManager, callOpts := getBlockManagerWithOpts(client, address)
	return blockManager.MaxAltBlocks(&callOpts)
}

func GetSortedProposedBlockId(client *ethclient.Client, address string, epoch uint32, index *big.Int) (uint8, error) {
	blockManager, callOpts := getBlockManagerWithOpts(client, address)
	return blockManager.SortedProposedBlockIds(&callOpts, epoch, index)
}

func GetSortedProposedBlockIds(client *ethclient.Client, address string, epoch uint32) ([]uint8, error) {
	numberOfProposedBlocks, err := GetNumberOfProposedBlocks(client, address, epoch)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	var sortedProposedBlockIds []uint8
	for i := 0; i < int(numberOfProposedBlocks); i++ {
		id, err := GetSortedProposedBlockId(client, address, epoch, big.NewInt(int64(i)))
		if err != nil {
			log.Error(err)
			return nil, err
		}
		sortedProposedBlockIds = append(sortedProposedBlockIds, id)
	}
	return sortedProposedBlockIds, nil
}
