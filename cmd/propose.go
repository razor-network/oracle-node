package cmd

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/ethclient"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	log "github.com/sirupsen/logrus"
	"math"
	"math/big"
	"modernc.org/sortutil"
	"razor/core/types"
	"razor/utils"
)

func Propose(client *ethclient.Client, account types.Account, config types.Configurations, stakerId *big.Int, epoch *big.Int) {
	if state, err := utils.GetDelayedState(client); err != nil || state != 2 {
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

	biggestStake, biggestStakerId, err := getBiggestStakeAndId(client, account.Address)
	if err != nil {
		log.Error("Error in calculating biggest staker: ", err)
		return
	}
	blockHashes, err := utils.GetBlockHashes(client, account.Address)
	if err != nil {
		log.Error("Error in fetching block hashes: ", blockHashes)
		return
	}
	log.Info("Biggest Staker Id: ", biggestStakerId)
	log.Infof("Biggest stake: %s, Stake: %s, Staker Id: %s, Number of Stakers: %s, Blockhashes: %s", biggestStake, staker.Stake, stakerId, numStakers, hex.EncodeToString(blockHashes))

	iteration := getIteration(types.ElectedProposer{
		Stake:           staker.Stake,
		StakerId:        stakerId,
		BiggestStake:    biggestStake,
		NumberOfStakers: numStakers,
		BlockHashes:     blockHashes,
	})

	log.Info("Iteration: ", iteration)

	if iteration == -1 {
		return
	}

	medians, lowerCutOffs, higherCutOffs, err := MakeBlock(client, account.Address, epoch)
	if err != nil {
		log.Error(err)
		return
	}

	log.Infof("\nMedians: %s Lower Cut Offs: %s Higher Cut Offs: %s \n", medians, lowerCutOffs, higherCutOffs)

	jobs, err := utils.GetActiveJobs(client, account.Address)
	var jobIds []*big.Int
	for _, job := range jobs {
		jobIds = append(jobIds, job.Id)
	}
	if err != nil {
		log.Error(err)
		return
	}
	txnOpts := utils.GetTxnOpts(types.TransactionOptions{
		Client:         client,
		Password:       account.Password,
		AccountAddress: account.Address,
		ChainId:        config.ChainId,
		GasMultiplier:  config.GasMultiplier,
	})
	blockManager := utils.GetBlockManager(client)

	log.Infof("\nEpoch: %s Medians: %s Lower Cut Offs: %s Higher Cut Offs: %s", epoch, medians, lowerCutOffs, higherCutOffs)
	log.Infof("Job Ids: %s Iteration: %d Biggest Staker Id: %s\n", jobIds, iteration, biggestStakerId)
	txn, err := blockManager.Propose(txnOpts, epoch, jobIds, medians, lowerCutOffs, higherCutOffs, big.NewInt(int64(iteration)), biggestStakerId)
	if err != nil {
		log.Error(err)
		return
	}
	log.Infof("Proposed Block\n%s", txn.Hash())
	utils.WaitForBlockCompletion(client, txn.Hash().String())
}

func getBiggestStakeAndId(client *ethclient.Client, address string) (*big.Int, *big.Int, error) {
	numberOfStakers, err := utils.GetNumberOfStakers(client, address)
	if err != nil {
		return nil, nil, err
	}
	var biggestStakerId *big.Int
	biggestStake := big.NewInt(0)
	for i := 0; i < int(numberOfStakers.Int64()); i++ {
		staker, err := utils.GetStaker(client, address, big.NewInt(int64(i)))
		if err != nil {
			return nil, nil, err
		}
		if staker.Stake.Cmp(biggestStake) > 0 {
			biggestStake = staker.Stake
			biggestStakerId = staker.Id
		}
	}
	return biggestStake, biggestStakerId, nil
}

func getIteration(proposer types.ElectedProposer) int {
	// TODO: Check why 10000000000
	for i := 0; i < 10000000000; i++ {
		proposer.Iteration = i
		isElected := isElectedProposer(proposer)
		if isElected {
			return i
		}
	}
	return -1
}

func isElectedProposer(proposer types.ElectedProposer) bool {
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

	biggestRandomStake := big.NewInt(1).Mul(randomHashNumber, proposer.BiggestStake)
	stake := big.NewInt(1).Mul(proposer.Stake, big.NewInt(int64(math.Exp2(32))))
	if biggestRandomStake.Cmp(stake) > 0 {
		return false
	}
	return true
}

func pseudoRandomNumberGenerator(seed []byte, max *big.Int, blockHashes []byte) *big.Int {
	hash := solsha3.SoliditySHA3([]string{"bytes32", "bytes32"}, []interface{}{"0x" + hex.EncodeToString(blockHashes), "0x" + hex.EncodeToString(seed)})
	sum := big.NewInt(0).SetBytes(hash)
	return sum.Mod(sum, max)
}

func MakeBlock(client *ethclient.Client, address string, epoch *big.Int) ([]*big.Int, []*big.Int, []*big.Int, error) {
	jobs, err := utils.GetActiveJobs(client, address)
	if err != nil {
		return nil, nil, nil, err
	}
	var (
		medians       []*big.Int
		lowerCutOffs  []*big.Int
		higherCutOffs []*big.Int
	)
	for assetId := 0; assetId < len(jobs); assetId++ {
		sortedWeights, sortedVotes, err := getSortedVotes(client, address, assetId, epoch)
		if err != nil {
			log.Error(err)
			continue
		}
		log.Info("Sorted Votes: ", sortedVotes)
		log.Info("Sorted Weights: ", sortedWeights)
		median, lowerCutOff, higherCutOff := weightedMedianAndCutOffs(sortedVotes, sortedWeights)
		log.Infof("Median: %s, Lower Cut Off: %s, Higher Cut Off: %s", median, lowerCutOff, higherCutOff)
		medians = append(medians, median)
		lowerCutOffs = append(lowerCutOffs, lowerCutOff)
		higherCutOffs = append(higherCutOffs, higherCutOff)
	}
	return medians, lowerCutOffs, higherCutOffs, nil
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
		vote, err := utils.GetVotes(client, address, epoch, big.NewInt(int64(i)), big.NewInt(int64(assetId)))
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
		weight, err := utils.GetVoteWeights(client, address, epoch, big.NewInt(int64(assetId)), value)
		if err != nil {
			log.Error(err)
			continue
		}
		voteWeights = append(voteWeights, weight)
	}
	return voteWeights, voteValues, nil
}

