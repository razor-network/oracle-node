package cmd

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"math/big"
	"razor/core/types"
)

type UtilsMock struct{}

type TokenManagerMock struct{}

type TransactionMock struct{}

type AssetManagerMock struct{}

type StakeManagerMock struct{}

type AccountMock struct{}

type KeystoreMock struct{}

type FlagSetMock struct{}

type VoteManagerMock struct{}

type BlockManagerMock struct{}

type CryptoMock struct{}

var GetOptionsMock func(bool, string, string) bind.CallOpts

var GetTxnOptsMock func(types.TransactionOptions) *bind.TransactOpts

var WaitForBlockCompletionMock func(*ethclient.Client, string) int

var WaitForCommitStateMock func(*ethclient.Client, string, string) (uint32, error)

var GetDefaultPathMock func() (string, error)

var AssignPasswordMock func(*pflag.FlagSet) string

var FetchBalanceMock func(*ethclient.Client, string) (*big.Int, error)

var AssignAmountInWeiMock func(flagSet *pflag.FlagSet) *big.Int

var CheckAmountAndBalanceMock func(amountInWei *big.Int, balance *big.Int) *big.Int

var GetAmountInDecimalMock func(*big.Int) *big.Float

var ConnectToClientMock func(string) *ethclient.Client

var GetDelayedStateMock func(*ethclient.Client, int32) (int64, error)

var GetEpochMock func(*ethclient.Client, string) (uint32, error)

var GetCommitmentsMock func(*ethclient.Client, string) ([32]byte, error)

var AllZeroMock func([32]byte) bool

var GetEpochLastCommittedMock func(*ethclient.Client, string, uint32) (uint32, error)

var GetActiveAssetsDataMock func(*ethclient.Client, string, uint32) ([]*big.Int, error)

var ConvertUintArrayToUint8ArrayMock func([]uint) []uint8

var WaitForDisputeOrConfirmStateMock func(*ethclient.Client, string, string) (uint32, error)

var PrivateKeyPromptMock func() string

var PasswordPromptMock func() string

var AllowanceMock func(*ethclient.Client, *bind.CallOpts, common.Address, common.Address) (*big.Int, error)

var ApproveMock func(*ethclient.Client, *bind.TransactOpts, common.Address, *big.Int) (*Types.Transaction, error)

var TransferMock func(*ethclient.Client, *bind.TransactOpts, common.Address, *big.Int) (*Types.Transaction, error)

var HashMock func(*Types.Transaction) common.Hash

var StakeMock func(*ethclient.Client, *bind.TransactOpts, uint32, *big.Int) (*Types.Transaction, error)

var ResetLockMock func(*ethclient.Client, *bind.TransactOpts, uint32) (*Types.Transaction, error)

var DelegateMock func(*ethclient.Client, *bind.TransactOpts, uint32, uint32, *big.Int) (*Types.Transaction, error)

var CreateAccountMock func(string, string) accounts.Account

var CreateJobMock func(*ethclient.Client, *bind.TransactOpts, int8, string, string, string) (*Types.Transaction, error)

var UpdateJobMock func(*ethclient.Client, *bind.TransactOpts, uint8, int8, string, string) (*Types.Transaction, error)

var UpdateCollectionMock func(*ethclient.Client, *bind.TransactOpts, uint8, uint32, int8) (*Types.Transaction, error)

var CreateCollectionMock func(*ethclient.Client, *bind.TransactOpts, []uint8, uint32, int8, string) (*Types.Transaction, error)

var AccountsMock func(string) []accounts.Account

var ImportECDSAMock func(string, *ecdsa.PrivateKey, string) (accounts.Account, error)

var GetStringFromMock func(*pflag.FlagSet) (string, error)

var GetStringToMock func(*pflag.FlagSet) (string, error)

var AddJobToCollectionMock func(*ethclient.Client, *bind.TransactOpts, uint8, uint8) (*Types.Transaction, error)

var RemoveJobFromCollectionMock func(*ethclient.Client, *bind.TransactOpts, uint8, uint8) (*Types.Transaction, error)

var GetStringAddressMock func(*pflag.FlagSet) (string, error)

var GetUint32StakerIdMock func(*pflag.FlagSet) (uint32, error)

var GetStringNameMock func(*pflag.FlagSet) (string, error)

var GetStringUrlMock func(*pflag.FlagSet) (string, error)

var GetStringSelectorMock func(*pflag.FlagSet) (string, error)

var GetInt8PowerMock func(*pflag.FlagSet) (int8, error)

var RevealMock func(*ethclient.Client, *bind.TransactOpts, uint32, []*big.Int, [32]byte) (*Types.Transaction, error)

var CommitMock func(*ethclient.Client, *bind.TransactOpts, uint32, [32]byte) (*Types.Transaction, error)

var ClaimBlockRewardMock func(*ethclient.Client, *bind.TransactOpts) (*Types.Transaction, error)

