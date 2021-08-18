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
	Short: "transfer razors from your account to others' account",
	Long: `transfer command allows user to transfer their razors to another account if they've not staked those razors

Example:
  ./razor transfer --amount 100 --to 0x91b1E6488307450f4c0442a1c35Bc314A505293e --from 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		if err != nil {
			log.Fatal("Error in getting config: ", err)
		}
		password := utils.PasswordPrompt()
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
		Amount string
		From   string
		To     string
	)

	transferCmd.Flags().StringVarP(&Amount, "amount", "a", "0", "amount to transfer (in Wei)")
	transferCmd.Flags().StringVarP(&From, "from", "", "", "transfer from")
	transferCmd.Flags().StringVarP(&To, "to", "", "", "transfer to")

	amountErr := transferCmd.MarkFlagRequired("amount")
	utils.CheckError("Amount error: ", amountErr)
	fromErr := transferCmd.MarkFlagRequired("from")
	utils.CheckError("From address error: ", fromErr)
	toErr := transferCmd.MarkFlagRequired("to")
	utils.CheckError("To address error: ", toErr)

}
