package cmd

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"razor/accounts"
	"razor/core"
	"razor/core/types"
	"razor/logger"
	"razor/pkg/bindings"
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
	address, err := flagSetUtils.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)

	logger.Address = address

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

	account := types.Account{Address: address, Password: password}

	cmdUtils.HandleExit()

	if err := cmdUtils.Vote(context.Background(), config, client, rogueData, account); err != nil {
		log.Errorf("%s\n", err)
		osUtils.Exit(1)
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
	_commitData      types.CommitData
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
	staker, err := razorUtils.GetStaker(client, stakerId)
	if err != nil {
		log.Error(err)
		return
	}
	stakedAmount := staker.Stake

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
	log.Infof("Block: %d Epoch: %d State: %s Staker ID: %d Stake: %f Eth Balance: %f", blockNumber, epoch, razorUtils.GetStateName(state), stakerId, actualStake, actualBalance)
	if stakedAmount.Cmp(minStakeAmount) < 0 {
		log.Error("Stake is below minimum required. Cannot vote.")
		if stakedAmount.Cmp(big.NewInt(0)) == 0 {
			log.Error("Stopped voting as total stake is already withdrawn.")
		} else {
			log.Debug("Auto starting Unstake followed by InitiateWithdraw")
			cmdUtils.AutoUnstakeAndWithdraw(client, account, stakedAmount, config)
			log.Error("Stopped voting as total stake is withdrawn now")
		}
		osUtils.Exit(0)
	}

	if staker.IsSlashed {
		log.Error("Staker is slashed.... cannot continue to vote!")
		osUtils.Exit(0)
	}

	switch state {
	case 0:
		err := cmdUtils.InitiateCommit(client, config, account, epoch, stakerId, rogueData)
		if err != nil {
			log.Error(err)
			break
		}
	case 1:
		err := cmdUtils.InitiateReveal(client, config, account, epoch, staker, rogueData)
		if err != nil {
			log.Error(err)
			break
		}
	case 2:
		err := InitiatePropose(client, config, account, epoch, staker, blockNumber, rogueData)
		if err != nil {
			log.Error(err)
			break
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
				ABI:             bindings.BlockManagerABI,
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
			timeUtils.Sleep(5 * time.Second)
			return
		}
	}
	razorUtils.WaitTillNextNSecs(config.WaitTime)
	fmt.Println()
}

func (*UtilsStruct) InitiateCommit(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, stakerId uint32, rogueData types.Rogue) error {
	lastCommit, err := razorUtils.GetEpochLastCommitted(client, stakerId)
	if err != nil {
		return errors.New("Error in fetching last commit: " + err.Error())
	}
	if lastCommit >= epoch {
		log.Debugf("Cannot commit in epoch %d because last committed epoch is %d", epoch, lastCommit)
		return nil
	}

	secret, err := cmdUtils.CalculateSecret(account, epoch)
	if err != nil {
		return err
	}

	log.Debugf("Secret: %s", hex.EncodeToString(secret))

	salt, err := cmdUtils.GetSalt(client, epoch)
	if err != nil {
		return err
	}
	log.Debugf("Salt: %s", hex.EncodeToString(salt[:]))

	seed := solsha3.SoliditySHA3([]string{"bytes32", "bytes32"}, []interface{}{"0x" + hex.EncodeToString(salt[:]), "0x" + hex.EncodeToString(secret)})
	log.Debugf("Seed: %s", hex.EncodeToString(seed[:]))

	commitData, err := cmdUtils.HandleCommitState(client, epoch, seed, rogueData)
	if err != nil {
		return errors.New("Error in getting active assets: " + err.Error())
	}

	_commitData = commitData

	merkleTree := utils.MerkleInterface.CreateMerkle(commitData.Leaves)
	commitTxn, err := cmdUtils.Commit(client, seed, utils.MerkleInterface.GetMerkleRoot(merkleTree), epoch, account, config)
	if err != nil {
		return errors.New("Error in committing data: " + err.Error())
	}
	if commitTxn != core.NilHash {
		razorUtils.WaitForBlockCompletion(client, commitTxn.String())
	}

	//TODO: Need to save the entire commitData, which includes AssignedCollections, SeqAllottedCollections and Leaves to construct merkle tree
	log.Debug("Saving committed data for recovery")
	fileName, err := cmdUtils.GetCommitDataFileName(account.Address)
	if err != nil {
		return errors.New("Error in getting file name to save committed data: " + err.Error())
	}

	err = razorUtils.SaveDataToFile(fileName, epoch, commitData.Leaves)
	if err != nil {
		return errors.New("Error in saving data to file" + fileName + ": " + err.Error())
	}
	log.Debug("Data saved!")
	return nil
}

