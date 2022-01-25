package cmd

import (
	"razor/core"
	"razor/core/types"

	"github.com/ethereum/go-ethereum/common"
)

var blockManagerUtils blockManagerInterface

func (utilsStruct UtilsStruct) ClaimBlockReward(options types.TransactionOptions) (common.Hash, error) {
	epoch, err := utilsStruct.razorUtils.GetEpoch(options.Client)
	if err != nil {
		log.Error("Error in getting epoch: ", err)
		return core.NilHash, err
	}

	sortedProposedBlockIds, err := utilsStruct.razorUtils.GetSortedProposedBlockIds(options.Client, options.AccountAddress, epoch)
	if err != nil {
		log.Error("Error in getting sortedProposedBlockIds: ", err)
		return core.NilHash, err
	}

	if sortedProposedBlockIds == nil {
		log.Debug("No blocks proposed in this epoch")
		return core.NilHash, nil
	}

	stakerID, err := utilsStruct.razorUtils.GetStakerId(options.Client, options.AccountAddress)
	if err != nil {
		log.Error("Error in getting stakerId: ", err)
		return core.NilHash, err
	}

	selectedProposedBlock, err := utilsStruct.razorUtils.GetProposedBlock(options.Client, options.AccountAddress, epoch, sortedProposedBlockIds[0])
	if err != nil {
		log.Error("Error in getting selectedProposedBlock: ", err)
		return core.NilHash, err
	}

	if selectedProposedBlock.ProposerId == stakerID {
		log.Info("Claiming block reward...")
		txnOpts := utilsStruct.razorUtils.GetTxnOpts(options)
		txn, err := utilsStruct.blockManagerUtils.ClaimBlockReward(options.Client, txnOpts)
		if err != nil {
			log.Error("Error in claiming block reward: ", err)
			return core.NilHash, err
		}
		log.Info("Txn Hash: ", utilsStruct.transactionUtils.Hash(txn).Hex())
		return utilsStruct.transactionUtils.Hash(txn), nil
	}

	log.Debug("Only selected block proposer can claim block reward")
	return core.NilHash, nil
}
