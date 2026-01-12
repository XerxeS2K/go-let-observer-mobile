//go:build android || mobile
// +build android mobile

package main

import (
	"log"

	"github.com/ebitengine/gomobile/app"
	"github.com/ebitengine/gomobile/mobile"
)

func main() {
	log.Println("go-let-observer Android starting")

	app.Main(func(a app.App) {
		game := NewAndroidApp()
		mobile.SetGame(game)
	})
}
