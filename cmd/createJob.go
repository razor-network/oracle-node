package cmd

import (
	"razor/core"
	"razor/core/types"
	"razor/utils"

	"github.com/spf13/cobra"
)

var createJobCmd = &cobra.Command{
	Use:   "createJob",
	Short: "createJob can be used to create a job",
	Long: `A job consists of a URL and a selector to fetch the exact data from the URL. The createJob command can be used to create a job that the stakers can vote upon.

Example:
  ./razor createJob -a 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c -n btcusd_gemini -r true -s last -u https://api.gemini.com/v1/pubticker/btcusd

Note: 
  This command only works for the admin.
`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		if err != nil {
			log.Fatal("Error in getting config: ", err)
		}

		password := utils.AssignPassword(cmd.Flags())
		address, _ := cmd.Flags().GetString("address")
		name, _ := cmd.Flags().GetString("name")
		url, _ := cmd.Flags().GetString("url")
		selector, _ := cmd.Flags().GetString("selector")
		power, _ := cmd.Flags().GetInt8("power")

		client := utils.ConnectToClient(config.Provider)
		txnOpts := utils.GetTxnOpts(types.TransactionOptions{
			Client:         client,
			Password:       password,
			AccountAddress: address,
			ChainId:        core.ChainId,
			Config:         config,
		})

		assetManager := utils.GetAssetManager(client)
		log.Info("Creating Job...")
		txn, err := assetManager.CreateJob(txnOpts, power, name, selector, url)
		if err != nil {
			log.Fatal(err)
		}
		log.Info("Job creation transaction sent.")
		log.Info("Transaction Hash: ", txn.Hash())
		utils.WaitForBlockCompletion(client, txn.Hash().String())
	},
}

func init() {
	rootCmd.AddCommand(createJobCmd)

	var (
		URL      string
		Selector string
		Name     string
		Power    int8
		Account  string
		Password string
	)

	createJobCmd.Flags().StringVarP(&URL, "url", "u", "", "url of job")
	createJobCmd.Flags().StringVarP(&Selector, "selector", "s", "", "selector (jsonPath selector)")
	createJobCmd.Flags().StringVarP(&Name, "name", "n", "", "name of job")
	createJobCmd.Flags().Int8VarP(&Power, "power", "", 0, "power")
	createJobCmd.Flags().StringVarP(&Account, "address", "a", "", "address of the job creator")
	createJobCmd.Flags().StringVarP(&Password, "password", "", "", "password path of job creator to protect the keystore")

	urlErr := createJobCmd.MarkFlagRequired("url")
	utils.CheckError("URL error: ", urlErr)
	selectorErr := createJobCmd.MarkFlagRequired("selector")
	utils.CheckError("Selector error: ", selectorErr)
	nameErr := createJobCmd.MarkFlagRequired("name")
	utils.CheckError("Name error: ", nameErr)
	addrErr := createJobCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
	powErr := createJobCmd.MarkFlagRequired("power")
	utils.CheckError("Power error: ", powErr)
}
