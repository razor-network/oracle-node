package cmd

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {

	var flags *pflag.FlagSet
	//var args []string
	razorUtils := UtilsMock{}
	razorAccounts := AccountsMock{}

	AssignPasswordMock = func(flagset *pflag.FlagSet) string {
		return "test"
	}

	GetDefaultPathMock = func() string {
		return "test/path"
	}

	CreateAccountMock = func(string, string) accounts.Account {
		return accounts.Account{
			Address: common.HexToAddress("0x000000000000000000000000000000000000dead"),
			URL:     accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
		}
	}

	account := Create(flags, razorUtils, razorAccounts)
	log.Info("Account address: ", account.Address)
	log.Info("Keystore Path: ", account.URL)

	assert.Equal(t, account.Address, common.HexToAddress("0x000000000000000000000000000000000000dead"), "Created Address should match the supplied Address.")
}
