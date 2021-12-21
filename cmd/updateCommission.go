package cmd

import (
	"errors"
	"github.com/spf13/pflag"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"

	"github.com/spf13/cobra"
)

// updateCommissionCmd represents the updateCommission command
var updateCommissionCmd = &cobra.Command{
	Use:   "updateCommission",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		utilsStruct := UtilsStruct{
			razorUtils:        razorUtils,
			stakeManagerUtils: stakeManagerUtils,
			cmdUtils:          cmdUtils,
			transactionUtils:  transactionUtils,
			flagSetUtils:      flagSetUtils,
			packageUtils:      packageUtils,
		}
		err := utilsStruct.UpdateCommission(cmd.Flags())
		utils.CheckError("SetDelegation error: ", err)
	},
}

func (utilsStruct UtilsStruct) UpdateCommission(flagSet *pflag.FlagSet) error {
	config, err := utilsStruct.razorUtils.GetConfigData(utilsStruct)
	if err != nil {
		log.Error("Error in getting config")
		return err
	}
	client := utilsStruct.razorUtils.ConnectToClient(config.Provider)
	password := utilsStruct.razorUtils.AssignPassword(flagSet)
	address, err := utilsStruct.flagSetUtils.GetStringAddress(flagSet)
	if err != nil {
		return err
	}
	commission, err := utilsStruct.flagSetUtils.GetUint8Commission(flagSet)
	if err != nil {
		return err
	}

	stakerId, err := utilsStruct.razorUtils.GetStakerId(client, address)
	if err != nil {
		log.Error("Error in fetching staker id")
		return err
	}

	stakerInfo, err := utilsStruct.razorUtils.GetStaker(client, address, stakerId)
	if err != nil {
		log.Error("Error in fetching staker info")
		return err
	}

	maxCommission, err := utilsStruct.razorUtils.GetMaxCommission(client)
	if err != nil {
		return err
	}

	if commission == 0 || commission > maxCommission {
		return errors.New("commission out of range")
	}

	epochLimitForUpdateCommission, err := utilsStruct.razorUtils.GetEpochLimitForUpdateCommission(client)
	if err != nil {
		return err
	}

	epoch, err := utilsStruct.razorUtils.GetEpoch(client)
	if err != nil {
		return err
	}

	if (stakerInfo.EpochCommissionLastUpdated + uint32(epochLimitForUpdateCommission)) >= epoch {
		return errors.New("invalid epoch for update")
	}
	txnOpts := types.TransactionOptions{
		Client:          client,
		Password:        password,
		AccountAddress:  address,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.StakeManagerAddress,
		ABI:             bindings.StakeManagerABI,
		MethodName:      "updateCommission",
		Parameters:      []interface{}{commission},
	}

	updateCommissionTxnOpts := utilsStruct.razorUtils.GetTxnOpts(txnOpts, utilsStruct.packageUtils)
	log.Infof("Setting the commission value of Staker %d to %d%%", stakerId, commission)
	txn, err := utilsStruct.stakeManagerUtils.UpdateCommission(client, updateCommissionTxnOpts, commission)
	if err != nil {
		log.Error("Error in setting commission")
		return err
	}
	txnHash := utilsStruct.transactionUtils.Hash(txn)
	log.Infof("Transaction hash: %s", txnHash)
	utilsStruct.razorUtils.WaitForBlockCompletion(client, txnHash.String())
	return nil
}

func init() {
	razorUtils = Utils{}
	stakeManagerUtils = StakeManagerUtils{}
	transactionUtils = TransactionUtils{}
	flagSetUtils = FlagSetUtils{}
	cmdUtils = UtilsCmd{}
	packageUtils = utils.RazorUtils{}

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
