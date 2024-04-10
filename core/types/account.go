package types

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts"
)

//go:generate mockery --name=AccountManagerInterface --output=../../accounts/mocks --case=underscore

type Account struct {
	Address        string
	Password       string
	AccountManager AccountManagerInterface
}

type AccountManagerInterface interface {
	CreateAccount(keystorePath, password string) accounts.Account
	GetPrivateKey(address, password string) (*ecdsa.PrivateKey, error)
	SignData(hash []byte, address string, password string) ([]byte, error)
	NewAccount(passphrase string) (accounts.Account, error)
}
