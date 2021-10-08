package main

import (
	"fmt"

	"emperror.dev/emperror"
	"emperror.dev/errors"
)

var panicHandler = emperror.ErrorHandlerFunc(func(err error) {
	if err == nil {
		return
	}
	// Логируем ошибку, например.
	fmt.Println("panic handler called:", err)
})

func main() {
	// Recover паники и обработка её как ошибки.
	defer emperror.HandleRecover(panicHandler)

	// nil-ошибка панику не вызовет.
	emperror.Panic(nil)

	// Будет паника, если doSomething вернёт ошибку.
	// emperror.Panic полезно использовать при инициализации компонентов приложения на его старте,
	// когда при возникновении ошибки надо падать.
	emperror.Panic(doSomething())
}

func doSomething() error {
	return errors.New("error from doSomething")
}
