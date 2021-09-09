package cmd

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/spf13/pflag"
)

type UtilsMock struct{}

type AccountsMock struct{}

var AssignPasswordMock func(flagset *pflag.FlagSet) string

var GetDefaultPathMock func() string

var CreateAccountMock func(string, string) accounts.Account

func (r UtilsMock) AssignPassword(flagset *pflag.FlagSet) string {
	return AssignPasswordMock(flagset)
}

func (r UtilsMock) GetDefaultPath() string {
	return GetDefaultPathMock()
}

func (r AccountsMock) CreateAccount(path string, password string) accounts.Account {
	return CreateAccountMock(path, password)
}
