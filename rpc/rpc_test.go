// File: manager_integration_test.go
package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/undefinedlabs/go-mpatch"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

// jsonRPCHandler mimics a minimal Ethereum JSON–RPC endpoint.
// It expects a request for "eth_blockNumber" and returns block number 100 (hex "0x64").
func jsonRPCHandler(w http.ResponseWriter, r *http.Request) {
	var req map[string]interface{}
	// Decode the incoming JSON–RPC request.
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	// Verify that the method is "eth_blockNumber"
	method, ok := req["method"].(string)
	if !ok || method != "eth_blockNumber" {
		http.Error(w, "unsupported method", http.StatusBadRequest)
		return
	}

	// Build and send a JSON–RPC response with block number 100.
	resp := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      req["id"],
		"result":  "0x64", // 0x64 in hexadecimal equals 100 in decimal
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

// errorRPCHandler simulates a failing RPC endpoint.
// It returns an HTTP error (status 500) so that calls to it fail.
func errorRPCHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

// TestRPCManagerWithJSONRPC starts a test HTTP server using jsonRPCHandler.
// It creates an RPCManager instance with the test server's URL as its endpoint.
// After calling RefreshEndpoints (which internally calls updateAndSortEndpoints),
// it verifies that the endpoint's metrics (block number and latency) have been set.
func TestRPCManagerWithJSONRPC(t *testing.T) {
	// Start a test HTTP server with our JSON–RPC handler.
	ts := httptest.NewServer(http.HandlerFunc(jsonRPCHandler))
	defer ts.Close()

	// Create an RPCManager instance with one endpoint pointing to our test server.
	manager := &RPCManager{
		Endpoints: []*RPCEndpoint{
			{URL: ts.URL},
		},
	}

	// Call RefreshEndpoints, which will update metrics by dialing the endpoint.
	if err := manager.RefreshEndpoints(); err != nil {
		t.Fatalf("RefreshEndpoints failed: %v", err)
	}

	// Retrieve the best endpoint's URL.
	bestURL, err := manager.GetBestEndpointURL()
	if err != nil {
		t.Fatalf("GetBestEndpointURL failed: %v", err)
	}
	if bestURL != ts.URL {
		t.Errorf("Expected best endpoint to be %s, got %s", ts.URL, bestURL)
	}

	// Verify that the endpoint’s BlockNumber has been set correctly (should be 100).
	if manager.BestEndpoint.BlockNumber != 100 {
		t.Errorf("Expected block number 100, got %d", manager.BestEndpoint.BlockNumber)
	}

	// (Optional) Verify that latency has been measured (should be > 0).
	if manager.BestEndpoint.Latency <= 0 {
		t.Errorf("Expected positive latency, got %f", manager.BestEndpoint.Latency)
	}
}

// delayedJSONRPCHandler returns a handler that simulates a response delay
// and returns a specified block number in hexadecimal.
func delayedJSONRPCHandler(delay time.Duration, blockNumber uint64) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Simulate network or processing delay.
		time.Sleep(delay)

		var req map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		// Build a hex string for the provided block number.
		hexBlock := "0x" + strconv.FormatUint(blockNumber, 16)
		resp := map[string]interface{}{
			"jsonrpc": "2.0",
			"id":      req["id"],
			"result":  hexBlock,
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	}
}

