//go:build android
// +build android

package main

import (
	"log"

	"github.com/ebitengine/gomobile/app"
	"github.com/ebitengine/gomobile/mobile"
)

func main() {
	log.Println("go-let-observer Android starting")

	game := NewAndroidGame()

	app.Main(func(_ app.App) {
		mobile.SetGame(game)
	})
}
