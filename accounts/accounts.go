package accounts

import (
	"crypto/ecdsa"
	"io/ioutil"
	"razor/core/types"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	log "github.com/sirupsen/logrus"
)

func CreateAccount(path string, password string) accounts.Account {
	ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: false}, ks)
	newAcc, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal("Error in creating account: ", err)
	}
	return newAcc
}

func getPrivateKeyFromKeystore(keystorePath string, password string) *ecdsa.PrivateKey {
	jsonBytes, err := ioutil.ReadFile(keystorePath)
	if err != nil {
		log.Fatal("Error in reading keystore: ", err)
	}
	key, err := keystore.DecryptKey(jsonBytes, password)
	if err != nil {
		log.Fatal("Error in fetching private key: ", err)
	}
	return key.PrivateKey
}

func GetPrivateKey(address string, password string, keystorePath string) *ecdsa.PrivateKey {
	ks := keystore.NewKeyStore(keystorePath, keystore.StandardScryptN, keystore.StandardScryptP)
	allAccounts := ks.Accounts()
	for _, account := range allAccounts {
		if strings.EqualFold(account.Address.Hex(), address) {
			return getPrivateKeyFromKeystore(account.URL.Path, password)
		}
	}
	return nil
}

func Sign(hash []byte, account types.Account, defaultPath string) ([]byte, error) {
	privateKey := GetPrivateKey(account.Address, account.Password, defaultPath)
	return crypto.Sign(hash, privateKey)
}
