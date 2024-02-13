package cmd

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
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
func startListener(client *ethclient.Client, fromBlock *big.Int, interval time.Duration, listenerFunc func(*ethclient.Client, *big.Int, *big.Int)) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		toBlock, err := clientUtils.GetLatestBlockWithRetry(client)
		if err != nil {
			log.Error("Error in getting latest block to start event listener: ", err)
			continue
		}

		listenerFunc(client, fromBlock, toBlock.Number)
		// Update fromBlock for the next interval
		fromBlock = new(big.Int).Add(toBlock.Number, big.NewInt(1))
	}
}

func listenForJobUpdates(client *ethclient.Client, fromBlock, toBlock *big.Int) {
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
		log.Errorf("Failed to fetch logs for JobUpdated event: %v", err)
		return
	}

	contractAbi, err := abi.JSON(strings.NewReader(bindings.CollectionManagerMetaData.ABI))
	if err != nil {
		log.Errorf("Failed to parse contract ABI: %v", err)
		return
	}

	for _, vLog := range logs {
		// Check if the log is a JobUpdated event
		if len(vLog.Topics) > 0 && vLog.Topics[0].Hex() == contractAbi.Events["JobUpdated"].ID.Hex() {
			topics := vLog.Topics

			// topics[1] gives job id in data type common.Hash
			jobId := utils.ConvertHashToUint16(topics[1])
			updatedJob, err := utils.UtilsInterface.GetActiveJob(client, jobId)
			if err != nil {
				log.Errorf("Error in getting job with job Id %v: %v", jobId, err)
				return
			}

			log.Debugf("RECEIVED ASSET UPDATE: Updating the job with Id %v with details %+v...", jobId, updatedJob)
			cache.UpdateJobCache(jobId, updatedJob)
		}
	}
}

func listenForCollectionUpdates(client *ethclient.Client, fromBlock, toBlock *big.Int) {
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
		log.Errorf("Failed to fetch logs for CollectionUpdated event: %v", err)
		return
	}

	contractAbi, err := abi.JSON(strings.NewReader(bindings.CollectionManagerMetaData.ABI))
	if err != nil {
		log.Errorf("Failed to parse contract ABI: %v", err)
		return
	}

	for _, vLog := range logs {
		// Check if the log is a CollectionUpdated event
		if len(vLog.Topics) > 0 && vLog.Topics[0].Hex() == contractAbi.Events["CollectionUpdated"].ID.Hex() {
			topics := vLog.Topics

			// topics[1] gives collection id in data type common.Hash
			collectionId := utils.ConvertHashToUint16(topics[1])
			updatedCollection, err := utils.UtilsInterface.GetCollection(client, collectionId)
			if err != nil {
				log.Errorf("Error in getting updated collection with collection Id %v: %v", collectionId, err)
				return
			}

			log.Debugf("RECEIVED ASSET UPDATE: Updating the collection with Id %v with details %+v...", collectionId, updatedCollection)
			cache.UpdateCollectionCache(collectionId, updatedCollection)
		} else if len(vLog.Topics) > 0 && vLog.Topics[0].Hex() == contractAbi.Events["CollectionActivityStatus"].ID.Hex() {
			topics := vLog.Topics

			// topics[1] gives collection id in data type common.Hash
			collectionId := utils.ConvertHashToUint16(topics[1])
			updatedCollection, err := utils.UtilsInterface.GetCollection(client, collectionId)
			if err != nil {
				log.Errorf("Error in getting updated collection with collection Id %v: %v", collectionId, err)
				return
			}

			log.Debugf("RECEIVED ASSET UPDATE: Updating the activity status for collection with ID %v with details %+v", collectionId, updatedCollection)
			cache.UpdateCollectionCache(collectionId, updatedCollection)
		}
	}
}

func listenForAssetCreation(client *ethclient.Client, fromBlock, toBlock *big.Int) {
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
		log.Errorf("Failed to fetch logs for JobCreated event: %v", err)
		return
	}

	contractAbi, err := abi.JSON(strings.NewReader(bindings.CollectionManagerMetaData.ABI))
	if err != nil {
		log.Errorf("Failed to parse contract ABI: %v", err)
		return
	}

	for _, vLog := range logs {
		// Check if the log is a JobCreated event
		if len(vLog.Topics) > 0 && vLog.Topics[0].Hex() == contractAbi.Events["JobCreated"].ID.Hex() {
			topics := vLog.Topics

			// topics[1] gives job id in data type common.Hash
			jobId := utils.ConvertHashToUint16(topics[1])
			newJob, err := utils.UtilsInterface.GetActiveJob(client, jobId)
			if err != nil {
				log.Errorf("Error in getting job with job Id %v: %v", jobId, err)
				return
			}

			log.Debugf("RECEIVED ASSET UPDATE: New JobCreated event detected for job ID %v with details %+v", jobId, newJob)
			cache.UpdateJobCache(jobId, newJob)
		} else if len(vLog.Topics) > 0 && vLog.Topics[0].Hex() == contractAbi.Events["CollectionCreated"].ID.Hex() {
			topics := vLog.Topics

			// topics[1] gives collection id in data type common.Hash
			collectionId := utils.ConvertHashToUint16(topics[1])
			newCollection, err := utils.UtilsInterface.GetCollection(client, collectionId)
			if err != nil {
				log.Errorf("Error in getting collection with collection Id %v: %v", collectionId, err)
				return
			}

			log.Debugf("RECEIVED ASSET UPDATE: New CollectionCreated event detected for collection ID %v with details %+v", collectionId, newCollection)
			cache.UpdateCollectionCache(collectionId, newCollection)
		}
	}
}
