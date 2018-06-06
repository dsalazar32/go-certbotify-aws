package command

import (
	"github.com/mitchellh/cli"
	"flag"
)

type Meta struct {
	Ui cli.Ui
}

func (m *Meta) flagSet(n string) *flag.FlagSet {
	f := flag.NewFlagSet(n, flag.ContinueOnError)
	return f
}
