//Package cmd provides all functions related to command line
package cmd

import (
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/rpc"
	"razor/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// delegateCmd represents the delegate command
var delegateCmd = &cobra.Command{
	Use:   "delegate",
	Short: "delegate can be used by delegator to stake coins on the network without setting up a node",
	Long: `If a user has Razors with them, and wants to stake them but doesn't want to set up a node, they can use the delegate command.

Example:
  ./razor delegate --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --value 1000 --stakerId 1 --logFile delegateLogs
`,
	Run: initialiseDelegate,
}

//This function initialises the ExecuteDelegate function
func initialiseDelegate(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteDelegate(cmd.Flags())
}

//This function sets the flags appropriately and executes the Delegate function
func (*UtilsStruct) ExecuteDelegate(flagSet *pflag.FlagSet) {
	config, rpcParameters, account, err := InitializeCommandDependencies(flagSet)
	utils.CheckError("Error in initialising command dependencies: ", err)

	stakerId, err := flagSetUtils.GetUint32StakerId(flagSet)
	utils.CheckError("Error in getting stakerId: ", err)
	log.Debug("ExecuteDelegate: Staker Id: ", stakerId)
	balance, err := razorUtils.FetchBalance(rpcParameters, account.Address)
	utils.CheckError("Error in fetching razor balance for account "+account.Address+": ", err)
	log.Debug("ExecuteDelegate: Balance: ", balance)

	log.Debug("Getting amount in wei...")
	valueInWei, err := cmdUtils.AssignAmountInWei(flagSet)
	utils.CheckError("Error in getting amount: ", err)

	log.Debug("Checking for sufficient balance...")
	razorUtils.CheckAmountAndBalance(valueInWei, balance)

	log.Debug("Checking whether sFUEL balance is not 0...")
	razorUtils.CheckEthBalanceIsZero(rpcParameters, account.Address)

	txnArgs := types.TransactionOptions{
		Amount:  valueInWei,
		ChainId: core.ChainId,
		Config:  config,
		Account: account,
	}

	approveTxnHash, err := cmdUtils.Approve(rpcParameters, txnArgs)
	utils.CheckError("Approve error: ", err)

	if approveTxnHash != core.NilHash {
		err = razorUtils.WaitForBlockCompletion(rpcParameters, approveTxnHash.Hex())
		utils.CheckError("Error in WaitForBlockCompletion for approve: ", err)
	}

	log.Debug("ExecuteDelegate:Calling Delegate() with stakerId: ", stakerId)
	delegateTxnHash, err := cmdUtils.Delegate(rpcParameters, txnArgs, stakerId)
	utils.CheckError("Delegate error: ", err)
	err = razorUtils.WaitForBlockCompletion(rpcParameters, delegateTxnHash.Hex())
	utils.CheckError("Error in WaitForBlockCompletion for delegate: ", err)
}

//This function allows the delegator to stake coins without setting up a node
func (*UtilsStruct) Delegate(rpcParameters rpc.RPCParameters, txnArgs types.TransactionOptions, stakerId uint32) (common.Hash, error) {
	log.Infof("Delegating %g razors to Staker %d", utils.GetAmountInDecimal(txnArgs.Amount), stakerId)
	txnArgs.ContractAddress = core.StakeManagerAddress
	txnArgs.MethodName = "delegate"
	txnArgs.ABI = bindings.StakeManagerMetaData.ABI
	txnArgs.Parameters = []interface{}{stakerId, txnArgs.Amount}
	delegationTxnOpts, err := razorUtils.GetTxnOpts(rpcParameters, txnArgs)
	if err != nil {
		return core.NilHash, err
	}
	log.Info("Sending Delegate transaction...")

	client, err := rpcParameters.RPCManager.GetBestRPCClient()
	if err != nil {
		return core.NilHash, err
	}

	log.Debugf("Executing Delegate transaction with stakerId = %d, amount = %s", stakerId, txnArgs.Amount)
	txn, err := stakeManagerUtils.Delegate(client, delegationTxnOpts, stakerId, txnArgs.Amount)
	if err != nil {
		return core.NilHash, err
	}
	txnHash := transactionUtils.Hash(txn)
	log.Info("Txn Hash: ", txnHash.Hex())
	return txnHash, nil
}

func init() {
	rootCmd.AddCommand(delegateCmd)
	var (
		Amount   string
		Address  string
		StakerId uint32
		Password string
		WeiRazor bool
	)

	delegateCmd.Flags().StringVarP(&Amount, "value", "v", "0", "amount to stake (in Wei)")
	delegateCmd.Flags().StringVarP(&Address, "address", "a", "", "your account address")
	delegateCmd.Flags().Uint32VarP(&StakerId, "stakerId", "", 0, "staker id")
	delegateCmd.Flags().StringVarP(&Password, "password", "", "", "password path to protect the keystore")
	delegateCmd.Flags().BoolVarP(&WeiRazor, "weiRazor", "", false, "value can be passed in wei")

	valueErr := delegateCmd.MarkFlagRequired("value")
	utils.CheckError("Value error: ", valueErr)
	addrErr := delegateCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
	stakerIdErr := delegateCmd.MarkFlagRequired("stakerId")
	utils.CheckError("StakerId error: ", stakerIdErr)

}
