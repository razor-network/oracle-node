package cmd

import (
	"crypto/ecdsa"
	"github.com/avast/retry-go"
	ethAccounts "github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"math/big"
	"razor/core/types"
	"razor/path"
	"razor/pkg/bindings"
	"razor/utils"
	"strconv"
	"time"
)

func (u Utils) GetConfigFilePath() (string, error) {
	return path.GetConfigFilePath()
}

func (u Utils) ViperWriteConfigAs(path string) error {
	return viper.WriteConfigAs(path)
}

func (u Utils) GetEpoch(client *ethclient.Client) (uint32, error) {
	return utils.GetEpoch(client)
}

func (u Utils) GetUpdatedEpoch(client *ethclient.Client) (uint32, error) {
	return utils.GetEpoch(client)
}

func (u Utils) GetOptions() bind.CallOpts {
	return utils.UtilsInterface.GetOptions()
}

func (u Utils) CalculateBlockTime(client *ethclient.Client) int64 {
	return utils.CalculateBlockTime(client)
}

func (u Utils) Sleep(duration time.Duration) {
	utils.Sleep(duration)
}

func (u Utils) GetTxnOpts(transactionData types.TransactionOptions) *bind.TransactOpts {
	utilsInterface := utils.StartRazor(utils.OptionsPackageStruct{
		Options:        utils.Options,
		UtilsInterface: utils.UtilsInterface,
	})
	utils.Options = &utils.OptionsStruct{}
	utils.UtilsInterface = &utils.UtilsStruct{}
	return utilsInterface.GetTxnOpts(transactionData)
}

func (u Utils) GetConfigData() (types.Configurations, error) {
	return cmdUtils.GetConfigData()
}

func (u Utils) AssignPassword(flagSet *pflag.FlagSet) string {
	return utils.AssignPassword(flagSet)
}

func (u Utils) GetStringAddress(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("address")
}

func (u Utils) GetUint32BountyId(flagSet *pflag.FlagSet) (uint32, error) {
	return flagSet.GetUint32("bountyId")
}

func (u Utils) ConnectToClient(provider string) *ethclient.Client {
	return utils.ConnectToClient(provider)
}

func (u Utils) WaitForBlockCompletion(client *ethclient.Client, hashToRead string) int {
	return utils.WaitForBlockCompletion(client, hashToRead)
}

func (u Utils) GetNumActiveAssets(client *ethclient.Client) (*big.Int, error) {
	return utils.UtilsInterface.GetNumActiveAssets(client)
}

func (u Utils) GetRogueRandomValue(value int) *big.Int {
	return utils.GetRogueRandomValue(value)
}

func (u Utils) GetActiveAssetsData(client *ethclient.Client, epoch uint32) ([]*big.Int, error) {
	return utils.UtilsInterface.GetActiveAssetsData(client, epoch)
}

func (u Utils) GetDelayedState(client *ethclient.Client, buffer int32) (int64, error) {
	return utils.GetDelayedState(client, buffer)
}

func (u Utils) GetDefaultPath() (string, error) {
	return path.GetDefaultPath()
}

func (u Utils) GetJobFilePath() (string, error) {
	return path.GetJobFilePath()
}

func (u Utils) FetchBalance(client *ethclient.Client, accountAddress string) (*big.Int, error) {
	return utils.FetchBalance(client, accountAddress)
}

func (u Utils) IsFlagPassed(name string) bool {
	return utils.IsFlagPassed(name)
}

func (u Utils) GetFractionalAmountInWei(amount *big.Int, power string) (*big.Int, error) {
	return utils.GetFractionalAmountInWei(amount, power)
}

func (u Utils) GetAmountInWei(amount *big.Int) *big.Int {
	return utils.GetAmountInWei(amount)
}

func (u Utils) CheckAmountAndBalance(amountInWei *big.Int, balance *big.Int) *big.Int {
	return utils.CheckAmountAndBalance(amountInWei, balance)
}

