package test

import (
	"path/filepath"
	"testing"
	"yangon/tools"
)

func TestString(t *testing.T) {
	t.Log(tools.StrFirstToUpper("id"))
	t.Log(tools.UnStrFirstToUpper("Id"))
	t.Log(tools.UnStrFirstToUpper("v21"))

	t.Log(filepath.Join("{{path}}","{{table}}.go"))
}
