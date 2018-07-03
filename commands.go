package main

import (
	"github.com/dsalazar32/go-gen-ssl/command"
	"github.com/dsalazar32/go-gen-ssl/command/certbot"
	"github.com/mitchellh/cli"
)

var Commands map[string]cli.CommandFactory

func initCommands() map[string]cli.CommandFactory {
	meta := command.Meta{
		Ui: Ui,
	}

	cbot := &certbot.Certbot{
		CertbotFlags: certbot.CertbotFlags{
			{"-n", ""},
			{"--dns-route53", ""},
			{"--agree-tos", ""},
		},
	}

	Commands = map[string]cli.CommandFactory{
		"generate": func() (cli.Command, error) {
			return &command.SSLGenerator{
				Meta:    meta,
				Certbot: *cbot,
			}, nil
		},
	}

	return Commands
}
