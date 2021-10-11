package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/pflag"
	"razor/core"
	"razor/core/types"
	"razor/utils"

	"github.com/spf13/cobra"
)

var addJobToCollectionCmd = &cobra.Command{
	Use:   "addJobToCollection",
	Short: "addJobToCollection can be used to add a particular job to an existing collection",
	Long: `If there are existing jobs and collections, this command can be used to add a job to a collection.

Example: 
  ./razor addJobToCollection --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --collectionId 6 --jobId 7 

Note: 
  This command only works for the admin.
`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		if err != nil {
			log.Fatal("Error in getting config: ", err)
		}
		txn, err := addJobToCollection(cmd.Flags(), config, razorUtils, assetManagerUtils, transactionUtils, flagSetUtils)
		utils.CheckError("AddJobToCollection error: ", err)
		utils.WaitForBlockCompletion(utils.ConnectToClient(config.Provider), txn.String())
	},
}

func addJobToCollection(flagSet *pflag.FlagSet, config types.Configurations, razorUtils utilsInterface, assetManagerUtils assetManagerInterface, transactionUtils transactionInterface, flagSetUtils flagSetInterface) (common.Hash, error) {

	password := razorUtils.AssignPassword(flagSet)
	address, err := flagSetUtils.GetStringAddress(flagSet)
	if err != nil {
		return core.NilHash, err
	}
	jobId, err := flagSetUtils.GetUint8JobId(flagSet)
	if err != nil {
		return core.NilHash, err
	}
	collectionId, err := flagSetUtils.GetUint8CollectionId(flagSet)
	if err != nil {
		return core.NilHash, err
	}
	client := razorUtils.ConnectToClient(config.Provider)

	txnOpts := razorUtils.GetTxnOpts(types.TransactionOptions{
		Client:         client,
		Password:       password,
		AccountAddress: address,
		ChainId:        core.ChainId,
		Config:         config,
	})

	log.Infof("Adding Job %d to collection %d", jobId, collectionId)
	txn, err := assetManagerUtils.AddJobToCollection(client, txnOpts, collectionId, jobId)
	if err != nil {
		return core.NilHash, err
	}
	log.Info("Txn Hash: ", transactionUtils.Hash(txn))
	return transactionUtils.Hash(txn), nil
}

func init() {

	razorUtils = Utils{}
	assetManagerUtils = AssetManagerUtils{}
	transactionUtils = TransactionUtils{}
	flagSetUtils = FlagSetUtils{}

	rootCmd.AddCommand(addJobToCollectionCmd)

	var (
		Account      string
		JobId        uint8
		CollectionId uint8
	)

	addJobToCollectionCmd.Flags().StringVarP(&Account, "address", "a", "", "address of the job creator")
	addJobToCollectionCmd.Flags().Uint8VarP(&JobId, "jobId", "", 0, "job id to add to the  collection")
	addJobToCollectionCmd.Flags().Uint8VarP(&CollectionId, "collectionId", "", 0, "collection id")

	addrErr := addJobToCollectionCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
	jobIdErr := addJobToCollectionCmd.MarkFlagRequired("jobId")
	utils.CheckError("Job Id error: ", jobIdErr)
	collectionIdErr := addJobToCollectionCmd.MarkFlagRequired("collectionId")
	utils.CheckError("Collection Id error: ", collectionIdErr)
}
