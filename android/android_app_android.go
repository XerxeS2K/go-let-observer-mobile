//go:build android || mobile
// +build android mobile

package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zMoooooritz/go-let-observer/pkg/ui"
	"github.com/zMoooooritz/go-let-observer/pkg/ui/shared"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

// NewAndroidApp creates the Ebiten Game used by gomobile.
// mobile.SetGame(...) will drive Update/Draw/Layout automatically.
func NewAndroidApp() ebiten.Game {
	// Match the desktop default config path.
	// This file exists in your repo at configs/config.yaml.
	//
	// If you later want an Android-specific config path, we can add it,
	// but this keeps behavior consistent with the upstream app.
	const configPath = "configs/config.yaml"

	if err := util.InitConfig(configPath); err != nil {
		// Don’t panic; log and continue. The UI might still show something useful.
		log.Printf("InitConfig(%s) failed: %v", configPath, err)
	}

	// PresentationMode in this project means viewer/replay/record etc (not “touch mode”).
	// Android touch support is handled by Ebiten’s mobile input automatically.
	game := ui.NewUI(shared.MODE_VIEWER)

	// Optional safety: ensure we satisfy ebiten.Game at compile time
	var _ ebiten.Game = game

	return game
}
