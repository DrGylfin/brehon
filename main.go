package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2/pkg/options/mac"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed frontend/src
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {

	// Create an instance of the app structure
	app := NewApp()

	// Struct for file handling
	// TODO: populate struct's values from preferences or defaults
	fileHandler := &FileHandler{}

	// Create application with options
	err := wails.Run(&options.App{
		Title:             "brehon",
		Width:             1280,
		Height:            740,
		MinWidth:          720,
		MinHeight:         570,
		MaxWidth:          1280,
		MaxHeight:         740,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false,
		RGBA:              &options.RGBA{R: 33, G: 37, B: 43, A: 255},
		Assets:            assets,
		LogLevel:          logger.DEBUG,
		OnStartup:         app.startup,
		OnDomReady:        app.domReady,
		OnShutdown:        app.shutdown,
		Bind: []interface{}{
			app, fileHandler,
		},
		// Windows platform specific options
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
		},
		Mac: &mac.Options{
			TitleBar:             mac.TitleBarHiddenInset(),
			WebviewIsTransparent: false,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "Brehon",
				Message: "A to-do list tracker and chooser",
				Icon:    icon,
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
