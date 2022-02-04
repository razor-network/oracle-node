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
	GetPrivateKey(address string, password string, keystorePath string, accountUtils accounts.AccountInterface) *ecdsa.PrivateKey
	NewKeyedTransactorWithChainID(key *ecdsa.PrivateKey, chainID *big.Int) (*bind.TransactOpts, error)
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
	GetNumAssets(*ethclient.Client, *bind.CallOpts) (uint16, error)
	GetNumActiveCollections(*ethclient.Client, *bind.CallOpts) (*big.Int, error)
	GetAsset(*ethclient.Client, *bind.CallOpts, uint16) (types.Asset, error)
	GetActiveCollections(*ethclient.Client, *bind.CallOpts) ([]uint16, error)
	Jobs(*ethclient.Client, *bind.CallOpts, uint16) (bindings.StructsJob, error)
	ReadJSONData(string) (map[string]*types.StructsJob, error)
	ConvertToNumber(interface{}) (*big.Float, error)
	ReadAll(io.ReadCloser) ([]byte, error)
	NewAssetManager(common.Address, *ethclient.Client) (*bindings.AssetManager, error)
	NewRAZOR(address common.Address, client *ethclient.Client) (*bindings.RAZOR, error)
	NewStakeManager(address common.Address, client *ethclient.Client) (*bindings.StakeManager, error)
	NewVoteManager(address common.Address, client *ethclient.Client) (*bindings.VoteManager, error)
	NewBlockManager(address common.Address, client *ethclient.Client) (*bindings.BlockManager, error)
	NewStakedToken(address common.Address, client *ethclient.Client) (*bindings.StakedToken, error)
	Dial(string) (*ethclient.Client, error)
	TransactionReceipt(*ethclient.Client, context.Context, common.Hash) (*Types.Receipt, error)
	OpenFile(string, int, fs.FileMode) (*os.File, error)
	Open(string) (*os.File, error)
	Atoi(string) (int, error)
	NewScanner(*os.File) *bufio.Scanner
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
	GetStakeManager(client *ethclient.Client) *bindings.StakeManager
	GetVoteManager(client *ethclient.Client) *bindings.VoteManager
	GetStakedToken(client *ethclient.Client, tokenAddress common.Address) *bindings.StakedToken
	ConnectToClient(provider string) *ethclient.Client
	FetchBalance(client *ethclient.Client, accountAddress string) (*big.Int, error)
	GetDelayedState(client *ethclient.Client, buffer int32) (int64, error)
	CheckTransactionReceipt(client *ethclient.Client, _txHash string) int
	WaitForBlockCompletion(client *ethclient.Client, hashToRead string) int
	CheckEthBalanceIsZero(client *ethclient.Client, address string)
	GetStateName(stateNumber int64) string
	AssignStakerId(flagSet *pflag.FlagSet, client *ethclient.Client, address string) (uint32, error)
	GetEpoch(client *ethclient.Client) (uint32, error)
	SaveCommittedDataToFile(fileName string, epoch uint32, committedData []*big.Int) error
	ReadCommittedDataFromFile(fileName string) (uint32, []*big.Int, error)
	CalculateBlockTime(client *ethclient.Client) int64
}

type OptionsStruct struct{}
type UtilsStruct struct{}

type OptionsPackageStruct struct {
	Options        OptionUtils
	UtilsInterface Utils
}
