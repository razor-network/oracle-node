package cmd

import (
	"encoding/hex"
	"fmt"
	"github.com/spf13/pflag"
	"math/big"
	"os"
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
	utilsStruct := UtilsStruct{
		razorUtils:        razorUtils,
		proposeUtils:      proposeUtils,
		transactionUtils:  transactionUtils,
		blockManagerUtils: blockManagerUtils,
		voteManagerUtils:  voteManagerUtils,
		cmdUtils:          cmdUtils,
		flagSetUtils:      flagSetUtils,
	}
	utilsStruct.executeVote(cmd.Flags())
}

func (utilsStruct UtilsStruct) executeVote(flagSet *pflag.FlagSet) {
	config, err := cmdUtilsMockery.GetConfigData()
	utils.CheckError("Error in fetching config details: ", err)

	password := utils.AssignPassword(flagSet)
	isRogue, _ := flagSet.GetBool("rogue")
	rogueMode, _ := flagSet.GetStringSlice("rogueMode")
	rogueData := types.Rogue{
		IsRogue:   isRogue,
		RogueMode: rogueMode,
	}
	client := utils.ConnectToClient(config.Provider)
	address, _ := flagSet.GetString("address")
	account := types.Account{Address: address, Password: password}
	utilsStruct.vote(config, client, rogueData, account)

}

func (utilsStruct UtilsStruct) vote(config types.Configurations, client *ethclient.Client, rogueData types.Rogue, account types.Account) {
	header, err := razorUtils.GetLatestBlock(client)
	utils.CheckError("Error in getting block: ", err)

	for {
		latestHeader, err := utils.UtilsInterface.GetLatestBlockWithRetry(client)
		if err != nil {
			log.Error("Error in fetching block: ", err)
			continue
		}
		if latestHeader.Number.Cmp(header.Number) != 0 {
			header = latestHeader
			handleBlock(client, account, latestHeader.Number, config, rogueData, utilsStruct)
		}
	}
}

var (
	_committedData   []*big.Int
	lastVerification uint32
	blockConfirmed   uint32
)

