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

func (u UtilsMockery) GetConfigFilePath() (string, error) {
	return path.GetConfigFilePath()
}

func (u UtilsMockery) ViperWriteConfigAs(path string) error {
	return viper.WriteConfigAs(path)
}

func (u UtilsMockery) GetEpoch(client *ethclient.Client) (uint32, error) {
	return utils.GetEpoch(client)
}

func (u UtilsMockery) GetUpdatedEpoch(client *ethclient.Client) (uint32, error) {
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

func (u UtilsMockery) GetNumActiveAssets(client *ethclient.Client) (*big.Int, error) {
	return utils.GetNumActiveAssets(client)
}

func (u UtilsMockery) GetRogueRandomValue(value int) *big.Int {
	return utils.GetRogueRandomValue(value)
}

func (u UtilsMockery) GetActiveAssetsData(client *ethclient.Client, epoch uint32) ([]*big.Int, error) {
	return utils.GetActiveAssetsData(client, epoch)
}

func (u UtilsMockery) GetDelayedState(client *ethclient.Client, buffer int32) (int64, error) {
	return utils.GetDelayedState(client, buffer)
}

func (u UtilsMockery) GetDefaultPath() (string, error) {
	return path.GetDefaultPath()
}

func (u UtilsMockery) FetchBalance(client *ethclient.Client, accountAddress string) (*big.Int, error) {
	return utils.FetchBalance(client, accountAddress)
}

func (u UtilsMockery) IsFlagPassed(name string) bool {
	return utils.IsFlagPassed(name)
}

func (u UtilsMockery) GetFractionalAmountInWei(amount *big.Int, power string) (*big.Int, error) {
	return utils.GetFractionalAmountInWei(amount, power)
}

func (u UtilsMockery) GetAmountInWei(amount *big.Int) *big.Int {
	return utils.GetAmountInWei(amount)
}

func (u UtilsMockery) CheckAmountAndBalance(amountInWei *big.Int, balance *big.Int) *big.Int {
	return utils.CheckAmountAndBalance(amountInWei, balance)
}

func (u UtilsMockery) GetAmountInDecimal(amountInWei *big.Int) *big.Float {
	return utils.GetAmountInDecimal(amountInWei)
}

func (u UtilsMockery) GetEpochLastCommitted(client *ethclient.Client, stakerId uint32) (uint32, error) {
	return utils.GetEpochLastCommitted(client, stakerId)
}

func (u UtilsMockery) GetCommitments(client *ethclient.Client, address string) ([32]byte, error) {
	return utils.GetCommitments(client, address)
}

func (u UtilsMockery) AllZero(bytesValue [32]byte) bool {
	return utils.AllZero(bytesValue)
}

func (u UtilsMockery) ConvertUintArrayToUint16Array(uintArr []uint) []uint16 {
	return utils.ConvertUintArrayToUint16Array(uintArr)
}

func (u UtilsMockery) GetStateName(state int64) string {
	return utils.GetStateName(state)
}

func (u UtilsMockery) GetJobs(client *ethclient.Client) ([]bindings.StructsJob, error) {
	return utils.GetJobs(client)
}

func (u UtilsMockery) CheckEthBalanceIsZero(client *ethclient.Client, address string) {
	utils.CheckEthBalanceIsZero(client, address)
}

func (u UtilsMockery) AssignStakerId(flagSet *pflag.FlagSet, client *ethclient.Client, address string) (uint32, error) {
	return utils.AssignStakerId(flagSet, client, address)
}

func (u UtilsMockery) GetLock(client *ethclient.Client, address string, stakerId uint32) (types.Locks, error) {
	return utils.GetLock(client, address, stakerId)
}

func (u UtilsMockery) GetStaker(client *ethclient.Client, address string, stakerId uint32) (bindings.StructsStaker, error) {
	return utils.GetStaker(client, address, stakerId)
}

func (u UtilsMockery) GetUpdatedStaker(client *ethclient.Client, address string, stakerId uint32) (bindings.StructsStaker, error) {
	return utils.GetStaker(client, address, stakerId)
}

func (u UtilsMockery) GetStakedToken(client *ethclient.Client, address common.Address) *bindings.StakedToken {
	return utils.GetStakedToken(client, address)
}

func (u UtilsMockery) ConvertSRZRToRZR(sAmount *big.Int, currentStake *big.Int, totalSupply *big.Int) *big.Int {
	return utils.ConvertSRZRToRZR(sAmount, currentStake, totalSupply)
}

func (u UtilsMockery) ConvertRZRToSRZR(sAmount *big.Int, currentStake *big.Int, totalSupply *big.Int) (*big.Int, error) {
	return utils.ConvertRZRToSRZR(sAmount, currentStake, totalSupply)
}

func (u UtilsMockery) GetWithdrawReleasePeriod(client *ethclient.Client, address string) (uint8, error) {
	return utils.GetWithdrawReleasePeriod(client, address)
}

func (u UtilsMockery) GetInfluenceSnapshot(client *ethclient.Client, stakerId uint32, epoch uint32) (*big.Int, error) {
	return utils.GetInfluenceSnapshot(client, stakerId, epoch)
}

func (u UtilsMockery) GetCollections(client *ethclient.Client) ([]bindings.StructsCollection, error) {
	return utils.GetCollections(client)
}

func (u UtilsMockery) GetNumberOfStakers(client *ethclient.Client, address string) (uint32, error) {
	return utils.GetNumberOfStakers(client, address)
}

func (u UtilsMockery) GetRandaoHash(client *ethclient.Client) ([32]byte, error) {
	return utils.GetRandaoHash(client)
}

func (u UtilsMockery) GetNumberOfProposedBlocks(client *ethclient.Client, epoch uint32) (uint8, error) {
	return utils.UtilsInterface.GetNumberOfProposedBlocks(client, epoch)
}

func (u UtilsMockery) GetMaxAltBlocks(client *ethclient.Client) (uint8, error) {
	return utils.UtilsInterface.GetMaxAltBlocks(client)
}

func (u UtilsMockery) GetProposedBlock(client *ethclient.Client, epoch uint32, proposedBlockId uint32) (bindings.StructsBlock, error) {
	return utils.UtilsInterface.GetProposedBlock(client, epoch, proposedBlockId)
}

func (u UtilsMockery) GetEpochLastRevealed(client *ethclient.Client, address string, stakerId uint32) (uint32, error) {
	return utils.GetEpochLastRevealed(client, address, stakerId)
}

func (u UtilsMockery) GetVoteValue(client *ethclient.Client, assetId uint16, stakerId uint32) (*big.Int, error) {
	return utils.GetVoteValue(client, assetId, stakerId)
}

func (u UtilsMockery) GetTotalInfluenceRevealed(client *ethclient.Client, epoch uint32) (*big.Int, error) {
	return utils.GetTotalInfluenceRevealed(client, epoch)
}

func (u UtilsMockery) ConvertBigIntArrayToUint32Array(bigIntArray []*big.Int) []uint32 {
	return utils.ConvertBigIntArrayToUint32Array(bigIntArray)
}

func (u UtilsMockery) GetActiveAssetIds(client *ethclient.Client) ([]uint16, error) {
	return utils.GetActiveAssetIds(client)
}

func (u UtilsMockery) GetBlockManager(client *ethclient.Client) *bindings.BlockManager {
	return utils.UtilsInterface.GetBlockManager(client)
}

func (u UtilsMockery) GetVotes(client *ethclient.Client, stakerId uint32) (bindings.StructsVote, error) {
	return utils.GetVotes(client, stakerId)
}

func (u UtilsMockery) GetSortedProposedBlockIds(client *ethclient.Client, epoch uint32) ([]uint32, error) {
	return utils.UtilsInterface.GetSortedProposedBlockIds(client, epoch)
}

func (u UtilsMockery) ParseBool(str string) (bool, error) {
	return strconv.ParseBool(str)
}

func (u UtilsMockery) GetStakerId(client *ethclient.Client, address string) (uint32, error) {
	return utils.GetStakerId(client, address)
}

func (u UtilsMockery) PrivateKeyPrompt() string {
	return utils.PrivateKeyPrompt()
}

func (u UtilsMockery) PasswordPrompt() string {
	return utils.PasswordPrompt()
}

func (u UtilsMockery) GetMaxCommission(client *ethclient.Client) (uint8, error) {
	return utils.GetMaxCommission(client)
}

func (u UtilsMockery) GetEpochLimitForUpdateCommission(client *ethclient.Client) (uint16, error) {
	return utils.GetEpochLimitForUpdateCommission(client)
}

func (transactionUtils TransactionUtilsMockery) Hash(txn *Types.Transaction) common.Hash {
	return txn.Hash()
}

func (stakeManagerUtils StakeManagerUtilsMockery) Stake(client *ethclient.Client, txnOpts *bind.TransactOpts, epoch uint32, amount *big.Int) (*Types.Transaction, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.Stake(txnOpts, epoch, amount)
}

func (stakeManagerUtils StakeManagerUtilsMockery) ExtendLock(client *ethclient.Client, opts *bind.TransactOpts, stakerId uint32) (*Types.Transaction, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.ExtendLock(opts, stakerId)
}

func (stakeManagerUtils StakeManagerUtilsMockery) Delegate(client *ethclient.Client, opts *bind.TransactOpts, stakerId uint32, amount *big.Int) (*Types.Transaction, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.Delegate(opts, stakerId, amount)
}

func (stakeManagerUtils StakeManagerUtilsMockery) Withdraw(client *ethclient.Client, opts *bind.TransactOpts, stakerId uint32) (*Types.Transaction, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.Withdraw(opts, stakerId)
}

func (stakeManagerUtils StakeManagerUtilsMockery) SetDelegationAcceptance(client *ethclient.Client, opts *bind.TransactOpts, status bool) (*Types.Transaction, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.SetDelegationAcceptance(opts, status)
}

func (stakeManagerUtils StakeManagerUtilsMockery) UpdateCommission(client *ethclient.Client, opts *bind.TransactOpts, commission uint8) (*Types.Transaction, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.UpdateCommission(opts, commission)
}

func (stakeManagerUtils StakeManagerUtilsMockery) Unstake(client *ethclient.Client, opts *bind.TransactOpts, stakerId uint32, sAmount *big.Int) (*Types.Transaction, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.Unstake(opts, stakerId, sAmount)
}

func (stakeManagerUtils StakeManagerUtilsMockery) RedeemBounty(client *ethclient.Client, opts *bind.TransactOpts, bountyId uint32) (*Types.Transaction, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.RedeemBounty(opts, bountyId)
}

func (stakeManagerUtils StakeManagerUtilsMockery) StakerInfo(client *ethclient.Client, opts *bind.CallOpts, stakerId uint32) (types.Staker, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.Stakers(opts, stakerId)
}

func (stakeManagerUtils StakeManagerUtilsMockery) GetMaturity(client *ethclient.Client, opts *bind.CallOpts, age uint32) (uint16, error) {
	stakeManager := utils.GetStakeManager(client)
	index := age / 10000
	return stakeManager.Maturities(opts, big.NewInt(int64(index)))
}

func (stakeManagerUtils StakeManagerUtilsMockery) GetBountyLock(client *ethclient.Client, opts *bind.CallOpts, bountyId uint32) (types.BountyLock, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.BountyLocks(opts, bountyId)
}

func (stakeManagerUtils StakeManagerUtilsMockery) BalanceOf(stakedToken *bindings.StakedToken, callOpts *bind.CallOpts, address common.Address) (*big.Int, error) {
	return stakedToken.BalanceOf(callOpts, address)
}

func (stakeManagerUtils StakeManagerUtilsMockery) GetTotalSupply(token *bindings.StakedToken, callOpts *bind.CallOpts) (*big.Int, error) {
	return token.TotalSupply(callOpts)
}

func (blockManagerUtils BlockManagerUtilsMockery) ClaimBlockReward(client *ethclient.Client, opts *bind.TransactOpts) (*Types.Transaction, error) {
	blockManager := utils.UtilsInterface.GetBlockManager(client)
	return blockManager.ClaimBlockReward(opts)
}

func (blockManagerUtils BlockManagerUtilsMockery) FinalizeDispute(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, blockIndex uint8) (*Types.Transaction, error) {
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

func (blockManagerUtils BlockManagerUtilsMockery) DisputeBiggestInfluenceProposed(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, blockIndex uint8, correctBiggestInfluencerId uint32) (*Types.Transaction, error) {
	blockManager := utils.UtilsInterface.GetBlockManager(client)
	var (
		txn *Types.Transaction
		err error
	)
	err = retry.Do(func() error {
		txn, err = blockManager.DisputeBiggestInfluenceProposed(opts, epoch, blockIndex, correctBiggestInfluencerId)
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

func (blockManagerUtils BlockManagerUtilsMockery) Propose(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, medians []uint32, iteration *big.Int, biggestInfluencerId uint32) (*Types.Transaction, error) {
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

func (voteManagerUtils VoteManagerUtilsMockery) Reveal(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, values []*big.Int, secret [32]byte) (*Types.Transaction, error) {
	voteManager := utils.GetVoteManager(client)
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

func (voteManagerUtils VoteManagerUtilsMockery) Commit(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, commitment [32]byte) (*Types.Transaction, error) {
	voteManager := utils.GetVoteManager(client)
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

func (tokenManagerUtils TokenManagerUtilsMockery) Allowance(client *ethclient.Client, opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	tokenManager := utils.GetTokenManager(client)
	return tokenManager.Allowance(opts, owner, spender)
}

func (tokenManagerUtils TokenManagerUtilsMockery) Approve(client *ethclient.Client, opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*Types.Transaction, error) {
	tokenManager := utils.GetTokenManager(client)
	return tokenManager.Approve(opts, spender, amount)
}

func (tokenManagerUtils TokenManagerUtilsMockery) Transfer(client *ethclient.Client, opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*Types.Transaction, error) {
	tokenManager := utils.GetTokenManager(client)
	return tokenManager.Transfer(opts, recipient, amount)
}

func (assetManagerUtils AssetManagerUtilsMockery) CreateJob(client *ethclient.Client, opts *bind.TransactOpts, weight uint8, power int8, selectorType uint8, name string, selector string, url string) (*Types.Transaction, error) {
	assetManager := utils.GetAssetManager(client)
	return assetManager.CreateJob(opts, weight, power, selectorType, name, selector, url)
}

func (assetManagerUtils AssetManagerUtilsMockery) SetCollectionStatus(client *ethclient.Client, opts *bind.TransactOpts, assetStatus bool, id uint16) (*Types.Transaction, error) {
	assetManager := utils.GetAssetManager(client)
	return assetManager.SetCollectionStatus(opts, assetStatus, id)
}

func (assetManagerUtils AssetManagerUtilsMockery) GetActiveStatus(client *ethclient.Client, opts *bind.CallOpts, id uint16) (bool, error) {
	assetMananger := utils.GetAssetManager(client)
	return assetMananger.GetCollectionStatus(opts, id)
}

func (assetManagerUtils AssetManagerUtilsMockery) UpdateJob(client *ethclient.Client, opts *bind.TransactOpts, jobId uint16, weight uint8, power int8, selectorType uint8, selector string, url string) (*Types.Transaction, error) {
	assetManager := utils.GetAssetManager(client)
	return assetManager.UpdateJob(opts, jobId, weight, power, selectorType, selector, url)
}

func (assetManagerUtils AssetManagerUtilsMockery) CreateCollection(client *ethclient.Client, opts *bind.TransactOpts, jobIDs []uint16, aggregationMethod uint32, power int8, name string) (*Types.Transaction, error) {
	assetManager := utils.GetAssetManager(client)
	return assetManager.CreateCollection(opts, jobIDs, aggregationMethod, power, name)
}

func (assetManagerUtils AssetManagerUtilsMockery) UpdateCollection(client *ethclient.Client, opts *bind.TransactOpts, collectionId uint16, aggregationMethod uint32, power int8, jobIds []uint16) (*Types.Transaction, error) {
	assetManager := utils.GetAssetManager(client)
	return assetManager.UpdateCollection(opts, collectionId, aggregationMethod, power, jobIds)
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

func (flagSetUtils FLagSetUtilsMockery) GetStringValue(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("value")
}

func (flagSetUtils FLagSetUtilsMockery) GetStringPow(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("pow")
}

func (KeystoreUtils KeystoreUtilsMockery) Accounts(path string) []ethAccounts.Account {
	ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	return ks.Accounts()
}

func (keystoreUtils KeystoreUtilsMockery) ImportECDSA(path string, priv *ecdsa.PrivateKey, passphrase string) (ethAccounts.Account, error) {
	ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	return ks.ImportECDSA(priv, passphrase)
}

func (c CryptoUtils) HexToECDSA(hexKey string) (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(hexKey)
}

func (*UtilsStructMockery) GiveSorted(client *ethclient.Client, blockManager *bindings.BlockManager, txnOpts *bind.TransactOpts, epoch uint32, assetId uint16, sortedStakers []uint32) {
	GiveSorted(client, blockManager, txnOpts, epoch, assetId, sortedStakers)
}
