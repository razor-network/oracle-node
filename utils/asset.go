package utils

import (
	"encoding/json"
	"errors"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"regexp"

	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getAssetManagerWithOpts(client *ethclient.Client) (*bindings.AssetManager, bind.CallOpts) {
	return GetAssetManager(client), GetOptions()
}

func GetNumAssets(client *ethclient.Client) (uint16, error) {
	assetManager, callOpts := getAssetManagerWithOpts(client)
	var (
		numAssets uint16
		err       error
	)
	err = retry.Do(
		func() error {
			numAssets, err = assetManager.GetNumAssets(&callOpts)
			if err != nil {
				log.Error("Error in fetching numAssets.... Retrying")
				return err
			}
			return nil
		}, retry.Attempts(core.MaxRetries))
	if err != nil {
		return 0, err
	}
	return numAssets, nil
}

func GetNumActiveAssets(client *ethclient.Client) (*big.Int, error) {
	assetManager, callOpts := getAssetManagerWithOpts(client)
	var (
		numActiveAssets *big.Int
		err             error
	)
	err = retry.Do(
		func() error {
			numActiveAssets, err = assetManager.GetNumActiveCollections(&callOpts)
			if err != nil {
				log.Error("Error in fetching active assets.... Retrying")
				return err
			}
			return nil
		}, retry.Attempts(core.MaxRetries))
	if err != nil {
		return nil, err
	}
	return numActiveAssets, nil
}

func GetAssetType(client *ethclient.Client, assetId uint16) (uint8, error) {
	assetManager, callOpts := getAssetManagerWithOpts(client)
	var (
		activeAsset types.Asset
		err         error
	)
	err = retry.Do(
		func() error {
			activeAsset, err = assetManager.GetAsset(&callOpts, assetId)
			if err != nil {
				log.Error("Error in fetching asset.... Retrying")
				return err
			}
			return nil
		}, retry.Attempts(core.MaxRetries))
	if err != nil {
		return 0, err
	}
	if activeAsset.Job.Id == 0 {
		return 2, nil
	}
	return 1, nil
}

func GetCollection(client *ethclient.Client, collectionId uint16) (bindings.StructsCollection, error) {
	assetManager, callOpts := getAssetManagerWithOpts(client)
	var (
		asset types.Asset
		err   error
	)
	err = retry.Do(
		func() error {
			asset, err = assetManager.GetAsset(&callOpts, collectionId)
			if err != nil {
				log.Error("Error in fetching collection.... Retrying")
				return err
			}
			return nil
		}, retry.Attempts(core.MaxRetries))
	if err != nil {
		return bindings.StructsCollection{}, err
	}
	return asset.Collection, nil
}

func GetActiveAssetIds(client *ethclient.Client) ([]uint16, error) {
	assetManager, callOpts := getAssetManagerWithOpts(client)
	var (
		activeAssetIds []uint16
		err            error
	)
	err = retry.Do(
		func() error {
			activeAssetIds, err = assetManager.GetActiveCollections(&callOpts)
			if err != nil {
				log.Error("Error in fetching active assets.... Retrying")
				return err
			}
			return nil
		}, retry.Attempts(core.MaxRetries))
	if err != nil {
		return nil, err
	}
	return activeAssetIds, nil
}

func GetActiveAssetsData(client *ethclient.Client, epoch uint32) ([]*big.Int, error) {
	var data []*big.Int

	numOfAssets, err := GetNumAssets(client)
	if err != nil {
		return data, err
	}

	for assetIndex := 1; assetIndex <= int(numOfAssets); assetIndex++ {
		assetType, err := GetAssetType(client, uint16(assetIndex))
		if err != nil {
			log.Error("Error in fetching asset type: ", assetType)
			return nil, err
		}
		if assetType == 2 {
			activeCollection, err := GetActiveCollection(client, uint16(assetIndex))
			if err != nil {
				log.Error(err)
				if err == errors.New("collection inactive") {
					continue
				}
				return nil, err
			}
			//Supply previous epoch to Aggregate in case if last reported value is required.
			collectionData, aggregationError := Aggregate(client, epoch-1, activeCollection)
			if aggregationError != nil {
				return nil, aggregationError
			}
			data = append(data, collectionData)
		}
	}
	return data, nil
}

func Aggregate(client *ethclient.Client, previousEpoch uint32, collection bindings.StructsCollection) (*big.Int, error) {
	if len(collection.JobIDs) == 0 {
		return nil, errors.New("no jobs present in the collection")
	}
	var jobs []bindings.StructsJob
	for _, id := range collection.JobIDs {
		job, err := GetActiveJob(client, id)
		if err != nil {
			log.Errorf("Error in fetching job %d: %s", id, err)
			continue
		}
		jobs = append(jobs, job)
	}
	dataToCommit, weight, err := GetDataToCommitFromJobs(jobs)
	if err != nil || len(dataToCommit) == 0 {
		prevCommitmentData, err := FetchPreviousValue(client, previousEpoch, collection.Id)
		if err != nil {
			return nil, err
		}
		return big.NewInt(int64(prevCommitmentData)), nil
	}
	return performAggregation(dataToCommit, weight, collection.AggregationMethod)
}

func GetActiveJob(client *ethclient.Client, jobId uint16) (bindings.StructsJob, error) {
	assetManager := GetAssetManager(client)
	callOpts := GetOptions()
	var (
		job bindings.StructsJob
		err error
	)
	err = retry.Do(
		func() error {
			job, err = assetManager.Jobs(&callOpts, jobId)
			if err != nil {
				log.Errorf("Error in fetching job %d.... Retrying", jobId)
				return err
			}
			return nil
		}, retry.Attempts(core.MaxRetries))
	if err != nil {
		return bindings.StructsJob{}, err
	}
	return job, nil
}

func GetActiveCollection(client *ethclient.Client, collectionId uint16) (bindings.StructsCollection, error) {
	collection, err := GetCollection(client, collectionId)
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
	var parsedData interface{}
	if job.SelectorType == 0 {
		apiErr = retry.Do(
			func() error {
				response, apiErr = GetDataFromAPI(job.Url)
				if apiErr != nil {
					log.Error("Error in fetching data from API: ", apiErr)
					return apiErr
				}
				return nil
			}, retry.Attempts(core.MaxRetries))

		err := json.Unmarshal(response, &parsedJSON)
		if err != nil {
			log.Error("Error in parsing data from API: ", err)
			return nil, err
		}
		parsedData, err = GetDataFromJSON(parsedJSON, job.Selector)
		if err != nil {
			log.Error("Error in fetching value from parsed data: ", err)
			return nil, err
		}
	} else {
		//TODO: Add retry here.
		dataPoint, err := GetDataFromHTML(job.Url, job.Selector)
		if err != nil {
			log.Error("Error in fetching value from parsed XHTML: ", err)
			return nil, err
		}
		// remove "," and currency symbols
		parsedData = regexp.MustCompile(`[\p{Sc},]`).ReplaceAllString(dataPoint, "")
	}

	datum, err := ConvertToNumber(parsedData)
	if err != nil {
		log.Error("Result is not a number")
		return nil, err
	}

	return MultiplyWithPower(datum, job.Power), err
}
