package internal

import (
	"fmt"

	"github.com/LouisBrunner/lemmy/library/backend/api"
	"github.com/LouisBrunner/lemmy/library/backend/internal/application"
	"github.com/LouisBrunner/lemmy/library/backend/internal/config"
	"github.com/LouisBrunner/lemmy/library/backend/internal/logging"
	"github.com/LouisBrunner/lemmy/library/backend/internal/mywails"
	"github.com/emersion/go-appdir"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

type worker[Bindings, Config any] struct {
	opts api.Options[Bindings, Config]
}

func NewWorker[Bindings, Config any](opts api.Options[Bindings, Config]) *worker[Bindings, Config] {
	return &worker[Bindings, Config]{
		opts: opts,
	}
}

func (me *worker[Bindings, Config]) Work() error {
	wailsConfig, err := mywails.GetConfig(me.opts.WailsJSON)
	if err != nil {
		return err
	}

	dirs := appdir.New(wailsConfig.Name)

	log := logging.NewLogger(dirs.UserLogs())
	cfgMng := config.NewManager[Config](log, dirs.UserConfig())

	userApp, err := me.opts.AppMaker(log, cfgMng)
	if err != nil {
		return fmt.Errorf("user app failed to initialize: %w", err)
	}

	app := application.New[Bindings, Config](log, cfgMng, userApp)
	appConfig := cfgMng.Internal()

	return wails.Run(&options.App{
		Title:              wailsConfig.Name,
		MinWidth:           800,
		MinHeight:          600,
		Width:              appConfig.LastWidth,
		Height:             appConfig.LastHeight,
		Assets:             me.opts.Assets,
		LogLevel:           logger.DEBUG,
		LogLevelProduction: logger.INFO,
		OnStartup:          app.OnStartup,
		OnDomReady:         app.OnDomReady,
		OnBeforeClose:      app.OnBeforeClose,
		OnShutdown:         app.OnShutdown,
		Logger:             logging.NewAdapter(log),
		Bind: []interface{}{
			userApp.Bindings(),
		},
		Mac: &mac.Options{
			TitleBar:             mac.TitleBarHiddenInset(),
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   wailsConfig.Name,
				Message: fmt.Sprintf("© Copyright %s", wailsConfig.Author.Name),
				Icon:    me.opts.Icon,
			},
		},
		Linux: &linux.Options{
			Icon: me.opts.Icon,
		},
	})
}
