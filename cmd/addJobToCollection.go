package cmd

import (
	log "github.com/sirupsen/logrus"
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

		password := utils.PasswordPrompt()

		address, _ := cmd.Flags().GetString("address")
		jobId, _ := cmd.Flags().GetUint8("jobId")
		collectionId, _ := cmd.Flags().GetUint8("collectionId")

		client := utils.ConnectToClient(config.Provider)

		txnOpts := utils.GetTxnOpts(types.TransactionOptions{
			Client:         client,
			Password:       password,
			AccountAddress: address,
			ChainId:        core.ChainId,
			Config:         config,
		})

		assetManager := utils.GetAssetManager(client)
		log.Infof("Adding Job %d to collection %d", jobId, collectionId)
		txn, err := assetManager.AddJobToCollection(txnOpts, collectionId, jobId)
		if err != nil {
			log.Fatal(err)
		}
		utils.WaitForBlockCompletion(client, txn.Hash().String())
	},
}

func init() {
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
