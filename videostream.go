package main

import (
	"log"
	"os"

	"github.com/krishpranav/videostream/data"
	"github.com/krishpranav/videostream/net"
)

func initConfig() *data.Config {
	args := os.Args
	if len(args) > 1 {
		configPath := args[1]

		customConfig, err := data.ParseJsonConfig(configPath)
		if err != nil {
			log.Fatalln("Failed to load config at ", configPath, err)
		}

		return customConfig
	}

	defaultConfig := data.DefaultConfig()
	log.Println("Using default config ", defaultConfig)
	return defaultConfig
}

func main() {
	c := initConfig()
	net.RunServer(c)
}
