package block

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"razor/rpc"
)

func makeBlockHandler(blockNumber uint64, blockTime uint64) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		method, ok := req["method"].(string)
		if !ok {
			http.Error(w, "missing method", http.StatusBadRequest)
			return
		}

		var result interface{}

		switch method {
		case "eth_blockNumber":
			result = fmt.Sprintf("0x%x", blockNumber)

		case "eth_getBlockByNumber":
			// Use go-ethereum's types.Header struct to ensure all required fields are included
			header := &types.Header{
				Number:      big.NewInt(int64(blockNumber)),
				Time:        blockTime,
				ParentHash:  common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
				UncleHash:   types.EmptyUncleHash,
				Coinbase:    common.HexToAddress("0x0000000000000000000000000000000000000000"),
				Root:        common.HexToHash("0x2222222222222222222222222222222222222222222222222222222222222222"),
				TxHash:      common.HexToHash("0x3333333333333333333333333333333333333333333333333333333333333333"),
				ReceiptHash: common.HexToHash("0x4444444444444444444444444444444444444444444444444444444444444444"),
				Bloom:       types.Bloom{},
				Difficulty:  big.NewInt(2),
				GasLimit:    30000000,
				GasUsed:     0,
				Extra:       []byte{},
			}

			// Serialize the header to JSON format
			result = header
		default:
			http.Error(w, "unsupported method", http.StatusBadRequest)
			return
		}

		resp := map[string]interface{}{
			"jsonrpc": "2.0",
			"id":      req["id"],
			"result":  result,
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	}
}

// createManagerFromServers creates an RPCManager from a list of test servers.
func createManagerFromServers(servers ...*httptest.Server) *rpc.RPCManager {
	endpoints := make([]*rpc.RPCEndpoint, len(servers))
	for i, s := range servers {
		endpoints[i] = &rpc.RPCEndpoint{URL: s.URL}
	}
	return &rpc.RPCManager{Endpoints: endpoints}
}

// TestBlockMonitorUpdateBlock tests `updateLatestBlock` behavior with different block numbers.
func TestBlockMonitorUpdateBlock(t *testing.T) {
	// Simulate two endpoints: one returning an outdated block, another returning an up-to-date block.
	blockNumber := uint64(100)
	outdatedBlockNumber := uint64(90)

	ts1 := httptest.NewServer(makeBlockHandler(blockNumber, uint64(time.Now().Unix())))
	t.Cleanup(func() { ts1.Close() })

	ts2 := httptest.NewServer(makeBlockHandler(outdatedBlockNumber, uint64(time.Now().Unix())))
	t.Cleanup(func() { ts2.Close() })

	// Create an RPC manager with both endpoints.
	manager := createManagerFromServers(ts1, ts2)

	// Refresh endpoints so that the manager pings each endpoint.
	if err := manager.RefreshEndpoints(); err != nil {
		t.Fatalf("RefreshEndpoints failed: %v", err)
	}

	// Initialize BlockMonitor.
	client, _ := manager.GetBestRPCClient()
	bm := NewBlockMonitor(client, manager, 1, 10)

	// Simulate fetching the latest block.
	bm.updateLatestBlock()

	if bm.latestBlock == nil {
		t.Fatal("Expected latest block to be set, but got nil")
	}

	// Ensure that the latest block number is from the correct, up-to-date endpoint.
	if bm.latestBlock.Number.Uint64() != blockNumber {
		t.Errorf("Expected block number %d, got %d", blockNumber, bm.latestBlock.Number.Uint64())
	}

	// Simulate outdated block number being reported.
	bm.latestBlock.Number = big.NewInt(int64(outdatedBlockNumber))
	bm.updateLatestBlock()

	// The block number should remain at the latest, and an RPC switch should have occurred.
	newBestURL, _ := manager.GetBestEndpointURL()
	if newBestURL != ts1.URL {
		t.Errorf("Expected best endpoint to be %s after outdated block, got %s", ts1.URL, newBestURL)
	}
}

