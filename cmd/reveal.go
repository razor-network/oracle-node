package cmd

import (
	"errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"
	"strings"
)

func (*UtilsStruct) HandleRevealState(client *ethclient.Client, staker bindings.StructsStaker, epoch uint32) error {
	epochLastCommitted, err := razorUtils.GetEpochLastCommitted(client, staker.Id)
	if err != nil {
		return err
	}
	log.Debug("Staker last epoch committed: ", epochLastCommitted)
	if epochLastCommitted != epoch {
		return errors.New("commitment for this epoch not found on network.... aborting reveal")
	}
	return nil
}

func (*UtilsStruct) Reveal(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, commitData types.CommitData, secret []byte) (common.Hash, error) {
	if state, err := razorUtils.GetDelayedState(client, config.BufferPercent); err != nil || state != 1 {
		log.Error("Not reveal state")
		return core.NilHash, err
	}

	merkleTree := utils.MerkleInterface.CreateMerkle(commitData.Leaves)
	treeRevealData := cmdUtils.GenerateTreeRevealData(merkleTree, commitData)
	secretBytes32 := [32]byte{}
	copy(secretBytes32[:], secret)

	log.Debugf("Revealing vote for epoch: %d secret: %s  commitAccount: %s, treeRevealData: %t",
		epoch,
		"0x"+common.Bytes2Hex(secret),
		account.Address,
		treeRevealData)
	log.Info("Revealing votes...")

	txnOpts := razorUtils.GetTxnOpts(types.TransactionOptions{
		Client:          client,
		Password:        account.Password,
		AccountAddress:  account.Address,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.VoteManagerAddress,
		ABI:             bindings.VoteManagerABI,
		MethodName:      "reveal",
		Parameters:      []interface{}{epoch, treeRevealData, secretBytes32},
	})
	txn, err := voteManagerUtils.Reveal(client, txnOpts, epoch, treeRevealData, secretBytes32)
	if err != nil {
		log.Error(err)
		return core.NilHash, err
	}
	log.Info("Txn Hash: ", transactionUtils.Hash(txn))
	return transactionUtils.Hash(txn), nil
}

func (*UtilsStruct) GenerateTreeRevealData(merkleTree [][][]byte, commitData types.CommitData) bindings.StructsMerkleTree {
	if merkleTree == nil || commitData.SeqAllottedCollections == nil || commitData.Leaves == nil {
		log.Error("No data present for construction of StructsMerkleTree")
		return bindings.StructsMerkleTree{}
	}
	var (
		values []bindings.StructsAssignedAsset
		proofs [][][32]byte
	)

	for i := 0; i < len(commitData.SeqAllottedCollections); i++ {
		value := bindings.StructsAssignedAsset{
			LeafId: uint16(commitData.SeqAllottedCollections[i].Uint64()),
			Value:  uint32(commitData.Leaves[commitData.SeqAllottedCollections[i].Uint64()].Uint64()),
		}
		proof := utils.MerkleInterface.GetProofPath(merkleTree, value.LeafId)
		values = append(values, value)
		proofs = append(proofs, proof)
	}

	return bindings.StructsMerkleTree{
		Values: values,
		Proofs: proofs,
		Root:   utils.MerkleInterface.GetMerkleRoot(merkleTree),
	}
}

func (*UtilsStruct) IndexRevealEventsOfCurrentEpoch(client *ethclient.Client, blockNumber *big.Int, epoch uint32) ([]types.RevealedStruct, error) {
	numberOfBlocks := int64(core.StateLength) * core.NumberOfStates
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0).Sub(blockNumber, big.NewInt(numberOfBlocks)),
		ToBlock:   blockNumber,
		Addresses: []common.Address{
			common.HexToAddress(core.VoteManagerAddress),
		},
	}
	logs, err := utils.UtilsInterface.FilterLogsWithRetry(client, query)
	if err != nil {
		return nil, err
	}
	contractAbi, err := utils.ABIInterface.Parse(strings.NewReader(bindings.VoteManagerABI))
	if err != nil {
		return nil, err
	}
	var revealedData []types.RevealedStruct
	for _, vLog := range logs {
		data, unpackErr := abiUtils.Unpack(contractAbi, "Revealed", vLog.Data)
		if unpackErr != nil {
			log.Error(unpackErr)
			continue
		}
		if epoch == data[0].(uint32) {
			treeValues := data[3].([]struct {
				LeafId uint16 `json:"leafId"`
				Value  uint32 `json:"value"`
			})
			var revealedValues []types.AssignedAsset
			for _, value := range treeValues {
				revealedValues = append(revealedValues, types.AssignedAsset{
					LeafId: value.LeafId,
					Value:  value.Value,
				})
			}
			consolidatedRevealedData := types.RevealedStruct{
				RevealedValues: revealedValues,
				Influence:      data[2].(*big.Int),
			}
			revealedData = append(revealedData, consolidatedRevealedData)
		}
	}
	return revealedData, nil
}
