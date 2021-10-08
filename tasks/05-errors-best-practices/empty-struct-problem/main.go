package main

import (
	"errors"
	"fmt"

	"github.com/www-golang-courses-ru/advanced-dealing-with-errors-in-go/tasks/05-errors-best-practices/empty-struct-problem/pkga"
	"github.com/www-golang-courses-ru/advanced-dealing-with-errors-in-go/tasks/05-errors-best-practices/empty-struct-problem/pkgb"
)

var (
	a error = pkga.EOF{}
	b error = pkgb.EOF{}
	c error = new(pkgb.EOF)
	d error = new(pkgb.EOF)
)

func main() {
	fmt.Println(a == b)          // false
	fmt.Println(c == d)          // true
	fmt.Println(errors.Is(a, c)) // false
	fmt.Println(errors.Is(c, d)) // true

	// 0x10303ec80 0x10303ec90
	// 0x1030733b8 0x1030733b8 <- Одинаковые!
	fmt.Printf("%p %p \n%p %p\n", &a, &b, c, d)
}
