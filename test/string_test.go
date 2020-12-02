package test

import (
	"testing"
	"yangon/tools"
)

func TestString(t *testing.T) {
	t.Log(tools.StrFirstToUpper("id"))
	t.Log(tools.UnStrFirstToUpper("Id"))
}
