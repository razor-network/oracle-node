package cmd

import (
	"encoding/hex"
	"math"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/utils"

	"github.com/ethereum/go-ethereum/ethclient"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	log "github.com/sirupsen/logrus"
	"modernc.org/sortutil"
)

func Propose(client *ethclient.Client, account types.Account, config types.Configurations, stakerId *big.Int, epoch *big.Int) {
	if state, err := utils.GetDelayedState(client, config.BufferPercent); err != nil || state != 2 {
		log.Error("Not propose state")
		return
	}
	staker, err := utils.GetStaker(client, account.Address, stakerId)
	if err != nil {
		log.Error("Error in fetching staker: ", err)
		return
	}
	numStakers, err := utils.GetNumberOfStakers(client, account.Address)
	if err != nil {
		log.Error("Error in fetching number of stakers: ", err)
		return
	}
	log.Info("Stake: ", staker.Stake)

	biggestInfluence, biggestInfluenceId, err := getBiggestInfluenceAndId(client, account.Address)
	if err != nil {
		log.Error("Error in calculating biggest staker: ", err)
		return
	}
	blockHashes, err := utils.GetBlockHashes(client, account.Address)
	if err != nil {
		log.Error("Error in fetching block hashes: ", blockHashes)
		return
	}
	log.Info("Biggest Influence Id: ", biggestInfluenceId)
	log.Infof("Biggest influence: %s, Stake: %s, Staker Id: %s, Number of Stakers: %s, Blockhashes: %s", biggestInfluence, staker.Stake, stakerId, numStakers, hex.EncodeToString(blockHashes))

	iteration := getIteration(client, account.Address, types.ElectedProposer{
		Stake:            staker.Stake,
		StakerId:         stakerId,
		BiggestInfluence: biggestInfluence,
		NumberOfStakers:  numStakers,
		BlockHashes:      blockHashes,
	})

	log.Info("Iteration: ", iteration)

	if iteration == -1 {
		return
	}

	medians, err := MakeBlock(client, account.Address, epoch)
	if err != nil {
		log.Error(err)
		return
	}

	log.Infof("\nMedians: %s", medians)

	var ids []*big.Int
	numAssets, err := utils.GetNumAssets(client, account.Address)
	if err != nil {
		log.Error(err)
		return
	}
	for assetIndex := 1; assetIndex <= int(numAssets.Int64()); assetIndex++ {
		ids = append(ids, big.NewInt(int64(assetIndex)))
	}
	txnOpts := utils.GetTxnOpts(types.TransactionOptions{
		Client:         client,
		Password:       account.Password,
		AccountAddress: account.Address,
		ChainId:        core.ChainId,
		GasMultiplier:  config.GasMultiplier,
	})
	blockManager := utils.GetBlockManager(client)

	log.Infof("\nEpoch: %s Medians: %s", epoch, medians)
	log.Infof("Asset Ids: %s Iteration: %d Biggest Influence Id: %s\n", ids, iteration, biggestInfluenceId)
	txn, err := blockManager.Propose(txnOpts, epoch, ids, medians, big.NewInt(int64(iteration)), biggestInfluenceId)
	if err != nil {
		log.Error(err)
		return
	}
	log.Infof("Proposed Block\n%s", txn.Hash())
	utils.WaitForBlockCompletion(client, txn.Hash().String())
}

func getBiggestInfluenceAndId(client *ethclient.Client, address string) (*big.Int, *big.Int, error) {
	numberOfStakers, err := utils.GetNumberOfStakers(client, address)
	if err != nil {
		return nil, nil, err
	}
	var biggestInfluenceId *big.Int
	biggestInfluence := big.NewInt(0)
	for i := 1; i <= int(numberOfStakers.Int64()); i++ {
		influence, err := utils.GetInfluence(client, address, big.NewInt(int64(i)))
		if err != nil {
			return nil, nil, err
		}
		if influence.Cmp(biggestInfluence) > 0 {
			biggestInfluence = influence
			biggestInfluenceId = big.NewInt(int64(i))
		}
	}
	return biggestInfluence, biggestInfluenceId, nil
}

func getIteration(client *ethclient.Client, address string, proposer types.ElectedProposer) int {
	for i := 0; i < 10000000000; i++ {
		proposer.Iteration = i
		isElected := isElectedProposer(client, address, proposer)
		if isElected {
			return i
		}
	}
	return -1
}

