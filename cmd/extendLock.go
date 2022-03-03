package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"razor/core"
	"razor/core/types"
	"razor/logger"
	"razor/pkg/bindings"
	"razor/utils"

	"github.com/spf13/cobra"
)

var extendLockCmd = &cobra.Command{
	Use:   "extendLock",
	Short: "extendLock can be used to reset the lock once the withdraw lock period is over",
	Long: `If the withdrawal period is over, then the lock must be reset otherwise the user cannot unstake. This can be done by extendLock command.

Example:
  ./razor extendLock --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c 
`,
	Run: initialiseExtendLock,
}

func initialiseExtendLock(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteExtendLock(cmd.Flags())
}

func (*UtilsStruct) ExecuteExtendLock(flagSet *pflag.FlagSet) {
	address, err := flagSetUtils.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)

	logger.Address = address

	config, err := cmdUtils.GetConfigData()
	utils.CheckError("Error in getting config data: ", err)

	password := razorUtils.AssignPassword(flagSet)

	client := razorUtils.ConnectToClient(config.Provider)

	stakerId, err := razorUtils.AssignStakerId(flagSet, client, address)
	utils.CheckError("Error in getting stakerId: ", err)

	extendLockInput := types.ExtendLockInput{
		Address:  address,
		Password: password,
		StakerId: stakerId,
	}
	txn, err := cmdUtils.ExtendLock(client, config, extendLockInput)
	utils.CheckError("Error in extending lock: ", err)
	razorUtils.WaitForBlockCompletion(client, txn.String())
}

func (*UtilsStruct) ExtendLock(client *ethclient.Client, config types.Configurations, extendLockInput types.ExtendLockInput) (common.Hash, error) {
	txnOpts := razorUtils.GetTxnOpts(types.TransactionOptions{
		Client:          client,
		Password:        extendLockInput.Password,
		AccountAddress:  extendLockInput.Address,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.StakeManagerAddress,
		MethodName:      "extendLock",
		Parameters:      []interface{}{extendLockInput.StakerId},
		ABI:             bindings.StakeManagerABI,
	})

	log.Info("Extending lock...")
	txn, err := stakeManagerUtils.ExtendLock(client, txnOpts, extendLockInput.StakerId)
	if err != nil {
		return core.NilHash, err
	}
	log.Info("Txn Hash: ", transactionUtils.Hash(txn))
	return transactionUtils.Hash(txn), nil
}

func init() {
	razorUtils = Utils{}
	stakeManagerUtils = StakeManagerUtils{}
	transactionUtils = TransactionUtils{}
	flagSetUtils = FLagSetUtils{}
	cmdUtils = &UtilsStruct{}
	InitializeUtils()

	rootCmd.AddCommand(extendLockCmd)

	var (
		Address  string
		Password string
		StakerId uint32
	)

	extendLockCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the user")
	extendLockCmd.Flags().StringVarP(&Password, "password", "", "", "password path of the user to protect the keystore")
	extendLockCmd.Flags().Uint32VarP(&StakerId, "stakerId", "", 0, "staker id")

	addrErr := extendLockCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
}
