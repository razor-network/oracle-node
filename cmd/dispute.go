package cmd

import (
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/utils"
)

func HandleDispute(client *ethclient.Client, config types.Configurations, account types.Account, epoch *big.Int) {
	numberOfProposedBlocks, err := utils.GetNumberOfProposedBlocks(client, account.Address, epoch)
	if err != nil {
		log.Error(err)
		return
	}
	if numberOfProposedBlocks.Cmp(big.NewInt(5)) > 0 {
		numberOfProposedBlocks = big.NewInt(5)
	}
	for i := 0; i < int(numberOfProposedBlocks.Int64()); i++ {
		proposedBlock, err := utils.GetProposedBlock(client, account.Address, epoch, big.NewInt(int64(i)))
		if err != nil {
			log.Error(err)
			continue
		}
		log.Info("Values in the block")
		log.Infof("Medians: %s, LowerCutOffs: %s, HigherCutOffs: %s", proposedBlock.BlockMedians, proposedBlock.LowerCutoffs, proposedBlock.HigherCutoffs)
		medians, lowerCutOffs, higherCutOffs, err := MakeBlock(client, account.Address, epoch)
		if err != nil {
			log.Error(err)
			continue
		}
		log.Info("Locally calculated data:")
		log.Infof("Medians: %s, LowerCutOffs: %s, HigherCutOffs: %s\n", medians, lowerCutOffs, higherCutOffs)
		if !utils.IsEqual(proposedBlock.BlockMedians, medians) ||
			!utils.IsEqual(proposedBlock.LowerCutoffs, lowerCutOffs) ||
			!utils.IsEqual(proposedBlock.HigherCutoffs, higherCutOffs) {
			log.Warn("BLOCK NOT MATCHING WITH LOCAL CALCULATIONS.")
			log.Info("Block Values: ", proposedBlock.BlockMedians, proposedBlock.LowerCutoffs, proposedBlock.HigherCutoffs)
			log.Info("Local Calculations: ", medians, lowerCutOffs, higherCutOffs)
			err := Dispute(client, config, account, epoch, big.NewInt(int64(i)))
			if err != nil {
				log.Error("Error in disputing...", err)
				continue
			}
		} else {
			log.Info("Proposed median matches with local calculations. Will not open dispute.")
		}
	}
}

func Dispute(client *ethclient.Client, config types.Configurations, account types.Account, epoch *big.Int, blockId *big.Int) error {
	// TODO: Check what should be the value of asset id
	sortedVotes, _, err := getSortedVotes(client, account.Address, 0, epoch)
	if err != nil {
		return err
	}
	iter := int(math.Ceil(float64(len(sortedVotes))/1000))
	blockManager := utils.GetBlockManager(client)
	for i := 0; i < iter; i++ {
		 log.Info(epoch, sortedVotes[i*1000:i*1000+1000])
		 txnOpts := utils.GetTxnOpts(types.TransactionOptions{
			 Client:         client,
			 Password:       account.Password,
			 AccountAddress: account.Address,
			 ChainId:        core.ChainId,
			 GasMultiplier:  config.GasMultiplier,
		 })
		 txn, err := blockManager.GiveSorted(txnOpts, epoch, big.NewInt(0),sortedVotes[i*1000: i*1000+1])
		 if err != nil {
		 	log.Error(err)
		 	continue
		 }
		 log.Info("Txn Hash: ", txn.Hash())
		 utils.WaitForBlockCompletion(client, txn.Hash().String())
	}

	log.Info("Sending finalized dispute...")
	txnOpts := utils.GetTxnOpts(types.TransactionOptions{
		Client:         client,
		Password:       account.Password,
		AccountAddress: account.Address,
		ChainId:        core.ChainId,
		GasMultiplier:  config.GasMultiplier,
	})
	txn, err :=blockManager.FinalizeDispute(txnOpts, epoch, blockId)
	if err != nil {
		return err
	}
	log.Info("Txn Hash: ", txn.Hash())
	utils.WaitForBlockCompletion(client, txn.Hash().String())
	return nil
}