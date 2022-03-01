package cmd

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"razor/accounts"
	"razor/core"
	"razor/core/types"
	jobManager "razor/pkg/bindings"
	"razor/utils"
	"strings"
	"time"

	"github.com/spf13/pflag"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
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
	Run: initializeVote,
}

func initializeVote(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteVote(cmd.Flags())
}

func (*UtilsStruct) ExecuteVote(flagSet *pflag.FlagSet) {
	config, err := cmdUtils.GetConfigData()
	utils.CheckError("Error in fetching config details: ", err)

	password := razorUtils.AssignPassword(flagSet)
	isRogue, err := flagSetUtils.GetBoolRogue(flagSet)
	utils.CheckError("Error in getting rogue status: ", err)

	rogueMode, err := flagSetUtils.GetStringSliceRogueMode(flagSet)
	utils.CheckError("Error in getting rogue modes: ", err)

	rogueData := types.Rogue{
		IsRogue:   isRogue,
		RogueMode: rogueMode,
	}
	client := razorUtils.ConnectToClient(config.Provider)
	address, err := flagSetUtils.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)

	account := types.Account{Address: address, Password: password}

	cmdUtils.HandleExit()

	if err := cmdUtils.Vote(context.Background(), config, client, rogueData, account); err != nil {
		log.Errorf("%s\n", err)
		razorUtils.Exit(1)
	}
}

func (*UtilsStruct) HandleExit() {
	// listen for CTRL+C
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()

	go func() {
		select {
		case <-signalChan:
			log.Warn("If you don't unstake and withdraw your coins, you may get inactivity penalty!")
		case <-ctx.Done():
		}
		<-signalChan // second signal, hard exit
		os.Exit(2)
	}()
}

func (*UtilsStruct) Vote(ctx context.Context, config types.Configurations, client *ethclient.Client, rogueData types.Rogue, account types.Account) error {
	header, err := utils.UtilsInterface.GetLatestBlockWithRetry(client)
	utils.CheckError("Error in getting block: ", err)
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			latestHeader, err := utils.UtilsInterface.GetLatestBlockWithRetry(client)
			if err != nil {
				log.Error("Error in fetching block: ", err)
				continue
			}
			if latestHeader.Number.Cmp(header.Number) != 0 {
				header = latestHeader
				cmdUtils.HandleBlock(client, account, latestHeader.Number, config, rogueData)
			}
		}
	}
}

var (
	_committedData   []*big.Int
	lastVerification uint32
	blockConfirmed   uint32
)

