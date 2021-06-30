package cmd

import (
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/utils"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// withdrawCmd represents the withdraw command
var withdrawCmd = &cobra.Command{
	Use:   "withdraw",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		if err != nil {
			log.Fatal("Error in getting config: ", err)
		}
		password := utils.PasswordPrompt()
		address, _ := cmd.Flags().GetString("address")

		client := utils.ConnectToClient(config.Provider)

		balance, err := utils.FetchBalance(client, address)
		if err != nil {
			log.Fatalf("Error in fetching balance for account %s: %e", balance, err)
		}
		if balance.Cmp(big.NewInt(0)) == 0 {
			log.Fatal("Balance is 0. Aborting...")
		}

		epoch, err := WaitForCommitState(client, address, "withdraw")
		stakeManager := utils.GetStakeManager(client)
		txnOpts := utils.GetTxnOpts(types.TransactionOptions{
			Client:         client,
			Password:       password,
			AccountAddress: address,
			ChainId:        core.ChainId,
			GasMultiplier:  config.GasMultiplier,
		})
		log.Info("Withdrawing funds...")
		stakerId, err := utils.GetStakerId(client, address)
		if err != nil {
			log.Fatal(err)
		}
		txn, err := stakeManager.Withdraw(txnOpts, epoch, stakerId)
		if err != nil {
			log.Fatal(err)
		}
		log.Info("Withdraw Transaction sent.")
		log.Info("Txn Hash: ", txn.Hash())
		utils.WaitForBlockCompletion(client, txn.Hash().String())
	},
}

func init() {
	rootCmd.AddCommand(withdrawCmd)

	var Address string

	withdrawCmd.Flags().StringVarP(&Address, "address", "", "", "address of the staker")

	withdrawCmd.MarkFlagRequired("address")
}
