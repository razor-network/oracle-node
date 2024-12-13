// Package cmd provides all functions related to command line
package cmd

import (
	"crypto/ecdsa"
	"math/big"
	"razor/cache"
	"razor/core/types"
	"razor/path"
	"razor/pkg/bindings"
	"razor/rpc"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
)

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
	ApproveUnstake(client *ethclient.Client, opts *bind.TransactOpts, stakerTokenAddress common.Address, amount *big.Int) (*Types.Transaction, error)
	ClaimStakerReward(client *ethclient.Client, opts *bind.TransactOpts) (*Types.Transaction, error)
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
	GiveSorted(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, leafId uint16, sortedValues []*big.Int) (*Types.Transaction, error)
	ResetDispute(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32) (*Types.Transaction, error)
}

type VoteManagerInterface interface {
	Commit(client *ethclient.Client, txnOpts *bind.TransactOpts, epoch uint32, commitment [32]byte) (*Types.Transaction, error)
	Reveal(client *ethclient.Client, txnOpts *bind.TransactOpts, epoch uint32, tree bindings.StructsMerkleTree, signature []byte) (*Types.Transaction, error)
}

type TokenManagerInterface interface {
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
	FetchFlagInput(flagSet *pflag.FlagSet, flagKeyword string, dataType string) (interface{}, error)
	FetchRootFlagInput(flagName string, dataType string) (interface{}, error)
	Changed(flagSet *pflag.FlagSet, flagName string) bool
	GetUint32BountyId(flagSet *pflag.FlagSet) (uint32, error)
	GetStringFrom(flagSet *pflag.FlagSet) (string, error)
	GetStringTo(flagSet *pflag.FlagSet) (string, error)
	GetStringAddress(flagSet *pflag.FlagSet) (string, error)
	GetUint32StakerId(flagSet *pflag.FlagSet) (uint32, error)
	GetStringName(flagSet *pflag.FlagSet) (string, error)
	GetStringUrl(flagSet *pflag.FlagSet) (string, error)
	GetStringSelector(flagSet *pflag.FlagSet) (string, error)
	GetInt8Power(flagSet *pflag.FlagSet) (int8, error)
	GetUint8Weight(flagSet *pflag.FlagSet) (uint8, error)
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
	GetStringSliceBackupNode(flagSet *pflag.FlagSet) ([]string, error)
	GetStringExposeMetrics(flagSet *pflag.FlagSet) (string, error)
	GetStringCertFile(flagSet *pflag.FlagSet) (string, error)
	GetStringCertKey(flagSet *pflag.FlagSet) (string, error)
	GetIntLogFileMaxSize(flagSet *pflag.FlagSet) (int, error)
	GetIntLogFileMaxBackups(flagSet *pflag.FlagSet) (int, error)
	GetIntLogFileMaxAge(flagSet *pflag.FlagSet) (int, error)
}

