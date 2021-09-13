package cmd

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
	"razor/path"
	"razor/utils"
	"strings"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "import can be used to import existing accounts into razor-go",
	Long: `If the user has their private key of an account, they can import that account into razor-go to perform further operations with razor-go.
Example:
  ./razor import`,
	Run: func(cmd *cobra.Command, args []string) {
		privateKey := utils.PrivateKeyPrompt()
		// Remove 0x from the private key
		privateKey = strings.TrimPrefix(privateKey, "0x")
		log.Info("Enter password to protect keystore file")
		password := utils.PasswordPrompt()
		path, err := path.GetDefaultPath()
		utils.CheckError("Error in fetching .razor directory: ", err)
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
