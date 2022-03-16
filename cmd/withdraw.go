package cmd

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/logger"
	"razor/pkg/bindings"
	"razor/utils"
	"strconv"
	"time"
)

var withdrawCmd = &cobra.Command{
	Use:   "withdraw",
	Short: "withdraw your razors once you've unstaked",
	Long: `withdraw command can be used once the user has unstaked their token and the withdraw period is upon them.

Example:
  ./razor withdraw --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --stakerId 1
`,
	Run: initialiseWithdraw,
}

func initialiseWithdraw(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteWithdraw(cmd.Flags())
}

func (*UtilsStruct) ExecuteWithdraw(flagSet *pflag.FlagSet) {
	address, err := flagSetUtils.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)

	logger.Address = address

	config, err := cmdUtils.GetConfigData()
	utils.CheckError("Error in getting config: ", err)

	password := razorUtils.AssignPassword(flagSet)

	client := razorUtils.ConnectToClient(config.Provider)

	razorUtils.CheckEthBalanceIsZero(client, address)

	stakerId, err := razorUtils.AssignStakerId(flagSet, client, address)
	utils.CheckError("Error in fetching stakerId:  ", err)

	txn, err := cmdUtils.WithdrawFunds(client, types.Account{
		Address:  address,
		Password: password,
	}, config, stakerId)

	utils.CheckError("Withdraw error: ", err)
	if txn != core.NilHash {
		razorUtils.WaitForBlockCompletion(client, txn.String())
	}
}

func (*UtilsStruct) WithdrawFunds(client *ethclient.Client, account types.Account, configurations types.Configurations, stakerId uint32) (common.Hash, error) {

	lock, err := razorUtils.GetLock(client, account.Address, stakerId)
	if err != nil {
		log.Error("Error in fetching lock")
		return core.NilHash, err
	}

	if lock.WithdrawAfter.Cmp(big.NewInt(0)) == 0 {
		log.Info("Please unstake Razors before withdrawing.")
		return core.NilHash, nil
	}

	withdrawReleasePeriod, err := razorUtils.GetWithdrawReleasePeriod(client)
	if err != nil {
		log.Error("Error in fetching withdraw release period")
		return core.NilHash, err
	}
	withdrawBefore := big.NewInt(0).Add(lock.WithdrawAfter, big.NewInt(int64(withdrawReleasePeriod)))
	txnArgs := types.TransactionOptions{
		Client:          client,
		Password:        account.Password,
		AccountAddress:  account.Address,
		ChainId:         core.ChainId,
		Config:          configurations,
		ContractAddress: core.StakeManagerAddress,
		MethodName:      "withdraw",
		ABI:             bindings.StakeManagerABI,
	}
	epoch, err := razorUtils.GetEpoch(client)
	if err != nil {
		log.Error("Error in fetching epoch")
		return core.NilHash, err
	}
	if big.NewInt(int64(epoch)).Cmp(withdrawBefore) > 0 {
		log.Info("Withdrawal period has passed. Cannot withdraw now, please reset the lock!")
		return core.NilHash, nil
	}

	waitFor := big.NewInt(0).Sub(lock.WithdrawAfter, big.NewInt(int64(epoch)))
	if waitFor.Cmp(big.NewInt(0)) > 0 {
		timeRemaining := (time.Duration(int64(uint32(waitFor.Int64()))*core.EpochLength*razorUtils.CalculateBlockTime(client)) * time.Second) / (1000000000)
		value, err := strconv.Atoi(big.NewInt(int64(timeRemaining)).String())
		if err != nil {
			return core.NilHash, err
		}
		if waitFor.Cmp(big.NewInt(1)) == 0 {
			log.Infof("Withdrawal period not reached. Cannot withdraw now, please wait for %d epoch! (approximately %s)", waitFor, razorUtils.SecondsToReadableTime(value))
		} else {
			log.Infof("Withdrawal period not reached. Cannot withdraw now, please wait for %d epochs! (approximately %s)", waitFor, razorUtils.SecondsToReadableTime(value))
		}
		return core.NilHash, nil
	}

	txnArgs.Parameters = []interface{}{stakerId}
	txnOpts := razorUtils.GetTxnOpts(txnArgs)

	if big.NewInt(int64(epoch)).Cmp(lock.WithdrawAfter) >= 0 && big.NewInt(int64(epoch)).Cmp(withdrawBefore) <= 0 {
		return cmdUtils.Withdraw(client, txnOpts, stakerId)
	}
	return core.NilHash, nil
}

func (*UtilsStruct) Withdraw(client *ethclient.Client, txnOpts *bind.TransactOpts, stakerId uint32) (common.Hash, error) {
	log.Info("Withdrawing funds...")

	txn, err := stakeManagerUtils.Withdraw(client, txnOpts, stakerId)
	if err != nil {
		log.Error("Error in withdrawing funds")
		return core.NilHash, err
	}

	log.Info("Txn Hash: ", transactionUtils.Hash(txn))

	return transactionUtils.Hash(txn), nil
}

func init() {

	rootCmd.AddCommand(withdrawCmd)

	var (
		Address  string
		Password string
		StakerId uint32
	)

	withdrawCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the user")
	withdrawCmd.Flags().StringVarP(&Password, "password", "", "", "password path of user to protect the keystore")
	withdrawCmd.Flags().Uint32VarP(&StakerId, "stakerId", "", 0, "password path of user to protect the keystore")

	addrErr := withdrawCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
}
