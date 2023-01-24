package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"razor/cache"
	"razor/core"
	"strings"
	"time"

	"io/ioutil"
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
		var body []byte
		if dataSourceURLStruct.Type == "GET" {
			err := retry.Do(
				func() error {
					response, err := client.Get(dataSourceURLStruct.URL)
					if err != nil {
						return err
					}
					defer response.Body.Close()
					if response.StatusCode != 200 {
						log.Errorf("API: %s responded with status code %d", dataSourceURLStruct.URL, response.StatusCode)
						return errors.New("unable to reach API")
					}
					body, err = IOInterface.ReadAll(response.Body)
					if err != nil {
						return err
					}
					return nil
				}, retry.Attempts(2), retry.Delay(time.Second*2))
			if err != nil {
				return nil, err
			}
		}
		if dataSourceURLStruct.Type == "POST" {
			postBody, err := json.Marshal(dataSourceURLStruct.Body)
			if err != nil {
				log.Errorf("Error in marshalling body of a POST request URL %s: %v", dataSourceURLStruct.URL, err)
			}
			responseBody := bytes.NewBuffer(postBody)
			err = retry.Do(
				func() error {
					request, err := http.NewRequest("POST", dataSourceURLStruct.URL, responseBody)
					if err != nil {
						log.Errorf("Error in creating a POST request for URL %s: %v", dataSourceURLStruct.URL, err)
						return err
					}
					requestWithHeader, err := AddHeaderToPostRequest(request, dataSourceURLStruct.Header)
					if err != nil {
						log.Error("Error in adding header to post request: ", err)
						return err
					}
					response, err := client.Do(requestWithHeader)
					if err != nil {
						log.Errorf("Error sending POST request URL %s: %v", dataSourceURLStruct.URL, err)
						return err
					}
					defer response.Body.Close()
					if response.StatusCode != 200 {
						log.Errorf("URL: %s responded with status code %d", dataSourceURLStruct.URL, response.StatusCode)
						return errors.New("unable to reach API")
					}
					body, err = ioutil.ReadAll(response.Body)
					if err != nil {
						return err
					}
					return nil
				}, retry.Attempts(2), retry.Delay(time.Second*2))
			if err != nil {
				return nil, err
			}
		}
		dataToCache := cache.Data{
			Result: body,
		}
		localCache.Update(dataToCache, dataSourceURLStruct.URL, time.Now().Add(time.Second*time.Duration(core.StateLength)).Unix())
		return body, nil
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

func AddHeaderToPostRequest(request *http.Request, headerMap map[string]string) (*http.Request, error) {
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
