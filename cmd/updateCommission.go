package cmd

import (
	"errors"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"

	"github.com/spf13/cobra"
)

// updateCommissionCmd represents the updateCommission command
var updateCommissionCmd = &cobra.Command{
	Use:   "updateCommission",
	Short: "updateCommission allows a staker to add/update the commission value",
	Long: `Using updateCommission stakers can add or update the commission charged by them
Example:
  ./razor updateCommission --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --commission 10`,
	Run: initialiseUpdateCommission,
}

func initialiseUpdateCommission(cmd *cobra.Command, args []string) {
	cmdUtilsMockery.ExecuteUpdateCommission(cmd.Flags())
}

func (*UtilsStructMockery) ExecuteUpdateCommission(flagSet *pflag.FlagSet) {
	config, err := cmdUtilsMockery.GetConfigData()
	utils.CheckError("Error in getting config", err)

	client := razorUtilsMockery.ConnectToClient(config.Provider)
	password := razorUtilsMockery.AssignPassword(flagSet)
	address, err := flagSetUtilsMockery.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address", err)

	commission, err := flagSetUtilsMockery.GetUint8Commission(flagSet)
	utils.CheckError("Error in getting commission", err)

	stakerId, err := razorUtilsMockery.GetStakerId(client, address)
	utils.CheckError("Error in getting stakerId", err)

	updateCommissionInput := types.UpdateCommissionInput{
		Address:    address,
		Password:   password,
		StakerId:   stakerId,
		Commission: commission,
	}

	err = cmdUtilsMockery.UpdateCommission(config, client, updateCommissionInput)
	utils.CheckError("SetDelegation error: ", err)
}

func (*UtilsStructMockery) UpdateCommission(config types.Configurations, client *ethclient.Client, updateCommissionInput types.UpdateCommissionInput) error {

	stakerInfo, err := razorUtilsMockery.GetStaker(client, updateCommissionInput.Address, updateCommissionInput.StakerId)
	if err != nil {
		log.Error("Error in fetching staker info")
		return err
	}

	maxCommission, err := razorUtilsMockery.GetMaxCommission(client)
	if err != nil {
		return err
	}

	if updateCommissionInput.Commission == 0 || updateCommissionInput.Commission > maxCommission {
		return errors.New("commission out of range")
	}

	epochLimitForUpdateCommission, err := razorUtilsMockery.GetEpochLimitForUpdateCommission(client)
	if err != nil {
		return err
	}

	epoch, err := razorUtilsMockery.GetEpoch(client)
	if err != nil {
		return err
	}

	if (stakerInfo.EpochCommissionLastUpdated + uint32(epochLimitForUpdateCommission)) >= epoch {
		return errors.New("invalid epoch for update")
	}
	txnOpts := types.TransactionOptions{
		Client:          client,
		Password:        updateCommissionInput.Password,
		AccountAddress:  updateCommissionInput.Address,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.StakeManagerAddress,
		ABI:             bindings.StakeManagerABI,
		MethodName:      "updateCommission",
		Parameters:      []interface{}{updateCommissionInput.Commission},
	}

	updateCommissionTxnOpts := razorUtilsMockery.GetTxnOpts(txnOpts)
	log.Infof("Setting the commission value of Staker %d to %d%%", updateCommissionInput.StakerId, updateCommissionInput.Commission)
	txn, err := stakeManagerUtilsMockery.UpdateCommission(client, updateCommissionTxnOpts, updateCommissionInput.Commission)
	if err != nil {
		log.Error("Error in setting commission")
		return err
	}
	txnHash := transactionUtilsMockery.Hash(txn)
	log.Infof("Transaction hash: %s", txnHash)
	razorUtilsMockery.WaitForBlockCompletion(client, txnHash.String())
	return nil
}

func init() {
	razorUtilsMockery = UtilsMockery{}
	stakeManagerUtilsMockery = StakeManagerUtilsMockery{}
	transactionUtilsMockery = TransactionUtilsMockery{}
	flagSetUtilsMockery = FLagSetUtilsMockery{}
	cmdUtilsMockery = &UtilsStructMockery{}

	var (
		Address    string
		Commission uint8
		Password   string
	)

	rootCmd.AddCommand(updateCommissionCmd)

	updateCommissionCmd.Flags().StringVarP(&Address, "address", "a", "", "your account address")
	updateCommissionCmd.Flags().Uint8VarP(&Commission, "commission", "c", 0, "commission")
	updateCommissionCmd.Flags().StringVarP(&Password, "password", "", "", "password path to protect the keystore")

	addrErr := updateCommissionCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
}
