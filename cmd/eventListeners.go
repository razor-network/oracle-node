package cmd

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"razor/cache"
	"razor/core"
	"razor/pkg/bindings"
	"razor/utils"
	"strings"
	"time"
)

func initAssetCache(client *ethclient.Client) error {
	log.Info("INITIALIZING JOBS AND COLLECTIONS CACHE...")
	if err := utils.InitJobsCache(client); err != nil {
		log.Error("Error in initializing jobs cache: ", err)
		return err
	}
	if err := utils.InitCollectionsCache(client); err != nil {
		log.Error("Error in initializing collections cache: ", err)
		return err
	}

	latestHeader, err := clientUtils.GetLatestBlockWithRetry(client)
	if err != nil {
		log.Error("Error in fetching block: ", err)
		return err
	}
	log.Debugf("initAssetCache: Latest header value: %d", latestHeader.Number)

	fromBlock, err := razorUtils.EstimateBlockNumberAtEpochBeginning(client, latestHeader.Number)
	if err != nil {
		log.Error("Error in estimating block number at epoch beginning: ", err)
		return err
	}

	// Start listeners for job and collection updates
	go startListener(client, fromBlock, time.Second*time.Duration(core.AssetUpdateListenerInterval), listenForJobUpdates)
	go startListener(client, fromBlock, time.Second*time.Duration(core.AssetUpdateListenerInterval), listenForCollectionUpdates)
	go startListener(client, fromBlock, time.Second*time.Duration(core.AssetUpdateListenerInterval), listenForAssetCreation)

	return nil
}

// startListener starts a generic listener for blockchain events.
func startListener(client *ethclient.Client, fromBlock *big.Int, interval time.Duration, listenerFunc func(*ethclient.Client, abi.ABI, *big.Int, *big.Int)) {
	collectionManagerContractABI, err := abi.JSON(strings.NewReader(bindings.CollectionManagerMetaData.ABI))
	if err != nil {
		log.Errorf("Failed to parse contract ABI: %v", err)
		return
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		toBlock, err := clientUtils.GetLatestBlockWithRetry(client)
		if err != nil {
			log.Error("Error in getting latest block to start event listener: ", err)
			continue
		}

		listenerFunc(client, collectionManagerContractABI, fromBlock, toBlock.Number)
		// Update fromBlock for the next interval
		fromBlock = new(big.Int).Add(toBlock.Number, big.NewInt(1))
	}
}

// listenForJobUpdates listens and processes job update events.
func listenForJobUpdates(client *ethclient.Client, collectionManagerContractABI abi.ABI, fromBlock, toBlock *big.Int) {
	err := processEventLogs(client, collectionManagerContractABI, fromBlock, toBlock, "JobUpdated", func(topics []common.Hash, vLog Types.Log) {
		jobId := utils.ConvertHashToUint16(topics[1])
		updatedJob, err := utils.UtilsInterface.GetActiveJob(client, jobId)
		if err != nil {
			log.Errorf("Error in getting job with job Id %v: %v", jobId, err)
			return
		}
		log.Debugf("RECEIVED ASSET UPDATE: Updating the job with Id %v with details %+v...", jobId, updatedJob)
		cache.UpdateJobCache(jobId, updatedJob)
	})

	if err != nil {
		log.Errorf("Error processing JobUpdated events: %v", err)
		return
	}
}