func (u Utils) GetAmountInDecimal(amountInWei *big.Int) *big.Float {
	return utils.GetAmountInDecimal(amountInWei)
}

func (u Utils) GetEpochLastCommitted(client *ethclient.Client, stakerId uint32) (uint32, error) {
	return utils.UtilsInterface.GetEpochLastCommitted(client, stakerId)
}

func (u Utils) GetCommitments(client *ethclient.Client, address string) ([32]byte, error) {
	return utils.UtilsInterface.GetCommitments(client, address)
}

func (u Utils) AllZero(bytesValue [32]byte) bool {
	return utils.AllZero(bytesValue)
}

func (u Utils) ConvertUintArrayToUint16Array(uintArr []uint) []uint16 {
	return utils.ConvertUintArrayToUint16Array(uintArr)
}

func (u Utils) GetStateName(state int64) string {
	return utils.GetStateName(state)
}

func (u Utils) GetJobs(client *ethclient.Client) ([]bindings.StructsJob, error) {
	return utils.UtilsInterface.GetJobs(client)
}

func (u Utils) CheckEthBalanceIsZero(client *ethclient.Client, address string) {
	utils.CheckEthBalanceIsZero(client, address)
}

func (u Utils) AssignStakerId(flagSet *pflag.FlagSet, client *ethclient.Client, address string) (uint32, error) {
	return utils.AssignStakerId(flagSet, client, address)
}

func (u Utils) GetLock(client *ethclient.Client, address string, stakerId uint32) (types.Locks, error) {
	return utils.GetLock(client, address, stakerId)
}

func (u Utils) GetStaker(client *ethclient.Client, address string, stakerId uint32) (bindings.StructsStaker, error) {
	return utils.GetStaker(client, address, stakerId)
}

func (u Utils) GetUpdatedStaker(client *ethclient.Client, address string, stakerId uint32) (bindings.StructsStaker, error) {
	return utils.GetStaker(client, address, stakerId)
}

func (u Utils) GetStakedToken(client *ethclient.Client, address common.Address) *bindings.StakedToken {
	return utils.GetStakedToken(client, address)
}

func (u Utils) ConvertSRZRToRZR(sAmount *big.Int, currentStake *big.Int, totalSupply *big.Int) *big.Int {
	return utils.ConvertSRZRToRZR(sAmount, currentStake, totalSupply)
}

func (u Utils) ConvertRZRToSRZR(sAmount *big.Int, currentStake *big.Int, totalSupply *big.Int) (*big.Int, error) {
	return utils.ConvertRZRToSRZR(sAmount, currentStake, totalSupply)
}

func (u Utils) GetWithdrawReleasePeriod(client *ethclient.Client, address string) (uint8, error) {
	return utils.GetWithdrawReleasePeriod(client, address)
}

func (u Utils) GetInfluenceSnapshot(client *ethclient.Client, stakerId uint32, epoch uint32) (*big.Int, error) {
	return utils.UtilsInterface.GetInfluenceSnapshot(client, stakerId, epoch)
}

func (u Utils) GetCollections(client *ethclient.Client) ([]bindings.StructsCollection, error) {
	return utils.UtilsInterface.GetCollections(client)
}

func (u Utils) GetNumberOfStakers(client *ethclient.Client, address string) (uint32, error) {
	return utils.GetNumberOfStakers(client, address)
}

func (u Utils) GetRandaoHash(client *ethclient.Client) ([32]byte, error) {
	return utils.UtilsInterface.GetRandaoHash(client)
}

//TODO: Check direct usage from utils package without implementing it here

func (u Utils) GetNumberOfProposedBlocks(client *ethclient.Client, epoch uint32) (uint8, error) {
	return utils.UtilsInterface.GetNumberOfProposedBlocks(client, epoch)
}

func (u Utils) GetMaxAltBlocks(client *ethclient.Client) (uint8, error) {
	return utils.UtilsInterface.GetMaxAltBlocks(client)
}

