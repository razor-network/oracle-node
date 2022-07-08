//Package cmd provides all functions related to command line
package cmd

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	Accounts "razor/accounts"
	"razor/core/types"
	"razor/path"
	"razor/pkg/bindings"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
)

//go:generate mockery --name UtilsInterface --output ./mocks/ --case=underscore
//go:generate mockery --name FlagSetInterface --output ./mocks/ --case=underscore
//go:generate mockery --name UtilsCmdInterface --output ./mocks/ --case=underscore
//go:generate mockery --name StakeManagerInterface --output ./mocks/ --case=underscore
//go:generate mockery --name TransactionInterface --output ./mocks/ --case=underscore
//go:generate mockery --name BlockManagerInterface --output ./mocks/ --case=underscore
//go:generate mockery --name VoteManagerInterface --output ./mocks/ --case=underscore
//go:generate mockery --name KeystoreInterface --output ./mocks/ --case=underscore
//go:generate mockery --name TokenManagerInterface --output ./mocks/ --case=underscore
//go:generate mockery --name AssetManagerInterface --output ./mocks/ --case=underscore
//go:generate mockery --name CryptoInterface --output ./mocks/ --case=underscore
//go:generate mockery --name ViperInterface --output ./mocks/ --case=underscore
//go:generate mockery --name TimeInterface --output ./mocks/ --case=underscore
//go:generate mockery --name StringInterface --output ./mocks/ --case=underscore
//go:generate mockery --name AbiInterface --output ./mocks/ --case=underscore
//go:generate mockery --name OSInterface --output ./mocks/ --case=underscore

var razorUtils UtilsInterface
var flagSetUtils FlagSetInterface
var cmdUtils UtilsCmdInterface
var stakeManagerUtils StakeManagerInterface
var transactionUtils TransactionInterface
var blockManagerUtils BlockManagerInterface
var voteManagerUtils VoteManagerInterface
var keystoreUtils KeystoreInterface
var tokenManagerUtils TokenManagerInterface
var assetManagerUtils AssetManagerInterface
var cryptoUtils CryptoInterface
var viperUtils ViperInterface
var timeUtils TimeInterface
var stringUtils StringInterface
var abiUtils AbiInterface
var osUtils OSInterface

