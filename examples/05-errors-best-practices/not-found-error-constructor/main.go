package main

import (
	"fmt"
	"net/http"
)

type HTTPError interface {
	error
	StatusCode() int
}

type NotFoundError struct{}

func (n *NotFoundError) Error() string {
	return "Not Found"
}

func (n NotFoundError) StatusCode() int {
	return http.StatusNotFound
}

func NewNotFoundError() (*NotFoundError, error) {
	return new(NotFoundError), nil
}

func main() {
	var httpErr HTTPError
	httpErr, _ = NewNotFoundError()
	fmt.Println(httpErr.StatusCode(), httpErr.Error()) // 404 Not Found
}
