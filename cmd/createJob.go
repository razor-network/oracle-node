//Package cmd provides all functions related to command line
package cmd

import (
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/rpc"
	"razor/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/pflag"

	"github.com/spf13/cobra"
)

var createJobCmd = &cobra.Command{
	Use:   "createJob",
	Short: "[ADMIN ONLY]createJob can be used to create a job",
	Long: `A job consists of a URL and a selector to fetch the exact data from the URL. The createJob command can be used to create a job that the stakers can vote upon.

Example:
  ./razor createJob -a 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c -n btcusd_gemini -p 2 -s last --selectorType 1 -u https://api.gemini.com/v1/pubticker/btcusd

Note: 
  This command only works for the admin.
`,
	Run: initialiseCreateJob,
}

//This function initialises the ExecuteCreateJob function
func initialiseCreateJob(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteCreateJob(cmd.Flags())
}

//This function sets the flags appropriately and executes the CreateJob function
func (*UtilsStruct) ExecuteCreateJob(flagSet *pflag.FlagSet) {
	config, rpcParameters, account, err := InitializeCommandDependencies(flagSet)
	utils.CheckError("Error in initialising command dependencies: ", err)

	name, err := flagSetUtils.GetStringName(flagSet)
	utils.CheckError("Error in getting name: ", err)

	url, err := flagSetUtils.GetStringUrl(flagSet)
	utils.CheckError("Error in getting url: ", err)

	selector, err := flagSetUtils.GetStringSelector(flagSet)
	utils.CheckError("Error in getting selector: ", err)

	power, err := flagSetUtils.GetInt8Power(flagSet)
	utils.CheckError("Error in getting power: ", err)

	weight, err := flagSetUtils.GetUint8Weight(flagSet)
	utils.CheckError("Error in getting weight: ", err)

	selectorType, err := flagSetUtils.GetUint8SelectorType(flagSet)
	utils.CheckError("Error in getting selectorType: ", err)

	jobInput := types.CreateJobInput{
		Url:          url,
		Name:         name,
		Selector:     selector,
		SelectorType: selectorType,
		Weight:       weight,
		Power:        power,
		Account:      account,
	}

	txn, err := cmdUtils.CreateJob(rpcParameters, config, jobInput)
	utils.CheckError("CreateJob error: ", err)
	err = razorUtils.WaitForBlockCompletion(rpcParameters, txn.Hex())
	utils.CheckError("Error in WaitForBlockCompletion for createJob: ", err)
}

//This function allows the admin to create the job
func (*UtilsStruct) CreateJob(rpcParameters rpc.RPCParameters, config types.Configurations, jobInput types.CreateJobInput) (common.Hash, error) {
	txnArgs := types.TransactionOptions{
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.CollectionManagerAddress,
		MethodName:      "createJob",
		Parameters:      []interface{}{jobInput.Weight, jobInput.Power, jobInput.SelectorType, jobInput.Name, jobInput.Selector, jobInput.Url},
		ABI:             bindings.CollectionManagerMetaData.ABI,
		Account:         jobInput.Account,
	}

	txnOpts, err := razorUtils.GetTxnOpts(rpcParameters, txnArgs)
	if err != nil {
		return core.NilHash, err
	}
	log.Info("Creating Job...")

	client, err := rpcParameters.RPCManager.GetBestRPCClient()
	if err != nil {
		return core.NilHash, err
	}

	log.Debugf("CreateJob: Executing CreateJob transaction with weight = %d, power = %d, selector type = %d, name = %s, selector = %s, URl = %s", jobInput.Weight, jobInput.Power, jobInput.SelectorType, jobInput.Name, jobInput.Selector, jobInput.Url)
	txn, err := assetManagerUtils.CreateJob(client, txnOpts, jobInput.Weight, jobInput.Power, jobInput.SelectorType, jobInput.Name, jobInput.Selector, jobInput.Url)
	if err != nil {
		return core.NilHash, err
	}
	txnHash := transactionUtils.Hash(txn)
	log.Info("Txn Hash: ", txnHash.Hex())
	return txnHash, nil
}

func init() {
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
	createJobCmd.Flags().Uint8VarP(&SelectorType, "selectorType", "", 0, "selector type (0 for json, 1 for XHTML)")
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
