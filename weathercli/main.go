package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
	"strings"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`

	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`

	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				TempC     float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}
	apiKey := os.Getenv("API_KEY")

	var city string
	fmt.Println("Enter the city you would like to query:")
	fmt.Scanln(&city)

	city = strings.TrimSpace(city)
	city = strings.ReplaceAll(city, " ", "%20")

	url := fmt.Sprintf("https://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=1&aqi=no&alerts=no", apiKey, city)
	//res, err := http.Get("https://api.weatherapi.com/v1/forecast.json?key=3794d64bae224dcb81a150812253004&q=Ankara&days=1&aqi=no&alerts=no")
	/*
		res, err := http.Get("https://api.weatherapi.com/v1/forecast.json?key=3794d64bae224dcb81a150812253004&q=Duzce&days=1&aqi=no&alerts=no")
		if err != nil {
			panic(err)
		}*/
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("weather api not available")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}
	/*location, current, _ := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour

	fmt.Printf("%s, %s: %.0fC, %s \n",
		location.Name,
		location.Country,
		current.TempC,
		current.Condition.Text,
	)*/
	fmt.Printf("\n%s, %s weather forecast:\n%.1f°C, %s\n",
		weather.Location.Name,
		weather.Location.Country,
		weather.Current.TempC,
		weather.Current.Condition.Text,
	)

}
