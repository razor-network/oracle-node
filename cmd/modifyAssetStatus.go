package cmd

import (
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"razor/core"
	"razor/core/types"
	"razor/utils"
	"strconv"

	"github.com/spf13/cobra"
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

		address, _ := cmd.Flags().GetString("address")
		assetId, _ := cmd.Flags().GetUint8("assetId")
		statusString, _ := cmd.Flags().GetString("status")

		status, err := strconv.ParseBool(statusString)
		utils.CheckError("Error in parsing status to boolean: ", err)

		password := utils.PasswordPrompt()

		client := utils.ConnectToClient(config.Provider)

		currentStatus, err := CheckCurrentStatus(client, address, assetId)
		utils.CheckError("Error in fetching active status: ", err)
		if currentStatus == status {
			log.Fatalf("Asset %d has the active status already set to %t", assetId, status)
		}

		err = ModifyAssetStatus(types.TransactionOptions{
			Client:         client,
			Password:       password,
			AccountAddress: address,
			ChainId:        core.ChainId,
			Config:         config,
		}, assetId, status)
		utils.CheckError("Error in changing asset active status: ", err)
	},
}

func CheckCurrentStatus(client *ethclient.Client, address string, assetId uint8) (bool, error) {
	assetManager := utils.GetAssetManager(client)
	callOpts := utils.GetOptions(false, address, "")
	return assetManager.GetActiveStatus(&callOpts, assetId)
}

func ModifyAssetStatus(transactionOpts types.TransactionOptions, assetId uint8, status bool) error {
	assetManager := utils.GetAssetManager(transactionOpts.Client)
	txnOpts := utils.GetTxnOpts(transactionOpts)
	log.Infof("Changing active status of asset: %d from %t to %t", assetId, !status, status)
	txn, err := assetManager.SetAssetStatus(txnOpts, status, assetId)
	if err != nil {
		return err
	}
	log.Info("Asset active status changed")
	log.Info("Txn Hash: ", txn.Hash().String())
	utils.WaitForBlockCompletion(transactionOpts.Client, txn.Hash().String())
	return nil
}

func init() {
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
