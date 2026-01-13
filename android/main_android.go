//go:build android
// +build android

package main

import (
	"log"

	gomobileapp "github.com/ebitengine/gomobile/app"
	"github.com/hajimehoshi/ebiten/v2/mobile"
)

func main() {
	log.Println("go-let-observer Android starting...")
	gomobileapp.Main(func(a gomobileapp.App) {
		mobile.SetGame(NewAndroidApp())
	})
}
