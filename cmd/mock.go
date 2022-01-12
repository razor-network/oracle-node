package cmd

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"razor/core/types"
)

type UtilsMock struct{}

type TransactionMock struct{}

type StakeManagerMock struct{}

var GetTxnOptsMock func(types.TransactionOptions) *bind.TransactOpts

var GetEpochMock func(*ethclient.Client) (uint32, error)

var HashMock func(*Types.Transaction) common.Hash

var StakeMock func(*ethclient.Client, *bind.TransactOpts, uint32, *big.Int) (*Types.Transaction, error)

func (u UtilsMock) GetTxnOpts(transactionData types.TransactionOptions) *bind.TransactOpts {
	return GetTxnOptsMock(transactionData)
}

func (u UtilsMock) GetEpoch(client *ethclient.Client) (uint32, error) {
	return GetEpochMock(client)
}

func (transactionMock TransactionMock) Hash(txn *Types.Transaction) common.Hash {
	return HashMock(txn)
}

func (stakeManagerMock StakeManagerMock) Stake(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, amount *big.Int) (*Types.Transaction, error) {
	return StakeMock(client, opts, epoch, amount)
}
