package main

import (
	newAPi "yangon/command/api"
	"yangon/command/model"
	newApp "yangon/command/new"
	"yangon/command/version"

	"github.com/coder2z/g-saber/xflag"
)

func main() {
	xflag.NewRootCommand(&xflag.CommandNode{
		Name: "Yangon",
		Command: &xflag.Command{
			Use:                "Yangon",
			DisableSuggestions: false,
		},
	})
	xflag.Register(
		newApp.Rpc,
		model.Model,
		version.Version,
		newAPi.Api,
	)
	_ = xflag.Parse()
}
