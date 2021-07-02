package cmd

import (
	"fmt"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/utils"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var unstakeCmd = &cobra.Command{
	Use:   "unstake",
	Short: "Unstake your razors",
	Long: `unstake allows user to unstake their sRzrs in the razor network
	For ex:
	unstake --address <address> --amount <amount_of_sRazors>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		utils.CheckError("Error in getting config: ", err)

		address, _ := cmd.Flags().GetString("address")
		amount, _ := cmd.Flags().GetString("amount")
		stakerId, _ := cmd.Flags().GetString("stakerId")
		password := utils.PasswordPrompt()

		client := utils.ConnectToClient(config.Provider)

		balance, err := utils.FetchBalance(client, address)
		if err != nil {
			log.Fatal("Error in fetching balance: ", err)
		}

		amountInWei := utils.GetAmountWithChecks(amount, balance)

		epoch, err := WaitForCommitState(client, address, "unstake")
		utils.CheckError("Error in fetching epoch: ", err)
		_stakerId, ok := new(big.Int).SetString(stakerId, 10)
		if !ok {
			log.Fatal("Set string error in converting staker id")
		}

		lock, err := utils.GetLock(client, address, _stakerId)
		utils.CheckError("Error in getting lock: ", err)
		if lock.Amount.Cmp(big.NewInt(0)) != 0 {
			log.Fatal("Existing lock")
		}

		stakeManager := utils.GetStakeManager(client)
		txnOpts := utils.GetTxnOpts(types.TransactionOptions{
			Client:         client,
			Password:       password,
			AccountAddress: address,
			ChainId:        core.ChainId,
			GasMultiplier:  config.GasMultiplier,
		})

		log.Info("Unstaking coins")
		txn, err := stakeManager.Unstake(txnOpts, epoch, _stakerId, amountInWei)
		if err != nil {
			log.Fatal("Error in un-staking: ", err)
		}
		log.Infof("Successfully unstaked %s sRazors", amountInWei)
		log.Info("Transaction hash: ", txn.Hash())
		utils.WaitForBlockCompletion(client, fmt.Sprintf("%s", txn.Hash()))
	},
}

func init() {
	rootCmd.AddCommand(unstakeCmd)

	var (
		Address         string
		StakerId		string
		AmountToUnStake string
	)

	unstakeCmd.Flags().StringVarP(&Address, "address", "", "", "user's address")
	unstakeCmd.Flags().StringVarP(&StakerId, "stakerId", "", "", "staker id")
	unstakeCmd.Flags().StringVarP(&AmountToUnStake, "amount", "a", "0", "amount of sRazors to un-stake")

	unstakeCmd.MarkFlagRequired("address")
	unstakeCmd.MarkFlagRequired("stakerId")
	unstakeCmd.MarkFlagRequired("amount")

}
