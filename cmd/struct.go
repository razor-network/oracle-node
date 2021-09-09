package cmd

import (
	ethereumAccounts "github.com/ethereum/go-ethereum/accounts"
	"github.com/spf13/pflag"
	"razor/accounts"
	"razor/utils"
)

type Utils struct{}
type Accounts struct{}

func (r Utils) AssignPassword(flags *pflag.FlagSet) string {
	return utils.AssignPassword(flags)
}

func (r Utils) GetDefaultPath() string {
	return utils.GetDefaultPath()
}

func (r Accounts) CreateAccount(path string, password string) ethereumAccounts.Account {
	return accounts.CreateAccount(path, password)
}
