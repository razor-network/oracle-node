package utils

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"io"
	"io/fs"
	"io/ioutil"
	"math/big"
	"razor/accounts"
	coretypes "razor/core/types"
	"razor/path"
	"razor/pkg/bindings"
)

func StartRazor(optionsPackageStruct OptionsPackageStruct) Utils {
	Options = optionsPackageStruct.Options
	UtilsInterface = optionsPackageStruct.UtilsInterface
	return &UtilsStruct{}
}

func (o OptionsStruct) Parse(reader io.Reader) (abi.ABI, error) {
	return abi.JSON(reader)
}

func (o OptionsStruct) Pack(parsedData abi.ABI, name string, args ...interface{}) ([]byte, error) {
	return parsedData.Pack(name, args...)
}

func (o OptionsStruct) GetDefaultPath() (string, error) {
	return path.PathUtilsInterface.GetDefaultPath()
}

func (o OptionsStruct) GetJobFilePath() (string, error) {
	return path.PathUtilsInterface.GetJobFilePath()
}

func (o OptionsStruct) GetPrivateKey(address string, password string, keystorePath string) *ecdsa.PrivateKey {
	return accounts.AccountUtilsInterface.GetPrivateKey(address, password, keystorePath)
}

func (o OptionsStruct) NewKeyedTransactorWithChainID(key *ecdsa.PrivateKey, chainID *big.Int) (*bind.TransactOpts, error) {
	return bind.NewKeyedTransactorWithChainID(key, chainID)
}

func (o OptionsStruct) RetryAttempts(numberOfAttempts uint) retry.Option {
	return retry.Attempts(numberOfAttempts)
}

func (o OptionsStruct) HeaderByNumber(client *ethclient.Client, ctx context.Context, number *big.Int) (*types.Header, error) {
	return client.HeaderByNumber(ctx, number)
}

func (o OptionsStruct) PendingNonceAt(client *ethclient.Client, ctx context.Context, account common.Address) (uint64, error) {
	return client.PendingNonceAt(ctx, account)
}

func (o OptionsStruct) SuggestGasPrice(client *ethclient.Client, ctx context.Context) (*big.Int, error) {
	return client.SuggestGasPrice(ctx)
}

func (o OptionsStruct) EstimateGas(client *ethclient.Client, ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	return client.EstimateGas(ctx, msg)
}

func (o OptionsStruct) FilterLogs(client *ethclient.Client, ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	return client.FilterLogs(ctx, q)
}

func (o OptionsStruct) BalanceAt(client *ethclient.Client, ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return client.BalanceAt(ctx, account, blockNumber)
}

func (o OptionsStruct) GetNumProposedBlocks(client *ethclient.Client, epoch uint32) (uint8, error) {
	blockManager, opts := UtilsInterface.GetBlockManagerWithOpts(client)
	return blockManager.GetNumProposedBlocks(&opts, epoch)
}

func (o OptionsStruct) GetProposedBlock(client *ethclient.Client, epoch uint32, proposedBlock uint32) (bindings.StructsBlock, error) {
	blockManager, opts := UtilsInterface.GetBlockManagerWithOpts(client)
	return blockManager.GetProposedBlock(&opts, epoch, proposedBlock)
}

func (o OptionsStruct) GetBlock(client *ethclient.Client, epoch uint32) (bindings.StructsBlock, error) {
	blockManager, opts := UtilsInterface.GetBlockManagerWithOpts(client)
	return blockManager.GetBlock(&opts, epoch)
}

func (o OptionsStruct) MinStake(client *ethclient.Client) (*big.Int, error) {
	blockManager, opts := UtilsInterface.GetBlockManagerWithOpts(client)
	return blockManager.MinStake(&opts)
}

func (o OptionsStruct) MaxAltBlocks(client *ethclient.Client) (uint8, error) {
	blockManager, opts := UtilsInterface.GetBlockManagerWithOpts(client)
	return blockManager.MaxAltBlocks(&opts)
}

func (o OptionsStruct) SortedProposedBlockIds(client *ethclient.Client, opts *bind.CallOpts, arg0 uint32, arg1 *big.Int) (uint32, error) {
	blockManager := UtilsInterface.GetBlockManager(client)
	return blockManager.SortedProposedBlockIds(opts, arg0, arg1)
}

func (o OptionsStruct) GetStakerId(client *ethclient.Client, address common.Address) (uint32, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	return stakeManager.GetStakerId(&opts, address)
}

func (o OptionsStruct) GetNumStakers(client *ethclient.Client) (uint32, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	return stakeManager.GetNumStakers(&opts)
}

func (o OptionsStruct) Locks(client *ethclient.Client, opts *bind.CallOpts, address common.Address, address1 common.Address) (coretypes.Locks, error) {
	stakeManager := UtilsInterface.GetStakeManager(client)
	return stakeManager.Locks(opts, address, address1)
}

func (o OptionsStruct) WithdrawInitiationPeriod(client *ethclient.Client) (uint8, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	return stakeManager.WithdrawInitiationPeriod(&opts)
}

func (o OptionsStruct) MaxCommission(client *ethclient.Client) (uint8, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	return stakeManager.MaxCommission(&opts)
}