type UtilsInterface interface {
	GetConfigFilePath() (string, error)
	GetEpoch(client *ethclient.Client) (uint32, error)
	GetOptions() bind.CallOpts
	CalculateBlockTime(client *ethclient.Client) int64
	GetTxnOpts(transactionData types.TransactionOptions) *bind.TransactOpts
	AssignPassword() string
	GetStringAddress(flagSet *pflag.FlagSet) (string, error)
	GetUint32BountyId(flagSet *pflag.FlagSet) (uint32, error)
	ConnectToClient(provider string) *ethclient.Client
	WaitForBlockCompletion(client *ethclient.Client, hashToRead string) error
	GetNumActiveCollections(client *ethclient.Client) (uint16, error)
	GetRogueRandomValue(value int) *big.Int
	GetRogueRandomMedianValue() uint32
	GetAggregatedDataOfCollection(client *ethclient.Client, collectionId uint16, epoch uint32) (*big.Int, error)
	GetDelayedState(client *ethclient.Client, buffer int32) (int64, error)
	GetDefaultPath() (string, error)
	GetJobFilePath() (string, error)
	FetchBalance(client *ethclient.Client, accountAddress string) (*big.Int, error)
	IsFlagPassed(name string) bool
	GetAmountInWei(amount *big.Int) *big.Int
	CheckAmountAndBalance(amountInWei *big.Int, balance *big.Int) *big.Int
	GetAmountInDecimal(amountInWei *big.Int) *big.Float
	GetEpochLastCommitted(client *ethclient.Client, stakerId uint32) (uint32, error)
	GetCommitments(client *ethclient.Client, address string) ([32]byte, error)
	AllZero(bytesValue [32]byte) bool
	ConvertUintArrayToUint16Array(uintArr []uint) []uint16
	ConvertUint32ArrayToBigIntArray(uint32Array []uint32) []*big.Int
	GetJobs(client *ethclient.Client) ([]bindings.StructsJob, error)
	CheckEthBalanceIsZero(client *ethclient.Client, address string)
	AssignStakerId(flagSet *pflag.FlagSet, client *ethclient.Client, address string) (uint32, error)
	GetLock(client *ethclient.Client, address string, stakerId uint32, lockType uint8) (types.Locks, error)
	GetStaker(client *ethclient.Client, stakerId uint32) (bindings.StructsStaker, error)
	GetUpdatedStaker(client *ethclient.Client, stakerId uint32) (bindings.StructsStaker, error)
	GetStakedToken(client *ethclient.Client, address common.Address) *bindings.StakedToken
	ConvertSRZRToRZR(sAmount *big.Int, currentStake *big.Int, totalSupply *big.Int) *big.Int
	ConvertRZRToSRZR(sAmount *big.Int, currentStake *big.Int, totalSupply *big.Int) (*big.Int, error)
	GetWithdrawInitiationPeriod(client *ethclient.Client) (uint8, error)
	GetCollections(client *ethclient.Client) ([]bindings.StructsCollection, error)
	GetInfluenceSnapshot(client *ethclient.Client, stakerId uint32, epoch uint32) (*big.Int, error)
	GetStakerId(client *ethclient.Client, address string) (uint32, error)
	GetNumberOfStakers(client *ethclient.Client) (uint32, error)
	GetNumberOfProposedBlocks(client *ethclient.Client, epoch uint32) (uint8, error)
	GetMaxAltBlocks(client *ethclient.Client) (uint8, error)
	GetProposedBlock(client *ethclient.Client, epoch uint32, proposedBlockId uint32) (bindings.StructsBlock, error)
	GetEpochLastRevealed(client *ethclient.Client, stakerId uint32) (uint32, error)
	GetVoteValue(client *ethclient.Client, epoch uint32, stakerId uint32, medianIndex uint16) (*big.Int, error)
	GetTotalInfluenceRevealed(client *ethclient.Client, epoch uint32, medianIndex uint16) (*big.Int, error)
	GetActiveCollections(client *ethclient.Client) ([]uint16, error)
	GetBlockManager(client *ethclient.Client) *bindings.BlockManager
	GetSortedProposedBlockIds(client *ethclient.Client, epoch uint32) ([]uint32, error)
	PrivateKeyPrompt() string
	PasswordPrompt() string
	GetMaxCommission(client *ethclient.Client) (uint8, error)
	GetEpochLimitForUpdateCommission(client *ethclient.Client) (uint16, error)
	GetStakeSnapshot(client *ethclient.Client, stakerId uint32, epoch uint32) (*big.Int, error)
	GetStake(client *ethclient.Client, stakerId uint32) (*big.Int, error)
	ConvertWeiToEth(data *big.Int) (*big.Float, error)
	WaitTillNextNSecs(seconds int32)
	DeleteJobFromJSON(s string, jobId string) error
	AddJobToJSON(s string, job *types.StructsJob) error
	GetStakerSRZRBalance(client *ethclient.Client, staker bindings.StructsStaker) (*big.Int, error)
	SecondsToReadableTime(time int) string
	SaveDataToCommitJsonFile(flePath string, epoch uint32, commitFileData types.CommitData) error
	ReadFromCommitJsonFile(filePath string) (types.CommitFileData, error)
	SaveDataToProposeJsonFile(flePath string, epoch uint32, proposeFileData types.ProposeData) error
	ReadFromProposeJsonFile(filePath string) (types.ProposeFileData, error)
	SaveDataToDisputeJsonFile(filePath string, bountyIdQueue []uint32) error
	ReadFromDisputeJsonFile(filePath string) (types.DisputeFileData, error)
	AssignLogFile(flagSet *pflag.FlagSet)
	GetCommitDataFileName(address string) (string, error)
	GetProposeDataFileName(address string) (string, error)
	GetDisputeDataFileName(address string) (string, error)
}

