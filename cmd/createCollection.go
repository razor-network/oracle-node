package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"
)

var createCollectionCmd = &cobra.Command{
	Use:   "createCollection",
	Short: "createCollection can be used to create collections if existing jobs are present",
	Long: `A collection is a group of jobs that reports the aggregated value of jobs. createCollection can be used to club multiple jobs into one collection bound by an aggregation method.

Example: 
  ./razor createCollection --name btcCollectionMean --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --jobIds 1,2 --aggregation 2 --power 2

Note: 
  This command only works for the admin.
`,
	Run: func(cmd *cobra.Command, args []string) {
		utilsStruct := UtilsStruct{
			razorUtils:        razorUtils,
			assetManagerUtils: assetManagerUtils,
			transactionUtils:  transactionUtils,
			flagSetUtils:      flagSetUtils,
			cmdUtils:          cmdUtils,
			keystoreUtils:     keystoreUtils,
		}
		config, err := GetConfigData()
		utils.CheckError("Error in getting config: ", err)

		txn, err := utilsStruct.createCollection(cmd.Flags(), config)
		utils.CheckError("CreateCollection error: ", err)
		utils.WaitForBlockCompletion(utils.ConnectToClient(config.Provider), txn.String())
	},
}

func (utilsStruct UtilsStruct) createCollection(flagSet *pflag.FlagSet, config types.Configurations) (common.Hash, error) {
	password := utilsStruct.razorUtils.AssignPassword(flagSet)
	name, err := utilsStruct.flagSetUtils.GetStringName(flagSet)
	if err != nil {
		return core.NilHash, err
	}
	address, err := utilsStruct.flagSetUtils.GetStringAddress(flagSet)
	if err != nil {
		return core.NilHash, err
	}
	jobIdInUint, err := utilsStruct.flagSetUtils.GetUintSliceJobIds(flagSet)
	if err != nil {
		return core.NilHash, err
	}
	aggregation, err := utilsStruct.flagSetUtils.GetUint32Aggregation(flagSet)
	if err != nil {
		return core.NilHash, err
	}
	power, err := utilsStruct.flagSetUtils.GetInt8Power(flagSet)
	if err != nil {
		return core.NilHash, err
	}

	client := utilsStruct.razorUtils.ConnectToClient(config.Provider)

	jobIds := utilsStruct.razorUtils.ConvertUintArrayToUint8Array(jobIdInUint)
	_, err = utilsStruct.cmdUtils.WaitForAppropriateState(client, address, "create collection", utilsStruct, 4)
	if err != nil {
		log.Error("Error in fetching state")
		return core.NilHash, err
	}
	txnOpts := utilsStruct.razorUtils.GetTxnOpts(types.TransactionOptions{
		Client:          client,
		Password:        password,
		AccountAddress:  address,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.AssetManagerAddress,
		MethodName:      "createCollection",
		Parameters:      []interface{}{jobIds, aggregation, power, name},
		ABI:             bindings.AssetManagerABI,
	}, utilsStruct)
	txn, err := utilsStruct.assetManagerUtils.CreateCollection(client, txnOpts, jobIds, aggregation, power, name)
	if err != nil {
		log.Error("Error in creating collection")
		return core.NilHash, err
	}
	log.Info("Creating collection...")
	log.Info("Txn Hash: ", utilsStruct.transactionUtils.Hash(txn))
	return utilsStruct.transactionUtils.Hash(txn), nil
}

func init() {

	razorUtils = Utils{}
	assetManagerUtils = AssetManagerUtils{}
	transactionUtils = TransactionUtils{}
	flagSetUtils = FlagSetUtils{}
	cmdUtils = UtilsCmd{}
	keystoreUtils = KeystoreUtils{}

	rootCmd.AddCommand(createCollectionCmd)

	var (
		Name              string
		Account           string
		JobIds            []uint
		AggregationMethod uint32
		Password          string
		Power             int8
	)

	createCollectionCmd.Flags().StringVarP(&Name, "name", "n", "", "name of the collection")
	createCollectionCmd.Flags().StringVarP(&Account, "address", "a", "", "address of the job creator")
	createCollectionCmd.Flags().UintSliceVarP(&JobIds, "jobIds", "", []uint{}, "job ids for the  collection")
	createCollectionCmd.Flags().Uint32VarP(&AggregationMethod, "aggregation", "", 1, "aggregation method to be used")
	createCollectionCmd.Flags().Int8VarP(&Power, "power", "", 0, "multiplier for the collection")
	createCollectionCmd.Flags().StringVarP(&Password, "password", "", "", "password path of job creator to protect the keystore")

	nameErr := createCollectionCmd.MarkFlagRequired("name")
	utils.CheckError("Name error: ", nameErr)
	addrErr := createCollectionCmd.MarkFlagRequired("address")
	utils.CheckError("Address Error: ", addrErr)
	jobIdErr := createCollectionCmd.MarkFlagRequired("jobIds")
	utils.CheckError("Job Id Error: ", jobIdErr)
	powerErr := createCollectionCmd.MarkFlagRequired("power")
	utils.CheckError("Power Error: ", powerErr)
}