// TestBlockMonitorStaleBlock checks if the stale block detection works correctly.
func TestBlockMonitorStaleBlock(t *testing.T) {
	currentTime := uint64(time.Now().Unix())
	staleTime := uint64(time.Now().Add(-15 * time.Second).Unix())

	ts1 := httptest.NewServer(makeBlockHandler(120, currentTime))
	t.Cleanup(func() { ts1.Close() })

	ts2 := httptest.NewServer(makeBlockHandler(110, staleTime))
	t.Cleanup(func() { ts2.Close() })

	// Create an RPC manager with both endpoints.
	manager := createManagerFromServers(ts1, ts2)

	// Refresh endpoints so that the manager pings each endpoint.
	if err := manager.RefreshEndpoints(); err != nil {
		t.Fatalf("RefreshEndpoints failed: %v", err)
	}

	// Initialize BlockMonitor.
	client, _ := manager.GetBestRPCClient()
	bm := NewBlockMonitor(client, manager, 1, 10)

	// Fetch the latest block (stale one)
	bm.updateLatestBlock()
	bm.checkForStaleBlock()

	if bm.latestBlock.Number.Uint64() != 120 {
		t.Errorf("Expected block number 120 after detecting stale block, got %d", bm.latestBlock.Number.Uint64())
	}
}

// TestBlockMonitorSwitchOnStale tests switching to a better endpoint when a stale block is detected.
func TestBlockMonitorSwitchOnStale(t *testing.T) {
	latestBlock := uint64(150)
	staleBlock := uint64(140)

	var blockNumber atomic.Uint64
	blockNumber.Store(staleBlock)

	// Simulate a server that starts with a stale block but updates to a fresh block.
	ts1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		makeBlockHandler(blockNumber.Load(), uint64(time.Now().Unix()))(w, r)
	}))
	t.Cleanup(func() { ts1.Close() })

	ts2 := httptest.NewServer(makeBlockHandler(latestBlock, uint64(time.Now().Unix())))
	t.Cleanup(func() { ts2.Close() })

	// Create an RPC manager.
	manager := createManagerFromServers(ts1, ts2)

	// Refresh endpoints so that the manager pings each endpoint.
	if err := manager.RefreshEndpoints(); err != nil {
		t.Fatalf("RefreshEndpoints failed: %v", err)
	}

	// Initialize BlockMonitor.
	client, _ := manager.GetBestRPCClient()
	bm := NewBlockMonitor(client, manager, 1, 10)

	// Start with a stale block.
	blockNumber.Store(staleBlock)
	bm.updateLatestBlock()
	bm.checkForStaleBlock()

	// Ensure the switch occurred to the better endpoint.
	bestURL, _ := manager.GetBestEndpointURL()
	if bestURL != ts2.URL {
		t.Errorf("Expected best endpoint to switch to %s, got %s", ts2.URL, bestURL)
	}

	// Now update the first endpoint to a fresh block.
	blockNumber.Store(latestBlock)
	bm.updateLatestBlock()
	bm.checkForStaleBlock()

	// The monitor should detect the updated block.
	if bm.latestBlock.Number.Uint64() != latestBlock {
		t.Errorf("Expected latest block number %d, got %d", latestBlock, bm.latestBlock.Number.Uint64())
	}
}

// TestBlockMonitorSwitchFailure tests when no alternate endpoints are available.
func TestBlockMonitorSwitchFailure(t *testing.T) {
	staleBlock := uint64(80)

	ts1 := httptest.NewServer(makeBlockHandler(staleBlock, uint64(time.Now().Add(-20*time.Second).Unix())))
	t.Cleanup(func() { ts1.Close() })

	// Create an RPC manager with a single stale endpoint.
	manager := createManagerFromServers(ts1)

	// Refresh endpoints so that the manager pings each endpoint.
	if err := manager.RefreshEndpoints(); err != nil {
		t.Fatalf("RefreshEndpoints failed: %v", err)
	}

	// Initialize BlockMonitor.
	client, _ := manager.GetBestRPCClient()
	bm := NewBlockMonitor(client, manager, 1, 10)

	// Start with a stale block.
	bm.updateLatestBlock()
	bm.checkForStaleBlock()

	// Since no alternate endpoints are available, the best endpoint should remain unchanged.
	bestURL, _ := manager.GetBestEndpointURL()
	if bestURL != ts1.URL {
		t.Errorf("Expected best endpoint to remain %s, got %s", ts1.URL, bestURL)
	}
}
