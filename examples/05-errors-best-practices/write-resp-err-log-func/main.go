package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hijack-err", func(w http.ResponseWriter, r *http.Request) {
		hj, ok := w.(http.Hijacker)
		if !ok {
			http.Error(w, "no hijacking support", http.StatusInternalServerError)
			return
		}

		conn, _, err := hj.Hijack()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		w.Header().Set("content-type", "application/json")
		logWriteErr(w.Write([]byte(`{"msg": "OK"}`))) // http: connection has been hijacked
	})

	http.HandleFunc("/body-not-allowed-err", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		w.Header().Set("content-type", "application/json")
		logWriteErr(w.Write([]byte(`{"msg": "OK"}`))) // http: request method or response status code does not allow body
	})

	http.HandleFunc("/content-length-err", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.Header().Set("content-length", "1")
		logWriteErr(w.Write([]byte(`{"msg": "OK"}`))) // wrote more than the declared Content-Length
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		logWriteErr(w.Write([]byte(`{"msg": "OK"}`)))
	})

	if err := http.ListenAndServe(":8080", http.DefaultServeMux); err != nil {
		log.Fatal(err)
	}
}

func logWriteErr(_ int, err error) {
	if err != nil {
		log.Println("cannot write response: " + err.Error())
	}
}
