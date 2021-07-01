package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// setDelegationCmd represents the setDelegation command
var setDelegationCmd = &cobra.Command{
	Use:   "setDelegation",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("setDelegation called")
	},
}

func init() {
	rootCmd.AddCommand(setDelegationCmd)

	var (
		State      bool
		Address    string
		Commission string
	)
	setDelegationCmd.Flags().BoolVarP(&State, "state", "a", true, "true for accepting delegation and false for not accepting")
	setDelegationCmd.Flags().StringVarP(&Address, "address", "", "", "your account address")
	setDelegationCmd.Flags().StringVarP(&Commission, "commissioned", "", "0", "commission")

	setDelegationCmd.MarkFlagRequired("amount")
	setDelegationCmd.MarkFlagRequired("address")
	setDelegationCmd.MarkFlagRequired("stakerId")

}
