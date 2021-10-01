package cmd

import (
	"github.com/spf13/pflag"
	"razor/core"
	"razor/core/types"
	"razor/utils"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
		err := SetDelegation(cmd.Flags(), razorUtils, stakeManagerUtils, transactionUtils, flagSetUtils)
		utils.CheckError("SetDelegation error: ", err)
	},
}

func SetDelegation(flagSet *pflag.FlagSet, razorUtils utilsInterface, stakeManagerUtils stakeManagerInterface, transactionUtils transactionInterface, flagSetUtils flagSetInterface) error {

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

	status, err := strconv.ParseBool(statusString)
	utils.CheckError("Error in parsing status to boolean: ", err)

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

	txnOpts := razorUtils.GetTxnOpts(types.TransactionOptions{
		Client:         client,
		Password:       password,
		AccountAddress: address,
		ChainId:        core.ChainId,
		Config:         config,
	})

	if stakerInfo.AcceptDelegation != status {
		log.Infof("Setting delegation acceptance of Staker %d to %t", stakerId, status)
		delegationAcceptanceTxn, err := stakeManagerUtils.SetDelegationAcceptance(client, txnOpts, status)
		if err != nil {
			log.Error("Error in setting delegation acceptance")
			return err
		}
		log.Infof("Transaction hash: %s", transactionUtils.Hash(delegationAcceptanceTxn))
		razorUtils.WaitForBlockCompletion(client, transactionUtils.Hash(delegationAcceptanceTxn).String())
	}

	// Fetch updated stakerInfo
	stakerInfo, err = razorUtils.GetStaker(client, address, stakerId)
	if err != nil {
		log.Error("Error in fetching staker info")
		return err
	}
	if commission != 0 && stakerInfo.AcceptDelegation {
		// Call SetCommission if the commission value is provided and the staker hasn't already set commission
		if stakerInfo.Commission == 0 {
			err = cmdUtils.SetCommission(client, stakerId, txnOpts, commission, razorUtils, stakeManagerUtils, transactionUtils)
			if err != nil {
				return err
			}
		}

		// Call DecreaseCommission if the commission value is provided and the staker has already set commission
		if stakerInfo.Commission > 0 && stakerInfo.Commission > commission {
			err = cmdUtils.DecreaseCommission(client, stakerId, txnOpts, commission, razorUtils, stakeManagerUtils, transactionUtils)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func SetCommission(client *ethclient.Client, stakerId uint32, txnOpts *bind.TransactOpts, commission uint8, razorUtils utilsInterface, stakeManagerUtils stakeManagerInterface, transactionUtils transactionInterface) error {
	log.Infof("Setting the commission value of Staker %d to %d%%", stakerId, commission)
	commissionTxn, err := stakeManagerUtils.SetCommission(client, txnOpts, commission)
	if err != nil {
		log.Error("Error in setting commission")
		return err
	}
	log.Infof("Transaction hash: %s", transactionUtils.Hash(commissionTxn))
	razorUtils.WaitForBlockCompletion(client, transactionUtils.Hash(commissionTxn).String())
	return nil
}

func DecreaseCommission(client *ethclient.Client, stakerId uint32, txnOpts *bind.TransactOpts, commission uint8, razorUtils utilsInterface, stakeManagerUtils stakeManagerInterface, transactionUtils transactionInterface) error {
	log.Infof("Decreasing the commission value of Staker %d to %d%%", stakerId, commission)
	prompt := promptui.Prompt{
		Label:     "Decrease Commission? Once decreased, your commission cannot be increased.",
		IsConfirm: true,
	}
	result, err := prompt.Run()
	utils.CheckError(result, err)
	if strings.ToLower(result) == "y" {
		log.Info("Sending DecreaseCommission transaction...")
		decreaseCommissionTxn, err := stakeManagerUtils.DecreaseCommission(client, txnOpts, commission)
		if err != nil {
			log.Error("Error in decreasing commission")
			return err
		}
		log.Infof("Transaction hash: %s", transactionUtils.Hash(decreaseCommissionTxn))
		razorUtils.WaitForBlockCompletion(client, transactionUtils.Hash(decreaseCommissionTxn).String())
	}
	return nil
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
