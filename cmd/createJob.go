package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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
	Run: initialiseCreateJob,
}

func initialiseCreateJob(cmd *cobra.Command, args []string) {
	cmdUtilsMockery.ExecuteCreateJob(cmd.Flags())
}

func (*UtilsStructMockery) ExecuteCreateJob(flagSet *pflag.FlagSet) {
	config, err := cmdUtilsMockery.GetConfigData()
	utils.CheckError("Error in getting config: ", err)

	password := razorUtilsMockery.AssignPassword(flagSet)
	address, err := flagSetUtilsMockery.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)

	name, err := flagSetUtilsMockery.GetStringName(flagSet)
	utils.CheckError("Error in getting name: ", err)

	url, err := flagSetUtilsMockery.GetStringUrl(flagSet)
	utils.CheckError("Error in getting url: ", err)

	selector, err := flagSetUtilsMockery.GetStringSelector(flagSet)
	utils.CheckError("Error in getting selector: ", err)

	power, err := flagSetUtilsMockery.GetInt8Power(flagSet)
	utils.CheckError("Error in getting power: ", err)

	weight, err := flagSetUtilsMockery.GetUint8Weight(flagSet)
	utils.CheckError("Error in getting weight: ", err)

	selectorType, err := flagSetUtilsMockery.GetUint8SelectorType(flagSet)
	utils.CheckError("Error in getting selectorType: ", err)

	client := razorUtilsMockery.ConnectToClient(config.Provider)

	jobInput := types.CreateJobInput{
		Address:      address,
		Password:     password,
		Url:          url,
		Name:         name,
		Selector:     selector,
		SelectorType: selectorType,
		Weight:       weight,
		Power:        power,
	}

	txn, err := cmdUtilsMockery.CreateJob(client, config, jobInput)
	utils.CheckError("CreateJob error: ", err)
	razorUtilsMockery.WaitForBlockCompletion(client, txn.String())
}

func (*UtilsStructMockery) CreateJob(client *ethclient.Client, config types.Configurations, jobInput types.CreateJobInput) (common.Hash, error) {
	txnArgs := types.TransactionOptions{
		Client:          client,
		Password:        jobInput.Password,
		AccountAddress:  jobInput.Address,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.AssetManagerAddress,
		MethodName:      "createJob",
		Parameters:      []interface{}{jobInput.Weight, jobInput.Power, jobInput.SelectorType, jobInput.Name, jobInput.Selector, jobInput.Url},
		ABI:             bindings.AssetManagerABI,
	}

	txnOpts := razorUtilsMockery.GetTxnOpts(txnArgs)
	log.Info("Creating Job...")
	txn, err := assetManagerUtilsMockery.CreateJob(txnArgs.Client, txnOpts, jobInput.Weight, jobInput.Power, jobInput.SelectorType, jobInput.Name, jobInput.Selector, jobInput.Url)
	if err != nil {
		return core.NilHash, err
	}
	log.Info("Transaction Hash: ", transactionUtilsMockery.Hash(txn))
	return transactionUtilsMockery.Hash(txn), nil
}

func init() {

	razorUtils = Utils{}
	assetManagerUtils = AssetManagerUtils{}
	transactionUtils = TransactionUtils{}
	flagSetUtils = FlagSetUtils{}
	cmdUtilsMockery = &UtilsStructMockery{}
	razorUtilsMockery = &UtilsMockery{}
	assetManagerUtilsMockery = &AssetManagerUtilsMockery{}
	transactionUtilsMockery = &TransactionUtilsMockery{}
	flagSetUtilsMockery = &FLagSetUtilsMockery{}

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
