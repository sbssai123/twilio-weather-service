package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	weatherApiSecret := os.Getenv("OPEN_WEATHER_API_KEY")
	forecastData, _ := getTodaysForecast("Boston", weatherApiSecret)
	fmt.Printf("%+v\n", forecastData)
	// TODO: Send to Twilio
}
