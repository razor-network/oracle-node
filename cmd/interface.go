package cmd

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"razor/core/types"
)

type utilsInterface interface {
	GetTxnOpts(types.TransactionOptions) *bind.TransactOpts
	GetEpoch(*ethclient.Client) (uint32, error)
}

type transactionInterface interface {
	Hash(*Types.Transaction) common.Hash
}

type stakeManagerInterface interface {
	Stake(*ethclient.Client, *bind.TransactOpts, uint32, *big.Int) (*Types.Transaction, error)
}
