package utils

import (
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
)

func getAssetManagerWithOpts(client *ethclient.Client, address string) (*bindings.AssetManager, bind.CallOpts) {
	return GetAssetManager(client), GetOptions(false, address, "")
}

func GetNumAssets(client *ethclient.Client, address string) (uint8, error) {
	assetManager, callOpts := getAssetManagerWithOpts(client, address)
	var (
		numAssets uint8
		err       error
	)
	for retry := 1; retry <= core.MaxRetries; retry++ {
		numAssets, err = assetManager.GetNumAssets(&callOpts)
		if err != nil {
			Retry(retry, "Error in fetching numAssets: ", err)
			continue
		}
		break
	}
	if err != nil {
		return 0, err
	}
	return numAssets, nil
}

func GetNumActiveAssets(client *ethclient.Client, address string) (*big.Int, error) {
	assetManager, callOpts := getAssetManagerWithOpts(client, address)
	var (
		numActiveAssets *big.Int
		err             error
	)
	for retry := 1; retry <= core.MaxRetries; retry++ {
		numActiveAssets, err = assetManager.GetNumActiveAssets(&callOpts)
		if err != nil {
			Retry(retry, "Error in fetching numActiveAssets: ", err)
			continue
		}
		break
	}
	if err != nil {
		return nil, err
	}
	return numActiveAssets, nil
}

func GetAssetType(client *ethclient.Client, address string, assetId uint8) (uint8, error) {
	assetManager, callOpts := getAssetManagerWithOpts(client, address)
	var (
		activeAsset types.Asset
		err         error
	)
	for retry := 1; retry <= core.MaxRetries; retry++ {
		activeAsset, err = assetManager.GetAsset(&callOpts, assetId)
		if err != nil {
			Retry(retry, "Error in fetching asset type: ", err)
			continue
		}
		break
	}
	if err != nil {
		return 0, err
	}
	if activeAsset.Job.Id == 0 {
		return 2, nil
	}
	return 1, nil
}

func GetCollection(client *ethclient.Client, address string, collectionId uint8) (bindings.StructsCollection, error) {
	assetManager, callOpts := getAssetManagerWithOpts(client, address)
	var (
		asset types.Asset
		err   error
	)
	for retry := 1; retry <= core.MaxRetries; retry++ {
		asset, err = assetManager.GetAsset(&callOpts, collectionId)
		if err != nil {
			Retry(retry, "Error in fetching collection "+string(collectionId)+": ", err)
			continue
		}
		break
	}
	if err != nil {
		return bindings.StructsCollection{}, err
	}
	return asset.Collection, nil
}

func GetActiveAssetsData(client *ethclient.Client, address string, epoch uint32) ([]*big.Int, error) {
	var data []*big.Int

	numOfAssets, err := GetNumAssets(client, address)
	if err != nil {
		return data, err
	}

	for assetIndex := 1; assetIndex <= int(numOfAssets); assetIndex++ {
		assetType, err := GetAssetType(client, address, uint8(assetIndex))
		if err != nil {
			log.Error("Error in fetching asset type: ", assetType)
			return nil, err
		}
		if assetType == 2 {
			activeCollection, err := GetActiveCollection(client, address, uint8(assetIndex))
			if err != nil {
				log.Error(err)
				if err == errors.New("collection inactive") {
					continue
				}
				return nil, err
			}
			//Supply previous epoch to Aggregate in case if last reported value is required.
			collectionData, aggregationError := Aggregate(client, address, epoch-1, activeCollection)
			if aggregationError != nil {
				return nil, aggregationError
			}
			data = append(data, collectionData)
		}
	}
	return data, nil
}

func Aggregate(client *ethclient.Client, address string, previousEpoch uint32, collection bindings.StructsCollection) (*big.Int, error) {
	if len(collection.JobIDs) == 0 {
		return nil, errors.New("no jobs present in the collection")
	}
	var jobs []bindings.StructsJob
	for _, id := range collection.JobIDs {
		job, err := GetActiveJob(client, address, id)
		if err != nil {
			log.Errorf("Error in fetching job %d: %s", id, err)
			continue
		}
		jobs = append(jobs, job)
	}
	dataToCommit, weight, err := GetDataToCommitFromJobs(jobs)
	if err != nil || len(dataToCommit) == 0 {
		prevCommitmentData, err := FetchPreviousValue(client, address, previousEpoch, collection.Id)
		if err != nil {
			return nil, err
		}
		return big.NewInt(int64(prevCommitmentData)), nil
	}
	return performAggregation(dataToCommit, weight, collection.AggregationMethod)
}

func GetActiveJob(client *ethclient.Client, address string, jobId uint8) (bindings.StructsJob, error) {
	assetManager := GetAssetManager(client)
	callOpts := GetOptions(false, address, "")
	var (
		job bindings.StructsJob
		err error
	)
	for retry := 1; retry <= core.MaxRetries; retry++ {
		job, err = assetManager.Jobs(&callOpts, jobId)
		if err != nil {
			Retry(retry, "Error in fetching job "+string(jobId)+": ", err)
			continue
		}
		break
	}

	if err != nil {
		return bindings.StructsJob{}, err
	}
	return job, nil
}

func GetActiveCollection(client *ethclient.Client, address string, collectionId uint8) (bindings.StructsCollection, error) {
	collection, err := GetCollection(client, address, collectionId)
	if err != nil {
		return bindings.StructsCollection{}, err
	}
	if !collection.Active {
		return bindings.StructsCollection{}, errors.New("collection inactive")
	}
	return collection, nil
}

func GetDataToCommitFromJobs(jobs []bindings.StructsJob) ([]*big.Int, []uint8, error) {
	var (
		data   []*big.Int
		weight []uint8
	)
	for _, job := range jobs {
		dataToAppend, err := GetDataToCommitFromJob(job)
		if err != nil {
			continue
		}
		data = append(data, dataToAppend)
		weight = append(weight, job.Weight)
	}
	return data, weight, nil
}

func GetDataToCommitFromJob(job bindings.StructsJob) (*big.Int, error) {
	var parsedJSON map[string]interface{}
	var (
		response []byte
		apiErr   error
	)

	// Fetch data from API with retry mechanism
	for retry := 1; retry <= core.MaxRetries; retry++ {
		response, apiErr = GetDataFromAPI(job.Url)
		if apiErr != nil {
			Retry(retry, "Error in fetching data from API: ", apiErr)
			continue
		}
		break
	}

	err := json.Unmarshal(response, &parsedJSON)
	if err != nil {
		log.Error("Error in parsing data from API: ", err)
		return nil, err
	}
	var parsedData interface{}
	if job.SelectorType == 1 {
		parsedData, err = GetDataFromJSON(parsedJSON, job.Selector)
		if err != nil {
			log.Error("Error in fetching value from parsed data: ", err)
			return nil, err
		}
	}
	//else {
	//TODO: Add support for XHTML selector
	//}

	datum, err := ConvertToNumber(parsedData)
	if err != nil {
		log.Error("Result is not a number")
		return nil, err
	}

	return MultiplyWithPower(datum, job.Power), err
}
