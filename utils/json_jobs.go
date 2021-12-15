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
		// If file is blank, do nothing
		if err.Error() == "unexpected end of JSON input" {
			return map[string]*types.StructsJob{}, nil
		}
		return nil, err
	}
	return data, nil
}

func writeDataToJSON(fileName string, data map[string]*types.StructsJob) error {
	jsonString, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fileName, jsonString, 0600)
	if err != nil {
		return err
	}
	return nil
}

func GetJobFromJSON(fileName string, jobId string) (*types.StructsJob, error) {
	data, err := readJSONData(fileName)
	if err != nil {
		return nil, err
	}
	if data[jobId] != nil {
		return data[jobId], nil
	}
	return nil, nil
}

func DeleteJobFromJSON(fileName string, jobId string) error {
	data, err := readJSONData(fileName)
	if err != nil {
		return err
	}
	if data[jobId] != nil {
		delete(data, jobId)
	}
	return writeDataToJSON(fileName, data)
}

func AddJobToJSON(fileName string, job *types.StructsJob) error {
	data, err := readJSONData(fileName)
	if err != nil {
		log.Error(err)
	}
	data[strconv.Itoa(int(job.Id))] = job
	return writeDataToJSON(fileName, data)
}
