//Package cmd provides all functions related to command line
package cmd

import (
	"razor/core"
	"razor/core/types"
	"razor/rpc"

	"github.com/ethereum/go-ethereum/common"
)

//This function allows the user to claim the block reward and returns the hash
func (*UtilsStruct) ClaimBlockReward(rpcParameters rpc.RPCParameters, options types.TransactionOptions) (common.Hash, error) {
	epoch, err := razorUtils.GetEpoch(rpcParameters)
	if err != nil {
		log.Error("Error in getting epoch: ", err)
		return core.NilHash, err
	}
	log.Debug("ClaimBlockReward: Epoch: ", epoch)

	sortedProposedBlockIds, err := razorUtils.GetSortedProposedBlockIds(rpcParameters, epoch)
	if err != nil {
		log.Error("Error in getting sortedProposedBlockIds: ", err)
		return core.NilHash, err
	}
	log.Debug("ClaimBlockReward: Sorted proposed block Ids: ", sortedProposedBlockIds)

	if sortedProposedBlockIds == nil {
		log.Debug("No blocks proposed in this epoch")
		return core.NilHash, nil
	}

	stakerID, err := razorUtils.GetStakerId(rpcParameters, options.Account.Address)
	if err != nil {
		log.Error("Error in getting stakerId: ", err)
		return core.NilHash, err
	}
	log.Debug("ClaimBlockReward: Staker Id: ", stakerID)

	selectedProposedBlock, err := razorUtils.GetProposedBlock(rpcParameters, epoch, sortedProposedBlockIds[0])
	if err != nil {
		log.Error("Error in getting selectedProposedBlock: ", err)
		return core.NilHash, err
	}
	log.Debug("ClaimBlockReward: Selected proposed block: ", selectedProposedBlock)

	if selectedProposedBlock.ProposerId == stakerID {
		log.Info("Claiming block reward...")
		txnOpts, err := razorUtils.GetTxnOpts(rpcParameters, options)
		if err != nil {
			return core.NilHash, err
		}

		client, err := rpcParameters.RPCManager.GetBestRPCClient()
		if err != nil {
			return core.NilHash, err
		}

		log.Debug("Executing ClaimBlockReward transaction...")
		txn, err := blockManagerUtils.ClaimBlockReward(client, txnOpts)
		if err != nil {
			log.Error("Error in claiming block reward: ", err)
			return core.NilHash, err
		}
		txnHash := transactionUtils.Hash(txn)
		log.Info("Txn Hash: ", txnHash.Hex())
		return txnHash, nil
	}

	log.Debug("Only selected block proposer can claim block reward")
	return core.NilHash, nil
}
