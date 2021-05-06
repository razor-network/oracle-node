package cmd

import (
	"github.com/ethereum/go-ethereum/accounts"
	log "github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func createAccount(path string, password string) accounts.Account {
	ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: false}, ks)
	newAcc, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal("Error in creating account: ", err)
	}
	return newAcc
}
