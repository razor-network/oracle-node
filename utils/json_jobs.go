package utils

import (
	"errors"
	"razor/core/types"
	"strconv"
)

func (*UtilsStruct) ReadJSONData(fileName string) (map[string]*types.StructsJob, error) {
	var data = map[string]*types.StructsJob{}
	file, err := OS.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	err = JsonInterface.Unmarshal(file, &data)
	if err != nil {
		// If file is blank, do nothing
		if err.Error() == "unexpected end of JSON input" {
			return map[string]*types.StructsJob{}, nil
		}
		return nil, err
	}
	return data, nil
}

func (*UtilsStruct) WriteDataToJSON(fileName string, data map[string]*types.StructsJob) error {
	jsonString, err := JsonInterface.Marshal(data)
	if err != nil {
		return err
	}
	err = OS.WriteFile(fileName, jsonString, 0600)
	if err != nil {
		return err
	}
	return nil
}

func (*UtilsStruct) DeleteJobFromJSON(fileName string, jobId string) error {
	data, err := UtilsInterface.ReadJSONData(fileName)
	if err != nil {
		return err
	}
	if data[jobId] != nil {
		delete(data, jobId)
	} else {
		return errors.New("No job with jobId = " + jobId + " found")
	}
	return UtilsInterface.WriteDataToJSON(fileName, data)
}

func (*UtilsStruct) AddJobToJSON(fileName string, job *types.StructsJob) error {
	data, err := UtilsInterface.ReadJSONData(fileName)
	if err != nil {
		log.Error(err)
		return err
	}
	data[strconv.Itoa(int(job.Id))] = job
	return UtilsInterface.WriteDataToJSON(fileName, data)
}
