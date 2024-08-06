package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/fatih/color"
)

type Weather struct {
	Location struct{
		Name string `json:"name"`
		Region string `json:"region"`
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

// ReplaceWhitespacesWithUnderscore replaces all whitespaces in the input string with underscores.
func ReplaceWhitespacesWithUnderscore(s string) string {
    // Create a new builder to build the resulting string
    var result strings.Builder
    
    // Iterate through each character in the input string
    for _, ch := range s {
        // Check if the character is a whitespace
        if unicode.IsSpace(ch) {
            result.WriteRune('_')
        } else {
            result.WriteRune(ch)
        }
    }
    
    // Return the resulting string
    return result.String()
}

func OutputLocation(args []string) string {
	var q string
	if len(os.Args) >= 2 {
		length := len(os.Args)
		for i := 1; i < length; i++ {
			if (i == 1) {
				q = os.Args[i]
				continue
			}
			q += "_" + os.Args[i]
		}
		fmt.Println(q)
	} else {
		q = "San_Jose_Costa_Rica"
	}		

	return q
}

func main() {
	q := OutputLocation(os.Args)
	
	res, err := http.Get("https://api.weatherapi.com/v1/forecast.json?key=f339f0fe0fa9486da65230533240508&q=" + q + "&days=7")

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

	var weather Weather
	err = json.Unmarshal(body, &weather)

	if err != nil {
		panic(err)
	}

	location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour

	fmt.Printf("%s, %s, %s: %.0fC, %s\n", location.Name, location.Region, location.Country, current.TempC, current.Condition.Text)

	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)

		if date.Before(time.Now()) {
			continue
		}

		message := fmt.Sprintf("%s - %.0fC, %.0f, %s\n", date.Format("15:04"), hour.TempC, hour.ChanceOfRain, hour.Condition.Text)

		if hour.ChanceOfRain < 40 {
			fmt.Print(message)
		} else {
			color.Red(message)
		}
	}
}