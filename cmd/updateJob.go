package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"razor/core"
	"razor/core/types"
	"razor/utils"
)

var updateJobCmd = &cobra.Command{
	Use:   "updateJob",
	Short: "updateJob can be used to update an existing job",
	Long: `A job consists of a URL and a selector to fetch the exact data from the URL. The updateJob command can be used to update an existing job that the stakers can vote upon.

Example:
  ./razor updateJob -a 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --jobId 1 -r true -s last -u https://api.gemini.com/v1/pubticker/btcusd

Note: 
  This command only works for the admin.
`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		utils.CheckError("Error in getting config: ", err)
		txn, err := updateJob(cmd.Flags(), config, razorUtils, assetManagerUtils, transactionUtils, flagSetUtils)
		utils.CheckError("UpdateJob error: ", err)
		utils.WaitForBlockCompletion(utils.ConnectToClient(config.Provider), txn.String())
	},
}

func updateJob(flagSet *pflag.FlagSet, config types.Configurations, razorUtils utilsInterface, assetManagerUtils assetManagerInterface, transactionUtils transactionInterface, flagSetUtils flagSetInterface) (common.Hash, error) {
	password := razorUtils.AssignPassword(flagSet)
	address, err := flagSetUtils.GetStringAddress(flagSet)
	if err != nil {
		return core.NilHash, err
	}
	jobId, err := flagSetUtils.GetUint8JobId(flagSet)
	if err != nil {
		return core.NilHash, err
	}
	power, err := flagSetUtils.GetInt8Power(flagSet)
	if err != nil {
		return core.NilHash, err
	}
	selector, err := flagSetUtils.GetStringSelector(flagSet)
	if err != nil {
		return core.NilHash, err
	}
	url, err := flagSetUtils.GetStringUrl(flagSet)
	if err != nil {
		return core.NilHash, err
	}
	client := razorUtils.ConnectToClient(config.Provider)
	txnArgs := razorUtils.GetTxnOpts(types.TransactionOptions{
		Client:          client,
		Password:        password,
		AccountAddress:  address,
		ChainId:         core.ChainId,
		Config:          config,
	})
	txn, err := assetManagerUtils.UpdateJob(client, txnArgs, jobId, power, selector, url)
	if err != nil {
		return core.NilHash, err
	}
	return transactionUtils.Hash(txn), nil
}

func init() {
	rootCmd.AddCommand(updateJobCmd)

	razorUtils = Utils{}
	assetManagerUtils = AssetManagerUtils{}
	transactionUtils = TransactionUtils{}
	flagSetUtils = FlagSetUtils{}

	var (
		URL      string
		Selector string
		Power    int8
		Account  string
		Password string
	)

	updateJobCmd.Flags().StringVarP(&URL, "url", "u", "", "url of job")
	updateJobCmd.Flags().StringVarP(&Selector, "selector", "s", "", "selector (jsonPath selector)")
	updateJobCmd.Flags().Int8VarP(&Power, "power", "", 0, "power")
	updateJobCmd.Flags().StringVarP(&Account, "address", "a", "", "address of the job creator")
	updateJobCmd.Flags().StringVarP(&Password, "password", "", "", "password path of job creator to protect the keystore")

	urlErr := updateJobCmd.MarkFlagRequired("url")
	utils.CheckError("URL error: ", urlErr)
	selectorErr := updateJobCmd.MarkFlagRequired("selector")
	utils.CheckError("Selector error: ", selectorErr)
	addrErr := updateJobCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
	powErr := updateJobCmd.MarkFlagRequired("power")
	utils.CheckError("Power error: ", powErr)
}
