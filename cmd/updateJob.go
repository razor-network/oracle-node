//Package cmd provides all functions related to command line
package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"razor/core"
	"razor/core/types"
	"razor/logger"
	"razor/pkg/bindings"
	"razor/utils"
)

var updateJobCmd = &cobra.Command{
	Use:   "updateJob",
	Short: "[ADMIN ONLY]updateJob can be used to update an existing job",
	Long: `A job consists of a URL and a selector to fetch the exact data from the URL. The updateJob command can be used to update an existing job that the stakers can vote upon.

Example:
  ./razor updateJob -a 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --jobId 1 -r true -s last --selectorType 1 -u https://api.gemini.com/v1/pubticker/btcusd

Note: 
  This command only works for the admin.
`,
	Run: initialiseUpdateJob,
}

//This function initialises the ExecuteUpdateJob function
func initialiseUpdateJob(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteUpdateJob(cmd.Flags())
}

//This function sets the flag appropriately and executes the UpdateJob function
func (*UtilsStruct) ExecuteUpdateJob(flagSet *pflag.FlagSet) {
	razorUtils.AssignLogFile(flagSet)
	address, err := flagSetUtils.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)

	logger.Address = address

	config, err := cmdUtils.GetConfigData()
	utils.CheckError("Error in getting config: ", err)

	password := razorUtils.AssignPassword()

	jobId, err := flagSetUtils.GetUint16JobId(flagSet)
	utils.CheckError("Error in getting jobId: ", err)

	power, err := flagSetUtils.GetInt8Power(flagSet)
	utils.CheckError("Error in getting power: ", err)

	selector, err := flagSetUtils.GetStringSelector(flagSet)
	utils.CheckError("Error in getting selector: ", err)

	url, err := flagSetUtils.GetStringUrl(flagSet)
	utils.CheckError("Error in getting url: ", err)

	weight, err := flagSetUtils.GetUint8Weight(flagSet)
	utils.CheckError("Error in getting weight: ", err)

	selectorType, err := flagSetUtils.GetUint8SelectorType(flagSet)
	utils.CheckError("Error in getting selector type: ", err)

	jobInput := types.CreateJobInput{
		Address:      address,
		Password:     password,
		Power:        power,
		Selector:     selector,
		Url:          url,
		Weight:       weight,
		SelectorType: selectorType,
	}

	client := razorUtils.ConnectToClient(config.Provider)

	txn, err := cmdUtils.UpdateJob(client, config, jobInput, jobId)
	utils.CheckError("UpdateJob error: ", err)
	err = razorUtils.WaitForBlockCompletion(client, txn.String())
	utils.CheckError("Error in WaitForBlockCompletion for updateJob: ", err)
}

//This function allows the admin to update an existing job
func (*UtilsStruct) UpdateJob(client *ethclient.Client, config types.Configurations, jobInput types.CreateJobInput, jobId uint16) (common.Hash, error) {

	_, err := cmdUtils.WaitIfCommitState(client, "update job")
	if err != nil {
		log.Error("Error in fetching state")
		return core.NilHash, err
	}
	txnArgs := razorUtils.GetTxnOpts(types.TransactionOptions{
		Client:          client,
		Password:        jobInput.Password,
		AccountAddress:  jobInput.Address,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.CollectionManagerAddress,
		MethodName:      "updateJob",
		Parameters:      []interface{}{jobId, jobInput.Weight, jobInput.Power, jobInput.SelectorType, jobInput.Selector, jobInput.Url},
		ABI:             bindings.CollectionManagerABI,
	})
	txn, err := assetManagerUtils.UpdateJob(client, txnArgs, jobId, jobInput.Weight, jobInput.Power, jobInput.SelectorType, jobInput.Selector, jobInput.Url)
	if err != nil {
		return core.NilHash, err
	}
	return transactionUtils.Hash(txn), nil
}

func init() {
	rootCmd.AddCommand(updateJobCmd)

	var (
		JobId        uint16
		URL          string
		Selector     string
		SelectorType uint8
		Power        int8
		Weight       uint8
		Account      string
	)

	updateJobCmd.Flags().Uint16VarP(&JobId, "jobId", "", 0, "job id")
	updateJobCmd.Flags().StringVarP(&URL, "url", "u", "", "url of job")
	updateJobCmd.Flags().StringVarP(&Selector, "selector", "s", "", "selector (jsonPath/XHTML selector)")
	updateJobCmd.Flags().Uint8VarP(&SelectorType, "selectorType", "", 1, "selector type (1 for json, 2 for XHTML)")
	updateJobCmd.Flags().Int8VarP(&Power, "power", "", 0, "power")
	updateJobCmd.Flags().Uint8VarP(&Weight, "weight", "", 0, "weight")
	updateJobCmd.Flags().StringVarP(&Account, "address", "a", "", "address of the job creator")

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
