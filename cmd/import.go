package cmd

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"razor/utils"
	"strings"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "import can be used to import existing accounts into razor-go",
	Long: `If the user has their private key of an account, they can import that account into razor-go to perform further operations with razor-go.
Example:
  ./razor import`,
	Run: initialiseImport,
}

func initialiseImport(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteImport()
}

func (*UtilsStruct) ExecuteImport() {
	account, err := cmdUtils.ImportAccount()
	utils.CheckError("Import error: ", err)
	log.Info("Account Address: ", account.Address)
	log.Info("Keystore Path: ", account.URL)
}

func (*UtilsStruct) ImportAccount() (accounts.Account, error) {
	privateKey := razorUtils.PrivateKeyPrompt()
	// Remove 0x from the private key
	privateKey = strings.TrimPrefix(privateKey, "0x")
	log.Info("Enter password to protect keystore file")
	password := razorUtils.PasswordPrompt()
	path, err := razorUtils.GetDefaultPath()
	if err != nil {
		log.Error("Error in fetching .razor directory")
		return accounts.Account{Address: common.Address{0x00}}, err
	}
	priv, err := cryptoUtils.HexToECDSA(privateKey)
	if err != nil {
		log.Error("Error in parsing private key")
		return accounts.Account{Address: common.Address{0x00}}, err
	}
	account, err := keystoreUtils.ImportECDSA(path, priv, password)
	if err != nil {
		log.Error("Error in importing account")
		return accounts.Account{Address: common.Address{0x00}}, err
	}
	log.Info("Account imported...")
	return account, nil
}

func init() {
	razorUtils = Utils{}
	keystoreUtils = KeystoreUtils{}
	cryptoUtils = CryptoUtils{}

	rootCmd.AddCommand(importCmd)
}
