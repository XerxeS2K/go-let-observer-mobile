//go:build android || mobile
// +build android mobile

package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// AndroidApp is the root Ebiten game for Android.
// It MUST implement ebiten.Game.
type AndroidApp struct {
	width  int
	height int
}

// NewAndroidApp creates the Android game instance.
// This is what mobile.SetGame(...) receives.
func NewAndroidApp() *AndroidApp {
	log.Println("NewAndroidApp initialised")
	return &AndroidApp{
		width:  1280,
		height: 720,
	}
}

// Update is called every tick by Ebiten.
func (a *AndroidApp) Update() error {
	// TODO:
	// - Touch handling
	// - RCON UI state updates
	// - Networking / polling
	return nil
}

// Draw renders the frame.
func (a *AndroidApp) Draw(screen *ebiten.Image) {
	// TODO:
	// - Draw map
	// - Draw UI controls
	// - Draw overlays / debug text
}

// Layout tells Ebiten the logical resolution.
func (a *AndroidApp) Layout(outsideWidth, outsideHeight int) (int, int) {
	if outsideWidth > 0 && outsideHeight > 0 {
		a.width = outsideWidth
		a.height = outsideHeight
	}
	return a.width, a.height
}
