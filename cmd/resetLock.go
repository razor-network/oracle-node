package cmd

import (
	"razor/core"
	"razor/core/types"
	"razor/utils"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var resetLockCmd = &cobra.Command{
	Use:   "resetLock",
	Short: "resetLock can be used to reset the lock once the withdraw lock period is over",
	Long: `If the withdrawal period is over, then the lock must be reset otherwise the user cannot unstake. This can be done by resetLock command.

Example:
  ./razor resetLock --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c 
`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		utils.CheckError("Error in getting config data: ", err)

		password := utils.AssignPassword(cmd.Flags())
		address, _ := cmd.Flags().GetString("address")

		client := utils.ConnectToClient(config.Provider)
		stakerId, err := utils.GetStakerId(client, address)
		utils.CheckError("Error in fetching staker id: ", err)
		stakeManager := utils.GetStakeManager(client)

		txnOpts := utils.GetTxnOpts(types.TransactionOptions{
			Client:         client,
			Password:       password,
			AccountAddress: address,
			ChainId:        core.ChainId,
			Config:         config,
		})

		log.Info("Resetting lock...")
		txn, err := stakeManager.ResetLock(txnOpts, stakerId)
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
		Password string
	)

	resetLockCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the user")
	resetLockCmd.Flags().StringVarP(&Password, "password", "", "", "password path of the user to protect the keystore")

	addrErr := resetLockCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
}
