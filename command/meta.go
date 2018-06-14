package command

import (
	"flag"
	"github.com/mitchellh/cli"
)

type Meta struct {
	Ui cli.Ui
}

func (m *Meta) flagSet(n string) *flag.FlagSet {
	f := flag.NewFlagSet(n, flag.ContinueOnError)
	return f
}
