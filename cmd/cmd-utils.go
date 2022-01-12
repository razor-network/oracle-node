package cmd

import (
	"errors"
	"math/big"
	"razor/utils"
	"time"

	"github.com/spf13/pflag"

	"github.com/ethereum/go-ethereum/ethclient"
)

func (*UtilsStructMockery) GetEpochAndState(client *ethclient.Client) (uint32, int64, error) {
	epoch, err := razorUtilsMockery.GetEpoch(client)
	if err != nil {
		return 0, 0, err
	}
	bufferPercent, err := cmdUtilsMockery.GetBufferPercent()
	if err != nil {
		return 0, 0, err
	}
	state, err := razorUtilsMockery.GetDelayedState(client, bufferPercent)
	if err != nil {
		return 0, 0, err
	}
	log.Debug("Epoch ", epoch)
	log.Debug("State ", razorUtilsMockery.GetStateName(state))
	return epoch, state, nil
}

func (*UtilsStructMockery) WaitForAppropriateState(client *ethclient.Client, action string, states ...int) (uint32, error) {
	for {
		epoch, state, err := cmdUtilsMockery.GetEpochAndState(client)
		if err != nil {
			log.Error("Error in fetching epoch and state: ", err)
			return epoch, err
		}
		if !utils.Contains(states, int(state)) {
			log.Debugf("Can only %s during %d state(s). Retrying in 5 seconds...", action, states)
			time.Sleep(5 * time.Second)
		} else {
			return epoch, nil
		}
	}
}

func (*UtilsStructMockery) WaitIfCommitState(client *ethclient.Client, action string) (uint32, error) {
	for {
		epoch, state, err := cmdUtilsMockery.GetEpochAndState(client)
		if err != nil {
			log.Error("Error in fetching epoch and state: ", err)
			return epoch, err
		}
		if state == 0 || state == -1 {
			log.Debugf("Cannot perform %s during commit state. Retrying in 5 seconds...", action)
			razorUtilsMockery.Sleep(5 * time.Second)
		} else {
			return epoch, nil
		}
	}
}

func (*UtilsStructMockery) AssignAmountInWei(flagSet *pflag.FlagSet) (*big.Int, error) {
	amount, err := flagSetUtilsMockery.GetStringValue(flagSet)
	if err != nil {
		log.Error("Error in reading value: ", err)
		return nil, err
	}
	_amount, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		return nil, errors.New("SetString: error")
	}
	var amountInWei *big.Int
	if razorUtilsMockery.IsFlagPassed("pow") {
		power, err := flagSetUtilsMockery.GetStringPow(flagSet)
		if err != nil {
			log.Error("Error in getting power: ", err)
			return nil, err
		}
		amountInWei, err = razorUtilsMockery.GetFractionalAmountInWei(_amount, power)
		if err != nil {
			return nil, err
		}
	} else {
		amountInWei = razorUtilsMockery.GetAmountInWei(_amount)
	}
	return amountInWei, nil
}
