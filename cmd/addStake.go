//Package cmd provides all functions related to command line
package cmd

import (
	"razor/core"
	"razor/core/types"
	"razor/logger"
	"razor/pkg/bindings"
	"razor/utils"

	"github.com/spf13/pflag"

	"github.com/ethereum/go-ethereum/common"

	"github.com/spf13/cobra"
)

var stakeCmd = &cobra.Command{
	Use:   "addStake",
	Short: "Stake some razors",
	Long: `addStake allows user to stake razors in the razor network

Example:
  ./razor addStake --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --value 1000 --logFile addStake`,
	Run: initialiseStake,
}

//This function initialises the ExecuteStake function
func initialiseStake(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteStake(cmd.Flags())
}

//This function sets the flags appropriately and executes the StakeCoins function
func (*UtilsStruct) ExecuteStake(flagSet *pflag.FlagSet) {
	config, err := cmdUtils.GetConfigData()
	utils.CheckError("Error in getting config: ", err)

	client := razorUtils.ConnectToClient(config.Provider)

	address, err := flagSetUtils.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)

	logger.SetLoggerParameters(client, address)
	razorUtils.AssignLogFile(flagSet)

	password := razorUtils.AssignPassword(flagSet)

	balance, err := razorUtils.FetchBalance(client, address)
	utils.CheckError("Error in fetching razor balance for account: "+address, err)
	valueInWei, err := cmdUtils.AssignAmountInWei(flagSet)
	utils.CheckError("Error in getting amount: ", err)

	utils.CheckAmountAndBalance(valueInWei, balance)

	razorUtils.CheckEthBalanceIsZero(client, address)

	minSafeRazor, err := utils.UtilsInterface.GetMinSafeRazor(client)
	utils.CheckError("Error in getting minimum safe razor amount: ", err)

	stakerId, err := razorUtils.GetStakerId(client, address)
	utils.CheckError("Error in getting stakerId: ", err)

	if valueInWei.Cmp(minSafeRazor) < 0 && stakerId == 0 {
		log.Fatal("The amount of razors entered is below min safe value.")
	}

	if stakerId != 0 {
		staker, err := razorUtils.GetStaker(client, stakerId)
		utils.CheckError("Error in getting staker: ", err)

		if staker.IsSlashed {
			log.Fatal("Staker is slashed, cannot stake")
		}
	}

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
		err = razorUtils.WaitForBlockCompletion(txnArgs.Client, approveTxnHash.Hex())
		utils.CheckError("Error in WaitForBlockCompletion for approve: ", err)
	}

	stakeTxnHash, err := cmdUtils.StakeCoins(txnArgs)
	utils.CheckError("Stake error: ", err)

	err = razorUtils.WaitForBlockCompletion(txnArgs.Client, stakeTxnHash.Hex())
	utils.CheckError("Error in WaitForBlockCompletion for stake: ", err)
}

//This function allows the user to stake razors in the razor network and returns the hash
func (*UtilsStruct) StakeCoins(txnArgs types.TransactionOptions) (common.Hash, error) {
	epoch, err := razorUtils.GetEpoch(txnArgs.Client)
	if err != nil {
		return core.NilHash, err
	}

	log.Info("Sending stake transactions...")
	txnArgs.ContractAddress = core.StakeManagerAddress
	txnArgs.MethodName = "stake"
	txnArgs.Parameters = []interface{}{epoch, txnArgs.Amount}
	txnArgs.ABI = bindings.StakeManagerMetaData.ABI
	txnOpts := razorUtils.GetTxnOpts(txnArgs)
	txn, err := stakeManagerUtils.Stake(txnArgs.Client, txnOpts, epoch, txnArgs.Amount)
	if err != nil {
		return core.NilHash, err
	}
	txnHash := transactionUtils.Hash(txn)
	log.Info("Txn Hash: ", txnHash.Hex())
	return txnHash, nil
}

func init() {
	rootCmd.AddCommand(stakeCmd)
	var (
		Amount   string
		Address  string
		Password string
		WeiRazor bool
	)

	stakeCmd.Flags().StringVarP(&Amount, "value", "v", "0", "amount of Razors to stake")
	stakeCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the staker")
	stakeCmd.Flags().StringVarP(&Password, "password", "", "", "password path of staker to protect the keystore")
	stakeCmd.Flags().BoolVarP(&WeiRazor, "weiRazor", "", false, "value can be passed in wei")

	amountErr := stakeCmd.MarkFlagRequired("value")
	utils.CheckError("Value error: ", amountErr)
	addrErr := stakeCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)

}
