package weather

import (
	"github.com/luigizhou/weathercli/internal/api"
	"github.com/spf13/cobra"
)

var cityCmd = &cobra.Command{
	Use:   "city [city-name]",
	Short: "Get weather by city",
	Long:  "Get the weather forecast for a specific city.",
	Args:  cobra.ExactArgs(1), // The command expects exactly one argument (city name)
	Run: func(cmd *cobra.Command, args []string) {
		city := args[0]
		api.FetchWeatherData(city)
	},
}

func init() {
	// Add the 'city' subcommand to the RootCmd
	RootCmd.AddCommand(cityCmd)
}
