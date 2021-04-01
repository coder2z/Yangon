package tools

import (
	"os"
)

func WriteToFile(filename, content string) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	os.Chmod(filename, 0777)
	MustCheck(err)
	defer CloseFile(f)
	_, err = f.WriteString(content)
	MustCheck(err)
}

func MakeAllPath(path string) error {
	err := os.MkdirAll(path, 0777)
	MustCheck(err)
	os.Chmod(path, 0777)
	return nil
}

func CloseFile(f *os.File) {
	err := f.Close()
	MustCheck(err)
}

func RemoveAllList(paths ...string) (err error) {
	for _, path := range paths {
		_ = os.RemoveAll(path)
	}
	return err
}

func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
