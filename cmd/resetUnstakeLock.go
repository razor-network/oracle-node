//Package cmd provides all functions related to command line
package cmd

import (
	"razor/core"
	"razor/core/types"
	"razor/logger"
	"razor/pkg/bindings"
	"razor/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"

	"github.com/spf13/cobra"
)

var extendUnstakeLockCmd = &cobra.Command{
	Use:   "extendLock",
	Short: "extendLock can be used to reset the lock once the withdraw lock period is over",
	Long: `If the withdrawal period is over, then the lock must be reset otherwise the user cannot unstake. This can be done by extendLock command.

Example:
  ./razor extendLock --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c 
`,
	Run: initialiseExtendLock,
}

//This function initialises the ExtendLock function
func initialiseExtendLock(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteExtendLock(cmd.Flags())
}

//This function sets the flags appropriately and executes the ResetUnstakeLock function
func (*UtilsStruct) ExecuteExtendLock(flagSet *pflag.FlagSet) {
	razorUtils.AssignLogFile(flagSet)
	address, err := flagSetUtils.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)

	logger.Address = address

	config, err := cmdUtils.GetConfigData()
	utils.CheckError("Error in getting config data: ", err)

	password := razorUtils.AssignPassword(flagSet)

	client := razorUtils.ConnectToClient(config.Provider)

	stakerId, err := razorUtils.AssignStakerId(flagSet, client, address)
	utils.CheckError("Error in getting stakerId: ", err)

	extendLockInput := types.ExtendLockInput{
		Address:  address,
		Password: password,
		StakerId: stakerId,
	}
	txn, err := cmdUtils.ResetUnstakeLock(client, config, extendLockInput)
	utils.CheckError("Error in extending lock: ", err)
	razorUtils.WaitForBlockCompletion(client, txn.String())
}

//This function is used to reset the lock once the withdraw lock period is over
func (*UtilsStruct) ResetUnstakeLock(client *ethclient.Client, config types.Configurations, extendLockInput types.ExtendLockInput) (common.Hash, error) {
	txnOpts := razorUtils.GetTxnOpts(types.TransactionOptions{
		Client:          client,
		Password:        extendLockInput.Password,
		AccountAddress:  extendLockInput.Address,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.StakeManagerAddress,
		MethodName:      "resetUnstakeLock",
		Parameters:      []interface{}{extendLockInput.StakerId},
		ABI:             bindings.StakeManagerABI,
	})

	log.Info("Extending lock...")
	txn, err := stakeManagerUtils.ResetUnstakeLock(client, txnOpts, extendLockInput.StakerId)
	if err != nil {
		return core.NilHash, err
	}
	log.Info("Txn Hash: ", transactionUtils.Hash(txn))
	return transactionUtils.Hash(txn), nil
}

func init() {
	rootCmd.AddCommand(extendUnstakeLockCmd)

	var (
		Address  string
		Password string
		StakerId uint32
	)

	extendUnstakeLockCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the user")
	extendUnstakeLockCmd.Flags().StringVarP(&Password, "password", "", "", "password path of the user to protect the keystore")
	extendUnstakeLockCmd.Flags().Uint32VarP(&StakerId, "stakerId", "", 0, "staker id")

	addrErr := extendUnstakeLockCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
}
