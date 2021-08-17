package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	port = 8888

	shutdownTimeout = 1 * time.Second
)

var (
	shutdownSignals = []os.Signal{syscall.SIGINT, syscall.SIGTERM}

	middlewares = []func(http.Handler) http.Handler{
		NewPanicMiddleware,
	}
)

func listen(ch chan error, server *http.Server) {
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		ch <- err
	}
}

func shutdown(server *http.Server) error {
	var cancel func()
	ctx := context.Background()

	ctx, cancel = context.WithTimeout(ctx, shutdownTimeout)
	defer cancel()

	server.SetKeepAlivesEnabled(false)
	return server.Shutdown(ctx)
}

func internalServerError(w http.ResponseWriter, errMsg string) {
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(errMsg))
}

func ok(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(msg))
}

func Handle(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		internalServerError(w, "error while reading request body")
		return
	}

	switch string(body) {
	case "panic":
		panic("client wants me to panic")
	case "error":
		ok(w, "internal server error")
	default:
		ok(w, "ok")
	}
}

func main() {
	mux := http.DefaultServeMux
	mux.Handle("/", http.HandlerFunc(Handle))

	var handler http.Handler = mux
	for _, mw := range middlewares {
		handler = mw(handler)
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handler,
	}

	errsChan := make(chan error)
	signalsChan := make(chan os.Signal, 2)
	signal.Notify(signalsChan, shutdownSignals...)

	go listen(errsChan, server)

	select {
	case err := <-errsChan:
		fmt.Printf("error happened: %v\n", err)
		_ = shutdown(server)
	case s := <-signalsChan:
		fmt.Printf("signal received: %v\n", s)
		signal.Stop(signalsChan)
		_ = shutdown(server)
	}
}
