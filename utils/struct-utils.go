package utils

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"io"
	"io/fs"
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

func StartRazor(optionsPackageStruct OptionsPackageStruct) Utils {
	UtilsInterface = optionsPackageStruct.UtilsInterface
	EthClient = optionsPackageStruct.EthClient
	ClientInterface = optionsPackageStruct.ClientInterface
	Time = optionsPackageStruct.Time
	OS = optionsPackageStruct.OS
	Bufio = optionsPackageStruct.Bufio
	CoinInterface = optionsPackageStruct.CoinInterface
	IOInterface = optionsPackageStruct.IOInterface
	ABIInterface = optionsPackageStruct.ABIInterface
	PathInterface = optionsPackageStruct.PathInterface
	BindInterface = optionsPackageStruct.BindInterface
	AccountsInterface = optionsPackageStruct.AccountsInterface
	BlockManagerInterface = optionsPackageStruct.BlockManagerInterface
	StakeManagerInterface = optionsPackageStruct.StakeManagerInterface
	AssetManagerInterface = optionsPackageStruct.AssetManagerInterface
	VoteManagerInterface = optionsPackageStruct.VoteManagerInterface
	BindingsInterface = optionsPackageStruct.BindingsInterface
	JsonInterface = optionsPackageStruct.JsonInterface
	StakedTokenInterface = optionsPackageStruct.StakedTokenInterface
	RetryInterface = optionsPackageStruct.RetryInterface
	MerkleInterface = optionsPackageStruct.MerkleInterface
	FlagSetInterface = optionsPackageStruct.FlagSetInterface
	return &UtilsStruct{}
}

func (b BlockManagerStruct) GetBlockIndexToBeConfirmed(client *ethclient.Client) (int8, error) {
	blockManager, opts := UtilsInterface.GetBlockManagerWithOpts(client)
	return blockManager.BlockIndexToBeConfirmed(&opts)
}

func (s StakeManagerStruct) WithdrawInitiationPeriod(client *ethclient.Client) (uint16, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	return stakeManager.WithdrawInitiationPeriod(&opts)
}

func (a AssetManagerStruct) GetNumJobs(client *ethclient.Client) (uint16, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	return collectionManager.GetNumJobs(&opts)
}

func (a AssetManagerStruct) GetCollection(client *ethclient.Client, id uint16) (bindings.StructsCollection, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	return collectionManager.GetCollection(&opts, id)
}

func (a AssetManagerStruct) GetJob(client *ethclient.Client, id uint16) (bindings.StructsJob, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	return collectionManager.GetJob(&opts, id)
}

func (a AssetManagerStruct) GetCollectionIdFromIndex(client *ethclient.Client, index uint16) (uint16, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	return collectionManager.LeafIdToCollectionIdRegistry(&opts, index)
}

func (a AssetManagerStruct) GetCollectionIdFromLeafId(client *ethclient.Client, leafId uint16) (uint16, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	return collectionManager.GetCollectionIdFromLeafId(&opts, leafId)
}

func (a AssetManagerStruct) GetLeafIdOfACollection(client *ethclient.Client, collectionId uint16) (uint16, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	return collectionManager.GetLeafIdOfCollection(&opts, collectionId)
}

func (v VoteManagerStruct) ToAssign(client *ethclient.Client) (uint16, error) {
	voteManager, opts := UtilsInterface.GetVoteManagerWithOpts(client)
	return voteManager.ToAssign(&opts)
}

func (v VoteManagerStruct) GetSaltFromBlockchain(client *ethclient.Client) ([32]byte, error) {
	voteManager, opts := UtilsInterface.GetVoteManagerWithOpts(client)
	return voteManager.GetSalt(&opts)
}

func (a AccountsStruct) GetPrivateKey(address string, password string, keystorePath string) (*ecdsa.PrivateKey, error) {
	return accounts.AccountUtilsInterface.GetPrivateKey(address, password, keystorePath)
}

func (b BlockManagerStruct) GetNumProposedBlocks(client *ethclient.Client, epoch uint32) (uint8, error) {
	blockManager, opts := UtilsInterface.GetBlockManagerWithOpts(client)
	return blockManager.GetNumProposedBlocks(&opts, epoch)
}

func (b BlockManagerStruct) GetProposedBlock(client *ethclient.Client, epoch uint32, proposedBlock uint32) (bindings.StructsBlock, error) {
	blockManager, opts := UtilsInterface.GetBlockManagerWithOpts(client)
	return blockManager.GetProposedBlock(&opts, epoch, proposedBlock)
}

func (b BlockManagerStruct) GetBlock(client *ethclient.Client, epoch uint32) (bindings.StructsBlock, error) {
	blockManager, opts := UtilsInterface.GetBlockManagerWithOpts(client)
	return blockManager.GetBlock(&opts, epoch)
}

func (b BlockManagerStruct) MinStake(client *ethclient.Client) (*big.Int, error) {
	blockManager, opts := UtilsInterface.GetBlockManagerWithOpts(client)
	return blockManager.MinStake(&opts)
}

