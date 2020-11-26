package main

import (
	"github.com/spf13/cobra"
	newApp "yangon/command/new"
)

func main() {
	rootCmd := &cobra.Command{Use: "Yangon"}
	rootCmd.AddCommand(newApp.App)
	_ = rootCmd.Execute()
}