var GetUintSliceJobIdsMock func(*pflag.FlagSet) ([]uint, error)

var GetUint32AggregationMock func(*pflag.FlagSet) (uint32, error)

var GetUint8JobIdMock func(*pflag.FlagSet) (uint8, error)

var GetUint8CollectionIdMock func(*pflag.FlagSet) (uint8, error)

var HexToECDSAMock func(string) (*ecdsa.PrivateKey, error)

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

func (u UtilsMock) GetCommitments(client *ethclient.Client, address string) ([32]byte, error) {
	return GetCommitmentsMock(client, address)
}

func (u UtilsMock) AllZero(bytesValue [32]byte) bool {
	return AllZeroMock(bytesValue)
}

func (u UtilsMock) GetEpochLastCommitted(client *ethclient.Client, address string, stakerId uint32) (uint32, error) {
	return GetEpochLastCommittedMock(client, address, stakerId)
}

func (u UtilsMock) GetActiveAssetsData(client *ethclient.Client, address string, epoch uint32) ([]*big.Int, error) {
	return GetActiveAssetsDataMock(client, address, epoch)
}

func (u UtilsMock) ConvertUintArrayToUint8Array(uintArr []uint) []uint8 {
	return ConvertUintArrayToUint8ArrayMock(uintArr)
}

func (u UtilsMock) WaitForDisputeOrConfirmState(client *ethclient.Client, accountAddress string, action string) (uint32, error) {
	return WaitForDisputeOrConfirmStateMock(client, accountAddress, action)
}

func (u UtilsMock) PrivateKeyPrompt() string {
	return PrivateKeyPromptMock()
}

func (u UtilsMock) PasswordPrompt() string {
	return PasswordPromptMock()
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
	return CreateJobMock(client, opts, power, name, selector, url)
}

func (assetManagerMock AssetManagerMock) CreateCollection(client *ethclient.Client, opts *bind.TransactOpts, jobIDs []uint8, aggregationMethod uint32, power int8, name string) (*Types.Transaction, error) {
	return CreateCollectionMock(client, opts, jobIDs, aggregationMethod, power, name)
}

func (assetManagerMock AssetManagerMock) AddJobToCollection(client *ethclient.Client, opts *bind.TransactOpts, collectionID uint8, jobID uint8) (*Types.Transaction, error) {
	return AddJobToCollectionMock(client, opts, collectionID, jobID)
}

func (assetManagerMock AssetManagerMock) RemoveJobFromCollection(client *ethclient.Client, opts *bind.TransactOpts, collectionID uint8, jobID uint8) (*Types.Transaction, error) {
	return RemoveJobFromCollectionMock(client, opts, collectionID, jobID)
}

func (assetManagerMock AssetManagerMock) UpdateJob(client *ethclient.Client, opts *bind.TransactOpts, jobId uint8, power int8, selector string, url string) (*Types.Transaction, error) {
	return UpdateJobMock(client, opts, jobId, power, selector, url)
}

func (assetManagerMock AssetManagerMock) UpdateCollection(client *ethclient.Client, opts *bind.TransactOpts, collectionId uint8, aggregationMethod uint32, power int8) (*Types.Transaction, error) {
	return UpdateCollectionMock(client, opts, collectionId, aggregationMethod, power)
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

func (ks KeystoreMock) ImportECDSA(path string, priv *ecdsa.PrivateKey, passphrase string) (accounts.Account, error) {
	return ImportECDSAMock(path, priv, passphrase)
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

func (voteManagerMock VoteManagerMock) Reveal(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, values []*big.Int, secret [32]byte) (*Types.Transaction, error) {
	return RevealMock(client, opts, epoch, values, secret)
}

func (voteManagerMock VoteManagerMock) Commit(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, commitment [32]byte) (*Types.Transaction, error) {
	return CommitMock(client, opts, epoch, commitment)
}

func (blockManagerMock BlockManagerMock) ClaimBlockReward(client *ethclient.Client, opts *bind.TransactOpts) (*Types.Transaction, error) {
	return ClaimBlockRewardMock(client, opts)
}

func (flagSetMock FlagSetMock) GetUintSliceJobIds(flagSet *pflag.FlagSet) ([]uint, error) {
	return GetUintSliceJobIdsMock(flagSet)
}

func (flagSetMock FlagSetMock) GetUint32Aggregation(flagSet *pflag.FlagSet) (uint32, error) {
	return GetUint32AggregationMock(flagSet)
}

func (flagSetMock FlagSetMock) GetUint8JobId(flagSet *pflag.FlagSet) (uint8, error) {
	return GetUint8JobIdMock(flagSet)
}

func (flagSetMock FlagSetMock) GetUint8CollectionId(flagSet *pflag.FlagSet) (uint8, error) {
	return GetUint8CollectionIdMock(flagSet)
}

func (c CryptoMock) HexToECDSA(hexKey string) (*ecdsa.PrivateKey, error) {
	return HexToECDSAMock(hexKey)
}
