package utils

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"razor/accounts"
	"razor/core/types"
	"strings"

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

	gasPrice := getGasPrice(transactionData.Client, transactionData.Config)

	txnOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, transactionData.ChainId)
	CheckError("Error in getting transactor: ", err)
	txnOpts.Nonce = big.NewInt(int64(nonce))
	txnOpts.GasPrice = gasPrice
	txnOpts.Value = transactionData.EtherValue
	gasLimit, err := getGasLimit(transactionData, txnOpts)
	CheckError("Error in getting gas limit: ", err)
	log.Info("Gas Limit: ", gasLimit)
	txnOpts.GasLimit = (gasLimit*3)/20 + gasLimit
	log.Info("Gas Limit after increment: ", txnOpts.GasLimit)
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

func getGasLimit(transactionData types.TransactionOptions, txnOpts *bind.TransactOpts) (uint64, error) {
	if transactionData.MethodName == "" {
		return 0, nil
	}
	parsed, err := abi.JSON(strings.NewReader(transactionData.ABI))
	inputData, err := parsed.Pack(transactionData.MethodName, transactionData.Parameters...)
	if err != nil {
		log.Error("Error in calculating inputData: ", err)
		return 0, err
	}
	contractAddress := common.HexToAddress(transactionData.ContractAddress)
	msg := ethereum.CallMsg{
		From:     common.HexToAddress(transactionData.AccountAddress),
		To:       &contractAddress,
		GasPrice: txnOpts.GasPrice,
		Value:    txnOpts.Value,
		Data:     inputData,
	}
	return transactionData.Client.EstimateGas(context.Background(), msg)
}
