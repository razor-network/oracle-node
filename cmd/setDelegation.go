package cmd

import (
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/manifoldco/promptui"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// setDelegationCmd represents the setDelegation command
var setDelegationCmd = &cobra.Command{
	Use:   "setDelegation",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		utils.CheckError("Error in getting config: ", err)

		password := utils.PasswordPrompt()
		address, _ := cmd.Flags().GetString("address")
		status, _ := cmd.Flags().GetBool("status")
		commission, _ := cmd.Flags().GetString("commission")

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
			GasMultiplier:  config.GasMultiplier,
		})

		if stakerInfo.AcceptDelegation != status {
			log.Infof("Setting delegation acceptance of Staker %s to %t", stakerId, status)
			delegationTxn, err := stakeManager.SetDelegationAcceptance(txnOpts, status)
			utils.CheckError("Error in setting delegation acceptance: ", err)
			log.Info("Sending SetDelegationAcceptance transaction...")
			log.Infof("Transaction hash: %s", delegationTxn.Hash())
			utils.WaitForBlockCompletion(client, delegationTxn.Hash().String())
		}

		_commission, ok := new(big.Int).SetString(commission, 10)
		commissionAmountInWei := big.NewInt(1).Mul(_commission, big.NewInt(1e18))
		if !ok {
			log.Fatal("Set string: error")
		}

		// Fetch updated stakerInfo
		stakerInfo, err = utils.GetStaker(client, address, stakerId)
		utils.CheckError("Error in fetching staker info: ", err)
		if commission != "0" && stakerInfo.AcceptDelegation {
			// Call SetCommission if the commission value is provided and the staker hasn't already set commission
			if stakerInfo.Commission.Cmp(big.NewInt(0)) == 0 {
				SetCommission(client, stakeManager, stakerId, txnOpts, commissionAmountInWei)
			}

			// Call DecreaseCommission if the commission value is provided and the staker has already set commission
			if stakerInfo.Commission.Cmp(big.NewInt(0)) > 0 && stakerInfo.Commission.Cmp(commissionAmountInWei) > 0 {
				DecreaseCommission(client, stakeManager, stakerId, txnOpts, commissionAmountInWei)
			}
		}

	},
}

func SetCommission(client *ethclient.Client, stakeManager *bindings.StakeManager, stakerId *big.Int, txnOpts *bind.TransactOpts, commissionAmountInWei *big.Int) {
	log.Infof("Setting the commission value of Staker %s to %s", stakerId, commissionAmountInWei)
	commissionTxn, err := stakeManager.SetCommission(txnOpts, commissionAmountInWei)
	utils.CheckError("Error in setting commission: ", err)
	log.Info("Sending SetCommission transaction...")
	log.Infof("Transaction hash: %s", commissionTxn.Hash())
	utils.WaitForBlockCompletion(client, commissionTxn.Hash().String())
}

func DecreaseCommission(client *ethclient.Client, stakeManager *bindings.StakeManager, stakerId *big.Int, txnOpts *bind.TransactOpts, commissionAmountInWei *big.Int) {
	log.Infof("Decreasing the commission value of Staker %s to %s", stakerId, commissionAmountInWei)
	prompt := promptui.Prompt{
		Label:     "Decrease Commission? Once decreased, your commission cannot be increased.",
		IsConfirm: true,
	}
	result, err := prompt.Run()
	utils.CheckError(result, err)
	if strings.ToLower(result) == "yes" || strings.ToLower(result) == "y" {
		decreaseCommissionTxn, err := stakeManager.DecreaseCommission(txnOpts, commissionAmountInWei)
		utils.CheckError("Error in decreasing commission: ", err)
		log.Info("Sending DecreaseCommission transaction...")
		log.Infof("Transaction hash: %s", decreaseCommissionTxn.Hash())
		utils.WaitForBlockCompletion(client, decreaseCommissionTxn.Hash().String())
	}
}

func init() {
	rootCmd.AddCommand(setDelegationCmd)

	var (
		Status     bool
		Address    string
		Commission string
	)
	setDelegationCmd.Flags().BoolVarP(&Status, "status", "s", true, "true for accepting delegation and false for not accepting")
	setDelegationCmd.Flags().StringVarP(&Address, "address", "", "", "your account address")
	setDelegationCmd.Flags().StringVarP(&Commission, "commission", "c", "0", "commission")

	addrErr := setDelegationCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
}
