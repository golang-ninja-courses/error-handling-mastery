package main

import (
	"errors"
	"fmt"

	"emperror.dev/emperror"
)

var (
	err1     = errors.New("error 1")
	err2     = errors.New("error 2")
	errsSkip = []error{err1, err2}

	errHandler = emperror.ErrorHandlerFunc(func(err error) {
		if err == nil {
			return
		}

		// Логируем ошибку, например
		fmt.Printf("error handler called: %v\n", err)
	})

	errMatcher = emperror.ErrorMatcher(func(err error) bool {
		found := false
		for i := range errsSkip {
			if errors.Is(err, errsSkip[i]) {
				found = true
				break
			}
		}
		return found
	})
)

func main() {
	var handler = emperror.WithFilter(errHandler, errMatcher)

	handler.Handle(err1) // обработчик не сработает

	handler.Handle(errors.New("an error")) // обработчик сработает
}
