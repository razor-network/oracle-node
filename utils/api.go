package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"razor/cache"
	"razor/core"
	"regexp"
	"time"

	"razor/core/types"

	"github.com/PaesslerAG/jsonpath"
	"github.com/avast/retry-go"
	"github.com/gocolly/colly"
)

func GetDataFromAPI(dataSourceURLStruct types.DataSourceURL, localCache *cache.LocalCache) ([]byte, error) {
	client := http.Client{
		Timeout: time.Duration(HTTPTimeout) * time.Second,
	}

	cachedData, found := localCache.Read(dataSourceURLStruct.URL)
	if found {
		log.Debugf("Getting Data for URL %s from local cache...", dataSourceURLStruct.URL)
		return cachedData, nil
	}

	response, err := makeAPIRequest(client, dataSourceURLStruct)
	if err != nil {
		return nil, err
	}

	cacheKey, err := generateCacheKey(dataSourceURLStruct.URL, dataSourceURLStruct.Body)
	if err != nil {
		log.Errorf("Error in generating cache key for API %s: %v", dataSourceURLStruct.URL, err)
		return nil, err
	}

	// Storing the data into cache
	localCache.Update(response, cacheKey, time.Now().Add(time.Second*time.Duration(core.StateLength)).Unix())
	return response, nil
}

func makeAPIRequest(client http.Client, dataSourceURLStruct types.DataSourceURL) ([]byte, error) {
	var requestBody io.Reader // Using the broader io.Reader interface here

	switch dataSourceURLStruct.Type {
	case "GET":
		// For HTTP GET requests, there is typically no request body.
		// So we explicitly set the requestBody to nil to indicate this absence.
		requestBody = nil
	case "POST":
		postBody, err := json.Marshal(dataSourceURLStruct.Body)
		if err != nil {
			log.Errorf("Error in marshalling body of a POST request URL %s: %v", dataSourceURLStruct.URL, err)
			return nil, err
		}
		requestBody = bytes.NewBuffer(postBody)
	default:
		return nil, errors.New("invalid request type")
	}

	var response []byte
	err := retry.Do(
		func() error {
			responseBody, err := ProcessRequest(client, dataSourceURLStruct, requestBody)
			if err != nil {
				log.Errorf("Error in processing %s request: %v", dataSourceURLStruct.Type, err)
				return err
			}
			response = responseBody
			return nil
		}, retry.Attempts(core.ProcessRequestRetryAttempts), retry.Delay(time.Second*time.Duration(core.ProcessRequestRetryDelay)))

	if err != nil {
		return nil, err
	}

	return response, nil
}

func GetDataFromJSON(jsonObject map[string]interface{}, selector string) (interface{}, error) {
	if selector[0] == '[' {
		selector = "$" + selector
	} else {
		selector = "$." + selector
	}
	return jsonpath.Get(selector, jsonObject)
}

func GetDataFromXHTML(dataSourceURLStruct types.DataSourceURL, selector string) (string, error) {
	c := colly.NewCollector()
	var priceData string
	c.OnXML(selector, func(e *colly.XMLElement) {
		priceData = e.Text
	})
	err := c.Visit(dataSourceURLStruct.URL)
	if err != nil {
		return "", err
	}
	return priceData, nil
}

func processHeaderValue(value string, re *regexp.Regexp) string {
	// check if any API authentication is required
	if re.MatchString(value) {
		return ReplaceValueWithDataFromENVFile(re, value)
	}
	return value
}

func addHeaderToRequest(request *http.Request, headerMap map[string]string) *http.Request {
	re := regexp.MustCompile(core.APIKeyRegex)
	for key, value := range headerMap {
		processedValue := processHeaderValue(value, re)
		log.Debugf("Adding key: %s, value: %s pair to header", key, value)
		request.Header.Add(key, processedValue)
	}
	return request
}

func ProcessRequest(client http.Client, dataSourceURLStruct types.DataSourceURL, requestBody io.Reader) ([]byte, error) {
	request, err := http.NewRequest(dataSourceURLStruct.Type, dataSourceURLStruct.URL, requestBody)
	if err != nil {
		return nil, err
	}
	requestWithHeader := addHeaderToRequest(request, dataSourceURLStruct.Header)
	response, err := client.Do(requestWithHeader)
	if err != nil {
		log.Errorf("Error sending %s request URL %s: %v", dataSourceURLStruct.Type, dataSourceURLStruct.URL, err)
		return nil, err
	}
	defer response.Body.Close()
	// Success is indicated with 2xx status codes:
	statusOK := response.StatusCode >= 200 && response.StatusCode < 300
	if !statusOK {
		log.Errorf("API: %s responded with status code %d", dataSourceURLStruct.URL, response.StatusCode)
		return nil, fmt.Errorf("HTTP request failed with status code %d: %s", response.StatusCode, response.Status)
	}
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return responseBody, nil
}
