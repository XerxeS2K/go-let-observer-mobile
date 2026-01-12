//go:build android
// +build android

package main

import (
	"log"

	"golang.org/x/mobile/app"
)

func main() {
	log.Println("go-let-observer Android starting")
	app.Main(func(a app.App) {
		NewAndroidApp().Run(a)
	})
}
