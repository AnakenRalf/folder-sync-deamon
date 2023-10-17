// config/config.go
package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	SourcePath  string `json:"source_path"`
	NetworkPath string `json:"network_path"`
	// Add other configuration fields as needed
}

func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Config{}
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
