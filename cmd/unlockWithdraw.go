//Package cmd provides all functions related to command line
package cmd

import (
	"errors"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/logger"
	"razor/pkg/bindings"
	"razor/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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
	razorUtils.AssignLogFile(flagSet)
	address, err := flagSetUtils.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)

	logger.Address = address

	config, err := cmdUtils.GetConfigData()
	utils.CheckError("Error in getting config: ", err)

	password := razorUtils.AssignPassword()

	client := razorUtils.ConnectToClient(config.Provider)

	razorUtils.CheckEthBalanceIsZero(client, address)

	stakerId, err := razorUtils.AssignStakerId(flagSet, client, address)
	utils.CheckError("Error in fetching stakerId:  ", err)

	txn, err := cmdUtils.HandleWithdrawLock(client, types.Account{
		Address:  address,
		Password: password,
	}, config, stakerId)

	utils.CheckError("UnlockWithdraw error: ", err)
	if txn != core.NilHash {
		err = razorUtils.WaitForBlockCompletion(client, txn.Hex())
		utils.CheckError("Error in WaitForBlockCompletion for unlockWithdraw: ", err)
	}
}

//This function handles the Withdraw lock
func (*UtilsStruct) HandleWithdrawLock(client *ethclient.Client, account types.Account, configurations types.Configurations, stakerId uint32) (common.Hash, error) {
	withdrawLock, err := razorUtils.GetLock(client, account.Address, stakerId, 1)
	if err != nil {
		return core.NilHash, err
	}

	if withdrawLock.UnlockAfter.Cmp(big.NewInt(0)) == 0 {
		log.Error("initiateWithdrawCmd command not called before unlocking razors!")
		return core.NilHash, errors.New("initiate withdrawal of Razors before unlocking withdraw")
	}

	epoch, err := razorUtils.GetEpoch(client)
	if err != nil {
		log.Error("Error in fetching epoch")
		return core.NilHash, err
	}

	if big.NewInt(int64(epoch)).Cmp(withdrawLock.UnlockAfter) >= 0 {
		txnArgs := types.TransactionOptions{
			Client:          client,
			Password:        account.Password,
			AccountAddress:  account.Address,
			ChainId:         core.ChainId,
			Config:          configurations,
			ContractAddress: core.StakeManagerAddress,
			MethodName:      "unlockWithdraw",
			ABI:             bindings.StakeManagerMetaData.ABI,
			Parameters:      []interface{}{stakerId},
		}
		txnOpts := razorUtils.GetTxnOpts(txnArgs)
		return cmdUtils.UnlockWithdraw(client, txnOpts, stakerId)
	}
	return core.NilHash, errors.New("withdrawLock period not over yet! Please try after some time")
}

//This function withdraws your razor once withdraw lock has passed
func (*UtilsStruct) UnlockWithdraw(client *ethclient.Client, txnOpts *bind.TransactOpts, stakerId uint32) (common.Hash, error) {
	log.Info("Unlocking funds...")

	txn, err := stakeManagerUtils.UnlockWithdraw(client, txnOpts, stakerId)
	txnHash := transactionUtils.Hash(txn)
	if err != nil {
		log.Error("Error in unlocking funds")
		return core.NilHash, err
	}

	log.Info("Txn Hash: ", txnHash.Hex())

	return txnHash, nil
}

func init() {
	rootCmd.AddCommand(unlockWithdrawCmd)
	var (
		Address  string
		StakerId uint32
	)

	unlockWithdrawCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the user")
	unlockWithdrawCmd.Flags().Uint32VarP(&StakerId, "stakerId", "", 0, "password path of user to protect the keystore")

	addrErr := unlockWithdrawCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
}
