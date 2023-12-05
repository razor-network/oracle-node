//Package cmd provides all functions related to command line
package cmd

import (
	"encoding/hex"
	"errors"
	"math"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"
	"sort"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

var globalProposedDataStruct types.ProposeFileData

// Index reveal events of staker's
// Reveal Event would have two things, activeCollectionIndex/medianIndex and values
// Loop
// Medians Path : Find Medians on basis of weight (influence)
// Ids path : collectionManager.getIndexToIdRegistry(activeCollectionIndex)

// Find iteration using salt as seed

//This functions handles the propose state
func (*UtilsStruct) Propose(client *ethclient.Client, config types.Configurations, account types.Account, staker bindings.StructsStaker, epoch uint32, blockNumber *big.Int, rogueData types.Rogue) error {
	if state, err := razorUtils.GetBufferedState(client, config.BufferPercent); err != nil || state != 2 {
		log.Error("Not propose state")
		return err
	}
	numStakers, err := razorUtils.GetNumberOfStakers(client)
	if err != nil {
		log.Error("Error in fetching number of stakers: ", err)
		return err
	}
	log.Debug("Propose: Number of stakers: ", numStakers)
	log.Debug("Propose: Stake: ", staker.Stake)

	var (
		biggestStake     *big.Int
		biggestStakerId  uint32
		biggestStakerErr error
	)

	if rogueData.IsRogue && utils.Contains(rogueData.RogueMode, "biggestStakerId") {
		log.Warn("YOU ARE PROPOSING IN ROGUE MODE, THIS CAN INCUR PENALTIES!")
		// If staker is going rogue with biggestStakerId than we do biggestStakerId = smallestStakerId
		smallestStake, smallestStakerId, smallestStakerErr := cmdUtils.GetSmallestStakeAndId(client, epoch)
		if smallestStakerErr != nil {
			log.Error("Error in calculating smallest staker: ", smallestStakerErr)
			return smallestStakerErr
		}
		biggestStake = smallestStake
		biggestStakerId = smallestStakerId
		log.Debugf("Propose: In rogue mode, Biggest Stake: %s, Biggest Staker Id: %d", biggestStake, biggestStakerId)
	} else {
		biggestStake, biggestStakerId, biggestStakerErr = cmdUtils.GetBiggestStakeAndId(client, account.Address, epoch)
		if biggestStakerErr != nil {
			log.Error("Error in calculating biggest staker: ", biggestStakerErr)
			return biggestStakerErr
		}
	}

	log.Debugf("Getting Salt for current epoch %d...", epoch)
	salt, err := cmdUtils.GetSalt(client, epoch)
	if err != nil {
		return err
	}

	log.Debugf("Biggest staker Id: %d Biggest stake: %s, Stake: %s, Staker Id: %d, Number of Stakers: %d, Salt: %s", biggestStakerId, biggestStake, staker.Stake, staker.Id, numStakers, hex.EncodeToString(salt[:]))

	bufferPercent, err := cmdUtils.GetBufferPercent()
	if err != nil {
		return err
	}
	proposer := types.ElectedProposer{
		Stake:           staker.Stake,
		StakerId:        staker.Id,
		BiggestStake:    biggestStake,
		NumberOfStakers: numStakers,
		Salt:            salt,
		Epoch:           epoch,
	}
	log.Debugf("Propose: Calling GetIteration with arguments proposer = %+v, buffer percent = %d", proposer, bufferPercent)
	iteration := cmdUtils.GetIteration(client, proposer, bufferPercent)

	log.Debug("Iteration: ", iteration)

	if iteration == -1 {
		return nil
	}
	numOfProposedBlocks, err := razorUtils.GetNumberOfProposedBlocks(client, epoch)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Debug("Propose: Number of proposed blocks: ", numOfProposedBlocks)
	maxAltBlocks, err := razorUtils.GetMaxAltBlocks(client)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Debug("Propose: Number of maximum alternative blocks: ", maxAltBlocks)
	if numOfProposedBlocks >= maxAltBlocks {
		log.Debugf("Number of blocks proposed: %d, which is equal or greater than maximum alternative blocks allowed", numOfProposedBlocks)
		log.Debug("Comparing  iterations...")
		sortedProposedBlocks, err := razorUtils.GetSortedProposedBlockIds(client, epoch)
		if err != nil {
			log.Error("Error in fetching sorted proposed block ids")
			return err
		}
		log.Debug("Propose: Sorted proposed blocks: ", sortedProposedBlocks)
		lastBlockIndex := sortedProposedBlocks[numOfProposedBlocks-1]
		log.Debug("Propose: Last block index: ", lastBlockIndex)
		lastProposedBlockStruct, err := razorUtils.GetProposedBlock(client, epoch, lastBlockIndex)
		if err != nil {
			log.Error(err)
			return err
		}
		log.Debugf("Propose: Last proposed block: %+v", lastProposedBlockStruct)
		lastIteration := lastProposedBlockStruct.Iteration
		log.Debug("Propose: Iteration of last proposed block: ", lastIteration)
		if lastIteration.Cmp(big.NewInt(int64(iteration))) < 0 {
			log.Info("Current iteration is greater than iteration of last proposed block, cannot propose")
			return nil
		}
		log.Info("Current iteration is less than iteration of last proposed block, can propose")
	}
	log.Debugf("Propose: Calling MakeBlock() with arguments blockNumber = %s, epoch = %d, rogueData = %+v", blockNumber, epoch, rogueData)
	medians, ids, revealedDataMaps, err := cmdUtils.MakeBlock(client, blockNumber, epoch, rogueData)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Debugf("Propose: Medians: %d", medians)
	log.Debugf("Propose: Epoch: %d Medians: %d", epoch, medians)
	log.Debugf("Propose: Iteration: %d Biggest Staker Id: %d", iteration, biggestStakerId)
	log.Info("Proposing block...")

	txnOpts := razorUtils.GetTxnOpts(types.TransactionOptions{
		Client:          client,
		Password:        account.Password,
		AccountAddress:  account.Address,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.BlockManagerAddress,
		ABI:             bindings.BlockManagerMetaData.ABI,
		MethodName:      "propose",
		Parameters:      []interface{}{epoch, ids, medians, big.NewInt(int64(iteration)), biggestStakerId},
	})

	log.Debugf("Executing Propose transaction with epoch = %d, Ids = %v, medians = %s, iteration = %s, biggestStakerId = %d", epoch, ids, medians, big.NewInt(int64(iteration)), biggestStakerId)
	txn, err := blockManagerUtils.Propose(client, txnOpts, epoch, ids, medians, big.NewInt(int64(iteration)), biggestStakerId)
	if err != nil {
		log.Error(err)
		return err
	}
	proposeTxn := transactionUtils.Hash(txn)
	log.Info("Txn Hash: ", proposeTxn.Hex())
	if proposeTxn != core.NilHash {
		waitForBlockCompletionErr := razorUtils.WaitForBlockCompletion(client, proposeTxn.Hex())
		if waitForBlockCompletionErr != nil {
			log.Error("Error in WaitForBlockCompletionErr for propose: ", waitForBlockCompletionErr)
			return waitForBlockCompletionErr
		} else {
			// Saving proposed data after successful propose
			log.Debug("Updating global propose data struct...")
			updateGlobalProposedDataStruct(types.ProposeFileData{
				MediansData:           medians,
				RevealedDataMaps:      revealedDataMaps,
				RevealedCollectionIds: ids,
				Epoch:                 epoch,
			})
			log.Debugf("Propose: Global propose data struct: %+v", globalProposedDataStruct)

			log.Debug("Saving proposed data for recovery...")
			fileName, err := pathUtils.GetProposeDataFileName(account.Address)
			if err != nil {
				log.Error("Error in getting file name to save median data: ", err)
				return err
			}
			log.Debug("Propose: Propose data file path: ", fileName)
			err = fileUtils.SaveDataToProposeJsonFile(fileName, globalProposedDataStruct)
			if err != nil {
				log.Errorf("Error in saving data to file %s: %v", fileName, err)
				return err
			}
			log.Debug("Data saved!")
		}
	}

	return nil
}

//This function returns the biggest stake and Id of it
func (*UtilsStruct) GetBiggestStakeAndId(client *ethclient.Client, address string, epoch uint32) (*big.Int, uint32, error) {
	numberOfStakers, err := razorUtils.GetNumberOfStakers(client)
	if err != nil {
		return nil, 0, err
	}
	log.Debug("GetBiggestStakeAndId: Number of Stakers: ", numberOfStakers)
	if numberOfStakers == 0 {
		return nil, 0, errors.New("numberOfStakers is 0")
	}

	bufferPercent, err := cmdUtils.GetBufferPercent()
	if err != nil {
		return nil, 0, err
	}
	log.Debug("GetBiggestStakeAndId: Buffer Percent: ", bufferPercent)

	stateRemainingTime, err := razorUtils.GetRemainingTimeOfCurrentState(client, bufferPercent)
	if err != nil {
		return nil, 0, err
	}
	log.Debug("GetBiggestStakeAndId: State remaining time: ", stateRemainingTime)
	stateTimeout := time.NewTimer(time.Second * time.Duration(stateRemainingTime))

	var (
		biggestStake    = big.NewInt(0)
		biggestStakerId uint32
		mu              sync.Mutex
		wg              sync.WaitGroup
		errChan         = make(chan error, 1)
	)

	log.Debug("Iterating over all the stakers...")
	for i := 1; i <= int(numberOfStakers); i++ {
		wg.Add(1)
		go func(stakerId int) {
			defer wg.Done()

			select {
			case <-stateTimeout.C:
				errChan <- errors.New("state timeout error")
				return
			default:
				stake, err := razorUtils.GetStakeSnapshot(client, uint32(stakerId), epoch)
				if err != nil {
					errChan <- err
					return
				}

				mu.Lock()
				defer mu.Unlock()
				if stake.Cmp(biggestStake) > 0 {
					biggestStake = stake
					biggestStakerId = uint32(stakerId)
				}
			}
		}(i)
	}

	wg.Wait()

	select {
	case err := <-errChan:
		return nil, 0, err
	default:
	}

	if err != nil {
		return nil, 0, err
	}

	log.Debug("Propose: BiggestStake: ", biggestStake)
	log.Debug("Propose: Biggest Staker Id: ", biggestStakerId)
	return biggestStake, biggestStakerId, nil
}

//This function returns the iteration of the proposer if he is elected
func (*UtilsStruct) GetIteration(client *ethclient.Client, proposer types.ElectedProposer, bufferPercent int32) int {
	stake, err := razorUtils.GetStakeSnapshot(client, proposer.StakerId, proposer.Epoch)
	if err != nil {
		log.Error("Error in fetching influence of staker: ", err)
		return -1
	}
	log.Debug("GetIteration: Stake: ", stake)
	currentStakerStake := big.NewInt(1).Mul(stake, big.NewInt(int64(math.Exp2(32))))
	stateRemainingTime, err := razorUtils.GetRemainingTimeOfCurrentState(client, bufferPercent)
	if err != nil {
		return -1
	}
	stateTimeout := time.NewTimer(time.Second * time.Duration(stateRemainingTime))
	log.Debug("GetIteration: State remaining time: ", stateRemainingTime)
	log.Debug("Calculating Iteration...")
	log.Debugf("GetIteration: Calling IsElectedProposer() to find iteration...")
loop:
	for i := 0; i < 10000000; i++ {
		select {
		case <-stateTimeout.C:
			log.Error("State timeout!")
			break loop
		default:
			proposer.Iteration = i
			isElected := cmdUtils.IsElectedProposer(proposer, currentStakerStake)
			if isElected {
				return i
			}
		}
	}
	return -1
}

//This function returns if the elected staker is proposer or not
func (*UtilsStruct) IsElectedProposer(proposer types.ElectedProposer, currentStakerStake *big.Int) bool {
	seed := solsha3.SoliditySHA3([]string{"uint256"}, []interface{}{big.NewInt(int64(proposer.Iteration))})

	pseudoRandomNumber := pseudoRandomNumberGenerator(seed, proposer.NumberOfStakers, proposer.Salt[:])
	//add +1 since prng returns 0 to max-1 and staker start from 1
	pseudoRandomNumber = pseudoRandomNumber.Add(pseudoRandomNumber, big.NewInt(1))
	if pseudoRandomNumber.Cmp(big.NewInt(int64(proposer.StakerId))) != 0 {
		return false
	}
	seed2 := solsha3.SoliditySHA3([]string{"uint256", "uint256"}, []interface{}{big.NewInt(int64(proposer.StakerId)), big.NewInt(int64(proposer.Iteration))})
	log.Debug("IsElectedProposer: Seed: ", seed)
	log.Debug("IsElectedProposer: Pseudo random number: ", pseudoRandomNumber)
	log.Debug("IsElectedProposer: Seed 2: ", seed2)
	randomHash := solsha3.SoliditySHA3([]string{"bytes32", "bytes32"}, []interface{}{"0x" + hex.EncodeToString(proposer.Salt[:]), "0x" + hex.EncodeToString(seed2)})
	randomHashNumber := big.NewInt(0).SetBytes(randomHash)
	randomHashNumber = randomHashNumber.Mod(randomHashNumber, big.NewInt(int64(math.Exp2(32))))
	log.Debug("IsElectedProposer: Random hash number: ", randomHashNumber)
	biggestStake := big.NewInt(1).Mul(randomHashNumber, proposer.BiggestStake)
	log.Debug("IsElectedProposer: Biggest stake: ", biggestStake)
	isElected := biggestStake.Cmp(currentStakerStake) <= 0
	log.Debug("IsElectedProposer: Is staker elected: ", isElected)
	return isElected
}

//This function returns the pseudo random number
func pseudoRandomNumberGenerator(seed []byte, max uint32, blockHashes []byte) *big.Int {
	hash := solsha3.SoliditySHA3([]string{"bytes32", "bytes32"}, []interface{}{"0x" + hex.EncodeToString(blockHashes), "0x" + hex.EncodeToString(seed)})
	sum := big.NewInt(0).SetBytes(hash)
	return sum.Mod(sum, big.NewInt(int64(max)))
}

//This function returns the sorted revealed values
func (*UtilsStruct) GetSortedRevealedValues(client *ethclient.Client, blockNumber *big.Int, epoch uint32) (*types.RevealedDataMaps, error) {
	log.Debugf("GetSortedRevealedValues: Calling IndexRevealEventsOfCurrentEpoch with arguments blockNumber = %s, epoch = %d", blockNumber, epoch)
	assignedAsset, err := cmdUtils.IndexRevealEventsOfCurrentEpoch(client, blockNumber, epoch)
	if err != nil {
		return nil, err
	}
	log.Debugf("GetSortedRevealedValues: Revealed Data: %+v", assignedAsset)
	revealedValuesWithIndex := make(map[uint16][]*big.Int)
	voteWeights := make(map[string]*big.Int)
	influenceSum := make(map[uint16]*big.Int)
	log.Debug("Calculating sorted revealed values, vote weights and influence sum...")
	for _, asset := range assignedAsset {
		for _, assetValue := range asset.RevealedValues {
			if revealedValuesWithIndex[assetValue.LeafId] == nil {
				revealedValuesWithIndex[assetValue.LeafId] = []*big.Int{assetValue.Value}
			} else {
				if !utils.ContainsBigInteger(revealedValuesWithIndex[assetValue.LeafId], assetValue.Value) {
					revealedValuesWithIndex[assetValue.LeafId] = append(revealedValuesWithIndex[assetValue.LeafId], assetValue.Value)
				}
			}
			//Calculate vote weights
			if voteWeights[assetValue.Value.String()] == nil {
				voteWeights[assetValue.Value.String()] = big.NewInt(0)
			}
			voteWeights[assetValue.Value.String()] = big.NewInt(0).Add(voteWeights[assetValue.Value.String()], asset.Influence)

			//Calculate influence sum
			if influenceSum[assetValue.LeafId] == nil {
				influenceSum[assetValue.LeafId] = big.NewInt(0)
			}
			influenceSum[assetValue.LeafId] = big.NewInt(0).Add(influenceSum[assetValue.LeafId], asset.Influence)
		}
	}
	//sort revealed values
	for _, element := range revealedValuesWithIndex {
		sort.Slice(element, func(i, j int) bool {
			return element[i].Cmp(element[j]) == -1
		})
	}
	return &types.RevealedDataMaps{
		SortedRevealedValues: revealedValuesWithIndex,
		VoteWeights:          voteWeights,
		InfluenceSum:         influenceSum,
	}, nil
}

//This function returns the medians, idsRevealedInThisEpoch and revealedDataMaps
func (*UtilsStruct) MakeBlock(client *ethclient.Client, blockNumber *big.Int, epoch uint32, rogueData types.Rogue) ([]*big.Int, []uint16, *types.RevealedDataMaps, error) {
	log.Debugf("MakeBlock: Calling GetSortedRevealedValues with arguments blockNumber = %s, epoch = %d", blockNumber, epoch)
	revealedDataMaps, err := cmdUtils.GetSortedRevealedValues(client, blockNumber, epoch)
	if err != nil {
		return nil, nil, nil, err
	}
	log.Debugf("MakeBlock: Revealed data map: %+v", revealedDataMaps)

	activeCollections, err := razorUtils.GetActiveCollectionIds(client)
	if err != nil {
		return nil, nil, nil, err
	}
	log.Debug("MakeBlock: Active collections: ", activeCollections)

	var (
		medians                []*big.Int
		idsRevealedInThisEpoch []uint16
	)

	log.Debug("Iterating over all the active collections for medians calculation....")
	for leafId := uint16(0); leafId < uint16(len(activeCollections)); leafId++ {
		influenceSum := revealedDataMaps.InfluenceSum[leafId]
		if influenceSum != nil && influenceSum.Cmp(big.NewInt(0)) != 0 {
			idsRevealedInThisEpoch = append(idsRevealedInThisEpoch, activeCollections[leafId])
			if rogueData.IsRogue && utils.Contains(rogueData.RogueMode, "medians") {
				medians = append(medians, razorUtils.GetRogueRandomValue(10000000))
				continue
			}
			accWeight := big.NewInt(0)
			for i := 0; i < len(revealedDataMaps.SortedRevealedValues[leafId]); i++ {
				revealedValue := revealedDataMaps.SortedRevealedValues[leafId][i]
				accWeight = accWeight.Add(accWeight, revealedDataMaps.VoteWeights[revealedValue.String()])
				if accWeight.Cmp(influenceSum.Div(influenceSum, big.NewInt(2))) > 0 {
					medians = append(medians, revealedValue)
					break
				}
			}
		}
	}
	if rogueData.IsRogue && utils.Contains(rogueData.RogueMode, "missingIds") {
		log.Warn("YOU ARE PROPOSING IDS REVEALED IN ROGUE MODE, THIS CAN INCUR PENALTIES!")
		//Replacing the last ID: id with id+1 in idsRevealed array if rogueMode == missingIds
		idsRevealedInThisEpoch[len(idsRevealedInThisEpoch)-1] = idsRevealedInThisEpoch[len(idsRevealedInThisEpoch)-1] + 1
	}
	if rogueData.IsRogue && utils.Contains(rogueData.RogueMode, "extraIds") {
		log.Warn("YOU ARE PROPOSING IDS REVEALED IN ROGUE MODE, THIS CAN INCUR PENALTIES!")
		//Adding a dummy median and appending extra id to idsRevealed array if rogueMode == extraIds
		medians = append(medians, razorUtils.GetRogueRandomValue(10000000))
		idsRevealedInThisEpoch = append(idsRevealedInThisEpoch, idsRevealedInThisEpoch[len(idsRevealedInThisEpoch)-1]+1)
	}
	if rogueData.IsRogue && utils.Contains(rogueData.RogueMode, "unsortedIds") && len(idsRevealedInThisEpoch) > 1 {
		log.Warn("YOU ARE PROPOSING IDS REVEALED IN ROGUE MODE, THIS CAN INCUR PENALTIES!")
		//Interchanging the first 2 elements of idsRevealed array
		temp := idsRevealedInThisEpoch[0]
		idsRevealedInThisEpoch[0] = idsRevealedInThisEpoch[1]
		idsRevealedInThisEpoch[1] = temp
	}
	return medians, idsRevealedInThisEpoch, revealedDataMaps, nil
}

func (*UtilsStruct) GetSmallestStakeAndId(client *ethclient.Client, epoch uint32) (*big.Int, uint32, error) {
	numberOfStakers, err := razorUtils.GetNumberOfStakers(client)
	if err != nil {
		return nil, 0, err
	}
	if numberOfStakers == 0 {
		return nil, 0, errors.New("numberOfStakers is 0")
	}
	var smallestStakerId uint32
	smallestStake := big.NewInt(1).Mul(big.NewInt(1e18), big.NewInt(1e18))

	for i := 1; i <= int(numberOfStakers); i++ {
		stake, err := razorUtils.GetStakeSnapshot(client, uint32(i), epoch)
		if err != nil {
			return nil, 0, err
		}
		if stake.Cmp(smallestStake) < 0 {
			smallestStake = stake
			smallestStakerId = uint32(i)
		}
	}
	return smallestStake, smallestStakerId, nil
}

func updateGlobalProposedDataStruct(proposedData types.ProposeFileData) types.ProposeFileData {
	globalProposedDataStruct.MediansData = proposedData.MediansData
	globalProposedDataStruct.RevealedDataMaps = proposedData.RevealedDataMaps
	globalProposedDataStruct.RevealedCollectionIds = proposedData.RevealedCollectionIds
	globalProposedDataStruct.Epoch = proposedData.Epoch
	return globalProposedDataStruct
}
