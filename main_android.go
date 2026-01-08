//go:build android
// +build android

package main

import (
	"log"

	"github.com/ebitengine/gomobile/app"
	"github.com/hajimehoshi/ebiten/v2/mobile"
)

func main() {
	log.Println("go-let-observer Android starting...")
	app.Main(func(a app.App) {
		mobile.SetGame(NewAndroidApp())
	})
}
