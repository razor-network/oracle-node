package utils

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"io"
	"io/fs"
	"math/big"
	"os"
	"razor/core/types"
	"razor/pkg/bindings"
	"time"

	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
)

//go:generate mockery --name Utils --output ./mocks/ --case=underscore
//go:generate mockery --name EthClientUtils --output ./mocks --case=underscore
//go:generate mockery --name ClientUtils --output ./mocks --case=underscore
//go:generate mockery --name TimeUtils --output ./mocks --case=underscore
//go:generate mockery --name OSUtils --output ./mocks --case=underscore
//go:generate mockery --name BufioUtils --output ./mocks --case=underscore
//go:generate mockery --name CoinUtils --output ./mocks --case=underscore
//go:generate mockery --name IOUtils --output ./mocks --case=underscore
//go:generate mockery --name ABIUtils --output ./mocks --case=underscore
//go:generate mockery --name PathUtils --output ./mocks --case=underscore
//go:generate mockery --name BindUtils --output ./mocks --case=underscore
//go:generate mockery --name AccountsUtils --output ./mocks --case=underscore
//go:generate mockery --name BlockManagerUtils --output ./mocks --case=underscore
//go:generate mockery --name AssetManagerUtils --output ./mocks --case=underscore
//go:generate mockery --name VoteManagerUtils --output ./mocks --case=underscore
//go:generate mockery --name StakeManagerUtils --output ./mocks --case=underscore
//go:generate mockery --name BindingsUtils --output ./mocks --case=underscore
//go:generate mockery --name JsonUtils --output ./mocks --case=underscore
//go:generate mockery --name StakedTokenUtils --output ./mocks --case=underscore
//go:generate mockery --name RetryUtils --output ./mocks --case=underscore
//go:generate mockery --name MerkleTreeInterface --output ./mocks --case=underscore
//go:generate mockery --name FlagSetUtils --output ./mocks --case=underscore

var UtilsInterface Utils
var EthClient EthClientUtils
var ClientInterface ClientUtils
var Time TimeUtils
var OS OSUtils
var Bufio BufioUtils
var CoinInterface CoinUtils
var IOInterface IOUtils
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
var MerkleInterface MerkleTreeInterface
var FlagSetInterface FlagSetUtils