func (b BlockManagerStruct) StateBuffer(client *ethclient.Client) (uint8, error) {
	blockManager, opts := UtilsInterface.GetBlockManagerWithOpts(client)
	return blockManager.Buffer(&opts)
}

func (b BlockManagerStruct) MaxAltBlocks(client *ethclient.Client) (uint8, error) {
	blockManager, opts := UtilsInterface.GetBlockManagerWithOpts(client)
	return blockManager.MaxAltBlocks(&opts)
}

func (b BlockManagerStruct) SortedProposedBlockIds(client *ethclient.Client, arg0 uint32, arg1 *big.Int) (uint32, error) {
	blockManager, opts := UtilsInterface.GetBlockManagerWithOpts(client)
	return blockManager.SortedProposedBlockIds(&opts, arg0, arg1)
}

func (b BlockManagerStruct) GetEpochLastProposed(client *ethclient.Client, stakerId uint32) (uint32, error) {
	blockManager, opts := UtilsInterface.GetBlockManagerWithOpts(client)
	return blockManager.EpochLastProposed(&opts, stakerId)
}

func (s StakeManagerStruct) GetStakerId(client *ethclient.Client, address common.Address) (uint32, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	return stakeManager.GetStakerId(&opts, address)
}

func (s StakeManagerStruct) GetNumStakers(client *ethclient.Client) (uint32, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	return stakeManager.GetNumStakers(&opts)
}

func (s StakeManagerStruct) MinSafeRazor(client *ethclient.Client) (*big.Int, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	return stakeManager.MinSafeRazor(&opts)
}

func (s StakeManagerStruct) Locks(client *ethclient.Client, address common.Address, address1 common.Address, lockType uint8) (coretypes.Locks, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	return stakeManager.Locks(&opts, address, address1, lockType)
}

func (s StakeManagerStruct) MaxCommission(client *ethclient.Client) (uint8, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	return stakeManager.MaxCommission(&opts)
}

func (s StakeManagerStruct) EpochLimitForUpdateCommission(client *ethclient.Client) (uint16, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	return stakeManager.EpochLimitForUpdateCommission(&opts)
}

func (s StakeManagerStruct) GetStaker(client *ethclient.Client, stakerId uint32) (bindings.StructsStaker, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	return stakeManager.GetStaker(&opts, stakerId)
}

func (a AssetManagerStruct) GetNumCollections(client *ethclient.Client) (uint16, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	return collectionManager.GetNumCollections(&opts)
}

func (a AssetManagerStruct) GetNumActiveCollections(client *ethclient.Client) (uint16, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	return collectionManager.GetNumActiveCollections(&opts)
}

func (a AssetManagerStruct) GetActiveCollections(client *ethclient.Client) ([]uint16, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	return collectionManager.GetActiveCollections(&opts)
}

func (a AssetManagerStruct) Jobs(client *ethclient.Client, id uint16) (bindings.StructsJob, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	return collectionManager.Jobs(&opts, id)
}

func (v VoteManagerStruct) Commitments(client *ethclient.Client, stakerId uint32) (coretypes.Commitment, error) {
	voteManager, opts := UtilsInterface.GetVoteManagerWithOpts(client)
	return voteManager.Commitments(&opts, stakerId)
}

func (v VoteManagerStruct) GetVoteValue(client *ethclient.Client, epoch uint32, stakerId uint32, medianIndex uint16) (*big.Int, error) {
	voteManager, opts := UtilsInterface.GetVoteManagerWithOpts(client)
	return voteManager.GetVoteValue(&opts, epoch, stakerId, medianIndex)
}

func (v VoteManagerStruct) GetInfluenceSnapshot(client *ethclient.Client, epoch uint32, stakerId uint32) (*big.Int, error) {
	voteManager, opts := UtilsInterface.GetVoteManagerWithOpts(client)
	return voteManager.GetInfluenceSnapshot(&opts, epoch, stakerId)
}

func (v VoteManagerStruct) GetStakeSnapshot(client *ethclient.Client, epoch uint32, stakerId uint32) (*big.Int, error) {
	voteManager, opts := UtilsInterface.GetVoteManagerWithOpts(client)
	return voteManager.GetStakeSnapshot(&opts, epoch, stakerId)
}

func (v VoteManagerStruct) GetTotalInfluenceRevealed(client *ethclient.Client, epoch uint32, medianIndex uint16) (*big.Int, error) {
	voteManager, opts := UtilsInterface.GetVoteManagerWithOpts(client)
	return voteManager.GetTotalInfluenceRevealed(&opts, epoch, medianIndex)
}

func (v VoteManagerStruct) GetEpochLastCommitted(client *ethclient.Client, stakerId uint32) (uint32, error) {
	voteManager, opts := UtilsInterface.GetVoteManagerWithOpts(client)
	return voteManager.GetEpochLastCommitted(&opts, stakerId)
}

