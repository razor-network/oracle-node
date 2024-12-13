//Package cmd provides all functions related to command line
package cmd

import (
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/rpc"
	"razor/utils"

	"github.com/ethereum/go-ethereum/common"
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
	config, rpcParameters, account, err := InitializeCommandDependencies(flagSet)
	utils.CheckError("Error in initialising command dependencies: ", err)

	stakerId, err := razorUtils.AssignStakerId(rpcParameters, flagSet, account.Address)
	utils.CheckError("Error in getting stakerId: ", err)

	extendLockInput := types.ExtendLockInput{
		StakerId: stakerId,
		Account:  account,
	}

	txn, err := cmdUtils.ResetUnstakeLock(rpcParameters, config, extendLockInput)
	utils.CheckError("Error in extending lock: ", err)
	err = razorUtils.WaitForBlockCompletion(rpcParameters, txn.Hex())
	utils.CheckError("Error in WaitForBlockCompletion for resetUnstakeLock: ", err)
}

//This function is used to reset the lock once the withdraw lock period is over
func (*UtilsStruct) ResetUnstakeLock(rpcParameters rpc.RPCParameters, config types.Configurations, extendLockInput types.ExtendLockInput) (common.Hash, error) {
	txnOpts, err := razorUtils.GetTxnOpts(rpcParameters, types.TransactionOptions{
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.StakeManagerAddress,
		MethodName:      "resetUnstakeLock",
		Parameters:      []interface{}{extendLockInput.StakerId},
		ABI:             bindings.StakeManagerMetaData.ABI,
		Account:         extendLockInput.Account,
	})
	if err != nil {
		return core.NilHash, err
	}

	log.Info("Extending lock...")
	client, err := rpcParameters.RPCManager.GetBestRPCClient()
	if err != nil {
		return core.NilHash, err
	}

	log.Debug("Executing ResetUnstakeLock transaction with stakerId = ", extendLockInput.StakerId)
	txn, err := stakeManagerUtils.ResetUnstakeLock(client, txnOpts, extendLockInput.StakerId)
	if err != nil {
		return core.NilHash, err
	}
	txnHash := transactionUtils.Hash(txn)
	log.Info("Txn Hash: ", txnHash.Hex())
	return txnHash, nil
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
