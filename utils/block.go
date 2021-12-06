package utils

import (
	"github.com/avast/retry-go"
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
	err = retry.Do(
		func() error {
			numProposedBlocks, err = blockManager.GetNumProposedBlocks(&callOpts, epoch)
			if err != nil {
				log.Error("Error in fetching numProposedBlocks.... Retrying")
				return err
			}
			return nil
		}, retry.Attempts(core.MaxRetries))
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
	err = retry.Do(
		func() error {
			proposedBlock, err = blockManager.GetProposedBlock(&callOpts, epoch, proposedBlockId)
			if err != nil {
				log.Error("Error in fetching proposed block.... Retrying")
				return err
			}
			return nil
		}, retry.Attempts(core.MaxRetries))
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
	err = retry.Do(
		func() error {
			block, err = blockManager.GetBlock(&callOpts, epoch)
			if err != nil {
				log.Error("Error in fetching proposed block.... Retrying")
				return err
			}
			return nil
		}, retry.Attempts(core.MaxRetries))
	if err != nil {
		return 0, err
	}
	return block.Medians[assetId-1], nil
}

func GetMinStakeAmount(client *ethclient.Client, address string) (*big.Int, error) {
	blockManager, callOpts := getBlockManagerWithOpts(client, address)
	var (
		minStake *big.Int
		err      error
	)
	err = retry.Do(
		func() error {
			minStake, err = blockManager.MinStake(&callOpts)
			if err != nil {
				log.Error("Error in fetching minimum stake amount.... Retrying")
				return err
			}
			return nil
		}, retry.Attempts(core.MaxRetries))
	if err != nil {
		return nil, err
	}
	return minStake, nil
}

func GetMaxAltBlocks(client *ethclient.Client, address string) (uint8, error) {
	blockManager, callOpts := getBlockManagerWithOpts(client, address)
	var (
		maxAltBlocks uint8
		err          error
	)
	err = retry.Do(
		func() error {
			maxAltBlocks, err = blockManager.MaxAltBlocks(&callOpts)
			if err != nil {
				log.Error("Error in fetching max alt blocks.... Retrying")
				return err
			}
			return nil
		}, retry.Attempts(core.MaxRetries))
	if err != nil {
		return 0, err
	}
	return maxAltBlocks, nil
}

func GetSortedProposedBlockId(client *ethclient.Client, address string, epoch uint32, index *big.Int) (uint8, error) {
	blockManager, callOpts := getBlockManagerWithOpts(client, address)
	var (
		sortedProposedBlockId uint8
		err                   error
	)
	err = retry.Do(
		func() error {
			sortedProposedBlockId, err = blockManager.SortedProposedBlockIds(&callOpts, epoch, index)
			if err != nil {
				log.Error("Error in fetching max alt blocks.... Retrying")
				return err
			}
			return nil
		}, retry.Attempts(core.MaxRetries))
	if err != nil {
		return 0, err
	}
	return sortedProposedBlockId, nil
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
