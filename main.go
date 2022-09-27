package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"os"

	"github.com/LouisBrunner/lemmy/backend/application"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed wails.json
var rawConfig []byte

type wailsConfig struct {
	Name string `json:"name"`
}

func work() error {
	app := application.New()

	config := wailsConfig{}
	err := json.Unmarshal(rawConfig, &config)
	if err != nil {
		return fmt.Errorf("failed to parse internal config: %w", err)
	}

	return wails.Run(&options.App{
		Title:            config.Name,
		Width:            1024,
		Height:           768,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.OnStartup,
		Bind: []interface{}{
			app.Bindings(),
		},
	})
}

func main() {
	err := work()
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
