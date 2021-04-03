package run

import (
	"github.com/coder2z/g-saber/xflag"
)

var Run *xflag.Command

func init() {
	var options *RunOptions
	Run = &xflag.Command{
		Use:   "run",
		Short: "run all app",
		Long:  `run all app`,
		RunE: func(cmd *xflag.Command, args []string) error {
			options.Run()
			return nil
		},
	}
	Run.DisableSuggestions = true
	options = NewRunOptions(Run)
	options.Flags()
}

type RunOptions struct {
	c *xflag.Command
}

func NewRunOptions(c *xflag.Command) *RunOptions {
	s := &RunOptions{c: c}
	return s
}

func (options *RunOptions) Flags() () {
}
