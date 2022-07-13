//Package cmd provides all functions related to command line
package cmd

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"os"
	"os/signal"
	"path"
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

//This function initialises the ExecuteVote function
func initializeVote(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteVote(cmd.Flags())
}

//This function sets the flag appropriately and executes the Vote function
func (*UtilsStruct) ExecuteVote(flagSet *pflag.FlagSet) {
	razorUtils.AssignLogFile(flagSet)
	address, err := flagSetUtils.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)

	logger.Address = address

	config, err := cmdUtils.GetConfigData()
	utils.CheckError("Error in fetching config details: ", err)

	password := razorUtils.AssignPassword()
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

//This function handles the exit and listens for CTRL+C
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
			log.Info("Press CTRL+C again to terminate.")
		case <-ctx.Done():
		}
		<-signalChan // second signal, hard exit
		os.Exit(2)
	}()
}

//This function handles all the states of voting
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
	disputeData      types.DisputeFileData
)

//This function handles the block
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

	sRZRBalance, err := razorUtils.GetStakerSRZRBalance(client, staker)
	if err != nil {
		log.Error("Error in getting sRZR balance for staker: ", err)
		return
	}

	var sRZRInEth *big.Float
	if sRZRBalance.Cmp(big.NewInt(0)) == 0 {
		sRZRInEth = big.NewFloat(0)
	} else {
		sRZRInEth, err = razorUtils.ConvertWeiToEth(sRZRBalance)
		if err != nil {
			log.Error(err)
			return
		}
	}

	log.Infof("Block: %d Epoch: %d State: %s Staker ID: %d Stake: %f sRZR Balance: %f Eth Balance: %f", blockNumber, epoch, utils.UtilsInterface.GetStateName(state), stakerId, actualStake, sRZRInEth, actualBalance)

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
		err := cmdUtils.InitiatePropose(client, config, account, epoch, staker, blockNumber, rogueData)
		if err != nil {
			log.Error(err)
			break
		}
	case 3:
		if lastVerification >= epoch {
			break
		}

		err := cmdUtils.HandleDispute(client, config, account, epoch, blockNumber, rogueData)
		if err != nil {
			log.Error(err)
			break
		}

		lastVerification = epoch

		if utilsInterface.IsFlagPassed("autoClaimBounty") {
			err = cmdUtils.HandleClaimBounty(client, config, account)
			if err != nil {
				log.Error(err)
				break
			}
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
				waitForBlockCompletionErr := razorUtils.WaitForBlockCompletion(client, txn.Hex())
				if waitForBlockCompletionErr != nil {
					log.Error("Error in WaitForBlockCompletion for claimBlockReward: ", err)
					break
				}
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

//This function initiates the commit
func (*UtilsStruct) InitiateCommit(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, stakerId uint32, rogueData types.Rogue) error {
	staker, err := razorUtils.GetStaker(client, stakerId)
	if err != nil {
		log.Error(err)
		return err
	}
	stakedAmount := staker.Stake
	minStakeAmount, err := utils.UtilsInterface.GetMinStakeAmount(client)
	if err != nil {
		log.Error("Error in getting minimum stake amount: ", err)
		return err
	}
	if stakedAmount.Cmp(minStakeAmount) < 0 {
		log.Error("Stake is below minimum required. Kindly add stake to continue voting.")
		return nil
	}
	lastCommit, err := razorUtils.GetEpochLastCommitted(client, stakerId)
	if err != nil {
		return errors.New("Error in fetching last commit: " + err.Error())
	}
	if lastCommit >= epoch {
		log.Debugf("Cannot commit in epoch %d because last committed epoch is %d", epoch, lastCommit)
		return nil
	}
	razorPath, err := razorUtils.GetDefaultPath()
	if err != nil {
		return err
	}
	keystorePath := path.Join(razorPath, "keystore_files")

	_, secret, err := cmdUtils.CalculateSecret(account, epoch, keystorePath, core.ChainId)
	if err != nil {
		return err
	}

	salt, err := cmdUtils.GetSalt(client, epoch)
	if err != nil {
		return err
	}

	seed := solsha3.SoliditySHA3([]string{"bytes32", "bytes32"}, []interface{}{"0x" + hex.EncodeToString(salt[:]), "0x" + hex.EncodeToString(secret)})

	commitData, err := cmdUtils.HandleCommitState(client, epoch, seed, rogueData)
	if err != nil {
		return errors.New("Error in getting active assets: " + err.Error())
	}

	_commitData = commitData

	merkleTree := utils.MerkleInterface.CreateMerkle(commitData.Leaves)
	commitTxn, err := cmdUtils.Commit(client, config, account, epoch, seed, utils.MerkleInterface.GetMerkleRoot(merkleTree))
	if err != nil {
		return errors.New("Error in committing data: " + err.Error())
	}
	if commitTxn != core.NilHash {
		waitForBlockCompletionErr := razorUtils.WaitForBlockCompletion(client, commitTxn.String())
		if waitForBlockCompletionErr != nil {
			log.Error("Error in WaitForBlockCompletion for commit: ", err)
			return errors.New("error in sending commit transaction")
		}
	}

	log.Debug("Saving committed data for recovery")
	fileName, err := razorUtils.GetCommitDataFileName(account.Address)
	if err != nil {
		return errors.New("Error in getting file name to save committed data: " + err.Error())
	}

	err = razorUtils.SaveDataToCommitJsonFile(fileName, epoch, commitData)
	if err != nil {
		return errors.New("Error in saving data to file" + fileName + ": " + err.Error())
	}
	log.Debug("Data saved!")
	return nil
}

//This function initiates the reveal
func (*UtilsStruct) InitiateReveal(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, staker bindings.StructsStaker, rogueData types.Rogue) error {
	stakedAmount := staker.Stake
	minStakeAmount, err := utils.UtilsInterface.GetMinStakeAmount(client)
	if err != nil {
		log.Error("Error in getting minimum stake amount: ", err)
		return err
	}
	if stakedAmount.Cmp(minStakeAmount) < 0 {
		log.Error("Stake is below minimum required. Kindly add stake to continue voting.")
		return nil
	}
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
		return err
	}
	log.Debug("Epoch last revealed: ", lastReveal)

	if _commitData.AssignedCollections == nil && _commitData.SeqAllottedCollections == nil && _commitData.Leaves == nil {
		fileName, err := razorUtils.GetCommitDataFileName(account.Address)
		if err != nil {
			log.Error("Error in getting file name to save committed data: ", err)
			return err
		}
		committedDataFromFile, err := razorUtils.ReadFromCommitJsonFile(fileName)
		if err != nil {
			log.Errorf("Error in getting committed data from file %s: %t", fileName, err)
			return err
		}
		if committedDataFromFile.Epoch != epoch {
			log.Errorf("File %s doesn't contain latest committed data: %t", fileName, err)
			return errors.New("commit data file doesn't contain latest committed data")
		}
		_commitData.AssignedCollections = committedDataFromFile.AssignedCollections
		_commitData.SeqAllottedCollections = committedDataFromFile.SeqAllottedCollections
		_commitData.Leaves = committedDataFromFile.Leaves
	}
	if rogueData.IsRogue && utils.Contains(rogueData.RogueMode, "reveal") {
		var rogueCommittedData []*big.Int
		for i := 0; i < len(_commitData.Leaves); i++ {
			rogueCommittedData = append(rogueCommittedData, razorUtils.GetRogueRandomValue(10000000))
		}
		_commitData.Leaves = rogueCommittedData
	}

	razorPath, err := razorUtils.GetDefaultPath()
	if err != nil {
		return err
	}
	keystorePath := path.Join(razorPath, "keystore_files")

	signature, _, err := cmdUtils.CalculateSecret(account, epoch, keystorePath, core.ChainId)
	if err != nil {
		return err
	}
	revealTxn, err := cmdUtils.Reveal(client, config, account, epoch, _commitData, signature)
	if err != nil {
		return errors.New("Reveal error: " + err.Error())
	}
	if revealTxn != core.NilHash {
		waitForBlockCompletionErr := razorUtils.WaitForBlockCompletion(client, revealTxn.String())
		if waitForBlockCompletionErr != nil {
			log.Error("Error in WaitForBlockCompletionErr for reveal: ", err)
			return err
		}
	}
	return nil
}

