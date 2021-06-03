package cmd

import (
	"context"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/logrusorgru/aurora/v3"
	"github.com/miguelmota/go-solidity-sha3"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"math/big"
	"razor/accounts"
	"razor/core/types"
	"razor/utils"
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
		account := types.Account{Address: address, Password: password}
		for {
			latestHeader, err := client.HeaderByNumber(context.Background(), nil)
			if err != nil {
				log.Error("Error in fetching block: ", err)
				continue
			}
			if latestHeader.Number.Cmp(header.Number) != 0 {
				header = latestHeader
				handleBlock(client, account, latestHeader.Number, config)
			}
		}
	},
}

var (
	_committedData   []*big.Int
	lastCommit       *big.Int
	lastReveal       *big.Int
	lastProposal     *big.Int
	lastElection     *big.Int
	lastVerification *big.Int
)

func handleBlock(client *ethclient.Client, account types.Account, blockNumber *big.Int, config types.Configurations) {
	state, err := utils.GetDelayedState(client)
	if err != nil {
		log.Error("Error in getting state: ", err)
	}
	epoch, err := utils.GetEpoch(client, account.Address)
	if err != nil {
		log.Error("Error in getting epoch: ", err)
	}
	stakerId, err := utils.GetStakerId(client, account.Address)
	if err != nil {
		log.Error("Error in getting staker id: ", err)
	}
	stakedAmount, err := utils.GetStake(client, account.Address, stakerId)
	if err != nil {
		log.Error("Error in getting staked amount: ", err)
	}
	ethBalance, err := client.BalanceAt(context.Background(), common.HexToAddress(account.Address), nil)
	if err != nil {
		log.Errorf("Error in fetching balance of the account: %s\n%s", account.Address, err)
	}
	minStakeAmount, err := utils.GetMinStakeAmount(client, account.Address)
	if err != nil {
		log.Error("Error in getting minimum stake amount: ", err)
	}
	log.Info(aurora.Red("🔲 Block:"), aurora.Red(blockNumber), aurora.Yellow("⌛ Epoch:"), aurora.Yellow(epoch), aurora.Green("⏱️ State:"), aurora.Green(state), aurora.Blue("📒:"), aurora.Blue(account.Address), aurora.BrightBlue("👤 Staker ID:"), aurora.BrightBlue(stakerId), aurora.Cyan("💰Stake:"), aurora.Cyan(stakedAmount), aurora.Magenta("Ξ:"), aurora.Magenta(ethBalance))
	if stakedAmount.Cmp(minStakeAmount) < 0 {
		log.Error("Stake is below minimum required. Cannot vote.")
		return
	}

	switch state {
	case 0:
		if lastCommit != nil && lastCommit.Cmp(epoch) >= 0 {
			break
		}
		secret := calculateSecret(account, epoch)
		if secret == nil {
			break
		}
		data := HandleCommitState(client, account.Address)
		if err := Commit(client, data, secret, account, config); err != nil {
			log.Error("Error in committing data: ", err)
			break
		}
		_committedData = data
		lastCommit = epoch
		break

	case 1:
		if _committedData == nil || (lastReveal != nil && lastReveal.Cmp(epoch) >= 0) {
			break
		}
		lastReveal = epoch
		secret := calculateSecret(account, epoch)
		if secret == nil {
			break
		}
		if err := HandleRevealState(client, account.Address, stakerId, epoch); err != nil {
			log.Error(err)
			break
		}
		Reveal(client, _committedData, secret, account, account.Address, config)
		break

	case 2:
		if lastElection != nil && lastElection.Cmp(epoch) >= 0 {
			break
		}
		lastElection = epoch
		if lastProposal != nil && lastProposal.Cmp(epoch) >= 0 {
			break
		}
		lastProposal = epoch
		log.Info("Proposing block....")
		Propose(client, account, config, stakerId, epoch)
		break

	case 3:
		if lastVerification != nil && lastVerification.Cmp(epoch) >= 0 {
			break
		}
		lastVerification = epoch
		HandleDispute(client, config, account, epoch)
		break
	}

}

func calculateSecret(account types.Account, epoch *big.Int) []byte {
	hash := solsha3.SoliditySHA3([]string{"address", "uint256"}, []interface{}{account.Address, epoch.String()})
	signedData, err := accounts.Sign(hash, account, utils.GetDefaultPath())
	if err != nil {
		log.Error("Error in signing the data: ", err)
		return nil
	}
	secret := solsha3.SoliditySHA3([]string{"string"}, []interface{}{hex.EncodeToString(signedData)})
	return secret
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
