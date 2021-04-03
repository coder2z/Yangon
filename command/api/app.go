package newAPi

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
	dir, _ := os.Getwd()
	defer func() {
		_ = tools.RemoveAllList(filepath.Join(dir, options.ProjectName, "newApi"), filepath.Join(dir, "temporary"))
	}()
	tools.MustCheck(tools.GitClone(constant.GitUrl, filepath.Join(dir, "temporary", options.ProjectName)))
	_ = os.Mkdir(filepath.Join(dir, "temporary", options.ProjectName, "newApi"), 0777)
	_ = filepath.Walk(filepath.Join(dir, "temporary", options.ProjectName, "newApi"), func(path string, info os.FileInfo, err error) error {
		newPath := tools.ReplaceAllData(path, map[string]string{
			"newApi":    "",
			".tmpl":     "",
			"temporary": "",
		})

		if regexp.MustCompile(`.git`).MatchString(newPath) && !info.IsDir() {
			return nil
		}
		if info.IsDir() {
			_ = os.MkdirAll(newPath, 0777)
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
}
