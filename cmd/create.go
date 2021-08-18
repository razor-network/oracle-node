package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"razor/accounts"
	"razor/utils"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create command can be used to create new accounts",
	Long:  `For a new user to start doing anything, an account is required. This command helps the user to create a new account secured by a password so that only that user would be able to use the account
	USAGE: ./razor create`,
	Run: func(cmd *cobra.Command, args []string) {
		path := utils.GetDefaultPath()
		password := utils.PasswordPrompt()
		account := accounts.CreateAccount(path, password)
		log.Info("Account address: ", account.Address)
		log.Info("Keystore Path: ", account.URL)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
