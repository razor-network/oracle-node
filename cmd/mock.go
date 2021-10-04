package cmd

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"math/big"
	"razor/core/types"
	"razor/pkg/bindings"
)

type UtilsMock struct{}

type TokenManagerMock struct{}

type TransactionMock struct{}

type AssetManagerMock struct{}

type StakeManagerMock struct{}

type AccountMock struct{}

type KeystoreMock struct{}

type FlagSetMock struct{}

type ProposeUtilsMock struct{}

type BlockManagerMock struct{}

var GetOptionsMock func(bool, string, string) bind.CallOpts

var GetTxnOptsMock func(types.TransactionOptions) *bind.TransactOpts

var WaitForBlockCompletionMock func(*ethclient.Client, string) int

var WaitForCommitStateMock func(*ethclient.Client, string, string) (uint32, error)

var GetDefaultPathMock func() (string, error)

var AssignPasswordMock func(*pflag.FlagSet) string

var FetchBalanceMock func(*ethclient.Client, string) (*big.Int, error)

var AssignAmountInWeiMock func(flagSet *pflag.FlagSet) *big.Int

var CheckAmountAndBalanceMock func(amountInWei *big.Int, balance *big.Int) *big.Int

var GetAmountInDecimalMock func(amountInWei *big.Int) *big.Float

var ConnectToClientMock func(string) *ethclient.Client

var GetDelayedStateMock func(*ethclient.Client, int32) (int64, error)

var GetEpochMock func(*ethclient.Client, string) (uint32, error)

var GetStakerMock func(*ethclient.Client, string, uint32) (bindings.StructsStaker, error)

var GetNumberOfStakersMock func(*ethclient.Client, string) (uint32, error)

var GetRandaoHashMock func(*ethclient.Client, string) ([32]byte, error)

var GetNumberOfProposedBlocksMock func(*ethclient.Client, string, uint32) (uint8, error)

var GetMaxAltBlocksMock func(*ethclient.Client, string) (uint8, error)

var GetProposedBlockMock func(*ethclient.Client, string, uint32, uint8) (types.Block, error)

var AllowanceMock func(*ethclient.Client, *bind.CallOpts, common.Address, common.Address) (*big.Int, error)

var ApproveMock func(*ethclient.Client, *bind.TransactOpts, common.Address, *big.Int) (*Types.Transaction, error)

var TransferMock func(*ethclient.Client, *bind.TransactOpts, common.Address, *big.Int) (*Types.Transaction, error)

var HashMock func(*Types.Transaction) common.Hash

var StakeMock func(*ethclient.Client, *bind.TransactOpts, uint32, *big.Int) (*Types.Transaction, error)

var ResetLockMock func(*ethclient.Client, *bind.TransactOpts, uint32) (*Types.Transaction, error)

var DelegateMock func(*ethclient.Client, *bind.TransactOpts, uint32, uint32, *big.Int) (*Types.Transaction, error)

var CreateAccountMock func(string, string) accounts.Account

var AccountsMock func(string) []accounts.Account

var GetStringFromMock func(*pflag.FlagSet) (string, error)

var GetStringToMock func(*pflag.FlagSet) (string, error)

var CreateJobMock func(*bind.TransactOpts, int8, string, string, string) (*Types.Transaction, error)

var GetStringAddressMock func(*pflag.FlagSet) (string, error)

var GetUint32StakerIdMock func(*pflag.FlagSet) (uint32, error)

var GetStringNameMock func(*pflag.FlagSet) (string, error)

var GetStringUrlMock func(*pflag.FlagSet) (string, error)

var GetStringSelectorMock func(*pflag.FlagSet) (string, error)

var GetInt8PowerMock func(*pflag.FlagSet) (int8, error)

var getBiggestInfluenceAndIdMock func(*ethclient.Client, string) (*big.Int, uint32, error)

var getIterationMock func(*ethclient.Client, string, types.ElectedProposer) int

var isElectedProposerMock func(*ethclient.Client, string, types.ElectedProposer) bool

var pseudoRandomNumberGeneratorMock func([]byte, uint32, []byte) *big.Int

var MakeBlockMock func(*ethclient.Client, string, bool) ([]uint32, error)

