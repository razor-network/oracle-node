//Package cmd provides all functions related to command line
package cmd

import (
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/rpc"
	"razor/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var modifyCollectionStatusCmd = &cobra.Command{
	Use:   "modifyCollectionStatus",
	Short: "[ADMIN ONLY]modify the active status of an collection",
	Long: `modifyCollectionStatus can be used by admins to change the status of an collection
Example:	
  ./razor modifyCollectionStatus --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --collectionId 1 --status true --logLevel modifyLogs`,
	Run: initialiseModifyCollectionStatus,
}

//This function initialises the ExecuteModifyCollectionStatus function
func initialiseModifyCollectionStatus(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteModifyCollectionStatus(cmd.Flags())
}

//This function sets the flags appropriately and executes the ModifyCollectionStatus function
func (*UtilsStruct) ExecuteModifyCollectionStatus(flagSet *pflag.FlagSet) {
	config, rpcParameters, account, err := InitializeCommandDependencies(flagSet)
	utils.CheckError("Error in initialising command dependencies: ", err)

	collectionId, err := flagSetUtils.GetUint16CollectionId(flagSet)
	utils.CheckError("Error in getting collectionId: ", err)

	statusString, err := flagSetUtils.GetStringStatus(flagSet)
	utils.CheckError("Error in getting status: ", err)

	status, err := stringUtils.ParseBool(statusString)
	utils.CheckError("Error in parsing status: ", err)

	modifyCollectionInput := types.ModifyCollectionInput{
		Status:       status,
		CollectionId: collectionId,
		Account:      account,
	}

	txn, err := cmdUtils.ModifyCollectionStatus(rpcParameters, config, modifyCollectionInput)
	utils.CheckError("Error in changing collection active status: ", err)
	if txn != core.NilHash {
		err = razorUtils.WaitForBlockCompletion(rpcParameters, txn.Hex())
		utils.CheckError("Error in WaitForBlockCompletion for modifyCollectionStatus: ", err)
	}
}

//This function allows the admin to modify the active status of collection
func (*UtilsStruct) ModifyCollectionStatus(rpcParameters rpc.RPCParameters, config types.Configurations, modifyCollectionInput types.ModifyCollectionInput) (common.Hash, error) {
	currentStatus, err := razorUtils.GetActiveStatus(rpcParameters, modifyCollectionInput.CollectionId)
	if err != nil {
		log.Error("Error in fetching active status")
		return core.NilHash, err
	}
	log.Debug("ModifyCollectionStatus: Current status of collection: ", currentStatus)
	if currentStatus == modifyCollectionInput.Status {
		log.Errorf("Collection %d has the active status already set to %t", modifyCollectionInput.CollectionId, modifyCollectionInput.Status)
		return core.NilHash, nil
	}
	_, err = cmdUtils.WaitForAppropriateState(rpcParameters, "modify collection status", 4)
	if err != nil {
		return core.NilHash, err
	}

	txnArgs := types.TransactionOptions{
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.CollectionManagerAddress,
		MethodName:      "setCollectionStatus",
		Parameters:      []interface{}{modifyCollectionInput.Status, modifyCollectionInput.CollectionId},
		ABI:             bindings.CollectionManagerMetaData.ABI,
		Account:         modifyCollectionInput.Account,
	}

	txnOpts, err := razorUtils.GetTxnOpts(rpcParameters, txnArgs)
	if err != nil {
		return core.NilHash, err
	}
	log.Infof("Changing active status of collection: %d from %t to %t", modifyCollectionInput.CollectionId, !modifyCollectionInput.Status, modifyCollectionInput.Status)
	client, err := rpcParameters.RPCManager.GetBestRPCClient()
	if err != nil {
		return core.NilHash, err
	}

	log.Debugf("Executing SetCollectionStatus transaction with status = %v, collectionId = %d", modifyCollectionInput.Status, modifyCollectionInput.CollectionId)
	txn, err := assetManagerUtils.SetCollectionStatus(client, txnOpts, modifyCollectionInput.Status, modifyCollectionInput.CollectionId)
	if err != nil {
		return core.NilHash, err
	}
	txnHash := transactionUtils.Hash(txn)
	log.Info("Txn Hash: ", txnHash.Hex())
	return txnHash, nil
}

func init() {
	rootCmd.AddCommand(modifyCollectionStatusCmd)
	var (
		Address      string
		CollectionId uint16
		Status       string
		Password     string
	)

	modifyCollectionStatusCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the user")
	modifyCollectionStatusCmd.Flags().Uint16VarP(&CollectionId, "collectionId", "", 0, "collectionId of the collection")
	modifyCollectionStatusCmd.Flags().StringVarP(&Status, "status", "", "true", "active status of the collection")
	modifyCollectionStatusCmd.Flags().StringVarP(&Password, "password", "", "", "password path of user to protect the keystore")

	addressErr := modifyCollectionStatusCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addressErr)
	collectionIdErr := modifyCollectionStatusCmd.MarkFlagRequired("collectionId")
	utils.CheckError("Collection Id error: ", collectionIdErr)
	statusErr := modifyCollectionStatusCmd.MarkFlagRequired("status")
	utils.CheckError("Status error: ", statusErr)
}
