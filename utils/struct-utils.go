//Package utils provides the utils functions
package utils

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"io"
	"io/fs"
	"io/ioutil"
	"math/big"
	"os"
	"razor/accounts"
	coretypes "razor/core/types"
	"razor/path"
	"razor/pkg/bindings"
	"time"

	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
)

//This function initialises the Razor
func StartRazor(optionsPackageStruct OptionsPackageStruct) Utils {
	UtilsInterface = optionsPackageStruct.UtilsInterface
	EthClient = optionsPackageStruct.EthClient
	ClientInterface = optionsPackageStruct.ClientInterface
	Time = optionsPackageStruct.Time
	OS = optionsPackageStruct.OS
	Bufio = optionsPackageStruct.Bufio
	CoinInterface = optionsPackageStruct.CoinInterface
	IoutilInterface = optionsPackageStruct.IoutilInterface
	ABIInterface = optionsPackageStruct.ABIInterface
	PathInterface = optionsPackageStruct.PathInterface
	BindInterface = optionsPackageStruct.BindInterface
	AccountsInterface = optionsPackageStruct.AccountsInterface
	BlockManagerInterface = optionsPackageStruct.BlockManagerInterface
	StakeManagerInterface = optionsPackageStruct.StakeManagerInterface
	AssetManagerInterface = optionsPackageStruct.AssetManagerInterface
	BondManagerInterface = optionsPackageStruct.BondManagerInterface
	VoteManagerInterface = optionsPackageStruct.VoteManagerInterface
	BindingsInterface = optionsPackageStruct.BindingsInterface
	JsonInterface = optionsPackageStruct.JsonInterface
	StakedTokenInterface = optionsPackageStruct.StakedTokenInterface
	RetryInterface = optionsPackageStruct.RetryInterface
	MerkleInterface = optionsPackageStruct.MerkleInterface
	FlagSetInterface = optionsPackageStruct.FlagSetInterface
	return &UtilsStruct{}
}

//This function returns the block Index to be confirmed
func (b BlockManagerStruct) GetBlockIndexToBeConfirmed(client *ethclient.Client) (int8, error) {
	blockManager, opts := UtilsInterface.GetBlockManagerWithOpts(client)
	return blockManager.BlockIndexToBeConfirmed(&opts)
}

//This function returns the withdraw initiation period
func (s StakeManagerStruct) WithdrawInitiationPeriod(client *ethclient.Client) (uint8, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	return stakeManager.WithdrawInitiationPeriod(&opts)
}

//This function returns the withdraw lock period
func (s StakeManagerStruct) WithdrawLockPeriod(client *ethclient.Client) (uint8, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	return stakeManager.WithdrawLockPeriod(&opts)
}

//This function returns the number of jobs
func (a AssetManagerStruct) GetNumJobs(client *ethclient.Client) (uint16, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	return collectionManager.GetNumJobs(&opts)
}

//This function returns the collection
func (a AssetManagerStruct) GetCollection(client *ethclient.Client, id uint16) (bindings.StructsCollection, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	return collectionManager.GetCollection(&opts, id)
}

//This function returns the job
func (a AssetManagerStruct) GetJob(client *ethclient.Client, id uint16) (bindings.StructsJob, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	return collectionManager.GetJob(&opts, id)
}

//This function returns the collection Id from index
func (a AssetManagerStruct) GetCollectionIdFromIndex(client *ethclient.Client, index uint16) (uint16, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	return collectionManager.LeafIdToCollectionIdRegistry(&opts, index)
}

//This function returns the collection Id from leaf Id
func (a AssetManagerStruct) GetCollectionIdFromLeafId(client *ethclient.Client, leafId uint16) (uint16, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	return collectionManager.GetCollectionIdFromLeafId(&opts, leafId)
}

//This function returns the leaf Id of a collection
func (a AssetManagerStruct) GetLeafIdOfACollection(client *ethclient.Client, collectionId uint16) (uint16, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	return collectionManager.GetLeafIdOfCollection(&opts, collectionId)
}

//This function returns where to assign
func (v VoteManagerStruct) ToAssign(client *ethclient.Client) (uint16, error) {
	voteManager, opts := UtilsInterface.GetVoteManagerWithOpts(client)
	return voteManager.ToAssign(&opts)
}

