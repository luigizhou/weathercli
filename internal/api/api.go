package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type WeatherResponse struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

const weatherAPIURL = "https://api.openweathermap.org/data/2.5/weather" // TODO: Make it configurable
const apiKey = ""                                                       // TODO: use VIPER or pass it as parameter to cli

func FetchWeatherData(city string) (string, error) {
	url := fmt.Sprintf("%s?q=%s&appid=%s", weatherAPIURL, city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Errorf(err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("failed to fetch weather data. Status: %s", resp.Status)
		os.Exit(1)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf(err.Error())
		os.Exit(1)
	}
	return string(body), nil
}