//This function initiates the propose
func (*UtilsStruct) InitiatePropose(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, staker bindings.StructsStaker, blockNumber *big.Int, rogueData types.Rogue) error {
	stakedAmount := staker.Stake
	minStakeAmount, err := utils.UtilsInterface.GetMinStakeAmount(client)
	if err != nil {
		log.Error("Error in getting minimum stake amount: ", err)
		return err
	}
	if stakedAmount.Cmp(minStakeAmount) < 0 {
		log.Error("Stake is below minimum required. Kindly add stake to continue voting.")
		return nil
	}
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
		waitForBlockCompletionErr := razorUtils.WaitForBlockCompletion(client, proposeTxn.String())
		if waitForBlockCompletionErr != nil {
			log.Error("Error in WaitForBlockCompletionErr for propose: ", err)
			return err
		}
	}
	return nil
}

//This function returns the last proposed epoch
func (*UtilsStruct) GetLastProposedEpoch(client *ethclient.Client, blockNumber *big.Int, stakerId uint32) (uint32, error) {
	fromBlock, err := utils.UtilsInterface.CalculateBlockNumberAtEpochBeginning(client, core.EpochLength, blockNumber)
	if err != nil {
		return 0, errors.New("Not able to Fetch Block: " + err.Error())
	}
	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   blockNumber,
		Addresses: []common.Address{
			common.HexToAddress(core.BlockManagerAddress),
		},
	}
	logs, err := utils.UtilsInterface.FilterLogsWithRetry(client, query)
	if err != nil {
		return 0, err
	}
	contractAbi, err := utils.ABIInterface.Parse(strings.NewReader(bindings.BlockManagerABI))
	if err != nil {
		return 0, err
	}
	epochLastProposed := uint32(0)

	bufferPercent, err := cmdUtils.GetBufferPercent()
	if err != nil {
		return 0, err
	}

	stateRemainingTime, err := utilsInterface.GetRemainingTimeOfCurrentState(client, bufferPercent)
	if err != nil {
		return 0, err
	}
	stateTimeout := time.NewTimer(time.Second * time.Duration(stateRemainingTime))

