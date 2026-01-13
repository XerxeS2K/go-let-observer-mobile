//go:build android
// +build android

package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/zMoooooritz/go-let-observer/pkg/ui"
	"github.com/zMoooooritz/go-let-observer/pkg/ui/shared"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

func NewAndroidApp() *ui.UI {
	log.Println("NewAndroidApp(): initializing config and UI")

	// Try to load a persisted config if it exists, otherwise fall back to defaults.
	if cfgDir, err := os.UserConfigDir(); err == nil && cfgDir != "" {
		cfgPath := filepath.Join(cfgDir, "go-let-observer", "config.yaml")
		if _, statErr := os.Stat(cfgPath); statErr == nil {
			if err := util.InitConfig(cfgPath); err != nil {
				log.Printf("InitConfig(%s) failed, using defaults: %v", cfgPath, err)
				_ = util.InitConfig("")
			}
		} else {
			_ = util.InitConfig("")
		}
	} else {
		_ = util.InitConfig("")
	}

	// Viewer mode (same intent as desktop viewer, but mobile-driven)
	return ui.NewUI(shared.MODE_VIEWER)
}
