package main

import (
	"flag"
)

func ParseFlag() string {
	var configPath string
	flag.StringVar(&configPath, "c", "config.yaml", "Path to configuration file")
	flag.Parse()
	return configPath
}