//This function retusn the salt from blockchain
func (v VoteManagerStruct) GetSaltFromBlockchain(client *ethclient.Client) ([32]byte, error) {
	voteManager, opts := UtilsInterface.GetVoteManagerWithOpts(client)
	return voteManager.GetSalt(&opts)
}

//This function returns the private key
func (a AccountsStruct) GetPrivateKey(address string, password string, keystorePath string) *ecdsa.PrivateKey {
	return accounts.AccountUtilsInterface.GetPrivateKey(address, password, keystorePath)
}

//This function returns the number if proposed blocks
func (b BlockManagerStruct) GetNumProposedBlocks(client *ethclient.Client, epoch uint32) (uint8, error) {
	blockManager, opts := UtilsInterface.GetBlockManagerWithOpts(client)
	return blockManager.GetNumProposedBlocks(&opts, epoch)
}

//This function returns the proposed block
func (b BlockManagerStruct) GetProposedBlock(client *ethclient.Client, epoch uint32, proposedBlock uint32) (bindings.StructsBlock, error) {
	blockManager, opts := UtilsInterface.GetBlockManagerWithOpts(client)
	return blockManager.GetProposedBlock(&opts, epoch, proposedBlock)
}

//This function returns the block
func (b BlockManagerStruct) GetBlock(client *ethclient.Client, epoch uint32) (bindings.StructsBlock, error) {
	blockManager, opts := UtilsInterface.GetBlockManagerWithOpts(client)
	return blockManager.GetBlock(&opts, epoch)
}

//This function returns the minimum stake
func (b BlockManagerStruct) MinStake(client *ethclient.Client) (*big.Int, error) {
	blockManager, opts := UtilsInterface.GetBlockManagerWithOpts(client)
	return blockManager.MinStake(&opts)
}

//This function returns the maximum alt blocks
func (b BlockManagerStruct) MaxAltBlocks(client *ethclient.Client) (uint8, error) {
	blockManager, opts := UtilsInterface.GetBlockManagerWithOpts(client)
	return blockManager.MaxAltBlocks(&opts)
}

//This function returns the sorted proposed block Ids
func (b BlockManagerStruct) SortedProposedBlockIds(client *ethclient.Client, arg0 uint32, arg1 *big.Int) (uint32, error) {
	blockManager, opts := UtilsInterface.GetBlockManagerWithOpts(client)
	return blockManager.SortedProposedBlockIds(&opts, arg0, arg1)
}

//This function returns the stakerId
func (s StakeManagerStruct) GetStakerId(client *ethclient.Client, address common.Address) (uint32, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	return stakeManager.GetStakerId(&opts, address)
}

//This function returns the number of stakers
func (s StakeManagerStruct) GetNumStakers(client *ethclient.Client) (uint32, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	return stakeManager.GetNumStakers(&opts)
}

//This function returns the minimum safe razor
func (s StakeManagerStruct) MinSafeRazor(client *ethclient.Client) (*big.Int, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	return stakeManager.MinSafeRazor(&opts)
}

//This function returns the locks
func (s StakeManagerStruct) Locks(client *ethclient.Client, address common.Address, address1 common.Address, lockType uint8) (coretypes.Locks, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	return stakeManager.Locks(&opts, address, address1, lockType)
}

//This function returns the maximum commission
func (s StakeManagerStruct) MaxCommission(client *ethclient.Client) (uint8, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	return stakeManager.MaxCommission(&opts)
}

//This function returns the epoch limit for update commission
func (s StakeManagerStruct) EpochLimitForUpdateCommission(client *ethclient.Client) (uint16, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	return stakeManager.EpochLimitForUpdateCommission(&opts)
}

//This function returns the staker
func (s StakeManagerStruct) GetStaker(client *ethclient.Client, stakerId uint32) (bindings.StructsStaker, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	return stakeManager.GetStaker(&opts, stakerId)
}

//This function returns the number of collection
func (a AssetManagerStruct) GetNumCollections(client *ethclient.Client) (uint16, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	return collectionManager.GetNumCollections(&opts)
}

//This function returns the number of active collections
func (a AssetManagerStruct) GetNumActiveCollections(client *ethclient.Client) (uint16, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	return collectionManager.GetNumActiveCollections(&opts)
}

//This function returns the active collection
func (a AssetManagerStruct) GetActiveCollections(client *ethclient.Client) ([]uint16, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	return collectionManager.GetActiveCollections(&opts)
}

