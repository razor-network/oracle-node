//Package account provides all account related functions
package accounts

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"razor/core/types"
)

//go:generate mockery --name AccountInterface --output ./mocks/ --case=underscore

var AccountUtilsInterface AccountInterface

type AccountInterface interface {
	CreateAccount(path string, password string) accounts.Account
	GetPrivateKey(address string, password string, keystorePath string) (*ecdsa.PrivateKey, error)
	SignData(hash []byte, account types.Account, defaultPath string) ([]byte, error)
	Accounts(path string) []accounts.Account
	NewAccount(path string, passphrase string) (accounts.Account, error)
	Sign(digestHash []byte, prv *ecdsa.PrivateKey) ([]byte, error)
}

type AccountUtils struct{}

//This function returns all the accounts in form of array
func (accountUtils AccountUtils) Accounts(path string) []accounts.Account {
	ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	return ks.Accounts()
}

//This function takes path and pass phrase as input and returns the new account
func (accountUtils AccountUtils) NewAccount(path string, passphrase string) (accounts.Account, error) {
	ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: false}, ks)
	return ks.NewAccount(passphrase)
}

//This function takes hash in form of byte array and private key as input and returns signature as byte array
func (accountUtils AccountUtils) Sign(digestHash []byte, prv *ecdsa.PrivateKey) (sig []byte, err error) {
	return crypto.Sign(digestHash, prv)
}
