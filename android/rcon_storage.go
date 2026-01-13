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
	if err != nil || dir == "" {
		// Fallback: current dir (should still compile even if not writable)
		return "rcon_config.json", nil
	}

	base := filepath.Join(dir, "go-let-observer")
	if mkErr := os.MkdirAll(base, 0o700); mkErr != nil {
		return "", mkErr
	}

	return filepath.Join(base, "rcon_config.json"), nil
}

func LoadRCONConfig() (RCONConfig, error) {
	var cfg RCONConfig

	p, err := rconConfigPath()
	if err != nil {
		return cfg, err
	}

	b, err := os.ReadFile(p)
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return cfg, err
	}

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
