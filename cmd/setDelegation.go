package cmd

import (
	"github.com/spf13/pflag"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"

	"github.com/spf13/cobra"
)

var cmdUtils utilsCmdInterface

var setDelegationCmd = &cobra.Command{
	Use:   "setDelegation",
	Short: "setDelegation allows a staker to start accepting/rejecting delegation requests",
	Long: `Using setDelegation, a staker can accept delegation from delegators and charge a commission from them.

Example:
  ./razor setDelegation --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --status true 
`,
	Run: func(cmd *cobra.Command, args []string) {
		utilsStruct := UtilsStruct{
			razorUtils:        razorUtils,
			stakeManagerUtils: stakeManagerUtils,
			cmdUtils:          cmdUtils,
			transactionUtils:  transactionUtils,
			flagSetUtils:      flagSetUtils,
			packageUtils:      packageUtils,
		}
		err := utilsStruct.SetDelegation(cmd.Flags())
		utils.CheckError("SetDelegation error: ", err)
	},
}

func (utilsStruct UtilsStruct) SetDelegation(flagSet *pflag.FlagSet) error {

	config, err := utilsStruct.razorUtils.GetConfigData(utilsStruct)
	if err != nil {
		log.Error("Error in getting config")
		return err
	}
	password := utilsStruct.razorUtils.AssignPassword(flagSet)
	address, err := utilsStruct.flagSetUtils.GetStringAddress(flagSet)
	if err != nil {
		return err
	}
	statusString, err := utilsStruct.flagSetUtils.GetStringStatus(flagSet)
	if err != nil {
		return err
	}
	status, err := utilsStruct.razorUtils.ParseBool(statusString)
	if err != nil {
		log.Error("Error in parsing status to boolean")
		return err
	}

	client := utilsStruct.razorUtils.ConnectToClient(config.Provider)

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

	txnOpts := types.TransactionOptions{
		Client:          client,
		Password:        password,
		AccountAddress:  address,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.StakeManagerAddress,
		ABI:             bindings.StakeManagerABI,
		MethodName:      "setDelegationAcceptance",
		Parameters:      []interface{}{status},
	}

	if stakerInfo.AcceptDelegation == status {
		log.Infof("Delegation status already set to %t", status)
		return nil
	}
	log.Infof("Setting delegation acceptance of Staker %d to %t", stakerId, status)
	setDelegationAcceptanceTxnOpts := utilsStruct.razorUtils.GetTxnOpts(txnOpts, utilsStruct.packageUtils)
	delegationAcceptanceTxn, err := utilsStruct.stakeManagerUtils.SetDelegationAcceptance(client, setDelegationAcceptanceTxnOpts, status)
	if err != nil {
		log.Error("Error in setting delegation acceptance")
		return err
	}
	log.Infof("Transaction hash: %s", utilsStruct.transactionUtils.Hash(delegationAcceptanceTxn))
	utilsStruct.razorUtils.WaitForBlockCompletion(client, utilsStruct.transactionUtils.Hash(delegationAcceptanceTxn).String())
	return nil
}

func init() {

	razorUtils = Utils{}
	stakeManagerUtils = StakeManagerUtils{}
	transactionUtils = TransactionUtils{}
	flagSetUtils = FlagSetUtils{}
	cmdUtils = UtilsCmd{}
	packageUtils = utils.PackageUtils{}

	rootCmd.AddCommand(setDelegationCmd)

	var (
		Status   string
		Address  string
		Password string
	)

	setDelegationCmd.Flags().StringVarP(&Status, "status", "s", "true", "true for accepting delegation and false for not accepting")
	setDelegationCmd.Flags().StringVarP(&Address, "address", "a", "", "your account address")
	setDelegationCmd.Flags().StringVarP(&Password, "password", "", "", "password path to protect the keystore")

	addrErr := setDelegationCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
}
