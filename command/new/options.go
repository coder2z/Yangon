package newApp

import "github.com/spf13/cobra"

var App *cobra.Command

func init() {
	var options *RunOptions
	App = &cobra.Command{
		Use:   "new",
		Short: "Generate app scaffolding",
		Long:  `Quickly generate app scaffolding`,
		Run: func(cmd *cobra.Command, args []string) {
			options.Run()
		},
	}
	App.DisableSuggestions = true
	options = NewRunOptions(App)
	options.Flags()
}

type RunOptions struct {
	c                    *cobra.Command
	AppName, ProjectName string
}

func NewRunOptions(c *cobra.Command) *RunOptions {
	s := &RunOptions{
		c:           c,
		AppName:     "",
		ProjectName: "",
	}
	return s
}

func (options *RunOptions) Flags() () {
	options.c.Flags().StringVarP(&options.AppName, "AppName", "a", "demoApp", "app name")
	options.c.Flags().StringVarP(&options.ProjectName, "ProjectName", "p", "demoProject", "project name")
}
