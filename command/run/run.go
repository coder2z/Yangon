package run

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func (options *RunOptions) Run() {
	_ = filepath.Walk("cmd", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if regexp.MustCompile(`main\.go`).MatchString(path) {
				fmt.Println(path)
				return nil
			}
		}
		return nil
	})
}
