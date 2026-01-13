//go:build android || mobile
// +build android mobile

package main

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type RCONConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
}

func rconConfigPath() (string, error) {
	// On Android (gomobile), UserConfigDir normally resolves to app-private storage.
	dir, err := os.UserConfigDir()
	if err != nil || dir == "" {
		// fallback: current dir
		dir = "."
	}
	cfgDir := filepath.Join(dir, "go-let-observer")
	if err := os.MkdirAll(cfgDir, 0o700); err != nil {
		return "", err
	}
	return filepath.Join(cfgDir, "rcon.json"), nil
}

func LoadRCONConfig() (RCONConfig, error) {
	p, err := rconConfigPath()
	if err != nil {
		return RCONConfig{}, err
	}

	b, err := os.ReadFile(p)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return RCONConfig{}, os.ErrNotExist
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
