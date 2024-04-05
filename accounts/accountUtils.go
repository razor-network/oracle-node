//Package account provides all account related functions
package accounts

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"razor/core/types"
)

//type AccountInterface interface {
//	CreateAccount(keystorePath, password string) accounts.Account
//	GetPrivateKey(address, password string) (*ecdsa.PrivateKey, error)
//	SignData(hash []byte, account types.Account) ([]byte, error)
//	NewAccount(passphrase string) (accounts.Account, error)
//}

type AccountManager struct {
	Keystore *keystore.KeyStore
}

func NewAccountManager(keystorePath string) *AccountManager {
	ks := keystore.NewKeyStore(keystorePath, keystore.StandardScryptN, keystore.StandardScryptP)
	return &AccountManager{
		Keystore: ks,
	}
}

// InitAccountStruct initializes an Account struct with provided details.
func InitAccountStruct(address, password string, accountManager types.AccountManagerInterface) types.Account {
	return types.Account{
		Address:        address,
		Password:       password,
		AccountManager: accountManager,
	}
}
