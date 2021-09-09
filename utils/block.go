package utils

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"razor/pkg/bindings"
)

func GetNumberOfProposedBlocks(client *ethclient.Client, address string, epoch uint32) (uint8, error) {
	blockManager := GetBlockManager(client)
	callOpts := GetOptions(false, address, "")
	return blockManager.GetNumProposedBlocks(&callOpts, epoch)
}

func GetProposedBlock(client *ethclient.Client, address string, epoch uint32, proposedBlock uint8) (struct {
	Block        bindings.StructsBlock
	BlockMedians []uint32
}, error) {
	blockManager := GetBlockManager(client)
	callOpts := GetOptions(false, address, "")
	return blockManager.GetProposedBlock(&callOpts, epoch, proposedBlock)
}
