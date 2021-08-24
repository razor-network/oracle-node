package cmd

import (
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/utils"
	"strconv"

	"github.com/spf13/cobra"
)

var modifyAssetStatusCmd = &cobra.Command{
	Use:   "modifyAssetStatus",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		utils.CheckError("Error in fetching config data: ", err)

		address, _ := cmd.Flags().GetString("address")
		assetId, _ := cmd.Flags().GetString("assetId")
		statusString, _ := cmd.Flags().GetString("status")

		status, err := strconv.ParseBool(statusString)
		utils.CheckError("Error in parsing status to boolean: ", err)

		assetIdInBigInt, ok := new(big.Int).SetString(assetId, 10)
		if !ok {
			log.Fatal("SetString: error")
		}

		password := utils.PasswordPrompt()

		client := utils.ConnectToClient(config.Provider)

		currentStatus, err := CheckCurrentStatus(client, address, assetIdInBigInt)
		utils.CheckError("Error in fetching active status: ", err)
		if currentStatus == status {
			log.Fatalf("Asset %s has the active status already set to %t", assetIdInBigInt, status)
		}

		err = ModifyAssetStatus(types.TransactionOptions{
			Client:         client,
			Password:       password,
			AccountAddress: address,
			ChainId:        core.ChainId,
			Config:         config,
		}, assetIdInBigInt, status)
		utils.CheckError("Error in changing asset active status: ", err)
	},
}

func CheckCurrentStatus(client *ethclient.Client, address string, assetId *big.Int) (bool, error) {
	assetManager := utils.GetAssetManager(client)
	callOpts := utils.GetOptions(false, address, "")
	return assetManager.GetActiveStatus(&callOpts, assetId)
}

func ModifyAssetStatus(transactionOpts types.TransactionOptions, assetId *big.Int, status bool) error {
	assetManager := utils.GetAssetManager(transactionOpts.Client)
	txnOpts := utils.GetTxnOpts(transactionOpts)
	log.Infof("Changing active status of asset: %s from %t to %t", assetId, !status, status)
	txn, err := assetManager.SetAssetStatus(txnOpts, assetId, status)
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
		AssetId string
		Status string
	)

	modifyAssetStatusCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the user")
	modifyAssetStatusCmd.Flags().StringVarP(&AssetId, "assetId", "", "", "assetId of the asset")
	modifyAssetStatusCmd.Flags().StringVarP(&Status, "status", "", "true", "active status of the asset")

	addressErr := modifyAssetStatusCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addressErr)
	assetIdErr := modifyAssetStatusCmd.MarkFlagRequired("assetId")
	utils.CheckError("Asset Id error: ", assetIdErr)
	statusErr := modifyAssetStatusCmd.MarkFlagRequired("status")
	utils.CheckError("Status error: ", statusErr)
}
