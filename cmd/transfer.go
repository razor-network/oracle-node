//Package cmd provides all functions related to command line
package cmd

import (
	"razor/accounts"
	"razor/core"
	"razor/core/types"
	"razor/logger"
	"razor/pkg/bindings"
	"razor/utils"

	"github.com/ethereum/go-ethereum/ethclient"

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

//This function initialises the ExecuteTransfer function
func initialiseTransfer(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteTransfer(cmd.Flags())
}

//This function sets the flag appropriately and executes the Transfer function
func (*UtilsStruct) ExecuteTransfer(flagSet *pflag.FlagSet) {
	config, err := cmdUtils.GetConfigData()
	utils.CheckError("Error in getting config: ", err)
	log.Debugf("ExecuteTransfer: Config: %+v", config)

	client := razorUtils.ConnectToClient(config.Provider)

	fromAddress, err := flagSetUtils.GetStringFrom(flagSet)
	utils.CheckError("Error in getting fromAddress: ", err)

	logger.SetLoggerParameters(client, fromAddress)

	log.Debug("Checking to assign log file...")
	fileUtils.AssignLogFile(flagSet, config)

	log.Debug("Getting password...")
	password := razorUtils.AssignPassword(flagSet)

	accountManager, err := razorUtils.AccountManagerForKeystore()
	utils.CheckError("Error in getting accounts manager for keystore: ", err)

	account := accounts.InitAccountStruct(fromAddress, password, accountManager)

	err = razorUtils.CheckPassword(account)
	utils.CheckError("Error in fetching private key from given password: ", err)

	toAddress, err := flagSetUtils.GetStringTo(flagSet)
	utils.CheckError("Error in getting toAddress: ", err)

	balance, err := razorUtils.FetchBalance(client, fromAddress)
	utils.CheckError("Error in fetching razor balance: ", err)

	log.Debug("Getting amount in wei...")
	valueInWei, err := cmdUtils.AssignAmountInWei(flagSet)
	utils.CheckError("Error in getting amount: ", err)

	transferInput := types.TransferInput{
		ToAddress:  toAddress,
		ValueInWei: valueInWei,
		Balance:    balance,
		Account:    account,
	}

	txn, err := cmdUtils.Transfer(client, config, transferInput)
	utils.CheckError("Transfer error: ", err)

	err = razorUtils.WaitForBlockCompletion(client, txn.Hex())
	utils.CheckError("Error in WaitForBlockCompletion for transfer: ", err)
}

//This function transfers the razors from your account to others account
func (*UtilsStruct) Transfer(client *ethclient.Client, config types.Configurations, transferInput types.TransferInput) (common.Hash, error) {
	log.Debug("Checking for sufficient balance...")
	razorUtils.CheckAmountAndBalance(transferInput.ValueInWei, transferInput.Balance)

	txnOpts := razorUtils.GetTxnOpts(types.TransactionOptions{
		Client:          client,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.RAZORAddress,
		MethodName:      "transfer",
		Parameters:      []interface{}{common.HexToAddress(transferInput.ToAddress), transferInput.ValueInWei},
		ABI:             bindings.RAZORMetaData.ABI,
		Account:         transferInput.Account,
	})
	log.Infof("Transferring %g tokens from %s to %s", utils.GetAmountInDecimal(transferInput.ValueInWei), transferInput.Account.Address, transferInput.ToAddress)

	log.Debugf("Executing Transfer transaction with toAddress: %s, amount: %s", transferInput.ToAddress, transferInput.ValueInWei)
	txn, err := tokenManagerUtils.Transfer(client, txnOpts, common.HexToAddress(transferInput.ToAddress), transferInput.ValueInWei)
	if err != nil {
		log.Errorf("Error in transferring tokens ")
		return core.NilHash, err
	}

	txnHash := transactionUtils.Hash(txn)
	log.Info("Txn Hash: ", txnHash.Hex())
	return txnHash, nil
}

func init() {
	rootCmd.AddCommand(transferCmd)
	var (
		Amount   string
		From     string
		To       string
		Password string
		WeiRazor bool
	)

	transferCmd.Flags().StringVarP(&Amount, "value", "v", "0", "value to transfer")
	transferCmd.Flags().StringVarP(&From, "from", "", "", "transfer from")
	transferCmd.Flags().StringVarP(&To, "to", "", "", "transfer to")
	transferCmd.Flags().StringVarP(&Password, "password", "", "", "password path to protect the keystore")
	transferCmd.Flags().BoolVarP(&WeiRazor, "weiRazor", "", false, "value can be passed in wei")

	amountErr := transferCmd.MarkFlagRequired("value")
	utils.CheckError("Value error: ", amountErr)
	fromErr := transferCmd.MarkFlagRequired("from")
	utils.CheckError("From address error: ", fromErr)
	toErr := transferCmd.MarkFlagRequired("to")
	utils.CheckError("To address error: ", toErr)

}
