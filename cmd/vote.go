package cmd

import (
	"context"
	"encoding/json"
	"math/big"
	"razor/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/logrusorgru/aurora/v3"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var voteCmd = &cobra.Command{
	Use:   "vote",
	Short: "Start monitoring contract, commit, vote, propose and dispute automatically",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		if err != nil {
			log.Fatal("Error in fetching config details: ", err)
		}
		client := utils.ConnectToClient(config.Provider)
		header, err := client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			log.Fatal(err)
		}
		password, _ := cmd.Flags().GetString("password")
		address, _ := cmd.Flags().GetString("address")
		for {
			latestHeader, err := client.HeaderByNumber(context.Background(), nil)
			if err != nil {
				log.Error("Error in fetching block: ", err)
			}
			if latestHeader.Number.Cmp(header.Number) != 0 {
				header = latestHeader
				handleBlock(client, address, password, latestHeader.Number)
			}
		}
	},
}

func handleBlock(client *ethclient.Client, address string, password string, blockNumber *big.Int) {
	state, err := utils.GetDelayedState(client)
	if err != nil {
		log.Error("Error in getting state: ", err)
	}
	epoch, err := utils.GetEpoch(client, address)
	if err != nil {
		log.Error("Error in getting epoch: ", err)
	}
	stakerId, err := utils.GetStakerId(client, address)
	if err != nil {
		log.Error("Error in getting staker id: ", err)
	}
	stakedAmount, err := utils.GetStake(client, address, stakerId)
	if err != nil {
		log.Error("Error in getting staked amount: ", err)
	}
	ethBalance, err := client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		log.Errorf("Error in fetching balance of the account: %s\n%s", address, err)
	}
	minStakeAmount, err := utils.GetMinStakeAmount(client, address)
	if err != nil {
		log.Error("Error in getting minimum stake amount: ", err)
	}
	if stakedAmount.Cmp(minStakeAmount) < 0 {
		log.Error("Stake is below minimum required. Cannot vote.")
	}
	log.Info(aurora.Red("🔲 Block:"), aurora.Red(blockNumber), aurora.Yellow("⌛ Epoch:"), aurora.Yellow(epoch), aurora.Green("⏱️ State:"), aurora.Green(state), aurora.Blue("📒:"), aurora.Blue(address), aurora.BrightBlue("👤 Staker ID:"), aurora.BrightBlue(stakerId), aurora.Cyan("💰Stake:"), aurora.Cyan(stakedAmount), aurora.Magenta("Ξ:"), aurora.Magenta(ethBalance))

	switch state {
	case 0:
		handleCommitState(client, address, password)
		break
	case 1:
		handleRevealState()
		break
	case 2:
		handleBlockProposalState()
		break
	case 3:
		handleDisputeState()
		break
	}

}

func handleCommitState(client *ethclient.Client, address string, password string) {
	jobs, err := utils.GetActiveJobs(client, address)
	if err != nil {
		log.Error("Error in getting active jobs: ", err)
	}
	var datum []interface{}
	for jobIndex := 0; jobIndex < len(jobs); jobIndex++ {
		response, err := utils.GetDataFromAPI(jobs[jobIndex].Url)
		data := big.NewFloat(0)
		if err != nil {
			log.Error(err)
		}
		var parsedJSON map[string]interface{}
		err = json.Unmarshal(response, &parsedJSON)
		if err != nil {
			log.Error("Error in parsing data from API: ", err)
		}
		data, err = utils.ConvertToNumber(parsedJSON[jobs[jobIndex].Selector])
		if err != nil {
			log.Error("Result is not a number")
		}
		data = utils.MultiplyToEightDecimals(data)
		datum = append(datum, data)
	}
	if len(datum) == 0 {
		datum = append(datum, big.NewFloat(0))
	}
	log.Info("Datum", datum)
	commit(client, datum, "", address, password)
}

func handleRevealState() {

}

func handleBlockProposalState() {

}

func handleDisputeState() {

}

func commit(client *ethclient.Client, data []interface{}, secret string, account string, password string) {
	state, err := utils.GetDelayedState(client)
	if err != nil {
		log.Error("Error in fetching state: ", state)
	}
	if state != 0 {
		log.Error("Not commit state")
		return
	}
	//TODO: Figure out what to do for the merkle function
	//tree, err := merkletree.NewUsing([][]byte("data"), keccak256.New(), nil)
	//tree.Root()
}

func init() {
	rootCmd.AddCommand(voteCmd)

	var (
		Address  string
		Password string
	)

	voteCmd.Flags().StringVarP(&Address, "address", "", "", "address of the staker")
	voteCmd.Flags().StringVarP(&Password, "password", "", "", "password to unlock account")

	voteCmd.MarkFlagRequired("address")
	voteCmd.MarkFlagRequired("password")
}
