package utils

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"razor/pkg/bindings"
)

func GetNumberOfProposedBlocks(client *ethclient.Client, address string, epoch *big.Int) (*big.Int, error) {
	blockManager := GetBlockManager(client)
	callOpts := GetOptions(false, address, "")
	return blockManager.GetNumProposedBlocks(&callOpts, epoch)
}

func GetProposedBlock(client *ethclient.Client, address string, epoch *big.Int, proposedBlock *big.Int) (struct {
	Block         bindings.StructsBlock
	BlockMedians  []*big.Int
	LowerCutoffs  []*big.Int
	HigherCutoffs []*big.Int
}, error) {
	blockManager := GetBlockManager(client)
	callOpts := GetOptions(false, address, "")
	return blockManager.GetProposedBlock(&callOpts, epoch, proposedBlock)
}
