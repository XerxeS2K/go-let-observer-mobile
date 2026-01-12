//go:build android || mobile
// +build android mobile

package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type RCONConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
}

func rconConfigPath() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	appDir := filepath.Join(dir, "go-let-observer")
	err = os.MkdirAll(appDir, 0700)
	if err != nil {
		return "", err
	}

	return filepath.Join(appDir, "rcon.json"), nil
}

func SaveRCONConfig(cfg RCONConfig) error {
	path, err := rconConfigPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0600)
}

func LoadRCONConfig() (RCONConfig, error) {
	var cfg RCONConfig

	path, err := rconConfigPath()
	if err != nil {
		return cfg, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}

	err = json.Unmarshal(data, &cfg)
	return cfg, err
}
