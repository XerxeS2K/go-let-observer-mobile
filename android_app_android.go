//go:build android
// +build android

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/zMoooooritz/go-let-observer/pkg/ui"
	"github.com/zMoooooritz/go-let-observer/pkg/ui/shared"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

type appMode int

const (
	modeSetup appMode = iota
	modeViewer
)

type AndroidApp struct {
	mode appMode

	setup *RCONSetupUI
	errMsg string

	viewer ebiten.Game
	wrap   *TouchWrapper
}

func NewAndroidApp() ebiten.Game {
	_ = util.InitConfig("") // defaults only

	a := &AndroidApp{
		mode:  modeSetup,
		setup: NewRCONSetupUI(),
	}

	if cfg, err := LoadRCONConfig(); err == nil && cfg != nil {
		a.setup.Host = cfg.Host
		a.setup.Port = cfg.Port
		a.setup.Password = cfg.Password
		if err := a.tryStartViewer(); err == nil {
			a.mode = modeViewer
		} else {
			a.errMsg = err.Error()
		}
	}

	return a
}

func (a *AndroidApp) Update() error {
	switch a.mode {
	case modeSetup:
		a.setup.Update()
		if a.setup.SavePressed {
			a.setup.SavePressed = false
			cfg := &RCONConfig{
				Host:     strings.TrimSpace(a.setup.Host),
				Port:     strings.TrimSpace(a.setup.Port),
				Password: a.setup.Password,
			}
			if err := SaveRCONConfig(cfg); err != nil {
				a.errMsg = "Save failed: " + err.Error()
				return nil
			}
			if err := a.tryStartViewer(); err != nil {
				a.errMsg = err.Error()
				return nil
			}
			a.errMsg = ""
			a.mode = modeViewer
		}
		return nil

	case modeViewer:
		// Settings tap zone (top-left)
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			if x >= 0 && x < 170 && y >= 0 && y < 70 {
				a.mode = modeSetup
				return nil
			}
		}
		return a.wrap.Update()
	}
	return nil
}

func (a *AndroidApp) Draw(screen *ebiten.Image) {
	switch a.mode {
	case modeSetup:
		a.setup.Draw(screen, a.errMsg)
	case modeViewer:
		a.wrap.Draw(screen)
		ebitenutil.DebugPrintAt(screen, "[Settings]", 10, 10)
	}
}

func (a *AndroidApp) Layout(outsideW, outsideH int) (int, int) {
	// Stable internal render size; scales to device.
	return 1000, 720
}

func (a *AndroidApp) tryStartViewer() error {
	host := strings.TrimSpace(a.setup.Host)
	portStr := strings.TrimSpace(a.setup.Port)
	pass := a.setup.Password

	if host == "" {
		return fmt.Errorf("Host is required")
	}
	if portStr == "" {
		return fmt.Errorf("Port is required")
	}
	if _, err := strconv.Atoi(portStr); err != nil {
		return fmt.Errorf("Port must be a number")
	}
	if pass == "" {
		return fmt.Errorf("Password is required")
	}

	util.Config.ServerCredentials.Host = host
	util.Config.ServerCredentials.Port = portStr
	util.Config.ServerCredentials.Password = pass

	// Viewer-only
	viewerMode := shared.MODE_VIEWER

	screenSize := util.Config.UIOptions.ScreenSize
	screenSize = util.Clamp(screenSize, shared.MIN_SCREEN_SIZE, shared.MAX_SCREEN_SIZE)
	util.Config.UIOptions.ScreenSize = screenSize
	util.InitializeFonts(screenSize)

	a.viewer = ui.NewUI(viewerMode)
	a.wrap = NewTouchWrapper(a.viewer)

	return nil
}

type TouchWrapper struct {
	base ebiten.Game

	lastTwo bool
	lastDist float64
	lastCX, lastCY float64
}

func NewTouchWrapper(base ebiten.Game) *TouchWrapper {
	return &TouchWrapper{base: base}
}

func (w *TouchWrapper) Update() error {
	// HUD toggles first (consume taps)
	if w.handleHUDTap() {
		return w.base.Update()
	}

	w.handlePinchZoom()
	return w.base.Update()
}

func (w *TouchWrapper) Draw(screen *ebiten.Image) {
	w.base.Draw(screen)
	w.drawHUD(screen)
}

func (w *TouchWrapper) Layout(wi, he int) (int,int) { return w.base.Layout(wi,he) }

