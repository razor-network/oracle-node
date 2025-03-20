//Package cmd provides all functions related to command line
package cmd

import (
	"path/filepath"
	"razor/utils"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var listAccountsCmd = &cobra.Command{
	Use:   "listAccounts",
	Short: "listAccounts command can be used to list all accessible accounts",
	Long: `If the user wants to see what all accounts are existing in the razor-go environment, they can use this command to list down all the accounts.
Example:
  ./razor listAccounts`,
	Run: initialiseListAccounts,
}

//This function initialises the ExecuteListAccounts function
func initialiseListAccounts(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteListAccounts(cmd.Flags())
}

//This function sets the flag appropriately and executes the ListAccounts function
func (*UtilsStruct) ExecuteListAccounts(flagSet *pflag.FlagSet) {
	log.Debug("ExecuteListAccounts: Calling ListAccounts()...")
	allAccounts, err := cmdUtils.ListAccounts()
	utils.CheckError("ListAccounts error: ", err)
	log.Info("ExecuteListAccounts: The available accounts are: ")
	for _, account := range allAccounts {
		log.Infof("%s", account.Address.String())
	}
}

//This function is used to list all accessible accounts
func (*UtilsStruct) ListAccounts() ([]accounts.Account, error) {
	path, err := pathUtils.GetDefaultPath()
	if err != nil {
		log.Error("Error in fetching .razor directory")
		return nil, err
	}
	log.Debug("ListAccounts: .razor directory path: ", path)

	keystorePath := filepath.Join(path, "keystore_files")
	log.Debug("ListAccounts: Keystore path: ", keystorePath)
	return keystoreUtils.Accounts(keystorePath), nil
}

func init() {
	rootCmd.AddCommand(listAccountsCmd)
}
