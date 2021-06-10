package cmd

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	log "github.com/sirupsen/logrus"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/utils"
)

func HandleCommitState(client *ethclient.Client, address string) []*big.Int {
	jobs, err := utils.GetActiveJobs(client, address)
	if err != nil {
		log.Error("Error in getting active jobs: ", err)
		return nil
	}

	data := getDataToCommitFromJobs(jobs)

	log.Info("Data", data)

	return data
}

func getDataToCommitFromJobs(jobs []types.Job) []*big.Int {
	var data []*big.Int
	for _, job := range jobs {
		datum := big.NewFloat(0)
		var parsedJSON map[string]interface{}

		response, err := utils.GetDataFromAPI(job.Url)
		if err != nil {
			log.Error(err)
			data = append(data, big.NewInt(0))
			continue
		}

		err = json.Unmarshal(response, &parsedJSON)
		if err != nil {
			log.Error("Error in parsing data from API: ", err)
			data = append(data, big.NewInt(0))
			continue
		}
		parsedData, err := utils.GetDataFromJSON(parsedJSON, job.Selector)
		if err != nil {
			log.Error("Error in fetching value from parsed data ", err)
			continue
		}
		datum, err = utils.ConvertToNumber(parsedData)
		if err != nil {
			log.Error("Result is not a number")
			data = append(data, big.NewInt(0))
			continue
		}

		dataToAppend := utils.MultiplyToEightDecimals(datum)
		data = append(data, dataToAppend)
	}
	if len(data) == 0 {
		data = append(data, big.NewInt(0))
	}
	return data
}

func Commit(client *ethclient.Client, data []*big.Int, secret []byte, account types.Account, config types.Configurations) error {
	if state, err := utils.GetDelayedState(client); err != nil || state != 0 {
		log.Error("Not commit state")
		return err
	}

	root, err := utils.GetMerkleTreeRoot(data)
	if err != nil {
		return err
	}

	epoch, err := utils.GetEpoch(client, account.Address)
	if err != nil {
		return err
	}

	// FIXME: Not required
	commitments, err := utils.GetCommitments(client, account.Address, epoch)
	if err != nil {
		return err
	}
	if !utils.AllZero(commitments) {
		return errors.New("already committed")
	}

	commitment := solsha3.SoliditySHA3([]string{"uint256", "bytes32", "bytes32"}, []interface{}{epoch.String(), "0x"+hex.EncodeToString(root), "0x"+hex.EncodeToString(secret)})

	voteManager := utils.GetVoteManager(client)
	txnOpts := utils.GetTxnOpts(types.TransactionOptions{
		Client:         client,
		Password:       account.Password,
		AccountAddress: account.Address,
		ChainId:        core.ChainId,
		GasMultiplier:  config.GasMultiplier,
	})
	commitmentToSend := [32]byte{}
	copy(commitmentToSend[:], commitment)

	log.Infof("Committing: epoch: %s, root: %s, commitment: %s, secret: %s, account: %s", epoch, "0x"+hex.EncodeToString(root), "0x"+hex.EncodeToString(commitment), "0x"+hex.EncodeToString(secret), account.Address)

	txn, err := voteManager.Commit(txnOpts, epoch, commitmentToSend)
	if err != nil {
		return err
	}

	log.Info("Commitment sent...\nTxn hash: ", txn.Hash())

	if utils.WaitForBlockCompletion(client, fmt.Sprintf("%s", txn.Hash())) == 0 {
		log.Error("Commit failed....")
	}
	return nil
}