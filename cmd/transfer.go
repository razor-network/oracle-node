package cmd

import (
	"razor/core"
	"razor/core/types"
	"razor/utils"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var transferCmd = &cobra.Command{
	Use:   "transfer",
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

		var password string
		if utils.IsFlagPassed("password") {
			passwordPath, _ := cmd.Flags().GetString("password")
			password = utils.GetPasswordFromFile(passwordPath)
		} else {
			password = utils.PasswordPrompt()
		}

		fromAddress, _ := cmd.Flags().GetString("from")
		toAddress, _ := cmd.Flags().GetString("to")

		amount, err := cmd.Flags().GetString("amount")
		if err != nil {
			log.Fatal("Error in reading amount", err)
		}

		client := utils.ConnectToClient(config.Provider)

		balance, err := utils.FetchBalance(client, fromAddress)
		if err != nil {
			log.Fatalf("Error in fetching balance for account %s: %e", balance, err)
		}

		amountInWei := utils.GetAmountWithChecks(amount, balance)

		tokenManager := utils.GetTokenManager(client)
		txnOpts := utils.GetTxnOpts(types.TransactionOptions{
			Client:         client,
			Password:       password,
			AccountAddress: fromAddress,
			ChainId:        core.ChainId,
			Config:         config,
		})
		log.Infof("Transferring %s tokens from %s to %s", amount, fromAddress, toAddress)

		txn, err := tokenManager.Transfer(txnOpts, common.HexToAddress(toAddress), amountInWei)
		if err != nil {
			log.Fatal(err)
		}

		log.Info("Transfer transaction sent.")
		log.Info("Transaction Hash: ", txn.Hash())
		utils.WaitForBlockCompletion(client, txn.Hash().String())
	},
}

func init() {
	rootCmd.AddCommand(transferCmd)
	var (
		Amount   string
		From     string
		To       string
		Password string
	)

	transferCmd.Flags().StringVarP(&Amount, "amount", "a", "0", "amount to transfer (in Wei)")
	transferCmd.Flags().StringVarP(&From, "from", "", "", "transfer from")
	transferCmd.Flags().StringVarP(&To, "to", "", "", "transfer to")
	transferCmd.Flags().StringVarP(&Password, "password", "", "", "your account password")

	amountErr := transferCmd.MarkFlagRequired("amount")
	utils.CheckError("Amount error: ", amountErr)
	fromErr := transferCmd.MarkFlagRequired("from")
	utils.CheckError("From address error: ", fromErr)
	toErr := transferCmd.MarkFlagRequired("to")
	utils.CheckError("To address error: ", toErr)

}
