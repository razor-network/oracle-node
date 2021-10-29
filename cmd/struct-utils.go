package cmd

import (
	"crypto/ecdsa"
	"math/big"
	"razor/accounts"
	"razor/core/types"
	"razor/path"
	"razor/pkg/bindings"
	"razor/utils"
	"strconv"

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
type AccountUtils struct{}
type KeystoreUtils struct{}
type FlagSetUtils struct{}
type ProposeUtils struct{}
type UtilsCmd struct{}
type VoteManagerUtils struct{}
type BlockManagerUtils struct{}
type CryptoUtils struct{}

func (u Utils) ConnectToClient(provider string) *ethclient.Client {
	return utils.ConnectToClient(provider)
}

func (u Utils) GetOptions(pending bool, from string, blockNumber string) bind.CallOpts {
	return utils.GetOptions(pending, from, blockNumber)
}

func (u Utils) GetTxnOpts(transactionData types.TransactionOptions) *bind.TransactOpts {
	return utils.GetTxnOpts(transactionData)
}

func (u Utils) WaitForBlockCompletion(client *ethclient.Client, hashToRead string) int {
	return utils.WaitForBlockCompletion(client, hashToRead)
}

func (u Utils) WaitForCommitState(client *ethclient.Client, accountAddress string, action string) (uint32, error) {
	return WaitForCommitState(client, accountAddress, action)
}

func (u Utils) WaitForCommitStateAgain(client *ethclient.Client, accountAddress string, action string) (uint32, error) {
	return WaitForCommitState(client, accountAddress, action)
}

func (u Utils) AssignPassword(flagSet *pflag.FlagSet) string {
	return utils.AssignPassword(flagSet)
}

func (u Utils) FetchBalance(client *ethclient.Client, accountAddress string) (*big.Int, error) {
	return utils.FetchBalance(client, accountAddress)
}

func (u Utils) AssignAmountInWei(flagSet *pflag.FlagSet) *big.Int {
	return utils.AssignAmountInWei(flagSet)
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

func (u Utils) GetConfigData() (types.Configurations, error) {
	return GetConfigData()
}

func (u Utils) ParseBool(str string) (bool, error) {
	return strconv.ParseBool(str)
}

func (u Utils) GetDelayedState(client *ethclient.Client, buffer int32) (int64, error) {
	return utils.GetDelayedState(client, buffer)
}

func (u Utils) GetEpoch(client *ethclient.Client, address string) (uint32, error) {
	return utils.GetEpoch(client, address)
}

func (u Utils) GetActiveAssetsData(client *ethclient.Client, address string, epoch uint32) ([]*big.Int, error) {
	return utils.GetActiveAssetsData(client, address, epoch)
}

func (u Utils) ConvertUintArrayToUint8Array(uintArr []uint) []uint8 {
	return utils.ConvertUintArrayToUint8Array(uintArr)
}

func (u Utils) WaitForConfirmState(client *ethclient.Client, accountAddress string, action string) (uint32, error) {
	return WaitForConfirmState(client, accountAddress, action)
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
	return utils.GetStakerId(client, address)
}

func (u Utils) GetRandaoHash(client *ethclient.Client, address string) ([32]byte, error) {
	return utils.GetRandaoHash(client, address)
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

func (u Utils) GetInfluence(client *ethclient.Client, address string, stakerId uint32) (*big.Int, error) {
	return utils.GetInfluence(client, address, stakerId)
}

func (u Utils) GetEpochLastRevealed(client *ethclient.Client, address string, stakerId uint32) (uint32, error) {
	return utils.GetEpochLastRevealed(client, address, stakerId)
}

func (u Utils) GetVoteValue(client *ethclient.Client, address string, assetId uint8, stakerId uint32) (*big.Int, error) {
	return utils.GetVoteValue(client, address, assetId, stakerId)
}

func (u Utils) GetInfluenceSnapshot(client *ethclient.Client, address string, stakerId uint32, epoch uint32) (*big.Int, error) {
	return utils.GetInfluenceSnapshot(client, address, stakerId, epoch)
}

func (u Utils) GetNumActiveAssets(client *ethclient.Client, address string) (*big.Int, error) {
	return utils.GetNumActiveAssets(client, address)
}

func (u Utils) GetTotalInfluenceRevealed(client *ethclient.Client, address string, epoch uint32) (*big.Int, error) {
	return utils.GetTotalInfluenceRevealed(client, address, epoch)
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

func (u Utils) GetEpochLastCommitted(client *ethclient.Client, address string, stakerId uint32) (uint32, error) {
	return utils.GetEpochLastCommitted(client, address, stakerId)
}

func (u Utils) GetConfigFilePath() (string, error) {
	return path.GetConfigFilePath()
}

func (u Utils) ViperWriteConfigAs(path string) error {
	return viper.WriteConfigAs(path)
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

func (account AccountUtils) CreateAccount(path string, password string) ethAccounts.Account {
	return accounts.CreateAccount(path, password)
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

func (flagSetUtils FlagSetUtils) GetStringStatus(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("status")
}

func (voteManagerUtils VoteManagerUtils) Reveal(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, values []*big.Int, secret [32]byte) (*Types.Transaction, error) {
	voteManager := utils.GetVoteManager(client)
	return voteManager.Reveal(opts, epoch, values, secret)
}

func (voteManagerUtils VoteManagerUtils) Commit(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, commitment [32]byte) (*Types.Transaction, error) {
	voteManager := utils.GetVoteManager(client)
	return voteManager.Commit(opts, epoch, commitment)
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

func (proposeUtils ProposeUtils) getBiggestInfluenceAndId(client *ethclient.Client, address string, razorUtils utilsInterface) (*big.Int, uint32, error) {
	return getBiggestInfluenceAndId(client, address, razorUtils)
}

func (proposeUtils ProposeUtils) getIteration(client *ethclient.Client, address string, proposer types.ElectedProposer, proposeUtil proposeUtilsInterface) int {
	return getIteration(client, address, proposer, proposeUtil)
}

func (proposeUtils ProposeUtils) isElectedProposer(client *ethclient.Client, address string, proposer types.ElectedProposer) bool {
	return isElectedProposer(client, address, proposer)
}

func (proposeUtils ProposeUtils) pseudoRandomNumberGenerator(seed []byte, max uint32, blockHashes []byte) *big.Int {
	return pseudoRandomNumberGenerator(seed, max, blockHashes)
}

func (proposeUtils ProposeUtils) MakeBlock(client *ethclient.Client, address string, rogueMode bool, razorUtils utilsInterface, proposeUtil proposeUtilsInterface) ([]uint32, error) {
	return MakeBlock(client, address, rogueMode, razorUtils, proposeUtil)
}

func (proposeUtils ProposeUtils) getSortedVotes(client *ethclient.Client, address string, assetId uint8, epoch uint32, razorUtils utilsInterface) ([]*big.Int, error) {
	return getSortedVotes(client, address, assetId, epoch, razorUtils)
}

func (proposeUtils ProposeUtils) influencedMedian(sortedVotes []*big.Int, totalInfluenceRevealed *big.Int) *big.Int {
	return influencedMedian(sortedVotes, totalInfluenceRevealed)
}

func (blockManagerUtils BlockManagerUtils) Propose(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, medians []uint32, iteration *big.Int, biggestInfluencerId uint32) (*Types.Transaction, error) {
	blockManager := utils.GetBlockManager(client)
	return blockManager.Propose(opts, epoch, medians, iteration, biggestInfluencerId)
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

func (cmdUtils UtilsCmd) SetCommission(client *ethclient.Client, stakerId uint32, txnOpts *bind.TransactOpts, commission uint8, razorUtils utilsInterface, stakeManagerUtils stakeManagerInterface, transactionUtils transactionInterface) error {
	return SetCommission(client, stakerId, txnOpts, commission, razorUtils, stakeManagerUtils, transactionUtils)
}

func (cmdUtils UtilsCmd) DecreaseCommission(client *ethclient.Client, stakerId uint32, txnOpts *bind.TransactOpts, commission uint8, razorUtils utilsInterface, stakeManagerUtils stakeManagerInterface, transactionUtils transactionInterface, cmdUtil utilsCmdInterface) error {
	return DecreaseCommission(client, stakerId, txnOpts, commission, razorUtils, stakeManagerUtils, transactionUtils, cmdUtil)
}

func (cmdUtils UtilsCmd) DecreaseCommissionPrompt() bool {
	return DecreaseCommissionPrompt()
}

func (cmdUtils UtilsCmd) Withdraw(client *ethclient.Client, txnOpts *bind.TransactOpts, epoch uint32, stakerId uint32, stakeManagerUtils stakeManagerInterface, transactionUtils transactionInterface) (common.Hash, error) {
	return withdraw(client, txnOpts, epoch, stakerId, stakeManagerUtils, transactionUtils)
}

func (cmdUtils UtilsCmd) CheckCurrentStatus(client *ethclient.Client, address string, assetId uint8, razorUtils utilsInterface, assetManagerUtils assetManagerInterface) (bool, error) {
	return CheckCurrentStatus(client, address, assetId, razorUtils, assetManagerUtils)
}

func (blockManagerUtils BlockManagerUtils) ClaimBlockReward(client *ethclient.Client, opts *bind.TransactOpts) (*Types.Transaction, error) {
	blockManager := utils.GetBlockManager(client)
	return blockManager.ClaimBlockReward(opts)
}

func (c CryptoUtils) HexToECDSA(hexKey string) (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(hexKey)
}
