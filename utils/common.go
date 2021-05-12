package utils

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
	"os"
	"time"
)

func ConnectToClient(provider string) *ethclient.Client {
	client, err := ethclient.Dial(provider)
	if err != nil {
		log.Fatal("Error in connecting...\n", err)
	}
	log.Info("Connected to: ", provider)
	return client
}

func FetchBalance(client *ethclient.Client, accountAddress string) *big.Int {
	address := common.HexToAddress(accountAddress)
	coinContract := GetCoinContract(client)
	opts := GetOptions(false, accountAddress, "")

	balance, err := coinContract.BalanceOf(&opts, address)
	if err != nil {
		log.Fatalf("Error in getting balance for account: %s\n%s", accountAddress, err)
	}
	return balance
}

func GetDefaultPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	defaultPath := home + "/.razor"
	if _, err := os.Stat(defaultPath); os.IsNotExist(err) {
		os.Mkdir(defaultPath, 0777)
	}
	return defaultPath
}

func GetEpoch(client *ethclient.Client, address string) *big.Int {
	stateManager := GetStateManager(client)
	callOpts := GetOptions(false, address, "")
	epoch, err := stateManager.GetEpoch(&callOpts)
	if err != nil {
		log.Fatal("Error in fetching epoch: ", err)
	}
	return epoch
}

func checkTransactionReceipt(client *ethclient.Client, _txHash string) int {
	txHash := common.HexToHash(_txHash)
	tx, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return -1
	}
	return int(tx.Status)
}

func WaitForBlockCompletion(client *ethclient.Client, hashToRead string) int {
	for {
		log.Info("Checking if transaction is mined....\n")
		transactionStatus := checkTransactionReceipt(client, hashToRead)
		if transactionStatus == 0 {
			log.Info("Transaction mining unsuccessful")
			return 0
		} else if transactionStatus == 1 {
			log.Info("Transaction mined successfully\n")
			return 1
		}
		time.Sleep(2 * time.Second)
	}
}
