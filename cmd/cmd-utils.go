package cmd

import (
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

func GetEpochAndState(client *ethclient.Client, accountAddress string, utilsStruct UtilsStruct) (uint32, int64, error) {
	epoch, err := utilsStruct.razorUtils.GetEpoch(client)
	if err != nil {
		return 0, 0, err
	}
	bufferPercent, err := utilsStruct.razorUtils.getBufferPercent()
	if err != nil {
		return 0, 0, err
	}
	state, err := utilsStruct.razorUtils.GetDelayedState(client, bufferPercent)
	if err != nil {
		return 0, 0, err
	}
	log.Debug("Epoch ", epoch)
	log.Debug("State ", utilsStruct.razorUtils.GetStateName(state))
	return epoch, state, nil
}

func WaitForAppropriateState(client *ethclient.Client, accountAddress string, action string, utilsStruct UtilsStruct, states ...int) (uint32, error) {
	for {
		epoch, state, err := utilsStruct.cmdUtils.GetEpochAndState(client, accountAddress, utilsStruct)
		if err != nil {
			log.Error("Error in fetching epoch and state: ", err)
			return epoch, err
		}
		if !utilsStruct.razorUtils.Contains(states, int(state)) {
			log.Debugf("Can only %s during %d state(s). Retrying in 5 seconds...", action, states)
			time.Sleep(5 * time.Second)
		} else {
			return epoch, nil
		}
	}
}

func WaitIfCommitState(client *ethclient.Client, accountAddress string, action string, utilsStruct UtilsStruct) (uint32, error) {
	for {
		epoch, state, err := utilsStruct.cmdUtils.GetEpochAndState(client, accountAddress, utilsStruct)
		if err != nil {
			log.Error("Error in fetching epoch and state: ", err)
			return epoch, err
		}
		if state == 0 || state == -1 {
			log.Debugf("Cannot perform %s during commit state. Retrying in 5 seconds...", action)
			time.Sleep(5 * time.Second)
		} else {
			return epoch, nil
		}
	}
}