type StakeManagerInterface interface {
	Stake(client *ethclient.Client, txnOpts *bind.TransactOpts, epoch uint32, amount *big.Int) (*Types.Transaction, error)
	ResetUnstakeLock(client *ethclient.Client, opts *bind.TransactOpts, stakerId uint32) (*Types.Transaction, error)
	Delegate(client *ethclient.Client, opts *bind.TransactOpts, stakerId uint32, amount *big.Int) (*Types.Transaction, error)
	InitiateWithdraw(client *ethclient.Client, opts *bind.TransactOpts, stakerId uint32) (*Types.Transaction, error)
	UnlockWithdraw(client *ethclient.Client, opts *bind.TransactOpts, stakerId uint32) (*Types.Transaction, error)
	SetDelegationAcceptance(client *ethclient.Client, opts *bind.TransactOpts, status bool) (*Types.Transaction, error)
	Unstake(client *ethclient.Client, opts *bind.TransactOpts, stakerId uint32, sAmount *big.Int) (*Types.Transaction, error)
	RedeemBounty(client *ethclient.Client, opts *bind.TransactOpts, bountyId uint32) (*Types.Transaction, error)
	UpdateCommission(client *ethclient.Client, opts *bind.TransactOpts, commission uint8) (*Types.Transaction, error)
	ApproveUnstake(client *ethclient.Client, opts *bind.TransactOpts, staker bindings.StructsStaker, amount *big.Int) (*Types.Transaction, error)
	ClaimStakeReward(client *ethclient.Client, opts *bind.TransactOpts) (*Types.Transaction, error)

	//Getter methods
	StakerInfo(client *ethclient.Client, opts *bind.CallOpts, stakerId uint32) (types.Staker, error)
	GetMaturity(client *ethclient.Client, opts *bind.CallOpts, age uint32) (uint16, error)
	GetBountyLock(client *ethclient.Client, opts *bind.CallOpts, bountyId uint32) (types.BountyLock, error)
}

type KeystoreInterface interface {
	Accounts(path string) []accounts.Account
	ImportECDSA(path string, priv *ecdsa.PrivateKey, passphrase string) (accounts.Account, error)
}

type BlockManagerInterface interface {
	ClaimBlockReward(client *ethclient.Client, opts *bind.TransactOpts) (*Types.Transaction, error)
	Propose(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, ids []uint16, medians []*big.Int, iteration *big.Int, biggestInfluencerId uint32) (*Types.Transaction, error)
	FinalizeDispute(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, blockIndex uint8, positionOfCollectionInBlock *big.Int) (*Types.Transaction, error)
	DisputeBiggestStakeProposed(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, blockIndex uint8, correctBiggestStakerId uint32) (*Types.Transaction, error)
	DisputeOnOrderOfIds(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, blockIndex uint8, index0 *big.Int, index1 *big.Int) (*Types.Transaction, error)
	DisputeCollectionIdShouldBeAbsent(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, blockIndex uint8, id uint16, positionOfCollectionInBlock *big.Int) (*Types.Transaction, error)
	DisputeCollectionIdShouldBePresent(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, blockIndex uint8, id uint16) (*Types.Transaction, error)
	GiveSorted(blockManager *bindings.BlockManager, opts *bind.TransactOpts, epoch uint32, leafId uint16, sortedValues []*big.Int) (*Types.Transaction, error)
	ResetDispute(blockManager *bindings.BlockManager, opts *bind.TransactOpts, epoch uint32) (*Types.Transaction, error)
}

type VoteManagerInterface interface {
	Commit(client *ethclient.Client, txnOpts *bind.TransactOpts, epoch uint32, commitment [32]byte) (*Types.Transaction, error)
	Reveal(client *ethclient.Client, txnOpts *bind.TransactOpts, epoch uint32, tree bindings.StructsMerkleTree, secret [32]byte, signature []byte) (*Types.Transaction, error)
}

