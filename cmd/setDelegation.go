package cmd

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"
	"strings"

	"github.com/spf13/pflag"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/manifoldco/promptui"

	"github.com/spf13/cobra"
)

var cmdUtils utilsCmdInterface

var setDelegationCmd = &cobra.Command{
	Use:   "setDelegation",
	Short: "setDelegation allows a staker to start accepting/rejecting delegation requests",
	Long: `Using setDelegation, a staker can accept delegation from delegators and charge a commission from them.

Example:
  ./razor setDelegation --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --status true --commission 100
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
	commission, err := utilsStruct.flagSetUtils.GetUint8Commission(flagSet)
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
	}

	if commission != 0 {
		// Call SetCommission if the commission value is provided and the staker hasn't already set commission
		if stakerInfo.Commission == 0 {
			txnOpts.MethodName = "setCommission"
			txnOpts.Parameters = []interface{}{commission}
			setCommissionTxnOpts := utilsStruct.razorUtils.GetTxnOpts(txnOpts, utilsStruct.packageUtils)
			err = utilsStruct.cmdUtils.SetCommission(client, stakerId, setCommissionTxnOpts, commission, utilsStruct)
			if err != nil {
				return err
			}
		}

		// Call DecreaseCommission if the commission value is provided and the staker has already set commission
		if stakerInfo.Commission > 0 && stakerInfo.Commission > commission {
			txnOpts.MethodName = "decreaseCommission"
			txnOpts.Parameters = []interface{}{commission}
			decreaseCommissionTxnOpts := utilsStruct.razorUtils.GetTxnOpts(txnOpts, utilsStruct.packageUtils)
			err = utilsStruct.cmdUtils.DecreaseCommission(client, stakerId, decreaseCommissionTxnOpts, commission, utilsStruct)
			if err != nil {
				return err
			}
		}
	}

	// Fetch updated stakerInfo
	stakerInfo, err = utilsStruct.razorUtils.GetUpdatedStaker(client, address, stakerId)
	if err != nil {
		log.Error("Error in fetching staker info")
		return err
	}

	if stakerInfo.AcceptDelegation != status {
		log.Infof("Setting delegation acceptance of Staker %d to %t", stakerId, status)
		txnOpts.MethodName = "setDelegationAcceptance"
		txnOpts.Parameters = []interface{}{status}
		setDelegationAcceptanceTxnOpts := utilsStruct.razorUtils.GetTxnOpts(txnOpts, utilsStruct.packageUtils)
		delegationAcceptanceTxn, err := utilsStruct.stakeManagerUtils.SetDelegationAcceptance(client, setDelegationAcceptanceTxnOpts, status)
		if err != nil {
			log.Error("Error in setting delegation acceptance")
			return err
		}
		log.Infof("Transaction hash: %s", utilsStruct.transactionUtils.Hash(delegationAcceptanceTxn))
		utilsStruct.razorUtils.WaitForBlockCompletion(client, utilsStruct.transactionUtils.Hash(delegationAcceptanceTxn).String())
	}
	return nil
}

func SetCommission(client *ethclient.Client, stakerId uint32, setCommissionTxnOpts *bind.TransactOpts, commission uint8, utilsStruct UtilsStruct) error {
	log.Infof("Setting the commission value of Staker %d to %d%%", stakerId, commission)
	txn, err := utilsStruct.stakeManagerUtils.SetCommission(client, setCommissionTxnOpts, commission)
	if err != nil {
		log.Error("Error in setting commission")
		return err
	}
	log.Infof("Transaction hash: %s", utilsStruct.transactionUtils.Hash(txn))
	utilsStruct.razorUtils.WaitForBlockCompletion(client, utilsStruct.transactionUtils.Hash(txn).String())
	return nil
}

func DecreaseCommission(client *ethclient.Client, stakerId uint32, decreaseCommissionTxnOpts *bind.TransactOpts, commission uint8, utilsStruct UtilsStruct) error {
	log.Infof("Decreasing the commission value of Staker %d to %d%%", stakerId, commission)
	if utilsStruct.cmdUtils.DecreaseCommissionPrompt() {
		log.Info("Sending DecreaseCommission transaction...")
		decreaseCommissionTxn, err := utilsStruct.stakeManagerUtils.DecreaseCommission(client, decreaseCommissionTxnOpts, commission)
		if err != nil {
			log.Error("Error in decreasing commission")
			return err
		}
		log.Infof("Transaction hash: %s", utilsStruct.transactionUtils.Hash(decreaseCommissionTxn))
		utilsStruct.razorUtils.WaitForBlockCompletion(client, utilsStruct.transactionUtils.Hash(decreaseCommissionTxn).String())
	}
	return nil
}

func DecreaseCommissionPrompt() bool {
	prompt := promptui.Prompt{
		Label:     "Decrease Commission? Once decreased, your commission cannot be increased.",
		IsConfirm: true,
	}
	result, err := prompt.Run()
	utils.CheckError(result, err)
	return strings.ToLower(result) == "y"
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
		Status     string
		Address    string
		Commission uint8
		Password   string
	)

	setDelegationCmd.Flags().StringVarP(&Status, "status", "s", "true", "true for accepting delegation and false for not accepting")
	setDelegationCmd.Flags().StringVarP(&Address, "address", "a", "", "your account address")
	setDelegationCmd.Flags().Uint8VarP(&Commission, "commission", "c", 0, "commission")
	setDelegationCmd.Flags().StringVarP(&Password, "password", "", "", "password path to protect the keystore")

	addrErr := setDelegationCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
}
