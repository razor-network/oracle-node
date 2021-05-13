package utils

import (
	"io/ioutil"
	"net/http"
	"time"
)

func GetDataFromAPI(url string) ([]byte, error) {
	client := http.Client{
		Timeout: 60 * time.Second,
	}
	response, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(response.Body)
}
