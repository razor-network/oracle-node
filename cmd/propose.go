package cmd

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"math"
	"math/big"
	"math/rand"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"
	"sort"
)

var _mediansData []*big.Int

// Index reveal events of staker's
// Reveal Event would have two things, activeCollectionIndex/medianIndex and values
// Loop
// Medians Path : Find Medians on basis of weight (influence)
// Ids path : collectionManager.getIndexToIdRegistry(activeCollectionIndex)

// Find iteration using salt as seed

func (*UtilsStruct) Propose(client *ethclient.Client, config types.Configurations, account types.Account, staker bindings.StructsStaker, epoch uint32, blockNumber *big.Int, rogueData types.Rogue) (common.Hash, error) {
	if state, err := razorUtils.GetDelayedState(client, config.BufferPercent); err != nil || state != 2 {
		log.Error("Not propose state")
		return core.NilHash, err
	}
	numStakers, err := razorUtils.GetNumberOfStakers(client)
	if err != nil {
		log.Error("Error in fetching number of stakers: ", err)
		return core.NilHash, err
	}
	log.Debug("Stake: ", staker.Stake)

	biggestStake, biggestStakerId, err := cmdUtils.GetBiggestStakeAndId(client, account.Address, epoch)
	if err != nil {
		log.Error("Error in calculating biggest staker: ", err)
		return core.NilHash, err
	}

	salt, err := cmdUtils.GetSalt(client, epoch)
	if err != nil {
		return core.NilHash, err
	}

	log.Debugf("Biggest staker Id: %d Biggest stake: %s, Stake: %s, Staker Id: %d, Number of Stakers: %d, Salt: %s", biggestStakerId, biggestStake, staker.Stake, staker.Id, numStakers, hex.EncodeToString(salt[:]))

	iteration := cmdUtils.GetIteration(client, types.ElectedProposer{
		Stake:           staker.Stake,
		StakerId:        staker.Id,
		BiggestStake:    biggestStake,
		NumberOfStakers: numStakers,
		Salt:            salt,
		Epoch:           epoch,
	})

	log.Debug("Iteration: ", iteration)

	if iteration == -1 {
		return core.NilHash, nil
	}
	numOfProposedBlocks, err := razorUtils.GetNumberOfProposedBlocks(client, epoch)
	if err != nil {
		log.Error(err)
		return core.NilHash, err
	}
	maxAltBlocks, err := razorUtils.GetMaxAltBlocks(client)
	if err != nil {
		log.Error(err)
		return core.NilHash, err
	}
	if numOfProposedBlocks >= maxAltBlocks {
		log.Debugf("Number of blocks proposed: %d, which is equal or greater than maximum alternative blocks allowed", numOfProposedBlocks)
		log.Debug("Comparing  iterations...")
		lastBlockIndex := numOfProposedBlocks - 1
		lastProposedBlockStruct, err := razorUtils.GetProposedBlock(client, epoch, uint32(lastBlockIndex))
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
	medians, err := cmdUtils.MakeBlock(client, blockNumber, epoch, rogueData)
	if err != nil {
		log.Error(err)
		return core.NilHash, err
	}

	_mediansData = razorUtils.ConvertUint32ArrayToBigIntArray(medians)

	log.Debug("Saving median data for recovery")
	fileName, err := cmdUtils.GetMedianDataFileName(account.Address)
	if err != nil {
		log.Error("Error in getting file name to save median data: ", err)
		return core.NilHash, nil
	}
	err = razorUtils.SaveDataToFile(fileName, epoch, _mediansData)
	if err != nil {
		log.Errorf("Error in saving data to file %s: %t", fileName, err)
		return core.NilHash, nil
	}
	log.Debug("Data saved!")

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
		Parameters:      []interface{}{epoch, medians, big.NewInt(int64(iteration)), biggestStakerId},
	})

	log.Debugf("Epoch: %d Medians: %d", epoch, medians)
	log.Debugf("Iteration: %d Biggest Staker Id: %d", iteration, biggestStakerId)
	log.Info("Proposing block...")
	//TODO: Check if this is correct for ids
	ids, err := razorUtils.GetActiveCollectionIds(client)
	if err != nil {
		return core.NilHash, err
	}

	txn, err := blockManagerUtils.Propose(client, txnOpts, epoch, ids, medians, big.NewInt(int64(iteration)), biggestStakerId)
	if err != nil {
		log.Error(err)
		return core.NilHash, err
	}
	log.Info("Txn Hash: ", transactionUtils.Hash(txn))
	return transactionUtils.Hash(txn), nil
}

