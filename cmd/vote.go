package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wealdtech/go-merkletree/keccak256"
	"math/big"
	"razor/accounts"
	"razor/core/types"
	"razor/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/logrusorgru/aurora/v3"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wealdtech/go-merkletree"
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
				handleBlock(client, address, password, latestHeader.Number, config)
			}
		}
	},
}

func handleBlock(client *ethclient.Client, address string, password string, blockNumber *big.Int, config types.Configurations) {
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
		hash := utils.GetKeccak256Hash([]byte(address), epoch.Bytes())
		signedData, err := accounts.Sign(hash, address, password, utils.GetDefaultPath())
		if err != nil {
			log.Error("Error in signing the data: ", err)
		}
		secret := utils.GetKeccak256Hash(signedData)
		handleCommitState(client, address, password, secret, config)
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

func handleCommitState(client *ethclient.Client, address string, password string, secret []byte, config types.Configurations) {
	jobs, err := utils.GetActiveJobs(client, address)
	if err != nil {
		log.Error("Error in getting active jobs: ", err)
	}
	var data []*big.Int
	for jobIndex := 0; jobIndex < len(jobs); jobIndex++ {
		response, err := utils.GetDataFromAPI(jobs[jobIndex].Url)
		datum := big.NewFloat(0)
		if err != nil {
			log.Error(err)
			continue
		}
		var parsedJSON map[string]interface{}
		err = json.Unmarshal(response, &parsedJSON)
		if err != nil {
			log.Error("Error in parsing data from API: ", err)
		}
		datum, err = utils.ConvertToNumber(parsedJSON[jobs[jobIndex].Selector])
		if err != nil {
			log.Error("Result is not a number")
		}
		dataToAppend := utils.MultiplyToEightDecimals(datum)
		data = append(data, dataToAppend)
	}
	if len(data) == 0 {
		data = append(data, big.NewInt(0))
	}
	log.Info("Data", data)

	if err := commit(client, data, secret, address, password, config); err != nil {
		log.Error("Error in committing data: ", err)
	}
}

func handleRevealState() {

}

func handleBlockProposalState() {

}

func handleDisputeState() {

}

func commit(client *ethclient.Client, data []*big.Int, secret []byte, account string, password string, config types.Configurations) error {
	state, err := utils.GetDelayedState(client)
	if err != nil {
		return err
	}
	if state != 0 {
		return errors.New("not commit state")
	}
	bytesData := utils.GetDataInBytes(data)
	tree, err := merkletree.NewUsing(bytesData, keccak256.New(), nil)
	if err != nil {
		return err
	}
	root := tree.Root()
	epoch, err := utils.GetEpoch(client, account)
	if err != nil {
		return err
	}
	commitments, err := utils.GetCommitments(client, account, epoch)
	if err != nil {
		return err
	}
	if !utils.AllZero(commitments) {
		return errors.New("already committed")
	}
	commitment := utils.GetKeccak256Hash(epoch.Bytes(), root, secret)
	voteManager := utils.GetVoteManager(client)
	txnOpts := utils.GetTxnOpts(types.TransactionOptions{
		Client:         client,
		Password:       password,
		AccountAddress: account,
		ChainId:        config.ChainId,
		GasMultiplier:  config.GasMultiplier,
	})
	// TODO: Find a better approach for commitmentToSend
	commitmentToSend := [32]byte{}
	copy(commitmentToSend[:], commitment)
	log.Infof("Committing: epoch: %s, root: %s, commitment: %s, secret: %s, account: %s", epoch, common.Bytes2Hex(root), common.Bytes2Hex(commitment), common.Bytes2Hex(secret), account)
	txn, err := voteManager.Commit(txnOpts, epoch, commitmentToSend)
	if err != nil {
		return err
	}
	log.Info("Commitment sent...\nTxn hash: ", txn.Hash())
	if utils.WaitForBlockCompletion(client, fmt.Sprintf("%s", txn.Hash())) == 0 {
		log.Error("Commit failed....")
	}
	return nil
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
