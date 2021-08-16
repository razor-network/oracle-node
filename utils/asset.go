package utils

import (
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
	"razor/core/types"
)

func GetNumAssets(client *ethclient.Client, address string) (*big.Int, error) {
	assetManager := GetAssetManager(client)
	callOpts := GetOptions(false, address, "")
	return assetManager.GetNumAssets(&callOpts)
}

func GetActiveAssetIds(client *ethclient.Client, address string) ([]*big.Int, error) {
	numAssets, err := GetNumAssets(client, address)
	if err != nil {
		return nil, err
	}
	assetManager := GetAssetManager(client)
	callOpts := GetOptions(false, address, "")
	var activeAssets []*big.Int
	for assetId := 1; assetId <= int(numAssets.Int64()); assetId++ {
		assetType, err := assetManager.GetAssetType(&callOpts, big.NewInt(int64(assetId)))
		if err != nil {
			log.Error("Error in fetching asset: ", assetId)
			continue
		}
		if assetType.Cmp(big.NewInt(1)) == 0 {
			activeJob, err := GetActiveJob(client, address, big.NewInt(int64(assetId)))
			if err != nil {
				log.Error(err)
				continue
			}
			activeAssets = append(activeAssets, activeJob.Id)
		} else {
			activeCollection, err := GetActiveCollection(client, address, big.NewInt(int64(assetId)))
			if err != nil {
				log.Error(err)
				continue
			}
			activeAssets = append(activeAssets, activeCollection.Id)
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

	for assetIndex := 1; assetIndex <= int(numOfAssets.Int64()); assetIndex++ {
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
			data = append(data, GetDataToCommitFromJob(activeJob))
		} else {
			activeCollection, err := GetActiveCollection(client, address, big.NewInt(int64(assetIndex)))
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
	if !collection.Active {
		return types.Collection{}, nil
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
		log.Error("Error in fetching value from parsed data ", err)
		return big.NewInt(1)
	}

	datum, err := ConvertToNumber(parsedData)
	if err != nil {
		log.Error("Result is not a number")
		return big.NewInt(1)
	}

	return MultiplyToEightDecimals(datum)
}
