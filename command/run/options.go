package run

import (
	"github.com/spf13/cobra"
)

var Run *cobra.Command

func init() {
	var options *RunOptions
	Run = &cobra.Command{
		Use:   "run",
		Short: "run all app",
		Long:  `run all app`,
		RunE: func(cmd *cobra.Command, args []string) error {
			options.Run()
			return nil
		},
	}
	Run.DisableSuggestions = true
	options = NewRunOptions(Run)
	options.Flags()
}

type RunOptions struct {
	c *cobra.Command
}

func NewRunOptions(c *cobra.Command) *RunOptions {
	s := &RunOptions{c: c}
	return s
}

func (options *RunOptions) Flags() () {
}
