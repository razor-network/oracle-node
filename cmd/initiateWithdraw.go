//Package cmd provides all functions related to command line
package cmd

import (
	"errors"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/rpc"
	"razor/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var initiateWithdrawCmd = &cobra.Command{
	Use:   "initiateWithdraw",
	Short: "initiateWithdraw for your razors once you've unstaked",
	Long: `initiateWithdraw command can be used once the user has unstaked their token and the withdraw period is upon them.

Example:
  ./razor initiateWithdraw --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --stakerId 1
`,
	Run: func(cmd *cobra.Command, args []string) {
		cmdUtils.ExecuteInitiateWithdraw(cmd.Flags())
	},
}

//This function sets the flags appropriately and executes the InitiateWithdraw function
func (*UtilsStruct) ExecuteInitiateWithdraw(flagSet *pflag.FlagSet) {
	config, rpcParameters, account, err := InitializeCommandDependencies(flagSet)
	utils.CheckError("Error in initialising command dependencies: ", err)

	stakerId, err := razorUtils.AssignStakerId(rpcParameters, flagSet, account.Address)
	utils.CheckError("Error in fetching stakerId:  ", err)
	log.Debug("ExecuteInitiateWithdraw: Staker Id: ", stakerId)

	log.Debugf("ExecuteInitiateWithdraw: Calling HandleUnstakeLock() with arguments account address: %s, stakerId: %d", account.Address, stakerId)
	txn, err := cmdUtils.HandleUnstakeLock(rpcParameters, account, config, stakerId)
	utils.CheckError("InitiateWithdraw error: ", err)

	if txn != core.NilHash {
		err := razorUtils.WaitForBlockCompletion(rpcParameters, txn.Hex())
		utils.CheckError("Error in WaitForBlockCompletion for initiateWithdraw: ", err)
	}
}

//This function handles the unstake lock
func (*UtilsStruct) HandleUnstakeLock(rpcParameters rpc.RPCParameters, account types.Account, configurations types.Configurations, stakerId uint32) (common.Hash, error) {
	unstakeLock, err := razorUtils.GetLock(rpcParameters, account.Address, stakerId, 0)
	if err != nil {
		log.Error("Error in fetching unstakeLock")
		return core.NilHash, err
	}
	log.Debugf("HandleUnstakeLock: Unstake lock: %+v", unstakeLock)

	if unstakeLock.UnlockAfter.Cmp(big.NewInt(0)) == 0 {
		log.Error("Unstake command not called before initiating withdrawal!")
		return core.NilHash, errors.New("unstake Razors before withdrawing")
	}

	withdrawInitiationPeriod, err := razorUtils.GetWithdrawInitiationPeriod(rpcParameters)
	if err != nil {
		log.Error("Error in fetching withdraw release period")
		return core.NilHash, err
	}
	log.Debug("HandleUnstakeLock: Withdraw initiation period: ", withdrawInitiationPeriod)

	withdrawBefore := big.NewInt(0).Add(unstakeLock.UnlockAfter, big.NewInt(int64(withdrawInitiationPeriod)))
	log.Debug("HandleUnstakeLock: Withdraw before epoch: ", withdrawBefore)
	epoch, err := razorUtils.GetEpoch(rpcParameters)
	if err != nil {
		log.Error("Error in fetching epoch")
		return core.NilHash, err
	}
	log.Debug("HandleUnstakeLock: Current epoch: ", epoch)

	if big.NewInt(int64(epoch)).Cmp(withdrawBefore) > 0 {
		log.Info("Withdraw initiation period has passed. Cannot withdraw now, please reset the unstakeLock!")
		return core.NilHash, nil
	}

	waitFor := big.NewInt(0).Sub(unstakeLock.UnlockAfter, big.NewInt(int64(epoch)))
	if waitFor.Cmp(big.NewInt(0)) > 0 {
		timeRemaining := uint64(waitFor.Int64()) * core.EpochLength
		log.Infof("Withdrawal Initiation period not reached. Cannot initiate withdraw now, please wait for %d epoch(s)! (approximately %s)", waitFor, razorUtils.SecondsToReadableTime(int(timeRemaining)))
		return core.NilHash, nil
	}

	log.Debug("Waiting for appropriate state to initiate withdraw...")
	_, err = cmdUtils.WaitForAppropriateState(rpcParameters, "initiateWithdraw", 0, 1, 4)
	if err != nil {
		log.Error("Error in fetching state: ", err)
		return core.NilHash, err
	}

	txnArgs := types.TransactionOptions{
		ChainId:         core.ChainId,
		Config:          configurations,
		ContractAddress: core.StakeManagerAddress,
		MethodName:      "initiateWithdraw",
		ABI:             bindings.StakeManagerMetaData.ABI,
		Parameters:      []interface{}{stakerId},
		Account:         account,
	}
	txnOpts, err := razorUtils.GetTxnOpts(rpcParameters, txnArgs)
	if err != nil {
		return core.NilHash, err
	}

	if big.NewInt(int64(epoch)).Cmp(unstakeLock.UnlockAfter) >= 0 && big.NewInt(int64(epoch)).Cmp(withdrawBefore) <= 0 {
		log.Debug("Calling InitiateWithdraw() with arguments stakerId: ", stakerId)
		return cmdUtils.InitiateWithdraw(rpcParameters, txnOpts, stakerId)
	}
	return core.NilHash, errors.New("unstakeLock period not over yet! Please try after some time")
}

//This function initiate withdraw for your razors once you've unstaked
func (*UtilsStruct) InitiateWithdraw(rpcParameters rpc.RPCParameters, txnOpts *bind.TransactOpts, stakerId uint32) (common.Hash, error) {
	log.Info("Initiating withdrawal of funds...")

	client, err := rpcParameters.RPCManager.GetBestRPCClient()
	if err != nil {
		return core.NilHash, err
	}

	log.Debug("Executing InitiateWithdraw transaction for stakerId = ", stakerId)
	txn, err := stakeManagerUtils.InitiateWithdraw(client, txnOpts, stakerId)
	if err != nil {
		log.Error("Error in initiating withdrawal of funds")
		return core.NilHash, err
	}

	txnHash := transactionUtils.Hash(txn)
	log.Info("Txn Hash: ", txnHash.Hex())
	return txnHash, nil
}

func init() {
	rootCmd.AddCommand(initiateWithdrawCmd)

	var (
		Address  string
		Password string
		StakerId uint32
	)

	initiateWithdrawCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the user")
	initiateWithdrawCmd.Flags().StringVarP(&Password, "password", "", "", "password path of user to protect the keystore")
	initiateWithdrawCmd.Flags().Uint32VarP(&StakerId, "stakerId", "", 0, "password path of user to protect the keystore")

	addrErr := initiateWithdrawCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
}
