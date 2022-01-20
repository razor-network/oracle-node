package cmd

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"math/big"
	"razor/core/types"
	"razor/pkg/bindings"
	"time"
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

type UtilsInterface interface {
	GetConfigFilePath() (string, error)
	ViperWriteConfigAs(string) error
	GetEpoch(*ethclient.Client) (uint32, error)
	GetUpdatedEpoch(*ethclient.Client) (uint32, error)
	GetOptions() bind.CallOpts
	CalculateBlockTime(*ethclient.Client) int64
	Sleep(time.Duration)
	GetTxnOpts(types.TransactionOptions) *bind.TransactOpts
	AssignPassword(*pflag.FlagSet) string
	GetStringAddress(*pflag.FlagSet) (string, error)
	GetUint32BountyId(*pflag.FlagSet) (uint32, error)
	ConnectToClient(string) *ethclient.Client
	WaitForBlockCompletion(*ethclient.Client, string) int
	GetNumActiveAssets(*ethclient.Client) (*big.Int, error)
	GetRogueRandomValue(int) *big.Int
	GetActiveAssetsData(*ethclient.Client, uint32) ([]*big.Int, error)
	GetDelayedState(*ethclient.Client, int32) (int64, error)
	GetDefaultPath() (string, error)
	GetJobFilePath() (string, error)
	FetchBalance(*ethclient.Client, string) (*big.Int, error)
	IsFlagPassed(string) bool
	GetFractionalAmountInWei(*big.Int, string) (*big.Int, error)
	GetAmountInWei(*big.Int) *big.Int
	CheckAmountAndBalance(*big.Int, *big.Int) *big.Int
	GetAmountInDecimal(*big.Int) *big.Float
	GetEpochLastCommitted(*ethclient.Client, uint32) (uint32, error)
	GetCommitments(*ethclient.Client, string) ([32]byte, error)
	AllZero([32]byte) bool
	ConvertUintArrayToUint16Array(uintArr []uint) []uint16
	GetStateName(int64) string
	GetJobs(*ethclient.Client) ([]bindings.StructsJob, error)
	CheckEthBalanceIsZero(*ethclient.Client, string)
	AssignStakerId(*pflag.FlagSet, *ethclient.Client, string) (uint32, error)
	GetLock(*ethclient.Client, string, uint32) (types.Locks, error)
	GetStaker(*ethclient.Client, string, uint32) (bindings.StructsStaker, error)
	GetUpdatedStaker(*ethclient.Client, string, uint32) (bindings.StructsStaker, error)
	GetStakedToken(*ethclient.Client, common.Address) *bindings.StakedToken
	ConvertSRZRToRZR(*big.Int, *big.Int, *big.Int) *big.Int
	ConvertRZRToSRZR(*big.Int, *big.Int, *big.Int) (*big.Int, error)
	GetWithdrawReleasePeriod(*ethclient.Client, string) (uint8, error)
	GetCollections(*ethclient.Client) ([]bindings.StructsCollection, error)
	GetInfluenceSnapshot(*ethclient.Client, uint32, uint32) (*big.Int, error)
	ParseBool(str string) (bool, error)
	GetStakerId(*ethclient.Client, string) (uint32, error)
	GetNumberOfStakers(*ethclient.Client, string) (uint32, error)
	GetRandaoHash(*ethclient.Client) ([32]byte, error)
	GetNumberOfProposedBlocks(*ethclient.Client, uint32) (uint8, error)
	GetMaxAltBlocks(*ethclient.Client) (uint8, error)
	GetProposedBlock(*ethclient.Client, uint32, uint32) (bindings.StructsBlock, error)
	GetEpochLastRevealed(*ethclient.Client, string, uint32) (uint32, error)
	GetVoteValue(*ethclient.Client, uint16, uint32) (*big.Int, error)
	GetTotalInfluenceRevealed(*ethclient.Client, uint32) (*big.Int, error)
	ConvertBigIntArrayToUint32Array([]*big.Int) []uint32
	GetActiveAssetIds(*ethclient.Client) ([]uint16, error)
	GetBlockManager(*ethclient.Client) *bindings.BlockManager
	GetVotes(*ethclient.Client, uint32) (bindings.StructsVote, error)
	GetSortedProposedBlockIds(*ethclient.Client, uint32) ([]uint32, error)
	PrivateKeyPrompt() string
	PasswordPrompt() string
	GetMaxCommission(client *ethclient.Client) (uint8, error)
	GetEpochLimitForUpdateCommission(client *ethclient.Client) (uint16, error)
	GetStakeSnapshot(*ethclient.Client, uint32, uint32) (*big.Int, error)
	DeleteJobFromJSON(string, string) error
	AddJobToJSON(string, *types.StructsJob) error
}

