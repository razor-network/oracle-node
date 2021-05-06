package cmd

import (
	"context"
	log "github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func connectToClient(provider string) *ethclient.Client {
	client, err := ethclient.Dial(provider)
	if err != nil {
		log.Fatal("Error in connecting...\n", err)
	}
	log.Info("Connected to: ", provider)
	return client
}

func getOptions(pending bool, from string, blockNumber string) bind.CallOpts {
	block, _ := new(big.Int).SetString(blockNumber, 10)
	return bind.CallOpts{
		Pending:     pending,
		From:        common.HexToAddress(from),
		BlockNumber: block,
		Context:     context.Background(),
	}
}

func getTxnOpts(client *ethclient.Client, from string, gasMultiplier float32) bind.TransactOpts {
	nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(from))
	if err != nil {
		log.Fatal("Error in fetching pending nonce: ", err)
	}

	gasPrice := getGasPrice(client, gasMultiplier)

	return bind.TransactOpts{
		From:     common.HexToAddress(from),
		Nonce:    new(big.Int).SetInt64(int64(nonce)),
		Signer:   nil,
		Value:    nil,
		GasPrice: gasPrice,
		GasLimit: 0,
		Context:  context.Background(),
		NoSend:   false,
	}
}

func getGasPrice(client *ethclient.Client, gasMultiplier float32) *big.Int {
	gas, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	gasValue := new(big.Float).SetInt(gas)
	multiplier := big.NewFloat(float64(gasMultiplier))
	gasPrice := new(big.Float).Mul(gasValue, multiplier).String()
	gasPriceValue, ok := new(big.Int).SetString(gasPrice, 10)
	if !ok {
		log.Fatal("Error in calculating gas price")
	}
	return gasPriceValue
}

func fetchBalance(client *ethclient.Client, accountAddress string) *big.Int {
	address := common.HexToAddress(accountAddress)
	coinContract := getCoinContract(client)
	opts := getOptions(false, accountAddress, "")

	balance, err := coinContract.BalanceOf(&opts, address)
	if err != nil {
		log.Fatalf("Error in getting balance for account: %s\n%s", accountAddress, err)
	}
	return balance
}
