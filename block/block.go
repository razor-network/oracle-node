package block

import (
	"context"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"razor/rpc"
)

// BlockMonitor monitors the latest block and handles stale blocks.
type BlockMonitor struct {
	client         *ethclient.Client
	rpcManager     *rpc.RPCManager
	latestBlock    *types.Header
	mu             sync.Mutex
	checkInterval  time.Duration
	staleThreshold time.Duration
}

// NewBlockMonitor initializes a BlockMonitor with RPC integration.
func NewBlockMonitor(client *ethclient.Client, rpcManager *rpc.RPCManager, checkInterval, staleThreshold time.Duration) *BlockMonitor {
	return &BlockMonitor{
		client:         client,
		rpcManager:     rpcManager,
		checkInterval:  time.Second * checkInterval,
		staleThreshold: time.Second * staleThreshold,
	}
}

// Start begins the block monitoring process.
func (bm *BlockMonitor) Start() {
	go func() {
		for {
			bm.updateLatestBlock()
			bm.checkForStaleBlock()
			time.Sleep(bm.checkInterval)
		}
	}()
}

// GetLatestBlock retrieves the most recent block header safely.
func (bm *BlockMonitor) GetLatestBlock() *types.Header {
	bm.mu.Lock()
	defer bm.mu.Unlock()
	return bm.latestBlock
}

// updateLatestBlock fetches the latest block and updates the state.
func (bm *BlockMonitor) updateLatestBlock() {
	if bm.client == nil {
		return
	}

	header, err := bm.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		logrus.Errorf("Error fetching latest block: %v", err)
		return
	}

	bm.mu.Lock()
	defer bm.mu.Unlock()

	// Check if the fetched block number is less than the last recorded block number.
	if bm.latestBlock != nil && header.Number.Uint64() < bm.latestBlock.Number.Uint64() {
		logrus.Warnf("Fetched block number %d is less than the last recorded block number %d. Switching to the next best RPC endpoint.",
			header.Number.Uint64(), bm.latestBlock.Number.Uint64())

		// Attempt to switch to the next best RPC endpoint.
		if bm.rpcManager != nil {
			switched, err := bm.rpcManager.SwitchToNextBestRPCClient()
			if err != nil {
				logrus.Errorf("Failed to switch RPC endpoint: %v", err)
			} else if switched {
				logrus.Info("Switched to the next best RPC endpoint.")
				bm.updateClient()
			} else {
				logrus.Warn("Retaining the current best RPC endpoint as no valid alternate was found.")
			}
		}
		return
	}

	// Update the latest block only if it changes.
	if bm.latestBlock == nil || header.Number.Uint64() != bm.latestBlock.Number.Uint64() {
		bm.latestBlock = header
	}
}

// checkForStaleBlock detects stale blocks and triggers appropriate actions.
func (bm *BlockMonitor) checkForStaleBlock() {
	if bm.staleThreshold == 0 {
		return
	}

	bm.mu.Lock()
	defer bm.mu.Unlock()

	if bm.latestBlock == nil || time.Since(time.Unix(int64(bm.latestBlock.Time), 0)) >= bm.staleThreshold {
		logrus.Warnf("Stale block detected: Block %d is stale for %s", bm.latestBlock.Number.Uint64(), bm.staleThreshold)

		// Switch to the next best RPC endpoint if stale block detected.
		if bm.rpcManager != nil {
			switched, err := bm.rpcManager.SwitchToNextBestRPCClient()
			if err != nil {
				logrus.Errorf("Failed to switch RPC endpoint: %v", err)
			} else if switched {
				logrus.Info("Switched to the next best RPC endpoint.")
				bm.updateClient()
			} else {
				logrus.Warn("Retaining the current best RPC endpoint as no valid alternate was found.")
			}
		}
	}
}

// updateClient updates the Ethereum client to use the new best RPC endpoint.
func (bm *BlockMonitor) updateClient() {
	if bm.rpcManager == nil {
		return
	}

	newClient, err := bm.rpcManager.GetBestRPCClient()
	if err != nil {
		return
	}

	bm.client = newClient
	logrus.Info("Client in logger updated with the new best RPC endpoint.")
}
