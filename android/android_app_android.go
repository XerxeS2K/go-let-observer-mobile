//go:build android || mobile
// +build android mobile

package main

import "github.com/hajimehoshi/ebiten/v2"

type AndroidApp struct{}

func NewAndroidApp() *AndroidApp {
	return &AndroidApp{}
}

// Ebiten Game interface
func (a *AndroidApp) Update() error {
	return nil
}

func (a *AndroidApp) Draw(screen *ebiten.Image) {
}

func (a *AndroidApp) Layout(w, h int) (int, int) {
	return w, h
}