func (*UtilsStruct) InitiateReveal(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, staker bindings.StructsStaker, rogueData types.Rogue) error {
	lastReveal, err := razorUtils.GetEpochLastRevealed(client, staker.Id)
	if err != nil {
		return errors.New("Error in fetching last reveal: " + err.Error())
	}
	if lastReveal >= epoch {
		log.Debugf("Since last reveal was at epoch: %d, won't reveal again in epoch: %d", lastReveal, epoch)
		return nil
	}

	if err := cmdUtils.HandleRevealState(client, staker, epoch); err != nil {
		log.Error(err)
	}
	log.Debug("Epoch last revealed: ", lastReveal)

	//TODO: Reveal rogue and fetch committed data from file

	//if _committedData == nil {
	//	fileName, err := cmdUtils.GetCommitDataFileName(account.Address)
	//	if err != nil {
	//		log.Error("Error in getting file name to save committed data: ", err)
	//		break
	//	}
	//	epochInFile, committedDataFromFile, err := razorUtils.ReadDataFromFile(fileName)
	//	if err != nil {
	//		log.Errorf("Error in getting committed data from file %s: %t", fileName, err)
	//		break
	//	}
	//	if epochInFile != epoch {
	//		log.Errorf("File %s doesn't contain latest committed data: %t", fileName, err)
	//		break
	//	}
	//	_committedData = committedDataFromFile
	//}
	//if rogueData.IsRogue && utils.Contains(rogueData.RogueMode, "reveal") {
	//	var rogueCommittedData []*big.Int
	//	for i := 0; i < len(_committedData); i++ {
	//		rogueCommittedData = append(rogueCommittedData, razorUtils.GetRogueRandomValue(10000000))
	//	}
	//	_committedData = rogueCommittedData
	//}

	secret, err := cmdUtils.CalculateSecret(account, epoch)
	if err != nil {
		return err
	}

	revealTxn, err := cmdUtils.Reveal(client, config, account, epoch, _commitData, secret)
	if err != nil {
		return errors.New("Reveal error: " + err.Error())
	}
	if revealTxn != core.NilHash {
		razorUtils.WaitForBlockCompletion(client, revealTxn.String())
	}
	return nil
}

func InitiatePropose(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, staker bindings.StructsStaker, blockNumber *big.Int, rogueData types.Rogue) error {
	lastProposal, err := cmdUtils.GetLastProposedEpoch(client, blockNumber, staker.Id)
	if err != nil {
		return errors.New("Error in fetching last proposal: " + err.Error())
	}
	if lastProposal >= epoch {
		log.Debugf("Since last propose was at epoch: %d, won't propose again in epoch: %d", epoch, lastProposal)
		return nil
	}
	lastReveal, err := razorUtils.GetEpochLastRevealed(client, staker.Id)
	if err != nil {
		return errors.New("Error in fetching last reveal: " + err.Error())
	}
	if lastReveal < epoch {
		log.Debugf("Cannot propose in epoch %d because last reveal was in epoch %d", epoch, lastReveal)
		return nil
	}

	proposeTxn, err := cmdUtils.Propose(client, config, account, staker, epoch, blockNumber, rogueData)
	if err != nil {
		return errors.New("Propose error: " + err.Error())
	}
	if proposeTxn != core.NilHash {
		razorUtils.WaitForBlockCompletion(client, proposeTxn.String())
	}
	return nil
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
	contractAbi, err := utils.Options.Parse(strings.NewReader(bindings.BlockManagerABI))
	if err != nil {
		return 0, err
	}
	epochLastProposed := uint32(0)
	for _, vLog := range logs {
		data, unpackErr := abiUtils.Unpack(contractAbi, "Proposed", vLog.Data)
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

func (*UtilsStruct) CalculateSecret(account types.Account, epoch uint32) ([]byte, error) {
	hash := solsha3.SoliditySHA3([]string{"address", "uint32", "uint256", "string"}, []interface{}{account.Address, epoch, core.ChainId.String(), "razororacle"})
	razorPath, err := razorUtils.GetDefaultPath()
	if err != nil {
		return nil, errors.New("Error in fetching .razor directory: " + err.Error())
	}
	signedData, err := accounts.AccountUtilsInterface.SignData(hash, account, razorPath)
	if err != nil {
		return nil, errors.New("Error in signing the data: " + err.Error())
	}
	secret := solsha3.SoliditySHA3([]string{"string"}, []interface{}{hex.EncodeToString(signedData)})
	return secret, nil
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
	cmdUtils = &UtilsStruct{}
	blockManagerUtils = BlockManagerUtils{}
	voteManagerUtils = VoteManagerUtils{}
	transactionUtils = TransactionUtils{}
	flagSetUtils = FLagSetUtils{}
	abiUtils = AbiUtils{}
	timeUtils = TimeUtils{}
	stringUtils = StringUtils{}
	osUtils = OSUtils{}
	InitializeUtils()
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
