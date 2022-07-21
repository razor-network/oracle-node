//Package cmd provides all functions related to command line
package cmd

import (
	"errors"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

//This function handles the reveal state
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

//This function checks if the state is reveal or not and then reveals the votes
func (*UtilsStruct) Reveal(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, commitData types.CommitData, signature []byte) (common.Hash, error) {
	if state, err := razorUtils.GetDelayedState(client, config.BufferPercent); err != nil || state != 1 {
		log.Error("Not reveal state")
		return core.NilHash, err
	}

	merkleTree := utils.MerkleInterface.CreateMerkle(commitData.Leaves)
	treeRevealData := cmdUtils.GenerateTreeRevealData(merkleTree, commitData)

	log.Debugf("Revealing vote for epoch: %d, commitAccount: %s, treeRevealData: %v, root: %v",
		epoch,
		account.Address,
		treeRevealData.Values,
		common.Bytes2Hex(treeRevealData.Root[:]),
	)

	log.Info("Revealing votes...")

	txnOpts := razorUtils.GetTxnOpts(types.TransactionOptions{
		Client:          client,
		Password:        account.Password,
		AccountAddress:  account.Address,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.VoteManagerAddress,
		ABI:             bindings.VoteManagerMetaData.ABI,
		MethodName:      "reveal",
		Parameters:      []interface{}{epoch, treeRevealData, signature},
	})
	txn, err := voteManagerUtils.Reveal(client, txnOpts, epoch, treeRevealData, signature)
	txnHash := transactionUtils.Hash(txn)
	if err != nil {
		log.Error(err)
		return core.NilHash, err
	}
	log.Info("Txn Hash: ", txnHash.Hex())
	return txnHash, nil
}

//This function generates the tree reveal data
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
			Value:  big.NewInt(commitData.Leaves[commitData.SeqAllottedCollections[i].Uint64()].Int64()),
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

//This function indexes the reveal events of current epoch
func (*UtilsStruct) IndexRevealEventsOfCurrentEpoch(client *ethclient.Client, blockNumber *big.Int, epoch uint32) ([]types.RevealedStruct, error) {
	fromBlock, err := utils.UtilsInterface.CalculateBlockNumberAtEpochBeginning(client, core.EpochLength, blockNumber)
	if err != nil {
		return nil, errors.New("Not able to Fetch Block: " + err.Error())
	}
	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   blockNumber,
		Addresses: []common.Address{
			common.HexToAddress(core.VoteManagerAddress),
		},
	}
	logs, err := utils.UtilsInterface.FilterLogsWithRetry(client, query)
	if err != nil {
		return nil, err
	}
	contractAbi, err := utils.ABIInterface.Parse(strings.NewReader(bindings.VoteManagerMetaData.ABI))
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
			treeValues := data[2].([]struct {
				LeafId uint16   `json:"leafId"`
				Value  *big.Int `json:"value"`
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
				Influence:      data[1].(*big.Int),
			}
			revealedData = append(revealedData, consolidatedRevealedData)
		}
	}
	log.Debug("Revealed values: ", revealedData)
	return revealedData, nil
}
