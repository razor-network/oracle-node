package cmd

import (
	"razor/core"
	"razor/core/types"
	"razor/utils"

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
		stakerId, _ := cmd.Flags().GetUint32("stakerId")

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
		txn, err := stakeManager.ResetLock(txnOpts, stakerId)
		utils.CheckError("Error in resetting lock: ", err)
		log.Infof("Transaction Hash: %s", txn.Hash())
		utils.WaitForBlockCompletion(client, txn.Hash().String())
	},
}

func init() {
	rootCmd.AddCommand(resetLockCmd)

	var (
		Address  string
		Password string
		StakerId uint32
	)

	resetLockCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the user")
	resetLockCmd.Flags().StringVarP(&Password, "password", "", "", "password path of the user to protect the keystore")
	resetLockCmd.Flags().Uint32VarP(&StakerId, "stakerId", "", 0, "staker id")

	addrErr := resetLockCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
}