type TokenManagerInterface interface {
	Allowance(client *ethclient.Client, opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error)
	Approve(client *ethclient.Client, opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*Types.Transaction, error)
	Transfer(client *ethclient.Client, opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*Types.Transaction, error)
}

type AssetManagerInterface interface {
	CreateJob(client *ethclient.Client, opts *bind.TransactOpts, weight uint8, power int8, selectorType uint8, name string, selector string, url string) (*Types.Transaction, error)
	SetCollectionStatus(client *ethclient.Client, opts *bind.TransactOpts, assetStatus bool, id uint16) (*Types.Transaction, error)
	GetActiveStatus(client *ethclient.Client, opts *bind.CallOpts, id uint16) (bool, error)
	CreateCollection(client *ethclient.Client, opts *bind.TransactOpts, tolerance uint32, power int8, aggregationMethod uint32, jobIDs []uint16, name string) (*Types.Transaction, error)
	UpdateJob(client *ethclient.Client, opts *bind.TransactOpts, jobId uint16, weight uint8, power int8, selectorType uint8, selector string, url string) (*Types.Transaction, error)
	UpdateCollection(client *ethclient.Client, opts *bind.TransactOpts, collectionId uint16, tolerance uint32, aggregationMethod uint32, power int8, jobIds []uint16) (*Types.Transaction, error)
}

type FlagSetInterface interface {
	GetStringProvider(flagSet *pflag.FlagSet) (string, error)
	GetFloat32GasMultiplier(flagSet *pflag.FlagSet) (float32, error)
	GetInt32Buffer(flagSet *pflag.FlagSet) (int32, error)
	GetInt32Wait(flagSet *pflag.FlagSet) (int32, error)
	GetInt32GasPrice(flagSet *pflag.FlagSet) (int32, error)
	GetFloat32GasLimit(flagSet *pflag.FlagSet) (float32, error)
	GetStringLogLevel(flagSet *pflag.FlagSet) (string, error)
	GetUint32BountyId(flagSet *pflag.FlagSet) (uint32, error)
	GetRootStringProvider() (string, error)
	GetRootFloat32GasMultiplier() (float32, error)
	GetRootInt32Buffer() (int32, error)
	GetRootInt32Wait() (int32, error)
	GetRootInt32GasPrice() (int32, error)
	GetRootStringLogLevel() (string, error)
	GetRootFloat32GasLimit() (float32, error)
	GetStringFrom(flagSet *pflag.FlagSet) (string, error)
	GetStringTo(flagSet *pflag.FlagSet) (string, error)
	GetStringAddress(flagSet *pflag.FlagSet) (string, error)
	GetUint32StakerId(flagSet *pflag.FlagSet) (uint32, error)
	GetStringName(flagSet *pflag.FlagSet) (string, error)
	GetStringUrl(flagSet *pflag.FlagSet) (string, error)
	GetStringSelector(flagSet *pflag.FlagSet) (string, error)
	GetInt8Power(flagSet *pflag.FlagSet) (int8, error)
	GetUint8Weight(flagSet *pflag.FlagSet) (uint8, error)
	GetUint16AssetId(flagSet *pflag.FlagSet) (uint16, error)
	GetUint8SelectorType(flagSet *pflag.FlagSet) (uint8, error)
	GetStringStatus(flagSet *pflag.FlagSet) (string, error)
	GetUint8Commission(flagSet *pflag.FlagSet) (uint8, error)
	GetUintSliceJobIds(flagSet *pflag.FlagSet) ([]uint, error)
	GetUint32Aggregation(flagSet *pflag.FlagSet) (uint32, error)
	GetUint16JobId(flagSet *pflag.FlagSet) (uint16, error)
	GetUint16CollectionId(flagSet *pflag.FlagSet) (uint16, error)
	GetStringValue(flagSet *pflag.FlagSet) (string, error)
	GetBoolWeiRazor(flagSet *pflag.FlagSet) (bool, error)
	GetUint32Tolerance(flagSet *pflag.FlagSet) (uint32, error)
	GetBoolRogue(flagSet *pflag.FlagSet) (bool, error)
	GetStringSliceRogueMode(flagSet *pflag.FlagSet) ([]string, error)
	GetStringExposeMetrics(flagSet *pflag.FlagSet) (string, error)
	GetStringCertFile(flagSet *pflag.FlagSet) (string, error)
	GetStringCertKey(flagSet *pflag.FlagSet) (string, error)
}

