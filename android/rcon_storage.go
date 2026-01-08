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
	dir = filepath.Join(dir, "go-let-observer")
	return filepath.Join(dir, "rcon.json"), nil
}

func LoadRCONConfig() (*RCONConfig, error) {
	p, err := rconConfigPath()
	if err != nil {
		return nil, err
	}
	b, err := os.ReadFile(p)
	if err != nil {
		return nil, err
	}
	var cfg RCONConfig
	if err := json.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func SaveRCONConfig(cfg *RCONConfig) error {
	p, err := rconConfigPath()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(p), 0o700); err != nil {
		return err
	}
	b, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(p, b, 0o600)
}
