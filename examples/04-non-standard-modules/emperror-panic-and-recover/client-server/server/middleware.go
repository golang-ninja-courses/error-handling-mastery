package main

import (
	"emperror.dev/emperror"
	"fmt"
	"net/http"
)

func NewPanicMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Recover паники и обработка её как ошибки
		defer emperror.HandleRecover(emperror.ErrorHandlerFunc(func(err error) {
			fmt.Printf("panic handler called: %v\n", err)
			internalServerError(w, "panic happened")
		}))
		h.ServeHTTP(w, req)
	})
}
