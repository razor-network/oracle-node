package cmd

import (
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"
	"time"

	"github.com/spf13/cobra"
)

var unstakeCmd = &cobra.Command{
	Use:   "unstake",
	Short: "Unstake your razors",
	Long: `unstake allows user to unstake their sRzrs in the razor network

Example:	
  ./razor unstake --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --value 1000 --autoWithdraw
	`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		utils.CheckError("Error in getting config: ", err)

		password := utils.AssignPassword(cmd.Flags())
		address, _ := cmd.Flags().GetString("address")
		autoWithdraw, _ := cmd.Flags().GetBool("autoWithdraw")

		client := utils.ConnectToClient(config.Provider)

		valueInWei, err := AssignAmountInWei(cmd.Flags())
		utils.CheckError("Error in getting amount: ", err)

		utils.CheckEthBalanceIsZero(client, address)

		stakerId, err := utils.AssignStakerId(cmd.Flags(), client, address)
		utils.CheckError("StakerId error: ", err)

		lock, err := utils.GetLock(client, address, stakerId)
		utils.CheckError("Error in getting lock: ", err)
		if lock.Amount.Cmp(big.NewInt(0)) != 0 {
			log.Fatal("Existing lock")
		}

		utilsStruct := UtilsStruct{
			stakeManagerUtils: stakeManagerUtils,
			transactionUtils:  transactionUtils,
			cmdUtils:          cmdUtils,
			razorUtils:        razorUtils,
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

		Unstake(txnOptions, stakerId)

		if autoWithdraw {
			AutoWithdraw(txnOptions, stakerId, utilsStruct)
		}
	},
}

func Unstake(txnArgs types.TransactionOptions, stakerId uint32) {
	lock, err := utils.GetLock(txnArgs.Client, txnArgs.AccountAddress, stakerId)
	utils.CheckError("Error in getting lock: ", err)
	if lock.Amount.Cmp(big.NewInt(0)) != 0 {
		log.Fatal("Existing lock")
	}

	stakeManager := utils.GetStakeManager(txnArgs.Client)

	epoch, err := WaitForAppropriateState(txnArgs.Client, txnArgs.AccountAddress, "unstake", 0, 1, 4)
	txnArgs.Parameters = []interface{}{epoch, stakerId, txnArgs.Amount}
	txnOpts := GetTxnOpts(txnArgs)
	utils.CheckError("Error in fetching epoch: ", err)
	log.Info("Unstaking coins")
	txn, err := stakeManager.Unstake(txnOpts, epoch, stakerId, txnArgs.Amount)
	utils.CheckError("Error in un-staking: ", err)
	log.Info("Transaction hash: ", txn.Hash())
	utils.WaitForBlockCompletion(txnArgs.Client, txn.Hash().String())
}

func AutoWithdraw(txnArgs types.TransactionOptions, stakerId uint32, utilsStruct UtilsStruct) {
	log.Info("Starting withdrawal now...")
	time.Sleep(time.Duration(core.EpochLength) * time.Second)
	txn, err := utilsStruct.withdrawFunds(txnArgs.Client, types.Account{
		Address:  txnArgs.AccountAddress,
		Password: txnArgs.Password,
	}, txnArgs.Config, stakerId)
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
