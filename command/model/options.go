package model

import (
	"github.com/coder2z/g-saber/xflag"
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
		},
	}

	Api := xflag.CommandNode{
		Name: "apim",
		Command: &xflag.Command{
			Use:   "api",
			Short: "api db,handle,server,route code production",
			Long:  `Quickly api db,handle,server,route code code production`,
			Run: func(cmd *xflag.Command, args []string) {
				options.ApiRun()
			},
		},
		Flags: func(command *xflag.Command) {
			command.Flags().StringVarP(&options.ProjectName, "ProjectName", "p", "", "project name (required)")
			command.Flags().StringVarP(&options.Version, "ApiVersion", "v", "", "api version (required)")
			command.Flags().StringVarP(&options.dbKey, "dbKey", "k", "mysql", "dbKey")
			command.Flags().StringVarP(&options.dbLabel, "dbLabel", "l", "main", "dbLabel")
			command.Flags().StringP("xcfg", "c", "config/config.toml", "配置文件")
			_ = command.MarkFlagRequired("ProjectName")
			_ = command.MarkFlagRequired("ApiVersion")
		},
	}

	Rpc := xflag.CommandNode{
		Name: "rpc",
		Command: &xflag.Command{
			Use:   "rpc",
			Short: "rpc db,handle,server,route code production",
			Long:  `Quickly rpc db,handle,server,route code code production`,
			Run: func(cmd *xflag.Command, args []string) {
				options.RpcRun()
			},
		},
		Flags: func(command *xflag.Command) {
			command.Flags().StringVarP(&options.AppName, "AppName", "a", "", "app name (required)")
			command.Flags().StringVarP(&options.ProjectName, "ProjectName", "p", "", "project name (required)")
			command.Flags().StringVarP(&options.Version, "ApiVersion", "v", "", "api version (required)")
			command.Flags().StringVarP(&options.dbKey, "dbKey", "k", "mysql", "dbKey")
			command.Flags().StringVarP(&options.dbLabel, "dbLabel", "l", "main", "dbLabel")
			command.Flags().StringP("xcfg", "c", "config/config.toml", "配置文件")
			_ = command.MarkFlagRequired("AppName")
			_ = command.MarkFlagRequired("ProjectName")
			_ = command.MarkFlagRequired("ApiVersion")
		},
	}
	xflag.RegisterSpecify(&Model,Rpc,Api)
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
