//Package utils provides the utils functions
package utils

import (
	"math/big"
	"razor/core"
	"razor/pkg/bindings"

	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

//This function returns the block manager with opts
func (*UtilsStruct) GetBlockManagerWithOpts(client *ethclient.Client) (*bindings.BlockManager, bind.CallOpts) {
	return UtilsInterface.GetBlockManager(client), UtilsInterface.GetOptions()
}

//This function returns the number of proposed blocks
func (*UtilsStruct) GetNumberOfProposedBlocks(client *ethclient.Client, epoch uint32) (uint8, error) {
	var (
		numProposedBlocks uint8
		err               error
	)
	err = retry.Do(
		func() error {
			numProposedBlocks, err = BlockManagerInterface.GetNumProposedBlocks(client, epoch)
			if err != nil {
				log.Error("Error in fetching numProposedBlocks.... Retrying")
				return err
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if err != nil {
		return 0, err
	}
	return numProposedBlocks, nil
}

//This function returns the proposed block
func (*UtilsStruct) GetProposedBlock(client *ethclient.Client, epoch uint32, proposedBlockId uint32) (bindings.StructsBlock, error) {
	var (
		proposedBlock bindings.StructsBlock
		err           error
	)
	err = retry.Do(
		func() error {
			proposedBlock, err = BlockManagerInterface.GetProposedBlock(client, epoch, proposedBlockId)
			if err != nil {
				log.Error("Error in fetching proposed block.... Retrying")
				return err
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if err != nil {
		return bindings.StructsBlock{}, err
	}
	return proposedBlock, nil
}

//This function fetches the previous value
func (*UtilsStruct) FetchPreviousValue(client *ethclient.Client, epoch uint32, assetId uint16) (*big.Int, error) {
	block, err := UtilsInterface.GetBlock(client, epoch)
	if err != nil {
		return big.NewInt(0), err
	}
	return block.Medians[assetId-1], nil
}

//This function returns the block
func (*UtilsStruct) GetBlock(client *ethclient.Client, epoch uint32) (bindings.StructsBlock, error) {
	var (
		block bindings.StructsBlock
		err   error
	)
	err = retry.Do(
		func() error {
			block, err = BlockManagerInterface.GetBlock(client, epoch)
			if err != nil {
				log.Error("Error in fetching proposed block.... Retrying")
				return err
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if err != nil {
		return bindings.StructsBlock{}, err
	}
	return block, nil
}

//This function returns the minimum stake amount
func (*UtilsStruct) GetMinStakeAmount(client *ethclient.Client) (*big.Int, error) {
	var (
		minStake *big.Int
		err      error
	)
	err = retry.Do(
		func() error {
			minStake, err = BlockManagerInterface.MinStake(client)
			if err != nil {
				log.Error("Error in fetching minimum stake amount.... Retrying")
				return err
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if err != nil {
		return nil, err
	}
	return minStake, nil
}

//This function returns the maximum alt blocks
func (*UtilsStruct) GetMaxAltBlocks(client *ethclient.Client) (uint8, error) {
	var (
		maxAltBlocks uint8
		err          error
	)
	err = retry.Do(
		func() error {
			maxAltBlocks, err = BlockManagerInterface.MaxAltBlocks(client)
			if err != nil {
				log.Error("Error in fetching max alt blocks.... Retrying")
				return err
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if err != nil {
		return 0, err
	}
	return maxAltBlocks, nil
}

//This function returns the sorted proposed block Id
func (*UtilsStruct) GetSortedProposedBlockId(client *ethclient.Client, epoch uint32, index *big.Int) (uint32, error) {
	var (
		sortedProposedBlockId uint32
		err                   error
	)
	err = retry.Do(
		func() error {
			sortedProposedBlockId, err = BlockManagerInterface.SortedProposedBlockIds(client, epoch, index)
			if err != nil {
				log.Error("Error in fetching sorted proposed blockId.... Retrying")
				return err
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if err != nil {
		return 0, err
	}
	return sortedProposedBlockId, nil
}

//This function returns the sorted proposed block Ids
func (*UtilsStruct) GetSortedProposedBlockIds(client *ethclient.Client, epoch uint32) ([]uint32, error) {
	numberOfProposedBlocks, err := UtilsInterface.GetNumberOfProposedBlocks(client, epoch)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	var sortedProposedBlockIds []uint32
	for i := 0; i < int(numberOfProposedBlocks); i++ {
		id, err := UtilsInterface.GetSortedProposedBlockId(client, epoch, big.NewInt(int64(i)))
		if err != nil {
			log.Error(err)
			return nil, err
		}
		sortedProposedBlockIds = append(sortedProposedBlockIds, id)
	}
	return sortedProposedBlockIds, nil
}

//This function returns the block index which has to be confirmed
func (*UtilsStruct) GetBlockIndexToBeConfirmed(client *ethclient.Client) (int8, error) {
	var (
		blockIndex int8
		err        error
	)
	err = retry.Do(
		func() error {
			blockIndex, err = BlockManagerInterface.GetBlockIndexToBeConfirmed(client)
			if err != nil {
				log.Error("Error in fetching blockIndexToBeConfirmed....Retrying")
				return err
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if err != nil {
		return 0, err
	}
	return blockIndex, nil
}

func (*UtilsStruct) IsBlockConfirmed(client *ethclient.Client, epoch uint32) (bool, error) {
	var (
		isConfirmed bool
		err         error
	)
	err = retry.Do(
		func() error {
			isConfirmed, err = BlockManagerInterface.IsBlockConfirmed(client, epoch)
			if err != nil {
				log.Error("Error in fetching isBlockConfirmed!")
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if err != nil {
		return false, err
	}
	return isConfirmed, nil
}
