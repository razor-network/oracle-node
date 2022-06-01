//Package cmd provides all functions related to command line
package cmd

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	pathPkg "path"
	"razor/path"
	"razor/utils"
	"strings"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "import can be used to import existing accounts into razor-go",
	Long: `If the user has their private key of an account, they can import that account into razor-go to perform further operations with razor-go.
Example:
  ./razor import --logFile importLogs`,
	Run: initialiseImport,
}

//This function initialises the ExecuteImport function
func initialiseImport(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteImport(cmd.Flags())
}

//This function sets the flags appropriately and executes the ImportAccount function
func (*UtilsStruct) ExecuteImport(flagSet *pflag.FlagSet) {
	razorUtils.AssignLogFile(flagSet)
	account, err := cmdUtils.ImportAccount()
	utils.CheckError("Import error: ", err)
	log.Info("Account Address: ", account.Address)
	log.Info("Keystore Path: ", account.URL)
}

//This function is used to import existing accounts into razor-go
func (*UtilsStruct) ImportAccount() (accounts.Account, error) {
	privateKey := razorUtils.PrivateKeyPrompt()
	// Remove 0x from the private key
	privateKey = strings.TrimPrefix(privateKey, "0x")
	log.Info("Enter password to protect keystore file")
	password := razorUtils.PasswordPrompt()
	razorPath, err := razorUtils.GetDefaultPath()
	if err != nil {
		log.Error("Error in fetching .razor directory")
		return accounts.Account{Address: common.Address{0x00}}, err
	}

	keystoreDir := pathPkg.Join(razorPath, "keystoreFiles")
	if _, err := path.OSUtilsInterface.Stat(keystoreDir); path.OSUtilsInterface.IsNotExist(err) {
		mkdirErr := path.OSUtilsInterface.Mkdir(keystoreDir, 0700)
		if mkdirErr != nil {
			return accounts.Account{Address: common.Address{0x00}}, mkdirErr
		}
	}

	priv, err := cryptoUtils.HexToECDSA(privateKey)
	if err != nil {
		log.Error("Error in parsing private key")
		return accounts.Account{Address: common.Address{0x00}}, err
	}
	account, err := keystoreUtils.ImportECDSA(keystoreDir, priv, password)
	if err != nil {
		log.Error("Error in importing account")
		return accounts.Account{Address: common.Address{0x00}}, err
	}
	log.Info("Account imported...")
	return account, nil
}

func init() {
	rootCmd.AddCommand(importCmd)

}
