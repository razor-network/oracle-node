//Package cmd provides all functions related to command line
package cmd

import (
	"context"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"os"
	"razor/RPC"
	"razor/core/types"
	"razor/logger"
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

//This function initialises the ExecuteStakerInfo function
func initialiseStakerInfo(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteStakerinfo(cmd.Flags())
}

//This function sets the flag appropriately and executes the GetStakerInfo function
func (*UtilsStruct) ExecuteStakerinfo(flagSet *pflag.FlagSet) {
	config, err := cmdUtils.GetConfigData()
	utils.CheckError("Error in getting config: ", err)
	log.Debugf("ExecuteStakerinfo: Config: %+v", config)

	client := razorUtils.ConnectToClient(config.Provider)
	logger.SetLoggerParameters(client, "")

	stakerId, err := flagSetUtils.GetUint32StakerId(flagSet)
	utils.CheckError("Error in getting stakerId: ", err)
	log.Debug("ExecuteStakerinfo: StakerId: ", stakerId)

	rpcManager, err := RPC.InitializeRPCManager(config.Provider)
	utils.CheckError("Error in initializing RPC Manager: ", err)

	rpcParameters := types.RPCParameters{
		RPCManager: rpcManager,
		Ctx:        context.Background(),
	}

	log.Debug("ExecuteStakerinfo: Calling GetStakerInfo() with argument stakerId = ", stakerId)
	err = cmdUtils.GetStakerInfo(rpcParameters, stakerId)
	utils.CheckError("Error in getting staker info: ", err)

}

//This function provides the staker details like age, stake, maturity etc.
func (*UtilsStruct) GetStakerInfo(rpcParameters types.RPCParameters, stakerId uint32) error {
	stakerInfo, err := razorUtils.StakerInfo(rpcParameters, stakerId)
	if err != nil {
		return err
	}
	maturity, err := razorUtils.GetMaturity(rpcParameters, stakerInfo.Age)
	if err != nil {
		return err
	}
	epoch, err := razorUtils.GetEpoch(rpcParameters)
	if err != nil {
		return err
	}
	influence, err := razorUtils.GetInfluenceSnapshot(rpcParameters, stakerId, epoch)
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
	rootCmd.AddCommand(stakerInfoCmd)

	var (
		StakerId uint32
	)

	stakerInfoCmd.Flags().Uint32VarP(&StakerId, "stakerId", "", 0, "staker id")
}
