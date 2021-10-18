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

var removeJobFromCollectionCmd = &cobra.Command{
	Use:   "removeJobFromCollection",
	Short: "removeJobFromCollection can be used to remove a particular job from an existing collection",
	Long: `A collection is a group of jobs that reports the aggregated value of jobs. updateCollection can be used to modify an already existing collection.
Example: 
  ./razor removeJobFromCollection --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --collectionId 4 --jobId 3 
Note: 
  This command only works for the admin.
`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		utils.CheckError("Error in getting config: ", err)

		txn, err := removeJobFromCollection(cmd.Flags(), config, razorUtils, assetManagerUtils, transactionUtils, flagSetUtils)
		utils.CheckError("removeJobFromCollection error: ", err)
		utils.WaitForBlockCompletion(utils.ConnectToClient(config.Provider), txn.String())
	},
}

func removeJobFromCollection(flagSet *pflag.FlagSet, config types.Configurations, razorUtils utilsInterface, assetManagerUtils assetManagerInterface, transactionUtils transactionInterface, flagSetUtils flagSetInterface) (common.Hash, error) {
	password := razorUtils.AssignPassword(flagSet)
	address, err := flagSetUtils.GetStringAddress(flagSet)
	if err != nil {
		return core.NilHash, err
	}
	collectionId, err := flagSetUtils.GetUint8CollectionId(flagSet)
	if err != nil {
		return core.NilHash, err
	}
	jobId, err := flagSetUtils.GetUint8JobId(flagSet)
	if err != nil {
		return core.NilHash, err
	}

	client := razorUtils.ConnectToClient(config.Provider)

	txnOpts := razorUtils.GetTxnOpts(types.TransactionOptions{
		Client:          client,
		Password:        password,
		AccountAddress:  address,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.AssetManagerAddress,
		MethodName:      "removeJobFromCollection",
		Parameters:      []interface{}{collectionId, jobId},
		ABI:             bindings.AssetManagerABI,
	})

	txn, err := assetManagerUtils.RemoveJobFromCollection(client, txnOpts, collectionId, jobId)
	if err != nil {
		log.Error("Error in removing job from collection")
		return core.NilHash, err
	}
	log.Info("Removing job from collection...")
	log.Info("Txn Hash: ", transactionUtils.Hash(txn))
	return transactionUtils.Hash(txn), nil
}

func init() {
	razorUtils = Utils{}
	assetManagerUtils = AssetManagerUtils{}
	transactionUtils = TransactionUtils{}
	flagSetUtils = FlagSetUtils{}

	rootCmd.AddCommand(removeJobFromCollectionCmd)

	var (
		Account      string
		CollectionId uint8
		JobId        uint8
		Password     string
	)

	removeJobFromCollectionCmd.Flags().StringVarP(&Account, "address", "a", "", "address of the job creator")
	removeJobFromCollectionCmd.Flags().Uint8VarP(&CollectionId, "collectionId", "", 0, "collection id to be modified")
	removeJobFromCollectionCmd.Flags().Uint8VarP(&JobId, "jobId", "", 0, "job id to be removed")
	removeJobFromCollectionCmd.Flags().StringVarP(&Password, "password", "", "", "password path of job creator to protect the keystore")

	collectionIdErr := removeJobFromCollectionCmd.MarkFlagRequired("collectionId")
	utils.CheckError("Collection Id error: ", collectionIdErr)
	addrErr := removeJobFromCollectionCmd.MarkFlagRequired("address")
	utils.CheckError("Address Error: ", addrErr)
	jobIdErr := removeJobFromCollectionCmd.MarkFlagRequired("jobId")
	utils.CheckError("JobId Error: ", jobIdErr)

}
