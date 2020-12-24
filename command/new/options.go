package newApp

import (
	"github.com/myxy99/component/pkg/xflag"
	"github.com/spf13/cobra"
)

var App xflag.CommandNode

func init() {
	options := NewRunOptions()
	App = xflag.CommandNode{
		Name: "new",
		Command: &xflag.Command{
			Use:   "new",
			Short: "Generate app scaffolding",
			Long:  `Quickly generate app scaffolding`,
			Run: func(cmd *cobra.Command, args []string) {
				options.Run()
			},
		},
		Flags: func(command *xflag.Command) {
			command.Flags().StringVarP(&options.AppName, "AppName", "a", "", "app name (required)")
			command.Flags().StringVarP(&options.ProjectName, "ProjectName", "p", "", "project name (required)")
			command.Flags().BoolVarP(&options.Backup, "Backup", "b", false, "backup file")
			_ = command.MarkFlagRequired("AppName")
			_ = command.MarkFlagRequired("ProjectName")
		},
	}
}

type RunOptions struct {
	AppName, ProjectName string
	Backup               bool
}

func NewRunOptions() *RunOptions {
	s := &RunOptions{
		AppName:     "",
		ProjectName: "",
		Backup:      true,
	}
	return s
}
