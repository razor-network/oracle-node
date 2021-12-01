package cmd

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"razor/utils"
)

var accountUtils accountInterface

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create command can be used to create new accounts",
	Long: `For a new user to start doing anything, an account is required. This command helps the user to create a new account secured by a password so that only that user would be able to use the account

Example: 
  ./razor create`,
	Run: initialiseCreate,
}

func initialiseCreate(cmd *cobra.Command, args []string) {
	utilsStruct := UtilsStruct{
		razorUtils:   razorUtils,
		accountUtils: accountUtils,
		cmdUtils: cmdUtils,
	}
	utilsStruct.executeCreate(cmd.Flags())
}

func (utilsStruct UtilsStruct) executeCreate(flagSet *pflag.FlagSet) {
	password := utilsStruct.razorUtils.AssignPassword(flagSet)
	account, err := utilsStruct.cmdUtils.Create(password, utilsStruct)
	utils.CheckError("Create error: ", err)
	log.Info("Account address: ", account.Address)
	log.Info("Keystore Path: ", account.URL)
}

func Create(password string, utilsStruct UtilsStruct) (accounts.Account, error) {
	path, err := utilsStruct.razorUtils.GetDefaultPath()
	if err != nil {
		log.Error("Error in fetching .razor directory")
		return accounts.Account{Address: common.Address{0x00}}, err
	}
	account := utilsStruct.accountUtils.CreateAccount(path, password)
	return account, nil
}

func init() {
	razorUtils = Utils{}
	accountUtils = AccountUtils{}
	cmdUtils = UtilsCmd{}

	rootCmd.AddCommand(createCmd)

	var (
		Password string
	)

	createCmd.Flags().StringVarP(&Password, "password", "", "", "password file path to protect the keystore")
}
