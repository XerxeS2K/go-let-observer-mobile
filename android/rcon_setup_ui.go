//go:build android
// +build android

package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type focusField int

const (
	focusHost focusField = iota
	focusPort
	focusPass
)

type RCONSetupUI struct {
	Host     string
	Port     string
	Password string

	SavePressed bool

	focus focusField
}

func NewRCONSetupUI() *RCONSetupUI {
	return &RCONSetupUI{focus: focusHost}
}

func (u *RCONSetupUI) Update() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		if hit(x, y, 60, 160, 880, 80) {
			u.focus = focusHost
		} else if hit(x, y, 60, 270, 880, 80) {
			u.focus = focusPort
		} else if hit(x, y, 60, 380, 880, 80) {
			u.focus = focusPass
		} else if hit(x, y, 60, 500, 880, 90) {
			u.SavePressed = true
		}
	}

	var chars []rune
	chars = ebiten.AppendInputChars(chars)
	for _, r := range chars {
		if r == '\n' || r == '\r' || r == '\t' {
			continue
		}
		switch u.focus {
		case focusHost:
			u.Host += string(r)
		case focusPort:
			if r >= '0' && r <= '9' {
				u.Port += string(r)
			}
		case focusPass:
			// password visible (user choice A)
			u.Password += string(r)
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
		switch u.focus {
		case focusHost:
			u.Host = backspace(u.Host)
		case focusPort:
			u.Port = backspace(u.Port)
		case focusPass:
			u.Password = backspace(u.Password)
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		u.SavePressed = true
	}
}

func (u *RCONSetupUI) Draw(screen *ebiten.Image, errMsg string) {
	ebitenutil.DebugPrintAt(screen, "RCON Setup", 60, 60)
	ebitenutil.DebugPrintAt(screen, "Type Host / Port / Password. Saved locally on this device.", 60, 95)

	drawField(screen, 60, 160, 880, 80, "Host", u.Host, u.focus == focusHost)
	drawField(screen, 60, 270, 880, 80, "Port", u.Port, u.focus == focusPort)
	drawField(screen, 60, 380, 880, 80, "Password", u.Password, u.focus == focusPass)

	drawButton(screen, 60, 500, 880, 90, "Save & Connect")

	if errMsg != "" {
		ebitenutil.DebugPrintAt(screen, errMsg, 60, 620)
	}
}

func backspace(s string) string {
	r := []rune(s)
	if len(r) == 0 {
		return s
	}
	return string(r[:len(r)-1])
}

func hit(x, y, bx, by, bw, bh int) bool {
	return x >= bx && x < bx+bw && y >= by && y < by+bh
}

func drawField(screen *ebiten.Image, x, y, w, h int, label, value string, focused bool) {
	bg := color.RGBA{30, 30, 35, 230}
	if focused {
		bg = color.RGBA{50, 50, 70, 240}
	}
	ebitenutil.DrawRect(screen, float64(x), float64(y), float64(w), float64(h), bg)
	ebitenutil.DebugPrintAt(screen, label+": "+value, x+16, y+28)
}

func drawButton(screen *ebiten.Image, x, y, w, h int, label string) {
	ebitenutil.DrawRect(screen, float64(x), float64(y), float64(w), float64(h), color.RGBA{0, 0, 0, 160})
	ebitenutil.DebugPrintAt(screen, label, x+16, y+32)
}