func (u Utils) GetProposedBlock(client *ethclient.Client, epoch uint32, proposedBlockId uint32) (bindings.StructsBlock, error) {
	return utils.UtilsInterface.GetProposedBlock(client, epoch, proposedBlockId)
}

func (u Utils) GetEpochLastRevealed(client *ethclient.Client, stakerId uint32) (uint32, error) {
	return utils.UtilsInterface.GetEpochLastRevealed(client, stakerId)
}

func (u Utils) GetVoteValue(client *ethclient.Client, assetId uint16, stakerId uint32) (*big.Int, error) {
	return utils.UtilsInterface.GetVoteValue(client, assetId, stakerId)
}

func (u Utils) GetTotalInfluenceRevealed(client *ethclient.Client, epoch uint32) (*big.Int, error) {
	return utils.UtilsInterface.GetTotalInfluenceRevealed(client, epoch)
}

func (u Utils) ConvertBigIntArrayToUint32Array(bigIntArray []*big.Int) []uint32 {
	return utils.ConvertBigIntArrayToUint32Array(bigIntArray)
}

func (u Utils) GetActiveAssetIds(client *ethclient.Client) ([]uint16, error) {
	return utils.UtilsInterface.GetActiveAssetIds(client)
}

func (u Utils) GetBlockManager(client *ethclient.Client) *bindings.BlockManager {
	return utils.UtilsInterface.GetBlockManager(client)
}

func (u Utils) GetVotes(client *ethclient.Client, stakerId uint32) (bindings.StructsVote, error) {
	return utils.UtilsInterface.GetVotes(client, stakerId)
}

func (u Utils) GetSortedProposedBlockIds(client *ethclient.Client, epoch uint32) ([]uint32, error) {
	return utils.UtilsInterface.GetSortedProposedBlockIds(client, epoch)
}

func (u Utils) ParseBool(str string) (bool, error) {
	return strconv.ParseBool(str)
}

func (u Utils) GetStakerId(client *ethclient.Client, address string) (uint32, error) {
	return utils.GetStakerId(client, address)
}

func (u Utils) PrivateKeyPrompt() string {
	return utils.PrivateKeyPrompt()
}

func (u Utils) PasswordPrompt() string {
	return utils.PasswordPrompt()
}

func (u Utils) GetMaxCommission(client *ethclient.Client) (uint8, error) {
	return utils.GetMaxCommission(client)
}

func (u Utils) GetEpochLimitForUpdateCommission(client *ethclient.Client) (uint16, error) {
	return utils.GetEpochLimitForUpdateCommission(client)
}

func (u Utils) GetStakeSnapshot(client *ethclient.Client, stakerId uint32, epoch uint32) (*big.Int, error) {
	return utils.UtilsInterface.GetStakeSnapshot(client, stakerId, epoch)
}

func (u Utils) DeleteJobFromJSON(s string, jobId string) error {
	return utils.DeleteJobFromJSON(s, jobId)
}

func (u Utils) AddJobToJSON(s string, job *types.StructsJob) error {
	return utils.AddJobToJSON(s, job)
}

func (transactionUtils TransactionUtils) Hash(txn *Types.Transaction) common.Hash {
	return txn.Hash()
}

func (stakeManagerUtils StakeManagerUtils) Stake(client *ethclient.Client, txnOpts *bind.TransactOpts, epoch uint32, amount *big.Int) (*Types.Transaction, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.Stake(txnOpts, epoch, amount)
}

func (stakeManagerUtils StakeManagerUtils) ExtendLock(client *ethclient.Client, opts *bind.TransactOpts, stakerId uint32) (*Types.Transaction, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.ExtendLock(opts, stakerId)
}

func (stakeManagerUtils StakeManagerUtils) Delegate(client *ethclient.Client, opts *bind.TransactOpts, stakerId uint32, amount *big.Int) (*Types.Transaction, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.Delegate(opts, stakerId, amount)
}

