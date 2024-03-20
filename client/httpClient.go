package client

import (
	"net/http"
	"razor/core"
	"time"
)

var httpClient *http.Client

func InitHttpClient(httpTimeout int64) {
	httpClient = &http.Client{
		Timeout: time.Duration(httpTimeout) * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        core.HTTPClientMaxIdleConns,
			MaxIdleConnsPerHost: core.HTTPClientMaxIdleConnsPerHost,
		},
	}
}

func GetHttpClient() *http.Client {
	return httpClient
}
