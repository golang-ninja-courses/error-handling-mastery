package main

import (
	"fmt"

	"emperror.dev/emperror"
	"emperror.dev/errors"
)

var (
	errHandler = emperror.ErrorHandlerFunc(func(err error) {
		if err == nil {
			return
		}

		// Логируем ошибку, например
		fmt.Printf("error handler called: %v\n", err)
	})
)

func doSomething() error {
	return errors.New("error")
}

func main() {
	// Recover паники и обработка её как ошибки
	defer emperror.HandleRecover(errHandler)

	// nil-ошибка панику не вызовет
	emperror.Panic(nil)

	// Будет паника, если doSomething() вернет не nil.
	// emperror.Panic() полезно использовать при инициализации структур на старте приложения, когда при возникновении
	// ошибки надо падать.
	emperror.Panic(doSomething())
}
