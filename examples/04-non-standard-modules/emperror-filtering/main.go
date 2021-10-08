package main

import (
	"errors"
	"fmt"

	"emperror.dev/emperror"
)

var (
	err1       = errors.New("error 1")
	err2       = errors.New("error 2")
	errsToSkip = []error{err1, err2}

	errHandler = emperror.ErrorHandlerFunc(func(err error) {
		if err == nil {
			return
		}
		fmt.Printf("error handler called: %v\n", err)
	})

	errMatcher = emperror.ErrorMatcher(func(err error) bool {
		for i := range errsToSkip {
			if needDiscard := errors.Is(err, errsToSkip[i]); needDiscard {
				return true
			}
		}
		return false
	})
)

func main() {
	handler := emperror.WithFilter(errHandler, errMatcher)

	handler.Handle(err1)                        // Обработчик не сработает.
	handler.Handle(errors.New("unknown error")) // Обработчик сработает.
}
