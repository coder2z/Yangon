package newApp

import (
	"fmt"
	"github.com/coder2z/g-saber/xconsole"
	"os"
	"path/filepath"
	"regexp"
	"yangon/constant"
	"yangon/tools"
)

func (options *RunOptions) Run() {
	tools.MustCheck(tools.GitClone(constant.GitUrl, "tmp\\"+options.ProjectName))
	_ = filepath.Walk("tmp\\"+options.ProjectName+"\\new", func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return nil
		}
		newPath := tools.ReplaceAllData(path, map[string]string{
			"{{AppName}}": options.AppName,
			"new\\":       "",
			"tmp\\":       "",
			".tmpl":       "",
		})
		if regexp.MustCompile(`\\.git`).MatchString(newPath) && !info.IsDir() {
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
		xconsole.Green(newPath)
		return nil
	})

	_ = tools.RemoveAllList(options.ProjectName+"/new", "tmp")
}
