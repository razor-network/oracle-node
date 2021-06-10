package cmd

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/utils"
)

var stakeCmd = &cobra.Command{
	Use:   "stake",
	Short: "Stake some razors",
	Long: `Stake allows user to stake razors in the razor network
	For ex:
	stake -a <amount> --address <address> --password <password>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		if err != nil {
			log.Fatal("Error in getting config: ", err)
		}

		password := utils.PasswordPrompt()
		address, _ := cmd.Flags().GetString("address")
		client := utils.ConnectToClient(config.Provider)
		balance, err := utils.FetchBalance(client, address)
		if err != nil {
			log.Fatalf("Error in fetching balance for account %s: %e", address, err)
		}

		amount, err := cmd.Flags().GetString("amount")
		if err != nil {
			log.Fatal("Error in reading amount", err)
		}

		amountInWei := utils.GetAmountWithChecks(amount, balance)

		accountBalance, err := client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
		if err != nil {
			log.Fatalf("Error in fetching balance of the account: %s\n%s", address, err)
		}

		// FIXME: Not required, remove
		if accountBalance.Cmp(big.NewInt(1e16)) < 0 {
			log.Fatal("Please make sure you hold at least 0.01 ether in your account")
		}

		txnArgs := types.TransactionOptions{
			Client:         client,
			AccountAddress: address,
			Password:       password,
			Amount:         amountInWei,
			ChainId:        core.ChainId,
			GasMultiplier:  config.GasMultiplier,
		}
		approve(txnArgs)
		stakeCoins(txnArgs)
	},
}

func approve(txnArgs types.TransactionOptions) {
	tokenManager := utils.GetTokenManager(txnArgs.Client)
	opts := utils.GetOptions(false, txnArgs.AccountAddress, "")
	allowance, err := tokenManager.Allowance(&opts, common.HexToAddress(txnArgs.AccountAddress), common.HexToAddress(core.StakeManagerAddress))
	if err != nil {
		log.Fatal(err)
	}
	if allowance.Cmp(txnArgs.Amount) >= 0 {
		log.Info("Sufficient allowance, no need to increase")
	} else {
		log.Info("Sending Approve transaction...")
		txnOpts := utils.GetTxnOpts(txnArgs)
		txn, err := tokenManager.Approve(txnOpts, common.HexToAddress(core.StakeManagerAddress), txnArgs.Amount)
		if err != nil {
			log.Fatal("Error in approving: ", err)
		}
		log.Info("Approve transaction sent...\nTxn Hash: ", txn.Hash())
		log.Info("Waiting for transaction to be mined....")
		utils.WaitForBlockCompletion(txnArgs.Client, fmt.Sprintf("%s", txn.Hash()))
	}
}

func stakeCoins(txnArgs types.TransactionOptions) {
	stakeManager := utils.GetStakeManager(txnArgs.Client)
	log.Info("Sending stake transactions...")
	epoch, err := WaitForCommitState(txnArgs.Client, txnArgs.AccountAddress, "stake")
	txnOpts := utils.GetTxnOpts(txnArgs)
	if err != nil {
		log.Fatal("Error in getting commit state: ", err)
	}
	tx, err := stakeManager.Stake(txnOpts, epoch, txnArgs.Amount)
	if err != nil {
		log.Fatal("Error in staking: ", err)
	}
	log.Info("Staked\nTxn Hash: ", tx.Hash())
	utils.WaitForBlockCompletion(txnArgs.Client, fmt.Sprintf("%s", tx.Hash()))
}

func init() {
	rootCmd.AddCommand(stakeCmd)
	var (
		Amount   string
		Address  string
	)

	stakeCmd.Flags().StringVarP(&Amount, "amount", "a", "0", "amount to stake (in Wei)")
	stakeCmd.Flags().StringVarP(&Address, "address", "", "", "address of the staker")

	stakeCmd.MarkFlagRequired("amount")
	stakeCmd.MarkFlagRequired("address")

}
