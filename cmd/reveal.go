package cmd

import (
	"encoding/hex"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"github.com/wealdtech/go-merkletree"
	"math/big"
	"razor/core/types"
	"razor/utils"
)

func HandleRevealState(client *ethclient.Client, address string, stakerId *big.Int, epoch *big.Int) error {
	staker, err := utils.GetStaker(client, address, stakerId)
	if err != nil {
		log.Error("Error in fetching staker")
		return err
	}
	log.Info("Staker last epoch committed: ", staker.EpochLastCommitted)
	if staker.EpochLastCommitted.Cmp(epoch) != 0 {
		return errors.New("commitment for this epoch not found on network.... aborting reveal")
	}
	log.Info("Staker last epoch revealed: ", staker.EpochLastRevealed)
	return nil
}

func Reveal(client *ethclient.Client, committedData []*big.Int, secret []byte, account types.Account, commitAccount string, config types.Configurations) {
	if state, err := utils.GetDelayedState(client); err != nil || state != 1 {
		log.Error("Not reveal state")
		return
	}

	epoch, err := utils.GetEpoch(client, account.Address)
	if err != nil {
		log.Error(err)
		return
	}
	commitments, err := utils.GetCommitments(client, account.Address, epoch)
	if err != nil {
		log.Error(err)
		return
	}
	if utils.AllZero(commitments) {
		log.Error("Did not commit")
		return
	}

	// TODO Check if already revealed
	revealed, err := utils.GetVotes(client, account.Address, epoch)
	if err != nil {
		log.Error(err)
		return
	}
	log.Info(revealed)

	tree, err := utils.GetMerkleTree(committedData)
	if err != nil {
		log.Error(err)
		return
	}

	proofs := getProofs(tree, committedData)

	txnOpts := utils.GetTxnOpts(types.TransactionOptions{
		Client:         client,
		Password:       account.Password,
		AccountAddress: account.Address,
		ChainId:        config.ChainId,
		GasMultiplier:  config.GasMultiplier,
	})

	voteManager := utils.GetVoteManager(client)

	root := [32]byte{}
	copy(root[:], tree.Root())

	secretBytes32 := [32]byte{}
	copy(secretBytes32[:], secret)
	log.Infof("Revealing vote for epoch: %s  votes: %s  root: %s  proof: %s  secret: %s  commitAccount: %s",
		epoch,
		committedData,
		"0x"+common.Bytes2Hex(tree.Root()),
		proofs,
		"0x"+common.Bytes2Hex(secret),
		commitAccount,
	)
	txn, err := voteManager.Reveal(txnOpts, epoch, root, committedData, proofs, secretBytes32, common.HexToAddress(commitAccount))
	if err != nil {
		log.Error(err)
		return
	}
	utils.WaitForBlockCompletion(client, txn.Hash().String())
}

func getProofs(tree *merkletree.MerkleTree, data []*big.Int) [][][32]byte {
	var proofs []*merkletree.Proof
	for _, datum := range data {
		proof, err := tree.GenerateProof(datum.Bytes())
		if err != nil {
			log.Error("Error in calculating merkle proof: ", err)
			continue
		}
		proofs = append(proofs, proof)
	}
	var finalProofs [][][32]byte
	for _, proof := range proofs {
		var proofHash [][32]byte
		for _, nestedProof := range proof.Hashes {
			if nestedProof != nil {
				nestedProofBytes32 := [32]byte{}
				copy(nestedProofBytes32[:], nestedProof)
				log.Info("Proof: ", hex.EncodeToString(nestedProof))
				proofHash = append(proofHash, nestedProofBytes32)
			}
		}
		finalProofs = append(finalProofs, proofHash)
	}
	return finalProofs
}
