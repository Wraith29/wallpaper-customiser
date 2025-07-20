package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

const baseUrl = "http://api.weatherapi.com/v1"

type Weather string

const (
	Sun  Weather = "Sun"
	Rain Weather = "Rain"
)

// TODO: Make this actually check the code properly
// That does rely on having more than just sunny / rainy weather states
// Weather Conditions Map: https://www.weatherapi.com/docs/weather_conditions.json
func WeatherFromCode(code int) Weather {
	if code == 1000 {
		return Sun
	}
	return Rain
}

type WeatherClient struct {
	ApiKey, Latitude, Longitude string
}

func NewWeatherClient() (*WeatherClient, error) {
	apiKey := os.Getenv("WC_WEATHER_API_KEY")
	if apiKey == "" {
		return nil, errors.New("missing required env var")
	}

	lat := os.Getenv("WC_LATITUDE")
	if lat == "" {
		return nil, errors.New("missing required env var")
	}
	long := os.Getenv("WC_LONGITUDE")
	if long == "" {
		return nil, errors.New("missing required env var")
	}

	return &WeatherClient{
		ApiKey:    apiKey,
		Latitude:  lat,
		Longitude: long,
	}, nil
}

type condition struct {
	Code int `json:"code"`
}

type current struct {
	Condition condition `json:"condition"`
}

type weatherApiResponse struct {
	Current current `json:"current"`
}

func (w *WeatherClient) GetWeather() (Weather, error) {
	apiClient := http.Client{}

	endpoint := fmt.Sprintf(
		"%s/current.json?key=%s&q=auto:ip",
		baseUrl, w.ApiKey,
	)

	res, err := apiClient.Get(endpoint)
	if err != nil {
		return "", err
	}

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var response weatherApiResponse
	if err := json.Unmarshal(buf, &response); err != nil {
		return "", err
	}

	return WeatherFromCode(response.Current.Condition.Code), nil
}
