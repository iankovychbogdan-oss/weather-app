package main

import (
	"apiscanner/apiscanner"
	"encoding/json"
	"fmt"
	"net/http"
)

type Weather struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	WeatherDesc []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func getWeather(city string, apiKey string) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching weather:", err)
		return
	}
	defer resp.Body.Close()

	var w Weather
	json.NewDecoder(resp.Body).Decode(&w)

	fmt.Println("City:", w.Name)
	fmt.Println("Temperature:", w.Main.Temp, "ºC")
	fmt.Println("Description:", w.WeatherDesc[0].Description)
}

func main() {
	var city string
	fmt.Print("Enter City: ")
	fmt.Scan(&city)

	apiKey := apiscanner.GetApiKey()
	getWeather(city, apiKey)
}