func isElectedProposer(client *ethclient.Client, address string, proposer types.ElectedProposer) bool {
	seed := solsha3.SoliditySHA3([]string{"uint256"}, []interface{}{big.NewInt(int64(proposer.Iteration))})
	pseudoRandomNumber := pseudoRandomNumberGenerator(seed, proposer.NumberOfStakers, proposer.BlockHashes)
	//add +1 since prng returns 0 to max-1 and staker start from 1
	pseudoRandomNumber = pseudoRandomNumber.Add(pseudoRandomNumber, big.NewInt(1))
	if pseudoRandomNumber.Cmp(proposer.StakerId) != 0 {
		return false
	}
	seed2 := solsha3.SoliditySHA3([]string{"uint256", "uint256"}, []interface{}{proposer.StakerId, big.NewInt(int64(proposer.Iteration))})
	randomHash := solsha3.SoliditySHA3([]string{"bytes32", "bytes32"}, []interface{}{"0x" + hex.EncodeToString(proposer.BlockHashes), "0x" + hex.EncodeToString(seed2)})
	randomHashNumber := big.NewInt(0).SetBytes(randomHash)
	randomHashNumber = randomHashNumber.Mod(randomHashNumber, big.NewInt(int64(math.Exp2(32))))

	influence, err := utils.GetInfluence(client, address, proposer.StakerId)
	if err != nil {
		log.Error("Error in fetching influence of staker: ", err)
		return false
	}
	biggestInfluence := big.NewInt(1).Mul(randomHashNumber, proposer.BiggestInfluence)
	stakerInfluence := big.NewInt(1).Mul(influence, big.NewInt(int64(math.Exp2(32))))
	return biggestInfluence.Cmp(stakerInfluence) <= 0
}

func pseudoRandomNumberGenerator(seed []byte, max *big.Int, blockHashes []byte) *big.Int {
	hash := solsha3.SoliditySHA3([]string{"bytes32", "bytes32"}, []interface{}{"0x" + hex.EncodeToString(blockHashes), "0x" + hex.EncodeToString(seed)})
	sum := big.NewInt(0).SetBytes(hash)
	return sum.Mod(sum, max)
}

func MakeBlock(client *ethclient.Client, address string, epoch *big.Int) ([]*big.Int, error) {
	numAssets, err := utils.GetNumAssets(client, address)
	if err != nil {
		return nil, err
	}

	var medians []*big.Int

	for assetId := 1; assetId <= int(numAssets.Int64()); assetId++ {
		sortedWeights, sortedVotes, err := getSortedVotes(client, address, assetId, epoch)
		if err != nil {
			log.Error(err)
			continue
		}
		log.Info("Sorted Votes: ", sortedVotes)
		log.Info("Sorted Weights: ", sortedWeights)
		median := weightedMedian(sortedVotes, sortedWeights)
		log.Infof("Median: %s", median)
		medians = append(medians, median)
	}
	return medians, nil
}

func getSortedVotes(client *ethclient.Client, address string, assetId int, epoch *big.Int) ([]*big.Int, []*big.Int, error) {
	numberOfStakers, err := utils.GetNumberOfStakers(client, address)
	if err != nil {
		return nil, nil, err
	}
	var (
		voteValues  []*big.Int
		voteWeights []*big.Int
	)
	for i := 1; i <= int(numberOfStakers.Int64()); i++ {
		vote, err := utils.GetVotes(client, address, epoch, big.NewInt(int64(i)), big.NewInt(int64(assetId-1)))
		if err != nil {
			return nil, nil, err
		}
		if vote.Value.Cmp(big.NewInt(0)) == 0 {
			continue
		}
		if !utils.Contains(voteValues, vote.Value) {
			voteValues = append(voteValues, vote.Value)
		}
	}
	sortutil.BigIntSlice.Sort(voteValues)
	for _, value := range voteValues {
		weight, err := utils.GetVoteWeights(client, address, epoch, big.NewInt(int64(assetId-1)), value)
		if err != nil {
			log.Error(err)
			continue
		}
		voteWeights = append(voteWeights, weight)
	}
	return voteWeights, voteValues, nil
}

func weightedMedian(sortedVotes, sortedWeights []*big.Int) *big.Int {
	totalWeight := big.NewInt(0)
	for _, weight := range sortedWeights {
		totalWeight.Add(totalWeight, weight)
	}
	medianWeight := big.NewInt(1).Div(totalWeight, big.NewInt(2))

	weight := big.NewInt(0)
	median := big.NewInt(0)

	for i, vote := range sortedVotes {
		weight = weight.Add(weight, sortedWeights[i])
		if weight.Cmp(medianWeight) >= 0 && median.Cmp(big.NewInt(0)) == 0 {
			median = vote
		}
	}

	if median.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(1)
	}
	return median
}
