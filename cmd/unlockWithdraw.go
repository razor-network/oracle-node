package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// unlockWithdrawCmd represents the unlockWithdraw command
var unlockWithdrawCmd = &cobra.Command{
	Use:   "unlockWithdraw",
	Short: "InitiateWithdraw your razors once withdraw lock has passed",
	Long:  `unlockWithdraw has to be called once the withdraw lock period is over to get back all the razor tokens into your account`,
	Run:   initializeUnlockWithdraw,
}

func initializeUnlockWithdraw(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteUnlockWithdraw(cmd.Flags())
}

func (*UtilsStruct) ExecuteUnlockWithdraw(flagSet *pflag.FlagSet) {

}

func init() {
	rootCmd.AddCommand(unlockWithdrawCmd)

}
