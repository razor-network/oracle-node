//Package cmd provides all functions related to command line
package cmd

import (
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/rpc"
	"razor/utils"

	"github.com/ethereum/go-ethereum/common"
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
	config, rpcParameters, account, err := InitializeCommandDependencies(flagSet)
	utils.CheckError("Error in initialising command dependencies: ", err)

	statusString, err := flagSetUtils.GetStringStatus(flagSet)
	utils.CheckError("Error in getting status: ", err)

	status, err := stringUtils.ParseBool(statusString)
	utils.CheckError("Error in parsing status to boolean: ", err)

	stakerId, err := razorUtils.GetStakerId(rpcParameters, account.Address)
	utils.CheckError("StakerId error: ", err)

	commission, err := flagSetUtils.GetUint8Commission(flagSet)
	utils.CheckError("Error in fetching commission: ", err)

	delegationInput := types.SetDelegationInput{
		Status:       status,
		StatusString: statusString,
		StakerId:     stakerId,
		Commission:   commission,
		Account:      account,
	}

	txn, err := cmdUtils.SetDelegation(rpcParameters, config, delegationInput)
	utils.CheckError("SetDelegation error: ", err)
	if txn != core.NilHash {
		err = razorUtils.WaitForBlockCompletion(rpcParameters, txn.Hex())
		utils.CheckError("Error in WaitForBlockCompletion for setDelegation: ", err)
	}
}

//This function allows the staker to start accepting/rejecting delegation requests
func (*UtilsStruct) SetDelegation(rpcParameters rpc.RPCParameters, config types.Configurations, delegationInput types.SetDelegationInput) (common.Hash, error) {
	stakerInfo, err := razorUtils.GetStaker(rpcParameters, delegationInput.StakerId)
	if err != nil {
		return core.NilHash, err
	}
	log.Debugf("SetDelegation: Staker Info: %+v", stakerInfo)
	if delegationInput.Commission != 0 {
		updateCommissionInput := types.UpdateCommissionInput{
			StakerId:   delegationInput.StakerId,
			Commission: delegationInput.Commission,
			Account:    delegationInput.Account,
		}
		err = cmdUtils.UpdateCommission(rpcParameters, config, updateCommissionInput)
		if err != nil {
			return core.NilHash, err
		}
	}

	txnOpts := types.TransactionOptions{
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.StakeManagerAddress,
		ABI:             bindings.StakeManagerMetaData.ABI,
		MethodName:      "setDelegationAcceptance",
		Parameters:      []interface{}{delegationInput.Status},
		Account:         delegationInput.Account,
	}

	if stakerInfo.AcceptDelegation == delegationInput.Status {
		log.Infof("Delegation status already set to %t", delegationInput.Status)
		return core.NilHash, nil
	}
	log.Infof("Setting delegation acceptance of Staker %d to %t", delegationInput.StakerId, delegationInput.Status)
	setDelegationAcceptanceTxnOpts, err := razorUtils.GetTxnOpts(rpcParameters, txnOpts)
	if err != nil {
		return core.NilHash, err
	}

	client, err := rpcParameters.RPCManager.GetBestRPCClient()
	if err != nil {
		return core.NilHash, err
	}

	log.Debug("Executing SetDelegationAcceptance transaction with status ", delegationInput.Status)
	delegationAcceptanceTxn, err := stakeManagerUtils.SetDelegationAcceptance(client, setDelegationAcceptanceTxnOpts, delegationInput.Status)
	if err != nil {
		log.Error("Error in setting delegation acceptance")
		return core.NilHash, err
	}
	delegationAcceptanceTxnHash := transactionUtils.Hash(delegationAcceptanceTxn)
	log.Info("Txn Hash: ", delegationAcceptanceTxnHash.Hex())
	return delegationAcceptanceTxnHash, nil
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