type UtilsCmdInterface interface {
	SetConfig(flagSet *pflag.FlagSet) error
	GetProvider() (string, error)
	GetMultiplier() (float32, error)
	GetWaitTime() (int32, error)
	GetGasPrice() (int32, error)
	GetLogLevel() (string, error)
	GetGasLimit() (float32, error)
	GetBufferPercent() (int32, error)
	GetConfigData() (types.Configurations, error)
	ExecuteClaimBounty(flagSet *pflag.FlagSet)
	ClaimBounty(config types.Configurations, client *ethclient.Client, redeemBountyInput types.RedeemBountyInput) (common.Hash, error)
	ClaimBlockReward(options types.TransactionOptions) (common.Hash, error)
	GetSalt(client *ethclient.Client, epoch uint32) ([32]byte, error)
	HandleCommitState(client *ethclient.Client, epoch uint32, seed []byte, rogueData types.Rogue) (types.CommitData, error)
	Commit(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, seed []byte, root [32]byte) (common.Hash, error)
	ListAccounts() ([]accounts.Account, error)
	AssignAmountInWei(flagSet *pflag.FlagSet) (*big.Int, error)
	ExecuteTransfer(flagSet *pflag.FlagSet)
	Transfer(client *ethclient.Client, config types.Configurations, transferInput types.TransferInput) (common.Hash, error)
	HandleRevealState(client *ethclient.Client, staker bindings.StructsStaker, epoch uint32) error
	Reveal(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, commitData types.CommitData, secret []byte, signature []byte) (common.Hash, error)
	GenerateTreeRevealData(merkleTree [][][]byte, commitData types.CommitData) bindings.StructsMerkleTree
	IndexRevealEventsOfCurrentEpoch(client *ethclient.Client, blockNumber *big.Int, epoch uint32) ([]types.RevealedStruct, error)
	ExecuteCreateJob(flagSet *pflag.FlagSet)
	CreateJob(client *ethclient.Client, config types.Configurations, jobInput types.CreateJobInput) (common.Hash, error)
	ExecuteCreateCollection(flagSet *pflag.FlagSet)
	CreateCollection(client *ethclient.Client, config types.Configurations, collectionInput types.CreateCollectionInput) (common.Hash, error)
	GetEpochAndState(client *ethclient.Client) (uint32, int64, error)
	WaitForAppropriateState(client *ethclient.Client, action string, states ...int) (uint32, error)
	ExecuteJobList(flagSet *pflag.FlagSet)
	GetJobList(client *ethclient.Client) error
	ExecuteUnstake(flagSet *pflag.FlagSet)
	Unstake(config types.Configurations, client *ethclient.Client, input types.UnstakeInput) (common.Hash, error)
	ApproveUnstake(client *ethclient.Client, staker bindings.StructsStaker, txnArgs types.TransactionOptions) (common.Hash, error)
	ExecuteInitiateWithdraw(flagSet *pflag.FlagSet)
	ExecuteUnlockWithdraw(flagSet *pflag.FlagSet)
	InitiateWithdraw(client *ethclient.Client, txnOpts *bind.TransactOpts, stakerId uint32) (common.Hash, error)
	UnlockWithdraw(client *ethclient.Client, txnOpts *bind.TransactOpts, stakerId uint32) (common.Hash, error)
	HandleUnstakeLock(client *ethclient.Client, account types.Account, configurations types.Configurations, stakerId uint32) (common.Hash, error)
	HandleWithdrawLock(client *ethclient.Client, account types.Account, configurations types.Configurations, stakerId uint32) (common.Hash, error)
	ExecuteUpdateJob(flagSet *pflag.FlagSet)
	UpdateJob(client *ethclient.Client, config types.Configurations, jobInput types.CreateJobInput, jobId uint16) (common.Hash, error)
	WaitIfCommitState(client *ethclient.Client, action string) (uint32, error)
	ExecuteCollectionList(flagSet *pflag.FlagSet)
	GetCollectionList(client *ethclient.Client) error
	ExecuteStakerinfo(flagSet *pflag.FlagSet)
	ExecuteSetDelegation(flagSet *pflag.FlagSet)
	SetDelegation(client *ethclient.Client, config types.Configurations, delegationInput types.SetDelegationInput) (common.Hash, error)
	GetStakerInfo(client *ethclient.Client, stakerId uint32) error
	ExecuteUpdateCollection(flagSet *pflag.FlagSet)
	UpdateCollection(client *ethclient.Client, config types.Configurations, collectionInput types.CreateCollectionInput, collectionId uint16) (common.Hash, error)
	InfluencedMedian(sortedVotes []*big.Int, totalInfluenceRevealed *big.Int) *big.Int
	MakeBlock(client *ethclient.Client, blockNumber *big.Int, epoch uint32, rogueData types.Rogue) ([]*big.Int, []uint16, *types.RevealedDataMaps, error)
	IsElectedProposer(proposer types.ElectedProposer, currentStakerStake *big.Int) bool
	GetSortedRevealedValues(client *ethclient.Client, blockNumber *big.Int, epoch uint32) (*types.RevealedDataMaps, error)
	GetIteration(client *ethclient.Client, proposer types.ElectedProposer, bufferPercent int32) int
	Propose(client *ethclient.Client, config types.Configurations, account types.Account, staker bindings.StructsStaker, epoch uint32, blockNumber *big.Int, rogueData types.Rogue) (common.Hash, error)
	GiveSorted(client *ethclient.Client, blockManager *bindings.BlockManager, txnOpts *bind.TransactOpts, epoch uint32, assetId uint16, sortedStakers []*big.Int)
	GetLocalMediansData(client *ethclient.Client, account types.Account, epoch uint32, blockNumber *big.Int, rogueData types.Rogue) ([]*big.Int, []uint16, *types.RevealedDataMaps, error)
	CheckDisputeForIds(client *ethclient.Client, transactionOpts types.TransactionOptions, epoch uint32, blockIndex uint8, idsInProposedBlock []uint16, revealedCollectionIds []uint16) (*Types.Transaction, error)
	Dispute(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, blockIndex uint8, proposedBlock bindings.StructsBlock, leafId uint16, sortedValues []*big.Int) error
	GetCollectionIdPositionInBlock(client *ethclient.Client, leafId uint16, proposedBlock bindings.StructsBlock) *big.Int
	HandleDispute(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, blockNumber *big.Int, rogueData types.Rogue) error
	ExecuteExtendLock(flagSet *pflag.FlagSet)
	ResetUnstakeLock(client *ethclient.Client, config types.Configurations, extendLockInput types.ExtendLockInput) (common.Hash, error)
	CheckCurrentStatus(client *ethclient.Client, collectionId uint16) (bool, error)
	ExecuteModifyCollectionStatus(flagSet *pflag.FlagSet)
	ModifyCollectionStatus(client *ethclient.Client, config types.Configurations, modifyCollectionInput types.ModifyCollectionInput) (common.Hash, error)
	Approve(txnArgs types.TransactionOptions) (common.Hash, error)
	ExecuteDelegate(flagSet *pflag.FlagSet)
	Delegate(txnArgs types.TransactionOptions, stakerId uint32) (common.Hash, error)
	ExecuteCreate(flagSet *pflag.FlagSet)
	Create(password string) (accounts.Account, error)
	ExecuteImport(flagSet *pflag.FlagSet)
	ImportAccount() (accounts.Account, error)
	ExecuteUpdateCommission(flagSet *pflag.FlagSet)
	UpdateCommission(config types.Configurations, client *ethclient.Client, updateCommissionInput types.UpdateCommissionInput) error
	GetBiggestStakeAndId(client *ethclient.Client, address string, epoch uint32) (*big.Int, uint32, error)
	StakeCoins(txnArgs types.TransactionOptions) (common.Hash, error)
	CalculateSecret(account types.Account, epoch uint32) ([]byte, []byte, error)
	GetLastProposedEpoch(client *ethclient.Client, blockNumber *big.Int, stakerId uint32) (uint32, error)
	HandleBlock(client *ethclient.Client, account types.Account, blockNumber *big.Int, config types.Configurations, rogueData types.Rogue)
	ExecuteVote(flagSet *pflag.FlagSet)
	Vote(ctx context.Context, config types.Configurations, client *ethclient.Client, rogueData types.Rogue, account types.Account) error
	HandleExit()
	ExecuteListAccounts(flagSet *pflag.FlagSet)
	ClaimCommission(flagSet *pflag.FlagSet)
	ExecuteStake(flagSet *pflag.FlagSet)
	InitiateCommit(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, stakerId uint32, rogueData types.Rogue) error
	InitiateReveal(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, staker bindings.StructsStaker, rogueData types.Rogue) error
	InitiatePropose(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, staker bindings.StructsStaker, blockNumber *big.Int, rogueData types.Rogue) error
	GetBountyIdFromEvents(client *ethclient.Client, blockNumber *big.Int, bountyHunter string) (uint32, error)
	HandleClaimBounty(client *ethclient.Client, config types.Configurations, account types.Account) error
	ExecuteContractAddresses(flagSet *pflag.FlagSet)
	ContractAddresses()
	ResetDispute(client *ethclient.Client, blockManager *bindings.BlockManager, txnOpts *bind.TransactOpts, epoch uint32)
	StoreBountyId(client *ethclient.Client, account types.Account) error
}