type Utils interface {
	SuggestGasPriceWithRetry(client *ethclient.Client) (*big.Int, error)
	MultiplyFloatAndBigInt(bigIntVal *big.Int, floatingVal float64) *big.Int
	GetNonceAtWithRetry(client *ethclient.Client, accountAddress common.Address, blockNumber *big.Int) (uint64, error)
	GetGasPrice(client *ethclient.Client, config types.Configurations) *big.Int
	GetTxnOpts(transactionData types.TransactionOptions) *bind.TransactOpts
	GetGasLimit(transactionData types.TransactionOptions, txnOpts *bind.TransactOpts) (uint64, error)
	EstimateGasWithRetry(client *ethclient.Client, message ethereum.CallMsg) (uint64, error)
	IncreaseGasLimitValue(client *ethclient.Client, gasLimit uint64, gasLimitMultiplier float32) (uint64, error)
	GetLatestBlockWithRetry(client *ethclient.Client) (*Types.Header, error)
	FilterLogsWithRetry(client *ethclient.Client, query ethereum.FilterQuery) ([]Types.Log, error)
	BalanceAtWithRetry(client *ethclient.Client, account common.Address) (*big.Int, error)
	GetBlockManager(client *ethclient.Client) *bindings.BlockManager
	GetOptions() bind.CallOpts
	GetNumberOfProposedBlocks(client *ethclient.Client, epoch uint32) (uint8, error)
	GetSortedProposedBlockId(client *ethclient.Client, epoch uint32, index *big.Int) (uint32, error)
	FetchPreviousValue(client *ethclient.Client, epoch uint32, assetId uint16) (*big.Int, error)
	GetBlock(client *ethclient.Client, epoch uint32) (bindings.StructsBlock, error)
	GetMaxAltBlocks(client *ethclient.Client) (uint8, error)
	GetMinSafeRazor(client *ethclient.Client) (*big.Int, error)
	GetMinStakeAmount(client *ethclient.Client) (*big.Int, error)
	GetStateBuffer(client *ethclient.Client) (uint64, error)
	GetProposedBlock(client *ethclient.Client, epoch uint32, proposedBlockId uint32) (bindings.StructsBlock, error)
	GetSortedProposedBlockIds(client *ethclient.Client, epoch uint32) ([]uint32, error)
	GetBlockIndexToBeConfirmed(client *ethclient.Client) (int8, error)
	GetBlockManagerWithOpts(client *ethclient.Client) (*bindings.BlockManager, bind.CallOpts)
	GetStakeManager(client *ethclient.Client) *bindings.StakeManager
	GetStakeManagerWithOpts(client *ethclient.Client) (*bindings.StakeManager, bind.CallOpts)
	GetStaker(client *ethclient.Client, stakerId uint32) (bindings.StructsStaker, error)
	GetStake(client *ethclient.Client, stakerId uint32) (*big.Int, error)
	GetStakerId(client *ethclient.Client, address string) (uint32, error)
	GetNumberOfStakers(client *ethclient.Client) (uint32, error)
	GetLock(client *ethclient.Client, address string, stakerId uint32, lockType uint8) (types.Locks, error)
	GetWithdrawInitiationPeriod(client *ethclient.Client) (uint16, error)
	GetMaxCommission(client *ethclient.Client) (uint8, error)
	GetEpochLimitForUpdateCommission(client *ethclient.Client) (uint16, error)
	GetVoteManagerWithOpts(client *ethclient.Client) (*bindings.VoteManager, bind.CallOpts)
	GetCommitments(client *ethclient.Client, address string) ([32]byte, error)
	GetVoteValue(client *ethclient.Client, epoch uint32, stakerId uint32, medianIndex uint16) (*big.Int, error)
	GetInfluenceSnapshot(client *ethclient.Client, stakerId uint32, epoch uint32) (*big.Int, error)
	GetStakeSnapshot(client *ethclient.Client, stakerId uint32, epoch uint32) (*big.Int, error)
	GetTotalInfluenceRevealed(client *ethclient.Client, epoch uint32, medianIndex uint16) (*big.Int, error)
	GetEpochLastCommitted(client *ethclient.Client, stakerId uint32) (uint32, error)
	GetEpochLastRevealed(client *ethclient.Client, stakerId uint32) (uint32, error)
	GetVoteManager(client *ethclient.Client) *bindings.VoteManager
	GetCollectionManager(client *ethclient.Client) *bindings.CollectionManager
	GetCollectionManagerWithOpts(client *ethclient.Client) (*bindings.CollectionManager, bind.CallOpts)
	GetNumCollections(client *ethclient.Client) (uint16, error)
	GetActiveJob(client *ethclient.Client, jobId uint16) (bindings.StructsJob, error)
	GetCollection(client *ethclient.Client, collectionId uint16) (bindings.StructsCollection, error)
	GetActiveCollection(client *ethclient.Client, collectionId uint16) (bindings.StructsCollection, error)
	Aggregate(client *ethclient.Client, previousEpoch uint32, collection bindings.StructsCollection) (*big.Int, error)
	GetDataToCommitFromJobs(jobs []bindings.StructsJob) ([]*big.Int, []uint8, error)
	GetDataToCommitFromJob(job bindings.StructsJob) (*big.Int, error)
	GetAssignedCollections(client *ethclient.Client, numActiveCollections uint16, seed []byte) (map[int]bool, []*big.Int, error)
	GetLeafIdOfACollection(client *ethclient.Client, collectionId uint16) (uint16, error)
	GetCollectionIdFromIndex(client *ethclient.Client, medianIndex uint16) (uint16, error)
	GetCollectionIdFromLeafId(client *ethclient.Client, leafId uint16) (uint16, error)
	GetNumActiveCollections(client *ethclient.Client) (uint16, error)
	GetAggregatedDataOfCollection(client *ethclient.Client, collectionId uint16, epoch uint32) (*big.Int, error)
	GetJobs(client *ethclient.Client) ([]bindings.StructsJob, error)
	GetAllCollections(client *ethclient.Client) ([]bindings.StructsCollection, error)
	GetActiveCollectionIds(client *ethclient.Client) ([]uint16, error)
	GetDataFromAPI(url string) ([]byte, error)
	GetDataFromJSON(jsonObject map[string]interface{}, selector string) (interface{}, error)
	HandleOfficialJobsFromJSONFile(client *ethclient.Client, collection bindings.StructsCollection, dataString string) ([]bindings.StructsJob, []uint16)
	GetDataFromXHTML(url string, selector string) (string, error)
	ConnectToClient(provider string) *ethclient.Client
	FetchBalance(client *ethclient.Client, accountAddress string) (*big.Int, error)
	GetDelayedState(client *ethclient.Client, buffer int32) (int64, error)
	WaitForBlockCompletion(client *ethclient.Client, hashToRead string) error
	CheckEthBalanceIsZero(client *ethclient.Client, address string)
	AssignStakerId(flagSet *pflag.FlagSet, client *ethclient.Client, address string) (uint32, error)
	GetEpoch(client *ethclient.Client) (uint32, error)
	SaveDataToCommitJsonFile(filePath string, epoch uint32, commitData types.CommitData) error
	ReadFromCommitJsonFile(filePath string) (types.CommitFileData, error)
	SaveDataToProposeJsonFile(filePath string, proposeData types.ProposeFileData) error
	ReadFromProposeJsonFile(filePath string) (types.ProposeFileData, error)
	SaveDataToDisputeJsonFile(filePath string, bountyIdQueue []uint32) error
	ReadFromDisputeJsonFile(filePath string) (types.DisputeFileData, error)
	CalculateBlockTime(client *ethclient.Client) int64
	IsFlagPassed(name string) bool
	GetTokenManager(client *ethclient.Client) *bindings.RAZOR
	GetStakedToken(client *ethclient.Client, tokenAddress common.Address) *bindings.StakedToken
	GetUint32(flagSet *pflag.FlagSet, name string) (uint32, error)
	WaitTillNextNSecs(waitTime int32)
	ReadJSONData(fileName string) (map[string]*types.StructsJob, error)
	WriteDataToJSON(fileName string, data map[string]*types.StructsJob) error
	DeleteJobFromJSON(fileName string, jobId string) error
	AddJobToJSON(fileName string, job *types.StructsJob) error
	CheckTransactionReceipt(client *ethclient.Client, _txHash string) int
	CalculateSalt(epoch uint32, medians []*big.Int) [32]byte
	ToAssign(client *ethclient.Client) (uint16, error)
	Prng(max uint32, prngHashes []byte) *big.Int
	GetSaltFromBlockchain(client *ethclient.Client) ([32]byte, error)
	GetStakerSRZRBalance(client *ethclient.Client, staker bindings.StructsStaker) (*big.Int, error)
	GetRemainingTimeOfCurrentState(client *ethclient.Client, bufferPercent int32) (int64, error)
	ConvertToNumber(num interface{}) (*big.Float, error)
	SecondsToReadableTime(input int) string
	AssignLogFile(flagSet *pflag.FlagSet)
	CalculateBlockNumberAtEpochBeginning(client *ethclient.Client, epochLength int64, currentBlockNumber *big.Int) (*big.Int, error)
	GetStateName(stateNumber int64) string
	Shuffle(slice []uint32) []uint32
	GetEpochLastProposed(client *ethclient.Client, stakerId uint32) (uint32, error)
}

