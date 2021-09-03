package utils

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"razor/pkg/bindings"
)

func getVoteManagerWithOpts(client *ethclient.Client, address string) (*bindings.VoteManager, bind.CallOpts) {
	return GetVoteManager(client), GetOptions(false, address, "")
}

func GetCommitments(client *ethclient.Client, address string, epoch uint32) ([32]byte, error) {
	voteManager, callOpts := getVoteManagerWithOpts(client, address)
	commitments, err := voteManager.Commitments(&callOpts, epoch)
	if err != nil {
		return [32]byte{}, err
	}
	return commitments.CommitmentHash, nil
}

func GetVotes(client *ethclient.Client, address string, stakerId uint32) (struct {
	Epoch uint32
	Values []*big.Int
}, error) {
	voteManager, callOpts := getVoteManagerWithOpts(client, address)
	return voteManager.GetVote(&callOpts, stakerId)
}

func GetInfluenceSnapshot(client *ethclient.Client, address string, epoch uint32) (*big.Int, error) {
	voteManager, callOpts := getVoteManagerWithOpts(client, address)
	stakerId, err := GetStakerId(client, address)
	if err != nil {
		return nil, err
	}
	return voteManager.GetInfluenceSnapshot(&callOpts, epoch, stakerId)
}

func GetTotalInfluenceRevealed(client *ethclient.Client, address string, epoch uint32) (*big.Int, error) {
	voteManager, callOpts := getVoteManagerWithOpts(client, address)
	return voteManager.GetTotalInfluenceRevealed(&callOpts, epoch)
}

func GetRandaoHash(client *ethclient.Client, address string) ([32]byte, error) {
	voteManager, callOpts := getVoteManagerWithOpts(client, address)
	return voteManager.GetRandaoHash(&callOpts)
}

func GetEpochLastCommitted(client *ethclient.Client, address string, stakerId uint32) (uint32, error) {
	voteManager, callOpts := getVoteManagerWithOpts(client, address)
	return voteManager.GetEpochLastCommitted(&callOpts, stakerId)
}

func GetEpochLastRevealed(client *ethclient.Client, address string, stakerId uint32) (uint32, error) {
	voteManager, callOpts := getVoteManagerWithOpts(client, address)
	return voteManager.GetEpochLastRevealed(&callOpts, stakerId)
}