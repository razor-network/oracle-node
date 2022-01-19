package cmd

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	razorAccounts "razor/accounts"

	"razor/utils"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create command can be used to create new accounts",
	Long: `For a new user to start doing anything, an account is required. This command helps the user to create a new account secured by a password so that only that user would be able to use the account

Example: 
  ./razor create`,
	Run: initialiseCreate,
}

func initialiseCreate(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteCreate(cmd.Flags())
}

func (*UtilsStruct) ExecuteCreate(flagSet *pflag.FlagSet) {
	password := razorUtils.AssignPassword(flagSet)
	account, err := cmdUtils.Create(password)
	utils.CheckError("Create error: ", err)
	log.Info("Account address: ", account.Address)
	log.Info("Keystore Path: ", account.URL)
}

func (*UtilsStruct) Create(password string) (accounts.Account, error) {
	path, err := razorUtils.GetDefaultPath()
	if err != nil {
		log.Error("Error in fetching .razor directory")
		return accounts.Account{Address: common.Address{0x00}}, err
	}
	account := razorAccounts.AccountUtilsInterface.CreateAccount(path, password)
	return account, nil
}

func init() {
	razorUtils = Utils{}
	cmdUtils = &UtilsStruct{}
	razorAccounts.AccountUtilsInterface = razorAccounts.AccountUtils{}

	rootCmd.AddCommand(createCmd)

	var (
		Password string
	)

	createCmd.Flags().StringVarP(&Password, "password", "", "", "password file path to protect the keystore")
}
