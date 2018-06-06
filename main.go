package main

import (
	"github.com/mitchellh/cli"
	"os"
	"log"
)

const (
	ErrPrefix = "e: "
	OutPrefix = "o: "
	InfPrefix = "i: "
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
