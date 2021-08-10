package cmd

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"razor/utils"
)

//listAccountsCmd represents listAccounts command
var listAccountsCmd = &cobra.Command{
	Use:   "listAccounts",
	Short: "listAccounts command can be used to list all accessible accounts",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		path := utils.GetDefaultPath()
		ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
		allAccounts := ks.Accounts()
		log.Info("The available accounts are: ")
		for _, account := range allAccounts {
			log.Infof("%s\n", account.Address.String())
		}
	},
}

func init() {
	rootCmd.AddCommand(listAccountsCmd)
}
