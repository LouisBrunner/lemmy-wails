package main

import (
	"context"
	"flag"
	"fmt"

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
	// for the templating
	name        string
	repo        string
	version     string
	authorName  string
	authorEmail string
}

func (me *initCmd) Name() string {
	return "init"
}

func (me *initCmd) Synopsis() string {
	return "Initialize the current folder to use the boilerplate."
}

func (me *initCmd) Usage() string {
	return ""
}

func (me *initCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&me.force, "force", false, "init even if the repository is already setup")
	f.StringVar(&me.name, "name", "", "name for your project, e.g. Cool New App")
	f.StringVar(&me.repo, "repo", "", "where your code will be stored, e.g. github.com/me/newapp")
	f.StringVar(&me.version, "version", "main", "what version of the boilerplate to use")
	f.StringVar(&me.authorName, "authorName", "", "your name")
	f.StringVar(&me.authorEmail, "authorEmail", "", "your email address")
}

func (me *initCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	for _, flag := range []struct {
		name  string
		value string
	}{
		{name: "name", value: me.name},
		{name: "repo", value: me.repo},
		{name: "version", value: me.version},
		{name: "authorName", value: me.authorName},
		{name: "authorEmail", value: me.authorEmail},
	} {
		if flag.value == "" {
			return me.fromError(fmt.Errorf("missing `%s` flag", flag.name))
		}
	}

	data := internal.InitData{
		Name:               me.name,
		Repo:               me.repo,
		BoilerplateVersion: me.version,
		Author: internal.InitDataAuthor{
			Name:  me.authorName,
			Email: me.authorEmail,
		},
	}
	return me.fromError(internal.Init(me.logger, data, me.cfg.folder, me.force))
}
