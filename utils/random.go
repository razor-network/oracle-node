package utils

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"razor/core"
)

func GetBlockHashes(client *ethclient.Client, address string) ([]byte, error) {
	randomClient := GetRandomClient(client)
	callOpts := GetOptions(false, address, "")
	// TODO: Check NumberOfBlocks value here. Should it be 10?
	blockHashes, err := randomClient.BlockHashes(&callOpts, uint8(core.NumberOfBlocks))
	if err != nil {
		return nil, err
	}
	return blockHashes[:], err
}

