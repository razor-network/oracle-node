package types

import (
	"context"
	"razor/RPC"
)

type RPCParameters struct {
	Ctx        context.Context // Context with timeout for handling unresponsive RPC calls
	RPCManager *RPC.RPCManager // RPC manager for client selection and contract calls
}
