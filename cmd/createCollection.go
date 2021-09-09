package cmd

import (
	"github.com/spf13/cobra"
	"razor/core"
	"razor/core/types"
	"razor/utils"
)

var createCollectionCmd = &cobra.Command{
	Use:   "createCollection",
	Short: "createCollection can be used to create collections if existing jobs are present",
	Long: `A collection is a group of jobs that reports the aggregated value of jobs. createCollection can be used to club multiple jobs into one collection bound by an aggregation method.

Example: 
  ./razor createCollection --name btcCollectionMean --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --jobIds 1,2 --aggregation 2

Note: 
  This command only works for the admin.
`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		if err != nil {
			log.Fatal("Error in getting config: ", err)
		}

		password := utils.AssignPassword(cmd.Flags())
		name, _ := cmd.Flags().GetString("name")
		address, _ := cmd.Flags().GetString("address")
		jobIdInUint, _ := cmd.Flags().GetUintSlice("jobIds")
		aggregation, _ := cmd.Flags().GetUint32("aggregation")
		power, _ := cmd.Flags().GetInt8("power")

		client := utils.ConnectToClient(config.Provider)

		txnOpts := utils.GetTxnOpts(types.TransactionOptions{
			Client:         client,
			Password:       password,
			AccountAddress: address,
			ChainId:        core.ChainId,
			Config:         config,
		})
		jobIds := utils.ConvertUintArrayToUint8Array(jobIdInUint)
		assetManager := utils.GetAssetManager(client)
		_, err = WaitForDisputeOrConfirmState(client, address, "create collection")
		utils.CheckError("Error in fetching state: ", err)
		txn, err := assetManager.CreateCollection(txnOpts, jobIds, aggregation, power, name)
		utils.CheckError("Error in creating collection: ", err)
		log.Info("Creating collection...")
		utils.WaitForBlockCompletion(client, txn.Hash().String())
	},
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
	)

	createCollectionCmd.Flags().StringVarP(&Name, "name", "n", "", "name of the collection")
	createCollectionCmd.Flags().StringVarP(&Account, "address", "a", "", "address of the job creator")
	createCollectionCmd.Flags().UintSliceVarP(&JobIds, "jobIds", "", []uint{}, "job ids for the  collection")
	createCollectionCmd.Flags().Uint32VarP(&AggregationMethod, "aggregation", "", 1, "aggregation method to be used")
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
}
