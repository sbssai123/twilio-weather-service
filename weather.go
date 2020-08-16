package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	BASE_API_URL = "http://api.openweathermap.org/data/2.5/weather"
)

func getTodaysForecast(cityName string) {

	req, err := http.NewRequest("GET", BASE_API_URL, nil)
	q := req.URL.Query()
	q.Add("q", cityName)
	q.Add("appid", "<API_KEY>")
	req.URL.RawQuery = q.Encode()
	req.Close = true

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Request failed with error %s \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(data))
	}
}
