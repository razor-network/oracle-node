package utils

import (
	"github.com/PaesslerAG/jsonpath"
	"github.com/avast/retry-go"
	"io/ioutil"
	"net/http"
	"razor/core"
	"time"
)

func GetDataFromAPI(url string) ([]byte, error) {
	client := http.Client{
		Timeout: 60 * time.Second,
	}
	var body []byte
	err := retry.Do(
		func() error {
			response, err := client.Get(url)
			if err != nil {
				return err
			}
			defer response.Body.Close()
			body, err = ioutil.ReadAll(response.Body)
			if err != nil {
				return err
			}
			return nil
		}, retry.Attempts(core.MaxRetries))
	if err != nil {
		return nil, err
	}
	return body, nil
}

func GetDataFromJSON(jsonObject map[string]interface{}, selector string) (interface{}, error) {
	if selector[0] == '[' {
		selector = "$" + selector
	} else {
		selector = "$." + selector
	}
	return jsonpath.Get(selector, jsonObject)
}
