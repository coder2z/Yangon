package model

import (
	"github.com/spf13/cobra"
)

var Model *cobra.Command

func init() {
	var options *RunOptions
	Model = &cobra.Command{
		Use:   "go",
		Short: "db,handle,server,route code production",
		Long:  `Quickly db,handle,server,route code code production`,
		RunE: func(cmd *cobra.Command, args []string) error {
			options.Run()
			return nil
		},
	}
	Model.DisableSuggestions = true
	options = NewRunOptions(Model)
	options.Flags()
}

type RunOptions struct {
	c                             *cobra.Command
	AppName, ProjectName, Version string
}

func NewRunOptions(c *cobra.Command) *RunOptions {
	s := &RunOptions{
		c:       c,
		AppName: "",
	}
	return s
}

func (options *RunOptions) Flags() () {
	options.c.Flags().StringVarP(&options.AppName, "AppName", "a", "", "app name (required)")
	options.c.Flags().StringVarP(&options.ProjectName, "ProjectName", "p", "", "project name (required)")
	options.c.Flags().StringVarP(&options.Version, "ApiVersion", "v", "", "api version (required)")
	_ = options.c.MarkFlagRequired("AppName")
	_ = options.c.MarkFlagRequired("ProjectName")
	_ = options.c.MarkFlagRequired("ApiVersion")
}
