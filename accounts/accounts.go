package accounts

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts"
	"razor/core/types"
	"razor/logger"
	"strings"
)

var log = logger.NewLogger()

func (AccountUtils) CreateAccount(path string, password string) accounts.Account {
	newAcc, err := AccountUtilsInterface.NewAccount(path, password)
	if err != nil {
		log.Fatal("Error in creating account: ", err)
	}
	return newAcc
}

func (AccountUtils) GetPrivateKeyFromKeystore(keystorePath string, password string) *ecdsa.PrivateKey {
	jsonBytes, err := AccountUtilsInterface.ReadFile(keystorePath)
	if err != nil {
		log.Fatal("Error in reading keystore: ", err)
	}
	key, err := AccountUtilsInterface.DecryptKey(jsonBytes, password)
	if err != nil {
		log.Fatal("Error in fetching private key: ", err)
	}
	return key.PrivateKey
}

func (AccountUtils) GetPrivateKey(address string, password string, keystorePath string) *ecdsa.PrivateKey {
	allAccounts := AccountUtilsInterface.Accounts(keystorePath)
	for _, account := range allAccounts {
		if strings.EqualFold(account.Address.Hex(), address) {
			return AccountUtilsInterface.GetPrivateKeyFromKeystore(account.URL.Path, password)
		}
	}
	return nil
}

func (AccountUtils) SignAccount(hash []byte, account types.Account, defaultPath string) ([]byte, error) {
	privateKey := AccountUtilsInterface.GetPrivateKey(account.Address, account.Password, defaultPath)
	return AccountUtilsInterface.Sign(hash, privateKey)
}
