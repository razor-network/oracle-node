package utils

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"io"
	"io/fs"
	"math/big"
	"os"
	"razor/accounts"
	"razor/core/types"
	"razor/pkg/bindings"
	"time"
)

//go:generate mockery --name Utils --output ./mocks/ --case=underscore
//go:generate mockery --name EthClientUtils --output ./mocks --case=underscore
//go:generate mockery --name ClientUtils --output ./mocks --case=underscore
//go:generate mockery --name TimeUtils --output ./mocks --case=underscore
//go:generate mockery --name OSUtils --output ./mocks --case=underscore
//go:generate mockery --name BufioUtils --output ./mocks --case=underscore
//go:generate mockery --name CoinUtils --output ./mocks --case=underscore
//go:generate mockery --name IoutilUtils --output ./mocks --case=underscore
//go:generate mockery --name ABIUtils --output ./mocks --case=underscore
//go:generate mockery --name PathUtils --output ./mocks --case=underscore
//go:generate mockery --name BindUtils --output ./mocks --case=underscore
//go:generate mockery --name AccountsUtils --output ./mocks --case=underscore
//go:generate mockery --name BlockManagerUtils --output ./mocks --case=underscore
//go:generate mockery --name AssetManagerUtils --output ./mocks --case=underscore
//go:generate mockery --name VoteManagerUtils --output ./mocks --case=underscore
//go:generate mockery --name BindingsUtils --output ./mocks --case=underscore
//go:generate mockery --name JsonUtils --output ./mocks --case=underscore
//go:generate mockery --name StakedTokenUtils --output ./mocks --case=underscore
//go:generate mockery --name RetryUtils --output ./mocks --case=underscore

var UtilsInterface Utils
var EthClient EthClientUtils
var ClientInterface ClientUtils
var Time TimeUtils
var OS OSUtils
var Bufio BufioUtils
var CoinInterface CoinUtils
var IoutilInterface IoutilUtils
var ABIInterface ABIUtils
var PathInterface PathUtils
var BindInterface BindUtils
var AccountsInterface AccountsUtils
var BlockManagerInterface BlockManagerUtils
var StakeManagerInterface StakeManagerUtils
var AssetManagerInterface AssetManagerUtils
var VoteManagerInterface VoteManagerUtils
var BindingsInterface BindingsUtils
var JsonInterface JsonUtils
var StakedTokenInterface StakedTokenUtils
var RetryInterface RetryUtils

