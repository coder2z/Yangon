package model

import "github.com/spf13/cobra"

var Model *cobra.Command

func init() {
	var options *RunOptions
	Model = &cobra.Command{
		Use:   "go",
		Short: "db,handle,server,route code production",
		Long:  `Quickly db,handle,server,route code code production`,
		Run: func(cmd *cobra.Command, args []string) {
			options.Run()
		},
	}
	Model.DisableSuggestions = true
	options = NewRunOptions(Model)
	options.Flags()
}

type RunOptions struct {
	c                    *cobra.Command
	AppName, ProjectName string
}

func NewRunOptions(c *cobra.Command) *RunOptions {
	s := &RunOptions{
		c:       c,
		AppName: "",
	}
	return s
}

func (options *RunOptions) Flags() () {
	options.c.Flags().StringVarP(&options.AppName, "AppName", "a", "demoApp", "app name")
	options.c.Flags().StringVarP(&options.ProjectName, "ProjectName", "p", "demoProject", "project name")
}
