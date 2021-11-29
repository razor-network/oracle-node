package cmd

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/pflag"
	"math/big"
	"razor/core/types"
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
			log.Debugf("Can only %s during %d state(s). Retrying in 5 seconds...", action, states)
			time.Sleep(5 * time.Second)
		} else {
			return epoch, nil
		}
	}
}

func WaitIfCommitState(client *ethclient.Client, accountAddress string, action string) (uint32, error) {
	for {
		epoch, state, err := GetEpochAndState(client, accountAddress)
		utils.CheckError("Error in fetching epoch and state: ", err)
		if state == 0 || state == -1 {
			log.Debugf("Cannot perform %s during commit state. Retrying in 5 seconds...", action)
			time.Sleep(5 * time.Second)
		} else {
			return epoch, nil
		}
	}
}

func AssignAmountInWei(flagSet *pflag.FlagSet) (*big.Int, error) {
	amount, err := flagSet.GetString("value")
	if err != nil {
		log.Error("Error in reading value: ", err)
		return nil, err
	}
	_amount, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		return nil, errors.New("SetString: error")
	}
	var amountInWei *big.Int
	if utils.IsFlagPassed("pow") {
		power, err := flagSet.GetString("pow")
		if err != nil {
			log.Error("Error in getting power: ", err)
			return nil, err
		}
		amountInWei, err = utils.GetFractionalAmountInWei(_amount, power)
		if err != nil {
			return nil, err
		}
	} else {
		amountInWei = utils.GetAmountInWei(_amount)
	}
	return amountInWei, nil
}

func GetTxnOpts(transactionData types.TransactionOptions, utilsStruct UtilsStruct) *bind.TransactOpts {
	defaultPath, err := utilsStruct.razorUtils.GetDefaultPath()
	utils.CheckError("Error in fetching default path: ", err)
	privateKey := utilsStruct.cmdUtils.GetPrivateKey(transactionData.AccountAddress, transactionData.Password, defaultPath, utilsStruct)
	if privateKey == nil {
		utils.CheckError("Error in fetching private key: ", errors.New(transactionData.AccountAddress+" not present in razor-go"))
	}
	nonce, err := utilsStruct.razorUtils.PendingNonceAt(context.Background(), common.HexToAddress(transactionData.AccountAddress), transactionData)
	utils.CheckError("Error in fetching pending nonce: ", err)

	gasPrice := utilsStruct.razorUtils.GetGasPrice(transactionData.Client, transactionData.Config)

	txnOpts, err := utilsStruct.razorUtils.NewKeyedTransactorWithChainID(privateKey, transactionData.ChainId)
	utils.CheckError("Error in getting transactor: ", err)
	txnOpts.Nonce = big.NewInt(int64(nonce))
	txnOpts.GasPrice = gasPrice
	txnOpts.Value = transactionData.EtherValue

	gasLimit, err := utilsStruct.razorUtils.GetGasLimit(transactionData, txnOpts)
	if err != nil {
		log.Error("Error in getting gas limit: ", err)
	}
	log.Debug("Gas after increment: ", gasLimit)
	txnOpts.GasLimit = gasLimit
	return txnOpts
}
