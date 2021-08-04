package utils

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func GetCommitments(client *ethclient.Client, address string, epoch *big.Int) ([32]byte, error) {
	voteManager := GetVoteManager(client)
	callOpts := GetOptions(false, address, "")
	stakerId, err := GetStakerId(client, address)
	if err != nil {
		return [32]byte{}, err
	}
	return voteManager.Commitments(&callOpts, epoch, stakerId)
}

func GetVotes(client *ethclient.Client, address string, epoch *big.Int, stakerId *big.Int, assetId *big.Int) (struct {
	Value  *big.Int
	Weight *big.Int
}, error) {
	voteManager := GetVoteManager(client)
	callOpts := GetOptions(false, address, "")
	if assetId == nil {
		assetId = big.NewInt(0)
	}

	return voteManager.GetVote(&callOpts, epoch, stakerId, assetId)
}

func GetVoteWeights(client *ethclient.Client, address string, epoch *big.Int, assetId *big.Int, voteValue *big.Int) (*big.Int, error) {
	voteManager := GetVoteManager(client)
	callOpts := GetOptions(false, address, "")
	if assetId == nil {
		assetId = big.NewInt(0)
	}
	return voteManager.VoteWeights(&callOpts, epoch, assetId, voteValue)
}
