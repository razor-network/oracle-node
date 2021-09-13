package cmd

import (
	"razor/utils"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

func GetEpochAndState(client *ethclient.Client, accountAddress string) (uint32, int64, error) {
	epoch, err := utils.GetEpoch(client, accountAddress)
	if err != nil {
		return 0, 0, err
	}
	bufferPercent, err := getBufferPercent()
	if err != nil {
		return 0, 0, err
	}
	state, err := utils.GetDelayedState(client, bufferPercent)
	if err != nil {
		return 0, 0, err
	}
	log.Debug("Epoch ", epoch)
	log.Debug("State ", state)
	return epoch, state, nil
}

func WaitForCommitState(client *ethclient.Client, accountAddress string, action string) (uint32, error) {
	for {
		epoch, state, err := GetEpochAndState(client, accountAddress)
		utils.CheckError("Error in fetching epoch and state: ", err)
		if state != 0 {
			log.Debugf("Can only %s during state 0 (commit). Retrying in 5 second...", action)
			time.Sleep(5 * time.Second)
		} else {
			return epoch, nil
		}
	}
}

func WaitForDisputeOrConfirmState(client *ethclient.Client, accountAddress string, action string) (uint32, error) {
	for {
		epoch, state, err := GetEpochAndState(client, accountAddress)
		utils.CheckError("Error in fetching epoch and state: ", err)
		if state != 3 && state != 4 {
			log.Debugf("Can only %s during dispute or confirm state. Retrying in 5 seconds...", action)
			time.Sleep(5 * time.Second)
		} else {
			return epoch, nil
		}
	}
}
