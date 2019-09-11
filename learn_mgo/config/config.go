package config

import (
	"encoding/json"
	"os"
)

type MgoConfig struct {
	Host       string `json:"host"`
	Port       int    `json:"port"`
	Database   string `json:"database"`
	Collection string `json:"collection"`
}

func LoadConfig(path string, c interface{}) error {
	source, err := os.Open(path)
	if err != nil {
		return os.ErrNotExist
	}
	return json.NewDecoder(source).Decode(c)
}
