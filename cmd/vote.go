package cmd

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	types2 "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/logrusorgru/aurora/v3"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"math"
	"math/big"
	"razor/accounts"
	"razor/core"
	"razor/core/types"
	jobManager "razor/pkg/bindings"
	"razor/utils"
	"strings"
	"time"
)

var voteCmd = &cobra.Command{
	Use:   "vote",
	Short: "Start monitoring contract, commit, reveal, propose and dispute automatically",
	Long: `vote command allows you to participate in the voting of assets and earn rewards.

Example:
  ./razor vote --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		if err != nil {
			log.Fatal("Error in fetching config details: ", err)
		}

		password := utils.AssignPassword(cmd.Flags())
		rogueMode, _ := cmd.Flags().GetBool("rogue")
		client := utils.ConnectToClient(config.Provider)
		header, err := client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			log.Fatal(err)
		}
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
				handleBlock(client, account, latestHeader.Number, config, rogueMode)
			}

		}
	},
}

var (
	_committedData   []*big.Int
	lastVerification *big.Int
)

func handleBlock(client *ethclient.Client, account types.Account, blockNumber *big.Int, config types.Configurations, rogueMode bool) {
	state, err := utils.GetDelayedState(client, config.BufferPercent)
	if err != nil {
		log.Error("Error in getting state: ", err)
		return
	}
	epoch, err := utils.GetEpoch(client, account.Address)
	if err != nil {
		log.Error("Error in getting epoch: ", err)
		return
	}
	stakerId, err := utils.GetStakerId(client, account.Address)
	if err != nil {
		log.Error("Error in getting staker id: ", err)
		return
	}
	stakedAmount, err := utils.GetStake(client, account.Address, stakerId)
	if err != nil {
		log.Error("Error in getting staked amount: ", err)
		return
	}
	ethBalance, err := client.BalanceAt(context.Background(), common.HexToAddress(account.Address), nil)
	if err != nil {
		log.Errorf("Error in fetching balance of the account: %s\n%s", account.Address, err)
		return
	}
	minStakeAmount, err := utils.GetMinStakeAmount(client, account.Address)
	if err != nil {
		log.Error("Error in getting minimum stake amount: ", err)
		return
	}
	log.Info(aurora.Red("🔲 Block:"), aurora.Red(blockNumber), aurora.Yellow("⌛ Epoch:"), aurora.Yellow(epoch), aurora.Green("⏱️ State:"), aurora.Green(state), aurora.Blue("📒:"), aurora.Blue(account.Address), aurora.BrightBlue("👤 Staker ID:"), aurora.BrightBlue(stakerId), aurora.Cyan("💰Stake:"), aurora.Cyan(stakedAmount), aurora.Magenta("Ξ:"), aurora.Magenta(ethBalance))
	if stakedAmount.Cmp(minStakeAmount) < 0 {
		log.Error("Stake is below minimum required. Cannot vote.")
		AutoUnstakeAndWithdraw(client, account, stakedAmount, config)
		log.Fatal("Stopped voting as total stake is withdrawn now")
	}

	staker, err := utils.GetStaker(client, account.Address, stakerId)
	if err != nil {
		log.Error(err)
		return
	}

	switch state {
	case 0:
		lastCommit := staker.EpochLastCommitted
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
	case 1:
		lastReveal := staker.EpochLastRevealed
		if _committedData == nil || (lastReveal != nil && lastReveal.Cmp(epoch) >= 0) {
			break
		}
		secret := calculateSecret(account, epoch)
		if secret == nil {
			break
		}
		if err := HandleRevealState(staker, epoch); err != nil {
			log.Error(err)
			break
		}
		Reveal(client, _committedData, secret, account, account.Address, config)
	case 2:
		lastProposal, err := getLastProposedEpoch(client, blockNumber, stakerId)
		if err != nil {
			log.Error("Error in fetching last proposal: ", err)
			break
		}
		if lastProposal != nil && lastProposal.Cmp(epoch) >= 0 {
			break
		}
		lastProposal = epoch
		log.Info("Proposing block....")
		Propose(client, account, config, stakerId, epoch, rogueMode)
	case 3:
		if lastVerification != nil && lastVerification.Cmp(epoch) >= 0 {
			break
		}
		lastVerification = epoch
		HandleDispute(client, config, account, epoch)
	case -1:
		if config.WaitTime > 5 {
			time.Sleep(5 * time.Second)
			return
		}
	}
	utils.WaitTillNextNSecs(config.WaitTime)
	fmt.Println()
}

func getLastProposedEpoch(client *ethclient.Client, blockNumber *big.Int, stakerId *big.Int) (*big.Int, error) {
	maxRetries := 3
	numberOfBlocks := int64(core.StateLength) * core.NumberOfStates
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0).Sub(blockNumber, big.NewInt(numberOfBlocks)),
		ToBlock:   blockNumber,
		Addresses: []common.Address{
			common.HexToAddress(core.BlockManagerAddress),
		},
	}
	var (
		logs []types2.Log
		err  error
	)
	for retry := 1; retry <= maxRetries; retry++ {
		logs, err = client.FilterLogs(context.Background(), query)
		if err != nil {
			log.Error("Error in fetching logs: ", err)
			retryingIn := math.Pow(2, float64(retry))
			log.Infof("Retrying in %f seconds.....", retryingIn)
			time.Sleep(time.Duration(retryingIn) * time.Second)
			continue
		}
		break
	}
	if err != nil {
		return big.NewInt(0), err
	}
	contractAbi, err := abi.JSON(strings.NewReader(jobManager.BlockManagerABI))
	if err != nil {
		return big.NewInt(0), err
	}
	epochLastProposed := big.NewInt(0)
	for _, vLog := range logs {
		data, unpackErr := contractAbi.Unpack("Proposed", vLog.Data)
		if unpackErr != nil {
			log.Error(unpackErr)
			continue
		}
		if stakerId.Cmp(data[1].(*big.Int)) == 0 {
			epochLastProposed = data[0].(*big.Int)
		}
	}
	return epochLastProposed, nil
}

func calculateSecret(account types.Account, epoch *big.Int) []byte {
	hash := solsha3.SoliditySHA3([]string{"address", "uint256", "uint256", "string"}, []interface{}{account.Address, epoch.String(), core.ChainId.String(), "razororacle"})
	signedData, err := accounts.Sign(hash, account, utils.GetDefaultPath())
	if err != nil {
		log.Error("Error in signing the data: ", err)
		return nil
	}
	secret := solsha3.SoliditySHA3([]string{"string"}, []interface{}{hex.EncodeToString(signedData)})
	return secret
}

func AutoUnstakeAndWithdraw(client *ethclient.Client, account types.Account, amount *big.Int, config types.Configurations) {
	stakeManager := utils.GetStakeManager(client)
	txnArgs := types.TransactionOptions{
		Client:         client,
		AccountAddress: account.Address,
		Password:       account.Password,
		Amount:         amount,
		ChainId:        core.ChainId,
		Config:         config,
	}

	txnOpts := utils.GetTxnOpts(txnArgs)

	epoch, err := WaitForCommitState(client, account.Address, "unstake")
	utils.CheckError("Error in fetching epoch: ", err)

	stakerId, err := utils.GetStakerId(client, account.Address)
	utils.CheckError("Error in getting staker id: ", err)

	log.Info("Auto starting Unstake followed by Withdraw")
	log.Info("Unstaking coins")
	txn, err := stakeManager.Unstake(txnOpts, epoch, stakerId, txnArgs.Amount)
	utils.CheckError("Error in un-staking: ", err)
	log.Infof("Successfully unstaked %s sRazors", txnArgs.Amount)
	log.Info("Transaction hash: ", txn.Hash())
	utils.WaitForBlockCompletion(client, fmt.Sprintf("%s", txn.Hash()))

	log.Info("Starting withdrawal now...")
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Start()
	time.Sleep(time.Duration(core.EpochLength) * time.Second)
	s.Stop()
	checkForCommitStateAndWithdraw(client, account, config, stakerId)

}

func init() {
	rootCmd.AddCommand(voteCmd)

	var (
		Address  string
		Rogue    bool
		Password string
	)

	voteCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the staker")
	voteCmd.Flags().BoolVarP(&Rogue, "rogue", "r", false, "enable rogue mode to report wrong values")
	voteCmd.Flags().StringVarP(&Password, "password", "", "", "password path of the staker to protect the keystore")

	addrErr := voteCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
}
