package cmd

import (
	log "github.com/sirupsen/logrus"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/utils"

	"github.com/spf13/cobra"
)

// addJobToCollectionCmd represents the addJobToCollection command
var addJobToCollectionCmd = &cobra.Command{
	Use:   "addJobToCollection",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

	addJobToCollectionCmd.Flags().StringVarP(&Account, "address", "", "", "address of the job creator")
	addJobToCollectionCmd.Flags().StringVarP(&JobId, "jobId", "", "", "job id to add to the  collection")
	addJobToCollectionCmd.Flags().StringVarP(&CollectionId, "collectionId", "", "", "collection id")

	addrErr := addJobToCollectionCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
	jobIdErr := addJobToCollectionCmd.MarkFlagRequired("jobId")
	utils.CheckError("Job Id error: ", jobIdErr)
	collectionIdErr := addJobToCollectionCmd.MarkFlagRequired("collectionId")
	utils.CheckError("Collection Id error: ", collectionIdErr)
}
