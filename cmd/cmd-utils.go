package cmd

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math"
	"math/big"
	"razor/core"
	"razor/utils"
	"time"
)

func WaitForCommitState(client *ethclient.Client, accountAddress string, _type string) *big.Int {
	var epoch *big.Int
	for true {
		epoch = utils.GetEpoch(client, accountAddress)
		state := getDelayedState(client)
		log.Info("Epoch ", epoch)
		log.Info("State ", state)
		if state != 0 {
			log.Infof("Can only %s during state 0 (commit). Retrying in 1 second...", _type)
			time.Sleep(1 * time.Second)
		} else {
			break
		}
	}
	return epoch
}

func getDelayedState(client *ethclient.Client) int64 {
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal("Error in fetching latest block number: ", err)
	}

	if blockNumber%(core.BlockDivider) > 7 || blockNumber%(core.BlockDivider) < 1 {
		return -1
	}
	state := math.Floor(float64(blockNumber / core.BlockDivider))
	return int64(state) % core.NumberOfStates
}
