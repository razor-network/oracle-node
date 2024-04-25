//This function add the following command to the root command
package cmd

import (
	"math/big"
	"razor/core"
	"razor/core/types"
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
	log.Debugf("ClaimCommission: Config: %+v", config)

	client := razorUtils.ConnectToClient(config.Provider)

	address, err := flagSetUtils.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)
	log.Debug("ClaimCommission: Address: ", address)

	log.SetLoggerParameters(client, address)

	log.Debug("Checking to assign log file...")
	fileUtils.AssignLogFile(flagSet, config)

	log.Debug("Getting password...")
	password := razorUtils.AssignPassword(flagSet)

	err = razorUtils.CheckPassword(address, password)
	utils.CheckError("Error in fetching private key from given password: ", err)

	stakerId, err := razorUtils.GetStakerId(client, address)
	utils.CheckError("Error in getting stakerId: ", err)
	log.Debug("ClaimCommission: Staker Id: ", stakerId)
	callOpts := razorUtils.GetOptions()

	stakerInfo, err := stakeManagerUtils.StakerInfo(client, &callOpts, stakerId)
	utils.CheckError("Error in getting stakerInfo: ", err)
	log.Debugf("ClaimCommission: Staker Info: %+v", stakerInfo)

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
			ABI:             bindings.StakeManagerMetaData.ABI,
		})

		log.Info("Claiming commission...")

		log.Debug("Executing ClaimStakeReward transaction...")
		txn, err := stakeManagerUtils.ClaimStakerReward(client, txnOpts)
		utils.CheckError("Error in claiming stake reward: ", err)

		txnHash := transactionUtils.Hash(txn)
		log.Info("Txn Hash: ", txnHash.Hex())
		err = razorUtils.WaitForBlockCompletion(client, txnHash.Hex())
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
