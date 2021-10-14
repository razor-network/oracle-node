package cmd

import (
	"context"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"razor/accounts"
	"razor/core"
	"razor/core/types"
	"razor/path"
	jobManager "razor/pkg/bindings"
	"razor/utils"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	types2 "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/spf13/cobra"
)

var voteCmd = &cobra.Command{
	Use:   "vote",
	Short: "Start monitoring contract, commit, reveal, propose and dispute automatically",
	Long: `vote command allows you to participate in the voting of assets and earn rewards.

Example:
  ./razor vote --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := GetConfigData()
		utils.CheckError("Error in fetching config details: ", err)

		password := utils.AssignPassword(cmd.Flags())
		rogueMode, _ := cmd.Flags().GetBool("rogue")
		client := utils.ConnectToClient(config.Provider)
		header, err := client.HeaderByNumber(context.Background(), nil)
		utils.CheckError("Error in getting block: ", err)

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
	lastVerification uint32
	blockConfirmed   uint32
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
	if stakerId == 0 {
		log.Error("Staker doesn't exist")
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
	actualStake, err := utils.ConvertWeiToEth(stakedAmount)
	if err != nil {
		log.Error("Error in converting stakedAmount from wei denomination")
	}
	actualBalance, err := utils.ConvertWeiToEth(ethBalance)
	if err != nil {
		log.Error("Error in converting ethBalance from wei denomination")
	}
	log.Debug("Block:", blockNumber, " Epoch:", epoch, " State:", utils.GetStateName(state), " Address:", account.Address, " Staker ID:", stakerId, " Stake:", actualStake, " Eth Balance:", actualBalance)
	if stakedAmount.Cmp(minStakeAmount) < 0 {
		log.Error("Stake is below minimum required. Cannot vote.")
		if stakedAmount.Cmp(big.NewInt(0)) == 0 {
			log.Error("Stopped voting as total stake is already withdrawn.")
		} else {
			log.Debug("Auto starting Unstake followed by Withdraw")
			AutoUnstakeAndWithdraw(client, account, stakedAmount, config)
			log.Error("Stopped voting as total stake is withdrawn now")
		}
		return
	}

	staker, err := utils.GetStaker(client, account.Address, stakerId)
	if err != nil {
		log.Error(err)
		return
	}

	switch state {
	case 0:
		lastCommit, err := utils.GetEpochLastCommitted(client, account.Address, stakerId)
		if err != nil {
			log.Error("Error in fetching last commit: ", err)
			break
		}
		if lastCommit >= epoch {
			log.Warnf("Cannot commit in epoch %d because last committed epoch is %d", epoch, lastCommit)
			break
		}
		secret := calculateSecret(account, epoch)
		if secret == nil {
			break
		}
		data, err := HandleCommitState(client, account.Address, epoch, razorUtils)
		if err != nil {
			log.Error("Error in getting active assets: ", err)
			break
		}
		commitTxn, err := Commit(client, data, secret, account, config, razorUtils, voteManagerUtils, transactionUtils)
		if err != nil {
			log.Error("Error in committing data: ", err)
			break
		}
		utils.WaitForBlockCompletion(client, commitTxn.String())
		_committedData = data
	case 1:
		lastReveal, err := utils.GetEpochLastRevealed(client, account.Address, stakerId)
		if err != nil {
			log.Error("Error in fetching last reveal: ", err)
			break
		}
		if _committedData == nil || lastReveal >= epoch {
			log.Warnf("Last reveal: %d", lastReveal)
			log.Warnf("Cannot reveal in epoch %d", epoch)
			break
		}
		secret := calculateSecret(account, epoch)
		if secret == nil {
			break
		}
		if err := HandleRevealState(client, account.Address, staker, epoch); err != nil {
			log.Error(err)
			break
		}
		log.Debug("Epoch last revealed: ", lastReveal)
		Reveal(client, _committedData, secret, account, account.Address, config)
	case 2:
		lastProposal, err := getLastProposedEpoch(client, blockNumber, stakerId)
		if err != nil {
			log.Error("Error in fetching last proposal: ", err)
			break
		}
		if lastProposal >= epoch {
			log.Warnf("Cannot propose in epoch %d because last proposed epoch is %d", epoch, lastProposal)
			break
		}
		lastReveal, err := utils.GetEpochLastRevealed(client, account.Address, stakerId)
		if err != nil {
			log.Error("Error in fetching last reveal: ", err)
			break
		}
		if lastReveal < epoch {
			log.Warnf("Cannot propose in epoch %d because last reveal was in epoch %d", epoch, lastReveal)
			break
		}
		proposeTxn, err := Propose(client, account, config, stakerId, epoch, rogueMode, razorUtils, proposeUtils, blockManagerUtils, transactionUtils)
		if err != nil {
			log.Error("Propose error: ", err)
			break
		}
		if proposeTxn != core.NilHash {
			utils.WaitForBlockCompletion(client, proposeTxn.String())
		}
	case 3:
		if lastVerification >= epoch {
			break
		}
		if rogueMode {
			log.Warn("Won't dispute in rogue mode..")
			break
		}
		lastVerification = epoch
		HandleDispute(client, config, account, epoch, razorUtils, proposeUtils)
	case 4:
		if lastVerification == epoch && blockConfirmed < epoch {
			txn, err := ClaimBlockReward(types.TransactionOptions{
				Client:          client,
				Password:        account.Password,
				AccountAddress:  account.Address,
				ChainId:         core.ChainId,
				Config:          config,
				ContractAddress: core.BlockManagerAddress,
				MethodName:      "claimBlockReward",
				ABI:             jobManager.BlockManagerABI,
			}, razorUtils, blockManagerUtils, transactionUtils)

			if err != nil {
				log.Error("ClaimBlockReward error: ", err)
				break
			}
			utils.WaitForBlockCompletion(client, txn.Hex())
			blockConfirmed = epoch
		}
	case -1:
		if config.WaitTime > 5 {
			time.Sleep(5 * time.Second)
			return
		}
	}
	utils.WaitTillNextNSecs(config.WaitTime)
	fmt.Println()
}

func getLastProposedEpoch(client *ethclient.Client, blockNumber *big.Int, stakerId uint32) (uint32, error) {
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
	for retry := 1; retry <= core.MaxRetries; retry++ {
		logs, err = client.FilterLogs(context.Background(), query)
		if err != nil {
			log.Error("Error in fetching logs: ", err)
			retryingIn := math.Pow(2, float64(retry))
			log.Debugf("Retrying in %f seconds.....", retryingIn)
			time.Sleep(time.Duration(retryingIn) * time.Second)
			continue
		}
		break
	}
	if err != nil {
		return 0, err
	}
	contractAbi, err := abi.JSON(strings.NewReader(jobManager.BlockManagerABI))
	if err != nil {
		return 0, err
	}
	epochLastProposed := uint32(0)
	for _, vLog := range logs {
		data, unpackErr := contractAbi.Unpack("Proposed", vLog.Data)
		if unpackErr != nil {
			log.Error(unpackErr)
			continue
		}
		if stakerId == data[1].(uint32) {
			epochLastProposed = data[0].(uint32)
		}
	}
	return epochLastProposed, nil
}

func calculateSecret(account types.Account, epoch uint32) []byte {
	hash := solsha3.SoliditySHA3([]string{"address", "uint32", "uint256", "string"}, []interface{}{account.Address, epoch, core.ChainId.String(), "razororacle"})
	razorPath, err := path.GetDefaultPath()
	if err != nil {
		log.Error("Error in fetching .razor directory: ", err)
	}
	signedData, err := accounts.Sign(hash, account, razorPath)
	if err != nil {
		log.Error("Error in signing the data: ", err)
		return nil
	}
	secret := solsha3.SoliditySHA3([]string{"string"}, []interface{}{hex.EncodeToString(signedData)})
	return secret
}

func AutoUnstakeAndWithdraw(client *ethclient.Client, account types.Account, amount *big.Int, config types.Configurations) {
	txnArgs := types.TransactionOptions{
		Client:         client,
		AccountAddress: account.Address,
		Password:       account.Password,
		Amount:         amount,
		ChainId:        core.ChainId,
		Config:         config,
	}
	stakerId, err := utils.GetStakerId(client, account.Address)
	utils.CheckError("Error in getting staker id: ", err)
	Unstake(txnArgs, stakerId)
	AutoWithdraw(txnArgs, stakerId)
}

func init() {

	razorUtils = Utils{}
	proposeUtils = ProposeUtils{}
	voteManagerUtils = VoteManagerUtils{}
	blockManagerUtils = BlockManagerUtils{}
	transactionUtils = TransactionUtils{}

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
