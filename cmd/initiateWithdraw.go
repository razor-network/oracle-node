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

	txn, err := cmdUtils.HandleUnstakeLock(client, types.Account{
		Address:  address,
		Password: password,
	}, config, stakerId)

	utils.CheckError("InitiateWithdraw error: ", err)
	if txn != core.NilHash {
		err := razorUtils.WaitForBlockCompletion(client, txn.Hex())
		utils.CheckError("Error in WaitForBlockCompletion for initiateWithdraw: ", err)
	}

	if txn != core.NilHash && autoWithdraw {
		err = cmdUtils.AutoWithdraw(client, types.Account{
			Address:  address,
			Password: password,
		}, config, stakerId)
		utils.CheckError("AutoWithdraw Error: ", err)
	}
}

//This function handles the unstake lock
func (*UtilsStruct) HandleUnstakeLock(client *ethclient.Client, account types.Account, configurations types.Configurations, stakerId uint32) (common.Hash, error) {
	_, err := cmdUtils.WaitForAppropriateState(client, "initiateWithdraw", 0, 1, 4)
	if err != nil {
		log.Error("Error in fetching epoch: ", err)
		return core.NilHash, err
	}

	unstakeLock, err := razorUtils.GetLock(client, account.Address, stakerId, 0)
	if err != nil {
		log.Error("Error in fetching unstakeLock")
		return core.NilHash, err
	}

	if unstakeLock.UnlockAfter.Cmp(big.NewInt(0)) == 0 {
		log.Error("Unstake command not called before initiating withdrawal!")
		return core.NilHash, errors.New("unstake Razors before withdrawing")
	}

	withdrawInitiationPeriod, err := razorUtils.GetWithdrawInitiationPeriod(client)
	if err != nil {
		log.Error("Error in fetching withdraw release period")
		return core.NilHash, err
	}

	withdrawBefore := big.NewInt(0).Add(unstakeLock.UnlockAfter, big.NewInt(int64(withdrawInitiationPeriod)))
	epoch, err := razorUtils.GetEpoch(client)
	if err != nil {
		log.Error("Error in fetching epoch")
		return core.NilHash, err
	}

	if big.NewInt(int64(epoch)).Cmp(withdrawBefore) > 0 {
		log.Info("Withdraw initiation period has passed. Cannot withdraw now, please reset the unstakeLock!")
		return core.NilHash, nil
	}

	waitFor := big.NewInt(0).Sub(unstakeLock.UnlockAfter, big.NewInt(int64(epoch)))
	if waitFor.Cmp(big.NewInt(0)) > 0 {
		timeRemaining := uint64(waitFor.Int64()) * core.EpochLength
		log.Infof("Withdrawal Initiation period not reached. Cannot initiate withdraw now, please wait for %d epoch(s)! (approximately %s)", waitFor, razorUtils.SecondsToReadableTime(int(timeRemaining)))
		return core.NilHash, nil
	}

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
	txnOpts := razorUtils.GetTxnOpts(txnArgs)

	return cmdUtils.InitiateWithdraw(client, txnOpts, stakerId)
}

//This function initiate withdraw for your razors once you've unstaked
func (*UtilsStruct) InitiateWithdraw(client *ethclient.Client, txnOpts *bind.TransactOpts, stakerId uint32) (common.Hash, error) {
	log.Info("Initiating withdrawal of funds...")

	txn, err := stakeManagerUtils.InitiateWithdraw(client, txnOpts, stakerId)
	txnHash := transactionUtils.Hash(txn)
	if err != nil {
		log.Error("Error in initiating withdrawal of funds")
		return core.NilHash, err
	}

	log.Info("Txn Hash: ", txnHash.Hex())

	return txnHash, nil
}

//	This function helps the user to auto withdraw the razors after initiating withdraw
func (*UtilsStruct) AutoWithdraw(client *ethclient.Client, account types.Account, configurations types.Configurations, stakerId uint32) error {
	log.Info("Starting withdrawal...")
	withdrawLock, err := razorUtils.GetLock(client, account.Address, stakerId, 1)
	if err != nil {
		log.Error("Error in fetching withdrawLock")
		return err
	}
	epoch, state, err := cmdUtils.GetEpochAndState(client)
	if err != nil {
		log.Error("Error in fetching epoch")
		return err
	}

	waitFor := big.NewInt(0).Sub(withdrawLock.UnlockAfter, big.NewInt(int64(epoch)))
	timeRemaining := (uint64(waitFor.Int64()-1) * core.EpochLength) + (uint64(6-state) * core.EpochLength / 5) + 5
	log.Infof("Waiting for lock to get over... please wait for approximately %s", razorUtils.SecondsToReadableTime(int(timeRemaining)))

	timeUtils.Sleep((time.Duration(timeRemaining) * time.Second))
	log.Info("Lock period completed")

	txn, err := cmdUtils.HandleWithdrawLock(client, types.Account{
		Address:  account.Address,
		Password: account.Password,
	}, configurations, stakerId)
	utils.CheckError("UnlockWithdraw error: ", err)

	if txn != core.NilHash {
		err = razorUtils.WaitForBlockCompletion(client, txn.Hex())
		utils.CheckError("Error in WaitForBlockCompletion for unlockWithdraw: ", err)
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