func (*UtilsStruct) HandleBlock(client *ethclient.Client, account types.Account, blockNumber *big.Int, config types.Configurations, rogueData types.Rogue) {
	state, err := razorUtils.GetDelayedState(client, config.BufferPercent)
	if err != nil {
		log.Error("Error in getting state: ", err)
		return
	}
	epoch, err := razorUtils.GetEpoch(client)
	if err != nil {
		log.Error("Error in getting epoch: ", err)
		return
	}

	stakerId, err := razorUtils.GetStakerId(client, account.Address)
	if err != nil {
		log.Error("Error in getting staker id: ", err)
		return
	}
	if stakerId == 0 {
		log.Error("Staker doesn't exist")
		return
	}

	stakedAmount, err := razorUtils.GetStake(client, stakerId)
	if err != nil {
		log.Error("Error in getting staked amount: ", err)
		return
	}
	ethBalance, err := utils.UtilsInterface.BalanceAtWithRetry(client, common.HexToAddress(account.Address))
	if err != nil {
		log.Errorf("Error in fetching balance of the account: %s\n%s", account.Address, err)
		return
	}
	minStakeAmount, err := utils.UtilsInterface.GetMinStakeAmount(client)
	if err != nil {
		log.Error("Error in getting minimum stake amount: ", err)
		return
	}
	actualStake, err := razorUtils.ConvertWeiToEth(stakedAmount)
	if err != nil {
		log.Error("Error in converting stakedAmount from wei denomination: ", err)
		return
	}
	actualBalance, err := razorUtils.ConvertWeiToEth(ethBalance)
	if err != nil {
		log.Error("Error in converting ethBalance from wei denomination: ", err)
		return
	}
	log.Infof("Block: %d Epoch: %d State: %s Address: %s Staker ID: %d Stake: %f Eth Balance: %f", blockNumber, epoch, razorUtils.GetStateName(state), account.Address, stakerId, actualStake, actualBalance)
	if stakedAmount.Cmp(minStakeAmount) < 0 {
		log.Error("Stake is below minimum required. Cannot vote.")
		if stakedAmount.Cmp(big.NewInt(0)) == 0 {
			log.Error("Stopped voting as total stake is already withdrawn.")
		} else {
			log.Debug("Auto starting Unstake followed by Withdraw")
			cmdUtils.AutoUnstakeAndWithdraw(client, account, stakedAmount, config)
			log.Error("Stopped voting as total stake is withdrawn now")
		}
		razorUtils.Exit(0)
	}

	staker, err := razorUtils.GetStaker(client, stakerId)
	if err != nil {
		log.Error(err)
		return
	}
	if staker.IsSlashed {
		log.Error("Staker is slashed.... cannot continue to vote!")
		razorUtils.Exit(0)
	}
	switch state {
	case 0:
		lastCommit, err := razorUtils.GetEpochLastCommitted(client, stakerId)
		if err != nil {
			log.Error("Error in fetching last commit: ", err)
			break
		}
		if lastCommit >= epoch {
			log.Debugf("Cannot commit in epoch %d because last committed epoch is %d", epoch, lastCommit)
			break
		}
		secret := cmdUtils.CalculateSecret(account, epoch)
		if secret == nil {
			break
		}
		data, err := cmdUtils.HandleCommitState(client, epoch, rogueData)
		if err != nil {
			log.Error("Error in getting active assets: ", err)
			break
		}
		commitTxn, err := cmdUtils.Commit(client, data, secret, account, config)
		if err != nil {
			log.Error("Error in committing data: ", err)
			break
		}
		if commitTxn != core.NilHash {
			razorUtils.WaitForBlockCompletion(client, commitTxn.String())
		}
		_committedData = data
		log.Debug("Saving committed data for recovery")
		fileName, err := cmdUtils.GetCommitDataFileName(account.Address)
		if err != nil {
			log.Error("Error in getting file name to save committed data: ", err)
			break
		}
		err = razorUtils.SaveDataToFile(fileName, epoch, _committedData)
		if err != nil {
			log.Errorf("Error in saving data to file %s: %t", fileName, err)
			break
		}
		log.Debug("Data saved!")
	case 1:
		lastReveal, err := razorUtils.GetEpochLastRevealed(client, stakerId)
		if err != nil {
			log.Error("Error in fetching last reveal: ", err)
			break
		}
		if lastReveal >= epoch {
			log.Debugf("Last reveal: %d", lastReveal)
			log.Debugf("Cannot reveal in epoch %d", epoch)
			break
		}
		if err := cmdUtils.HandleRevealState(client, staker, epoch); err != nil {
			log.Error(err)
			break
		}
		log.Debug("Epoch last revealed: ", lastReveal)
		if _committedData == nil {
			fileName, err := cmdUtils.GetCommitDataFileName(account.Address)
			if err != nil {
				log.Error("Error in getting file name to save committed data: ", err)
				break
			}
			epochInFile, committedDataFromFile, err := razorUtils.ReadDataFromFile(fileName)
			if err != nil {
				log.Errorf("Error in getting committed data from file %s: %t", fileName, err)
				break
			}
			if epochInFile != epoch {
				log.Errorf("File %s doesn't contain latest committed data: %t", fileName, err)
				break
			}
			_committedData = committedDataFromFile
		}
		secret := cmdUtils.CalculateSecret(account, epoch)
		if secret == nil {
			break
		}

		// Reveal wrong data if rogueMode contains reveal
		if rogueData.IsRogue && utils.Contains(rogueData.RogueMode, "reveal") {
			var rogueCommittedData []*big.Int
			for i := 0; i < len(_committedData); i++ {
				rogueCommittedData = append(rogueCommittedData, razorUtils.GetRogueRandomValue(10000000))
			}
			_committedData = rogueCommittedData
		}

		revealTxn, err := cmdUtils.Reveal(client, _committedData, secret, account, account.Address, config)
		if err != nil {
			log.Error("Reveal error: ", err)
			break
		}
		if revealTxn != core.NilHash {
			razorUtils.WaitForBlockCompletion(client, revealTxn.String())
		}
	case 2:
		lastProposal, err := cmdUtils.GetLastProposedEpoch(client, blockNumber, stakerId)
		if err != nil {
			log.Error("Error in fetching last proposal: ", err)
			break
		}
		if lastProposal >= epoch {
			log.Debugf("Cannot propose in epoch %d because last proposed epoch is %d", epoch, lastProposal)
			break
		}
		lastReveal, err := razorUtils.GetEpochLastRevealed(client, stakerId)
		if err != nil {
			log.Error("Error in fetching last reveal: ", err)
			break
		}
		if lastReveal < epoch {
			log.Debugf("Cannot propose in epoch %d because last reveal was in epoch %d", epoch, lastReveal)
			break
		}
		proposeTxn, err := cmdUtils.Propose(client, account, config, stakerId, epoch, rogueData)
		if err != nil {
			log.Error("Propose error: ", err)
			break
		}
		if proposeTxn != core.NilHash {
			razorUtils.WaitForBlockCompletion(client, proposeTxn.String())
		}
	case 3:
		if lastVerification >= epoch {
			break
		}
		lastVerification = epoch
		err := cmdUtils.HandleDispute(client, config, account, epoch, rogueData)
		if err != nil {
			log.Error(err)
			break
		}
	case 4:
		if lastVerification == epoch && blockConfirmed < epoch {
			txn, err := cmdUtils.ClaimBlockReward(types.TransactionOptions{
				Client:          client,
				Password:        account.Password,
				AccountAddress:  account.Address,
				ChainId:         core.ChainId,
				Config:          config,
				ContractAddress: core.BlockManagerAddress,
				MethodName:      "claimBlockReward",
				ABI:             jobManager.BlockManagerABI,
			})

			if err != nil {
				log.Error("ClaimBlockReward error: ", err)
				break
			}
			if txn != core.NilHash {
				razorUtils.WaitForBlockCompletion(client, txn.Hex())
				blockConfirmed = epoch
			}
		}
	case -1:
		if config.WaitTime > 5 {
			razorUtils.Sleep(5 * time.Second)
			return
		}
	}
	razorUtils.WaitTillNextNSecs(config.WaitTime)
	fmt.Println()
}