type EthClientUtils interface {
	Dial(rawurl string) (*ethclient.Client, error)
}

type ClientUtils interface {
	TransactionReceipt(client *ethclient.Client, ctx context.Context, txHash common.Hash) (*Types.Receipt, error)
	BalanceAt(client *ethclient.Client, ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error)
	HeaderByNumber(client *ethclient.Client, ctx context.Context, number *big.Int) (*Types.Header, error)
	NonceAt(client *ethclient.Client, ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error)
	SuggestGasPrice(client *ethclient.Client, ctx context.Context) (*big.Int, error)
	EstimateGas(client *ethclient.Client, ctx context.Context, msg ethereum.CallMsg) (uint64, error)
	FilterLogs(client *ethclient.Client, ctx context.Context, q ethereum.FilterQuery) ([]Types.Log, error)
}

type TimeUtils interface {
	Sleep(duration time.Duration)
}

type OSUtils interface {
	OpenFile(name string, flag int, perm fs.FileMode) (*os.File, error)
	Open(name string) (*os.File, error)
	WriteFile(name string, data []byte, perm fs.FileMode) error
	ReadFile(filename string) ([]byte, error)
}

type BufioUtils interface {
	NewScanner(r io.Reader) *bufio.Scanner
}

type CoinUtils interface {
	BalanceOf(coinContract *bindings.RAZOR, opts *bind.CallOpts, account common.Address) (*big.Int, error)
}

