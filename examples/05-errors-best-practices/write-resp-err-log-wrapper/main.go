package main

import (
	"log"
	"net/http"
)

type LogWriter struct {
	http.ResponseWriter
}

/*
func (w LogWriter) Write(p []byte) (n int, err error) {
	n, err = w.ResponseWriter.Write(p)
	if err != nil {
		log.Println("cannot write response: " + err.Error())
	}
	return
}
*/

func (w LogWriter) Write(p []byte) {
	_, err := w.ResponseWriter.Write(p)
	if err != nil {
		log.Println("cannot write response: " + err.Error())
	}
	return
}

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
		LogWriter{w}.Write([]byte(`{"msg": "OK"}`))
	})

	http.HandleFunc("/body-not-allowed-err", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		w.Header().Set("content-type", "application/json")
		LogWriter{w}.Write([]byte(`{"msg": "OK"}`))
	})

	http.HandleFunc("/content-length-err", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.Header().Set("content-length", "1")
		LogWriter{w}.Write([]byte(`{"msg": "OK"}`))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		LogWriter{w}.Write([]byte(`{"msg": "OK"}`))
		// writeOK(LogWriter{w})
	})

	if err := http.ListenAndServe(":8080", http.DefaultServeMux); err != nil {
		log.Fatal(err)
	}
}

func writeOK(w http.ResponseWriter) {
	w.Write([]byte(`{"msg": "OK"}`))
}
