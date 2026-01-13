//go:build android || mobile
// +build android mobile

package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/zMoooooritz/go-let-observer/pkg/ui"
	"github.com/zMoooooritz/go-let-observer/pkg/ui/shared"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

// NewAndroidApp creates the Ebiten UI in viewer mode for Android.
// It also ensures we have a writable per-app config location (for future use).
func NewAndroidApp() *ui.UI {
	log.Println("Android app init...")

	// Ensure app config directory exists (we'll use this later for RCON config files etc.)
	if dir, err := os.UserConfigDir(); err == nil && dir != "" {
		_ = os.MkdirAll(filepath.Join(dir, "go-let-observer"), 0o700)
	}

	// Load normal app config (safe even if no config file exists)
	// If you later want to force a specific config file path, pass it here.
	_ = util.InitConfig("")

	// Start UI in viewer mode (same pattern as desktop, but touch driven)
	return ui.NewUI(shared.MODE_VIEWER)
}
