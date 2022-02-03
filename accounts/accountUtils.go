package accounts

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"razor/core/types"
)

//go:generate mockery --name AccountInterface --output ./mocks/ --case=underscore

var AccountUtilsInterface AccountInterface

type AccountInterface interface {
	CreateAccount(string, string) accounts.Account
	GetPrivateKeyFromKeystore(string, string) *ecdsa.PrivateKey
	GetPrivateKey(string, string, string) *ecdsa.PrivateKey
	SignData([]byte, types.Account, string) ([]byte, error)
	Accounts(string) []accounts.Account
	NewAccount(string, string) (accounts.Account, error)
	DecryptKey([]byte, string) (*keystore.Key, error)
	Sign([]byte, *ecdsa.PrivateKey) ([]byte, error)
	ReadFile(string) ([]byte, error)
}

type AccountUtils struct{}

func (accountUtils AccountUtils) Accounts(path string) []accounts.Account {
	ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	return ks.Accounts()
}

func (accountUtils AccountUtils) NewAccount(path string, passphrase string) (accounts.Account, error) {
	ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: false}, ks)
	return ks.NewAccount(passphrase)
}

func (accountUtils AccountUtils) DecryptKey(jsonBytes []byte, password string) (*keystore.Key, error) {
	return keystore.DecryptKey(jsonBytes, password)
}

func (accountUtils AccountUtils) Sign(digestHash []byte, prv *ecdsa.PrivateKey) (sig []byte, err error) {
	return crypto.Sign(digestHash, prv)
}

func (accountUtils AccountUtils) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}
