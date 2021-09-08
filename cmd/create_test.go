package cmd

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

var AssignPasswordMock func(flagset *pflag.FlagSet) string

var GetDefaultPathMock func() string

var CreateAccountMock func(string, string) accounts.Account

type UtilsMock struct{}

type AccountsMock struct{}

func (r UtilsMock) AssignPassword(flagset *pflag.FlagSet) string {
	return AssignPasswordMock(flagset)
}

func (r UtilsMock) GetDefaultPath() string {
	return GetDefaultPathMock()
}

func (r AccountsMock) CreateAccount(path string, password string) accounts.Account {
	return CreateAccountMock(path, password)

}

func TestCreate(t *testing.T) {

	var flags *pflag.FlagSet
	var args []string
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
			Address: common.HexToAddress("0x000000000000000000000000000000000000dea1"),
			URL:     accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
		}
	}

	account := Create(flags, args, razorUtils, razorAccounts)
	log.Info("Account address: ", account.Address)
	log.Info("Keystore Path: ", account.URL)

	assert.Equal(t, account.Address, common.HexToAddress("0x000000000000000000000000000000000000dead"), "Created Address should match the supplied Address.")
}
