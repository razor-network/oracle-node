package utils

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func GetMinStakeAmount(client *ethclient.Client, address string) (*big.Int, error) {
	constantsManager := GetConstantsManager(client)
	callOpts := GetOptions(false, address, "")
	return constantsManager.MinStake(&callOpts)
}