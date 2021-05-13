package utils

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"razor/accounts"
	"razor/core"
	"razor/core/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
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
	if err != nil {
		log.Fatal("Error in fetching pending nonce: ", err)
	}

	if err != nil {
		log.Fatal("Error in fetching gas multiplier", err)
	}
	gasPrice := getGasPrice(transactionData.Client, transactionData.GasMultiplier)

	txnOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, transactionData.ChainId)
	txnOpts.Nonce = big.NewInt(int64(nonce))
	txnOpts.GasLimit = uint64(core.GasLimit)
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