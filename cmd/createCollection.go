package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"
)

var createCollectionCmd = &cobra.Command{
	Use:   "createCollection",
	Short: "createCollection can be used to create collections if existing jobs are present",
	Long: `A collection is a group of jobs that reports the aggregated value of jobs. createCollection can be used to club multiple jobs into one collection bound by an aggregation method.

Example: 
  ./razor createCollection --name btcCollectionMean --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --jobIds 1,2 --aggregation 2 --power 2

Note: 
  This command only works for the admin.
`,
	Run: initialiseCreateCollection,
}

func initialiseCreateCollection(cmd *cobra.Command, args []string) {
	cmdUtilsMockery.ExecuteCreateCollection(cmd.Flags())
}

func (*UtilsStructMockery) ExecuteCreateCollection(flagSet *pflag.FlagSet) {
	config, err := cmdUtilsMockery.GetConfigData()
	utils.CheckError("Error in getting config: ", err)

	password := razorUtilsMockery.AssignPassword(flagSet)
	name, err := flagSetUtilsMockery.GetStringName(flagSet)
	utils.CheckError("Error in getting name: ", err)

	address, err := flagSetUtilsMockery.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)

	jobIdInUint, err := flagSetUtilsMockery.GetUintSliceJobIds(flagSet)
	utils.CheckError("Error in getting jobId: ", err)

	aggregation, err := flagSetUtilsMockery.GetUint32Aggregation(flagSet)
	utils.CheckError("Error in getting aggregation method: ", err)

	power, err := flagSetUtilsMockery.GetInt8Power(flagSet)
	utils.CheckError("Error in getting power: ", err)

	client := razorUtilsMockery.ConnectToClient(config.Provider)

	collectionInput := types.CreateCollectionInput{
		Address:     address,
		Password:    password,
		Power:       power,
		Name:        name,
		Aggregation: aggregation,
		JobIds:      jobIdInUint,
	}

	txn, err := cmdUtilsMockery.CreateCollection(client, config, collectionInput)
	utils.CheckError("CreateCollection error: ", err)
	razorUtilsMockery.WaitForBlockCompletion(client, txn.String())
}

func (*UtilsStructMockery) CreateCollection(client *ethclient.Client, config types.Configurations, collectionInput types.CreateCollectionInput) (common.Hash, error) {
	jobIds := razorUtilsMockery.ConvertUintArrayToUint16Array(collectionInput.JobIds)
	_, err := cmdUtilsMockery.WaitForAppropriateState(client, "create collection", 4)
	if err != nil {
		log.Error("Error in fetching state")
		return core.NilHash, err
	}
	txnOpts := razorUtilsMockery.GetTxnOpts(types.TransactionOptions{
		Client:          client,
		Password:        collectionInput.Password,
		AccountAddress:  collectionInput.Address,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.AssetManagerAddress,
		MethodName:      "createCollection",
		Parameters:      []interface{}{jobIds, collectionInput.Aggregation, collectionInput.Power, collectionInput.Name},
		ABI:             bindings.AssetManagerABI,
	})
	txn, err := assetManagerUtilsMockery.CreateCollection(client, txnOpts, jobIds, collectionInput.Aggregation, collectionInput.Power, collectionInput.Name)
	if err != nil {
		log.Error("Error in creating collection")
		return core.NilHash, err
	}
	log.Info("Creating collection...")
	log.Info("Txn Hash: ", transactionUtilsMockery.Hash(txn))
	return transactionUtilsMockery.Hash(txn), nil
}

func init() {

	razorUtils = Utils{}
	assetManagerUtils = AssetManagerUtils{}
	transactionUtils = TransactionUtils{}
	flagSetUtils = FlagSetUtils{}
	cmdUtils = UtilsCmd{}
	utils.Options = &utils.OptionsStruct{}
	utils.UtilsInterface = &utils.UtilsStruct{}
	cmdUtilsMockery = &UtilsStructMockery{}
	razorUtilsMockery = UtilsMockery{}
	assetManagerUtilsMockery = AssetManagerUtilsMockery{}
	transactionUtilsMockery = TransactionUtilsMockery{}
	flagSetUtilsMockery = FLagSetUtilsMockery{}

	rootCmd.AddCommand(createCollectionCmd)

	var (
		Name              string
		Account           string
		JobIds            []uint
		AggregationMethod uint32
		Password          string
		Power             int8
	)

	createCollectionCmd.Flags().StringVarP(&Name, "name", "n", "", "name of the collection")
	createCollectionCmd.Flags().StringVarP(&Account, "address", "a", "", "address of the job creator")
	createCollectionCmd.Flags().UintSliceVarP(&JobIds, "jobIds", "", []uint{}, "job ids for the  collection")
	createCollectionCmd.Flags().Uint32VarP(&AggregationMethod, "aggregation", "", 1, "aggregation method to be used")
	createCollectionCmd.Flags().Int8VarP(&Power, "power", "", 0, "multiplier for the collection")
	createCollectionCmd.Flags().StringVarP(&Password, "password", "", "", "password path of job creator to protect the keystore")

	nameErr := createCollectionCmd.MarkFlagRequired("name")
	utils.CheckError("Name error: ", nameErr)
	addrErr := createCollectionCmd.MarkFlagRequired("address")
	utils.CheckError("Address Error: ", addrErr)
	jobIdErr := createCollectionCmd.MarkFlagRequired("jobIds")
	utils.CheckError("Job Id Error: ", jobIdErr)
	powerErr := createCollectionCmd.MarkFlagRequired("power")
	utils.CheckError("Power Error: ", powerErr)
}
