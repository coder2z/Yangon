package tools

import (
	"bytes"
	"text/template"
)

func ParseTmplFile(tmpl string, project interface{}) (string, error) {
	tmp, err := template.ParseFiles(tmpl)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err = tmp.Execute(&buf, project); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func ParseTmpl(tmpl string, project interface{}) (string, error) {
	tmp, err := template.New("").Parse(tmpl)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err = tmp.Execute(&buf, project); err != nil {
		return "", err
	}
	return buf.String(), nil
}