func (stakeManagerUtils StakeManagerUtils) Withdraw(client *ethclient.Client, opts *bind.TransactOpts, stakerId uint32) (*Types.Transaction, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.Withdraw(opts, stakerId)
}

func (stakeManagerUtils StakeManagerUtils) SetDelegationAcceptance(client *ethclient.Client, opts *bind.TransactOpts, status bool) (*Types.Transaction, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.SetDelegationAcceptance(opts, status)
}

func (stakeManagerUtils StakeManagerUtils) UpdateCommission(client *ethclient.Client, opts *bind.TransactOpts, commission uint8) (*Types.Transaction, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.UpdateCommission(opts, commission)
}

func (stakeManagerUtils StakeManagerUtils) Unstake(client *ethclient.Client, opts *bind.TransactOpts, stakerId uint32, sAmount *big.Int) (*Types.Transaction, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.Unstake(opts, stakerId, sAmount)
}

func (stakeManagerUtils StakeManagerUtils) RedeemBounty(client *ethclient.Client, opts *bind.TransactOpts, bountyId uint32) (*Types.Transaction, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.RedeemBounty(opts, bountyId)
}

func (stakeManagerUtils StakeManagerUtils) StakerInfo(client *ethclient.Client, opts *bind.CallOpts, stakerId uint32) (types.Staker, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.Stakers(opts, stakerId)
}

func (stakeManagerUtils StakeManagerUtils) GetMaturity(client *ethclient.Client, opts *bind.CallOpts, age uint32) (uint16, error) {
	stakeManager := utils.GetStakeManager(client)
	index := age / 10000
	return stakeManager.Maturities(opts, big.NewInt(int64(index)))
}

func (stakeManagerUtils StakeManagerUtils) GetBountyLock(client *ethclient.Client, opts *bind.CallOpts, bountyId uint32) (types.BountyLock, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.BountyLocks(opts, bountyId)
}

func (stakeManagerUtils StakeManagerUtils) BalanceOf(stakedToken *bindings.StakedToken, callOpts *bind.CallOpts, address common.Address) (*big.Int, error) {
	return stakedToken.BalanceOf(callOpts, address)
}

func (stakeManagerUtils StakeManagerUtils) GetTotalSupply(token *bindings.StakedToken, callOpts *bind.CallOpts) (*big.Int, error) {
	return token.TotalSupply(callOpts)
}

func (blockManagerUtils BlockManagerUtils) ClaimBlockReward(client *ethclient.Client, opts *bind.TransactOpts) (*Types.Transaction, error) {
	blockManager := utils.UtilsInterface.GetBlockManager(client)
	return blockManager.ClaimBlockReward(opts)
}

func (blockManagerUtils BlockManagerUtils) FinalizeDispute(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, blockIndex uint8) (*Types.Transaction, error) {
	blockManager := utils.UtilsInterface.GetBlockManager(client)
	var (
		txn *Types.Transaction
		err error
	)
	err = retry.Do(func() error {
		txn, err = blockManager.FinalizeDispute(opts, epoch, blockIndex)
		if err != nil {
			log.Error("Error in finalizing dispute.. Retrying")
			return err
		}
		return nil
	}, retry.Attempts(3))
	if err != nil {
		return nil, err
	}
	return txn, nil
}

func (blockManagerUtils BlockManagerUtils) DisputeBiggestStakeProposed(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, blockIndex uint8, correctBiggestStakerId uint32) (*Types.Transaction, error) {
	blockManager := utils.UtilsInterface.GetBlockManager(client)
	var (
		txn *Types.Transaction
		err error
	)
	err = retry.Do(func() error {
		txn, err = blockManager.DisputeBiggestStakeProposed(opts, epoch, blockIndex, correctBiggestStakerId)
		if err != nil {
			log.Error("Error in disputing biggest influence proposed.. Retrying")
			return err
		}
		return nil
	}, retry.Attempts(3))
	if err != nil {
		return nil, err
	}
	return txn, nil
}

