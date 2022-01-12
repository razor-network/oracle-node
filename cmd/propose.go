package cmd

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"math"
	"math/big"
	"math/rand"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"

	"github.com/ethereum/go-ethereum/ethclient"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"modernc.org/sortutil"
)

var proposeUtils proposeUtilsInterface

func (utilsStruct UtilsStruct) Propose(client *ethclient.Client, account types.Account, config types.Configurations, stakerId uint32, epoch uint32, rogueData types.Rogue) (common.Hash, error) {
	if state, err := utilsStruct.razorUtils.GetDelayedState(client, config.BufferPercent); err != nil || state != 2 {
		log.Error("Not propose state")
		return core.NilHash, err
	}
	staker, err := utilsStruct.razorUtils.GetStaker(client, account.Address, stakerId)
	if err != nil {
		log.Error("Error in fetching staker: ", err)
		return core.NilHash, err
	}
	numStakers, err := utilsStruct.razorUtils.GetNumberOfStakers(client, account.Address)
	if err != nil {
		log.Error("Error in fetching number of stakers: ", err)
		return core.NilHash, err
	}
	log.Debug("Stake: ", staker.Stake)

	biggestStake, biggestStakerId, err := utilsStruct.proposeUtils.getBiggestStakeAndId(client, account.Address, epoch, utilsStruct)
	if err != nil {
		log.Error("Error in calculating biggest staker: ", err)
		return core.NilHash, err
	}

	randaoHash, err := utilsStruct.razorUtils.GetRandaoHash(client)
	if err != nil {
		log.Error("Error in fetching random hash: ", err)
		return core.NilHash, err
	}
	log.Debug("Biggest Staker Id: ", biggestStakerId)
	log.Debugf("Biggest stake: %s, Stake: %s, Staker Id: %d, Number of Stakers: %d, Randao Hash: %s", biggestStake, staker.Stake, stakerId, numStakers, hex.EncodeToString(randaoHash[:]))

	iteration := utilsStruct.proposeUtils.getIteration(client, types.ElectedProposer{
		Stake:           staker.Stake,
		StakerId:        stakerId,
		BiggestStake:    biggestStake,
		NumberOfStakers: numStakers,
		RandaoHash:      randaoHash,
		Epoch:           epoch,
	}, utilsStruct)

	log.Debug("Iteration: ", iteration)

	if iteration == -1 {
		return core.NilHash, nil
	}
	numOfProposedBlocks, err := utilsStruct.razorUtils.GetNumberOfProposedBlocks(client, account.Address, epoch)
	if err != nil {
		log.Error(err)
		return core.NilHash, err
	}
	maxAltBlocks, err := utilsStruct.razorUtils.GetMaxAltBlocks(client, account.Address)
	if err != nil {
		log.Error(err)
		return core.NilHash, err
	}
	if numOfProposedBlocks >= maxAltBlocks {
		log.Debugf("Number of blocks proposed: %d, which is equal or greater than maximum alternative blocks allowed", numOfProposedBlocks)
		log.Debug("Comparing  iterations...")
		lastBlockIndex := numOfProposedBlocks - 1
		lastProposedBlockStruct, err := utilsStruct.razorUtils.GetProposedBlock(client, account.Address, epoch, uint32(lastBlockIndex))
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
	medians, err := utilsStruct.proposeUtils.MakeBlock(client, account.Address, rogueData, utilsStruct)
	if err != nil {
		log.Error(err)
		return core.NilHash, err
	}

	log.Debugf("Medians: %d", medians)

	txnOpts := utilsStruct.razorUtils.GetTxnOpts(types.TransactionOptions{
		Client:          client,
		Password:        account.Password,
		AccountAddress:  account.Address,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.BlockManagerAddress,
		ABI:             bindings.BlockManagerABI,
		MethodName:      "propose",
		Parameters:      []interface{}{epoch, medians, big.NewInt(int64(iteration)), biggestStakerId},
	})

	log.Debugf("Epoch: %d Medians: %d", epoch, medians)
	log.Debugf("Iteration: %d Biggest Staker Id: %d", iteration, biggestStakerId)
	log.Info("Proposing block...")
	txn, err := utilsStruct.blockManagerUtils.Propose(client, txnOpts, epoch, medians, big.NewInt(int64(iteration)), biggestStakerId)
	if err != nil {
		log.Error(err)
		return core.NilHash, err
	}
	log.Info("Txn Hash: ", utilsStruct.transactionUtils.Hash(txn))
	return utilsStruct.transactionUtils.Hash(txn), nil
}

func getBiggestStakeAndId(client *ethclient.Client, address string, epoch uint32, utilsStruct UtilsStruct) (*big.Int, uint32, error) {
	numberOfStakers, err := utilsStruct.razorUtils.GetNumberOfStakers(client, address)
	if err != nil {
		return nil, 0, err
	}
	var biggestStakerId uint32
	biggestStake := big.NewInt(0)
	for i := 1; i <= int(numberOfStakers); i++ {
		stake, err := utilsStruct.razorUtils.GetStakeSnapshot(client, uint32(i), epoch)
		if err != nil {
			return nil, 0, err
		}
		if stake.Cmp(biggestStake) > 0 {
			biggestStake = stake
			biggestStakerId = uint32(i)
		}
	}
	return biggestStake, biggestStakerId, nil
}

func getIteration(client *ethclient.Client, proposer types.ElectedProposer, utilsStruct UtilsStruct) int {
	for i := 0; i < 10000000; i++ {
		proposer.Iteration = i
		isElected := utilsStruct.proposeUtils.isElectedProposer(client, proposer, utilsStruct)
		if isElected {
			return i
		}
	}
	return -1
}

func isElectedProposer(client *ethclient.Client, proposer types.ElectedProposer, utilsStruct UtilsStruct) bool {
	seed := solsha3.SoliditySHA3([]string{"uint256"}, []interface{}{big.NewInt(int64(proposer.Iteration))})
	pseudoRandomNumber := pseudoRandomNumberGenerator(seed, proposer.NumberOfStakers, proposer.RandaoHash[:])
	//add +1 since prng returns 0 to max-1 and staker start from 1
	pseudoRandomNumber = pseudoRandomNumber.Add(pseudoRandomNumber, big.NewInt(1))
	if pseudoRandomNumber.Cmp(big.NewInt(int64(proposer.StakerId))) != 0 {
		return false
	}
	seed2 := solsha3.SoliditySHA3([]string{"uint256", "uint256"}, []interface{}{big.NewInt(int64(proposer.StakerId)), big.NewInt(int64(proposer.Iteration))})
	randomHash := solsha3.SoliditySHA3([]string{"bytes32", "bytes32"}, []interface{}{"0x" + hex.EncodeToString(proposer.RandaoHash[:]), "0x" + hex.EncodeToString(seed2)})
	randomHashNumber := big.NewInt(0).SetBytes(randomHash)
	randomHashNumber = randomHashNumber.Mod(randomHashNumber, big.NewInt(int64(math.Exp2(32))))

	stake, err := utilsStruct.razorUtils.GetStakeSnapshot(client, proposer.StakerId, proposer.Epoch)
	if err != nil {
		log.Error("Error in fetching influence of staker: ", err)
		return false
	}
	biggestStake := big.NewInt(1).Mul(randomHashNumber, proposer.BiggestStake)
	currentStakerStake := big.NewInt(1).Mul(stake, big.NewInt(int64(math.Exp2(32))))
	return biggestStake.Cmp(currentStakerStake) <= 0
}

func pseudoRandomNumberGenerator(seed []byte, max uint32, blockHashes []byte) *big.Int {
	hash := solsha3.SoliditySHA3([]string{"bytes32", "bytes32"}, []interface{}{"0x" + hex.EncodeToString(blockHashes), "0x" + hex.EncodeToString(seed)})
	sum := big.NewInt(0).SetBytes(hash)
	return sum.Mod(sum, big.NewInt(int64(max)))
}

func MakeBlock(client *ethclient.Client, address string, rogueData types.Rogue, utilsStruct UtilsStruct) ([]uint32, error) {
	numAssets, err := utilsStruct.razorUtils.GetNumActiveAssets(client)
	if err != nil {
		return nil, err
	}

	var medians []*big.Int

	epoch, err := utilsStruct.razorUtils.GetEpoch(client)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	for assetId := 1; assetId <= int(numAssets.Int64()); assetId++ {
		sortedVotes, err := utilsStruct.proposeUtils.getSortedVotes(client, address, uint16(assetId), epoch, utilsStruct)
		if err != nil {
			log.Error(err)
			continue
		}

		totalInfluenceRevealed, err := utilsStruct.razorUtils.GetTotalInfluenceRevealed(client, epoch)
		if err != nil {
			log.Error(err)
			continue
		}

		log.Debug("Sorted Votes: ", sortedVotes)
		log.Debug("Total influence revealed: ", totalInfluenceRevealed)

		var median *big.Int
		if rogueData.IsRogue && utils.Contains(rogueData.RogueMode, "propose") {
			median = big.NewInt(int64(rand.Intn(10000000)))
		} else {
			median = utilsStruct.proposeUtils.influencedMedian(sortedVotes, totalInfluenceRevealed)
		}
		log.Debugf("Median: %s", median)
		medians = append(medians, median)
	}
	mediansInUint32 := utilsStruct.razorUtils.ConvertBigIntArrayToUint32Array(medians)
	return mediansInUint32, nil
}

func getSortedVotes(client *ethclient.Client, address string, assetId uint16, epoch uint32, utilsStruct UtilsStruct) ([]*big.Int, error) {
	numberOfStakers, err := utilsStruct.razorUtils.GetNumberOfStakers(client, address)
	if err != nil {
		return nil, err
	}
	var weightedVoteValues []*big.Int
	for i := 1; i <= int(numberOfStakers); i++ {
		epochLastRevealed, err := utilsStruct.razorUtils.GetEpochLastRevealed(client, address, uint32(i))
		if err != nil {
			return nil, err
		}
		if epoch == epochLastRevealed {
			vote, err := utilsStruct.razorUtils.GetVoteValue(client, assetId, uint32(i))
			if err != nil {
				return nil, err
			}
			influence, err := utilsStruct.razorUtils.GetInfluenceSnapshot(client, uint32(i), epoch)
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
