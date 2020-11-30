package tools

import "os"

func WriteToFile(filename, content string) {
	f, err := os.Create(filename)
	MustCheck(err)
	defer CloseFile(f)
	_, err = f.WriteString(content)
	MustCheck(err)
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
