package block

import (
	"context"
	"razor/core"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

var latestBlock *types.Header
var mu = sync.Mutex{}

func GetLatestBlock() *types.Header {
	mu.Lock()
	defer mu.Unlock()
	return latestBlock
}

func SetLatestBlock(block *types.Header) {
	mu.Lock()
	latestBlock = block
	mu.Unlock()
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
