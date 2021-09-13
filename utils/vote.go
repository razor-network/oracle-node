package utils

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
)

func getVoteManagerWithOpts(client *ethclient.Client, address string) (*bindings.VoteManager, bind.CallOpts) {
	return GetVoteManager(client), GetOptions(false, address, "")
}

func GetCommitments(client *ethclient.Client, address string) ([32]byte, error) {
	voteManager, callOpts := getVoteManagerWithOpts(client, address)
	stakerId, err := GetStakerId(client, address)
	if err != nil {
		return [32]byte{}, err
	}
	var (
		commitments   types.Commitment
		commitmentErr error
	)
	for retry := 1; retry <= core.MaxRetries; retry++ {
		commitments, commitmentErr = voteManager.Commitments(&callOpts, stakerId)
		if commitmentErr != nil {
			Retry(retry, "Error in fetching commitment: ", commitmentErr)
			continue
		}
		break
	}
	if commitmentErr != nil {
		return [32]byte{}, err
	}
	return commitments.CommitmentHash, nil
}

func GetVoteValue(client *ethclient.Client, address string, assetId uint8, stakerId uint32) (*big.Int, error) {
	voteManager, callOpts := getVoteManagerWithOpts(client, address)
	var (
		voteValue    *big.Int
		voteValueErr error
	)
	for retry := 1; retry <= core.MaxRetries; retry++ {
		voteValue, voteValueErr = voteManager.GetVoteValue(&callOpts, assetId, stakerId)
		if voteValueErr != nil {
			Retry(retry, "Error in fetching last vote: ", voteValueErr)
			continue
		}
		break
	}
	if voteValueErr != nil {
		return nil, voteValueErr
	}
	return voteValue, nil
}

func GetVoteValue(client *ethclient.Client, address string, assetId uint8, stakerId uint32) (*big.Int, error) {
	voteManager, callOpts := getVoteManagerWithOpts(client, address)
	return voteManager.GetVoteValue(&callOpts, assetId, stakerId)
}

func GetInfluenceSnapshot(client *ethclient.Client, address string, epoch uint32) (*big.Int, error) {
	voteManager, callOpts := getVoteManagerWithOpts(client, address)
	stakerId, err := GetStakerId(client, address)
	if err != nil {
		return nil, err
	}
	var (
		influenceSnapshot *big.Int
		influenceErr      error
	)
	for retry := 1; retry <= core.MaxRetries; retry++ {
		influenceSnapshot, influenceErr = voteManager.GetInfluenceSnapshot(&callOpts, epoch, stakerId)
		if influenceErr != nil {
			Retry(retry, "Error in fetching influence snapshot: ", influenceErr)
			continue
		}
		break
	}
	if influenceErr != nil {
		return nil, influenceErr
	}
	return influenceSnapshot, nil
}

func GetTotalInfluenceRevealed(client *ethclient.Client, address string, epoch uint32) (*big.Int, error) {
	voteManager, callOpts := getVoteManagerWithOpts(client, address)
	var (
		totalInfluenceRevealed *big.Int
		influenceErr           error
	)
	for retry := 1; retry <= core.MaxRetries; retry++ {
		totalInfluenceRevealed, influenceErr = voteManager.GetTotalInfluenceRevealed(&callOpts, epoch)
		if influenceErr != nil {
			Retry(retry, "Error in fetching influence snapshot: ", influenceErr)
			continue
		}
		break
	}
	if influenceErr != nil {
		return nil, influenceErr
	}
	return totalInfluenceRevealed, nil
}

func GetRandaoHash(client *ethclient.Client, address string) ([32]byte, error) {
	voteManager, callOpts := getVoteManagerWithOpts(client, address)
	var (
		randaoHash [32]byte
		randaoErr  error
	)
	for retry := 1; retry <= core.MaxRetries; retry++ {
		randaoHash, randaoErr = voteManager.GetRandaoHash(&callOpts)
		if randaoErr != nil {
			Retry(retry, "Error in fetching randao hash: ", randaoErr)
			continue
		}
		break
	}
	if randaoErr != nil {
		return [32]byte{}, randaoErr
	}
	return randaoHash, nil
}

func GetEpochLastCommitted(client *ethclient.Client, address string, stakerId uint32) (uint32, error) {
	voteManager, callOpts := getVoteManagerWithOpts(client, address)
	var (
		epochLastCommitted uint32
		err                error
	)
	for retry := 1; retry <= core.MaxRetries; retry++ {
		epochLastCommitted, err = voteManager.GetEpochLastCommitted(&callOpts, stakerId)
		if err != nil {
			Retry(retry, "Error in fetching epoch last committed: ", err)
			continue
		}
		break
	}
	if err != nil {
		return 0, err
	}
	return epochLastCommitted, nil
}

func GetEpochLastRevealed(client *ethclient.Client, address string, stakerId uint32) (uint32, error) {
	voteManager, callOpts := getVoteManagerWithOpts(client, address)
	var (
		epochLastRevealed uint32
		err               error
	)
	for retry := 1; retry <= core.MaxRetries; retry++ {
		epochLastRevealed, err = voteManager.GetEpochLastRevealed(&callOpts, stakerId)
		if err != nil {
			Retry(retry, "Error in fetching epoch last revealed: ", err)
			continue
		}
		break
	}
	if err != nil {
		return 0, err
	}
	return epochLastRevealed, nil
}
