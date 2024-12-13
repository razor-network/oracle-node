//Package cmd provides all functions related to command line
package cmd

import (
	"errors"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/rpc"
	"razor/utils"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
)

//This function checks for epoch last committed
func (*UtilsStruct) CheckForLastCommitted(rpcParameters rpc.RPCParameters, staker bindings.StructsStaker, epoch uint32) error {
	epochLastCommitted, err := razorUtils.GetEpochLastCommitted(rpcParameters, staker.Id)
	if err != nil {
		return err
	}
	log.Debug("CheckForLastCommitted: Staker last epoch committed: ", epochLastCommitted)
	if epochLastCommitted != epoch {
		return errors.New("commitment for this epoch not found on network.... aborting reveal")
	}
	return nil
}

//This function checks if the state is reveal or not and then reveals the votes
func (*UtilsStruct) Reveal(rpcParameters rpc.RPCParameters, config types.Configurations, account types.Account, epoch uint32, latestHeader *Types.Header, stateBuffer uint64, commitData types.CommitData, signature []byte) (common.Hash, error) {
	if state, err := razorUtils.GetBufferedState(latestHeader, stateBuffer, config.BufferPercent); err != nil || state != 1 {
		log.Error("Not reveal state")
		return core.NilHash, err
	}

	log.Debug("Creating merkle tree...")
	merkleTree, err := merkleUtils.CreateMerkle(commitData.Leaves)

	if err != nil {
		log.Error("Error in getting merkle tree: ", err)
		return core.NilHash, err
	}
	log.Debug("Generating tree reveal data...")
	treeRevealData := cmdUtils.GenerateTreeRevealData(merkleTree, commitData)

	log.Debugf("Revealing vote for epoch: %d, commitAccount: %s, treeRevealData: %v, root: %v",
		epoch,
		account.Address,
		treeRevealData.Values,
		common.Bytes2Hex(treeRevealData.Root[:]),
	)

	log.Info("Revealing votes...")

	txnOpts, err := razorUtils.GetTxnOpts(rpcParameters, types.TransactionOptions{
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.VoteManagerAddress,
		ABI:             bindings.VoteManagerMetaData.ABI,
		MethodName:      "reveal",
		Parameters:      []interface{}{epoch, treeRevealData, signature},
		Account:         account,
	})
	if err != nil {
		log.Error(err)
		return core.NilHash, err
	}

	client, err := rpcParameters.RPCManager.GetBestRPCClient()
	if err != nil {
		log.Error(err)
		return core.NilHash, err
	}

	log.Debugf("Executing Reveal transaction wih epoch = %d, treeRevealData = %v, signature = %v", epoch, treeRevealData, signature)
	txn, err := voteManagerUtils.Reveal(client, txnOpts, epoch, treeRevealData, signature)
	if err != nil {
		log.Error(err)
		return core.NilHash, err
	}
	txnHash := transactionUtils.Hash(txn)
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
		proof := merkleUtils.GetProofPath(merkleTree, value.LeafId)
		values = append(values, value)
		proofs = append(proofs, proof)
	}

	root, err := merkleUtils.GetMerkleRoot(merkleTree)
	if err != nil {
		log.Error("Error in getting root: ", err)
		return bindings.StructsMerkleTree{}
	}
	log.Debugf("GenerateTreeRevealData: values = %+v, proofs = %+v, root = %v", values, proofs, root)

	return bindings.StructsMerkleTree{
		Values: values,
		Proofs: proofs,
		Root:   root,
	}
}

//This function indexes the reveal events of current epoch
func (*UtilsStruct) IndexRevealEventsOfCurrentEpoch(rpcParameters rpc.RPCParameters, blockNumber *big.Int, epoch uint32) ([]types.RevealedStruct, error) {
	log.Debug("Fetching reveal events of current epoch...")
	fromBlock, err := razorUtils.EstimateBlockNumberAtEpochBeginning(rpcParameters, blockNumber)
	if err != nil {
		return nil, errors.New("Not able to Fetch Block: " + err.Error())
	}
	log.Debugf("IndexRevealEventsOfCurrentEpoch: Checking for events from block: %s to block: %s", fromBlock, blockNumber)
	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   blockNumber,
		Addresses: []common.Address{
			common.HexToAddress(core.VoteManagerAddress),
		},
	}
	log.Debugf("IndexRevealEventsOfCurrentEpoch: Query to send in filter logs: %+v", query)
	logs, err := clientUtils.FilterLogsWithRetry(rpcParameters, query)
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
			log.Debug(unpackErr)
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
	log.Debug("IndexRevealEventsOfCurrentEpoch: Revealed values: ", revealedData)
	return revealedData, nil
}
