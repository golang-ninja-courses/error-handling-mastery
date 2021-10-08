package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

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

func internalServerError(w http.ResponseWriter, errMsg string) {
	w.WriteHeader(http.StatusInternalServerError)
	safeWrite(w, errMsg)
}

func ok(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusOK)
	safeWrite(w, msg)
}

func safeWrite(w http.ResponseWriter, msg string) {
	if _, err := w.Write([]byte(msg)); err != nil {
		log.Println(err)
	}
}
