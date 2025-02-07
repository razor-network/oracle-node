package rpc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

// decodeRPCRequest decodes the JSON–RPC request body into a map.
func decodeRPCRequest(r *http.Request) (map[string]interface{}, error) {
	var req map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// writeRPCResponse writes the provided response as JSON with the appropriate header.
func writeRPCResponse(w http.ResponseWriter, response map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

// jsonRPCHandler simulates an Ethereum JSON–RPC endpoint that returns block number 100.
func jsonRPCHandler(w http.ResponseWriter, r *http.Request) {
	req, err := decodeRPCRequest(r)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	// Verify that the method is "eth_blockNumber".
	if method, ok := req["method"].(string); !ok || method != "eth_blockNumber" {
		http.Error(w, "unsupported method", http.StatusBadRequest)
		return
	}

	// Build and send a JSON–RPC response with block number 100 (hex "0x64").
	resp := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      req["id"],
		"result":  "0x64",
	}
	writeRPCResponse(w, resp)
}

// errorRPCHandler simulates a failing RPC endpoint by returning status 500.
func errorRPCHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

// delayedJSONRPCHandler returns a handler that simulates a delay and returns the given block number.
func delayedJSONRPCHandler(delay time.Duration, blockNumber uint64) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Simulate network delay.
		time.Sleep(delay)

		req, err := decodeRPCRequest(r)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		hexBlock := "0x" + strconv.FormatUint(blockNumber, 16)
		resp := map[string]interface{}{
			"jsonrpc": "2.0",
			"id":      req["id"],
			"result":  hexBlock,
		}
		writeRPCResponse(w, resp)
	}
}

// makeHandler returns an HTTP handler that responds with the provided block number.
func makeHandler(blockNumber uint64) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := decodeRPCRequest(r)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		hexBlock := fmt.Sprintf("0x%x", blockNumber)
		resp := map[string]interface{}{
			"jsonrpc": "2.0",
			"id":      req["id"],
			"result":  hexBlock,
		}
		writeRPCResponse(w, resp)
	}
}

// createManagerFromServers creates an RPCManager from a list of test servers.
func createManagerFromServers(servers ...*httptest.Server) *RPCManager {
	endpoints := make([]*RPCEndpoint, len(servers))
	for i, s := range servers {
		endpoints[i] = &RPCEndpoint{URL: s.URL}
	}
	return &RPCManager{Endpoints: endpoints}
}

// TestRPCManagerWithJSONRPC verifies that RefreshEndpoints correctly pings multiple endpoints,
// updates their metrics, and selects the best (fastest healthy) endpoint.
func TestRPCManagerWithJSONRPC(t *testing.T) {
	// Create three test servers:
	// ts1: healthy with no artificial delay (using jsonRPCHandler which has block number 100 defined)
	ts1 := httptest.NewServer(http.HandlerFunc(jsonRPCHandler))
	defer ts1.Close()

	// ts2: healthy but with a 200ms delay (using delayedJSONRPCHandler returning block 100)
	ts2 := httptest.NewServer(delayedJSONRPCHandler(200*time.Millisecond, 100))
	defer ts2.Close()

	// ts3: unhealthy (using errorRPCHandler)
	ts3 := httptest.NewServer(http.HandlerFunc(errorRPCHandler))
	defer ts3.Close()

	// Create an RPCManager with all three endpoints.
	manager := createManagerFromServers(ts1, ts2, ts3)

	// Refresh endpoints so that the manager pings each endpoint.
	if err := manager.RefreshEndpoints(); err != nil {
		t.Fatalf("RefreshEndpoints failed: %v", err)
	}

	// Since ts1 and ts2 are healthy and return the same block number (100),
	// the tie-breaker is latency. ts1 should be faster (no delay) than ts2.
	bestURL, err := manager.GetBestEndpointURL()
	if err != nil {
		t.Fatalf("GetBestEndpointURL failed: %v", err)
	}

	if bestURL != ts1.URL {
		t.Errorf("Expected best endpoint to be %s, got %s", ts1.URL, bestURL)
	}

	if manager.BestEndpoint.BlockNumber != 100 {
		t.Errorf("Expected block number 100, got %d", manager.BestEndpoint.BlockNumber)
	}

	if manager.BestEndpoint.Latency <= 0 {
		t.Errorf("Expected positive latency, got %f", manager.BestEndpoint.Latency)
	}
}

// TestRPCManagerWithDelayedJSONRPC verifies that a delayed response is handled correctly,
// with latency measured properly and block number updated.
func TestRPCManagerWithDelayedJSONRPC(t *testing.T) {
	delay := 2 * time.Second
	blockNumber := uint64(150)

	ts := httptest.NewServer(delayedJSONRPCHandler(delay, blockNumber))
	defer ts.Close()

	manager := &RPCManager{
		Endpoints: []*RPCEndpoint{
			{URL: ts.URL},
		},
	}

	start := time.Now()
	if err := manager.RefreshEndpoints(); err != nil {
		t.Fatalf("RefreshEndpoints failed: %v", err)
	}
	elapsed := time.Since(start)
	if elapsed < delay {
		t.Errorf("Expected elapsed time at least %v, got %v", delay, elapsed)
	}

	bestURL, err := manager.GetBestEndpointURL()
	if err != nil {
		t.Fatalf("GetBestEndpointURL failed: %v", err)
	}
	if bestURL != ts.URL {
		t.Errorf("Expected best endpoint URL %s, got %s", ts.URL, bestURL)
	}

	if manager.BestEndpoint.BlockNumber != blockNumber {
		t.Errorf("Expected block number %d, got %d", blockNumber, manager.BestEndpoint.BlockNumber)
	}

	// Ensure the measured latency reflects the simulated delay.
	if manager.BestEndpoint.Latency < delay.Seconds() {
		t.Errorf("Expected measured latency at least %f, got %f", delay.Seconds(), manager.BestEndpoint.Latency)
	}
}

