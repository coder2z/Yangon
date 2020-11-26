package newApp

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"yangon/tools"
)

var App *cobra.Command

func init() {
	var options *RunOptions
	App = &cobra.Command{
		Use:   "new",
		Short: "Generate app scaffolding",
		Long:  `Quickly generate app scaffolding`,
		Run: func(cmd *cobra.Command, args []string) {
			Run(options)
		},
	}
	App.DisableSuggestions = true
	options = NewRunOptions(App)
	options.Flags()
}

func Run(options *RunOptions) {
	tools.MustCheck(tools.GitClone("https://github.com/myxy99/Yangon-tpl.git", "tmp\\"+options.ProjectName))
	_ = filepath.Walk("tmp\\"+options.ProjectName, func(path string, info os.FileInfo, err error) error {
		newPath := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(path, "{{AppName}}", options.AppName), "new\\", ""), "tmp\\", "")
		if regexp.MustCompile(`\\.git`).MatchString(newPath) {
			return nil
		}
		if info.IsDir() {
			_ = os.MkdirAll(newPath, 777)
		} else {
			f, err := ioutil.ReadFile(path)
			tools.MustCheck(err)
			tools.WriteToFile(newPath, strings.ReplaceAll(strings.ReplaceAll(string(f), "{{AppName}}", options.AppName), "{{ProjectName}}", options.ProjectName))
		}
		fmt.Println(newPath)
		return nil
	})

	_ = tools.RemoveAllList(options.ProjectName+"/new", "tmp")

}
