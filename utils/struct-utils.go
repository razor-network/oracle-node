package utils

import (
	"bufio"
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
	"github.com/spf13/pflag"
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
)

func StartRazor(optionsPackageStruct OptionsPackageStruct) Utils {
	Options = optionsPackageStruct.Options
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
	VoteManagerInterface = optionsPackageStruct.VoteManagerInterface
	BindingsInterface = optionsPackageStruct.BindingsInterface
	JsonInterface = optionsPackageStruct.JsonInterface
	return &UtilsStruct{}
}

func (a AccountsStruct) GetPrivateKey(address string, password string, keystorePath string, accountUtils accounts.AccountInterface) *ecdsa.PrivateKey {
	return accounts.AccountUtilsInterface.GetPrivateKey(address, password, keystorePath)
}

func (o OptionsStruct) RetryAttempts(numberOfAttempts uint) retry.Option {
	return retry.Attempts(numberOfAttempts)
}

func (o OptionsStruct) ConvertToNumber(num interface{}) (*big.Float, error) {
	return ConvertToNumber(num)
}

func (b BlockManagerStruct) GetNumProposedBlocks(client *ethclient.Client, opts *bind.CallOpts, epoch uint32) (uint8, error) {
	blockManager := UtilsInterface.GetBlockManager(client)
	return blockManager.GetNumProposedBlocks(opts, epoch)
}

func (b BlockManagerStruct) GetProposedBlock(client *ethclient.Client, opts *bind.CallOpts, epoch uint32, proposedBlock uint32) (bindings.StructsBlock, error) {
	blockManager := UtilsInterface.GetBlockManager(client)
	return blockManager.GetProposedBlock(opts, epoch, proposedBlock)
}

func (b BlockManagerStruct) GetBlock(client *ethclient.Client, opts *bind.CallOpts, epoch uint32) (bindings.StructsBlock, error) {
	blockManager := UtilsInterface.GetBlockManager(client)
	return blockManager.GetBlock(opts, epoch)
}

func (b BlockManagerStruct) MinStake(client *ethclient.Client, opts *bind.CallOpts) (*big.Int, error) {
	blockManager := UtilsInterface.GetBlockManager(client)
	return blockManager.MinStake(opts)
}

func (b BlockManagerStruct) MaxAltBlocks(client *ethclient.Client, opts *bind.CallOpts) (uint8, error) {
	blockManager := UtilsInterface.GetBlockManager(client)
	return blockManager.MaxAltBlocks(opts)
}

func (b BlockManagerStruct) SortedProposedBlockIds(client *ethclient.Client, opts *bind.CallOpts, arg0 uint32, arg1 *big.Int) (uint32, error) {
	blockManager := UtilsInterface.GetBlockManager(client)
	return blockManager.SortedProposedBlockIds(opts, arg0, arg1)
}

func (s StakeManagerStruct) GetStakerId(client *ethclient.Client, opts *bind.CallOpts, address common.Address) (uint32, error) {
	stakeManager := UtilsInterface.GetStakeManager(client)
	return stakeManager.GetStakerId(opts, address)
}

func (s StakeManagerStruct) GetNumStakers(client *ethclient.Client, opts *bind.CallOpts) (uint32, error) {
	stakeManager := UtilsInterface.GetStakeManager(client)
	return stakeManager.GetNumStakers(opts)
}

func (s StakeManagerStruct) Locks(client *ethclient.Client, opts *bind.CallOpts, address common.Address, address1 common.Address) (coretypes.Locks, error) {
	stakeManager := UtilsInterface.GetStakeManager(client)
	return stakeManager.Locks(opts, address, address1)
}

func (s StakeManagerStruct) WithdrawReleasePeriod(client *ethclient.Client, opts *bind.CallOpts) (uint8, error) {
	stakeManager := UtilsInterface.GetStakeManager(client)
	return stakeManager.WithdrawReleasePeriod(opts)
}

func (s StakeManagerStruct) MaxCommission(client *ethclient.Client, opts *bind.CallOpts) (uint8, error) {
	stakeManager := UtilsInterface.GetStakeManager(client)
	return stakeManager.MaxCommission(opts)
}

func (s StakeManagerStruct) EpochLimitForUpdateCommission(client *ethclient.Client, opts *bind.CallOpts) (uint16, error) {
	stakeManager := UtilsInterface.GetStakeManager(client)
	return stakeManager.EpochLimitForUpdateCommission(opts)
}

func (s StakeManagerStruct) GetStaker(client *ethclient.Client, opts *bind.CallOpts, stakerId uint32) (bindings.StructsStaker, error) {
	stakeManager := UtilsInterface.GetStakeManager(client)
	return stakeManager.GetStaker(opts, stakerId)
}

