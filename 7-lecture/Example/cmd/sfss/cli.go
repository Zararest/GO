package main

import "flag"

var (
	// ConfigPathFlag is flag with path to yaml config file.
	ConfigPathFlag = flag.String("config", "configs/default.yaml", "Path to YAML config")
)

func init() {
	flag.Parse()
}
