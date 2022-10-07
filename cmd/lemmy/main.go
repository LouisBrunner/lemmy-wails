package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
	"github.com/sirupsen/logrus"
)

type config struct {
	folder string
	debug  bool
}

type commandMaker func(logger *logrus.Logger, cfg *config) subcommands.Command

var commands = []commandMaker{
	newInitCmd,
	newUpdateCmd,
	newSyncCmd,
}

func main() {
	cfg := &config{}
	logger := logrus.New()

	commander := subcommands.DefaultCommander
	commander.Register(subcommands.HelpCommand(), "")
	commander.Register(subcommands.FlagsCommand(), "")
	for _, commandMaker := range commands {
		commander.Register(commandMaker(logger, cfg), "")
	}
	flag.StringVar(&cfg.folder, "folder", ".", "which folder to run commands in")
	flag.BoolVar(&cfg.debug, "debug", false, "run in debug mode")

	flag.Parse()
	if cfg.debug {
		logger.SetLevel(logrus.DebugLevel)
	}
	ctx := context.Background()
	os.Exit(int(commander.Execute(ctx)))
}
