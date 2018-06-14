package main

import (
	"github.com/mitchellh/cli"
	"log"
	"os"
)

const (
	OutPrefix = "[Go-Certbot] "
	InfPrefix = "[Go-Certbot Info] "
	ErrPrefix = "[Go-Certbot Error] "
)

var Ui cli.Ui

func init() {
	Ui = &cli.PrefixedUi{
		OutputPrefix: OutPrefix,
		InfoPrefix:   InfPrefix,
		ErrorPrefix:  ErrPrefix,
		Ui:           &cli.BasicUi{Writer: os.Stdout},
	}
}

func main() {
	c := cli.NewCLI("go-certbotify-aws", "0.0.1")
	c.Args = os.Args[1:]
	c.Commands = initCommands()

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
