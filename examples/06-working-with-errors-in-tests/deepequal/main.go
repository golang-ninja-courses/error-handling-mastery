package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
)

var MyEOF = errors.New(io.EOF.Error())

type Errors []error            // Комментарий для поддержки отсутствия новой строки между определением типа и методом.
func (e Errors) Error() string { return "errors" }

func main() {
	cases := []struct {
		name   string
		lhs    error
		rhs    error
		deepEq bool
	}{
		{
			name:   "io.EOF, io.EOF",
			lhs:    io.EOF,
			rhs:    io.EOF,
			deepEq: true,
		},
		{
			name:   "io.EOF, io.ErrUnexpectedEOF",
			lhs:    io.EOF,
			rhs:    io.ErrUnexpectedEOF,
			deepEq: false,
		},
		{
			name:   "io.EOF, nil",
			lhs:    io.EOF,
			rhs:    nil,
			deepEq: false,
		},
		{
			name:   `errors.New("some error"), errors.New("some error")`,
			lhs:    errors.New("some error"),
			rhs:    errors.New("some error"),
			deepEq: true,
		},
		{
			name:   "MyEOF, io.EOF",
			lhs:    MyEOF,
			rhs:    io.EOF,
			deepEq: true,
		},
		{
			name:   "&os.PathError{Err: io.EOF}, &os.PathError{Err: MyEOF}",
			lhs:    &os.PathError{Err: io.EOF},
			rhs:    &os.PathError{Err: MyEOF},
			deepEq: true,
		},
		{
			name:   "Errors{}, Errors(nil)",
			lhs:    Errors{},
			rhs:    Errors(nil),
			deepEq: false,
		},
		{
			name:   `fmt.Errorf("some error: %d", 10), fmt.Errorf("some error: %d", 10)`,
			lhs:    fmt.Errorf("some error: %d", 10),
			rhs:    fmt.Errorf("some error: %d", 10),
			deepEq: true,
		},
		{
			name:   `fmt.Errorf("some error: %w", io.EOF), fmt.Errorf("some error: %w", MyEOF)`,
			lhs:    fmt.Errorf("some error: %w", io.EOF),
			rhs:    fmt.Errorf("some error: %w", MyEOF),
			deepEq: true,
		},
	}

	for _, c := range cases {
		de := reflect.DeepEqual(c.lhs, c.rhs)
		if de != c.deepEq {
			log.Fatal(c.name)
		}
		fmt.Printf("%-100s: %v\n", fmt.Sprintf("reflect.DeepEqual(%s)", c.name), c.deepEq)
	}
}
