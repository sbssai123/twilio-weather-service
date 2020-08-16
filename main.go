package main

import "fmt"

func main() {
	forecastData, _ := getTodaysForecast("Boston")
	fmt.Printf("%+v\n", forecastData)
	// TODO: Send to Twilio
}
