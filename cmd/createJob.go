package cmd

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/utils"

	"github.com/spf13/cobra"
)

// createJobCmd represents the createJob command
var createJobCmd = &cobra.Command{
	Use:   "createJob",
	Short: "Create Job is used to create a job on razor.network",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createJob called")
		config, err := GetConfigData()
		if err != nil {
			log.Fatal("Error in getting config: ", err)
		}

		password, _ := cmd.Flags().GetString("password")
		address, _ := cmd.Flags().GetString("address")
		fee, _ := cmd.Flags().GetString("fee")
		name, _ := cmd.Flags().GetString("name")
		repeat, _ := cmd.Flags().GetBool("repeat")
		url, _ := cmd.Flags().GetString("url")
		selector, _ := cmd.Flags().GetString("selector")

		client := utils.ConnectToClient(config.Provider)

		accountBalance, err := client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
		if err != nil {
			log.Errorf("Error in fetching balance of the account: %s", address)
			log.Fatal(err)
		}

		feeInBigInt, ok := new(big.Int).SetString(fee, 10)
		if !ok {
			log.Fatal("SetString: error")
		}

		//feeInWei := big.NewInt(1).Mul(feeInBigInt, big.NewInt(1e18))

		if accountBalance.Cmp(feeInBigInt) < 0 {
			log.Fatal("Please make sure you hold sufficient ether in your account")
		}

		txnOpts := utils.GetTxnOpts(types.TransactionOptions{
			Client:         client,
			Password:       password,
			EtherValue:     feeInBigInt,
			AccountAddress: address,
			ChainId:        core.ChainId,
			GasMultiplier:  config.GasMultiplier,
		})

		jobManager := utils.GetJobManager(client)
		log.Info("Creating Job...")
		txn, err := jobManager.CreateJob(txnOpts, url, selector, name, repeat)
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
		Fee      string
		Account  string
		Password string
	)

	createJobCmd.Flags().StringVarP(&URL, "url", "u", "", "url of job")
	createJobCmd.Flags().StringVarP(&Selector, "selector", "s", "", "selector (comma separated for nested values)")
	createJobCmd.Flags().StringVarP(&Name, "name", "n", "", "name of job")
	createJobCmd.Flags().BoolVarP(&Repeat, "repeat", "r", true, "repeat")
	createJobCmd.Flags().StringVarP(&Fee, "fee", "f", "0", "fee")
	createJobCmd.Flags().StringVarP(&Account, "address", "", "", "address of the job creator")
	createJobCmd.Flags().StringVarP(&Password, "password", "", "", "password of the ̰ob creator")

	createJobCmd.MarkFlagRequired("url")
	createJobCmd.MarkFlagRequired("selector")
	createJobCmd.MarkFlagRequired("name")
	createJobCmd.MarkFlagRequired("fee")
	createJobCmd.MarkFlagRequired("address")
	createJobCmd.MarkFlagRequired("password")

}
