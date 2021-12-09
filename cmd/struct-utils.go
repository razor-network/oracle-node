package cmd

import (
	"crypto/ecdsa"
	"github.com/avast/retry-go"
	"math/big"
	"razor/accounts"
	"razor/core/types"
	"razor/path"
	"razor/pkg/bindings"
	"razor/utils"
	"strconv"
	"time"

	ethAccounts "github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Utils struct{}
type TokenManagerUtils struct{}
type TransactionUtils struct{}
type StakeManagerUtils struct{}
type AssetManagerUtils struct{}
type KeystoreUtils struct{}
type FlagSetUtils struct{}
type ProposeUtils struct{}
type UtilsCmd struct{}
type VoteManagerUtils struct{}
type BlockManagerUtils struct{}
type CryptoUtils struct{}
type UtilsStruct struct {
	razorUtils        utilsInterface
	cmdUtils          utilsCmdInterface
	proposeUtils      proposeUtilsInterface
	blockManagerUtils blockManagerInterface
	stakeManagerUtils stakeManagerInterface
	transactionUtils  transactionInterface
	assetManagerUtils assetManagerInterface
	voteManagerUtils  voteManagerInterface
	tokenManagerUtils tokenManagerInterface
	keystoreUtils     keystoreInterface
	flagSetUtils      flagSetInterface
	cryptoUtils       cryptoInterface
	accountUtils      accounts.AccountInterface
	packageUtils      utils.Utils
}

func (u Utils) ConnectToClient(provider string) *ethclient.Client {
	return utils.ConnectToClient(provider)
}

func (u Utils) GetOptions() bind.CallOpts {
	return utils.GetOptions()
}

func (u Utils) GetTxnOpts(transactionData types.TransactionOptions, packageUtils utils.Utils) *bind.TransactOpts {
	return utils.GetTxnOpts(transactionData, packageUtils)
}

func (u Utils) WaitForBlockCompletion(client *ethclient.Client, hashToRead string) int {
	return utils.WaitForBlockCompletion(client, hashToRead)
}

func (u Utils) AssignPassword(flagSet *pflag.FlagSet) string {
	return utils.AssignPassword(flagSet)
}

func (u Utils) FetchBalance(client *ethclient.Client, accountAddress string) (*big.Int, error) {
	return utils.FetchBalance(client, accountAddress)
}

func (u Utils) CheckAmountAndBalance(amountInWei *big.Int, balance *big.Int) *big.Int {
	return utils.CheckAmountAndBalance(amountInWei, balance)
}

func (u Utils) GetAmountInDecimal(amountInWei *big.Int) *big.Float {
	return utils.GetAmountInDecimal(amountInWei)
}

func (u Utils) GetStakerId(client *ethclient.Client, address string) (uint32, error) {
	return utils.GetStakerId(client, address)
}

func (u Utils) GetStaker(client *ethclient.Client, address string, stakerId uint32) (bindings.StructsStaker, error) {
	return utils.GetStaker(client, address, stakerId)
}

func (u Utils) GetUpdatedStaker(client *ethclient.Client, address string, stakerId uint32) (bindings.StructsStaker, error) {
	return utils.GetStaker(client, address, stakerId)
}

func (u Utils) GetConfigData(utilsStruct UtilsStruct) (types.Configurations, error) {
	return GetConfigData(utilsStruct)
}

func (u Utils) ParseBool(str string) (bool, error) {
	return strconv.ParseBool(str)
}

func (u Utils) GetDelayedState(client *ethclient.Client, buffer int32, packageUtils utils.Utils) (int64, error) {
	return utils.GetDelayedState(client, buffer, packageUtils)
}

func (u Utils) GetEpoch(client *ethclient.Client, packageUtils utils.Utils) (uint32, error) {
	return utils.GetEpoch(client, packageUtils)
}

func (u Utils) GetActiveAssetsData(client *ethclient.Client, address string, epoch uint32) ([]*big.Int, error) {
	return utils.GetActiveAssetsData(client, address, epoch)
}

func (u Utils) ConvertUintArrayToUint8Array(uintArr []uint) []uint8 {
	return utils.ConvertUintArrayToUint8Array(uintArr)
}

func (u Utils) PrivateKeyPrompt() string {
	return utils.PrivateKeyPrompt()
}

func (u Utils) PasswordPrompt() string {
	return utils.PasswordPrompt()
}

func (u Utils) GetDefaultPath() (string, error) {
	return path.GetDefaultPath()
}

func (u Utils) GetNumberOfStakers(client *ethclient.Client, address string) (uint32, error) {
	return utils.GetNumberOfStakers(client, address)
}

func (u Utils) GetRandaoHash(client *ethclient.Client) ([32]byte, error) {
	return utils.GetRandaoHash(client)
}

func (u Utils) GetNumberOfProposedBlocks(client *ethclient.Client, address string, epoch uint32) (uint8, error) {
	return utils.GetNumberOfProposedBlocks(client, address, epoch)
}

func (u Utils) GetMaxAltBlocks(client *ethclient.Client, address string) (uint8, error) {
	return utils.GetMaxAltBlocks(client, address)
}

func (u Utils) GetProposedBlock(client *ethclient.Client, address string, epoch uint32, proposedBlockId uint8) (bindings.StructsBlock, error) {
	return utils.GetProposedBlock(client, address, epoch, proposedBlockId)
}

func (u Utils) GetEpochLastRevealed(client *ethclient.Client, address string, stakerId uint32) (uint32, error) {
	return utils.GetEpochLastRevealed(client, address, stakerId)
}

func (u Utils) GetVoteValue(client *ethclient.Client, assetId uint8, stakerId uint32) (*big.Int, error) {
	return utils.GetVoteValue(client, assetId, stakerId)
}

func (u Utils) GetInfluenceSnapshot(client *ethclient.Client, stakerId uint32, epoch uint32) (*big.Int, error) {
	return utils.GetInfluenceSnapshot(client, stakerId, epoch)
}

func (u Utils) GetNumActiveAssets(client *ethclient.Client, address string) (*big.Int, error) {
	return utils.GetNumActiveAssets(client, address)
}

func (u Utils) GetTotalInfluenceRevealed(client *ethclient.Client, epoch uint32) (*big.Int, error) {
	return utils.GetTotalInfluenceRevealed(client, epoch)
}

func (u Utils) ConvertBigIntArrayToUint32Array(bigIntArray []*big.Int) []uint32 {
	return utils.ConvertBigIntArrayToUint32Array(bigIntArray)
}

func (u Utils) GetLock(client *ethclient.Client, address string, stakerId uint32) (types.Locks, error) {
	return utils.GetLock(client, address, stakerId)
}

func (u Utils) GetWithdrawReleasePeriod(client *ethclient.Client, address string) (uint8, error) {
	return utils.GetWithdrawReleasePeriod(client, address)
}

func (u Utils) GetCommitments(client *ethclient.Client, address string) ([32]byte, error) {
	return utils.GetCommitments(client, address)
}

func (u Utils) AllZero(bytesValue [32]byte) bool {
	return utils.AllZero(bytesValue)
}

func (u Utils) GetEpochLastCommitted(client *ethclient.Client, stakerId uint32) (uint32, error) {
	return utils.GetEpochLastCommitted(client, stakerId)
}

func (u Utils) GetConfigFilePath() (string, error) {
	return path.GetConfigFilePath()
}

func (u Utils) ViperWriteConfigAs(path string) error {
	return viper.WriteConfigAs(path)
}

func (u Utils) IsEqual(arr1 []uint32, arr2 []uint32) (bool, int) {
	return utils.IsEqual(arr1, arr2)
}

func (u Utils) GetActiveAssetIds(client *ethclient.Client, address string) ([]uint8, error) {
	return utils.GetActiveAssetIds(client, address)
}

func (u Utils) GetBlockManager(client *ethclient.Client) *bindings.BlockManager {
	return utils.GetBlockManager(client)
}

func (u Utils) GetVotes(client *ethclient.Client, stakerId uint32) (bindings.StructsVote, error) {
	return utils.GetVotes(client, stakerId)
}

func (u Utils) Contains(arr []int, val int) bool {
	return utils.Contains(arr, val)
}

func (u Utils) CheckEthBalanceIsZero(client *ethclient.Client, address string) {
	utils.CheckEthBalanceIsZero(client, address)
}

func (u Utils) AssignStakerId(flagSet *pflag.FlagSet, client *ethclient.Client, address string) (uint32, error) {
	return utils.AssignStakerId(flagSet, client, address)
}

func (u Utils) GetLatestBlock(client *ethclient.Client, packageUtils utils.Utils) (*Types.Header, error) {
	return utils.GetLatestBlockWithRetry(client, packageUtils)
}

func (u Utils) GetSortedProposedBlockIds(client *ethclient.Client, address string, epoch uint32) ([]uint8, error) {
	return utils.GetSortedProposedBlockIds(client, address, epoch)
}

func (u Utils) CheckError(msg string, err error) {
	utils.CheckError(msg, err)
}

func (u Utils) GetUpdatedEpoch(client *ethclient.Client, packageUtils utils.Utils) (uint32, error) {
	return utils.GetEpoch(client, packageUtils)
}

func (u Utils) GetStateName(stateNumber int64) string {
	return utils.GetStateName(stateNumber)
}

func (u Utils) getBufferPercent(utilsStruct UtilsStruct) (int32, error) {
	return getBufferPercent(utilsStruct)
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

func (u Utils) Sleep(duration time.Duration) {
	utils.Sleep(duration)
}

func (u Utils) getProvider(utilsStruct UtilsStruct) (string, error) {
	return getProvider(utilsStruct)
}

func (u Utils) getMultiplier(utilsStruct UtilsStruct) (float32, error) {
	return getMultiplier(utilsStruct)
}

func (u Utils) getWaitTime(utilsStruct UtilsStruct) (int32, error) {
	return getWaitTime(utilsStruct)
}

func (u Utils) getGasPrice(utilsStruct UtilsStruct) (int32, error) {
	return getGasPrice(utilsStruct)
}

func (u Utils) getLogLevel(utilsStruct UtilsStruct) (string, error) {
	return getLogLevel(utilsStruct)
}

func (u Utils) getGasLimit(utilsStruct UtilsStruct) (float32, error) {
	return getGasLimit(utilsStruct)
}

func (u Utils) CalculateBlockTime(client *ethclient.Client, packageUtils utils.Utils) int64 {
	return utils.CalculateBlockTime(client, packageUtils)
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

func (stakeManagerUtils StakeManagerUtils) Delegate(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, stakerId uint32, amount *big.Int) (*Types.Transaction, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.Delegate(opts, epoch, stakerId, amount)
}

func (stakeManagerUtils StakeManagerUtils) Withdraw(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, stakerId uint32) (*Types.Transaction, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.Withdraw(opts, epoch, stakerId)
}

func (stakeManagerUtils StakeManagerUtils) SetDelegationAcceptance(client *ethclient.Client, opts *bind.TransactOpts, status bool) (*Types.Transaction, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.SetDelegationAcceptance(opts, status)
}

func (stakeManagerUtils StakeManagerUtils) SetCommission(client *ethclient.Client, opts *bind.TransactOpts, commission uint8) (*Types.Transaction, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.SetCommission(opts, commission)
}

func (stakeManagerUtils StakeManagerUtils) DecreaseCommission(client *ethclient.Client, opts *bind.TransactOpts, commission uint8) (*Types.Transaction, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.DecreaseCommission(opts, commission)
}

func (stakeManagerUtils StakeManagerUtils) Unstake(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, stakerId uint32, sAmount *big.Int) (*Types.Transaction, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.Unstake(opts, epoch, stakerId, sAmount)
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

func (assetManagerUtils AssetManagerUtils) CreateJob(client *ethclient.Client, opts *bind.TransactOpts, weight uint8, power int8, selectorType uint8, name string, selector string, url string) (*Types.Transaction, error) {
	assetManager := utils.GetAssetManager(client)
	return assetManager.CreateJob(opts, weight, power, selectorType, name, selector, url)
}

func (assetManagerUtils AssetManagerUtils) SetCollectionStatus(client *ethclient.Client, opts *bind.TransactOpts, assetStatus bool, id uint8) (*Types.Transaction, error) {
	assetManager := utils.GetAssetManager(client)
	return assetManager.SetCollectionStatus(opts, assetStatus, id)
}

func (assetManagerUtils AssetManagerUtils) GetActiveStatus(client *ethclient.Client, opts *bind.CallOpts, id uint8) (bool, error) {
	assetMananger := utils.GetAssetManager(client)
	return assetMananger.GetCollectionStatus(opts, id)
}

func (assetManagerUtils AssetManagerUtils) UpdateJob(client *ethclient.Client, opts *bind.TransactOpts, jobId uint8, weight uint8, power int8, selectorType uint8, selector string, url string) (*Types.Transaction, error) {
	assetManager := utils.GetAssetManager(client)
	return assetManager.UpdateJob(opts, jobId, weight, power, selectorType, selector, url)
}

func (assetManagerUtils AssetManagerUtils) CreateCollection(client *ethclient.Client, opts *bind.TransactOpts, jobIDs []uint8, aggregationMethod uint32, power int8, name string) (*Types.Transaction, error) {
	assetManager := utils.GetAssetManager(client)
	return assetManager.CreateCollection(opts, jobIDs, aggregationMethod, power, name)
}

func (assetManagerUtils AssetManagerUtils) UpdateCollection(client *ethclient.Client, opts *bind.TransactOpts, collectionId uint8, aggregationMethod uint32, power int8, jobIds []uint8) (*Types.Transaction, error) {
	assetManager := utils.GetAssetManager(client)
	return assetManager.UpdateCollection(opts, collectionId, aggregationMethod, power, jobIds)
}

func (keystoreUtils KeystoreUtils) Accounts(path string) []ethAccounts.Account {
	ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	return ks.Accounts()
}

func (keystoreUtils KeystoreUtils) ImportECDSA(path string, priv *ecdsa.PrivateKey, passphrase string) (ethAccounts.Account, error) {
	ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	return ks.ImportECDSA(priv, passphrase)
}

func (flagSetUtils FlagSetUtils) GetStringFrom(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("from")
}

func (flagSetUtils FlagSetUtils) GetStringTo(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("to")
}

func (flagSetUtils FlagSetUtils) GetStringAddress(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("address")
}

func (flagSetUtils FlagSetUtils) GetUint32StakerId(flagSet *pflag.FlagSet) (uint32, error) {
	return flagSet.GetUint32("stakerId")
}

func (flagSetUtils FlagSetUtils) GetStringName(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("name")
}

func (flagSetUtils FlagSetUtils) GetStringUrl(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("url")
}

func (flagSetUtils FlagSetUtils) GetStringSelector(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("selector")
}

func (flagSetUtils FlagSetUtils) GetInt8Power(flagSet *pflag.FlagSet) (int8, error) {
	return flagSet.GetInt8("power")
}

func (flagSetUtils FlagSetUtils) GetUint8Weight(flagSet *pflag.FlagSet) (uint8, error) {
	return flagSet.GetUint8("weight")
}

func (flagSetUtils FlagSetUtils) GetUint8AssetId(flagSet *pflag.FlagSet) (uint8, error) {
	return flagSet.GetUint8("assetId")
}

func (flagSetUtils FlagSetUtils) GetUint8SelectorType(flagSet *pflag.FlagSet) (uint8, error) {
	return flagSet.GetUint8("selectorType")
}

func (flagSetUtils FlagSetUtils) GetStringStatus(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("status")
}

func (flagSetUtils FlagSetUtils) GetFloat32GasLimit(flagSet *pflag.FlagSet) (float32, error) {
	return flagSet.GetFloat32("gasLimit")
}

func (voteManagerUtils VoteManagerUtils) Reveal(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, values []*big.Int, secret [32]byte) (*Types.Transaction, error) {
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

func (voteManagerUtils VoteManagerUtils) Commit(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, commitment [32]byte) (*Types.Transaction, error) {
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

func (flagSetUtils FlagSetUtils) GetUint8Commission(flagSet *pflag.FlagSet) (uint8, error) {
	return flagSet.GetUint8("commission")
}

func (flagSetUtils FlagSetUtils) GetUintSliceJobIds(flagSet *pflag.FlagSet) ([]uint, error) {
	return flagSet.GetUintSlice("jobIds")
}

func (flagSetUtils FlagSetUtils) GetUint32Aggregation(flagSet *pflag.FlagSet) (uint32, error) {
	return flagSet.GetUint32("aggregation")
}

func (flagSetUtils FlagSetUtils) GetUint8JobId(flagSet *pflag.FlagSet) (uint8, error) {
	return flagSet.GetUint8("jobId")
}

func (flagSetUtils FlagSetUtils) GetUint8CollectionId(flagSet *pflag.FlagSet) (uint8, error) {
	return flagSet.GetUint8("collectionId")
}

func (flagSetUtils FlagSetUtils) GetStringValue(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("value")
}

func (proposeUtils ProposeUtils) getBiggestInfluenceAndId(client *ethclient.Client, address string, epoch uint32, utilsStruct UtilsStruct) (*big.Int, uint32, error) {
	return getBiggestInfluenceAndId(client, address, epoch, utilsStruct)
}

func (proposeUtils ProposeUtils) getIteration(client *ethclient.Client, address string, proposer types.ElectedProposer, utilsStruct UtilsStruct) int {
	return getIteration(client, address, proposer, utilsStruct)
}

func (proposeUtils ProposeUtils) isElectedProposer(client *ethclient.Client, proposer types.ElectedProposer, utilsStruct UtilsStruct) bool {
	return isElectedProposer(client, proposer, utilsStruct)
}

func (proposeUtils ProposeUtils) pseudoRandomNumberGenerator(seed []byte, max uint32, blockHashes []byte) *big.Int {
	return pseudoRandomNumberGenerator(seed, max, blockHashes)
}

func (proposeUtils ProposeUtils) MakeBlock(client *ethclient.Client, address string, rogueMode bool, utilsStruct UtilsStruct) ([]uint32, error) {
	return MakeBlock(client, address, rogueMode, utilsStruct)
}

func (proposeUtils ProposeUtils) getSortedVotes(client *ethclient.Client, address string, assetId uint8, epoch uint32, utilsStruct UtilsStruct) ([]*big.Int, error) {
	return getSortedVotes(client, address, assetId, epoch, utilsStruct)
}

func (proposeUtils ProposeUtils) influencedMedian(sortedVotes []*big.Int, totalInfluenceRevealed *big.Int) *big.Int {
	return influencedMedian(sortedVotes, totalInfluenceRevealed)
}

func (blockManagerUtils BlockManagerUtils) Propose(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, medians []uint32, iteration *big.Int, biggestInfluencerId uint32) (*Types.Transaction, error) {
	blockManager := utils.GetBlockManager(client)
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

func (flagSetUtils FlagSetUtils) GetStringProvider(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("provider")
}

func (flagSetUtils FlagSetUtils) GetFloat32GasMultiplier(flagSet *pflag.FlagSet) (float32, error) {
	return flagSet.GetFloat32("gasmultiplier")
}

func (flagSetUtils FlagSetUtils) GetInt32Buffer(flagSet *pflag.FlagSet) (int32, error) {
	return flagSet.GetInt32("buffer")
}

func (flagSetUtils FlagSetUtils) GetInt32Wait(flagSet *pflag.FlagSet) (int32, error) {
	return flagSet.GetInt32("wait")
}

func (flagSetUtils FlagSetUtils) GetInt32GasPrice(flagSet *pflag.FlagSet) (int32, error) {
	return flagSet.GetInt32("gasprice")
}

func (flagSetUtils FlagSetUtils) GetStringLogLevel(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("logLevel")
}

func (flagSetUtils FlagSetUtils) GetStringPow(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("pow")
}

func (flagSetUtils FlagSetUtils) GetBoolAutoWithdraw(flagSet *pflag.FlagSet) (bool, error) {
	return flagSet.GetBool("autoWithdraw")
}

func (flagSetUtils FlagSetUtils) GetUint32BountyId(flagSet *pflag.FlagSet) (uint32, error) {
	return flagSet.GetUint32("bountyId")
}

func (flagSetUtils FlagSetUtils) GetRootStringProvider() (string, error) {
	return rootCmd.PersistentFlags().GetString("provider")
}

func (flagSetUtils FlagSetUtils) GetRootFloat32GasMultiplier() (float32, error) {
	return rootCmd.PersistentFlags().GetFloat32("gasmultiplier")
}

func (flagSetUtils FlagSetUtils) GetRootInt32Buffer() (int32, error) {
	return rootCmd.PersistentFlags().GetInt32("buffer")
}

func (flagSetUtils FlagSetUtils) GetRootInt32Wait() (int32, error) {
	return rootCmd.PersistentFlags().GetInt32("wait")
}

func (flagSetUtils FlagSetUtils) GetRootInt32GasPrice() (int32, error) {
	return rootCmd.PersistentFlags().GetInt32("gasprice")
}

func (flagSetUtils FlagSetUtils) getRootStringLogLevel() (string, error) {
	return rootCmd.PersistentFlags().GetString("logLevel")
}

func (flagSetUtils FlagSetUtils) GetRootFloat32GasLimit() (float32, error) {
	return rootCmd.PersistentFlags().GetFloat32("gasLimit")
}

func (cmdUtils UtilsCmd) SetCommission(client *ethclient.Client, stakerId uint32, txnOpts *bind.TransactOpts, commission uint8, utilsStruct UtilsStruct) error {
	return SetCommission(client, stakerId, txnOpts, commission, utilsStruct)
}

func (cmdUtils UtilsCmd) DecreaseCommission(client *ethclient.Client, stakerId uint32, txnOpts *bind.TransactOpts, commission uint8, utilsStruct UtilsStruct) error {
	return DecreaseCommission(client, stakerId, txnOpts, commission, utilsStruct)
}

func (cmdUtils UtilsCmd) DecreaseCommissionPrompt() bool {
	return DecreaseCommissionPrompt()
}

func (cmdUtils UtilsCmd) Withdraw(client *ethclient.Client, txnOpts *bind.TransactOpts, epoch uint32, stakerId uint32, utilsStruct UtilsStruct) (common.Hash, error) {
	return withdraw(client, txnOpts, epoch, stakerId, utilsStruct)
}

func (cmdUtils UtilsCmd) CheckCurrentStatus(client *ethclient.Client, assetId uint8, utilsStruct UtilsStruct) (bool, error) {
	return CheckCurrentStatus(client, assetId, utilsStruct)
}

func (cmdUtils UtilsCmd) Dispute(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, blockId uint8, assetId int, utilsStruct UtilsStruct) error {
	return Dispute(client, config, account, epoch, blockId, assetId, utilsStruct)
}

func (cmdUtils UtilsCmd) GiveSorted(client *ethclient.Client, blockManager *bindings.BlockManager, txnOpts *bind.TransactOpts, epoch uint32, assetId uint8, sortedStakers []uint32) {
	GiveSorted(client, blockManager, txnOpts, epoch, assetId, sortedStakers)
}

func (cmdUtils UtilsCmd) GetEpochAndState(client *ethclient.Client, accountAddress string, utilsStruct UtilsStruct) (uint32, int64, error) {
	return GetEpochAndState(client, accountAddress, utilsStruct)
}

func (cmdUtils UtilsCmd) WaitForAppropriateState(client *ethclient.Client, accountAddress string, action string, utilsStruct UtilsStruct, states ...int) (uint32, error) {
	return WaitForAppropriateState(client, accountAddress, action, utilsStruct, states...)
}

func (cmdUtils UtilsCmd) WaitIfCommitState(client *ethclient.Client, accountAddress string, action string, utilsStruct UtilsStruct) (uint32, error) {
	return WaitIfCommitState(client, accountAddress, action, utilsStruct)
}

func (cmdUtils UtilsCmd) AssignAmountInWei(flagSet *pflag.FlagSet, utilsStruct UtilsStruct) (*big.Int, error) {
	return AssignAmountInWei(flagSet, utilsStruct)
}

func (cmdUtils UtilsCmd) Unstake(config types.Configurations, client *ethclient.Client, unstakeInput types.UnstakeInput, utilsStruct UtilsStruct) (types.TransactionOptions, error) {
	return Unstake(config, client, unstakeInput, utilsStruct)
}

func (cmdUtils UtilsCmd) AutoWithdraw(txnArgs types.TransactionOptions, stakerId uint32, utilsStruct UtilsStruct) error {
	return AutoWithdraw(txnArgs, stakerId, utilsStruct)
}

func (cmdUtils UtilsCmd) withdrawFunds(client *ethclient.Client, account types.Account, configurations types.Configurations, stakerId uint32, utilsStruct UtilsStruct) (common.Hash, error) {
	return withdrawFunds(client, account, configurations, stakerId, utilsStruct)
}

func (cmdUtils UtilsCmd) Create(password string, utilsStruct UtilsStruct) (ethAccounts.Account, error) {
	return Create(password, utilsStruct)
}

func (cmdUtils UtilsCmd) claimBounty(config types.Configurations, client *ethclient.Client, redeemBountyInput types.RedeemBountyInput, utilsStruct UtilsStruct) (common.Hash, error) {
	return claimBounty(config, client, redeemBountyInput, utilsStruct)
}

func (cmdUtils UtilsCmd) GetAmountInSRZRs(client *ethclient.Client, address string, staker bindings.StructsStaker, amount *big.Int, utilsStruct UtilsStruct) (*big.Int, error) {
	return GetAmountInSRZRs(client, address, staker, amount, utilsStruct)
}

func (blockManagerUtils BlockManagerUtils) ClaimBlockReward(client *ethclient.Client, opts *bind.TransactOpts) (*Types.Transaction, error) {
	blockManager := utils.GetBlockManager(client)
	return blockManager.ClaimBlockReward(opts)
}

func (blockManagerUtils BlockManagerUtils) FinalizeDispute(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, blockIndex uint8) (*Types.Transaction, error) {
	blockManager := utils.GetBlockManager(client)
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

func (blockManagerUtils BlockManagerUtils) DisputeBiggestInfluenceProposed(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, blockIndex uint8, correctBiggestInfluencerId uint32) (*Types.Transaction, error) {
	blockManager := utils.GetBlockManager(client)
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

func (c CryptoUtils) HexToECDSA(hexKey string) (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(hexKey)
}
