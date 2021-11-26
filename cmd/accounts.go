package cmd

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts"
	"razor/core/types"
	"strings"
)

func CreateAccount(path string, password string, utilsStruct UtilsStruct) accounts.Account {
	newAcc, err := utilsStruct.keystoreUtils.NewAccount(path, password)
	if err != nil {
		log.Fatal("Error in creating account: ", err)
	}
	return newAcc
}

func getPrivateKeyFromKeystore(keystorePath string, password string, utilsStruct UtilsStruct) *ecdsa.PrivateKey {
	jsonBytes, err := utilsStruct.razorUtils.ReadFile(keystorePath)
	if err != nil {
		log.Fatal("Error in reading keystore: ", err)
	}
	key, err := utilsStruct.keystoreUtils.DecryptKey(jsonBytes, password)
	if err != nil {
		log.Fatal("Error in fetching private key: ", err)
	}
	return key.PrivateKey
}

func GetPrivateKey(address string, password string, keystorePath string, utilsStruct UtilsStruct) *ecdsa.PrivateKey {
	allAccounts := utilsStruct.keystoreUtils.Accounts(keystorePath)
	for _, account := range allAccounts {
		if strings.EqualFold(account.Address.Hex(), address) {
			return utilsStruct.cmdUtils.getPrivateKeyFromKeystore(account.URL.Path, password, utilsStruct)
		}
	}
	return nil
}

func Sign(hash []byte, account types.Account, defaultPath string, utilsStruct UtilsStruct) ([]byte, error) {
	privateKey := utilsStruct.cmdUtils.GetPrivateKey(account.Address, account.Password, defaultPath, utilsStruct)
	return utilsStruct.cryptoUtils.Sign(hash, privateKey)
}
