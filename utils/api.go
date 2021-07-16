package utils

import (
	"github.com/PaesslerAG/jsonpath"
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

func GetDataFromJSON(jsonObject map[string]interface{}, selector string) (interface{}, error) {
	if selector[0] == '[' {
		selector = "$" + selector
	} else {
		selector = "$." + selector
	}
	return jsonpath.Get(selector, jsonObject)
}