func (*UtilsStruct) GetLastProposedEpoch(client *ethclient.Client, blockNumber *big.Int, stakerId uint32) (uint32, error) {
	numberOfBlocks := int64(core.StateLength) * core.NumberOfStates
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0).Sub(blockNumber, big.NewInt(numberOfBlocks)),
		ToBlock:   blockNumber,
		Addresses: []common.Address{
			common.HexToAddress(core.BlockManagerAddress),
		},
	}
	logs, err := utils.UtilsInterface.FilterLogsWithRetry(client, query)
	if err != nil {
		return 0, err
	}
	contractAbi, err := utils.ABIInterface.Parse(strings.NewReader(jobManager.BlockManagerABI))
	if err != nil {
		return 0, err
	}
	epochLastProposed := uint32(0)
	for _, vLog := range logs {
		data, unpackErr := razorUtils.Unpack(contractAbi, "Proposed", vLog.Data)
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

func (*UtilsStruct) CalculateSecret(account types.Account, epoch uint32) []byte {
	hash := solsha3.SoliditySHA3([]string{"address", "uint32", "uint256", "string"}, []interface{}{account.Address, epoch, core.ChainId.String(), "razororacle"})
	razorPath, err := razorUtils.GetDefaultPath()
	if err != nil {
		log.Error("Error in fetching .razor directory: ", err)
	}
	signedData, err := accounts.AccountUtilsInterface.SignData(hash, account, razorPath)
	if err != nil {
		log.Error("Error in signing the data: ", err)
		return nil
	}
	secret := solsha3.SoliditySHA3([]string{"string"}, []interface{}{hex.EncodeToString(signedData)})
	return secret
}

func (*UtilsStruct) GetCommitDataFileName(address string) (string, error) {
	homeDir, err := razorUtils.GetDefaultPath()
	if err != nil {
		return "", err
	}
	return homeDir + "/" + address + "_data", nil
}

func (*UtilsStruct) AutoUnstakeAndWithdraw(client *ethclient.Client, account types.Account, amount *big.Int, config types.Configurations) {
	txnArgs := types.TransactionOptions{
		Client:         client,
		AccountAddress: account.Address,
		Password:       account.Password,
		Amount:         amount,
		ChainId:        core.ChainId,
		Config:         config,
	}

	stakerId, err := razorUtils.GetStakerId(client, account.Address)
	utils.CheckError("Error in getting staker id: ", err)

	_, err = cmdUtils.Unstake(config, client,
		types.UnstakeInput{
			Address:    account.Address,
			Password:   account.Password,
			ValueInWei: amount,
			StakerId:   stakerId,
		})
	utils.CheckError("Error in Unstake: ", err)
	err = cmdUtils.AutoWithdraw(txnArgs, stakerId)
	utils.CheckError("Error in AutoWithdraw: ", err)
}

func init() {

	razorUtils = Utils{}
	utils.Options = &utils.OptionsStruct{}
	utils.UtilsInterface = &utils.UtilsStruct{}
	cmdUtils = &UtilsStruct{}
	blockManagerUtils = BlockManagerUtils{}
	voteManagerUtils = VoteManagerUtils{}
	transactionUtils = TransactionUtils{}
	flagSetUtils = FLagSetUtils{}
	accounts.AccountUtilsInterface = accounts.AccountUtils{}

	rootCmd.AddCommand(voteCmd)

	var (
		Address   string
		Rogue     bool
		RogueMode []string
		Password  string
	)

	voteCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the staker")
	voteCmd.Flags().BoolVarP(&Rogue, "rogue", "r", false, "enable rogue mode to report wrong values")
	voteCmd.Flags().StringSliceVarP(&RogueMode, "rogueMode", "", []string{}, "type of rogue mode")
	voteCmd.Flags().StringVarP(&Password, "password", "", "", "password path of the staker to protect the keystore")

	addrErr := voteCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
}