// TestRPCManagerWithDelayedJSONRPC uses a test HTTP server that delays its response.
// It checks that the RPC manager measures latency correctly and that the block number is updated.
func TestRPCManagerWithDelayedJSONRPC(t *testing.T) {
	// Configure a delay and a custom block number (e.g., 150).
	delay := 2 * time.Second
	blockNumber := uint64(150)

	// Start a test HTTP server with the delayed handler.
	ts := httptest.NewServer(delayedJSONRPCHandler(delay, blockNumber))
	defer ts.Close()

	manager := &RPCManager{
		Endpoints: []*RPCEndpoint{
			{URL: ts.URL},
		},
	}

	// Record the start time.
	start := time.Now()
	if err := manager.RefreshEndpoints(); err != nil {
		t.Fatalf("RefreshEndpoints failed: %v", err)
	}
	elapsed := time.Since(start)

	// Ensure that the total elapsed time is at least the simulated delay.
	if elapsed < delay {
		t.Errorf("Expected elapsed time at least %v, got %v", delay, elapsed)
	}

	// Verify that the best endpoint is the test server.
	bestURL, err := manager.GetBestEndpointURL()
	if err != nil {
		t.Fatalf("GetBestEndpointURL failed: %v", err)
	}
	if bestURL != ts.URL {
		t.Errorf("Expected best endpoint to be %s, got %s", ts.URL, bestURL)
	}

	// Check that the BlockNumber was set to the custom value.
	if manager.BestEndpoint.BlockNumber != blockNumber {
		t.Errorf("Expected block number %d, got %d", blockNumber, manager.BestEndpoint.BlockNumber)
	}

	// Also verify that the measured latency is at least the delay.
	if manager.BestEndpoint.Latency < delay.Seconds() {
		t.Errorf("Expected measured latency at least %f, got %f", delay.Seconds(), manager.BestEndpoint.Latency)
	}
}

// makeHandler returns an HTTP handler that responds to a JSON–RPC "eth_blockNumber" request
// with the provided blockNumber (formatted as a hexadecimal string).
func makeHandler(blockNumber uint64) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode the incoming JSON–RPC request.
		var req map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		// Build a hex string for the block number.
		hexBlock := fmt.Sprintf("0x%x", blockNumber)
		resp := map[string]interface{}{
			"jsonrpc": "2.0",
			"id":      req["id"],
			"result":  hexBlock,
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	}
}

// fakeDialContext is our custom replacement for ethclient.DialContext.
// It ignores the context and calls ethclient.Dial instead.
func fakeDialContext(ctx context.Context, url string) (*ethclient.Client, error) {
	return ethclient.Dial(url)
}

func TestSelectBestEndpointAndSwitch(t *testing.T) {
	// Patch ethclient.DialContext using PatchMethod.
	patch, err := mpatch.PatchMethod(ethclient.DialContext, fakeDialContext)
	if err != nil {
		t.Fatalf("Failed to patch ethclient.DialContext: %v", err)
	}
	defer patch.Unpatch()

	// Now set up your test servers and RPCManager as before.
	ts1 := httptest.NewServer(makeHandler(100))
	defer ts1.Close()

	ts2 := httptest.NewServer(makeHandler(120))
	// We will close ts2 later.

	ts3 := httptest.NewServer(makeHandler(110))
	defer ts3.Close()

	manager := &RPCManager{
		Endpoints: []*RPCEndpoint{
			{URL: ts1.URL},
			{URL: ts2.URL},
			{URL: ts3.URL},
		},
	}

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
	t.Logf("Initially, best endpoint is: %s", bestURL)

	// Simulate failure of ts2.
	ts2.Close()
	time.Sleep(100 * time.Millisecond)

	switched, err := manager.SwitchToNextBestRPCClient()
	if err != nil {
		t.Fatalf("SwitchToNextBestRPCClient returned error: %v", err)
	}
	if !switched {
		t.Fatal("Expected a switch to next best RPC client, but switching did not occur")
	}

	newBestURL, err := manager.GetBestEndpointURL()
	if err != nil {
		t.Fatalf("GetBestEndpointURL failed after switch: %v", err)
	}
	if newBestURL != ts3.URL {
		t.Errorf("Expected new best endpoint to be %s, got %s", ts3.URL, newBestURL)
	} else {
		t.Logf("After switch, new best endpoint is: %s", newBestURL)
	}

	if manager.BestEndpoint.BlockNumber != 110 {
		t.Errorf("Expected block number 110 for new best endpoint, got %d", manager.BestEndpoint.BlockNumber)
	}
}
