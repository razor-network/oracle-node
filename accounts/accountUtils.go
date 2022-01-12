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
	SignAccount([]byte, types.Account, string) ([]byte, error)
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

//type AccountUtilsMock struct{}
//
//var CreateAccountMock func(string, string, AccountInterface) accounts.Account
//
//var getPrivateKeyFromKeystoreMock func(string, string, AccountInterface) *ecdsa.PrivateKey
//
//var GetPrivateKeyMock func(string, string, string, AccountInterface) *ecdsa.PrivateKey
//
//var AccountsMock func(string) []accounts.Account
//
//var NewAccountMock func(string, string) (accounts.Account, error)
//
//var DecryptKeyMock func(keyjson []byte, auth string) (*keystore.Key, error)
//
//var SignMock func([]byte, *ecdsa.PrivateKey) ([]byte, error)
//
//var ReadFileMock func(string) ([]byte, error)
//
//func (accountUtilsMock AccountUtilsMock) CreateAccount(path string, password string, accountUtil AccountInterface) accounts.Account {
//	return CreateAccountMock(path, password, accountUtil)
//}
//
//func (accountUtilsMock AccountUtilsMock) getPrivateKeyFromKeystore(keystorePath string, password string, accountUtil AccountInterface) *ecdsa.PrivateKey {
//	return getPrivateKeyFromKeystoreMock(keystorePath, password, accountUtil)
//}
//
//func (accountUtilsMock AccountUtilsMock) GetPrivateKey(address string, password string, keystorePath string, accountUtil AccountInterface) *ecdsa.PrivateKey {
//	return GetPrivateKeyMock(address, password, keystorePath, accountUtil)
//}
//
//func (accountUtilsMock AccountUtilsMock) Accounts(path string) []accounts.Account {
//	return AccountsMock(path)
//}
//
//func (accountUtilsMock AccountUtilsMock) NewAccount(path string, passphrase string) (accounts.Account, error) {
//	return NewAccountMock(path, passphrase)
//}
//
//func (accountUtilsMock AccountUtilsMock) DecryptKey(jsonBytes []byte, password string) (*keystore.Key, error) {
//	return DecryptKeyMock(jsonBytes, password)
//}
//
//func (accountUtilsMock AccountUtilsMock) Sign(digestHash []byte, prv *ecdsa.PrivateKey) (sig []byte, err error) {
//	return SignMock(digestHash, prv)
//}
//
//func (accountUtilsMock AccountUtilsMock) ReadFile(filename string) ([]byte, error) {
//	return ReadFileMock(filename)
//}
