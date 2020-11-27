package newApp

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"yangon/tools"
)

func (options *RunOptions) Run() {
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
