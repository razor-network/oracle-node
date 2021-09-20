package cmd

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"math/big"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"
)

type Utils struct{}
type TokenManagerUtils struct{}
type TransactionUtils struct{}
type StakeManagerUtils struct{}

func (u Utils) ConnectToClient(provider string) *ethclient.Client {
	return utils.ConnectToClient(provider)
}

func (u Utils) GetTokenManager(client *ethclient.Client) *bindings.RAZOR {
	return utils.GetTokenManager(client)
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

func (tokenManagerUtils TokenManagerUtils) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address, client *ethclient.Client) (*big.Int, error) {
	tokenManager := utils.GetTokenManager(client)
	return tokenManager.Allowance(opts, owner, spender)
}

func (tokenManagerUtils TokenManagerUtils) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int, client *ethclient.Client) (*Types.Transaction, error) {
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

func (stakeManagerUtils StakeManagerUtils) Stake(txnOpts *bind.TransactOpts, epoch uint32, amount *big.Int, client *ethclient.Client) (*Types.Transaction, error) {
	stakeManager := utils.GetStakeManager(client)
	return stakeManager.Stake(txnOpts, epoch, amount)
}
