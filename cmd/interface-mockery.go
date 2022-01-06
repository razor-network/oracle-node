package cmd

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"math/big"
	"razor/core/types"
	"time"
)

//go:generate mockery --name UtilsInterfaceMockery --output ./mocks/ --case=underscore
//go:generate mockery --name FlagSetInterfaceMockery --output ./mocks/ --case=underscore
//go:generate mockery --name UtilsCmdInterfaceMockery --output ./mocks/ --case=underscore
//go:generate mockery --name StakeManagerInterfaceMockery --output ./mocks/ --case=underscore
//go:generate mockery --name TransactionInterfaceMockery --output ./mocks/ --case=underscore
//go:generate mockery --name BlockManagerInterfaceMockery --output ./mocks/ --case=underscore
//go:generate mockery --name VoteManagerInterfaceMockery --output ./mocks/ --case=underscore

var razorUtilsMockery UtilsInterfaceMockery
var flagSetUtilsMockery FlagSetInterfaceMockery
var cmdUtilsMockery UtilsCmdInterfaceMockery
var stakeManagerUtilsMockery StakeManagerInterfaceMockery
var transactionUtilsMockery TransactionInterfaceMockery
var blockManagerUtilsMockery BlockManagerInterfaceMockery
var voteManagerUtilsMockery VoteManagerInterfaceMockery

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
