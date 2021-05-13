package cmd

import (
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
	"razor/utils"
	"time"
)

func WaitForCommitState(client *ethclient.Client, accountAddress string, _type string) (*big.Int, error) {
	for true {
		epoch, err := utils.GetEpoch(client, accountAddress)
		if err != nil {
			log.Fatal("Error in fetching epoch: ", err)
		}
		state, err := utils.GetDelayedState(client)
		if err != nil {
			log.Fatal("Error in fetching state: ", err)
		}
		log.Info("Epoch ", epoch)
		log.Info("State ", state)
		if state != 0 {
			log.Infof("Can only %s during state 0 (commit). Retrying in 1 second...", _type)
			time.Sleep(1 * time.Second)
		} else {
			return epoch, nil
		}
	}
	return nil, nil
}