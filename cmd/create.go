package cmd

import (
	"razor/accounts"

	ethereumAccounts "github.com/ethereum/go-ethereum/accounts"

	"razor/utils"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type utilsInterface interface {
	AssignPassword(*pflag.FlagSet) string
	GetDefaultPath() string
}

type accountsInterface interface {
	CreateAccount(string, string) ethereumAccounts.Account
}

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

var razorUtils utilsInterface
var razorAccounts accountsInterface

func Create(flags *pflag.FlagSet, args []string, razorUtils utilsInterface, razorAccounts accountsInterface) ethereumAccounts.Account {
	password := razorUtils.AssignPassword(flags)
	path := razorUtils.GetDefaultPath()
	account := razorAccounts.CreateAccount(path, password)
	log.Info("Account address: ", account.Address)
	log.Info("Keystore Path: ", account.URL)
	return account
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create command can be used to create new accounts",
	Long: `For a new user to start doing anything, an account is required. This command helps the user to create a new account secured by a password so that only that user would be able to use the account

Example: 
  ./razor create`,
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		Create(flags, args, razorUtils, razorAccounts)
	},
}

func init() {
	razorUtils = Utils{}
	razorAccounts = Accounts{}

	rootCmd.AddCommand(createCmd)

	var (
		Password string
	)

	createCmd.Flags().StringVarP(&Password, "password", "", "", "password file path to protect the keystore")
}
