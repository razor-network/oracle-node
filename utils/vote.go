package utils

import (
	"github.com/avast/retry-go"
	"math/big"
	"razor/core/types"
	"razor/pkg/bindings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getVoteManagerWithOpts(client *ethclient.Client) (*bindings.VoteManager, bind.CallOpts) {
	return GetVoteManager(client), UtilsInterface.GetOptions()
}

func GetCommitments(client *ethclient.Client, address string) ([32]byte, error) {
	voteManager, callOpts := getVoteManagerWithOpts(client)
	stakerId, err := GetStakerId(client, address)
	if err != nil {
		return [32]byte{}, err
	}
	var (
		commitments   types.Commitment
		commitmentErr error
	)
	commitmentErr = retry.Do(
		func() error {
			commitments, commitmentErr = voteManager.Commitments(&callOpts, stakerId)
			if commitmentErr != nil {
				log.Error("Error in fetching commitment....Retrying")
				return commitmentErr
			}
			return nil
		})
	if commitmentErr != nil {
		return [32]byte{}, err
	}
	return commitments.CommitmentHash, nil
}

func GetVoteValue(client *ethclient.Client, assetId uint16, stakerId uint32) (*big.Int, error) {
	voteManager, callOpts := getVoteManagerWithOpts(client)
	var (
		voteValue    *big.Int
		voteValueErr error
	)
	voteValueErr = retry.Do(
		func() error {
			voteValue, voteValueErr = voteManager.GetVoteValue(&callOpts, assetId-1, stakerId)
			if voteValueErr != nil {
				log.Error("Error in fetching last vote value....Retrying")
				return voteValueErr
			}
			return nil
		})
	if voteValueErr != nil {
		return nil, voteValueErr
	}
	return voteValue, nil
}

func GetVotes(client *ethclient.Client, stakerId uint32) (bindings.StructsVote, error) {
	voteManager, callOpts := getVoteManagerWithOpts(client)
	var (
		votes    bindings.StructsVote
		votesErr error
	)
	votesErr = retry.Do(
		func() error {
			votes, votesErr = voteManager.GetVote(&callOpts, stakerId)
			if votesErr != nil {
				log.Error("Error in fetching last vote value....Retrying")
				return votesErr
			}
			return nil
		})
	if votesErr != nil {
		return bindings.StructsVote{}, votesErr
	}
	return votes, nil
}

func GetInfluenceSnapshot(client *ethclient.Client, stakerId uint32, epoch uint32) (*big.Int, error) {
	voteManager, callOpts := getVoteManagerWithOpts(client)
	var (
		influenceSnapshot *big.Int
		influenceErr      error
	)
	influenceErr = retry.Do(
		func() error {
			influenceSnapshot, influenceErr = voteManager.GetInfluenceSnapshot(&callOpts, epoch, stakerId)
			if influenceErr != nil {
				log.Error("Error in fetching influence snapshot....Retrying")
				return influenceErr
			}
			return nil
		})
	if influenceErr != nil {
		return nil, influenceErr
	}
	return influenceSnapshot, nil
}

func GetTotalInfluenceRevealed(client *ethclient.Client, epoch uint32) (*big.Int, error) {
	voteManager, callOpts := getVoteManagerWithOpts(client)
	var (
		totalInfluenceRevealed *big.Int
		influenceErr           error
	)
	influenceErr = retry.Do(
		func() error {
			totalInfluenceRevealed, influenceErr = voteManager.GetTotalInfluenceRevealed(&callOpts, epoch)
			if influenceErr != nil {
				log.Error("Error in fetching total influence revealed....Retrying")
				return influenceErr
			}
			return nil
		})
	if influenceErr != nil {
		return nil, influenceErr
	}
	return totalInfluenceRevealed, nil
}

func GetRandaoHash(client *ethclient.Client) ([32]byte, error) {
	voteManager, callOpts := getVoteManagerWithOpts(client)
	var (
		randaoHash [32]byte
		randaoErr  error
	)
	randaoErr = retry.Do(
		func() error {
			randaoHash, randaoErr = voteManager.GetRandaoHash(&callOpts)
			if randaoErr != nil {
				log.Error("Error in fetching randao hash.....Retrying")
				return randaoErr
			}
			return nil
		})
	if randaoErr != nil {
		return [32]byte{}, randaoErr
	}
	return randaoHash, nil
}

func GetEpochLastCommitted(client *ethclient.Client, stakerId uint32) (uint32, error) {
	voteManager, callOpts := getVoteManagerWithOpts(client)
	var (
		epochLastCommitted uint32
		err                error
	)
	err = retry.Do(
		func() error {
			epochLastCommitted, err = voteManager.GetEpochLastCommitted(&callOpts, stakerId)
			if err != nil {
				log.Error("Error in fetching epoch last committed....Retrying")
				return err
			}
			return nil
		})
	if err != nil {
		return 0, err
	}
	return epochLastCommitted, nil
}

func GetEpochLastRevealed(client *ethclient.Client, address string, stakerId uint32) (uint32, error) {
	voteManager, callOpts := getVoteManagerWithOpts(client)
	var (
		epochLastRevealed uint32
		err               error
	)
	err = retry.Do(
		func() error {
			epochLastRevealed, err = voteManager.GetEpochLastRevealed(&callOpts, stakerId)
			if err != nil {
				log.Error("Error in fetching epoch last revealed....Retrying")
				return err
			}
			return nil
		})
	if err != nil {
		return 0, err
	}
	return epochLastRevealed, nil
}