type Utils interface {
	SuggestGasPriceWithRetry(*ethclient.Client) (*big.Int, error)
	MultiplyFloatAndBigInt(*big.Int, float64) *big.Int
	GetPendingNonceAtWithRetry(*ethclient.Client, common.Address) (uint64, error)
	GetGasPrice(*ethclient.Client, types.Configurations) *big.Int
	GetTxnOpts(types.TransactionOptions) *bind.TransactOpts
	GetGasLimit(types.TransactionOptions, *bind.TransactOpts) (uint64, error)
	EstimateGasWithRetry(*ethclient.Client, ethereum.CallMsg) (uint64, error)
	IncreaseGasLimitValue(*ethclient.Client, uint64, float32) (uint64, error)
	GetLatestBlockWithRetry(*ethclient.Client) (*Types.Header, error)
	FilterLogsWithRetry(*ethclient.Client, ethereum.FilterQuery) ([]Types.Log, error)
	BalanceAtWithRetry(*ethclient.Client, common.Address) (*big.Int, error)
	GetBlockManager(*ethclient.Client) *bindings.BlockManager
	GetOptions() bind.CallOpts
	GetNumberOfProposedBlocks(*ethclient.Client, uint32) (uint8, error)
	GetSortedProposedBlockId(*ethclient.Client, uint32, *big.Int) (uint32, error)
	FetchPreviousValue(*ethclient.Client, uint32, uint16) (uint32, error)
	GetMaxAltBlocks(*ethclient.Client) (uint8, error)
	GetMinStakeAmount(*ethclient.Client) (*big.Int, error)
	GetProposedBlock(*ethclient.Client, uint32, uint32) (bindings.StructsBlock, error)
	GetSortedProposedBlockIds(*ethclient.Client, uint32) ([]uint32, error)
	GetBlockManagerWithOpts(*ethclient.Client) (*bindings.BlockManager, bind.CallOpts)
	GetStakeManager(*ethclient.Client) *bindings.StakeManager
	GetStakeManagerWithOpts(*ethclient.Client) (*bindings.StakeManager, bind.CallOpts)
	GetStaker(*ethclient.Client, uint32) (bindings.StructsStaker, error)
	GetStake(*ethclient.Client, uint32) (*big.Int, error)
	GetStakerId(*ethclient.Client, string) (uint32, error)
	GetNumberOfStakers(*ethclient.Client) (uint32, error)
	GetLock(*ethclient.Client, string, uint32) (types.Locks, error)
	GetWithdrawReleasePeriod(*ethclient.Client) (uint8, error)
	GetMaxCommission(*ethclient.Client) (uint8, error)
	GetEpochLimitForUpdateCommission(*ethclient.Client) (uint16, error)
	GetVoteManagerWithOpts(*ethclient.Client) (*bindings.VoteManager, bind.CallOpts)
	GetCommitments(*ethclient.Client, string) ([32]byte, error)
	GetVoteValue(*ethclient.Client, uint16, uint32) (*big.Int, error)
	GetVotes(*ethclient.Client, uint32) (bindings.StructsVote, error)
	GetInfluenceSnapshot(*ethclient.Client, uint32, uint32) (*big.Int, error)
	GetStakeSnapshot(*ethclient.Client, uint32, uint32) (*big.Int, error)
	GetTotalInfluenceRevealed(*ethclient.Client, uint32) (*big.Int, error)
	GetRandaoHash(*ethclient.Client) ([32]byte, error)
	GetEpochLastCommitted(*ethclient.Client, uint32) (uint32, error)
	GetEpochLastRevealed(*ethclient.Client, uint32) (uint32, error)
	GetVoteManager(*ethclient.Client) *bindings.VoteManager
	GetAssetManager(*ethclient.Client) *bindings.AssetManager
	GetAssetManagerWithOpts(*ethclient.Client) (*bindings.AssetManager, bind.CallOpts)
	GetNumAssets(*ethclient.Client) (uint16, error)
	GetActiveJob(*ethclient.Client, uint16) (bindings.StructsJob, error)
	GetAssetType(*ethclient.Client, uint16) (uint8, error)
	GetCollection(*ethclient.Client, uint16) (bindings.StructsCollection, error)
	GetActiveCollection(*ethclient.Client, uint16) (bindings.StructsCollection, error)
	Aggregate(*ethclient.Client, uint32, bindings.StructsCollection) (*big.Int, error)
	GetDataToCommitFromJobs([]bindings.StructsJob) ([]*big.Int, []uint8, error)
	GetDataToCommitFromJob(bindings.StructsJob) (*big.Int, error)
	GetNumActiveAssets(*ethclient.Client) (*big.Int, error)
	GetActiveAssetsData(*ethclient.Client, uint32) ([]*big.Int, error)
	GetJobs(*ethclient.Client) ([]bindings.StructsJob, error)
	GetCollections(*ethclient.Client) ([]bindings.StructsCollection, error)
	GetActiveAssetIds(*ethclient.Client) ([]uint16, error)
	GetDataFromAPI(string) ([]byte, error)
	GetDataFromJSON(map[string]interface{}, string) (interface{}, error)
	GetDataFromXHTML(string, string) (string, error)
	ConnectToClient(string) *ethclient.Client
	FetchBalance(*ethclient.Client, string) (*big.Int, error)
	GetDelayedState(*ethclient.Client, int32) (int64, error)
	WaitForBlockCompletion(*ethclient.Client, string) int
	CheckEthBalanceIsZero(*ethclient.Client, string)
	AssignStakerId(*pflag.FlagSet, *ethclient.Client, string) (uint32, error)
	GetEpoch(*ethclient.Client) (uint32, error)
	SaveDataToFile(string, uint32, []*big.Int) error
	ReadDataFromFile(string) (uint32, []*big.Int, error)
	CalculateBlockTime(*ethclient.Client) int64
	IsFlagPassed(string) bool
	GetTokenManager(*ethclient.Client) *bindings.RAZOR
	GetStakedToken(*ethclient.Client, common.Address) *bindings.StakedToken
	GetUint32(*pflag.FlagSet, string) (uint32, error)
	WaitTillNextNSecs(int32)
	ReadJSONData(string) (map[string]*types.StructsJob, error)
	WriteDataToJSON(string, map[string]*types.StructsJob) error
	DeleteJobFromJSON(string, string) error
	AddJobToJSON(string, *types.StructsJob) error
	CheckTransactionReceipt(*ethclient.Client, string) int
	GetStakerSRZRBalance(*ethclient.Client, bindings.StructsStaker) (*big.Int, error)
	GetRemainingTimeOfCurrentState(*ethclient.Client) (int64, error)
	ConvertToNumber(interface{}) (*big.Float, error)
}

