package main

import (
	"fmt"
	"net/http"
)

func main() {
	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=f339f0fe0fa9486da65230533240508&q=London")

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("API request failed")
	}
	fmt.Println("ready to go")
}