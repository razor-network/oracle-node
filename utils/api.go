package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"razor/cache"
	"razor/core"
	"strings"
	"time"

	"razor/core/types"

	"github.com/PaesslerAG/jsonpath"
	"github.com/avast/retry-go"
	"github.com/gocolly/colly"
)

func (*UtilsStruct) GetDataFromAPI(dataSourceURLStruct types.DataSourceURL, localCache *cache.LocalCache) ([]byte, error) {
	client := http.Client{
		Timeout: time.Duration(HTTPTimeout) * time.Second,
	}
	cachedData, cachedErr := localCache.Read(dataSourceURLStruct.URL)
	if cachedErr != nil {
		var response []byte
		switch dataSourceURLStruct.Type {
		case "GET":
			err := retry.Do(
				func() error {
					responseBody, err := ProcessRequest(client, dataSourceURLStruct, nil)
					if err != nil {
						log.Error("Error in processing GET request: ", err)
						return err
					}
					response = responseBody
					return nil
				}, retry.Attempts(core.ProcessRequestRetryAttempts), retry.Delay(time.Second*time.Duration(core.ProcessRequestRetryDelay)))
			if err != nil {
				return nil, err
			}
		case "POST":
			postBody, err := json.Marshal(dataSourceURLStruct.Body)
			if err != nil {
				log.Errorf("Error in marshalling body of a POST request URL %s: %v", dataSourceURLStruct.URL, err)
				return nil, err
			}
			requestBody := bytes.NewBuffer(postBody)
			err = retry.Do(
				func() error {
					responseBody, err := ProcessRequest(client, dataSourceURLStruct, requestBody)
					if err != nil {
						log.Error("Error in processing POST request: ", err)
						return err
					}
					response = responseBody
					return nil
				}, retry.Attempts(core.ProcessRequestRetryAttempts), retry.Delay(time.Second*time.Duration(core.ProcessRequestRetryDelay)))
			if err != nil {
				return nil, err
			}
		default:
			return nil, errors.New("invalid request type")
		}

		dataToCache := cache.Data{
			Result: response,
		}
		localCache.Update(dataToCache, dataSourceURLStruct.URL, time.Now().Add(time.Second*time.Duration(core.StateLength)).Unix())
		return response, nil
	}
	log.Debugf("Getting Data for URL %s from local cache...", dataSourceURLStruct.URL)
	return cachedData.Result, nil
}

func (*UtilsStruct) GetDataFromJSON(jsonObject map[string]interface{}, selector string) (interface{}, error) {
	if selector[0] == '[' {
		selector = "$" + selector
	} else {
		selector = "$." + selector
	}
	return jsonpath.Get(selector, jsonObject)
}

func (*UtilsStruct) GetDataFromXHTML(dataSourceURLStruct types.DataSourceURL, selector string) (string, error) {
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

func AddHeaderToRequest(request *http.Request, headerMap map[string]string) (*http.Request, error) {
	for key, value := range headerMap {
		// If core.APIKeyRegex = `$` and if value starts with '$' then we need to fetch the respective value from env file
		if strings.HasPrefix(value, core.APIKeyRegex) {
			_, APIKey, err := GetKeyWordAndAPIKeyFromENVFile(value)
			if err != nil {
				log.Error("Error in getting value from env file: ", err)
				return nil, err
			}
			value = APIKey
		}
		log.Debugf("Adding key: %s, value: %s pair to header", key, value)
		request.Header.Add(key, value)
	}
	return request, nil
}

func ProcessRequest(client http.Client, dataSourceURLStruct types.DataSourceURL, requestBody io.Reader) ([]byte, error) {
	request, err := http.NewRequest(dataSourceURLStruct.Type, dataSourceURLStruct.URL, requestBody)
	if err != nil {
		return nil, err
	}
	requestWithHeader, err := AddHeaderToRequest(request, dataSourceURLStruct.Header)
	if err != nil {
		log.Errorf("Error in adding header to %s request: %v", dataSourceURLStruct.Type, err)
		return nil, err
	}
	response, err := client.Do(requestWithHeader)
	if err != nil {
		log.Errorf("Error sending %s request URL %s: %v", dataSourceURLStruct.Type, dataSourceURLStruct.URL, err)
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		log.Errorf("API: %s responded with status code %d", dataSourceURLStruct.URL, response.StatusCode)
		return nil, errors.New("unable to reach API")
	}
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return responseBody, nil
}
