package types

import "net/http"

type HttpClientConfig struct {
	Timeout                   int64
	MaxIdleConnections        int
	MaxIdleConnectionsPerHost int
}

type HttpClientInterface interface {
	Do(request *http.Request) (*http.Response, error)
}
