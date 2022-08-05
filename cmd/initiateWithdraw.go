//Package cmd provides all functions related to command line
package cmd

import (
	"errors"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/logger"
	"razor/pkg/bindings"
	"razor/utils"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var initiateWithdrawCmd = &cobra.Command{
	Use:   "initiateWithdraw",
	Short: "initiateWithdraw for your razors once you've unstaked",
	Long: `initiateWithdraw command can be used once the user has unstaked their token and the withdraw period is upon them.

Example:
  ./razor initiateWithdraw --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --stakerId 1
`,
	Run: func(cmd *cobra.Command, args []string) {
		cmdUtils.ExecuteInitiateWithdraw(cmd.Flags())
	},
}

//This function sets the flags appropriately and executes the InitiateWithdraw function
func (*UtilsStruct) ExecuteInitiateWithdraw(flagSet *pflag.FlagSet) {
	razorUtils.AssignLogFile(flagSet)
	address, err := flagSetUtils.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)

	logger.Address = address

	config, err := cmdUtils.GetConfigData()
	utils.CheckError("Error in getting config: ", err)

	password := razorUtils.AssignPassword()

	autoWithdraw, err := flagSetUtils.GetBoolAutoWithdraw(flagSet)
	utils.CheckError("Error in getting autoWithdraw status: ", err)

	client := razorUtils.ConnectToClient(config.Provider)

	razorUtils.CheckEthBalanceIsZero(client, address)

	stakerId, err := razorUtils.AssignStakerId(flagSet, client, address)
	utils.CheckError("Error in fetching stakerId:  ", err)

	txnOptions, err := cmdUtils.HandleUnstakeLock(client, types.Account{
		Address:  address,
		Password: password,
	}, config, stakerId)

	utils.CheckError("InitiateWithdraw error: ", err)
	if autoWithdraw {
		err = cmdUtils.AutoWithdraw(txnOptions, stakerId)
		utils.CheckError("AutoWithdraw Error: ", err)
	}
}

//This function handles the unstake lock
func (*UtilsStruct) HandleUnstakeLock(client *ethclient.Client, account types.Account, configurations types.Configurations, stakerId uint32) (types.TransactionOptions, error) {
	txnArgs := types.TransactionOptions{
		Client:          client,
		Password:        account.Password,
		AccountAddress:  account.Address,
		ChainId:         big.NewInt(configurations.ChainId),
		Config:          configurations,
		ContractAddress: core.StakeManagerAddress,
		MethodName:      "initiateWithdraw",
		ABI:             bindings.StakeManagerMetaData.ABI,
		Parameters:      []interface{}{stakerId},
	}
	_, err := cmdUtils.WaitForAppropriateState(txnArgs.Client, "initiateWithdraw", 0, 1, 4)
	if err != nil {
		log.Error("Error in fetching epoch: ", err)
		return txnArgs, err
	}

	unstakeLock, err := razorUtils.GetLock(txnArgs.Client, txnArgs.AccountAddress, stakerId, 0)
	if err != nil {
		log.Error("Error in fetching unstakeLock")
		return txnArgs, err
	}

	if unstakeLock.UnlockAfter.Cmp(big.NewInt(0)) == 0 {
		log.Error("Unstake command not called before initiating withdrawal!")
		return txnArgs, errors.New("unstake Razors before withdrawing")
	}

	withdrawInitiationPeriod, err := razorUtils.GetWithdrawInitiationPeriod(txnArgs.Client)
	if err != nil {
		log.Error("Error in fetching withdrawal Initiation period")
		return txnArgs, err
	}

	withdrawBefore := big.NewInt(0).Add(unstakeLock.UnlockAfter, big.NewInt(int64(withdrawInitiationPeriod)))
	epoch, err := razorUtils.GetEpoch(txnArgs.Client)
	if err != nil {
		log.Error("Error in fetching epoch")
		return txnArgs, err
	}

	if big.NewInt(int64(epoch)).Cmp(withdrawBefore) > 0 {
		log.Info("Withdraw initiation period has passed. Cannot withdraw now, please reset the unstakeLock!")
		return txnArgs, errors.New("withdrawal initiation period has passed")
	}

	waitFor := big.NewInt(0).Sub(unstakeLock.UnlockAfter, big.NewInt(int64(epoch)))
	if waitFor.Cmp(big.NewInt(0)) > 0 {
		timeRemaining := uint64(waitFor.Int64()) * core.EpochLength
		log.Infof("Withdrawal Initiation period not reached. Cannot initiate withdraw now, please wait for %d epoch(s)! (approximately %s)", waitFor, razorUtils.SecondsToReadableTime(int(timeRemaining)))
		return txnArgs, errors.New("withdrawal initiation period not reached")
	}

	txnOpts := razorUtils.GetTxnOpts(txnArgs)
	txn, err := cmdUtils.InitiateWithdraw(client, txnOpts, stakerId)
	if err != nil {
		log.Error("Error in initiating withdrawal of funds", err)
		return txnArgs, err
	}
	if txn != core.NilHash {
		err := razorUtils.WaitForBlockCompletion(client, txn.Hex())
		if err != nil {
			log.Error("Error in WaitForBlockCompletion for initiateWithdraw: ", err)
			return txnArgs, err
		}
	}
	return txnArgs, nil

}

