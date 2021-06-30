package utils

import (
	"math/big"
	"razor/core"

	"github.com/ethereum/go-ethereum/ethclient"
)

func GetBlockHashes(client *ethclient.Client, address string) ([]byte, error) {
	randomClient := GetRandomClient(client)
	callOpts := GetOptions(false, address, "")
	blockHashes, err := randomClient.BlockHashes(&callOpts, uint8(core.NumberOfBlocks), big.NewInt(core.EpochLength))
	if err != nil {
		return nil, err
	}
	return blockHashes[:], err
}
