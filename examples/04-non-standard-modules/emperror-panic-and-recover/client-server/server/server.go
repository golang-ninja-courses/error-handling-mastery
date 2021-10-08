package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	port            = 8888
	shutdownTimeout = time.Second
)

func main() {
	mux := http.DefaultServeMux
	mux.Handle("/", http.HandlerFunc(Handle))

	server := &http.Server{
		Addr:    fmt.Sprintf("127.0.0.1:%d", port),
		Handler: NewPanicMiddleware(mux),
	}

	errs := make(chan error)
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go listen(errs, server)

	select {
	case err := <-errs:
		log.Println("error happened:", err)
		_ = shutdown(server)

	case s := <-signals:
		log.Println("signal received:", s)
		signal.Stop(signals)
		_ = shutdown(server)
	}
}

func listen(errs chan<- error, server *http.Server) {
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		errs <- err
	}
}

func shutdown(server *http.Server) error {
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	server.SetKeepAlivesEnabled(false)
	return server.Shutdown(ctx)
}
