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

func doSomething() error {
	return &WithStacktraceError{
		message:    "something went wrong",
		stacktrace: debug.Stack(),
	}
}

func main() {
	if err := doSomething(); err != nil {
		if stacktraceErr, ok := err.(*WithStacktraceError); ok {
			fmt.Printf("%s\n%s", stacktraceErr.message, string(stacktraceErr.stacktrace))
		}
	}
}
