//Package account provides all account related functions
package accounts

import (
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"os"
	"razor/logger"
	"razor/path"
	"strings"
)

var log = logger.GetLogger()

//This function takes path and password as input and returns new account
func (am *AccountManager) CreateAccount(keystorePath string, password string) accounts.Account {
	if _, err := path.OSUtilsInterface.Stat(keystorePath); path.OSUtilsInterface.IsNotExist(err) {
		mkdirErr := path.OSUtilsInterface.Mkdir(keystorePath, 0700)
		if mkdirErr != nil {
			log.Fatal("Error in creating directory: ", mkdirErr)
		}
	}
	newAcc, err := am.NewAccount(password)
	if err != nil {
		log.Fatal("Error in creating account: ", err)
	}
	return newAcc
}

//This function takes path and pass phrase as input and returns the new account
func (am *AccountManager) NewAccount(passphrase string) (accounts.Account, error) {
	ks := am.Keystore
	accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: false}, ks)
	return ks.NewAccount(passphrase)
}

//This function takes address of account, password and keystore path as input and returns private key of account
func (am *AccountManager) GetPrivateKey(address string, password string) (*ecdsa.PrivateKey, error) {
	allAccounts := am.Keystore.Accounts()
	for _, account := range allAccounts {
		if strings.EqualFold(account.Address.Hex(), address) {
			return getPrivateKeyFromKeystore(account.URL.Path, password)
		}
	}
	return nil, errors.New("no keystore file found")
}

//This function takes and path of keystore and password as input and returns private key of account
func getPrivateKeyFromKeystore(keystoreFilePath string, password string) (*ecdsa.PrivateKey, error) {
	jsonBytes, err := os.ReadFile(keystoreFilePath)
	if err != nil {
		log.Error("Error in reading keystore: ", err)
		return nil, err
	}
	key, err := keystore.DecryptKey(jsonBytes, password)
	if err != nil {
		log.Error("Error in fetching private key: ", err)
		return nil, err
	}
	return key.PrivateKey, nil
}

//This function takes hash, account and path as input and returns the signed data as array of byte
func (am *AccountManager) SignData(hash []byte, address string, password string) ([]byte, error) {
	privateKey, err := am.GetPrivateKey(address, password)
	if err != nil {
		return nil, err
	}
	return crypto.Sign(hash, privateKey)
}