func weightedMedianAndCutOffs(sortedVotes, sortedWeights []*big.Int) (*big.Int, *big.Int, *big.Int) {
	//TODO: median calculation doesn't calculate median actually
	totalWeight := big.NewInt(0)
	for _, weight := range sortedWeights {
		totalWeight.Add(totalWeight, weight)
	}
	medianWeight := big.NewInt(1).Div(totalWeight, big.NewInt(2))
	// TODO: understand the lower cut off weight and higher cut off weight calculation
	lowerCutOffWeight := big.NewInt(1).Div(totalWeight, big.NewInt(4))
	intermediateHigherCutOffWeight := big.NewInt(1).Mul(totalWeight, big.NewInt(3))
	higherCutOffWeight := big.NewInt(1).Div(intermediateHigherCutOffWeight, big.NewInt(4))

	weight := big.NewInt(0)
	median := big.NewInt(0)
	lowerCutOff := big.NewInt(0)
	higherCutOff := big.NewInt(0)

	for i, vote := range sortedVotes {
		weight = weight.Add(weight, sortedWeights[i])
		if weight.Cmp(medianWeight) >= 0 && median.Cmp(big.NewInt(0)) == 0 {
			median = vote
		}
		if weight.Cmp(lowerCutOffWeight) >= 0 && lowerCutOff.Cmp(big.NewInt(0)) == 0 {
			lowerCutOff = vote
		}
		if weight.Cmp(higherCutOffWeight) >= 0 && higherCutOff.Cmp(big.NewInt(0)) == 0 {
			higherCutOff = vote
		}
	}

	return median, lowerCutOff, higherCutOff
}
