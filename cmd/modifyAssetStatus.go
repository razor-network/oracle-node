package cmd

import (
	"razor/core"
	"razor/core/types"
	"razor/logger"
	"razor/pkg/bindings"
	"razor/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var modifyAssetStatusCmd = &cobra.Command{
	Use:   "modifyAssetStatus",
	Short: "[ADMIN ONLY]modify the active status of an asset",
	Long: `modifyAssetStatus can be used by admins to change the status of an asset
Example:	
  ./razor modifyAssetStatus --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --assetId 1 --status true --logLevel modifyLogs`,
	Run: initialiseModifyAssetStatus,
}

func initialiseModifyAssetStatus(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteModifyAssetStatus(cmd.Flags())
}

func (*UtilsStruct) ExecuteModifyAssetStatus(flagSet *pflag.FlagSet) {
	cmdUtils.AssignLogFile(flagSet)
	address, err := flagSetUtils.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)

	logger.Address = address

	config, err := cmdUtils.GetConfigData()
	utils.CheckError("Error in fetching config data: ", err)

	assetId, err := flagSetUtils.GetUint16AssetId(flagSet)
	utils.CheckError("Error in getting assetId: ", err)

	statusString, err := flagSetUtils.GetStringStatus(flagSet)
	utils.CheckError("Error in getting status: ", err)

	status, err := stringUtils.ParseBool(statusString)
	utils.CheckError("Error in parsing status: ", err)

	password := razorUtils.AssignPassword(flagSet)

	client := razorUtils.ConnectToClient(config.Provider)

	modifyAssetInput := types.ModifyAssetInput{
		Address:  address,
		Password: password,
		Status:   status,
		AssetId:  assetId,
	}

	txn, err := cmdUtils.ModifyAssetStatus(client, config, modifyAssetInput)
	utils.CheckError("Error in changing asset active status: ", err)
	if txn != core.NilHash {
		razorUtils.WaitForBlockCompletion(client, txn.String())
	}
}

func (*UtilsStruct) CheckCurrentStatus(client *ethclient.Client, assetId uint16) (bool, error) {
	callOpts := razorUtils.GetOptions()
	return assetManagerUtils.GetActiveStatus(client, &callOpts, assetId)
}

func (*UtilsStruct) ModifyAssetStatus(client *ethclient.Client, config types.Configurations, modifyAssetInput types.ModifyAssetInput) (common.Hash, error) {
	currentStatus, err := cmdUtils.CheckCurrentStatus(client, modifyAssetInput.AssetId)
	if err != nil {
		log.Error("Error in fetching active status")
		return core.NilHash, err
	}
	if currentStatus == modifyAssetInput.Status {
		log.Errorf("Asset %d has the active status already set to %t", modifyAssetInput.AssetId, modifyAssetInput.Status)
		return core.NilHash, nil
	}
	_, err = cmdUtils.WaitForAppropriateState(client, "modify asset status", 4)
	if err != nil {
		return core.NilHash, err
	}

	txnArgs := types.TransactionOptions{
		Client:          client,
		Password:        modifyAssetInput.Password,
		AccountAddress:  modifyAssetInput.Address,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.AssetManagerAddress,
		MethodName:      "setCollectionStatus",
		Parameters:      []interface{}{modifyAssetInput.Status, modifyAssetInput.AssetId},
		ABI:             bindings.AssetManagerABI,
	}

	txnOpts := razorUtils.GetTxnOpts(txnArgs)
	log.Infof("Changing active status of asset: %d from %t to %t", modifyAssetInput.AssetId, !modifyAssetInput.Status, modifyAssetInput.Status)
	txn, err := assetManagerUtils.SetCollectionStatus(client, txnOpts, modifyAssetInput.Status, modifyAssetInput.AssetId)
	if err != nil {
		return core.NilHash, err
	}
	log.Info("Txn Hash: ", transactionUtils.Hash(txn).String())
	return transactionUtils.Hash(txn), nil
}

func init() {
	rootCmd.AddCommand(modifyAssetStatusCmd)
	var (
		Address string
		AssetId uint16
		Status  string
		LogFile string
	)

	modifyAssetStatusCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the user")
	modifyAssetStatusCmd.Flags().Uint16VarP(&AssetId, "assetId", "", 0, "assetId of the asset")
	modifyAssetStatusCmd.Flags().StringVarP(&Status, "status", "", "true", "active status of the asset")
	modifyAssetStatusCmd.Flags().StringVarP(&LogFile, "logFile", "", "", "name of log file")

	addressErr := modifyAssetStatusCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addressErr)
	assetIdErr := modifyAssetStatusCmd.MarkFlagRequired("assetId")
	utils.CheckError("Asset Id error: ", assetIdErr)
	statusErr := modifyAssetStatusCmd.MarkFlagRequired("status")
	utils.CheckError("Status error: ", statusErr)
}