func (blockManagerUtils BlockManagerUtils) Propose(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, medians []uint32, iteration *big.Int, biggestInfluencerId uint32) (*Types.Transaction, error) {
	blockManager := utils.UtilsInterface.GetBlockManager(client)
	var (
		txn *Types.Transaction
		err error
	)
	err = retry.Do(func() error {
		txn, err = blockManager.Propose(opts, epoch, medians, iteration, biggestInfluencerId)
		if err != nil {
			log.Error("Error in proposing... Retrying")
			return err
		}
		return nil
	}, retry.Attempts(3))
	if err != nil {
		return nil, err
	}
	return txn, nil
}

func (voteManagerUtils VoteManagerUtils) Reveal(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, values []*big.Int, secret [32]byte) (*Types.Transaction, error) {
	voteManager := utils.UtilsInterface.GetVoteManager(client)
	var (
		txn *Types.Transaction
		err error
	)
	err = retry.Do(func() error {
		txn, err = voteManager.Reveal(opts, epoch, values, secret)
		if err != nil {
			log.Error("Error in revealing... Retrying")
			return err
		}
		return nil
	}, retry.Attempts(3))
	if err != nil {
		return nil, err
	}
	return txn, nil
}

func (voteManagerUtils VoteManagerUtils) Commit(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, commitment [32]byte) (*Types.Transaction, error) {
	voteManager := utils.UtilsInterface.GetVoteManager(client)
	var (
		txn *Types.Transaction
		err error
	)
	err = retry.Do(func() error {
		txn, err = voteManager.Commit(opts, epoch, commitment)
		if err != nil {
			log.Error("Error in committing... Retrying")
			return err
		}
		return nil
	}, retry.Attempts(3))
	if err != nil {
		return nil, err
	}
	return txn, nil
}

func (tokenManagerUtils TokenManagerUtils) Allowance(client *ethclient.Client, opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	tokenManager := utils.GetTokenManager(client)
	return tokenManager.Allowance(opts, owner, spender)
}

func (tokenManagerUtils TokenManagerUtils) Approve(client *ethclient.Client, opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*Types.Transaction, error) {
	tokenManager := utils.GetTokenManager(client)
	return tokenManager.Approve(opts, spender, amount)
}

func (tokenManagerUtils TokenManagerUtils) Transfer(client *ethclient.Client, opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*Types.Transaction, error) {
	tokenManager := utils.GetTokenManager(client)
	return tokenManager.Transfer(opts, recipient, amount)
}

func (assetManagerUtils AssetManagerUtils) CreateJob(client *ethclient.Client, opts *bind.TransactOpts, weight uint8, power int8, selectorType uint8, name string, selector string, url string) (*Types.Transaction, error) {
	assetManager := utils.UtilsInterface.GetAssetManager(client)
	return assetManager.CreateJob(opts, weight, power, selectorType, name, selector, url)
}

func (assetManagerUtils AssetManagerUtils) SetCollectionStatus(client *ethclient.Client, opts *bind.TransactOpts, assetStatus bool, id uint16) (*Types.Transaction, error) {
	assetManager := utils.UtilsInterface.GetAssetManager(client)
	return assetManager.SetCollectionStatus(opts, assetStatus, id)
}

func (assetManagerUtils AssetManagerUtils) GetActiveStatus(client *ethclient.Client, opts *bind.CallOpts, id uint16) (bool, error) {
	assetMananger := utils.UtilsInterface.GetAssetManager(client)
	return assetMananger.GetCollectionStatus(opts, id)
}

func (assetManagerUtils AssetManagerUtils) UpdateJob(client *ethclient.Client, opts *bind.TransactOpts, jobId uint16, weight uint8, power int8, selectorType uint8, selector string, url string) (*Types.Transaction, error) {
	assetManager := utils.UtilsInterface.GetAssetManager(client)
	return assetManager.UpdateJob(opts, jobId, weight, power, selectorType, selector, url)
}

