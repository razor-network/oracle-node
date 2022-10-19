//Package cmd provides all functions related to command line
package cmd

import (
	"razor/core"
	"razor/core/types"
	"razor/logger"
	"razor/pkg/bindings"
	"razor/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"

	"github.com/spf13/cobra"
)

var setDelegationCmd = &cobra.Command{
	Use:   "setDelegation",
	Short: "setDelegation allows a staker to start accepting/rejecting delegation requests",
	Long: `Using setDelegation, a staker can accept delegation from delegators and charge a commission from them.

Example:
  ./razor setDelegation --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --status true
`,
	Run: initialiseSetDelegation,
}

//This function initialises the ExecuteSetDelegation function
func initialiseSetDelegation(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteSetDelegation(cmd.Flags())
}

//This function sets the flags appropriately and executes the SetDelegation function
func (*UtilsStruct) ExecuteSetDelegation(flagSet *pflag.FlagSet) {
	config, err := cmdUtils.GetConfigData()
	utils.CheckError("Error in getting config: ", err)

	client := razorUtils.ConnectToClient(config.Provider)

	address, err := flagSetUtils.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)

	logger.SetLoggerParameters(client, address)
	razorUtils.AssignLogFile(flagSet)

	password := razorUtils.AssignPassword(flagSet)
	statusString, err := flagSetUtils.GetStringStatus(flagSet)
	utils.CheckError("Error in getting status: ", err)

	status, err := stringUtils.ParseBool(statusString)
	utils.CheckError("Error in parsing status to boolean: ", err)

	stakerId, err := razorUtils.GetStakerId(client, address)
	utils.CheckError("StakerId error: ", err)

	commission, err := flagSetUtils.GetUint8Commission(flagSet)
	utils.CheckError("Error in fetching commission: ", err)

	delegationInput := types.SetDelegationInput{
		Address:      address,
		Password:     password,
		Status:       status,
		StatusString: statusString,
		StakerId:     stakerId,
		Commission:   commission,
	}

	txn, err := cmdUtils.SetDelegation(client, config, delegationInput)
	utils.CheckError("SetDelegation error: ", err)
	if txn != core.NilHash {
		err = razorUtils.WaitForBlockCompletion(client, txn.String())
		utils.CheckError("Error in WaitForBlockCompletion for setDelegation: ", err)
	}
}

//This function allows the staker to start accepting/rejecting delegation requests
func (*UtilsStruct) SetDelegation(client *ethclient.Client, config types.Configurations, delegationInput types.SetDelegationInput) (common.Hash, error) {
	stakerInfo, err := razorUtils.GetStaker(client, delegationInput.StakerId)
	if err != nil {
		return core.NilHash, err
	}
	if delegationInput.Commission != 0 {
		err = cmdUtils.UpdateCommission(config, client, types.UpdateCommissionInput{
			StakerId:   delegationInput.StakerId,
			Address:    delegationInput.Address,
			Password:   delegationInput.Password,
			Commission: delegationInput.Commission,
		})
		if err != nil {
			return core.NilHash, err
		}
	}

	txnOpts := types.TransactionOptions{
		Client:          client,
		Password:        delegationInput.Password,
		AccountAddress:  delegationInput.Address,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.StakeManagerAddress,
		ABI:             bindings.StakeManagerMetaData.ABI,
		MethodName:      "setDelegationAcceptance",
		Parameters:      []interface{}{delegationInput.Status},
	}

	if stakerInfo.AcceptDelegation == delegationInput.Status {
		log.Infof("Delegation status already set to %t", delegationInput.Status)
		return core.NilHash, nil
	}
	log.Infof("Setting delegation acceptance of Staker %d to %t", delegationInput.StakerId, delegationInput.Status)
	setDelegationAcceptanceTxnOpts := razorUtils.GetTxnOpts(txnOpts)
	delegationAcceptanceTxn, err := stakeManagerUtils.SetDelegationAcceptance(client, setDelegationAcceptanceTxnOpts, delegationInput.Status)
	if err != nil {
		log.Error("Error in setting delegation acceptance")
		return core.NilHash, err
	}
	log.Infof("Transaction hash: %s", transactionUtils.Hash(delegationAcceptanceTxn))
	return transactionUtils.Hash(delegationAcceptanceTxn), nil
}

func init() {

	rootCmd.AddCommand(setDelegationCmd)

	var (
		Status     string
		Address    string
		Password   string
		Commission uint8
	)

	setDelegationCmd.Flags().StringVarP(&Status, "status", "s", "true", "true for accepting delegation and false for not accepting")
	setDelegationCmd.Flags().StringVarP(&Address, "address", "a", "", "your account address")
	setDelegationCmd.Flags().StringVarP(&Password, "password", "", "", "password path to protect the keystore")
	setDelegationCmd.Flags().Uint8VarP(&Commission, "commission", "c", 0, "commission")

	addrErr := setDelegationCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
}
