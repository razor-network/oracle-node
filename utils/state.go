package utils

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func GetEpoch(client *ethclient.Client, address string) (*big.Int, error) {
	stateManager := GetStateManager(client)
	callOpts := GetOptions(false, address, "")
	return stateManager.GetEpoch(&callOpts)
}
