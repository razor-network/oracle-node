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

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"

	"github.com/spf13/cobra"
)

var unstakeCmd = &cobra.Command{
	Use:   "unstake",
	Short: "Unstake your razors",
	Long: `unstake allows user to unstake their sRzrs in the razor network

Example:	
  ./razor unstake --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --value 1000
	`,
	Run: initialiseUnstake,
}

//This function initialises the ExecuteUnstake function
func initialiseUnstake(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteUnstake(cmd.Flags())
}

//This function sets the flag appropriately and executes the Unstake function
func (*UtilsStruct) ExecuteUnstake(flagSet *pflag.FlagSet) {
	razorUtils.AssignLogFile(flagSet)
	address, err := flagSetUtils.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)

	logger.Address = address

	config, err := cmdUtils.GetConfigData()
	utils.CheckError("Error in getting config: ", err)

	password := razorUtils.AssignPassword(flagSet)

	client := razorUtils.ConnectToClient(config.Provider)

	valueInWei, err := cmdUtils.AssignAmountInWei(flagSet)
	utils.CheckError("Error in getting amountInWei: ", err)

	razorUtils.CheckEthBalanceIsZero(client, address)

	stakerId, err := razorUtils.AssignStakerId(flagSet, client, address)
	utils.CheckError("StakerId error: ", err)

	unstakeInput := types.UnstakeInput{
		Address:    address,
		Password:   password,
		ValueInWei: valueInWei,
		StakerId:   stakerId,
	}

	txnHash, err := cmdUtils.Unstake(config, client, unstakeInput)
	utils.CheckError("Unstake Error: ", err)
	if txnHash != core.NilHash {
		razorUtils.WaitForBlockCompletion(client, txnHash.String())
	}
}

//This function allows user to unstake their sRZRs in the razor network
func (*UtilsStruct) Unstake(config types.Configurations, client *ethclient.Client, input types.UnstakeInput) (common.Hash, error) {
	txnArgs := types.TransactionOptions{
		Client:         client,
		Password:       input.Password,
		AccountAddress: input.Address,
		Amount:         input.ValueInWei,
		ChainId:        core.ChainId,
		Config:         config,
	}
	stakerId := input.StakerId
	staker, err := razorUtils.GetStaker(client, stakerId)
	if err != nil {
		log.Error("Error in getting staker: ", err)
		return core.NilHash, err
	}
	approveHash, err := cmdUtils.ApproveUnstake(client, staker, txnArgs)
	if err != nil {
		return core.NilHash, err
	}

	if approveHash != core.NilHash {
		razorUtils.WaitForBlockCompletion(client, approveHash.String())
	}

	log.Info("Approved for unstake!")

	txnArgs.ContractAddress = core.StakeManagerAddress
	txnArgs.MethodName = "unstake"
	txnArgs.ABI = bindings.StakeManagerABI

	unstakeLock, err := razorUtils.GetLock(txnArgs.Client, txnArgs.AccountAddress, stakerId, 0)
	if err != nil {
		log.Error("Error in getting unstakeLock: ", err)
		return core.NilHash, err
	}

	if unstakeLock.Amount.Cmp(big.NewInt(0)) != 0 {
		err := errors.New("existing unstake lock")
		log.Error(err)
		return core.NilHash, err
	}

	_, err = cmdUtils.WaitForAppropriateState(txnArgs.Client, "unstake", 4)
	if err != nil {
		log.Error("Error in fetching epoch: ", err)
		return core.NilHash, err
	}
	txnArgs.Parameters = []interface{}{stakerId, txnArgs.Amount}
	txnOpts := razorUtils.GetTxnOpts(txnArgs)
	log.Info("Unstaking coins")
	txn, err := stakeManagerUtils.Unstake(txnArgs.Client, txnOpts, stakerId, txnArgs.Amount)
	if err != nil {
		log.Error("Error in un-staking: ", err)
		return core.NilHash, err
	}
	log.Info("Transaction hash: ", transactionUtils.Hash(txn))
	return transactionUtils.Hash(txn), nil
}

//This function approves the unstake
func (*UtilsStruct) ApproveUnstake(client *ethclient.Client, staker bindings.StructsStaker, txnArgs types.TransactionOptions) (common.Hash, error) {
	txnOpts := razorUtils.GetTxnOpts(txnArgs)
	log.Infof("Approving %d amount for unstake...", txnArgs.Amount)
	txn, err := stakeManagerUtils.ApproveUnstake(client, txnOpts, staker, txnArgs.Amount)
	if err != nil {
		log.Error("Error in approving for unstake")
		return core.NilHash, err
	}
	log.Info("Transaction Hash: ", transactionUtils.Hash(txn).String())
	return transactionUtils.Hash(txn), nil
}

//This function helps the user to auto withdraw the razors after unstaking
func (*UtilsStruct) AutoWithdraw(txnArgs types.TransactionOptions, stakerId uint32) error {
	log.Info("Starting withdrawal now...")
	timeUtils.Sleep(time.Duration(core.EpochLength) * time.Second)
	txn, err := cmdUtils.HandleUnstakeLock(txnArgs.Client, types.Account{
		Address:  txnArgs.AccountAddress,
		Password: txnArgs.Password,
	}, txnArgs.Config, stakerId)
	if err != nil {
		log.Error("HandleUnstakeLock error ", err)
		return err
	}
	if txn != core.NilHash {
		razorUtils.WaitForBlockCompletion(txnArgs.Client, txn.String())
	}
	return nil
}

//This function add the following command to the root command
func init() {
	rootCmd.AddCommand(unstakeCmd)

	var (
		Address         string
		AmountToUnStake string
		Password        string
		Power           string
		StakerId        uint32
	)

	unstakeCmd.Flags().StringVarP(&Address, "address", "a", "", "user's address")
	unstakeCmd.Flags().StringVarP(&AmountToUnStake, "value", "v", "0", "value of sRazors to un-stake")
	unstakeCmd.Flags().StringVarP(&Password, "password", "", "", "password path to protect the keystore")
	unstakeCmd.Flags().StringVarP(&Power, "pow", "", "", "power of 10")
	unstakeCmd.Flags().Uint32VarP(&StakerId, "stakerId", "", 0, "staker id")

	addrErr := unstakeCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
	valueErr := unstakeCmd.MarkFlagRequired("value")
	utils.CheckError("Value error: ", valueErr)

}
