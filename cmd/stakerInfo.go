package cmd

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"os"
	"razor/utils"
	"strconv"
)

var stakerInfoCmd = &cobra.Command{
	Use:   "stakerInfo",
	Short: "staker details",
	Long: `Provides the staker details like age, stake, maturity etc.

Example:
  ./razor stakerInfo --stakerId 2`,
	Run: initialiseStakerInfo,
}

func initialiseStakerInfo(cmd *cobra.Command, args []string) {
	cmdUtilsMockery.ExecuteStakerinfo(cmd.Flags())
}

func (*UtilsStructMockery) ExecuteStakerinfo(flagSet *pflag.FlagSet) {

	config, err := cmdUtilsMockery.GetConfigData()
	utils.CheckError("Error in getting config: ", err)

	client := razorUtilsMockery.ConnectToClient(config.Provider)

	stakerId, err := flagSetUtilsMockery.GetUint32StakerId(flagSet)
	utils.CheckError("Error in getting stakerId: ", err)

	err = cmdUtilsMockery.GetStakerInfo(client, stakerId)
	utils.CheckError("Error in getting staker info: ", err)

}

func (*UtilsStructMockery) GetStakerInfo(client *ethclient.Client, stakerId uint32) error {
	callOpts := razorUtilsMockery.GetOptions()
	stakerInfo, err := stakeManagerUtilsMockery.StakerInfo(client, &callOpts, stakerId)
	if err != nil {
		return err
	}
	maturity, err := stakeManagerUtilsMockery.GetMaturity(client, &callOpts, stakerInfo.Age)
	if err != nil {
		return err
	}
	epoch, err := razorUtilsMockery.GetEpoch(client)
	if err != nil {
		return err
	}
	influence, err := razorUtilsMockery.GetInfluenceSnapshot(client, stakerId, epoch)
	if err != nil {
		return err
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Staker Id", "Staker Address", "Stake", "Age", "Maturity", "Influence"})
	table.Append([]string{
		strconv.Itoa(int(stakerInfo.Id)),
		stakerInfo.Address.String(),
		stakerInfo.Stake.String(),
		strconv.Itoa(int(stakerInfo.Age)),
		strconv.Itoa(int(maturity)),
		influence.String(),
	})
	table.Render()
	return nil
}

func init() {
	razorUtilsMockery = &UtilsMockery{}
	stakeManagerUtilsMockery = StakeManagerUtilsMockery{}
	flagSetUtilsMockery = FLagSetUtilsMockery{}
	cmdUtilsMockery = &UtilsStructMockery{}

	rootCmd.AddCommand(stakerInfoCmd)

	var (
		StakerId uint32
	)

	stakerInfoCmd.Flags().Uint32VarP(&StakerId, "stakerId", "", 0, "staker id")
}
