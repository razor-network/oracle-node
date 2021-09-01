package cmd

import (
	"math/big"
	"razor/utils"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

func WaitForCommitState(client *ethclient.Client, accountAddress string, action string) (*big.Int, error) {
	for {
		epoch, err := utils.GetEpoch(client, accountAddress)
		if err != nil {
			log.Fatal("Error in fetching epoch: ", err)
		}
		bufferPercent, err := getBufferPercent()
		if err != nil {
			log.Fatal(err)
		}
		state, err := utils.GetDelayedState(client, bufferPercent)
		if err != nil {
			log.Fatal("Error in fetching state: ", err)
		}
		log.Info("Epoch ", epoch)
		switch state {
		case 0:
			log.Info("State 0: Commit")
		case 1:
			log.Info("State 1: Reveal")
		case 2:
			log.Info("State 2: Propose")
		case 3:
			log.Info("State 3: Dispute")
		default:
			log.Info("State: ", state)
		}

		if state != 0 {
			log.Infof("Can only %s during state 0 (commit). Retrying in 5 second...", action)
			time.Sleep(5 * time.Second)
		} else {
			return epoch, nil
		}
	}
}