// TestSwitchToNextBestClient verifies the behavior of SwitchToNextBestRPCClient under various conditions.
func TestSwitchToNextBestClient(t *testing.T) {
	t.Run("switches when best endpoint fails", func(t *testing.T) {
		// Create three test servers with different block numbers.
		ts1 := httptest.NewServer(makeHandler(100))
		defer ts1.Close()

		ts2 := httptest.NewServer(makeHandler(120))
		// Do not defer ts2.Close() here to simulate its failure.

		ts3 := httptest.NewServer(makeHandler(110))
		defer ts3.Close()

		manager := createManagerFromServers(ts1, ts2, ts3)

		if err := manager.RefreshEndpoints(); err != nil {
			t.Fatalf("RefreshEndpoints failed: %v", err)
		}

		bestURL, err := manager.GetBestEndpointURL()
		if err != nil {
			t.Fatalf("GetBestEndpointURL failed: %v", err)
		}
		if bestURL != ts2.URL {
			t.Fatalf("Expected best endpoint to be %s, got %s", ts2.URL, bestURL)
		}

		// Simulate failure of the best endpoint.
		ts2.Close()
		time.Sleep(100 * time.Millisecond)

		switched, err := manager.SwitchToNextBestRPCClient()
		if err != nil {
			t.Fatalf("SwitchToNextBestRPCClient returned error: %v", err)
		}
		if !switched {
			t.Fatal("Expected switch to next best endpoint, but no switch occurred")
		}

		newBestURL, err := manager.GetBestEndpointURL()
		if err != nil {
			t.Fatalf("GetBestEndpointURL failed after switch: %v", err)
		}
		if newBestURL != ts3.URL {
			t.Errorf("Expected new best endpoint to be %s, got %s", ts3.URL, newBestURL)
		}
		if manager.BestEndpoint.BlockNumber != 110 {
			t.Errorf("Expected block number 110 for new best endpoint, got %d", manager.BestEndpoint.BlockNumber)
		}
	})

	t.Run("does not switch when the rest of the endpoints are not healthy", func(t *testing.T) {
		ts1 := httptest.NewServer(makeHandler(100))
		ts2 := httptest.NewServer(makeHandler(120))
		ts3 := httptest.NewServer(makeHandler(110))
		defer ts1.Close()
		defer ts2.Close()
		defer ts3.Close()

		manager := createManagerFromServers(ts1, ts2, ts3)

		if err := manager.RefreshEndpoints(); err != nil {
			t.Fatalf("RefreshEndpoints failed: %v", err)
		}

		bestURL, err := manager.GetBestEndpointURL()
		if err != nil {
			t.Fatalf("GetBestEndpointURL failed: %v", err)
		}
		if bestURL != ts2.URL {
			t.Fatalf("Expected best endpoint to be %s, got %s", ts2.URL, bestURL)
		}

		// Simulate failure of ts1 and ts3.
		ts1.Close()
		ts3.Close()
		time.Sleep(100 * time.Millisecond)

		switched, err := manager.SwitchToNextBestRPCClient()
		if err != nil {
			t.Fatalf("SwitchToNextBestRPCClient returned error: %v", err)
		}
		if switched {
			t.Error("Did not expect a switch when the alternative endpoints are unhealthy")
		}

		newBestURL, err := manager.GetBestEndpointURL()
		if err != nil {
			t.Fatalf("GetBestEndpointURL failed after attempted switch: %v", err)
		}
		if newBestURL != ts2.URL {
			t.Errorf("Expected best endpoint to remain %s, got %s", ts2.URL, newBestURL)
		}
	})

	t.Run("returns false when only one endpoint is available", func(t *testing.T) {
		ts1 := httptest.NewServer(makeHandler(120))
		defer ts1.Close()

		manager := createManagerFromServers(ts1)
		if err := manager.RefreshEndpoints(); err != nil {
			t.Fatalf("RefreshEndpoints failed: %v", err)
		}

		switched, err := manager.SwitchToNextBestRPCClient()
		if err != nil {
			t.Fatalf("SwitchToNextBestRPCClient returned error: %v", err)
		}
		if switched {
			t.Error("Expected no switch when only one endpoint is available")
		}
	})

	t.Run("returns error when current best client is not found in the list", func(t *testing.T) {
		ts1 := httptest.NewServer(makeHandler(100))
		ts2 := httptest.NewServer(makeHandler(120))
		defer ts1.Close()
		defer ts2.Close()

		manager := createManagerFromServers(ts1, ts2)
		if err := manager.RefreshEndpoints(); err != nil {
			t.Fatalf("RefreshEndpoints failed: %v", err)
		}

		// Simulate a scenario where the current best client is not in the endpoints list.
		manager.BestEndpoint = &RPCEndpoint{URL: "http://nonexistent"}
		_, err := manager.SwitchToNextBestRPCClient()
		if err == nil {
			t.Error("Expected an error when current best client is not found, but got nil")
		}
	})
}
