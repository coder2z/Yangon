package main

import (
	"github.com/spf13/cobra"
	"yangon/command/model"
	newApp "yangon/command/new"
	"yangon/command/run"
)

func main() {
	rootCmd := &cobra.Command{Use: "Yangon"}
	rootCmd.AddCommand(
		newApp.App,
		model.Model,
		run.Run,
	)
	_ = rootCmd.Execute()
}
