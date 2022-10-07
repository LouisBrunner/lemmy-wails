package main

import (
	"context"
	"flag"

	"github.com/LouisBrunner/lemmy-wails/cmd/lemmy/internal"
	"github.com/google/subcommands"
	"github.com/sirupsen/logrus"
)

func newSyncCmd(logger *logrus.Logger, cfg *config) subcommands.Command {
	return &syncCmd{
		commandCommons: newCommons(logger, cfg),
	}
}

type syncCmd struct {
	commandCommons
}

func (me *syncCmd) Name() string {
	return "sync"
}

func (me *syncCmd) Synopsis() string {
	return "Sync the current directory with the Lemmy-Wails boilerplate."
}

func (me *syncCmd) Usage() string {
	return ""
}

func (me *syncCmd) SetFlags(f *flag.FlagSet) {
}

func (me *syncCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	return me.fromError(internal.Sync(me.logger, me.cfg.folder))
}
