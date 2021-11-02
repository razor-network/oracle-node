package cmd

import (
	"razor/utils"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

func GetEpochAndState(client *ethclient.Client, accountAddress string) (uint32, int64, error) {
	epoch, err := utils.GetEpoch(client)
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
	log.Debug("State ", utils.GetStateName(state))
	return epoch, state, nil
}

func WaitForAppropriateState(client *ethclient.Client, accountAddress string, action string, states ...int) (uint32, error) {
	for {
		epoch, state, err := GetEpochAndState(client, accountAddress)
		utils.CheckError("Error in fetching epoch and state: ", err)
		if !utils.Contains(states, int(state)) {
			log.Debugf("Can only %s during confirm state. Retrying in 5 seconds...", action)
			time.Sleep(5 * time.Second)
		} else {
			return epoch, nil
		}
	}
}
