package cmd

import (
	ethAccounts "github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"github.com/wealdtech/go-merkletree"
	"math/big"
	"razor/accounts"
	"razor/core/types"
	"razor/path"
	"razor/utils"
)

type Utils struct{}
type TokenManagerUtils struct{}
type TransactionUtils struct{}
type StakeManagerUtils struct{}
type AssetManagerUtils struct{}
type AccountUtils struct{}
type KeystoreUtils struct{}
type FlagSetUtils struct{}
type VoteManagerUtils struct{}
type TreeUtils struct{}

func (u Utils) ConnectToClient(provider string) *ethclient.Client {
	return utils.ConnectToClient(provider)
}

func (u Utils) GetOptions(pending bool, from string, blockNumber string) bind.CallOpts {
	return utils.GetOptions(pending, from, blockNumber)
}

func (u Utils) GetTxnOpts(transactionData types.TransactionOptions) *bind.TransactOpts {
	return utils.GetTxnOpts(transactionData)
}

func (u Utils) WaitForBlockCompletion(client *ethclient.Client, hashToRead string) int {
	return utils.WaitForBlockCompletion(client, hashToRead)
}

func (u Utils) WaitForCommitState(client *ethclient.Client, accountAddress string, action string) (uint32, error) {
	return WaitForCommitState(client, accountAddress, action)
}

func (u Utils) AssignPassword(flagSet *pflag.FlagSet) string {
	return utils.AssignPassword(flagSet)
}

func (u Utils) FetchBalance(client *ethclient.Client, accountAddress string) (*big.Int, error) {
	return utils.FetchBalance(client, accountAddress)
}

func (u Utils) AssignAmountInWei(flagSet *pflag.FlagSet) *big.Int {
	return utils.AssignAmountInWei(flagSet)
}

func (u Utils) CheckAmountAndBalance(amountInWei *big.Int, balance *big.Int) *big.Int {
	return utils.CheckAmountAndBalance(amountInWei, balance)
}

func (u Utils) GetAmountInDecimal(amountInWei *big.Int) *big.Float {
	return utils.GetAmountInDecimal(amountInWei)
}

func (u Utils) GetDefaultPath() (string, error) {
	return path.GetDefaultPath()
}

func (u Utils) GetDelayedState(client *ethclient.Client, buffer int32) (int64, error) {
	return utils.GetDelayedState(client, buffer)
}

func (u Utils) GetEpoch(client *ethclient.Client, address string) (uint32, error) {
	return utils.GetEpoch(client, address)
}

func (u Utils) GetCommitments(client *ethclient.Client, address string) ([32]byte, error) {
	return utils.GetCommitments(client, address)
}

func (u Utils) AllZero(bytesValue [32]byte) bool {
	return utils.AllZero(bytesValue)
}

func (u Utils) GetMerkleTree(data []*big.Int) (*merkletree.MerkleTree, error) {
	return utils.GetMerkleTree(data)
}

func (u Utils) GetEpochLastCommitted(client *ethclient.Client, address string, stakerId uint32) (uint32, error) {
	return utils.GetEpochLastCommitted(client, address, stakerId)
}

func (tokenManagerUtils TokenManagerUtils) Allowance(client *ethclient.Client, opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	tokenManager := utils.GetTokenManager(client)
	return tokenManager.Allowance(opts, owner, spender)
}

func (tokenManagerUtils TokenManagerUtils) Approve(client *ethclient.Client, opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*Types.Transaction, error) {
	tokenManager := utils.GetTokenManager(client)
	return tokenManager.Approve(opts, spender, amount)
}

func (tokenManagerUtils TokenManagerUtils) Transfer(client *ethclient.Client, opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*Types.Transaction, error) {
	tokenManager := utils.GetTokenManager(client)
	return tokenManager.Transfer(opts, recipient, amount)
}

func (transactionUtils TransactionUtils) Hash(txn *Types.Transaction) common.Hash {
	return txn.Hash()
}

func (stakeManagerUtils StakeManagerUtils) Stake(client *ethclient.Client, txnOpts *bind.TransactOpts, epoch uint32, amount *big.Int) (*Types.Transaction, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.Stake(txnOpts, epoch, amount)
}

func (stakeManagerUtils StakeManagerUtils) ResetLock(client *ethclient.Client, opts *bind.TransactOpts, stakerId uint32) (*Types.Transaction, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.ResetLock(opts, stakerId)
}

func (stakeManagerUtils StakeManagerUtils) Delegate(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, stakerId uint32, amount *big.Int) (*Types.Transaction, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.Delegate(opts, epoch, stakerId, amount)
}

func (assetManagerUtils AssetManagerUtils) CreateJob(client *ethclient.Client, opts *bind.TransactOpts, power int8, name string, selector string, url string) (*Types.Transaction, error) {
	assetManager := utils.GetAssetManager(client)
	return assetManager.CreateJob(opts, power, name, selector, url)
}

func (account AccountUtils) CreateAccount(path string, password string) ethAccounts.Account {
	return accounts.CreateAccount(path, password)
}

func (keystoreUtils KeystoreUtils) Accounts(path string) []ethAccounts.Account {
	ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	return ks.Accounts()
}

func (flagSetUtils FlagSetUtils) GetStringFrom(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("from")
}

func (flagSetUtils FlagSetUtils) GetStringTo(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("to")
}

func (flagSetUtils FlagSetUtils) GetStringAddress(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("address")
}

func (flagSetUtils FlagSetUtils) GetUint32StakerId(flagSet *pflag.FlagSet) (uint32, error) {
	return flagSet.GetUint32("stakerId")
}

func (flagSetUtils FlagSetUtils) GetStringName(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("name")
}

func (flagSetUtils FlagSetUtils) GetStringUrl(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("url")
}

func (flagSetUtils FlagSetUtils) GetStringSelector(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("selector")
}

func (flagSetUtils FlagSetUtils) GetInt8Power(flagSet *pflag.FlagSet) (int8, error) {
	return flagSet.GetInt8("power")
}

func (voteManagerUtils VoteManagerUtils) Reveal(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, values []*big.Int, secret [32]byte) (*Types.Transaction, error) {
	voteManager := utils.GetVoteManager(client)
	return voteManager.Reveal(opts, epoch, values, secret)
}

func (t TreeUtils) RootV1(tree *merkletree.MerkleTree) []byte {
	return tree.RootV1()
}
