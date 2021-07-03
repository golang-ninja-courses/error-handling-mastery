package main

import (
	"fmt"
	"io"

	"github.com/www-golang-courses-ru/advanced-dealing-with-errors-in-go/examples/05-errors-best-practices/constant-errors-problem/dirtyhacker"
)

func main() {
	err1 := io.EOF
	err2 := io.EOF
	fmt.Println(io.EOF)       // EOF
	fmt.Println(err1 == err2) // true

	// Меняем глобальный io.EOF,
	// в принципе это можно сделать в init пакета dirtyhacker.
	dirtyhacker.MutateEOF()

	fmt.Println(io.EOF)         // nil
	fmt.Println(err1 == io.EOF) // false
}
