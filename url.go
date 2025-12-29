package main

import (
	"net/url"
)

func CreateUrl(city string, apiKey string) (string, error) {
	url, err := url.Parse("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/")
	if err != nil {
		return "", err
	}

	url = url.JoinPath(city)
	q := url.Query()
	q.Set("key", apiKey)
	q.Set("unitGroup", "metric")
	url.RawQuery = q.Encode()

	return url.String(), nil
}
