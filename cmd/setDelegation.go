//Package cmd provides all functions related to command line
package cmd

import (
	"razor/accounts"
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
	log.Debugf("ExecuteSetDelegation: Config: %+v", config)

	client := razorUtils.ConnectToClient(config.Provider)

	address, err := flagSetUtils.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)

	logger.SetLoggerParameters(client, address)

	log.Debug("Checking to assign log file...")
	fileUtils.AssignLogFile(flagSet, config)

	log.Debug("Getting password...")
	password := razorUtils.AssignPassword(flagSet)

	accountManager, err := razorUtils.AccountManagerForKeystore()
	utils.CheckError("Error in getting accounts manager for keystore: ", err)

	account := accounts.InitAccountStruct(address, password, accountManager)

	err = razorUtils.CheckPassword(account)
	utils.CheckError("Error in fetching private key from given password: ", err)

	statusString, err := flagSetUtils.GetStringStatus(flagSet)
	utils.CheckError("Error in getting status: ", err)

	status, err := stringUtils.ParseBool(statusString)
	utils.CheckError("Error in parsing status to boolean: ", err)

	stakerId, err := razorUtils.GetStakerId(client, address)
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

	txn, err := cmdUtils.SetDelegation(client, config, delegationInput)
	utils.CheckError("SetDelegation error: ", err)
	if txn != core.NilHash {
		err = razorUtils.WaitForBlockCompletion(client, txn.Hex())
		utils.CheckError("Error in WaitForBlockCompletion for setDelegation: ", err)
	}
}

//This function allows the staker to start accepting/rejecting delegation requests
func (*UtilsStruct) SetDelegation(client *ethclient.Client, config types.Configurations, delegationInput types.SetDelegationInput) (common.Hash, error) {
	stakerInfo, err := razorUtils.GetStaker(client, delegationInput.StakerId)
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
		err = cmdUtils.UpdateCommission(config, client, updateCommissionInput)
		if err != nil {
			return core.NilHash, err
		}
	}

	txnOpts := types.TransactionOptions{
		Client:          client,
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
	setDelegationAcceptanceTxnOpts := razorUtils.GetTxnOpts(txnOpts)
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
