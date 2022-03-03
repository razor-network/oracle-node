package cmd

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/spf13/cobra"
	"razor/utils"
)

var listAccountsCmd = &cobra.Command{
	Use:   "listAccounts",
	Short: "listAccounts command can be used to list all accessible accounts",
	Long: `If the user wants to see what all accounts are existing in the razor-go environment, they can use this command to list down all the accounts.
Example:
  ./razor listAccounts`,
	Run: initialiseListAccounts,
}

func initialiseListAccounts(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteListAccounts()
}

func (*UtilsStruct) ExecuteListAccounts() {
	allAccounts, err := cmdUtils.ListAccounts()
	utils.CheckError("ListAccounts error: ", err)
	log.Info("The available accounts are: ")
	for _, account := range allAccounts {
		log.Infof("%s", account.Address.String())
	}
}

func (*UtilsStruct) ListAccounts() ([]accounts.Account, error) {
	path, err := razorUtils.GetDefaultPath()
	if err != nil {
		log.Error("Error in fetching .razor directory")
		return nil, err
	}

	return keystoreUtils.Accounts(path), nil
}

func init() {
	razorUtils = &Utils{}
	keystoreUtils = KeystoreUtils{}
	InitializeUtils()
	rootCmd.AddCommand(listAccountsCmd)
}
