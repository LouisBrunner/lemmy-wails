package application

import (
	"context"

	"github.com/LouisBrunner/lemmy/backend/bindings"
	"github.com/LouisBrunner/lemmy/backend/contracts"
	"github.com/LouisBrunner/lemmy/library/backend/api"
	"github.com/sirupsen/logrus"
)

type Config struct{}

type app struct {
	logger *logrus.Logger
}

func New(logger *logrus.Logger, cfg api.ConfigManager[Config]) api.App[contracts.Bindings] {
	return &app{
		logger: logger,
	}
}

func (me *app) OnStartup(ctx context.Context) {
}

func (me *app) OnDomReady(ctx context.Context, wailsCtx context.Context) {
}

func (me *app) OnBeforeClose(ctx context.Context, wailsCtx context.Context) (prevent bool) {
	return false
}

func (me *app) OnShutdown(ctx context.Context, wailsCtx context.Context) {
}

func (me *app) Bindings() contracts.Bindings {
	return bindings.New()
}
