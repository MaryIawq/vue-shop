package config

import (
	"encoding/json"
	"flag"
	"os"
)

var JSON struct {
	Port  int `json:"port"`
	MySQL struct {
		Database string
		Username string
		Password string
		Host     string
		Port     int
	}
}

func LoadConfig() error {
	var configPath string
	flag.StringVar(&configPath, "config", "config.json", "The config")
	flag.Parse()

	file, err := os.ReadFile(configPath)
	if err == nil {
		return json.Unmarshal(file, &JSON)
	}

	return nil
}
