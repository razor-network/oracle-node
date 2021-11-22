package cmd

import (
	"errors"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
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
	Run: initialiseUnstake,
}

func initialiseUnstake(cmd *cobra.Command, args []string) {
	utilsStruct := UtilsStruct{
		stakeManagerUtils: stakeManagerUtils,
		razorUtils:        razorUtils,
		transactionUtils:  transactionUtils,
		cmdUtils:          cmdUtils,
		flagSetUtils:      flagSetUtils,
	}
	utilsStruct.executeUnstake(cmd.Flags())
}

func (utilsStruct UtilsStruct) executeUnstake(flagSet *pflag.FlagSet) {
	config, err := utilsStruct.razorUtils.GetConfigData()
	utils.CheckError("Error in getting config: ", err)

	password := utilsStruct.razorUtils.AssignPassword(flagSet)
	address, err := utilsStruct.flagSetUtils.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)

	autoWithdraw, err := utilsStruct.flagSetUtils.GetBoolAutoWithdraw(flagSet)
	utils.CheckError("Error in getting autoWithdraw status: ", err)

	client := utilsStruct.razorUtils.ConnectToClient(config.Provider)

	valueInWei := utilsStruct.razorUtils.AssignAmountInWei(flagSet)

	utilsStruct.razorUtils.CheckEthBalanceIsZero(client, address)

	stakerId, err := utilsStruct.razorUtils.AssignStakerId(flagSet, client, address)
	utils.CheckError("StakerId error: ", err)

	lock, err := utilsStruct.razorUtils.GetLock(client, address, stakerId)
	utils.CheckError("Error in getting lock: ", err)

	if lock.Amount.Cmp(big.NewInt(0)) != 0 {
		err = errors.New("existing lock")
		log.Fatal(err)
	}

	unstakeInput := types.UnstakeInput{
		Address:    address,
		Password:   password,
		ValueInWei: valueInWei,
		StakerId:   stakerId,
	}

	txnOptions, err := utilsStruct.cmdUtils.Unstake(config, client, unstakeInput, utilsStruct)
	utils.CheckError("Unstake Error: ", err)

	if autoWithdraw {
		utilsStruct.cmdUtils.AutoWithdraw(txnOptions, stakerId, utilsStruct)
	}
}

func Unstake(config types.Configurations, client *ethclient.Client, input types.UnstakeInput, utilsStruct UtilsStruct) (types.TransactionOptions, error) {
	txnArgs := types.TransactionOptions{
		Client:          client,
		Password:        input.Password,
		AccountAddress:  input.Address,
		Amount:          input.ValueInWei,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.StakeManagerAddress,
		MethodName:      "unstake",
		ABI:             bindings.StakeManagerABI,
	}
	stakerId := input.StakerId
	lock, err := utilsStruct.razorUtils.GetLock(txnArgs.Client, txnArgs.AccountAddress, stakerId)
	if err != nil {
		log.Error("Error in getting lock: ", err)
		return txnArgs, err
	}
	if lock.Amount.Cmp(big.NewInt(0)) != 0 {
		err := errors.New("existing lock")
		log.Error(err)
		return txnArgs, err
	}

	epoch, err := utilsStruct.razorUtils.WaitForAppropriateState(txnArgs.Client, txnArgs.AccountAddress, "unstake", 0, 1, 4)
	if err != nil {
		log.Error("Error in fetching epoch: ", err)
		return txnArgs, err
	}
	txnArgs.Parameters = []interface{}{epoch, stakerId, txnArgs.Amount}
	txnOpts := utilsStruct.razorUtils.GetTxnOpts(txnArgs)
	log.Info("Unstaking coins")
	txn, err := utilsStruct.stakeManagerUtils.Unstake(txnArgs.Client, txnOpts, epoch, stakerId, txnArgs.Amount)
	if err != nil {
		log.Error("Error in un-staking: ", err)
		return txnArgs, err
	}
	log.Info("Transaction hash: ", utilsStruct.transactionUtils.Hash(txn))
	utilsStruct.razorUtils.WaitForBlockCompletion(txnArgs.Client, utilsStruct.transactionUtils.Hash(txn).String())
	return txnArgs, nil
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
		razorUtils.WaitForBlockCompletion(txnArgs.Client, txn.String())
	}
}

func init() {

	razorUtils = Utils{}
	transactionUtils = TransactionUtils{}
	stakeManagerUtils = StakeManagerUtils{}
	cmdUtils = UtilsCmd{}
	flagSetUtils = FlagSetUtils{}

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
