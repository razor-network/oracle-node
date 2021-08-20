package cmd

import (
	"math/big"
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

		password := utils.PasswordPrompt()
		address, _ := cmd.Flags().GetString("address")
		stakerId, _ := cmd.Flags().GetString("stakerId")
		value, _ := cmd.Flags().GetString("value")

		client := utils.ConnectToClient(config.Provider)

		balance, err := utils.FetchBalance(client, address)
		utils.CheckError("Error in fetching balance for account "+address+": ", err)

		valueInWei := utils.GetAmountWithChecks(value, balance)
		_stakerId, ok := new(big.Int).SetString(stakerId, 10)
		if !ok {
			log.Fatal("SetString error while converting stakerId")
		}

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

		log.Infof("Delegating %s razors to Staker %s", value, _stakerId)
		delegationTxnOpts := utils.GetTxnOpts(txnOpts)
		epoch, err := WaitForCommitState(client, address, "delegate")
		utils.CheckError("Error in fetching epoch: ", err)
		txn, err := stakeManager.Delegate(delegationTxnOpts, epoch, valueInWei, _stakerId)
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
		StakerId string
	)

	delegateCmd.Flags().StringVarP(&Amount, "value", "v", "0", "amount to stake (in Wei)")
	delegateCmd.Flags().StringVarP(&Address, "address", "a", "", "your account address")
	delegateCmd.Flags().StringVarP(&StakerId, "stakerId", "", "", "staker id")

	valueErr := delegateCmd.MarkFlagRequired("value")
	utils.CheckError("Value error: ", valueErr)
	addrErr := delegateCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
	stakerIdErr := delegateCmd.MarkFlagRequired("stakerId")
	utils.CheckError("StakerId error: ", stakerIdErr)

}
