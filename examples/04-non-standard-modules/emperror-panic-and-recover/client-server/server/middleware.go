package main

import (
	"log"
	"net/http"

	"emperror.dev/emperror"
)

func NewPanicMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Recover паники и обработка её как ошибки.
		defer emperror.HandleRecover(emperror.ErrorHandlerFunc(func(err error) {
			log.Println("panic handler called:", err)
			internalServerError(w, "panic happened")
		}))
		h.ServeHTTP(w, req)
	})
}
