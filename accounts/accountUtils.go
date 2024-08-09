//Package account provides all account related functions
package accounts

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"razor/core/types"
)

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
