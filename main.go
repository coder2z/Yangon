package main

import (
	"yangon/command/model"
	newApp "yangon/command/new"
	"yangon/command/version"

	"github.com/coder2z/g-saber/xflag"
)

func main() {
	xflag.NewRootCommand(&xflag.Command{
		Use: "Yangon",
	})
	xflag.Register(
		newApp.Rpc,
		model.Model,
		version.Version,
	)
	_ = xflag.Parse()
}