//This function initiate withdraw for your razors once you've unstaked
func (*UtilsStruct) InitiateWithdraw(client *ethclient.Client, txnOpts *bind.TransactOpts, stakerId uint32) (common.Hash, error) {
	log.Info("Initiating withdrawal of funds...")

	txn, err := stakeManagerUtils.InitiateWithdraw(client, txnOpts, stakerId)
	txnHash := transactionUtils.Hash(txn)
	if err != nil {
		return core.NilHash, err
	}

	log.Info("Txn Hash: ", txnHash.Hex())

	return txnHash, nil
}

//	This function helps the user to auto withdraw the razors after initiating withdraw
func (*UtilsStruct) AutoWithdraw(txnArgs types.TransactionOptions, stakerId uint32) error {
	log.Info("Starting withdrawal now...")
	withdrawLock, err := razorUtils.GetLock(txnArgs.Client, txnArgs.AccountAddress, stakerId, 1)
	if err != nil {
		log.Error("Error in fetching withdrawLock")
		return err
	}
	epoch, err := razorUtils.GetEpoch(txnArgs.Client)
	if err != nil {
		log.Error("Error in fetching epoch")
		return err
	}

	waitFor := big.NewInt(0).Sub(withdrawLock.UnlockAfter, big.NewInt(int64(epoch)))
	timeRemaining := uint64(waitFor.Int64()) * core.EpochLength

	timeUtils.Sleep(time.Duration(timeRemaining) * time.Second)
	txn, err := cmdUtils.HandleWithdrawLock(txnArgs.Client, types.Account{
		Address:  txnArgs.AccountAddress,
		Password: txnArgs.Password,
	}, txnArgs.Config, stakerId)
	if err != nil {
		log.Error("HandleWithdrawLock error ", err)
		return err
	}
	if txn != core.NilHash {
		err = razorUtils.WaitForBlockCompletion(txnArgs.Client, txn.String())
		if err != nil {
			return err
		}
	}
	return nil
}

func init() {
	rootCmd.AddCommand(initiateWithdrawCmd)

	var (
		Address               string
		StakerId              uint32
		WithdrawAutomatically bool
	)

	initiateWithdrawCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the user")
	initiateWithdrawCmd.Flags().Uint32VarP(&StakerId, "stakerId", "", 0, "password path of user to protect the keystore")
	initiateWithdrawCmd.Flags().BoolVarP(&WithdrawAutomatically, "autoWithdraw", "", false, "withdraw after un-stake automatically")

	addrErr := initiateWithdrawCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
}
