package utils

import (
	"math/big"
	"razor/core"
	"razor/pkg/bindings"

	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

func (*UtilsStruct) GetBlockManagerWithOpts(client *ethclient.Client) (*bindings.BlockManager, bind.CallOpts) {
	return UtilsInterface.GetBlockManager(client), UtilsInterface.GetOptions()
}

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

func (*UtilsStruct) FetchPreviousValue(client *ethclient.Client, epoch uint32, assetId uint16) (*big.Int, error) {
	block, err := UtilsInterface.GetBlock(client, epoch)
	if err != nil {
		return big.NewInt(0), err
	}
	return block.Medians[assetId-1], nil
}

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

func (*UtilsStruct) GetStateBuffer(client *ethclient.Client) (uint64, error) {
	var (
		stateBuffer uint64
		err         error
	)
	err = retry.Do(
		func() error {
			stateBufferUint8, err := BlockManagerInterface.StateBuffer(client)
			stateBuffer = uint64(stateBufferUint8)
			if err != nil {
				log.Error("Error in fetching state buffer.... Retrying")
				return err
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if err != nil {
		return 0, err
	}
	return stateBuffer, nil
}

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

func (*UtilsStruct) GetBlockIndexToBeConfirmed(client *ethclient.Client) (int8, error) {
	var (
		blockIndex int8
		err        error
	)
	err = retry.Do(
		func() error {
			blockIndex, err = BlockManagerInterface.GetBlockIndexToBeConfirmed(client)
			if err != nil {
				log.Error("Error in fetching salt....Retrying")
				return err
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if err != nil {
		return 0, err
	}
	return blockIndex, nil
}

func (*UtilsStruct) GetEpochLastProposed(client *ethclient.Client, stakerId uint32) (uint32, error) {
	var (
		epochLastProposed uint32
		err               error
	)
	err = retry.Do(
		func() error {
			epochLastProposed, err = BlockManagerInterface.GetEpochLastProposed(client, stakerId)
			if err != nil {
				log.Error("Error in fetching epoch last proposed....Retrying")
				return err
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if err != nil {
		return 0, err
	}
	return epochLastProposed, nil
}
