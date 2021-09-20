package cmd

import (
	"github.com/spf13/pflag"
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
		txn, err := transfer(cmd.Flags(), config)
		utils.CheckError("Transfer error: ", err)
		log.Info("Transaction Hash: ", txn)
		utils.WaitForBlockCompletion(utils.ConnectToClient(config.Provider), txn.String())
	},
}

func transfer(flagSet *pflag.FlagSet, config types.Configurations) (common.Hash, error) {

	password := razorUtils.AssignPassword(flagSet)
	fromAddress, _ := flagSet.GetString("from")
	toAddress, _ := flagSet.GetString("to")

	client := razorUtils.ConnectToClient(config.Provider)

	balance, err := razorUtils.FetchBalance(client, fromAddress)
	utils.CheckError("Error in fetching balance for account "+fromAddress+": ", err)

	valueInWei := razorUtils.AssignAmountInWei(flagSet)
	razorUtils.CheckAmountAndBalance(valueInWei, balance)

	txnOpts := razorUtils.GetTxnOpts(types.TransactionOptions{
		Client:         client,
		Password:       password,
		AccountAddress: fromAddress,
		ChainId:        core.ChainId,
		Config:         config,
	})
	log.Infof("Transferring %g tokens from %s to %s", razorUtils.GetAmountInDecimal(valueInWei), fromAddress, toAddress)

	txn, err := tokenManagerUtils.Transfer(client, txnOpts, common.HexToAddress(toAddress), valueInWei)
	utils.CheckError("Error in transferring tokens: ", err)

	return transactionUtils.Hash(txn), err
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
