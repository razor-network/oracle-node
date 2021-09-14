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

type CollectionStruct struct {
	Active            bool
	Power             int8
	JobIDs            []uint8
	AggregationMethod uint32
	Name              string
}

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

func GetNumActiveAssets(client *ethclient.Client, address string) (uint8, error) {
	assetManager, callOpts := getAssetManagerWithOpts(client, address)
	var (
		numActiveAssets uint8
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
		return 0, err
	}
	return numActiveAssets, nil
}

func GetAssetType(client *ethclient.Client, address string, assetId uint8) (uint8, error) {
	assetManager, callOpts := getAssetManagerWithOpts(client, address)
	var (
		numActiveAssets uint8
		err             error
	)
	for retry := 1; retry <= core.MaxRetries; retry++ {
		numActiveAssets, err = assetManager.GetAssetType(&callOpts, assetId)
		if err != nil {
			Retry(retry, "Error in fetching asset type: ", err)
			continue
		}
		break
	}
	if err != nil {
		return 0, err
	}
	return numActiveAssets, nil
}

func GetCollection(client *ethclient.Client, address string, collectionId uint8) (CollectionStruct, error) {
	assetManager, callOpts := getAssetManagerWithOpts(client, address)
	var (
		collection CollectionStruct
		err        error
	)
	for retry := 1; retry <= core.MaxRetries; retry++ {
		collection, err = assetManager.GetCollection(&callOpts, collectionId)
		if err != nil {
			Retry(retry, "Error in fetching collection "+string(collectionId)+": ", err)
			continue
		}
		break
	}
	if err != nil {
		return CollectionStruct{
			Active:            false,
			Power:             0,
			JobIDs:            nil,
			AggregationMethod: 0,
			Name:              "",
		}, err
	}
	return collection, nil
}

func GetActiveAssetsData(client *ethclient.Client, address string) ([]*big.Int, error) {
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
			collectionData, aggregationError := Aggregate(client, address, activeCollection)
			if aggregationError != nil {
				log.Error(aggregationError)
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
	var (
		job types.Job
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
		return types.Job{}, err
	}
	if job.Active {
		return job, nil
	}
	return types.Job{}, errors.New("job already fulfilled")
}

func GetActiveCollection(client *ethclient.Client, address string, collectionId uint8) (types.Collection, error) {
	collection, err := GetCollection(client, address, collectionId)
	if err != nil {
		return types.Collection{}, err
	}
	if !collection.Active {
		return types.Collection{}, errors.New("collection inactive")
	}
	return types.Collection{
		Id:                collectionId,
		Name:              collection.Name,
		AggregationMethod: collection.AggregationMethod,
		JobIDs:            collection.JobIDs,
		Power:             collection.Power,
	}, nil
}

func GetDataToCommitFromJobs(client *ethclient.Client, address string, jobs []types.Job) ([]*big.Int, error) {
	var data []*big.Int
	for _, job := range jobs {
		dataToAppend, err := GetDataToCommitFromJob(client, address, job)
		if err != nil {
			return nil, err
		}
		data = append(data, dataToAppend)
	}
	if len(data) == 0 {
		data = append(data, big.NewInt(1))
	}
	return data, nil
}

func GetDataToCommitFromJob(client *ethclient.Client, address string, job types.Job) (*big.Int, error) {
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

	// If the API still throws an error, fetch the last reported value
	if apiErr != nil {
		return FetchPreviousValue(client, address, job.Id)
	}

	err := json.Unmarshal(response, &parsedJSON)
	if err != nil {
		log.Error("Error in parsing data from API: ", err)
		return nil, err
	}

	parsedData, err := GetDataFromJSON(parsedJSON, job.Selector)
	if err != nil {
		log.Error("Error in fetching value from parsed data: ", err)
		return big.NewInt(1), nil
	}

	datum, err := ConvertToNumber(parsedData)
	if err != nil {
		log.Error("Result is not a number")
		return big.NewInt(1), err
	}

	return MultiplyWithPower(datum, job.Power), err
}

func FetchPreviousValue(client *ethclient.Client, address string, assetId uint8) (*big.Int, error) {
	stakerId, stakerErr := GetStakerId(client, address)
	if stakerErr != nil {
		return big.NewInt(1), nil
	}

	voteValue, voteValueErr := GetVoteValue(client, address, assetId, stakerId)
	if voteValueErr != nil {
		return big.NewInt(1), nil
	}
	return voteValue, nil
}
