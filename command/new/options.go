package newApp

import (
	"github.com/coder2z/g-saber/xflag"
)

var Rpc xflag.CommandNode

func init() {
	options := NewRunOptions()
	Rpc = xflag.CommandNode{
		Name: "rpc",
		Command: &xflag.Command{
			Use:   "rpc",
			Short: "Generate rpc app scaffolding",
			Long:  `Quickly generate rpc app scaffolding`,
			Run: func(cmd *xflag.Command, args []string) {
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