func handleBlock(client *ethclient.Client, account types.Account, blockNumber *big.Int, config types.Configurations, rogueData types.Rogue, utilsStruct UtilsStruct) {
	state, err := utils.GetDelayedState(client, config.BufferPercent)
	if err != nil {
		log.Error("Error in getting state: ", err)
		return
	}
	epoch, err := utils.GetEpoch(client)
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
	actualStake, err := utils.ConvertWeiToEth(stakedAmount)
	if err != nil {
		log.Error("Error in converting stakedAmount from wei denomination: ", err)
	}
	actualBalance, err := utils.ConvertWeiToEth(ethBalance)
	if err != nil {
		log.Error("Error in converting ethBalance from wei denomination: ", err)
	}
	log.Debug("Block:", blockNumber, " Epoch:", epoch, " State:", utils.GetStateName(state), " Address:", account.Address, " Staker ID:", stakerId, " Stake:", actualStake, " Eth Balance:", actualBalance)
	if stakedAmount.Cmp(minStakeAmount) < 0 {
		log.Error("Stake is below minimum required. Cannot vote.")
		if stakedAmount.Cmp(big.NewInt(0)) == 0 {
			log.Error("Stopped voting as total stake is already withdrawn.")
		} else {
			log.Debug("Auto starting Unstake followed by Withdraw")
			AutoUnstakeAndWithdraw(client, account, stakedAmount, config, utilsStruct)
			log.Error("Stopped voting as total stake is withdrawn now")
		}
		os.Exit(0)
	}

	staker, err := utils.GetStaker(client, account.Address, stakerId)
	if err != nil {
		log.Error(err)
		return
	}
	if staker.IsSlashed {
		log.Error("Staker is slashed.... cannot continue to vote!")
		os.Exit(0)
	}
	switch state {
	case 0:
		lastCommit, err := utils.GetEpochLastCommitted(client, stakerId)
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
		data, err := cmdUtilsMockery.HandleCommitState(client, epoch, rogueData)
		if err != nil {
			log.Error("Error in getting active assets: ", err)
			break
		}
		commitTxn, err := cmdUtilsMockery.Commit(client, data, secret, account, config)
		if err != nil {
			log.Error("Error in committing data: ", err)
			break
		}
		utils.WaitForBlockCompletion(client, commitTxn.String())
		_committedData = data
		log.Debug("Saving committed data for recovery")
		fileName, err := getCommitDataFileName(account.Address, utilsStruct)
		if err != nil {
			log.Error("Error in getting file name to save committed data: ", err)
			break
		}
		err = utils.SaveCommittedDataToFile(fileName, epoch, _committedData)
		if err != nil {
			log.Errorf("Error in saving data to file %s: %t", fileName, err)
			break
		}
		log.Debug("Data saved!")
	case 1:
		lastReveal, err := utils.GetEpochLastRevealed(client, account.Address, stakerId)
		if err != nil {
			log.Error("Error in fetching last reveal: ", err)
			break
		}
		if lastReveal >= epoch {
			log.Warnf("Last reveal: %d", lastReveal)
			log.Warnf("Cannot reveal in epoch %d", epoch)
			break
		}
		if _committedData == nil {
			fileName, err := getCommitDataFileName(account.Address, utilsStruct)
			if err != nil {
				log.Error("Error in getting file name to save committed data: ", err)
				break
			}
			epochInFile, committedDataFromFile, err := utils.ReadCommittedDataFromFile(fileName)
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
		secret := calculateSecret(account, epoch)
		if secret == nil {
			break
		}
		if err := cmdUtilsMockery.HandleRevealState(client, staker, epoch); err != nil {
			log.Error(err)
			break
		}
		log.Debug("Epoch last revealed: ", lastReveal)
		revealTxn, err := cmdUtilsMockery.Reveal(client, _committedData, secret, account, account.Address, config)
		if err != nil {
			log.Error("Reveal error: ", err)
			break
		}
		if revealTxn != core.NilHash {
			utils.WaitForBlockCompletion(client, revealTxn.String())
		}
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
		proposeTxn, err := utilsStruct.Propose(client, account, config, stakerId, epoch, rogueData)
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
		lastVerification = epoch
		err := utilsStruct.HandleDispute(client, config, account, epoch)
		if err != nil {
			log.Error(err)
			break
		}
	case 4:
		if lastVerification == epoch && blockConfirmed < epoch {
			txn, err := cmdUtilsMockery.ClaimBlockReward(types.TransactionOptions{
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
	logs, err := utils.UtilsInterface.FilterLogsWithRetry(client, query)
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
	signedData, err := accounts.Sign(hash, account, razorPath, accounts.AccountUtilsInterface)
	if err != nil {
		log.Error("Error in signing the data: ", err)
		return nil
	}
	secret := solsha3.SoliditySHA3([]string{"string"}, []interface{}{hex.EncodeToString(signedData)})
	return secret
}

func getCommitDataFileName(address string, utilsStruct UtilsStruct) (string, error) {
	homeDir, err := utilsStruct.razorUtils.GetDefaultPath()
	if err != nil {
		return "", err
	}
	return homeDir + "/" + address + "_data", nil
}

func AutoUnstakeAndWithdraw(client *ethclient.Client, account types.Account, amount *big.Int, config types.Configurations, utilsStruct UtilsStruct) {
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

	_, err = Unstake(config, client,
		types.UnstakeInput{
			Address:    account.Address,
			Password:   account.Password,
			ValueInWei: amount,
			StakerId:   stakerId,
		}, utilsStruct)
	utils.CheckError("Error in Unstake: ", err)
	err = AutoWithdraw(txnArgs, stakerId, utilsStruct)
	utils.CheckError("Error in AutoWithdraw: ", err)
}

func init() {

	razorUtils = Utils{}
	proposeUtils = ProposeUtils{}
	voteManagerUtils = VoteManagerUtils{}
	transactionUtils = TransactionUtils{}
	blockManagerUtils = BlockManagerUtils{}
	transactionUtils = TransactionUtils{}
	proposeUtils = ProposeUtils{}
	cmdUtils = UtilsCmd{}
	flagSetUtils = FlagSetUtils{}
	utils.Options = &utils.OptionsStruct{}
	utils.UtilsInterface = &utils.UtilsStruct{}
	cmdUtilsMockery = &UtilsStructMockery{}
	blockManagerUtilsMockery = BlockManagerUtilsMockery{}
	voteManagerUtilsMockery = VoteManagerUtilsMockery{}
	transactionUtilsMockery = TransactionUtilsMockery{}

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
