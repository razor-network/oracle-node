package RPC

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"os"
	"path/filepath"
	"razor/logger"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
)

var log = logger.NewLogger()

func (m *RPCManager) calculateMetrics(endpoint *RPCEndpoint) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, endpoint.URL)
	if err != nil {
		return fmt.Errorf("failed to connect to RPC: %w", err)
	}

	start := time.Now()
	blockNumber, err := client.BlockNumber(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return fmt.Errorf("RPC call timed out: %w", err)
		}
		return fmt.Errorf("RPC call failed: %w", err)
	}
	latency := time.Since(start).Seconds()

	endpoint.BlockNumber = blockNumber
	endpoint.Latency = latency
	endpoint.Client = client // Store the client for future use

	return nil
}

// updateAndSortEndpoints calculates metrics and sorts the endpoints
func (m *RPCManager) updateAndSortEndpoints() error {
	if len(m.Endpoints) == 0 {
		return fmt.Errorf("no endpoints available to update")
	}

	var wg sync.WaitGroup
	log.Debug("Starting concurrent metrics calculation for all endpoints...")

	for _, endpoint := range m.Endpoints {
		wg.Add(1)
		go func(ep *RPCEndpoint) {
			defer wg.Done()
			if err := m.calculateMetrics(ep); err != nil {
				log.Printf("Error calculating metrics for endpoint %s: %v", ep.URL, err)
			}
		}(endpoint)
	}
	wg.Wait()

	log.Debug("Concurrent metrics calculation complete. Sorting endpoints...")

	m.mutex.Lock()
	defer m.mutex.Unlock()

	sort.Slice(m.Endpoints, func(i, j int) bool {
		if m.Endpoints[i].BlockNumber == m.Endpoints[j].BlockNumber {
			return m.Endpoints[i].Latency < m.Endpoints[j].Latency
		}
		return m.Endpoints[i].BlockNumber > m.Endpoints[j].BlockNumber
	})

	// Update the best RPC client after sorting
	m.BestRPCClient = m.Endpoints[0].Client

	return nil
}

// RefreshEndpoints will update and sort the endpoints.
func (m *RPCManager) RefreshEndpoints() error {
	if err := m.updateAndSortEndpoints(); err != nil {
		return fmt.Errorf("failed to refresh endpoints: %w", err)
	}
	log.Infof("Endpoints refreshed successfully")
	return nil
}

func InitializeRPCManager(provider string) (*RPCManager, error) {
	// Locate the razor-go project directory dynamically
	_, file, _, ok := runtime.Caller(0) // Get the current file (rpc.go) path
	if !ok {
		return nil, fmt.Errorf("failed to get the caller information")
	}

	// Move up to the razor-go project root
	projectDir := filepath.Dir(filepath.Dir(file)) // This gets `razor-go` directory
	endpointsFile := filepath.Join(projectDir, "endpoints.json")

	fileData, err := os.ReadFile(endpointsFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read endpoints.json: %w", err)
	}

	// Unmarshal the JSON file into a list of RPC endpoints
	var rpcEndpointsList []string
	err = json.Unmarshal(fileData, &rpcEndpointsList)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal endpoints.json: %w", err)
	}

	// Normalize provider input and check if it's already in the list
	provider = strings.TrimSpace(provider) // Trim whitespace from the input
	providerFound := false
	for _, endpoint := range rpcEndpointsList {
		if endpoint == provider {
			providerFound = true
			break
		}
	}

	// If the provider is not found, add it to the list
	if !providerFound && provider != "" {
		log.Infof("Adding user-provided endpoint: %s", provider)
		rpcEndpointsList = append(rpcEndpointsList, provider)
	}

	// Initialize the RPC endpoints
	rpcEndpoints := make([]*RPCEndpoint, len(rpcEndpointsList))
	for i, url := range rpcEndpointsList {
		rpcEndpoints[i] = &RPCEndpoint{URL: url}
	}

	rpcManager := &RPCManager{
		Endpoints: rpcEndpoints,
	}

	// Pre-calculate metrics and set the best client on initialization
	if err := rpcManager.updateAndSortEndpoints(); err != nil {
		return nil, fmt.Errorf("failed to initialize RPC Manager: %w", err)
	}

	return rpcManager, nil
}

func (m *RPCManager) GetBestRPCClient() (*ethclient.Client, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	if m.BestRPCClient == nil {
		return nil, fmt.Errorf("no best RPC client available")
	}
	return m.BestRPCClient, nil
}

// SwitchToNextBestRPCClient switches to the next best available client after the current best client.
func (m *RPCManager) SwitchToNextBestRPCClient() error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	// If there are fewer than 2 endpoints, there are no alternate clients to switch to.
	if len(m.Endpoints) < 2 {
		return fmt.Errorf("no other RPC clients to switch to")
	}

	// Find the index of the current best client
	var currentIndex = -1
	for i, endpoint := range m.Endpoints {
		if endpoint.Client == m.BestRPCClient {
			currentIndex = i
			break
		}
	}

	// If current client is not found (which is rare), return an error
	if currentIndex == -1 {
		return fmt.Errorf("current best client not found in the list of endpoints")
	}

	// Calculate the next index by wrapping around if necessary
	nextIndex := (currentIndex + 1) % len(m.Endpoints)

	// Switch to the next client in the list
	m.BestRPCClient = m.Endpoints[nextIndex].Client
	log.Infof("Switched to the next best RPC client: %s", m.Endpoints[nextIndex].URL)
	return nil
}