type StakeManagerInterface interface {
	Stake(*ethclient.Client, *bind.TransactOpts, uint32, *big.Int) (*Types.Transaction, error)
	ExtendLock(*ethclient.Client, *bind.TransactOpts, uint32) (*Types.Transaction, error)
	Delegate(*ethclient.Client, *bind.TransactOpts, uint32, *big.Int) (*Types.Transaction, error)
	Withdraw(*ethclient.Client, *bind.TransactOpts, uint32) (*Types.Transaction, error)
	SetDelegationAcceptance(*ethclient.Client, *bind.TransactOpts, bool) (*Types.Transaction, error)
	Unstake(*ethclient.Client, *bind.TransactOpts, uint32, *big.Int) (*Types.Transaction, error)
	RedeemBounty(*ethclient.Client, *bind.TransactOpts, uint32) (*Types.Transaction, error)
	UpdateCommission(client *ethclient.Client, opts *bind.TransactOpts, commission uint8) (*Types.Transaction, error)

	//Getter methods
	StakerInfo(*ethclient.Client, *bind.CallOpts, uint32) (types.Staker, error)
	GetMaturity(*ethclient.Client, *bind.CallOpts, uint32) (uint16, error)
	GetBountyLock(*ethclient.Client, *bind.CallOpts, uint32) (types.BountyLock, error)
	BalanceOf(*bindings.StakedToken, *bind.CallOpts, common.Address) (*big.Int, error)
	GetTotalSupply(*bindings.StakedToken, *bind.CallOpts) (*big.Int, error)
}

type KeystoreInterface interface {
	Accounts(string) []accounts.Account
	ImportECDSA(string, *ecdsa.PrivateKey, string) (accounts.Account, error)
}

type BlockManagerInterface interface {
	ClaimBlockReward(*ethclient.Client, *bind.TransactOpts) (*Types.Transaction, error)
	Propose(*ethclient.Client, *bind.TransactOpts, uint32, []uint32, *big.Int, uint32) (*Types.Transaction, error)
	FinalizeDispute(*ethclient.Client, *bind.TransactOpts, uint32, uint8) (*Types.Transaction, error)
	DisputeBiggestStakeProposed(*ethclient.Client, *bind.TransactOpts, uint32, uint8, uint32) (*Types.Transaction, error)
}

type VoteManagerInterface interface {
	Commit(*ethclient.Client, *bind.TransactOpts, uint32, [32]byte) (*Types.Transaction, error)
	Reveal(*ethclient.Client, *bind.TransactOpts, uint32, []*big.Int, [32]byte) (*Types.Transaction, error)
}

type TokenManagerInterface interface {
	Allowance(*ethclient.Client, *bind.CallOpts, common.Address, common.Address) (*big.Int, error)
	Approve(*ethclient.Client, *bind.TransactOpts, common.Address, *big.Int) (*Types.Transaction, error)
	Transfer(*ethclient.Client, *bind.TransactOpts, common.Address, *big.Int) (*Types.Transaction, error)
}

type AssetManagerInterface interface {
	CreateJob(*ethclient.Client, *bind.TransactOpts, uint8, int8, uint8, string, string, string) (*Types.Transaction, error)
	SetCollectionStatus(*ethclient.Client, *bind.TransactOpts, bool, uint16) (*Types.Transaction, error)
	GetActiveStatus(*ethclient.Client, *bind.CallOpts, uint16) (bool, error)
	CreateCollection(*ethclient.Client, *bind.TransactOpts, uint16, int8, uint32, []uint16, string) (*Types.Transaction, error)
	UpdateJob(*ethclient.Client, *bind.TransactOpts, uint16, uint8, int8, uint8, string, string) (*Types.Transaction, error)
	UpdateCollection(*ethclient.Client, *bind.TransactOpts, uint16, uint16, uint32, int8, []uint16) (*Types.Transaction, error)
}

