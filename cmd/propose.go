package cmd

import (
	"encoding/hex"
	"math"
	"math/big"
	"math/rand"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"modernc.org/sortutil"
)

var proposeUtils proposeUtilsInterface

func Propose(client *ethclient.Client, account types.Account, config types.Configurations, stakerId uint32, epoch uint32, rogueMode bool, razorUtils utilsInterface, proposeUtils proposeUtilsInterface, blockManagerUtils blockManagerInterface, transactionUtils transactionInterface) (common.Hash, error) {
	if state, err := razorUtils.GetDelayedState(client, config.BufferPercent); err != nil || state != 2 {
		log.Error("Not propose state")
		return core.NilHash, err
	}
	staker, err := razorUtils.GetStaker(client, account.Address, stakerId)
	if err != nil {
		log.Error("Error in fetching staker: ", err)
		return core.NilHash, err
	}
	numStakers, err := razorUtils.GetNumberOfStakers(client, account.Address)
	if err != nil {
		log.Error("Error in fetching number of stakers: ", err)
		return core.NilHash, err
	}
	log.Debug("Stake: ", staker.Stake)

	biggestInfluence, biggestInfluenceId, err := proposeUtils.getBiggestInfluenceAndId(client, account.Address, epoch, razorUtils)
	if err != nil {
		log.Error("Error in calculating biggest staker: ", err)
		return core.NilHash, err
	}

	randaoHash, err := razorUtils.GetRandaoHash(client, account.Address)
	if err != nil {
		log.Error("Error in fetching random hash: ", err)
		return core.NilHash, err
	}
	log.Debug("Biggest Influence Id: ", biggestInfluenceId)
	log.Debugf("Biggest influence: %s, Stake: %s, Staker Id: %d, Number of Stakers: %d, Randao Hash: %s", biggestInfluence, staker.Stake, stakerId, numStakers, hex.EncodeToString(randaoHash[:]))

	iteration := proposeUtils.getIteration(client, account.Address, types.ElectedProposer{
		Stake:            staker.Stake,
		StakerId:         stakerId,
		BiggestInfluence: biggestInfluence,
		NumberOfStakers:  numStakers,
		RandaoHash:       randaoHash,
		Epoch:            epoch,
	}, proposeUtils)

	log.Debug("Iteration: ", iteration)

	if iteration == -1 {
		return core.NilHash, nil
	}
	numOfProposedBlocks, err := razorUtils.GetNumberOfProposedBlocks(client, account.Address, epoch)
	if err != nil {
		log.Error(err)
		return core.NilHash, err
	}
	maxAltBlocks, err := razorUtils.GetMaxAltBlocks(client, account.Address)
	if err != nil {
		log.Error(err)
		return core.NilHash, err
	}
	if numOfProposedBlocks >= maxAltBlocks {
		log.Debugf("Number of blocks proposed: %d, which is equal or greater than maximum alternative blocks allowed", numOfProposedBlocks)
		log.Debug("Comparing  iterations...")
		lastBlockIndex := numOfProposedBlocks - 1
		lastProposedBlockStruct, err := razorUtils.GetProposedBlock(client, account.Address, epoch, lastBlockIndex)
		if err != nil {
			log.Error(err)
			return core.NilHash, err
		}
		lastIteration := lastProposedBlockStruct.Iteration
		if lastIteration.Cmp(big.NewInt(int64(iteration))) < 0 {
			log.Info("Current iteration is greater than iteration of last proposed block, cannot propose")
			return core.NilHash, nil
		}
		log.Info("Current iteration is less than iteration of last proposed block, can propose")
	}
	medians, err := proposeUtils.MakeBlock(client, account.Address, rogueMode, razorUtils, proposeUtils)
	if err != nil {
		log.Error(err)
		return core.NilHash, err
	}

	log.Debugf("Medians: %d", medians)

	txnOpts := razorUtils.GetTxnOpts(types.TransactionOptions{
		Client:          client,
		Password:        account.Password,
		AccountAddress:  account.Address,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.BlockManagerAddress,
		ABI:             bindings.BlockManagerABI,
		MethodName:      "propose",
		Parameters:      []interface{}{epoch, medians, big.NewInt(int64(iteration)), biggestInfluenceId},
	})

	log.Debugf("Epoch: %d Medians: %d", epoch, medians)
	log.Debugf("Iteration: %d Biggest Influence Id: %d", iteration, biggestInfluenceId)
	log.Info("Proposing block...")
	txn, err := blockManagerUtils.Propose(client, txnOpts, epoch, medians, big.NewInt(int64(iteration)), biggestInfluenceId)
	if err != nil {
		log.Error(err)
		return core.NilHash, err
	}
	log.Info("Txn Hash: ", transactionUtils.Hash(txn))
	return transactionUtils.Hash(txn), nil
}

func getBiggestInfluenceAndId(client *ethclient.Client, address string, epoch uint32, razorUtils utilsInterface) (*big.Int, uint32, error) {
	numberOfStakers, err := razorUtils.GetNumberOfStakers(client, address)
	if err != nil {
		return nil, 0, err
	}
	var biggestInfluenceId uint32
	biggestInfluence := big.NewInt(0)
	for i := 1; i <= int(numberOfStakers); i++ {
		influence, err := razorUtils.GetInfluenceSnapshot(client, address, uint32(i), epoch)
		if err != nil {
			return nil, 0, err
		}
		if influence.Cmp(biggestInfluence) > 0 {
			biggestInfluence = influence
			biggestInfluenceId = uint32(i)
		}
	}
	return biggestInfluence, biggestInfluenceId, nil
}

