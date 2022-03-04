package cmd

import (
	"errors"
	"math/big"
	"razor/utils"
	"strconv"
	"time"

	"github.com/spf13/pflag"

	"github.com/ethereum/go-ethereum/ethclient"
)

func (*UtilsStruct) GetEpochAndState(client *ethclient.Client) (uint32, int64, error) {
	epoch, err := razorUtils.GetEpoch(client)
	if err != nil {
		return 0, 0, err
	}
	bufferPercent, err := cmdUtils.GetBufferPercent()
	if err != nil {
		return 0, 0, err
	}
	state, err := razorUtils.GetDelayedState(client, bufferPercent)
	if err != nil {
		return 0, 0, err
	}
	log.Debug("Epoch ", epoch)
	log.Debug("State ", razorUtils.GetStateName(state))
	return epoch, state, nil
}

func (*UtilsStruct) WaitForAppropriateState(client *ethclient.Client, action string, states ...int) (uint32, error) {
	var allowedStates string
	for i := 0; i < len(states); i++ {
		if i == len(states)-1 {
			allowedStates = allowedStates + strconv.Itoa(states[i]) + ":" + razorUtils.GetStateName(int64(states[i]))
		} else {
			allowedStates = allowedStates + strconv.Itoa(states[i]) + ":" + razorUtils.GetStateName(int64(states[i])) + ", "
		}
	}
	for {
		epoch, state, err := cmdUtils.GetEpochAndState(client)
		if err != nil {
			log.Error("Error in fetching epoch and state: ", err)
			return epoch, err
		}
		if !utils.Contains(states, int(state)) {
			log.Infof("Can only %s during %s state(s). Retrying in 5 seconds...", action, allowedStates)
			timeUtils.Sleep(5 * time.Second)
		} else {
			return epoch, nil
		}
	}
}

func (*UtilsStruct) WaitIfCommitState(client *ethclient.Client, action string) (uint32, error) {
	for {
		epoch, state, err := cmdUtils.GetEpochAndState(client)
		if err != nil {
			log.Error("Error in fetching epoch and state: ", err)
			return epoch, err
		}
		if state == 0 || state == -1 {
			log.Debugf("Cannot perform %s during commit state. Retrying in 5 seconds...", action)
			timeUtils.Sleep(5 * time.Second)
		} else {
			return epoch, nil
		}
	}
}

func (*UtilsStruct) AssignAmountInWei(flagSet *pflag.FlagSet) (*big.Int, error) {
	amount, err := flagSetUtils.GetStringValue(flagSet)
	if err != nil {
		log.Error("Error in reading value: ", err)
		return nil, err
	}
	_amount, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		return nil, errors.New("SetString: error")
	}
	var amountInWei *big.Int
	if razorUtils.IsFlagPassed("pow") {
		power, err := flagSetUtils.GetStringPow(flagSet)
		if err != nil {
			log.Error("Error in getting power: ", err)
			return nil, err
		}
		amountInWei, err = razorUtils.GetFractionalAmountInWei(_amount, power)
		if err != nil {
			return nil, err
		}
	} else {
		amountInWei = razorUtils.GetAmountInWei(_amount)
	}
	return amountInWei, nil
}
