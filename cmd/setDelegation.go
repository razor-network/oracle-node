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
		err := SetDelegation(cmd.Flags(), razorUtils, stakeManagerUtils, cmdUtils, transactionUtils, flagSetUtils)
		utils.CheckError("SetDelegation error: ", err)
	},
}

func SetDelegation(flagSet *pflag.FlagSet, razorUtils utilsInterface, stakeManagerUtils stakeManagerInterface, cmdUtils utilsCmdInterface, transactionUtils transactionInterface, flagSetUtils flagSetInterface) error {

	config, err := razorUtils.GetConfigData()
	if err != nil {
		log.Error("Error in getting config")
		return err
	}
	password := razorUtils.AssignPassword(flagSet)
	address, err := flagSetUtils.GetStringAddress(flagSet)
	if err != nil {
		return err
	}
	statusString, err := flagSetUtils.GetStringStatus(flagSet)
	if err != nil {
		return err
	}
	commission, err := flagSetUtils.GetUint8Commission(flagSet)
	if err != nil {
		return err
	}

	status, err := razorUtils.ParseBool(statusString)
	if err != nil {
		log.Error("Error in parsing status to boolean")
		return err
	}

	client := razorUtils.ConnectToClient(config.Provider)

	stakerId, err := razorUtils.GetStakerId(client, address)
	if err != nil {
		log.Error("Error in fetching staker id")
		return err
	}

	stakerInfo, err := razorUtils.GetStaker(client, address, stakerId)
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

	if stakerInfo.AcceptDelegation != status {
		log.Infof("Setting delegation acceptance of Staker %d to %t", stakerId, status)
		txnOpts.MethodName = "setDelegationAcceptance"
		txnOpts.Parameters = []interface{}{status}
		setDelegationAcceptanceTxnOpts := razorUtils.GetTxnOpts(txnOpts)
		delegationAcceptanceTxn, err := stakeManagerUtils.SetDelegationAcceptance(client, setDelegationAcceptanceTxnOpts, status)
		if err != nil {
			log.Error("Error in setting delegation acceptance")
			return err
		}
		log.Infof("Transaction hash: %s", transactionUtils.Hash(delegationAcceptanceTxn))
		razorUtils.WaitForBlockCompletion(client, transactionUtils.Hash(delegationAcceptanceTxn).String())
	}

	// Fetch updated stakerInfo
	stakerInfo, err = razorUtils.GetUpdatedStaker(client, address, stakerId)
	if err != nil {
		log.Error("Error in fetching staker info")
		return err
	}

	if commission != 0 && stakerInfo.AcceptDelegation {
		// Call SetCommission if the commission value is provided and the staker hasn't already set commission
		if stakerInfo.Commission == 0 {
			txnOpts.MethodName = "setCommission"
			txnOpts.Parameters = []interface{}{commission}
			setCommissionTxnOpts := razorUtils.GetTxnOpts(txnOpts)
			err = cmdUtils.SetCommission(client, stakerId, setCommissionTxnOpts, commission, razorUtils, stakeManagerUtils, transactionUtils)
			if err != nil {
				return err
			}
		}

		// Call DecreaseCommission if the commission value is provided and the staker has already set commission
		if stakerInfo.Commission > 0 && stakerInfo.Commission > commission {
			txnOpts.MethodName = "decreaseCommission"
			txnOpts.Parameters = []interface{}{commission}
			decreaseCommissionTxnOpts := razorUtils.GetTxnOpts(txnOpts)
			err = cmdUtils.DecreaseCommission(client, stakerId, decreaseCommissionTxnOpts, commission, razorUtils, stakeManagerUtils, transactionUtils, cmdUtils)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func SetCommission(client *ethclient.Client, stakerId uint32, setCommissionTxnOpts *bind.TransactOpts, commission uint8, razorUtils utilsInterface, stakeManagerUtils stakeManagerInterface, transactionUtils transactionInterface) error {
	log.Infof("Setting the commission value of Staker %d to %d%%", stakerId, commission)
	txn, err := stakeManagerUtils.SetCommission(client, setCommissionTxnOpts, commission)
	if err != nil {
		log.Error("Error in setting commission")
		return err
	}
	log.Infof("Transaction hash: %s", transactionUtils.Hash(txn))
	razorUtils.WaitForBlockCompletion(client, transactionUtils.Hash(txn).String())
	return nil
}

func DecreaseCommission(client *ethclient.Client, stakerId uint32, decreaseCommissionTxnOpts *bind.TransactOpts, commission uint8, razorUtils utilsInterface, stakeManagerUtils stakeManagerInterface, transactionUtils transactionInterface, cmdUtils utilsCmdInterface) error {
	log.Infof("Decreasing the commission value of Staker %d to %d%%", stakerId, commission)
	if cmdUtils.DecreaseCommissionPrompt() {
		log.Info("Sending DecreaseCommission transaction...")
		decreaseCommissionTxn, err := stakeManagerUtils.DecreaseCommission(client, decreaseCommissionTxnOpts, commission)
		if err != nil {
			log.Error("Error in decreasing commission")
			return err
		}
		log.Infof("Transaction hash: %s", transactionUtils.Hash(decreaseCommissionTxn))
		razorUtils.WaitForBlockCompletion(client, transactionUtils.Hash(decreaseCommissionTxn).String())
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
