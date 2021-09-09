package cmd

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/spf13/pflag"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var razorUtils utilsInterface
var razorAccounts accountsInterface

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create command can be used to create new accounts",
	Long: `For a new user to start doing anything, an account is required. This command helps the user to create a new account secured by a password so that only that user would be able to use the account

Example: 
  ./razor create`,
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		Create(flags, razorUtils, razorAccounts)
	},
}

func Create(flags *pflag.FlagSet, razorUtils utilsInterface, razorAccounts accountsInterface) accounts.Account {
	password := razorUtils.AssignPassword(flags)
	path := razorUtils.GetDefaultPath()
	account := razorAccounts.CreateAccount(path, password)
	log.Info("Account address: ", account.Address)
	log.Info("Keystore Path: ", account.URL)
	return account
}

func init() {
	rootCmd.AddCommand(createCmd)

	var (
		Password string
	)

	createCmd.Flags().StringVarP(&Password, "password", "", "", "password file path to protect the keystore")
}
