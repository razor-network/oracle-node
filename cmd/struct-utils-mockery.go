package cmd

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"razor/core/types"
	"razor/path"
	"razor/utils"
	"time"
)

func (u UtilsMockery) GetConfigFilePath() (string, error) {
	return path.GetConfigFilePath()
}

func (u UtilsMockery) ViperWriteConfigAs(path string) error {
	return viper.WriteConfigAs(path)
}

func (u UtilsMockery) GetEpoch(client *ethclient.Client) (uint32, error) {
	return utils.GetEpoch(client)
}

func (u UtilsMockery) GetOptions() bind.CallOpts {
	return utils.UtilsInterface.GetOptions()
}

func (u UtilsMockery) CalculateBlockTime(client *ethclient.Client) int64 {
	return utils.CalculateBlockTime(client)
}

func (u UtilsMockery) Sleep(duration time.Duration) {
	utils.Sleep(duration)
}

func (u UtilsMockery) GetTxnOpts(transactionData types.TransactionOptions) *bind.TransactOpts {
	utilsInterface := utils.StartRazor(utils.OptionsPackageStruct{
		Options:        utils.Options,
		UtilsInterface: utils.UtilsInterface,
	})
	utils.Options = &utils.OptionsStruct{}
	utils.UtilsInterface = &utils.UtilsStruct{}
	return utilsInterface.GetTxnOpts(transactionData)
}

func (u UtilsMockery) GetConfigData() (types.Configurations, error) {
	return cmdUtilsMockery.GetConfigData()
}

func (u UtilsMockery) AssignPassword(flagSet *pflag.FlagSet) string {
	return utils.AssignPassword(flagSet)
}

func (u UtilsMockery) GetStringAddress(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("address")
}

func (u UtilsMockery) GetUint32BountyId(flagSet *pflag.FlagSet) (uint32, error) {
	return flagSet.GetUint32("bountyId")
}

func (u UtilsMockery) ConnectToClient(provider string) *ethclient.Client {
	return utils.ConnectToClient(provider)
}

func (u UtilsMockery) WaitForBlockCompletion(client *ethclient.Client, hashToRead string) int {
	return utils.WaitForBlockCompletion(client, hashToRead)
}

func (transactionUtils TransactionUtilsMockery) Hash(txn *Types.Transaction) common.Hash {
	return txn.Hash()
}

func (stakeManagerUtils StakeManagerUtilsMockery) GetBountyLock(client *ethclient.Client, opts *bind.CallOpts, bountyId uint32) (types.BountyLock, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.BountyLocks(opts, bountyId)
}

func (stakeManagerUtils StakeManagerUtilsMockery) RedeemBounty(client *ethclient.Client, opts *bind.TransactOpts, bountyId uint32) (*Types.Transaction, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.RedeemBounty(opts, bountyId)
}

func (flagSetUtils FLagSetUtilsMockery) GetStringProvider(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("provider")
}

func (flagSetUtils FLagSetUtilsMockery) GetFloat32GasMultiplier(flagSet *pflag.FlagSet) (float32, error) {
	return flagSet.GetFloat32("gasmultiplier")
}

func (flagSetUtils FLagSetUtilsMockery) GetInt32Buffer(flagSet *pflag.FlagSet) (int32, error) {
	return flagSet.GetInt32("buffer")
}

func (flagSetUtils FLagSetUtilsMockery) GetInt32Wait(flagSet *pflag.FlagSet) (int32, error) {
	return flagSet.GetInt32("wait")
}

func (flagSetUtils FLagSetUtilsMockery) GetInt32GasPrice(flagSet *pflag.FlagSet) (int32, error) {
	return flagSet.GetInt32("gasprice")
}

func (flagSetUtils FLagSetUtilsMockery) GetStringLogLevel(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("logLevel")
}

func (flagSetUtils FLagSetUtilsMockery) GetFloat32GasLimit(flagSet *pflag.FlagSet) (float32, error) {
	return flagSet.GetFloat32("gasLimit")
}

func (flagSetUtils FLagSetUtilsMockery) GetBoolAutoWithdraw(flagSet *pflag.FlagSet) (bool, error) {
	return flagSet.GetBool("autoWithdraw")
}