func (o OptionsStruct) EpochLimitForUpdateCommission(client *ethclient.Client) (uint16, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	return stakeManager.EpochLimitForUpdateCommission(&opts)
}

func (o OptionsStruct) GetStaker(client *ethclient.Client, stakerId uint32) (bindings.StructsStaker, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	return stakeManager.GetStaker(&opts, stakerId)
}

func (o OptionsStruct) GetNumCollections(client *ethclient.Client) (uint16, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	return collectionManager.GetNumCollections(&opts)
}

func (o OptionsStruct) GetNumJobs(client *ethclient.Client) (uint16, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	return collectionManager.GetNumJobs(&opts)
}

func (o OptionsStruct) GetNumActiveCollections(client *ethclient.Client) (uint16, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	return collectionManager.GetNumActiveCollections(&opts)
}

func (o OptionsStruct) GetCollection(client *ethclient.Client, id uint16) (bindings.StructsCollection, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	return collectionManager.GetCollection(&opts, id)
}

func (o OptionsStruct) GetJob(client *ethclient.Client, id uint16) (bindings.StructsJob, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	return collectionManager.GetJob(&opts, id)
}

func (o OptionsStruct) GetActiveCollections(client *ethclient.Client) ([]uint16, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	return collectionManager.GetActiveCollections(&opts)
}

func (o OptionsStruct) Jobs(client *ethclient.Client, id uint16) (bindings.StructsJob, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	return collectionManager.Jobs(&opts, id)
}

func (o OptionsStruct) ConvertToNumber(num interface{}) (*big.Float, error) {
	return ConvertToNumber(num)
}

func (o OptionsStruct) ReadAll(body io.ReadCloser) ([]byte, error) {
	return ioutil.ReadAll(body)
}

func (o OptionsStruct) Commitments(client *ethclient.Client, opts *bind.CallOpts, stakerId uint32) (coretypes.Commitment, error) {
	voteManager := UtilsInterface.GetVoteManager(client)
	return voteManager.Commitments(opts, stakerId)
}

func (o OptionsStruct) GetVoteValue(client *ethclient.Client, opts *bind.CallOpts, epoch uint32, stakerId uint32, medianIndex uint16) (uint32, error) {
	voteManager := UtilsInterface.GetVoteManager(client)
	return voteManager.GetVoteValue(opts, epoch, stakerId, medianIndex)
}

func (o OptionsStruct) GetInfluenceSnapshot(client *ethclient.Client, opts *bind.CallOpts, epoch uint32, stakerId uint32) (*big.Int, error) {
	voteManager := UtilsInterface.GetVoteManager(client)
	return voteManager.GetInfluenceSnapshot(opts, epoch, stakerId)
}

func (o OptionsStruct) GetStakeSnapshot(client *ethclient.Client, opts *bind.CallOpts, epoch uint32, stakerId uint32) (*big.Int, error) {
	voteManager := UtilsInterface.GetVoteManager(client)
	return voteManager.GetStakeSnapshot(opts, epoch, stakerId)
}

func (o OptionsStruct) GetTotalInfluenceRevealed(client *ethclient.Client, opts *bind.CallOpts, epoch uint32, medianIndex uint16) (*big.Int, error) {
	voteManager := UtilsInterface.GetVoteManager(client)
	return voteManager.GetTotalInfluenceRevealed(opts, epoch, medianIndex)
}

func (o OptionsStruct) GetEpochLastCommitted(client *ethclient.Client, opts *bind.CallOpts, stakerId uint32) (uint32, error) {
	voteManager := UtilsInterface.GetVoteManager(client)
	return voteManager.GetEpochLastCommitted(opts, stakerId)
}

func (o OptionsStruct) GetEpochLastRevealed(client *ethclient.Client, opts *bind.CallOpts, stakerId uint32) (uint32, error) {
	voteManager := UtilsInterface.GetVoteManager(client)
	return voteManager.GetEpochLastRevealed(opts, stakerId)
}

func (o OptionsStruct) NewCollectionManager(address common.Address, client *ethclient.Client) (*bindings.CollectionManager, error) {
	return bindings.NewCollectionManager(address, client)
}

func (o OptionsStruct) NewRAZOR(address common.Address, client *ethclient.Client) (*bindings.RAZOR, error) {
	return bindings.NewRAZOR(address, client)
}

func (o OptionsStruct) NewStakeManager(address common.Address, client *ethclient.Client) (*bindings.StakeManager, error) {
	return bindings.NewStakeManager(address, client)
}

func (o OptionsStruct) NewVoteManager(address common.Address, client *ethclient.Client) (*bindings.VoteManager, error) {
	return bindings.NewVoteManager(address, client)
}

func (o OptionsStruct) NewBlockManager(address common.Address, client *ethclient.Client) (*bindings.BlockManager, error) {
	return bindings.NewBlockManager(address, client)
}

func (o OptionsStruct) NewStakedToken(address common.Address, client *ethclient.Client) (*bindings.StakedToken, error) {
	return bindings.NewStakedToken(address, client)
}

func (o OptionsStruct) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func (o OptionsStruct) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (o OptionsStruct) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (o OptionsStruct) WriteFile(filename string, data []byte, perm fs.FileMode) error {
	return ioutil.WriteFile(filename, data, perm)
}
