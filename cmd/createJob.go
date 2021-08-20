package cmd

import (
	"razor/core"
	"razor/core/types"
	"razor/utils"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// createJobCmd represents the createJob command
var createJobCmd = &cobra.Command{
	Use:   "createJob",
	Short: "Create Job is used to create a job on razor.network",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		if err != nil {
			log.Fatal("Error in getting config: ", err)
		}

		var password string
		if utils.IsFlagPassed("password") {
			passwordPath, _ := cmd.Flags().GetString("password")
			password = utils.GetPasswordFromFile(passwordPath)
		} else {
			password = utils.PasswordPrompt()
		}

		address, _ := cmd.Flags().GetString("address")
		name, _ := cmd.Flags().GetString("name")
		repeat, _ := cmd.Flags().GetBool("repeat")
		url, _ := cmd.Flags().GetString("url")
		selector, _ := cmd.Flags().GetString("selector")

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
		txn, err := assetManager.CreateJob(txnOpts, url, selector, name, repeat)
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
		Repeat   bool
		Account  string
		Password string
	)

	createJobCmd.Flags().StringVarP(&URL, "url", "u", "", "url of job")
	createJobCmd.Flags().StringVarP(&Selector, "selector", "s", "", "selector (jsonPath selector)")
	createJobCmd.Flags().StringVarP(&Name, "name", "n", "", "name of job")
	createJobCmd.Flags().BoolVarP(&Repeat, "repeat", "r", true, "repeat")
	createJobCmd.Flags().StringVarP(&Account, "address", "", "", "address of the job creator")
	createJobCmd.Flags().StringVarP(&Password, "password", "", "", "password of the job creator")

	urlErr := createJobCmd.MarkFlagRequired("url")
	utils.CheckError("URL error: ", urlErr)
	selectorErr := createJobCmd.MarkFlagRequired("selector")
	utils.CheckError("Selector error: ", selectorErr)
	nameErr := createJobCmd.MarkFlagRequired("name")
	utils.CheckError("Name error: ", nameErr)
	addrErr := createJobCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
}
