package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"os"

	"github.com/LouisBrunner/lemmy/backend/application"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

//go:embed wails.json
var rawConfig []byte

type wailsConfig struct {
	Name   string      `json:"name"`
	Author wailsAuthor `json:"author"`
}

type wailsAuthor struct {
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
		Title:     config.Name,
		MinWidth:  800,
		MinHeight: 600,
		Assets:    assets,
		OnStartup: app.OnStartup,
		Bind: []interface{}{
			app.Bindings(),
		},
		Mac: &mac.Options{
			TitleBar:             mac.TitleBarHiddenInset(),
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   config.Name,
				Message: fmt.Sprintf("© Copyright %s", config.Author.Name),
				Icon:    icon,
			},
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
