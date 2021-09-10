package cmd

import (
	"razor/core"
	"razor/core/types"
	"razor/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "transfer razors from your account to others' account",
	Long: `transfer command allows user to transfer their razors to another account if they've not staked those razors

Example:
  ./razor transfer --value 100 --to 0x91b1E6488307450f4c0442a1c35Bc314A505293e --from 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		utils.CheckError("Error in getting config: ", err)

		password := utils.AssignPassword(cmd.Flags())
		fromAddress, _ := cmd.Flags().GetString("from")
		toAddress, _ := cmd.Flags().GetString("to")

		client := utils.ConnectToClient(config.Provider)

		balance, err := utils.FetchBalance(client, fromAddress)
		utils.CheckError("Error in fetching balance for account "+fromAddress+": ", err)

		valueInWei := utils.AssignAmountInWei(cmd.Flags())
		utils.CheckAmountAndBalance(valueInWei, balance)

		tokenManager := utils.GetTokenManager(client)
		txnOpts := utils.GetTxnOpts(types.TransactionOptions{
			Client:         client,
			Password:       password,
			AccountAddress: fromAddress,
			ChainId:        core.ChainId,
			Config:         config,
		})
		log.Infof("Transferring %g tokens from %s to %s", utils.GetAmountInDecimal(valueInWei), fromAddress, toAddress)

		txn, err := tokenManager.Transfer(txnOpts, common.HexToAddress(toAddress), valueInWei)
		utils.CheckError("Error in transferring tokens: ", err)

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
		Power    string
	)

	transferCmd.Flags().StringVarP(&Amount, "value", "v", "0", "value to transfer")
	transferCmd.Flags().StringVarP(&From, "from", "", "", "transfer from")
	transferCmd.Flags().StringVarP(&To, "to", "", "", "transfer to")
	transferCmd.Flags().StringVarP(&Password, "password", "", "", "password path to protect the keystore")
	transferCmd.Flags().StringVarP(&Power, "pow", "", "", "power of 10")

	amountErr := transferCmd.MarkFlagRequired("value")
	utils.CheckError("Value error: ", amountErr)
	fromErr := transferCmd.MarkFlagRequired("from")
	utils.CheckError("From address error: ", fromErr)
	toErr := transferCmd.MarkFlagRequired("to")
	utils.CheckError("To address error: ", toErr)

}