type MerkleTreeInterface interface {
	CreateMerkle(values []*big.Int) [][][]byte
	GetProofPath(tree [][][]byte, assetId uint16) [][32]byte
	GetMerkleRoot(tree [][][]byte) [32]byte
}
type IOUtils interface {
	ReadAll(body io.ReadCloser) ([]byte, error)
}

type ABIUtils interface {
	Parse(reader io.Reader) (abi.ABI, error)
	Pack(parsedData abi.ABI, name string, args ...interface{}) ([]byte, error)
}

type PathUtils interface {
	GetDefaultPath() (string, error)
	GetJobFilePath() (string, error)
}

type BindUtils interface {
	NewKeyedTransactorWithChainID(key *ecdsa.PrivateKey, chainID *big.Int) (*bind.TransactOpts, error)
}

type AccountsUtils interface {
	GetPrivateKey(address string, password string, keystorePath string) (*ecdsa.PrivateKey, error)
}

type BlockManagerUtils interface {
	GetNumProposedBlocks(client *ethclient.Client, epoch uint32) (uint8, error)
	GetProposedBlock(client *ethclient.Client, epoch uint32, proposedBlock uint32) (bindings.StructsBlock, error)
	GetBlock(client *ethclient.Client, epoch uint32) (bindings.StructsBlock, error)
	MinStake(client *ethclient.Client) (*big.Int, error)
	StateBuffer(client *ethclient.Client) (uint8, error)
	MaxAltBlocks(client *ethclient.Client) (uint8, error)
	SortedProposedBlockIds(client *ethclient.Client, arg0 uint32, arg1 *big.Int) (uint32, error)
	GetBlockIndexToBeConfirmed(client *ethclient.Client) (int8, error)
	GetEpochLastProposed(client *ethclient.Client, stakerId uint32) (uint32, error)
}

type StakeManagerUtils interface {
	GetStakerId(client *ethclient.Client, address common.Address) (uint32, error)
	GetStaker(client *ethclient.Client, stakerId uint32) (bindings.StructsStaker, error)
	GetNumStakers(client *ethclient.Client) (uint32, error)
	MinSafeRazor(client *ethclient.Client) (*big.Int, error)
	Locks(client *ethclient.Client, address common.Address, address1 common.Address, lockType uint8) (types.Locks, error)
	MaxCommission(client *ethclient.Client) (uint8, error)
	EpochLimitForUpdateCommission(client *ethclient.Client) (uint16, error)
	WithdrawInitiationPeriod(client *ethclient.Client) (uint16, error)
}