type UtilsCmdInterface interface {
	SetConfig(flagSet *pflag.FlagSet) error
	GetProvider() (string, error)
	GetMultiplier() (float32, error)
	GetWaitTime() (int32, error)
	GetGasPrice() (int32, error)
	GetLogLevel() (string, error)
	GetGasLimit() (float32, error)
	GetGasLimitOverride() (uint64, error)
	GetBufferPercent() (int32, error)
	GetRPCTimeout() (int64, error)
	GetHTTPTimeout() (int64, error)
	GetLogFileMaxSize() (int, error)
	GetLogFileMaxBackups() (int, error)
	GetLogFileMaxAge() (int, error)
	GetConfigData() (types.Configurations, error)
	ExecuteClaimBounty(flagSet *pflag.FlagSet)
	ClaimBounty(rpcParameters rpc.RPCParameters, config types.Configurations, redeemBountyInput types.RedeemBountyInput) (common.Hash, error)
	ClaimBlockReward(rpcParameters rpc.RPCParameters, options types.TransactionOptions) (common.Hash, error)
	GetSalt(rpcParameters rpc.RPCParameters, epoch uint32) ([32]byte, error)
	HandleCommitState(rpcParameters rpc.RPCParameters, epoch uint32, seed []byte, commitParams *types.CommitParams, rogueData types.Rogue) (types.CommitData, error)
	Commit(rpcParameters rpc.RPCParameters, config types.Configurations, account types.Account, epoch uint32, latestHeader *Types.Header, stateBuffer uint64, commitment [32]byte) (common.Hash, error)
	ListAccounts() ([]accounts.Account, error)
	AssignAmountInWei(flagSet *pflag.FlagSet) (*big.Int, error)
	ExecuteTransfer(flagSet *pflag.FlagSet)
	Transfer(rpcParameters rpc.RPCParameters, config types.Configurations, transferInput types.TransferInput) (common.Hash, error)
	CheckForLastCommitted(rpcParameters rpc.RPCParameters, staker bindings.StructsStaker, epoch uint32) error
	Reveal(rpcParameters rpc.RPCParameters, config types.Configurations, account types.Account, epoch uint32, latestHeader *Types.Header, stateBuffer uint64, commitData types.CommitData, signature []byte) (common.Hash, error)
	GenerateTreeRevealData(merkleTree [][][]byte, commitData types.CommitData) bindings.StructsMerkleTree
	IndexRevealEventsOfCurrentEpoch(rpcParameters rpc.RPCParameters, blockNumber *big.Int, epoch uint32) ([]types.RevealedStruct, error)
	ExecuteCreateJob(flagSet *pflag.FlagSet)
	CreateJob(rpcParameter rpc.RPCParameters, config types.Configurations, jobInput types.CreateJobInput) (common.Hash, error)
	ExecuteCreateCollection(flagSet *pflag.FlagSet)
	CreateCollection(rpcParameters rpc.RPCParameters, config types.Configurations, collectionInput types.CreateCollectionInput) (common.Hash, error)
	GetEpochAndState(rpcParameter rpc.RPCParameters) (uint32, int64, error)
	WaitForAppropriateState(rpcParameter rpc.RPCParameters, action string, states ...int) (uint32, error)
	ExecuteJobList(flagSet *pflag.FlagSet)
	GetJobList(rpcParameters rpc.RPCParameters) error
	ExecuteUnstake(flagSet *pflag.FlagSet)
	Unstake(rpcParameters rpc.RPCParameters, config types.Configurations, input types.UnstakeInput) (common.Hash, error)
	ApproveUnstake(rpcParameters rpc.RPCParameters, stakerTokenAddress common.Address, txnArgs types.TransactionOptions) (common.Hash, error)
	ExecuteInitiateWithdraw(flagSet *pflag.FlagSet)
	ExecuteUnlockWithdraw(flagSet *pflag.FlagSet)
	InitiateWithdraw(rpcParameters rpc.RPCParameters, txnOpts *bind.TransactOpts, stakerId uint32) (common.Hash, error)
	UnlockWithdraw(rpcParameters rpc.RPCParameters, txnOpts *bind.TransactOpts, stakerId uint32) (common.Hash, error)
	HandleUnstakeLock(rpcParameters rpc.RPCParameters, account types.Account, configurations types.Configurations, stakerId uint32) (common.Hash, error)
	HandleWithdrawLock(rpcParameters rpc.RPCParameters, account types.Account, configurations types.Configurations, stakerId uint32) (common.Hash, error)
	ExecuteUpdateJob(flagSet *pflag.FlagSet)
	UpdateJob(rpcParameters rpc.RPCParameters, config types.Configurations, jobInput types.CreateJobInput, jobId uint16) (common.Hash, error)
	WaitIfCommitState(rpcParameter rpc.RPCParameters, action string) (uint32, error)
	ExecuteCollectionList(flagSet *pflag.FlagSet)
	GetCollectionList(rpcParameters rpc.RPCParameters) error
	ExecuteStakerinfo(flagSet *pflag.FlagSet)
	ExecuteSetDelegation(flagSet *pflag.FlagSet)
	SetDelegation(rpcParameters rpc.RPCParameters, config types.Configurations, delegationInput types.SetDelegationInput) (common.Hash, error)
	GetStakerInfo(rpcParameters rpc.RPCParameters, stakerId uint32) error
	ExecuteUpdateCollection(flagSet *pflag.FlagSet)
	UpdateCollection(rpcParameters rpc.RPCParameters, config types.Configurations, collectionInput types.CreateCollectionInput, collectionId uint16) (common.Hash, error)
	MakeBlock(rpcParameters rpc.RPCParameters, blockNumber *big.Int, epoch uint32, rogueData types.Rogue) ([]*big.Int, []uint16, *types.RevealedDataMaps, error)
	IsElectedProposer(proposer types.ElectedProposer, currentStakerStake *big.Int) bool
	GetSortedRevealedValues(rpcParameters rpc.RPCParameters, blockNumber *big.Int, epoch uint32) (*types.RevealedDataMaps, error)
	GetIteration(rpcParameters rpc.RPCParameters, proposer types.ElectedProposer, bufferPercent int32) int
	Propose(rpcParameters rpc.RPCParameters, config types.Configurations, account types.Account, staker bindings.StructsStaker, epoch uint32, latestHeader *Types.Header, stateBuffer uint64, rogueData types.Rogue) error
	GiveSorted(rpcParameters rpc.RPCParameters, txnArgs types.TransactionOptions, epoch uint32, assetId uint16, sortedStakers []*big.Int) error
	GetLocalMediansData(rpcParameters rpc.RPCParameters, account types.Account, epoch uint32, blockNumber *big.Int, rogueData types.Rogue) (types.ProposeFileData, error)
	CheckDisputeForIds(rpcParameters rpc.RPCParameters, transactionOpts types.TransactionOptions, epoch uint32, blockIndex uint8, idsInProposedBlock []uint16, revealedCollectionIds []uint16) (*Types.Transaction, error)
	Dispute(rpcParameters rpc.RPCParameters, config types.Configurations, account types.Account, epoch uint32, blockIndex uint8, proposedBlock bindings.StructsBlock, leafId uint16, sortedValues []*big.Int) error
	GetCollectionIdPositionInBlock(rpcParameters rpc.RPCParameters, leafId uint16, proposedBlock bindings.StructsBlock) *big.Int
	HandleDispute(rpcParameters rpc.RPCParameters, config types.Configurations, account types.Account, epoch uint32, blockNumber *big.Int, rogueData types.Rogue, backupNodeActionsToIgnore []string) error
	ExecuteExtendLock(flagSet *pflag.FlagSet)
	ResetUnstakeLock(rpcParameters rpc.RPCParameters, config types.Configurations, extendLockInput types.ExtendLockInput) (common.Hash, error)
	ExecuteModifyCollectionStatus(flagSet *pflag.FlagSet)
	ModifyCollectionStatus(rpcParameters rpc.RPCParameters, config types.Configurations, modifyCollectionInput types.ModifyCollectionInput) (common.Hash, error)
	Approve(rpcParameters rpc.RPCParameters, txnArgs types.TransactionOptions) (common.Hash, error)
	ExecuteDelegate(flagSet *pflag.FlagSet)
	Delegate(rpcParameters rpc.RPCParameters, txnArgs types.TransactionOptions, stakerId uint32) (common.Hash, error)
	ExecuteCreate(flagSet *pflag.FlagSet)
	Create(password string) (accounts.Account, error)
	ExecuteImport(flagSet *pflag.FlagSet)
	ImportAccount() (accounts.Account, error)
	ExecuteUpdateCommission(flagSet *pflag.FlagSet)
	UpdateCommission(rpcParameters rpc.RPCParameters, config types.Configurations, updateCommissionInput types.UpdateCommissionInput) error
	GetBiggestStakeAndId(rpcParameters rpc.RPCParameters, epoch uint32) (*big.Int, uint32, error)
	GetSmallestStakeAndId(rpcParameters rpc.RPCParameters, epoch uint32) (*big.Int, uint32, error)
	StakeCoins(rpcParameters rpc.RPCParameters, txnArgs types.TransactionOptions) (common.Hash, error)
	CalculateSecret(account types.Account, epoch uint32, keystorePath string, chainId *big.Int) ([]byte, []byte, error)
	HandleBlock(rpcParameters rpc.RPCParameters, account types.Account, stakerId uint32, header *Types.Header, config types.Configurations, commitParams *types.CommitParams, rogueData types.Rogue, backupNodeActionsToIgnore []string)
	ExecuteVote(flagSet *pflag.FlagSet)
	Vote(rpcParameters rpc.RPCParameters, config types.Configurations, account types.Account, stakerId uint32, commitParams *types.CommitParams, rogueData types.Rogue, backupNodeActionsToIgnore []string) error
	HandleExit()
	ExecuteListAccounts(flagSet *pflag.FlagSet)
	ClaimCommission(flagSet *pflag.FlagSet)
	ExecuteStake(flagSet *pflag.FlagSet)
	InitiateCommit(rpcParameters rpc.RPCParameters, config types.Configurations, account types.Account, epoch uint32, stakerId uint32, latestHeader *Types.Header, commitParams *types.CommitParams, stateBuffer uint64, rogueData types.Rogue) error
	InitiateReveal(rpcParameters rpc.RPCParameters, config types.Configurations, account types.Account, epoch uint32, staker bindings.StructsStaker, latestHeader *Types.Header, stateBuffer uint64, rogueData types.Rogue) error
	InitiatePropose(rpcParameters rpc.RPCParameters, config types.Configurations, account types.Account, epoch uint32, staker bindings.StructsStaker, latestHeader *Types.Header, stateBuffer uint64, rogueData types.Rogue) error
	GetBountyIdFromEvents(rpcParameters rpc.RPCParameters, blockNumber *big.Int, bountyHunter string) (uint32, error)
	HandleClaimBounty(rpcParameters rpc.RPCParameters, config types.Configurations, account types.Account) error
	ExecuteContractAddresses(flagSet *pflag.FlagSet)
	ContractAddresses()
	ResetDispute(rpcParameters rpc.RPCParameters, txnOpts *bind.TransactOpts, epoch uint32)
	StoreBountyId(rpcParameters rpc.RPCParameters, account types.Account) error
	CheckToDoResetDispute(rpcParameters rpc.RPCParameters, txnOpts *bind.TransactOpts, epoch uint32, sortedValues []*big.Int)
	InitJobAndCollectionCache(rpcParameters rpc.RPCParameters) (*cache.JobsCache, *cache.CollectionsCache, *big.Int, error)
	BatchGetStakeSnapshotCalls(rpcParameters rpc.RPCParameters, epoch uint32, numberOfStakers uint32) ([]*big.Int, error)
	ExecuteImportEndpoints()
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

	path.PathUtilsInterface = path.PathUtils{}
	path.OSUtilsInterface = path.OSUtils{}
	InitializeUtils()
}
