package newAPi

import (
	"github.com/coder2z/g-saber/xflag"
)

var Api xflag.CommandNode

func init() {
	options := NewRunOptions()
	Api = xflag.CommandNode{
		Name: "api",
		Command: &xflag.Command{
			Use:   "api",
			Short: "Generate api app scaffolding",
			Long:  `Quickly generate api app scaffolding`,
			Run: func(cmd *xflag.Command, args []string) {
				options.Run()
			},
		},
		Flags: func(command *xflag.Command) {
			command.Flags().StringVarP(&options.ProjectName, "ProjectName", "p", "", "project name (required)")
			command.Flags().BoolVarP(&options.Backup, "Backup", "b", false, "backup file")
			_ = command.MarkFlagRequired("ProjectName")
		},
	}
}

type RunOptions struct {
	ProjectName string
	Backup      bool
}

func NewRunOptions() *RunOptions {
	s := &RunOptions{
		ProjectName: "",
		Backup:      true,
	}
	return s
}
