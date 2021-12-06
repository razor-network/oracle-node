package utils

import (
	"errors"
	"io/ioutil"
	"net/http"
	"razor/core"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/avast/retry-go"
	"github.com/gocolly/colly"
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
			if response.StatusCode != 200 {
				return errors.New("unable to reach API")
			}
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

func GetDataFromHTML(url string, selector string) (string, error) {
	c := colly.NewCollector()
	var priceData string
	c.OnHTML(selector, func(e *colly.HTMLElement) {
		priceData = e.Text
	})
	err := retry.Do(func() error {
		visitErr := c.Visit(url)
		if visitErr != nil {
			log.Error("Error in visiting URL.... Retrying")
			return visitErr
		}
		return nil
	}, retry.Attempts(core.MaxRetries))
	if err != nil {
		return "", err
	}
	return priceData, nil
}
