package cmd

import (
	"razor/core"
	"razor/core/types"
	"razor/utils"

	log "github.com/sirupsen/logrus"
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
		config, err := GetConfigData()
		utils.CheckError("Error in getting config: ", err)

		password := utils.AssignPassword(cmd.Flags())
		address, _ := cmd.Flags().GetString("address")
		stakerId, _ := cmd.Flags().GetUint32("stakerId")

		client := utils.ConnectToClient(config.Provider)

		balance, err := utils.FetchBalance(client, address)
		utils.CheckError("Error in fetching balance for account "+address+": ", err)

		valueInWei := utils.AssignAmountInWei(cmd.Flags())
		utils.CheckAmountAndBalance(valueInWei, balance)

		utils.CheckEthBalanceIsZero(client, address)

		stakeManager := utils.GetStakeManager(client)
		txnOpts := types.TransactionOptions{
			Client:         client,
			Password:       password,
			Amount:         valueInWei,
			AccountAddress: address,
			ChainId:        core.ChainId,
			Config:         config,
		}

		approve(txnOpts)

		log.Infof("Delegating %g razors to Staker %d", utils.GetAmountInDecimal(valueInWei), stakerId)
		delegationTxnOpts := utils.GetTxnOpts(txnOpts)
		epoch, err := WaitForCommitState(client, address, "delegate")
		utils.CheckError("Error in fetching epoch: ", err)
		txn, err := stakeManager.Delegate(delegationTxnOpts, epoch, stakerId, valueInWei)
		utils.CheckError("Error in delegating: ", err)
		log.Infof("Sending Delegate transaction...")
		log.Infof("Transaction hash: %s", txn.Hash())
		utils.WaitForBlockCompletion(client, txn.Hash().String())
	},
}

func init() {
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
