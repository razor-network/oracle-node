package cmd

import (
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/utils"

	"github.com/ethereum/go-ethereum/ethclient"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// withdrawCmd represents the withdraw command
var withdrawCmd = &cobra.Command{
	Use:   "withdraw",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		utils.CheckError("Error in getting config: ", err)

		password := utils.AssignPassword(cmd.Flags())
		address, _ := cmd.Flags().GetString("address")
		stakerId, _ := cmd.Flags().GetString("stakerId")

		client := utils.ConnectToClient(config.Provider)

		balance, err := utils.FetchBalance(client, address)
		utils.CheckError("Error in fetching balance for account"+address+": ", err)

		if balance.Cmp(big.NewInt(0)) == 0 {
			log.Fatal("Balance is 0. Aborting...")
		}

		_stakerId, ok := new(big.Int).SetString(stakerId, 10)
		if !ok {
			log.Fatal("Set string error in converting staker id")
		}

		checkForCommitStateAndWithdraw(client, types.Account{
			Address:  address,
			Password: password,
		}, config, _stakerId)

	},
}

func checkForCommitStateAndWithdraw(client *ethclient.Client, account types.Account, configurations types.Configurations, stakerId *big.Int) {

	lock, err := utils.GetLock(client, account.Address, stakerId)
	utils.CheckError("Error in fetching lock: ", err)
	withdrawReleasePeriod, err := utils.GetWithdrawReleasePeriod(client, account.Address)
	utils.CheckError("Error in fetching withdraw release period", err)
	withdrawBefore := big.NewInt(0).Add(lock.WithdrawAfter, withdrawReleasePeriod)

	epoch, err := WaitForCommitState(client, account.Address, "withdraw")
	utils.CheckError("Error in fetching epoch: ", err)

	if epoch.Cmp(withdrawBefore) > 0 {
		log.Fatal("Withdrawal period has passed. Cannot withdraw now, please reset the lock!")
	}

	for i := epoch; i.Cmp(withdrawBefore) < 0; {
		if epoch.Cmp(lock.WithdrawAfter) >= 0 && epoch.Cmp(withdrawBefore) <= 0 {
			withdraw(client, types.TransactionOptions{
				Client:         client,
				Password:       account.Password,
				AccountAddress: account.Address,
				ChainId:        core.ChainId,
				Config:         configurations,
			}, epoch, stakerId)
			break
		} else {
			i, err = WaitForCommitState(client, account.Address, "withdraw")
			utils.CheckError("Error in fetching epoch: ", err)
		}
	}
}

func withdraw(client *ethclient.Client, txnOpts types.TransactionOptions, epoch *big.Int, stakerId *big.Int) {
	log.Info("Withdrawing funds...")

	stakeManager := utils.GetStakeManager(client)
	txn, err := stakeManager.Withdraw(utils.GetTxnOpts(txnOpts), epoch, stakerId)
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
		Password string
	)

	withdrawCmd.Flags().StringVarP(&Address, "address", "", "", "address of the user")
	withdrawCmd.Flags().StringVarP(&StakerId, "stakerId", "", "", "staker's id to withdraw")
	withdrawCmd.Flags().StringVarP(&Password, "password", "", "", "password path of user to protect the keystore")

	addrErr := withdrawCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
	stakerIdErr := withdrawCmd.MarkFlagRequired("stakerId")
	utils.CheckError("Staker id error: ", stakerIdErr)
}
