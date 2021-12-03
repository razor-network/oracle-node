package utils

import (
	"errors"
	"github.com/PaesslerAG/jsonpath"
	"github.com/gocolly/colly"
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
	if response.StatusCode != http.StatusOK {
		return nil, errors.New("API not responding")
	}
	defer response.Body.Close()
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

func GetDataFromHTML(url string, selector string) (string, error) {
	c := colly.NewCollector()
	var priceData string
	c.OnHTML(selector, func(e *colly.HTMLElement) {
		priceData = e.Text
	})
	err := c.Visit(url)
	if err != nil {
		return "", err
	}
	return priceData, nil
}
