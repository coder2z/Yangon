package tools

import (
	"gopkg.in/src-d/go-git.v4"
	"os"
)

func GitClone(url, path string) (err error) {
	_, err = git.PlainClone(path, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	return
}
