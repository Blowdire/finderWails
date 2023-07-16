package main

import (
	"embed"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/go-vgo/robotgo"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS
var isKeyPressed bool

// Function used to check key combination to raise the app
func listenKeyCombination(app *App) {
	fmt.Println("Listening for key combination")
	//if Ctrl-q is pressed raise window
	for {
		if ok := robotgo.AddEvents("q", "ctrl"); ok {
			fmt.Println("Pressed ")
			runtime.WindowShow(app.ctx)
		}

		robotgo.MilliSleep(100)
	}

}

func main() {
	// Create an instance of the app structure
	app := NewApp()
	go listenKeyCombination(app)
	// Create application with options
	err := wails.Run(&options.App{
		Title:  "finder",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
