package main

import (
	"os"

	"github.com/krishpranav/videostream/data"
	"github.com/krishpranav/videostream/net"
)

func initConfig() *data.Config {
	args := os.Args
	if len(args) > 1 {

	}

	defaultconfig := data.DefaultConfig()
	return defaultConfig
}

func main() {
	c := initConfig()
	net.RunServer(c)
}
