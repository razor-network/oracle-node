package accounts

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts"
	"razor/core/types"
	"razor/logger"
	"strings"
)

var log = logger.NewLogger()
var AccountUtilsInterface AccountInterface

func CreateAccount(path string, password string, accountUtils AccountInterface) accounts.Account {
	newAcc, err := accountUtils.NewAccount(path, password)
	if err != nil {
		log.Fatal("Error in creating account: ", err)
	}
	return newAcc
}

func getPrivateKeyFromKeystore(keystorePath string, password string, accountUtils AccountInterface) *ecdsa.PrivateKey {
	jsonBytes, err := accountUtils.ReadFile(keystorePath)
	if err != nil {
		log.Fatal("Error in reading keystore: ", err)
	}
	key, err := accountUtils.DecryptKey(jsonBytes, password)
	if err != nil {
		log.Fatal("Error in fetching private key: ", err)
	}
	return key.PrivateKey
}

func GetPrivateKey(address string, password string, keystorePath string, accountUtils AccountInterface) *ecdsa.PrivateKey {
	allAccounts := accountUtils.Accounts(keystorePath)
	for _, account := range allAccounts {
		if strings.EqualFold(account.Address.Hex(), address) {
			return accountUtils.getPrivateKeyFromKeystore(account.URL.Path, password, accountUtils)
		}
	}
	return nil
}

func Sign(hash []byte, account types.Account, defaultPath string, accountUtils AccountInterface) ([]byte, error) {
	privateKey := accountUtils.GetPrivateKey(account.Address, account.Password, defaultPath, accountUtils)
	return accountUtils.Sign(hash, privateKey)
}
