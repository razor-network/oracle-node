/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	jobManager "razor/pkg/bindings"
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
		fmt.Println("Job command called")

		client, err := ethclient.Dial("wss://goerli.infura.io/ws/v3/0d8f675eedad42cb82600b8e208b3465")
		if err != nil {
			log.Fatal(err)
		}

		header, err := client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			log.Fatal(err)
		}

		contractAddress := common.HexToAddress("0x264F48268dCEA2E264333B40dfe7d82fB8837E23")
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(0).Sub(header.Number, big.NewInt(1000)),
			ToBlock:   header.Number,
			Addresses: []common.Address{
				contractAddress,
			},
		}

		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			log.Fatal(err)
		}

		contractAbi, err := abi.JSON(strings.NewReader(string(jobManager.JobManagerABI)))
		if err != nil {
			log.Fatal(err)
		}

		for _, vLog := range logs {
			data, abc := contractAbi.Unpack("JobReported", vLog.Data)
			if abc != nil {
				log.Fatal(abc)
			}

			fmt.Println(data)

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
