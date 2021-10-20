package utils

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"razor/core"
	"razor/pkg/bindings"
)

func GetNumberOfProposedBlocks(client *ethclient.Client, address string, epoch uint32) (uint8, error) {
	blockManager := GetBlockManager(client)
	callOpts := GetOptions(false, address, "")
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
	callOpts := GetOptions(false, address, "")
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
	callOpts := GetOptions(false, address, "")
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
