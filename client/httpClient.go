package client

import (
	"net/http"
	"razor/core/types"
	"time"
)

type HttpClient struct {
	Client *http.Client
}

func NewHttpClient(config types.HttpClientConfig) *HttpClient {
	client := &http.Client{
		Timeout: time.Duration(config.Timeout) * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        config.MaxIdleConnections,
			MaxIdleConnsPerHost: config.MaxIdleConnectionsPerHost,
		},
	}
	return &HttpClient{client}
}

func (hc *HttpClient) Do(request *http.Request) (*http.Response, error) {
	return hc.Client.Do(request)
}
