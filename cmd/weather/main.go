package weather

import (
	"encoding/json"
	"fmt"
	"github.com/luigizhou/weathercli/internal/api"
	"github.com/luigizhou/weathercli/internal/cache"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "weather",
	Short: "A simple weather application",
	Long:  "A simple command-line weather application that provides location-based weather forecasts.",
	Run:   weather,
}

var city string

func weather(cmd *cobra.Command, args []string) {
	fmt.Println("Welcome to the Weather App!")
	var weatherResp api.WeatherResponse

	if city != "" {
		resp, found := cache.GetValueFromCache(city)
		if found {
			if err := json.Unmarshal([]byte(resp), &weatherResp); err != nil {
				fmt.Errorf(err.Error())
			}
		} else {
			resp, err := api.FetchWeatherData(city)
			if err != nil {
				fmt.Errorf(err.Error())
			}
			cache.SetValueInCache(city, resp)
			if err := json.Unmarshal([]byte(resp), &weatherResp); err != nil {
				fmt.Errorf(err.Error())
			}
		}

		fmt.Printf("We'll try getting the weather for %s\n", city)

		fmt.Println("Weather for City: ", city)
		fmt.Println("Temperature: ", weatherResp.Main.Temp)
		fmt.Println("Humidity: ", weatherResp.Main.Humidity)
		fmt.Println("Wind: ", weatherResp.Wind.Speed)
		fmt.Println("Description: ", weatherResp.Weather[0].Description)
	}

}

func init() {
	RootCmd.PersistentFlags().StringVar(&city, "city", "", "Name of the city for which you want to know the weather")
}
