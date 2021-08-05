package cmd

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"razor/utils"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		privateKey := utils.PrivateKeyPrompt()
		// Remove 0x from the private key
		privateKey = privateKey[2:]
		log.Info("Enter password to protect keystore file")
		password := utils.PasswordPrompt()
		path := utils.GetDefaultPath()
		priv, err := crypto.HexToECDSA(privateKey)
		utils.CheckError("Error in parsing private key: ", err)
		importAccount(path, password, priv)
	},
}

func importAccount(path string, passphrase string, priv *ecdsa.PrivateKey) {
	ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.ImportECDSA(priv, passphrase)
	utils.CheckError("Error in importing account: ", err)
	log.Info("Account imported...")
	log.Info("Account Address: ", account.Address)
	log.Info("Keystore Path: ", account.URL)
}

func init() {
	rootCmd.AddCommand(importCmd)
}
