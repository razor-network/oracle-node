package cmd

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/utils"

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
		config, err := GetConfigData()
		utils.CheckError("Error in getting config: ", err)

		password := utils.AssignPassword(cmd.Flags())
		address, _ := cmd.Flags().GetString("address")

		client := utils.ConnectToClient(config.Provider)
		utils.CheckError("Error in fetching staker id: ", err)

		utils.CheckEthBalanceIsZero(client, address)

		stakerId, err := utils.AssignStakerId(cmd.Flags(), client, address)
		utils.CheckError("StakerId error: ", err)

		txn, err := checkForCommitStateAndWithdraw(client, types.Account{
			Address:  address,
			Password: password,
		}, config, stakerId, razorUtils, cmdUtils, stakeManagerUtils, transactionUtils)

		utils.CheckError("Withdraw error: ", err)
		if txn != core.NilHash {
			utils.WaitForBlockCompletion(client, txn.String())
		}
	},
}

func checkForCommitStateAndWithdraw(client *ethclient.Client, account types.Account, configurations types.Configurations, stakerId uint32, razorUtils utilsInterface, cmdUtils utilsCmdInterface, stakeManagerUtils stakeManagerInterface, transactionUtils transactionInterface) (common.Hash, error) {

	lock, err := razorUtils.GetLock(client, account.Address, stakerId)
	if err != nil {
		log.Error("Error in fetching lock")
		return core.NilHash, err
	}

	if lock.WithdrawAfter.Cmp(big.NewInt(0)) == 0 {
		log.Info("Please unstake Razors before withdrawing.")
		return core.NilHash, nil
	}

	withdrawReleasePeriod, err := razorUtils.GetWithdrawReleasePeriod(client, account.Address)
	if err != nil {
		log.Error("Error in fetching withdraw release period")
		return core.NilHash, err
	}
	withdrawBefore := big.NewInt(0).Add(lock.WithdrawAfter, big.NewInt(int64(withdrawReleasePeriod)))
	txnOpts := razorUtils.GetTxnOpts(types.TransactionOptions{
		Client:         client,
		Password:       account.Password,
		AccountAddress: account.Address,
		ChainId:        core.ChainId,
		Config:         configurations,
	})

	epoch, err := razorUtils.GetEpoch(client, account.Address)
	if err != nil {
		log.Error("Error in fetching epoch")
		return core.NilHash, err
	}
	if big.NewInt(int64(epoch)).Cmp(withdrawBefore) > 0 {
		log.Info("Withdrawal period has passed. Cannot withdraw now, please reset the lock!")
		return core.NilHash, nil
	}

	commitStateEpoch, err := razorUtils.WaitForCommitState(client, account.Address, "withdraw")
	if err != nil {
		log.Error("Error in fetching epoch")
		return core.NilHash, err
	}
	for i := commitStateEpoch; big.NewInt(int64(i)).Cmp(withdrawBefore) < 0; {
		if big.NewInt(int64(commitStateEpoch)).Cmp(lock.WithdrawAfter) >= 0 && big.NewInt(int64(commitStateEpoch)).Cmp(withdrawBefore) <= 0 {
			return cmdUtils.Withdraw(client, txnOpts, commitStateEpoch, stakerId, stakeManagerUtils, transactionUtils)
		} else {
			i, err = razorUtils.WaitForCommitStateAgain(client, account.Address, "withdraw")
			if err != nil {
				log.Error("Error in fetching epoch")
				return core.NilHash, err
			}
		}
	}
	return core.NilHash, nil
}

func withdraw(client *ethclient.Client, txnOpts *bind.TransactOpts, epoch uint32, stakerId uint32, stakeManagerUtils stakeManagerInterface, transactionUtils transactionInterface) (common.Hash, error) {
	log.Info("Withdrawing funds...")

	txn, err := stakeManagerUtils.Withdraw(client, txnOpts, epoch, stakerId)
	if err != nil {
		log.Error("Error in withdrawing funds")
		return core.NilHash, err
	}

	log.Info("Txn Hash: ", transactionUtils.Hash(txn))

	return transactionUtils.Hash(txn), nil
}

func init() {

	razorUtils = Utils{}
	transactionUtils = TransactionUtils{}
	stakeManagerUtils = StakeManagerUtils{}
	cmdUtils = UtilsCmd{}

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
