//go:build android || mobile
// +build android mobile

package main

import "github.com/ebitengine/gomobile/app"

func main() {
	app.Main(func(a app.App) {
		NewAndroidApp().Run(a)
	})
}
