package api

import (
	"context"
	"embed"

	"github.com/sirupsen/logrus"
)

type ConfigManager[User any] interface {
	Load() (User, bool)
	Save(config User) error
}

type App[Bindings any] interface {
	OnStartup(ctx context.Context)
	OnDomReady(ctx context.Context, wailsCtx context.Context)
	OnBeforeClose(ctx context.Context, wailsCtx context.Context) (prevent bool)
	OnShutdown(ctx context.Context, wailsCtx context.Context)
	Bindings() Bindings
}

type AppMaker[Bindings, Config any] func(logger *logrus.Logger, cfg ConfigManager[Config]) (App[Bindings], error)

type Options[Bindings, Config any] struct {
	WailsJSON []byte
	Assets    embed.FS
	Icon      []byte
	AppMaker  AppMaker[Bindings, Config]
}
