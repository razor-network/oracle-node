//Package cmd provides all functions related to command line
package cmd

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"razor/cache"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/rpc"
	"razor/utils"
	"time"

	Types "github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/spf13/pflag"

	"github.com/ethereum/go-ethereum/common"
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
	config, rpcParameters, account, err := InitializeCommandDependencies(flagSet)
	utils.CheckError("Error in initialising command dependencies: ", err)

	err = ValidateBufferPercentLimit(rpcParameters, config.BufferPercent)
	utils.CheckError("Error in validating buffer percent: ", err)

	isRogue, err := flagSetUtils.GetBoolRogue(flagSet)
	utils.CheckError("Error in getting rogue status: ", err)
	log.Debug("ExecuteVote: IsRogue: ", isRogue)

	rogueMode, err := flagSetUtils.GetStringSliceRogueMode(flagSet)
	utils.CheckError("Error in getting rogue modes: ", err)
	log.Debug("ExecuteVote: RogueMode: ", rogueMode)

	backupNodeActionsToIgnore, err := flagSetUtils.GetStringSliceBackupNode(flagSet)
	utils.CheckError("Error in getting backupNode actions to ignore: ", err)
	log.Debug("ExecuteVote: Backup node actions to ignore: ", backupNodeActionsToIgnore)

	rogueData := types.Rogue{
		IsRogue:   isRogue,
		RogueMode: rogueMode,
	}

	if rogueData.IsRogue {
		log.Warn("YOU ARE RUNNING VOTE IN ROGUE MODE, THIS CAN INCUR PENALTIES!")
	}

	httpClient := &http.Client{
		Timeout: time.Duration(config.HTTPTimeout) * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        core.HTTPClientMaxIdleConns,
			MaxIdleConnsPerHost: core.HTTPClientMaxIdleConnsPerHost,
		},
	}

	stakerId, err := razorUtils.GetStakerId(rpcParameters, account.Address)
	utils.CheckError("Error in getting staker id: ", err)

	if stakerId == 0 {
		log.Fatal("Staker doesn't exist")
	}

	cmdUtils.HandleExit()

	jobsCache, collectionsCache, initCacheBlockNumber, err := cmdUtils.InitJobAndCollectionCache(rpcParameters)
	utils.CheckError("Error in initializing asset cache: ", err)

	commitParams := &types.CommitParams{
		LocalCache:                cache.NewLocalCache(), // Creating a local cache which will store API results
		JobsCache:                 jobsCache,
		CollectionsCache:          collectionsCache,
		HttpClient:                httpClient,
		FromBlockToCheckForEvents: initCacheBlockNumber,
	}

	log.Debugf("Calling Vote() with arguments rogueData = %+v, account address = %s, backup node actions to ignore = %s", rogueData, account.Address, backupNodeActionsToIgnore)
	if err := cmdUtils.Vote(rpcParameters, config, account, stakerId, commitParams, rogueData, backupNodeActionsToIgnore); err != nil {
		log.Errorf("%v\n", err)
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
func (*UtilsStruct) Vote(rpcParameters rpc.RPCParameters, config types.Configurations, account types.Account, stakerId uint32, commitParams *types.CommitParams, rogueData types.Rogue, backupNodeActionsToIgnore []string) error {
	header, err := clientUtils.GetLatestBlockWithRetry(rpcParameters)
	utils.CheckError("Error in getting block: ", err)
	for {
		select {
		case <-rpcParameters.Ctx.Done():
			return nil
		default:
			log.Debugf("Vote: Header value: %d", header.Number)
			latestHeader, err := clientUtils.GetLatestBlockWithRetry(rpcParameters)
			if err != nil {
				log.Error("Error in fetching block: ", err)
				continue
			}
			log.Debugf("Vote: Latest header value: %d", latestHeader.Number)
			if latestHeader.Number.Cmp(header.Number) != 0 {
				header = latestHeader
				cmdUtils.HandleBlock(rpcParameters, account, stakerId, latestHeader, config, commitParams, rogueData, backupNodeActionsToIgnore)
			}
			time.Sleep(time.Second * time.Duration(core.BlockNumberInterval))
		}
	}
}

var (
	globalCommitDataStruct types.CommitFileData
	lastVerification       uint32
	blockConfirmed         uint32
	disputeData            types.DisputeFileData
	lastRPCRefreshEpoch    uint32
)

//This function handles the block
func (*UtilsStruct) HandleBlock(rpcParameters rpc.RPCParameters, account types.Account, stakerId uint32, latestHeader *Types.Header, config types.Configurations, commitParams *types.CommitParams, rogueData types.Rogue, backupNodeActionsToIgnore []string) {
	stateBuffer, err := razorUtils.GetStateBuffer(rpcParameters)
	if err != nil {
		log.Error("Error in getting state buffer: ", err)
		return
	}
	remainingTimeOfTheCurrentState, err := razorUtils.GetRemainingTimeOfCurrentState(latestHeader, stateBuffer, config.BufferPercent)
	if err != nil {
		log.Error("Error in getting remaining time of the current state: ", err)
		return
	}
	if remainingTimeOfTheCurrentState <= 0 {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(remainingTimeOfTheCurrentState)*time.Second)
	defer cancel()

	// Replacing context with the context which timeouts after remainingTimeOfTheCurrentState seconds
	rpcParameters.Ctx = ctx

	state, err := razorUtils.GetBufferedState(latestHeader, stateBuffer, config.BufferPercent)
	if err != nil {
		log.Error("Error in getting state: ", err)
		return
	}
	epoch, err := razorUtils.GetEpoch(rpcParameters)
	if err != nil {
		log.Error("Error in getting epoch: ", err)
		return
	}

	staker, err := razorUtils.GetStaker(rpcParameters, stakerId)
	if err != nil {
		log.Error(err)
		return
	}
	stakedAmount := staker.Stake

	ethBalance, err := clientUtils.BalanceAtWithRetry(rpcParameters, common.HexToAddress(account.Address))
	if err != nil {
		log.Errorf("Error in fetching balance of the account: %s\n%v", account.Address, err)
		return
	}

	// Warning the staker if sFUEL balance is less than 0.001 sFUEL
	if ethBalance.Cmp(big.NewInt(1e15)) == -1 {
		log.Warn("sFUEL balance is lower than 0.001 sFUEL, kindly add more sFUEL to be safe for executing transactions successfully")
	}

	actualStake, err := utils.ConvertWeiToEth(stakedAmount)
	if err != nil {
		log.Error("Error in converting stakedAmount from wei denomination: ", err)
		return
	}
	actualBalance, err := utils.ConvertWeiToEth(ethBalance)
	if err != nil {
		log.Error("Error in converting ethBalance from wei denomination: ", err)
		return
	}

	sRZRBalance, err := razorUtils.GetStakerSRZRBalance(rpcParameters, staker)
	if err != nil {
		log.Error("Error in getting sRZR balance for staker: ", err)
		return
	}

	var sRZRInEth *big.Float
	if sRZRBalance.Cmp(big.NewInt(0)) == 0 {
		sRZRInEth = big.NewFloat(0)
	} else {
		sRZRInEth, err = utils.ConvertWeiToEth(sRZRBalance)
		if err != nil {
			log.Error(err)
			return
		}
	}

	log.Infof("State: %s Staker ID: %d Stake: %f sRZR Balance: %f sFUEL Balance: %f", utils.GetStateName(state), stakerId, actualStake, sRZRInEth, actualBalance)

	if staker.IsSlashed {
		log.Error("Staker is slashed.... cannot continue to vote!")
		osUtils.Exit(0)
	}

	switch state {
	case 0:
		log.Debugf("Starting commit...")
		err := cmdUtils.InitiateCommit(rpcParameters, config, account, epoch, stakerId, latestHeader, commitParams, stateBuffer, rogueData)
		if err != nil {
			log.Error(err)
			break
		}
	case 1:
		log.Debugf("Starting reveal...")
		err := cmdUtils.InitiateReveal(rpcParameters, config, account, epoch, staker, latestHeader, stateBuffer, rogueData)
		if err != nil {
			log.Error(err)
			break
		}
	case 2:
		log.Debugf("Starting propose...")
		err := cmdUtils.InitiatePropose(rpcParameters, config, account, epoch, staker, latestHeader, stateBuffer, rogueData)
		if err != nil {
			log.Error(err)
			break
		}
	case 3:
		log.Debugf("Last verification: %d", lastVerification)
		if lastVerification >= epoch {
			log.Debugf("Last verification (%d) is greater or equal to current epoch (%d)", lastVerification, epoch)
			log.Debugf("Won't dispute now")
			break
		}

		err := cmdUtils.HandleDispute(rpcParameters, config, account, epoch, latestHeader.Number, rogueData, backupNodeActionsToIgnore)
		if err != nil {
			log.Error(err)
			break
		}

		lastVerification = epoch

		if razorUtils.IsFlagPassed("autoClaimBounty") {
			log.Debugf("Automatically claiming bounty")
			err = cmdUtils.HandleClaimBounty(rpcParameters, config, account)
			if err != nil {
				log.Error(err)
				break
			}
		}

	case 4:
		log.Debugf("Last verification: %d", lastVerification)
		log.Debugf("Block confirmed: %d", blockConfirmed)

		if blockConfirmed >= epoch {
			log.Debug("Block is already confirmed for this epoch!")
			break
		}

		confirmedBlock, err := razorUtils.GetConfirmedBlocks(rpcParameters, epoch)
		if err != nil {
			log.Error(err)
			break
		}

		if confirmedBlock.ProposerId != 0 {
			log.Infof("Block is already confirmed, setting blockConfirmed (%d) to current epoch (%d)", blockConfirmed, epoch)
			blockConfirmed = epoch
			break
		}
		if (lastVerification == epoch || lastVerification == 0) && blockConfirmed < epoch {
			txn, err := cmdUtils.ClaimBlockReward(rpcParameters, types.TransactionOptions{
				ChainId:         core.ChainId,
				Config:          config,
				ContractAddress: core.BlockManagerAddress,
				MethodName:      "claimBlockReward",
				ABI:             bindings.BlockManagerMetaData.ABI,
				Account:         account,
			})

			if err != nil {
				log.Error("ClaimBlockReward error: ", err)
				break
			}
			if txn != core.NilHash {
				log.Info("Confirm Transaction Hash: ", txn)
			}

			if lastRPCRefreshEpoch < epoch {
				err = rpcParameters.RPCManager.RefreshEndpoints()
				if err != nil {
					log.Error("Error in refreshing RPC endpoints: ", err)
					break
				}
				bestEndpointURL, err := rpcParameters.RPCManager.GetBestEndpointURL()
				if err != nil {
					log.Error("Error in getting best RPC endpoint URL after refreshing: ", err)
					break
				}
				log.Info("Current best RPC endpoint URL: ", bestEndpointURL)

				lastRPCRefreshEpoch = epoch
			}
		}
	case -1:
		if config.WaitTime >= core.BufferStateSleepTime {
			timeUtils.Sleep(time.Second * time.Duration(core.BufferStateSleepTime))
			return
		}
	}
	razorUtils.WaitTillNextNSecs(config.WaitTime)
	fmt.Println()
}

//This function initiates the commit
func (*UtilsStruct) InitiateCommit(rpcParameters rpc.RPCParameters, config types.Configurations, account types.Account, epoch uint32, stakerId uint32, latestHeader *Types.Header, commitParams *types.CommitParams, stateBuffer uint64, rogueData types.Rogue) error {
	lastCommit, err := razorUtils.GetEpochLastCommitted(rpcParameters, stakerId)
	if err != nil {
		return errors.New("Error in fetching last commit: " + err.Error())
	}
	log.Debug("InitiateCommit: Epoch last committed: ", lastCommit)

	if lastCommit >= epoch {
		// Clearing up the cache storing API results as the staker has already committed successfully
		commitParams.LocalCache.ClearAll()
		return nil
	}

	err = CheckForJobAndCollectionEvents(rpcParameters, commitParams)
	if err != nil {
		log.Error("Error in checking for asset events: ", err)
		return err
	}

	staker, err := razorUtils.GetStaker(rpcParameters, stakerId)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Debug("InitiateCommit: Staker:", staker)
	stakedAmount := staker.Stake
	minStakeAmount, err := razorUtils.GetMinStakeAmount(rpcParameters)
	if err != nil {
		log.Error("Error in getting minimum stake amount: ", err)
		return err
	}
	log.Debug("InitiateCommit: Minimum stake amount: ", minStakeAmount)
	if stakedAmount.Cmp(minStakeAmount) < 0 {
		log.Error("Stake is below minimum required. Kindly add stake to continue voting.")
		return nil
	}
	razorPath, err := pathUtils.GetDefaultPath()
	if err != nil {
		return err
	}
	log.Debugf("InitiateCommit: .razor directory path: %s", razorPath)
	keystorePath := filepath.Join(razorPath, "keystore_files")
	log.Debugf("InitiateCommit: Keystore file path: %s", keystorePath)
	log.Debugf("InitiateCommit: Calling CalculateSeed() with arguments keystorePath = %s, epoch = %d", keystorePath, epoch)
	seed, err := CalculateSeed(rpcParameters, account, keystorePath, epoch)
	if err != nil {
		return errors.New("Error in getting seed: " + err.Error())
	}

	log.Debugf("InitiateCommit: Calling HandleCommitState with arguments epoch = %d, seed = %v, rogueData = %+v", epoch, seed, rogueData)
	commitData, err := cmdUtils.HandleCommitState(rpcParameters, epoch, seed, commitParams, rogueData)
	if err != nil {
		return errors.New("Error in getting active assets: " + err.Error())
	}
	log.Debug("InitiateCommit: Commit Data: ", commitData)

	commitmentToSend, err := CalculateCommitment(seed, commitData.Leaves)
	if err != nil {
		log.Error("Error in getting commitment: ", err)
		return err
	}

	commitTxn, err := cmdUtils.Commit(rpcParameters, config, account, epoch, latestHeader, stateBuffer, commitmentToSend)
	if err != nil {
		return errors.New("Error in committing data: " + err.Error())
	}
	log.Info("InitiateCommit: Commit Transaction Hash: ", commitTxn)
	if commitTxn != core.NilHash {
		log.Debug("Saving committed data for recovery...")
		fileName, err := pathUtils.GetCommitDataFileName(account.Address)
		if err != nil {
			return errors.New("Error in getting file name to save committed data: " + err.Error())
		}
		log.Debug("InitiateCommit: Commit data file path: ", fileName)

		err = fileUtils.SaveDataToCommitJsonFile(fileName, epoch, commitData, commitmentToSend)
		if err != nil {
			return errors.New("Error in saving data to file" + fileName + ": " + err.Error())
		}
		log.Debug("Data saved!")

		log.Debug("Updating GlobalCommitDataStruct with latest commitData and epoch...")
		updateGlobalCommitDataStruct(commitData, commitmentToSend, epoch)
		log.Debugf("InitiateCommit: Global commit data struct: %+v", globalCommitDataStruct)
	}
	return nil
}

//This function initiates the reveal
func (*UtilsStruct) InitiateReveal(rpcParameters rpc.RPCParameters, config types.Configurations, account types.Account, epoch uint32, staker bindings.StructsStaker, latestHeader *Types.Header, stateBuffer uint64, rogueData types.Rogue) error {
	stakedAmount := staker.Stake
	log.Debug("InitiateReveal: Staked Amount: ", stakedAmount)
	minStakeAmount, err := razorUtils.GetMinStakeAmount(rpcParameters)
	if err != nil {
		log.Error("Error in getting minimum stake amount: ", err)
		return err
	}
	log.Debug("InitiateReveal: Minimum Stake Amount: ", minStakeAmount)
	if stakedAmount.Cmp(minStakeAmount) < 0 {
		log.Error("Stake is below minimum required. Kindly add stake to continue voting.")
		return nil
	}
	lastReveal, err := razorUtils.GetEpochLastRevealed(rpcParameters, staker.Id)
	if err != nil {
		return errors.New("Error in fetching last reveal: " + err.Error())
	}
	log.Debug("InitiateReveal: Last reveal was at epoch ", lastReveal)

	if lastReveal >= epoch {
		log.Debugf("Since last reveal was at epoch: %d, won't reveal again in epoch: %d", lastReveal, epoch)
		return nil
	}

	log.Debugf("InitiateReveal: Calling CheckForLastCommitted with arguments staker = %+v, epoch = %d", staker, epoch)
	if err := cmdUtils.CheckForLastCommitted(rpcParameters, staker, epoch); err != nil {
		log.Error(err)
		return err
	}

	razorPath, err := pathUtils.GetDefaultPath()
	if err != nil {
		return err
	}
	log.Debugf("InitiateReveal: .razor directory path: %s", razorPath)
	keystorePath := filepath.Join(razorPath, "keystore_files")
	log.Debugf("InitiateReveal: Keystore file path: %s", keystorePath)

	// Consolidated commitment verification for commit data being fetched from memory or file
	commitData, err := GetCommittedDataForEpoch(rpcParameters, account, epoch, rogueData)
	if err != nil {
		return err
	}

	log.Debugf("InitiateReveal: Calling CalculateSecret() with argument epoch = %d, keystorePath = %s, chainId = %s", epoch, keystorePath, core.ChainId)
	signature, _, err := cmdUtils.CalculateSecret(account, epoch, keystorePath, core.ChainId)
	if err != nil {
		return err
	}
	log.Debug("InitiateReveal: Signature: ", signature)
	log.Debug("InitiateReveal: Assigned Collections: ", commitData.AssignedCollections)
	log.Debug("InitiateReveal: SeqAllottedCollections: ", commitData.SeqAllottedCollections)
	log.Debug("InitiateReveal: Leaves: ", commitData.Leaves)

	commitDataToSend := types.CommitData{
		Leaves:                 commitData.Leaves,
		AssignedCollections:    commitData.AssignedCollections,
		SeqAllottedCollections: commitData.SeqAllottedCollections,
	}
	log.Debugf("InitiateReveal: Calling Reveal() with arguments epoch = %d, commitDataToSend = %+v, signature = %v", epoch, commitDataToSend, signature)
	revealTxn, err := cmdUtils.Reveal(rpcParameters, config, account, epoch, latestHeader, stateBuffer, commitDataToSend, signature)
	if err != nil {
		return errors.New("Reveal error: " + err.Error())
	}
	log.Info("InitiateReveal: Reveal Transaction Hash: ", revealTxn)
	return nil
}

//This function initiates the propose
func (*UtilsStruct) InitiatePropose(rpcParameters rpc.RPCParameters, config types.Configurations, account types.Account, epoch uint32, staker bindings.StructsStaker, latestHeader *Types.Header, stateBuffer uint64, rogueData types.Rogue) error {
	stakedAmount := staker.Stake
	log.Debug("InitiatePropose: Staked Amount: ", stakedAmount)
	minStakeAmount, err := razorUtils.GetMinStakeAmount(rpcParameters)
	if err != nil {
		log.Error("Error in getting minimum stake amount: ", err)
		return err
	}
	log.Debug("InitiatePropose: Minimum Stake Amount: ", minStakeAmount)
	if stakedAmount.Cmp(minStakeAmount) < 0 {
		log.Error("Stake is below minimum required. Kindly add stake to continue voting.")
		return nil
	}
	lastProposal, err := razorUtils.GetEpochLastProposed(rpcParameters, staker.Id)
	if err != nil {
		return errors.New("Error in fetching last proposal: " + err.Error())
	}
	log.Debug("InitiatePropose: Last propose was in epoch ", lastProposal)
	if lastProposal >= epoch {
		log.Debugf("Since last propose was at epoch: %d, won't propose again in epoch: %d", epoch, lastProposal)
		return nil
	}
	lastReveal, err := razorUtils.GetEpochLastRevealed(rpcParameters, staker.Id)
	if err != nil {
		return errors.New("Error in fetching last reveal: " + err.Error())
	}
	log.Debug("InitiatePropose: Last reveal was in epoch ", lastReveal)
	if lastReveal < epoch {
		log.Debugf("Cannot propose in epoch %d because last reveal was in epoch %d", epoch, lastReveal)
		return nil
	}

	log.Debugf("InitiatePropose: Calling Propose() with arguments staker = %+v, epoch = %d, blockNumber = %s, rogueData = %+v", staker, epoch, latestHeader.Number, rogueData)
	err = cmdUtils.Propose(rpcParameters, config, account, staker, epoch, latestHeader, stateBuffer, rogueData)
	if err != nil {
		return errors.New("Propose error: " + err.Error())
	}
	return nil
}

//This function calculates the secret
func (*UtilsStruct) CalculateSecret(account types.Account, epoch uint32, keystorePath string, chainId *big.Int) ([]byte, []byte, error) {
	if chainId == nil {
		return nil, nil, errors.New("chainId is nil")
	}
	hash := solsha3.SoliditySHA3([]string{"address", "uint32", "uint256", "string"}, []interface{}{common.HexToAddress(account.Address), epoch, chainId, "razororacle"})
	log.Debug("CalculateSecret: Hash: ", hash)
	ethHash := utils.SignHash(hash)
	log.Debug("Hash generated for secret")
	log.Debug("CalculateSecret: Ethereum signed hash: ", ethHash)
	signedData, err := account.AccountManager.SignData(ethHash, account.Address, account.Password)
	if err != nil {
		return nil, nil, errors.New("Error in signing the data: " + err.Error())
	}
	log.Debug("CalculateSecret: SignedData: ", signedData)
	log.Debugf("Checking whether recovered address from Hash: %v and Signed data: %v is same as given address...", hash, signedData)
	recoveredAddress, err := utils.EcRecover(hash, signedData)
	if err != nil {
		return nil, nil, errors.New("Error in verifying: " + err.Error())
	}
	log.Debug("CalculateSecret: Recovered Address: ", recoveredAddress)
	if recoveredAddress != common.HexToAddress(account.Address) {
		return nil, nil, errors.New("invalid verification")
	}
	log.Debug("Address verified, generating secret....")
	if signedData[64] == 0 || signedData[64] == 1 {
		signedData[64] += 27
	}

	secret := crypto.Keccak256(signedData)
	log.Debug("Secret generated.")
	return signedData, secret, nil
}

func updateGlobalCommitDataStruct(commitData types.CommitData, commitment [32]byte, epoch uint32) types.CommitFileData {
	globalCommitDataStruct.Leaves = commitData.Leaves
	globalCommitDataStruct.AssignedCollections = commitData.AssignedCollections
	globalCommitDataStruct.SeqAllottedCollections = commitData.SeqAllottedCollections
	globalCommitDataStruct.Epoch = epoch
	globalCommitDataStruct.Commitment = commitment
	return globalCommitDataStruct
}

func init() {
	rootCmd.AddCommand(voteCmd)

	var (
		Address         string
		Rogue           bool
		RogueMode       []string
		Password        string
		AutoClaimBounty bool
		BackupNode      []string
	)

	voteCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the staker")
	voteCmd.Flags().BoolVarP(&Rogue, "rogue", "r", false, "enable rogue mode to report wrong values")
	voteCmd.Flags().StringSliceVarP(&RogueMode, "rogueMode", "", []string{}, "type of rogue mode")
	voteCmd.Flags().StringVarP(&Password, "password", "", "", "password path of the staker to protect the keystore")
	voteCmd.Flags().BoolVarP(&AutoClaimBounty, "autoClaimBounty", "", false, "auto claim bounty")
	voteCmd.Flags().StringSliceVarP(&BackupNode, "backupNode", "", []string{}, "actions that backup node will ignore")

	addrErr := voteCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
}
