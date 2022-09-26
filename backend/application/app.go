package application

import (
	"context"

	"github.com/LouisBrunner/lemmy/backend/bindings"
	"github.com/LouisBrunner/lemmy/backend/contracts"
)

type app struct {
	ctx      context.Context
	bindings contracts.Bindings
}

func New() contracts.App {
	return &app{
		bindings: bindings.New(),
	}
}

func (me *app) OnStartup(ctx context.Context) {
	me.ctx = ctx
}

func (me *app) Bindings() contracts.Bindings {
	return me.bindings
}
