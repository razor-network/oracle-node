package cmd

import (
	"razor/utils"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/spf13/cobra"
)

var keystoreUtils keystoreInterface

var listAccountsCmd = &cobra.Command{
	Use:   "listAccounts",
	Short: "listAccounts command can be used to list all accessible accounts",
	Long: `If the user wants to see what all accounts are existing in the razor-go environment, they can use this command to list down all the accounts.
Example:
  ./razor listAccounts`,
	Run: func(cmd *cobra.Command, args []string) {
		allAccounts, err := listAccounts(razorUtils, keystoreUtils)
		utils.CheckError("ListAccounts error: ", err)
		log.Info("The available accounts are: ")
		for _, account := range allAccounts {
			log.Infof("%s", account.Address.String())
		}
	},
}

func listAccounts(razorUtils utilsInterface, keystoreUtils keystoreInterface) ([]accounts.Account, error) {
	path, err := razorUtils.GetDefaultPath()
	if err != nil {
		log.Error("Error in fetching .razor directory")
		return nil, err
	}

	return keystoreUtils.Accounts(path), nil
}

func init() {
	razorUtils = Utils{}
	keystoreUtils = KeystoreUtils{}
	rootCmd.AddCommand(listAccountsCmd)
}