loop:
	for _, vLog := range logs {
		select {
		case <-stateTimeout.C:
			log.Error("State timeout!")
			err = errors.New("propose state timeout")
			break loop
		default:
			data, unpackErr := abiUtils.Unpack(contractAbi, "Proposed", vLog.Data)
			if unpackErr != nil {
				log.Error(unpackErr)
				continue
			}
			topics := vLog.Topics
			// topics[1] gives staker id in data type common.Hash
			// Converting uint32 staker id to common.Hash to compare with staker id from topics
			stakerIdInHash := common.BigToHash(big.NewInt(int64(stakerId)))
			if stakerIdInHash == topics[1] && len(topics) > 1 {
				epochLastProposed = data[0].(uint32)
			}
		}
	}
	if err != nil {
		return 0, err
	}
	return epochLastProposed, nil
}

//This function calculates the secret
func (*UtilsStruct) CalculateSecret(account types.Account, epoch uint32, keystorePath string, chainId *big.Int) ([]byte, []byte, error) {
	if chainId == nil {
		return nil, nil, errors.New("chainId is nil")
	}
	hash := solsha3.SoliditySHA3([]string{"address", "uint32", "uint256", "string"}, []interface{}{common.HexToAddress(account.Address), epoch, chainId, "razororacle"})
	ethHash := utils.SignHash(hash)

	signedData, err := accounts.AccountUtilsInterface.SignData(ethHash, account, keystorePath)
	if err != nil {
		return nil, nil, errors.New("Error in signing the data: " + err.Error())
	}
	recoveredAddress, err := utils.EcRecover(hash, signedData)
	if err != nil {
		return nil, nil, errors.New("Error in verifying: " + err.Error())
	}
	if recoveredAddress != common.HexToAddress(account.Address) {
		return nil, nil, errors.New("invalid verification")
	}
	if signedData[64] == 0 || signedData[64] == 1 {
		signedData[64] += 27
	}

	secret := crypto.Keccak256(signedData)
	return signedData, secret, nil
}

func init() {
	rootCmd.AddCommand(voteCmd)

	var (
		Address         string
		Rogue           bool
		RogueMode       []string
		AutoClaimBounty bool
	)

	voteCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the staker")
	voteCmd.Flags().BoolVarP(&Rogue, "rogue", "r", false, "enable rogue mode to report wrong values")
	voteCmd.Flags().StringSliceVarP(&RogueMode, "rogueMode", "", []string{}, "type of rogue mode")
	voteCmd.Flags().BoolVarP(&AutoClaimBounty, "autoClaimBounty", "", false, "auto claim bounty")

	addrErr := voteCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
}
