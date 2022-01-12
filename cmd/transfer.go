package cmd

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"

	"github.com/spf13/pflag"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "transfer razors from your account to others' account",
	Long: `transfer command allows user to transfer their razors to another account if they've not staked those razors

Example:
  ./razor transfer --value 100 --to 0x91b1E6488307450f4c0442a1c35Bc314A505293e --from 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c`,
	Run: initialiseTransfer,
}

func initialiseTransfer(cmd *cobra.Command, args []string) {
	cmdUtilsMockery.ExecuteTransfer(cmd.Flags())
}

func (*UtilsStructMockery) ExecuteTransfer(flagSet *pflag.FlagSet) {
	config, err := cmdUtilsMockery.GetConfigData()
	utils.CheckError("Error in getting config: ", err)
	password := razorUtilsMockery.AssignPassword(flagSet)
	fromAddress, err := flagSetUtilsMockery.GetStringFrom(flagSet)
	utils.CheckError("Error in getting fromAddress: ", err)
	toAddress, err := flagSetUtilsMockery.GetStringTo(flagSet)
	utils.CheckError("Error in getting toAddress: ", err)

	client := razorUtilsMockery.ConnectToClient(config.Provider)

	balance, err := razorUtilsMockery.FetchBalance(client, fromAddress)
	utils.CheckError("Error in fetching balance: ", err)

	valueInWei, err := cmdUtilsMockery.AssignAmountInWei(flagSet)
	utils.CheckError("Error in getting amount: ", err)

	transferInput := types.TransferInput{
		FromAddress: fromAddress,
		ToAddress:   toAddress,
		Password:    password,
		ValueInWei:  valueInWei,
		Balance:     balance,
	}

	txn, err := cmdUtilsMockery.Transfer(client, config, transferInput)
	utils.CheckError("Transfer error: ", err)
	log.Info("Transaction Hash: ", txn)
	razorUtilsMockery.WaitForBlockCompletion(client, txn.String())
}

func (*UtilsStructMockery) Transfer(client *ethclient.Client, config types.Configurations, transferInput types.TransferInput) (common.Hash, error) {

	razorUtilsMockery.CheckAmountAndBalance(transferInput.ValueInWei, transferInput.Balance)

	txnOpts := razorUtilsMockery.GetTxnOpts(types.TransactionOptions{
		Client:          client,
		Password:        transferInput.Password,
		AccountAddress:  transferInput.FromAddress,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.RAZORAddress,
		MethodName:      "transfer",
		Parameters:      []interface{}{common.HexToAddress(transferInput.ToAddress), transferInput.ValueInWei},
		ABI:             bindings.RAZORABI,
	})
	log.Infof("Transferring %g tokens from %s to %s", razorUtilsMockery.GetAmountInDecimal(transferInput.ValueInWei), transferInput.FromAddress, transferInput.ToAddress)

	txn, err := tokenManagerUtilsMockery.Transfer(client, txnOpts, common.HexToAddress(transferInput.ToAddress), transferInput.ValueInWei)
	if err != nil {
		log.Errorf("Error in transferring tokens ")
		return core.NilHash, err
	}

	return transactionUtilsMockery.Hash(txn), err
}

func init() {
	razorUtils = Utils{}
	transactionUtils = TransactionUtils{}
	cmdUtilsMockery = &UtilsStructMockery{}
	razorUtilsMockery = UtilsMockery{}
	tokenManagerUtilsMockery = TokenManagerUtilsMockery{}
	flagSetUtilsMockery = FLagSetUtilsMockery{}

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
