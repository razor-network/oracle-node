package cmd

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"
)

var voteManagerUtils voteManagerInterface

func (utilsStruct UtilsStruct) HandleCommitState(client *ethclient.Client, epoch uint32, rogueData types.Rogue) ([]*big.Int, error) {
	var (
		data []*big.Int
		err  error
	)
	//rogue mode
	if rogueData.IsRogue && utils.Contains(rogueData.RogueMode, "commit") {
		numActiveAssets, err := utilsStruct.razorUtils.GetNumActiveAssets(client)
		if err != nil {
			return nil, err
		}
		for i := 0; i < int(numActiveAssets.Int64()); i++ {
			rogueValue := utilsStruct.razorUtils.GetRogueRandomValue(10000000)
			data = append(data, rogueValue)
		}
		log.Debug("Data: ", data)
		return data, nil
	}

	//normal mode
	data, err = utilsStruct.razorUtils.GetActiveAssetsData(client, epoch)
	if err != nil {
		return nil, err
	}
	log.Debug("Data: ", data)
	return data, nil
}

func (utilsStruct UtilsStruct) Commit(client *ethclient.Client, data []*big.Int, secret []byte, account types.Account, config types.Configurations) (common.Hash, error) {
	if state, err := utilsStruct.razorUtils.GetDelayedState(client, config.BufferPercent); err != nil || state != 0 {
		log.Error("Not commit state")
		return core.NilHash, err
	}

	epoch, err := utilsStruct.razorUtils.GetEpoch(client)
	if err != nil {
		return core.NilHash, err
	}

	commitment := solsha3.SoliditySHA3([]string{"uint32", "uint256[]", "bytes32"}, []interface{}{epoch, data, "0x" + hex.EncodeToString(secret)})
	commitmentToSend := [32]byte{}
	copy(commitmentToSend[:], commitment)
	txnOpts := utilsStruct.razorUtils.GetTxnOpts(types.TransactionOptions{
		Client:          client,
		Password:        account.Password,
		AccountAddress:  account.Address,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.VoteManagerAddress,
		ABI:             bindings.VoteManagerABI,
		MethodName:      "commit",
		Parameters:      []interface{}{epoch, commitmentToSend},
	}, utilsStruct.packageUtils)

	log.Debugf("Committing: epoch: %d, commitment: %s, secret: %s, account: %s", epoch, "0x"+hex.EncodeToString(commitment), "0x"+hex.EncodeToString(secret), account.Address)

	log.Info("Commitment sent...")
	txn, err := utilsStruct.voteManagerUtils.Commit(client, txnOpts, epoch, commitmentToSend)
	if err != nil {
		return core.NilHash, err
	}
	log.Info("Txn Hash: ", utilsStruct.transactionUtils.Hash(txn))
	return utilsStruct.transactionUtils.Hash(txn), nil
}
