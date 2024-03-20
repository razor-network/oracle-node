//Package account provides all account related functions
package accounts

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"os"
	"path/filepath"
	"razor/core/types"
	"razor/logger"
	"razor/path"
	"regexp"
	"strings"
)

var log = logger.NewLogger()

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

//This function takes address of account, password and keystore path as input and returns private key of account
func (AccountUtils) GetPrivateKey(address string, password string, keystoreDirPath string) (*ecdsa.PrivateKey, error) {
	fileName, err := FindKeystoreFileForAddress(keystoreDirPath, address)
	if err != nil {
		log.Error("Error in finding keystore file for an address: ", err)
		return nil, err
	}

	keyJson, err := os.ReadFile(fileName)
	if err != nil {
		log.Error("Error in reading keystore: ", err)
		return nil, err
	}
	key, err := keystore.DecryptKey(keyJson, password)
	if err != nil {
		log.Error("Error in decrypting private key: ", err)
		return nil, err
	}

	// Check if the input address matches with address from keystore file in which password was matched
	if strings.EqualFold(key.Address.Hex(), address) {
		return key.PrivateKey, nil
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

// FindKeystoreFileForAddress matches the keystore file for the given address.
func FindKeystoreFileForAddress(keystoreDirPath, address string) (string, error) {
	normalizedAddress := strings.ToLower(strings.TrimPrefix(address, "0x"))
	regexPattern := fmt.Sprintf("^UTC--.*--%s$", regexp.QuoteMeta(normalizedAddress))
	re, err := regexp.Compile(regexPattern)
	if err != nil {
		log.Errorf("Error in compiling regex: %v", err)
		return "", err
	}

	var keystoreFilePath string
	err = filepath.WalkDir(keystoreDirPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err // Propagate errors encountered during traversal
		}
		if d.IsDir() {
			return nil // Skip directories, continue walking
		}
		if re.MatchString(d.Name()) { // Check if file name matches the regex
			keystoreFilePath = path
			return filepath.SkipDir // File found, no need to continue
		}
		return nil
	})

	if err != nil {
		log.Errorf("Error walking through keystore directory: %v", err)
		return "", err
	}

	if keystoreFilePath == "" {
		log.Errorf("No matching keystore file found for address %s", address)
		return "", errors.New("no matching keystore file found")
	}

	return keystoreFilePath, nil
}