var getSortedVotesMock func(*ethclient.Client, string, uint8, uint32) ([]*big.Int, error)

var influencedMedianMock func([]*big.Int, *big.Int) *big.Int

var ProposeMock func(*ethclient.Client, *bind.TransactOpts, uint32, []uint32, *big.Int, uint32) (*Types.Transaction, error)

func (u UtilsMock) GetOptions(pending bool, from string, blockNumber string) bind.CallOpts {
	return GetOptionsMock(pending, from, blockNumber)
}

func (u UtilsMock) GetTxnOpts(transactionData types.TransactionOptions) *bind.TransactOpts {
	return GetTxnOptsMock(transactionData)
}

func (u UtilsMock) WaitForBlockCompletion(client *ethclient.Client, hashToRead string) int {
	return WaitForBlockCompletionMock(client, hashToRead)
}

func (u UtilsMock) WaitForCommitState(client *ethclient.Client, accountAddress string, action string) (uint32, error) {
	return WaitForCommitStateMock(client, accountAddress, action)
}

func (u UtilsMock) AssignPassword(flagSet *pflag.FlagSet) string {
	return AssignPasswordMock(flagSet)
}

func (u UtilsMock) FetchBalance(client *ethclient.Client, accountAddress string) (*big.Int, error) {
	return FetchBalanceMock(client, accountAddress)
}

func (u UtilsMock) AssignAmountInWei(flagSet *pflag.FlagSet) *big.Int {
	return AssignAmountInWeiMock(flagSet)
}

func (u UtilsMock) CheckAmountAndBalance(amountInWei *big.Int, balance *big.Int) *big.Int {
	return CheckAmountAndBalanceMock(amountInWei, balance)
}

func (u UtilsMock) GetDefaultPath() (string, error) {
	return GetDefaultPathMock()
}

func (u UtilsMock) GetAmountInDecimal(amountInWei *big.Int) *big.Float {
	return GetAmountInDecimalMock(amountInWei)
}

func (u UtilsMock) ConnectToClient(provider string) *ethclient.Client {
	return ConnectToClientMock(provider)
}

func (u UtilsMock) GetDelayedState(client *ethclient.Client, buffer int32) (int64, error) {
	return GetDelayedStateMock(client, buffer)
}

func (u UtilsMock) GetEpoch(client *ethclient.Client, address string) (uint32, error) {
	return GetEpochMock(client, address)
}

func (u UtilsMock) GetStaker(client *ethclient.Client, address string, stakerId uint32) (bindings.StructsStaker, error) {
	return GetStakerMock(client, address, stakerId)
}

func (u UtilsMock) GetNumberOfStakers(client *ethclient.Client, address string) (uint32, error) {
	return GetNumberOfStakersMock(client, address)
}

func (u UtilsMock) GetRandaoHash(client *ethclient.Client, address string) ([32]byte, error) {
	return GetRandaoHashMock(client, address)
}

func (u UtilsMock) GetNumberOfProposedBlocks(client *ethclient.Client, address string, epoch uint32) (uint8, error) {
	return GetNumberOfProposedBlocksMock(client, address, epoch)
}

func (u UtilsMock) GetMaxAltBlocks(client *ethclient.Client, address string) (uint8, error) {
	return GetMaxAltBlocksMock(client, address)
}

func (u UtilsMock) GetProposedBlock(client *ethclient.Client, address string, epoch uint32, proposedBlockId uint8) (types.Block, error) {
	return GetProposedBlockMock(client, address, epoch, proposedBlockId)
}

func (tokenManagerMock TokenManagerMock) Allowance(client *ethclient.Client, opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	return AllowanceMock(client, opts, owner, spender)
}

func (tokenManagerMock TokenManagerMock) Approve(client *ethclient.Client, opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*Types.Transaction, error) {
	return ApproveMock(client, opts, spender, amount)
}

func (tokenManagerMock TokenManagerMock) Transfer(client *ethclient.Client, opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*Types.Transaction, error) {
	return TransferMock(client, opts, recipient, amount)
}

func (transactionMock TransactionMock) Hash(txn *Types.Transaction) common.Hash {
	return HashMock(txn)
}

