//Package cmd provides all functions related to command line
package cmd

import (
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/rpc"
	"razor/utils"

	"github.com/spf13/pflag"

	"github.com/ethereum/go-ethereum/common"

	"github.com/spf13/cobra"
)

var stakeCmd = &cobra.Command{
	Use:   "addStake",
	Short: "Stake some razors",
	Long: `addStake allows user to stake razors in the razor network

Example:
  ./razor addStake --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --value 1000 --logFile addStake`,
	Run: initialiseStake,
}

//This function initialises the ExecuteStake function
func initialiseStake(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteStake(cmd.Flags())
}

//This function sets the flags appropriately and executes the StakeCoins function
func (*UtilsStruct) ExecuteStake(flagSet *pflag.FlagSet) {
	config, rpcParameters, account, err := InitializeCommandDependencies(flagSet)
	utils.CheckError("Error in initialising command dependencies: ", err)

	balance, err := razorUtils.FetchBalance(rpcParameters, account.Address)
	utils.CheckError("Error in fetching razor balance for account: "+account.Address, err)
	log.Debug("Getting amount in wei...")
	valueInWei, err := cmdUtils.AssignAmountInWei(flagSet)
	utils.CheckError("Error in getting amount: ", err)
	log.Debug("ExecuteStake: Amount in wei: ", valueInWei)

	log.Debug("Checking for sufficient balance...")
	razorUtils.CheckAmountAndBalance(valueInWei, balance)

	log.Debug("Checking whether sFUEL balance is not 0...")
	razorUtils.CheckEthBalanceIsZero(rpcParameters, account.Address)

	minSafeRazor, err := razorUtils.GetMinSafeRazor(rpcParameters)
	utils.CheckError("Error in getting minimum safe razor amount: ", err)
	log.Debug("ExecuteStake: Minimum razor that you can stake for first time: ", minSafeRazor)

	stakerId, err := razorUtils.GetStakerId(rpcParameters, account.Address)
	utils.CheckError("Error in getting stakerId: ", err)
	log.Debug("ExecuteStake: Staker Id: ", stakerId)

	if valueInWei.Cmp(minSafeRazor) < 0 && stakerId == 0 {
		log.Fatal("The amount of razors entered is below min safe value.")
	}

	if stakerId != 0 {
		staker, err := razorUtils.GetStaker(rpcParameters, stakerId)
		utils.CheckError("Error in getting staker: ", err)

		if staker.IsSlashed {
			log.Fatal("Staker is slashed, cannot stake")
		}
	}

	txnArgs := types.TransactionOptions{
		Amount:  valueInWei,
		ChainId: core.ChainId,
		Config:  config,
		Account: account,
	}

	log.Debug("ExecuteStake: Calling Approve() for amount: ", txnArgs.Amount)
	approveTxnHash, err := cmdUtils.Approve(rpcParameters, txnArgs)
	utils.CheckError("Approve error: ", err)

	if approveTxnHash != core.NilHash {
		err = razorUtils.WaitForBlockCompletion(rpcParameters, approveTxnHash.Hex())
		utils.CheckError("Error in WaitForBlockCompletion for approve: ", err)
	}

	log.Debug("ExecuteStake: Calling StakeCoins() for amount: ", txnArgs.Amount)
	stakeTxnHash, err := cmdUtils.StakeCoins(rpcParameters, txnArgs)
	utils.CheckError("Stake error: ", err)

	err = razorUtils.WaitForBlockCompletion(rpcParameters, stakeTxnHash.Hex())
	utils.CheckError("Error in WaitForBlockCompletion for stake: ", err)
}

//This function allows the user to stake razors in the razor network and returns the hash
func (*UtilsStruct) StakeCoins(rpcParameters rpc.RPCParameters, txnArgs types.TransactionOptions) (common.Hash, error) {
	epoch, err := razorUtils.GetEpoch(rpcParameters)
	if err != nil {
		return core.NilHash, err
	}
	log.Debug("StakeCoins: Epoch: ", epoch)

	txnArgs.ContractAddress = core.StakeManagerAddress
	txnArgs.MethodName = "stake"
	txnArgs.Parameters = []interface{}{epoch, txnArgs.Amount}
	txnArgs.ABI = bindings.StakeManagerMetaData.ABI
	txnOpts, err := razorUtils.GetTxnOpts(rpcParameters, txnArgs)
	if err != nil {
		return core.NilHash, err
	}

	client, err := rpcParameters.RPCManager.GetBestRPCClient()
	if err != nil {
		return core.NilHash, err
	}

	log.Debugf("Executing Stake transaction with epoch = %d, amount = %d", epoch, txnArgs.Amount)
	txn, err := stakeManagerUtils.Stake(client, txnOpts, epoch, txnArgs.Amount)
	if err != nil {
		return core.NilHash, err
	}
	txnHash := transactionUtils.Hash(txn)
	log.Info("Txn Hash: ", txnHash.Hex())
	return txnHash, nil
}

func init() {
	rootCmd.AddCommand(stakeCmd)
	var (
		Amount   string
		Address  string
		Password string
		WeiRazor bool
	)

	stakeCmd.Flags().StringVarP(&Amount, "value", "v", "0", "amount of Razors to stake")
	stakeCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the staker")
	stakeCmd.Flags().StringVarP(&Password, "password", "", "", "password path of staker to protect the keystore")
	stakeCmd.Flags().BoolVarP(&WeiRazor, "weiRazor", "", false, "value can be passed in wei")

	amountErr := stakeCmd.MarkFlagRequired("value")
	utils.CheckError("Value error: ", amountErr)
	addrErr := stakeCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)

}
