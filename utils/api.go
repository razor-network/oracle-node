package utils

import (
	"io/ioutil"
	"net/http"
	"strings"
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
	if !strings.Contains(selector, ",") {
		return jsonObject[selector], nil
	}
	splitSelector := strings.Split(selector, ",")
	return getNestedDataFromJSON(jsonObject, splitSelector)
}

func getNestedDataFromJSON(jsonObject map[string]interface{}, splitSelector []string) (interface{}, error) {
	if len(splitSelector) == 1 {
		return jsonObject[splitSelector[0]], nil
	}
	for i, s := range splitSelector {
		return getNestedDataFromJSON(jsonObject[s].(map[string]interface{}), splitSelector[i+1:])
	}
	return nil, nil
}