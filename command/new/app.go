package newApp

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"yangon/tools"
)

func (options *RunOptions) Run() {
	tools.MustCheck(tools.GitClone("https://github.com/myxy99/Yangon-tpl.git", "tmp\\"+options.ProjectName))
	_ = filepath.Walk("tmp\\"+options.ProjectName+"\\new", func(path string, info os.FileInfo, err error) error {
		newPath := tools.ReplaceAllData(path, map[string]string{
			"{{AppName}}": options.AppName,
			"new\\":       "",
			"tmp\\":       "",
			".tmpl":       "",
		})
		if regexp.MustCompile(`\\.git`).MatchString(newPath) {
			return nil
		}
		if info.IsDir() {
			_ = os.MkdirAll(newPath, 777)
		} else {
			if tools.CheckFileIsExist(newPath) && options.Backup {
				tools.MustCheck(os.Rename(newPath, fmt.Sprintf("%s.bak", newPath)))
			}
			s, err := tools.ParseTmplFile(path, options)
			tools.MustCheck(err)
			tools.WriteToFile(newPath, s)
		}
		fmt.Println(newPath)
		return nil
	})

	_ = tools.RemoveAllList(options.ProjectName+"/new", "tmp")
}
