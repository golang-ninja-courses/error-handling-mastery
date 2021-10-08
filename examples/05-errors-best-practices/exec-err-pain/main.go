package main

import (
	"errors"
	"text/template"
)

func main() {
	var err error

	var t template.ExecError
	var tPtr *template.ExecError

	switch {
	case errors.As(err, &t):
	case errors.As(err, &tPtr):
	}
}

type MyExecError struct{}

func (m *MyExecError) Error() string {
	return "cool error"
}

func (m *MyExecError) Is(target error) {
	// Что выбрать?
	switch target.(type) {
	case *template.ExecError:
		// ...
	case template.ExecError:
		// ...
	}
}

func (m *MyExecError) As(target interface{}) {
	// Что выбрать?
	switch target.(type) {
	case *template.ExecError:
		// ...
	case **template.ExecError:
		// ...
	}
}
