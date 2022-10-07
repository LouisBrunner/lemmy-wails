package main

import (
	"github.com/google/subcommands"
	"github.com/sirupsen/logrus"
)

type commandCommons struct {
	logger *logrus.Logger
	cfg    *config
}

func newCommons(logger *logrus.Logger, cfg *config) commandCommons {
	return commandCommons{
		logger: logger,
		cfg:    cfg,
	}
}

func (me *commandCommons) fromError(err error) subcommands.ExitStatus {
	if err == nil {
		return subcommands.ExitSuccess
	}
	me.logger.WithError(err).Errorf("failed")
	return subcommands.ExitFailure
}
