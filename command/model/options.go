package model

import (
	"github.com/coder2z/component/pkg/xflag"
	"github.com/spf13/cobra"
)

var Model xflag.CommandNode

func init() {
	options := NewRunOptions()
	Model = xflag.CommandNode{
		Name: "go",
		Command: &xflag.Command{
			Use:   "go",
			Short: "db,handle,server,route code production",
			Long:  `Quickly db,handle,server,route code code production`,
			RunE: func(cmd *cobra.Command, args []string) error {
				options.Run()
				return nil
			},
		},
		Flags: func(command *xflag.Command) {
			command.Flags().StringVarP(&options.AppName, "AppName", "a", "", "app name (required)")
			command.Flags().StringVarP(&options.ProjectName, "ProjectName", "p", "", "project name (required)")
			command.Flags().StringVarP(&options.Version, "ApiVersion", "v", "", "api version (required)")
			command.Flags().StringVarP(&options.dbKey, "dbKey", "k", "mysql", "dbKey")
			command.Flags().StringVarP(&options.dbLabel, "dbLabel", "l", "main", "dbLabel")
			command.Flags().StringP("config", "c", "config/config.toml", "配置文件")
			_ = command.MarkFlagRequired("AppName")
			_ = command.MarkFlagRequired("ProjectName")
			_ = command.MarkFlagRequired("ApiVersion")
		},
	}
}

type RunOptions struct {
	AppName, ProjectName, Version, dbKey, dbLabel string
}

func NewRunOptions() *RunOptions {
	s := &RunOptions{
		AppName: "",
	}
	return s
}
