//This function add the following command to the root command
package cmd

import (
	"github.com/spf13/pflag"
	"razor/core"
	"razor/core/types"
	"razor/logger"
	"razor/pkg/bindings"
	"razor/utils"

	"github.com/spf13/cobra"
)

var claimCommissionCmd = &cobra.Command{
	Use:   "claimCommission",
	Short: "claim commission from staker",
	Long: `staker can claim the rewards earned from delegator's pool share as commission
Example:
  ./razor claimCommission --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --logFile claimComm`,
	Run: func(cmd *cobra.Command, args []string) {
		cmdUtils.ClaimCommission(cmd.Flags())
	},
}

//This function allows the staker to claim the rewards earned from delegator's pool share as commission
func (*UtilsStruct) ClaimCommission(flagSet *pflag.FlagSet) {
	razorUtils.AssignLogFile(flagSet)
	address, err := flagSetUtils.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)

	logger.Address = address

	config, err := cmdUtils.GetConfigData()
	utils.CheckError("Error in getting config: ", err)
	password := razorUtils.AssignPassword()
	client := razorUtils.ConnectToClient(config.Provider)

	razorUtils.CheckEthBalanceIsZero(client, address)

	txnOpts := razorUtils.GetTxnOpts(types.TransactionOptions{
		Client:          client,
		AccountAddress:  address,
		Password:        password,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.StakeManagerAddress,
		MethodName:      "claimStakerReward",
		Parameters:      []interface{}{},
		ABI:             bindings.StakeManagerABI,
	})

	log.Info("Claiming commission")

	txn, err := stakeManagerUtils.ClaimStakeReward(client, txnOpts)
	if err != nil {
		log.Fatal("Error in claiming stake reward: ", err)
	}

	razorUtils.WaitForBlockCompletion(client, transactionUtils.Hash(txn).String())

}

func init() {
	rootCmd.AddCommand(claimCommissionCmd)

	var (
		Address string
	)

	claimCommissionCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the staker")

	addrErr := claimCommissionCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
}
