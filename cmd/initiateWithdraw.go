package cmd

import (
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/logger"
	"razor/pkg/bindings"
	"razor/utils"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var initiateWithdraw = &cobra.Command{
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

func (*UtilsStruct) ExecuteInitiateWithdraw(flagSet *pflag.FlagSet) {
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

	utils.CheckError("InitiateWithdraw error: ", err)
	if txn != core.NilHash {
		razorUtils.WaitForBlockCompletion(client, txn.String())
	}
}

func (*UtilsStruct) WithdrawFunds(client *ethclient.Client, account types.Account, configurations types.Configurations, stakerId uint32) (common.Hash, error) {

	withdrawLock, err := razorUtils.GetLock(client, account.Address, stakerId, 1)
	if err != nil {
		log.Error("Error in fetching withdrawLock")
		return core.NilHash, err
	}

	if withdrawLock.UnlockAfter.Cmp(big.NewInt(0)) == 0 {
		log.Info("Please unstake Razors before withdrawing.")
		return core.NilHash, nil
	}

	withdrawReleasePeriod, err := razorUtils.GetWithdrawReleasePeriod(client)
	if err != nil {
		log.Error("Error in fetching withdraw release period")
		return core.NilHash, err
	}
	withdrawBefore := big.NewInt(0).Add(withdrawLock.UnlockAfter, big.NewInt(int64(withdrawReleasePeriod)))
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
		log.Info("Withdrawal period has passed. Cannot withdraw now, please reset the withdrawLock!")
		return core.NilHash, nil
	}

	txnArgs.Parameters = []interface{}{stakerId}
	txnOpts := razorUtils.GetTxnOpts(txnArgs)

	for i := epoch; big.NewInt(int64(i)).Cmp(withdrawBefore) < 0; {
		if big.NewInt(int64(epoch)).Cmp(withdrawLock.UnlockAfter) >= 0 && big.NewInt(int64(epoch)).Cmp(withdrawBefore) <= 0 {
			return cmdUtils.InitiateWithdraw(client, txnOpts, stakerId)
		}
		log.Debug("Waiting for withdrawLock period to get over....")
		// Wait for 30 seconds if withdrawLock period isn't over
		timeUtils.Sleep(30 * time.Second)
		epoch, err = razorUtils.GetUpdatedEpoch(client)
		if err != nil {
			log.Error("Error in fetching epoch")
			return core.NilHash, err
		}
	}
	return core.NilHash, nil
}

func (*UtilsStruct) InitiateWithdraw(client *ethclient.Client, txnOpts *bind.TransactOpts, stakerId uint32) (common.Hash, error) {
	log.Info("Initiating withdrawal of funds...")

	txn, err := stakeManagerUtils.InitiateWithdraw(client, txnOpts, stakerId)
	if err != nil {
		log.Error("Error in initiating withdrawal of funds")
		return core.NilHash, err
	}

	log.Info("Txn Hash: ", transactionUtils.Hash(txn))

	return transactionUtils.Hash(txn), nil
}

func init() {

	razorUtils = Utils{}
	transactionUtils = TransactionUtils{}
	stakeManagerUtils = StakeManagerUtils{}
	InitializeUtils()
	cmdUtils = &UtilsStruct{}

	rootCmd.AddCommand(initiateWithdraw)

	var (
		Address  string
		Password string
		StakerId uint32
	)

	initiateWithdraw.Flags().StringVarP(&Address, "address", "a", "", "address of the user")
	initiateWithdraw.Flags().StringVarP(&Password, "password", "", "", "password path of user to protect the keystore")
	initiateWithdraw.Flags().Uint32VarP(&StakerId, "stakerId", "", 0, "password path of user to protect the keystore")

	addrErr := initiateWithdraw.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
}