type AssetManagerUtils interface {
	GetNumCollections(client *ethclient.Client) (uint16, error)
	GetNumJobs(client *ethclient.Client) (uint16, error)
	GetNumActiveCollections(client *ethclient.Client) (uint16, error)
	GetJob(client *ethclient.Client, id uint16) (bindings.StructsJob, error)
	GetCollection(client *ethclient.Client, id uint16) (bindings.StructsCollection, error)
	GetActiveCollections(client *ethclient.Client) ([]uint16, error)
	Jobs(client *ethclient.Client, id uint16) (bindings.StructsJob, error)
	GetCollectionIdFromIndex(client *ethclient.Client, index uint16) (uint16, error)
	GetCollectionIdFromLeafId(client *ethclient.Client, leafId uint16) (uint16, error)
	GetLeafIdOfACollection(client *ethclient.Client, collectionId uint16) (uint16, error)
}

type VoteManagerUtils interface {
	Commitments(client *ethclient.Client, stakerId uint32) (types.Commitment, error)
	GetVoteValue(client *ethclient.Client, epoch uint32, stakerId uint32, medianIndex uint16) (*big.Int, error)
	GetInfluenceSnapshot(client *ethclient.Client, epoch uint32, stakerId uint32) (*big.Int, error)
	GetStakeSnapshot(client *ethclient.Client, epoch uint32, stakerId uint32) (*big.Int, error)
	GetTotalInfluenceRevealed(client *ethclient.Client, epoch uint32, medianIndex uint16) (*big.Int, error)
	GetEpochLastCommitted(client *ethclient.Client, stakerId uint32) (uint32, error)
	GetEpochLastRevealed(client *ethclient.Client, stakerId uint32) (uint32, error)
	ToAssign(client *ethclient.Client) (uint16, error)
	GetSaltFromBlockchain(client *ethclient.Client) ([32]byte, error)
}

type BindingsUtils interface {
	NewCollectionManager(address common.Address, client *ethclient.Client) (*bindings.CollectionManager, error)
	NewRAZOR(address common.Address, client *ethclient.Client) (*bindings.RAZOR, error)
	NewStakeManager(address common.Address, client *ethclient.Client) (*bindings.StakeManager, error)
	NewVoteManager(address common.Address, client *ethclient.Client) (*bindings.VoteManager, error)
	NewBlockManager(address common.Address, client *ethclient.Client) (*bindings.BlockManager, error)
	NewStakedToken(address common.Address, client *ethclient.Client) (*bindings.StakedToken, error)
}

type JsonUtils interface {
	Unmarshal(data []byte, v interface{}) error
	Marshal(v interface{}) ([]byte, error)
}

type StakedTokenUtils interface {
	BalanceOf(stakedToken *bindings.StakedToken, callOpts *bind.CallOpts, address common.Address) (*big.Int, error)
}

type RetryUtils interface {
	RetryAttempts(numberOfAttempts uint) retry.Option
}

type FlagSetUtils interface {
	GetLogFileName(flagSet *pflag.FlagSet) (string, error)
}

type UtilsStruct struct{}
type EthClientStruct struct{}
type ClientStruct struct{}
type TimeStruct struct{}
type OSStruct struct{}
type BufioStruct struct{}
type CoinStruct struct{}
type IOStruct struct{}
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
type MerkleTreeStruct struct{}
type FlagSetStruct struct{}

type OptionsPackageStruct struct {
	UtilsInterface        Utils
	EthClient             EthClientUtils
	ClientInterface       ClientUtils
	Time                  TimeUtils
	OS                    OSUtils
	Bufio                 BufioUtils
	CoinInterface         CoinUtils
	IOInterface           IOUtils
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
	MerkleInterface       MerkleTreeInterface
	FlagSetInterface      FlagSetUtils
}
