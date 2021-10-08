package main

import (
	"errors"
	"fmt"
	"net/url"
	"syscall"
)

func main() {
	e1 := url.InvalidHostError("www.golang-courses.ru")
	e2 := url.InvalidHostError("www.golang-courses.ru")
	fmt.Println(errors.Is(e1, e2)) // true

	e3 := syscall.Errno(0xd)
	e4 := syscall.Errno(0xd)
	fmt.Println(errors.Is(e3, e4)) // true
}
