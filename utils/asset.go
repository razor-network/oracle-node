package utils

import (
	"encoding/json"
	"errors"
	"math/big"
	"razor/core/types"

	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

func GetNumAssets(client *ethclient.Client, address string) (*big.Int, error) {
	assetManager := GetAssetManager(client)
	callOpts := GetOptions(false, address, "")
	return assetManager.GetNumAssets(&callOpts)
}

func GetActiveAssets(client *ethclient.Client, address string) ([]types.Job, []types.Collection, error) {
	var jobs []types.Job

	var collections []types.Collection

	numOfAssets, err := GetNumAssets(client, address)
	if err != nil {
		return jobs, collections, err
	}

	assetManager := GetAssetManager(client)
	callOpts := GetOptions(false, address, "")
	for assetIndex := 1; assetIndex < int(numOfAssets.Int64()); assetIndex++ {
		assetType, err := assetManager.GetAssetType(&callOpts, big.NewInt(int64(assetIndex)))
		if err != nil {
			log.Error("Error in fetching asset: ", assetIndex)
			continue
		}
		if assetType.Cmp(big.NewInt(1)) == 0 {
			activeJob, err := GetActiveJob(client, address, big.NewInt(int64(assetIndex)))
			if err != nil {
				log.Error(err)
				continue
			}
			jobs = append(jobs, activeJob)
		} else {
			activeCollection, err := GetActiveCollection(client, address, big.NewInt(int64(assetIndex)))
			if err != nil {
				log.Error(err)
				continue
			}
			collections = append(collections, activeCollection)
		}
	}
	return jobs, collections, nil
}

func GetActiveJob(client *ethclient.Client, address string, jobId *big.Int) (types.Job, error) {
	assetManager := GetAssetManager(client)
	callOpts := GetOptions(false, address, "")
	epoch, err := GetEpoch(client, address)

	if err != nil {
		return types.Job{}, err
	}
	job, err := assetManager.Jobs(&callOpts, jobId)
	if err != nil {
		return types.Job{}, err
	}
	if job.Active && job.Epoch.Cmp(epoch) < 0 {
		return job, nil
	}
	return types.Job{}, errors.New("job already fulfilled")
}

func GetActiveCollection(client *ethclient.Client, address string, collectionId *big.Int) (types.Collection, error) {
	assetManager := GetAssetManager(client)
	callOpts := GetOptions(false, address, "")
	collection, err := assetManager.GetCollection(&callOpts, collectionId)
	if err != nil {
		return types.Collection{}, err
	}
	return types.Collection{
		Id:                collectionId,
		Name:              collection.Name,
		AggregationMethod: collection.AggregationMethod,
		JobIDs:            collection.JobIDs,
		Result:            collection.Result,
	}, nil
}

func GetDataToCommitFromJobs(jobs []types.Job) []*big.Int {
	var data []*big.Int
	for _, job := range jobs {
		var parsedJSON map[string]interface{}

		response, err := GetDataFromAPI(job.Url)
		if err != nil {
			log.Error(err)
			data = append(data, big.NewInt(0))
			continue
		}

		err = json.Unmarshal(response, &parsedJSON)
		if err != nil {
			log.Error("Error in parsing data from API: ", err)
			data = append(data, big.NewInt(0))
			continue
		}
		parsedData, err := GetDataFromJSON(parsedJSON, job.Selector)
		if err != nil {
			log.Error("Error in fetching value from parsed data ", err)
			continue
		}
		datum, err := ConvertToNumber(parsedData)
		if err != nil {
			log.Error("Result is not a number")
			data = append(data, big.NewInt(0))
			continue
		}

		dataToAppend := MultiplyToEightDecimals(datum)
		data = append(data, dataToAppend)
	}
	if len(data) == 0 {
		data = append(data, big.NewInt(0))
	}
	return data
}
