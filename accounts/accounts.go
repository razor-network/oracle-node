//Package account provides all account related functions
package accounts

import (
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"os"
	"razor/core/types"
	"razor/logger"
	"razor/path"
	"strings"
)

var (
	log        = logger.NewLogger()
	ksInstance *keystore.KeyStore
)

// InitializeKeystore directly initializes the global keystore instance.
func initializeKeystore(keystorePath string) {
	log.Info("Initialising keystoreInstance...")
	ksInstance = keystore.NewKeyStore(keystorePath, keystore.StandardScryptN, keystore.StandardScryptP)
}

//This function takes path and password as input and returns new account
func (AccountUtils) CreateAccount(keystorePath string, password string) accounts.Account {
	if _, err := path.OSUtilsInterface.Stat(keystorePath); path.OSUtilsInterface.IsNotExist(err) {
		mkdirErr := path.OSUtilsInterface.Mkdir(keystorePath, 0700)
		if mkdirErr != nil {
			log.Fatal("Error in creating directory: ", mkdirErr)
		}
	}
	newAcc, err := AccountUtilsInterface.NewAccount(keystorePath, password)
	if err != nil {
		log.Fatal("Error in creating account: ", err)
	}
	return newAcc
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

//This function takes address of account, password and keystore path as input and returns private key of account
func (AccountUtils) GetPrivateKey(address string, password string, keystoreDirPath string) (*ecdsa.PrivateKey, error) {
	if ksInstance == nil {
		initializeKeystore(keystoreDirPath)
	}

	allAccounts := ksInstance.Accounts()
	for _, account := range allAccounts {
		if strings.EqualFold(account.Address.Hex(), address) {
			return getPrivateKeyFromKeystore(account.URL.Path, password)
		}
	}
	return nil, errors.New("no keystore file found")
}

//This function takes hash, account and path as input and returns the signed data as array of byte
func (AccountUtils) SignData(hash []byte, account types.Account, defaultPath string) ([]byte, error) {
	privateKey, err := AccountUtilsInterface.GetPrivateKey(account.Address, account.Password, defaultPath)
	if err != nil {
		return nil, err
	}
	return AccountUtilsInterface.Sign(hash, privateKey)
}
