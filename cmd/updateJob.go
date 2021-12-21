package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"
)

var updateJobCmd = &cobra.Command{
	Use:   "updateJob",
	Short: "updateJob can be used to update an existing job",
	Long: `A job consists of a URL and a selector to fetch the exact data from the URL. The updateJob command can be used to update an existing job that the stakers can vote upon.

Example:
  ./razor updateJob -a 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --jobId 1 -r true -s last --selectorType 1 -u https://api.gemini.com/v1/pubticker/btcusd

Note: 
  This command only works for the admin.
`,
	Run: func(cmd *cobra.Command, args []string) {
		utilsStruct := UtilsStruct{
			razorUtils:        razorUtils,
			assetManagerUtils: assetManagerUtils,
			transactionUtils:  transactionUtils,
			flagSetUtils:      flagSetUtils,
			cmdUtils:          cmdUtils,
			packageUtils:      packageUtils,
		}
		config, err := GetConfigData(utilsStruct)
		utils.CheckError("Error in getting config: ", err)
		txn, err := utilsStruct.updateJob(cmd.Flags(), config)
		utils.CheckError("UpdateJob error: ", err)
		utils.WaitForBlockCompletion(utils.ConnectToClient(config.Provider), txn.String())
	},
}

func (utilsStruct UtilsStruct) updateJob(flagSet *pflag.FlagSet, config types.Configurations) (common.Hash, error) {
	password := utilsStruct.razorUtils.AssignPassword(flagSet)
	address, err := utilsStruct.flagSetUtils.GetStringAddress(flagSet)
	if err != nil {
		return core.NilHash, err
	}
	jobId, err := utilsStruct.flagSetUtils.GetUint8JobId(flagSet)
	if err != nil {
		return core.NilHash, err
	}
	power, err := utilsStruct.flagSetUtils.GetInt8Power(flagSet)
	if err != nil {
		return core.NilHash, err
	}
	selector, err := utilsStruct.flagSetUtils.GetStringSelector(flagSet)
	if err != nil {
		return core.NilHash, err
	}
	url, err := utilsStruct.flagSetUtils.GetStringUrl(flagSet)
	if err != nil {
		return core.NilHash, err
	}
	weight, err := utilsStruct.flagSetUtils.GetUint8Weight(flagSet)
	if err != nil {
		return core.NilHash, err
	}
	client := utilsStruct.razorUtils.ConnectToClient(config.Provider)
	selectorType, err := utilsStruct.flagSetUtils.GetUint8SelectorType(flagSet)
	if err != nil {
		return core.NilHash, err
	}
	_, err = utilsStruct.cmdUtils.WaitIfCommitState(client, address, "update job", utilsStruct)
	if err != nil {
		log.Error("Error in fetching state")
		return core.NilHash, err
	}
	txnArgs := utilsStruct.razorUtils.GetTxnOpts(types.TransactionOptions{
		Client:          client,
		Password:        password,
		AccountAddress:  address,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.AssetManagerAddress,
		MethodName:      "updateJob",
		Parameters:      []interface{}{jobId, weight, power, selectorType, selector, url},
		ABI:             bindings.AssetManagerABI,
	})
	txn, err := utilsStruct.assetManagerUtils.UpdateJob(client, txnArgs, jobId, weight, power, selectorType, selector, url)
	if err != nil {
		return core.NilHash, err
	}
	return utilsStruct.transactionUtils.Hash(txn), nil
}

func init() {
	rootCmd.AddCommand(updateJobCmd)

	razorUtils = Utils{}
	assetManagerUtils = AssetManagerUtils{}
	transactionUtils = TransactionUtils{}
	flagSetUtils = FlagSetUtils{}
	cmdUtils = UtilsCmd{}

	var (
		JobId        uint8
		URL          string
		Selector     string
		SelectorType uint8
		Power        int8
		Weight       uint8
		Account      string
		Password     string
	)

	updateJobCmd.Flags().Uint8VarP(&JobId, "jobId", "", 0, "job id")
	updateJobCmd.Flags().StringVarP(&URL, "url", "u", "", "url of job")
	updateJobCmd.Flags().StringVarP(&Selector, "selector", "s", "", "selector (jsonPath/XHTML selector)")
	updateJobCmd.Flags().Uint8VarP(&SelectorType, "selectorType", "", 1, "selector type (1 for json, 2 for XHTML)")
	updateJobCmd.Flags().Int8VarP(&Power, "power", "", 0, "power")
	updateJobCmd.Flags().Uint8VarP(&Weight, "weight", "", 0, "weight")
	updateJobCmd.Flags().StringVarP(&Account, "address", "a", "", "address of the job creator")
	updateJobCmd.Flags().StringVarP(&Password, "password", "", "", "password path of job creator to protect the keystore")

	jobIdErr := updateJobCmd.MarkFlagRequired("jobId")
	utils.CheckError("Job Id error: ", jobIdErr)
	urlErr := updateJobCmd.MarkFlagRequired("url")
	utils.CheckError("URL error: ", urlErr)
	selectorErr := updateJobCmd.MarkFlagRequired("selector")
	utils.CheckError("Selector error: ", selectorErr)
	addrErr := updateJobCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
	powErr := updateJobCmd.MarkFlagRequired("power")
	utils.CheckError("Power error: ", powErr)
	weightErr := updateJobCmd.MarkFlagRequired("weight")
	utils.CheckError("Power error: ", weightErr)
}
