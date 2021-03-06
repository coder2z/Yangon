package main

import (
	"yangon/command/model"
	newApp "yangon/command/new"

	"github.com/coder2z/component/pkg/xflag"
)

func main() {
	xflag.NewRootCommand(&xflag.Command{
		Use: "Yangon",
	})
	xflag.Register(
		newApp.App,
		model.Model,
	)
	_ = xflag.Parse()
}