type EthClientUtils interface {
	Dial(string) (*ethclient.Client, error)
}

type ClientUtils interface {
	TransactionReceipt(*ethclient.Client, context.Context, common.Hash) (*Types.Receipt, error)
	BalanceAt(*ethclient.Client, context.Context, common.Address, *big.Int) (*big.Int, error)
	HeaderByNumber(*ethclient.Client, context.Context, *big.Int) (*Types.Header, error)
	PendingNonceAt(*ethclient.Client, context.Context, common.Address) (uint64, error)
	SuggestGasPrice(*ethclient.Client, context.Context) (*big.Int, error)
	EstimateGas(*ethclient.Client, context.Context, ethereum.CallMsg) (uint64, error)
	FilterLogs(*ethclient.Client, context.Context, ethereum.FilterQuery) ([]Types.Log, error)
}

type TimeUtils interface {
	Sleep(time.Duration)
}

type OSUtils interface {
	OpenFile(string, int, fs.FileMode) (*os.File, error)
	Open(string) (*os.File, error)
}

type BufioUtils interface {
	NewScanner(r io.Reader) *bufio.Scanner
}

type CoinUtils interface {
	BalanceOf(*bindings.RAZOR, *bind.CallOpts, common.Address) (*big.Int, error)
}

type IoutilUtils interface {
	ReadAll(io.ReadCloser) ([]byte, error)
	ReadFile(string) ([]byte, error)
	WriteFile(string, []byte, fs.FileMode) error
}

type ABIUtils interface {
	Parse(io.Reader) (abi.ABI, error)
	Pack(abi.ABI, string, ...interface{}) ([]byte, error)
}

type PathUtils interface {
	GetDefaultPath() (string, error)
	GetJobFilePath() (string, error)
}

type BindUtils interface {
	NewKeyedTransactorWithChainID(*ecdsa.PrivateKey, *big.Int) (*bind.TransactOpts, error)
}

type AccountsUtils interface {
	GetPrivateKey(string, string, string, accounts.AccountInterface) *ecdsa.PrivateKey
}

type BlockManagerUtils interface {
	GetNumProposedBlocks(*ethclient.Client, *bind.CallOpts, uint32) (uint8, error)
	GetProposedBlock(*ethclient.Client, *bind.CallOpts, uint32, uint32) (bindings.StructsBlock, error)
	GetBlock(*ethclient.Client, *bind.CallOpts, uint32) (bindings.StructsBlock, error)
	MinStake(*ethclient.Client, *bind.CallOpts) (*big.Int, error)
	MaxAltBlocks(*ethclient.Client, *bind.CallOpts) (uint8, error)
	SortedProposedBlockIds(*ethclient.Client, *bind.CallOpts, uint32, *big.Int) (uint32, error)
}

type StakeManagerUtils interface {
	GetStakerId(*ethclient.Client, *bind.CallOpts, common.Address) (uint32, error)
	GetStaker(*ethclient.Client, *bind.CallOpts, uint32) (bindings.StructsStaker, error)
	GetNumStakers(*ethclient.Client, *bind.CallOpts) (uint32, error)
	Locks(*ethclient.Client, *bind.CallOpts, common.Address, common.Address) (types.Locks, error)
	WithdrawReleasePeriod(*ethclient.Client, *bind.CallOpts) (uint8, error)
	MaxCommission(*ethclient.Client, *bind.CallOpts) (uint8, error)
	EpochLimitForUpdateCommission(*ethclient.Client, *bind.CallOpts) (uint16, error)
}

