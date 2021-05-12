package cmd

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	types2 "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"math"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/utils"
	"time"
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

		password, _ := cmd.Flags().GetString("password")
		address, _ := cmd.Flags().GetString("address")
		client := utils.ConnectToClient(config.Provider)
		balance := utils.FetchBalance(client, address)

		amount, err := cmd.Flags().GetString("amount")
		if err != nil {
			log.Fatal("Error in reading amount", err)
		}

		_amount, ok := new(big.Int).SetString(amount, 10)
		if !ok {
			log.Fatal("SetString: error")
		}
		amountInWei := big.NewInt(1).Mul(_amount, big.NewInt(1e18))

		if amountInWei.Cmp(balance) > 0 {
			log.Fatal("Not enough balance")
		}

		accountBalance, err := client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
		if err != nil {
			log.Fatalf("Error in fetching balance of the account: %s\n%s", address, err)
		}

		if accountBalance.Cmp(big.NewInt(1e16)) < 0 {
			log.Fatal("Please make sure you hold at least 0.01 ether in your account")
		}

		txnArgs := types.TransactionOptions{
			Client:         client,
			AccountAddress: address,
			Password:       password,
			Amount:         amountInWei,
			ChainId:        config.ChainId,
			GasMultiplier:  config.GasMultiplier,
		}
		approve(txnArgs)
		stakeCoins(txnArgs)
	},
}

func approve(txnArgs types.TransactionOptions) {
	coinContract := utils.GetCoinContract(txnArgs.Client)
	opts := utils.GetOptions(false, txnArgs.AccountAddress, "")
	allowance, err := coinContract.Allowance(&opts, common.HexToAddress(txnArgs.AccountAddress), common.HexToAddress(core.StakeManagerAddress))
	if err != nil {
		log.Fatal(err)
	}
	if allowance.Cmp(txnArgs.Amount) >= 0 {
		log.Info("Sufficient allowance, no need to increase")
	} else {
		log.Info("Sending Approve transaction...")
		txnOpts := utils.GetTxnOpts(txnArgs)
		txn, err := coinContract.Approve(txnOpts, common.HexToAddress(core.StakeManagerAddress), txnArgs.Amount)
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
	var epoch *big.Int
	for true {
		epoch = utils.GetEpoch(txnArgs.Client, txnArgs.AccountAddress)
		state := getDelayedState(txnArgs.Client)
		log.Info("Epoch ", epoch)
		log.Info("State ", state)
		if state != 0 {
			log.Info("Can only stake during state 0 (commit). Retrying in 1 second...")
			time.Sleep(1 * time.Second)
		} else {
			break
		}
	}
	log.Info("Sending stake transactions...")
	txnOpts := utils.GetTxnOpts(txnArgs)
	tx, err := stakeManager.Stake(txnOpts, epoch, txnArgs.Amount)
	if err != nil {
		log.Fatal("Error in staking: ", err)
	}
	log.Info("Staked\nTxn Hash: ", tx.Hash())
	utils.WaitForBlockCompletion(txnArgs.Client, fmt.Sprintf("%s", tx.Hash()))
}

func getDelayedState(client *ethclient.Client) int64 {
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal("Error in fetching latest block number: ", err)
	}

	if blockNumber%(core.BlockDivider) > 7 || blockNumber%(core.BlockDivider) < 1 {
		return -1
	}
	state := math.Floor(float64(blockNumber / core.BlockDivider))
	return int64(state) % core.NumberOfStates
}

func checkIfTxnIsSuccess(transaction *types2.Transaction) {

}

func init() {
	rootCmd.AddCommand(stakeCmd)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	var (
		Amount   string
		Address  string
		Password string
	)

	stakeCmd.Flags().StringVarP(&Amount, "amount", "a", "0", "amount to stake (in Wei)")
	stakeCmd.Flags().StringVarP(&Address, "address", "", "", "address of the staker")
	stakeCmd.Flags().StringVarP(&Password, "password", "", "", "password to unlock account")

	stakeCmd.MarkFlagRequired("amount")
	stakeCmd.MarkFlagRequired("address")
	stakeCmd.MarkFlagRequired("password")

}
