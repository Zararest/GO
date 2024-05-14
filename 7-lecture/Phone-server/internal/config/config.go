package config

import (
	"flag"
	"os"
	"io"
	"fmt"
	"gopkg.in/yaml.v3"
)

type LoggerConfig struct {
	Level uint `yaml:"level"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
}

type StorageConfig struct {
	Path string `yaml:"path"`
}

type GlobalConfig struct {
	Log    	LoggerConfig 	`yaml:"logger"`
	Server 	ServerConfig 	`yaml:"server"`
	Storage	StorageConfig	`yaml:"storage"`
}

var (
	ConfigPathFlag = flag.String("config", "configs/default.yaml", "Path to YAML config")
)

func init() {
	flag.Parse()
}

func Unmarshal(r io.Reader) (GlobalConfig, error) {
	var cfg GlobalConfig
	if err := yaml.NewDecoder(r).Decode(&cfg); err != nil {
		return GlobalConfig{}, fmt.Errorf("parse config error: %w", err)
	}

	return cfg, nil
}

func GetParameters() (GlobalConfig, error) {
	configFile, err := os.Open(*ConfigPathFlag)
	if err != nil {
		return GlobalConfig{}, err
	}

	cfg, err := Unmarshal(configFile)
	if err != nil {
		return GlobalConfig{}, err
	}
	
	return cfg, nil
}
