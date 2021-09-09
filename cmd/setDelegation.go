package cmd

import (
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/manifoldco/promptui"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var setDelegationCmd = &cobra.Command{
	Use:   "setDelegation",
	Short: "setDelegation allows a staker to start accepting/rejecting delegation requests",
	Long: `Using setDelegation, a staker can accept delegation from delegators and charge a commission from them.

Example:
  ./razor setDelegation --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --status true --commission 100
`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		utils.CheckError("Error in getting config: ", err)

		password := utils.AssignPassword(cmd.Flags())
		address, _ := cmd.Flags().GetString("address")
		statusString, _ := cmd.Flags().GetString("status")
		commission, _ := cmd.Flags().GetUint8("commission")

		status, err := strconv.ParseBool(statusString)
		utils.CheckError("Error in parsing status to boolean: ", err)

		client := utils.ConnectToClient(config.Provider)

		stakerId, err := utils.GetStakerId(client, address)
		utils.CheckError("Error in fetching staker id: ", err)

		stakerInfo, err := utils.GetStaker(client, address, stakerId)
		utils.CheckError("Error in fetching staker info: ", err)

		stakeManager := utils.GetStakeManager(client)
		txnOpts := utils.GetTxnOpts(types.TransactionOptions{
			Client:         client,
			Password:       password,
			AccountAddress: address,
			ChainId:        core.ChainId,
			Config:         config,
		})

		if stakerInfo.AcceptDelegation != status {
			log.Infof("Setting delegation acceptance of Staker %d to %t", stakerId, status)
			delegationTxn, err := stakeManager.SetDelegationAcceptance(txnOpts, status)
			utils.CheckError("Error in setting delegation acceptance: ", err)
			log.Info("Sending SetDelegationAcceptance transaction...")
			log.Infof("Transaction hash: %s", delegationTxn.Hash())
			utils.WaitForBlockCompletion(client, delegationTxn.Hash().String())
		}

		// Fetch updated stakerInfo
		stakerInfo, err = utils.GetStaker(client, address, stakerId)
		utils.CheckError("Error in fetching staker info: ", err)
		if commission != 0 && stakerInfo.AcceptDelegation {
			// Call SetCommission if the commission value is provided and the staker hasn't already set commission
			if stakerInfo.Commission == 0 {
				SetCommission(client, stakeManager, stakerId, txnOpts, commission)
			}

			// Call DecreaseCommission if the commission value is provided and the staker has already set commission
			if stakerInfo.Commission > 0 && stakerInfo.Commission > commission {
				DecreaseCommission(client, stakeManager, stakerId, txnOpts, commission)
			}
		}

	},
}

func SetCommission(client *ethclient.Client, stakeManager *bindings.StakeManager, stakerId uint32, txnOpts *bind.TransactOpts, commission uint8) {
	log.Infof("Setting the commission value of Staker %d to %d%%", stakerId, commission)
	commissionTxn, err := stakeManager.SetCommission(txnOpts, commission)
	utils.CheckError("Error in setting commission: ", err)
	log.Info("Sending SetCommission transaction...")
	log.Infof("Transaction hash: %s", commissionTxn.Hash())
	utils.WaitForBlockCompletion(client, commissionTxn.Hash().String())
}

func DecreaseCommission(client *ethclient.Client, stakeManager *bindings.StakeManager, stakerId uint32, txnOpts *bind.TransactOpts, commission uint8) {
	log.Infof("Decreasing the commission value of Staker %d to %d%%", stakerId, commission)
	prompt := promptui.Prompt{
		Label:     "Decrease Commission? Once decreased, your commission cannot be increased.",
		IsConfirm: true,
	}
	result, err := prompt.Run()
	utils.CheckError(result, err)
	if strings.ToLower(result) == "y" {
		decreaseCommissionTxn, err := stakeManager.DecreaseCommission(txnOpts, commission)
		utils.CheckError("Error in decreasing commission: ", err)
		log.Info("Sending DecreaseCommission transaction...")
		log.Infof("Transaction hash: %s", decreaseCommissionTxn.Hash())
		utils.WaitForBlockCompletion(client, decreaseCommissionTxn.Hash().String())
	}
}

func init() {
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
