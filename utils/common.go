package utils

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"github.com/wealdtech/go-merkletree"
	"github.com/wealdtech/go-merkletree/keccak256"
	"math"
	"math/big"
	"os"
	"razor/core"
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

func FetchBalance(client *ethclient.Client, accountAddress string) (*big.Int, error) {
	address := common.HexToAddress(accountAddress)
	coinContract := GetTokenManager(client)
	opts := GetOptions(false, accountAddress, "")
	return coinContract.BalanceOf(&opts, address)
}

func GetDefaultPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	defaultPath := home + "/.razor"
	if _, err := os.Stat(defaultPath); os.IsNotExist(err) {
		// TODO: Restrict permission
		os.Mkdir(defaultPath, 0777)
	}
	return defaultPath
}

func GetDelayedState(client *ethclient.Client) (int64, error) {
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		return -1, err
	}
	// TODO: EpochLength should be renamed StateLength
	// TODO: Buffer should be set in config
	if blockNumber%(core.EpochLength) > 7 || blockNumber%(core.EpochLength) < 1 {
		return -1, nil
	}
	state := math.Floor(float64(blockNumber / core.EpochLength))
	return int64(state) % core.NumberOfStates, nil
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
	// TODO: STATELENGTH * BLOCKTIME = TIMEOUT
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
		time.Sleep(3 * time.Second)
	}
}

func GetMerkleTree(data []*big.Int) (*merkletree.MerkleTree, error) {
	bytesData := GetDataInBytes(data)
	return merkletree.NewUsingV1(bytesData, keccak256.New(), nil)
}

func GetMerkleTreeRoot(data []*big.Int) ([]byte, error) {
	tree, err := GetMerkleTree(data)
	if err != nil {
		return nil, err
	}
	return tree.RootV1(), err
}