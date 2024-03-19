package client

import (
	"net/http"
	"time"
)

var httpClient *http.Client

func InitHttpClient(httpTimeout int64) {
	httpClient = &http.Client{
		Timeout: time.Duration(httpTimeout) * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        5,
			MaxIdleConnsPerHost: 5,
		},
	}
}

func GetHttpClient() *http.Client {
	return httpClient
}