type TransactionInterface interface {
	Hash(txn *Types.Transaction) common.Hash
}

type CryptoInterface interface {
	HexToECDSA(hexKey string) (*ecdsa.PrivateKey, error)
}

type ViperInterface interface {
	ViperWriteConfigAs(path string) error
}

type TimeInterface interface {
	Sleep(duration time.Duration)
}

type StringInterface interface {
	ParseBool(str string) (bool, error)
}

type AbiInterface interface {
	Unpack(abi abi.ABI, name string, data []byte) ([]interface{}, error)
}

type OSInterface interface {
	Exit(code int)
}

type Utils struct{}
type FLagSetUtils struct{}
type UtilsStruct struct{}
type StakeManagerUtils struct{}
type BlockManagerUtils struct{}
type TransactionUtils struct{}
type VoteManagerUtils struct{}
type KeystoreUtils struct{}
type TokenManagerUtils struct{}
type AssetManagerUtils struct{}
type CryptoUtils struct{}
type ViperUtils struct{}
type TimeUtils struct{}
type StringUtils struct{}
type AbiUtils struct{}
type OSUtils struct{}

func InitializeInterfaces() {
	razorUtils = Utils{}
	flagSetUtils = FLagSetUtils{}
	cmdUtils = &UtilsStruct{}
	stakeManagerUtils = StakeManagerUtils{}
	transactionUtils = TransactionUtils{}
	blockManagerUtils = BlockManagerUtils{}
	voteManagerUtils = VoteManagerUtils{}
	keystoreUtils = KeystoreUtils{}
	tokenManagerUtils = TokenManagerUtils{}
	assetManagerUtils = AssetManagerUtils{}
	cryptoUtils = CryptoUtils{}
	viperUtils = ViperUtils{}
	timeUtils = TimeUtils{}
	stringUtils = StringUtils{}
	abiUtils = AbiUtils{}
	osUtils = OSUtils{}

	Accounts.AccountUtilsInterface = Accounts.AccountUtils{}
	path.PathUtilsInterface = path.PathUtils{}
	path.OSUtilsInterface = path.OSUtils{}
	InitializeUtils()
}
