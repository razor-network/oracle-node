package utils

import (
	"errors"
	"math/big"
	"razor/pkg/bindings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

func (*UtilsStruct) GetBlockManagerWithOpts(client *ethclient.Client) (*bindings.BlockManager, bind.CallOpts) {
	return UtilsInterface.GetBlockManager(client), UtilsInterface.GetOptions()
}

func (*UtilsStruct) GetNumberOfProposedBlocks(client *ethclient.Client, epoch uint32) (uint8, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(BlockManagerInterface, "GetNumProposedBlocks", client, epoch)
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint8), nil
}

func (*UtilsStruct) GetProposedBlock(client *ethclient.Client, epoch uint32, proposedBlockId uint32) (bindings.StructsBlock, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(BlockManagerInterface, "GetProposedBlock", client, epoch, proposedBlockId)
	if err != nil {
		return bindings.StructsBlock{}, err
	}
	return returnedValues[0].Interface().(bindings.StructsBlock), nil
}

func (*UtilsStruct) FetchPreviousValue(client *ethclient.Client, epoch uint32, assetId uint16) (*big.Int, error) {
	block, err := UtilsInterface.GetBlock(client, epoch)
	if err != nil {
		return big.NewInt(0), err
	}
	if len(block.Medians) < int(assetId) {
		return big.NewInt(0), errors.New("value not found in previous block")
	}
	return block.Medians[assetId-1], nil
}

func (*UtilsStruct) GetBlock(client *ethclient.Client, epoch uint32) (bindings.StructsBlock, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(BlockManagerInterface, "GetBlock", client, epoch)
	if err != nil {
		return bindings.StructsBlock{}, err
	}
	return returnedValues[0].Interface().(bindings.StructsBlock), nil
}

func (*UtilsStruct) GetMinStakeAmount(client *ethclient.Client) (*big.Int, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(BlockManagerInterface, "MinStake", client)
	if err != nil {
		return nil, err
	}
	return returnedValues[0].Interface().(*big.Int), nil
}

func (*UtilsStruct) GetStateBuffer(client *ethclient.Client) (uint64, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(BlockManagerInterface, "StateBuffer", client)
	if err != nil {
		return 0, err
	}
	stateBufferUint8 := returnedValues[0].Interface().(uint8)
	return uint64(stateBufferUint8), nil
}

func (*UtilsStruct) GetMaxAltBlocks(client *ethclient.Client) (uint8, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(BlockManagerInterface, "MaxAltBlocks", client)
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint8), nil
}

func (*UtilsStruct) GetSortedProposedBlockId(client *ethclient.Client, epoch uint32, index *big.Int) (uint32, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(BlockManagerInterface, "SortedProposedBlockIds", client, epoch, index)
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint32), nil
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
	returnedValues, err := InvokeFunctionWithRetryAttempts(BlockManagerInterface, "GetBlockIndexToBeConfirmed", client)
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(int8), nil
}

func (*UtilsStruct) GetEpochLastProposed(client *ethclient.Client, stakerId uint32) (uint32, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(BlockManagerInterface, "GetEpochLastProposed", client, stakerId)
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint32), nil
}
