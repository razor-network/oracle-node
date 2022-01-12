package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/pflag"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"

	"github.com/spf13/cobra"
)

// delegateCmd represents the delegate command
var delegateCmd = &cobra.Command{
	Use:   "delegate",
	Short: "delegate can be used by delegator to stake coins on the network without setting up a node",
	Long: `If a user has Razors with them, and wants to stake them but doesn't want to set up a node, they can use the delegate command.

Example:
  ./razor delegate --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --value 1000 --stakerId 1
`,
	Run: initialiseDelegate,
}

func initialiseDelegate(cmd *cobra.Command, args []string) {
	cmdUtilsMockery.ExecuteDelegate(cmd.Flags())
}

func (*UtilsStructMockery) ExecuteDelegate(flagSet *pflag.FlagSet) {
	config, err := cmdUtilsMockery.GetConfigData()
	utils.CheckError("Error in getting config: ", err)

	password := razorUtilsMockery.AssignPassword(flagSet)
	address, err := flagSetUtilsMockery.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)

	stakerId, err := flagSetUtilsMockery.GetUint32StakerId(flagSet)
	utils.CheckError("Error in getting stakerId: ", err)

	client := razorUtilsMockery.ConnectToClient(config.Provider)

	balance, err := razorUtilsMockery.FetchBalance(client, address)
	utils.CheckError("Error in fetching balance for account "+address+": ", err)

	valueInWei, err := cmdUtilsMockery.AssignAmountInWei(flagSet)
	utils.CheckError("Error in getting amount: ", err)

	razorUtilsMockery.CheckAmountAndBalance(valueInWei, balance)

	razorUtilsMockery.CheckEthBalanceIsZero(client, address)

	txnArgs := types.TransactionOptions{
		Client:         client,
		Password:       password,
		Amount:         valueInWei,
		AccountAddress: address,
		ChainId:        core.ChainId,
		Config:         config,
	}

	approveTxnHash, err := cmdUtilsMockery.Approve(txnArgs)
	utils.CheckError("Approve error: ", err)

	if approveTxnHash != core.NilHash {
		razorUtilsMockery.WaitForBlockCompletion(txnArgs.Client, approveTxnHash.String())
	}

	delegateTxnHash, err := cmdUtilsMockery.Delegate(txnArgs, stakerId)
	utils.CheckError("Delegate error: ", err)
	razorUtilsMockery.WaitForBlockCompletion(client, delegateTxnHash.String())
}

func (*UtilsStructMockery) Delegate(txnArgs types.TransactionOptions, stakerId uint32) (common.Hash, error) {
	log.Infof("Delegating %g razors to Staker %d", razorUtilsMockery.GetAmountInDecimal(txnArgs.Amount), stakerId)
	epoch, err := razorUtilsMockery.GetEpoch(txnArgs.Client)
	if err != nil {
		return common.Hash{0x00}, err
	}
	txnArgs.ContractAddress = core.StakeManagerAddress
	txnArgs.MethodName = "delegate"
	txnArgs.ABI = bindings.StakeManagerABI
	txnArgs.Parameters = []interface{}{epoch, stakerId, txnArgs.Amount}
	delegationTxnOpts := razorUtilsMockery.GetTxnOpts(txnArgs)
	log.Info("Sending Delegate transaction...")
	txn, err := stakeManagerUtilsMockery.Delegate(txnArgs.Client, delegationTxnOpts, stakerId, txnArgs.Amount)
	if err != nil {
		return common.Hash{0x00}, err
	}
	log.Infof("Transaction hash: %s", transactionUtilsMockery.Hash(txn))
	return transactionUtilsMockery.Hash(txn), nil
}

func init() {
	razorUtilsMockery = UtilsMockery{}
	transactionUtilsMockery = TransactionUtilsMockery{}
	stakeManagerUtilsMockery = StakeManagerUtilsMockery{}
	flagSetUtilsMockery = FLagSetUtilsMockery{}
	cmdUtilsMockery = &UtilsStructMockery{}

	rootCmd.AddCommand(delegateCmd)
	var (
		Amount   string
		Address  string
		StakerId uint32
		Password string
		Power    string
	)

	delegateCmd.Flags().StringVarP(&Amount, "value", "v", "0", "amount to stake (in Wei)")
	delegateCmd.Flags().StringVarP(&Address, "address", "a", "", "your account address")
	delegateCmd.Flags().Uint32VarP(&StakerId, "stakerId", "", 0, "staker id")
	delegateCmd.Flags().StringVarP(&Password, "password", "", "", "password path to protect the keystore")
	delegateCmd.Flags().StringVarP(&Power, "pow", "", "", "power of 10")

	valueErr := delegateCmd.MarkFlagRequired("value")
	utils.CheckError("Value error: ", valueErr)
	addrErr := delegateCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
	stakerIdErr := delegateCmd.MarkFlagRequired("stakerId")
	utils.CheckError("StakerId error: ", stakerIdErr)

}
