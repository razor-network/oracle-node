//Package cmd provides all functions related to command line
package cmd

import (
	"errors"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/rpc"
	"razor/utils"

	"github.com/spf13/pflag"

	"github.com/spf13/cobra"
)

// updateCommissionCmd represents the updateCommission command
var updateCommissionCmd = &cobra.Command{
	Use:   "updateCommission",
	Short: "updateCommission allows a staker to add/update the commission value",
	Long: `Using updateCommission stakers can add or update the commission charged by them
Example:
  ./razor updateCommission --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --commission 10`,
	Run: initialiseUpdateCommission,
}

//This function initialises the ExecuteUpdateCommission function
func initialiseUpdateCommission(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteUpdateCommission(cmd.Flags())
}

//This function sets the flag appropriately and executes the UpdateCommission function
func (*UtilsStruct) ExecuteUpdateCommission(flagSet *pflag.FlagSet) {
	config, rpcParameters, account, err := InitializeCommandDependencies(flagSet)
	utils.CheckError("Error in initialising command dependencies: ", err)

	commission, err := flagSetUtils.GetUint8Commission(flagSet)
	utils.CheckError("Error in getting commission", err)

	stakerId, err := razorUtils.GetStakerId(rpcParameters, account.Address)
	utils.CheckError("Error in getting stakerId", err)

	updateCommissionInput := types.UpdateCommissionInput{
		Commission: commission,
		StakerId:   stakerId,
		Account:    account,
	}

	err = cmdUtils.UpdateCommission(rpcParameters, config, updateCommissionInput)
	utils.CheckError("UpdateCommission error: ", err)
}

//This function allows a staker to add/update the commission value
func (*UtilsStruct) UpdateCommission(rpcParameters rpc.RPCParameters, config types.Configurations, updateCommissionInput types.UpdateCommissionInput) error {
	stakerInfo, err := razorUtils.GetStaker(rpcParameters, updateCommissionInput.StakerId)
	if err != nil {
		log.Error("Error in fetching staker info")
		return err
	}
	log.Debugf("UpdateCommission: Staker Info: %+v", stakerInfo)

	maxCommission, err := razorUtils.GetMaxCommission(rpcParameters)
	if err != nil {
		return err
	}
	log.Debug("UpdateCommission: Maximum Commission: ", maxCommission)

	if updateCommissionInput.Commission == 0 || updateCommissionInput.Commission > maxCommission {
		return errors.New("commission out of range")
	}

	epochLimitForUpdateCommission, err := razorUtils.GetEpochLimitForUpdateCommission(rpcParameters)
	if err != nil {
		return err
	}
	log.Debug("UpdateCommission: Epoch limit to update commission: ", epochLimitForUpdateCommission)

	epoch, err := razorUtils.GetEpoch(rpcParameters)
	if err != nil {
		return err
	}
	log.Debug("UpdateCommission: Current epoch: ", epoch)

	if stakerInfo.EpochCommissionLastUpdated != 0 && (stakerInfo.EpochCommissionLastUpdated+uint32(epochLimitForUpdateCommission)) >= epoch {
		waitFor := uint32(epochLimitForUpdateCommission) - (epoch - stakerInfo.EpochCommissionLastUpdated) + 1
		timeRemaining := uint64(waitFor) * core.EpochLength
		if waitFor == 1 {
			log.Infof("Cannot update commission now. Please wait for %d epoch! (approximately %s)", waitFor, razorUtils.SecondsToReadableTime(int(timeRemaining)))
		} else {
			log.Infof("Cannot update commission now. Please wait for %d epochs! (approximately %s)", waitFor, razorUtils.SecondsToReadableTime(int(timeRemaining)))
		}

		return errors.New("invalid epoch for update")
	}
	txnOpts := types.TransactionOptions{
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.StakeManagerAddress,
		ABI:             bindings.StakeManagerMetaData.ABI,
		MethodName:      "updateCommission",
		Parameters:      []interface{}{updateCommissionInput.Commission},
		Account:         updateCommissionInput.Account,
	}
	updateCommissionTxnOpts, err := razorUtils.GetTxnOpts(rpcParameters, txnOpts)
	if err != nil {
		return err
	}
	log.Infof("Setting the commission value of Staker %d to %d%%", updateCommissionInput.StakerId, updateCommissionInput.Commission)
	client, err := rpcParameters.RPCManager.GetBestRPCClient()
	if err != nil {
		return err
	}

	log.Debug("Executing UpdateCommission transaction with commission = ", updateCommissionInput.Commission)
	txn, err := stakeManagerUtils.UpdateCommission(client, updateCommissionTxnOpts, updateCommissionInput.Commission)
	if err != nil {
		log.Error("Error in setting commission")
		return err
	}
	txnHash := transactionUtils.Hash(txn)
	log.Infof("Txn Hash: %s", txnHash.Hex())
	err = razorUtils.WaitForBlockCompletion(rpcParameters, txnHash.Hex())
	if err != nil {
		log.Error("Error in WaitForBlockCompletion for updateCommission: ", err)
		return err
	}
	return nil
}

func init() {
	var (
		Address    string
		Commission uint8
		Password   string
	)

	rootCmd.AddCommand(updateCommissionCmd)

	updateCommissionCmd.Flags().StringVarP(&Address, "address", "a", "", "your account address")
	updateCommissionCmd.Flags().Uint8VarP(&Commission, "commission", "c", 0, "commission")
	updateCommissionCmd.Flags().StringVarP(&Password, "password", "", "", "password path to protect the keystore")

	addrErr := updateCommissionCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
}