type AssetManagerUtils interface {
	GetNumAssets(*ethclient.Client, *bind.CallOpts) (uint16, error)
	GetNumActiveCollections(*ethclient.Client, *bind.CallOpts) (*big.Int, error)
	GetAsset(*ethclient.Client, *bind.CallOpts, uint16) (types.Asset, error)
	GetActiveCollections(*ethclient.Client, *bind.CallOpts) ([]uint16, error)
	Jobs(*ethclient.Client, *bind.CallOpts, uint16) (bindings.StructsJob, error)
}

type VoteManagerUtils interface {
	Commitments(*ethclient.Client, *bind.CallOpts, uint32) (types.Commitment, error)
	GetVoteValue(*ethclient.Client, *bind.CallOpts, uint16, uint32) (*big.Int, error)
	GetVote(*ethclient.Client, *bind.CallOpts, uint32) (bindings.StructsVote, error)
	GetInfluenceSnapshot(*ethclient.Client, *bind.CallOpts, uint32, uint32) (*big.Int, error)
	GetStakeSnapshot(*ethclient.Client, *bind.CallOpts, uint32, uint32) (*big.Int, error)
	GetTotalInfluenceRevealed(*ethclient.Client, *bind.CallOpts, uint32) (*big.Int, error)
	GetRandaoHash(*ethclient.Client, *bind.CallOpts) ([32]byte, error)
	GetEpochLastCommitted(*ethclient.Client, *bind.CallOpts, uint32) (uint32, error)
	GetEpochLastRevealed(*ethclient.Client, *bind.CallOpts, uint32) (uint32, error)
}

type BindingsUtils interface {
	NewAssetManager(common.Address, *ethclient.Client) (*bindings.AssetManager, error)
	NewRAZOR(common.Address, *ethclient.Client) (*bindings.RAZOR, error)
	NewStakeManager(common.Address, *ethclient.Client) (*bindings.StakeManager, error)
	NewVoteManager(common.Address, *ethclient.Client) (*bindings.VoteManager, error)
	NewBlockManager(common.Address, *ethclient.Client) (*bindings.BlockManager, error)
	NewStakedToken(common.Address, *ethclient.Client) (*bindings.StakedToken, error)
}

type JsonUtils interface {
	Unmarshal([]byte, interface{}) error
	Marshal(interface{}) ([]byte, error)
}

type StakedTokenUtils interface {
	BalanceOf(*bindings.StakedToken, *bind.CallOpts, common.Address) (*big.Int, error)
}

type RetryUtils interface {
	RetryAttempts(uint) retry.Option
}

type UtilsStruct struct{}
type EthClientStruct struct{}
type ClientStruct struct{}
type TimeStruct struct{}
type OSStruct struct{}
type BufioStruct struct{}
type CoinStruct struct{}
type IoutilStruct struct{}
type ABIStruct struct{}
type PathStruct struct{}
type BindStruct struct{}
type AccountsStruct struct{}
type BlockManagerStruct struct{}
type StakeManagerStruct struct{}
type AssetManagerStruct struct{}
type VoteManagerStruct struct{}
type BindingsStruct struct{}
type JsonStruct struct{}
type StakedTokenStruct struct{}
type RetryStruct struct{}

type OptionsPackageStruct struct {
	UtilsInterface        Utils
	EthClient             EthClientUtils
	ClientInterface       ClientUtils
	Time                  TimeUtils
	OS                    OSUtils
	Bufio                 BufioUtils
	CoinInterface         CoinUtils
	IoutilInterface       IoutilUtils
	ABIInterface          ABIUtils
	PathInterface         PathUtils
	BindInterface         BindUtils
	AccountsInterface     AccountsUtils
	BlockManagerInterface BlockManagerUtils
	StakeManagerInterface StakeManagerUtils
	AssetManagerInterface AssetManagerUtils
	VoteManagerInterface  VoteManagerUtils
	BindingsInterface     BindingsUtils
	JsonInterface         JsonUtils
	StakedTokenInterface  StakedTokenUtils
	RetryInterface        RetryUtils
}