func (assetManagerMock AssetManagerMock) CreateJob(client *ethclient.Client, opts *bind.TransactOpts, power int8, name string, selector string, url string) (*Types.Transaction, error) {
	return CreateJobMock(opts, power, name, selector, url)
}

func (stakeManagerMock StakeManagerMock) Stake(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, amount *big.Int) (*Types.Transaction, error) {
	return StakeMock(client, opts, epoch, amount)
}

func (stakeManagerMock StakeManagerMock) ResetLock(client *ethclient.Client, opts *bind.TransactOpts, stakerId uint32) (*Types.Transaction, error) {
	return ResetLockMock(client, opts, stakerId)
}

func (stakeManagerMock StakeManagerMock) Delegate(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, stakerId uint32, amount *big.Int) (*Types.Transaction, error) {
	return DelegateMock(client, opts, epoch, stakerId, amount)
}

func (account AccountMock) CreateAccount(path string, password string) accounts.Account {
	return CreateAccountMock(path, password)
}

func (ks KeystoreMock) Accounts(path string) []accounts.Account {
	return AccountsMock(path)
}

func (flagSetMock FlagSetMock) GetStringFrom(flagSet *pflag.FlagSet) (string, error) {
	return GetStringFromMock(flagSet)
}

func (flagSetMock FlagSetMock) GetStringTo(flagSet *pflag.FlagSet) (string, error) {
	return GetStringToMock(flagSet)
}

func (flagSetMock FlagSetMock) GetStringName(flagSet *pflag.FlagSet) (string, error) {
	return GetStringNameMock(flagSet)
}

func (flagSetMock FlagSetMock) GetUint32StakerId(flagset *pflag.FlagSet) (uint32, error) {
	return GetUint32StakerIdMock(flagset)
}

func (flagSetMock FlagSetMock) GetStringAddress(flagSet *pflag.FlagSet) (string, error) {
	return GetStringAddressMock(flagSet)
}

func (flagSetMock FlagSetMock) GetStringUrl(flagSet *pflag.FlagSet) (string, error) {
	return GetStringUrlMock(flagSet)
}

func (flagSetMock FlagSetMock) GetStringSelector(flagSet *pflag.FlagSet) (string, error) {
	return GetStringSelectorMock(flagSet)
}

func (flagSetMock FlagSetMock) GetInt8Power(flagSet *pflag.FlagSet) (int8, error) {
	return GetInt8PowerMock(flagSet)
}

func (proposeUtilsMock ProposeUtilsMock) getBiggestInfluenceAndId(client *ethclient.Client, address string) (*big.Int, uint32, error) {
	return getBiggestInfluenceAndIdMock(client, address)
}

func (proposeUtilsMock ProposeUtilsMock) getIteration(client *ethclient.Client, address string, proposer types.ElectedProposer) int {
	return getIterationMock(client, address, proposer)
}

func (proposeUtilsMock ProposeUtilsMock) isElectedProposer(client *ethclient.Client, address string, proposer types.ElectedProposer) bool {
	return isElectedProposerMock(client, address, proposer)
}

func (proposeUtilsMock ProposeUtilsMock) pseudoRandomNumberGenerator(seed []byte, max uint32, blockHashes []byte) *big.Int {
	return pseudoRandomNumberGeneratorMock(seed, max, blockHashes)
}

func (proposeUtilsMock ProposeUtilsMock) MakeBlock(client *ethclient.Client, address string, rogueMode bool) ([]uint32, error) {
	return MakeBlockMock(client, address, rogueMode)
}

func (proposeUtilsMock ProposeUtilsMock) getSortedVotes(client *ethclient.Client, address string, assetId uint8, epoch uint32) ([]*big.Int, error) {
	return getSortedVotesMock(client, address, assetId, epoch)
}

func (proposeUtilsMock ProposeUtilsMock) influencedMedian(sortedVotes []*big.Int, totalInfluenceRevealed *big.Int) *big.Int {
	return influencedMedianMock(sortedVotes, totalInfluenceRevealed)
}

func (blockManagerMock BlockManagerMock) Propose(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, medians []uint32, iteration *big.Int, biggestInfluencerId uint32) (*Types.Transaction, error) {
	return ProposeMock(client, opts, epoch, medians, iteration, biggestInfluencerId)
}
