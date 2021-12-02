package cmd

import (
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
	Run: func(cmd *cobra.Command, args []string) {
		utilsStruct := UtilsStruct{
			razorUtils:        razorUtils,
			tokenManagerUtils: tokenManagerUtils,
			transactionUtils:  transactionUtils,
			flagSetUtils:      flagSetUtils,
			cmdUtils:          cmdUtils,
		}
		config, err := GetConfigData()
		utils.CheckError("Error in getting config: ", err)
		txn, err := utilsStruct.transfer(cmd.Flags(), config)
		utils.CheckError("Transfer error: ", err)
		log.Info("Transaction Hash: ", txn)
		utils.WaitForBlockCompletion(utils.ConnectToClient(config.Provider), txn.String())
	},
}

func (utilsStruct UtilsStruct) transfer(flagSet *pflag.FlagSet, config types.Configurations) (common.Hash, error) {

	password := utilsStruct.razorUtils.AssignPassword(flagSet)
	fromAddress, err := utilsStruct.flagSetUtils.GetStringFrom(flagSet)
	if err != nil {
		return core.NilHash, err
	}
	toAddress, err := utilsStruct.flagSetUtils.GetStringTo(flagSet)
	if err != nil {
		return core.NilHash, err
	}

	client := utilsStruct.razorUtils.ConnectToClient(config.Provider)

	balance, err := utilsStruct.razorUtils.FetchBalance(client, fromAddress)
	if err != nil {
		log.Errorf("Error in fetching balance for account " + fromAddress)
		return core.NilHash, err
	}

	valueInWei, err := utilsStruct.cmdUtils.AssignAmountInWei(flagSet, utilsStruct)
	if err != nil {
		log.Error("Error in getting amount: ", err)
		return core.NilHash, err
	}

	utilsStruct.razorUtils.CheckAmountAndBalance(valueInWei, balance)

	txnOpts := utilsStruct.razorUtils.GetTxnOpts(types.TransactionOptions{
		Client:          client,
		Password:        password,
		AccountAddress:  fromAddress,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.RAZORAddress,
		MethodName:      "transfer",
		Parameters:      []interface{}{common.HexToAddress(toAddress), valueInWei},
		ABI:             bindings.RAZORABI,
	})
	log.Infof("Transferring %g tokens from %s to %s", utilsStruct.razorUtils.GetAmountInDecimal(valueInWei), fromAddress, toAddress)

	txn, err := utilsStruct.tokenManagerUtils.Transfer(client, txnOpts, common.HexToAddress(toAddress), valueInWei)
	if err != nil {
		log.Errorf("Error in transferring tokens ")
		return core.NilHash, err
	}

	return utilsStruct.transactionUtils.Hash(txn), err
}

func init() {
	razorUtils = Utils{}
	tokenManagerUtils = TokenManagerUtils{}
	transactionUtils = TransactionUtils{}
	flagSetUtils = FlagSetUtils{}
	cmdUtils = UtilsCmd{}

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
