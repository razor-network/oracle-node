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
	"log"
	"github.com/ethereum/go-ethereum/ethclient"
	"os"
  "github.com/ethereum/go-ethereum/common"
	"math"
	"math/big"
	"context"
	// "time"
	// "io/ioutil"
  // web3 "github.com/regcostajr/go-web3"
	// "github.com/regcostajr/go-web3/dto"
	// "github.com/regcostajr/go-web3/providers"
)

// stakeCmd represents the stake command
var stakeCmd = &cobra.Command{
	Use:   "stake",
	Short: "Stake some schells",
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
       fmt.Println("stake called")
			 client, err := ethclient.Dial("http://localhost:8545")
       if err != nil {
       log.Fatal(err)
       }

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
		// amount := args[0]
		address:= args[1]
		fmt.Println("account", address)
		// access schell balance from SchellingCoin
		// fmt.Println("schell balance", balance, "schells")
	  // if balance<amount {
		// 	fmt.Println("Not enough schells to stake")
		// 	os.Exit(0)
		// }
		account:=common.HexToAddress(address)
		weiBalance,err :=client.BalanceAt(context.Background(),account,nil)
		if err != nil {
			log.Fatal(err)
		}
		fbalance := new(big.Float)
    fbalance.SetString(weiBalance.String())
    ethBalance := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	  fmt.Println("ether balance", ethBalance, "eth")
    if ethBalance.Cmp(big.NewFloat(0.01)) == -1 {
			fmt.Println("Please fund this account with more ether to pay for tx fees")
			os.Exit(0)
		}
    // approve transaction
	for {
		//   epoch=??
		//   blocknumber,err := client.HeaderByNumber(context.Background(),nil)
		//   if err != nil {
		//   log.Fatal(err)
	  //     }
		//
		// fmt.Println("epoch",epoch)
		// fmt.Println("state",state)
		// if state!=0 {
		// 	fmt.Println("Can only stake during state 0 (commit). Retrying in 1 second...")
		// 	time.Sleep(time.Second)
		// }
	}
	fmt.Println("Sending stake transaction...")
  // nonce=?
  // let tx2 = await stakeManager.methods.stake(epoch, amountBN).send({
  //   from: account,
  // nonce: String(nonce)})
  // console.log(tx2.events)
  // return (tx2.events.Staked.event === 'Staked')
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
