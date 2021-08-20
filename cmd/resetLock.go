package cmd

import (
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/utils"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// resetLockCmd represents the resetLock command
var resetLockCmd = &cobra.Command{
	Use:   "resetLock",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		utils.CheckError("Error in getting config data: ", err)

		var password string
		if utils.IsFlagPassed("password") {
			passwordPath, _ := cmd.Flags().GetString("password")
			password = utils.GetPasswordFromFile(passwordPath)
		} else {
			password = utils.PasswordPrompt()
		}

		address, _ := cmd.Flags().GetString("address")
		stakerId, _ := cmd.Flags().GetString("stakerId")

		_stakerId, ok := new(big.Int).SetString(stakerId, 10)
		if !ok {
			log.Fatal("Set string error in converting staker id")
		}

		client := utils.ConnectToClient(config.Provider)

		stakeManager := utils.GetStakeManager(client)

		txnOpts := utils.GetTxnOpts(types.TransactionOptions{
			Client:         client,
			Password:       password,
			AccountAddress: address,
			ChainId:        core.ChainId,
			Config:         config,
		})

		log.Info("Resetting lock...")
		txn, err := stakeManager.ResetLock(txnOpts, _stakerId)
		utils.CheckError("Error in resetting lock: ", err)
		log.Info("Transaction sent..")
		log.Infof("Transaction Hash: %s", txn.Hash())
		utils.WaitForBlockCompletion(client, txn.Hash().String())
	},
}

func init() {
	rootCmd.AddCommand(resetLockCmd)

	var (
		Address  string
		StakerId string
		Password string
	)

	resetLockCmd.Flags().StringVarP(&Address, "address", "", "", "address of the user")
	resetLockCmd.Flags().StringVarP(&StakerId, "stakerId", "", "", "staker's id to reset lock")
	resetLockCmd.Flags().StringVarP(&Password, "password", "", "", "password of the user")

	addrErr := resetLockCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
	stakerIdErr := resetLockCmd.MarkFlagRequired("stakerId")
	utils.CheckError("Address error: ", stakerIdErr)
}
