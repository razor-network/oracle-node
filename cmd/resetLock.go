package cmd

import (
	"math/big"
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
  ./razor resetLock --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --stakerId 1
`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		utils.CheckError("Error in getting config data: ", err)

		password := utils.PasswordPrompt()
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
	)

	resetLockCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the user")
	resetLockCmd.Flags().StringVarP(&StakerId, "stakerId", "", "", "staker's id to reset lock")

	addrErr := resetLockCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
	stakerIdErr := resetLockCmd.MarkFlagRequired("stakerId")
	utils.CheckError("Address error: ", stakerIdErr)
}
