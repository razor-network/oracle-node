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

var updateCollectionCmd = &cobra.Command{
	Use:   "updateCollection",
	Short: "updateCollection can be used to update an existing collection",
	Long: `A collection is a group of jobs that reports the aggregated value of jobs. updateCollection can be used to modify an already existing collection.

Example: 
  ./razor updateCollection --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --aggregation 2 --power 3 --collectionId 3

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
			keystoreUtils:     keystoreUtils,
		}
		config, err := GetConfigData()
		utils.CheckError("Error in getting config: ", err)

		txn, err := utilsStruct.updateCollection(cmd.Flags(), config)
		utils.CheckError("Update Collection error: ", err)
		utils.WaitForBlockCompletion(utils.ConnectToClient(config.Provider), txn.String())
	},
}

func (utilsStruct UtilsStruct) updateCollection(flagSet *pflag.FlagSet, config types.Configurations) (common.Hash, error) {
	password := utilsStruct.razorUtils.AssignPassword(flagSet)
	address, err := utilsStruct.flagSetUtils.GetStringAddress(flagSet)
	if err != nil {
		return core.NilHash, err
	}
	collectionId, err := utilsStruct.flagSetUtils.GetUint8CollectionId(flagSet)
	if err != nil {
		return core.NilHash, err
	}
	aggregation, err := utilsStruct.flagSetUtils.GetUint32Aggregation(flagSet)
	if err != nil {
		return core.NilHash, err
	}
	power, err := utilsStruct.flagSetUtils.GetInt8Power(flagSet)
	if err != nil {
		return core.NilHash, err
	}
	jobIdInUint, err := utilsStruct.flagSetUtils.GetUintSliceJobIds(flagSet)
	if err != nil {
		return core.NilHash, err
	}
	jobIds := utilsStruct.razorUtils.ConvertUintArrayToUint8Array(jobIdInUint)
	client := utilsStruct.razorUtils.ConnectToClient(config.Provider)
	_, err = utilsStruct.razorUtils.WaitIfCommitState(client, address, "update collection")
	if err != nil {
		log.Error("Error in fetching state")
		return core.NilHash, err
	}
	txnOpts := utilsStruct.razorUtils.GetTxnOpts(types.TransactionOptions{
		Client:          client,
		Password:        password,
		AccountAddress:  address,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.AssetManagerAddress,
		MethodName:      "updateCollection",
		Parameters:      []interface{}{collectionId, aggregation, power},
		ABI:             bindings.AssetManagerABI,
	}, utilsStruct)

	txn, err := utilsStruct.assetManagerUtils.UpdateCollection(client, txnOpts, collectionId, aggregation, power, jobIds)
	if err != nil {
		log.Error("Error in updating collection")
		return core.NilHash, err
	}
	log.Info("Updating collection...")
	log.Info("Txn Hash: ", utilsStruct.transactionUtils.Hash(txn))
	return utilsStruct.transactionUtils.Hash(txn), nil
}

func init() {
	razorUtils = Utils{}
	assetManagerUtils = AssetManagerUtils{}
	transactionUtils = TransactionUtils{}
	flagSetUtils = FlagSetUtils{}
	cmdUtils = UtilsCmd{}
	keystoreUtils = KeystoreUtils{}

	rootCmd.AddCommand(updateCollectionCmd)

	var (
		Account           string
		CollectionId      uint8
		AggregationMethod uint32
		Password          string
		Power             int8
		JobIds            []uint
	)

	updateCollectionCmd.Flags().StringVarP(&Account, "address", "a", "", "address of the job creator")
	updateCollectionCmd.Flags().Uint8VarP(&CollectionId, "collectionId", "", 0, "collection id to be modified")
	updateCollectionCmd.Flags().Uint32VarP(&AggregationMethod, "aggregation", "", 1, "aggregation method to be used")
	updateCollectionCmd.Flags().Int8VarP(&Power, "power", "", 0, "multiplier for the collection")
	updateCollectionCmd.Flags().StringVarP(&Password, "password", "", "", "password path of job creator to protect the keystore")
	updateCollectionCmd.Flags().UintSliceVarP(&JobIds, "jobIds", "", []uint{}, "job ids for the  collection")

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
