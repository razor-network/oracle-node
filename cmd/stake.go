package cmd

import (
	"razor/core"
	"razor/core/types"
	"razor/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var stakeCmd = &cobra.Command{
	Use:   "stake",
	Short: "Stake some razors",
	Long: `Stake allows user to stake razors in the razor network

Example:
  ./razor stake --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --value 1000`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		utils.CheckError("Error in getting config: ", err)

		password := utils.AssignPassword(cmd.Flags())
		address, _ := cmd.Flags().GetString("address")
		client := utils.ConnectToClient(config.Provider)
		balance, err := utils.FetchBalance(client, address)
		utils.CheckError("Error in fetching balance for account: "+address, err)

		valueInWei := utils.AssignAmountInWei(cmd.Flags())
		utils.CheckAmountAndBalance(valueInWei, balance)

		utils.CheckEthBalanceIsZero(client, address)

		txnArgs := types.TransactionOptions{
			Client:         client,
			AccountAddress: address,
			Password:       password,
			Amount:         valueInWei,
			ChainId:        core.ChainId,
			Config:         config,
		}
		approve(txnArgs)
		stakeCoins(txnArgs)
	},
}

func approve(txnArgs types.TransactionOptions) {
	tokenManager := utils.GetTokenManager(txnArgs.Client)
	opts := utils.GetOptions(false, txnArgs.AccountAddress, "")
	allowance, err := tokenManager.Allowance(&opts, common.HexToAddress(txnArgs.AccountAddress), common.HexToAddress(core.StakeManagerAddress))
	utils.CheckError("Error in sending allowance: ", err)

	if allowance.Cmp(txnArgs.Amount) >= 0 {
		log.Debug("Sufficient allowance, no need to increase")
	} else {
		log.Info("Sending Approve transaction...")
		txnOpts := utils.GetTxnOpts(txnArgs)
		txn, err := tokenManager.Approve(txnOpts, common.HexToAddress(core.StakeManagerAddress), txnArgs.Amount)
		utils.CheckError("Error in approving", err)
		log.Info("Txn Hash: ", txn.Hash())
		utils.WaitForBlockCompletion(txnArgs.Client, txn.Hash().String())
	}
}

func stakeCoins(txnArgs types.TransactionOptions) {
	stakeManager := utils.GetStakeManager(txnArgs.Client)
	txnOpts := utils.GetTxnOpts(txnArgs)
	epoch, err := WaitForCommitState(txnArgs.Client, txnArgs.AccountAddress, "stake")
	utils.CheckError("Error in getting commit state: ", err)

	log.Info("Sending stake transactions...")
	tx, err := stakeManager.Stake(txnOpts, epoch, txnArgs.Amount)
	utils.CheckError("Error in staking: ", err)

	log.Info("Txn Hash: ", tx.Hash().Hex())
	utils.WaitForBlockCompletion(txnArgs.Client, tx.Hash().String())
}

func init() {
	rootCmd.AddCommand(stakeCmd)
	var (
		Amount   string
		Address  string
		Password string
		Power    string
	)

	stakeCmd.Flags().StringVarP(&Amount, "value", "v", "0", "amount of Razors to stake")
	stakeCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the staker")
	stakeCmd.Flags().StringVarP(&Password, "password", "", "", "password path of staker to protect the keystore")
	stakeCmd.Flags().StringVarP(&Power, "pow", "", "", "power of 10")

	amountErr := stakeCmd.MarkFlagRequired("value")
	utils.CheckError("Value error: ", amountErr)
	addrErr := stakeCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)

}
