package utils

import (
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
	"razor/core/types"
)

func GetActiveJobs(client *ethclient.Client, address string) ([]types.Job, error) {
	var jobs []types.Job
	jobManager := GetJobManager(client)
	callOpts := GetOptions(false, address, "")
	numOfJobs, err := jobManager.GetNumJobs(&callOpts)
	if err != nil {
		return jobs, err
	}
	epoch, err := GetEpoch(client, address)
	if err != nil {
		return jobs, err
	}
	for jobIndex := 1; jobIndex <= int(numOfJobs.Int64()); jobIndex++ {
		job, err := jobManager.Jobs(&callOpts, big.NewInt(int64(jobIndex)))
		if err != nil {
			log.Error("Error in fetching job", err)
		} else {
			if !job.Fulfilled && job.Epoch.Cmp(epoch) < 0 {
				jobs = append(jobs, job)
			}
		}
	}
	return jobs, nil
}