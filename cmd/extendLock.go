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
	cmdUtilsMockery.ExecuteExtendLock(cmd.Flags())
}

func (*UtilsStructMockery) ExecuteExtendLock(flagSet *pflag.FlagSet) {
	config, err := cmdUtilsMockery.GetConfigData()
	utils.CheckError("Error in getting config data: ", err)

	password := razorUtilsMockery.AssignPassword(flagSet)
	address, err := flagSetUtilsMockery.GetStringAddress(flagSet)
	utils.CheckError("Error in getting config data: ", err)

	stakerId, err := flagSetUtilsMockery.GetUint32StakerId(flagSet)
	utils.CheckError("Error in getting config data: ", err)

	client := razorUtilsMockery.ConnectToClient(config.Provider)

	extendLockInput := types.ExtendLockInput{
		Address:  address,
		Password: password,
		StakerId: stakerId,
	}
	txn, err := cmdUtilsMockery.ExtendLock(client, config, extendLockInput)
	utils.CheckError("Error in extending lock: ", err)
	razorUtilsMockery.WaitForBlockCompletion(client, txn.String())
}

func (*UtilsStructMockery) ExtendLock(client *ethclient.Client, config types.Configurations, extendLockInput types.ExtendLockInput) (common.Hash, error) {
	txnOpts := razorUtilsMockery.GetTxnOpts(types.TransactionOptions{
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
	txn, err := stakeManagerUtilsMockery.ExtendLock(client, txnOpts, extendLockInput.StakerId)
	if err != nil {
		return core.NilHash, err
	}
	log.Info("Txn Hash: ", transactionUtilsMockery.Hash(txn))
	return transactionUtilsMockery.Hash(txn), nil
}

func init() {
	razorUtilsMockery = UtilsMockery{}
	stakeManagerUtilsMockery = StakeManagerUtilsMockery{}
	transactionUtilsMockery = TransactionUtilsMockery{}
	flagSetUtilsMockery = FLagSetUtilsMockery{}
	cmdUtilsMockery = &UtilsStructMockery{}
	utils.Options = &utils.OptionsStruct{}
	utils.UtilsInterface = &utils.UtilsStruct{}

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
