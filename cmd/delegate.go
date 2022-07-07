//Package cmd provides all functions related to command line
package cmd

import (
	"razor/core"
	"razor/core/types"
	"razor/logger"
	"razor/pkg/bindings"
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
	razorUtils.AssignLogFile(flagSet)
	address, err := flagSetUtils.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)

	logger.Address = address

	config, err := cmdUtils.GetConfigData()
	utils.CheckError("Error in getting config: ", err)

	password := razorUtils.AssignPassword(flagSet)

	stakerId, err := flagSetUtils.GetUint32StakerId(flagSet)
	utils.CheckError("Error in getting stakerId: ", err)

	client := razorUtils.ConnectToClient(config.Provider)

	balance, err := razorUtils.FetchBalance(client, address)
	utils.CheckError("Error in fetching razor balance for account "+address+": ", err)

	valueInWei, err := cmdUtils.AssignAmountInWei(flagSet)
	utils.CheckError("Error in getting amount: ", err)

	razorUtils.CheckAmountAndBalance(valueInWei, balance)

	razorUtils.CheckEthBalanceIsZero(client, address)

	txnArgs := types.TransactionOptions{
		Client:         client,
		Password:       password,
		Amount:         valueInWei,
		AccountAddress: address,
		ChainId:        core.ChainId,
		Config:         config,
	}

	approveTxnHash, err := cmdUtils.Approve(txnArgs)
	utils.CheckError("Approve error: ", err)

	if approveTxnHash != core.NilHash {
		err = razorUtils.WaitForBlockCompletion(txnArgs.Client, approveTxnHash.String())
		utils.CheckError("Error in WaitForBlockCompletion for approve: ", err)
	}

	delegateTxnHash, err := cmdUtils.Delegate(txnArgs, stakerId)
	utils.CheckError("Delegate error: ", err)
	err = razorUtils.WaitForBlockCompletion(client, delegateTxnHash.String())
	utils.CheckError("Error in WaitForBlockCompletion for delegate: ", err)
}

//This function allows the delegator to stake coins without setting up a node
func (*UtilsStruct) Delegate(txnArgs types.TransactionOptions, stakerId uint32) (common.Hash, error) {
	log.Infof("Delegating %g razors to Staker %d", razorUtils.GetAmountInDecimal(txnArgs.Amount), stakerId)
	txnArgs.ContractAddress = core.StakeManagerAddress
	txnArgs.MethodName = "delegate"
	txnArgs.ABI = bindings.StakeManagerABI
	txnArgs.Parameters = []interface{}{stakerId, txnArgs.Amount}
	delegationTxnOpts := razorUtils.GetTxnOpts(txnArgs)
	log.Info("Sending Delegate transaction...")
	txn, err := stakeManagerUtils.Delegate(txnArgs.Client, delegationTxnOpts, stakerId, txnArgs.Amount)
	if err != nil {
		return common.Hash{0x00}, err
	}
	log.Infof("Transaction hash: %s", transactionUtils.Hash(txn))
	return transactionUtils.Hash(txn), nil
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
