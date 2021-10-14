package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/pflag"
	"razor/core"
	"razor/core/types"
	"razor/utils"

	"github.com/spf13/cobra"
)

var assetManagerUtils assetManagerInterface
var flagSetUtils flagSetInterface

var createJobCmd = &cobra.Command{
	Use:   "createJob",
	Short: "createJob can be used to create a job",
	Long: `A job consists of a URL and a selector to fetch the exact data from the URL. The createJob command can be used to create a job that the stakers can vote upon.

Example:
  ./razor createJob -a 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c -n btcusd_gemini -r true -s last -u https://api.gemini.com/v1/pubticker/btcusd

Note: 
  This command only works for the admin.
`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		utils.CheckError("Error in getting config: ", err)
		txn, err := createJob(cmd.Flags(), config, razorUtils, assetManagerUtils, transactionUtils, flagSetUtils)
		utils.CheckError("CreateJob error: ", err)
		utils.WaitForBlockCompletion(utils.ConnectToClient(config.Provider), txn.String())
	},
}

func createJob(flagSet *pflag.FlagSet, config types.Configurations, razorUtils utilsInterface, assetManagerUtils assetManagerInterface, transactionUtils transactionInterface, flagSetUtils flagSetInterface) (common.Hash, error) {
	password := razorUtils.AssignPassword(flagSet)
	address, err := flagSetUtils.GetStringAddress(flagSet)
	if err != nil {
		return core.NilHash, err
	}

	name, err := flagSetUtils.GetStringName(flagSet)
	if err != nil {
		return core.NilHash, err
	}

	url, err := flagSetUtils.GetStringUrl(flagSet)
	if err != nil {
		return core.NilHash, err
	}

	selector, err := flagSetUtils.GetStringSelector(flagSet)
	if err != nil {
		return core.NilHash, err
	}

	power, err := flagSetUtils.GetInt8Power(flagSet)
	if err != nil {
		return core.NilHash, err
	}

	weight, err := flagSetUtils.GetUint8Weight(flagSet)
	if err != nil {
		return core.NilHash, err
	}

	client := razorUtils.ConnectToClient(config.Provider)
	selectorType := 1
	txnArgs := types.TransactionOptions{
		Client:         client,
		Password:       password,
		AccountAddress: address,
		ChainId:        core.ChainId,
		Config:         config,
	}

	txnOpts := razorUtils.GetTxnOpts(txnArgs)
	log.Info("Creating Job...")
	txn, err := assetManagerUtils.CreateJob(txnArgs.Client, txnOpts, weight, power, uint8(selectorType), name, selector, url)
	if err != nil {
		return core.NilHash, err
	}
	log.Info("Transaction Hash: ", transactionUtils.Hash(txn))
	return transactionUtils.Hash(txn), nil
}

func init() {

	razorUtils = Utils{}
	assetManagerUtils = AssetManagerUtils{}
	transactionUtils = TransactionUtils{}
	flagSetUtils = FlagSetUtils{}

	rootCmd.AddCommand(createJobCmd)

	var (
		URL      string
		Selector string
		Name     string
		Power    int8
		Account  string
		Password string
		Weight uint8
	)

	createJobCmd.Flags().StringVarP(&URL, "url", "u", "", "url of job")
	createJobCmd.Flags().StringVarP(&Selector, "selector", "s", "", "selector (jsonPath selector)")
	createJobCmd.Flags().StringVarP(&Name, "name", "n", "", "name of job")
	createJobCmd.Flags().Int8VarP(&Power, "power", "", 0, "power")
	createJobCmd.Flags().Uint8VarP(&Weight, "weight", "", 0, "weight assigned to the job")
	createJobCmd.Flags().StringVarP(&Account, "address", "a", "", "address of the job creator")
	createJobCmd.Flags().StringVarP(&Password, "password", "", "", "password path of job creator to protect the keystore")

	urlErr := createJobCmd.MarkFlagRequired("url")
	utils.CheckError("URL error: ", urlErr)
	selectorErr := createJobCmd.MarkFlagRequired("selector")
	utils.CheckError("Selector error: ", selectorErr)
	nameErr := createJobCmd.MarkFlagRequired("name")
	utils.CheckError("Name error: ", nameErr)
	addrErr := createJobCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
	powErr := createJobCmd.MarkFlagRequired("power")
	utils.CheckError("Power error: ", powErr)
	weightErr := createJobCmd.MarkFlagRequired("weight")
	utils.CheckError("Weight error: ", weightErr)
}