//This function returns the jobs
func (a AssetManagerStruct) Jobs(client *ethclient.Client, id uint16) (bindings.StructsJob, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	return collectionManager.Jobs(&opts, id)
}

//This function returns the commitments
func (v VoteManagerStruct) Commitments(client *ethclient.Client, stakerId uint32) (coretypes.Commitment, error) {
	voteManager, opts := UtilsInterface.GetVoteManagerWithOpts(client)
	return voteManager.Commitments(&opts, stakerId)
}

//This function returns the vote value
func (v VoteManagerStruct) GetVoteValue(client *ethclient.Client, epoch uint32, stakerId uint32, medianIndex uint16) (*big.Int, error) {
	voteManager, opts := UtilsInterface.GetVoteManagerWithOpts(client)
	return voteManager.GetVoteValue(&opts, epoch, stakerId, medianIndex)
}

//This function returns the influence snapshot
func (v VoteManagerStruct) GetInfluenceSnapshot(client *ethclient.Client, epoch uint32, stakerId uint32) (*big.Int, error) {
	voteManager, opts := UtilsInterface.GetVoteManagerWithOpts(client)
	return voteManager.GetInfluenceSnapshot(&opts, epoch, stakerId)
}

//This function returns the stake snapshot
func (v VoteManagerStruct) GetStakeSnapshot(client *ethclient.Client, epoch uint32, stakerId uint32) (*big.Int, error) {
	voteManager, opts := UtilsInterface.GetVoteManagerWithOpts(client)
	return voteManager.GetStakeSnapshot(&opts, epoch, stakerId)
}

//This function returns the total influence revealed
func (v VoteManagerStruct) GetTotalInfluenceRevealed(client *ethclient.Client, epoch uint32, medianIndex uint16) (*big.Int, error) {
	voteManager, opts := UtilsInterface.GetVoteManagerWithOpts(client)
	return voteManager.GetTotalInfluenceRevealed(&opts, epoch, medianIndex)
}

//This function returns the spoch last committed
func (v VoteManagerStruct) GetEpochLastCommitted(client *ethclient.Client, stakerId uint32) (uint32, error) {
	voteManager, opts := UtilsInterface.GetVoteManagerWithOpts(client)
	return voteManager.GetEpochLastCommitted(&opts, stakerId)
}

//This function returns the epoch last revealed
func (v VoteManagerStruct) GetEpochLastRevealed(client *ethclient.Client, stakerId uint32) (uint32, error) {
	voteManager, opts := UtilsInterface.GetVoteManagerWithOpts(client)
	return voteManager.GetEpochLastRevealed(&opts, stakerId)
}

//This function returns the new collectiion manager
func (b BindingsStruct) NewCollectionManager(address common.Address, client *ethclient.Client) (*bindings.CollectionManager, error) {
	return bindings.NewCollectionManager(address, client)
}

//This function returns the new RAZOR
func (b BindingsStruct) NewRAZOR(address common.Address, client *ethclient.Client) (*bindings.RAZOR, error) {
	return bindings.NewRAZOR(address, client)
}

//This function returns the new stake manager
func (b BindingsStruct) NewStakeManager(address common.Address, client *ethclient.Client) (*bindings.StakeManager, error) {
	return bindings.NewStakeManager(address, client)
}

//This function returns the new vote manager
func (b BindingsStruct) NewVoteManager(address common.Address, client *ethclient.Client) (*bindings.VoteManager, error) {
	return bindings.NewVoteManager(address, client)
}

//This function returns the new block manager
func (b BindingsStruct) NewBlockManager(address common.Address, client *ethclient.Client) (*bindings.BlockManager, error) {
	return bindings.NewBlockManager(address, client)
}

//This function returns  the new staked token
func (b BindingsStruct) NewStakedToken(address common.Address, client *ethclient.Client) (*bindings.StakedToken, error) {
	return bindings.NewStakedToken(address, client)
}

//This function unmarshales the data
func (j JsonStruct) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

//This function marshals the interface
func (j JsonStruct) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

//This function returns the Uint32
func (u UtilsStruct) GetUint32(flagSet *pflag.FlagSet, name string) (uint32, error) {
	return flagSet.GetUint32(name)
}

//This function dials on a url
func (e EthClientStruct) Dial(rawurl string) (*ethclient.Client, error) {
	return ethclient.Dial(rawurl)
}

//This function sleeps for a particular duration
func (t TimeStruct) Sleep(duration time.Duration) {
	time.Sleep(duration)
}

