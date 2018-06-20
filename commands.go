package main

import (
	"github.com/dsalazar32/go-certbotify-aws/command"
	"github.com/mitchellh/cli"
)

var Commands map[string]cli.CommandFactory

func initCommands() map[string]cli.CommandFactory {
	meta := command.Meta{
		Ui: Ui,
	}

	Commands = map[string]cli.CommandFactory{
		"certbot": func() (cli.Command, error) {
			return &command.CertbotCommand{
				Meta:    meta,
				Command: []string{"certbot", "certonly"},
				CertbotDefaults: []string{
					"-n",
					"-agree-tos",
					"-dns-route53",
				},
			}, nil
		},
	}

	return Commands
}