func (flagSetUtils FLagSetUtilsMockery) GetUint32BountyId(flagSet *pflag.FlagSet) (uint32, error) {
	return flagSet.GetUint32("bountyId")
}

func (flagSetUtils FLagSetUtilsMockery) GetRootStringProvider() (string, error) {
	return rootCmd.PersistentFlags().GetString("provider")
}

func (flagSetUtils FLagSetUtilsMockery) GetRootFloat32GasMultiplier() (float32, error) {
	return rootCmd.PersistentFlags().GetFloat32("gasmultiplier")
}

func (flagSetUtils FLagSetUtilsMockery) GetRootInt32Buffer() (int32, error) {
	return rootCmd.PersistentFlags().GetInt32("buffer")
}

func (flagSetUtils FLagSetUtilsMockery) GetRootInt32Wait() (int32, error) {
	return rootCmd.PersistentFlags().GetInt32("wait")
}

func (flagSetUtils FLagSetUtilsMockery) GetRootInt32GasPrice() (int32, error) {
	return rootCmd.PersistentFlags().GetInt32("gasprice")
}

func (flagSetUtils FLagSetUtilsMockery) GetRootStringLogLevel() (string, error) {
	return rootCmd.PersistentFlags().GetString("logLevel")
}

func (flagSetUtils FLagSetUtilsMockery) GetRootFloat32GasLimit() (float32, error) {
	return rootCmd.PersistentFlags().GetFloat32("gasLimit")
}

func (flagSetUtils FLagSetUtilsMockery) GetStringFrom(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("from")
}

func (flagSetUtils FLagSetUtilsMockery) GetStringTo(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("to")
}

func (flagSetUtils FLagSetUtilsMockery) GetStringAddress(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("address")
}

func (flagSetUtils FLagSetUtilsMockery) GetUint32StakerId(flagSet *pflag.FlagSet) (uint32, error) {
	return flagSet.GetUint32("stakerId")
}

func (flagSetUtils FLagSetUtilsMockery) GetStringName(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("name")
}

func (flagSetUtils FLagSetUtilsMockery) GetStringUrl(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("url")
}

func (flagSetUtils FLagSetUtilsMockery) GetStringSelector(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("selector")
}

func (flagSetUtils FLagSetUtilsMockery) GetInt8Power(flagSet *pflag.FlagSet) (int8, error) {
	return flagSet.GetInt8("power")
}

func (flagSetUtils FLagSetUtilsMockery) GetUint8Weight(flagSet *pflag.FlagSet) (uint8, error) {
	return flagSet.GetUint8("weight")
}

func (flagSetUtils FLagSetUtilsMockery) GetUint16AssetId(flagSet *pflag.FlagSet) (uint16, error) {
	return flagSet.GetUint16("assetId")
}

func (flagSetUtils FLagSetUtilsMockery) GetUint8SelectorType(flagSet *pflag.FlagSet) (uint8, error) {
	return flagSet.GetUint8("selectorType")
}

func (flagSetUtils FLagSetUtilsMockery) GetStringStatus(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("status")
}

func (flagSetUtils FLagSetUtilsMockery) GetUint8Commission(flagSet *pflag.FlagSet) (uint8, error) {
	return flagSet.GetUint8("commission")
}

func (flagSetUtils FLagSetUtilsMockery) GetUintSliceJobIds(flagSet *pflag.FlagSet) ([]uint, error) {
	return flagSet.GetUintSlice("jobIds")
}

func (flagSetUtils FLagSetUtilsMockery) GetUint32Aggregation(flagSet *pflag.FlagSet) (uint32, error) {
	return flagSet.GetUint32("aggregation")
}

func (flagSetUtils FLagSetUtilsMockery) GetUint16JobId(flagSet *pflag.FlagSet) (uint16, error) {
	return flagSet.GetUint16("jobId")
}

func (flagSetUtils FLagSetUtilsMockery) GetUint16CollectionId(flagSet *pflag.FlagSet) (uint16, error) {
	return flagSet.GetUint16("collectionId")
}
