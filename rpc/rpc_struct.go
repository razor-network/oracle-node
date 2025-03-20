package rpc

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"sync"
)

type RPCEndpoint struct {
	URL         string
	BlockNumber uint64
	Latency     float64
	Client      *ethclient.Client
}

type RPCManager struct {
	Endpoints    []*RPCEndpoint
	mutex        sync.RWMutex
	BestEndpoint *RPCEndpoint // Holds the URL to current best RPC client
}

type RPCParameters struct {
	Ctx        context.Context // Context with timeout for handling unresponsive RPC calls
	RPCManager *RPCManager     // RPC manager for client selection and contract calls
}
