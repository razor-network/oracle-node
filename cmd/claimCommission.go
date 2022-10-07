//This function add the following command to the root command
package cmd

import (
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/logger"
	"razor/pkg/bindings"
	"razor/utils"

	"github.com/spf13/pflag"

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
	config, err := cmdUtils.GetConfigData()
	utils.CheckError("Error in getting config: ", err)

	client := razorUtils.ConnectToClient(config.Provider)

	address, err := flagSetUtils.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)

	logger.SetLoggerParameters(client, address)
	razorUtils.AssignLogFile(flagSet)

	password := razorUtils.AssignPassword(flagSet)

	razorUtils.CheckEthBalanceIsZero(client, address)

	stakerId, err := razorUtils.GetStakerId(client, address)
	utils.CheckError("Error in getting stakerId: ", err)
	callOpts := razorUtils.GetOptions()

	stakerInfo, err := stakeManagerUtils.StakerInfo(client, &callOpts, stakerId)
	utils.CheckError("Error in getting stakerInfo: ", err)

	if stakerInfo.StakerReward.Cmp(big.NewInt(0)) > 0 {
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
		utils.CheckError("Error in claiming stake reward: ", err)

		err = razorUtils.WaitForBlockCompletion(client, transactionUtils.Hash(txn).String())
		utils.CheckError("Error in WaitForBlockCompletion for claimCommission: ", err)
	} else {
		log.Error("no commission to claim")
		return
	}
}

func init() {
	rootCmd.AddCommand(claimCommissionCmd)

	var (
		Address  string
		Password string
	)

	claimCommissionCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the staker")
	claimCommissionCmd.Flags().StringVarP(&Password, "password", "", "", "password path of staker to protect the keystore")

	addrErr := claimCommissionCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
}
