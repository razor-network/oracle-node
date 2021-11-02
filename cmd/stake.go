package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"

	"github.com/spf13/cobra"
)

var razorUtils utilsInterface
var tokenManagerUtils tokenManagerInterface
var transactionUtils transactionInterface
var stakeManagerUtils stakeManagerInterface

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
		approveTxnHash, err := approve(txnArgs, razorUtils, tokenManagerUtils, transactionUtils)
		utils.CheckError("Approve error: ", err)

		if approveTxnHash != core.NilHash {
			razorUtils.WaitForBlockCompletion(txnArgs.Client, approveTxnHash.String())
		}

		stakeTxnHash, err := stakeCoins(txnArgs, razorUtils, stakeManagerUtils, transactionUtils)
		utils.CheckError("Stake error: ", err)
		razorUtils.WaitForBlockCompletion(txnArgs.Client, stakeTxnHash.String())

	},
}

func stakeCoins(txnArgs types.TransactionOptions, razorUtils utilsInterface, stakeManagerUtils stakeManagerInterface, transactionUtils transactionInterface) (common.Hash, error) {
	epoch, err := razorUtils.GetEpoch(txnArgs.Client)
	if err != nil {
		return common.Hash{0x00}, err
	}

	log.Info("Sending stake transactions...")
	txnArgs.ContractAddress = core.StakeManagerAddress
	txnArgs.MethodName = "stake"
	txnArgs.Parameters = []interface{}{epoch, txnArgs.Amount}
	txnArgs.ABI = bindings.StakeManagerABI
	txnOpts := razorUtils.GetTxnOpts(txnArgs)
	tx, err := stakeManagerUtils.Stake(txnArgs.Client, txnOpts, epoch, txnArgs.Amount)
	if err != nil {
		return common.Hash{0x00}, err
	}
	log.Info("Txn Hash: ", transactionUtils.Hash(tx).Hex())
	return transactionUtils.Hash(tx), nil
}

func init() {
	razorUtils = Utils{}
	tokenManagerUtils = TokenManagerUtils{}
	transactionUtils = TransactionUtils{}
	stakeManagerUtils = StakeManagerUtils{}

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
