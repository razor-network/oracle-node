package utils

import (
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"razor/core/types"
)

func GetNumAssets(client *ethclient.Client, address string) (uint8, error) {
	assetManager := GetAssetManager(client)
	callOpts := GetOptions(false, address, "")
	return assetManager.GetNumAssets(&callOpts)
}

func GetNumActiveAssets(client *ethclient.Client, address string) (uint8, error) {
	assetManager := GetAssetManager(client)
	callOpts := GetOptions(false, address, "")
	return assetManager.GetNumActiveAssets(&callOpts)
}

func GetActiveAssetIds(client *ethclient.Client, address string) ([]*big.Int, error) {
	numAssets, err := GetNumAssets(client, address)
	if err != nil {
		return nil, err
	}
	assetManager := GetAssetManager(client)
	callOpts := GetOptions(false, address, "")
	var activeAssets []*big.Int
	for assetId := 1; assetId <= int(numAssets); assetId++ {
		assetType, err := assetManager.GetAssetType(&callOpts, uint8(assetId))
		if err != nil {
			//TODO: Implement retry
			log.Error("Error in calling GetActiveStatus: ", err)
		}
		if assetType == 1 {
			continue
		}
		isActiveAsset, err := assetManager.GetActiveStatus(&callOpts, uint8(assetId))
		if err != nil {
			//TODO: Implement retry
			log.Error("Error in calling GetActiveStatus: ", err)
		}
		if isActiveAsset {
			activeAssets = append(activeAssets, big.NewInt(int64(assetId)))
		}
	}
	return activeAssets, nil
}

func GetActiveAssetsData(client *ethclient.Client, address string) ([]*big.Int, error) {
	var data []*big.Int

	numOfAssets, err := GetNumAssets(client, address)
	if err != nil {
		return data, err
	}

	assetManager := GetAssetManager(client)
	callOpts := GetOptions(false, address, "")

	for assetIndex := 1; assetIndex <= int(numOfAssets); assetIndex++ {
		assetType, err := assetManager.GetAssetType(&callOpts, uint8(assetIndex))
		if err != nil {
			log.Error("Error in fetching asset: ", assetIndex)
			continue
		}
		if assetType == 2 {
			activeCollection, err := GetActiveCollection(client, address, uint8(assetIndex))
			if err != nil {
				log.Error(err)
				continue
			}
			collectionData, err := Aggregate(client, address, activeCollection)
			if err != nil {
				collectionData = big.NewInt(1)
			}
			data = append(data, collectionData)
		}
	}
	return data, nil
}

func GetActiveJob(client *ethclient.Client, address string, jobId uint8) (types.Job, error) {
	assetManager := GetAssetManager(client)
	callOpts := GetOptions(false, address, "")
	job, err := assetManager.Jobs(&callOpts, jobId)
	if err != nil {
		return types.Job{}, err
	}
	if job.Active {
		return job, nil
	}
	return types.Job{}, errors.New("job already fulfilled")
}

func GetActiveCollection(client *ethclient.Client, address string, collectionId uint8) (types.Collection, error) {
	assetManager := GetAssetManager(client)
	callOpts := GetOptions(false, address, "")
	collection, err := assetManager.GetCollection(&callOpts, collectionId)
	if err != nil {
		return types.Collection{}, err
	}
	if !collection.Active {
		return types.Collection{}, nil
	}
	return types.Collection{
		Id:                collectionId,
		Name:              collection.Name,
		AggregationMethod: collection.AggregationMethod,
		JobIDs:            collection.JobIDs,
		Power:             collection.Power,
	}, nil
}

func GetDataToCommitFromJobs(jobs []types.Job) []*big.Int {
	var data []*big.Int
	for _, job := range jobs {
		dataToAppend := GetDataToCommitFromJob(job)
		data = append(data, dataToAppend)
	}
	if len(data) == 0 {
		data = append(data, big.NewInt(1))
	}
	return data
}

func GetDataToCommitFromJob(job types.Job) *big.Int {
	var parsedJSON map[string]interface{}

	response, err := GetDataFromAPI(job.Url)
	if err != nil {
		log.Error(err)
		return big.NewInt(1)
	}

	err = json.Unmarshal(response, &parsedJSON)
	if err != nil {
		log.Error("Error in parsing data from API: ", err)
		return big.NewInt(1)
	}

	parsedData, err := GetDataFromJSON(parsedJSON, job.Selector)
	if err != nil {
		log.Error("Error in fetching value from parsed data: ", err)
		return big.NewInt(1)
	}

	datum, err := ConvertToNumber(parsedData)
	if err != nil {
		log.Error("Result is not a number")
		return big.NewInt(1)
	}

	return MultiplyWithPower(datum, job.Power)
}