func (a AssetManagerStruct) GetNumAssets(client *ethclient.Client, opts *bind.CallOpts) (uint16, error) {
	assetManager := UtilsInterface.GetAssetManager(client)
	return assetManager.GetNumAssets(opts)
}

func (a AssetManagerStruct) GetNumActiveCollections(client *ethclient.Client, opts *bind.CallOpts) (*big.Int, error) {
	assetManager := UtilsInterface.GetAssetManager(client)
	return assetManager.GetNumActiveCollections(opts)
}

func (a AssetManagerStruct) GetAsset(client *ethclient.Client, opts *bind.CallOpts, id uint16) (coretypes.Asset, error) {
	assetManager := UtilsInterface.GetAssetManager(client)
	return assetManager.GetAsset(opts, id)
}

func (a AssetManagerStruct) GetActiveCollections(client *ethclient.Client, opts *bind.CallOpts) ([]uint16, error) {
	assetManager := UtilsInterface.GetAssetManager(client)
	return assetManager.GetActiveCollections(opts)
}

func (a AssetManagerStruct) Jobs(client *ethclient.Client, opts *bind.CallOpts, id uint16) (bindings.StructsJob, error) {
	assetManager := UtilsInterface.GetAssetManager(client)
	return assetManager.Jobs(opts, id)
}

func (v VoteManagerStruct) Commitments(client *ethclient.Client, opts *bind.CallOpts, stakerId uint32) (coretypes.Commitment, error) {
	voteManager := UtilsInterface.GetVoteManager(client)
	return voteManager.Commitments(opts, stakerId)
}

func (v VoteManagerStruct) GetVoteValue(client *ethclient.Client, opts *bind.CallOpts, assetIndex uint16, stakerId uint32) (*big.Int, error) {
	voteManager := UtilsInterface.GetVoteManager(client)
	return voteManager.GetVoteValue(opts, assetIndex, stakerId)
}

func (v VoteManagerStruct) GetVote(client *ethclient.Client, opts *bind.CallOpts, stakerId uint32) (bindings.StructsVote, error) {
	voteManager := UtilsInterface.GetVoteManager(client)
	return voteManager.GetVote(opts, stakerId)
}

func (v VoteManagerStruct) GetInfluenceSnapshot(client *ethclient.Client, opts *bind.CallOpts, epoch uint32, stakerId uint32) (*big.Int, error) {
	voteManager := UtilsInterface.GetVoteManager(client)
	return voteManager.GetInfluenceSnapshot(opts, epoch, stakerId)
}

func (v VoteManagerStruct) GetStakeSnapshot(client *ethclient.Client, opts *bind.CallOpts, epoch uint32, stakerId uint32) (*big.Int, error) {
	voteManager := UtilsInterface.GetVoteManager(client)
	return voteManager.GetStakeSnapshot(opts, epoch, stakerId)
}

func (v VoteManagerStruct) GetTotalInfluenceRevealed(client *ethclient.Client, opts *bind.CallOpts, epoch uint32) (*big.Int, error) {
	voteManager := UtilsInterface.GetVoteManager(client)
	return voteManager.GetTotalInfluenceRevealed(opts, epoch)
}

func (v VoteManagerStruct) GetRandaoHash(client *ethclient.Client, opts *bind.CallOpts) ([32]byte, error) {
	voteManager := UtilsInterface.GetVoteManager(client)
	return voteManager.GetRandaoHash(opts)
}

func (v VoteManagerStruct) GetEpochLastCommitted(client *ethclient.Client, opts *bind.CallOpts, stakerId uint32) (uint32, error) {
	voteManager := UtilsInterface.GetVoteManager(client)
	return voteManager.GetEpochLastCommitted(opts, stakerId)
}

func (v VoteManagerStruct) GetEpochLastRevealed(client *ethclient.Client, opts *bind.CallOpts, stakerId uint32) (uint32, error) {
	voteManager := UtilsInterface.GetVoteManager(client)
	return voteManager.GetEpochLastRevealed(opts, stakerId)
}

func (b BindingsStruct) NewAssetManager(address common.Address, client *ethclient.Client) (*bindings.AssetManager, error) {
	return bindings.NewAssetManager(address, client)
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

func (i IoutilStruct) ReadAll(body io.ReadCloser) ([]byte, error) {
	return ioutil.ReadAll(body)
}

func (i IoutilStruct) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func (i IoutilStruct) WriteFile(filename string, data []byte, perm fs.FileMode) error {
	return ioutil.WriteFile(filename, data, perm)
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
