/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"fmt"

	"github.com/spf13/cobra"

	"io/ioutil"

	web3 "github.com/regcostajr/go-web3"
	"github.com/regcostajr/go-web3/dto"
	"github.com/regcostajr/go-web3/providers"
)

// stakeCmd represents the stake command
var stakeCmd = &cobra.Command{
	Use:   "stake",
	Short: "Stake some schells",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stake called")

		// content, err := ioutil.ReadFile("../build/contracts/StakeManager.json")

		// type TruffleContract struct {
		// 	Abi      string `json:"abi"`
		// 	Bytecode string `json:"bytecode"`
		// }

		// var unmarshalResponse TruffleContract

		// json.Unmarshal(content, &unmarshalResponse)

		// var connection = web3.NewWeb3(providers.NewHTTPProvider("127.0.0.1:8545", 10, false))
		// bytecode := unmarshalResponse.Bytecode
		// contract, err := connection.Eth.NewContract(unmarshalResponse.Abi)
	},
}

func init() {
	rootCmd.AddCommand(stakeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stakeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stakeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
