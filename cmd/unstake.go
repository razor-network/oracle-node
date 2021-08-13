package cmd

import (
	"fmt"
	"github.com/briandowns/spinner"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/utils"
	"time"

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
		autoWithdraw, _ := cmd.Flags().GetBool("autoWithdraw")
		password := utils.PasswordPrompt()

		client := utils.ConnectToClient(config.Provider)

		_amount, ok := new(big.Int).SetString(amount, 10)
		amountInWei := big.NewInt(1).Mul(_amount, big.NewInt(1e18))

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
			Config:         config,
		})

		log.Info("Unstaking coins")
		txn, err := stakeManager.Unstake(txnOpts, epoch, _stakerId, amountInWei)
		utils.CheckError("Error in un-staking: ", err)
		log.Infof("Successfully unstaked %s sRazors", amountInWei)
		log.Info("Transaction hash: ", txn.Hash())
		utils.WaitForBlockCompletion(client, fmt.Sprintf("%s", txn.Hash()))

		if autoWithdraw {
			log.Info("Starting withdrawal now...")
			s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
			s.Start()
			time.Sleep(time.Duration(core.EpochLength) * time.Second)
			s.Stop()
			checkForCommitStateAndWithdraw(client, types.Account{
				Address:  address,
				Password: password,
			}, config, _stakerId)
		}

	},
}

func init() {
	rootCmd.AddCommand(unstakeCmd)

	var (
		Address               string
		StakerId              string
		AmountToUnStake       string
		WithdrawAutomatically bool
	)

	unstakeCmd.Flags().StringVarP(&Address, "address", "", "", "user's address")
	unstakeCmd.Flags().StringVarP(&StakerId, "stakerId", "", "", "staker id")
	unstakeCmd.Flags().StringVarP(&AmountToUnStake, "amount", "a", "0", "amount of sRazors to un-stake")
	unstakeCmd.Flags().BoolVarP(&WithdrawAutomatically, "autoWithdraw", "", false, "withdraw after un-stake automatically")

	addrErr := unstakeCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
	stakerIdErr := unstakeCmd.MarkFlagRequired("stakerId")
	utils.CheckError("Staker Id error: ", stakerIdErr)
	amountErr := unstakeCmd.MarkFlagRequired("amount")
	utils.CheckError("Amount error: ", amountErr)

}
