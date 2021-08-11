package cmd

import (
	"context"
	"math/big"
	"razor/utils"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	assetManager "razor/pkg/bindings"
)

// jobCmd represents the job command
var jobCmd = &cobra.Command{
	Use:   "jobs",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		if err != nil {
			log.Fatal("Error in getting config: ", err)
		}

		client := utils.ConnectToClient(config.Provider)

		header, err := client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			log.Fatal(err)
		}
		assetManagerAddress := utils.GetAssetManagerAddress()
		contractAddress := common.HexToAddress(assetManagerAddress)
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(0),
			ToBlock:   header.Number,
			Addresses: []common.Address{
				contractAddress,
			},
		}

		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			log.Fatal(err)
		}

		contractAbi, err := abi.JSON(strings.NewReader(string(assetManager.AssetManagerABI)))
		if err != nil {
			log.Fatal(err)
		}

		for _, vLog := range logs {
			data, unpackErr := contractAbi.Unpack("JobReported", vLog.Data)
			if unpackErr != nil {
				log.Error(unpackErr)
				continue
			}

			log.Info(data)

		}
	},
}

func init() {
	rootCmd.AddCommand(jobCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jobCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jobCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
