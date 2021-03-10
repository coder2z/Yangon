package version

import (
	"github.com/coder2z/g-saber/xcast"
	"github.com/coder2z/g-saber/xconsole"
	"github.com/coder2z/g-saber/xflag"
	"github.com/spf13/cobra"
	"yangon/constant"
)

var Version xflag.CommandNode

func init() {
	Version = xflag.CommandNode{
		Name: "version",
		Command: &xflag.Command{
			Use:   "version",
			Short: "app version",
			Long:  `app version`,
			RunE: func(cmd *cobra.Command, args []string) error {
				xconsole.Greenf("version:", xcast.ToString(constant.Version))
				return nil
			},
		},
		Flags: func(command *xflag.Command) {},
	}
}