type FlagSetInterface interface {
	GetStringProvider(*pflag.FlagSet) (string, error)
	GetFloat32GasMultiplier(*pflag.FlagSet) (float32, error)
	GetInt32Buffer(*pflag.FlagSet) (int32, error)
	GetInt32Wait(*pflag.FlagSet) (int32, error)
	GetInt32GasPrice(*pflag.FlagSet) (int32, error)
	GetFloat32GasLimit(set *pflag.FlagSet) (float32, error)
	GetStringLogLevel(*pflag.FlagSet) (string, error)
	GetBoolAutoWithdraw(*pflag.FlagSet) (bool, error)
	GetUint32BountyId(*pflag.FlagSet) (uint32, error)
	GetRootStringProvider() (string, error)
	GetRootFloat32GasMultiplier() (float32, error)
	GetRootInt32Buffer() (int32, error)
	GetRootInt32Wait() (int32, error)
	GetRootInt32GasPrice() (int32, error)
	GetRootStringLogLevel() (string, error)
	GetRootFloat32GasLimit() (float32, error)
	GetStringFrom(*pflag.FlagSet) (string, error)
	GetStringTo(*pflag.FlagSet) (string, error)
	GetStringAddress(*pflag.FlagSet) (string, error)
	GetUint32StakerId(*pflag.FlagSet) (uint32, error)
	GetStringName(*pflag.FlagSet) (string, error)
	GetStringUrl(*pflag.FlagSet) (string, error)
	GetStringSelector(*pflag.FlagSet) (string, error)
	GetInt8Power(*pflag.FlagSet) (int8, error)
	GetUint8Weight(*pflag.FlagSet) (uint8, error)
	GetUint16AssetId(*pflag.FlagSet) (uint16, error)
	GetUint8SelectorType(set *pflag.FlagSet) (uint8, error)
	GetStringStatus(*pflag.FlagSet) (string, error)
	GetUint8Commission(*pflag.FlagSet) (uint8, error)
	GetUintSliceJobIds(*pflag.FlagSet) ([]uint, error)
	GetUint32Aggregation(*pflag.FlagSet) (uint32, error)
	GetUint16JobId(*pflag.FlagSet) (uint16, error)
	GetUint16CollectionId(*pflag.FlagSet) (uint16, error)
	GetStringValue(*pflag.FlagSet) (string, error)
	GetStringPow(flagSet *pflag.FlagSet) (string, error)
	GetUint16Tolerance(set *pflag.FlagSet) (uint16, error)
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
	ExecuteClaimBounty(*pflag.FlagSet)
	ClaimBounty(types.Configurations, *ethclient.Client, types.RedeemBountyInput) (common.Hash, error)
	ClaimBlockReward(types.TransactionOptions) (common.Hash, error)
	HandleCommitState(*ethclient.Client, uint32, types.Rogue) ([]*big.Int, error)
	Commit(*ethclient.Client, []*big.Int, []byte, types.Account, types.Configurations) (common.Hash, error)
	ListAccounts() ([]accounts.Account, error)
	AssignAmountInWei(*pflag.FlagSet) (*big.Int, error)
	ExecuteTransfer(*pflag.FlagSet)
	Transfer(*ethclient.Client, types.Configurations, types.TransferInput) (common.Hash, error)
	HandleRevealState(*ethclient.Client, bindings.StructsStaker, uint32) error
	Reveal(*ethclient.Client, []*big.Int, []byte, types.Account, string, types.Configurations) (common.Hash, error)
	ExecuteCreateJob(*pflag.FlagSet)
	CreateJob(*ethclient.Client, types.Configurations, types.CreateJobInput) (common.Hash, error)
	ExecuteCreateCollection(*pflag.FlagSet)
	CreateCollection(*ethclient.Client, types.Configurations, types.CreateCollectionInput) (common.Hash, error)
	GetEpochAndState(*ethclient.Client) (uint32, int64, error)
	WaitForAppropriateState(*ethclient.Client, string, ...int) (uint32, error)
	ExecuteJobList()
	GetJobList(*ethclient.Client) error
	GetAmountInSRZRs(*ethclient.Client, string, bindings.StructsStaker, *big.Int) (*big.Int, error)
	ExecuteUnstake(*pflag.FlagSet)
	Unstake(types.Configurations, *ethclient.Client, types.UnstakeInput) (types.TransactionOptions, error)
	AutoWithdraw(types.TransactionOptions, uint32) error
	ExecuteWithdraw(*pflag.FlagSet)
	Withdraw(*ethclient.Client, *bind.TransactOpts, uint32) (common.Hash, error)
	WithdrawFunds(*ethclient.Client, types.Account, types.Configurations, uint32) (common.Hash, error)
	ExecuteUpdateJob(*pflag.FlagSet)
	UpdateJob(*ethclient.Client, types.Configurations, types.CreateJobInput, uint16) (common.Hash, error)
	WaitIfCommitState(*ethclient.Client, string) (uint32, error)
	ExecuteCollectionList()
	GetCollectionList(*ethclient.Client) error
	ExecuteStakerinfo(*pflag.FlagSet)
	ExecuteSetDelegation(*pflag.FlagSet)
	SetDelegation(*ethclient.Client, types.Configurations, types.SetDelegationInput) (common.Hash, error)
	GetStakerInfo(*ethclient.Client, uint32) error
	ExecuteUpdateCollection(*pflag.FlagSet)
	UpdateCollection(*ethclient.Client, types.Configurations, types.CreateCollectionInput, uint16) (common.Hash, error)
	InfluencedMedian([]*big.Int, *big.Int) *big.Int
	GetSortedVotes(*ethclient.Client, string, uint16, uint32) ([]*big.Int, error)
	MakeBlock(*ethclient.Client, string, types.Rogue) ([]uint32, error)
	IsElectedProposer(*ethclient.Client, types.ElectedProposer) bool
	GetIteration(*ethclient.Client, types.ElectedProposer) int
	Propose(*ethclient.Client, types.Account, types.Configurations, uint32, uint32, types.Rogue) (common.Hash, error)
	GiveSorted(*ethclient.Client, *bindings.BlockManager, *bind.TransactOpts, uint32, uint16, []uint32)
	Dispute(*ethclient.Client, types.Configurations, types.Account, uint32, uint8, int) error
	HandleDispute(*ethclient.Client, types.Configurations, types.Account, uint32) error
	ExecuteExtendLock(*pflag.FlagSet)
	ExtendLock(*ethclient.Client, types.Configurations, types.ExtendLockInput) (common.Hash, error)
	CheckCurrentStatus(*ethclient.Client, uint16) (bool, error)
	ExecuteModifyAssetStatus(*pflag.FlagSet)
	ModifyAssetStatus(*ethclient.Client, types.Configurations, types.ModifyAssetInput) (common.Hash, error)
	Approve(types.TransactionOptions) (common.Hash, error)
	ExecuteDelegate(*pflag.FlagSet)
	Delegate(types.TransactionOptions, uint32) (common.Hash, error)
	ExecuteCreate(*pflag.FlagSet)
	Create(string) (accounts.Account, error)
	ExecuteImport()
	ImportAccount() (accounts.Account, error)
	ExecuteUpdateCommission(*pflag.FlagSet)
	UpdateCommission(types.Configurations, *ethclient.Client, types.UpdateCommissionInput) error
	GetBiggestStakeAndId(*ethclient.Client, string, uint32) (*big.Int, uint32, error)
	ExecuteOverrideJob(*pflag.FlagSet)
	OverrideJob(*types.StructsJob) error
	ExecuteDeleteOverrideJob(*pflag.FlagSet)
	DeleteOverrideJob(uint16) error
	StakeCoins(types.TransactionOptions) (common.Hash, error)
	ExecuteListAccounts()
}

type TransactionInterface interface {
	Hash(*Types.Transaction) common.Hash
}

type CryptoInterface interface {
	HexToECDSA(string) (*ecdsa.PrivateKey, error)
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
