package tmpl

import (
	"html/template"
	"io"
)

var (
	errParseTemplate   error
	errExecuteTemplate error
)

func ParseAndExecuteTemplate(wr io.Writer, name, text string, data interface{}) {
	t, _ := template.New(name).Parse(text)
	t.Execute(wr, data)
}
