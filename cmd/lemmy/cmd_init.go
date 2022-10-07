package main

import (
	"context"
	"flag"

	"github.com/LouisBrunner/lemmy-wails/cmd/lemmy/internal"
	"github.com/google/subcommands"
	"github.com/sirupsen/logrus"
)

func newInitCmd(logger *logrus.Logger, cfg *config) subcommands.Command {
	return &initCmd{
		commandCommons: newCommons(logger, cfg),
	}
}

type initCmd struct {
	commandCommons
	force bool
}

func (me *initCmd) Name() string {
	return "init"
}

func (me *initCmd) Synopsis() string {
	return "Initialize the current folder to use Lemmy-Wails."
}

func (me *initCmd) Usage() string {
	return ""
}

func (me *initCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&me.force, "force", false, "init even if the repository is already setup")
}

func (me *initCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	return me.fromError(internal.Init(me.logger, me.cfg.folder, me.force))
}
