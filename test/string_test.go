package test

import (
	"fmt"
	"path/filepath"
	"sync"
	"testing"
	"yangon/tools"
)

func TestString(t *testing.T) {
	t.Log(tools.StrFirstToUpper("id"))
	t.Log(tools.UnStrFirstToUpper("Id"))
	t.Log(tools.UnStrFirstToUpper("v21"))

	t.Log(filepath.Join("{{path}}","{{table}}.go"))
}

func TestGO(t *testing.T) {
	var wg sync.WaitGroup
	for _,v:=range []string{"a","b","c"} {
		wg.Add(1)
		go func(a string) {
			fmt.Println(a)
			wg.Done()
		}(v)
	}

	wg.Wait()
}