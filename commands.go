package main

import (
	"github.com/mitchellh/cli"
	"github.com/dsalazar32/go-certbotify-aws/command"
)

var Commands map[string]cli.CommandFactory

func initCommands() map[string]cli.CommandFactory {
	meta := command.Meta{
		Ui: Ui,
	}

	Commands = map[string]cli.CommandFactory{
		"certbot": func() (cli.Command, error) {
			return &command.Certbot{
				Meta: meta,
			}, nil
		},
	}

	return Commands
}
