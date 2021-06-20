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

/*
Тип NotFoundError имеет множество методов T:
- StatusCode

Тип NotFoundError имеет множество методов T + *T:
- Error
- StatusCode

Как следствие именно *NotFoundError реализует интерфейс HTTPError.
*/
func main() {
	var err HTTPError = &NotFoundError{}
	fmt.Println(err.StatusCode(), err.Error()) // 404 Not Found

	// Не скопилируется:
	// var _ HTTPError = NotFoundError{}
	// cannot use NotFoundError{} (type NotFoundError) as type HTTPError in assignment:
	//     NotFoundError does not implement HTTPError (Error method has pointer receiver)
}