func (w *TouchWrapper) handlePinchZoom() {
	var ids []ebiten.TouchID
	ids = ebiten.AppendTouchIDs(ids)
	if len(ids) < 2 {
		w.lastTwo = false
		return
	}
	x1,y1 := ebiten.TouchPosition(ids[0])
	x2,y2 := ebiten.TouchPosition(ids[1])
	cx := (float64(x1)+float64(x2))/2
	cy := (float64(y1)+float64(y2))/2
	dist := math.Hypot(float64(x2-x1), float64(y2-y1))

	dim := shared.GlobalViewDimension
	if dim == nil {
		w.lastTwo = true
		w.lastDist = dist
		w.lastCX, w.lastCY = cx, cy
		return
	}

	if !w.lastTwo {
		w.lastTwo = true
		w.lastDist = dist
		w.lastCX, w.lastCY = cx, cy
		return
	}
	if w.lastDist <= 0 {
		w.lastDist = dist
		return
	}
	ratio := dist / w.lastDist
	if ratio < 0.90 || ratio > 1.10 {
		oldZoom := dim.ZoomLevel
		newZoom := oldZoom * ratio
		if newZoom < shared.MIN_ZOOM_LEVEL {
			newZoom = shared.MIN_ZOOM_LEVEL
		} else if newZoom > shared.MAX_ZOOM_LEVEL {
			newZoom = shared.MAX_ZOOM_LEVEL
		}
		// keep pinch center stable
		mouseWorldX := (cx - dim.PanX) / oldZoom
		mouseWorldY := (cy - dim.PanY) / oldZoom
		dim.ZoomLevel = newZoom
		dim.PanX -= mouseWorldX * (dim.ZoomLevel - oldZoom)
		dim.PanY -= mouseWorldY * (dim.ZoomLevel - oldZoom)

		if dim.ZoomLevel == shared.MIN_ZOOM_LEVEL {
			dim.PanX = 0
			dim.PanY = 0
		}
		// clamp
		dim.PanX = util.Clamp(dim.PanX, float64(dim.SizeX)*(shared.MIN_ZOOM_LEVEL-dim.ZoomLevel), 0)
		dim.PanY = util.Clamp(dim.PanY, float64(dim.SizeY)*(shared.MIN_ZOOM_LEVEL-dim.ZoomLevel), 0)

		w.lastDist = dist
	}
	w.lastCX, w.lastCY = cx, cy
}

func (w *TouchWrapper) handleHUDTap() bool {
	if !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return false
	}
	x,y := ebiten.CursorPosition()

	btnW, btnH := 220, 60
	startX := 1000 - btnW - 10
	startY := 80

	buttons := []struct{
		label string
		row int
		on *bool
	}{
		{"Players",0,&util.Config.UIOptions.ShowPlayers},
		{"Info",1,&util.Config.UIOptions.ShowPlayerInfo},
		{"Spawns",2,&util.Config.UIOptions.ShowSpawns},
		{"Tanks",3,&util.Config.UIOptions.ShowTanks},
		{"Grid",4,&util.Config.UIOptions.ShowGridOverlay},
		{"Header",5,&util.Config.UIOptions.ShowServerInfoOverlay},
	}

	for _, b := range buttons {
		bx := startX
		by := startY + b.row*(btnH+10)
		if hit(x,y,bx,by,btnW,btnH) {
			*b.on = !*b.on
			return true
		}
	}
	return false
}

func (w *TouchWrapper) drawHUD(screen *ebiten.Image) {
	// lightweight text HUD
	ebitenutil.DebugPrintAt(screen, "Touch HUD", 820, 55)
	btnW, btnH := 220, 60
	startX := 1000 - btnW - 10
	startY := 80

	draw := func(row int, label string, on bool) {
		bx := startX
		by := startY + row*(btnH+10)
		state := "OFF"
		if on { state = "ON" }
		ebitenutil.DebugPrintAt(screen, label+" ["+state+"]", bx+12, by+20)
	}
	draw(0,"Players", util.Config.UIOptions.ShowPlayers)
	draw(1,"Info", util.Config.UIOptions.ShowPlayerInfo)
	draw(2,"Spawns", util.Config.UIOptions.ShowSpawns)
	draw(3,"Tanks", util.Config.UIOptions.ShowTanks)
	draw(4,"Grid", util.Config.UIOptions.ShowGridOverlay)
	draw(5,"Header", util.Config.UIOptions.ShowServerInfoOverlay)
}
