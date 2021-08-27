package cmd

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/utils"

	"github.com/ethereum/go-ethereum/ethclient"

	log "github.com/sirupsen/logrus"
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

		password := utils.PasswordPrompt()
		address, _ := cmd.Flags().GetString("address")
		stakerId, _ := cmd.Flags().GetString("stakerId")

		client := utils.ConnectToClient(config.Provider)

		_stakerId, ok := new(big.Int).SetString(stakerId, 10)
		if !ok {
			log.Fatal("Set string error in converting staker id")
		}

		utils.CheckEthBalanceIsZero(client, address)

		checkForCommitStateAndWithdraw(client, types.Account{
			Address:  address,
			Password: password,
		}, config, _stakerId)

	},
}

func checkForCommitStateAndWithdraw(client *ethclient.Client, account types.Account, configurations types.Configurations, stakerId *big.Int) {

	lock, err := utils.GetLock(client, account.Address, stakerId)
	utils.CheckError("Error in fetching lock: ", err)
	log.Info(lock.WithdrawAfter)

	if lock.WithdrawAfter.Cmp(big.NewInt(0)) == 0 {
		log.Fatal("Please unstake Razors before withdrawing.")
	}

	withdrawReleasePeriod, err := utils.GetWithdrawReleasePeriod(client, account.Address)
	utils.CheckError("Error in fetching withdraw release period", err)
	withdrawBefore := big.NewInt(0).Add(lock.WithdrawAfter, withdrawReleasePeriod)
	txnOpts := utils.GetTxnOpts(types.TransactionOptions{
		Client:         client,
		Password:       account.Password,
		AccountAddress: account.Address,
		ChainId:        core.ChainId,
		Config:         configurations,
	})

	epoch, err := utils.GetEpoch(client, account.Address)
	utils.CheckError("Error in fetching epoch: ", err)
	if epoch.Cmp(withdrawBefore) > 0 {
		log.Fatal("Withdrawal period has passed. Cannot withdraw now, please reset the lock!")
	}

	commitStateEpoch, err := WaitForCommitState(client, account.Address, "withdraw")
	for i := commitStateEpoch; i.Cmp(withdrawBefore) < 0; {
		if commitStateEpoch.Cmp(lock.WithdrawAfter) >= 0 && commitStateEpoch.Cmp(withdrawBefore) <= 0 {
			utils.CheckError("Error in fetching epoch: ", err)
			withdraw(client, txnOpts, commitStateEpoch, stakerId)
			break
		} else {
			i, err = WaitForCommitState(client, account.Address, "withdraw")
			utils.CheckError("Error in fetching epoch: ", err)
		}
	}
}

func withdraw(client *ethclient.Client, txnOpts *bind.TransactOpts, epoch *big.Int, stakerId *big.Int) {
	log.Info("Withdrawing funds...")

	stakeManager := utils.GetStakeManager(client)
	txn, err := stakeManager.Withdraw(txnOpts, epoch, stakerId)
	utils.CheckError("Error in withdrawing funds: ", err)

	log.Info("Withdraw Transaction sent.")
	log.Info("Txn Hash: ", txn.Hash())

	utils.WaitForBlockCompletion(client, txn.Hash().String())
}

func init() {
	rootCmd.AddCommand(withdrawCmd)

	var (
		Address  string
		StakerId string
	)

	withdrawCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the user")
	withdrawCmd.Flags().StringVarP(&StakerId, "stakerId", "", "", "staker's id to withdraw")

	addrErr := withdrawCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
	stakerIdErr := withdrawCmd.MarkFlagRequired("stakerId")
	utils.CheckError("Staker id error: ", stakerIdErr)
}