func (v VoteManagerStruct) GetEpochLastRevealed(client *ethclient.Client, stakerId uint32) (uint32, error) {
	voteManager, opts := UtilsInterface.GetVoteManagerWithOpts(client)
	return voteManager.GetEpochLastRevealed(&opts, stakerId)
}

func (b BindingsStruct) NewCollectionManager(address common.Address, client *ethclient.Client) (*bindings.CollectionManager, error) {
	return bindings.NewCollectionManager(address, client)
}

func (b BindingsStruct) NewRAZOR(address common.Address, client *ethclient.Client) (*bindings.RAZOR, error) {
	return bindings.NewRAZOR(address, client)
}

func (b BindingsStruct) NewStakeManager(address common.Address, client *ethclient.Client) (*bindings.StakeManager, error) {
	return bindings.NewStakeManager(address, client)
}

func (b BindingsStruct) NewVoteManager(address common.Address, client *ethclient.Client) (*bindings.VoteManager, error) {
	return bindings.NewVoteManager(address, client)
}

func (b BindingsStruct) NewBlockManager(address common.Address, client *ethclient.Client) (*bindings.BlockManager, error) {
	return bindings.NewBlockManager(address, client)
}

func (b BindingsStruct) NewStakedToken(address common.Address, client *ethclient.Client) (*bindings.StakedToken, error) {
	return bindings.NewStakedToken(address, client)
}

func (j JsonStruct) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (j JsonStruct) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (u UtilsStruct) GetUint32(flagSet *pflag.FlagSet, name string) (uint32, error) {
	return flagSet.GetUint32(name)
}

func (e EthClientStruct) Dial(rawurl string) (*ethclient.Client, error) {
	return ethclient.Dial(rawurl)
}

func (t TimeStruct) Sleep(duration time.Duration) {
	time.Sleep(duration)
}

func (o OSStruct) OpenFile(name string, flag int, perm fs.FileMode) (*os.File, error) {
	return os.OpenFile(name, flag, perm)
}

func (o OSStruct) Open(name string) (*os.File, error) {
	return os.Open(name)
}

func (o OSStruct) WriteFile(name string, data []byte, perm fs.FileMode) error {
	return os.WriteFile(name, data, perm)
}

func (o OSStruct) ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func (i IOStruct) ReadAll(body io.ReadCloser) ([]byte, error) {
	return io.ReadAll(body)
}

func (c ClientStruct) TransactionReceipt(client *ethclient.Client, ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return client.TransactionReceipt(ctx, txHash)
}

func (c ClientStruct) BalanceAt(client *ethclient.Client, ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return client.BalanceAt(ctx, account, blockNumber)
}

func (c ClientStruct) HeaderByNumber(client *ethclient.Client, ctx context.Context, number *big.Int) (*types.Header, error) {
	return client.HeaderByNumber(ctx, number)
}

func (c ClientStruct) PendingNonceAt(client *ethclient.Client, ctx context.Context, account common.Address) (uint64, error) {
	return client.PendingNonceAt(ctx, account)
}

func (c ClientStruct) SuggestGasPrice(client *ethclient.Client, ctx context.Context) (*big.Int, error) {
	return client.SuggestGasPrice(ctx)
}

func (c ClientStruct) EstimateGas(client *ethclient.Client, ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	return client.EstimateGas(ctx, msg)
}

func (c ClientStruct) FilterLogs(client *ethclient.Client, ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	return client.FilterLogs(ctx, q)
}

func (b BufioStruct) NewScanner(r io.Reader) *bufio.Scanner {
	return bufio.NewScanner(r)
}

func (c CoinStruct) BalanceOf(coinContract *bindings.RAZOR, opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	return coinContract.BalanceOf(opts, account)
}

func (a ABIStruct) Parse(reader io.Reader) (abi.ABI, error) {
	return abi.JSON(reader)
}

func (a ABIStruct) Pack(parsedData abi.ABI, name string, args ...interface{}) ([]byte, error) {
	return parsedData.Pack(name, args...)
}

func (p PathStruct) GetDefaultPath() (string, error) {
	return path.PathUtilsInterface.GetDefaultPath()
}

func (p PathStruct) GetJobFilePath() (string, error) {
	return path.PathUtilsInterface.GetJobFilePath()
}

func (b BindStruct) NewKeyedTransactorWithChainID(key *ecdsa.PrivateKey, chainID *big.Int) (*bind.TransactOpts, error) {
	return bind.NewKeyedTransactorWithChainID(key, chainID)
}

func (s StakedTokenStruct) BalanceOf(stakedToken *bindings.StakedToken, callOpts *bind.CallOpts, address common.Address) (*big.Int, error) {
	return stakedToken.BalanceOf(callOpts, address)
}

func (r RetryStruct) RetryAttempts(numberOfAttempts uint) retry.Option {
	return retry.Attempts(numberOfAttempts)
}

func (f FlagSetStruct) GetLogFileName(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("logFile")
}
