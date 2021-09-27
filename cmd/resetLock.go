package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/pflag"
	"razor/core"
	"razor/core/types"
	"razor/utils"

	"github.com/spf13/cobra"
)

var flagSetUtils flagSetInterface

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
		txn, err := resetLock(cmd.Flags(), config, razorUtils, stakeManagerUtils, transactionUtils, flagSetUtils)
		utils.CheckError("Error in resetting lock: ", err)
		utils.WaitForBlockCompletion(utils.ConnectToClient(config.Provider), txn.String())
	},
}

func resetLock(flagSet *pflag.FlagSet, config types.Configurations, razorUtils utilsInterface, stakeManagerUtils stakeManagerInterface, transactionUtils transactionInterface, flagSetUtils flagSetInterface) (common.Hash, error) {
	password := razorUtils.AssignPassword(flagSet)
	address, err := flagSetUtils.GetStringAddress(flagSet)
	if err != nil {
		return core.NilHash, err
	}
	stakerId, err := flagSetUtils.GetUint32StakerId(flagSet)
	if err != nil {
		return core.NilHash, err
	}
	client := razorUtils.ConnectToClient(config.Provider)

	txnOpts := razorUtils.GetTxnOpts(types.TransactionOptions{
		Client:         client,
		Password:       password,
		AccountAddress: address,
		ChainId:        core.ChainId,
		Config:         config,
	})

	log.Info("Resetting lock...")
	txn, err := stakeManagerUtils.ResetLock(client, txnOpts, stakerId)
	if err != nil {
		return core.NilHash, err
	}
	log.Info("Txn Hash: ", transactionUtils.Hash(txn))
	return transactionUtils.Hash(txn), nil
}

func init() {
	razorUtils = Utils{}
	stakeManagerUtils = StakeManagerUtils{}
	transactionUtils = TransactionUtils{}
	flagSetUtils = FlagSetUtils{}
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
