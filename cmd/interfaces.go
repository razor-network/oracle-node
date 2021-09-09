package cmd

import (
	ethereumAccounts "github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
)

type utilsInterface interface {
	AssignPassword(*pflag.FlagSet) string
	GetDefaultPath() string
}

type accountsInterface interface {
	CreateAccount(string, string) ethereumAccounts.Account
}

type clientInterface interface {
	ConnectToClient(provider string) *ethclient.Client
}
