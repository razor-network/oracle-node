package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/utils"
)

// createCollectionCmd represents the createCollection command
var createCollectionCmd = &cobra.Command{
	Use:   "createCollection",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		if err != nil {
			log.Fatal("Error in getting config: ", err)
		}

		password := utils.PasswordPrompt()

		name, _ := cmd.Flags().GetString("name")
		fee, _ := cmd.Flags().GetString("fee")
		address, _ := cmd.Flags().GetString("address")
		jobIds, _ := cmd.Flags().GetStringSlice("jobIds")
		aggregation, _ := cmd.Flags().GetUint32("aggregation")

		client := utils.ConnectToClient(config.Provider)

		feeInBigInt, ok := new(big.Int).SetString(fee, 10)
		if !ok {
			log.Fatal("SetString: error")
		}

		jobIdsInBigInt := utils.ConvertToBigIntArray(jobIds)

		txnOpts := utils.GetTxnOpts(types.TransactionOptions{
			Client:         client,
			Password:       password,
			EtherValue:     feeInBigInt,
			AccountAddress: address,
			ChainId:        core.ChainId,
			GasMultiplier:  config.GasMultiplier,
		})

		assetManager := utils.GetAssetManager(client)
		txn, err := assetManager.CreateCollection(txnOpts, name, jobIdsInBigInt, aggregation)
		if err != nil {
			log.Fatal(err)
		}
		log.Info("Creating collection...")
		utils.WaitForBlockCompletion(client, txn.Hash().String())
	},
}

func init() {
	rootCmd.AddCommand(createCollectionCmd)

	var (
		Name              string
		Fee               string
		Account           string
		JobIds            []string
		AggregationMethod uint32
	)

	createCollectionCmd.Flags().StringVarP(&Name, "name", "n", "", "name of the collection")
	createCollectionCmd.Flags().StringVarP(&Fee, "fee", "f", "0", "fee")
	createCollectionCmd.Flags().StringVarP(&Account, "address", "", "", "address of the job creator")
	createCollectionCmd.Flags().StringSliceVarP(&JobIds, "jobIds", "", []string{}, "job ids for the  collection")
	createCollectionCmd.Flags().Uint32VarP(&AggregationMethod, "aggregation", "", 1, "aggregation method to be used")

	nameErr := createCollectionCmd.MarkFlagRequired("name")
	utils.CheckError("Name error: ", nameErr)
	feeErr := createCollectionCmd.MarkFlagRequired("fee")
	utils.CheckError("Fee error: ", feeErr)
	addrErr := createCollectionCmd.MarkFlagRequired("address")
	utils.CheckError("Address Error: ", addrErr)
	jobIdErr := createCollectionCmd.MarkFlagRequired("jobIds")
	utils.CheckError("Job Id Error: ", jobIdErr)
}
