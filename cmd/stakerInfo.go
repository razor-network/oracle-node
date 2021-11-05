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
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		utils.CheckError("Error in getting config: ", err)

		client := utils.ConnectToClient(config.Provider)

		address, _ := cmd.Flags().GetString("address")
		stakerId, _ := cmd.Flags().GetUint32("stakerId")
		err = GetStakerInfo(client, address, stakerId, stakeManagerUtils, razorUtils)
		if err != nil {
			log.Error("Error in getting staker info: ", err)
		}
	},
}

func GetStakerInfo(client *ethclient.Client, address string, stakerId uint32, stakeManagerUtils stakeManagerInterface, razorUtils utilsInterface) error {
	callOpts := razorUtils.GetOptions(false, address, "")
	stakerInfo, err := stakeManagerUtils.StakerInfo(client, &callOpts, stakerId)
	if err != nil {
		return err
	}
	maturity, err := stakeManagerUtils.GetMaturity(client, &callOpts, stakerInfo.Age)
	if err != nil {
		return err
	}
	//TODO: change this once v0.1.76 is merged
	epoch, err := razorUtils.GetEpoch(client, address)
	if err != nil {
		return err
	}
	influence, err := razorUtils.GetInfluenceSnapshot(client, address, stakerId, epoch)
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
