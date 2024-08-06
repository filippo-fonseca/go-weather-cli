package main

import (
	"fmt"
	"io"
	"net/http"
)

type Weather struct {
	Location struct{
		Name string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct{
		TempC float64 `json:"temp_c"`
		Condition struct{
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct{
		Forecastday []struct{
			Date string `json:"date"`
			Hour []struct{
				TimeEpoch int64 `json:"time_epoch"`
				TempC float64 `json:"temp_c"`
				Condition struct{
			Text string `json:"text"`
		} `json:"condition"`
		ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=f339f0fe0fa9486da65230533240508&q=London")

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("API request failed")
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}