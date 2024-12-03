package utils

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	Types "razor/core/types"
	"razor/pkg/bindings"
	"razor/rpc"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

func (*UtilsStruct) GetBlockManagerWithOpts(client *ethclient.Client) (*bindings.BlockManager, bind.CallOpts) {
	return UtilsInterface.GetBlockManager(client), UtilsInterface.GetOptions()
}

func (*UtilsStruct) GetNumberOfProposedBlocks(rpcParameters rpc.RPCParameters, epoch uint32) (uint8, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, BlockManagerInterface, "GetNumProposedBlocks", epoch)
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint8), nil
}

func (*UtilsStruct) GetProposedBlock(rpcParameters rpc.RPCParameters, epoch uint32, proposedBlockId uint32) (bindings.StructsBlock, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, BlockManagerInterface, "GetProposedBlock", epoch, proposedBlockId)
	if err != nil {
		return bindings.StructsBlock{}, err
	}
	return returnedValues[0].Interface().(bindings.StructsBlock), nil
}

func (*UtilsStruct) FetchPreviousValue(rpcParameters rpc.RPCParameters, epoch uint32, assetId uint16) (*big.Int, error) {
	block, err := UtilsInterface.GetBlock(rpcParameters, epoch)
	if err != nil {
		return big.NewInt(0), err
	}
	if len(block.Medians) < int(assetId) {
		return big.NewInt(0), errors.New("value not found in previous block")
	}
	return block.Medians[assetId-1], nil
}

func (*UtilsStruct) GetBlock(rpcParameters rpc.RPCParameters, epoch uint32) (bindings.StructsBlock, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, BlockManagerInterface, "GetBlock", epoch)
	if err != nil {
		return bindings.StructsBlock{}, err
	}
	return returnedValues[0].Interface().(bindings.StructsBlock), nil
}

func (*UtilsStruct) GetMinStakeAmount(rpcParameters rpc.RPCParameters) (*big.Int, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, BlockManagerInterface, "MinStake")
	if err != nil {
		return nil, err
	}
	return returnedValues[0].Interface().(*big.Int), nil
}

func (*UtilsStruct) GetStateBuffer(rpcParameters rpc.RPCParameters) (uint64, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, BlockManagerInterface, "StateBuffer")
	if err != nil {
		return 0, err
	}
	stateBufferUint8 := returnedValues[0].Interface().(uint8)
	return uint64(stateBufferUint8), nil
}

func (*UtilsStruct) GetMaxAltBlocks(rpcParameters rpc.RPCParameters) (uint8, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, BlockManagerInterface, "MaxAltBlocks")
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint8), nil
}

func (*UtilsStruct) GetSortedProposedBlockId(rpcParameters rpc.RPCParameters, epoch uint32, index *big.Int) (uint32, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, BlockManagerInterface, "SortedProposedBlockIds", epoch, index)
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint32), nil
}

func (*UtilsStruct) GetSortedProposedBlockIds(rpcParameters rpc.RPCParameters, epoch uint32) ([]uint32, error) {
	numberOfProposedBlocks, err := UtilsInterface.GetNumberOfProposedBlocks(rpcParameters, epoch)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	var sortedProposedBlockIds []uint32
	for i := 0; i < int(numberOfProposedBlocks); i++ {
		id, err := UtilsInterface.GetSortedProposedBlockId(rpcParameters, epoch, big.NewInt(int64(i)))
		if err != nil {
			log.Error(err)
			return nil, err
		}
		sortedProposedBlockIds = append(sortedProposedBlockIds, id)
	}
	return sortedProposedBlockIds, nil
}

func (*UtilsStruct) GetBlockIndexToBeConfirmed(rpcParameters rpc.RPCParameters) (int8, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, BlockManagerInterface, "GetBlockIndexToBeConfirmed")
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(int8), nil
}

func (*UtilsStruct) GetEpochLastProposed(rpcParameters rpc.RPCParameters, stakerId uint32) (uint32, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, BlockManagerInterface, "GetEpochLastProposed", stakerId)
	if err != nil {
		return 0, err
	}
	return returnedValues[0].Interface().(uint32), nil
}

func (*UtilsStruct) GetConfirmedBlocks(rpcParameters rpc.RPCParameters, epoch uint32) (Types.ConfirmedBlock, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, BlockManagerInterface, "GetConfirmedBlocks", epoch)
	if err != nil {
		return Types.ConfirmedBlock{}, err
	}
	return returnedValues[0].Interface().(Types.ConfirmedBlock), nil
}

func (*UtilsStruct) Disputes(rpcParameters rpc.RPCParameters, epoch uint32, address common.Address) (Types.DisputesStruct, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(rpcParameters, BlockManagerInterface, "Disputes", epoch, address)
	if err != nil {
		return Types.DisputesStruct{}, err
	}
	return returnedValues[0].Interface().(Types.DisputesStruct), nil
}
