//go:build android || mobile
// +build android mobile

package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type AndroidApp struct {
	width  int
	height int
}

func NewAndroidApp() *AndroidApp {
	return &AndroidApp{}
}

func (a *AndroidApp) Update() error {
	return nil
}

func (a *AndroidApp) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{20, 20, 20, 255})
}

func (a *AndroidApp) Layout(outsideWidth, outsideHeight int) (int, int) {
	a.width = outsideWidth
	a.height = outsideHeight
	return outsideWidth, outsideHeight
}
