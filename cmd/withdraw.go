package cmd

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/pflag"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/spf13/cobra"
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
	cmdUtilsMockery.ExecuteWithdraw(cmd.Flags())
}

func (*UtilsStructMockery) ExecuteWithdraw(flagSet *pflag.FlagSet) {
	config, err := cmdUtilsMockery.GetConfigData()
	utils.CheckError("Error in getting config: ", err)

	password := razorUtilsMockery.AssignPassword(flagSet)
	address, err := flagSetUtilsMockery.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)

	client := razorUtilsMockery.ConnectToClient(config.Provider)

	razorUtilsMockery.CheckEthBalanceIsZero(client, address)

	stakerId, err := razorUtilsMockery.AssignStakerId(flagSet, client, address)
	utils.CheckError("Error in fetching stakerId:  ", err)

	txn, err := cmdUtilsMockery.WithdrawFunds(client, types.Account{
		Address:  address,
		Password: password,
	}, config, stakerId)

	utils.CheckError("Withdraw error: ", err)
	if txn != core.NilHash {
		razorUtilsMockery.WaitForBlockCompletion(client, txn.String())
	}
}

func (*UtilsStructMockery) WithdrawFunds(client *ethclient.Client, account types.Account, configurations types.Configurations, stakerId uint32) (common.Hash, error) {

	lock, err := razorUtilsMockery.GetLock(client, account.Address, stakerId)
	if err != nil {
		log.Error("Error in fetching lock")
		return core.NilHash, err
	}

	if lock.WithdrawAfter.Cmp(big.NewInt(0)) == 0 {
		log.Info("Please unstake Razors before withdrawing.")
		return core.NilHash, nil
	}

	withdrawReleasePeriod, err := razorUtilsMockery.GetWithdrawReleasePeriod(client, account.Address)
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
	epoch, err := razorUtilsMockery.GetEpoch(client)
	if err != nil {
		log.Error("Error in fetching epoch")
		return core.NilHash, err
	}
	if big.NewInt(int64(epoch)).Cmp(withdrawBefore) > 0 {
		log.Info("Withdrawal period has passed. Cannot withdraw now, please reset the lock!")
		return core.NilHash, nil
	}

	txnArgs.Parameters = []interface{}{epoch, stakerId}
	txnOpts := razorUtilsMockery.GetTxnOpts(txnArgs)

	for i := epoch; big.NewInt(int64(i)).Cmp(withdrawBefore) < 0; {
		if big.NewInt(int64(epoch)).Cmp(lock.WithdrawAfter) >= 0 && big.NewInt(int64(epoch)).Cmp(withdrawBefore) <= 0 {
			return cmdUtilsMockery.Withdraw(client, txnOpts, stakerId)
		}
		log.Debug("Waiting for lock period to get over....")
		// Wait for 30 seconds if lock period isn't over
		razorUtilsMockery.Sleep(30 * time.Second)
		epoch, err = razorUtilsMockery.GetUpdatedEpoch(client)
		if err != nil {
			log.Error("Error in fetching epoch")
			return core.NilHash, err
		}
	}
	return core.NilHash, nil
}

func (*UtilsStructMockery) Withdraw(client *ethclient.Client, txnOpts *bind.TransactOpts, stakerId uint32) (common.Hash, error) {
	log.Info("Withdrawing funds...")

	txn, err := stakeManagerUtilsMockery.Withdraw(client, txnOpts, stakerId)
	if err != nil {
		log.Error("Error in withdrawing funds")
		return core.NilHash, err
	}

	log.Info("Txn Hash: ", transactionUtilsMockery.Hash(txn))

	return transactionUtilsMockery.Hash(txn), nil
}

func init() {

	razorUtils = Utils{}
	transactionUtils = TransactionUtils{}
	stakeManagerUtils = StakeManagerUtils{}
	cmdUtils = UtilsCmd{}
	flagSetUtils = FlagSetUtils{}
	utils.Options = &utils.OptionsStruct{}
	utils.UtilsInterface = &utils.UtilsStruct{}
	cmdUtilsMockery = &UtilsStructMockery{}

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
