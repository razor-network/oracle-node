package utils

import (
	"context"
	"github.com/spf13/pflag"
	"math"
	"math/big"
	"os"
	"razor/core"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wealdtech/go-merkletree"
	"github.com/wealdtech/go-merkletree/keccak256"
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

func GetDelayedState(client *ethclient.Client, buffer int32) (int64, error) {
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		return -1, err
	}
	lowerLimit := (core.StateLength * uint64(buffer)) / 100
	upperLimit := core.StateLength - (core.StateLength*uint64(buffer))/100
	if blockNumber%(core.StateLength) > upperLimit || blockNumber%(core.StateLength) < lowerLimit {
		return -1, nil
	}
	state := blockNumber / core.StateLength
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
	timeout := core.StateLength * 2
	for start := time.Now(); time.Since(start) < time.Duration(timeout)*time.Second; {
		log.Debug("Checking if transaction is mined....")
		transactionStatus := checkTransactionReceipt(client, hashToRead)
		if transactionStatus == 0 {
			log.Error("Transaction mining unsuccessful")
			return 0
		} else if transactionStatus == 1 {
			log.Info("Transaction mined successfully")
			return 1
		}
		time.Sleep(3 * time.Second)
	}
	log.Info("Timeout Passed")
	return 0
}

func WaitTillNextNSecs(waitTime int32) {
	if waitTime <= 0 {
		waitTime = 1
	}
	time.Sleep(time.Duration(waitTime) * time.Second)
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

func CheckError(msg string, err error) {
	if err != nil {
		log.Fatal(msg + err.Error())
	}
}

func IsFlagPassed(name string) bool {
	found := false
	for _, arg := range os.Args {
		if arg == "--"+name {
			found = true
		}
	}
	return found
}

func CheckEthBalanceIsZero(client *ethclient.Client, address string) {
	ethBalance, err := client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		log.Fatalf("Error in fetching eth balance of the account: %s\n%s", address, err)
	}
	if ethBalance.Cmp(big.NewInt(0)) == 0 {
		log.Fatal("Eth balance is 0, Aborting...")
	}
}

func Retry(retry int, errMsg string, err error) {
	log.Error(errMsg, err)
	retryingIn := math.Pow(2, float64(retry))
	log.Debugf("Retrying in %f seconds.....", retryingIn)
	time.Sleep(time.Duration(retryingIn) * time.Second)
}

func GetStateName(stateNumber int64) string {
	var stateName string
	switch stateNumber {
	case 0:
		stateName = "Commit"
	case 1:
		stateName = "Reveal"
	case 2:
		stateName = "Propose"
	case 3:
		stateName = "Dispute"
	case 4:
		stateName = "Confirm"
	default:
		stateName = "-1"
	}
	return stateName
}

func AssignStakerId(flagSet *pflag.FlagSet, client *ethclient.Client, address string) (uint32, error) {
	if IsFlagPassed("stakerId") {
		return flagSet.GetUint32("stakerId")
	}
	return GetStakerId(client, address)
}