// listenForCollectionUpdates listens and processes collection update and collection activity status events.
func listenForCollectionUpdates(client *ethclient.Client, collectionManagerContractABI abi.ABI, fromBlock, toBlock *big.Int) {
	// Process CollectionCreated event
	err := processEventLogs(client, collectionManagerContractABI, fromBlock, toBlock, "CollectionUpdated", func(topics []common.Hash, vLog Types.Log) {
		collectionId := utils.ConvertHashToUint16(topics[1])
		newCollection, err := utils.UtilsInterface.GetCollection(client, collectionId)
		if err != nil {
			log.Errorf("Error in getting collection with collection Id %v: %v", collectionId, err)
			return
		}
		log.Debugf("RECEIVED ASSET UPDATE: Updating the collection with ID %v with details %+v", collectionId, newCollection)
		cache.UpdateCollectionCache(collectionId, newCollection)
	})

	if err != nil {
		log.Errorf("Error processing CollectionCreated events: %v", err)
		return
	}

	// Process CollectionActivityStatus event
	err = processEventLogs(client, collectionManagerContractABI, fromBlock, toBlock, "CollectionActivityStatus", func(topics []common.Hash, vLog Types.Log) {
		collectionId := utils.ConvertHashToUint16(topics[1])
		updatedCollection, err := utils.UtilsInterface.GetCollection(client, collectionId)
		if err != nil {
			log.Errorf("Error in getting updated collection with collection Id %v: %v", collectionId, err)
			return
		}
		log.Debugf("RECEIVED ASSET UPDATE: Updating the activity status for collection with ID %v with details %+v", collectionId, updatedCollection)
		cache.UpdateCollectionCache(collectionId, updatedCollection)
	})

	if err != nil {
		log.Errorf("Error processing CollectionActivityStatus events: %v", err)
	}
}

// listenForAssetCreation listens and processes asset creation events.
func listenForAssetCreation(client *ethclient.Client, collectionManagerContractABI abi.ABI, fromBlock, toBlock *big.Int) {
	// Process JobCreated events
	err := processEventLogs(client, collectionManagerContractABI, fromBlock, toBlock, "JobCreated", func(topics []common.Hash, vLog Types.Log) {
		jobId := utils.ConvertHashToUint16(topics[1])
		newJob, err := utils.UtilsInterface.GetActiveJob(client, jobId)
		if err != nil {
			log.Errorf("Error in getting job with job Id %v: %v", jobId, err)
			return
		}
		log.Debugf("RECEIVED ASSET UPDATE: New JobCreated event detected for job ID %v with details %+v", jobId, newJob)
		cache.UpdateJobCache(jobId, newJob)
	})

	if err != nil {
		log.Errorf("Error processing JobCreated events: %v", err)
		return
	}

	// Process CollectionCreated events
	err = processEventLogs(client, collectionManagerContractABI, fromBlock, toBlock, "CollectionCreated", func(topics []common.Hash, vLog Types.Log) {
		collectionId := utils.ConvertHashToUint16(topics[1])
		newCollection, err := utils.UtilsInterface.GetCollection(client, collectionId)
		if err != nil {
			log.Errorf("Error in getting collection with collection Id %v: %v", collectionId, err)
			return
		}
		log.Debugf("RECEIVED ASSET UPDATE: New CollectionCreated event detected for collection ID %v with details %+v", collectionId, newCollection)
		cache.UpdateCollectionCache(collectionId, newCollection)
	})

	if err != nil {
		log.Errorf("Error processing CollectionCreated events: %v", err)
		return
	}
}

// getEventLogs is a utility function to fetch the event logs
func getEventLogs(client *ethclient.Client, fromBlock *big.Int, toBlock *big.Int) ([]Types.Log, error) {
	// Set up the query for filtering logs
	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Addresses: []common.Address{
			common.HexToAddress(core.CollectionManagerAddress),
		},
	}

	// Retrieve the logs
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Errorf("Error in filter logs: %v", err)
		return []Types.Log{}, nil
	}

	return logs, nil
}

// processEventLogs is a utility function to process the event logs using a provided handler function.
func processEventLogs(client *ethclient.Client, collectionManagerContractABI abi.ABI, fromBlock, toBlock *big.Int, eventName string, handler func(topics []common.Hash, vLog Types.Log)) error {
	logs, err := getEventLogs(client, fromBlock, toBlock)
	if err != nil {
		log.Errorf("Failed to fetch logs for %s event: %v", eventName, err)
		return err
	}

	eventID := collectionManagerContractABI.Events[eventName].ID.Hex()
	for _, vLog := range logs {
		if len(vLog.Topics) > 0 && vLog.Topics[0].Hex() == eventID {
			handler(vLog.Topics, vLog)
		}
	}

	return nil
}
