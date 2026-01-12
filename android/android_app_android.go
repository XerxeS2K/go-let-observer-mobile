//go:build android
// +build android

package android

import (
	"log"

	"github.com/ebitengine/gomobile/app"
	"github.com/ebitengine/gomobile/mobile"
)

// AndroidApp is the gomobile entry wrapper.
// It does NOT have Run(), Update(), etc.
// Ebiten handles the lifecycle internally.
type AndroidApp struct {
	game mobile.Game
}

// NewAndroidApp creates the Android Ebiten app wrapper.
func NewAndroidApp(game mobile.Game) *AndroidApp {
	return &AndroidApp{
		game: game,
	}
}

// Start registers the game with gomobile/app.
// This MUST be called from main().
func (a *AndroidApp) Start() {
	log.Println("AndroidApp starting (Ebiten gomobile)")
	app.Main(func(_ app.App) {
		mobile.SetGame(a.game)
	})
}
