package utils

import (
	"context"
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
	nonce, err := transactionData.Client.PendingNonceAt(context.Background(), common.HexToAddress(transactionData.AccountAddress))
	CheckError("Error in fetching pending nonce: ", err)

	gasPrice := getGasPrice(transactionData.Client, transactionData.GasMultiplier)

	txnOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, transactionData.ChainId)
	CheckError("Error in getting transactor: ", err)
	txnOpts.Nonce = big.NewInt(int64(nonce))
	txnOpts.GasPrice = gasPrice
	txnOpts.Value = transactionData.EtherValue

	return txnOpts
}

func getGasPrice(client *ethclient.Client, gasMultiplier float32) *big.Int {
	gas, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	gasPrice := MultiplyFloatAndBigInt(gas, float64(gasMultiplier))
	return gasPrice
}
