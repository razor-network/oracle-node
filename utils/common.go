package utils

import (
	"context"
	"errors"
	"math/big"
	"os"
	"razor/core"
	"razor/core/types"
	"razor/logger"
	"time"

	"github.com/avast/retry-go"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/spf13/pflag"
)

func (*UtilsStruct) ConnectToClient(provider string) *ethclient.Client {
	client, err := EthClient.Dial(provider)
	if err != nil {
		log.Fatal("Error in connecting...", err)
	}
	log.Info("Connected to: ", provider)
	return client
}

func (*UtilsStruct) FetchBalance(client *ethclient.Client, accountAddress string) (*big.Int, error) {
	var (
		balance *big.Int
		err     error
	)
	err = retry.Do(
		func() error {
			address := common.HexToAddress(accountAddress)
			erc20Contract := UtilsInterface.GetTokenManager(client)
			opts := UtilsInterface.GetOptions()
			balance, err = CoinInterface.BalanceOf(erc20Contract, &opts, address)
			if err != nil {
				log.Error("Error in fetching balance....Retrying")
				return err
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if err != nil {
		return big.NewInt(0), err
	}
	return balance, nil
}

func (*UtilsStruct) GetBufferedState(client *ethclient.Client, buffer int32) (int64, error) {
	block, err := ClientInterface.GetLatestBlockWithRetry(client)
	if err != nil {
		return -1, err
	}
	stateBuffer, err := UtilsInterface.GetStateBuffer(client)
	if err != nil {
		return -1, err
	}
	lowerLimit := (core.StateLength * uint64(buffer)) / 100
	upperLimit := core.StateLength - (core.StateLength*uint64(buffer))/100
	if block.Time%(core.StateLength) > upperLimit-stateBuffer || block.Time%(core.StateLength) < lowerLimit+stateBuffer {
		return -1, nil
	}
	state := block.Time / core.StateLength
	return int64(state % core.NumberOfStates), nil
}

func (*UtilsStruct) CheckTransactionReceipt(client *ethclient.Client, _txHash string) int {
	txHash := common.HexToHash(_txHash)
	tx, err := ClientInterface.TransactionReceipt(client, context.Background(), txHash)
	if err != nil {
		return -1
	}
	return int(tx.Status)
}

func (*UtilsStruct) WaitForBlockCompletion(client *ethclient.Client, hashToRead string) error {
	timeout := core.BlockCompletionTimeout
	for start := time.Now(); time.Since(start) < time.Duration(timeout)*time.Second; {
		log.Debug("Checking if transaction is mined....")
		transactionStatus := UtilsInterface.CheckTransactionReceipt(client, hashToRead)
		if transactionStatus == 0 {
			err := errors.New("transaction mining unsuccessful")
			log.Error(err)
			return err
		} else if transactionStatus == 1 {
			log.Info("Transaction mined successfully")
			return nil
		}
		Time.Sleep(time.Second)
	}
	log.Info("Timeout Passed")
	return errors.New("timeout passed for transaction mining")
}

func (*UtilsStruct) WaitTillNextNSecs(waitTime int32) {
	if waitTime <= 0 {
		waitTime = 1
	}
	Time.Sleep(time.Duration(waitTime) * time.Second)
}

func CheckError(msg string, err error) {
	if err != nil {
		log.Fatal(msg + err.Error())
	}
}

func IsValidAddress(address string) bool {
	if !common.IsHexAddress(address) {
		log.Error("Invalid Address")
		return false
	}
	return true
}

func ValidateAddress(address string) (string, error) {
	if !IsValidAddress(address) {
		return "", errors.New("invalid address")
	}
	return address, nil
}

func (*UtilsStruct) IsFlagPassed(name string) bool {
	found := false
	for _, arg := range os.Args {
		if arg == "--"+name {
			found = true
		}
	}
	return found
}

func (*UtilsStruct) CheckEthBalanceIsZero(client *ethclient.Client, address string) {
	ethBalance, err := ClientInterface.BalanceAt(client, context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		log.Fatalf("Error in fetching sFuel balance of the account: %s\n%s", address, err)
	}
	if ethBalance.Cmp(big.NewInt(0)) == 0 {
		log.Fatal("sFuel balance is 0, Aborting...")
	}
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
		stateName = "Buffer"
	}
	return stateName
}

func (*UtilsStruct) AssignStakerId(flagSet *pflag.FlagSet, client *ethclient.Client, address string) (uint32, error) {
	if UtilsInterface.IsFlagPassed("stakerId") {
		return UtilsInterface.GetUint32(flagSet, "stakerId")
	}
	return UtilsInterface.GetStakerId(client, address)
}

func (*UtilsStruct) GetEpoch(client *ethclient.Client) (uint32, error) {
	latestHeader, err := ClientInterface.GetLatestBlockWithRetry(client)
	if err != nil {
		log.Error("Error in fetching block: ", err)
		return 0, err
	}
	epoch := latestHeader.Time / core.EpochLength
	return uint32(epoch), nil
}

func (*UtilsStruct) CalculateBlockTime(client *ethclient.Client) int64 {
	latestBlock, err := ClientInterface.GetLatestBlockWithRetry(client)
	if err != nil {
		log.Fatalf("Error in fetching latest Block: %s", err)
	}
	latestBlockNumber := latestBlock.Number
	lastSecondBlock, err := ClientInterface.HeaderByNumber(client, context.Background(), big.NewInt(1).Sub(latestBlockNumber, big.NewInt(1)))
	if err != nil {
		log.Fatalf("Error in fetching last second Block: %s", err)
	}
	return int64(latestBlock.Time - lastSecondBlock.Time)
}

func (*UtilsStruct) GetRemainingTimeOfCurrentState(client *ethclient.Client, bufferPercent int32) (int64, error) {
	block, err := ClientInterface.GetLatestBlockWithRetry(client)
	if err != nil {
		return 0, err
	}
	stateBuffer, err := UtilsInterface.GetStateBuffer(client)
	if err != nil {
		return 0, err
	}
	timeRemaining := core.StateLength - (block.Time % core.StateLength)
	upperLimit := ((core.StateLength * uint64(bufferPercent)) / 100) + stateBuffer

	return int64(timeRemaining - upperLimit), nil
}

func (*UtilsStruct) CalculateSalt(epoch uint32, medians []*big.Int) [32]byte {
	salt := solsha3.SoliditySHA3([]string{"uint32", "uint256"}, []interface{}{epoch, medians})
	var saltInBytes32 [32]byte
	copy(saltInBytes32[:], salt)
	return saltInBytes32
}

func (*UtilsStruct) Prng(max uint32, prngHashes []byte) *big.Int {
	sum := big.NewInt(0).SetBytes(prngHashes)
	maxBigInt := big.NewInt(int64(max))
	return sum.Mod(sum, maxBigInt)
}

func (*UtilsStruct) EstimateBlockNumberAtEpochBeginning(client *ethclient.Client, currentBlockNumber *big.Int) (*big.Int, error) {
	block, err := ClientInterface.HeaderByNumber(client, context.Background(), currentBlockNumber)
	if err != nil {
		log.Error("Error in fetching block: ", err)
		return nil, err
	}
	currentEpoch := block.Time / core.EpochLength
	previousBlockNumber := block.Number.Uint64() - core.StateLength

	previousBlock, err := ClientInterface.HeaderByNumber(client, context.Background(), big.NewInt(int64(previousBlockNumber)))
	if err != nil {
		log.Error("Err in fetching Previous block: ", err)
		return nil, err
	}
	previousBlockActualTimestamp := previousBlock.Time
	previousBlockAssumedTimestamp := block.Time - core.EpochLength
	previousEpoch := previousBlockActualTimestamp / core.EpochLength
	if previousBlockActualTimestamp > previousBlockAssumedTimestamp && previousEpoch != currentEpoch-1 {
		return UtilsInterface.EstimateBlockNumberAtEpochBeginning(client, big.NewInt(int64(previousBlockNumber)))

	}
	return big.NewInt(int64(previousBlockNumber)), nil

}

func (*FileStruct) SaveDataToCommitJsonFile(filePath string, epoch uint32, commitData types.CommitData) error {

	var data types.CommitFileData
	data.Epoch = epoch
	data.AssignedCollections = commitData.AssignedCollections
	data.SeqAllottedCollections = commitData.SeqAllottedCollections
	data.Leaves = commitData.Leaves

	jsonData, err := JsonInterface.Marshal(data)
	if err != nil {
		return err
	}
	err = OS.WriteFile(filePath, jsonData, 0600)
	if err != nil {
		log.Error("Error in writing to file: ", err)
		return err
	}
	return nil
}

func (*FileStruct) ReadFromCommitJsonFile(filePath string) (types.CommitFileData, error) {
	jsonFile, err := OS.Open(filePath)
	if err != nil {
		log.Error("Error in opening json file: ", err)
		return types.CommitFileData{}, err
	}
	byteValue, err := IOInterface.ReadAll(jsonFile)
	if err != nil {
		log.Error("Error in reading data from json file: ", err)
		return types.CommitFileData{}, err
	}
	var commitedData types.CommitFileData

	err = JsonInterface.Unmarshal(byteValue, &commitedData)
	if err != nil {
		log.Error(" Unmarshal error: ", err)
		return types.CommitFileData{}, err
	}
	return commitedData, nil
}

func (*FileStruct) AssignLogFile(flagSet *pflag.FlagSet, configurations types.Configurations) {
	if UtilsInterface.IsFlagPassed("logFile") {
		fileName, err := FlagSetInterface.GetLogFileName(flagSet)
		if err != nil {
			log.Fatal("Error in getting file name: ", err)
		}
		log.Debug("Log file name: ", fileName)
		logger.InitializeLogger(fileName, configurations)
	} else {
		log.Debug("No `logFile` flag passed, not storing logs in any file")
	}
}

func (*FileStruct) SaveDataToProposeJsonFile(filePath string, proposeData types.ProposeFileData) error {

	var data types.ProposeFileData
	data.Epoch = proposeData.Epoch
	data.MediansData = proposeData.MediansData
	data.RevealedCollectionIds = proposeData.RevealedCollectionIds
	data.RevealedDataMaps = proposeData.RevealedDataMaps

	jsonData, err := JsonInterface.Marshal(data)
	if err != nil {
		return err
	}
	err = OS.WriteFile(filePath, jsonData, 0600)
	if err != nil {
		log.Error("Error in writing to file: ", err)
		return err
	}
	return nil
}

func (*FileStruct) ReadFromProposeJsonFile(filePath string) (types.ProposeFileData, error) {
	jsonFile, err := OS.Open(filePath)
	if err != nil {
		log.Error("Error in opening json file: ", err)
		return types.ProposeFileData{}, err
	}
	byteValue, err := IOInterface.ReadAll(jsonFile)
	if err != nil {
		log.Error("Error in reading data from json file: ", err)
		return types.ProposeFileData{}, err
	}
	var proposedData types.ProposeFileData

	err = JsonInterface.Unmarshal(byteValue, &proposedData)
	if err != nil {
		log.Error(" Unmarshal error: ", err)
		return types.ProposeFileData{}, err
	}
	return proposedData, nil
}

func (*FileStruct) SaveDataToDisputeJsonFile(filePath string, bountyIdQueue []uint32) error {
	var data types.DisputeFileData

	data.BountyIdQueue = bountyIdQueue
	jsonData, err := JsonInterface.Marshal(data)
	if err != nil {
		return err
	}
	err = OS.WriteFile(filePath, jsonData, 0600)
	if err != nil {
		log.Error("Error in writing to file: ", err)
		return err
	}
	return nil
}

func (*FileStruct) ReadFromDisputeJsonFile(filePath string) (types.DisputeFileData, error) {
	jsonFile, err := OS.Open(filePath)
	if err != nil {
		log.Error("Error in opening json file: ", err)
		return types.DisputeFileData{}, err
	}
	byteValue, err := IOInterface.ReadAll(jsonFile)
	if err != nil {
		log.Error("Error in reading data from json file: ", err)
		return types.DisputeFileData{}, err
	}
	var disputeData types.DisputeFileData

	err = JsonInterface.Unmarshal(byteValue, &disputeData)
	if err != nil {
		log.Error(" Unmarshal error: ", err)
		return types.DisputeFileData{}, err
	}
	return disputeData, nil
}