func (*UtilsStruct) GetBiggestStakeAndId(client *ethclient.Client, address string, epoch uint32) (*big.Int, uint32, error) {
	numberOfStakers, err := razorUtils.GetNumberOfStakers(client)
	if err != nil {
		return nil, 0, err
	}
	var biggestStakerId uint32
	biggestStake := big.NewInt(0)
	for i := 1; i <= int(numberOfStakers); i++ {
		stake, err := razorUtils.GetStakeSnapshot(client, uint32(i), epoch)
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

func (*UtilsStruct) GetIteration(client *ethclient.Client, proposer types.ElectedProposer) int {
	stake, err := razorUtils.GetStakeSnapshot(client, proposer.StakerId, proposer.Epoch)
	if err != nil {
		log.Error("Error in fetching influence of staker: ", err)
		return -1
	}
	currentStakerStake := big.NewInt(1).Mul(stake, big.NewInt(int64(math.Exp2(32))))
	for i := 0; i < 10000000; i++ {
		proposer.Iteration = i
		isElected := cmdUtils.IsElectedProposer(proposer, currentStakerStake)
		if isElected {
			return i
		}
	}
	return -1
}

func (*UtilsStruct) IsElectedProposer(proposer types.ElectedProposer, currentStakerStake *big.Int) bool {
	seed := solsha3.SoliditySHA3([]string{"uint256"}, []interface{}{big.NewInt(int64(proposer.Iteration))})
	pseudoRandomNumber := pseudoRandomNumberGenerator(seed, proposer.NumberOfStakers, proposer.Salt[:])
	//add +1 since prng returns 0 to max-1 and staker start from 1
	pseudoRandomNumber = pseudoRandomNumber.Add(pseudoRandomNumber, big.NewInt(1))
	if pseudoRandomNumber.Cmp(big.NewInt(int64(proposer.StakerId))) != 0 {
		return false
	}
	seed2 := solsha3.SoliditySHA3([]string{"uint256", "uint256"}, []interface{}{big.NewInt(int64(proposer.StakerId)), big.NewInt(int64(proposer.Iteration))})
	randomHash := solsha3.SoliditySHA3([]string{"bytes32", "bytes32"}, []interface{}{"0x" + hex.EncodeToString(proposer.Salt[:]), "0x" + hex.EncodeToString(seed2)})
	randomHashNumber := big.NewInt(0).SetBytes(randomHash)
	randomHashNumber = randomHashNumber.Mod(randomHashNumber, big.NewInt(int64(math.Exp2(32))))
	biggestStake := big.NewInt(1).Mul(randomHashNumber, proposer.BiggestStake)
	return biggestStake.Cmp(currentStakerStake) <= 0
}

func pseudoRandomNumberGenerator(seed []byte, max uint32, blockHashes []byte) *big.Int {
	hash := solsha3.SoliditySHA3([]string{"bytes32", "bytes32"}, []interface{}{"0x" + hex.EncodeToString(blockHashes), "0x" + hex.EncodeToString(seed)})
	sum := big.NewInt(0).SetBytes(hash)
	return sum.Mod(sum, big.NewInt(int64(max)))
}

func (*UtilsStruct) getSortedRevealedValues(client *ethclient.Client, blockNumber *big.Int, epoch uint32) (*types.RevealedDataMaps, error) {
	assignedAsset, err := cmdUtils.IndexRevealEventsOfCurrentEpoch(client, blockNumber, epoch)
	if err != nil {
		return nil, err
	}
	revealedValuesWithIndex := make(map[uint16][]uint32)
	voteWeights := make(map[uint32]*big.Int)
	influenceSum := make(map[uint16]*big.Int)
	for _, asset := range assignedAsset {
		if revealedValuesWithIndex[asset.MedianIndex] == nil {
			revealedValuesWithIndex[asset.MedianIndex] = []uint32{asset.Value}
		} else {
			if !utils.Contains(revealedValuesWithIndex[asset.MedianIndex], asset.Value) {
				revealedValuesWithIndex[asset.MedianIndex] = append(revealedValuesWithIndex[asset.MedianIndex], asset.Value)
			}
		}
		//Calculate vote weights
		if voteWeights[asset.Value] == nil {
			voteWeights[asset.Value] = big.NewInt(0)
		}
		voteWeights[asset.Value] = big.NewInt(0).Add(voteWeights[asset.Value], asset.Influence)

		//Calculate influence sum
		if influenceSum[asset.MedianIndex] == nil {
			influenceSum[asset.MedianIndex] = big.NewInt(0)
		}
		influenceSum[asset.MedianIndex] = big.NewInt(0).Add(influenceSum[asset.MedianIndex], asset.Influence)
	}
	//sort revealed values
	for _, element := range revealedValuesWithIndex {
		sort.Slice(element, func(i, j int) bool {
			return element[i] < element[j]
		})
	}
	return &types.RevealedDataMaps{
		SortedRevealedValues: revealedValuesWithIndex,
		VoteWeights:          voteWeights,
		InfluenceSum:         influenceSum,
	}, nil
}

func (*UtilsStruct) MakeBlock(client *ethclient.Client, blockNumber *big.Int, epoch uint32, rogueData types.Rogue) ([]uint32, error) {
	revealedDataMaps, err := cmdUtils.getSortedRevealedValues(client, blockNumber, epoch)
	if err != nil {
		return nil, err
	}

	numActiveCollections, err := razorUtils.GetNumActiveCollections(client)
	if err != nil {
		return nil, err
	}

	var medians []uint32

	//TODO: Check if 0 is the correct value for collectionId
	for collectionId := uint16(1); collectionId <= numActiveCollections; collectionId++ {
		influenceSum := revealedDataMaps.InfluenceSum[collectionId]
		if influenceSum.Cmp(big.NewInt(0)) != 0 {
			accWeight := big.NewInt(0)
			for i := 0; i < len(revealedDataMaps.SortedRevealedValues); i++ {
				revealedValue := revealedDataMaps.SortedRevealedValues[collectionId][i]
				if rogueData.IsRogue && utils.Contains(rogueData.RogueMode, "propose") {
					medians = append(medians, rand.Uint32())
				} else {
					accWeight = accWeight.Add(accWeight, revealedDataMaps.VoteWeights[revealedValue])
					if accWeight.Cmp(influenceSum.Div(influenceSum, big.NewInt(2))) > 0 {
						medians = append(medians, revealedValue)
					}
				}
			}
		}
	}
	return medians, nil
}

func (*UtilsStruct) GetMedianDataFileName(address string) (string, error) {
	homeDir, err := razorUtils.GetDefaultPath()
	if err != nil {
		return "", err
	}
	return homeDir + "/" + address + "_median", nil
}

func (*UtilsStruct) InfluencedMedian(sortedVotes []*big.Int, totalInfluenceRevealed *big.Int) *big.Int {
	accProd := big.NewInt(0)

	for _, vote := range sortedVotes {
		accProd = accProd.Add(accProd, vote)
	}
	if totalInfluenceRevealed.Cmp(big.NewInt(0)) == 0 {
		return accProd
	}
	return accProd.Div(accProd, totalInfluenceRevealed)
}
