package newApp

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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
	_, err := git.PlainClone(options.ProjectName, false, &git.CloneOptions{
		URL:      "https://github.com/myxy99/Yangon-tpl.git",
		Progress: os.Stdout,
	})
	MustCheck(err)
	_ = filepath.Walk(options.ProjectName, func(path string, info os.FileInfo, err error) error {
		newPath := strings.ReplaceAll(strings.ReplaceAll(path, "{{AppName}}", options.AppName), "new\\", "")
		if info.IsDir() {
			_ = os.MkdirAll(newPath, 777)
		} else {
			f, err := ioutil.ReadFile(path)
			MustCheck(err)
			WriteToFile(newPath, strings.ReplaceAll(strings.ReplaceAll(string(f), "{{AppName}}", options.AppName), "{{ProjectName}}", options.ProjectName))
		}
		fmt.Println(newPath)
		return nil
	})
	_ = os.RemoveAll(options.ProjectName + "/new")
}

func WriteToFile(filename, content string) {
	f, err := os.Create(filename)
	MustCheck(err)
	defer CloseFile(f)
	_, err = f.WriteString(content)
	MustCheck(err)
}
func MustCheck(err error) {
	if err != nil {
		panic(err)
	}
}
func CloseFile(f *os.File) {
	err := f.Close()
	MustCheck(err)
}
