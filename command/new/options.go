package newApp

import "github.com/spf13/cobra"

var App *cobra.Command

func init() {
	var options *RunOptions
	App = &cobra.Command{
		Use:   "new",
		Short: "Generate app scaffolding",
		Long:  `Quickly generate app scaffolding`,
		Run: func(cmd *cobra.Command, args []string) {
			options.Run()
		},
	}
	App.DisableSuggestions = true
	options = NewRunOptions(App)
	options.Flags()
}

type RunOptions struct {
	c                    *cobra.Command
	AppName, ProjectName string
	Backup               bool
}

func NewRunOptions(c *cobra.Command) *RunOptions {
	s := &RunOptions{
		c:           c,
		AppName:     "",
		ProjectName: "",
		Backup:      true,
	}
	return s
}

func (options *RunOptions) Flags() () {
	options.c.Flags().StringVarP(&options.AppName, "AppName", "a", "", "app name (required)")
	options.c.Flags().StringVarP(&options.ProjectName, "ProjectName", "p", "", "project name (required)")
	options.c.Flags().BoolVarP(&options.Backup, "Backup", "b", false, "backup file")
	_ = options.c.MarkFlagRequired("AppName")
	_ = options.c.MarkFlagRequired("ProjectName")
}
