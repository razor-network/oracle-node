package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/pflag"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
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
  ./razor createJob -a 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c -n btcusd_gemini -p 2 -s last --selectorType 1 -u https://api.gemini.com/v1/pubticker/btcusd

Note: 
  This command only works for the admin.
`,
	Run: func(cmd *cobra.Command, args []string) {
		utilsStruct := UtilsStruct{
			razorUtils:        razorUtils,
			assetManagerUtils: assetManagerUtils,
			transactionUtils:  transactionUtils,
			flagSetUtils:      flagSetUtils,
			packageUtils:      packageUtils,
		}
		config, err := GetConfigData(utilsStruct)
		utils.CheckError("Error in getting config: ", err)
		txn, err := utilsStruct.createJob(cmd.Flags(), config)
		utils.CheckError("CreateJob error: ", err)
		utils.WaitForBlockCompletion(utils.ConnectToClient(config.Provider), txn.String())
	},
}

func (utilsStruct UtilsStruct) createJob(flagSet *pflag.FlagSet, config types.Configurations) (common.Hash, error) {
	password := utilsStruct.razorUtils.AssignPassword(flagSet)
	address, err := utilsStruct.flagSetUtils.GetStringAddress(flagSet)
	if err != nil {
		return core.NilHash, err
	}

	name, err := utilsStruct.flagSetUtils.GetStringName(flagSet)
	if err != nil {
		return core.NilHash, err
	}

	url, err := utilsStruct.flagSetUtils.GetStringUrl(flagSet)
	if err != nil {
		return core.NilHash, err
	}

	selector, err := utilsStruct.flagSetUtils.GetStringSelector(flagSet)
	if err != nil {
		return core.NilHash, err
	}

	power, err := utilsStruct.flagSetUtils.GetInt8Power(flagSet)
	if err != nil {
		return core.NilHash, err
	}

	weight, err := utilsStruct.flagSetUtils.GetUint8Weight(flagSet)
	if err != nil {
		return core.NilHash, err
	}

	selectorType, err := utilsStruct.flagSetUtils.GetUint8SelectorType(flagSet)
	if err != nil {
		return core.NilHash, err
	}

	client := utilsStruct.razorUtils.ConnectToClient(config.Provider)
	txnArgs := types.TransactionOptions{
		Client:          client,
		Password:        password,
		AccountAddress:  address,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.AssetManagerAddress,
		MethodName:      "createJob",
		Parameters:      []interface{}{weight, power, selectorType, name, selector, url},
		ABI:             bindings.AssetManagerABI,
	}

	txnOpts := utilsStruct.razorUtils.GetTxnOpts(txnArgs, utilsStruct.packageUtils)
	log.Info("Creating Job...")
	txn, err := utilsStruct.assetManagerUtils.CreateJob(txnArgs.Client, txnOpts, weight, power, selectorType, name, selector, url)
	if err != nil {
		return core.NilHash, err
	}
	log.Info("Transaction Hash: ", utilsStruct.transactionUtils.Hash(txn))
	return utilsStruct.transactionUtils.Hash(txn), nil
}

func init() {

	razorUtils = Utils{}
	assetManagerUtils = AssetManagerUtils{}
	transactionUtils = TransactionUtils{}
	flagSetUtils = FlagSetUtils{}
	packageUtils = utils.RazorUtils{}

	rootCmd.AddCommand(createJobCmd)

	var (
		URL          string
		Selector     string
		SelectorType uint8
		Name         string
		Power        int8
		Account      string
		Password     string
		Weight       uint8
	)

	createJobCmd.Flags().StringVarP(&URL, "url", "u", "", "url of job")
	createJobCmd.Flags().StringVarP(&Selector, "selector", "s", "", "selector (jsonPath/XHTML selector)")
	createJobCmd.Flags().Uint8VarP(&SelectorType, "selectorType", "", 1, "selector type (1 for json, 2 for XHTML)")
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
