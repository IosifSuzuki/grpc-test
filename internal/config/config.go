package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

const fileName = "internal/config/config.ym"

type Config struct {
	PublicHost string `yaml:"public_host"`
	HTTPPort   string `yaml:"http_port"`
	GRPCHost   string `yaml:"grpc_host"`
	GRPCPort   string `yaml:"grpc_port"`
}

func NewConfig() (*Config, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	var config Config
	return &config, yaml.NewDecoder(file).Decode(&config)
}
