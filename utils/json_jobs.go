package utils

import (
	"encoding/json"
	"io/ioutil"
	"razor/core/types"
	"strconv"
)

func readJSONData(fileName string) (map[string]*types.StructsJob, error) {
	var data = map[string]*types.StructsJob{}
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func writeDataToJSON(fileName string, data map[string]*types.StructsJob) (bool, error) {
	jsonString, err := json.Marshal(data)
	if err != nil {
		return false, err
	}
	err = ioutil.WriteFile(fileName, jsonString, 0600)
	if err != nil {
		return false, err
	}
	return true, nil
}

func getJobFromJSON(fileName string, jobId string) (*types.StructsJob, error) {
	data, err := readJSONData(fileName)
	if err != nil {
		return nil, err
	}
	if data[jobId] != nil {
		return data[jobId], nil
	}
	return nil, nil
}

func deleteJobFromJSON(fileName string, jobId string) (bool, error) {
	data, err := readJSONData(fileName)
	if err != nil {
		return false, err
	}
	if data[jobId] != nil {
		delete(data, jobId)
	}
	return writeDataToJSON(fileName, data)
}

func addJobToJSON(fileName string, job *types.StructsJob) (bool, error) {
	data, err := readJSONData(fileName)
	if err != nil {
		return false, err
	}
	data[strconv.Itoa(int(job.Id))] = job
	return writeDataToJSON(fileName, data)
}
