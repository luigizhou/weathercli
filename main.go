package main

import (
	"fmt"
	"github.com/luigizhou/weathercli/cmd/weather"
	"os"
)

const apiKey = "413a0c18def7e7c6c11ecd30c1f7305a"
const API_ADDRESS = "http://api.openweathermap.org/data/2.5/forecast?id=524901&appid="

// WeatherResponse represents the JSON response from the weather API.

func main() {
	//city := "London"
	//url := fmt.Sprintf("%s?q=%s&appid=%s", weatherAPIURL, city, apiKey)
	//resp, err := http.Get(url)
	//if err != nil {
	//	fmt.Errorf(err.Error())
	//	os.Exit(1)
	//}
	//defer resp.Body.Close()
	//
	//if resp.StatusCode != http.StatusOK {
	//	fmt.Errorf("failed to fetch weather data. Status: %s", resp.Status)
	//	os.Exit(1)
	//}
	//
	//body, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Errorf(err.Error())
	//	os.Exit(1)
	//}
	//
	//var weatherResp WeatherResponse
	//if err := json.Unmarshal(body, &weatherResp); err != nil {
	//	fmt.Errorf(err.Error())
	//	os.Exit(1)
	//}
	//
	if err := weather.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
