//This function add the following command to the root command
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

//This function takes client as a parameter and returns the epoch and state
func (*UtilsStruct) GetEpochAndState(client *ethclient.Client) (uint32, int64, error) {
	epoch, err := razorUtils.GetEpoch(client)
	if err != nil {
		return 0, 0, err
	}
	bufferPercent, err := cmdUtils.GetBufferPercent()
	if err != nil {
		return 0, 0, err
	}
	latestHeader, err := clientUtils.GetLatestBlockWithRetry(client)
	if err != nil {
		log.Error("Error in fetching block: ", err)
		return 0, 0, err
	}
	state, err := razorUtils.GetBufferedState(client, latestHeader, bufferPercent)
	if err != nil {
		return 0, 0, err
	}
	log.Debug("Epoch ", epoch)
	log.Debug("State ", utils.GetStateName(state))
	return epoch, state, nil
}

//This function waits for the appropriate states which are required
func (*UtilsStruct) WaitForAppropriateState(client *ethclient.Client, action string, states ...int) (uint32, error) {
	statesAllowed := GetFormattedStateNames(states)
	for {
		epoch, state, err := cmdUtils.GetEpochAndState(client)
		if err != nil {
			log.Error("Error in fetching epoch and state: ", err)
			return epoch, err
		}
		if !utils.Contains(states, int(state)) {
			log.Infof("Can only %s during %s state(s). Retrying in 5 seconds...", action, statesAllowed)
			timeUtils.Sleep(5 * time.Second)
		} else {
			return epoch, nil
		}
	}
}

//This function wait if the state is commit state
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

//This function assignes amount in wei
func (*UtilsStruct) AssignAmountInWei(flagSet *pflag.FlagSet) (*big.Int, error) {
	amount, err := flagSetUtils.GetStringValue(flagSet)
	if err != nil {
		log.Error("Error in reading value: ", err)
		return nil, err
	}
	log.Debug("AssignAmountInWei: Amount: ", amount)
	_amount, ok := new(big.Int).SetString(amount, 10)

	if !ok {
		return nil, errors.New("SetString: error")
	}
	var amountInWei *big.Int
	if razorUtils.IsFlagPassed("weiRazor") {
		weiRazorPassed, err := flagSetUtils.GetBoolWeiRazor(flagSet)
		if err != nil {
			log.Error("Error in getting weiRazorBool Value: ", err)
			return nil, err
		}
		if weiRazorPassed {
			log.Debug("weiRazor flag is passed as true, considering teh value input in wei")
			amountInWei = _amount
		}
	} else {
		amountInWei = utils.GetAmountInWei(_amount)
	}
	return amountInWei, nil
}

//This function returns the states which are allowed
func GetFormattedStateNames(states []int) string {
	var statesAllowed string
	for i := 0; i < len(states); i++ {
		if i == len(states)-1 {
			statesAllowed = statesAllowed + strconv.Itoa(states[i]) + ":" + utils.GetStateName(int64(states[i]))
		} else {
			statesAllowed = statesAllowed + strconv.Itoa(states[i]) + ":" + utils.GetStateName(int64(states[i])) + ", "
		}
	}
	return statesAllowed
}
