package main

import (
	"fmt"
	"runtime/debug"
)

type WithStacktraceError struct {
	message    string
	stacktrace []byte
}

func (w *WithStacktraceError) Error() string {
	return w.message
}

func (w *WithStacktraceError) StackTrace() string {
	return string(w.stacktrace)
}

func doSomething() error {
	return &WithStacktraceError{
		message:    "something went wrong",
		stacktrace: debug.Stack(),
	}
}

func main() {
	if err := doSomething(); err != nil {
		type stackTracer interface {
			StackTrace() string
		}
		if st, ok := err.(stackTracer); ok {
			fmt.Printf("%s\n%s", err, st.StackTrace())
		}
	}
}
