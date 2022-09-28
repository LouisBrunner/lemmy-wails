package mywails

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	Name   string `json:"name"`
	Author Author `json:"author"`
}

type Author struct {
	Name string `json:"name"`
}

func GetConfig(rawConfig []byte) (*Config, error) {
	config := Config{}
	err := json.Unmarshal(rawConfig, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse internal config: %w", err)
	}
	return &config, nil
}
