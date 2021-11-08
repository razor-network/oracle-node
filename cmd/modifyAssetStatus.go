package cmd

import (
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var modifyAssetStatusCmd = &cobra.Command{
	Use:   "modifyAssetStatus",
	Short: "modify the active status of an asset",
	Long: `modifyAssetStatus can be used by admins to change the status of an asset
Example:	
  ./razor modifyAssetStatus --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --assetId 1 --status true`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		utils.CheckError("Error in fetching config data: ", err)

		txn, err := ModifyAssetStatus(cmd.Flags(), config, razorUtils, assetManagerUtils, cmdUtils, flagSetUtils, transactionUtils)
		utils.CheckError("Error in changing asset active status: ", err)
		if txn != core.NilHash {
			utils.WaitForBlockCompletion(utils.ConnectToClient(config.Provider), txn.String())
		}
	},
}

func CheckCurrentStatus(client *ethclient.Client, address string, assetId uint8, razorUtils utilsInterface, assetManagerUtils assetManagerInterface) (bool, error) {
	callOpts := razorUtils.GetOptions(false, address, "")
	return assetManagerUtils.GetActiveStatus(client, &callOpts, assetId)
}

func ModifyAssetStatus(flagSet *pflag.FlagSet, config types.Configurations, razorUtils utilsInterface, assetManagerUtils assetManagerInterface, cmdUtils utilsCmdInterface, flagSetUtils flagSetInterface, transactionUtils transactionInterface) (common.Hash, error) {
	address, err := flagSetUtils.GetStringAddress(flagSet)
	if err != nil {
		return core.NilHash, err
	}
	assetId, err := flagSetUtils.GetUint8AssetId(flagSet)
	if err != nil {
		return core.NilHash, err
	}
	statusString, err := flagSetUtils.GetStringStatus(flagSet)
	if err != nil {
		return core.NilHash, err
	}
	status, err := razorUtils.ParseBool(statusString)
	if err != nil {
		log.Error("Error in parsing status to boolean")
		return core.NilHash, err
	}

	password := razorUtils.PasswordPrompt()

	client := razorUtils.ConnectToClient(config.Provider)

	currentStatus, err := cmdUtils.CheckCurrentStatus(client, address, assetId, razorUtils, assetManagerUtils)
	if err != nil {
		log.Error("Error in fetching active status")
		return core.NilHash, err
	}
	if currentStatus == status {
		log.Errorf("Asset %d has the active status already set to %t", assetId, status)
		return core.NilHash, nil
	}
	_, err = razorUtils.WaitForAppropriateState(client, address, "modify asset status", 4)
	if err != nil {
		return core.NilHash, err
	}

	txnArgs := types.TransactionOptions{
		Client:          client,
		Password:        password,
		AccountAddress:  address,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.AssetManagerAddress,
		MethodName:      "setCollectionStatus",
		Parameters:      []interface{}{status, assetId},
		ABI:             bindings.AssetManagerABI,
	}

	txnOpts := razorUtils.GetTxnOpts(txnArgs)
	log.Infof("Changing active status of asset: %d from %t to %t", assetId, !status, status)
	txn, err := assetManagerUtils.SetCollectionStatus(client, txnOpts, status, assetId)
	if err != nil {
		return core.NilHash, err
	}
	log.Info("Txn Hash: ", transactionUtils.Hash(txn).String())
	return transactionUtils.Hash(txn), nil
}

func init() {

	razorUtils = Utils{}
	cmdUtils = UtilsCmd{}
	flagSetUtils = FlagSetUtils{}
	assetManagerUtils = AssetManagerUtils{}
	transactionUtils = TransactionUtils{}

	rootCmd.AddCommand(modifyAssetStatusCmd)

	var (
		Address string
		AssetId uint8
		Status  string
	)

	modifyAssetStatusCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the user")
	modifyAssetStatusCmd.Flags().Uint8VarP(&AssetId, "assetId", "", 0, "assetId of the asset")
	modifyAssetStatusCmd.Flags().StringVarP(&Status, "status", "", "true", "active status of the asset")

	addressErr := modifyAssetStatusCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addressErr)
	assetIdErr := modifyAssetStatusCmd.MarkFlagRequired("assetId")
	utils.CheckError("Asset Id error: ", assetIdErr)
	statusErr := modifyAssetStatusCmd.MarkFlagRequired("status")
	utils.CheckError("Status error: ", statusErr)
}
