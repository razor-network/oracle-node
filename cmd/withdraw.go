package cmd

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
	Run: func(cmd *cobra.Command, args []string) {
		utilsStruct := UtilsStruct{
			razorUtils:        razorUtils,
			stakeManagerUtils: stakeManagerUtils,
			cmdUtils:          cmdUtils,
			transactionUtils:  transactionUtils,
			flagSetUtils:      flagSetUtils,
		}

		config, err := cmdUtilsMockery.GetConfigData()
		utils.CheckError("Error in getting config: ", err)

		password := utils.AssignPassword(cmd.Flags())
		address, _ := cmd.Flags().GetString("address")

		client := utils.ConnectToClient(config.Provider)
		utils.CheckError("Error in fetching staker id: ", err)

		utils.CheckEthBalanceIsZero(client, address)

		stakerId, err := utils.AssignStakerId(cmd.Flags(), client, address)
		utils.CheckError("StakerId error: ", err)

		txn, err := withdrawFunds(client, types.Account{
			Address:  address,
			Password: password,
		}, config, stakerId, utilsStruct)

		utils.CheckError("Withdraw error: ", err)
		if txn != core.NilHash {
			utils.WaitForBlockCompletion(client, txn.String())
		}
	},
}

func withdrawFunds(client *ethclient.Client, account types.Account, configurations types.Configurations, stakerId uint32, utilsStruct UtilsStruct) (common.Hash, error) {

	lock, err := utilsStruct.razorUtils.GetLock(client, account.Address, stakerId)
	if err != nil {
		log.Error("Error in fetching lock")
		return core.NilHash, err
	}

	if lock.WithdrawAfter.Cmp(big.NewInt(0)) == 0 {
		log.Info("Please unstake Razors before withdrawing.")
		return core.NilHash, nil
	}

	withdrawReleasePeriod, err := utilsStruct.razorUtils.GetWithdrawReleasePeriod(client, account.Address)
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
	epoch, err := utilsStruct.razorUtils.GetEpoch(client)
	if err != nil {
		log.Error("Error in fetching epoch")
		return core.NilHash, err
	}
	if big.NewInt(int64(epoch)).Cmp(withdrawBefore) > 0 {
		log.Info("Withdrawal period has passed. Cannot withdraw now, please reset the lock!")
		return core.NilHash, nil
	}

	txnArgs.Parameters = []interface{}{epoch, stakerId}
	txnOpts := utilsStruct.razorUtils.GetTxnOpts(txnArgs)

	for i := epoch; big.NewInt(int64(i)).Cmp(withdrawBefore) < 0; {
		if big.NewInt(int64(epoch)).Cmp(lock.WithdrawAfter) >= 0 && big.NewInt(int64(epoch)).Cmp(withdrawBefore) <= 0 {
			return utilsStruct.cmdUtils.Withdraw(client, txnOpts, stakerId, utilsStruct)
		}
		log.Debug("Waiting for lock period to get over....")
		// Wait for 30 seconds if lock period isn't over
		utilsStruct.razorUtils.Sleep(30 * time.Second)
		epoch, err = utilsStruct.razorUtils.GetUpdatedEpoch(client)
		if err != nil {
			log.Error("Error in fetching epoch")
			return core.NilHash, err
		}
	}
	return core.NilHash, nil
}

func withdraw(client *ethclient.Client, txnOpts *bind.TransactOpts, stakerId uint32, utilsStruct UtilsStruct) (common.Hash, error) {
	log.Info("Withdrawing funds...")

	txn, err := utilsStruct.stakeManagerUtils.Withdraw(client, txnOpts, stakerId)
	if err != nil {
		log.Error("Error in withdrawing funds")
		return core.NilHash, err
	}

	log.Info("Txn Hash: ", utilsStruct.transactionUtils.Hash(txn))

	return utilsStruct.transactionUtils.Hash(txn), nil
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
