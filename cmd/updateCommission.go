//Package cmd provides all functions related to command line
package cmd

import (
	"errors"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/logger"
	"razor/pkg/bindings"
	"razor/utils"

	"github.com/ethereum/go-ethereum/ethclient"
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
	razorUtils.AssignLogFile(flagSet)
	address, err := flagSetUtils.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address", err)

	logger.Address = address

	config, err := cmdUtils.GetConfigData()
	utils.CheckError("Error in getting config", err)

	client := razorUtils.ConnectToClient(config.Provider)
	password := razorUtils.AssignPassword()

	commission, err := flagSetUtils.GetUint8Commission(flagSet)
	utils.CheckError("Error in getting commission", err)

	stakerId, err := razorUtils.GetStakerId(client, address)
	utils.CheckError("Error in getting stakerId", err)

	err = cmdUtils.UpdateCommission(config, client, types.UpdateCommissionInput{
		Commission: commission,
		Address:    address,
		Password:   password,
		StakerId:   stakerId,
	})
	utils.CheckError("SetDelegation error: ", err)
}

//This function allows a staker to add/update the commission value
func (*UtilsStruct) UpdateCommission(config types.Configurations, client *ethclient.Client, updateCommissionInput types.UpdateCommissionInput) error {
	stakerInfo, err := razorUtils.GetStaker(client, updateCommissionInput.StakerId)
	if err != nil {
		log.Error("Error in fetching staker info")
		return err
	}

	maxCommission, err := razorUtils.GetMaxCommission(client)
	if err != nil {
		return err
	}

	if updateCommissionInput.Commission == 0 || updateCommissionInput.Commission > maxCommission {
		return errors.New("commission out of range")
	}

	epochLimitForUpdateCommission, err := razorUtils.GetEpochLimitForUpdateCommission(client)
	if err != nil {
		return err
	}

	epoch, err := razorUtils.GetEpoch(client)
	if err != nil {
		return err
	}

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
		Client:          client,
		Password:        updateCommissionInput.Password,
		AccountAddress:  updateCommissionInput.Address,
		ChainId:         big.NewInt(config.ChainId),
		Config:          config,
		ContractAddress: core.StakeManagerAddress,
		ABI:             bindings.StakeManagerMetaData.ABI,
		MethodName:      "updateCommission",
		Parameters:      []interface{}{updateCommissionInput.Commission},
	}
	updateCommissionTxnOpts := razorUtils.GetTxnOpts(txnOpts)
	log.Infof("Setting the commission value of Staker %d to %d%%", updateCommissionInput.StakerId, updateCommissionInput.Commission)
	txn, err := stakeManagerUtils.UpdateCommission(client, updateCommissionTxnOpts, updateCommissionInput.Commission)
	if err != nil {
		log.Error("Error in setting commission")
		return err
	}
	txnHash := transactionUtils.Hash(txn)
	log.Infof("Transaction hash: %s", txnHash)
	err = razorUtils.WaitForBlockCompletion(client, txnHash.Hex())
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
	)

	rootCmd.AddCommand(updateCommissionCmd)

	updateCommissionCmd.Flags().StringVarP(&Address, "address", "a", "", "your account address")
	updateCommissionCmd.Flags().Uint8VarP(&Commission, "commission", "c", 0, "commission")

	addrErr := updateCommissionCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
}
