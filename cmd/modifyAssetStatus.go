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
	Run: initialiseModifyAssetStatus,
}

func initialiseModifyAssetStatus(cmd *cobra.Command, args []string) {
	cmdUtilsMockery.ExecuteModifyAssetStatus(cmd.Flags())
}

func (*UtilsStructMockery) ExecuteModifyAssetStatus(flagSet *pflag.FlagSet) {
	config, err := cmdUtilsMockery.GetConfigData()
	utils.CheckError("Error in fetching config data: ", err)

	address, err := flagSetUtilsMockery.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)

	assetId, err := flagSetUtilsMockery.GetUint16AssetId(flagSet)
	utils.CheckError("Error in getting assetId: ", err)

	statusString, err := flagSetUtilsMockery.GetStringStatus(flagSet)
	utils.CheckError("Error in getting status: ", err)

	status, err := razorUtilsMockery.ParseBool(statusString)
	utils.CheckError("Error in parsing status: ", err)

	password := razorUtilsMockery.AssignPassword(flagSet)

	client := razorUtilsMockery.ConnectToClient(config.Provider)

	modifyAssetInput := types.ModifyAssetInput{
		Address:  address,
		Password: password,
		Status:   status,
		AssetId:  assetId,
	}

	txn, err := cmdUtilsMockery.ModifyAssetStatus(client, config, modifyAssetInput)
	utils.CheckError("Error in changing asset active status: ", err)
	if txn != core.NilHash {
		razorUtilsMockery.WaitForBlockCompletion(client, txn.String())
	}
}

func (*UtilsStructMockery) CheckCurrentStatus(client *ethclient.Client, assetId uint16) (bool, error) {
	callOpts := razorUtilsMockery.GetOptions()
	return assetManagerUtilsMockery.GetActiveStatus(client, &callOpts, assetId)
}

func (*UtilsStructMockery) ModifyAssetStatus(client *ethclient.Client, config types.Configurations, modifyAssetInput types.ModifyAssetInput) (common.Hash, error) {
	currentStatus, err := cmdUtilsMockery.CheckCurrentStatus(client, modifyAssetInput.AssetId)
	if err != nil {
		log.Error("Error in fetching active status")
		return core.NilHash, err
	}
	if currentStatus == modifyAssetInput.Status {
		log.Errorf("Asset %d has the active status already set to %t", modifyAssetInput.AssetId, modifyAssetInput.Status)
		return core.NilHash, nil
	}
	_, err = cmdUtilsMockery.WaitForAppropriateState(client, "modify asset status", 4)
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

	txnOpts := razorUtilsMockery.GetTxnOpts(txnArgs)
	log.Infof("Changing active status of asset: %d from %t to %t", modifyAssetInput.AssetId, !modifyAssetInput.Status, modifyAssetInput.Status)
	txn, err := assetManagerUtilsMockery.SetCollectionStatus(client, txnOpts, modifyAssetInput.Status, modifyAssetInput.AssetId)
	if err != nil {
		return core.NilHash, err
	}
	log.Info("Txn Hash: ", transactionUtilsMockery.Hash(txn).String())
	return transactionUtilsMockery.Hash(txn), nil
}

func init() {

	razorUtilsMockery = UtilsMockery{}
	cmdUtilsMockery = &UtilsStructMockery{}
	flagSetUtilsMockery = FLagSetUtilsMockery{}
	assetManagerUtilsMockery = AssetManagerUtilsMockery{}
	transactionUtilsMockery = TransactionUtilsMockery{}
	utils.Options = &utils.OptionsStruct{}
	utils.UtilsInterface = &utils.UtilsStruct{}

	rootCmd.AddCommand(modifyAssetStatusCmd)

	var (
		Address string
		AssetId uint16
		Status  string
	)

	modifyAssetStatusCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the user")
	modifyAssetStatusCmd.Flags().Uint16VarP(&AssetId, "assetId", "", 0, "assetId of the asset")
	modifyAssetStatusCmd.Flags().StringVarP(&Status, "status", "", "true", "active status of the asset")

	addressErr := modifyAssetStatusCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addressErr)
	assetIdErr := modifyAssetStatusCmd.MarkFlagRequired("assetId")
	utils.CheckError("Asset Id error: ", assetIdErr)
	statusErr := modifyAssetStatusCmd.MarkFlagRequired("status")
	utils.CheckError("Status error: ", statusErr)
}
