package main

import (
	"errors"
	"net"
	"text/template"
)

func main() {
	var err error

	var (
		t    template.ExecError
		tPtr *template.ExecError
	)
	switch {
	case errors.As(err, &t):
	case errors.As(err, &tPtr):
	}

	var (
		n    net.DNSError
		nPtr *net.DNSError
		_    = n
	)
	switch {
	// case errors.As(err, &n): // Запаникует!
	case errors.As(err, &nPtr):
	}
}

type MyExecError struct{}

func (m *MyExecError) Error() string {
	return "cool error"
}

func (m *MyExecError) Is(target error) bool {
	// Что выбрать?
	switch target.(type) {
	case *template.ExecError:
		// ...
	case template.ExecError:
		// ...
	}
	return false
}

func (m *MyExecError) As(target interface{}) bool {
	// Что выбрать?
	switch target.(type) {
	case *template.ExecError:
		// ...
	case **template.ExecError:
		// ...
	}
	return false
}
