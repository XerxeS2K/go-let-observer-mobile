//go:build android
// +build android

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
	base := filepath.Join(dir, "go-let-observer")
	if err := os.MkdirAll(base, 0o755); err != nil {
		return "", err
	}
	return filepath.Join(base, "rcon.json"), nil
}

func LoadRCONConfig() (RCONConfig, error) {
	p, err := rconConfigPath()
	if err != nil {
		return RCONConfig{}, err
	}

	b, err := os.ReadFile(p)
	if err != nil {
		if os.IsNotExist(err) {
			return RCONConfig{}, nil
		}
		return RCONConfig{}, err
	}

	var cfg RCONConfig
	if err := json.Unmarshal(b, &cfg); err != nil {
		return RCONConfig{}, err
	}
	return cfg, nil
}

func SaveRCONConfig(cfg RCONConfig) error {
	p, err := rconConfigPath()
	if err != nil {
		return err
	}

	b, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(p, b, 0o600)
}
