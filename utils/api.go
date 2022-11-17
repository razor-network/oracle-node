package utils

import (
	"errors"
	"net/http"
	"razor/cache"
	"razor/core"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/avast/retry-go"
	"github.com/gocolly/colly"
)

func (*UtilsStruct) GetDataFromAPI(url string, localCache *cache.LocalCache) ([]byte, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	cachedData, err := localCache.Read(url)
	if err != nil {
		var body []byte
		err := retry.Do(
			func() error {
				response, err := client.Get(url)
				if err != nil {
					return err
				}
				defer response.Body.Close()
				if response.StatusCode != 200 {
					log.Errorf("API: %s responded with status code %d", url, response.StatusCode)
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
		dataToCache := cache.Data{
			Result: body,
		}
		localCache.Update(dataToCache, url, time.Now().Add(time.Second*time.Duration(core.StateLength)).Unix())
		return body, nil
	}
	log.Debugf("Getting Data for URL %s from local cache...", url)
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

func (*UtilsStruct) GetDataFromXHTML(url string, selector string) (string, error) {
	c := colly.NewCollector()
	var priceData string
	c.OnXML(selector, func(e *colly.XMLElement) {
		priceData = e.Text
	})
	err := c.Visit(url)
	if err != nil {
		return "", err
	}
	return priceData, nil
}
