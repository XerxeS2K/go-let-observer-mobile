//go:build android
// +build android

package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// RCONSetupUI is a minimal placeholder so the Android folder can evolve
// without blocking builds. It is not currently wired into the main UI.
type RCONSetupUI struct {
	Done   bool
	Config RCONConfig
}

func NewRCONSetupUI() *RCONSetupUI {
	return &RCONSetupUI{
		Done: false,
		Config: RCONConfig{
			Host:     "",
			Port:     "",
			Password: "",
		},
	}
}

func (u *RCONSetupUI) Update() error {
	// Placeholder: no-op
	return nil
}

func (u *RCONSetupUI) Draw(screen *ebiten.Image) {
	// Placeholder: no-op
	_ = screen
}

func (u *RCONSetupUI) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
