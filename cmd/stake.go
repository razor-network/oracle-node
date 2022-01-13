package cmd

import (
	"context"
	"razor/accounts"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"

	"github.com/ethereum/go-ethereum/common"

	"github.com/spf13/cobra"
)

//var razorUtils utilsInterface
//var transactionUtils transactionInterface
//var stakeManagerUtils stakeManagerInterface

var stakeCmd = &cobra.Command{
	Use:   "stake",
	Short: "Stake some razors",
	Long: `Stake allows user to stake razors in the razor network

Example:
  ./razor stake --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --value 1000`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := cmdUtils.GetConfigData()
		utils.CheckError("Error in getting config: ", err)

		password := utils.AssignPassword(cmd.Flags())
		address, _ := cmd.Flags().GetString("address")
		client := utils.ConnectToClient(config.Provider)
		balance, err := utils.FetchBalance(client, address)
		utils.CheckError("Error in fetching balance for account: "+address, err)

		valueInWei, err := cmdUtils.AssignAmountInWei(cmd.Flags())
		utils.CheckError("Error in getting amount: ", err)

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

		approveTxnHash, err := cmdUtils.Approve(txnArgs)
		utils.CheckError("Approve error: ", err)

		if approveTxnHash != core.NilHash {
			utils.WaitForBlockCompletion(txnArgs.Client, approveTxnHash.String())
		}

		stakeTxnHash, err := cmdUtils.StakeCoins(txnArgs)
		utils.CheckError("Stake error: ", err)
		utils.WaitForBlockCompletion(txnArgs.Client, stakeTxnHash.String())

		if utils.IsFlagPassed("autoVote") {
			isAutoVote, _ := cmd.Flags().GetBool("autoVote")
			if isAutoVote {
				log.Info("Staked!...Starting to vote now.")
				account := types.Account{Address: address, Password: password}
				isRogue, _ := cmd.Flags().GetBool("rogue")
				rogueMode, _ := cmd.Flags().GetStringSlice("rogueMode")
				rogueData := types.Rogue{
					IsRogue:   isRogue,
					RogueMode: rogueMode,
				}
				err := vote(context.Background(), config, client, rogueData, account)
				if err != nil {
					log.Fatal("Error in auto vote: ", err)
				}
			}
		}
	},
}

func (*UtilsStruct) StakeCoins(txnArgs types.TransactionOptions) (common.Hash, error) {
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
	tokenManagerUtils = TokenManagerUtils{}
	stakeManagerUtils = StakeManagerUtils{}
	razorUtils = Utils{}
	cmdUtils = &UtilsStruct{}
	blockManagerUtils = BlockManagerUtils{}
	voteManagerUtils = VoteManagerUtils{}
	transactionUtils = TransactionUtils{}
	flagSetUtils = FLagSetUtils{}
	utils.Options = &utils.OptionsStruct{}
	utils.UtilsInterface = &utils.UtilsStruct{}
	accounts.AccountUtilsInterface = accounts.AccountUtils{}

	rootCmd.AddCommand(stakeCmd)
	var (
		Amount            string
		Address           string
		Password          string
		Power             string
		VoteAutomatically bool
		Rogue             bool
		RogueMode         []string
	)

	stakeCmd.Flags().StringVarP(&Amount, "value", "v", "0", "amount of Razors to stake")
	stakeCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the staker")
	stakeCmd.Flags().StringVarP(&Password, "password", "", "", "password path of staker to protect the keystore")
	stakeCmd.Flags().StringVarP(&Power, "pow", "", "", "power of 10")
	stakeCmd.Flags().BoolVarP(&VoteAutomatically, "autoVote", "", false, "vote after stake automatically")
	stakeCmd.Flags().BoolVarP(&Rogue, "rogue", "r", false, "enable rogue mode to report wrong values")
	stakeCmd.Flags().StringSliceVarP(&RogueMode, "rogueMode", "", []string{}, "type of rogue mode")

	amountErr := stakeCmd.MarkFlagRequired("value")
	utils.CheckError("Value error: ", amountErr)
	addrErr := stakeCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)

}