func (assetManagerUtils AssetManagerUtils) CreateCollection(client *ethclient.Client, opts *bind.TransactOpts, tolerance uint16, power int8, aggregationMethod uint32, jobIDs []uint16, name string) (*Types.Transaction, error) {
	assetManager := utils.UtilsInterface.GetAssetManager(client)
	return assetManager.CreateCollection(opts, tolerance, power, aggregationMethod, jobIDs, name)
}

func (assetManagerUtils AssetManagerUtils) UpdateCollection(client *ethclient.Client, opts *bind.TransactOpts, collectionId uint16, tolerance uint16, aggregationMethod uint32, power int8, jobIds []uint16) (*Types.Transaction, error) {
	assetManager := utils.UtilsInterface.GetAssetManager(client)
	return assetManager.UpdateCollection(opts, collectionId, tolerance, aggregationMethod, power, jobIds)
}

func (flagSetUtils FLagSetUtils) GetStringProvider(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("provider")
}

func (flagSetUtils FLagSetUtils) GetFloat32GasMultiplier(flagSet *pflag.FlagSet) (float32, error) {
	return flagSet.GetFloat32("gasmultiplier")
}

func (flagSetUtils FLagSetUtils) GetInt32Buffer(flagSet *pflag.FlagSet) (int32, error) {
	return flagSet.GetInt32("buffer")
}

func (flagSetUtils FLagSetUtils) GetInt32Wait(flagSet *pflag.FlagSet) (int32, error) {
	return flagSet.GetInt32("wait")
}

func (flagSetUtils FLagSetUtils) GetInt32GasPrice(flagSet *pflag.FlagSet) (int32, error) {
	return flagSet.GetInt32("gasprice")
}

func (flagSetUtils FLagSetUtils) GetStringLogLevel(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("logLevel")
}

func (flagSetUtils FLagSetUtils) GetFloat32GasLimit(flagSet *pflag.FlagSet) (float32, error) {
	return flagSet.GetFloat32("gasLimit")
}

func (flagSetUtils FLagSetUtils) GetBoolAutoWithdraw(flagSet *pflag.FlagSet) (bool, error) {
	return flagSet.GetBool("autoWithdraw")
}

func (flagSetUtils FLagSetUtils) GetUint32BountyId(flagSet *pflag.FlagSet) (uint32, error) {
	return flagSet.GetUint32("bountyId")
}

func (flagSetUtils FLagSetUtils) GetRootStringProvider() (string, error) {
	return rootCmd.PersistentFlags().GetString("provider")
}

func (flagSetUtils FLagSetUtils) GetRootFloat32GasMultiplier() (float32, error) {
	return rootCmd.PersistentFlags().GetFloat32("gasmultiplier")
}

func (flagSetUtils FLagSetUtils) GetRootInt32Buffer() (int32, error) {
	return rootCmd.PersistentFlags().GetInt32("buffer")
}

func (flagSetUtils FLagSetUtils) GetRootInt32Wait() (int32, error) {
	return rootCmd.PersistentFlags().GetInt32("wait")
}

func (flagSetUtils FLagSetUtils) GetRootInt32GasPrice() (int32, error) {
	return rootCmd.PersistentFlags().GetInt32("gasprice")
}

func (flagSetUtils FLagSetUtils) GetRootStringLogLevel() (string, error) {
	return rootCmd.PersistentFlags().GetString("logLevel")
}

func (flagSetUtils FLagSetUtils) GetRootFloat32GasLimit() (float32, error) {
	return rootCmd.PersistentFlags().GetFloat32("gasLimit")
}

func (flagSetUtils FLagSetUtils) GetStringFrom(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("from")
}

func (flagSetUtils FLagSetUtils) GetStringTo(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("to")
}

