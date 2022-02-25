package utils

import (
	"context"
	"crypto/ecdsa"
	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"io"
	"io/fs"
	"math/big"
	"razor/accounts"
	"razor/core/types"
	"razor/pkg/bindings"
)

//go:generate mockery --name OptionUtils --output ./mocks/ --case=underscore
//go:generate mockery --name Utils --output ./mocks/ --case=underscore

var Options OptionUtils
var UtilsInterface Utils

type OptionUtils interface {
	Parse(io.Reader) (abi.ABI, error)
	Pack(abi.ABI, string, ...interface{}) ([]byte, error)
	GetDefaultPath() (string, error)
	GetJobFilePath() (string, error)
	GetPrivateKey(string, string, string, accounts.AccountInterface) *ecdsa.PrivateKey
	NewKeyedTransactorWithChainID(*ecdsa.PrivateKey, *big.Int) (*bind.TransactOpts, error)
	RetryAttempts(uint) retry.Option
	PendingNonceAt(*ethclient.Client, context.Context, common.Address) (uint64, error)
	HeaderByNumber(*ethclient.Client, context.Context, *big.Int) (*Types.Header, error)
	SuggestGasPrice(*ethclient.Client, context.Context) (*big.Int, error)
	EstimateGas(*ethclient.Client, context.Context, ethereum.CallMsg) (uint64, error)
	FilterLogs(*ethclient.Client, context.Context, ethereum.FilterQuery) ([]Types.Log, error)
	BalanceAt(*ethclient.Client, context.Context, common.Address, *big.Int) (*big.Int, error)
	GetNumProposedBlocks(*ethclient.Client, *bind.CallOpts, uint32) (uint8, error)
	GetProposedBlock(*ethclient.Client, *bind.CallOpts, uint32, uint32) (bindings.StructsBlock, error)
	GetBlock(*ethclient.Client, *bind.CallOpts, uint32) (bindings.StructsBlock, error)
	MinStake(*ethclient.Client, *bind.CallOpts) (*big.Int, error)
	MaxAltBlocks(*ethclient.Client, *bind.CallOpts) (uint8, error)
	SortedProposedBlockIds(*ethclient.Client, *bind.CallOpts, uint32, *big.Int) (uint32, error)
	GetStakerId(*ethclient.Client, *bind.CallOpts, common.Address) (uint32, error)
	GetStaker(*ethclient.Client, *bind.CallOpts, uint32) (bindings.StructsStaker, error)
	GetNumStakers(*ethclient.Client, *bind.CallOpts) (uint32, error)
	Locks(*ethclient.Client, *bind.CallOpts, common.Address, common.Address) (types.Locks, error)
	WithdrawReleasePeriod(*ethclient.Client, *bind.CallOpts) (uint8, error)
	MaxCommission(*ethclient.Client, *bind.CallOpts) (uint8, error)
	EpochLimitForUpdateCommission(*ethclient.Client, *bind.CallOpts) (uint16, error)
	Commitments(*ethclient.Client, *bind.CallOpts, uint32) (types.Commitment, error)
	GetVoteValue(*ethclient.Client, *bind.CallOpts, uint16, uint32) (*big.Int, error)
	GetVote(*ethclient.Client, *bind.CallOpts, uint32) (bindings.StructsVote, error)
	GetInfluenceSnapshot(*ethclient.Client, *bind.CallOpts, uint32, uint32) (*big.Int, error)
	GetStakeSnapshot(*ethclient.Client, *bind.CallOpts, uint32, uint32) (*big.Int, error)
	GetTotalInfluenceRevealed(*ethclient.Client, *bind.CallOpts, uint32) (*big.Int, error)
	GetRandaoHash(*ethclient.Client, *bind.CallOpts) ([32]byte, error)
	GetEpochLastCommitted(*ethclient.Client, *bind.CallOpts, uint32) (uint32, error)
	GetEpochLastRevealed(*ethclient.Client, *bind.CallOpts, uint32) (uint32, error)
	GetNumAssets(*ethclient.Client, *bind.CallOpts) (uint16, error)
	GetNumActiveCollections(*ethclient.Client, *bind.CallOpts) (*big.Int, error)
	GetAsset(*ethclient.Client, *bind.CallOpts, uint16) (types.Asset, error)
	GetActiveCollections(*ethclient.Client, *bind.CallOpts) ([]uint16, error)
	Jobs(*ethclient.Client, *bind.CallOpts, uint16) (bindings.StructsJob, error)
	ConvertToNumber(interface{}) (*big.Float, error)
	ReadAll(io.ReadCloser) ([]byte, error)
	NewAssetManager(common.Address, *ethclient.Client) (*bindings.AssetManager, error)
	NewRAZOR(address common.Address, client *ethclient.Client) (*bindings.RAZOR, error)
	NewStakeManager(address common.Address, client *ethclient.Client) (*bindings.StakeManager, error)
	NewVoteManager(address common.Address, client *ethclient.Client) (*bindings.VoteManager, error)
	NewBlockManager(address common.Address, client *ethclient.Client) (*bindings.BlockManager, error)
	NewStakedToken(address common.Address, client *ethclient.Client) (*bindings.StakedToken, error)
	ReadFile(filename string) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
	Marshal(v interface{}) ([]byte, error)
	WriteFile(filename string, data []byte, perm fs.FileMode) error
}

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
	GetDataFromHTML(string, string) (string, error)
	GetTokenManager(client *ethclient.Client) *bindings.RAZOR
	GetStakedToken(client *ethclient.Client, tokenAddress common.Address) *bindings.StakedToken
	ReadJSONData(fileName string) (map[string]*types.StructsJob, error)
	WriteDataToJSON(fileName string, data map[string]*types.StructsJob) error
	DeleteJobFromJSON(fileName string, jobId string) error
	AddJobToJSON(fileName string, job *types.StructsJob) error
	GetCollectionNames(client *ethclient.Client) ([]string, error)
	ConvertCustomJobToStructJob(customJob types.CustomJob) bindings.StructsJob
	GetCustomJobsFromJSONFile(collection string, jsonFileData string) ([]bindings.StructsJob, error)
	HandleOfficialJobsFromJSONFile(client *ethclient.Client, collection bindings.StructsCollection, dataString string) ([]bindings.StructsJob, []uint16)
}

type OptionsStruct struct{}
type UtilsStruct struct{}

type OptionsPackageStruct struct {
	Options        OptionUtils
	UtilsInterface Utils
}
