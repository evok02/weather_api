package main

import (
	"net/url"
	"time"
	"fmt"
)

func getCurrentDate() string {
	t := time.Now()
	y, m, d := t.Year(), t.Month(), t.Day()
	return fmt.Sprintf("%d-%d-%d", y, m, d)
}	

func CreateUrl(city string, apiKey string) (string, error) {
	rawUrl, err := url.Parse("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/")
	if err != nil {
		return "", err
	}

	url := rawUrl.JoinPath(city)
	if !isVerbose() {
		date := getCurrentDate()
		url = url.JoinPath(date)
	}

	q := url.Query()
	q.Set("key", apiKey)
	q.Set("unitGroup", "metric")
	url.RawQuery = q.Encode()

	return url.String(), nil
}
