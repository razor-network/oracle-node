package cmd

import (
	"errors"
	"github.com/ethereum/go-ethereum/ethclient"
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
	Short: "updateCommission allows a staker to add/update the commission value",
	Long: `Using updateCommission stakers can add or update the commission charged by them
Example:
  ./razor updateCommission --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --commission 10`,
	Run: func(cmd *cobra.Command, args []string) {
		utilsStruct := UtilsStruct{
			razorUtils:        razorUtils,
			stakeManagerUtils: stakeManagerUtils,
			cmdUtils:          cmdUtils,
			transactionUtils:  transactionUtils,
			flagSetUtils:      flagSetUtils,
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

	return utilsStruct.cmdUtils.ExecuteUpdateCommission(client, types.UpdateCommissionInput{
		Commission: commission,
		Config:     config,
		Address:    address,
		Password:   password,
		StakerId:   stakerId,
	}, utilsStruct)
}

func ExecuteUpdateCommission(client *ethclient.Client, input types.UpdateCommissionInput, utilsStruct UtilsStruct) error {
	stakerInfo, err := utilsStruct.razorUtils.GetStaker(client, input.Address, input.StakerId)
	if err != nil {
		log.Error("Error in fetching staker info")
		return err
	}

	maxCommission, err := utilsStruct.razorUtils.GetMaxCommission(client)
	if err != nil {
		return err
	}

	if input.Commission == 0 || input.Commission > maxCommission {
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

	if stakerInfo.EpochCommissionLastUpdated != 0 && (stakerInfo.EpochCommissionLastUpdated+uint32(epochLimitForUpdateCommission)) >= epoch {
		return errors.New("invalid epoch for update")
	}
	txnOpts := types.TransactionOptions{
		Client:          client,
		Password:        input.Password,
		AccountAddress:  input.Address,
		ChainId:         core.ChainId,
		Config:          input.Config,
		ContractAddress: core.StakeManagerAddress,
		ABI:             bindings.StakeManagerABI,
		MethodName:      "updateCommission",
		Parameters:      []interface{}{input.Commission},
	}
	updateCommissionTxnOpts := utilsStruct.razorUtils.GetTxnOpts(txnOpts)
	log.Infof("Setting the commission value of Staker %d to %d%%", input.StakerId, input.Commission)
	txn, err := utilsStruct.stakeManagerUtils.UpdateCommission(client, updateCommissionTxnOpts, input.Commission)
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
