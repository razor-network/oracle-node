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

var updateCollectionCmd = &cobra.Command{
	Use:   "updateCollection",
	Short: "updateCollection can be used to update an existing collection",
	Long: `A collection is a group of jobs that reports the aggregated value of jobs. updateCollection can be used to modify an already existing collection.

Example: 
  ./razor updateCollection --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --aggregation 2 --power 3 --collectionId 3

Note: 
  This command only works for the admin.
`,
	Run: initialiseUpdateCollection,
}

func initialiseUpdateCollection(cmd *cobra.Command, args []string) {
	cmdUtilsMockery.ExecuteUpdateCollection(cmd.Flags())
}

func (*UtilsStructMockery) ExecuteUpdateCollection(flagSet *pflag.FlagSet) {
	config, err := cmdUtilsMockery.GetConfigData()
	utils.CheckError("Error in getting config: ", err)

	password := razorUtilsMockery.AssignPassword(flagSet)
	address, err := flagSetUtilsMockery.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)

	collectionId, err := flagSetUtilsMockery.GetUint16CollectionId(flagSet)
	utils.CheckError("Error in getting collectionID: ", err)

	aggregation, err := flagSetUtilsMockery.GetUint32Aggregation(flagSet)
	utils.CheckError("Error in getting aggregation method: ", err)

	power, err := flagSetUtilsMockery.GetInt8Power(flagSet)
	utils.CheckError("Error in getting power: ", err)

	jobIdInUint, err := flagSetUtilsMockery.GetUintSliceJobIds(flagSet)
	utils.CheckError("Error in getting jobIds: ", err)

	client := razorUtilsMockery.ConnectToClient(config.Provider)

	tolerance, err := flagSetUtilsMockery.GetUint16Tolerance(flagSet)
	utils.CheckError("Error in getting tolerance: ", err)

	collectionInput := types.CreateCollectionInput{
		Address:     address,
		Password:    password,
		Aggregation: aggregation,
		Power:       power,
		JobIds:      jobIdInUint,
		Tolerance:   tolerance,
	}
	txn, err := cmdUtilsMockery.UpdateCollection(client, config, collectionInput, collectionId)
	utils.CheckError("Update Collection error: ", err)
	razorUtilsMockery.WaitForBlockCompletion(client, txn.String())
}

func (*UtilsStructMockery) UpdateCollection(client *ethclient.Client, config types.Configurations, collectionInput types.CreateCollectionInput, collectionId uint16) (common.Hash, error) {
	jobIds := razorUtilsMockery.ConvertUintArrayToUint16Array(collectionInput.JobIds)
	_, err := cmdUtilsMockery.WaitIfCommitState(client, "update collection")
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
		MethodName:      "updateCollection",
		Parameters:      []interface{}{collectionId, collectionInput.Tolerance, collectionInput.Aggregation, collectionInput.Power, jobIds},
		ABI:             bindings.AssetManagerABI,
	})
	txn, err := assetManagerUtilsMockery.UpdateCollection(client, txnOpts, collectionId, collectionInput.Tolerance, collectionInput.Aggregation, collectionInput.Power, jobIds)
	if err != nil {
		log.Error("Error in updating collection")
		return core.NilHash, err
	}
	log.Info("Updating collection...")
	log.Info("Txn Hash: ", transactionUtilsMockery.Hash(txn))
	return transactionUtilsMockery.Hash(txn), nil
}

func init() {
	razorUtilsMockery = UtilsMockery{}
	assetManagerUtilsMockery = AssetManagerUtilsMockery{}
	transactionUtilsMockery = TransactionUtilsMockery{}
	flagSetUtilsMockery = FLagSetUtilsMockery{}
	cmdUtilsMockery = &UtilsStructMockery{}

	rootCmd.AddCommand(updateCollectionCmd)

	var (
		Account           string
		CollectionId      uint16
		AggregationMethod uint32
		Password          string
		Power             int8
		JobIds            []uint
		Tolerance         uint16
	)

	updateCollectionCmd.Flags().StringVarP(&Account, "address", "a", "", "address of the job creator")
	updateCollectionCmd.Flags().Uint16VarP(&CollectionId, "collectionId", "", 0, "collection id to be modified")
	updateCollectionCmd.Flags().Uint32VarP(&AggregationMethod, "aggregation", "", 1, "aggregation method to be used")
	updateCollectionCmd.Flags().Int8VarP(&Power, "power", "", 0, "multiplier for the collection")
	updateCollectionCmd.Flags().StringVarP(&Password, "password", "", "", "password path of job creator to protect the keystore")
	updateCollectionCmd.Flags().UintSliceVarP(&JobIds, "jobIds", "", []uint{}, "job ids for the  collection")
	updateCollectionCmd.Flags().Uint16VarP(&Tolerance, "tolerance", "", 0, "tolerance")

	collectionIdErr := updateCollectionCmd.MarkFlagRequired("collectionId")
	utils.CheckError("Collection Id error: ", collectionIdErr)
	addrErr := updateCollectionCmd.MarkFlagRequired("address")
	utils.CheckError("Address Error: ", addrErr)
	powerErr := updateCollectionCmd.MarkFlagRequired("power")
	utils.CheckError("Power Error: ", powerErr)
	aggregationMethodErr := updateCollectionCmd.MarkFlagRequired("aggregation")
	utils.CheckError("Aggregation Method Error: ", aggregationMethodErr)
	jobIdErr := updateCollectionCmd.MarkFlagRequired("jobIds")
	utils.CheckError("Job Id Error: ", jobIdErr)
}
