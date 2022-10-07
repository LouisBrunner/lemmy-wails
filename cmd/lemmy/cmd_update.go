package main

import (
	"context"
	"flag"

	"github.com/LouisBrunner/lemmy-wails/cmd/lemmy/internal"
	"github.com/google/subcommands"
	"github.com/sirupsen/logrus"
)

func newUpdateCmd(logger *logrus.Logger, cfg *config) subcommands.Command {
	return &updateCmd{
		commandCommons: newCommons(logger, cfg),
	}
}

type updateCmd struct {
	commandCommons
	version string
}

func (me *updateCmd) Name() string {
	return "update"
}

func (me *updateCmd) Synopsis() string {
	return "Update to the latest (or provided) version of the boilerplate and sync it."
}

func (me *updateCmd) Usage() string {
	return ""
}

func (me *updateCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&me.version, "version", "latest", "which version to upgrade to")
}

func (me *updateCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	err := internal.Update(me.logger, me.version, me.cfg.folder)
	if err != nil {
		return me.fromError(err)
	}
	return me.fromError(internal.Sync(me.logger, me.cfg.folder))
}
