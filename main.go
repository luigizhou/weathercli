package main

import (
	"fmt"
	"github.com/luigizhou/weathercli/cmd/weather"
	"os"
)

func main() {
	if err := weather.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
