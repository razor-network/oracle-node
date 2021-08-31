package utils

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"razor/pkg/bindings"
)

func getParametersManagerWithOpts(client *ethclient.Client, address string) (*bindings.Parameters, bind.CallOpts) {
	return GetParametersManager(client), GetOptions(false, address, "")
}

func GetMinStakeAmount(client *ethclient.Client, address string) (*big.Int, error) {
	parametersManager, callOpts := getParametersManagerWithOpts(client, address)
	return parametersManager.MinStake(&callOpts)
}

func GetEpoch(client *ethclient.Client, address string) (*big.Int, error) {
	parametersManager, callOpts := getParametersManagerWithOpts(client, address)
	return parametersManager.GetEpoch(&callOpts)
}

func GetWithdrawReleasePeriod(client *ethclient.Client, address string) (*big.Int, error) {
	parametersManager, callOpts := getParametersManagerWithOpts(client, address)
	return parametersManager.WithdrawReleasePeriod(&callOpts)
}

func GetMaxAltBlocks(client *ethclient.Client, address string) (*big.Int, error) {
	parametersManager, callOpts := getParametersManagerWithOpts(client, address)
	return parametersManager.MaxAltBlocks(&callOpts)
}
