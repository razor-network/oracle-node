package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"

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

func initialiseSetDelegation(cmd *cobra.Command, args []string) {
	cmdUtilsMockery.ExecuteSetDelegation(cmd.Flags())
}

func (*UtilsStructMockery) ExecuteSetDelegation(flagSet *pflag.FlagSet) {

	config, err := cmdUtilsMockery.GetConfigData()
	utils.CheckError("Error in getting config: ", err)

	password := razorUtilsMockery.AssignPassword(flagSet)
	address, err := flagSetUtilsMockery.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)

	statusString, err := flagSetUtilsMockery.GetStringStatus(flagSet)
	utils.CheckError("Error in getting status: ", err)

	status, err := razorUtilsMockery.ParseBool(statusString)
	utils.CheckError("Error in parsing status to boolean: ", err)

	client := razorUtilsMockery.ConnectToClient(config.Provider)

	stakerId, err := flagSetUtilsMockery.GetUint32StakerId(flagSet)
	utils.CheckError("Error in fetching stakerId: ", err)

	delegationInput := types.SetDelegationInput{
		Address:      address,
		Password:     password,
		Status:       status,
		StatusString: statusString,
		StakerId:     stakerId,
	}

	txn, err := cmdUtilsMockery.SetDelegation(client, config, delegationInput)
	utils.CheckError("SetDelegation error: ", err)
	razorUtilsMockery.WaitForBlockCompletion(client, txn.String())
}

func (*UtilsStructMockery) SetDelegation(client *ethclient.Client, config types.Configurations, delegationInput types.SetDelegationInput) (common.Hash, error) {

	stakerInfo, err := razorUtilsMockery.GetStaker(client, delegationInput.Address, delegationInput.StakerId)
	if err != nil {
		return core.NilHash, nil
	}

	txnOpts := types.TransactionOptions{
		Client:          client,
		Password:        delegationInput.Password,
		AccountAddress:  delegationInput.Address,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.StakeManagerAddress,
		ABI:             bindings.StakeManagerABI,
		MethodName:      "setDelegationAcceptance",
		Parameters:      []interface{}{delegationInput.Status},
	}

	if stakerInfo.AcceptDelegation == delegationInput.Status {
		log.Infof("Delegation status already set to %t", delegationInput.Status)
		return core.NilHash, nil
	}
	log.Infof("Setting delegation acceptance of Staker %d to %t", delegationInput.StakerId, delegationInput.Status)
	setDelegationAcceptanceTxnOpts := razorUtilsMockery.GetTxnOpts(txnOpts)
	delegationAcceptanceTxn, err := stakeManagerUtilsMockery.SetDelegationAcceptance(client, setDelegationAcceptanceTxnOpts, delegationInput.Status)
	if err != nil {
		log.Error("Error in setting delegation acceptance")
		return core.NilHash, err
	}
	log.Infof("Transaction hash: %s", transactionUtilsMockery.Hash(delegationAcceptanceTxn))
	return transactionUtilsMockery.Hash(delegationAcceptanceTxn), nil
}

func init() {

	razorUtilsMockery = &UtilsMockery{}
	stakeManagerUtilsMockery = StakeManagerUtilsMockery{}
	transactionUtilsMockery = TransactionUtilsMockery{}
	flagSetUtilsMockery = FLagSetUtilsMockery{}
	cmdUtilsMockery = &UtilsStructMockery{}

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
