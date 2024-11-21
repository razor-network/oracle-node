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

// unlockWithdrawCmd represents the unlockWithdraw command
var unlockWithdrawCmd = &cobra.Command{
	Use:   "unlockWithdraw",
	Short: "UnlockWithdraw withdraws your razors once withdraw lock has passed",
	Long:  `unlockWithdraw has to be called once the withdraw lock period is over to get back all the razor tokens into your account`,
	Run:   initializeUnlockWithdraw,
}

//This function initialises the ExecuteUnlockWithdraw function
func initializeUnlockWithdraw(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteUnlockWithdraw(cmd.Flags())
}

//This function sets the flag appropriately and executes the UnlockWithdraw function
func (*UtilsStruct) ExecuteUnlockWithdraw(flagSet *pflag.FlagSet) {
	config, rpcParameters, account, err := InitializeCommandDependencies(flagSet)
	utils.CheckError("Error in initialising command dependencies: ", err)

	stakerId, err := razorUtils.AssignStakerId(rpcParameters, flagSet, account.Address)
	utils.CheckError("Error in fetching stakerId:  ", err)
	log.Debug("ExecuteUnlockWithdraw: StakerId: ", stakerId)

	txn, err := cmdUtils.HandleWithdrawLock(rpcParameters, account, config, stakerId)
	utils.CheckError("HandleWithdrawLock error: ", err)

	if txn != core.NilHash {
		err = razorUtils.WaitForBlockCompletion(rpcParameters, txn.Hex())
		utils.CheckError("Error in WaitForBlockCompletion for unlockWithdraw: ", err)
	}
}

//This function handles the Withdraw lock
func (*UtilsStruct) HandleWithdrawLock(rpcParameters rpc.RPCParameters, account types.Account, configurations types.Configurations, stakerId uint32) (common.Hash, error) {
	withdrawLock, err := razorUtils.GetLock(rpcParameters, account.Address, stakerId, 1)
	if err != nil {
		return core.NilHash, err
	}
	log.Debugf("HandleWithdrawLock: Withdraw lock: %+v", withdrawLock)

	if withdrawLock.UnlockAfter.Cmp(big.NewInt(0)) == 0 {
		log.Error("initiateWithdrawCmd command not called before unlocking razors!")
		return core.NilHash, errors.New("initiate withdrawal of Razors before unlocking withdraw")
	}

	epoch, err := razorUtils.GetEpoch(rpcParameters)
	if err != nil {
		log.Error("Error in fetching epoch")
		return core.NilHash, err
	}
	log.Debug("HandleWithdrawLock: Epoch: ", epoch)

	waitFor := big.NewInt(0).Sub(withdrawLock.UnlockAfter, big.NewInt(int64(epoch)))
	if waitFor.Cmp(big.NewInt(0)) > 0 {
		timeRemaining := uint64(waitFor.Int64()) * core.EpochLength
		log.Infof("Withdrawal period not reached. Cannot withdraw now, please wait for %d epoch(s)! (approximately %s)", waitFor, razorUtils.SecondsToReadableTime(int(timeRemaining)))
		return core.NilHash, nil
	}

	if big.NewInt(int64(epoch)).Cmp(withdrawLock.UnlockAfter) >= 0 {
		txnArgs := types.TransactionOptions{
			ChainId:         core.ChainId,
			Config:          configurations,
			ContractAddress: core.StakeManagerAddress,
			MethodName:      "unlockWithdraw",
			ABI:             bindings.StakeManagerMetaData.ABI,
			Parameters:      []interface{}{stakerId},
			Account:         account,
		}
		txnOpts, err := razorUtils.GetTxnOpts(rpcParameters, txnArgs)
		if err != nil {
			return core.NilHash, err
		}
		log.Debug("HandleWithdrawLock: Calling UnlockWithdraw() with arguments stakerId = ", stakerId)
		return cmdUtils.UnlockWithdraw(rpcParameters, txnOpts, stakerId)
	}
	return core.NilHash, errors.New("withdrawLock period not over yet! Please try after some time")
}

//This function withdraws your razor once withdraw lock has passed
func (*UtilsStruct) UnlockWithdraw(rpcParameters rpc.RPCParameters, txnOpts *bind.TransactOpts, stakerId uint32) (common.Hash, error) {
	log.Info("Unlocking funds...")
	client, err := rpcParameters.RPCManager.GetBestRPCClient()
	if err != nil {
		return core.NilHash, err
	}

	log.Debug("Executing UnlockWithdraw transaction with stakerId = ", stakerId)
	txn, err := stakeManagerUtils.UnlockWithdraw(client, txnOpts, stakerId)
	if err != nil {
		log.Error("Error in unlocking funds")
		return core.NilHash, err
	}

	txnHash := transactionUtils.Hash(txn)
	log.Info("Txn Hash: ", txnHash.Hex())
	return txnHash, nil
}

func init() {
	rootCmd.AddCommand(unlockWithdrawCmd)
	var (
		Address  string
		Password string
		StakerId uint32
	)

	unlockWithdrawCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the user")
	unlockWithdrawCmd.Flags().StringVarP(&Password, "password", "", "", "password path of user to protect the keystore")
	unlockWithdrawCmd.Flags().Uint32VarP(&StakerId, "stakerId", "", 0, "password path of user to protect the keystore")

	addrErr := unlockWithdrawCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
}
