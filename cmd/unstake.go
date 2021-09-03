package cmd

import (
	"fmt"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/utils"
	"time"

	"github.com/briandowns/spinner"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var unstakeCmd = &cobra.Command{
	Use:   "unstake",
	Short: "Unstake your razors",
	Long: `unstake allows user to unstake their sRzrs in the razor network

Example:	
  ./razor unstake --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --value 1000 --autoWithdraw
	`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		utils.CheckError("Error in getting config: ", err)

		password := utils.AssignPassword(cmd.Flags())
		address, _ := cmd.Flags().GetString("address")
		value, _ := cmd.Flags().GetString("value")
		autoWithdraw, _ := cmd.Flags().GetBool("autoWithdraw")

		client := utils.ConnectToClient(config.Provider)

		_value, ok := new(big.Int).SetString(value, 10)
		if !ok {
			log.Fatal("SetString: error")
		}
		valueInWei := big.NewInt(1).Mul(_value, big.NewInt(1e18))

		utils.CheckEthBalanceIsZero(client, address)

		stakerId, err := utils.GetStakerId(client, address)
		utils.CheckError("Error in fetching staker: ", err)

		lock, err := utils.GetLock(client, address, stakerId)
		utils.CheckError("Error in getting lock: ", err)
		if lock.Amount.Cmp(big.NewInt(0)) != 0 {
			log.Fatal("Existing lock")
		}

		stakeManager := utils.GetStakeManager(client)
		txnOpts := utils.GetTxnOpts(types.TransactionOptions{
			Client:         client,
			Password:       password,
			AccountAddress: address,
			Amount:         valueInWei,
			ChainId:        core.ChainId,
			Config:         config,
		})

		epoch, err := WaitForCommitState(client, address, "unstake")
		utils.CheckError("Error in fetching epoch: ", err)
		log.Info("Unstaking coins")
		txn, err := stakeManager.Unstake(txnOpts, epoch, stakerId, valueInWei)
		utils.CheckError("Error in un-staking: ", err)
		log.Infof("Successfully unstaked %s sRazors", valueInWei)
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
			}, config, stakerId)
		}
	},
}

func Unstake(txnArgs types.TransactionOptions, stakerId uint32) {
	lock, err := utils.GetLock(txnArgs.Client, txnArgs.AccountAddress, stakerId)
	utils.CheckError("Error in getting lock: ", err)
	if lock.Amount.Cmp(big.NewInt(0)) != 0 {
		log.Fatal("Existing lock")
	}

	stakeManager := utils.GetStakeManager(txnArgs.Client)
	txnOpts := utils.GetTxnOpts(txnArgs)

	epoch, err := WaitForCommitState(txnArgs.Client, txnArgs.AccountAddress, "unstake")
	utils.CheckError("Error in fetching epoch: ", err)
	log.Info("Unstaking coins")
	txn, err := stakeManager.Unstake(txnOpts, epoch, stakerId, txnArgs.Amount)
	utils.CheckError("Error in un-staking: ", err)
	log.Infof("Successfully unstaked %s sRazors", txnArgs.Amount)
	log.Info("Transaction hash: ", txn.Hash())
	utils.WaitForBlockCompletion(txnArgs.Client, fmt.Sprintf("%s", txn.Hash()))
}

func AutoWithdraw(txnArgs types.TransactionOptions, stakerId uint32) {
	log.Info("Starting withdrawal now...")
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Start()
	time.Sleep(time.Duration(core.EpochLength) * time.Second)
	s.Stop()
	checkForCommitStateAndWithdraw(txnArgs.Client, types.Account{
		Address:  txnArgs.AccountAddress,
		Password: txnArgs.Password,
	}, txnArgs.Config, stakerId)
}

func init() {
	rootCmd.AddCommand(unstakeCmd)

	var (
		Address               string
		AmountToUnStake       string
		WithdrawAutomatically bool
		Password              string
	)

	unstakeCmd.Flags().StringVarP(&Address, "address", "a", "", "user's address")
	unstakeCmd.Flags().StringVarP(&AmountToUnStake, "value", "v", "0", "value of sRazors to un-stake")
	unstakeCmd.Flags().BoolVarP(&WithdrawAutomatically, "autoWithdraw", "", false, "withdraw after un-stake automatically")
	unstakeCmd.Flags().StringVarP(&Password, "password", "", "", "password path to protect the keystore")

	addrErr := unstakeCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
	valueErr := unstakeCmd.MarkFlagRequired("value")
	utils.CheckError("Value error: ", valueErr)

}
