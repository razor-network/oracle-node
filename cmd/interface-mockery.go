package cmd

import (
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

//go:generate mockery --name UtilsInterfaceMockery --output ./mocks/ --case=underscore
//go:generate mockery --name FlagSetInterfaceMockery --output ./mocks/ --case=underscore
//go:generate mockery --name UtilsCmdInterfaceMockery --output ./mocks/ --case=underscore
//go:generate mockery --name StakeManagerInterfaceMockery --output ./mocks/ --case=underscore
//go:generate mockery --name TransactionInterfaceMockery --output ./mocks/ --case=underscore
//go:generate mockery --name BlockManagerInterfaceMockery --output ./mocks/ --case=underscore
//go:generate mockery --name VoteManagerInterfaceMockery --output ./mocks/ --case=underscore
//go:generate mockery --name TokenManagerInterfaceMockery --output ./mocks/ --case=underscore
//go:generate mockery --name AssetManagerInterfaceMockery --output ./mocks/ --case=underscore

var razorUtilsMockery UtilsInterfaceMockery
var flagSetUtilsMockery FlagSetInterfaceMockery
var cmdUtilsMockery UtilsCmdInterfaceMockery
var stakeManagerUtilsMockery StakeManagerInterfaceMockery
var transactionUtilsMockery TransactionInterfaceMockery
var blockManagerUtilsMockery BlockManagerInterfaceMockery
var voteManagerUtilsMockery VoteManagerInterfaceMockery
var tokenManagerUtilsMockery TokenManagerInterfaceMockery
var assetManagerUtilsMockery AssetManagerInterfaceMockery

type UtilsInterfaceMockery interface {
	GetConfigFilePath() (string, error)
	ViperWriteConfigAs(string) error
	GetEpoch(*ethclient.Client) (uint32, error)
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
	FetchBalance(*ethclient.Client, string) (*big.Int, error)
	IsFlagPassed(string) bool
	GetFractionalAmountInWei(*big.Int, string) (*big.Int, error)
	GetAmountInWei(*big.Int) *big.Int
	CheckAmountAndBalance(*big.Int, *big.Int) *big.Int
	GetAmountInDecimal(*big.Int) *big.Float
	GetEpochLastCommitted(*ethclient.Client, uint32) (uint32, error)
	GetCommitments(*ethclient.Client, string) ([32]byte, error)
	AllZero([32]byte) bool
}

type StakeManagerInterfaceMockery interface {
	GetBountyLock(*ethclient.Client, *bind.CallOpts, uint32) (types.BountyLock, error)
	RedeemBounty(*ethclient.Client, *bind.TransactOpts, uint32) (*Types.Transaction, error)
}

type BlockManagerInterfaceMockery interface {
	ClaimBlockReward(*ethclient.Client, *bind.TransactOpts) (*Types.Transaction, error)
	Propose(*ethclient.Client, *bind.TransactOpts, uint32, []uint32, *big.Int, uint32) (*Types.Transaction, error)
	FinalizeDispute(*ethclient.Client, *bind.TransactOpts, uint32, uint8) (*Types.Transaction, error)
	DisputeBiggestInfluenceProposed(*ethclient.Client, *bind.TransactOpts, uint32, uint8, uint32) (*Types.Transaction, error)
}

type VoteManagerInterfaceMockery interface {
	Commit(*ethclient.Client, *bind.TransactOpts, uint32, [32]byte) (*Types.Transaction, error)
	Reveal(*ethclient.Client, *bind.TransactOpts, uint32, []*big.Int, [32]byte) (*Types.Transaction, error)
}

type TokenManagerInterfaceMockery interface {
	Allowance(*ethclient.Client, *bind.CallOpts, common.Address, common.Address) (*big.Int, error)
	Approve(*ethclient.Client, *bind.TransactOpts, common.Address, *big.Int) (*Types.Transaction, error)
	Transfer(*ethclient.Client, *bind.TransactOpts, common.Address, *big.Int) (*Types.Transaction, error)
}

type AssetManagerInterfaceMockery interface {
	CreateJob(*ethclient.Client, *bind.TransactOpts, uint8, int8, uint8, string, string, string) (*Types.Transaction, error)
	SetCollectionStatus(*ethclient.Client, *bind.TransactOpts, bool, uint16) (*Types.Transaction, error)
	GetActiveStatus(*ethclient.Client, *bind.CallOpts, uint16) (bool, error)
	CreateCollection(client *ethclient.Client, opts *bind.TransactOpts, jobIDs []uint16, aggregationMethod uint32, power int8, name string) (*Types.Transaction, error)
	UpdateJob(*ethclient.Client, *bind.TransactOpts, uint16, uint8, int8, uint8, string, string) (*Types.Transaction, error)
	UpdateCollection(*ethclient.Client, *bind.TransactOpts, uint16, uint32, int8, []uint16) (*Types.Transaction, error)
}

type FlagSetInterfaceMockery interface {
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
}

type UtilsCmdInterfaceMockery interface {
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
	AssignAmountInWei(*pflag.FlagSet) (*big.Int, error)
	ExecuteTransfer(*pflag.FlagSet)
	Transfer(*ethclient.Client, types.Configurations, types.TransferInput) (common.Hash, error)
	HandleRevealState(*ethclient.Client, bindings.StructsStaker, uint32) error
	Reveal(*ethclient.Client, []*big.Int, []byte, types.Account, string, types.Configurations) (common.Hash, error)
	ExecuteCreateJob(*pflag.FlagSet)
	CreateJob(*ethclient.Client, types.Configurations, types.CreateJobInput) (common.Hash, error)
}

type TransactionInterfaceMockery interface {
	Hash(*Types.Transaction) common.Hash
}

type UtilsMockery struct{}
type FLagSetUtilsMockery struct{}
type UtilsStructMockery struct{}
type StakeManagerUtilsMockery struct{}
type BlockManagerUtilsMockery struct{}
type TransactionUtilsMockery struct{}
type VoteManagerUtilsMockery struct{}
type TokenManagerUtilsMockery struct{}
type AssetManagerUtilsMockery struct{}
