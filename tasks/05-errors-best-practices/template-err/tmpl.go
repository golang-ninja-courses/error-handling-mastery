package tmpl

import (
	"errors"
	"html/template"
	"io"
)

var (
	errParseTemplate   = errors.New("parse template error")
	errExecuteTemplate = errors.New("execute template error")
)

func ParseAndExecuteTemplate(wr io.Writer, name, text string, data interface{}) error {
	t, err := template.New(name).Parse(text)
	if err != nil {
		return errParseTemplate
	}

	err = t.Execute(wr, data)
	if err != nil {
		return errExecuteTemplate
	}

	return nil
}