func (flagSetUtils FLagSetUtils) GetStringAddress(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("address")
}

func (flagSetUtils FLagSetUtils) GetUint32StakerId(flagSet *pflag.FlagSet) (uint32, error) {
	return flagSet.GetUint32("stakerId")
}

func (flagSetUtils FLagSetUtils) GetStringName(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("name")
}

func (flagSetUtils FLagSetUtils) GetStringUrl(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("url")
}

func (flagSetUtils FLagSetUtils) GetStringSelector(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("selector")
}

func (flagSetUtils FLagSetUtils) GetInt8Power(flagSet *pflag.FlagSet) (int8, error) {
	return flagSet.GetInt8("power")
}

func (flagSetUtils FLagSetUtils) GetUint8Weight(flagSet *pflag.FlagSet) (uint8, error) {
	return flagSet.GetUint8("weight")
}

func (flagSetUtils FLagSetUtils) GetUint16AssetId(flagSet *pflag.FlagSet) (uint16, error) {
	return flagSet.GetUint16("assetId")
}

func (flagSetUtils FLagSetUtils) GetUint8SelectorType(flagSet *pflag.FlagSet) (uint8, error) {
	return flagSet.GetUint8("selectorType")
}

func (flagSetUtils FLagSetUtils) GetStringStatus(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("status")
}

func (flagSetUtils FLagSetUtils) GetUint8Commission(flagSet *pflag.FlagSet) (uint8, error) {
	return flagSet.GetUint8("commission")
}

func (flagSetUtils FLagSetUtils) GetUintSliceJobIds(flagSet *pflag.FlagSet) ([]uint, error) {
	return flagSet.GetUintSlice("jobIds")
}

func (flagSetUtils FLagSetUtils) GetUint32Aggregation(flagSet *pflag.FlagSet) (uint32, error) {
	return flagSet.GetUint32("aggregation")
}

func (flagSetUtils FLagSetUtils) GetUint16JobId(flagSet *pflag.FlagSet) (uint16, error) {
	return flagSet.GetUint16("jobId")
}

func (flagSetUtils FLagSetUtils) GetUint16CollectionId(flagSet *pflag.FlagSet) (uint16, error) {
	return flagSet.GetUint16("collectionId")
}

func (flagSetUtils FLagSetUtils) GetStringValue(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("value")
}

func (flagSetUtils FLagSetUtils) GetStringPow(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("pow")
}

func (flagSetUtils FLagSetUtils) GetUint16Tolerance(flagSet *pflag.FlagSet) (uint16, error) {
	return flagSet.GetUint16("tolerance")
}

func (flagSetUtils FLagSetUtils) GetBoolAutoVote(flagSet *pflag.FlagSet) (bool, error) {
	return flagSet.GetBool("autoVote")
}

func (flagSetUtils FLagSetUtils) GetBoolRogue(flagSet *pflag.FlagSet) (bool, error) {
	return flagSet.GetBool("rogue")
}

func (flagSetUtils FLagSetUtils) GetStringSliceRogueMode(flagSet *pflag.FlagSet) ([]string, error) {
	return flagSet.GetStringSlice("rogueMode")
}

func (KeystoreUtils KeystoreUtils) Accounts(path string) []ethAccounts.Account {
	ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	return ks.Accounts()
}

func (keystoreUtils KeystoreUtils) ImportECDSA(path string, priv *ecdsa.PrivateKey, passphrase string) (ethAccounts.Account, error) {
	ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	return ks.ImportECDSA(priv, passphrase)
}

func (c CryptoUtils) HexToECDSA(hexKey string) (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(hexKey)
}

func (*UtilsStruct) GiveSorted(client *ethclient.Client, blockManager *bindings.BlockManager, txnOpts *bind.TransactOpts, epoch uint32, assetId uint16, sortedStakers []uint32) {
	GiveSorted(client, blockManager, txnOpts, epoch, assetId, sortedStakers)
}
