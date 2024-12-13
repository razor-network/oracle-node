package utils

import (
	"context"
	"crypto/ecdsa"
	"io"
	"io/fs"
	"math/big"
	"os"
	"razor/cache"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/rpc"
	"time"

	RPC "github.com/ethereum/go-ethereum/rpc"

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
//go:generate mockery --name CoinUtils --output ./mocks --case=underscore
//go:generate mockery --name IOUtils --output ./mocks --case=underscore
//go:generate mockery --name ABIUtils --output ./mocks --case=underscore
//go:generate mockery --name PathUtils --output ./mocks --case=underscore
//go:generate mockery --name BindUtils --output ./mocks --case=underscore
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
//go:generate mockery --name GasUtils --output ./mocks --case=underscore
//go:generate mockery --name FileUtils --output ./mocks --case=underscore

var UtilsInterface Utils
var EthClient EthClientUtils
var ClientInterface ClientUtils
var Time TimeUtils
var OS OSUtils
var CoinInterface CoinUtils
var IOInterface IOUtils
var ABIInterface ABIUtils
var PathInterface PathUtils
var BindInterface BindUtils
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
var FileInterface FileUtils
var GasInterface GasUtils

type Utils interface {
	MultiplyFloatAndBigInt(bigIntVal *big.Int, floatingVal float64) *big.Int
	GetTxnOpts(rpcParameters rpc.RPCParameters, transactionData types.TransactionOptions) (*bind.TransactOpts, error)
	GetBlockManager(client *ethclient.Client) *bindings.BlockManager
	GetOptions() bind.CallOpts
	GetNumberOfProposedBlocks(rpcParameters rpc.RPCParameters, epoch uint32) (uint8, error)
	GetSortedProposedBlockId(rpcParameters rpc.RPCParameters, epoch uint32, index *big.Int) (uint32, error)
	FetchPreviousValue(rpcParameters rpc.RPCParameters, epoch uint32, assetId uint16) (*big.Int, error)
	GetBlock(rpcParameters rpc.RPCParameters, epoch uint32) (bindings.StructsBlock, error)
	GetMaxAltBlocks(rpcParameters rpc.RPCParameters) (uint8, error)
	GetMinSafeRazor(rpcParameters rpc.RPCParameters) (*big.Int, error)
	GetMinStakeAmount(rpcParameters rpc.RPCParameters) (*big.Int, error)
	GetStateBuffer(rpcParameters rpc.RPCParameters) (uint64, error)
	GetProposedBlock(rpcParameters rpc.RPCParameters, epoch uint32, proposedBlockId uint32) (bindings.StructsBlock, error)
	GetSortedProposedBlockIds(rpcParameters rpc.RPCParameters, epoch uint32) ([]uint32, error)
	GetBlockIndexToBeConfirmed(rpcParameters rpc.RPCParameters) (int8, error)
	GetBlockManagerWithOpts(client *ethclient.Client) (*bindings.BlockManager, bind.CallOpts)
	GetStakeManager(client *ethclient.Client) *bindings.StakeManager
	GetStakeManagerWithOpts(client *ethclient.Client) (*bindings.StakeManager, bind.CallOpts)
	GetStaker(rpcParameters rpc.RPCParameters, stakerId uint32) (bindings.StructsStaker, error)
	GetStake(rpcParameters rpc.RPCParameters, stakerId uint32) (*big.Int, error)
	GetStakerId(rpcParameters rpc.RPCParameters, address string) (uint32, error)
	GetNumberOfStakers(rpcParameters rpc.RPCParameters) (uint32, error)
	GetLock(rpcParameters rpc.RPCParameters, address string, stakerId uint32, lockType uint8) (types.Locks, error)
	GetWithdrawInitiationPeriod(rpcParameters rpc.RPCParameters) (uint16, error)
	GetMaxCommission(rpcParameters rpc.RPCParameters) (uint8, error)
	GetEpochLimitForUpdateCommission(rpcParameters rpc.RPCParameters) (uint16, error)
	StakerInfo(rpcParameters rpc.RPCParameters, stakerId uint32) (types.Staker, error)
	GetMaturity(rpcParameters rpc.RPCParameters, age uint32) (uint16, error)
	GetBountyLock(rpcParameters rpc.RPCParameters, bountyId uint32) (types.BountyLock, error)
	GetVoteManagerWithOpts(client *ethclient.Client) (*bindings.VoteManager, bind.CallOpts)
	GetCommitment(rpcParameters rpc.RPCParameters, address string) (types.Commitment, error)
	GetVoteValue(rpcParameters rpc.RPCParameters, epoch uint32, stakerId uint32, medianIndex uint16) (*big.Int, error)
	GetInfluenceSnapshot(rpcParameters rpc.RPCParameters, stakerId uint32, epoch uint32) (*big.Int, error)
	GetStakeSnapshot(rpcParameters rpc.RPCParameters, stakerId uint32, epoch uint32) (*big.Int, error)
	GetTotalInfluenceRevealed(rpcParameters rpc.RPCParameters, epoch uint32, medianIndex uint16) (*big.Int, error)
	GetEpochLastCommitted(rpcParameters rpc.RPCParameters, stakerId uint32) (uint32, error)
	GetEpochLastRevealed(rpcParameters rpc.RPCParameters, stakerId uint32) (uint32, error)
	GetVoteManager(client *ethclient.Client) *bindings.VoteManager
	GetCollectionManager(client *ethclient.Client) *bindings.CollectionManager
	GetCollectionManagerWithOpts(client *ethclient.Client) (*bindings.CollectionManager, bind.CallOpts)
	GetNumCollections(rpcParameters rpc.RPCParameters) (uint16, error)
	GetNumJobs(rpcParameters rpc.RPCParameters) (uint16, error)
	GetActiveJob(rpcParameters rpc.RPCParameters, jobId uint16) (bindings.StructsJob, error)
	GetCollection(rpcParameters rpc.RPCParameters, collectionId uint16) (bindings.StructsCollection, error)
	GetActiveCollection(collectionsCache *cache.CollectionsCache, collectionId uint16) (bindings.StructsCollection, error)
	Aggregate(rpcParameters rpc.RPCParameters, previousEpoch uint32, collection bindings.StructsCollection, commitParams *types.CommitParams) (*big.Int, error)
	GetDataToCommitFromJobs(jobs []bindings.StructsJob, commitParams *types.CommitParams) ([]*big.Int, []uint8)
	GetDataToCommitFromJob(job bindings.StructsJob, commitParams *types.CommitParams) (*big.Int, error)
	GetAssignedCollections(rpcParameters rpc.RPCParameters, numActiveCollections uint16, seed []byte) (map[int]bool, []*big.Int, error)
	GetLeafIdOfACollection(rpcParameters rpc.RPCParameters, collectionId uint16) (uint16, error)
	GetSaltFromBlockchain(rpcParameters rpc.RPCParameters) ([32]byte, error)
	GetCollectionIdFromIndex(rpcParameters rpc.RPCParameters, medianIndex uint16) (uint16, error)
	GetCollectionIdFromLeafId(rpcParameters rpc.RPCParameters, leafId uint16) (uint16, error)
	GetNumActiveCollections(rpcParameters rpc.RPCParameters) (uint16, error)
	GetAggregatedDataOfCollection(rpcParameters rpc.RPCParameters, collectionId uint16, epoch uint32, commitParams *types.CommitParams) (*big.Int, error)
	GetJobs(rpcParameters rpc.RPCParameters) ([]bindings.StructsJob, error)
	GetAllCollections(rpcParameters rpc.RPCParameters) ([]bindings.StructsCollection, error)
	GetActiveCollectionIds(rpcParameters rpc.RPCParameters) ([]uint16, error)
	GetActiveStatus(rpcParameters rpc.RPCParameters, id uint16) (bool, error)
	HandleOfficialJobsFromJSONFile(collection bindings.StructsCollection, dataString string, commitParams *types.CommitParams) ([]bindings.StructsJob, []uint16)
	ConnectToClient(provider string) *ethclient.Client
	FetchBalance(rpcParameters rpc.RPCParameters, accountAddress string) (*big.Int, error)
	Allowance(rpcParameters rpc.RPCParameters, owner common.Address, spender common.Address) (*big.Int, error)
	GetBufferedState(header *Types.Header, stateBuffer uint64, buffer int32) (int64, error)
	WaitForBlockCompletion(rpcManager rpc.RPCParameters, hashToRead string) error
	CheckEthBalanceIsZero(rpcParameters rpc.RPCParameters, address string)
	AssignStakerId(rpcParameters rpc.RPCParameters, flagSet *pflag.FlagSet, address string) (uint32, error)
	GetEpoch(rpcParameters rpc.RPCParameters) (uint32, error)
	CalculateBlockTime(rpcParameters rpc.RPCParameters) int64
	IsFlagPassed(name string) bool
	GetTokenManager(client *ethclient.Client) *bindings.RAZOR
	GetStakedToken(client *ethclient.Client, tokenAddress common.Address) *bindings.StakedToken
	GetUint32(flagSet *pflag.FlagSet, name string) (uint32, error)
	WaitTillNextNSecs(waitTime int32)
	ReadJSONData(fileName string) (map[string]*types.StructsJob, error)
	WriteDataToJSON(fileName string, data map[string]*types.StructsJob) error
	DeleteJobFromJSON(fileName string, jobId string) error
	AddJobToJSON(fileName string, job *types.StructsJob) error
	CheckTransactionReceipt(rpcManager rpc.RPCParameters, _txHash string) int
	CalculateSalt(epoch uint32, medians []*big.Int) [32]byte
	ToAssign(rpcParameters rpc.RPCParameters) (uint16, error)
	Prng(max uint32, prngHashes []byte) *big.Int
	GetRemainingTimeOfCurrentState(block *Types.Header, stateBuffer uint64, bufferPercent int32) (int64, error)
	SecondsToReadableTime(input int) string
	EstimateBlockNumberAtEpochBeginning(rpcParameters rpc.RPCParameters, currentBlockNumber *big.Int) (*big.Int, error)
	GetEpochLastProposed(rpcParameters rpc.RPCParameters, stakerId uint32) (uint32, error)
	GetConfirmedBlocks(rpcParameters rpc.RPCParameters, epoch uint32) (types.ConfirmedBlock, error)
	Disputes(rpcParameters rpc.RPCParameters, epoch uint32, address common.Address) (types.DisputesStruct, error)
	CheckAmountAndBalance(amountInWei *big.Int, balance *big.Int) *big.Int
	PasswordPrompt() string
	AssignPassword(flagSet *pflag.FlagSet) string
	PrivateKeyPrompt() string
	GetRogueRandomValue(value int) *big.Int
	GetStakedTokenManagerWithOpts(client *ethclient.Client, tokenAddress common.Address) (*bindings.StakedToken, bind.CallOpts)
	GetStakerSRZRBalance(rpcParameters rpc.RPCParameters, staker bindings.StructsStaker) (*big.Int, error)
	CheckPassword(account types.Account) error
	AccountManagerForKeystore() (types.AccountManagerInterface, error)
}

type EthClientUtils interface {
	Dial(rawurl string) (*ethclient.Client, error)
}

type ClientUtils interface {
	TransactionReceipt(client *ethclient.Client, ctx context.Context, txHash common.Hash) (*Types.Receipt, error)
	BalanceAt(client *ethclient.Client, ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error)
	HeaderByNumber(client *ethclient.Client, ctx context.Context, number *big.Int) (*Types.Header, error)
	NonceAt(client *ethclient.Client, ctx context.Context, account common.Address) (uint64, error)
	SuggestGasPrice(client *ethclient.Client, ctx context.Context) (*big.Int, error)
	EstimateGas(client *ethclient.Client, ctx context.Context, msg ethereum.CallMsg) (uint64, error)
	FilterLogs(client *ethclient.Client, ctx context.Context, q ethereum.FilterQuery) ([]Types.Log, error)
	SuggestGasPriceWithRetry(rpcParameters rpc.RPCParameters) (*big.Int, error)
	EstimateGasWithRetry(rpcParameters rpc.RPCParameters, message ethereum.CallMsg) (uint64, error)
	GetLatestBlockWithRetry(rpcParameters rpc.RPCParameters) (*Types.Header, error)
	GetBlockByNumberWithRetry(rpcParameters rpc.RPCParameters, blockNumber *big.Int) (*Types.Header, error)
	FilterLogsWithRetry(rpcParameters rpc.RPCParameters, query ethereum.FilterQuery) ([]Types.Log, error)
	BalanceAtWithRetry(rpcParameters rpc.RPCParameters, account common.Address) (*big.Int, error)
	GetNonceAtWithRetry(rpcParameters rpc.RPCParameters, accountAddress common.Address) (uint64, error)
	PerformBatchCall(rpcParameters rpc.RPCParameters, calls []RPC.BatchElem) error
	CreateBatchCalls(contractABI *abi.ABI, contractAddress, methodName string, args [][]interface{}) ([]RPC.BatchElem, error)
	BatchCall(rpcParameters rpc.RPCParameters, contractABI *abi.ABI, contractAddress, methodName string, args [][]interface{}) ([][]interface{}, error)
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

type CoinUtils interface {
	BalanceOf(client *ethclient.Client, account common.Address) (*big.Int, error)
	Allowance(client *ethclient.Client, owner common.Address, spender common.Address) (*big.Int, error)
}

type MerkleTreeInterface interface {
	CreateMerkle(values []*big.Int) ([][][]byte, error)
	GetProofPath(tree [][][]byte, assetId uint16) [][32]byte
	GetMerkleRoot(tree [][][]byte) ([32]byte, error)
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
	GetConfirmedBlocks(client *ethclient.Client, epoch uint32) (types.ConfirmedBlock, error)
	Disputes(client *ethclient.Client, epoch uint32, address common.Address) (types.DisputesStruct, error)
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
	StakerInfo(client *ethclient.Client, stakerId uint32) (types.Staker, error)
	GetMaturity(client *ethclient.Client, age uint32) (uint16, error)
	GetBountyLock(client *ethclient.Client, bountyId uint32) (types.BountyLock, error)
}

type AssetManagerUtils interface {
	GetNumCollections(client *ethclient.Client) (uint16, error)
	GetNumJobs(client *ethclient.Client) (uint16, error)
	GetNumActiveCollections(client *ethclient.Client) (uint16, error)
	GetJob(client *ethclient.Client, id uint16) (bindings.StructsJob, error)
	GetCollection(client *ethclient.Client, id uint16) (bindings.StructsCollection, error)
	GetActiveCollections(client *ethclient.Client) ([]uint16, error)
	GetActiveStatus(client *ethclient.Client, id uint16) (bool, error)
	Jobs(client *ethclient.Client, id uint16) (bindings.StructsJob, error)
	GetCollectionIdFromIndex(client *ethclient.Client, index uint16) (uint16, error)
	GetCollectionIdFromLeafId(client *ethclient.Client, leafId uint16) (uint16, error)
	GetLeafIdOfACollection(client *ethclient.Client, collectionId uint16) (uint16, error)
}

type VoteManagerUtils interface {
	GetCommitment(client *ethclient.Client, stakerId uint32) (types.Commitment, error)
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
	BalanceOf(client *ethclient.Client, tokenAddress common.Address, address common.Address) (*big.Int, error)
}

type RetryUtils interface {
	RetryAttempts(numberOfAttempts uint) retry.Option
}

type FlagSetUtils interface {
	GetLogFileName(flagSet *pflag.FlagSet) (string, error)
}

type FileUtils interface {
	SaveDataToCommitJsonFile(filePath string, epoch uint32, commitData types.CommitData, commitment [32]byte) error
	ReadFromCommitJsonFile(filePath string) (types.CommitFileData, error)
	SaveDataToProposeJsonFile(filePath string, proposeData types.ProposeFileData) error
	ReadFromProposeJsonFile(filePath string) (types.ProposeFileData, error)
	SaveDataToDisputeJsonFile(filePath string, bountyIdQueue []uint32) error
	ReadFromDisputeJsonFile(filePath string) (types.DisputeFileData, error)
	AssignLogFile(flagSet *pflag.FlagSet, configurations types.Configurations)
}

type GasUtils interface {
	GetGasPrice(rpcParameters rpc.RPCParameters, config types.Configurations) *big.Int
	GetGasLimit(rpcParameters rpc.RPCParameters, transactionData types.TransactionOptions, txnOpts *bind.TransactOpts) (uint64, error)
	IncreaseGasLimitValue(rpcParameters rpc.RPCParameters, gasLimit uint64, gasLimitMultiplier float32) (uint64, error)
}

type UtilsStruct struct{}
type EthClientStruct struct{}
type ClientStruct struct{}
type TimeStruct struct{}
type OSStruct struct{}
type CoinStruct struct{}
type IOStruct struct{}
type ABIStruct struct{}
type PathStruct struct{}
type BindStruct struct{}
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
type FileStruct struct{}
type GasStruct struct{}

type OptionsPackageStruct struct {
	UtilsInterface        Utils
	EthClient             EthClientUtils
	ClientInterface       ClientUtils
	Time                  TimeUtils
	OS                    OSUtils
	CoinInterface         CoinUtils
	IOInterface           IOUtils
	ABIInterface          ABIUtils
	PathInterface         PathUtils
	BindInterface         BindUtils
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
	FileInterface         FileUtils
	GasInterface          GasUtils
}
