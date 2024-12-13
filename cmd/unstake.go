//Package cmd provides all functions related to command line
package cmd

import (
	"errors"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/rpc"
	"razor/utils"

	"github.com/ethereum/go-ethereum/common"

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
	config, rpcParameters, account, err := InitializeCommandDependencies(flagSet)
	utils.CheckError("Error in initialising command dependencies: ", err)

	log.Debug("Getting amount in wei...")
	valueInWei, err := cmdUtils.AssignAmountInWei(flagSet)
	utils.CheckError("Error in getting amountInWei: ", err)

	stakerId, err := razorUtils.AssignStakerId(rpcParameters, flagSet, account.Address)
	utils.CheckError("StakerId error: ", err)

	unstakeInput := types.UnstakeInput{
		ValueInWei: valueInWei,
		StakerId:   stakerId,
		Account:    account,
	}

	txnHash, err := cmdUtils.Unstake(rpcParameters, config, unstakeInput)
	utils.CheckError("Unstake Error: ", err)
	if txnHash != core.NilHash {
		err = razorUtils.WaitForBlockCompletion(rpcParameters, txnHash.Hex())
		utils.CheckError("Error in WaitForBlockCompletion for unstake: ", err)
	}
}

//This function allows user to unstake their sRZRs in the razor network
func (*UtilsStruct) Unstake(rpcParameters rpc.RPCParameters, config types.Configurations, input types.UnstakeInput) (common.Hash, error) {
	txnArgs := types.TransactionOptions{
		Amount:  input.ValueInWei,
		ChainId: core.ChainId,
		Config:  config,
		Account: input.Account,
	}
	stakerId := input.StakerId
	staker, err := razorUtils.GetStaker(rpcParameters, stakerId)
	if err != nil {
		log.Error("Error in getting staker: ", err)
		return core.NilHash, err
	}

	log.Debugf("Unstake: Staker info: %+v", staker)
	log.Debug("Unstake: Calling ApproveUnstake()...")
	approveHash, err := cmdUtils.ApproveUnstake(rpcParameters, staker.TokenAddress, txnArgs)
	if err != nil {
		return core.NilHash, err
	}

	if approveHash != core.NilHash {
		err = razorUtils.WaitForBlockCompletion(rpcParameters, approveHash.Hex())
		if err != nil {
			return core.NilHash, err
		}
	}

	log.Info("Approved for unstake!")

	txnArgs.ContractAddress = core.StakeManagerAddress
	txnArgs.MethodName = "unstake"
	txnArgs.ABI = bindings.StakeManagerMetaData.ABI

	unstakeLock, err := razorUtils.GetLock(rpcParameters, txnArgs.Account.Address, stakerId, 0)
	if err != nil {
		log.Error("Error in getting unstakeLock: ", err)
		return core.NilHash, err
	}
	log.Debugf("Unstake: Unstake lock: %+v", unstakeLock)

	if unstakeLock.Amount.Cmp(big.NewInt(0)) != 0 {
		err := errors.New("existing unstake lock")
		log.Error(err)
		return core.NilHash, err
	}

	txnArgs.Parameters = []interface{}{stakerId, txnArgs.Amount}
	txnOpts, err := razorUtils.GetTxnOpts(rpcParameters, txnArgs)
	if err != nil {
		return core.NilHash, err
	}
	log.Info("Unstaking coins")
	client, err := rpcParameters.RPCManager.GetBestRPCClient()
	if err != nil {
		return core.NilHash, err
	}

	log.Debugf("Executing Unstake transaction with stakerId = %d, amount = %s", stakerId, txnArgs.Amount)
	txn, err := stakeManagerUtils.Unstake(client, txnOpts, stakerId, txnArgs.Amount)
	if err != nil {
		log.Error("Error in un-staking: ", err)
		return core.NilHash, err
	}
	txnHash := transactionUtils.Hash(txn)
	log.Info("Txn Hash: ", txnHash.Hex())
	return txnHash, nil
}

//This function approves the unstake
func (*UtilsStruct) ApproveUnstake(rpcParameters rpc.RPCParameters, stakerTokenAddress common.Address, txnArgs types.TransactionOptions) (common.Hash, error) {
	txnOpts, err := razorUtils.GetTxnOpts(rpcParameters, txnArgs)
	if err != nil {
		return core.NilHash, err
	}
	client, err := rpcParameters.RPCManager.GetBestRPCClient()
	if err != nil {
		return core.NilHash, err
	}

	log.Infof("Approving %d amount for unstake...", txnArgs.Amount)
	txn, err := stakeManagerUtils.ApproveUnstake(client, txnOpts, stakerTokenAddress, txnArgs.Amount)
	if err != nil {
		log.Error("Error in approving for unstake")
		return core.NilHash, err
	}
	txnHash := transactionUtils.Hash(txn)
	log.Info("Txn Hash: ", txnHash.Hex())
	return txnHash, nil
}

func init() {
	rootCmd.AddCommand(unstakeCmd)

	var (
		Address         string
		AmountToUnStake string
		Password        string
		WeiRazor        bool
		StakerId        uint32
	)

	unstakeCmd.Flags().StringVarP(&Address, "address", "a", "", "user's address")
	unstakeCmd.Flags().StringVarP(&AmountToUnStake, "value", "v", "0", "value of sRazors to un-stake")
	unstakeCmd.Flags().StringVarP(&Password, "password", "", "", "password path to protect the keystore")
	unstakeCmd.Flags().BoolVarP(&WeiRazor, "weiRazor", "", false, "value can be passed in wei")
	unstakeCmd.Flags().Uint32VarP(&StakerId, "stakerId", "", 0, "staker id")

	addrErr := unstakeCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
	valueErr := unstakeCmd.MarkFlagRequired("value")
	utils.CheckError("Value error: ", valueErr)

}
