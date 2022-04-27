package main

import (
	"log"
	"net/http"

	"github.com/www-golang-courses-ru/advanced-dealing-with-errors-in-go/tasks/05-errors-best-practices/json-writer/server"
)

func main() {
	s := server.New(stdLogger{}, simpleMsgProvider{})
	http.HandleFunc("/", s.HandleIndex)

	if err := http.ListenAndServe(":8080", http.DefaultServeMux); err != nil {
		log.Fatal(err)
	}
}

type stdLogger struct{}

func (l stdLogger) Error(msg string) {
	log.Println(msg)
}

type simpleMsgProvider struct{}

func (s simpleMsgProvider) Data() any {
	return struct {
		Msg string `json:"msg"`
	}{Msg: "OK"}
}
