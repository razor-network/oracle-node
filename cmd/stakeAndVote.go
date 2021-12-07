package cmd

import (
	"razor/utils"

	"github.com/spf13/cobra"
)

var stakeAndVoteCmd = &cobra.Command{
	Use:   "stakeAndVote",
	Short: "Quick command to stake and start voting",
	Long: `This command is a combination of stake and vote command. User can use this command directly to stake and start voting.

Example:
  ./razor stakeAndVote --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --value 1000`,
	Run: initialiseStakeAndVote,
}

func initialiseStakeAndVote(cmd *cobra.Command, args []string) {
	utilsStruct := UtilsStruct{
		stakeManagerUtils: stakeManagerUtils,
		razorUtils:        razorUtils,
		transactionUtils:  transactionUtils,
		cmdUtils:          cmdUtils,
		flagSetUtils:      flagSetUtils,
		tokenManagerUtils: tokenManagerUtils,
	}
	password := utils.AssignPassword(cmd.Flags())
	utilsStruct.executeStake(cmd.Flags(), password)
	//TODO: Update this when vote is ready to be tested.
	executeVote(cmd.Flags(), password)
}

func init() {

	razorUtils = Utils{}
	tokenManagerUtils = TokenManagerUtils{}
	transactionUtils = TransactionUtils{}
	stakeManagerUtils = StakeManagerUtils{}
	flagSetUtils = FlagSetUtils{}

	rootCmd.AddCommand(stakeAndVoteCmd)
	var (
		Amount   string
		Address  string
		Password string
		Power    string
		Rogue    bool
	)

	stakeAndVoteCmd.Flags().StringVarP(&Amount, "value", "v", "0", "amount of Razors to stake")
	stakeAndVoteCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the staker")
	stakeAndVoteCmd.Flags().StringVarP(&Password, "password", "", "", "password path of staker to protect the keystore")
	stakeAndVoteCmd.Flags().StringVarP(&Power, "pow", "", "", "power of 10")
	stakeAndVoteCmd.Flags().BoolVarP(&Rogue, "rogue", "r", false, "enable rogue mode to report wrong values")

	amountErr := stakeAndVoteCmd.MarkFlagRequired("value")
	utils.CheckError("Value error: ", amountErr)
	addrErr := stakeAndVoteCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
}
