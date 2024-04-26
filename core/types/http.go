package types

type HttpClientConfig struct {
	Timeout                   int64
	MaxIdleConnections        int
	MaxIdleConnectionsPerHost int
}
