package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/utils"
)

// delegateCmd represents the delegate command
var delegateCmd = &cobra.Command{
	Use:   "delegate",
	Short: "delegate is used by delegator to stake coins on the network without setting up a node",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		if err != nil {
			log.Fatal("Error in getting config: ", err)
		}

		password := utils.PasswordPrompt()
		address, _ := cmd.Flags().GetString("address")
		stakerId, _ := cmd.Flags().GetString("stakerId")
		amount, _ := cmd.Flags().GetString("amount")

		client := utils.ConnectToClient(config.Provider)

		balance, err := utils.FetchBalance(client, address)
		if err != nil {
			log.Fatalf("Error in fetching balance for account %s: %e", address, err)
		}

		amountInWei := utils.GetAmountWithChecks(amount, balance)
		epoch, err := WaitForCommitState(client, address, "delegate")
		if err != nil {
			log.Fatal("Error in fetching epoch: ", epoch)
		}
		_stakerId, ok := new(big.Int).SetString(stakerId, 10)
		if !ok {
			log.Fatal("SetString error while converting stakerId")
		}

		stakeManager := utils.GetStakeManager(client)
		txnOpts := types.TransactionOptions{
			Client:         client,
			Password:       password,
			Amount:         amountInWei,
			AccountAddress: address,
			ChainId:        core.ChainId,
			GasMultiplier:  config.GasMultiplier,
		}

		approve(txnOpts)

		log.Infof("Delegating %s razors to Staker %s", amount, _stakerId)
		txn, err := stakeManager.Delegate(utils.GetTxnOpts(txnOpts), epoch, amountInWei, _stakerId)
		if err != nil {
			log.Fatal(err)
		}
		log.Infof("Sending Delegate transaction...")
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

	delegateCmd.Flags().StringVarP(&Amount, "amount", "a", "0", "amount to stake (in Wei)")
	delegateCmd.Flags().StringVarP(&Address, "address", "", "", "your account address")
	delegateCmd.Flags().StringVarP(&StakerId, "stakerId", "", "", "staker id")

	delegateCmd.MarkFlagRequired("amount")
	delegateCmd.MarkFlagRequired("address")
	delegateCmd.MarkFlagRequired("stakerId")

}
