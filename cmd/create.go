package cmd

import (
	"github.com/spf13/cobra"
	"razor/accounts"
	"razor/path"
	"razor/utils"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create command can be used to create new accounts",
	Long: `For a new user to start doing anything, an account is required. This command helps the user to create a new account secured by a password so that only that user would be able to use the account

Example: 
  ./razor create`,
	Run: func(cmd *cobra.Command, args []string) {

		password := utils.AssignPassword(cmd.Flags())
		path, err := path.GetDefaultPath()
		utils.CheckError("Error in fetching .razor directory", err)
		account := accounts.CreateAccount(path, password)
		log.Info("Account address: ", account.Address)
		log.Info("Keystore Path: ", account.URL)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	var (
		Password string
	)

	createCmd.Flags().StringVarP(&Password, "password", "", "", "password file path to protect the keystore")
}
