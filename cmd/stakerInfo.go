package cmd

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
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
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		utils.CheckError("Error in getting config: ", err)

		client := utils.ConnectToClient(config.Provider)

		stakerId, _ := cmd.Flags().GetUint32("stakerId")
		utilsStruct := &UtilsStruct{
			razorUtils:        razorUtils,
			stakeManagerUtils: stakeManagerUtils,
		}
		err = utilsStruct.GetStakerInfo(client, stakerId)
		if err != nil {
			log.Error("Error in getting staker info: ", err)
		}
	},
}

func (utilsStruct *UtilsStruct) GetStakerInfo(client *ethclient.Client, stakerId uint32) error {
	callOpts := razorUtils.GetOptions(false, "", "")
	stakerInfo, err := utilsStruct.stakeManagerUtils.StakerInfo(client, &callOpts, stakerId)
	if err != nil {
		return err
	}
	maturity, err := utilsStruct.stakeManagerUtils.GetMaturity(client, &callOpts, stakerInfo.Age)
	if err != nil {
		return err
	}
	//TODO: change this once v0.1.76 is merged
	epoch, err := utilsStruct.razorUtils.GetEpoch(client)
	if err != nil {
		return err
	}
	influence, err := utilsStruct.razorUtils.GetInfluenceSnapshot(client, "", stakerId, epoch)
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
	razorUtils = Utils{}
	stakeManagerUtils = StakeManagerUtils{}

	rootCmd.AddCommand(stakerInfoCmd)

	var (
		StakerId uint32
	)

	stakerInfoCmd.Flags().Uint32VarP(&StakerId, "stakerId", "", 0, "staker id")
}
