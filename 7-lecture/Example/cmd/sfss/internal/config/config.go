package config

import (
	"fmt"
	"io"

	"gitlab.com/lgp/http-server-example/internal/repository/file"
	"gitlab.com/lgp/http-server-example/internal/server"
	"gopkg.in/yaml.v3"
)

// Config of app.
type Config struct {
	Server         server.Config             `yaml:"server"`
	FileRepository file.FileRepositoryConfig `yaml:"file"`
}

// Unmarshal app config using file reader.
func Unmarshal(r io.Reader) (Config, error) {
	var cfg Config
	if err := yaml.NewDecoder(r).Decode(&cfg); err != nil {
		return Config{}, fmt.Errorf("parse config error: %w", err)
	}

	return cfg, nil
}
