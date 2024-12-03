package rpc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"razor/core"
	"razor/path"
	"sort"
	"strings"
	"sync"
	"time"
)

func (m *RPCManager) calculateMetrics(endpoint *RPCEndpoint) error {
	ctx, cancel := context.WithTimeout(context.Background(), core.EndpointsContextTimeout*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, endpoint.URL)
	if err != nil {
		return fmt.Errorf("failed to connect to RPC: %w", err)
	}

	start := time.Now()
	blockNumber, err := m.fetchBlockNumberWithTimeout(ctx, client)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return fmt.Errorf("RPC call timed out: %w", err)
		}
		return fmt.Errorf("RPC call failed: %w", err)
	}
	latency := time.Since(start).Seconds()

	endpoint.BlockNumber = blockNumber
	endpoint.Latency = latency
	endpoint.Client = client

	return nil
}

// updateAndSortEndpoints calculates metrics and sorts the endpoints
func (m *RPCManager) updateAndSortEndpoints() error {
	if len(m.Endpoints) == 0 {
		return fmt.Errorf("no endpoints available to update")
	}

	var wg sync.WaitGroup
	logrus.Debug("Starting concurrent metrics calculation for all endpoints...")

	for _, endpoint := range m.Endpoints {
		wg.Add(1)
		go func(ep *RPCEndpoint) {
			defer wg.Done()
			if err := m.calculateMetrics(ep); err != nil {
				logrus.Errorf("Error calculating metrics for endpoint %s: %v", ep.URL, err)
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

	// Update the best RPC endpoint after sorting
	m.BestEndpoint = m.Endpoints[0]

	logrus.Infof("Best RPC endpoint updated: %s (BlockNumber: %d, Latency: %.2f)",
		m.BestEndpoint.URL, m.BestEndpoint.BlockNumber, m.BestEndpoint.Latency)

	return nil
}

// RefreshEndpoints will update and sort the endpoints.
func (m *RPCManager) RefreshEndpoints() error {
	if err := m.updateAndSortEndpoints(); err != nil {
		return fmt.Errorf("failed to refresh endpoints: %w", err)
	}
	logrus.Infof("Endpoints refreshed successfully")
	return nil
}

func InitializeRPCManager(provider string) (*RPCManager, error) {
	defaultPath, err := path.PathUtilsInterface.GetDefaultPath()
	if err != nil {
		return nil, fmt.Errorf("failed to get .razor path: %w", err)
	}

	endpointsFile := filepath.Join(defaultPath, "endpoints.json")
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
		logrus.Infof("Adding user-provided endpoint: %s", provider)
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

	if m.BestEndpoint.Client == nil {
		return nil, fmt.Errorf("no best RPC client available")
	}
	return m.BestEndpoint.Client, nil
}

// GetBestEndpointURL returns the URL of the best endpoint
func (m *RPCManager) GetBestEndpointURL() (string, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	if m.BestEndpoint == nil {
		return "", fmt.Errorf("no best endpoint available")
	}
	return m.BestEndpoint.URL, nil
}

// SwitchToNextBestRPCClient switches to the next best available client after the current best client.
// If no valid next best client is found, it retains the current best client.
func (m *RPCManager) SwitchToNextBestRPCClient() (bool, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	// If there are fewer than 2 endpoints, there are no alternate clients to switch to.
	if len(m.Endpoints) < 2 {
		logrus.Warn("No alternate RPC clients available. Retaining the current best client.")
		return false, nil
	}

	// Find the index of the current best client
	var currentIndex = -1
	for i, endpoint := range m.Endpoints {
		if endpoint.Client == m.BestEndpoint.Client {
			currentIndex = i
			break
		}
	}

	// If the current client is not found (which is rare), return an error
	if currentIndex == -1 {
		return false, fmt.Errorf("current best client not found in the list of endpoints")
	}

	// Iterate through the remaining endpoints to find a valid next best client
	for i := 1; i < len(m.Endpoints); i++ {
		nextIndex := (currentIndex + i) % len(m.Endpoints)
		nextEndpoint := m.Endpoints[nextIndex]

		// Check if we can connect to the next endpoint
		ctx, cancel := context.WithTimeout(context.Background(), core.EndpointsContextTimeout*time.Second)
		cancel()

		client, err := ethclient.DialContext(ctx, nextEndpoint.URL)
		if err != nil {
			logrus.Errorf("Failed to connect to RPC endpoint %s: %v", nextEndpoint.URL, err)
			continue
		}

		// Try fetching block number to validate the endpoint
		_, err = m.fetchBlockNumberWithTimeout(ctx, client)
		if err != nil {
			logrus.Errorf("Failed to fetch block number for endpoint %s: %v", nextEndpoint.URL, err)
			continue
		}

		// Successfully connected and validated, update the best client and endpoint
		m.BestEndpoint.Client = client
		m.BestEndpoint = nextEndpoint

		logrus.Infof("Switched to the next best RPC endpoint: %s (BlockNumber: %d, Latency: %.2f)",
			m.BestEndpoint.URL, m.BestEndpoint.BlockNumber, m.BestEndpoint.Latency)
		return true, nil
	}

	// If no valid endpoint is found, retain the current best client
	logrus.Warn("No valid next-best RPC client found. Retaining the current best client.")
	return false, nil
}

func (m *RPCManager) fetchBlockNumberWithTimeout(ctx context.Context, client *ethclient.Client) (uint64, error) {
	blockNumber, err := client.BlockNumber(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return 0, fmt.Errorf("RPC call timed out: %w", err)
		}
		return 0, fmt.Errorf("RPC call failed: %w", err)
	}
	return blockNumber, nil
}
