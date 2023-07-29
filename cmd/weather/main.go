package weather

import (
	"fmt"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "weather",
	Short: "A simple weather application",
	Long:  "A simple command-line weather application that provides location-based weather forecasts.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the Weather App!")
	},
}
