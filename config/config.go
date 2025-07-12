package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	LastFM struct {
		APIKey   string `json:"api_key"`
		Username string `json:"username"`
	} `json:"lastfm"`
}

func Load() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(homeDir, ".config", "portfolio-tui", "config.json")
	
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return &Config{}, nil
	}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}