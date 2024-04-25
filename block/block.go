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

// BlockManager manages the latest block information
type BlockManager struct {
	latestBlock *types.Header
	mu          sync.Mutex
	client      *ethclient.Client
}

// NewBlockManager creates a new BlockManager instance
func NewBlockManager(client *ethclient.Client) *BlockManager {
	return &BlockManager{
		client: client,
	}
}

func (bm *BlockManager) GetLatestBlock() *types.Header {
	bm.mu.Lock()
	defer bm.mu.Unlock()
	return bm.latestBlock
}

func (bm *BlockManager) SetLatestBlock(block *types.Header) {
	bm.mu.Lock()
	bm.latestBlock = block
	bm.mu.Unlock()
}

func (bm *BlockManager) CalculateLatestBlock() {
	for {
		if bm.client != nil {
			latestHeader, err := bm.client.HeaderByNumber(context.Background(), nil)
			if err != nil {
				logrus.Error("CalculateLatestBlock: Error in fetching block: ", err)
				continue
			}
			bm.SetLatestBlock(latestHeader)
		}
		time.Sleep(time.Second * time.Duration(core.BlockNumberInterval))
	}
}
