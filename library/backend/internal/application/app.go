package application

import (
	"context"

	"github.com/LouisBrunner/lemmy/library/backend/api"
	"github.com/LouisBrunner/lemmy/library/backend/internal/contracts"
	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type app[Bindings, UserConfig any] struct {
	ctx    context.Context
	logger *logrus.Logger
	config contracts.ConfigManager[UserConfig]
	user   api.App[Bindings]
}

func New[Bindings, UserConfig any](logger *logrus.Logger, config contracts.ConfigManager[UserConfig], user api.App[Bindings]) *app[Bindings, UserConfig] {
	return &app[Bindings, UserConfig]{
		logger: logger,
		config: config,
		user:   user,
	}
}

func (me *app[Bindings, UserConfig]) OnStartup(ctx context.Context) {
	me.ctx = ctx

	appConfig := me.config.Internal()
	if appConfig.LastX > 0 && appConfig.LastY > 0 {
		me.logger.Debugf("changing window position to: %vx%v", appConfig.LastX, appConfig.LastY)
		runtime.WindowSetPosition(me.ctx, appConfig.LastX, appConfig.LastY)
	}

	me.user.OnStartup(me.ctx)
}

func (me *app[Bindings, UserConfig]) OnDomReady(ctx context.Context) {
	me.user.OnDomReady(ctx, me.ctx)
}
func (me *app[Bindings, UserConfig]) OnBeforeClose(ctx context.Context) bool {
	return me.user.OnBeforeClose(ctx, me.ctx)
}

func (me *app[Bindings, UserConfig]) OnShutdown(ctx context.Context) {
	width, height := runtime.WindowGetSize(me.ctx)
	x, y := runtime.WindowGetPosition(me.ctx)
	err := me.config.SaveInternal(contracts.ConfigInternal{
		LastWidth:  width,
		LastHeight: height,
		LastX:      x,
		LastY:      y,
	})
	if err != nil {
		me.logger.Errorf("failed to write config: %v", err)
	}

	me.user.OnShutdown(ctx, me.ctx)
}
