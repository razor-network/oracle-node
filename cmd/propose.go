package cmd

import (
	"encoding/hex"
	"math"
	"math/big"
	"math/rand"
	"razor/core"
	"razor/core/types"
	"razor/utils"

	"github.com/ethereum/go-ethereum/ethclient"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	log "github.com/sirupsen/logrus"
	"modernc.org/sortutil"
)

func Propose(client *ethclient.Client, account types.Account, config types.Configurations, stakerId uint32, epoch uint32, rogueMode bool) {
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

	randaoHash, err := utils.GetRandaoHash(client, account.Address)
	if err != nil {
		log.Error("Error in fetching random hash: ", err)
		return
	}
	log.Info("Biggest Influence Id: ", biggestInfluenceId)
	log.Infof("Biggest influence: %s, Stake: %s, Staker Id: %d, Number of Stakers: %d, Randao Hash: %s", biggestInfluence, staker.Stake, stakerId, numStakers, hex.EncodeToString(randaoHash[:]))

	iteration := getIteration(client, account.Address, types.ElectedProposer{
		Stake:            staker.Stake,
		StakerId:         stakerId,
		BiggestInfluence: biggestInfluence,
		NumberOfStakers:  numStakers,
		RandaoHash:       randaoHash,
	})

	log.Info("Iteration: ", iteration)

	if iteration == -1 {
		return
	}
	numOfProposedBlocks, err := utils.GetNumberOfProposedBlocks(client, account.Address, epoch)
	if err != nil {
		log.Error(err)
	}
	maxAltBlocks, err := utils.GetMaxAltBlocks(client, account.Address)
	if err != nil {
		log.Error(err)
	}
	if numOfProposedBlocks >= maxAltBlocks {
		log.Infof("Number of blocks proposed: %d, which is equal or greater than maximum alternative blocks allowed", numOfProposedBlocks)
		log.Info("Comparing  iterations...")
		lastBlockIndex := numOfProposedBlocks - 1
		lastProposedBlockStruct, err := utils.GetProposedBlock(client, account.Address, epoch, lastBlockIndex)
		if err != nil {
			log.Error(err)
		}
		lastIteration := lastProposedBlockStruct.Block.Iteration
		if lastIteration.Cmp(big.NewInt(int64(iteration))) < 0 {
			log.Info("Current iteration is greater than iteration of last proposed block, cannot propose")
			return
		}
		log.Info("Current iteration is less than iteration of last proposed block, can propose")
	}
	medians, err := MakeBlock(client, account.Address, rogueMode)
	if err != nil {
		log.Error(err)
		return
	}

	log.Infof("Medians: %d", medians)

	ids, err := utils.GetActiveAssetIds(client, account.Address)
	if err != nil {
		log.Error(err)
		return
	}
	txnOpts := utils.GetTxnOpts(types.TransactionOptions{
		Client:         client,
		Password:       account.Password,
		AccountAddress: account.Address,
		ChainId:        core.ChainId,
		Config:         config,
	})
	blockManager := utils.GetBlockManager(client)

	log.Infof("Epoch: %d Medians: %d", epoch, medians)
	log.Infof("Asset Ids: %s Iteration: %d Biggest Influence Id: %d", ids, iteration, biggestInfluenceId)
	txn, err := blockManager.Propose(txnOpts, epoch, medians, big.NewInt(int64(iteration)), biggestInfluenceId)
	if err != nil {
		log.Error(err)
		return
	}
	log.Info("Block Proposed...")
	log.Info("Txn Hash: ", txn.Hash())
	utils.WaitForBlockCompletion(client, txn.Hash().String())
}

func getBiggestInfluenceAndId(client *ethclient.Client, address string) (*big.Int, uint32, error) {
	numberOfStakers, err := utils.GetNumberOfStakers(client, address)
	if err != nil {
		return nil, 0, err
	}
	var biggestInfluenceId uint32
	biggestInfluence := big.NewInt(0)
	for i := 1; i <= int(numberOfStakers); i++ {
		influence, err := utils.GetInfluence(client, address, uint32(i))
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

	influence, err := utils.GetInfluence(client, address, proposer.StakerId)
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

func MakeBlock(client *ethclient.Client, address string, rogueMode bool) ([]uint32, error) {
	numAssets, err := utils.GetNumAssets(client, address)
	if err != nil {
		return nil, err
	}

	var medians []*big.Int

	epoch, err := utils.GetEpoch(client, address)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	for assetId := 1; assetId <= int(numAssets); assetId++ {
		sortedVotes, err := getSortedVotes(client, address)
		if err != nil {
			log.Error(err)
			// TODO: Add retry mechanism
			continue
		}

		influenceSnapshot, err := utils.GetInfluenceSnapshot(client, address, epoch)
		if err != nil {
			log.Error(err)
			// TODO: Add retry mechanism
			continue
		}

		totalInfluenceRevealed, err := utils.GetTotalInfluenceRevealed(client, address, epoch)
		if err != nil {
			log.Error(err)
			// TODO: Add retry mechanism
			continue
		}

		log.Info("Sorted Votes: ", sortedVotes)
		log.Info("Influence Snapshot: ", influenceSnapshot)
		log.Debug("Total influence revealed: ", totalInfluenceRevealed)

		var median *big.Int
		if rogueMode {
			median = big.NewInt(int64(rand.Intn(10000000)))
		} else {
			median = influencedMedian(sortedVotes, influenceSnapshot, totalInfluenceRevealed)
		}
		log.Infof("Median: %s", median)
		medians = append(medians, median)
	}
	mediansInUint32 := utils.ConvertBigIntArrayToUint32Array(medians)
	return mediansInUint32, nil
}

func getSortedVotes(client *ethclient.Client, address string) ([]*big.Int, error) {
	numberOfStakers, err := utils.GetNumberOfStakers(client, address)
	if err != nil {
		return nil, err
	}
	var voteValues []*big.Int

	for i := 1; i <= int(numberOfStakers); i++ {
		vote, err := utils.GetVotes(client, address, uint32(i))
		if err != nil {
			return nil, err
		}
		voteValues = append(voteValues, vote.Values...)
	}

	sortutil.BigIntSlice.Sort(voteValues)
	return voteValues, nil
}

func influencedMedian(sortedVotes []*big.Int, influence *big.Int, totalInfluenceRevealed *big.Int) *big.Int {
	accProd := big.NewInt(0)

	for _, vote := range sortedVotes {
		accProd = accProd.Add(accProd, vote.Mul(vote, influence))
	}
	if totalInfluenceRevealed.Cmp(big.NewInt(0)) == 0 {
		return accProd
	}
	return accProd.Div(accProd, totalInfluenceRevealed)
}
