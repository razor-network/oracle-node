//Package cmd provides all functions related to command line
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"razor/core"
)

// contractAddressesCmd represents the contractAddresses command
var contractAddressesCmd = &cobra.Command{
	Use:   "contractAddresses",
	Short: "contractAddresses command can be used to list all contract addresses",
	Long:  `Provides the list of all contract addresses`,
	Run:   initialiseContractAddresses,
}

//This function initialises the ExecuteContractAddresses function
func initialiseContractAddresses(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteContractAddresses(cmd.Flags())
}

//This function sets the flag appropriatley and executes the ContractAddresses function
func (*UtilsStruct) ExecuteContractAddresses(flagSet *pflag.FlagSet) {
	log.Debug("Checking to assign log file...")
	razorUtils.AssignLogFile(flagSet)
	fmt.Println("The contract addresses are: ")
	cmdUtils.ContractAddresses()

}

//This function provides the list of all contract addresses
func (*UtilsStruct) ContractAddresses() {
	fmt.Println("StakeManagerAddress :", core.StakeManagerAddress)
	fmt.Println("RAZORAddress :", core.RAZORAddress)
	fmt.Println("CollectionManagerAddress :", core.CollectionManagerAddress)
	fmt.Println("VoteManagerAddress :", core.VoteManagerAddress)
	fmt.Println("BlockManagerAddress :", core.BlockManagerAddress)
}

func init() {
	rootCmd.AddCommand(contractAddressesCmd)
}
