package cmd

import (
	"github.com/ethereum/go-ethereum/common"
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
	Run: func(cmd *cobra.Command, args []string) {
		utilsStruct := UtilsStruct{
			razorUtils:        razorUtils,
			tokenManagerUtils: tokenManagerUtils,
			transactionUtils:  transactionUtils,
			stakeManagerUtils: stakeManagerUtils,
			flagSetUtils:      flagSetUtils,
			packageUtils:      packageUtils,
		}
		config, err := GetConfigData(utilsStruct)
		utils.CheckError("Error in getting config: ", err)

		password := utils.AssignPassword(cmd.Flags())
		address, _ := cmd.Flags().GetString("address")
		stakerId, _ := cmd.Flags().GetUint32("stakerId")

		client := utils.ConnectToClient(config.Provider)

		balance, err := utils.FetchBalance(client, address)
		utils.CheckError("Error in fetching balance for account "+address+": ", err)

		valueInWei, err := AssignAmountInWei(cmd.Flags(), utilsStruct)
		utils.CheckError("Error in getting amount: ", err)

		utils.CheckAmountAndBalance(valueInWei, balance)

		utils.CheckEthBalanceIsZero(client, address)

		txnArgs := types.TransactionOptions{
			Client:         client,
			Password:       password,
			Amount:         valueInWei,
			AccountAddress: address,
			ChainId:        core.ChainId,
			Config:         config,
		}

		approveTxnHash, err := utilsStruct.approve(txnArgs)
		utils.CheckError("Approve error: ", err)

		if approveTxnHash != core.NilHash {
			razorUtils.WaitForBlockCompletion(txnArgs.Client, approveTxnHash.String())
		}

		delegateTxnHash, err := utilsStruct.delegate(txnArgs, stakerId)
		utils.CheckError("Delegate error: ", err)
		utils.WaitForBlockCompletion(client, delegateTxnHash.String())
	},
}

func (utilsStruct UtilsStruct) delegate(txnArgs types.TransactionOptions, stakerId uint32) (common.Hash, error) {
	log.Infof("Delegating %g razors to Staker %d", utilsStruct.razorUtils.GetAmountInDecimal(txnArgs.Amount), stakerId)
	epoch, err := utilsStruct.razorUtils.GetEpoch(txnArgs.Client)
	if err != nil {
		return common.Hash{0x00}, err
	}
	txnArgs.ContractAddress = core.StakeManagerAddress
	txnArgs.MethodName = "delegate"
	txnArgs.ABI = bindings.StakeManagerABI
	txnArgs.Parameters = []interface{}{epoch, stakerId, txnArgs.Amount}
	delegationTxnOpts := utilsStruct.razorUtils.GetTxnOpts(txnArgs)
	log.Info("Sending Delegate transaction...")
	txn, err := utilsStruct.stakeManagerUtils.Delegate(txnArgs.Client, delegationTxnOpts, epoch, stakerId, txnArgs.Amount)
	if err != nil {
		return common.Hash{0x00}, err
	}
	log.Infof("Transaction hash: %s", utilsStruct.transactionUtils.Hash(txn))
	return utilsStruct.transactionUtils.Hash(txn), nil
}

func init() {
	razorUtils = Utils{}
	transactionUtils = TransactionUtils{}
	stakeManagerUtils = StakeManagerUtils{}
	flagSetUtils = FlagSetUtils{}

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
