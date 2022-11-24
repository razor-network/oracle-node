//Package cmd provides all functions related to command line
package cmd

import (
	"fmt"
	"razor/core"
	"razor/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
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
	config, err := cmdUtils.GetConfigData()
	utils.CheckError("Error in getting config: ", err)
	log.Debug("Checking to assign log file...")
	fileUtils.AssignLogFile(flagSet, config)
	fmt.Println("The contract addresses are: ")
	cmdUtils.ContractAddresses()

}

//This function provides the list of all contract addresses
func (*UtilsStruct) ContractAddresses() {
	log.Info("StakeManagerAddress :", core.StakeManagerAddress)
	log.Info("RAZORAddress :", core.RAZORAddress)
	log.Info("CollectionManagerAddress :", core.CollectionManagerAddress)
	log.Info("VoteManagerAddress :", core.VoteManagerAddress)
	log.Info("BlockManagerAddress :", core.BlockManagerAddress)
}

func init() {
	rootCmd.AddCommand(contractAddressesCmd)
}
