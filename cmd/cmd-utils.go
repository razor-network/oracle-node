//This function add the following command to the root command
package cmd

import (
	"context"
	"errors"
	"math/big"
	"razor/RPC"
	"razor/accounts"
	"razor/core/types"
	"razor/logger"
	"razor/utils"
	"strconv"
	"time"

	"github.com/spf13/pflag"
)

//This function takes client as a parameter and returns the epoch and state
func (*UtilsStruct) GetEpochAndState(rpcParameter RPC.RPCParameters) (uint32, int64, error) {
	epoch, err := razorUtils.GetEpoch(rpcParameter)
	if err != nil {
		return 0, 0, err
	}
	bufferPercent, err := cmdUtils.GetBufferPercent()
	if err != nil {
		return 0, 0, err
	}
	err = ValidateBufferPercentLimit(rpcParameter, bufferPercent)
	if err != nil {
		return 0, 0, err
	}
	latestHeader, err := clientUtils.GetLatestBlockWithRetry(rpcParameter)
	if err != nil {
		log.Error("Error in fetching block: ", err)
		return 0, 0, err
	}
	stateBuffer, err := razorUtils.GetStateBuffer(rpcParameter)
	if err != nil {
		log.Error("Error in getting state buffer: ", err)
		return 0, 0, err
	}
	state, err := razorUtils.GetBufferedState(latestHeader, stateBuffer, bufferPercent)
	if err != nil {
		return 0, 0, err
	}
	log.Debug("Epoch ", epoch)
	log.Debug("State ", utils.GetStateName(state))
	return epoch, state, nil
}

//This function waits for the appropriate states which are required
func (*UtilsStruct) WaitForAppropriateState(rpcParameter RPC.RPCParameters, action string, states ...int) (uint32, error) {
	statesAllowed := GetFormattedStateNames(states)
	for {
		epoch, state, err := cmdUtils.GetEpochAndState(rpcParameter)
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
func (*UtilsStruct) WaitIfCommitState(rpcParameter RPC.RPCParameters, action string) (uint32, error) {
	for {
		epoch, state, err := cmdUtils.GetEpochAndState(rpcParameter)
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

func InitializeCommandDependencies(flagSet *pflag.FlagSet) (types.Configurations, RPC.RPCParameters, types.Account, error) {
	config, err := cmdUtils.GetConfigData()
	if err != nil {
		log.Error("Error in getting config: ", err)
		return types.Configurations{}, RPC.RPCParameters{}, types.Account{}, err
	}
	log.Debugf("Config: %+v", config)

	address, err := flagSetUtils.GetStringAddress(flagSet)
	if err != nil {
		log.Error("Error in getting address: ", err)
		return types.Configurations{}, RPC.RPCParameters{}, types.Account{}, err
	}
	log.Debugf("Address: %v", address)

	log.Debug("Checking to assign log file...")
	fileUtils.AssignLogFile(flagSet, config)

	log.Debug("Getting password...")
	password := razorUtils.AssignPassword(flagSet)

	accountManager, err := razorUtils.AccountManagerForKeystore()
	if err != nil {
		log.Error("Error in getting accounts manager for keystore: ", err)
		return types.Configurations{}, RPC.RPCParameters{}, types.Account{}, err
	}

	account := accounts.InitAccountStruct(address, password, accountManager)

	err = razorUtils.CheckPassword(account)
	if err != nil {
		log.Error("Error in fetching private key from given password: ", err)
		return types.Configurations{}, RPC.RPCParameters{}, types.Account{}, err
	}

	rpcManager, err := RPC.InitializeRPCManager(config.Provider)
	if err != nil {
		log.Error("Error in initializing RPC Manager: ", err)
		return types.Configurations{}, RPC.RPCParameters{}, types.Account{}, err
	}

	rpcParameters := RPC.RPCParameters{
		RPCManager: rpcManager,
		Ctx:        context.Background(),
	}

	client, err := rpcManager.GetBestRPCClient()
	if err != nil {
		log.Error("Error in getting best RPC client: ", err)
		return types.Configurations{}, RPC.RPCParameters{}, types.Account{}, err
	}

	logger.SetLoggerParameters(client, address)

	return config, rpcParameters, account, nil
}
