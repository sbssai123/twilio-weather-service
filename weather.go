package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	BASE_API_URL = "http://api.openweathermap.org/data/2.5/weather"
)

type WeatherData struct {
	Temperature     TemperatureData `json:"main"`
	CityName        string          `json:"name"`
	FullDescription []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

type TemperatureData struct {
	CurrentTemperature   float64 `json:"temp,omitempty"`
	FeelsLikeTemperature float64 `json:"feels_like,omitempty"`
	MinTemperature       float64 `json:"temp_min,omitempty"`
	MaxTemperature       float64 `json:"temp_max,omitempty"`
	Humidity             int     `json:"humidity,omitempty"`
}

func getTodaysForecast(cityName string) (WeatherData, error) {

	req, err := http.NewRequest("GET", BASE_API_URL, nil)
	q := req.URL.Query()
	q.Add("q", cityName)
	q.Add("appid", "973a83b25899377e24c6d563fcb84ad2")
	q.Add("units", "imperial")
	req.URL.RawQuery = q.Encode()
	req.Close = true

	resp, err := http.DefaultClient.Do(req)
	var weatherData WeatherData
	if err != nil {
		return weatherData, errors.New("Unable to execute request to Weather API")
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		json.Unmarshal(body, &weatherData)
		return weatherData, nil
	}
}