func getIteration(client *ethclient.Client, address string, proposer types.ElectedProposer, proposeUtils proposeUtilsInterface) int {
	for i := 0; i < 10000000000; i++ {
		proposer.Iteration = i
		isElected := proposeUtils.isElectedProposer(client, address, proposer)
		if isElected {
			return i
		}
	}
	return -1
}

func isElectedProposer(client *ethclient.Client, address string, proposer types.ElectedProposer) bool {
	seed := solsha3.SoliditySHA3([]string{"uint256"}, []interface{}{big.NewInt(int64(proposer.Iteration))})
	pseudoRandomNumber := pseudoRandomNumberGenerator(seed, proposer.NumberOfStakers, proposer.RandaoHash[:])
	//add +1 since prng returns 0 to max-1 and staker start from 1
	pseudoRandomNumber = pseudoRandomNumber.Add(pseudoRandomNumber, big.NewInt(1))
	if pseudoRandomNumber.Cmp(big.NewInt(int64(proposer.StakerId))) != 0 {
		return false
	}
	seed2 := solsha3.SoliditySHA3([]string{"uint32", "uint256"}, []interface{}{proposer.StakerId, big.NewInt(int64(proposer.Iteration))})
	randomHash := solsha3.SoliditySHA3([]string{"bytes32", "bytes32"}, []interface{}{"0x" + hex.EncodeToString(proposer.RandaoHash[:]), "0x" + hex.EncodeToString(seed2)})
	randomHashNumber := big.NewInt(0).SetBytes(randomHash)
	randomHashNumber = randomHashNumber.Mod(randomHashNumber, big.NewInt(int64(math.Exp2(32))))

	influence, err := utils.GetInfluenceSnapshot(client, address, proposer.StakerId, proposer.Epoch)
	if err != nil {
		log.Error("Error in fetching influence of staker: ", err)
		return false
	}
	biggestInfluence := big.NewInt(1).Mul(randomHashNumber, proposer.BiggestInfluence)
	stakerInfluence := big.NewInt(1).Mul(influence, big.NewInt(int64(math.Exp2(32))))
	return biggestInfluence.Cmp(stakerInfluence) <= 0
}

func pseudoRandomNumberGenerator(seed []byte, max uint32, blockHashes []byte) *big.Int {
	hash := solsha3.SoliditySHA3([]string{"bytes32", "bytes32"}, []interface{}{"0x" + hex.EncodeToString(blockHashes), "0x" + hex.EncodeToString(seed)})
	sum := big.NewInt(0).SetBytes(hash)
	return sum.Mod(sum, big.NewInt(int64(max)))
}

func MakeBlock(client *ethclient.Client, address string, rogueMode bool, razorUtils utilsInterface, proposeUtils proposeUtilsInterface) ([]uint32, error) {
	numAssets, err := razorUtils.GetNumActiveAssets(client, address)
	if err != nil {
		return nil, err
	}

	var medians []*big.Int

	epoch, err := razorUtils.GetEpoch(client)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	for assetId := 1; assetId <= int(numAssets.Int64()); assetId++ {
		sortedVotes, err := proposeUtils.getSortedVotes(client, address, uint8(assetId), epoch, razorUtils)
		if err != nil {
			log.Error(err)
			continue
		}

		totalInfluenceRevealed, err := razorUtils.GetTotalInfluenceRevealed(client, address, epoch)
		if err != nil {
			log.Error(err)
			continue
		}

		log.Debug("Sorted Votes: ", sortedVotes)
		log.Debug("Total influence revealed: ", totalInfluenceRevealed)

		var median *big.Int
		if rogueMode {
			median = big.NewInt(int64(rand.Intn(10000000)))
		} else {
			median = proposeUtils.influencedMedian(sortedVotes, totalInfluenceRevealed)
		}
		log.Debugf("Median: %s", median)
		medians = append(medians, median)
	}
	mediansInUint32 := razorUtils.ConvertBigIntArrayToUint32Array(medians)
	return mediansInUint32, nil
}

func getSortedVotes(client *ethclient.Client, address string, assetId uint8, epoch uint32, razorUtils utilsInterface) ([]*big.Int, error) {
	numberOfStakers, err := razorUtils.GetNumberOfStakers(client, address)
	if err != nil {
		return nil, err
	}
	var weightedVoteValues []*big.Int
	for i := 1; i <= int(numberOfStakers); i++ {
		epochLastRevealed, err := razorUtils.GetEpochLastRevealed(client, address, uint32(i))
		if err != nil {
			return nil, err
		}
		if epoch == epochLastRevealed {
			vote, err := razorUtils.GetVoteValue(client, address, assetId, uint32(i))
			if err != nil {
				return nil, err
			}
			influence, err := razorUtils.GetInfluenceSnapshot(client, address, uint32(i), epoch)
			if err != nil {
				return nil, err
			}
			log.Debugf("Vote Value of staker %v: %v", i, vote)
			log.Debugf("Influence snapshot of staker %v: %v", i, influence)
			weightedVote := big.NewInt(1).Mul(vote, influence)
			log.Debugf("Weighted vote of staker %v: %v", i, weightedVote)
			weightedVoteValues = append(weightedVoteValues, weightedVote)
		}
	}

	sortutil.BigIntSlice.Sort(weightedVoteValues)
	return weightedVoteValues, nil
}

func influencedMedian(sortedVotes []*big.Int, totalInfluenceRevealed *big.Int) *big.Int {
	accProd := big.NewInt(0)

	for _, vote := range sortedVotes {
		accProd = accProd.Add(accProd, vote)
	}
	if totalInfluenceRevealed.Cmp(big.NewInt(0)) == 0 {
		return accProd
	}
	return accProd.Div(accProd, totalInfluenceRevealed)
}
