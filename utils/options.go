package utils

import (
	"context"
	"errors"
	"razor/accounts"
	"razor/core/types"

	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"

	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func GetOptions(pending bool, from string, blockNumber string) bind.CallOpts {
	block, _ := new(big.Int).SetString(blockNumber, 10)
	return bind.CallOpts{
		Pending:     pending,
		From:        common.HexToAddress(from),
		BlockNumber: block,
		Context:     context.Background(),
	}
}

func GetTxnOpts(transactionData types.TransactionOptions) *bind.TransactOpts {
	privateKey := accounts.GetPrivateKey(transactionData.AccountAddress, transactionData.Password, GetDefaultPath())
	if privateKey == nil {
		CheckError("Error in fetching private key: ", errors.New(transactionData.AccountAddress+" not present in razor-go"))
	}
	nonce, err := transactionData.Client.PendingNonceAt(context.Background(), common.HexToAddress(transactionData.AccountAddress))
	CheckError("Error in fetching pending nonce: ", err)

	gasPrice := getGasPrice(transactionData.Client, transactionData.Config)

	txnOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, transactionData.ChainId)
	CheckError("Error in getting transactor: ", err)
	txnOpts.Nonce = big.NewInt(int64(nonce))
	txnOpts.GasPrice = gasPrice
	txnOpts.Value = transactionData.EtherValue

	return txnOpts
}

func getGasPrice(client *ethclient.Client, config types.Configurations) *big.Int {
	var gas *big.Int
	if config.GasPrice != 0 {
		gas = big.NewInt(1).Mul(big.NewInt(int64(config.GasPrice)), big.NewInt(1e9))
	} else {
		var err error
		gas, err = client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}
	gasPrice := MultiplyFloatAndBigInt(gas, float64(config.GasMultiplier))
	return gasPrice
}
