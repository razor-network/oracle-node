//Package cmd provides all functions related to command line
package cmd

import (
	"path/filepath"
	"razor/utils"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create command can be used to create new accounts",
	Long: `For a new user to start doing anything, an account is required. This command helps the user to create a new account secured by a password so that only that user would be able to use the account

Example: 
  ./razor create --logFile createLogs`,
	Run: initialiseCreate,
}

//This function initialises the ExecuteCreate function
func initialiseCreate(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteCreate(cmd.Flags())
}

//This function sets the flags appropriately and executes the Create function
func (*UtilsStruct) ExecuteCreate(flagSet *pflag.FlagSet) {
	_, _, _, err := InitializeCommandDependencies(flagSet)
	utils.CheckError("Error in initialising command dependencies: ", err)
	log.Info("The password should be of minimum 8 characters containing least 1 uppercase, lowercase, digit and special character.")
	password := razorUtils.AssignPassword(flagSet)
	log.Debug("ExecuteCreate: Calling Create() with argument as input password")
	account, err := cmdUtils.Create(password)
	utils.CheckError("Create error: ", err)
	log.Info("ExecuteCreate: Account address: ", account.Address)
	log.Info("ExecuteCreate: Keystore Path: ", account.URL)
}

//This function is used to create the new account
func (*UtilsStruct) Create(password string) (accounts.Account, error) {
	razorPath, err := pathUtils.GetDefaultPath()
	if err != nil {
		log.Error("Error in fetching .razor directory")
		return accounts.Account{Address: common.Address{0x00}}, err
	}
	log.Debug("Create: .razor directory path: ", razorPath)
	accountManager, err := razorUtils.AccountManagerForKeystore()
	if err != nil {
		log.Error("Error in getting accounts manager for keystore: ", err)
		return accounts.Account{Address: common.Address{0x00}}, err
	}

	keystorePath := filepath.Join(razorPath, "keystore_files")
	account := accountManager.CreateAccount(keystorePath, password)
	return account, nil
}

func init() {
	rootCmd.AddCommand(createCmd)

	var (
		Password string
	)

	createCmd.Flags().StringVarP(&Password, "password", "", "", "password file path to protect the keystore")
}
