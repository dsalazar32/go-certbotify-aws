package main

import (
	"github.com/dsalazar32/go-gen-ssl/command"
	"github.com/mitchellh/cli"
	"github.com/dsalazar32/go-gen-ssl/command/certbot"
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

	// TODO: default just executes certbot
	// TODO: s3 sends output to s3 bucket
	Commands = map[string]cli.CommandFactory{
		"local": func() (cli.Command, error) {
			return &command.CertbotCommand{
				Meta:    meta,
				Certbot: *cbot,
			}, nil
		},
	}

	return Commands
}
