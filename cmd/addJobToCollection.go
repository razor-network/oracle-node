package cmd

import (
	log "github.com/sirupsen/logrus"
	"math/big"
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
		jobId, _ := cmd.Flags().GetString("jobId")
		collectionId, _ := cmd.Flags().GetString("collectionId")

		client := utils.ConnectToClient(config.Provider)

		jobIdInBigInt, ok := new(big.Int).SetString(jobId, 10)
		if !ok {
			log.Fatal("SetString: error")
		}

		collectionIdInBigInt, ok := new(big.Int).SetString(collectionId, 10)
		if !ok {
			log.Fatal("SetString: error")
		}

		txnOpts := utils.GetTxnOpts(types.TransactionOptions{
			Client:         client,
			Password:       password,
			AccountAddress: address,
			ChainId:        core.ChainId,
			Config:         config,
		})

		assetManager := utils.GetAssetManager(client)
		log.Infof("Adding Job %s to collection %s", jobIdInBigInt, collectionIdInBigInt)
		txn, err := assetManager.AddJobToCollection(txnOpts, collectionIdInBigInt, jobIdInBigInt)
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
		JobId        string
		CollectionId string
	)

	addJobToCollectionCmd.Flags().StringVarP(&Account, "address", "a", "", "address of the job creator")
	addJobToCollectionCmd.Flags().StringVarP(&JobId, "jobId", "", "", "job id to add to the  collection")
	addJobToCollectionCmd.Flags().StringVarP(&CollectionId, "collectionId", "", "", "collection id")

	addrErr := addJobToCollectionCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
	jobIdErr := addJobToCollectionCmd.MarkFlagRequired("jobId")
	utils.CheckError("Job Id error: ", jobIdErr)
	collectionIdErr := addJobToCollectionCmd.MarkFlagRequired("collectionId")
	utils.CheckError("Collection Id error: ", collectionIdErr)
}
