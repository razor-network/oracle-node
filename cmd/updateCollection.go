//Package cmd provides all functions related to command line
package cmd

import (
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/rpc"
	"razor/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var updateCollectionCmd = &cobra.Command{
	Use:   "updateCollection",
	Short: "[ADMIN ONLY]updateCollection can be used to update an existing collection",
	Long: `A collection is a group of jobs that reports the aggregated value of jobs. updateCollection can be used to modify an already existing collection.

Example: 
  ./razor updateCollection --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --aggregation 2 --power 3 --collectionId 3

Note: 
  This command only works for the admin.
`,
	Run: initialiseUpdateCollection,
}

//This function initialises the ExecuteUpdateCollection function
func initialiseUpdateCollection(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteUpdateCollection(cmd.Flags())
}

//This function sets the flag appropriately and executes the UpdateCollection function
func (*UtilsStruct) ExecuteUpdateCollection(flagSet *pflag.FlagSet) {
	config, rpcParameters, account, err := InitializeCommandDependencies(flagSet)
	utils.CheckError("Error in initialising command dependencies: ", err)

	collectionId, err := flagSetUtils.GetUint16CollectionId(flagSet)
	utils.CheckError("Error in getting collectionID: ", err)

	aggregation, err := flagSetUtils.GetUint32Aggregation(flagSet)
	utils.CheckError("Error in getting aggregation method: ", err)

	power, err := flagSetUtils.GetInt8Power(flagSet)
	utils.CheckError("Error in getting power: ", err)

	jobIdInUint, err := flagSetUtils.GetUintSliceJobIds(flagSet)
	utils.CheckError("Error in getting jobIds: ", err)

	tolerance, err := flagSetUtils.GetUint32Tolerance(flagSet)
	utils.CheckError("Error in getting tolerance: ", err)

	collectionInput := types.CreateCollectionInput{
		Aggregation: aggregation,
		Power:       power,
		JobIds:      jobIdInUint,
		Tolerance:   tolerance,
		Account:     account,
	}
	txn, err := cmdUtils.UpdateCollection(rpcParameters, config, collectionInput, collectionId)
	utils.CheckError("Update Collection error: ", err)
	err = razorUtils.WaitForBlockCompletion(rpcParameters, txn.Hex())
	utils.CheckError("Error in WaitForBlockCompletion for updateCollection: ", err)
}

//This function allows the admin to update an existing collection
func (*UtilsStruct) UpdateCollection(rpcParameters rpc.RPCParameters, config types.Configurations, collectionInput types.CreateCollectionInput, collectionId uint16) (common.Hash, error) {
	jobIds := utils.ConvertUintArrayToUint16Array(collectionInput.JobIds)
	log.Debug("UpdateCollection: Uint16 jobIds: ", jobIds)
	_, err := cmdUtils.WaitIfCommitState(rpcParameters, "update collection")
	if err != nil {
		log.Error("Error in fetching state")
		return core.NilHash, err
	}
	txnOpts, err := razorUtils.GetTxnOpts(rpcParameters, types.TransactionOptions{
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.CollectionManagerAddress,
		MethodName:      "updateCollection",
		Parameters:      []interface{}{collectionId, collectionInput.Tolerance, collectionInput.Aggregation, collectionInput.Power, jobIds},
		ABI:             bindings.CollectionManagerMetaData.ABI,
		Account:         collectionInput.Account,
	})
	if err != nil {
		return core.NilHash, err
	}
	log.Info("Updating collection...")
	client, err := rpcParameters.RPCManager.GetBestRPCClient()
	if err != nil {
		return core.NilHash, err
	}

	log.Debugf("Executing UpdateCollection transaction with collectionId = %d, tolerance = %d, aggregation method = %d, power = %d, jobIds = %v", collectionId, collectionInput.Tolerance, collectionInput.Aggregation, collectionInput.Power, jobIds)
	txn, err := assetManagerUtils.UpdateCollection(client, txnOpts, collectionId, collectionInput.Tolerance, collectionInput.Aggregation, collectionInput.Power, jobIds)
	if err != nil {
		log.Error("Error in updating collection")
		return core.NilHash, err
	}
	txnHash := transactionUtils.Hash(txn)
	log.Info("Txn Hash: ", txnHash.Hex())
	return txnHash, nil
}

func init() {
	rootCmd.AddCommand(updateCollectionCmd)

	var (
		Account           string
		CollectionId      uint16
		AggregationMethod uint32
		Password          string
		Power             int8
		JobIds            []uint
		Tolerance         uint32
	)

	updateCollectionCmd.Flags().StringVarP(&Account, "address", "a", "", "address of the job creator")
	updateCollectionCmd.Flags().Uint16VarP(&CollectionId, "collectionId", "", 0, "collection id to be modified")
	updateCollectionCmd.Flags().Uint32VarP(&AggregationMethod, "aggregation", "", 1, "aggregation method to be used")
	updateCollectionCmd.Flags().Int8VarP(&Power, "power", "", 0, "multiplier for the collection")
	updateCollectionCmd.Flags().StringVarP(&Password, "password", "", "", "password path of job creator to protect the keystore")
	updateCollectionCmd.Flags().UintSliceVarP(&JobIds, "jobIds", "", []uint{}, "job ids for the  collection")
	updateCollectionCmd.Flags().Uint32VarP(&Tolerance, "tolerance", "", 0, "tolerance")

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
