package block

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"razor/core"
	"time"
)

var latestBlock *types.Header

func GetLatestBlock() *types.Header {
	return latestBlock
}

func SetLatestBlock(block *types.Header) {
	latestBlock = block
}

func CalculateLatestBlock(client *ethclient.Client) {
	for {
		if client != nil {
			latestHeader, err := client.HeaderByNumber(context.Background(), nil)
			if err != nil {
				logrus.Error("CalculateBlockNumber: Error in fetching block: ", err)
				continue
			}
			SetLatestBlock(latestHeader)
		}
		time.Sleep(time.Second * time.Duration(core.BlockNumberInterval))
	}
}
