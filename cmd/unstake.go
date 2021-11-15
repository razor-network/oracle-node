package cmd

import (
	"errors"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"
	"time"

	"github.com/spf13/cobra"
)

var utilsStruct UtilsStruct

var unstakeCmd = &cobra.Command{
	Use:   "unstake",
	Short: "Unstake your razors",
	Long: `unstake allows user to unstake their sRzrs in the razor network

Example:	
  ./razor unstake --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --value 1000 --autoWithdraw
	`,
	Run: func(cmd *cobra.Command, args []string) {
		err := utilsStruct.initialiseUnstake(cmd)
		utils.CheckError("Error in initialising unstake function: ", err)
	},
}

//Run:  utilsStruct.initialiseUnstake(cmd *cobra.Command, args []string) {
//err := utilsStruct.initialiseUnstake(cmd)
//utils.CheckError("Error in initialising unstake function: ", err)
//},
//}

func (utilsStruct UtilsStruct) initialiseUnstake(cmd *cobra.Command) error {
	config, err := utilsStruct.razorUtils.GetConfigData()
	flagSet := cmd.Flags()
	if err != nil {
		log.Error("Error in getting config: ", err)
		return err
	}

	password := utilsStruct.razorUtils.AssignPassword(flagSet)
	address, err := utilsStruct.flagSetUtils.GetStringAddress(flagSet)
	if err != nil {
		log.Error("Error in getting address: ", err)
		return err
	}
	autoWithdraw, err := utilsStruct.flagSetUtils.GetBoolAutoWithdraw(flagSet)
	if err != nil {
		log.Error("Error in getting autoWithdraw status: ", err)
		return err
	}

	client := utilsStruct.razorUtils.ConnectToClient(config.Provider)

	valueInWei := utilsStruct.razorUtils.AssignAmountInWei(flagSet)

	utilsStruct.razorUtils.CheckEthBalanceIsZero(client, address)

	stakerId, err := utilsStruct.razorUtils.AssignStakerId(flagSet, client, address)
	if err != nil {
		log.Error("StakerId error: ", err)
		return err
	}

	lock, err := utilsStruct.razorUtils.GetLock(client, address, stakerId)
	if err != nil {
		log.Error("Error in getting lock: ", err)
		return err
	}

	if lock.Amount.Cmp(big.NewInt(0)) != 0 {
		err = errors.New("existing lock")
		log.Error(err)
		return err
	}

	txnOptions := types.TransactionOptions{
		Client:          client,
		Password:        password,
		AccountAddress:  address,
		Amount:          valueInWei,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.StakeManagerAddress,
		MethodName:      "unstake",
		ABI:             bindings.StakeManagerABI,
	}

	err = utilsStruct.cmdUtils.Unstake(txnOptions, stakerId, utilsStruct)
	if err != nil {
		log.Error("Unstake Error: ", err)
		return err
	}

	if autoWithdraw {
		AutoWithdraw(txnOptions, stakerId)
	}

	return nil
}

func Unstake(txnArgs types.TransactionOptions, stakerId uint32, utilsStruct UtilsStruct) error {
	lock, err := utilsStruct.razorUtils.GetLock(txnArgs.Client, txnArgs.AccountAddress, stakerId)
	if err != nil {
		log.Error("Error in getting lock: ", err)
		return err
	}
	if lock.Amount.Cmp(big.NewInt(0)) != 0 {
		err := errors.New("existing lock")
		log.Error(err)
		return err
	}

	epoch, err := utilsStruct.razorUtils.WaitForAppropriateState(txnArgs.Client, txnArgs.AccountAddress, "unstake", 0, 1, 4)
	if err != nil {
		log.Error("Error in fetching epoch: ", err)
		return err
	}
	txnArgs.Parameters = []interface{}{epoch, stakerId, txnArgs.Amount}
	txnOpts := utilsStruct.razorUtils.GetTxnOpts(txnArgs)
	log.Info("Unstaking coins")
	txn, err := utilsStruct.stakeManagerUtils.Unstake(txnArgs.Client, txnOpts, epoch, stakerId, txnArgs.Amount)
	if err != nil {
		log.Error("Error in un-staking: ", err)
		return err
	}
	log.Info("Transaction hash: ", utilsStruct.transactionUtils.Hash(txn))
	utilsStruct.razorUtils.WaitForBlockCompletion(txnArgs.Client, utilsStruct.transactionUtils.Hash(txn).String())
	return nil
}

func AutoWithdraw(txnArgs types.TransactionOptions, stakerId uint32) {
	log.Info("Starting withdrawal now...")
	time.Sleep(time.Duration(core.EpochLength) * time.Second)
	txn, err := withdrawFunds(txnArgs.Client, types.Account{
		Address:  txnArgs.AccountAddress,
		Password: txnArgs.Password,
	}, txnArgs.Config, stakerId, razorUtils, cmdUtils, stakeManagerUtils, transactionUtils)
	if err != nil {
		log.Error("WithdrawFunds error ", err)
	}
	if txn != core.NilHash {
		utils.WaitForBlockCompletion(txnArgs.Client, txn.String())
	}
}

func init() {
	rootCmd.AddCommand(unstakeCmd)

	var (
		Address               string
		AmountToUnStake       string
		WithdrawAutomatically bool
		Password              string
		Power                 string
		StakerId              uint32
	)

	unstakeCmd.Flags().StringVarP(&Address, "address", "a", "", "user's address")
	unstakeCmd.Flags().StringVarP(&AmountToUnStake, "value", "v", "0", "value of sRazors to un-stake")
	unstakeCmd.Flags().BoolVarP(&WithdrawAutomatically, "autoWithdraw", "", false, "withdraw after un-stake automatically")
	unstakeCmd.Flags().StringVarP(&Password, "password", "", "", "password path to protect the keystore")
	unstakeCmd.Flags().StringVarP(&Power, "pow", "", "", "power of 10")
	unstakeCmd.Flags().Uint32VarP(&StakerId, "stakerId", "", 0, "staker id")

	addrErr := unstakeCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
	valueErr := unstakeCmd.MarkFlagRequired("value")
	utils.CheckError("Value error: ", valueErr)

}
