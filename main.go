package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/LouisBrunner/lemmy/backend/application"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := application.New()

	err := wails.Run(&options.App{
		Title:            "lemmy",
		Width:            1024,
		Height:           768,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.OnStartup,
		Bind: []interface{}{
			app.Bindings(),
		},
	})
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
