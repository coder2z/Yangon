package newApp

import "github.com/spf13/cobra"

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

func (s *RunOptions) Flags() () {
	s.c.Flags().StringVarP(&s.AppName, "AppName", "a", "demoApp", "app name")
	s.c.Flags().StringVarP(&s.ProjectName, "ProjectName", "p", "demoProject", "project name")
}
