//Package cmd provides all functions related to command line
package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/logger"
	"razor/pkg/bindings"
	"razor/utils"
)

var createCollectionCmd = &cobra.Command{
	Use:   "createCollection",
	Short: "[ADMIN ONLY]createCollection can be used to create collections if existing jobs are present",
	Long: `A collection is a group of jobs that reports the aggregated value of jobs. createCollection can be used to club multiple jobs into one collection bound by an aggregation method.

Example: 
  ./razor createCollection --name btcCollectionMean --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --jobIds 1,2 --aggregation 2 --power 2

Note: 
  This command only works for the admin.
`,
	Run: initialiseCreateCollection,
}

//This function initialises the ExecuteCreateCollction function
func initialiseCreateCollection(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteCreateCollection(cmd.Flags())
}

//This function sets the flags appropriately and executes the CreateCollection function
func (*UtilsStruct) ExecuteCreateCollection(flagSet *pflag.FlagSet) {
	razorUtils.AssignLogFile(flagSet)
	address, err := flagSetUtils.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)

	logger.Address = address

	config, err := cmdUtils.GetConfigData()
	utils.CheckError("Error in getting config: ", err)

	password := razorUtils.AssignPassword(flagSet)
	name, err := flagSetUtils.GetStringName(flagSet)
	utils.CheckError("Error in getting name: ", err)

	jobIdInUint, err := flagSetUtils.GetUintSliceJobIds(flagSet)
	utils.CheckError("Error in getting jobId: ", err)

	aggregation, err := flagSetUtils.GetUint32Aggregation(flagSet)
	utils.CheckError("Error in getting aggregation method: ", err)

	power, err := flagSetUtils.GetInt8Power(flagSet)
	utils.CheckError("Error in getting power: ", err)

	client := razorUtils.ConnectToClient(config.Provider)

	tolerance, err := flagSetUtils.GetUint32Tolerance(flagSet)
	utils.CheckError("Error in getting tolerance: ", err)

	collectionInput := types.CreateCollectionInput{
		Address:     address,
		Password:    password,
		Power:       power,
		Name:        name,
		Aggregation: aggregation,
		JobIds:      jobIdInUint,
		Tolerance:   tolerance,
	}

	txn, err := cmdUtils.CreateCollection(client, config, collectionInput)
	utils.CheckError("CreateCollection error: ", err)
	razorUtils.WaitForBlockCompletion(client, txn.String())
}

//This function allows the admin to create collction if existing jobs are present
func (*UtilsStruct) CreateCollection(client *ethclient.Client, config types.Configurations, collectionInput types.CreateCollectionInput) (common.Hash, error) {
	jobIds := razorUtils.ConvertUintArrayToUint16Array(collectionInput.JobIds)
	_, err := cmdUtils.WaitForAppropriateState(client, "create collection", 4)
	if err != nil {
		log.Error("Error in fetching state")
		return core.NilHash, err
	}
	txnOpts := razorUtils.GetTxnOpts(types.TransactionOptions{
		Client:          client,
		Password:        collectionInput.Password,
		AccountAddress:  collectionInput.Address,
		ChainId:         big.NewInt(config.ChainId),
		Config:          config,
		ContractAddress: core.CollectionManagerAddress,
		MethodName:      "createCollection",
		Parameters:      []interface{}{collectionInput.Tolerance, collectionInput.Power, collectionInput.Aggregation, jobIds, collectionInput.Name},
		ABI:             bindings.CollectionManagerABI,
	})
	txn, err := assetManagerUtils.CreateCollection(client, txnOpts, collectionInput.Tolerance, collectionInput.Power, collectionInput.Aggregation, jobIds, collectionInput.Name)
	if err != nil {
		log.Error("Error in creating collection")
		return core.NilHash, err
	}
	log.Info("Creating collection...")
	log.Info("Txn Hash: ", transactionUtils.Hash(txn))
	return transactionUtils.Hash(txn), nil
}

func init() {
	rootCmd.AddCommand(createCollectionCmd)

	var (
		Name              string
		Account           string
		JobIds            []uint
		AggregationMethod uint32
		Password          string
		Power             int8
		Tolerance         uint32
	)

	createCollectionCmd.Flags().StringVarP(&Name, "name", "n", "", "name of the collection")
	createCollectionCmd.Flags().StringVarP(&Account, "address", "a", "", "address of the job creator")
	createCollectionCmd.Flags().UintSliceVarP(&JobIds, "jobIds", "", []uint{}, "job ids for the  collection")
	createCollectionCmd.Flags().Uint32VarP(&AggregationMethod, "aggregation", "", 1, "aggregation method to be used")
	createCollectionCmd.Flags().Uint32VarP(&Tolerance, "tolerance", "", 0, "tolerance")
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
	toleranceErr := createCollectionCmd.MarkFlagRequired("tolerance")
	utils.CheckError("Tolerance Error: ", toleranceErr)
}
