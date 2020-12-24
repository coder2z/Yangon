package main

import (
	"github.com/myxy99/component/pkg/xflag"
	"yangon/command/model"
	newApp "yangon/command/new"
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
