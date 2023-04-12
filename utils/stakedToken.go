package utils

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"razor/pkg/bindings"
)

func (*UtilsStruct) GetStakedTokenWithOpts(client *ethclient.Client, tokenAddress common.Address) (*bindings.StakedToken, bind.CallOpts) {
	return UtilsInterface.GetStakedToken(client, tokenAddress), UtilsInterface.GetOptions()
}

func (*UtilsStruct) GetStakerSRZRBalance(client *ethclient.Client, staker bindings.StructsStaker) (*big.Int, error) {
	returnedValues, err := InvokeFunctionWithRetryAttempts(StakedTokenInterface, "BalanceOf", client, staker.TokenAddress, staker.Address)
	if err != nil {
		log.Error("Error in getting sRZRBalance: ", err)
		return nil, err
	}
	return returnedValues[0].Interface().(*big.Int), nil
}