//This function is used to open the file and this is generalized open call
func (o OSStruct) OpenFile(name string, flag int, perm fs.FileMode) (*os.File, error) {
	return os.OpenFile(name, flag, perm)
}

//This function is used to open the file
func (o OSStruct) Open(name string) (*os.File, error) {
	return os.Open(name)
}

//This function is used to write on the file
func (o OSStruct) WriteFile(name string, data []byte, perm fs.FileMode) error {
	return os.WriteFile(name, data, perm)
}

//This function returns the transaction receipt
func (c ClientStruct) TransactionReceipt(client *ethclient.Client, ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return client.TransactionReceipt(ctx, txHash)
}

//This function returns the balance of client at particular block number
func (c ClientStruct) BalanceAt(client *ethclient.Client, ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return client.BalanceAt(ctx, account, blockNumber)
}

//This function returns the header by number
func (c ClientStruct) HeaderByNumber(client *ethclient.Client, ctx context.Context, number *big.Int) (*types.Header, error) {
	return client.HeaderByNumber(ctx, number)
}

//This function returns the pending nonce of particular account
func (c ClientStruct) PendingNonceAt(client *ethclient.Client, ctx context.Context, account common.Address) (uint64, error) {
	return client.PendingNonceAt(ctx, account)
}

//This function suggests the gas price
func (c ClientStruct) SuggestGasPrice(client *ethclient.Client, ctx context.Context) (*big.Int, error) {
	return client.SuggestGasPrice(ctx)
}

//This function estimates the gas
func (c ClientStruct) EstimateGas(client *ethclient.Client, ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	return client.EstimateGas(ctx, msg)
}

//This function filter the logs
func (c ClientStruct) FilterLogs(client *ethclient.Client, ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	return client.FilterLogs(ctx, q)
}

//NewScanner returns a new Scanner to read from r. The split function defaults to ScanLines.
func (b BufioStruct) NewScanner(r io.Reader) *bufio.Scanner {
	return bufio.NewScanner(r)
}

//This function returns the balance of account address
func (c CoinStruct) BalanceOf(coinContract *bindings.RAZOR, opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	return coinContract.BalanceOf(opts, account)
}

//This function is used to parse the data
func (a ABIStruct) Parse(reader io.Reader) (abi.ABI, error) {
	return abi.JSON(reader)
}

//This function is used to pack the parsed data
func (a ABIStruct) Pack(parsedData abi.ABI, name string, args ...interface{}) ([]byte, error) {
	return parsedData.Pack(name, args...)
}

//This function reads all the data from body
func (i IoutilStruct) ReadAll(body io.ReadCloser) ([]byte, error) {
	return ioutil.ReadAll(body)
}

//This function is used to read the file
func (i IoutilStruct) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

//This function is used to write data to file
func (i IoutilStruct) WriteFile(filename string, data []byte, perm fs.FileMode) error {
	return ioutil.WriteFile(filename, data, perm)
}

//This function returns the default path
func (p PathStruct) GetDefaultPath() (string, error) {
	return path.PathUtilsInterface.GetDefaultPath()
}

//This function returns the job file path
func (p PathStruct) GetJobFilePath() (string, error) {
	return path.PathUtilsInterface.GetJobFilePath()
}

//This function returns the keyes transactor with chain Id
func (b BindStruct) NewKeyedTransactorWithChainID(key *ecdsa.PrivateKey, chainID *big.Int) (*bind.TransactOpts, error) {
	return bind.NewKeyedTransactorWithChainID(key, chainID)
}

//This function returns the balance of account address
func (s StakedTokenStruct) BalanceOf(stakedToken *bindings.StakedToken, callOpts *bind.CallOpts, address common.Address) (*big.Int, error) {
	return stakedToken.BalanceOf(callOpts, address)
}

//This function helps in retrying the functionality of code
func (r RetryStruct) RetryAttempts(numberOfAttempts uint) retry.Option {
	return retry.Attempts(numberOfAttempts)
}

//This function returns the log file name
func (f FlagSetStruct) GetLogFileName(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("logFile")
}

func (b BlockManagerStruct) IsBlockConfirmed(client *ethclient.Client, epoch uint32) (bool, error) {
	blockManager, opts := UtilsInterface.GetBlockManagerWithOpts(client)
	return blockManager.IsBlockConfirmed(&opts, epoch)
}
