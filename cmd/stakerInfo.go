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
  ./razor stakerInfo --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --stakerId 2`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		utils.CheckError("Error in getting config: ", err)

		client := utils.ConnectToClient(config.Provider)

		address, _ := cmd.Flags().GetString("address")
		stakerId, _ := cmd.Flags().GetUint32("stakerId")
		utilsStruct := &UtilsStruct{
			razorUtils:        razorUtils,
			stakeManagerUtils: stakeManagerUtils,
		}
		err = utilsStruct.GetStakerInfo(client, address, stakerId)
		if err != nil {
			log.Error("Error in getting staker info: ", err)
		}
	},
}

func (utilsStruct *UtilsStruct) GetStakerInfo(client *ethclient.Client, address string, stakerId uint32) error {
	callOpts := razorUtils.GetOptions(false, address, "")
	stakerInfo, err := utilsStruct.stakeManagerUtils.StakerInfo(client, &callOpts, stakerId)
	if err != nil {
		return err
	}
	maturity, err := utilsStruct.stakeManagerUtils.GetMaturity(client, &callOpts, stakerInfo.Age)
	if err != nil {
		return err
	}
	//TODO: change this once v0.1.76 is merged
	epoch, err := utilsStruct.razorUtils.GetEpoch(client, address)
	if err != nil {
		return err
	}
	influence, err := utilsStruct.razorUtils.GetInfluenceSnapshot(client, address, stakerId, epoch)
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
		Address  string
		StakerId uint32
	)

	stakerInfoCmd.Flags().StringVarP(&Address, "address", "a", "", "user's address")
	stakerInfoCmd.Flags().Uint32VarP(&StakerId, "stakerId", "", 0, "staker id")
}
