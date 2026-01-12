//go:build android
// +build android

package main

import (
	"log"

	"golang.org/x/mobile/app"
)

type AndroidApp struct{}

func NewAndroidApp() *AndroidApp {
	return &AndroidApp{}
}

func (a *AndroidApp) Run(app.App) {
	log.Println("Android app running (touch-only stub)")
	select {} // keep app alive
}
